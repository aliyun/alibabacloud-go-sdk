// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteAfterImageUpload(v string) *CreateImageRequest
	GetDeleteAfterImageUpload() *string
	SetImageName(v string) *CreateImageRequest
	GetImageName() *string
	SetInstanceId(v string) *CreateImageRequest
	GetInstanceId() *string
	SetSnapshotId(v string) *CreateImageRequest
	GetSnapshotId() *string
	SetTargetOSSRegionId(v string) *CreateImageRequest
	GetTargetOSSRegionId() *string
	SetWithDataDisks(v bool) *CreateImageRequest
	GetWithDataDisks() *bool
}

type CreateImageRequest struct {
	// Specifies whether to automatically release the instance after the image is packaged and uploaded. Only image builders are supported. Default value: false. Valid values:
	//
	// 	- true: The image is released when the instance is released.
	//
	// 	- false: The image is retained when the instance is released.
	//
	// 	- If you leave this property empty, false is used by default.
	//
	// example:
	//
	// false
	DeleteAfterImageUpload *string `json:"DeleteAfterImageUpload,omitempty" xml:"DeleteAfterImageUpload,omitempty"`
	// The name of the image. The name must be 2 to 128 characters in length. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter but cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5rr1bnyrc4tswr8cq3w6y****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The region of the destination OSS bucket where the image is to be stored.
	//
	// example:
	//
	// cn-beijing
	TargetOSSRegionId *string `json:"TargetOSSRegionId,omitempty" xml:"TargetOSSRegionId,omitempty"`
	// Specifies whether to include data disk snapshots in the custom image.
	//
	// example:
	//
	// Value true false (default)
	WithDataDisks *bool `json:"WithDataDisks,omitempty" xml:"WithDataDisks,omitempty"`
}

func (s CreateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) GetDeleteAfterImageUpload() *string {
	return s.DeleteAfterImageUpload
}

func (s *CreateImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateImageRequest) GetTargetOSSRegionId() *string {
	return s.TargetOSSRegionId
}

func (s *CreateImageRequest) GetWithDataDisks() *bool {
	return s.WithDataDisks
}

func (s *CreateImageRequest) SetDeleteAfterImageUpload(v string) *CreateImageRequest {
	s.DeleteAfterImageUpload = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotId(v string) *CreateImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequest) SetTargetOSSRegionId(v string) *CreateImageRequest {
	s.TargetOSSRegionId = &v
	return s
}

func (s *CreateImageRequest) SetWithDataDisks(v bool) *CreateImageRequest {
	s.WithDataDisks = &v
	return s
}

func (s *CreateImageRequest) Validate() error {
	return dara.Validate(s)
}
