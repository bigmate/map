package maps

import (
	"testing"
)

type potato struct {
	id   int
	sign int
}

func TestFromSlice(t *testing.T) {
	pts := make([]potato, 0, 10)

	for i := 0; i < 10; i++ {
		p := potato{
			id:   i + 1,
			sign: -1,
		}
		pts = append(pts, p, p)
	}

	m := FromSlice(pts, func(v potato) int {
		return v.id
	})

	for i := 0; i < 10; i++ {
		if !m.Has(i + 1) {
			t.Errorf("key does not exist")
		}
	}

	for i := 100; i <= 300; i += 100 {
		p := potato{id: i}

		m.Set(i, p)

		if m.Get(i) != p {
			t.Errorf("invalid value")
		}
	}

	ptsp := make([]*potato, 0, len(pts))

	for i := range pts {
		ptsp = append(ptsp, &pts[i])
	}

	mp := FromSlice(ptsp, func(v *potato) int {
		return v.id
	})

	for _, p := range ptsp {
		if !mp.Has(p.id) {
			t.Errorf("key does not exist")
		}
	}

	t.Log(m)
	t.Log(mp)
}
