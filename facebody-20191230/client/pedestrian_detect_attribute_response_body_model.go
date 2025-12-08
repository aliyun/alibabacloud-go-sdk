// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPedestrianDetectAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PedestrianDetectAttributeResponseBodyData) *PedestrianDetectAttributeResponseBody
	GetData() *PedestrianDetectAttributeResponseBodyData
	SetRequestId(v string) *PedestrianDetectAttributeResponseBody
	GetRequestId() *string
}

type PedestrianDetectAttributeResponseBody struct {
	Data *PedestrianDetectAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4F609A30-F645-481E-A513-CADAA14DFB0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PedestrianDetectAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBody) GetData() *PedestrianDetectAttributeResponseBodyData {
	return s.Data
}

func (s *PedestrianDetectAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PedestrianDetectAttributeResponseBody) SetData(v *PedestrianDetectAttributeResponseBodyData) *PedestrianDetectAttributeResponseBody {
	s.Data = v
	return s
}

func (s *PedestrianDetectAttributeResponseBody) SetRequestId(v string) *PedestrianDetectAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PedestrianDetectAttributeResponseBodyData struct {
	Attributes []*PedestrianDetectAttributeResponseBodyDataAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Boxes      []*PedestrianDetectAttributeResponseBodyDataBoxes      `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// example:
	//
	// 584
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1
	PersonNumber *int32 `json:"PersonNumber,omitempty" xml:"PersonNumber,omitempty"`
	// example:
	//
	// 264
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyData) GetAttributes() []*PedestrianDetectAttributeResponseBodyDataAttributes {
	return s.Attributes
}

func (s *PedestrianDetectAttributeResponseBodyData) GetBoxes() []*PedestrianDetectAttributeResponseBodyDataBoxes {
	return s.Boxes
}

func (s *PedestrianDetectAttributeResponseBodyData) GetHeight() *int64 {
	return s.Height
}

func (s *PedestrianDetectAttributeResponseBodyData) GetPersonNumber() *int32 {
	return s.PersonNumber
}

func (s *PedestrianDetectAttributeResponseBodyData) GetWidth() *int64 {
	return s.Width
}

func (s *PedestrianDetectAttributeResponseBodyData) SetAttributes(v []*PedestrianDetectAttributeResponseBodyDataAttributes) *PedestrianDetectAttributeResponseBodyData {
	s.Attributes = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetBoxes(v []*PedestrianDetectAttributeResponseBodyDataBoxes) *PedestrianDetectAttributeResponseBodyData {
	s.Boxes = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetHeight(v int64) *PedestrianDetectAttributeResponseBodyData {
	s.Height = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetPersonNumber(v int32) *PedestrianDetectAttributeResponseBodyData {
	s.PersonNumber = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) SetWidth(v int64) *PedestrianDetectAttributeResponseBodyData {
	s.Width = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyData) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Boxes != nil {
		for _, item := range s.Boxes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PedestrianDetectAttributeResponseBodyDataAttributes struct {
	Age         *PedestrianDetectAttributeResponseBodyDataAttributesAge         `json:"Age,omitempty" xml:"Age,omitempty" type:"Struct"`
	Backpack    *PedestrianDetectAttributeResponseBodyDataAttributesBackpack    `json:"Backpack,omitempty" xml:"Backpack,omitempty" type:"Struct"`
	Gender      *PedestrianDetectAttributeResponseBodyDataAttributesGender      `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	Glasses     *PedestrianDetectAttributeResponseBodyDataAttributesGlasses     `json:"Glasses,omitempty" xml:"Glasses,omitempty" type:"Struct"`
	Handbag     *PedestrianDetectAttributeResponseBodyDataAttributesHandbag     `json:"Handbag,omitempty" xml:"Handbag,omitempty" type:"Struct"`
	Hat         *PedestrianDetectAttributeResponseBodyDataAttributesHat         `json:"Hat,omitempty" xml:"Hat,omitempty" type:"Struct"`
	LowerColor  *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor  `json:"LowerColor,omitempty" xml:"LowerColor,omitempty" type:"Struct"`
	LowerWear   *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear   `json:"LowerWear,omitempty" xml:"LowerWear,omitempty" type:"Struct"`
	Orient      *PedestrianDetectAttributeResponseBodyDataAttributesOrient      `json:"Orient,omitempty" xml:"Orient,omitempty" type:"Struct"`
	ShoulderBag *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag `json:"ShoulderBag,omitempty" xml:"ShoulderBag,omitempty" type:"Struct"`
	UpperColor  *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor  `json:"UpperColor,omitempty" xml:"UpperColor,omitempty" type:"Struct"`
	UpperWear   *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear   `json:"UpperWear,omitempty" xml:"UpperWear,omitempty" type:"Struct"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributes) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetAge() *PedestrianDetectAttributeResponseBodyDataAttributesAge {
	return s.Age
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetBackpack() *PedestrianDetectAttributeResponseBodyDataAttributesBackpack {
	return s.Backpack
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetGender() *PedestrianDetectAttributeResponseBodyDataAttributesGender {
	return s.Gender
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetGlasses() *PedestrianDetectAttributeResponseBodyDataAttributesGlasses {
	return s.Glasses
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetHandbag() *PedestrianDetectAttributeResponseBodyDataAttributesHandbag {
	return s.Handbag
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetHat() *PedestrianDetectAttributeResponseBodyDataAttributesHat {
	return s.Hat
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetLowerColor() *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor {
	return s.LowerColor
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetLowerWear() *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear {
	return s.LowerWear
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetOrient() *PedestrianDetectAttributeResponseBodyDataAttributesOrient {
	return s.Orient
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetShoulderBag() *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag {
	return s.ShoulderBag
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetUpperColor() *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor {
	return s.UpperColor
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) GetUpperWear() *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear {
	return s.UpperWear
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetAge(v *PedestrianDetectAttributeResponseBodyDataAttributesAge) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Age = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetBackpack(v *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Backpack = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetGender(v *PedestrianDetectAttributeResponseBodyDataAttributesGender) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Gender = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetGlasses(v *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Glasses = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetHandbag(v *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Handbag = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetHat(v *PedestrianDetectAttributeResponseBodyDataAttributesHat) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Hat = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetLowerColor(v *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.LowerColor = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetLowerWear(v *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.LowerWear = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetOrient(v *PedestrianDetectAttributeResponseBodyDataAttributesOrient) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.Orient = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetShoulderBag(v *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.ShoulderBag = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetUpperColor(v *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.UpperColor = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) SetUpperWear(v *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) *PedestrianDetectAttributeResponseBodyDataAttributes {
	s.UpperWear = v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributes) Validate() error {
	if s.Age != nil {
		if err := s.Age.Validate(); err != nil {
			return err
		}
	}
	if s.Backpack != nil {
		if err := s.Backpack.Validate(); err != nil {
			return err
		}
	}
	if s.Gender != nil {
		if err := s.Gender.Validate(); err != nil {
			return err
		}
	}
	if s.Glasses != nil {
		if err := s.Glasses.Validate(); err != nil {
			return err
		}
	}
	if s.Handbag != nil {
		if err := s.Handbag.Validate(); err != nil {
			return err
		}
	}
	if s.Hat != nil {
		if err := s.Hat.Validate(); err != nil {
			return err
		}
	}
	if s.LowerColor != nil {
		if err := s.LowerColor.Validate(); err != nil {
			return err
		}
	}
	if s.LowerWear != nil {
		if err := s.LowerWear.Validate(); err != nil {
			return err
		}
	}
	if s.Orient != nil {
		if err := s.Orient.Validate(); err != nil {
			return err
		}
	}
	if s.ShoulderBag != nil {
		if err := s.ShoulderBag.Validate(); err != nil {
			return err
		}
	}
	if s.UpperColor != nil {
		if err := s.UpperColor.Validate(); err != nil {
			return err
		}
	}
	if s.UpperWear != nil {
		if err := s.UpperWear.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PedestrianDetectAttributeResponseBodyDataAttributesAge struct {
	// example:
	//
	// Age18-60
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.99590516090393066
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesAge) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesAge) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesAge {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesAge {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesAge) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesBackpack struct {
	// example:
	//
	// No
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.96486538648605347
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesBackpack) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesBackpack) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesBackpack {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesBackpack {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesBackpack) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesGender struct {
	// example:
	//
	// female
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.97989875078201294
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGender) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGender) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesGender {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesGender {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGender) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesGlasses struct {
	// example:
	//
	// No
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.87284471094608307
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGlasses) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesGlasses) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesGlasses {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesGlasses {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesGlasses) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesHandbag struct {
	// example:
	//
	// Yes
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.55011671781539917
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHandbag) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHandbag) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesHandbag {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesHandbag {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHandbag) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesHat struct {
	// example:
	//
	// No
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.98272394016385078
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHat) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesHat) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesHat {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesHat {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesHat) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesLowerColor struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.69961744546890259
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerColor) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesLowerWear struct {
	// example:
	//
	// Trousers
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.6424860954284668
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesLowerWear) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesOrient struct {
	// example:
	//
	// Front
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.97838658094406128
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesOrient) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesOrient) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesOrient {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesOrient {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesOrient) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag struct {
	// example:
	//
	// No
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.91198787838220596
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesShoulderBag) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesUpperColor struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.97796273231506348
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperColor) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataAttributesUpperWear struct {
	// example:
	//
	// ShortSleeve
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.89291918277740479
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) GetName() *string {
	return s.Name
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) SetName(v string) *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear {
	s.Name = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataAttributesUpperWear) Validate() error {
	return dara.Validate(s)
}

type PedestrianDetectAttributeResponseBodyDataBoxes struct {
	// example:
	//
	// 584
	BottomRightX *float32 `json:"BottomRightX,omitempty" xml:"BottomRightX,omitempty"`
	// example:
	//
	// 218
	BottomRightY *float32 `json:"BottomRightY,omitempty" xml:"BottomRightY,omitempty"`
	// example:
	//
	// 0.88381063938140869
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 36
	TopLeftX *float32 `json:"TopLeftX,omitempty" xml:"TopLeftX,omitempty"`
	// example:
	//
	// 27
	TopLeftY *float32 `json:"TopLeftY,omitempty" xml:"TopLeftY,omitempty"`
}

func (s PedestrianDetectAttributeResponseBodyDataBoxes) String() string {
	return dara.Prettify(s)
}

func (s PedestrianDetectAttributeResponseBodyDataBoxes) GoString() string {
	return s.String()
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) GetBottomRightX() *float32 {
	return s.BottomRightX
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) GetBottomRightY() *float32 {
	return s.BottomRightY
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) GetScore() *float32 {
	return s.Score
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) GetTopLeftX() *float32 {
	return s.TopLeftX
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) GetTopLeftY() *float32 {
	return s.TopLeftY
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetBottomRightX(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.BottomRightX = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetBottomRightY(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.BottomRightY = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetScore(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.Score = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetTopLeftX(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.TopLeftX = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) SetTopLeftY(v float32) *PedestrianDetectAttributeResponseBodyDataBoxes {
	s.TopLeftY = &v
	return s
}

func (s *PedestrianDetectAttributeResponseBodyDataBoxes) Validate() error {
	return dara.Validate(s)
}
