// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectLivingFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectLivingFaceResponseBodyData) *DetectLivingFaceResponseBody
	GetData() *DetectLivingFaceResponseBodyData
	SetRequestId(v string) *DetectLivingFaceResponseBody
	GetRequestId() *string
}

type DetectLivingFaceResponseBody struct {
	Data *DetectLivingFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 2EEA0396-BD84-5729-B8BD-D60776FCEF2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectLivingFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBody) GetData() *DetectLivingFaceResponseBodyData {
	return s.Data
}

func (s *DetectLivingFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectLivingFaceResponseBody) SetData(v *DetectLivingFaceResponseBodyData) *DetectLivingFaceResponseBody {
	s.Data = v
	return s
}

func (s *DetectLivingFaceResponseBody) SetRequestId(v string) *DetectLivingFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectLivingFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectLivingFaceResponseBodyData struct {
	Elements []*DetectLivingFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectLivingFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyData) GetElements() []*DetectLivingFaceResponseBodyDataElements {
	return s.Elements
}

func (s *DetectLivingFaceResponseBodyData) SetElements(v []*DetectLivingFaceResponseBodyDataElements) *DetectLivingFaceResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectLivingFaceResponseBodyData) Validate() error {
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

type DetectLivingFaceResponseBodyDataElements struct {
	// example:
	//
	// 1
	FaceNumber *int64 `json:"FaceNumber,omitempty" xml:"FaceNumber,omitempty"`
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectLivingFace/DetectLivingFace4.jpg
	ImageURL *string                                            `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Results  []*DetectLivingFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// img1owc8WGskNm78OEAPJTZal-1****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElements) GetFaceNumber() *int64 {
	return s.FaceNumber
}

func (s *DetectLivingFaceResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectLivingFaceResponseBodyDataElements) GetResults() []*DetectLivingFaceResponseBodyDataElementsResults {
	return s.Results
}

func (s *DetectLivingFaceResponseBodyDataElements) GetTaskId() *string {
	return s.TaskId
}

func (s *DetectLivingFaceResponseBodyDataElements) SetFaceNumber(v int64) *DetectLivingFaceResponseBodyDataElements {
	s.FaceNumber = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetImageURL(v string) *DetectLivingFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetResults(v []*DetectLivingFaceResponseBodyDataElementsResults) *DetectLivingFaceResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) SetTaskId(v string) *DetectLivingFaceResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElements) Validate() error {
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

type DetectLivingFaceResponseBodyDataElementsResults struct {
	Frames []*DetectLivingFaceResponseBodyDataElementsResultsFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
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
	// 76.22
	Rate *float32                                             `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Rect *DetectLivingFaceResponseBodyDataElementsResultsRect `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Struct"`
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResults) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) GetFrames() []*DetectLivingFaceResponseBodyDataElementsResultsFrames {
	return s.Frames
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) GetLabel() *string {
	return s.Label
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) GetMessageTips() *string {
	return s.MessageTips
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) GetRate() *float32 {
	return s.Rate
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) GetRect() *DetectLivingFaceResponseBodyDataElementsResultsRect {
	return s.Rect
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetFrames(v []*DetectLivingFaceResponseBodyDataElementsResultsFrames) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Frames = v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetLabel(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetMessageTips(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.MessageTips = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetRate(v float32) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetRect(v *DetectLivingFaceResponseBodyDataElementsResultsRect) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Rect = v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) SetSuggestion(v string) *DetectLivingFaceResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResults) Validate() error {
	if s.Frames != nil {
		for _, item := range s.Frames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rect != nil {
		if err := s.Rect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectLivingFaceResponseBodyDataElementsResultsFrames struct {
	// example:
	//
	// 84.83
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// example:
	//
	// http://aligreen-shanghai.oss-cn-shanghai.aliyuncs.com/prod/hammal/26210da42/28118541_TB1urBOQFXXXXbMXFXXXXXXXXXX-1442-257.png?Expires=1582703593&OSSAccessKeyId=H4sp5QfNbuDg****&Signature=%2B8iUkb5YjomYR8ujV2c8wMAavs****&x-oss-process=image%2Fcrop%2Cx_0%2Cw_288%2Fauto-orient%2C0
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResultsFrames) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResultsFrames) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) GetRate() *float32 {
	return s.Rate
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) GetUrl() *string {
	return s.Url
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) SetRate(v float32) *DetectLivingFaceResponseBodyDataElementsResultsFrames {
	s.Rate = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) SetUrl(v string) *DetectLivingFaceResponseBodyDataElementsResultsFrames {
	s.Url = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsFrames) Validate() error {
	return dara.Validate(s)
}

type DetectLivingFaceResponseBodyDataElementsResultsRect struct {
	// example:
	//
	// 60
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 20
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 30
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 50
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectLivingFaceResponseBodyDataElementsResultsRect) String() string {
	return dara.Prettify(s)
}

func (s DetectLivingFaceResponseBodyDataElementsResultsRect) GoString() string {
	return s.String()
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) GetHeight() *int64 {
	return s.Height
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) GetLeft() *int64 {
	return s.Left
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) GetTop() *int64 {
	return s.Top
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) GetWidth() *int64 {
	return s.Width
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetHeight(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Height = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetLeft(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Left = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetTop(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Top = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) SetWidth(v int64) *DetectLivingFaceResponseBodyDataElementsResultsRect {
	s.Width = &v
	return s
}

func (s *DetectLivingFaceResponseBodyDataElementsResultsRect) Validate() error {
	return dara.Validate(s)
}
