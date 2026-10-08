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
	// The O&M task IDs. Multiple IDs are separated with commas (,).
	//
	// example:
	//
	// 11111,22222
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9ADB6BE-1598-57FC-B86D-D7F4C69B****
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
