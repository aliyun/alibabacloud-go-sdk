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
	// The enterprise network drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-352282****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The list of end user IDs.
	//
	// This parameter is required.
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum storage capacity of the user\\"s personal drive. This value cannot exceed the remaining allocatable capacity of the enterprise network drive. Unit: bytes.
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
