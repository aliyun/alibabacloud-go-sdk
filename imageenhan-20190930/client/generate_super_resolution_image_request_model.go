// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSuperResolutionImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *GenerateSuperResolutionImageRequest
	GetImageUrl() *string
	SetOutputFormat(v string) *GenerateSuperResolutionImageRequest
	GetOutputFormat() *string
	SetOutputQuality(v int32) *GenerateSuperResolutionImageRequest
	GetOutputQuality() *int32
	SetScale(v int32) *GenerateSuperResolutionImageRequest
	GetScale() *int32
	SetUserData(v string) *GenerateSuperResolutionImageRequest
	GetUserData() *string
}

type GenerateSuperResolutionImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/xxx/1025.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 95
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// example:
	//
	// jpg
	OutputQuality *int32 `json:"OutputQuality,omitempty" xml:"OutputQuality,omitempty"`
	// example:
	//
	// 2
	Scale    *int32  `json:"Scale,omitempty" xml:"Scale,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GenerateSuperResolutionImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateSuperResolutionImageRequest) GoString() string {
	return s.String()
}

func (s *GenerateSuperResolutionImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GenerateSuperResolutionImageRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *GenerateSuperResolutionImageRequest) GetOutputQuality() *int32 {
	return s.OutputQuality
}

func (s *GenerateSuperResolutionImageRequest) GetScale() *int32 {
	return s.Scale
}

func (s *GenerateSuperResolutionImageRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateSuperResolutionImageRequest) SetImageUrl(v string) *GenerateSuperResolutionImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *GenerateSuperResolutionImageRequest) SetOutputFormat(v string) *GenerateSuperResolutionImageRequest {
	s.OutputFormat = &v
	return s
}

func (s *GenerateSuperResolutionImageRequest) SetOutputQuality(v int32) *GenerateSuperResolutionImageRequest {
	s.OutputQuality = &v
	return s
}

func (s *GenerateSuperResolutionImageRequest) SetScale(v int32) *GenerateSuperResolutionImageRequest {
	s.Scale = &v
	return s
}

func (s *GenerateSuperResolutionImageRequest) SetUserData(v string) *GenerateSuperResolutionImageRequest {
	s.UserData = &v
	return s
}

func (s *GenerateSuperResolutionImageRequest) Validate() error {
	return dara.Validate(s)
}
