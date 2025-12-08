// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeepfakeFaceResponseBodyData) *DeepfakeFaceResponseBody
	GetData() *DeepfakeFaceResponseBodyData
	SetRequestId(v string) *DeepfakeFaceResponseBody
	GetRequestId() *string
}

type DeepfakeFaceResponseBody struct {
	Data *DeepfakeFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 8E6F65D5-62A1-1E5B-BC0B-00508034AC92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeepfakeFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceResponseBody) GetData() *DeepfakeFaceResponseBodyData {
	return s.Data
}

func (s *DeepfakeFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeepfakeFaceResponseBody) SetData(v *DeepfakeFaceResponseBodyData) *DeepfakeFaceResponseBody {
	s.Data = v
	return s
}

func (s *DeepfakeFaceResponseBody) SetRequestId(v string) *DeepfakeFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeepfakeFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeepfakeFaceResponseBodyData struct {
	Elements []*DeepfakeFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DeepfakeFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceResponseBodyData) GetElements() []*DeepfakeFaceResponseBodyDataElements {
	return s.Elements
}

func (s *DeepfakeFaceResponseBodyData) SetElements(v []*DeepfakeFaceResponseBodyDataElements) *DeepfakeFaceResponseBodyData {
	s.Elements = v
	return s
}

func (s *DeepfakeFaceResponseBodyData) Validate() error {
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

type DeepfakeFaceResponseBodyDataElements struct {
	// example:
	//
	// 1
	FaceNumber *int64 `json:"FaceNumber,omitempty" xml:"FaceNumber,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DeepfakeFace/DeepfakeFace1.jpg
	ImageURL *string                                        `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Results  []*DeepfakeFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 8E6F65D5-62A1-1E5B-BC0B-00508034AC92
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeepfakeFaceResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceResponseBodyDataElements) GetFaceNumber() *int64 {
	return s.FaceNumber
}

func (s *DeepfakeFaceResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *DeepfakeFaceResponseBodyDataElements) GetResults() []*DeepfakeFaceResponseBodyDataElementsResults {
	return s.Results
}

func (s *DeepfakeFaceResponseBodyDataElements) GetTaskId() *string {
	return s.TaskId
}

func (s *DeepfakeFaceResponseBodyDataElements) SetFaceNumber(v int64) *DeepfakeFaceResponseBodyDataElements {
	s.FaceNumber = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElements) SetImageURL(v string) *DeepfakeFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElements) SetResults(v []*DeepfakeFaceResponseBodyDataElementsResults) *DeepfakeFaceResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElements) SetTaskId(v string) *DeepfakeFaceResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElements) Validate() error {
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

type DeepfakeFaceResponseBodyDataElementsResults struct {
	// example:
	//
	// 36.6455
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// normalface
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// blurred, occluded or large angle face, please check.
	MessageTips *string                                          `json:"MessageTips,omitempty" xml:"MessageTips,omitempty"`
	Rect        *DeepfakeFaceResponseBodyDataElementsResultsRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
}

func (s DeepfakeFaceResponseBodyDataElementsResults) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) GetConfidence() *float32 {
	return s.Confidence
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) GetLabel() *string {
	return s.Label
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) GetMessageTips() *string {
	return s.MessageTips
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) GetRect() *DeepfakeFaceResponseBodyDataElementsResultsRect {
	return s.Rect
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) SetConfidence(v float32) *DeepfakeFaceResponseBodyDataElementsResults {
	s.Confidence = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) SetLabel(v string) *DeepfakeFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) SetMessageTips(v string) *DeepfakeFaceResponseBodyDataElementsResults {
	s.MessageTips = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) SetRect(v *DeepfakeFaceResponseBodyDataElementsResultsRect) *DeepfakeFaceResponseBodyDataElementsResults {
	s.Rect = v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResults) Validate() error {
	if s.Rect != nil {
		if err := s.Rect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeepfakeFaceResponseBodyDataElementsResultsRect struct {
	// example:
	//
	// 284
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 373
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 111
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 207
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DeepfakeFaceResponseBodyDataElementsResultsRect) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeFaceResponseBodyDataElementsResultsRect) GoString() string {
	return s.String()
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) GetHeight() *int64 {
	return s.Height
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) GetLeft() *int64 {
	return s.Left
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) GetTop() *int64 {
	return s.Top
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) GetWidth() *int64 {
	return s.Width
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) SetHeight(v int64) *DeepfakeFaceResponseBodyDataElementsResultsRect {
	s.Height = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) SetLeft(v int64) *DeepfakeFaceResponseBodyDataElementsResultsRect {
	s.Left = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) SetTop(v int64) *DeepfakeFaceResponseBodyDataElementsResultsRect {
	s.Top = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) SetWidth(v int64) *DeepfakeFaceResponseBodyDataElementsResultsRect {
	s.Width = &v
	return s
}

func (s *DeepfakeFaceResponseBodyDataElementsResultsRect) Validate() error {
	return dara.Validate(s)
}
