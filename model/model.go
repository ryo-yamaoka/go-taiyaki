package model

// Taiyaki は鯛焼きを抽象化したインターフェース
type Taiyaki interface {
	GetNakami() string
	SetNakami(nakami string)
}
