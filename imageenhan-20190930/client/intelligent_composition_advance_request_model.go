// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iIntelligentCompositionAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *IntelligentCompositionAdvanceRequest
	GetImageURLObject() io.Reader
	SetNumBoxes(v int32) *IntelligentCompositionAdvanceRequest
	GetNumBoxes() *int32
}

type IntelligentCompositionAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/IntelligentComposition/IntelligentComposition3.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// 5
	NumBoxes *int32 `json:"NumBoxes,omitempty" xml:"NumBoxes,omitempty"`
}

func (s IntelligentCompositionAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s IntelligentCompositionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *IntelligentCompositionAdvanceRequest) GetNumBoxes() *int32 {
	return s.NumBoxes
}

func (s *IntelligentCompositionAdvanceRequest) SetImageURLObject(v io.Reader) *IntelligentCompositionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *IntelligentCompositionAdvanceRequest) SetNumBoxes(v int32) *IntelligentCompositionAdvanceRequest {
	s.NumBoxes = &v
	return s
}

func (s *IntelligentCompositionAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
