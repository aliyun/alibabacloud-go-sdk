// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleICongestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectVehicleICongestionResponseBodyData) *DetectVehicleICongestionResponseBody
	GetData() *DetectVehicleICongestionResponseBodyData
	SetRequestId(v string) *DetectVehicleICongestionResponseBody
	GetRequestId() *string
}

type DetectVehicleICongestionResponseBody struct {
	Data *DetectVehicleICongestionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4FC381BB-04F2-50F4-B54B-593042BCF3C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVehicleICongestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBody) GetData() *DetectVehicleICongestionResponseBodyData {
	return s.Data
}

func (s *DetectVehicleICongestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectVehicleICongestionResponseBody) SetData(v *DetectVehicleICongestionResponseBodyData) *DetectVehicleICongestionResponseBody {
	s.Data = v
	return s
}

func (s *DetectVehicleICongestionResponseBody) SetRequestId(v string) *DetectVehicleICongestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVehicleICongestionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionResponseBodyData struct {
	Elements                []*DetectVehicleICongestionResponseBodyDataElements                `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	RegionIntersectFeatures []*DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures `json:"RegionIntersectFeatures,omitempty" xml:"RegionIntersectFeatures,omitempty" type:"Repeated"`
	RegionIntersectMatched  []*DetectVehicleICongestionResponseBodyDataRegionIntersectMatched  `json:"RegionIntersectMatched,omitempty" xml:"RegionIntersectMatched,omitempty" type:"Repeated"`
	RegionIntersects        []*DetectVehicleICongestionResponseBodyDataRegionIntersects        `json:"RegionIntersects,omitempty" xml:"RegionIntersects,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyData) GetElements() []*DetectVehicleICongestionResponseBodyDataElements {
	return s.Elements
}

func (s *DetectVehicleICongestionResponseBodyData) GetRegionIntersectFeatures() []*DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures {
	return s.RegionIntersectFeatures
}

func (s *DetectVehicleICongestionResponseBodyData) GetRegionIntersectMatched() []*DetectVehicleICongestionResponseBodyDataRegionIntersectMatched {
	return s.RegionIntersectMatched
}

func (s *DetectVehicleICongestionResponseBodyData) GetRegionIntersects() []*DetectVehicleICongestionResponseBodyDataRegionIntersects {
	return s.RegionIntersects
}

func (s *DetectVehicleICongestionResponseBodyData) SetElements(v []*DetectVehicleICongestionResponseBodyDataElements) *DetectVehicleICongestionResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyData) SetRegionIntersectFeatures(v []*DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) *DetectVehicleICongestionResponseBodyData {
	s.RegionIntersectFeatures = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyData) SetRegionIntersectMatched(v []*DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) *DetectVehicleICongestionResponseBodyData {
	s.RegionIntersectMatched = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyData) SetRegionIntersects(v []*DetectVehicleICongestionResponseBodyDataRegionIntersects) *DetectVehicleICongestionResponseBodyData {
	s.RegionIntersects = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionResponseBodyDataElements struct {
	Boxes []*DetectVehicleICongestionResponseBodyDataElementsBoxes `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0.962890625
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// vehicle
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DetectVehicleICongestionResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataElements) GetBoxes() []*DetectVehicleICongestionResponseBodyDataElementsBoxes {
	return s.Boxes
}

func (s *DetectVehicleICongestionResponseBodyDataElements) GetId() *int64 {
	return s.Id
}

func (s *DetectVehicleICongestionResponseBodyDataElements) GetScore() *float32 {
	return s.Score
}

func (s *DetectVehicleICongestionResponseBodyDataElements) GetTypeName() *string {
	return s.TypeName
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetBoxes(v []*DetectVehicleICongestionResponseBodyDataElementsBoxes) *DetectVehicleICongestionResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetId(v int64) *DetectVehicleICongestionResponseBodyDataElements {
	s.Id = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetScore(v float32) *DetectVehicleICongestionResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) SetTypeName(v string) *DetectVehicleICongestionResponseBodyDataElements {
	s.TypeName = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionResponseBodyDataElementsBoxes struct {
	// example:
	//
	// 576
	Bottom *int64 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// example:
	//
	// 341
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 589
	Right *int64 `json:"Right,omitempty" xml:"Right,omitempty"`
	// example:
	//
	// 434
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DetectVehicleICongestionResponseBodyDataElementsBoxes) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataElementsBoxes) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) GetBottom() *int64 {
	return s.Bottom
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) GetLeft() *int64 {
	return s.Left
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) GetRight() *int64 {
	return s.Right
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) GetTop() *int64 {
	return s.Top
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetBottom(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Bottom = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetLeft(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Left = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetRight(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Right = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) SetTop(v int64) *DetectVehicleICongestionResponseBodyDataElementsBoxes {
	s.Top = &v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataElementsBoxes) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures struct {
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) GetFeatures() []*string {
	return s.Features
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) SetFeatures(v []*string) *DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures {
	s.Features = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectFeatures) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionResponseBodyDataRegionIntersectMatched struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) GetIds() []*int64 {
	return s.Ids
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) SetIds(v []*int64) *DetectVehicleICongestionResponseBodyDataRegionIntersectMatched {
	s.Ids = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersectMatched) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionResponseBodyDataRegionIntersects struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersects) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionResponseBodyDataRegionIntersects) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersects) GetIds() []*int64 {
	return s.Ids
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersects) SetIds(v []*int64) *DetectVehicleICongestionResponseBodyDataRegionIntersects {
	s.Ids = v
	return s
}

func (s *DetectVehicleICongestionResponseBodyDataRegionIntersects) Validate() error {
	return dara.Validate(s)
}
