// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarehouseScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteWarehouseScheduleTaskResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteWarehouseScheduleTaskResponseBody
	GetRequestId() *string
}

type DeleteWarehouseScheduleTaskResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D1303CD4-AA70-5998-8025-F55B22C50840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWarehouseScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarehouseScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWarehouseScheduleTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteWarehouseScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWarehouseScheduleTaskResponseBody) SetData(v bool) *DeleteWarehouseScheduleTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteWarehouseScheduleTaskResponseBody) SetRequestId(v string) *DeleteWarehouseScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWarehouseScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
