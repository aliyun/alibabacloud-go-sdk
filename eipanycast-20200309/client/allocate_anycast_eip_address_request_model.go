// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateAnycastEipAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *AllocateAnycastEipAddressRequest
	GetBandwidth() *string
	SetClientToken(v string) *AllocateAnycastEipAddressRequest
	GetClientToken() *string
	SetDescription(v string) *AllocateAnycastEipAddressRequest
	GetDescription() *string
	SetInstanceChargeType(v string) *AllocateAnycastEipAddressRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *AllocateAnycastEipAddressRequest
	GetInternetChargeType() *string
	SetName(v string) *AllocateAnycastEipAddressRequest
	GetName() *string
	SetResourceGroupId(v string) *AllocateAnycastEipAddressRequest
	GetResourceGroupId() *string
	SetServiceLocation(v string) *AllocateAnycastEipAddressRequest
	GetServiceLocation() *string
}

type AllocateAnycastEipAddressRequest struct {
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	//
	// Valid values: **200*	- to **1000**.
	//
	// Default value: **1000**.
	//
	// > The maximum bandwidth is not a guaranteed service and is for reference only.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 200
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Anycast EIP.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// docdesc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the Anycast EIP.
	//
	// Set the value to **PostPaid**, which specifies the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the Anycast EIP.
	//
	// Set the value to **PayByTraffic**, which specifies the pay-by-data-transfer metering method.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The name of the Anycast EIP.
	//
	// The name must be 0 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// doctest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfm3obzjukv53a
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The access area of the Anycast EIP.
	//
	// Set the value to **international**, which specifies the areas outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// international
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
}

func (s AllocateAnycastEipAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateAnycastEipAddressRequest) GoString() string {
	return s.String()
}

func (s *AllocateAnycastEipAddressRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *AllocateAnycastEipAddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateAnycastEipAddressRequest) GetDescription() *string {
	return s.Description
}

func (s *AllocateAnycastEipAddressRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *AllocateAnycastEipAddressRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *AllocateAnycastEipAddressRequest) GetName() *string {
	return s.Name
}

func (s *AllocateAnycastEipAddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AllocateAnycastEipAddressRequest) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *AllocateAnycastEipAddressRequest) SetBandwidth(v string) *AllocateAnycastEipAddressRequest {
	s.Bandwidth = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetClientToken(v string) *AllocateAnycastEipAddressRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetDescription(v string) *AllocateAnycastEipAddressRequest {
	s.Description = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetInstanceChargeType(v string) *AllocateAnycastEipAddressRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetInternetChargeType(v string) *AllocateAnycastEipAddressRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetName(v string) *AllocateAnycastEipAddressRequest {
	s.Name = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetResourceGroupId(v string) *AllocateAnycastEipAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) SetServiceLocation(v string) *AllocateAnycastEipAddressRequest {
	s.ServiceLocation = &v
	return s
}

func (s *AllocateAnycastEipAddressRequest) Validate() error {
	return dara.Validate(s)
}
