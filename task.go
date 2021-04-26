package Task

import (
	"reflect"
)

// NewTask: Create task for running in thread pool
func NewTask(iFace interface{}, fn string, cb TaskCallback, p ...interface{}) *Task {
	t := Task{
		IFace:            iFace,
		RunnerFunction:   fn,
		CallbackFunction: cb,
		RunnerParam:      p,
	}
	return &t
}

// InvokeTask: Call task function
func (t *Task) InvokeTask() []interface{} {
	// parse parameters
	inputs := make([]reflect.Value, len(t.RunnerParam))
	for i := range t.RunnerParam {
		inputs[i] = reflect.ValueOf(t.RunnerParam[i])
	}

	// call function with reflect
	res := reflect.ValueOf(t.IFace).MethodByName(t.RunnerFunction).Call(inputs)

	// set invoke function return values
	returnSlice := make([]interface{}, len(res), cap(res))
	for i := range res {
		returnSlice[i] = reflect.ValueOf(res[i]).Interface()
	}

	// result callback
	if nil != t.CallbackFunction {
		t.CallbackFunction(returnSlice)
	} // else: callback is null. do nothing.

	// result return
	return returnSlice
}

// Close: Close task
func (t *Task) Close() {
	t.IFace = nil
	t.RunnerFunction = ""
	t.CallbackFunction = nil
	t.RunnerParam = nil
}
