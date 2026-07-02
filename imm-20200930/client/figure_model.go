// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFigure interface {
	dara.Model
	String() string
	GoString() string
	SetAge(v int64) *Figure
	GetAge() *int64
	SetAgeSD(v float32) *Figure
	GetAgeSD() *float32
	SetAttractive(v float32) *Figure
	GetAttractive() *float32
	SetBeard(v string) *Figure
	GetBeard() *string
	SetBeardConfidence(v float32) *Figure
	GetBeardConfidence() *float32
	SetBoundary(v *Boundary) *Figure
	GetBoundary() *Boundary
	SetEmotion(v string) *Figure
	GetEmotion() *string
	SetEmotionConfidence(v float32) *Figure
	GetEmotionConfidence() *float32
	SetFaceQuality(v float32) *Figure
	GetFaceQuality() *float32
	SetFigureClusterConfidence(v float32) *Figure
	GetFigureClusterConfidence() *float32
	SetFigureClusterId(v string) *Figure
	GetFigureClusterId() *string
	SetFigureConfidence(v float32) *Figure
	GetFigureConfidence() *float32
	SetFigureId(v string) *Figure
	GetFigureId() *string
	SetFigureType(v string) *Figure
	GetFigureType() *string
	SetGender(v string) *Figure
	GetGender() *string
	SetGenderConfidence(v float32) *Figure
	GetGenderConfidence() *float32
	SetGlasses(v string) *Figure
	GetGlasses() *string
	SetGlassesConfidence(v float32) *Figure
	GetGlassesConfidence() *float32
	SetHat(v string) *Figure
	GetHat() *string
	SetHatConfidence(v float32) *Figure
	GetHatConfidence() *float32
	SetHeadPose(v *HeadPose) *Figure
	GetHeadPose() *HeadPose
	SetMask(v string) *Figure
	GetMask() *string
	SetMaskConfidence(v float32) *Figure
	GetMaskConfidence() *float32
	SetMouth(v string) *Figure
	GetMouth() *string
	SetMouthConfidence(v float32) *Figure
	GetMouthConfidence() *float32
	SetSharpness(v float32) *Figure
	GetSharpness() *float32
}

