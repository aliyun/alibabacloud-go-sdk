// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToMarkdownJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNameExtension(v string) *SubmitConvertImageToMarkdownJobShrinkRequest
	GetImageNameExtension() *string
	SetImageNamesShrink(v string) *SubmitConvertImageToMarkdownJobShrinkRequest
	GetImageNamesShrink() *string
	SetImageUrlsShrink(v string) *SubmitConvertImageToMarkdownJobShrinkRequest
	GetImageUrlsShrink() *string
	SetOssBucket(v string) *SubmitConvertImageToMarkdownJobShrinkRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToMarkdownJobShrinkRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToMarkdownJobShrinkRequest struct {
	// example:
	//
	// jpg
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) GetImageNamesShrink() *string {
	return s.ImageNamesShrink
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) GetImageUrlsShrink() *string {
	return s.ImageUrlsShrink
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetOssBucket(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) SetOssEndpoint(v string) *SubmitConvertImageToMarkdownJobShrinkRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
