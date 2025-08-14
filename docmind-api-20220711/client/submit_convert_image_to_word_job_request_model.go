// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToWordJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageNameExtension(v string) *SubmitConvertImageToWordJobRequest
	GetImageNameExtension() *string
	SetImageNames(v []*string) *SubmitConvertImageToWordJobRequest
	GetImageNames() []*string
	SetImageUrls(v []*string) *SubmitConvertImageToWordJobRequest
	GetImageUrls() []*string
	SetOssBucket(v string) *SubmitConvertImageToWordJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToWordJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToWordJobRequest struct {
	// example:
	//
	// jpg
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	OssBucket          *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToWordJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToWordJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToWordJobRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToWordJobRequest) GetImageNames() []*string {
	return s.ImageNames
}

func (s *SubmitConvertImageToWordJobRequest) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *SubmitConvertImageToWordJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToWordJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToWordJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToWordJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToWordJobRequest) SetImageNames(v []*string) *SubmitConvertImageToWordJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToWordJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToWordJobRequest {
	s.ImageUrls = v
	return s
}

func (s *SubmitConvertImageToWordJobRequest) SetOssBucket(v string) *SubmitConvertImageToWordJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToWordJobRequest) SetOssEndpoint(v string) *SubmitConvertImageToWordJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToWordJobRequest) Validate() error {
	return dara.Validate(s)
}
