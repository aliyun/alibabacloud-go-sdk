// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomImageNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpdateCustomImageNameRequest
	GetImageId() *string
	SetImageName(v string) *UpdateCustomImageNameRequest
	GetImageName() *string
}

type UpdateCustomImageNameRequest struct {
	// The ID of the image.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// imagename
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s UpdateCustomImageNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomImageNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomImageNameRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateCustomImageNameRequest) GetImageName() *string {
	return s.ImageName
}

func (s *UpdateCustomImageNameRequest) SetImageId(v string) *UpdateCustomImageNameRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateCustomImageNameRequest) SetImageName(v string) *UpdateCustomImageNameRequest {
	s.ImageName = &v
	return s
}

func (s *UpdateCustomImageNameRequest) Validate() error {
	return dara.Validate(s)
}
