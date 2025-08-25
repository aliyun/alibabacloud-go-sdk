// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateSuperResolutionImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrlObject(v io.Reader) *GenerateSuperResolutionImageAdvanceRequest
	GetImageUrlObject() io.Reader
	SetOutputFormat(v string) *GenerateSuperResolutionImageAdvanceRequest
	GetOutputFormat() *string
	SetOutputQuality(v int32) *GenerateSuperResolutionImageAdvanceRequest
	GetOutputQuality() *int32
	SetScale(v int32) *GenerateSuperResolutionImageAdvanceRequest
	GetScale() *int32
	SetUserData(v string) *GenerateSuperResolutionImageAdvanceRequest
	GetUserData() *string
}

type GenerateSuperResolutionImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test/xxx/1025.jpg
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s GenerateSuperResolutionImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateSuperResolutionImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateSuperResolutionImageAdvanceRequest) GetImageUrlObject() io.Reader {
	return s.ImageUrlObject
}

func (s *GenerateSuperResolutionImageAdvanceRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *GenerateSuperResolutionImageAdvanceRequest) GetOutputQuality() *int32 {
	return s.OutputQuality
}

func (s *GenerateSuperResolutionImageAdvanceRequest) GetScale() *int32 {
	return s.Scale
}

func (s *GenerateSuperResolutionImageAdvanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateSuperResolutionImageAdvanceRequest) SetImageUrlObject(v io.Reader) *GenerateSuperResolutionImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *GenerateSuperResolutionImageAdvanceRequest) SetOutputFormat(v string) *GenerateSuperResolutionImageAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *GenerateSuperResolutionImageAdvanceRequest) SetOutputQuality(v int32) *GenerateSuperResolutionImageAdvanceRequest {
	s.OutputQuality = &v
	return s
}

func (s *GenerateSuperResolutionImageAdvanceRequest) SetScale(v int32) *GenerateSuperResolutionImageAdvanceRequest {
	s.Scale = &v
	return s
}

func (s *GenerateSuperResolutionImageAdvanceRequest) SetUserData(v string) *GenerateSuperResolutionImageAdvanceRequest {
	s.UserData = &v
	return s
}

func (s *GenerateSuperResolutionImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
