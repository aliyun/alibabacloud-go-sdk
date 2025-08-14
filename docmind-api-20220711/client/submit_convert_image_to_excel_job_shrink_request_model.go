// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertImageToExcelJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceMergeExcel(v bool) *SubmitConvertImageToExcelJobShrinkRequest
	GetForceMergeExcel() *bool
	SetImageNameExtension(v string) *SubmitConvertImageToExcelJobShrinkRequest
	GetImageNameExtension() *string
	SetImageNamesShrink(v string) *SubmitConvertImageToExcelJobShrinkRequest
	GetImageNamesShrink() *string
	SetImageUrlsShrink(v string) *SubmitConvertImageToExcelJobShrinkRequest
	GetImageUrlsShrink() *string
	SetOssBucket(v string) *SubmitConvertImageToExcelJobShrinkRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertImageToExcelJobShrinkRequest
	GetOssEndpoint() *string
}

type SubmitConvertImageToExcelJobShrinkRequest struct {
	ForceMergeExcel *bool `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	// example:
	//
	// jpg
	ImageNameExtension *string `json:"ImageNameExtension,omitempty" xml:"ImageNameExtension,omitempty"`
	ImageNamesShrink   *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	ImageUrlsShrink    *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertImageToExcelJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertImageToExcelJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) GetForceMergeExcel() *bool {
	return s.ForceMergeExcel
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) GetImageNameExtension() *string {
	return s.ImageNameExtension
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) GetImageNamesShrink() *string {
	return s.ImageNamesShrink
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) GetImageUrlsShrink() *string {
	return s.ImageUrlsShrink
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetForceMergeExcel(v bool) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ForceMergeExcel = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetImageNameExtension(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ImageNameExtension = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetImageNamesShrink(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetImageUrlsShrink(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.ImageUrlsShrink = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetOssBucket(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) SetOssEndpoint(v string) *SubmitConvertImageToExcelJobShrinkRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertImageToExcelJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
