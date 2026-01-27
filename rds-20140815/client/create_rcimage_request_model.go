// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageName(v string) *CreateRCImageRequest
	GetImageName() *string
	SetInstanceId(v string) *CreateRCImageRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateRCImageRequest
	GetRegionId() *string
	SetSnapshotId(v string) *CreateRCImageRequest
	GetSnapshotId() *string
}

type CreateRCImageRequest struct {
	// The name of the custom image.
	//
	// example:
	//
	// Created_from_rc-vma9w5z699x9********
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the RDS Custom instance.
	//
	// example:
	//
	// rc-vma9w5z699x93204****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot from which to create the custom image. You can call the DescribeRCSnapshots operation to query the snapshot ID.
	//
	// example:
	//
	// rcds-c9bjdl79vz5dx********
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateRCImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRCImageRequest) GoString() string {
	return s.String()
}

func (s *CreateRCImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateRCImageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRCImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRCImageRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateRCImageRequest) SetImageName(v string) *CreateRCImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateRCImageRequest) SetInstanceId(v string) *CreateRCImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRCImageRequest) SetRegionId(v string) *CreateRCImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRCImageRequest) SetSnapshotId(v string) *CreateRCImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateRCImageRequest) Validate() error {
	return dara.Validate(s)
}
