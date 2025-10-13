// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ScaleHoloWarehouseResponseBody
	GetData() *bool
	SetRequestId(v string) *ScaleHoloWarehouseResponseBody
	GetRequestId() *string
}

type ScaleHoloWarehouseResponseBody struct {
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

func (s ScaleHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleHoloWarehouseResponseBody) GetData() *bool {
	return s.Data
}

func (s *ScaleHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleHoloWarehouseResponseBody) SetData(v bool) *ScaleHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *ScaleHoloWarehouseResponseBody) SetRequestId(v string) *ScaleHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
