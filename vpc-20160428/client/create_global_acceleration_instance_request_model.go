// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalAccelerationInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *CreateGlobalAccelerationInstanceRequest
	GetBandwidth() *string
	SetBandwidthType(v string) *CreateGlobalAccelerationInstanceRequest
	GetBandwidthType() *string
	SetClientToken(v string) *CreateGlobalAccelerationInstanceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateGlobalAccelerationInstanceRequest
	GetDescription() *string
	SetName(v string) *CreateGlobalAccelerationInstanceRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateGlobalAccelerationInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateGlobalAccelerationInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateGlobalAccelerationInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateGlobalAccelerationInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateGlobalAccelerationInstanceRequest
	GetResourceOwnerId() *int64
	SetServiceLocation(v string) *CreateGlobalAccelerationInstanceRequest
	GetServiceLocation() *string
}

type CreateGlobalAccelerationInstanceRequest struct {
	// The maximum bandwidth of the GA instance. Set the value to **10**. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth type. Valid values:
	//
	// 	- **Sharing**
	//
	// 	- **Exclusive**
	//
	// example:
	//
	// Exclusive
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the GA instance.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// My GA
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the GA instance.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// GA-1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the GA instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The acceleration area. Valid values:
	//
	// 	- **china-mainland**
	//
	// 	- **north-america**
	//
	// 	- **asia-pacific**
	//
	// 	- **europe**
	//
	// This parameter is required.
	//
	// example:
	//
	// china-mainland
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
}

func (s CreateGlobalAccelerationInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalAccelerationInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalAccelerationInstanceRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *CreateGlobalAccelerationInstanceRequest) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *CreateGlobalAccelerationInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGlobalAccelerationInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGlobalAccelerationInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateGlobalAccelerationInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateGlobalAccelerationInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGlobalAccelerationInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGlobalAccelerationInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateGlobalAccelerationInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateGlobalAccelerationInstanceRequest) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *CreateGlobalAccelerationInstanceRequest) SetBandwidth(v string) *CreateGlobalAccelerationInstanceRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetBandwidthType(v string) *CreateGlobalAccelerationInstanceRequest {
	s.BandwidthType = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetClientToken(v string) *CreateGlobalAccelerationInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetDescription(v string) *CreateGlobalAccelerationInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetName(v string) *CreateGlobalAccelerationInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetOwnerAccount(v string) *CreateGlobalAccelerationInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetOwnerId(v int64) *CreateGlobalAccelerationInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetRegionId(v string) *CreateGlobalAccelerationInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetResourceOwnerAccount(v string) *CreateGlobalAccelerationInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetResourceOwnerId(v int64) *CreateGlobalAccelerationInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) SetServiceLocation(v string) *CreateGlobalAccelerationInstanceRequest {
	s.ServiceLocation = &v
	return s
}

func (s *CreateGlobalAccelerationInstanceRequest) Validate() error {
	return dara.Validate(s)
}
