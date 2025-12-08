// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoLivingFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectVideoLivingFaceResponseBodyData) *DetectVideoLivingFaceResponseBody
	GetData() *DetectVideoLivingFaceResponseBodyData
	SetRequestId(v string) *DetectVideoLivingFaceResponseBody
	GetRequestId() *string
}

type DetectVideoLivingFaceResponseBody struct {
	Data *DetectVideoLivingFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 8E92F28B-F889-48CB-9FCC-3207CFA2215E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVideoLivingFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBody) GetData() *DetectVideoLivingFaceResponseBodyData {
	return s.Data
}

func (s *DetectVideoLivingFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectVideoLivingFaceResponseBody) SetData(v *DetectVideoLivingFaceResponseBodyData) *DetectVideoLivingFaceResponseBody {
	s.Data = v
	return s
}

func (s *DetectVideoLivingFaceResponseBody) SetRequestId(v string) *DetectVideoLivingFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVideoLivingFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectVideoLivingFaceResponseBodyData struct {
	Elements []*DetectVideoLivingFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectVideoLivingFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBodyData) GetElements() []*DetectVideoLivingFaceResponseBodyDataElements {
	return s.Elements
}

func (s *DetectVideoLivingFaceResponseBodyData) SetElements(v []*DetectVideoLivingFaceResponseBodyDataElements) *DetectVideoLivingFaceResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyData) Validate() error {
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

type DetectVideoLivingFaceResponseBodyDataElements struct {
	// example:
	//
	// 0.84644949436187744
	FaceConfidence *float32 `json:"FaceConfidence,omitempty" xml:"FaceConfidence,omitempty"`
	// example:
	//
	// 0.72464925050735474
	LiveConfidence *float32 `json:"LiveConfidence,omitempty" xml:"LiveConfidence,omitempty"`
	Rect           []*int32 `json:"Rect,omitempty" xml:"Rect,omitempty" type:"Repeated"`
}

func (s DetectVideoLivingFaceResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoLivingFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) GetFaceConfidence() *float32 {
	return s.FaceConfidence
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) GetLiveConfidence() *float32 {
	return s.LiveConfidence
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) GetRect() []*int32 {
	return s.Rect
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetFaceConfidence(v float32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.FaceConfidence = &v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetLiveConfidence(v float32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.LiveConfidence = &v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) SetRect(v []*int32) *DetectVideoLivingFaceResponseBodyDataElements {
	s.Rect = v
	return s
}

func (s *DetectVideoLivingFaceResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
