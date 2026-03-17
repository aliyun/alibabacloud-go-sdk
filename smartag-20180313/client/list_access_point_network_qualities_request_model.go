// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointNetworkQualitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ListAccessPointNetworkQualitiesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListAccessPointNetworkQualitiesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListAccessPointNetworkQualitiesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListAccessPointNetworkQualitiesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAccessPointNetworkQualitiesRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ListAccessPointNetworkQualitiesRequest
	GetSmartAGId() *string
}

type ListAccessPointNetworkQualitiesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-p86e06z4geaji1****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ListAccessPointNetworkQualitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointNetworkQualitiesRequest) GoString() string {
	return s.String()
}

func (s *ListAccessPointNetworkQualitiesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListAccessPointNetworkQualitiesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAccessPointNetworkQualitiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccessPointNetworkQualitiesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAccessPointNetworkQualitiesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAccessPointNetworkQualitiesRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ListAccessPointNetworkQualitiesRequest) SetOwnerAccount(v string) *ListAccessPointNetworkQualitiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesRequest) SetOwnerId(v int64) *ListAccessPointNetworkQualitiesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesRequest) SetRegionId(v string) *ListAccessPointNetworkQualitiesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesRequest) SetResourceOwnerAccount(v string) *ListAccessPointNetworkQualitiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesRequest) SetResourceOwnerId(v int64) *ListAccessPointNetworkQualitiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesRequest) SetSmartAGId(v string) *ListAccessPointNetworkQualitiesRequest {
	s.SmartAGId = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesRequest) Validate() error {
	return dara.Validate(s)
}
