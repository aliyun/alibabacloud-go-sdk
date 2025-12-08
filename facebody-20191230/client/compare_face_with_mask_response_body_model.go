// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceWithMaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CompareFaceWithMaskResponseBodyData) *CompareFaceWithMaskResponseBody
	GetData() *CompareFaceWithMaskResponseBodyData
	SetRequestId(v string) *CompareFaceWithMaskResponseBody
	GetRequestId() *string
}

type CompareFaceWithMaskResponseBody struct {
	Data *CompareFaceWithMaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C6499918-E932-41B3-96F5-A18F50D262DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CompareFaceWithMaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceWithMaskResponseBody) GoString() string {
	return s.String()
}

func (s *CompareFaceWithMaskResponseBody) GetData() *CompareFaceWithMaskResponseBodyData {
	return s.Data
}

func (s *CompareFaceWithMaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareFaceWithMaskResponseBody) SetData(v *CompareFaceWithMaskResponseBodyData) *CompareFaceWithMaskResponseBody {
	s.Data = v
	return s
}

func (s *CompareFaceWithMaskResponseBody) SetRequestId(v string) *CompareFaceWithMaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareFaceWithMaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareFaceWithMaskResponseBodyData struct {
	Confidence     *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	IsMaskA        *int64   `json:"IsMaskA,omitempty" xml:"IsMaskA,omitempty"`
	IsMaskB        *int64   `json:"IsMaskB,omitempty" xml:"IsMaskB,omitempty"`
	LandmarksAList []*int64 `json:"LandmarksAList,omitempty" xml:"LandmarksAList,omitempty" type:"Repeated"`
	LandmarksBList []*int64 `json:"LandmarksBList,omitempty" xml:"LandmarksBList,omitempty" type:"Repeated"`
	// example:
	//
	// imageB quality score less threshold
	MessageTips   *string  `json:"MessageTips,omitempty" xml:"MessageTips,omitempty"`
	QualityScoreA *float32 `json:"QualityScoreA,omitempty" xml:"QualityScoreA,omitempty"`
	QualityScoreB *float32 `json:"QualityScoreB,omitempty" xml:"QualityScoreB,omitempty"`
	RectAList     []*int64 `json:"RectAList,omitempty" xml:"RectAList,omitempty" type:"Repeated"`
	RectBList     []*int64 `json:"RectBList,omitempty" xml:"RectBList,omitempty" type:"Repeated"`
	Thresholds    []*int64 `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Repeated"`
}

func (s CompareFaceWithMaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceWithMaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompareFaceWithMaskResponseBodyData) GetConfidence() *float32 {
	return s.Confidence
}

func (s *CompareFaceWithMaskResponseBodyData) GetIsMaskA() *int64 {
	return s.IsMaskA
}

func (s *CompareFaceWithMaskResponseBodyData) GetIsMaskB() *int64 {
	return s.IsMaskB
}

func (s *CompareFaceWithMaskResponseBodyData) GetLandmarksAList() []*int64 {
	return s.LandmarksAList
}

func (s *CompareFaceWithMaskResponseBodyData) GetLandmarksBList() []*int64 {
	return s.LandmarksBList
}

func (s *CompareFaceWithMaskResponseBodyData) GetMessageTips() *string {
	return s.MessageTips
}

func (s *CompareFaceWithMaskResponseBodyData) GetQualityScoreA() *float32 {
	return s.QualityScoreA
}

func (s *CompareFaceWithMaskResponseBodyData) GetQualityScoreB() *float32 {
	return s.QualityScoreB
}

func (s *CompareFaceWithMaskResponseBodyData) GetRectAList() []*int64 {
	return s.RectAList
}

func (s *CompareFaceWithMaskResponseBodyData) GetRectBList() []*int64 {
	return s.RectBList
}

func (s *CompareFaceWithMaskResponseBodyData) GetThresholds() []*int64 {
	return s.Thresholds
}

func (s *CompareFaceWithMaskResponseBodyData) SetConfidence(v float32) *CompareFaceWithMaskResponseBodyData {
	s.Confidence = &v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetIsMaskA(v int64) *CompareFaceWithMaskResponseBodyData {
	s.IsMaskA = &v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetIsMaskB(v int64) *CompareFaceWithMaskResponseBodyData {
	s.IsMaskB = &v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetLandmarksAList(v []*int64) *CompareFaceWithMaskResponseBodyData {
	s.LandmarksAList = v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetLandmarksBList(v []*int64) *CompareFaceWithMaskResponseBodyData {
	s.LandmarksBList = v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetMessageTips(v string) *CompareFaceWithMaskResponseBodyData {
	s.MessageTips = &v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetQualityScoreA(v float32) *CompareFaceWithMaskResponseBodyData {
	s.QualityScoreA = &v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetQualityScoreB(v float32) *CompareFaceWithMaskResponseBodyData {
	s.QualityScoreB = &v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetRectAList(v []*int64) *CompareFaceWithMaskResponseBodyData {
	s.RectAList = v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetRectBList(v []*int64) *CompareFaceWithMaskResponseBodyData {
	s.RectBList = v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) SetThresholds(v []*int64) *CompareFaceWithMaskResponseBodyData {
	s.Thresholds = v
	return s
}

func (s *CompareFaceWithMaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
