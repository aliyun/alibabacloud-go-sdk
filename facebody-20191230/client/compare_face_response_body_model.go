// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CompareFaceResponseBodyData) *CompareFaceResponseBody
	GetData() *CompareFaceResponseBodyData
	SetRequestId(v string) *CompareFaceResponseBody
	GetRequestId() *string
}

type CompareFaceResponseBody struct {
	Data      *CompareFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompareFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFaceResponseBody) GetData() *CompareFaceResponseBodyData {
	return s.Data
}

func (s *CompareFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareFaceResponseBody) SetData(v *CompareFaceResponseBodyData) *CompareFaceResponseBody {
	s.Data = v
	return s
}

func (s *CompareFaceResponseBody) SetRequestId(v string) *CompareFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareFaceResponseBodyData struct {
	Confidence     *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	IsMaskA        *int64   `json:"IsMaskA,omitempty" xml:"IsMaskA,omitempty"`
	IsMaskB        *int64   `json:"IsMaskB,omitempty" xml:"IsMaskB,omitempty"`
	LandmarksAList []*int64 `json:"LandmarksAList,omitempty" xml:"LandmarksAList,omitempty" type:"Repeated"`
	LandmarksBList []*int64 `json:"LandmarksBList,omitempty" xml:"LandmarksBList,omitempty" type:"Repeated"`
	// example:
	//
	// imageB quality score less threshold
	MessageTips *string `json:"MessageTips,omitempty" xml:"MessageTips,omitempty"`
	// example:
	//
	// 75.16
	QualityScoreA *float32 `json:"QualityScoreA,omitempty" xml:"QualityScoreA,omitempty"`
	// example:
	//
	// 75.20
	QualityScoreB *float32 `json:"QualityScoreB,omitempty" xml:"QualityScoreB,omitempty"`
	// 1
	RectAList []*int32 `json:"RectAList,omitempty" xml:"RectAList,omitempty" type:"Repeated"`
	// 1
	RectBList []*int32 `json:"RectBList,omitempty" xml:"RectBList,omitempty" type:"Repeated"`
	// 1
	Thresholds []*float32 `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Repeated"`
}

func (s CompareFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompareFaceResponseBodyData) GetConfidence() *float32 {
	return s.Confidence
}

func (s *CompareFaceResponseBodyData) GetIsMaskA() *int64 {
	return s.IsMaskA
}

func (s *CompareFaceResponseBodyData) GetIsMaskB() *int64 {
	return s.IsMaskB
}

func (s *CompareFaceResponseBodyData) GetLandmarksAList() []*int64 {
	return s.LandmarksAList
}

func (s *CompareFaceResponseBodyData) GetLandmarksBList() []*int64 {
	return s.LandmarksBList
}

func (s *CompareFaceResponseBodyData) GetMessageTips() *string {
	return s.MessageTips
}

func (s *CompareFaceResponseBodyData) GetQualityScoreA() *float32 {
	return s.QualityScoreA
}

func (s *CompareFaceResponseBodyData) GetQualityScoreB() *float32 {
	return s.QualityScoreB
}

func (s *CompareFaceResponseBodyData) GetRectAList() []*int32 {
	return s.RectAList
}

func (s *CompareFaceResponseBodyData) GetRectBList() []*int32 {
	return s.RectBList
}

func (s *CompareFaceResponseBodyData) GetThresholds() []*float32 {
	return s.Thresholds
}

func (s *CompareFaceResponseBodyData) SetConfidence(v float32) *CompareFaceResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetIsMaskA(v int64) *CompareFaceResponseBodyData {
	s.IsMaskA = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetIsMaskB(v int64) *CompareFaceResponseBodyData {
	s.IsMaskB = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetLandmarksAList(v []*int64) *CompareFaceResponseBodyData {
	s.LandmarksAList = v
	return s
}

func (s *CompareFaceResponseBodyData) SetLandmarksBList(v []*int64) *CompareFaceResponseBodyData {
	s.LandmarksBList = v
	return s
}

func (s *CompareFaceResponseBodyData) SetMessageTips(v string) *CompareFaceResponseBodyData {
	s.MessageTips = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetQualityScoreA(v float32) *CompareFaceResponseBodyData {
	s.QualityScoreA = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetQualityScoreB(v float32) *CompareFaceResponseBodyData {
	s.QualityScoreB = &v
	return s
}

func (s *CompareFaceResponseBodyData) SetRectAList(v []*int32) *CompareFaceResponseBodyData {
	s.RectAList = v
	return s
}

func (s *CompareFaceResponseBodyData) SetRectBList(v []*int32) *CompareFaceResponseBodyData {
	s.RectBList = v
	return s
}

func (s *CompareFaceResponseBodyData) SetThresholds(v []*float32) *CompareFaceResponseBodyData {
	s.Thresholds = v
	return s
}

func (s *CompareFaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
