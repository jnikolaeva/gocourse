package main

import "testing"

func TestGetEnding(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "ка"},
		{2, "ки"},
		{3, "ки"},
		{4, "ки"},
		{5, "ок"},
		{6, "ок"},
		{7, "ок"},
		{8, "ок"},
		{9, "ок"},
		{10, "ок"},
		{11, "ок"},
		{12, "ок"},
		{13, "ок"},
		{14, "ок"},
		{15, "ок"},
		{16, "ок"},
		{17, "ок"},
		{18, "ок"},
		{19, "ок"},
		{20, "ок"},
		{21, "ка"},
		{22, "ки"},
		{23, "ки"},
		{24, "ки"},
		{25, "ок"},
		{26, "ок"},
		{27, "ок"},
		{28, "ок"},
		{29, "ок"},
		{30, "ок"},
		{100, "ок"},
		{101, "ка"},
		{102, "ки"},
		{103, "ки"},
		{104, "ки"},
		{105, "ок"},
		{106, "ок"},
		{107, "ок"},
		{108, "ок"},
		{109, "ок"},
		{110, "ок"},
		{111, "ок"},
		{112, "ок"},
		{113, "ок"},
		{114, "ок"},
		{115, "ок"},
		{116, "ок"},
		{117, "ок"},
		{118, "ок"},
		{119, "ок"},
		{120, "ок"},
		{121, "ка"},
		{1200, "ок"},
		{1201, "ка"},
		{1202, "ки"},
		{1203, "ки"},
		{1204, "ки"},
		{1205, "ок"},
		{1206, "ок"},
		{1207, "ок"},
		{1208, "ок"},
		{1209, "ок"},
		{1210, "ок"},
		{1211, "ок"},
		{1212, "ок"},
		{1213, "ок"},
		{1214, "ок"},
		{1215, "ок"},
		{1216, "ок"},
		{1217, "ок"},
		{1218, "ок"},
		{1219, "ок"},
		{1220, "ок"},
		{1221, "ка"},
	}
	for _, c := range cases {
		got := getEnding(c.in)
		if got != c.want {
			t.Errorf("getEnding(\"%v\") == %v, want %v\n", c.in, got, c.want)
		}
	}
}