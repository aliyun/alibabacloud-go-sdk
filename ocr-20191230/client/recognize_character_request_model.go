// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeCharacterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *RecognizeCharacterRequest
	GetImageURL() *string
	SetMinHeight(v int32) *RecognizeCharacterRequest
	GetMinHeight() *int32
	SetOutputProbability(v bool) *RecognizeCharacterRequest
	GetOutputProbability() *bool
}

type RecognizeCharacterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/ocr/RecognizeCharacter/RecognizeCharacter5.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s RecognizeCharacterRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizeCharacterRequest) GetMinHeight() *int32 {
	return s.MinHeight
}

func (s *RecognizeCharacterRequest) GetOutputProbability() *bool {
	return s.OutputProbability
}

func (s *RecognizeCharacterRequest) SetImageURL(v string) *RecognizeCharacterRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeCharacterRequest) SetMinHeight(v int32) *RecognizeCharacterRequest {
	s.MinHeight = &v
	return s
}

func (s *RecognizeCharacterRequest) SetOutputProbability(v bool) *RecognizeCharacterRequest {
	s.OutputProbability = &v
	return s
}

func (s *RecognizeCharacterRequest) Validate() error {
	return dara.Validate(s)
}
