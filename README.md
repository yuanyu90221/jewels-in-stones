# jewels-and-stones

## 題目解讀：

### 題目來源:
[jewels-and-stones](https://leetcode.com/problems/jewels-and-stones/)
### 原文:
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

### 解讀：
給定一個 字串 J 其中每個字元代表一個種類的寶石,其中每個種類都只有一種
給定一個 字串 S 其中每個字元代表一個種類的石頭

想要找到一個演算法可以

算出 S中 有多少個石頭是寶石

## 初步解法:
### 初步觀察:

首先S中的每個字元都是石頭

J則代表所有種類的寶石

所以只找把每個字元去檢查是否有在J的寶石種類字串即可

### 初步設計:

Given a string J, a string S

Step 1: let size = length of S

step 2: let index = 0, count = 0

step 3: check if index == size go to step 6

step 4: check if J includes S[index] then count++

step 5: index=index+1 go to step 3

step 6: return count
## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package jewels_stone

import "strings"

func numsJewelsInStone(J string, S string) int {
	count := 0
	for _, r := range S {
		temp_s := string(r)
		if strings.Contains(J, temp_s) {
			count++
		}
	}
	return count
}

```
## 測資的撰寫
```golang
package jewels_stone

import "testing"

func Test_numsJewelsInStone(t *testing.T) {
	type args struct {
		J string
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				J: "aA",
				S: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "Example2",
			args: args{
				J: "z",
				S: "ZZ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numsJewelsInStone(tt.args.J, tt.args.S); got != tt.want {
				t.Errorf("numsJewelsInStone() = %v, want %v", got, tt.want)
			}
		})
	}
}
```
## my self record
[leetcode golang 30day 7th day](https://hackmd.io/L4lBmEWKT4KT-M5QIwIiLQ?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)