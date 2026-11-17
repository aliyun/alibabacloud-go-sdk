// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeWuyingServerDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ResizeWuyingServerDiskResponseBodyData) *ResizeWuyingServerDiskResponseBody
	GetData() *ResizeWuyingServerDiskResponseBodyData
	SetRequestId(v string) *ResizeWuyingServerDiskResponseBody
	GetRequestId() *string
}

type ResizeWuyingServerDiskResponseBody struct {
	// The returned data object.
	Data *ResizeWuyingServerDiskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeWuyingServerDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeWuyingServerDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeWuyingServerDiskResponseBody) GetData() *ResizeWuyingServerDiskResponseBodyData {
	return s.Data
}

func (s *ResizeWuyingServerDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeWuyingServerDiskResponseBody) SetData(v *ResizeWuyingServerDiskResponseBodyData) *ResizeWuyingServerDiskResponseBody {
	s.Data = v
	return s
}

func (s *ResizeWuyingServerDiskResponseBody) SetRequestId(v string) *ResizeWuyingServerDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeWuyingServerDiskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResizeWuyingServerDiskResponseBodyData struct {
	// The ID of the order.
	//
	// example:
	//
	// 23429322113****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ResizeWuyingServerDiskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResizeWuyingServerDiskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResizeWuyingServerDiskResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeWuyingServerDiskResponseBodyData) SetOrderId(v string) *ResizeWuyingServerDiskResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ResizeWuyingServerDiskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
