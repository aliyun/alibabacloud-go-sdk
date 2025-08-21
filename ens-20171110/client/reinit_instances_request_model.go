// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinitInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ReinitInstancesRequest
	GetImageId() *string
	SetInstanceIds(v []*string) *ReinitInstancesRequest
	GetInstanceIds() []*string
	SetPassword(v string) *ReinitInstancesRequest
	GetPassword() *string
}

type ReinitInstancesRequest struct {
	ImageId     *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	Password    *string   `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ReinitInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ReinitInstancesRequest) GoString() string {
	return s.String()
}

func (s *ReinitInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ReinitInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ReinitInstancesRequest) GetPassword() *string {
	return s.Password
}

func (s *ReinitInstancesRequest) SetImageId(v string) *ReinitInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *ReinitInstancesRequest) SetInstanceIds(v []*string) *ReinitInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ReinitInstancesRequest) SetPassword(v string) *ReinitInstancesRequest {
	s.Password = &v
	return s
}

func (s *ReinitInstancesRequest) Validate() error {
	return dara.Validate(s)
}
