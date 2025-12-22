// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeCharacterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeCharacterResponseBodyData) *RecognizeCharacterResponseBody
	GetData() *RecognizeCharacterResponseBodyData
	SetRequestId(v string) *RecognizeCharacterResponseBody
	GetRequestId() *string
}

type RecognizeCharacterResponseBody struct {
	Data *RecognizeCharacterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7A9BC7FE-2D42-57AF-93BC-09A229DD2F1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCharacterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBody) GetData() *RecognizeCharacterResponseBodyData {
	return s.Data
}

func (s *RecognizeCharacterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeCharacterResponseBody) SetData(v *RecognizeCharacterResponseBodyData) *RecognizeCharacterResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeCharacterResponseBody) SetRequestId(v string) *RecognizeCharacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeCharacterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeCharacterResponseBodyData struct {
	Results []*RecognizeCharacterResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RecognizeCharacterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyData) GetResults() []*RecognizeCharacterResponseBodyDataResults {
	return s.Results
}

func (s *RecognizeCharacterResponseBodyData) SetResults(v []*RecognizeCharacterResponseBodyDataResults) *RecognizeCharacterResponseBodyData {
	s.Results = v
	return s
}

func (s *RecognizeCharacterResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeCharacterResponseBodyDataResults struct {
	// example:
	//
	// 0.99
	Probability    *float32                                                 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	Text           *string                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
	TextRectangles *RecognizeCharacterResponseBodyDataResultsTextRectangles `json:"TextRectangles,omitempty" xml:"TextRectangles,omitempty" type:"Struct"`
}

func (s RecognizeCharacterResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyDataResults) GetProbability() *float32 {
	return s.Probability
}

func (s *RecognizeCharacterResponseBodyDataResults) GetText() *string {
	return s.Text
}

func (s *RecognizeCharacterResponseBodyDataResults) GetTextRectangles() *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	return s.TextRectangles
}

func (s *RecognizeCharacterResponseBodyDataResults) SetProbability(v float32) *RecognizeCharacterResponseBodyDataResults {
	s.Probability = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResults) SetText(v string) *RecognizeCharacterResponseBodyDataResults {
	s.Text = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResults) SetTextRectangles(v *RecognizeCharacterResponseBodyDataResultsTextRectangles) *RecognizeCharacterResponseBodyDataResults {
	s.TextRectangles = v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResults) Validate() error {
	if s.TextRectangles != nil {
		if err := s.TextRectangles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeCharacterResponseBodyDataResultsTextRectangles struct {
	// example:
	//
	// -65
	Angle *int32 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// 409
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 511
	Left *int32                                                        `json:"Left,omitempty" xml:"Left,omitempty"`
	Pos  []*RecognizeCharacterResponseBodyDataResultsTextRectanglesPos `json:"Pos,omitempty" xml:"Pos,omitempty" type:"Repeated"`
	// example:
	//
	// 150
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 77
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectangles) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectangles) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) GetAngle() *int32 {
	return s.Angle
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) GetHeight() *int32 {
	return s.Height
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) GetLeft() *int32 {
	return s.Left
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) GetPos() []*RecognizeCharacterResponseBodyDataResultsTextRectanglesPos {
	return s.Pos
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) GetTop() *int32 {
	return s.Top
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) GetWidth() *int32 {
	return s.Width
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetAngle(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Angle = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetHeight(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Height = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetLeft(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Left = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetPos(v []*RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Pos = v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetTop(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Top = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetWidth(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Width = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) Validate() error {
	if s.Pos != nil {
		for _, item := range s.Pos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeCharacterResponseBodyDataResultsTextRectanglesPos struct {
	X *int32 `json:"x,omitempty" xml:"x,omitempty"`
	Y *int32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) GetX() *int32 {
	return s.X
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) GetY() *int32 {
	return s.Y
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) SetX(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectanglesPos {
	s.X = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) SetY(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectanglesPos {
	s.Y = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectanglesPos) Validate() error {
	return dara.Validate(s)
}
