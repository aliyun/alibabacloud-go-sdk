// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpdateInstanceImageRequest
	GetImageId() *string
	SetInstanceIdList(v []*string) *UpdateInstanceImageRequest
	GetInstanceIdList() []*string
}

type UpdateInstanceImageRequest struct {
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId        *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s UpdateInstanceImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateInstanceImageRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *UpdateInstanceImageRequest) SetImageId(v string) *UpdateInstanceImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceImageRequest) SetInstanceIdList(v []*string) *UpdateInstanceImageRequest {
	s.InstanceIdList = v
	return s
}

func (s *UpdateInstanceImageRequest) Validate() error {
	return dara.Validate(s)
}
