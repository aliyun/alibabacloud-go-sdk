// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeActionResponseBodyData) *RecognizeActionResponseBody
	GetData() *RecognizeActionResponseBodyData
	SetRequestId(v string) *RecognizeActionResponseBody
	GetRequestId() *string
}

type RecognizeActionResponseBody struct {
	Data *RecognizeActionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E9C40AF5-A7F8-49D5-8A0C-B21F15A07F17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBody) GetData() *RecognizeActionResponseBodyData {
	return s.Data
}

func (s *RecognizeActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeActionResponseBody) SetData(v *RecognizeActionResponseBodyData) *RecognizeActionResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeActionResponseBody) SetRequestId(v string) *RecognizeActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeActionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizeActionResponseBodyData struct {
	Elements []*RecognizeActionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeActionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyData) GetElements() []*RecognizeActionResponseBodyDataElements {
	return s.Elements
}

func (s *RecognizeActionResponseBodyData) SetElements(v []*RecognizeActionResponseBodyDataElements) *RecognizeActionResponseBodyData {
	s.Elements = v
	return s
}

func (s *RecognizeActionResponseBodyData) Validate() error {
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

type RecognizeActionResponseBodyDataElements struct {
	Boxes  []*RecognizeActionResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Labels []*string                                       `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Scores []*float32                                      `json:"Scores,omitempty" xml:"Scores,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s RecognizeActionResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyDataElements) GetBoxes() []*RecognizeActionResponseBodyDataElementsBoxes {
	return s.Boxes
}

func (s *RecognizeActionResponseBodyDataElements) GetLabels() []*string {
	return s.Labels
}

func (s *RecognizeActionResponseBodyDataElements) GetScores() []*float32 {
	return s.Scores
}

func (s *RecognizeActionResponseBodyDataElements) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *RecognizeActionResponseBodyDataElements) SetBoxes(v []*RecognizeActionResponseBodyDataElementsBoxes) *RecognizeActionResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetLabels(v []*string) *RecognizeActionResponseBodyDataElements {
	s.Labels = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetScores(v []*float32) *RecognizeActionResponseBodyDataElements {
	s.Scores = v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) SetTimestamp(v int32) *RecognizeActionResponseBodyDataElements {
	s.Timestamp = &v
	return s
}

func (s *RecognizeActionResponseBodyDataElements) Validate() error {
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

type RecognizeActionResponseBodyDataElementsBoxes struct {
	Box []*int32 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
}

func (s RecognizeActionResponseBodyDataElementsBoxes) String() string {
	return dara.Prettify(s)
}

func (s RecognizeActionResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *RecognizeActionResponseBodyDataElementsBoxes) GetBox() []*int32 {
	return s.Box
}

func (s *RecognizeActionResponseBodyDataElementsBoxes) SetBox(v []*int32) *RecognizeActionResponseBodyDataElementsBoxes {
	s.Box = v
	return s
}

func (s *RecognizeActionResponseBodyDataElementsBoxes) Validate() error {
	return dara.Validate(s)
}
