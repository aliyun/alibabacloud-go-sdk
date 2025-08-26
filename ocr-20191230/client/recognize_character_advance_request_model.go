// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeCharacterAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *RecognizeCharacterAdvanceRequest
	GetImageURLObject() io.Reader
	SetMinHeight(v int32) *RecognizeCharacterAdvanceRequest
	GetMinHeight() *int32
	SetOutputProbability(v bool) *RecognizeCharacterAdvanceRequest
	GetOutputProbability() *bool
}

type RecognizeCharacterAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeCharacter/RecognizeCharacter5.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MinHeight *int32 `json:"MinHeight,omitempty" xml:"MinHeight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	OutputProbability *bool `json:"OutputProbability,omitempty" xml:"OutputProbability,omitempty"`
}

func (s RecognizeCharacterAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RecognizeCharacterAdvanceRequest) GetMinHeight() *int32 {
	return s.MinHeight
}

func (s *RecognizeCharacterAdvanceRequest) GetOutputProbability() *bool {
	return s.OutputProbability
}

func (s *RecognizeCharacterAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeCharacterAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeCharacterAdvanceRequest) SetMinHeight(v int32) *RecognizeCharacterAdvanceRequest {
	s.MinHeight = &v
	return s
}

func (s *RecognizeCharacterAdvanceRequest) SetOutputProbability(v bool) *RecognizeCharacterAdvanceRequest {
	s.OutputProbability = &v
	return s
}

func (s *RecognizeCharacterAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
