// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminUserIds(v []*string) *CreateCloudDriveGroupRequest
	GetAdminUserIds() []*string
	SetCdsId(v string) *CreateCloudDriveGroupRequest
	GetCdsId() *string
	SetGroupId(v string) *CreateCloudDriveGroupRequest
	GetGroupId() *string
	SetRegionId(v string) *CreateCloudDriveGroupRequest
	GetRegionId() *string
	SetTotalSize(v int64) *CreateCloudDriveGroupRequest
	GetTotalSize() *int64
}

type CreateCloudDriveGroupRequest struct {
	// if can be null:
	// true
	AdminUserIds []*string `json:"AdminUserIds,omitempty" xml:"AdminUserIds,omitempty" type:"Repeated"`
	// The ID of the cloud disk in Cloud Drive Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-532033****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the team.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7f4bbccda0cf40bb85981b65fb5e****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The size of the cloud disk in Cloud Drive Service. Unit: bytes. The size is unlimited.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024000
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s CreateCloudDriveGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveGroupRequest) GetAdminUserIds() []*string {
	return s.AdminUserIds
}

func (s *CreateCloudDriveGroupRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *CreateCloudDriveGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateCloudDriveGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudDriveGroupRequest) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *CreateCloudDriveGroupRequest) SetAdminUserIds(v []*string) *CreateCloudDriveGroupRequest {
	s.AdminUserIds = v
	return s
}

func (s *CreateCloudDriveGroupRequest) SetCdsId(v string) *CreateCloudDriveGroupRequest {
	s.CdsId = &v
	return s
}

func (s *CreateCloudDriveGroupRequest) SetGroupId(v string) *CreateCloudDriveGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateCloudDriveGroupRequest) SetRegionId(v string) *CreateCloudDriveGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCloudDriveGroupRequest) SetTotalSize(v int64) *CreateCloudDriveGroupRequest {
	s.TotalSize = &v
	return s
}

func (s *CreateCloudDriveGroupRequest) Validate() error {
	return dara.Validate(s)
}
