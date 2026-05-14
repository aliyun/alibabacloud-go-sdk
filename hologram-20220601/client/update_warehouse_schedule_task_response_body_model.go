// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarehouseScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateWarehouseScheduleTaskResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateWarehouseScheduleTaskResponseBody
	GetRequestId() *string
}

type UpdateWarehouseScheduleTaskResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWarehouseScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarehouseScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWarehouseScheduleTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateWarehouseScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWarehouseScheduleTaskResponseBody) SetData(v bool) *UpdateWarehouseScheduleTaskResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskResponseBody) SetRequestId(v string) *UpdateWarehouseScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
