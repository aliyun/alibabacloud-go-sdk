// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleIllegalParkingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectVehicleIllegalParkingResponseBodyData) *DetectVehicleIllegalParkingResponseBody
	GetData() *DetectVehicleIllegalParkingResponseBodyData
	SetRequestId(v string) *DetectVehicleIllegalParkingResponseBody
	GetRequestId() *string
}

type DetectVehicleIllegalParkingResponseBody struct {
	Data *DetectVehicleIllegalParkingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DB882EDD-991A-5A0C-A19B-CC7C4BA65E35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVehicleIllegalParkingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBody) GetData() *DetectVehicleIllegalParkingResponseBodyData {
	return s.Data
}

func (s *DetectVehicleIllegalParkingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectVehicleIllegalParkingResponseBody) SetData(v *DetectVehicleIllegalParkingResponseBodyData) *DetectVehicleIllegalParkingResponseBody {
	s.Data = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBody) SetRequestId(v string) *DetectVehicleIllegalParkingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingResponseBodyData struct {
	Elements         []*DetectVehicleIllegalParkingResponseBodyDataElements         `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	RegionIntersects []*DetectVehicleIllegalParkingResponseBodyDataRegionIntersects `json:"RegionIntersects,omitempty" xml:"RegionIntersects,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyData) GetElements() []*DetectVehicleIllegalParkingResponseBodyDataElements {
	return s.Elements
}

func (s *DetectVehicleIllegalParkingResponseBodyData) GetRegionIntersects() []*DetectVehicleIllegalParkingResponseBodyDataRegionIntersects {
	return s.RegionIntersects
}

func (s *DetectVehicleIllegalParkingResponseBodyData) SetElements(v []*DetectVehicleIllegalParkingResponseBodyDataElements) *DetectVehicleIllegalParkingResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyData) SetRegionIntersects(v []*DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) *DetectVehicleIllegalParkingResponseBodyData {
	s.RegionIntersects = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingResponseBodyDataElements struct {
	Boxes []*DetectVehicleIllegalParkingResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0.9599609375
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// vehicle
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DetectVehicleIllegalParkingResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) GetBoxes() []*DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	return s.Boxes
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) GetId() *int64 {
	return s.Id
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) GetScore() *float32 {
	return s.Score
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) GetTypeName() *string {
	return s.TypeName
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetBoxes(v []*DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetId(v int64) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.Id = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetScore(v float32) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) SetTypeName(v string) *DetectVehicleIllegalParkingResponseBodyDataElements {
	s.TypeName = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingResponseBodyDataElementsBoxes struct {
	// example:
	//
	// 268
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// example:
	//
	// 413
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 499
	Right *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	// example:
	//
	// 138
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) GetBottom() *int64 {
	return s.Bottom
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) GetLeft() *int64 {
	return s.Left
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) GetRight() *int64 {
	return s.Right
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) GetTop() *int64 {
	return s.Top
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetBottom(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Bottom = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetLeft(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Left = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetRight(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Right = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) SetTop(v int64) *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes {
	s.Top = &v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataElementsBoxes) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingResponseBodyDataRegionIntersects struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) GetIds() []*int64 {
	return s.Ids
}

func (s *DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) SetIds(v []*int64) *DetectVehicleIllegalParkingResponseBodyDataRegionIntersects {
	s.Ids = v
	return s
}

func (s *DetectVehicleIllegalParkingResponseBodyDataRegionIntersects) Validate() error {
	return dara.Validate(s)
}
