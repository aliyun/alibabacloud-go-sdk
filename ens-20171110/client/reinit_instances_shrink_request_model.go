// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinitInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ReinitInstancesShrinkRequest
	GetImageId() *string
	SetInstanceIdsShrink(v string) *ReinitInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetPassword(v string) *ReinitInstancesShrinkRequest
	GetPassword() *string
}

type ReinitInstancesShrinkRequest struct {
	ImageId           *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ReinitInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReinitInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReinitInstancesShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ReinitInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *ReinitInstancesShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *ReinitInstancesShrinkRequest) SetImageId(v string) *ReinitInstancesShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *ReinitInstancesShrinkRequest) SetInstanceIdsShrink(v string) *ReinitInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *ReinitInstancesShrinkRequest) SetPassword(v string) *ReinitInstancesShrinkRequest {
	s.Password = &v
	return s
}

func (s *ReinitInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
