// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeSuperResolutionImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *MakeSuperResolutionImageRequest
	GetMode() *string
	SetOutputFormat(v string) *MakeSuperResolutionImageRequest
	GetOutputFormat() *string
	SetOutputQuality(v int64) *MakeSuperResolutionImageRequest
	GetOutputQuality() *int64
	SetUpscaleFactor(v int64) *MakeSuperResolutionImageRequest
	GetUpscaleFactor() *int64
	SetUrl(v string) *MakeSuperResolutionImageRequest
	GetUrl() *string
}

type MakeSuperResolutionImageRequest struct {
	// example:
	//
	// base
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// png
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// example:
	//
	// 95
	OutputQuality *int64 `json:"OutputQuality,omitempty" xml:"OutputQuality,omitempty"`
	// example:
	//
	// 2
	UpscaleFactor *int64 `json:"UpscaleFactor,omitempty" xml:"UpscaleFactor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/MakeSuperResolutionImage/MakeSuperResolutionImage5.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageRequest) String() string {
	return dara.Prettify(s)
}

func (s MakeSuperResolutionImageRequest) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageRequest) GetMode() *string {
	return s.Mode
}

func (s *MakeSuperResolutionImageRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *MakeSuperResolutionImageRequest) GetOutputQuality() *int64 {
	return s.OutputQuality
}

func (s *MakeSuperResolutionImageRequest) GetUpscaleFactor() *int64 {
	return s.UpscaleFactor
}

func (s *MakeSuperResolutionImageRequest) GetUrl() *string {
	return s.Url
}

func (s *MakeSuperResolutionImageRequest) SetMode(v string) *MakeSuperResolutionImageRequest {
	s.Mode = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetOutputFormat(v string) *MakeSuperResolutionImageRequest {
	s.OutputFormat = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetOutputQuality(v int64) *MakeSuperResolutionImageRequest {
	s.OutputQuality = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetUpscaleFactor(v int64) *MakeSuperResolutionImageRequest {
	s.UpscaleFactor = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) SetUrl(v string) *MakeSuperResolutionImageRequest {
	s.Url = &v
	return s
}

func (s *MakeSuperResolutionImageRequest) Validate() error {
	return dara.Validate(s)
}
