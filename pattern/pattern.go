package pattern

type (
	Observer interface {
		Update(v interface{})
	}

	Publisher interface {
		RegisterObserver(o Observer)
		RemoveObserver(o Observer)
		Notify()
	}
)

