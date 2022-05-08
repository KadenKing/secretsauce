package complicated

// secretSauce is an unexported type that must be passed as a type param in order to make protected types.
type secretSauce struct{}

// ComplicatedType must be instantiated using NewComplicatedType.
type ComplicatedType[T secretSauce] struct {
	veryImportantMustBeValid string
}

// DoComplicatedThings requires ComplicatedType to have been constructed with NewComplicatedType otherwise it panics.
func (c *ComplicatedType[T]) DoComplicatedThings() {
	if c.veryImportantMustBeValid != "hey" {
		panic("oh no everything is wrong")
	}
}

// NewComplicatedType returns a new ComplicatedType
func NewComplicatedType() ComplicatedType[secretSauce] {
	return ComplicatedType[secretSauce]{
		veryImportantMustBeValid: "hey",
	}
}
