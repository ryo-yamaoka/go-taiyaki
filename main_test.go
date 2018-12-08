package main

import "testing"

type testTaiyaki struct {
	Nakami string
}

func (t *testTaiyaki) GetNakami() string {
	return t.Nakami
}

func (t *testTaiyaki) SetNakami(nakami string) {
	t.Nakami = nakami
}

func TestMain(t *testing.T) {
	patterns := []struct {
		nakami string
	}{
		{nakami: "あんこ"},
		{nakami: "クリーム"},
		{nakami: "いちごジャム"},
		{nakami: "謎ジャム"},
	}

	for i, pattern := range patterns {
		taiyaki := &testTaiyaki{}
		makeTaiyaki(taiyaki, pattern.nakami)
		if taiyaki.GetNakami() != pattern.nakami {
			t.Errorf("unexpected nakami(%d): %s != %s", i, taiyaki.GetNakami(), pattern.nakami)
		}
	}
}
