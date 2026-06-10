// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *CreateCloudDriveUsersRequest
	GetCdsId() *string
	SetEndUserId(v []*string) *CreateCloudDriveUsersRequest
	GetEndUserId() []*string
	SetRegionId(v string) *CreateCloudDriveUsersRequest
	GetRegionId() *string
	SetUserMaxSize(v int64) *CreateCloudDriveUsersRequest
	GetUserMaxSize() *int64
}

type CreateCloudDriveUsersRequest struct {
	// Enterprise cloud drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-352282****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// List of end user IDs.
	//
	// This parameter is required.
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The ID of the region. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to obtain a list of regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Maximum storage size for a user\\"s personal cloud drive. This value must not exceed the remaining available capacity in the enterprise cloud drive. Unit: byte.
	//
	// This parameter is required.
	//
	// example:
	//
	// 209715200
	UserMaxSize *int64 `json:"UserMaxSize,omitempty" xml:"UserMaxSize,omitempty"`
}

func (s CreateCloudDriveUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveUsersRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveUsersRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *CreateCloudDriveUsersRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *CreateCloudDriveUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudDriveUsersRequest) GetUserMaxSize() *int64 {
	return s.UserMaxSize
}

func (s *CreateCloudDriveUsersRequest) SetCdsId(v string) *CreateCloudDriveUsersRequest {
	s.CdsId = &v
	return s
}

func (s *CreateCloudDriveUsersRequest) SetEndUserId(v []*string) *CreateCloudDriveUsersRequest {
	s.EndUserId = v
	return s
}

func (s *CreateCloudDriveUsersRequest) SetRegionId(v string) *CreateCloudDriveUsersRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCloudDriveUsersRequest) SetUserMaxSize(v int64) *CreateCloudDriveUsersRequest {
	s.UserMaxSize = &v
	return s
}

func (s *CreateCloudDriveUsersRequest) Validate() error {
	return dara.Validate(s)
}
