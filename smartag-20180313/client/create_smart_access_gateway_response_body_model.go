// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSmartAccessGatewayResponseBody
	GetDescription() *string
	SetName(v string) *CreateSmartAccessGatewayResponseBody
	GetName() *string
	SetOrderId(v string) *CreateSmartAccessGatewayResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateSmartAccessGatewayResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateSmartAccessGatewayResponseBody
	GetResourceGroupId() *string
	SetSmartAGId(v string) *CreateSmartAccessGatewayResponseBody
	GetSmartAGId() *string
}

type CreateSmartAccessGatewayResponseBody struct {
	// The description of the SAG instance.
	//
	// example:
	//
	// testdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the SAG instance.
	//
	// example:
	//
	// testname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20337777****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A6B9EB0F-57DB-4843-A372-04678ABF490E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the SAG instance belongs.
	//
	// example:
	//
	// rg-acfm2iu4fnc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-nylv14tghsk26c*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s CreateSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewayResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateSmartAccessGatewayResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateSmartAccessGatewayResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmartAccessGatewayResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSmartAccessGatewayResponseBody) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *CreateSmartAccessGatewayResponseBody) SetDescription(v string) *CreateSmartAccessGatewayResponseBody {
	s.Description = &v
	return s
}

func (s *CreateSmartAccessGatewayResponseBody) SetName(v string) *CreateSmartAccessGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSmartAccessGatewayResponseBody) SetOrderId(v string) *CreateSmartAccessGatewayResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSmartAccessGatewayResponseBody) SetRequestId(v string) *CreateSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmartAccessGatewayResponseBody) SetResourceGroupId(v string) *CreateSmartAccessGatewayResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSmartAccessGatewayResponseBody) SetSmartAGId(v string) *CreateSmartAccessGatewayResponseBody {
	s.SmartAGId = &v
	return s
}

func (s *CreateSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
