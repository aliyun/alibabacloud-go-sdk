// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iMakeSuperResolutionImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *MakeSuperResolutionImageAdvanceRequest
	GetMode() *string
	SetOutputFormat(v string) *MakeSuperResolutionImageAdvanceRequest
	GetOutputFormat() *string
	SetOutputQuality(v int64) *MakeSuperResolutionImageAdvanceRequest
	GetOutputQuality() *int64
	SetUpscaleFactor(v int64) *MakeSuperResolutionImageAdvanceRequest
	GetUpscaleFactor() *int64
	SetUrlObject(v io.Reader) *MakeSuperResolutionImageAdvanceRequest
	GetUrlObject() io.Reader
}

type MakeSuperResolutionImageAdvanceRequest struct {
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
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s MakeSuperResolutionImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MakeSuperResolutionImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *MakeSuperResolutionImageAdvanceRequest) GetMode() *string {
	return s.Mode
}

func (s *MakeSuperResolutionImageAdvanceRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *MakeSuperResolutionImageAdvanceRequest) GetOutputQuality() *int64 {
	return s.OutputQuality
}

func (s *MakeSuperResolutionImageAdvanceRequest) GetUpscaleFactor() *int64 {
	return s.UpscaleFactor
}

func (s *MakeSuperResolutionImageAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetMode(v string) *MakeSuperResolutionImageAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetOutputFormat(v string) *MakeSuperResolutionImageAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetOutputQuality(v int64) *MakeSuperResolutionImageAdvanceRequest {
	s.OutputQuality = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetUpscaleFactor(v int64) *MakeSuperResolutionImageAdvanceRequest {
	s.UpscaleFactor = &v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) SetUrlObject(v io.Reader) *MakeSuperResolutionImageAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *MakeSuperResolutionImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
