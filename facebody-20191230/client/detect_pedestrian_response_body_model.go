// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectPedestrianResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectPedestrianResponseBodyData) *DetectPedestrianResponseBody
	GetData() *DetectPedestrianResponseBodyData
	SetRequestId(v string) *DetectPedestrianResponseBody
	GetRequestId() *string
}

type DetectPedestrianResponseBody struct {
	Data *DetectPedestrianResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 45DBA494-6250-42C4-80D1-7CF19BDD2CB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectPedestrianResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectPedestrianResponseBody) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBody) GetData() *DetectPedestrianResponseBodyData {
	return s.Data
}

func (s *DetectPedestrianResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectPedestrianResponseBody) SetData(v *DetectPedestrianResponseBodyData) *DetectPedestrianResponseBody {
	s.Data = v
	return s
}

func (s *DetectPedestrianResponseBody) SetRequestId(v string) *DetectPedestrianResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectPedestrianResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectPedestrianResponseBodyData struct {
	Elements []*DetectPedestrianResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// example:
	//
	// 599
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 899
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectPedestrianResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectPedestrianResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBodyData) GetElements() []*DetectPedestrianResponseBodyDataElements {
	return s.Elements
}

func (s *DetectPedestrianResponseBodyData) GetHeight() *int32 {
	return s.Height
}

func (s *DetectPedestrianResponseBodyData) GetWidth() *int32 {
	return s.Width
}

func (s *DetectPedestrianResponseBodyData) SetElements(v []*DetectPedestrianResponseBodyDataElements) *DetectPedestrianResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectPedestrianResponseBodyData) SetHeight(v int32) *DetectPedestrianResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectPedestrianResponseBodyData) SetWidth(v int32) *DetectPedestrianResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectPedestrianResponseBodyData) Validate() error {
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

type DetectPedestrianResponseBodyDataElements struct {
	// 1
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// example:
	//
	// 0.999
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// person
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectPedestrianResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectPedestrianResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectPedestrianResponseBodyDataElements) GetBoxes() []*int32 {
	return s.Boxes
}

func (s *DetectPedestrianResponseBodyDataElements) GetScore() *float32 {
	return s.Score
}

func (s *DetectPedestrianResponseBodyDataElements) GetType() *string {
	return s.Type
}

func (s *DetectPedestrianResponseBodyDataElements) SetBoxes(v []*int32) *DetectPedestrianResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectPedestrianResponseBodyDataElements) SetScore(v float32) *DetectPedestrianResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectPedestrianResponseBodyDataElements) SetType(v string) *DetectPedestrianResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectPedestrianResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
