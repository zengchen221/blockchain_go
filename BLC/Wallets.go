package BLC

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"log"
	"crypto/elliptic"
	"io/ioutil"
	"os"
)

const walletsFile = "Wallets_%s.dat"//存储钱包数据的本地文件名


//定义一个钱包的集合，存储多个钱包对象
type Wallets struct {
	WalletMap map[string]*Wallet
}

//提供一个函数，用于创建一个钱包的集合
/*
思路：修改该方法：
	读取本地的钱包文件，如果文件存在，直接获取
	如果文件不存在，创建钱包对象
 */
func NewWallets(nodeID string) *Wallets{
	//wallerts:=&Wallets{}
	//wallerts.WalletMap = make(map[string]*Wallet)
	//return wallerts

	/*
	格式化钱包文件的名字

	 */

	 walletsFile := fmt.Sprintf(walletsFile,nodeID)

	//step1：钱包文件不存在
	if _, err := os.Stat(walletsFile); os.IsNotExist(err) {
		fmt.Println("钱包文件不存在。。")
		wallets:=&Wallets{}
		wallets.WalletMap = make(map[string]*Wallet)
		return wallets
	}
	//step2:钱包文件存在：读取本地的钱包数据-->钱包对象
	//读取本地的钱包文件中的数据---->反序列得到钱包集合对象
	wsBytes,err:=ioutil.ReadFile(walletsFile)
	if err != nil{
		log.Panic(err)
	}
	//将数据，变成钱包集合对象
	gob.Register(elliptic.P256()) //Curve

	var wallets Wallets
	reader:=bytes.NewReader(wsBytes)
	decoder:=gob.NewDecoder(reader)
	err = decoder.Decode(&wallets)
	if err != nil{
		log.Panic(err)
	}
	return &wallets


}

func (ws *Wallets) CreateNewWallet(nodeID string){
	wallet:=NewWallet()

	address := wallet.GetAddress()
	fmt.Printf("创建的钱包地址：%s\n",address)

	ws.WalletMap[string(address)] = wallet

	//将钱包集合，存入到本地文件中
	ws.saveFile(nodeID)
}


//将钱包对象，存入到本地文件中
func (ws *Wallets) saveFile(nodeID string){

	//格式化文件名
	walletsFile:=fmt.Sprintf(walletsFile,nodeID)


	//1.将ws对象的数据--->byte[]
	var buf bytes.Buffer
	//序列化的过程中：被序列化的对象 中包含了接口，那么将口需要注册
	gob.Register(elliptic.P256()) //Curve
	encoder:=gob.NewEncoder(&buf)
	err :=encoder.Encode(ws)
	if err != nil{
		log.Panic(err)
	}

	wsBytes:=buf.Bytes()

	//2.将数据，存储到文件中，
	//注意：该方法的实现：ioutil.WriteFile，覆盖写数据
	err = ioutil.WriteFile(walletsFile,wsBytes,0644)
	if err != nil{
		log.Panic(err)
	}
}
