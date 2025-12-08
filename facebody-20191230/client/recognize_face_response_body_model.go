// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeFaceResponseBodyData) *RecognizeFaceResponseBody
	GetData() *RecognizeFaceResponseBodyData
	SetRequestId(v string) *RecognizeFaceResponseBody
	GetRequestId() *string
}

type RecognizeFaceResponseBody struct {
	Data *RecognizeFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 8251C88E-8273-4DBF-94FB-A6BCB268CEA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBody) GetData() *RecognizeFaceResponseBodyData {
	return s.Data
}

func (s *RecognizeFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeFaceResponseBody) SetData(v *RecognizeFaceResponseBodyData) *RecognizeFaceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeFaceResponseBody) SetRequestId(v string) *RecognizeFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeFaceResponseBodyData struct {
	// 1
	AgeList []*int32 `json:"AgeList,omitempty" xml:"AgeList,omitempty" type:"Repeated"`
	// 1
	BeautyList []*float32 `json:"BeautyList,omitempty" xml:"BeautyList,omitempty" type:"Repeated"`
	// example:
	//
	// 1024
	DenseFeatureLength *int32 `json:"DenseFeatureLength,omitempty" xml:"DenseFeatureLength,omitempty"`
	// 1
	DenseFeatures []*string `json:"DenseFeatures,omitempty" xml:"DenseFeatures,omitempty" type:"Repeated"`
	// 1
	Expressions []*int32 `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	FaceCount *int32 `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	// 1
	FaceProbabilityList []*float32 `json:"FaceProbabilityList,omitempty" xml:"FaceProbabilityList,omitempty" type:"Repeated"`
	// 1
	FaceRectangles []*int32 `json:"FaceRectangles,omitempty" xml:"FaceRectangles,omitempty" type:"Repeated"`
	// 1
	GenderList []*int32 `json:"GenderList,omitempty" xml:"GenderList,omitempty" type:"Repeated"`
	// 1
	Glasses []*int32 `json:"Glasses,omitempty" xml:"Glasses,omitempty" type:"Repeated"`
	// 1
	HatList []*int32 `json:"HatList,omitempty" xml:"HatList,omitempty" type:"Repeated"`
	// example:
	//
	// 105
	LandmarkCount *int32 `json:"LandmarkCount,omitempty" xml:"LandmarkCount,omitempty"`
	// 1
	Landmarks []*float32 `json:"Landmarks,omitempty" xml:"Landmarks,omitempty" type:"Repeated"`
	// 1
	Masks []*int64 `json:"Masks,omitempty" xml:"Masks,omitempty" type:"Repeated"`
	// 1
	PoseList []*float32 `json:"PoseList,omitempty" xml:"PoseList,omitempty" type:"Repeated"`
	// 1
	Pupils    []*float32                              `json:"Pupils,omitempty" xml:"Pupils,omitempty" type:"Repeated"`
	Qualities *RecognizeFaceResponseBodyDataQualities `json:"Qualities,omitempty" xml:"Qualities,omitempty" type:"Struct"`
}

func (s RecognizeFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBodyData) GetAgeList() []*int32 {
	return s.AgeList
}

func (s *RecognizeFaceResponseBodyData) GetBeautyList() []*float32 {
	return s.BeautyList
}

func (s *RecognizeFaceResponseBodyData) GetDenseFeatureLength() *int32 {
	return s.DenseFeatureLength
}

func (s *RecognizeFaceResponseBodyData) GetDenseFeatures() []*string {
	return s.DenseFeatures
}

func (s *RecognizeFaceResponseBodyData) GetExpressions() []*int32 {
	return s.Expressions
}

func (s *RecognizeFaceResponseBodyData) GetFaceCount() *int32 {
	return s.FaceCount
}

func (s *RecognizeFaceResponseBodyData) GetFaceProbabilityList() []*float32 {
	return s.FaceProbabilityList
}

func (s *RecognizeFaceResponseBodyData) GetFaceRectangles() []*int32 {
	return s.FaceRectangles
}

func (s *RecognizeFaceResponseBodyData) GetGenderList() []*int32 {
	return s.GenderList
}

func (s *RecognizeFaceResponseBodyData) GetGlasses() []*int32 {
	return s.Glasses
}

func (s *RecognizeFaceResponseBodyData) GetHatList() []*int32 {
	return s.HatList
}

func (s *RecognizeFaceResponseBodyData) GetLandmarkCount() *int32 {
	return s.LandmarkCount
}

