// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetAutoSnapshotPolicyId() *string
	SetDiskId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetDiskId() *string
	SetMaxResults(v int32) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetResourceOwnerId() *int64
}

type DescribeAutoSnapshotPolicyAssociationsRequest struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	DiskId               *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	MaxResults           *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetDiskId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetMaxResults(v int32) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetRegionId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
