// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeIdentityCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeIdentityCardResponseBodyData) *RecognizeIdentityCardResponseBody
	GetData() *RecognizeIdentityCardResponseBodyData
	SetRequestId(v string) *RecognizeIdentityCardResponseBody
	GetRequestId() *string
}

type RecognizeIdentityCardResponseBody struct {
	Data *RecognizeIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D3F5BA69-79C4-46A4-B02B-58C4EEBC4C33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeIdentityCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBody) GetData() *RecognizeIdentityCardResponseBodyData {
	return s.Data
}

func (s *RecognizeIdentityCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeIdentityCardResponseBody) SetData(v *RecognizeIdentityCardResponseBodyData) *RecognizeIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeIdentityCardResponseBody) SetRequestId(v string) *RecognizeIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeIdentityCardResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyData struct {
	BackResult  *RecognizeIdentityCardResponseBodyDataBackResult  `json:"BackResult,omitempty" xml:"BackResult,omitempty" type:"Struct"`
	FrontResult *RecognizeIdentityCardResponseBodyDataFrontResult `json:"FrontResult,omitempty" xml:"FrontResult,omitempty" type:"Struct"`
}

func (s RecognizeIdentityCardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyData) GetBackResult() *RecognizeIdentityCardResponseBodyDataBackResult {
	return s.BackResult
}

func (s *RecognizeIdentityCardResponseBodyData) GetFrontResult() *RecognizeIdentityCardResponseBodyDataFrontResult {
	return s.FrontResult
}

func (s *RecognizeIdentityCardResponseBodyData) SetBackResult(v *RecognizeIdentityCardResponseBodyDataBackResult) *RecognizeIdentityCardResponseBodyData {
	s.BackResult = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyData) SetFrontResult(v *RecognizeIdentityCardResponseBodyDataFrontResult) *RecognizeIdentityCardResponseBodyData {
	s.FrontResult = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyDataBackResult struct {
	// example:
	//
	// 19800101
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Issue   *string `json:"Issue,omitempty" xml:"Issue,omitempty"`
	// example:
	//
	// 19970101
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataBackResult) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataBackResult) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataBackResult) GetEndDate() *string {
	return s.EndDate
}

func (s *RecognizeIdentityCardResponseBodyDataBackResult) GetIssue() *string {
	return s.Issue
}

func (s *RecognizeIdentityCardResponseBodyDataBackResult) GetStartDate() *string {
	return s.StartDate
}

func (s *RecognizeIdentityCardResponseBodyDataBackResult) SetEndDate(v string) *RecognizeIdentityCardResponseBodyDataBackResult {
	s.EndDate = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataBackResult) SetIssue(v string) *RecognizeIdentityCardResponseBodyDataBackResult {
	s.Issue = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataBackResult) SetStartDate(v string) *RecognizeIdentityCardResponseBodyDataBackResult {
	s.StartDate = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataBackResult) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyDataFrontResult struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 19960111
	BirthDate        *string                                                             `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	CardAreas        []*RecognizeIdentityCardResponseBodyDataFrontResultCardAreas        `json:"CardAreas,omitempty" xml:"CardAreas,omitempty" type:"Repeated"`
	FaceRectVertices []*RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices `json:"FaceRectVertices,omitempty" xml:"FaceRectVertices,omitempty" type:"Repeated"`
	FaceRectangle    *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle      `json:"FaceRectangle,omitempty" xml:"FaceRectangle,omitempty" type:"Struct"`
	Gender           *string                                                             `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// 310228199601115411
	IDNumber    *string `json:"IDNumber,omitempty" xml:"IDNumber,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResult) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResult) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetAddress() *string {
	return s.Address
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetBirthDate() *string {
	return s.BirthDate
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetCardAreas() []*RecognizeIdentityCardResponseBodyDataFrontResultCardAreas {
	return s.CardAreas
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetFaceRectVertices() []*RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices {
	return s.FaceRectVertices
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetFaceRectangle() *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle {
	return s.FaceRectangle
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetGender() *string {
	return s.Gender
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetIDNumber() *string {
	return s.IDNumber
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetName() *string {
	return s.Name
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) GetNationality() *string {
	return s.Nationality
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetAddress(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Address = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetBirthDate(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.BirthDate = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetCardAreas(v []*RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.CardAreas = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetFaceRectVertices(v []*RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.FaceRectVertices = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetFaceRectangle(v *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.FaceRectangle = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetGender(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Gender = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetIDNumber(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.IDNumber = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetName(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Name = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetNationality(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Nationality = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyDataFrontResultCardAreas struct {
	// example:
	//
	// 40
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 81
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) GetX() *float32 {
	return s.X
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) GetY() *float32 {
	return s.Y
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas {
	s.X = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas {
	s.Y = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices struct {
	// example:
	//
	// 429.46124267578125
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 164.23321533203125
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) GetX() *float32 {
	return s.X
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) GetY() *float32 {
	return s.Y
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices {
	s.X = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices {
	s.Y = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle struct {
	// example:
	//
	// -87.710586547851562
	Angle  *float32                                                             `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Center *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter `json:"Center,omitempty" xml:"Center,omitempty" type:"Struct"`
	Size   *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize   `json:"Size,omitempty" xml:"Size,omitempty" type:"Struct"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) GetAngle() *float32 {
	return s.Angle
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) GetCenter() *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter {
	return s.Center
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) GetSize() *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize {
	return s.Size
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) SetAngle(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle {
	s.Angle = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) SetCenter(v *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle {
	s.Center = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) SetSize(v *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle {
	s.Size = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter struct {
	// example:
	//
	// 475.59390258789062
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 225.20643615722656
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) GetX() *float32 {
	return s.X
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) GetY() *float32 {
	return s.Y
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter {
	s.X = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter {
	s.Y = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) Validate() error {
	return dara.Validate(s)
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize struct {
	// example:
	//
	// 97.063156127929688
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 118.16333770751953
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) GetHeight() *float32 {
	return s.Height
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) GetWidth() *float32 {
	return s.Width
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) SetHeight(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize {
	s.Height = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) SetWidth(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize {
	s.Width = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) Validate() error {
	return dara.Validate(s)
}