type Figure struct {
	// The age.
	//
	// example:
	//
	// 29
	Age *int64 `json:"Age,omitempty" xml:"Age,omitempty"`
	// The age standard deviation.
	//
	// example:
	//
	// 5
	AgeSD *float32 `json:"AgeSD,omitempty" xml:"AgeSD,omitempty"`
	// The face attractiveness score. A higher score indicates greater attractiveness. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.96
	Attractive *float32 `json:"Attractive,omitempty" xml:"Attractive,omitempty"`
	// Indicates whether the person has a beard.
	//
	// example:
	//
	// none
	Beard *string `json:"Beard,omitempty" xml:"Beard,omitempty"`
	// The confidence score for whether the person has a beard. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 1
	BeardConfidence *float32 `json:"BeardConfidence,omitempty" xml:"BeardConfidence,omitempty"`
	// The face boundary information.
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// The emotion.
	//
	// example:
	//
	// happiness
	Emotion *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// The emotion confidence score. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 0.707845687866210
	EmotionConfidence *float32 `json:"EmotionConfidence,omitempty" xml:"EmotionConfidence,omitempty"`
	// The face quality score, which evaluates whether the image quality is suitable for face recognition. A higher score indicates better quality. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.960875928401947
	FaceQuality *float32 `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// The clustering confidence score. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 1
	FigureClusterConfidence *float32 `json:"FigureClusterConfidence,omitempty" xml:"FigureClusterConfidence,omitempty"`
	// The figure cluster ID. The following reserved special cluster ID names are included:
	//
	// - figure-cluster-id-independent: The face does not currently belong to any suitable group. After new images are added to the dataset and face clustering is performed, the face may be assigned to a group.
	//
	// - figure-cluster-id-unavailable: The face has not been clustered yet. This means that after new images were added to the dataset, face clustering has not been performed.
	//
	// example:
	//
	// Cluster-dbe72fec-b84c-4ab6-885b-3678e64****
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	// The confidence score of the figure object. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 1
	FigureConfidence *float32 `json:"FigureConfidence,omitempty" xml:"FigureConfidence,omitempty"`
	// The ID of the figure object.
	//
	// example:
	//
	// 868a9e74-cde5-4c7a-9013-28bb984****
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// The figure type.
	//
	// example:
	//
	// face
	FigureType *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
	// The gender.
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The gender confidence score. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 1
	GenderConfidence *float32 `json:"GenderConfidence,omitempty" xml:"GenderConfidence,omitempty"`
	// Indicates whether the person is wearing glasses.
	//
	// example:
	//
	// none
	Glasses *string `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	// The confidence score for whether the person is wearing glasses. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 0.8
	GlassesConfidence *float32 `json:"GlassesConfidence,omitempty" xml:"GlassesConfidence,omitempty"`
	// Indicates whether the person is wearing a hat.
	//
	// example:
	//
	// none
	Hat *string `json:"Hat,omitempty" xml:"Hat,omitempty"`
	// The confidence score for whether the person is wearing a hat. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 1
	HatConfidence *float32 `json:"HatConfidence,omitempty" xml:"HatConfidence,omitempty"`
	// The head pose information.
	HeadPose *HeadPose `json:"HeadPose,omitempty" xml:"HeadPose,omitempty"`
	// Indicates whether the person is wearing a mask.
	//
	// example:
	//
	// none
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The confidence score for whether the person is wearing a mask. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 1
	MaskConfidence *float32 `json:"MaskConfidence,omitempty" xml:"MaskConfidence,omitempty"`
	// Indicates whether the mouth is open.
	//
	// example:
	//
	// close
	Mouth *string `json:"Mouth,omitempty" xml:"Mouth,omitempty"`
	// The confidence score for whether the mouth is open. Valid values: 0 (lowest confidence) to 1 (highest confidence).
	//
	// example:
	//
	// 1
	MouthConfidence *float32 `json:"MouthConfidence,omitempty" xml:"MouthConfidence,omitempty"`
	// The sharpness score. A higher score indicates a clearer face. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.7
	Sharpness *float32 `json:"Sharpness,omitempty" xml:"Sharpness,omitempty"`
}

func (s Figure) String() string {
	return dara.Prettify(s)
}

func (s Figure) GoString() string {
	return s.String()
}

func (s *Figure) GetAge() *int64 {
	return s.Age
}

func (s *Figure) GetAgeSD() *float32 {
	return s.AgeSD
}

func (s *Figure) GetAttractive() *float32 {
	return s.Attractive
}

func (s *Figure) GetBeard() *string {
	return s.Beard
}

func (s *Figure) GetBeardConfidence() *float32 {
	return s.BeardConfidence
}

func (s *Figure) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *Figure) GetEmotion() *string {
	return s.Emotion
}

func (s *Figure) GetEmotionConfidence() *float32 {
	return s.EmotionConfidence
}

func (s *Figure) GetFaceQuality() *float32 {
	return s.FaceQuality
}

func (s *Figure) GetFigureClusterConfidence() *float32 {
	return s.FigureClusterConfidence
}

func (s *Figure) GetFigureClusterId() *string {
	return s.FigureClusterId
}

func (s *Figure) GetFigureConfidence() *float32 {
	return s.FigureConfidence
}

func (s *Figure) GetFigureId() *string {
	return s.FigureId
}

func (s *Figure) GetFigureType() *string {
	return s.FigureType
}

func (s *Figure) GetGender() *string {
	return s.Gender
}

func (s *Figure) GetGenderConfidence() *float32 {
	return s.GenderConfidence
}

func (s *Figure) GetGlasses() *string {
	return s.Glasses
}

