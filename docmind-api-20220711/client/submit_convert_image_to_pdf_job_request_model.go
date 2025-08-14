// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToPdfJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNameExtension(v string) *SubmitConvertImageToPdfJobRequest
	GetImageNameExtension() *string
	SetImageNames(v []*string) *SubmitConvertImageToPdfJobRequest
	GetImageNames() []*string
	SetImageUrls(v []*string) *SubmitConvertImageToPdfJobRequest
	GetImageUrls() []*string
	SetOssBucket(v string) *SubmitConvertImageToPdfJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToPdfJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToPdfJobRequest struct {
	// example:
	//
	// JPG
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	OssBucket          *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToPdfJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToPdfJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToPdfJobRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToPdfJobRequest) GetImageNames() []*string {
	return s.ImageNames
}

func (s *SubmitConvertImageToPdfJobRequest) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *SubmitConvertImageToPdfJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToPdfJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToPdfJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToPdfJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToPdfJobRequest) SetImageNames(v []*string) *SubmitConvertImageToPdfJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToPdfJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToPdfJobRequest {
	s.ImageUrls = v
	return s
}

func (s *SubmitConvertImageToPdfJobRequest) SetOssBucket(v string) *SubmitConvertImageToPdfJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToPdfJobRequest) SetOssEndpoint(v string) *SubmitConvertImageToPdfJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToPdfJobRequest) Validate() error {
	return dara.Validate(s)
}
