package BLC

import "fmt"

func (cli *CLI) GetAddressLists(nodeID string) {
	fmt.Println("打印所有的钱包地址。。")
	//获取钱包的集合，遍历，依次输出
	wallets := NewWallets(nodeID)
	for address, _ := range wallets.WalletMap {
		fmt.Println("address：", address)
	}
}
