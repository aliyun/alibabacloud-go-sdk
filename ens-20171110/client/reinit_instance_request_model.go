// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinitInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ReinitInstanceRequest
	GetImageId() *string
	SetInstanceId(v string) *ReinitInstanceRequest
	GetInstanceId() *string
	SetPassword(v string) *ReinitInstanceRequest
	GetPassword() *string
}

type ReinitInstanceRequest struct {
	// The ID of the image file that is used to reset the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-5wn1dhz5syoo9b48f440ntvad
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// i-5vn4n3y4laeb2ii9zxxltlvzi
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password of the instance.
	//
	// It must be 8 to 30 characters in length. It must include at least three of the following characters types: uppercase letters, lowercase letters, digits, and special characters. The following special character are supported: `()\\"~! @#$%^&*-_+={}[]:;\\"<>,.?/`
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ReinitInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReinitInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReinitInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ReinitInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReinitInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *ReinitInstanceRequest) SetImageId(v string) *ReinitInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *ReinitInstanceRequest) SetInstanceId(v string) *ReinitInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReinitInstanceRequest) SetPassword(v string) *ReinitInstanceRequest {
	s.Password = &v
	return s
}

func (s *ReinitInstanceRequest) Validate() error {
	return dara.Validate(s)
}
