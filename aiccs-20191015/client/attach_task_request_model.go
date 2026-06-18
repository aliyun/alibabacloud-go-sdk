// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallString(v string) *AttachTaskRequest
	GetCallString() *string
	SetOwnerId(v int64) *AttachTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AttachTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *AttachTaskRequest
	GetTaskId() *int64
}

type AttachTaskRequest struct {
	// The calling string (callee information and parameter list). Valid values:
	//
	// - **LIST**: Use this type when the script has no input variables. In this case, only the callee numbers need to be provided. Example: `0571****5678,0571****5679`.
	//
	// - **JSON**: Use this type when the script includes input variables. You must provide the variable names, callee numbers, and variable values. Example: `{"ParamNames":["name","age"],"CalleeList":[{"Callee":"181****0000","Params":["Zhang San","20"]},{"Callee":"181****0001","Params":["Li Si","21"]}]}`. **ParamNames*	- represents the list of parameter names; **Params*	- represents the list of parameter values.
	//
	// > You can view the script input variables on the [**Script Management**](https://aiccs.console.aliyun.com/patter/list) > **View*	- > **Input and Output Parameters*	- interface.
	//
	// example:
	//
	// {
	//
	//   "ParamNames": [
	//
	//     "name",
	//
	//     "age"
	//
	//   ],
	//
	//   "CalleeList": [
	//
	//     {
	//
	//       "Callee": "181****0000",
	//
	//       "Params": [
	//
	//         "张三",
	//
	//         "20"
	//
	//       ]
	//
	//     },
	//
	//     {
	//
	//       "Callee": "181****0001",
	//
	//       "Params": [
	//
	//         "李四",
	//
	//         "21"
	//
	//       ]
	//
	//     }
	//
	//   ]
	//
	// }
	CallString           *string `json:"CallString,omitempty" xml:"CallString,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The job ID. You can obtain the job ID from the [Task Management](https://aiccs.console.aliyun.com/job/list) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AttachTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachTaskRequest) GoString() string {
	return s.String()
}

func (s *AttachTaskRequest) GetCallString() *string {
	return s.CallString
}

func (s *AttachTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AttachTaskRequest) SetCallString(v string) *AttachTaskRequest {
	s.CallString = &v
	return s
}

func (s *AttachTaskRequest) SetOwnerId(v int64) *AttachTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachTaskRequest) SetResourceOwnerAccount(v string) *AttachTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachTaskRequest) SetResourceOwnerId(v int64) *AttachTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachTaskRequest) SetTaskId(v int64) *AttachTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AttachTaskRequest) Validate() error {
	return dara.Validate(s)
}
