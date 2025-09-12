// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewLindormInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *RenewLindormInstanceRequest
	GetDuration() *int32
	SetInstanceId(v string) *RenewLindormInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RenewLindormInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RenewLindormInstanceRequest
	GetOwnerId() *int64
	SetPricingCycle(v string) *RenewLindormInstanceRequest
	GetPricingCycle() *string
	SetRegionId(v string) *RenewLindormInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RenewLindormInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RenewLindormInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *RenewLindormInstanceRequest
	GetSecurityToken() *string
}

type RenewLindormInstanceRequest struct {
	// The subscription duration of the instance. The valid values of this parameter depend on the value of the PricingCycle parameter.
	//
	// 	- If PricingCycle is set to **Month**, set this parameter to an integer that ranges from **1*	- to **9**.
	//
	// 	- If PricingCycle is set to **Year**, set this parameter to an integer that ranges from **1*	- to **3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance that you want to renew. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The period based on which you are charged for the instance. Valid values:
	//
	// 	- **Month**: You are charged for the instance based on months.
	//
	// 	- **Year**: You are charged for the instance based on years.
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the region in which the instance that you want to renew is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RenewLindormInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewLindormInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewLindormInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewLindormInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RenewLindormInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewLindormInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RenewLindormInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewLindormInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RenewLindormInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RenewLindormInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RenewLindormInstanceRequest) SetDuration(v int32) *RenewLindormInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetInstanceId(v string) *RenewLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetOwnerAccount(v string) *RenewLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetOwnerId(v int64) *RenewLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetPricingCycle(v string) *RenewLindormInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetRegionId(v string) *RenewLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetResourceOwnerAccount(v string) *RenewLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetResourceOwnerId(v int64) *RenewLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetSecurityToken(v string) *RenewLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *RenewLindormInstanceRequest) Validate() error {
	return dara.Validate(s)
}