func (s *Figure) GetGlassesConfidence() *float32 {
	return s.GlassesConfidence
}

func (s *Figure) GetHat() *string {
	return s.Hat
}

func (s *Figure) GetHatConfidence() *float32 {
	return s.HatConfidence
}

func (s *Figure) GetHeadPose() *HeadPose {
	return s.HeadPose
}

func (s *Figure) GetMask() *string {
	return s.Mask
}

func (s *Figure) GetMaskConfidence() *float32 {
	return s.MaskConfidence
}

func (s *Figure) GetMouth() *string {
	return s.Mouth
}

func (s *Figure) GetMouthConfidence() *float32 {
	return s.MouthConfidence
}

func (s *Figure) GetSharpness() *float32 {
	return s.Sharpness
}

func (s *Figure) SetAge(v int64) *Figure {
	s.Age = &v
	return s
}

func (s *Figure) SetAgeSD(v float32) *Figure {
	s.AgeSD = &v
	return s
}

func (s *Figure) SetAttractive(v float32) *Figure {
	s.Attractive = &v
	return s
}

func (s *Figure) SetBeard(v string) *Figure {
	s.Beard = &v
	return s
}

func (s *Figure) SetBeardConfidence(v float32) *Figure {
	s.BeardConfidence = &v
	return s
}

func (s *Figure) SetBoundary(v *Boundary) *Figure {
	s.Boundary = v
	return s
}

func (s *Figure) SetEmotion(v string) *Figure {
	s.Emotion = &v
	return s
}

func (s *Figure) SetEmotionConfidence(v float32) *Figure {
	s.EmotionConfidence = &v
	return s
}

func (s *Figure) SetFaceQuality(v float32) *Figure {
	s.FaceQuality = &v
	return s
}

func (s *Figure) SetFigureClusterConfidence(v float32) *Figure {
	s.FigureClusterConfidence = &v
	return s
}

func (s *Figure) SetFigureClusterId(v string) *Figure {
	s.FigureClusterId = &v
	return s
}

func (s *Figure) SetFigureConfidence(v float32) *Figure {
	s.FigureConfidence = &v
	return s
}

func (s *Figure) SetFigureId(v string) *Figure {
	s.FigureId = &v
	return s
}

func (s *Figure) SetFigureType(v string) *Figure {
	s.FigureType = &v
	return s
}

func (s *Figure) SetGender(v string) *Figure {
	s.Gender = &v
	return s
}

func (s *Figure) SetGenderConfidence(v float32) *Figure {
	s.GenderConfidence = &v
	return s
}

func (s *Figure) SetGlasses(v string) *Figure {
	s.Glasses = &v
	return s
}

func (s *Figure) SetGlassesConfidence(v float32) *Figure {
	s.GlassesConfidence = &v
	return s
}

func (s *Figure) SetHat(v string) *Figure {
	s.Hat = &v
	return s
}

func (s *Figure) SetHatConfidence(v float32) *Figure {
	s.HatConfidence = &v
	return s
}

func (s *Figure) SetHeadPose(v *HeadPose) *Figure {
	s.HeadPose = v
	return s
}

func (s *Figure) SetMask(v string) *Figure {
	s.Mask = &v
	return s
}

func (s *Figure) SetMaskConfidence(v float32) *Figure {
	s.MaskConfidence = &v
	return s
}

func (s *Figure) SetMouth(v string) *Figure {
	s.Mouth = &v
	return s
}

func (s *Figure) SetMouthConfidence(v float32) *Figure {
	s.MouthConfidence = &v
	return s
}

func (s *Figure) SetSharpness(v float32) *Figure {
	s.Sharpness = &v
	return s
}

func (s *Figure) Validate() error {
	if s.Boundary != nil {
		if err := s.Boundary.Validate(); err != nil {
			return err
		}
	}
	if s.HeadPose != nil {
		if err := s.HeadPose.Validate(); err != nil {
			return err
		}
	}
	return nil
}
