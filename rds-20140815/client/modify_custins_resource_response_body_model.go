// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustinsResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCustinsResourceResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *ModifyCustinsResourceResponseBody
	GetTaskId() *int32
}

type ModifyCustinsResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6B5A6839-31A7-58D4-9F96-772BFAFD1CB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 507******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyCustinsResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustinsResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustinsResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustinsResourceResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyCustinsResourceResponseBody) SetRequestId(v string) *ModifyCustinsResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustinsResourceResponseBody) SetTaskId(v int32) *ModifyCustinsResourceResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyCustinsResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
