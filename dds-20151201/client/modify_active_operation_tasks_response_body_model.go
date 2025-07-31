// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyActiveOperationTasksResponseBody
	GetIds() *string
	SetRequestId(v string) *ModifyActiveOperationTasksResponseBody
	GetRequestId() *string
}

type ModifyActiveOperationTasksResponseBody struct {
	// The IDs of the O\\&M tasks. Multiple task IDs are separated by commas (,).
	//
	// example:
	//
	// 11111,22222
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CAAE9DDA-65FD-584C-A378-F1F24676****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksResponseBody) GetIds() *string {
	return s.Ids
}

func (s *ModifyActiveOperationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyActiveOperationTasksResponseBody) SetIds(v string) *ModifyActiveOperationTasksResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTasksResponseBody) SetRequestId(v string) *ModifyActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActiveOperationTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
