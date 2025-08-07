// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManuallyStartDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ManuallyStartDBClusterRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ManuallyStartDBClusterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ManuallyStartDBClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ManuallyStartDBClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ManuallyStartDBClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ManuallyStartDBClusterRequest
	GetResourceOwnerId() *int64
}

type ManuallyStartDBClusterRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-xxxxxxxxxxxxx
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ManuallyStartDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ManuallyStartDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ManuallyStartDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ManuallyStartDBClusterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ManuallyStartDBClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ManuallyStartDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ManuallyStartDBClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ManuallyStartDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ManuallyStartDBClusterRequest) SetDBClusterId(v string) *ManuallyStartDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ManuallyStartDBClusterRequest) SetOwnerAccount(v string) *ManuallyStartDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ManuallyStartDBClusterRequest) SetOwnerId(v int64) *ManuallyStartDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ManuallyStartDBClusterRequest) SetRegionId(v string) *ManuallyStartDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ManuallyStartDBClusterRequest) SetResourceOwnerAccount(v string) *ManuallyStartDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ManuallyStartDBClusterRequest) SetResourceOwnerId(v int64) *ManuallyStartDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ManuallyStartDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
