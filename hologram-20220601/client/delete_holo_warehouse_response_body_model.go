// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteHoloWarehouseResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteHoloWarehouseResponseBody
	GetRequestId() *string
}

type DeleteHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHoloWarehouseResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHoloWarehouseResponseBody) SetData(v bool) *DeleteHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteHoloWarehouseResponseBody) SetRequestId(v string) *DeleteHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
