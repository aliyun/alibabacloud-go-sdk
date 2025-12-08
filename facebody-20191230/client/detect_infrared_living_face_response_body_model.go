// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectInfraredLivingFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectInfraredLivingFaceResponseBodyData) *DetectInfraredLivingFaceResponseBody
	GetData() *DetectInfraredLivingFaceResponseBodyData
	SetRequestId(v string) *DetectInfraredLivingFaceResponseBody
	GetRequestId() *string
}

type DetectInfraredLivingFaceResponseBody struct {
	Data *DetectInfraredLivingFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 68DF6EC7-3B1D-11EE-9FA7-1122F1AE92DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectInfraredLivingFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceResponseBody) GetData() *DetectInfraredLivingFaceResponseBodyData {
	return s.Data
}

func (s *DetectInfraredLivingFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectInfraredLivingFaceResponseBody) SetData(v *DetectInfraredLivingFaceResponseBodyData) *DetectInfraredLivingFaceResponseBody {
	s.Data = v
	return s
}

func (s *DetectInfraredLivingFaceResponseBody) SetRequestId(v string) *DetectInfraredLivingFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectInfraredLivingFaceResponseBodyData struct {
	Elements []*DetectInfraredLivingFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectInfraredLivingFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceResponseBodyData) GetElements() []*DetectInfraredLivingFaceResponseBodyDataElements {
	return s.Elements
}

func (s *DetectInfraredLivingFaceResponseBodyData) SetElements(v []*DetectInfraredLivingFaceResponseBodyDataElements) *DetectInfraredLivingFaceResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyData) Validate() error {
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

type DetectInfraredLivingFaceResponseBodyDataElements struct {
	// example:
	//
	// 1
	FaceNumber *int64 `json:"FaceNumber,omitempty" xml:"FaceNumber,omitempty"`
	// example:
	//
	// http://viapi-demo.oss-cn-shanghai.aliyuncs.com/viapi-demo/images/SegmentCommonImage/segmentimage-src-hu.jpeg
	ImageURL *string                                                    `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Results  []*DetectInfraredLivingFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetectInfraredLivingFaceResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceResponseBodyDataElements) GetFaceNumber() *int64 {
	return s.FaceNumber
}

func (s *DetectInfraredLivingFaceResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectInfraredLivingFaceResponseBodyDataElements) GetResults() []*DetectInfraredLivingFaceResponseBodyDataElementsResults {
	return s.Results
}

func (s *DetectInfraredLivingFaceResponseBodyDataElements) SetFaceNumber(v int64) *DetectInfraredLivingFaceResponseBodyDataElements {
	s.FaceNumber = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElements) SetImageURL(v string) *DetectInfraredLivingFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElements) SetResults(v []*DetectInfraredLivingFaceResponseBodyDataElementsResults) *DetectInfraredLivingFaceResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElements) Validate() error {
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

type DetectInfraredLivingFaceResponseBodyDataElementsResults struct {
	// example:
	//
	// liveness
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// blurred, occluded or large angle face, please check.
	MessageTips *string `json:"MessageTips,omitempty" xml:"MessageTips,omitempty"`
	// example:
	//
	// 83.3848
	Rate *float32                                                     `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Rect *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DetectInfraredLivingFaceResponseBodyDataElementsResults) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) GetLabel() *string {
	return s.Label
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) GetMessageTips() *string {
	return s.MessageTips
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) GetRate() *float32 {
	return s.Rate
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) GetRect() *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect {
	return s.Rect
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) SetLabel(v string) *DetectInfraredLivingFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) SetMessageTips(v string) *DetectInfraredLivingFaceResponseBodyDataElementsResults {
	s.MessageTips = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) SetRate(v float32) *DetectInfraredLivingFaceResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) SetRect(v *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) *DetectInfraredLivingFaceResponseBodyDataElementsResults {
	s.Rect = v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) SetSuggestion(v string) *DetectInfraredLivingFaceResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResults) Validate() error {
	if s.Rect != nil {
		if err := s.Rect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectInfraredLivingFaceResponseBodyDataElementsResultsRect struct {
	// example:
	//
	// 20
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 60
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 50
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 30
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) String() string {
	return dara.Prettify(s)
}

func (s DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) GoString() string {
	return s.String()
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) GetHeight() *int64 {
	return s.Height
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) GetLeft() *int64 {
	return s.Left
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) GetTop() *int64 {
	return s.Top
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) GetWidth() *int64 {
	return s.Width
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) SetHeight(v int64) *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect {
	s.Height = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) SetLeft(v int64) *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect {
	s.Left = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) SetTop(v int64) *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect {
	s.Top = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) SetWidth(v int64) *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect {
	s.Width = &v
	return s
}

func (s *DetectInfraredLivingFaceResponseBodyDataElementsResultsRect) Validate() error {
	return dara.Validate(s)
}
