// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewMobileAgentPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewMobileAgentPackageResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewMobileAgentPackageResponseBody
	GetRequestId() *string
}

type RenewMobileAgentPackageResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 223684716098****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B9EFA4F-4305-5968-BAEE-BD8B8DE5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewMobileAgentPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewMobileAgentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *RenewMobileAgentPackageResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewMobileAgentPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewMobileAgentPackageResponseBody) SetOrderId(v string) *RenewMobileAgentPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewMobileAgentPackageResponseBody) SetRequestId(v string) *RenewMobileAgentPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewMobileAgentPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
