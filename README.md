# task

### Package Name : Task

#### abstract
##### function container. using reflect

#### type definitions
###### Task struct
```go
  type Task struct {
    IFace             interface{}
    RunnerFunction    string
    CallbackFunction  TaskCallback
    RunnerParam       []interface{}
  }
```
###### TaskCallback struct: result callback function
```go
  type TaskCallback func(cbResult []interface{})
```

#### APIs
###### NewTask(): Create new task
```go
  func NewTask(iFace interface{}, fn string, cb TaskCallback, p ...interface{}) *Task
```

##### Task
###### InvokeTask(): Execute Task
```go
  func (t *Task) InvokeTask() []interface{}
```

###### Close(): Close task
```go
  func (t *Task) Close()
```
