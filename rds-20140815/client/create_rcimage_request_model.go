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
	ImageName  *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
