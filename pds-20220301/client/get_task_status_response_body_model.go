// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *GetTaskStatusResponseBody
	GetStatus() *string
}

type GetTaskStatusResponseBody struct {
	// The state of the task.
	//
	// Valid values:
	//
	// 	- running
	//
	//     <!-- -->
	//
	//     : The task is
	//
	//     <!-- -->
	//
	//     running
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- failed
	//
	//     <!-- -->
	//
	//     : The task
	//
	//     <!-- -->
	//
	//     fails
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- succeeded
	//
	//     <!-- -->
	//
	//     : The task is
	//
	//     <!-- -->
	//
	//     successful
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTaskStatusResponseBody) SetStatus(v string) *GetTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
