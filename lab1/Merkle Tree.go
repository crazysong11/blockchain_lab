package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Block struct {
	index  int
	hash   string
	lchild *Block
	rchild *Block
}

func calSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

/*
// 由于个人写法，可以直接遍历切片
func compareMerkleTree(list1, list2 []Block) int {
	for i := 0; i < 16; i++ {
		if list1[i] != list2[i] {
			return i
		}
	}
	return -1
}
*/

// 遍历Merkle树
func compareMerkleTree(root1, root2 *Block) int {
	var p1, p2 *Block
	p1 = root1
	p2 = root2
	for {
		if p1.lchild == nil {
			return p1.index
		}
		if p1.lchild.hash != p2.lchild.hash {
			p1 = p1.lchild
			p2 = p2.lchild
		} else {
			p1 = p1.rchild
			p2 = p2.rchild
		}
	}
}

// 前序遍历输出树节点
func preorder(p *Block) {
	if p == nil {
		return
	}

	//输出信息
	fmt.Print("index: " + strconv.Itoa(p.index) + "\n")
	fmt.Print("hash: " + p.hash + "\n")
	if p.lchild != nil {
		fmt.Print("lchild's index: " + strconv.Itoa(p.lchild.index) + "\n")
		fmt.Print("rchild's index: " + strconv.Itoa(p.rchild.index) + "\n")
	} else {
		fmt.Print("lchild's index: no left child\n")
		fmt.Print("rchild's index: no left child\n")
	}

	preorder(p.lchild)
	preorder(p.rchild)
}

// 给定16个叶子节点生成Merkle树
func genMerkleTree(blockList []Block) ([]Block, *Block) {
	// 遍历切片生成父节点，随着切片元素的遍历，元素个数同时在增多
	for i := 0; i < 30; i = i + 2 {
		var newBlock Block
		newBlock.index = len(blockList)
		newBlock.hash = calSha256(calSha256(blockList[i].hash) + calSha256(blockList[i+1].hash))
		newBlock.lchild = &blockList[i]
		newBlock.rchild = &blockList[i+1]
		blockList = append(blockList, newBlock)
	}
	// 生成完毕，现在一共31个区块，index from 0 to 30
	fmt.Print("Merkle树生成完毕!\n")
	// 根节点是切片最后一个元素
	root := &blockList[len(blockList)-1]
	/* 输出生成Merkle树的信息
	fmt.Print("刚刚生成的Merkle树节点分别为（前序遍历）：\n")
	preorder(root)
	*/
	return blockList, root
}

func main() {
	var blockList []Block
	// 装16个叶节点
	for i := 0; i < 16; i++ {
		var s string
		fmt.Scanf("%s\n", &s)
		// 初始化区块
		block := Block{
			index:  i,
			hash:   calSha256(s),
			lchild: nil,
			rchild: nil,
		}
		// 将区块装入blocklist
		blockList = append(blockList, block)
	}
	_, root1 := genMerkleTree(blockList)

	//再生成一个树
	var blocklist []Block
	for i := 0; i < 16; i++ {
		var s string
		fmt.Scanf("%s\n", &s)
		// 初始化区块
		block := Block{
			index:  i,
			hash:   calSha256(s),
			lchild: nil,
			rchild: nil,
		}
		// 将区块装入blocklist
		blocklist = append(blocklist, block)
	}
	_, root2 := genMerkleTree(blocklist)
	fmt.Print("不同位置为: " + strconv.Itoa(compareMerkleTree(root1, root2)))
}
