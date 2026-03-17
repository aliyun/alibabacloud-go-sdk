// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartAccessGatewaySoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateSmartAccessGatewaySoftwareResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateSmartAccessGatewaySoftwareResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateSmartAccessGatewaySoftwareResponseBody
	GetResourceGroupId() *string
	SetSmartAGId(v string) *CreateSmartAccessGatewaySoftwareResponseBody
	GetSmartAGId() *string
}

type CreateSmartAccessGatewaySoftwareResponseBody struct {
	// The ID of the order that you placed to purchase the SAG app instance.
	//
	// example:
	//
	// 203000000000000
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 24675405-74DF-4C94-82C6-B749580C498E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the SAG app instance belongs.
	//
	// example:
	//
	// rg-acfm2iu4fnc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the SAG app instance.
	//
	// example:
	//
	// sag-gnhe6sywtare5******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s CreateSmartAccessGatewaySoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartAccessGatewaySoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) SetOrderId(v string) *CreateSmartAccessGatewaySoftwareResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) SetRequestId(v string) *CreateSmartAccessGatewaySoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) SetResourceGroupId(v string) *CreateSmartAccessGatewaySoftwareResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) SetSmartAGId(v string) *CreateSmartAccessGatewaySoftwareResponseBody {
	s.SmartAGId = &v
	return s
}

func (s *CreateSmartAccessGatewaySoftwareResponseBody) Validate() error {
	return dara.Validate(s)
}
