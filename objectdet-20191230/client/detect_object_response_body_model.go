// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectObjectResponseBodyData) *DetectObjectResponseBody
	GetData() *DetectObjectResponseBodyData
	SetRequestId(v string) *DetectObjectResponseBody
	GetRequestId() *string
}

type DetectObjectResponseBody struct {
	Data *DetectObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 6EF97B44-2763-4EAD-8737-FB9F5EE25FE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBody) GetData() *DetectObjectResponseBodyData {
	return s.Data
}

func (s *DetectObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectObjectResponseBody) SetData(v *DetectObjectResponseBodyData) *DetectObjectResponseBody {
	s.Data = v
	return s
}

func (s *DetectObjectResponseBody) SetRequestId(v string) *DetectObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectObjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectObjectResponseBodyData struct {
	Elements []*DetectObjectResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// example:
	//
	// 300
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 533
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBodyData) GetElements() []*DetectObjectResponseBodyDataElements {
	return s.Elements
}

func (s *DetectObjectResponseBodyData) GetHeight() *int32 {
	return s.Height
}

func (s *DetectObjectResponseBodyData) GetWidth() *int32 {
	return s.Width
}

func (s *DetectObjectResponseBodyData) SetElements(v []*DetectObjectResponseBodyDataElements) *DetectObjectResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectObjectResponseBodyData) SetHeight(v int32) *DetectObjectResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectObjectResponseBodyData) SetWidth(v int32) *DetectObjectResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectObjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectObjectResponseBodyDataElements struct {
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// example:
	//
	// 0.266
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// chair
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectObjectResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectObjectResponseBodyDataElements) GetBoxes() []*int32 {
	return s.Boxes
}

func (s *DetectObjectResponseBodyDataElements) GetScore() *float32 {
	return s.Score
}

func (s *DetectObjectResponseBodyDataElements) GetType() *string {
	return s.Type
}

func (s *DetectObjectResponseBodyDataElements) SetBoxes(v []*int32) *DetectObjectResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectObjectResponseBodyDataElements) SetScore(v float32) *DetectObjectResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectObjectResponseBodyDataElements) SetType(v string) *DetectObjectResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectObjectResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
