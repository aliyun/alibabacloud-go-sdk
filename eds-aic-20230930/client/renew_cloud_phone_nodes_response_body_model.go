// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCloudPhoneNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewCloudPhoneNodesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewCloudPhoneNodesResponseBody
	GetRequestId() *string
}

type RenewCloudPhoneNodesResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 22365781890****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewCloudPhoneNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewCloudPhoneNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewCloudPhoneNodesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewCloudPhoneNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewCloudPhoneNodesResponseBody) SetOrderId(v string) *RenewCloudPhoneNodesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewCloudPhoneNodesResponseBody) SetRequestId(v string) *RenewCloudPhoneNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewCloudPhoneNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
