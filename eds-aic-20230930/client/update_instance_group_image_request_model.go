// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceGroupImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpdateInstanceGroupImageRequest
	GetImageId() *string
	SetInstanceGroupIds(v []*string) *UpdateInstanceGroupImageRequest
	GetInstanceGroupIds() []*string
}

type UpdateInstanceGroupImageRequest struct {
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-075cllfeuazh****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The IDs of the instance groups.
	//
	// This parameter is required.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
}

func (s UpdateInstanceGroupImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceGroupImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceGroupImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateInstanceGroupImageRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *UpdateInstanceGroupImageRequest) SetImageId(v string) *UpdateInstanceGroupImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceGroupImageRequest) SetInstanceGroupIds(v []*string) *UpdateInstanceGroupImageRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *UpdateInstanceGroupImageRequest) Validate() error {
	return dara.Validate(s)
}
