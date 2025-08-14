// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToExcelJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceMergeExcel(v bool) *SubmitConvertImageToExcelJobRequest
	GetForceMergeExcel() *bool
	SetImageNameExtension(v string) *SubmitConvertImageToExcelJobRequest
	GetImageNameExtension() *string
	SetImageNames(v []*string) *SubmitConvertImageToExcelJobRequest
	GetImageNames() []*string
	SetImageUrls(v []*string) *SubmitConvertImageToExcelJobRequest
	GetImageUrls() []*string
	SetOssBucket(v string) *SubmitConvertImageToExcelJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToExcelJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToExcelJobRequest struct {
	ForceMergeExcel *bool `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	// example:
	//
	// jpg
	ImageNameExtension *string   `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNames         []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageUrls          []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	OssBucket          *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToExcelJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToExcelJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobRequest) GetForceMergeExcel() *bool {
	return s.ForceMergeExcel
}

func (s *SubmitConvertImageToExcelJobRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToExcelJobRequest) GetImageNames() []*string {
	return s.ImageNames
}

func (s *SubmitConvertImageToExcelJobRequest) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *SubmitConvertImageToExcelJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToExcelJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToExcelJobRequest) SetForceMergeExcel(v bool) *SubmitConvertImageToExcelJobRequest {
	s.ForceMergeExcel = &v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetImageNameExtension(v string) *SubmitConvertImageToExcelJobRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetImageNames(v []*string) *SubmitConvertImageToExcelJobRequest {
	s.ImageNames = v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetImageUrls(v []*string) *SubmitConvertImageToExcelJobRequest {
	s.ImageUrls = v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetOssBucket(v string) *SubmitConvertImageToExcelJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) SetOssEndpoint(v string) *SubmitConvertImageToExcelJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToExcelJobRequest) Validate() error {
	return dara.Validate(s)
}
