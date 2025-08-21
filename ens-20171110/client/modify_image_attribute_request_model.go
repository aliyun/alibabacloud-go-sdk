// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ModifyImageAttributeRequest
	GetImageId() *string
	SetImageName(v string) *ModifyImageAttributeRequest
	GetImageName() *string
}

type ModifyImageAttributeRequest struct {
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-5t4xwkfkbs0uxv0kymdb6uip7
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// Image Name
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s ModifyImageAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageAttributeRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageAttributeRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ModifyImageAttributeRequest) SetImageId(v string) *ModifyImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageAttributeRequest) SetImageName(v string) *ModifyImageAttributeRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyImageAttributeRequest) Validate() error {
	return dara.Validate(s)
}
