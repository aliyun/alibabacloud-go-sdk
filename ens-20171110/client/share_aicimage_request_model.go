// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareAICImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ShareAICImageRequest
	GetImageId() *string
	SetUsers(v []*string) *ShareAICImageRequest
	GetUsers() []*string
}

type ShareAICImageRequest struct {
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// mykey
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The user groups.
	//
	// This parameter is required.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ShareAICImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ShareAICImageRequest) GoString() string {
	return s.String()
}

func (s *ShareAICImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ShareAICImageRequest) GetUsers() []*string {
	return s.Users
}

func (s *ShareAICImageRequest) SetImageId(v string) *ShareAICImageRequest {
	s.ImageId = &v
	return s
}

func (s *ShareAICImageRequest) SetUsers(v []*string) *ShareAICImageRequest {
	s.Users = v
	return s
}

func (s *ShareAICImageRequest) Validate() error {
	return dara.Validate(s)
}
