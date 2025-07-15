// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DeleteCloudDriveUsersRequest
	GetCdsId() *string
	SetEndUserId(v []*string) *DeleteCloudDriveUsersRequest
	GetEndUserId() []*string
	SetRegionId(v string) *DeleteCloudDriveUsersRequest
	GetRegionId() *string
}

type DeleteCloudDriveUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-64326*****
	CdsId     *string   `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCloudDriveUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveUsersRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DeleteCloudDriveUsersRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *DeleteCloudDriveUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCloudDriveUsersRequest) SetCdsId(v string) *DeleteCloudDriveUsersRequest {
	s.CdsId = &v
	return s
}

func (s *DeleteCloudDriveUsersRequest) SetEndUserId(v []*string) *DeleteCloudDriveUsersRequest {
	s.EndUserId = v
	return s
}

func (s *DeleteCloudDriveUsersRequest) SetRegionId(v string) *DeleteCloudDriveUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCloudDriveUsersRequest) Validate() error {
	return dara.Validate(s)
}
