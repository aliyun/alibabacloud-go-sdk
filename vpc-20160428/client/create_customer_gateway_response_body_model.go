// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *CreateCustomerGatewayResponseBody
	GetCreateTime() *int64
	SetCustomerGatewayId(v string) *CreateCustomerGatewayResponseBody
	GetCustomerGatewayId() *string
	SetDescription(v string) *CreateCustomerGatewayResponseBody
	GetDescription() *string
	SetIpAddress(v string) *CreateCustomerGatewayResponseBody
	GetIpAddress() *string
	SetName(v string) *CreateCustomerGatewayResponseBody
	GetName() *string
	SetRequestId(v string) *CreateCustomerGatewayResponseBody
	GetRequestId() *string
}

type CreateCustomerGatewayResponseBody struct {
	// The timestamp generated when the customer gateway was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1493363486000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the customer gateway.
	//
	// example:
	//
	// cgw-bp1jrawp82av6bws9****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// The description of the customer gateway.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The static IP address of the gateway device in the on-premises data center.
	//
	// example:
	//
	// 101.12.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the customer gateway.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D32B3C26-6C6C-4988-93E9-D2A6444CE6AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomerGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerGatewayResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateCustomerGatewayResponseBody) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *CreateCustomerGatewayResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomerGatewayResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateCustomerGatewayResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateCustomerGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomerGatewayResponseBody) SetCreateTime(v int64) *CreateCustomerGatewayResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateCustomerGatewayResponseBody) SetCustomerGatewayId(v string) *CreateCustomerGatewayResponseBody {
	s.CustomerGatewayId = &v
	return s
}

func (s *CreateCustomerGatewayResponseBody) SetDescription(v string) *CreateCustomerGatewayResponseBody {
	s.Description = &v
	return s
}

func (s *CreateCustomerGatewayResponseBody) SetIpAddress(v string) *CreateCustomerGatewayResponseBody {
	s.IpAddress = &v
	return s
}

func (s *CreateCustomerGatewayResponseBody) SetName(v string) *CreateCustomerGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *CreateCustomerGatewayResponseBody) SetRequestId(v string) *CreateCustomerGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
