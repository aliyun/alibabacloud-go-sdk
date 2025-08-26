// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeDriverLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeDriverLicenseResponseBodyData) *RecognizeDriverLicenseResponseBody
	GetData() *RecognizeDriverLicenseResponseBodyData
	SetRequestId(v string) *RecognizeDriverLicenseResponseBody
	GetRequestId() *string
}

type RecognizeDriverLicenseResponseBody struct {
	Data *RecognizeDriverLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 374D8C7E-9ECC-4750-A87F-50571702F175
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDriverLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBody) GetData() *RecognizeDriverLicenseResponseBodyData {
	return s.Data
}

func (s *RecognizeDriverLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeDriverLicenseResponseBody) SetData(v *RecognizeDriverLicenseResponseBodyData) *RecognizeDriverLicenseResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeDriverLicenseResponseBody) SetRequestId(v string) *RecognizeDriverLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeDriverLicenseResponseBodyData struct {
	BackResult *RecognizeDriverLicenseResponseBodyDataBackResult `json:"BackResult,omitempty" xml:"BackResult,omitempty" type:"Struct"`
	FaceResult *RecognizeDriverLicenseResponseBodyDataFaceResult `json:"FaceResult,omitempty" xml:"FaceResult,omitempty" type:"Struct"`
}

func (s RecognizeDriverLicenseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBodyData) GetBackResult() *RecognizeDriverLicenseResponseBodyDataBackResult {
	return s.BackResult
}

func (s *RecognizeDriverLicenseResponseBodyData) GetFaceResult() *RecognizeDriverLicenseResponseBodyDataFaceResult {
	return s.FaceResult
}

func (s *RecognizeDriverLicenseResponseBodyData) SetBackResult(v *RecognizeDriverLicenseResponseBodyDataBackResult) *RecognizeDriverLicenseResponseBodyData {
	s.BackResult = v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyData) SetFaceResult(v *RecognizeDriverLicenseResponseBodyDataFaceResult) *RecognizeDriverLicenseResponseBodyData {
	s.FaceResult = v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeDriverLicenseResponseBodyDataBackResult struct {
	// example:
	//
	// 130601473955
	ArchiveNumber *string `json:"ArchiveNumber,omitempty" xml:"ArchiveNumber,omitempty"`
	// example:
	//
	// 210288898898898888
	CardNumber *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Record     *string `json:"Record,omitempty" xml:"Record,omitempty"`
}

func (s RecognizeDriverLicenseResponseBodyDataBackResult) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBodyDataBackResult) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) GetArchiveNumber() *string {
	return s.ArchiveNumber
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) GetCardNumber() *string {
	return s.CardNumber
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) GetName() *string {
	return s.Name
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) GetRecord() *string {
	return s.Record
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) SetArchiveNumber(v string) *RecognizeDriverLicenseResponseBodyDataBackResult {
	s.ArchiveNumber = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) SetCardNumber(v string) *RecognizeDriverLicenseResponseBodyDataBackResult {
	s.CardNumber = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) SetName(v string) *RecognizeDriverLicenseResponseBodyDataBackResult {
	s.Name = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) SetRecord(v string) *RecognizeDriverLicenseResponseBodyDataBackResult {
	s.Record = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) Validate() error {
	return dara.Validate(s)
}

type RecognizeDriverLicenseResponseBodyDataFaceResult struct {
	Address   *string `json:"Address,omitempty" xml:"Address,omitempty"`
	BirthDate *string `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	// example:
	//
	// 20190201
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Gender  *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// example:
	//
	// 20130208
	IssueDate *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	IssueUnit *string `json:"IssueUnit,omitempty" xml:"IssueUnit,omitempty"`
	// example:
	//
	// 210288898898898888
	LicenseNumber *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality   *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	// example:
	//
	// 20130208
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// C1
	VehicleType *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
}

func (s RecognizeDriverLicenseResponseBodyDataFaceResult) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBodyDataFaceResult) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetAddress() *string {
	return s.Address
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetBirthDate() *string {
	return s.BirthDate
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetEndDate() *string {
	return s.EndDate
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetGender() *string {
	return s.Gender
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetIssueDate() *string {
	return s.IssueDate
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetIssueUnit() *string {
	return s.IssueUnit
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetLicenseNumber() *string {
	return s.LicenseNumber
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetName() *string {
	return s.Name
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetNationality() *string {
	return s.Nationality
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetStartDate() *string {
	return s.StartDate
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) GetVehicleType() *string {
	return s.VehicleType
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetAddress(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.Address = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetBirthDate(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.BirthDate = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetEndDate(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.EndDate = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetGender(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.Gender = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetIssueDate(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.IssueDate = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetIssueUnit(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.IssueUnit = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetLicenseNumber(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.LicenseNumber = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetName(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.Name = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetNationality(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.Nationality = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetStartDate(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.StartDate = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetVehicleType(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.VehicleType = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) Validate() error {
	return dara.Validate(s)
}
