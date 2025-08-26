// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeTableAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssureDirection(v bool) *RecognizeTableAdvanceRequest
	GetAssureDirection() *bool
	SetHasLine(v bool) *RecognizeTableAdvanceRequest
	GetHasLine() *bool
	SetImageURLObject(v io.Reader) *RecognizeTableAdvanceRequest
	GetImageURLObject() io.Reader
	SetOutputFormat(v string) *RecognizeTableAdvanceRequest
	GetOutputFormat() *string
	SetSkipDetection(v bool) *RecognizeTableAdvanceRequest
	GetSkipDetection() *bool
	SetUseFinanceModel(v bool) *RecognizeTableAdvanceRequest
	GetUseFinanceModel() *bool
}

type RecognizeTableAdvanceRequest struct {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s RecognizeTableAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableAdvanceRequest) GetAssureDirection() *bool {
	return s.AssureDirection
}

func (s *RecognizeTableAdvanceRequest) GetHasLine() *bool {
	return s.HasLine
}

func (s *RecognizeTableAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeTableAdvanceRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *RecognizeTableAdvanceRequest) GetSkipDetection() *bool {
	return s.SkipDetection
}

func (s *RecognizeTableAdvanceRequest) GetUseFinanceModel() *bool {
	return s.UseFinanceModel
}

func (s *RecognizeTableAdvanceRequest) SetAssureDirection(v bool) *RecognizeTableAdvanceRequest {
	s.AssureDirection = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetHasLine(v bool) *RecognizeTableAdvanceRequest {
	s.HasLine = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTableAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetOutputFormat(v string) *RecognizeTableAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetSkipDetection(v bool) *RecognizeTableAdvanceRequest {
	s.SkipDetection = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetUseFinanceModel(v bool) *RecognizeTableAdvanceRequest {
	s.UseFinanceModel = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
