// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrls(v []*string) *AnalyzeImageRequest
	GetImageUrls() []*string
	SetResponseFormatType(v string) *AnalyzeImageRequest
	GetResponseFormatType() *string
	SetResultTypes(v []*string) *AnalyzeImageRequest
	GetResultTypes() []*string
	SetStream(v bool) *AnalyzeImageRequest
	GetStream() *bool
}

type AnalyzeImageRequest struct {
	ImageUrls          []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	ResponseFormatType *string   `json:"responseFormatType,omitempty" xml:"responseFormatType,omitempty"`
	ResultTypes        []*string `json:"resultTypes,omitempty" xml:"resultTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s AnalyzeImageRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeImageRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeImageRequest) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *AnalyzeImageRequest) GetResponseFormatType() *string {
	return s.ResponseFormatType
}

func (s *AnalyzeImageRequest) GetResultTypes() []*string {
	return s.ResultTypes
}

func (s *AnalyzeImageRequest) GetStream() *bool {
	return s.Stream
}

func (s *AnalyzeImageRequest) SetImageUrls(v []*string) *AnalyzeImageRequest {
	s.ImageUrls = v
	return s
}

func (s *AnalyzeImageRequest) SetResponseFormatType(v string) *AnalyzeImageRequest {
	s.ResponseFormatType = &v
	return s
}

func (s *AnalyzeImageRequest) SetResultTypes(v []*string) *AnalyzeImageRequest {
	s.ResultTypes = v
	return s
}

func (s *AnalyzeImageRequest) SetStream(v bool) *AnalyzeImageRequest {
	s.Stream = &v
	return s
}

func (s *AnalyzeImageRequest) Validate() error {
	return dara.Validate(s)
}
