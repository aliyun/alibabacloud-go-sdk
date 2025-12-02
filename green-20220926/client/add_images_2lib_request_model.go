// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImages2LibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImgUrl(v string) *AddImages2LibRequest
	GetImgUrl() *string
	SetLibId(v string) *AddImages2LibRequest
	GetLibId() *string
	SetRegionId(v string) *AddImages2LibRequest
	GetRegionId() *string
}

type AddImages2LibRequest struct {
	// URL of the image to be uploaded.
	//
	// example:
	//
	// upload/ea7a98f9-f8bd-4905-a79b-963c9da419c5.jpg
	ImgUrl *string `json:"ImgUrl,omitempty" xml:"ImgUrl,omitempty"`
	// The ID of image library.
	//
	// example:
	//
	// xxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddImages2LibRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImages2LibRequest) GoString() string {
	return s.String()
}

func (s *AddImages2LibRequest) GetImgUrl() *string {
	return s.ImgUrl
}

func (s *AddImages2LibRequest) GetLibId() *string {
	return s.LibId
}

func (s *AddImages2LibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddImages2LibRequest) SetImgUrl(v string) *AddImages2LibRequest {
	s.ImgUrl = &v
	return s
}

func (s *AddImages2LibRequest) SetLibId(v string) *AddImages2LibRequest {
	s.LibId = &v
	return s
}

func (s *AddImages2LibRequest) SetRegionId(v string) *AddImages2LibRequest {
	s.RegionId = &v
	return s
}

func (s *AddImages2LibRequest) Validate() error {
	return dara.Validate(s)
}
