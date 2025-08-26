// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssureDirection(v bool) *RecognizeTableRequest
	GetAssureDirection() *bool
	SetHasLine(v bool) *RecognizeTableRequest
	GetHasLine() *bool
	SetImageURL(v string) *RecognizeTableRequest
	GetImageURL() *string
	SetOutputFormat(v string) *RecognizeTableRequest
	GetOutputFormat() *string
	SetSkipDetection(v bool) *RecognizeTableRequest
	GetSkipDetection() *bool
	SetUseFinanceModel(v bool) *RecognizeTableRequest
	GetUseFinanceModel() *bool
}

type RecognizeTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	AssureDirection *bool `json:"AssureDirection,omitempty" xml:"AssureDirection,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	HasLine *bool `json:"HasLine,omitempty" xml:"HasLine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeTable/RecognizeTable4.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// json
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	SkipDetection *bool `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	UseFinanceModel *bool `json:"UseFinanceModel,omitempty" xml:"UseFinanceModel,omitempty"`
}

func (s RecognizeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableRequest) GetAssureDirection() *bool {
	return s.AssureDirection
}

func (s *RecognizeTableRequest) GetHasLine() *bool {
	return s.HasLine
}

func (s *RecognizeTableRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeTableRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RecognizeTableRequest) GetSkipDetection() *bool {
	return s.SkipDetection
}

func (s *RecognizeTableRequest) GetUseFinanceModel() *bool {
	return s.UseFinanceModel
}

func (s *RecognizeTableRequest) SetAssureDirection(v bool) *RecognizeTableRequest {
	s.AssureDirection = &v
	return s
}

func (s *RecognizeTableRequest) SetHasLine(v bool) *RecognizeTableRequest {
	s.HasLine = &v
	return s
}

func (s *RecognizeTableRequest) SetImageURL(v string) *RecognizeTableRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeTableRequest) SetOutputFormat(v string) *RecognizeTableRequest {
	s.OutputFormat = &v
	return s
}

func (s *RecognizeTableRequest) SetSkipDetection(v bool) *RecognizeTableRequest {
	s.SkipDetection = &v
	return s
}

func (s *RecognizeTableRequest) SetUseFinanceModel(v bool) *RecognizeTableRequest {
	s.UseFinanceModel = &v
	return s
}

func (s *RecognizeTableRequest) Validate() error {
	return dara.Validate(s)
}
