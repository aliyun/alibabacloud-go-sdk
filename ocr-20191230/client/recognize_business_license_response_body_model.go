// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeBusinessLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeBusinessLicenseResponseBodyData) *RecognizeBusinessLicenseResponseBody
	GetData() *RecognizeBusinessLicenseResponseBodyData
	SetRequestId(v string) *RecognizeBusinessLicenseResponseBody
	GetRequestId() *string
}

type RecognizeBusinessLicenseResponseBody struct {
	Data *RecognizeBusinessLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// F34D031B-02BD-4A59-BA35-EE068DD6F6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBody) GetData() *RecognizeBusinessLicenseResponseBodyData {
	return s.Data
}

func (s *RecognizeBusinessLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeBusinessLicenseResponseBody) SetData(v *RecognizeBusinessLicenseResponseBodyData) *RecognizeBusinessLicenseResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeBusinessLicenseResponseBodyData struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 0
	Angle    *float32                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Business *string                                         `json:"Business,omitempty" xml:"Business,omitempty"`
	Capital  *string                                         `json:"Capital,omitempty" xml:"Capital,omitempty"`
	Emblem   *RecognizeBusinessLicenseResponseBodyDataEmblem `json:"Emblem,omitempty" xml:"Emblem,omitempty" type:"Struct"`
	// example:
	//
	// 20150504
	EstablishDate *string                                         `json:"EstablishDate,omitempty" xml:"EstablishDate,omitempty"`
	LegalPerson   *string                                         `json:"LegalPerson,omitempty" xml:"LegalPerson,omitempty"`
	Name          *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	QRCode        *RecognizeBusinessLicenseResponseBodyDataQRCode `json:"QRCode,omitempty" xml:"QRCode,omitempty" type:"Struct"`
	// example:
	//
	// 91500108320423****
	RegisterNumber *string                                        `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Stamp          *RecognizeBusinessLicenseResponseBodyDataStamp `json:"Stamp,omitempty" xml:"Stamp,omitempty" type:"Struct"`
	Title          *RecognizeBusinessLicenseResponseBodyDataTitle `json:"Title,omitempty" xml:"Title,omitempty" type:"Struct"`
	Type           *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 29991231
	ValidPeriod *string `json:"ValidPeriod,omitempty" xml:"ValidPeriod,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetAngle() *float32 {
	return s.Angle
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetBusiness() *string {
	return s.Business
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetCapital() *string {
	return s.Capital
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetEmblem() *RecognizeBusinessLicenseResponseBodyDataEmblem {
	return s.Emblem
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetEstablishDate() *string {
	return s.EstablishDate
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetLegalPerson() *string {
	return s.LegalPerson
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetName() *string {
	return s.Name
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetQRCode() *RecognizeBusinessLicenseResponseBodyDataQRCode {
	return s.QRCode
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetRegisterNumber() *string {
	return s.RegisterNumber
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetStamp() *RecognizeBusinessLicenseResponseBodyDataStamp {
	return s.Stamp
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetTitle() *RecognizeBusinessLicenseResponseBodyDataTitle {
	return s.Title
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetType() *string {
	return s.Type
}

func (s *RecognizeBusinessLicenseResponseBodyData) GetValidPeriod() *string {
	return s.ValidPeriod
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetAddress(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Address = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetAngle(v float32) *RecognizeBusinessLicenseResponseBodyData {
	s.Angle = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetBusiness(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Business = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetCapital(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Capital = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetEmblem(v *RecognizeBusinessLicenseResponseBodyDataEmblem) *RecognizeBusinessLicenseResponseBodyData {
	s.Emblem = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetEstablishDate(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.EstablishDate = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetLegalPerson(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.LegalPerson = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetName(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetQRCode(v *RecognizeBusinessLicenseResponseBodyDataQRCode) *RecognizeBusinessLicenseResponseBodyData {
	s.QRCode = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetRegisterNumber(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.RegisterNumber = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetStamp(v *RecognizeBusinessLicenseResponseBodyDataStamp) *RecognizeBusinessLicenseResponseBodyData {
	s.Stamp = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetTitle(v *RecognizeBusinessLicenseResponseBodyDataTitle) *RecognizeBusinessLicenseResponseBodyData {
	s.Title = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetType(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Type = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetValidPeriod(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.ValidPeriod = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeBusinessLicenseResponseBodyDataEmblem struct {
	// example:
	//
	// 163
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 366
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 8
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 162
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataEmblem) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataEmblem) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) GetHeight() *int32 {
	return s.Height
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) GetLeft() *int32 {
	return s.Left
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) GetTop() *int32 {
	return s.Top
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) GetWidth() *int32 {
	return s.Width
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Left = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) Validate() error {
	return dara.Validate(s)
}

type RecognizeBusinessLicenseResponseBodyDataQRCode struct {
	// example:
	//
	// 132
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 156
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 914
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 126
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataQRCode) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataQRCode) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) GetHeight() *int32 {
	return s.Height
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) GetLeft() *int32 {
	return s.Left
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) GetTop() *int32 {
	return s.Top
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) GetWidth() *int32 {
	return s.Width
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Left = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) Validate() error {
	return dara.Validate(s)
}

type RecognizeBusinessLicenseResponseBodyDataStamp struct {
	// example:
	//
	// 154
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 650
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 1030
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 154
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataStamp) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataStamp) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) GetHeight() *int32 {
	return s.Height
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) GetLeft() *int32 {
	return s.Left
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) GetTop() *int32 {
	return s.Top
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) GetWidth() *int32 {
	return s.Width
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Left = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) Validate() error {
	return dara.Validate(s)
}

type RecognizeBusinessLicenseResponseBodyDataTitle struct {
	// example:
	//
	// 10
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 10
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 10
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 10
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataTitle) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataTitle) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) GetHeight() *int32 {
	return s.Height
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) GetLeft() *int32 {
	return s.Left
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) GetTop() *int32 {
	return s.Top
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) GetWidth() *int32 {
	return s.Width
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Left = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) Validate() error {
	return dara.Validate(s)
}
