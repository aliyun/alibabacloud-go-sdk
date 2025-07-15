// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkPackageId(v string) *CreateNetworkPackageResponseBody
	GetNetworkPackageId() *string
	SetOrderId(v string) *CreateNetworkPackageResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateNetworkPackageResponseBody
	GetRequestId() *string
}

type CreateNetworkPackageResponseBody struct {
	// The ID of the premium bandwidth plan.
	//
	// example:
	//
	// np-amtp8e8q1o9e4****
	NetworkPackageId *string `json:"NetworkPackageId,omitempty" xml:"NetworkPackageId,omitempty"`
	// The ID of the bill.
	//
	// example:
	//
	// 234526262716724
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkPackageResponseBody) GetNetworkPackageId() *string {
	return s.NetworkPackageId
}

func (s *CreateNetworkPackageResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateNetworkPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkPackageResponseBody) SetNetworkPackageId(v string) *CreateNetworkPackageResponseBody {
	s.NetworkPackageId = &v
	return s
}

func (s *CreateNetworkPackageResponseBody) SetOrderId(v string) *CreateNetworkPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateNetworkPackageResponseBody) SetRequestId(v string) *CreateNetworkPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
