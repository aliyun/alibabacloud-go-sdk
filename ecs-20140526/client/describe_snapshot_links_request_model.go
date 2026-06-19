// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotLinksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIds(v string) *DescribeSnapshotLinksRequest
	GetDiskIds() *string
	SetInstanceId(v string) *DescribeSnapshotLinksRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeSnapshotLinksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSnapshotLinksRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeSnapshotLinksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnapshotLinksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSnapshotLinksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotLinksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSnapshotLinksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSnapshotLinksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnapshotLinksRequest
	GetResourceOwnerId() *int64
	SetSnapshotLinkIds(v string) *DescribeSnapshotLinksRequest
	GetSnapshotLinkIds() *string
}

type DescribeSnapshotLinksRequest struct {
	// The IDs of disks. You can specify up to 100 disk IDs at a time. The DiskIds parameter is a JSON array. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["d-bp1d6tsvznfghy7y****", "d-bp1ippxbaql9zet7****", … "d-bp1ib7bcz07lcxa9****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1h6jmbefj2cyqs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page for a paged query. Maximum value: 100.
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
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the disk status list. Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disk. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of snapshot chains. You can specify up to 100 snapshot chain IDs at a time. The SnapshotLinkIds parameter is a JSON array. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["sl-bp1grgphbcc9brb5****", "sl-bp1c4izumvq0i5bs****", … "sl-bp1akk7isz866dds****"]
	SnapshotLinkIds *string `json:"SnapshotLinkIds,omitempty" xml:"SnapshotLinkIds,omitempty"`
}

func (s DescribeSnapshotLinksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotLinksRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotLinksRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeSnapshotLinksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotLinksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnapshotLinksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotLinksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnapshotLinksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnapshotLinksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotLinksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotLinksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotLinksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnapshotLinksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnapshotLinksRequest) GetSnapshotLinkIds() *string {
	return s.SnapshotLinkIds
}

func (s *DescribeSnapshotLinksRequest) SetDiskIds(v string) *DescribeSnapshotLinksRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetInstanceId(v string) *DescribeSnapshotLinksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetMaxResults(v int32) *DescribeSnapshotLinksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetNextToken(v string) *DescribeSnapshotLinksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetOwnerAccount(v string) *DescribeSnapshotLinksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetOwnerId(v int64) *DescribeSnapshotLinksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetPageNumber(v int32) *DescribeSnapshotLinksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetPageSize(v int32) *DescribeSnapshotLinksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetRegionId(v string) *DescribeSnapshotLinksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetResourceOwnerAccount(v string) *DescribeSnapshotLinksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetResourceOwnerId(v int64) *DescribeSnapshotLinksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetSnapshotLinkIds(v string) *DescribeSnapshotLinksRequest {
	s.SnapshotLinkIds = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) Validate() error {
	return dara.Validate(s)
}
