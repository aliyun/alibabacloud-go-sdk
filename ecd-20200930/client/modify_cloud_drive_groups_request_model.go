// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDriveGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ModifyCloudDriveGroupsRequest
	GetCdsId() *string
	SetGroupId(v []*string) *ModifyCloudDriveGroupsRequest
	GetGroupId() []*string
	SetRegionId(v string) *ModifyCloudDriveGroupsRequest
	GetRegionId() *string
	SetStatus(v string) *ModifyCloudDriveGroupsRequest
	GetStatus() *string
	SetTotalSize(v int64) *ModifyCloudDriveGroupsRequest
	GetTotalSize() *int64
}

type ModifyCloudDriveGroupsRequest struct {
	// The ID of the cloud disk in Cloud Drive Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+cds-135515****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The IDs of the teams.
	GroupId []*string `json:"GroupId,omitempty" xml:"GroupId,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the team space. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: enabled.
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total capacity of the team space.
	//
	// example:
	//
	// 32212254720
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ModifyCloudDriveGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDriveGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudDriveGroupsRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ModifyCloudDriveGroupsRequest) GetGroupId() []*string {
	return s.GroupId
}

func (s *ModifyCloudDriveGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudDriveGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyCloudDriveGroupsRequest) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ModifyCloudDriveGroupsRequest) SetCdsId(v string) *ModifyCloudDriveGroupsRequest {
	s.CdsId = &v
	return s
}

func (s *ModifyCloudDriveGroupsRequest) SetGroupId(v []*string) *ModifyCloudDriveGroupsRequest {
	s.GroupId = v
	return s
}

func (s *ModifyCloudDriveGroupsRequest) SetRegionId(v string) *ModifyCloudDriveGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudDriveGroupsRequest) SetStatus(v string) *ModifyCloudDriveGroupsRequest {
	s.Status = &v
	return s
}

func (s *ModifyCloudDriveGroupsRequest) SetTotalSize(v int64) *ModifyCloudDriveGroupsRequest {
	s.TotalSize = &v
	return s
}

func (s *ModifyCloudDriveGroupsRequest) Validate() error {
	return dara.Validate(s)
}