func (s *RecognizeFaceResponseBodyData) GetLandmarks() []*float32 {
	return s.Landmarks
}

func (s *RecognizeFaceResponseBodyData) GetMasks() []*int64 {
	return s.Masks
}

func (s *RecognizeFaceResponseBodyData) GetPoseList() []*float32 {
	return s.PoseList
}

func (s *RecognizeFaceResponseBodyData) GetPupils() []*float32 {
	return s.Pupils
}

func (s *RecognizeFaceResponseBodyData) GetQualities() *RecognizeFaceResponseBodyDataQualities {
	return s.Qualities
}

func (s *RecognizeFaceResponseBodyData) SetAgeList(v []*int32) *RecognizeFaceResponseBodyData {
	s.AgeList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetBeautyList(v []*float32) *RecognizeFaceResponseBodyData {
	s.BeautyList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetDenseFeatureLength(v int32) *RecognizeFaceResponseBodyData {
	s.DenseFeatureLength = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetDenseFeatures(v []*string) *RecognizeFaceResponseBodyData {
	s.DenseFeatures = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetExpressions(v []*int32) *RecognizeFaceResponseBodyData {
	s.Expressions = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceCount(v int32) *RecognizeFaceResponseBodyData {
	s.FaceCount = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceProbabilityList(v []*float32) *RecognizeFaceResponseBodyData {
	s.FaceProbabilityList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetFaceRectangles(v []*int32) *RecognizeFaceResponseBodyData {
	s.FaceRectangles = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetGenderList(v []*int32) *RecognizeFaceResponseBodyData {
	s.GenderList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetGlasses(v []*int32) *RecognizeFaceResponseBodyData {
	s.Glasses = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetHatList(v []*int32) *RecognizeFaceResponseBodyData {
	s.HatList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetLandmarkCount(v int32) *RecognizeFaceResponseBodyData {
	s.LandmarkCount = &v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetLandmarks(v []*float32) *RecognizeFaceResponseBodyData {
	s.Landmarks = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetMasks(v []*int64) *RecognizeFaceResponseBodyData {
	s.Masks = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetPoseList(v []*float32) *RecognizeFaceResponseBodyData {
	s.PoseList = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetPupils(v []*float32) *RecognizeFaceResponseBodyData {
	s.Pupils = v
	return s
}

func (s *RecognizeFaceResponseBodyData) SetQualities(v *RecognizeFaceResponseBodyDataQualities) *RecognizeFaceResponseBodyData {
	s.Qualities = v
	return s
}

func (s *RecognizeFaceResponseBodyData) Validate() error {
	if s.Qualities != nil {
		if err := s.Qualities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeFaceResponseBodyDataQualities struct {
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

func (s RecognizeFaceResponseBodyDataQualities) String() string {
	return dara.Prettify(s)
}

func (s RecognizeFaceResponseBodyDataQualities) GoString() string {
	return s.String()
}

func (s *RecognizeFaceResponseBodyDataQualities) GetBlurList() []*float32 {
	return s.BlurList
}

func (s *RecognizeFaceResponseBodyDataQualities) GetFnfList() []*float32 {
	return s.FnfList
}

func (s *RecognizeFaceResponseBodyDataQualities) GetGlassList() []*float32 {
	return s.GlassList
}

func (s *RecognizeFaceResponseBodyDataQualities) GetIlluList() []*float32 {
	return s.IlluList
}

func (s *RecognizeFaceResponseBodyDataQualities) GetMaskList() []*float32 {
	return s.MaskList
}

func (s *RecognizeFaceResponseBodyDataQualities) GetNoiseList() []*float32 {
	return s.NoiseList
}

func (s *RecognizeFaceResponseBodyDataQualities) GetPoseList() []*float32 {
	return s.PoseList
}

func (s *RecognizeFaceResponseBodyDataQualities) GetScoreList() []*float32 {
	return s.ScoreList
}

func (s *RecognizeFaceResponseBodyDataQualities) SetBlurList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.BlurList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetFnfList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.FnfList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetGlassList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.GlassList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetIlluList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.IlluList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetMaskList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.MaskList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetNoiseList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.NoiseList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetPoseList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.PoseList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) SetScoreList(v []*float32) *RecognizeFaceResponseBodyDataQualities {
	s.ScoreList = v
	return s
}

func (s *RecognizeFaceResponseBodyDataQualities) Validate() error {
	return dara.Validate(s)
}
