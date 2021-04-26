package Task

import (
	"fmt"
	"testing"
)

// Create struct for Task
type TaskforTest struct{}

// Create Task Function
func (t *TaskforTest) Sum_value(x, y int) (int, error) {
	return x + y, nil
}

// Create Task Callback Function
var LocalCallback TaskCallback = func(cbResult []interface{}) {
	// Print callback result
	fmt.Println("### LOG ### callback.Result:")
	fmt.Println(cbResult...)
}

func Test_main(t *testing.T) {
	// Create New Task
	t1 := NewTask(&TaskforTest{}, "Sum_value", LocalCallback, 10, 20)
	// Close Task
	defer t1.Close()

	// Invoke Task
	result := t1.InvokeTask()

	// Print invoke result
	fmt.Println("### LOG ### main.Result:")
	for i := range result {
		fmt.Println(result[i])
	}
}
