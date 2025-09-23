// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectWorkwearResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectWorkwearResponseBodyData) *DetectWorkwearResponseBody
	GetData() *DetectWorkwearResponseBodyData
	SetRequestId(v string) *DetectWorkwearResponseBody
	GetRequestId() *string
}

type DetectWorkwearResponseBody struct {
	Data *DetectWorkwearResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 16CA8094-D7BC-51D4-8D55-6AC59AB20E0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectWorkwearResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearResponseBody) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBody) GetData() *DetectWorkwearResponseBodyData {
	return s.Data
}

func (s *DetectWorkwearResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectWorkwearResponseBody) SetData(v *DetectWorkwearResponseBodyData) *DetectWorkwearResponseBody {
	s.Data = v
	return s
}

func (s *DetectWorkwearResponseBody) SetRequestId(v string) *DetectWorkwearResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectWorkwearResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectWorkwearResponseBodyData struct {
	Elements []*DetectWorkwearResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectWorkwearResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyData) GetElements() []*DetectWorkwearResponseBodyDataElements {
	return s.Elements
}

func (s *DetectWorkwearResponseBodyData) SetElements(v []*DetectWorkwearResponseBodyDataElements) *DetectWorkwearResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectWorkwearResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectWorkwearResponseBodyDataElements struct {
	Property   []*DetectWorkwearResponseBodyDataElementsProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
	Rectangles *DetectWorkwearResponseBodyDataElementsRectangles `json:"Rectangles,omitempty" xml:"Rectangles,omitempty" type:"Struct"`
	// example:
	//
	// 0.63913465
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectWorkwearResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElements) GetProperty() []*DetectWorkwearResponseBodyDataElementsProperty {
	return s.Property
}

func (s *DetectWorkwearResponseBodyDataElements) GetRectangles() *DetectWorkwearResponseBodyDataElementsRectangles {
	return s.Rectangles
}

func (s *DetectWorkwearResponseBodyDataElements) GetScore() *float64 {
	return s.Score
}

func (s *DetectWorkwearResponseBodyDataElements) GetType() *string {
	return s.Type
}

func (s *DetectWorkwearResponseBodyDataElements) SetProperty(v []*DetectWorkwearResponseBodyDataElementsProperty) *DetectWorkwearResponseBodyDataElements {
	s.Property = v
	return s
}

func (s *DetectWorkwearResponseBodyDataElements) SetRectangles(v *DetectWorkwearResponseBodyDataElementsRectangles) *DetectWorkwearResponseBodyDataElements {
	s.Rectangles = v
	return s
}

func (s *DetectWorkwearResponseBodyDataElements) SetScore(v float64) *DetectWorkwearResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElements) SetType(v string) *DetectWorkwearResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type DetectWorkwearResponseBodyDataElementsProperty struct {
	// example:
	//
	// hat
	Label       *string                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	Probability *DetectWorkwearResponseBodyDataElementsPropertyProbability `json:"Probability,omitempty" xml:"Probability,omitempty" type:"Struct"`
}

func (s DetectWorkwearResponseBodyDataElementsProperty) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElementsProperty) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElementsProperty) GetLabel() *string {
	return s.Label
}

func (s *DetectWorkwearResponseBodyDataElementsProperty) GetProbability() *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	return s.Probability
}

func (s *DetectWorkwearResponseBodyDataElementsProperty) SetLabel(v string) *DetectWorkwearResponseBodyDataElementsProperty {
	s.Label = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsProperty) SetProbability(v *DetectWorkwearResponseBodyDataElementsPropertyProbability) *DetectWorkwearResponseBodyDataElementsProperty {
	s.Probability = v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsProperty) Validate() error {
	return dara.Validate(s)
}

type DetectWorkwearResponseBodyDataElementsPropertyProbability struct {
	// example:
	//
	// 0.00036084422
	No *float64 `json:"No,omitempty" xml:"No,omitempty"`
	// example:
	//
	// 0
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 0.0006906331
	Unknown *float64 `json:"Unknown,omitempty" xml:"Unknown,omitempty"`
	// example:
	//
	// 0.9989485
	Yes *float64 `json:"Yes,omitempty" xml:"Yes,omitempty"`
}

func (s DetectWorkwearResponseBodyDataElementsPropertyProbability) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElementsPropertyProbability) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) GetNo() *float64 {
	return s.No
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) GetThreshold() *int64 {
	return s.Threshold
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) GetUnknown() *float64 {
	return s.Unknown
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) GetYes() *float64 {
	return s.Yes
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetNo(v float64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.No = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetThreshold(v int64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.Threshold = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetUnknown(v float64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.Unknown = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) SetYes(v float64) *DetectWorkwearResponseBodyDataElementsPropertyProbability {
	s.Yes = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsPropertyProbability) Validate() error {
	return dara.Validate(s)
}

type DetectWorkwearResponseBodyDataElementsRectangles struct {
	// example:
	//
	// 96
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1067
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 426
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 88
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectWorkwearResponseBodyDataElementsRectangles) String() string {
	return dara.Prettify(s)
}

func (s DetectWorkwearResponseBodyDataElementsRectangles) GoString() string {
	return s.String()
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) GetHeight() *int64 {
	return s.Height
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) GetLeft() *int64 {
	return s.Left
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) GetTop() *int64 {
	return s.Top
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) GetWidth() *int64 {
	return s.Width
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetHeight(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Height = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetLeft(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Left = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetTop(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Top = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) SetWidth(v int64) *DetectWorkwearResponseBodyDataElementsRectangles {
	s.Width = &v
	return s
}

func (s *DetectWorkwearResponseBodyDataElementsRectangles) Validate() error {
	return dara.Validate(s)
}
