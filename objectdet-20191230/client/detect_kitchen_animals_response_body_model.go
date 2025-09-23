// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectKitchenAnimalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectKitchenAnimalsResponseBodyData) *DetectKitchenAnimalsResponseBody
	GetData() *DetectKitchenAnimalsResponseBodyData
	SetRequestId(v string) *DetectKitchenAnimalsResponseBody
	GetRequestId() *string
}

type DetectKitchenAnimalsResponseBody struct {
	Data *DetectKitchenAnimalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C306F16F-30E1-54F4-93DF-A52CCF6664D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectKitchenAnimalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBody) GetData() *DetectKitchenAnimalsResponseBodyData {
	return s.Data
}

func (s *DetectKitchenAnimalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectKitchenAnimalsResponseBody) SetData(v *DetectKitchenAnimalsResponseBodyData) *DetectKitchenAnimalsResponseBody {
	s.Data = v
	return s
}

func (s *DetectKitchenAnimalsResponseBody) SetRequestId(v string) *DetectKitchenAnimalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectKitchenAnimalsResponseBodyData struct {
	Elements []*DetectKitchenAnimalsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectKitchenAnimalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBodyData) GetElements() []*DetectKitchenAnimalsResponseBodyDataElements {
	return s.Elements
}

func (s *DetectKitchenAnimalsResponseBodyData) SetElements(v []*DetectKitchenAnimalsResponseBodyDataElements) *DetectKitchenAnimalsResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectKitchenAnimalsResponseBodyDataElements struct {
	Rectangles *DetectKitchenAnimalsResponseBodyDataElementsRectangles `json:"Rectangles,omitempty" xml:"Rectangles,omitempty" type:"Struct"`
	// example:
	//
	// 0.75105053
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// mouse
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectKitchenAnimalsResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) GetRectangles() *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	return s.Rectangles
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) GetScore() *float32 {
	return s.Score
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) GetType() *string {
	return s.Type
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) SetRectangles(v *DetectKitchenAnimalsResponseBodyDataElementsRectangles) *DetectKitchenAnimalsResponseBodyDataElements {
	s.Rectangles = v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) SetScore(v float32) *DetectKitchenAnimalsResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) SetType(v string) *DetectKitchenAnimalsResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type DetectKitchenAnimalsResponseBodyDataElementsRectangles struct {
	// example:
	//
	// 64
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 292
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 1048
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 64
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectKitchenAnimalsResponseBodyDataElementsRectangles) String() string {
	return dara.Prettify(s)
}

func (s DetectKitchenAnimalsResponseBodyDataElementsRectangles) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) GetHeight() *int64 {
	return s.Height
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) GetLeft() *int64 {
	return s.Left
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) GetTop() *int64 {
	return s.Top
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) GetWidth() *int64 {
	return s.Width
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetHeight(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Height = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetLeft(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Left = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetTop(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Top = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) SetWidth(v int64) *DetectKitchenAnimalsResponseBodyDataElementsRectangles {
	s.Width = &v
	return s
}

func (s *DetectKitchenAnimalsResponseBodyDataElementsRectangles) Validate() error {
	return dara.Validate(s)
}
