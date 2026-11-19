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
	// The URL of the image. This value must be the Key field returned by GetUploadInfo, in the format of image/upload/<hash>. Passing an external image URL returns a 400 image.url.invalid error.
	//
	// example:
	//
	// image/upload/755c31906fc0b0b5007ecead9660c2ab
	ImgUrl *string `json:"ImgUrl,omitempty" xml:"ImgUrl,omitempty"`
	// The ID of the image library.
	//
	// example:
	//
	// custom_imglib_001
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The region ID.
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
