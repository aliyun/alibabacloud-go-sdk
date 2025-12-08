// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntelligentCompositionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *IntelligentCompositionResponseBodyData) *IntelligentCompositionResponseBody
	GetData() *IntelligentCompositionResponseBodyData
	SetRequestId(v string) *IntelligentCompositionResponseBody
	GetRequestId() *string
}

type IntelligentCompositionResponseBody struct {
	Data *IntelligentCompositionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C1D52018-D67A-46AD-9AAA-031750A6E770
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IntelligentCompositionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntelligentCompositionResponseBody) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBody) GetData() *IntelligentCompositionResponseBodyData {
	return s.Data
}

func (s *IntelligentCompositionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntelligentCompositionResponseBody) SetData(v *IntelligentCompositionResponseBodyData) *IntelligentCompositionResponseBody {
	s.Data = v
	return s
}

func (s *IntelligentCompositionResponseBody) SetRequestId(v string) *IntelligentCompositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntelligentCompositionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntelligentCompositionResponseBodyData struct {
	Elements []*IntelligentCompositionResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s IntelligentCompositionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s IntelligentCompositionResponseBodyData) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBodyData) GetElements() []*IntelligentCompositionResponseBodyDataElements {
	return s.Elements
}

func (s *IntelligentCompositionResponseBodyData) SetElements(v []*IntelligentCompositionResponseBodyDataElements) *IntelligentCompositionResponseBodyData {
	s.Elements = v
	return s
}

func (s *IntelligentCompositionResponseBodyData) Validate() error {
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

type IntelligentCompositionResponseBodyDataElements struct {
	// example:
	//
	// 981
	MaxX *int32 `json:"MaxX,omitempty" xml:"MaxX,omitempty"`
	// example:
	//
	// 672
	MaxY *int32 `json:"MaxY,omitempty" xml:"MaxY,omitempty"`
	// example:
	//
	// 43
	MinX *int32 `json:"MinX,omitempty" xml:"MinX,omitempty"`
	// example:
	//
	// 96
	MinY *int32 `json:"MinY,omitempty" xml:"MinY,omitempty"`
	// example:
	//
	// 3.6567564
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s IntelligentCompositionResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s IntelligentCompositionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *IntelligentCompositionResponseBodyDataElements) GetMaxX() *int32 {
	return s.MaxX
}

func (s *IntelligentCompositionResponseBodyDataElements) GetMaxY() *int32 {
	return s.MaxY
}

func (s *IntelligentCompositionResponseBodyDataElements) GetMinX() *int32 {
	return s.MinX
}

func (s *IntelligentCompositionResponseBodyDataElements) GetMinY() *int32 {
	return s.MinY
}

func (s *IntelligentCompositionResponseBodyDataElements) GetScore() *float32 {
	return s.Score
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMaxX(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MaxX = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMaxY(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MaxY = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMinX(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MinX = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetMinY(v int32) *IntelligentCompositionResponseBodyDataElements {
	s.MinY = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) SetScore(v float32) *IntelligentCompositionResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *IntelligentCompositionResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
