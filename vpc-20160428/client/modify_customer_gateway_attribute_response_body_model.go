// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomerGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *ModifyCustomerGatewayAttributeResponseBody
	GetCreateTime() *int64
	SetCustomerGatewayId(v string) *ModifyCustomerGatewayAttributeResponseBody
	GetCustomerGatewayId() *string
	SetDescription(v string) *ModifyCustomerGatewayAttributeResponseBody
	GetDescription() *string
	SetIpAddress(v string) *ModifyCustomerGatewayAttributeResponseBody
	GetIpAddress() *string
	SetName(v string) *ModifyCustomerGatewayAttributeResponseBody
	GetName() *string
	SetRequestId(v string) *ModifyCustomerGatewayAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *ModifyCustomerGatewayAttributeResponseBody
	GetResourceGroupId() *string
}

type ModifyCustomerGatewayAttributeResponseBody struct {
	// The timestamp generated when the customer gateway was created.
	//
	// example:
	//
	// 1492747187000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the customer gateway.
	//
	// example:
	//
	// cgw-bp1pvpl9r9adju6l5****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// The description of the customer gateway.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The public IP address of the gateway device in the data center.
	//
	// example:
	//
	// 139.32.XX.XX
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
	// 8AA5CE21-2E6A-4530-BDF5-F055849476E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the customer gateway belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyCustomerGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomerGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomerGatewayAttributeResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ModifyCustomerGatewayAttributeResponseBody) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *ModifyCustomerGatewayAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ModifyCustomerGatewayAttributeResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *ModifyCustomerGatewayAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifyCustomerGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomerGatewayAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyCustomerGatewayAttributeResponseBody) SetCreateTime(v int64) *ModifyCustomerGatewayAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponseBody) SetCustomerGatewayId(v string) *ModifyCustomerGatewayAttributeResponseBody {
	s.CustomerGatewayId = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponseBody) SetDescription(v string) *ModifyCustomerGatewayAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponseBody) SetIpAddress(v string) *ModifyCustomerGatewayAttributeResponseBody {
	s.IpAddress = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponseBody) SetName(v string) *ModifyCustomerGatewayAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponseBody) SetRequestId(v string) *ModifyCustomerGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponseBody) SetResourceGroupId(v string) *ModifyCustomerGatewayAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyCustomerGatewayAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
