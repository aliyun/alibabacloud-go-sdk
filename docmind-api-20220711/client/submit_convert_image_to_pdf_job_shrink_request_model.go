// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToPdfJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNameExtension(v string) *SubmitConvertImageToPdfJobShrinkRequest
	GetImageNameExtension() *string
	SetImageNamesShrink(v string) *SubmitConvertImageToPdfJobShrinkRequest
	GetImageNamesShrink() *string
	SetImageUrlsShrink(v string) *SubmitConvertImageToPdfJobShrinkRequest
	GetImageUrlsShrink() *string
	SetOssBucket(v string) *SubmitConvertImageToPdfJobShrinkRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToPdfJobShrinkRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToPdfJobShrinkRequest struct {
	// example:
	//
	// JPG
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToPdfJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToPdfJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) GetImageNamesShrink() *string {
	return s.ImageNamesShrink
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) GetImageUrlsShrink() *string {
	return s.ImageUrlsShrink
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetOssBucket(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) SetOssEndpoint(v string) *SubmitConvertImageToPdfJobShrinkRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToPdfJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
