// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDriveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ModifyCloudDriveUsersRequest
	GetCdsId() *string
	SetEndUserId(v []*string) *ModifyCloudDriveUsersRequest
	GetEndUserId() []*string
	SetRegionId(v string) *ModifyCloudDriveUsersRequest
	GetRegionId() *string
	SetStatus(v string) *ModifyCloudDriveUsersRequest
	GetStatus() *string
	SetUserMaxSize(v int64) *ModifyCloudDriveUsersRequest
	GetUserMaxSize() *int64
}

type ModifyCloudDriveUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-596198****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// This parameter is required.
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of Cloud Drive Service users.
	//
	// Valid values:
	//
	// 	- disabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     unavailable
	//
	//     <!-- -->
	//
	// 	- enabled
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     available
	//
	//     <!-- -->
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The maximum storage space of a user. Unit: bytes.
	//
	// example:
	//
	// 1024
	UserMaxSize *int64 `json:"UserMaxSize,omitempty" xml:"UserMaxSize,omitempty"`
}

func (s ModifyCloudDriveUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDriveUsersRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudDriveUsersRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ModifyCloudDriveUsersRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *ModifyCloudDriveUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudDriveUsersRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyCloudDriveUsersRequest) GetUserMaxSize() *int64 {
	return s.UserMaxSize
}

func (s *ModifyCloudDriveUsersRequest) SetCdsId(v string) *ModifyCloudDriveUsersRequest {
	s.CdsId = &v
	return s
}

func (s *ModifyCloudDriveUsersRequest) SetEndUserId(v []*string) *ModifyCloudDriveUsersRequest {
	s.EndUserId = v
	return s
}

func (s *ModifyCloudDriveUsersRequest) SetRegionId(v string) *ModifyCloudDriveUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudDriveUsersRequest) SetStatus(v string) *ModifyCloudDriveUsersRequest {
	s.Status = &v
	return s
}

func (s *ModifyCloudDriveUsersRequest) SetUserMaxSize(v int64) *ModifyCloudDriveUsersRequest {
	s.UserMaxSize = &v
	return s
}

func (s *ModifyCloudDriveUsersRequest) Validate() error {
	return dara.Validate(s)
}
