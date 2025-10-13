// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateHoloWarehouseResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateHoloWarehouseResponseBody
	GetRequestId() *string
}

type CreateHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
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

func (s CreateHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHoloWarehouseResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHoloWarehouseResponseBody) SetData(v bool) *CreateHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateHoloWarehouseResponseBody) SetRequestId(v string) *CreateHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
