// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizePdfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizePdfResponseBodyData) *RecognizePdfResponseBody
	GetData() *RecognizePdfResponseBodyData
	SetRequestId(v string) *RecognizePdfResponseBody
	GetRequestId() *string
}

type RecognizePdfResponseBody struct {
	Data *RecognizePdfResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CD9A9659-ABEE-4A7D-837F-9FDF40879A97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePdfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizePdfResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponseBody) GetData() *RecognizePdfResponseBodyData {
	return s.Data
}

func (s *RecognizePdfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizePdfResponseBody) SetData(v *RecognizePdfResponseBodyData) *RecognizePdfResponseBody {
	s.Data = v
	return s
}

func (s *RecognizePdfResponseBody) SetRequestId(v string) *RecognizePdfResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizePdfResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizePdfResponseBodyData struct {
	// example:
	//
	// 0
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// 788
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 610
	OrgHeight *int64 `json:"OrgHeight,omitempty" xml:"OrgHeight,omitempty"`
	// example:
	//
	// 394
	OrgWidth *int64 `json:"OrgWidth,omitempty" xml:"OrgWidth,omitempty"`
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 1220
	Width     *int64                                   `json:"Width,omitempty" xml:"Width,omitempty"`
	WordsInfo []*RecognizePdfResponseBodyDataWordsInfo `json:"WordsInfo,omitempty" xml:"WordsInfo,omitempty" type:"Repeated"`
}

func (s RecognizePdfResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizePdfResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponseBodyData) GetAngle() *int64 {
	return s.Angle
}

func (s *RecognizePdfResponseBodyData) GetHeight() *int64 {
	return s.Height
}

func (s *RecognizePdfResponseBodyData) GetOrgHeight() *int64 {
	return s.OrgHeight
}

func (s *RecognizePdfResponseBodyData) GetOrgWidth() *int64 {
	return s.OrgWidth
}

func (s *RecognizePdfResponseBodyData) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *RecognizePdfResponseBodyData) GetWidth() *int64 {
	return s.Width
}

func (s *RecognizePdfResponseBodyData) GetWordsInfo() []*RecognizePdfResponseBodyDataWordsInfo {
	return s.WordsInfo
}

func (s *RecognizePdfResponseBodyData) SetAngle(v int64) *RecognizePdfResponseBodyData {
	s.Angle = &v
	return s
}

func (s *RecognizePdfResponseBodyData) SetHeight(v int64) *RecognizePdfResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizePdfResponseBodyData) SetOrgHeight(v int64) *RecognizePdfResponseBodyData {
	s.OrgHeight = &v
	return s
}

func (s *RecognizePdfResponseBodyData) SetOrgWidth(v int64) *RecognizePdfResponseBodyData {
	s.OrgWidth = &v
	return s
}

func (s *RecognizePdfResponseBodyData) SetPageIndex(v int64) *RecognizePdfResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *RecognizePdfResponseBodyData) SetWidth(v int64) *RecognizePdfResponseBodyData {
	s.Width = &v
	return s
}

func (s *RecognizePdfResponseBodyData) SetWordsInfo(v []*RecognizePdfResponseBodyDataWordsInfo) *RecognizePdfResponseBodyData {
	s.WordsInfo = v
	return s
}

func (s *RecognizePdfResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizePdfResponseBodyDataWordsInfo struct {
	// example:
	//
	// 0
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// 16
	Height    *int64                                            `json:"Height,omitempty" xml:"Height,omitempty"`
	Positions []*RecognizePdfResponseBodyDataWordsInfoPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
	// example:
	//
	// 205
	Width *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Word  *string `json:"Word,omitempty" xml:"Word,omitempty"`
	// example:
	//
	// 863
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 46
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizePdfResponseBodyDataWordsInfo) String() string {
	return dara.Prettify(s)
}

func (s RecognizePdfResponseBodyDataWordsInfo) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponseBodyDataWordsInfo) GetAngle() *int64 {
	return s.Angle
}

func (s *RecognizePdfResponseBodyDataWordsInfo) GetHeight() *int64 {
	return s.Height
}

func (s *RecognizePdfResponseBodyDataWordsInfo) GetPositions() []*RecognizePdfResponseBodyDataWordsInfoPositions {
	return s.Positions
}

func (s *RecognizePdfResponseBodyDataWordsInfo) GetWidth() *int64 {
	return s.Width
}

func (s *RecognizePdfResponseBodyDataWordsInfo) GetWord() *string {
	return s.Word
}

func (s *RecognizePdfResponseBodyDataWordsInfo) GetX() *int64 {
	return s.X
}

func (s *RecognizePdfResponseBodyDataWordsInfo) GetY() *int64 {
	return s.Y
}

func (s *RecognizePdfResponseBodyDataWordsInfo) SetAngle(v int64) *RecognizePdfResponseBodyDataWordsInfo {
	s.Angle = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfo) SetHeight(v int64) *RecognizePdfResponseBodyDataWordsInfo {
	s.Height = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfo) SetPositions(v []*RecognizePdfResponseBodyDataWordsInfoPositions) *RecognizePdfResponseBodyDataWordsInfo {
	s.Positions = v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfo) SetWidth(v int64) *RecognizePdfResponseBodyDataWordsInfo {
	s.Width = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfo) SetWord(v string) *RecognizePdfResponseBodyDataWordsInfo {
	s.Word = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfo) SetX(v int64) *RecognizePdfResponseBodyDataWordsInfo {
	s.X = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfo) SetY(v int64) *RecognizePdfResponseBodyDataWordsInfo {
	s.Y = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfo) Validate() error {
	return dara.Validate(s)
}

type RecognizePdfResponseBodyDataWordsInfoPositions struct {
	// example:
	//
	// 863
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 43
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizePdfResponseBodyDataWordsInfoPositions) String() string {
	return dara.Prettify(s)
}

func (s RecognizePdfResponseBodyDataWordsInfoPositions) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponseBodyDataWordsInfoPositions) GetX() *int64 {
	return s.X
}

func (s *RecognizePdfResponseBodyDataWordsInfoPositions) GetY() *int64 {
	return s.Y
}

func (s *RecognizePdfResponseBodyDataWordsInfoPositions) SetX(v int64) *RecognizePdfResponseBodyDataWordsInfoPositions {
	s.X = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfoPositions) SetY(v int64) *RecognizePdfResponseBodyDataWordsInfoPositions {
	s.Y = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfoPositions) Validate() error {
	return dara.Validate(s)
}
