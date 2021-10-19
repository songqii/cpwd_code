package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var banner = `
                       _ 
                      | |
  ___ _ ____      ____| |
 / __| '_ \ \ /\ / / _  |
| (__| |_) \ V  V / (_| |
 \___| .__/ \_/\_/ \__,_|
     | |
     |_|

`

var long = fmt.Sprintf("%s \n %s ", purple(banner), `
cpwd can fast create password
default create 16 length password

`)

var Random *rand.Rand
var lowerCharacter = `abcdefghijklmopqrstuvwxyz`
var upperCharacter = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
var charMap = `~!@#$%^&*()_+|}{:"?><,.;'[]=-'"\/`
var number = `0123456789`

var length int
var nochar bool
var nonumber bool

func init() {
	Random = rand.New(rand.NewSource(time.Now().UnixNano()))
	rootCmd.PersistentFlags().IntVar(&length, "length", 16, "you need password char length, default 16")
	rootCmd.PersistentFlags().BoolVar(&nochar, "nochar", false, "if you no need spechars, set true, default false")
	rootCmd.PersistentFlags().BoolVar(&nonumber, "nonumber", false, "if you no need number, set true, default false")
}

var rootCmd = &cobra.Command{
	Use:     "cpwd",
	Version: "v0.0.1",
	Short:   "cpwd is create password tool",
	Long:    long,
	Run: func(cmd *cobra.Command, args []string) {
		pwd := createPwd(length)
		if err := clipboard.WriteAll(pwd); err != nil {
			fmt.Println(red(fmt.Sprintf("[ERROR]%v", err)))
		}
		fmt.Println(green(`Your create password (It had copy to your clipboard yet. you can cmd+v / ctrl+v to copy your password :)`))
		fmt.Println(pwd)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(red(fmt.Sprintf("[ERROR]%v", err)))
	}
}

func createPwd(num int) string {
	var all []string
	all = append(all, lowerCharacter)
	all = append(all, upperCharacter)
	if !nochar {
		all = append(all, charMap)
	}
	if !nonumber {
		all = append(all, number)
	}
	var res []rune
	for i := 0; i <= num-1; i++ {
		s1 := all[Random.Intn(len(all))]
		r2 := Random.Intn(len(s1))
		res = append(res, rune(s1[r2]))
	}
	return string(res)
}
