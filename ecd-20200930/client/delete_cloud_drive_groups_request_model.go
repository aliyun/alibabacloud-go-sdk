// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DeleteCloudDriveGroupsRequest
	GetCdsId() *string
	SetGroupId(v []*string) *DeleteCloudDriveGroupsRequest
	GetGroupId() []*string
	SetRegionId(v string) *DeleteCloudDriveGroupsRequest
	GetRegionId() *string
}

type DeleteCloudDriveGroupsRequest struct {
	// The enterprise network drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-532033****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The list of team space IDs.
	GroupId []*string `json:"GroupId,omitempty" xml:"GroupId,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCloudDriveGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveGroupsRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DeleteCloudDriveGroupsRequest) GetGroupId() []*string {
	return s.GroupId
}

func (s *DeleteCloudDriveGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCloudDriveGroupsRequest) SetCdsId(v string) *DeleteCloudDriveGroupsRequest {
	s.CdsId = &v
	return s
}

func (s *DeleteCloudDriveGroupsRequest) SetGroupId(v []*string) *DeleteCloudDriveGroupsRequest {
	s.GroupId = v
	return s
}

func (s *DeleteCloudDriveGroupsRequest) SetRegionId(v string) *DeleteCloudDriveGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCloudDriveGroupsRequest) Validate() error {
	return dara.Validate(s)
}
