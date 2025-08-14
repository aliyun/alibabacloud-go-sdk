// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToWordJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNameExtension(v string) *SubmitConvertImageToWordJobShrinkRequest
	GetImageNameExtension() *string
	SetImageNamesShrink(v string) *SubmitConvertImageToWordJobShrinkRequest
	GetImageNamesShrink() *string
	SetImageUrlsShrink(v string) *SubmitConvertImageToWordJobShrinkRequest
	GetImageUrlsShrink() *string
	SetOssBucket(v string) *SubmitConvertImageToWordJobShrinkRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToWordJobShrinkRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToWordJobShrinkRequest struct {
	// example:
	//
	// jpg
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToWordJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToWordJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobShrinkRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToWordJobShrinkRequest) GetImageNamesShrink() *string {
	return s.ImageNamesShrink
}

func (s *SubmitConvertImageToWordJobShrinkRequest) GetImageUrlsShrink() *string {
	return s.ImageUrlsShrink
}

func (s *SubmitConvertImageToWordJobShrinkRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToWordJobShrinkRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetOssBucket(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToWordJobShrinkRequest) SetOssEndpoint(v string) *SubmitConvertImageToWordJobShrinkRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToWordJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
