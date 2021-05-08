/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/echovl/cardano-wallet/db"
	"github.com/echovl/cardano-wallet/wallet"
	"github.com/spf13/cobra"
)

// listWalletCmd represents the listWallet command
var listWalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Print a list of known wallets",
	Run: func(cmd *cobra.Command, args []string) {
		bdb := db.NewBadgerDB()
		defer bdb.Close()

		wallets := wallet.GetWallets(bdb)

		fmt.Printf("%-18v %-9v %-9v\n", "ID", "NAME", "ADDRESS")
		for _, v := range wallets {
			fmt.Printf("%-18v %-9v %-9v\n", v.ID, v.Name, len(v.Keys()))
		}
	},
}

func init() {
	listCmd.AddCommand(listWalletCmd)
}