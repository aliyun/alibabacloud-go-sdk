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
	// The ID of the cloud disk in Cloud Drive Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-135515****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// testDirectoryName
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The team space status. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: enabled.
	//
	// example:
	//
	// enabled
	DriveStatus *string `json:"DriveStatus,omitempty" xml:"DriveStatus,omitempty"`
	// Specifies whether the space is increased.
	//
	// 	- binding: increased
	//
	// 	- unbound: not increased
	//
	// Default value: null. The default value indicates that all spaces are queried.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// binding
	DriveType *string `json:"DriveType,omitempty" xml:"DriveType,omitempty"`
	// The team ID.
	GroupId []*string `json:"GroupId,omitempty" xml:"GroupId,omitempty" type:"Repeated"`
	// The team name for fuzzy search.
	//
	// example:
	//
	// Test Team 1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The team type.
	//
	// 	- org: organizational structure
	//
	// 	- directory: workspace
	//
	// Default value: null. The default value indicates that all types of teams are queried.
	//
	// example:
	//
	// org
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- Valid values: 1 to 100
	//
	// 	- Default value: 20
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lY3I2VNHLwy+nIoSXh****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the parent node. If a parent node ID is specified, the subnodes are queried. If you set the value of this parameter to root, the root node is queried.
	//
	// Default value: null. The default value indicates that all nodes are queried.
	//
	// example:
	//
	// cg-e70ga4ixp30ur****
	ParentGroupId *string `json:"ParentGroupId,omitempty" xml:"ParentGroupId,omitempty"`
	// The region ID.
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
