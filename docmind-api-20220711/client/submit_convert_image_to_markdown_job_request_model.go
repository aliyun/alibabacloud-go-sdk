// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToMarkdownJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNameExtension(v string) *SubmitConvertImageToMarkdownJobRequest
	GetImageNameExtension() *string
	SetImageNames(v []*string) *SubmitConvertImageToMarkdownJobRequest
	GetImageNames() []*string
	SetImageUrls(v []*string) *SubmitConvertImageToMarkdownJobRequest
	GetImageUrls() []*string
	SetOssBucket(v string) *SubmitConvertImageToMarkdownJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToMarkdownJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToMarkdownJobRequest struct {
	// example:
	//
	// jpg
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	OssBucket          *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToMarkdownJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToMarkdownJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToMarkdownJobRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToMarkdownJobRequest) GetImageNames() []*string {
	return s.ImageNames
}

func (s *SubmitConvertImageToMarkdownJobRequest) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *SubmitConvertImageToMarkdownJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToMarkdownJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToMarkdownJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetImageNames(v []*string) *SubmitConvertImageToMarkdownJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToMarkdownJobRequest {
	s.ImageUrls = v
	return s
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetOssBucket(v string) *SubmitConvertImageToMarkdownJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobRequest) SetOssEndpoint(v string) *SubmitConvertImageToMarkdownJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToMarkdownJobRequest) Validate() error {
	return dara.Validate(s)
}
