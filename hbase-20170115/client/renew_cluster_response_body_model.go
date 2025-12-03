// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *RenewClusterResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *RenewClusterResponseBody
	GetRequestId() *string
}

type RenewClusterResponseBody struct {
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RenewClusterResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *RenewClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewClusterResponseBody) SetOrderId(v int64) *RenewClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewClusterResponseBody) SetRequestId(v string) *RenewClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
