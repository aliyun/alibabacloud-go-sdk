// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateWarehouseScheduleTaskResponseBody
	GetData() *string
	SetRequestId(v string) *CreateWarehouseScheduleTaskResponseBody
	GetRequestId() *string
}

type CreateWarehouseScheduleTaskResponseBody struct {
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWarehouseScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarehouseScheduleTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateWarehouseScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWarehouseScheduleTaskResponseBody) SetData(v string) *CreateWarehouseScheduleTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarehouseScheduleTaskResponseBody) SetRequestId(v string) *CreateWarehouseScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarehouseScheduleTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
