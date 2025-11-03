// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DescribeCloudDiskGroupsRequest
	GetCdsId() *string
	SetGroupId(v []*string) *DescribeCloudDiskGroupsRequest
	GetGroupId() []*string
	SetGroupName(v string) *DescribeCloudDiskGroupsRequest
	GetGroupName() *string
	SetParentOrgId(v string) *DescribeCloudDiskGroupsRequest
	GetParentOrgId() *string
	SetRegionId(v string) *DescribeCloudDiskGroupsRequest
	GetRegionId() *string
}

type DescribeCloudDiskGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-0456357992
	CdsId     *string   `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	GroupId   []*string `json:"GroupId,omitempty" xml:"GroupId,omitempty" type:"Repeated"`
	GroupName *string   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// root
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudDiskGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupsRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DescribeCloudDiskGroupsRequest) GetGroupId() []*string {
	return s.GroupId
}

func (s *DescribeCloudDiskGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeCloudDiskGroupsRequest) GetParentOrgId() *string {
	return s.ParentOrgId
}

func (s *DescribeCloudDiskGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudDiskGroupsRequest) SetCdsId(v string) *DescribeCloudDiskGroupsRequest {
	s.CdsId = &v
	return s
}

func (s *DescribeCloudDiskGroupsRequest) SetGroupId(v []*string) *DescribeCloudDiskGroupsRequest {
	s.GroupId = v
	return s
}

func (s *DescribeCloudDiskGroupsRequest) SetGroupName(v string) *DescribeCloudDiskGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeCloudDiskGroupsRequest) SetParentOrgId(v string) *DescribeCloudDiskGroupsRequest {
	s.ParentOrgId = &v
	return s
}

func (s *DescribeCloudDiskGroupsRequest) SetRegionId(v string) *DescribeCloudDiskGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudDiskGroupsRequest) Validate() error {
	return dara.Validate(s)
}
