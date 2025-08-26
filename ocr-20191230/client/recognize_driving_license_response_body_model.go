// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeDrivingLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeDrivingLicenseResponseBodyData) *RecognizeDrivingLicenseResponseBody
	GetData() *RecognizeDrivingLicenseResponseBodyData
	SetRequestId(v string) *RecognizeDrivingLicenseResponseBody
	GetRequestId() *string
}

type RecognizeDrivingLicenseResponseBody struct {
	Data *RecognizeDrivingLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 1DD989C1-4E08-4E04-9D5D-314901E91226
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBody) GetData() *RecognizeDrivingLicenseResponseBodyData {
	return s.Data
}

func (s *RecognizeDrivingLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeDrivingLicenseResponseBody) SetData(v *RecognizeDrivingLicenseResponseBodyData) *RecognizeDrivingLicenseResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetRequestId(v string) *RecognizeDrivingLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeDrivingLicenseResponseBodyData struct {
	BackResult *RecognizeDrivingLicenseResponseBodyDataBackResult `json:"BackResult,omitempty" xml:"BackResult,omitempty" type:"Struct"`
	FaceResult *RecognizeDrivingLicenseResponseBodyDataFaceResult `json:"FaceResult,omitempty" xml:"FaceResult,omitempty" type:"Struct"`
}

func (s RecognizeDrivingLicenseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBodyData) GetBackResult() *RecognizeDrivingLicenseResponseBodyDataBackResult {
	return s.BackResult
}

func (s *RecognizeDrivingLicenseResponseBodyData) GetFaceResult() *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	return s.FaceResult
}

func (s *RecognizeDrivingLicenseResponseBodyData) SetBackResult(v *RecognizeDrivingLicenseResponseBodyDataBackResult) *RecognizeDrivingLicenseResponseBodyData {
	s.BackResult = v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyData) SetFaceResult(v *RecognizeDrivingLicenseResponseBodyDataFaceResult) *RecognizeDrivingLicenseResponseBodyData {
	s.FaceResult = v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeDrivingLicenseResponseBodyDataBackResult struct {
	// example:
	//
	// 300
	ApprovedLoad *string `json:"ApprovedLoad,omitempty" xml:"ApprovedLoad,omitempty"`
	// example:
	//
	// 5
	ApprovedPassengerCapacity *string `json:"ApprovedPassengerCapacity,omitempty" xml:"ApprovedPassengerCapacity,omitempty"`
	// example:
	//
	// -
	EnergyType *string `json:"EnergyType,omitempty" xml:"EnergyType,omitempty"`
	FileNumber *string `json:"FileNumber,omitempty" xml:"FileNumber,omitempty"`
	// example:
	//
	// 2205
	GrossMass        *string `json:"GrossMass,omitempty" xml:"GrossMass,omitempty"`
	InspectionRecord *string `json:"InspectionRecord,omitempty" xml:"InspectionRecord,omitempty"`
	// example:
	//
	// 4945x1845x1480
	OverallDimension *string `json:"OverallDimension,omitempty" xml:"OverallDimension,omitempty"`
	PlateNumber      *string `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	// example:
	//
	// 100
	TractionMass *string `json:"TractionMass,omitempty" xml:"TractionMass,omitempty"`
	// example:
	//
	// 2000
	UnladenMass *string `json:"UnladenMass,omitempty" xml:"UnladenMass,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBodyDataBackResult) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyDataBackResult) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetApprovedLoad() *string {
	return s.ApprovedLoad
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetApprovedPassengerCapacity() *string {
	return s.ApprovedPassengerCapacity
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetEnergyType() *string {
	return s.EnergyType
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetFileNumber() *string {
	return s.FileNumber
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetGrossMass() *string {
	return s.GrossMass
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetInspectionRecord() *string {
	return s.InspectionRecord
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetOverallDimension() *string {
	return s.OverallDimension
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetPlateNumber() *string {
	return s.PlateNumber
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetTractionMass() *string {
	return s.TractionMass
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) GetUnladenMass() *string {
	return s.UnladenMass
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetApprovedLoad(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.ApprovedLoad = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetApprovedPassengerCapacity(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.ApprovedPassengerCapacity = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetEnergyType(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.EnergyType = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetFileNumber(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.FileNumber = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetGrossMass(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.GrossMass = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetInspectionRecord(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.InspectionRecord = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetOverallDimension(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.OverallDimension = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetPlateNumber(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.PlateNumber = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetTractionMass(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.TractionMass = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetUnladenMass(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.UnladenMass = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) Validate() error {
	return dara.Validate(s)
}

type RecognizeDrivingLicenseResponseBodyDataFaceResult struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 111111
	EngineNumber *string `json:"EngineNumber,omitempty" xml:"EngineNumber,omitempty"`
	// example:
	//
	// 20180313
	IssueDate   *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PlateNumber *string `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	// example:
	//
	// 20180312
	RegisterDate *string `json:"RegisterDate,omitempty" xml:"RegisterDate,omitempty"`
	UseCharacter *string `json:"UseCharacter,omitempty" xml:"UseCharacter,omitempty"`
	VehicleType  *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	// example:
	//
	// SSVUDDTT2J2022555
	Vin *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBodyDataFaceResult) String() string {
	return dara.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyDataFaceResult) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetAddress() *string {
	return s.Address
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetEngineNumber() *string {
	return s.EngineNumber
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetIssueDate() *string {
	return s.IssueDate
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetModel() *string {
	return s.Model
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetOwner() *string {
	return s.Owner
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetPlateNumber() *string {
	return s.PlateNumber
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetRegisterDate() *string {
	return s.RegisterDate
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetUseCharacter() *string {
	return s.UseCharacter
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetVehicleType() *string {
	return s.VehicleType
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) GetVin() *string {
	return s.Vin
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetAddress(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Address = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetEngineNumber(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.EngineNumber = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetIssueDate(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.IssueDate = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetModel(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Model = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetOwner(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Owner = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetPlateNumber(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.PlateNumber = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetRegisterDate(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.RegisterDate = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetUseCharacter(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.UseCharacter = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetVehicleType(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.VehicleType = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetVin(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Vin = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) Validate() error {
	return dara.Validate(s)
}
