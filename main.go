package main

import (
	"fmt"

	"github.com/ryo-yamaoka/go-taiyaki/model"
	"github.com/ryo-yamaoka/go-taiyaki/normal"
	"github.com/ryo-yamaoka/go-taiyaki/white"
)

func main() {
	t1 := normal.NewNormalTaiyaki()
	t2 := white.NewWhiteTaiyaki()

	makeTaiyaki(t1, "あんこ")
	makeTaiyaki(t2, "クリーム")

	fmt.Println(t1.GetNakami())
	fmt.Println(t2.GetNakami())
}

func makeTaiyaki(t model.Taiyaki, nakami string) model.Taiyaki {
	t.SetNakami(nakami)
	return t
}
