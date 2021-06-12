package club

import (
	"fmt"
	"github.com/chipgo/observer/pattern"
	"reflect"
	"time"
)

type (
	FanClub struct {
		name    string
		point   int
		fanList []pattern.Observer
	}

	Fan struct {
		name     string
		birthDay time.Time
		address  string
	}
)

func (f *Fan) Update(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Int {
		fmt.Printf("%s nhận được cập nhật kết quả mới : %d\n", f.name, v.Int())
	}
}

func (f *FanClub) RegisterObserver(fan pattern.Observer) {
	f.fanList = append(f.fanList, fan)
}

func (f *FanClub) RemoveObserver(fan pattern.Observer) {
	var found bool
	var i int

	for ; i < len(f.fanList); i++ {
		if f.fanList[i] == fan {
			found = true
			break
		}
	}

	if found {
		f.fanList = append(f.fanList[:i], f.fanList[i+1:]...)
	}
}

func (f *FanClub) Notify() {
	for _, fan := range f.fanList {
		fan.Update(f.point)
	}
}

func (f *FanClub) SetPoint(point int) {
	f.point = point
	f.Notify()
}

func (f *FanClub) RegisterObservers(fans ...pattern.Observer) {
	for _, fan := range fans {
		f.RegisterObserver(fan)
	}
}

func NewFanClub(name string) *FanClub {
	return &FanClub{
		name:    name,
		point:   0,
		fanList: make([]pattern.Observer, 0), // []pattern.Observer{}
	}
}

func NewFan(name string) *Fan {
	return &Fan{
		name:     name,
		birthDay: time.Now(),
		address:  "",
	}
}
