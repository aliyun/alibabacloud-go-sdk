// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *ConvertClusterResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *ConvertClusterResponseBody
	GetRequestId() *string
}

type ConvertClusterResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertClusterResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ConvertClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertClusterResponseBody) SetOrderId(v int64) *ConvertClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertClusterResponseBody) SetRequestId(v string) *ConvertClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
