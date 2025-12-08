// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeExpressionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeExpressionResponseBodyData) *RecognizeExpressionResponseBody
	GetData() *RecognizeExpressionResponseBodyData
	SetRequestId(v string) *RecognizeExpressionResponseBody
	GetRequestId() *string
}

type RecognizeExpressionResponseBody struct {
	Data *RecognizeExpressionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E1C4C576-1799-4079-A934-15BC406A54EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeExpressionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeExpressionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBody) GetData() *RecognizeExpressionResponseBodyData {
	return s.Data
}

func (s *RecognizeExpressionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeExpressionResponseBody) SetData(v *RecognizeExpressionResponseBodyData) *RecognizeExpressionResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeExpressionResponseBody) SetRequestId(v string) *RecognizeExpressionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeExpressionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeExpressionResponseBodyData struct {
	Elements []*RecognizeExpressionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeExpressionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeExpressionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyData) GetElements() []*RecognizeExpressionResponseBodyDataElements {
	return s.Elements
}

func (s *RecognizeExpressionResponseBodyData) SetElements(v []*RecognizeExpressionResponseBodyDataElements) *RecognizeExpressionResponseBodyData {
	s.Elements = v
	return s
}

func (s *RecognizeExpressionResponseBodyData) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeExpressionResponseBodyDataElements struct {
	// example:
	//
	// surprise
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// 0.99651491641998291
	FaceProbability *float32                                                  `json:"FaceProbability,omitempty" xml:"FaceProbability,omitempty"`
	FaceRectangle   *RecognizeExpressionResponseBodyDataElementsFaceRectangle `json:"FaceRectangle,omitempty" xml:"FaceRectangle,omitempty" type:"Struct"`
}

func (s RecognizeExpressionResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s RecognizeExpressionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyDataElements) GetExpression() *string {
	return s.Expression
}

func (s *RecognizeExpressionResponseBodyDataElements) GetFaceProbability() *float32 {
	return s.FaceProbability
}

func (s *RecognizeExpressionResponseBodyDataElements) GetFaceRectangle() *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	return s.FaceRectangle
}

func (s *RecognizeExpressionResponseBodyDataElements) SetExpression(v string) *RecognizeExpressionResponseBodyDataElements {
	s.Expression = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElements) SetFaceProbability(v float32) *RecognizeExpressionResponseBodyDataElements {
	s.FaceProbability = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElements) SetFaceRectangle(v *RecognizeExpressionResponseBodyDataElementsFaceRectangle) *RecognizeExpressionResponseBodyDataElements {
	s.FaceRectangle = v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElements) Validate() error {
	if s.FaceRectangle != nil {
		if err := s.FaceRectangle.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeExpressionResponseBodyDataElementsFaceRectangle struct {
	// example:
	//
	// 174
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 196
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 41
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 121
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeExpressionResponseBodyDataElementsFaceRectangle) String() string {
	return dara.Prettify(s)
}

func (s RecognizeExpressionResponseBodyDataElementsFaceRectangle) GoString() string {
	return s.String()
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) GetHeight() *int32 {
	return s.Height
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) GetLeft() *int32 {
	return s.Left
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) GetTop() *int32 {
	return s.Top
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) GetWidth() *int32 {
	return s.Width
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetHeight(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Height = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetLeft(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Left = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetTop(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Top = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) SetWidth(v int32) *RecognizeExpressionResponseBodyDataElementsFaceRectangle {
	s.Width = &v
	return s
}

func (s *RecognizeExpressionResponseBodyDataElementsFaceRectangle) Validate() error {
	return dara.Validate(s)
}
