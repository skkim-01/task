package Task

type TaskCallback func(cbResult []interface{})

// Task struct definition
type Task struct {
	// Task Running interface : struct
	IFace interface{}

	// Task Runner Function Name
	RunnerFunction string

	// Callback Function
	CallbackFunction TaskCallback

	// Function Parameters
	RunnerParam []interface{}
}
