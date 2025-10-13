// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartHoloWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RestartHoloWarehouseResponseBody
	GetData() *bool
	SetRequestId(v string) *RestartHoloWarehouseResponseBody
	GetRequestId() *string
}

type RestartHoloWarehouseResponseBody struct {
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

func (s RestartHoloWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *RestartHoloWarehouseResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestartHoloWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartHoloWarehouseResponseBody) SetData(v bool) *RestartHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *RestartHoloWarehouseResponseBody) SetRequestId(v string) *RestartHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartHoloWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
