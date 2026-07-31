// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLockedSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DescribeLockedSnapshotsRequest
	GetDryRun() *bool
	SetLockStatus(v string) *DescribeLockedSnapshotsRequest
	GetLockStatus() *string
	SetMaxResults(v int32) *DescribeLockedSnapshotsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLockedSnapshotsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeLockedSnapshotsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLockedSnapshotsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLockedSnapshotsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLockedSnapshotsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLockedSnapshotsRequest
	GetResourceOwnerId() *int64
	SetSnapshotIds(v []*string) *DescribeLockedSnapshotsRequest
	GetSnapshotIds() []*string
}

type DescribeLockedSnapshotsRequest struct {
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run. The system checks the request for potential issues, including required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - false (default): performs a dry run and sends the request. If the request passes the dry run, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The lock status. Valid values:
	//
	// - compliance-cooloff: The snapshot is locked in compliance mode but is still within the cooling-off period. The snapshot cannot be deleted, but users with the required RAM permissions can unlock the snapshot, extend or shorten the cooling-off period, or extend or shorten the lock duration.
	//
	// - compliance: The snapshot is locked in compliance mode and the cooling-off period has ended. The snapshot cannot be unlocked or deleted, but users with the required RAM permissions can extend the lock duration.
	//
	// - expired: The snapshot was previously locked, but the lock duration has ended and the lock has expired. The snapshot is currently unlocked and can be deleted.
	//
	// example:
	//
	// compliance-cooloff
	LockStatus *string `json:"LockStatus,omitempty" xml:"LockStatus,omitempty"`
	// The maximum number of entries per page for a paging query. Maximum value: 100.
	//
	// Default value:
	//
	// - If this parameter is not specified or is set to a value less than 10, the default value is 10.
	//
	// - If this parameter is set to a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the `NextToken` value returned in the previous API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-ecs-2014-05-26-describeregions) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 158704318252****
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 158704318252****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of snapshot IDs. Valid values of array length: 1 to 100.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty" type:"Repeated"`
}

func (s DescribeLockedSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLockedSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLockedSnapshotsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeLockedSnapshotsRequest) GetLockStatus() *string {
	return s.LockStatus
}

func (s *DescribeLockedSnapshotsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLockedSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLockedSnapshotsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLockedSnapshotsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLockedSnapshotsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLockedSnapshotsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLockedSnapshotsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLockedSnapshotsRequest) GetSnapshotIds() []*string {
	return s.SnapshotIds
}

func (s *DescribeLockedSnapshotsRequest) SetDryRun(v bool) *DescribeLockedSnapshotsRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetLockStatus(v string) *DescribeLockedSnapshotsRequest {
	s.LockStatus = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetMaxResults(v int32) *DescribeLockedSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetNextToken(v string) *DescribeLockedSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetOwnerAccount(v string) *DescribeLockedSnapshotsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetOwnerId(v int64) *DescribeLockedSnapshotsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetRegionId(v string) *DescribeLockedSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetResourceOwnerAccount(v string) *DescribeLockedSnapshotsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetResourceOwnerId(v int64) *DescribeLockedSnapshotsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLockedSnapshotsRequest) SetSnapshotIds(v []*string) *DescribeLockedSnapshotsRequest {
	s.SnapshotIds = v
	return s
}

func (s *DescribeLockedSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
