// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDriveGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DescribeCloudDriveGroupsRequest
	GetCdsId() *string
	SetDirectoryId(v string) *DescribeCloudDriveGroupsRequest
	GetDirectoryId() *string
	SetDirectoryName(v string) *DescribeCloudDriveGroupsRequest
	GetDirectoryName() *string
	SetDriveStatus(v string) *DescribeCloudDriveGroupsRequest
	GetDriveStatus() *string
	SetDriveType(v string) *DescribeCloudDriveGroupsRequest
	GetDriveType() *string
	SetGroupId(v []*string) *DescribeCloudDriveGroupsRequest
	GetGroupId() []*string
	SetGroupName(v string) *DescribeCloudDriveGroupsRequest
	GetGroupName() *string
	SetGroupType(v string) *DescribeCloudDriveGroupsRequest
	GetGroupType() *string
	SetMaxResults(v int32) *DescribeCloudDriveGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCloudDriveGroupsRequest
	GetNextToken() *string
	SetParentGroupId(v string) *DescribeCloudDriveGroupsRequest
	GetParentGroupId() *string
	SetRegionId(v string) *DescribeCloudDriveGroupsRequest
	GetRegionId() *string
}

type DescribeCloudDriveGroupsRequest struct {
	// The enterprise cloud drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-135515****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The office network name.
	//
	// example:
	//
	// testDirectoryName
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The team space status.
	//
	// example:
	//
	// enabled
	DriveStatus *string `json:"DriveStatus,omitempty" xml:"DriveStatus,omitempty"`
	// Specifies whether a space is added. Default value: empty, which indicates that all types are queried.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// binding
	DriveType *string `json:"DriveType,omitempty" xml:"DriveType,omitempty"`
	// The list of team space IDs.
	GroupId []*string `json:"GroupId,omitempty" xml:"GroupId,omitempty" type:"Repeated"`
	// The team space name. Fuzzy match is supported.
	//
	// example:
	//
	// TestTeam1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The team type. Default value: empty, which indicates that all types are queried.
	//
	// example:
	//
	// org
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The number of entries per page for a paged query.
	//
	// - Maximum value: 100.
	//
	// - Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of `NextToken` that was returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lY3I2VNHLwy+nIoSXh****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the parent node of the object to query. To query the root node, set this parameter to `root`.
	//
	// Default value: empty, which indicates that all team spaces are queried.
	//
	// example:
	//
	// cg-e70ga4ixp30ur****
	ParentGroupId *string `json:"ParentGroupId,omitempty" xml:"ParentGroupId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudDriveGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveGroupsRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DescribeCloudDriveGroupsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeCloudDriveGroupsRequest) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *DescribeCloudDriveGroupsRequest) GetDriveStatus() *string {
	return s.DriveStatus
}

func (s *DescribeCloudDriveGroupsRequest) GetDriveType() *string {
	return s.DriveType
}

func (s *DescribeCloudDriveGroupsRequest) GetGroupId() []*string {
	return s.GroupId
}

func (s *DescribeCloudDriveGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCloudDriveGroupsRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeCloudDriveGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCloudDriveGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudDriveGroupsRequest) GetParentGroupId() *string {
	return s.ParentGroupId
}

func (s *DescribeCloudDriveGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudDriveGroupsRequest) SetCdsId(v string) *DescribeCloudDriveGroupsRequest {
	s.CdsId = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetDirectoryId(v string) *DescribeCloudDriveGroupsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetDirectoryName(v string) *DescribeCloudDriveGroupsRequest {
	s.DirectoryName = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetDriveStatus(v string) *DescribeCloudDriveGroupsRequest {
	s.DriveStatus = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetDriveType(v string) *DescribeCloudDriveGroupsRequest {
	s.DriveType = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetGroupId(v []*string) *DescribeCloudDriveGroupsRequest {
	s.GroupId = v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetGroupName(v string) *DescribeCloudDriveGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetGroupType(v string) *DescribeCloudDriveGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetMaxResults(v int32) *DescribeCloudDriveGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetNextToken(v string) *DescribeCloudDriveGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetParentGroupId(v string) *DescribeCloudDriveGroupsRequest {
	s.ParentGroupId = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) SetRegionId(v string) *DescribeCloudDriveGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudDriveGroupsRequest) Validate() error {
	return dara.Validate(s)
}
