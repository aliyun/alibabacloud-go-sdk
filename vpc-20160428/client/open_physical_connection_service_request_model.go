// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPhysicalConnectionServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *OpenPhysicalConnectionServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenPhysicalConnectionServiceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenPhysicalConnectionServiceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *OpenPhysicalConnectionServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenPhysicalConnectionServiceRequest
	GetResourceOwnerId() *int64
}

type OpenPhysicalConnectionServiceRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Express Connect circuit is deployed.
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
}

func (s OpenPhysicalConnectionServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenPhysicalConnectionServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenPhysicalConnectionServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenPhysicalConnectionServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenPhysicalConnectionServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenPhysicalConnectionServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenPhysicalConnectionServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenPhysicalConnectionServiceRequest) SetOwnerAccount(v string) *OpenPhysicalConnectionServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenPhysicalConnectionServiceRequest) SetOwnerId(v int64) *OpenPhysicalConnectionServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenPhysicalConnectionServiceRequest) SetRegionId(v string) *OpenPhysicalConnectionServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenPhysicalConnectionServiceRequest) SetResourceOwnerAccount(v string) *OpenPhysicalConnectionServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenPhysicalConnectionServiceRequest) SetResourceOwnerId(v int64) *OpenPhysicalConnectionServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenPhysicalConnectionServiceRequest) Validate() error {
	return dara.Validate(s)
}
