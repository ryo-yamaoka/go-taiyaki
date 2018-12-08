package white

import "github.com/ryo-yamaoka/go-taiyaki/model"

type whiteTaiyaki struct {
	Nakami string
}

// NewWhiteTaiyaki は新しい白い鯛焼きを生成する
func NewWhiteTaiyaki() model.Taiyaki {
	return &whiteTaiyaki{}
}

func (t *whiteTaiyaki) GetNakami() string {
	return t.Nakami
}

func (t *whiteTaiyaki) SetNakami(nakami string) {
	t.Nakami = nakami
}
