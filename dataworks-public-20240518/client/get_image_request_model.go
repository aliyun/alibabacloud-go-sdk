// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetImageRequest
	GetId() *string
	SetImageVersion(v string) *GetImageRequest
	GetImageVersion() *string
}

type GetImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
}

func (s GetImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) GetId() *string {
	return s.Id
}

func (s *GetImageRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *GetImageRequest) SetId(v string) *GetImageRequest {
	s.Id = &v
	return s
}

func (s *GetImageRequest) SetImageVersion(v string) *GetImageRequest {
	s.ImageVersion = &v
	return s
}

func (s *GetImageRequest) Validate() error {
	return dara.Validate(s)
}
