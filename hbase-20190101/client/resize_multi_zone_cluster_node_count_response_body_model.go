// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeMultiZoneClusterNodeCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResizeMultiZoneClusterNodeCountResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResizeMultiZoneClusterNodeCountResponseBody
	GetRequestId() *string
}

type ResizeMultiZoneClusterNodeCountResponseBody struct {
	// example:
	//
	// 1234123412****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// E2B7E9DA-1575-4B9D-A0E4-9468BAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) SetOrderId(v string) *ResizeMultiZoneClusterNodeCountResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) SetRequestId(v string) *ResizeMultiZoneClusterNodeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponseBody) Validate() error {
	return dara.Validate(s)
}
