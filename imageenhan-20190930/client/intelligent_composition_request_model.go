// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntelligentCompositionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *IntelligentCompositionRequest
	GetImageURL() *string
	SetNumBoxes(v int32) *IntelligentCompositionRequest
	GetNumBoxes() *int32
}

type IntelligentCompositionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/IntelligentComposition/IntelligentComposition3.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 5
	NumBoxes *int32 `json:"NumBoxes,omitempty" xml:"NumBoxes,omitempty"`
}

func (s IntelligentCompositionRequest) String() string {
	return dara.Prettify(s)
}

func (s IntelligentCompositionRequest) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *IntelligentCompositionRequest) GetNumBoxes() *int32 {
	return s.NumBoxes
}

func (s *IntelligentCompositionRequest) SetImageURL(v string) *IntelligentCompositionRequest {
	s.ImageURL = &v
	return s
}

func (s *IntelligentCompositionRequest) SetNumBoxes(v int32) *IntelligentCompositionRequest {
	s.NumBoxes = &v
	return s
}

func (s *IntelligentCompositionRequest) Validate() error {
	return dara.Validate(s)
}
