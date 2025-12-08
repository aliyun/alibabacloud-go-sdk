// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectFaceResponseBodyData) *DetectFaceResponseBody
	GetData() *DetectFaceResponseBodyData
	SetRequestId(v string) *DetectFaceResponseBody
	GetRequestId() *string
}

type DetectFaceResponseBody struct {
	Data      *DetectFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBody) GetData() *DetectFaceResponseBodyData {
	return s.Data
}

func (s *DetectFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectFaceResponseBody) SetData(v *DetectFaceResponseBodyData) *DetectFaceResponseBody {
	s.Data = v
	return s
}

func (s *DetectFaceResponseBody) SetRequestId(v string) *DetectFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectFaceResponseBodyData struct {
	// example:
	//
	// 1
	FaceCount *int32 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	// 1
	FaceProbabilityList []*float32 `json:"FaceProbabilityList,omitempty" xml:"FaceProbabilityList,omitempty" type:"Repeated"`
	// 1
	FaceRectangles []*int32 `json:"FaceRectangles,omitempty" xml:"FaceRectangles,omitempty" type:"Repeated"`
	// example:
	//
	// 105
	LandmarkCount *int32     `json:"LandmarkCount,omitempty" xml:"LandmarkCount,omitempty"`
	LandmarkScore []*float32 `json:"LandmarkScore,omitempty" xml:"LandmarkScore,omitempty" type:"Repeated"`
	// 1
	Landmarks []*float32 `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Repeated"`
	// 1
	PoseList []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	// 1
	Pupils    []*float64                           `json:"Pupils,omitempty" xml:"Pupils,omitempty" type:"Repeated"`
	Qualities *DetectFaceResponseBodyDataQualities `json:"Qualities,omitempty" xml:"Qualities,omitempty" type:"Struct"`
}

func (s DetectFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBodyData) GetFaceCount() *int32 {
	return s.FaceCount
}

func (s *DetectFaceResponseBodyData) GetFaceProbabilityList() []*float32 {
	return s.FaceProbabilityList
}

func (s *DetectFaceResponseBodyData) GetFaceRectangles() []*int32 {
	return s.FaceRectangles
}

func (s *DetectFaceResponseBodyData) GetLandmarkCount() *int32 {
	return s.LandmarkCount
}

func (s *DetectFaceResponseBodyData) GetLandmarkScore() []*float32 {
	return s.LandmarkScore
}

func (s *DetectFaceResponseBodyData) GetLandmarks() []*float32 {
	return s.Landmarks
}

func (s *DetectFaceResponseBodyData) GetPoseList() []*float32 {
	return s.PoseList
}

func (s *DetectFaceResponseBodyData) GetPupils() []*float64 {
	return s.Pupils
}

func (s *DetectFaceResponseBodyData) GetQualities() *DetectFaceResponseBodyDataQualities {
	return s.Qualities
}

func (s *DetectFaceResponseBodyData) SetFaceCount(v int32) *DetectFaceResponseBodyData {
	s.FaceCount = &v
	return s
}

func (s *DetectFaceResponseBodyData) SetFaceProbabilityList(v []*float32) *DetectFaceResponseBodyData {
	s.FaceProbabilityList = v
	return s
}

func (s *DetectFaceResponseBodyData) SetFaceRectangles(v []*int32) *DetectFaceResponseBodyData {
	s.FaceRectangles = v
	return s
}

func (s *DetectFaceResponseBodyData) SetLandmarkCount(v int32) *DetectFaceResponseBodyData {
	s.LandmarkCount = &v
	return s
}

func (s *DetectFaceResponseBodyData) SetLandmarkScore(v []*float32) *DetectFaceResponseBodyData {
	s.LandmarkScore = v
	return s
}

func (s *DetectFaceResponseBodyData) SetLandmarks(v []*float32) *DetectFaceResponseBodyData {
	s.Landmarks = v
	return s
}

func (s *DetectFaceResponseBodyData) SetPoseList(v []*float32) *DetectFaceResponseBodyData {
	s.PoseList = v
	return s
}

func (s *DetectFaceResponseBodyData) SetPupils(v []*float64) *DetectFaceResponseBodyData {
	s.Pupils = v
	return s
}

func (s *DetectFaceResponseBodyData) SetQualities(v *DetectFaceResponseBodyDataQualities) *DetectFaceResponseBodyData {
	s.Qualities = v
	return s
}

func (s *DetectFaceResponseBodyData) Validate() error {
	if s.Qualities != nil {
		if err := s.Qualities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectFaceResponseBodyDataQualities struct {
	// 1
	BlurList []*float32 `json:"BlurList,omitempty" xml:"BlurList,omitempty" type:"Repeated"`
	// 1
	FnfList []*float32 `json:"FnfList,omitempty" xml:"FnfList,omitempty" type:"Repeated"`
	// 1
	GlassList []*float32 `json:"GlassList,omitempty" xml:"GlassList,omitempty" type:"Repeated"`
	// 1
	IlluList []*float32 `json:"IlluList,omitempty" xml:"IlluList,omitempty" type:"Repeated"`
	// 1
	MaskList []*float32 `json:"MaskList,omitempty" xml:"MaskList,omitempty" type:"Repeated"`
	// 1
	NoiseList []*float32 `json:"NoiseList,omitempty" xml:"NoiseList,omitempty" type:"Repeated"`
	// 1
	PoseList []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	// 1
	ScoreList []*float32 `json:"ScoreList,omitempty" xml:"ScoreList,omitempty" type:"Repeated"`
}

func (s DetectFaceResponseBodyDataQualities) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceResponseBodyDataQualities) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBodyDataQualities) GetBlurList() []*float32 {
	return s.BlurList
}

func (s *DetectFaceResponseBodyDataQualities) GetFnfList() []*float32 {
	return s.FnfList
}

func (s *DetectFaceResponseBodyDataQualities) GetGlassList() []*float32 {
	return s.GlassList
}

func (s *DetectFaceResponseBodyDataQualities) GetIlluList() []*float32 {
	return s.IlluList
}

func (s *DetectFaceResponseBodyDataQualities) GetMaskList() []*float32 {
	return s.MaskList
}

func (s *DetectFaceResponseBodyDataQualities) GetNoiseList() []*float32 {
	return s.NoiseList
}

func (s *DetectFaceResponseBodyDataQualities) GetPoseList() []*float32 {
	return s.PoseList
}

func (s *DetectFaceResponseBodyDataQualities) GetScoreList() []*float32 {
	return s.ScoreList
}

func (s *DetectFaceResponseBodyDataQualities) SetBlurList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.BlurList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetFnfList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.FnfList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetGlassList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.GlassList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetIlluList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.IlluList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetMaskList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.MaskList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetNoiseList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.NoiseList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetPoseList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.PoseList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) SetScoreList(v []*float32) *DetectFaceResponseBodyDataQualities {
	s.ScoreList = v
	return s
}

func (s *DetectFaceResponseBodyDataQualities) Validate() error {
	return dara.Validate(s)
}
