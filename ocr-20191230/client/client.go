// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type RecognizeDrivingLicenseRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Side      *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDrivingLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseRequest) SetImageType(v int32) *RecognizeDrivingLicenseRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeDrivingLicenseRequest) SetImageURL(v string) *RecognizeDrivingLicenseRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeDrivingLicenseRequest) SetSide(v string) *RecognizeDrivingLicenseRequest {
	s.Side = &v
	return s
}

type RecognizeDrivingLicenseAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Side           *string   `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDrivingLicenseAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeDrivingLicenseAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeDrivingLicenseAdvanceRequest) SetImageType(v int32) *RecognizeDrivingLicenseAdvanceRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeDrivingLicenseAdvanceRequest) SetSide(v string) *RecognizeDrivingLicenseAdvanceRequest {
	s.Side = &v
	return s
}

type RecognizeDrivingLicenseResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeDrivingLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeDrivingLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBody) SetRequestId(v string) *RecognizeDrivingLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetData(v *RecognizeDrivingLicenseResponseBodyData) *RecognizeDrivingLicenseResponseBody {
	s.Data = v
	return s
}

type RecognizeDrivingLicenseResponseBodyData struct {
	BackResult *RecognizeDrivingLicenseResponseBodyDataBackResult `json:"BackResult,omitempty" xml:"BackResult,omitempty" type:"Struct"`
	FaceResult *RecognizeDrivingLicenseResponseBodyDataFaceResult `json:"FaceResult,omitempty" xml:"FaceResult,omitempty" type:"Struct"`
}

func (s RecognizeDrivingLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBodyData) SetBackResult(v *RecognizeDrivingLicenseResponseBodyDataBackResult) *RecognizeDrivingLicenseResponseBodyData {
	s.BackResult = v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyData) SetFaceResult(v *RecognizeDrivingLicenseResponseBodyDataFaceResult) *RecognizeDrivingLicenseResponseBodyData {
	s.FaceResult = v
	return s
}

type RecognizeDrivingLicenseResponseBodyDataBackResult struct {
	OverallDimension          *string `json:"OverallDimension,omitempty" xml:"OverallDimension,omitempty"`
	InspectionRecord          *string `json:"InspectionRecord,omitempty" xml:"InspectionRecord,omitempty"`
	UnladenMass               *string `json:"UnladenMass,omitempty" xml:"UnladenMass,omitempty"`
	FileNumber                *string `json:"FileNumber,omitempty" xml:"FileNumber,omitempty"`
	TractionMass              *string `json:"TractionMass,omitempty" xml:"TractionMass,omitempty"`
	GrossMass                 *string `json:"GrossMass,omitempty" xml:"GrossMass,omitempty"`
	PlateNumber               *string `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	ApprovedPassengerCapacity *string `json:"ApprovedPassengerCapacity,omitempty" xml:"ApprovedPassengerCapacity,omitempty"`
	EnergyType                *string `json:"EnergyType,omitempty" xml:"EnergyType,omitempty"`
	ApprovedLoad              *string `json:"ApprovedLoad,omitempty" xml:"ApprovedLoad,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBodyDataBackResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyDataBackResult) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetOverallDimension(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.OverallDimension = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetInspectionRecord(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.InspectionRecord = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetUnladenMass(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.UnladenMass = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetFileNumber(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.FileNumber = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetTractionMass(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.TractionMass = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetGrossMass(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.GrossMass = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetPlateNumber(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.PlateNumber = &v
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

func (s *RecognizeDrivingLicenseResponseBodyDataBackResult) SetApprovedLoad(v string) *RecognizeDrivingLicenseResponseBodyDataBackResult {
	s.ApprovedLoad = &v
	return s
}

type RecognizeDrivingLicenseResponseBodyDataFaceResult struct {
	IssueDate    *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	Model        *string `json:"Model,omitempty" xml:"Model,omitempty"`
	VehicleType  *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	Owner        *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	EngineNumber *string `json:"EngineNumber,omitempty" xml:"EngineNumber,omitempty"`
	PlateNumber  *string `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	UseCharacter *string `json:"UseCharacter,omitempty" xml:"UseCharacter,omitempty"`
	Vin          *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
	RegisterDate *string `json:"RegisterDate,omitempty" xml:"RegisterDate,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBodyDataFaceResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyDataFaceResult) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetIssueDate(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.IssueDate = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetModel(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Model = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetVehicleType(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.VehicleType = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetOwner(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Owner = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetEngineNumber(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.EngineNumber = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetPlateNumber(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.PlateNumber = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetAddress(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Address = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetUseCharacter(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.UseCharacter = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetVin(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.Vin = &v
	return s
}

func (s *RecognizeDrivingLicenseResponseBodyDataFaceResult) SetRegisterDate(v string) *RecognizeDrivingLicenseResponseBodyDataFaceResult {
	s.RegisterDate = &v
	return s
}

type RecognizeDrivingLicenseResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeDrivingLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeDrivingLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponse) SetHeaders(v map[string]*string) *RecognizeDrivingLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetBody(v *RecognizeDrivingLicenseResponseBody) *RecognizeDrivingLicenseResponse {
	s.Body = v
	return s
}

type RecognizeChinapassportRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeChinapassportRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinapassportRequest) GoString() string {
	return s.String()
}

func (s *RecognizeChinapassportRequest) SetImageURL(v string) *RecognizeChinapassportRequest {
	s.ImageURL = &v
	return s
}

type RecognizeChinapassportAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s RecognizeChinapassportAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinapassportAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeChinapassportAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeChinapassportAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeChinapassportResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeChinapassportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeChinapassportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinapassportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeChinapassportResponseBody) SetRequestId(v string) *RecognizeChinapassportResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeChinapassportResponseBody) SetData(v *RecognizeChinapassportResponseBodyData) *RecognizeChinapassportResponseBody {
	s.Data = v
	return s
}

type RecognizeChinapassportResponseBodyData struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Sex            *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
	Authority      *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	IssuePlaceRaw  *string `json:"IssuePlaceRaw,omitempty" xml:"IssuePlaceRaw,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	LineOne        *string `json:"LineOne,omitempty" xml:"LineOne,omitempty"`
	ExpiryDate     *string `json:"ExpiryDate,omitempty" xml:"ExpiryDate,omitempty"`
	BirthPlace     *string `json:"BirthPlace,omitempty" xml:"BirthPlace,omitempty"`
	PassportNo     *string `json:"PassportNo,omitempty" xml:"PassportNo,omitempty"`
	BirthPlaceRaw  *string `json:"BirthPlaceRaw,omitempty" xml:"BirthPlaceRaw,omitempty"`
	IssueDate      *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	SourceCountry  *string `json:"SourceCountry,omitempty" xml:"SourceCountry,omitempty"`
	BirthDate      *string `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	NameChineseRaw *string `json:"NameChineseRaw,omitempty" xml:"NameChineseRaw,omitempty"`
	IssuePlace     *string `json:"IssuePlace,omitempty" xml:"IssuePlace,omitempty"`
	NameChinese    *string `json:"NameChinese,omitempty" xml:"NameChinese,omitempty"`
	LineZero       *string `json:"LineZero,omitempty" xml:"LineZero,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	BirthDay       *string `json:"BirthDay,omitempty" xml:"BirthDay,omitempty"`
	ExpiryDay      *string `json:"ExpiryDay,omitempty" xml:"ExpiryDay,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonId       *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
}

func (s RecognizeChinapassportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinapassportResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeChinapassportResponseBodyData) SetType(v string) *RecognizeChinapassportResponseBodyData {
	s.Type = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetSex(v string) *RecognizeChinapassportResponseBodyData {
	s.Sex = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetAuthority(v string) *RecognizeChinapassportResponseBodyData {
	s.Authority = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetIssuePlaceRaw(v string) *RecognizeChinapassportResponseBodyData {
	s.IssuePlaceRaw = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetSuccess(v bool) *RecognizeChinapassportResponseBodyData {
	s.Success = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetLineOne(v string) *RecognizeChinapassportResponseBodyData {
	s.LineOne = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetExpiryDate(v string) *RecognizeChinapassportResponseBodyData {
	s.ExpiryDate = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthPlace(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthPlace = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetPassportNo(v string) *RecognizeChinapassportResponseBodyData {
	s.PassportNo = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthPlaceRaw(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthPlaceRaw = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetIssueDate(v string) *RecognizeChinapassportResponseBodyData {
	s.IssueDate = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetSourceCountry(v string) *RecognizeChinapassportResponseBodyData {
	s.SourceCountry = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthDate(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthDate = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetNameChineseRaw(v string) *RecognizeChinapassportResponseBodyData {
	s.NameChineseRaw = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetIssuePlace(v string) *RecognizeChinapassportResponseBodyData {
	s.IssuePlace = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetNameChinese(v string) *RecognizeChinapassportResponseBodyData {
	s.NameChinese = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetLineZero(v string) *RecognizeChinapassportResponseBodyData {
	s.LineZero = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetCountry(v string) *RecognizeChinapassportResponseBodyData {
	s.Country = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthDay(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthDay = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetExpiryDay(v string) *RecognizeChinapassportResponseBodyData {
	s.ExpiryDay = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetName(v string) *RecognizeChinapassportResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetPersonId(v string) *RecognizeChinapassportResponseBodyData {
	s.PersonId = &v
	return s
}

type RecognizeChinapassportResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeChinapassportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeChinapassportResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinapassportResponse) GoString() string {
	return s.String()
}

func (s *RecognizeChinapassportResponse) SetHeaders(v map[string]*string) *RecognizeChinapassportResponse {
	s.Headers = v
	return s
}

func (s *RecognizeChinapassportResponse) SetBody(v *RecognizeChinapassportResponseBody) *RecognizeChinapassportResponse {
	s.Body = v
	return s
}

type TrimDocumentRequest struct {
	FileURL    *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	FileType   *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	Async      *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s TrimDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentRequest) GoString() string {
	return s.String()
}

func (s *TrimDocumentRequest) SetFileURL(v string) *TrimDocumentRequest {
	s.FileURL = &v
	return s
}

func (s *TrimDocumentRequest) SetFileType(v string) *TrimDocumentRequest {
	s.FileType = &v
	return s
}

func (s *TrimDocumentRequest) SetOutputType(v string) *TrimDocumentRequest {
	s.OutputType = &v
	return s
}

func (s *TrimDocumentRequest) SetAsync(v bool) *TrimDocumentRequest {
	s.Async = &v
	return s
}

type TrimDocumentAdvanceRequest struct {
	FileURLObject io.Reader `json:"FileURLObject,omitempty" xml:"FileURLObject,omitempty" require:"true"`
	FileType      *string   `json:"FileType,omitempty" xml:"FileType,omitempty"`
	OutputType    *string   `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	Async         *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s TrimDocumentAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TrimDocumentAdvanceRequest) SetFileURLObject(v io.Reader) *TrimDocumentAdvanceRequest {
	s.FileURLObject = v
	return s
}

func (s *TrimDocumentAdvanceRequest) SetFileType(v string) *TrimDocumentAdvanceRequest {
	s.FileType = &v
	return s
}

func (s *TrimDocumentAdvanceRequest) SetOutputType(v string) *TrimDocumentAdvanceRequest {
	s.OutputType = &v
	return s
}

func (s *TrimDocumentAdvanceRequest) SetAsync(v bool) *TrimDocumentAdvanceRequest {
	s.Async = &v
	return s
}

type TrimDocumentResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *TrimDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s TrimDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *TrimDocumentResponseBody) SetRequestId(v string) *TrimDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrimDocumentResponseBody) SetData(v *TrimDocumentResponseBodyData) *TrimDocumentResponseBody {
	s.Data = v
	return s
}

type TrimDocumentResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s TrimDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *TrimDocumentResponseBodyData) SetContent(v string) *TrimDocumentResponseBodyData {
	s.Content = &v
	return s
}

type TrimDocumentResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TrimDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TrimDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentResponse) GoString() string {
	return s.String()
}

func (s *TrimDocumentResponse) SetHeaders(v map[string]*string) *TrimDocumentResponse {
	s.Headers = v
	return s
}

func (s *TrimDocumentResponse) SetBody(v *TrimDocumentResponseBody) *TrimDocumentResponse {
	s.Body = v
	return s
}

type RecognizeTableRequest struct {
	ImageType       *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutputFormat    *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	UseFinanceModel *bool   `json:"UseFinanceModel,omitempty" xml:"UseFinanceModel,omitempty"`
	AssureDirection *bool   `json:"AssureDirection,omitempty" xml:"AssureDirection,omitempty"`
	HasLine         *bool   `json:"HasLine,omitempty" xml:"HasLine,omitempty"`
	SkipDetection   *bool   `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
}

func (s RecognizeTableRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableRequest) SetImageType(v int32) *RecognizeTableRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeTableRequest) SetImageURL(v string) *RecognizeTableRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeTableRequest) SetOutputFormat(v string) *RecognizeTableRequest {
	s.OutputFormat = &v
	return s
}

func (s *RecognizeTableRequest) SetUseFinanceModel(v bool) *RecognizeTableRequest {
	s.UseFinanceModel = &v
	return s
}

func (s *RecognizeTableRequest) SetAssureDirection(v bool) *RecognizeTableRequest {
	s.AssureDirection = &v
	return s
}

func (s *RecognizeTableRequest) SetHasLine(v bool) *RecognizeTableRequest {
	s.HasLine = &v
	return s
}

func (s *RecognizeTableRequest) SetSkipDetection(v bool) *RecognizeTableRequest {
	s.SkipDetection = &v
	return s
}

type RecognizeTableAdvanceRequest struct {
	ImageURLObject  io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType       *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	OutputFormat    *string   `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	UseFinanceModel *bool     `json:"UseFinanceModel,omitempty" xml:"UseFinanceModel,omitempty"`
	AssureDirection *bool     `json:"AssureDirection,omitempty" xml:"AssureDirection,omitempty"`
	HasLine         *bool     `json:"HasLine,omitempty" xml:"HasLine,omitempty"`
	SkipDetection   *bool     `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
}

func (s RecognizeTableAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTableAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetImageType(v int32) *RecognizeTableAdvanceRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetOutputFormat(v string) *RecognizeTableAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetUseFinanceModel(v bool) *RecognizeTableAdvanceRequest {
	s.UseFinanceModel = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetAssureDirection(v bool) *RecognizeTableAdvanceRequest {
	s.AssureDirection = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetHasLine(v bool) *RecognizeTableAdvanceRequest {
	s.HasLine = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetSkipDetection(v bool) *RecognizeTableAdvanceRequest {
	s.SkipDetection = &v
	return s
}

type RecognizeTableResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBody) SetRequestId(v string) *RecognizeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTableResponseBody) SetData(v *RecognizeTableResponseBodyData) *RecognizeTableResponseBody {
	s.Data = v
	return s
}

type RecognizeTableResponseBodyData struct {
	FileContent *string                                 `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	Tables      []*RecognizeTableResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s RecognizeTableResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyData) SetFileContent(v string) *RecognizeTableResponseBodyData {
	s.FileContent = &v
	return s
}

func (s *RecognizeTableResponseBodyData) SetTables(v []*RecognizeTableResponseBodyDataTables) *RecognizeTableResponseBodyData {
	s.Tables = v
	return s
}

type RecognizeTableResponseBodyDataTables struct {
	Head      []*string                                        `json:"Head,omitempty" xml:"Head,omitempty" type:"Repeated"`
	Tail      []*string                                        `json:"Tail,omitempty" xml:"Tail,omitempty" type:"Repeated"`
	TableRows []*RecognizeTableResponseBodyDataTablesTableRows `json:"TableRows,omitempty" xml:"TableRows,omitempty" type:"Repeated"`
}

func (s RecognizeTableResponseBodyDataTables) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyDataTables) SetHead(v []*string) *RecognizeTableResponseBodyDataTables {
	s.Head = v
	return s
}

func (s *RecognizeTableResponseBodyDataTables) SetTail(v []*string) *RecognizeTableResponseBodyDataTables {
	s.Tail = v
	return s
}

func (s *RecognizeTableResponseBodyDataTables) SetTableRows(v []*RecognizeTableResponseBodyDataTablesTableRows) *RecognizeTableResponseBodyDataTables {
	s.TableRows = v
	return s
}

type RecognizeTableResponseBodyDataTablesTableRows struct {
	TableColumns []*RecognizeTableResponseBodyDataTablesTableRowsTableColumns `json:"TableColumns,omitempty" xml:"TableColumns,omitempty" type:"Repeated"`
}

func (s RecognizeTableResponseBodyDataTablesTableRows) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponseBodyDataTablesTableRows) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyDataTablesTableRows) SetTableColumns(v []*RecognizeTableResponseBodyDataTablesTableRowsTableColumns) *RecognizeTableResponseBodyDataTablesTableRows {
	s.TableColumns = v
	return s
}

type RecognizeTableResponseBodyDataTablesTableRowsTableColumns struct {
	EndRow      *int32    `json:"EndRow,omitempty" xml:"EndRow,omitempty"`
	EndColumn   *int32    `json:"EndColumn,omitempty" xml:"EndColumn,omitempty"`
	Width       *int32    `json:"Width,omitempty" xml:"Width,omitempty"`
	Height      *int32    `json:"Height,omitempty" xml:"Height,omitempty"`
	Texts       []*string `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	StartRow    *int32    `json:"StartRow,omitempty" xml:"StartRow,omitempty"`
	StartColumn *int32    `json:"StartColumn,omitempty" xml:"StartColumn,omitempty"`
}

func (s RecognizeTableResponseBodyDataTablesTableRowsTableColumns) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetEndRow(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.EndRow = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetEndColumn(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.EndColumn = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetWidth(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Width = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetHeight(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Height = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetTexts(v []*string) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Texts = v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetStartRow(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.StartRow = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetStartColumn(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.StartColumn = &v
	return s
}

type RecognizeTableResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTableResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponse) SetHeaders(v map[string]*string) *RecognizeTableResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTableResponse) SetBody(v *RecognizeTableResponseBody) *RecognizeTableResponse {
	s.Body = v
	return s
}

type RecognizeIdentityCardRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Side      *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardRequest) SetImageType(v int32) *RecognizeIdentityCardRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeIdentityCardRequest) SetImageURL(v string) *RecognizeIdentityCardRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeIdentityCardRequest) SetSide(v string) *RecognizeIdentityCardRequest {
	s.Side = &v
	return s
}

type RecognizeIdentityCardAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Side           *string   `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeIdentityCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeIdentityCardAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeIdentityCardAdvanceRequest) SetImageType(v int32) *RecognizeIdentityCardAdvanceRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeIdentityCardAdvanceRequest) SetSide(v string) *RecognizeIdentityCardAdvanceRequest {
	s.Side = &v
	return s
}

type RecognizeIdentityCardResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBody) SetRequestId(v string) *RecognizeIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeIdentityCardResponseBody) SetData(v *RecognizeIdentityCardResponseBodyData) *RecognizeIdentityCardResponseBody {
	s.Data = v
	return s
}

type RecognizeIdentityCardResponseBodyData struct {
	BackResult  *RecognizeIdentityCardResponseBodyDataBackResult  `json:"BackResult,omitempty" xml:"BackResult,omitempty" type:"Struct"`
	FrontResult *RecognizeIdentityCardResponseBodyDataFrontResult `json:"FrontResult,omitempty" xml:"FrontResult,omitempty" type:"Struct"`
}

func (s RecognizeIdentityCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyData) SetBackResult(v *RecognizeIdentityCardResponseBodyDataBackResult) *RecognizeIdentityCardResponseBodyData {
	s.BackResult = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyData) SetFrontResult(v *RecognizeIdentityCardResponseBodyDataFrontResult) *RecognizeIdentityCardResponseBodyData {
	s.FrontResult = v
	return s
}

type RecognizeIdentityCardResponseBodyDataBackResult struct {
	EndDate   *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Issue     *string `json:"Issue,omitempty" xml:"Issue,omitempty"`
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataBackResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataBackResult) GoString() string {
	return s.String()
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

type RecognizeIdentityCardResponseBodyDataFrontResult struct {
	FaceRectangle    *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle      `json:"FaceRectangle,omitempty" xml:"FaceRectangle,omitempty" type:"Struct"`
	BirthDate        *string                                                             `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	Gender           *string                                                             `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Address          *string                                                             `json:"Address,omitempty" xml:"Address,omitempty"`
	FaceRectVertices []*RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices `json:"FaceRectVertices,omitempty" xml:"FaceRectVertices,omitempty" type:"Repeated"`
	CardAreas        []*RecognizeIdentityCardResponseBodyDataFrontResultCardAreas        `json:"CardAreas,omitempty" xml:"CardAreas,omitempty" type:"Repeated"`
	Nationality      *string                                                             `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	Name             *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	IDNumber         *string                                                             `json:"IDNumber,omitempty" xml:"IDNumber,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResult) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetFaceRectangle(v *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.FaceRectangle = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetBirthDate(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.BirthDate = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetGender(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Gender = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetAddress(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Address = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetFaceRectVertices(v []*RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.FaceRectVertices = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetCardAreas(v []*RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.CardAreas = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetNationality(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Nationality = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetName(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.Name = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResult) SetIDNumber(v string) *RecognizeIdentityCardResponseBodyDataFrontResult {
	s.IDNumber = &v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle struct {
	Size   *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize   `json:"Size,omitempty" xml:"Size,omitempty" type:"Struct"`
	Angle  *float32                                                             `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Center *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter `json:"Center,omitempty" xml:"Center,omitempty" type:"Struct"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) SetSize(v *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle {
	s.Size = v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) SetAngle(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle {
	s.Angle = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) SetCenter(v *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle {
	s.Center = v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize struct {
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) SetWidth(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize {
	s.Width = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) SetHeight(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize {
	s.Height = &v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter {
	s.Y = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter {
	s.X = &v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices {
	s.Y = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices {
	s.X = &v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultCardAreas struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas {
	s.Y = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas {
	s.X = &v
	return s
}

type RecognizeIdentityCardResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeIdentityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIdentityCardResponse) SetBody(v *RecognizeIdentityCardResponseBody) *RecognizeIdentityCardResponse {
	s.Body = v
	return s
}

type RecognizeBusinessLicenseRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseRequest) SetImageType(v int32) *RecognizeBusinessLicenseRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeBusinessLicenseRequest) SetImageURL(v string) *RecognizeBusinessLicenseRequest {
	s.ImageURL = &v
	return s
}

type RecognizeBusinessLicenseAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeBusinessLicenseAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeBusinessLicenseAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeBusinessLicenseAdvanceRequest) SetImageType(v int32) *RecognizeBusinessLicenseAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeBusinessLicenseResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeBusinessLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetData(v *RecognizeBusinessLicenseResponseBodyData) *RecognizeBusinessLicenseResponseBody {
	s.Data = v
	return s
}

type RecognizeBusinessLicenseResponseBodyData struct {
	Type           *string                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	Stamp          *RecognizeBusinessLicenseResponseBodyDataStamp  `json:"Stamp,omitempty" xml:"Stamp,omitempty" type:"Struct"`
	EstablishDate  *string                                         `json:"EstablishDate,omitempty" xml:"EstablishDate,omitempty"`
	ValidPeriod    *string                                         `json:"ValidPeriod,omitempty" xml:"ValidPeriod,omitempty"`
	Business       *string                                         `json:"Business,omitempty" xml:"Business,omitempty"`
	Angle          *float32                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	RegisterNumber *string                                         `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Address        *string                                         `json:"Address,omitempty" xml:"Address,omitempty"`
	Capital        *string                                         `json:"Capital,omitempty" xml:"Capital,omitempty"`
	Title          *RecognizeBusinessLicenseResponseBodyDataTitle  `json:"Title,omitempty" xml:"Title,omitempty" type:"Struct"`
	Emblem         *RecognizeBusinessLicenseResponseBodyDataEmblem `json:"Emblem,omitempty" xml:"Emblem,omitempty" type:"Struct"`
	Name           *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	QRCode         *RecognizeBusinessLicenseResponseBodyDataQRCode `json:"QRCode,omitempty" xml:"QRCode,omitempty" type:"Struct"`
	LegalPerson    *string                                         `json:"LegalPerson,omitempty" xml:"LegalPerson,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetType(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Type = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetStamp(v *RecognizeBusinessLicenseResponseBodyDataStamp) *RecognizeBusinessLicenseResponseBodyData {
	s.Stamp = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetEstablishDate(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.EstablishDate = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetValidPeriod(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.ValidPeriod = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetBusiness(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Business = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetAngle(v float32) *RecognizeBusinessLicenseResponseBodyData {
	s.Angle = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetRegisterNumber(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.RegisterNumber = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetAddress(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Address = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetCapital(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.Capital = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetTitle(v *RecognizeBusinessLicenseResponseBodyDataTitle) *RecognizeBusinessLicenseResponseBodyData {
	s.Title = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyData) SetEmblem(v *RecognizeBusinessLicenseResponseBodyDataEmblem) *RecognizeBusinessLicenseResponseBodyData {
	s.Emblem = v
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

func (s *RecognizeBusinessLicenseResponseBodyData) SetLegalPerson(v string) *RecognizeBusinessLicenseResponseBodyData {
	s.LegalPerson = &v
	return s
}

type RecognizeBusinessLicenseResponseBodyDataStamp struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataStamp) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataStamp) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataStamp) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataStamp {
	s.Left = &v
	return s
}

type RecognizeBusinessLicenseResponseBodyDataTitle struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataTitle) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataTitle) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataTitle) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataTitle {
	s.Left = &v
	return s
}

type RecognizeBusinessLicenseResponseBodyDataEmblem struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataEmblem) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataEmblem) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataEmblem) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataEmblem {
	s.Left = &v
	return s
}

type RecognizeBusinessLicenseResponseBodyDataQRCode struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataQRCode) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataQRCode) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetTop(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Top = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetWidth(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Width = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetHeight(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Height = &v
	return s
}

func (s *RecognizeBusinessLicenseResponseBodyDataQRCode) SetLeft(v int32) *RecognizeBusinessLicenseResponseBodyDataQRCode {
	s.Left = &v
	return s
}

type RecognizeBusinessLicenseResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBusinessLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponse) SetHeaders(v map[string]*string) *RecognizeBusinessLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetBody(v *RecognizeBusinessLicenseResponseBody) *RecognizeBusinessLicenseResponse {
	s.Body = v
	return s
}

type RecognizeBankCardRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBankCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardRequest) SetImageType(v int32) *RecognizeBankCardRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeBankCardRequest) SetImageURL(v string) *RecognizeBankCardRequest {
	s.ImageURL = &v
	return s
}

type RecognizeBankCardAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeBankCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeBankCardAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeBankCardAdvanceRequest) SetImageType(v int32) *RecognizeBankCardAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeBankCardResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeBankCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeBankCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBody) SetRequestId(v string) *RecognizeBankCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBankCardResponseBody) SetData(v *RecognizeBankCardResponseBodyData) *RecognizeBankCardResponseBody {
	s.Data = v
	return s
}

type RecognizeBankCardResponseBodyData struct {
	CardNumber *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ValidDate  *string `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
	BankName   *string `json:"BankName,omitempty" xml:"BankName,omitempty"`
}

func (s RecognizeBankCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBodyData) SetCardNumber(v string) *RecognizeBankCardResponseBodyData {
	s.CardNumber = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) SetValidDate(v string) *RecognizeBankCardResponseBodyData {
	s.ValidDate = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) SetBankName(v string) *RecognizeBankCardResponseBodyData {
	s.BankName = &v
	return s
}

type RecognizeBankCardResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBankCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBankCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponse) SetHeaders(v map[string]*string) *RecognizeBankCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankCardResponse) SetBody(v *RecognizeBankCardResponseBody) *RecognizeBankCardResponse {
	s.Body = v
	return s
}

type RecognizeVerificationcodeRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVerificationcodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVerificationcodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVerificationcodeRequest) SetImageURL(v string) *RecognizeVerificationcodeRequest {
	s.ImageURL = &v
	return s
}

type RecognizeVerificationcodeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s RecognizeVerificationcodeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVerificationcodeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVerificationcodeAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeVerificationcodeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeVerificationcodeResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeVerificationcodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeVerificationcodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVerificationcodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVerificationcodeResponseBody) SetRequestId(v string) *RecognizeVerificationcodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVerificationcodeResponseBody) SetData(v *RecognizeVerificationcodeResponseBodyData) *RecognizeVerificationcodeResponseBody {
	s.Data = v
	return s
}

type RecognizeVerificationcodeResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RecognizeVerificationcodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVerificationcodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVerificationcodeResponseBodyData) SetContent(v string) *RecognizeVerificationcodeResponseBodyData {
	s.Content = &v
	return s
}

type RecognizeVerificationcodeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVerificationcodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVerificationcodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVerificationcodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVerificationcodeResponse) SetHeaders(v map[string]*string) *RecognizeVerificationcodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVerificationcodeResponse) SetBody(v *RecognizeVerificationcodeResponseBody) *RecognizeVerificationcodeResponse {
	s.Body = v
	return s
}

type RecognizeAccountPageRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeAccountPageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageRequest) SetImageType(v int32) *RecognizeAccountPageRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeAccountPageRequest) SetImageURL(v string) *RecognizeAccountPageRequest {
	s.ImageURL = &v
	return s
}

type RecognizeAccountPageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeAccountPageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeAccountPageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeAccountPageAdvanceRequest) SetImageType(v int32) *RecognizeAccountPageAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeAccountPageResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeAccountPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeAccountPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBody) SetRequestId(v string) *RecognizeAccountPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeAccountPageResponseBody) SetData(v *RecognizeAccountPageResponseBodyData) *RecognizeAccountPageResponseBody {
	s.Data = v
	return s
}

type RecognizeAccountPageResponseBodyData struct {
	Gender              *string                                                    `json:"Gender,omitempty" xml:"Gender,omitempty"`
	TitleArea           *RecognizeAccountPageResponseBodyDataTitleArea             `json:"TitleArea,omitempty" xml:"TitleArea,omitempty" type:"Struct"`
	Angle               *float32                                                   `json:"Angle,omitempty" xml:"Angle,omitempty"`
	RegisterStampAreas  []*RecognizeAccountPageResponseBodyDataRegisterStampAreas  `json:"RegisterStampAreas,omitempty" xml:"RegisterStampAreas,omitempty" type:"Repeated"`
	Nationality         *string                                                    `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	BirthPlace          *string                                                    `json:"BirthPlace,omitempty" xml:"BirthPlace,omitempty"`
	OtherStampAreas     []*RecognizeAccountPageResponseBodyDataOtherStampAreas     `json:"OtherStampAreas,omitempty" xml:"OtherStampAreas,omitempty" type:"Repeated"`
	UndertakeStampAreas []*RecognizeAccountPageResponseBodyDataUndertakeStampAreas `json:"UndertakeStampAreas,omitempty" xml:"UndertakeStampAreas,omitempty" type:"Repeated"`
	BirthDate           *string                                                    `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	Relation            *string                                                    `json:"Relation,omitempty" xml:"Relation,omitempty"`
	NativePlace         *string                                                    `json:"NativePlace,omitempty" xml:"NativePlace,omitempty"`
	Name                *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	IDNumber            *string                                                    `json:"IDNumber,omitempty" xml:"IDNumber,omitempty"`
	InvalidStampAreas   []*RecognizeAccountPageResponseBodyDataInvalidStampAreas   `json:"InvalidStampAreas,omitempty" xml:"InvalidStampAreas,omitempty" type:"Repeated"`
}

func (s RecognizeAccountPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyData) SetGender(v string) *RecognizeAccountPageResponseBodyData {
	s.Gender = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetTitleArea(v *RecognizeAccountPageResponseBodyDataTitleArea) *RecognizeAccountPageResponseBodyData {
	s.TitleArea = v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetAngle(v float32) *RecognizeAccountPageResponseBodyData {
	s.Angle = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetRegisterStampAreas(v []*RecognizeAccountPageResponseBodyDataRegisterStampAreas) *RecognizeAccountPageResponseBodyData {
	s.RegisterStampAreas = v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetNationality(v string) *RecognizeAccountPageResponseBodyData {
	s.Nationality = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetBirthPlace(v string) *RecognizeAccountPageResponseBodyData {
	s.BirthPlace = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetOtherStampAreas(v []*RecognizeAccountPageResponseBodyDataOtherStampAreas) *RecognizeAccountPageResponseBodyData {
	s.OtherStampAreas = v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetUndertakeStampAreas(v []*RecognizeAccountPageResponseBodyDataUndertakeStampAreas) *RecognizeAccountPageResponseBodyData {
	s.UndertakeStampAreas = v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetBirthDate(v string) *RecognizeAccountPageResponseBodyData {
	s.BirthDate = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetRelation(v string) *RecognizeAccountPageResponseBodyData {
	s.Relation = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetNativePlace(v string) *RecognizeAccountPageResponseBodyData {
	s.NativePlace = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetName(v string) *RecognizeAccountPageResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetIDNumber(v string) *RecognizeAccountPageResponseBodyData {
	s.IDNumber = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetInvalidStampAreas(v []*RecognizeAccountPageResponseBodyDataInvalidStampAreas) *RecognizeAccountPageResponseBodyData {
	s.InvalidStampAreas = v
	return s
}

type RecognizeAccountPageResponseBodyDataTitleArea struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataTitleArea) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataTitleArea) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetTop(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Width = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Left = &v
	return s
}

type RecognizeAccountPageResponseBodyDataRegisterStampAreas struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataRegisterStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataRegisterStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Width = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Left = &v
	return s
}

type RecognizeAccountPageResponseBodyDataOtherStampAreas struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataOtherStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataOtherStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Width = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Left = &v
	return s
}

type RecognizeAccountPageResponseBodyDataUndertakeStampAreas struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataUndertakeStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataUndertakeStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Width = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Left = &v
	return s
}

type RecognizeAccountPageResponseBodyDataInvalidStampAreas struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataInvalidStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataInvalidStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Width = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Left = &v
	return s
}

type RecognizeAccountPageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeAccountPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeAccountPageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponse) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponse) SetHeaders(v map[string]*string) *RecognizeAccountPageResponse {
	s.Headers = v
	return s
}

func (s *RecognizeAccountPageResponse) SetBody(v *RecognizeAccountPageResponseBody) *RecognizeAccountPageResponse {
	s.Body = v
	return s
}

type RecognizeCharacterRequest struct {
	ImageType         *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL          *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	MinHeight         *int32  `json:"MinHeight,omitempty" xml:"MinHeight,omitempty"`
	OutputProbability *bool   `json:"OutputProbability,omitempty" xml:"OutputProbability,omitempty"`
}

func (s RecognizeCharacterRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterRequest) SetImageType(v int32) *RecognizeCharacterRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeCharacterRequest) SetImageURL(v string) *RecognizeCharacterRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeCharacterRequest) SetMinHeight(v int32) *RecognizeCharacterRequest {
	s.MinHeight = &v
	return s
}

func (s *RecognizeCharacterRequest) SetOutputProbability(v bool) *RecognizeCharacterRequest {
	s.OutputProbability = &v
	return s
}

type RecognizeCharacterAdvanceRequest struct {
	ImageURLObject    io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType         *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	MinHeight         *int32    `json:"MinHeight,omitempty" xml:"MinHeight,omitempty"`
	OutputProbability *bool     `json:"OutputProbability,omitempty" xml:"OutputProbability,omitempty"`
}

func (s RecognizeCharacterAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeCharacterAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeCharacterAdvanceRequest) SetImageType(v int32) *RecognizeCharacterAdvanceRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeCharacterAdvanceRequest) SetMinHeight(v int32) *RecognizeCharacterAdvanceRequest {
	s.MinHeight = &v
	return s
}

func (s *RecognizeCharacterAdvanceRequest) SetOutputProbability(v bool) *RecognizeCharacterAdvanceRequest {
	s.OutputProbability = &v
	return s
}

type RecognizeCharacterResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeCharacterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeCharacterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBody) SetRequestId(v string) *RecognizeCharacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeCharacterResponseBody) SetData(v *RecognizeCharacterResponseBodyData) *RecognizeCharacterResponseBody {
	s.Data = v
	return s
}

type RecognizeCharacterResponseBodyData struct {
	Results []*RecognizeCharacterResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RecognizeCharacterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyData) SetResults(v []*RecognizeCharacterResponseBodyDataResults) *RecognizeCharacterResponseBodyData {
	s.Results = v
	return s
}

type RecognizeCharacterResponseBodyDataResults struct {
	TextRectangles *RecognizeCharacterResponseBodyDataResultsTextRectangles `json:"TextRectangles,omitempty" xml:"TextRectangles,omitempty" type:"Struct"`
	Text           *string                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
	Probability    *float32                                                 `json:"Probability,omitempty" xml:"Probability,omitempty"`
}

func (s RecognizeCharacterResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyDataResults) SetTextRectangles(v *RecognizeCharacterResponseBodyDataResultsTextRectangles) *RecognizeCharacterResponseBodyDataResults {
	s.TextRectangles = v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResults) SetText(v string) *RecognizeCharacterResponseBodyDataResults {
	s.Text = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResults) SetProbability(v float32) *RecognizeCharacterResponseBodyDataResults {
	s.Probability = &v
	return s
}

type RecognizeCharacterResponseBodyDataResultsTextRectangles struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Angle  *int32 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectangles) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectangles) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetTop(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Top = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetWidth(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Width = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetHeight(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Height = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetAngle(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Angle = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetLeft(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Left = &v
	return s
}

type RecognizeCharacterResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeCharacterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeCharacterResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponse) SetHeaders(v map[string]*string) *RecognizeCharacterResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCharacterResponse) SetBody(v *RecognizeCharacterResponseBody) *RecognizeCharacterResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	Async *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetAsync(v bool) *GetAsyncJobResultRequest {
	s.Async = &v
	return s
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type RecognizeTakeoutOrderRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTakeoutOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderRequest) SetImageURL(v string) *RecognizeTakeoutOrderRequest {
	s.ImageURL = &v
	return s
}

type RecognizeTakeoutOrderAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s RecognizeTakeoutOrderAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTakeoutOrderAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeTakeoutOrderResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeTakeoutOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeTakeoutOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderResponseBody) SetRequestId(v string) *RecognizeTakeoutOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTakeoutOrderResponseBody) SetData(v *RecognizeTakeoutOrderResponseBodyData) *RecognizeTakeoutOrderResponseBody {
	s.Data = v
	return s
}

type RecognizeTakeoutOrderResponseBodyData struct {
	Elements []*RecognizeTakeoutOrderResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeTakeoutOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderResponseBodyData) SetElements(v []*RecognizeTakeoutOrderResponseBodyDataElements) *RecognizeTakeoutOrderResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeTakeoutOrderResponseBodyDataElements struct {
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RecognizeTakeoutOrderResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetValue(v string) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Value = &v
	return s
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetBoxes(v []*int32) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetScore(v float32) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetName(v string) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Name = &v
	return s
}

type RecognizeTakeoutOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTakeoutOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTakeoutOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderResponse) SetHeaders(v map[string]*string) *RecognizeTakeoutOrderResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTakeoutOrderResponse) SetBody(v *RecognizeTakeoutOrderResponseBody) *RecognizeTakeoutOrderResponse {
	s.Body = v
	return s
}

type RecognizeBusinessCardRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBusinessCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardRequest) SetImageType(v int32) *RecognizeBusinessCardRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeBusinessCardRequest) SetImageURL(v string) *RecognizeBusinessCardRequest {
	s.ImageURL = &v
	return s
}

type RecognizeBusinessCardAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeBusinessCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeBusinessCardAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeBusinessCardAdvanceRequest) SetImageType(v int32) *RecognizeBusinessCardAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeBusinessCardResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeBusinessCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeBusinessCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardResponseBody) SetRequestId(v string) *RecognizeBusinessCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeBusinessCardResponseBody) SetData(v *RecognizeBusinessCardResponseBodyData) *RecognizeBusinessCardResponseBody {
	s.Data = v
	return s
}

type RecognizeBusinessCardResponseBodyData struct {
	Companies          []*string `json:"Companies,omitempty" xml:"Companies,omitempty" type:"Repeated"`
	Titles             []*string `json:"Titles,omitempty" xml:"Titles,omitempty" type:"Repeated"`
	Emails             []*string `json:"Emails,omitempty" xml:"Emails,omitempty" type:"Repeated"`
	Departments        []*string `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	OfficePhoneNumbers []*string `json:"OfficePhoneNumbers,omitempty" xml:"OfficePhoneNumbers,omitempty" type:"Repeated"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	CellPhoneNumbers   []*string `json:"CellPhoneNumbers,omitempty" xml:"CellPhoneNumbers,omitempty" type:"Repeated"`
	Addresses          []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
}

func (s RecognizeBusinessCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardResponseBodyData) SetCompanies(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Companies = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetTitles(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Titles = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetEmails(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Emails = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetDepartments(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Departments = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetOfficePhoneNumbers(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.OfficePhoneNumbers = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetName(v string) *RecognizeBusinessCardResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetCellPhoneNumbers(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.CellPhoneNumbers = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetAddresses(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Addresses = v
	return s
}

type RecognizeBusinessCardResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeBusinessCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeBusinessCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardResponse) SetHeaders(v map[string]*string) *RecognizeBusinessCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBusinessCardResponse) SetBody(v *RecognizeBusinessCardResponseBody) *RecognizeBusinessCardResponse {
	s.Body = v
	return s
}

type DetectCardScreenshotRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectCardScreenshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotRequest) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotRequest) SetImageURL(v string) *DetectCardScreenshotRequest {
	s.ImageURL = &v
	return s
}

type DetectCardScreenshotAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s DetectCardScreenshotAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotAdvanceRequest) SetImageURLObject(v io.Reader) *DetectCardScreenshotAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectCardScreenshotResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectCardScreenshotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectCardScreenshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotResponseBody) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotResponseBody) SetRequestId(v string) *DetectCardScreenshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectCardScreenshotResponseBody) SetData(v *DetectCardScreenshotResponseBodyData) *DetectCardScreenshotResponseBody {
	s.Data = v
	return s
}

type DetectCardScreenshotResponseBodyData struct {
	IsBlur      *bool                                            `json:"IsBlur,omitempty" xml:"IsBlur,omitempty"`
	IsCard      *bool                                            `json:"IsCard,omitempty" xml:"IsCard,omitempty"`
	SpoofResult *DetectCardScreenshotResponseBodyDataSpoofResult `json:"SpoofResult,omitempty" xml:"SpoofResult,omitempty" type:"Struct"`
}

func (s DetectCardScreenshotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotResponseBodyData) SetIsBlur(v bool) *DetectCardScreenshotResponseBodyData {
	s.IsBlur = &v
	return s
}

func (s *DetectCardScreenshotResponseBodyData) SetIsCard(v bool) *DetectCardScreenshotResponseBodyData {
	s.IsCard = &v
	return s
}

func (s *DetectCardScreenshotResponseBodyData) SetSpoofResult(v *DetectCardScreenshotResponseBodyDataSpoofResult) *DetectCardScreenshotResponseBodyData {
	s.SpoofResult = v
	return s
}

type DetectCardScreenshotResponseBodyDataSpoofResult struct {
	IsSpoof   *bool                                                     `json:"IsSpoof,omitempty" xml:"IsSpoof,omitempty"`
	ResultMap *DetectCardScreenshotResponseBodyDataSpoofResultResultMap `json:"ResultMap,omitempty" xml:"ResultMap,omitempty" type:"Struct"`
}

func (s DetectCardScreenshotResponseBodyDataSpoofResult) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotResponseBodyDataSpoofResult) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotResponseBodyDataSpoofResult) SetIsSpoof(v bool) *DetectCardScreenshotResponseBodyDataSpoofResult {
	s.IsSpoof = &v
	return s
}

func (s *DetectCardScreenshotResponseBodyDataSpoofResult) SetResultMap(v *DetectCardScreenshotResponseBodyDataSpoofResultResultMap) *DetectCardScreenshotResponseBodyDataSpoofResult {
	s.ResultMap = v
	return s
}

type DetectCardScreenshotResponseBodyDataSpoofResultResultMap struct {
	ScreenThreshold *float32 `json:"ScreenThreshold,omitempty" xml:"ScreenThreshold,omitempty"`
	ScreenScore     *float32 `json:"ScreenScore,omitempty" xml:"ScreenScore,omitempty"`
}

func (s DetectCardScreenshotResponseBodyDataSpoofResultResultMap) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotResponseBodyDataSpoofResultResultMap) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotResponseBodyDataSpoofResultResultMap) SetScreenThreshold(v float32) *DetectCardScreenshotResponseBodyDataSpoofResultResultMap {
	s.ScreenThreshold = &v
	return s
}

func (s *DetectCardScreenshotResponseBodyDataSpoofResultResultMap) SetScreenScore(v float32) *DetectCardScreenshotResponseBodyDataSpoofResultResultMap {
	s.ScreenScore = &v
	return s
}

type DetectCardScreenshotResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectCardScreenshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectCardScreenshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotResponse) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotResponse) SetHeaders(v map[string]*string) *DetectCardScreenshotResponse {
	s.Headers = v
	return s
}

func (s *DetectCardScreenshotResponse) SetBody(v *DetectCardScreenshotResponseBody) *DetectCardScreenshotResponse {
	s.Body = v
	return s
}

type RecognizeDriverLicenseRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Side      *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDriverLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseRequest) SetImageType(v int32) *RecognizeDriverLicenseRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeDriverLicenseRequest) SetImageURL(v string) *RecognizeDriverLicenseRequest {
	s.ImageURL = &v
	return s
}

func (s *RecognizeDriverLicenseRequest) SetSide(v string) *RecognizeDriverLicenseRequest {
	s.Side = &v
	return s
}

type RecognizeDriverLicenseAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Side           *string   `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDriverLicenseAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeDriverLicenseAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeDriverLicenseAdvanceRequest) SetImageType(v int32) *RecognizeDriverLicenseAdvanceRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeDriverLicenseAdvanceRequest) SetSide(v string) *RecognizeDriverLicenseAdvanceRequest {
	s.Side = &v
	return s
}

type RecognizeDriverLicenseResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeDriverLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeDriverLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBody) SetRequestId(v string) *RecognizeDriverLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBody) SetData(v *RecognizeDriverLicenseResponseBodyData) *RecognizeDriverLicenseResponseBody {
	s.Data = v
	return s
}

type RecognizeDriverLicenseResponseBodyData struct {
	BackResult *RecognizeDriverLicenseResponseBodyDataBackResult `json:"BackResult,omitempty" xml:"BackResult,omitempty" type:"Struct"`
	FaceResult *RecognizeDriverLicenseResponseBodyDataFaceResult `json:"FaceResult,omitempty" xml:"FaceResult,omitempty" type:"Struct"`
}

func (s RecognizeDriverLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBodyData) SetBackResult(v *RecognizeDriverLicenseResponseBodyDataBackResult) *RecognizeDriverLicenseResponseBodyData {
	s.BackResult = v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyData) SetFaceResult(v *RecognizeDriverLicenseResponseBodyDataFaceResult) *RecognizeDriverLicenseResponseBodyData {
	s.FaceResult = v
	return s
}

type RecognizeDriverLicenseResponseBodyDataBackResult struct {
	ArchiveNumber *string `json:"ArchiveNumber,omitempty" xml:"ArchiveNumber,omitempty"`
}

func (s RecognizeDriverLicenseResponseBodyDataBackResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBodyDataBackResult) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBodyDataBackResult) SetArchiveNumber(v string) *RecognizeDriverLicenseResponseBodyDataBackResult {
	s.ArchiveNumber = &v
	return s
}

type RecognizeDriverLicenseResponseBodyDataFaceResult struct {
	VehicleType   *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	IssueDate     *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Gender        *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	LicenseNumber *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RecognizeDriverLicenseResponseBodyDataFaceResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBodyDataFaceResult) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetVehicleType(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.VehicleType = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetIssueDate(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.IssueDate = &v
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

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetAddress(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.Address = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetStartDate(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.StartDate = &v
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

type RecognizeDriverLicenseResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeDriverLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeDriverLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseResponse) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponse) SetHeaders(v map[string]*string) *RecognizeDriverLicenseResponse {
	s.Headers = v
	return s
}

func (s *RecognizeDriverLicenseResponse) SetBody(v *RecognizeDriverLicenseResponseBody) *RecognizeDriverLicenseResponse {
	s.Body = v
	return s
}

type RecognizeLicensePlateRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeLicensePlateRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateRequest) SetImageType(v int32) *RecognizeLicensePlateRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeLicensePlateRequest) SetImageURL(v string) *RecognizeLicensePlateRequest {
	s.ImageURL = &v
	return s
}

type RecognizeLicensePlateAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeLicensePlateAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeLicensePlateAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeLicensePlateAdvanceRequest) SetImageType(v int32) *RecognizeLicensePlateAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeLicensePlateResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeLicensePlateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeLicensePlateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBody) SetRequestId(v string) *RecognizeLicensePlateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeLicensePlateResponseBody) SetData(v *RecognizeLicensePlateResponseBodyData) *RecognizeLicensePlateResponseBody {
	s.Data = v
	return s
}

type RecognizeLicensePlateResponseBodyData struct {
	Plates []*RecognizeLicensePlateResponseBodyDataPlates `json:"Plates,omitempty" xml:"Plates,omitempty" type:"Repeated"`
}

func (s RecognizeLicensePlateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyData) SetPlates(v []*RecognizeLicensePlateResponseBodyDataPlates) *RecognizeLicensePlateResponseBodyData {
	s.Plates = v
	return s
}

type RecognizeLicensePlateResponseBodyDataPlates struct {
	PlateTypeConfidence *float32                                        `json:"PlateTypeConfidence,omitempty" xml:"PlateTypeConfidence,omitempty"`
	PlateType           *string                                         `json:"PlateType,omitempty" xml:"PlateType,omitempty"`
	Confidence          *float32                                        `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	PlateNumber         *string                                         `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	Roi                 *RecognizeLicensePlateResponseBodyDataPlatesRoi `json:"Roi,omitempty" xml:"Roi,omitempty" type:"Struct"`
}

func (s RecognizeLicensePlateResponseBodyDataPlates) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlates) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateTypeConfidence(v float32) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateTypeConfidence = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateType(v string) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateType = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetConfidence(v float32) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Confidence = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateNumber(v string) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateNumber = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetRoi(v *RecognizeLicensePlateResponseBodyDataPlatesRoi) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Roi = v
	return s
}

type RecognizeLicensePlateResponseBodyDataPlatesRoi struct {
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s RecognizeLicensePlateResponseBodyDataPlatesRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlatesRoi) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetW(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.W = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetH(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.H = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetY(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.Y = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetX(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.X = &v
	return s
}

type RecognizeLicensePlateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeLicensePlateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeLicensePlateResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponse) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponse) SetHeaders(v map[string]*string) *RecognizeLicensePlateResponse {
	s.Headers = v
	return s
}

func (s *RecognizeLicensePlateResponse) SetBody(v *RecognizeLicensePlateResponseBody) *RecognizeLicensePlateResponse {
	s.Body = v
	return s
}

type RecognizeStampRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeStampRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampRequest) GoString() string {
	return s.String()
}

func (s *RecognizeStampRequest) SetImageType(v int32) *RecognizeStampRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeStampRequest) SetImageURL(v string) *RecognizeStampRequest {
	s.ImageURL = &v
	return s
}

type RecognizeStampAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeStampAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeStampAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeStampAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeStampAdvanceRequest) SetImageType(v int32) *RecognizeStampAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeStampResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeStampResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeStampResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBody) SetRequestId(v string) *RecognizeStampResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeStampResponseBody) SetData(v *RecognizeStampResponseBodyData) *RecognizeStampResponseBody {
	s.Data = v
	return s
}

type RecognizeStampResponseBodyData struct {
	Results []*RecognizeStampResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RecognizeStampResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBodyData) SetResults(v []*RecognizeStampResponseBodyDataResults) *RecognizeStampResponseBodyData {
	s.Results = v
	return s
}

type RecognizeStampResponseBodyDataResults struct {
	Text        *RecognizeStampResponseBodyDataResultsText          `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
	Roi         *RecognizeStampResponseBodyDataResultsRoi           `json:"Roi,omitempty" xml:"Roi,omitempty" type:"Struct"`
	GeneralText []*RecognizeStampResponseBodyDataResultsGeneralText `json:"GeneralText,omitempty" xml:"GeneralText,omitempty" type:"Repeated"`
}

func (s RecognizeStampResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBodyDataResults) SetText(v *RecognizeStampResponseBodyDataResultsText) *RecognizeStampResponseBodyDataResults {
	s.Text = v
	return s
}

func (s *RecognizeStampResponseBodyDataResults) SetRoi(v *RecognizeStampResponseBodyDataResultsRoi) *RecognizeStampResponseBodyDataResults {
	s.Roi = v
	return s
}

func (s *RecognizeStampResponseBodyDataResults) SetGeneralText(v []*RecognizeStampResponseBodyDataResultsGeneralText) *RecognizeStampResponseBodyDataResults {
	s.GeneralText = v
	return s
}

type RecognizeStampResponseBodyDataResultsText struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Content    *string  `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RecognizeStampResponseBodyDataResultsText) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBodyDataResultsText) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBodyDataResultsText) SetConfidence(v float32) *RecognizeStampResponseBodyDataResultsText {
	s.Confidence = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsText) SetContent(v string) *RecognizeStampResponseBodyDataResultsText {
	s.Content = &v
	return s
}

type RecognizeStampResponseBodyDataResultsRoi struct {
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
}

func (s RecognizeStampResponseBodyDataResultsRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBodyDataResultsRoi) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetTop(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Top = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetWidth(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Width = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetHeight(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Height = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetLeft(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Left = &v
	return s
}

type RecognizeStampResponseBodyDataResultsGeneralText struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Content    *string  `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RecognizeStampResponseBodyDataResultsGeneralText) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBodyDataResultsGeneralText) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBodyDataResultsGeneralText) SetConfidence(v float32) *RecognizeStampResponseBodyDataResultsGeneralText {
	s.Confidence = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsGeneralText) SetContent(v string) *RecognizeStampResponseBodyDataResultsGeneralText {
	s.Content = &v
	return s
}

type RecognizeStampResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeStampResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeStampResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponse) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponse) SetHeaders(v map[string]*string) *RecognizeStampResponse {
	s.Headers = v
	return s
}

func (s *RecognizeStampResponse) SetBody(v *RecognizeStampResponseBody) *RecognizeStampResponse {
	s.Body = v
	return s
}

type RecognizeTaxiInvoiceRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTaxiInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceRequest) SetImageType(v int32) *RecognizeTaxiInvoiceRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeTaxiInvoiceRequest) SetImageURL(v string) *RecognizeTaxiInvoiceRequest {
	s.ImageURL = &v
	return s
}

type RecognizeTaxiInvoiceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeTaxiInvoiceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTaxiInvoiceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTaxiInvoiceAdvanceRequest) SetImageType(v int32) *RecognizeTaxiInvoiceAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeTaxiInvoiceResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeTaxiInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeTaxiInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBody) SetRequestId(v string) *RecognizeTaxiInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetData(v *RecognizeTaxiInvoiceResponseBodyData) *RecognizeTaxiInvoiceResponseBody {
	s.Data = v
	return s
}

type RecognizeTaxiInvoiceResponseBodyData struct {
	Invoices []*RecognizeTaxiInvoiceResponseBodyDataInvoices `json:"Invoices,omitempty" xml:"Invoices,omitempty" type:"Repeated"`
}

func (s RecognizeTaxiInvoiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyData) SetInvoices(v []*RecognizeTaxiInvoiceResponseBodyDataInvoices) *RecognizeTaxiInvoiceResponseBodyData {
	s.Invoices = v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoices struct {
	Items      []*RecognizeTaxiInvoiceResponseBodyDataInvoicesItems    `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RotateType *int32                                                  `json:"RotateType,omitempty" xml:"RotateType,omitempty"`
	InvoiceRoi *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi `json:"InvoiceRoi,omitempty" xml:"InvoiceRoi,omitempty" type:"Struct"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoices) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoices) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetItems(v []*RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.Items = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetRotateType(v int32) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.RotateType = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetInvoiceRoi(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.InvoiceRoi = v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItems struct {
	ItemRoi *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi `json:"ItemRoi,omitempty" xml:"ItemRoi,omitempty" type:"Struct"`
	Text    *string                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) SetItemRoi(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems {
	s.ItemRoi = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) SetText(v string) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItems {
	s.Text = &v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi struct {
	Size   *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize   `json:"Size,omitempty" xml:"Size,omitempty" type:"Struct"`
	Angle  *float32                                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Center *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter `json:"Center,omitempty" xml:"Center,omitempty" type:"Struct"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetSize(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Size = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetAngle(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Angle = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetCenter(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Center = v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize struct {
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) SetW(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize {
	s.W = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) SetH(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize {
	s.H = &v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter struct {
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) SetY(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter {
	s.Y = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) SetX(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter {
	s.X = &v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi struct {
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetW(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.W = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetH(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.H = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetY(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.Y = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetX(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.X = &v
	return s
}

type RecognizeTaxiInvoiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTaxiInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTaxiInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTaxiInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetBody(v *RecognizeTaxiInvoiceResponseBody) *RecognizeTaxiInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeVATInvoiceRequest struct {
	FileURL  *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s RecognizeVATInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceRequest) SetFileURL(v string) *RecognizeVATInvoiceRequest {
	s.FileURL = &v
	return s
}

func (s *RecognizeVATInvoiceRequest) SetFileType(v string) *RecognizeVATInvoiceRequest {
	s.FileType = &v
	return s
}

type RecognizeVATInvoiceAdvanceRequest struct {
	FileURLObject io.Reader `json:"FileURLObject,omitempty" xml:"FileURLObject,omitempty" require:"true"`
	FileType      *string   `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s RecognizeVATInvoiceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceAdvanceRequest) SetFileURLObject(v io.Reader) *RecognizeVATInvoiceAdvanceRequest {
	s.FileURLObject = v
	return s
}

func (s *RecognizeVATInvoiceAdvanceRequest) SetFileType(v string) *RecognizeVATInvoiceAdvanceRequest {
	s.FileType = &v
	return s
}

type RecognizeVATInvoiceResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeVATInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeVATInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBody) SetRequestId(v string) *RecognizeVATInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBody) SetData(v *RecognizeVATInvoiceResponseBodyData) *RecognizeVATInvoiceResponseBody {
	s.Data = v
	return s
}

type RecognizeVATInvoiceResponseBodyData struct {
	Box     *RecognizeVATInvoiceResponseBodyDataBox     `json:"Box,omitempty" xml:"Box,omitempty" type:"Struct"`
	Content *RecognizeVATInvoiceResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
}

func (s RecognizeVATInvoiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyData) SetBox(v *RecognizeVATInvoiceResponseBodyDataBox) *RecognizeVATInvoiceResponseBodyData {
	s.Box = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyData) SetContent(v *RecognizeVATInvoiceResponseBodyDataContent) *RecognizeVATInvoiceResponseBodyData {
	s.Content = v
	return s
}

type RecognizeVATInvoiceResponseBodyDataBox struct {
	PayerRegisterNoes []*float32 `json:"PayerRegisterNoes,omitempty" xml:"PayerRegisterNoes,omitempty" type:"Repeated"`
	PayeeAddresses    []*float32 `json:"PayeeAddresses,omitempty" xml:"PayeeAddresses,omitempty" type:"Repeated"`
	PayeeBankNames    []*float32 `json:"PayeeBankNames,omitempty" xml:"PayeeBankNames,omitempty" type:"Repeated"`
	Checkers          []*float32 `json:"Checkers,omitempty" xml:"Checkers,omitempty" type:"Repeated"`
	TaxAmounts        []*float32 `json:"TaxAmounts,omitempty" xml:"TaxAmounts,omitempty" type:"Repeated"`
	SumAmounts        []*float32 `json:"SumAmounts,omitempty" xml:"SumAmounts,omitempty" type:"Repeated"`
	Clerks            []*float32 `json:"Clerks,omitempty" xml:"Clerks,omitempty" type:"Repeated"`
	InvoiceNoes       []*float32 `json:"InvoiceNoes,omitempty" xml:"InvoiceNoes,omitempty" type:"Repeated"`
	InvoiceDates      []*float32 `json:"InvoiceDates,omitempty" xml:"InvoiceDates,omitempty" type:"Repeated"`
	InvoiceCodes      []*float32 `json:"InvoiceCodes,omitempty" xml:"InvoiceCodes,omitempty" type:"Repeated"`
	InvoiceFakeCodes  []*float32 `json:"InvoiceFakeCodes,omitempty" xml:"InvoiceFakeCodes,omitempty" type:"Repeated"`
	PayerNames        []*float32 `json:"PayerNames,omitempty" xml:"PayerNames,omitempty" type:"Repeated"`
	PayerBankNames    []*float32 `json:"PayerBankNames,omitempty" xml:"PayerBankNames,omitempty" type:"Repeated"`
	Payees            []*float32 `json:"Payees,omitempty" xml:"Payees,omitempty" type:"Repeated"`
	PayeeNames        []*float32 `json:"PayeeNames,omitempty" xml:"PayeeNames,omitempty" type:"Repeated"`
	InvoiceAmounts    []*float32 `json:"InvoiceAmounts,omitempty" xml:"InvoiceAmounts,omitempty" type:"Repeated"`
	WithoutTaxAmounts []*float32 `json:"WithoutTaxAmounts,omitempty" xml:"WithoutTaxAmounts,omitempty" type:"Repeated"`
	PayerAddresses    []*float32 `json:"PayerAddresses,omitempty" xml:"PayerAddresses,omitempty" type:"Repeated"`
	PayeeRegisterNoes []*float32 `json:"PayeeRegisterNoes,omitempty" xml:"PayeeRegisterNoes,omitempty" type:"Repeated"`
}

func (s RecognizeVATInvoiceResponseBodyDataBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyDataBox) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerRegisterNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerRegisterNoes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeAddresses(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeAddresses = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeBankNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeBankNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetCheckers(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Checkers = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetTaxAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.TaxAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetSumAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.SumAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetClerks(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Clerks = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceNoes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceDates(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceDates = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceCodes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceCodes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceFakeCodes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceFakeCodes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerBankNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerBankNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayees(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Payees = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetWithoutTaxAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.WithoutTaxAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerAddresses(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerAddresses = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeRegisterNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeRegisterNoes = v
	return s
}

type RecognizeVATInvoiceResponseBodyDataContent struct {
	PayerAddress     *string `json:"PayerAddress,omitempty" xml:"PayerAddress,omitempty"`
	PayeeRegisterNo  *string `json:"PayeeRegisterNo,omitempty" xml:"PayeeRegisterNo,omitempty"`
	PayeeBankName    *string `json:"PayeeBankName,omitempty" xml:"PayeeBankName,omitempty"`
	InvoiceNo        *string `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	PayerRegisterNo  *string `json:"PayerRegisterNo,omitempty" xml:"PayerRegisterNo,omitempty"`
	Checker          *string `json:"Checker,omitempty" xml:"Checker,omitempty"`
	TaxAmount        *string `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	InvoiceDate      *string `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	WithoutTaxAmount *string `json:"WithoutTaxAmount,omitempty" xml:"WithoutTaxAmount,omitempty"`
	InvoiceAmount    *string `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	AntiFakeCode     *string `json:"AntiFakeCode,omitempty" xml:"AntiFakeCode,omitempty"`
	PayerName        *string `json:"PayerName,omitempty" xml:"PayerName,omitempty"`
	Payee            *string `json:"Payee,omitempty" xml:"Payee,omitempty"`
	SumAmount        *string `json:"SumAmount,omitempty" xml:"SumAmount,omitempty"`
	PayerBankName    *string `json:"PayerBankName,omitempty" xml:"PayerBankName,omitempty"`
	Clerk            *string `json:"Clerk,omitempty" xml:"Clerk,omitempty"`
	PayeeName        *string `json:"PayeeName,omitempty" xml:"PayeeName,omitempty"`
	PayeeAddress     *string `json:"PayeeAddress,omitempty" xml:"PayeeAddress,omitempty"`
	InvoiceCode      *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
}

func (s RecognizeVATInvoiceResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerAddress(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerAddress = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeRegisterNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeRegisterNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeBankName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeBankName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerRegisterNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerRegisterNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetChecker(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Checker = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetTaxAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.TaxAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceDate(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceDate = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetWithoutTaxAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.WithoutTaxAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetAntiFakeCode(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.AntiFakeCode = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayee(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Payee = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetSumAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.SumAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerBankName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerBankName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetClerk(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Clerk = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeAddress(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeAddress = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceCode(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceCode = &v
	return s
}

type RecognizeVATInvoiceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVATInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVATInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeVATInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVATInvoiceResponse) SetBody(v *RecognizeVATInvoiceResponseBody) *RecognizeVATInvoiceResponse {
	s.Body = v
	return s
}

type RecognizePassportMRZRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizePassportMRZRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZRequest) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZRequest) SetImageURL(v string) *RecognizePassportMRZRequest {
	s.ImageURL = &v
	return s
}

type RecognizePassportMRZAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s RecognizePassportMRZAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizePassportMRZAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizePassportMRZResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizePassportMRZResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizePassportMRZResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZResponseBody) SetRequestId(v string) *RecognizePassportMRZResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizePassportMRZResponseBody) SetData(v *RecognizePassportMRZResponseBodyData) *RecognizePassportMRZResponseBody {
	s.Data = v
	return s
}

type RecognizePassportMRZResponseBodyData struct {
	Regions []*RecognizePassportMRZResponseBodyDataRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s RecognizePassportMRZResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZResponseBodyData) SetRegions(v []*RecognizePassportMRZResponseBodyDataRegions) *RecognizePassportMRZResponseBodyData {
	s.Regions = v
	return s
}

type RecognizePassportMRZResponseBodyDataRegions struct {
	RecognitionScore *float32   `json:"RecognitionScore,omitempty" xml:"RecognitionScore,omitempty"`
	DetectionScore   *float32   `json:"DetectionScore,omitempty" xml:"DetectionScore,omitempty"`
	Name             *string    `json:"Name,omitempty" xml:"Name,omitempty"`
	Content          *string    `json:"Content,omitempty" xml:"Content,omitempty"`
	BandBoxes        []*float32 `json:"BandBoxes,omitempty" xml:"BandBoxes,omitempty" type:"Repeated"`
}

func (s RecognizePassportMRZResponseBodyDataRegions) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZResponseBodyDataRegions) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZResponseBodyDataRegions) SetRecognitionScore(v float32) *RecognizePassportMRZResponseBodyDataRegions {
	s.RecognitionScore = &v
	return s
}

func (s *RecognizePassportMRZResponseBodyDataRegions) SetDetectionScore(v float32) *RecognizePassportMRZResponseBodyDataRegions {
	s.DetectionScore = &v
	return s
}

func (s *RecognizePassportMRZResponseBodyDataRegions) SetName(v string) *RecognizePassportMRZResponseBodyDataRegions {
	s.Name = &v
	return s
}

func (s *RecognizePassportMRZResponseBodyDataRegions) SetContent(v string) *RecognizePassportMRZResponseBodyDataRegions {
	s.Content = &v
	return s
}

func (s *RecognizePassportMRZResponseBodyDataRegions) SetBandBoxes(v []*float32) *RecognizePassportMRZResponseBodyDataRegions {
	s.BandBoxes = v
	return s
}

type RecognizePassportMRZResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizePassportMRZResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePassportMRZResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZResponse) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZResponse) SetHeaders(v map[string]*string) *RecognizePassportMRZResponse {
	s.Headers = v
	return s
}

func (s *RecognizePassportMRZResponse) SetBody(v *RecognizePassportMRZResponseBody) *RecognizePassportMRZResponse {
	s.Body = v
	return s
}

type RecognizeTrainTicketRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTrainTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketRequest) SetImageType(v int32) *RecognizeTrainTicketRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeTrainTicketRequest) SetImageURL(v string) *RecognizeTrainTicketRequest {
	s.ImageURL = &v
	return s
}

type RecognizeTrainTicketAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeTrainTicketAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTrainTicketAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTrainTicketAdvanceRequest) SetImageType(v int32) *RecognizeTrainTicketAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeTrainTicketResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeTrainTicketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeTrainTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponseBody) SetRequestId(v string) *RecognizeTrainTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeTrainTicketResponseBody) SetData(v *RecognizeTrainTicketResponseBodyData) *RecognizeTrainTicketResponseBody {
	s.Data = v
	return s
}

type RecognizeTrainTicketResponseBodyData struct {
	Price            *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	Destination      *string  `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DepartureStation *string  `json:"DepartureStation,omitempty" xml:"DepartureStation,omitempty"`
	Date             *string  `json:"Date,omitempty" xml:"Date,omitempty"`
	Number           *string  `json:"Number,omitempty" xml:"Number,omitempty"`
	Seat             *string  `json:"Seat,omitempty" xml:"Seat,omitempty"`
	Name             *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Level            *string  `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s RecognizeTrainTicketResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponseBodyData) SetPrice(v float32) *RecognizeTrainTicketResponseBodyData {
	s.Price = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetDestination(v string) *RecognizeTrainTicketResponseBodyData {
	s.Destination = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetDepartureStation(v string) *RecognizeTrainTicketResponseBodyData {
	s.DepartureStation = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetDate(v string) *RecognizeTrainTicketResponseBodyData {
	s.Date = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetNumber(v string) *RecognizeTrainTicketResponseBodyData {
	s.Number = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetSeat(v string) *RecognizeTrainTicketResponseBodyData {
	s.Seat = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetName(v string) *RecognizeTrainTicketResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetLevel(v string) *RecognizeTrainTicketResponseBodyData {
	s.Level = &v
	return s
}

type RecognizeTrainTicketResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeTrainTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTrainTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponse) SetHeaders(v map[string]*string) *RecognizeTrainTicketResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTrainTicketResponse) SetBody(v *RecognizeTrainTicketResponseBody) *RecognizeTrainTicketResponse {
	s.Body = v
	return s
}

type RecognizePoiNameRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizePoiNameRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameRequest) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameRequest) SetImageURL(v string) *RecognizePoiNameRequest {
	s.ImageURL = &v
	return s
}

type RecognizePoiNameAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
}

func (s RecognizePoiNameAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizePoiNameAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizePoiNameResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizePoiNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizePoiNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBody) SetRequestId(v string) *RecognizePoiNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizePoiNameResponseBody) SetData(v *RecognizePoiNameResponseBodyData) *RecognizePoiNameResponseBody {
	s.Data = v
	return s
}

type RecognizePoiNameResponseBodyData struct {
	Summary    *RecognizePoiNameResponseBodyDataSummary      `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
	Signboards []*RecognizePoiNameResponseBodyDataSignboards `json:"Signboards,omitempty" xml:"Signboards,omitempty" type:"Repeated"`
}

func (s RecognizePoiNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBodyData) SetSummary(v *RecognizePoiNameResponseBodyDataSummary) *RecognizePoiNameResponseBodyData {
	s.Summary = v
	return s
}

func (s *RecognizePoiNameResponseBodyData) SetSignboards(v []*RecognizePoiNameResponseBodyDataSignboards) *RecognizePoiNameResponseBodyData {
	s.Signboards = v
	return s
}

type RecognizePoiNameResponseBodyDataSummary struct {
	Brand *string  `json:"Brand,omitempty" xml:"Brand,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s RecognizePoiNameResponseBodyDataSummary) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBodyDataSummary) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBodyDataSummary) SetBrand(v string) *RecognizePoiNameResponseBodyDataSummary {
	s.Brand = &v
	return s
}

func (s *RecognizePoiNameResponseBodyDataSummary) SetScore(v float32) *RecognizePoiNameResponseBodyDataSummary {
	s.Score = &v
	return s
}

type RecognizePoiNameResponseBodyDataSignboards struct {
	Texts []*RecognizePoiNameResponseBodyDataSignboardsTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s RecognizePoiNameResponseBodyDataSignboards) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBodyDataSignboards) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBodyDataSignboards) SetTexts(v []*RecognizePoiNameResponseBodyDataSignboardsTexts) *RecognizePoiNameResponseBodyDataSignboards {
	s.Texts = v
	return s
}

type RecognizePoiNameResponseBodyDataSignboardsTexts struct {
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Score  *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Tag    *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Label  *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Points []*int32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s RecognizePoiNameResponseBodyDataSignboardsTexts) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBodyDataSignboardsTexts) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetType(v string) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Type = &v
	return s
}

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetScore(v float32) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Score = &v
	return s
}

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetTag(v string) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Tag = &v
	return s
}

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetLabel(v string) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Label = &v
	return s
}

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetPoints(v []*int32) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Points = v
	return s
}

type RecognizePoiNameResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizePoiNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePoiNameResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponse) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponse) SetHeaders(v map[string]*string) *RecognizePoiNameResponse {
	s.Headers = v
	return s
}

func (s *RecognizePoiNameResponse) SetBody(v *RecognizePoiNameResponseBody) *RecognizePoiNameResponse {
	s.Body = v
	return s
}

type RecognizeQrCodeRequest struct {
	Tasks []*RecognizeQrCodeRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s RecognizeQrCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeRequest) SetTasks(v []*RecognizeQrCodeRequestTasks) *RecognizeQrCodeRequest {
	s.Tasks = v
	return s
}

type RecognizeQrCodeRequestTasks struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeQrCodeRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeRequestTasks) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeRequestTasks) SetImageURL(v string) *RecognizeQrCodeRequestTasks {
	s.ImageURL = &v
	return s
}

type RecognizeQrCodeResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeQrCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeQrCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBody) SetRequestId(v string) *RecognizeQrCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeQrCodeResponseBody) SetData(v *RecognizeQrCodeResponseBodyData) *RecognizeQrCodeResponseBody {
	s.Data = v
	return s
}

type RecognizeQrCodeResponseBodyData struct {
	Elements []*RecognizeQrCodeResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizeQrCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBodyData) SetElements(v []*RecognizeQrCodeResponseBodyDataElements) *RecognizeQrCodeResponseBodyData {
	s.Elements = v
	return s
}

type RecognizeQrCodeResponseBodyDataElements struct {
	ImageURL *string                                           `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	TaskId   *string                                           `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Results  []*RecognizeQrCodeResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RecognizeQrCodeResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBodyDataElements) SetImageURL(v string) *RecognizeQrCodeResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElements) SetTaskId(v string) *RecognizeQrCodeResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElements) SetResults(v []*RecognizeQrCodeResponseBodyDataElementsResults) *RecognizeQrCodeResponseBodyDataElements {
	s.Results = v
	return s
}

type RecognizeQrCodeResponseBodyDataElementsResults struct {
	Suggestion  *string   `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	QrCodesData []*string `json:"QrCodesData,omitempty" xml:"QrCodesData,omitempty" type:"Repeated"`
	Label       *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	Rate        *float32  `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s RecognizeQrCodeResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetSuggestion(v string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetQrCodesData(v []*string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.QrCodesData = v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetLabel(v string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetRate(v float32) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

type RecognizeQrCodeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeQrCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeQrCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponse) SetHeaders(v map[string]*string) *RecognizeQrCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeQrCodeResponse) SetBody(v *RecognizeQrCodeResponseBody) *RecognizeQrCodeResponse {
	s.Body = v
	return s
}

type RecognizeVINCodeRequest struct {
	ImageType *int32  `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURL  *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVINCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVINCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeRequest) SetImageType(v int32) *RecognizeVINCodeRequest {
	s.ImageType = &v
	return s
}

func (s *RecognizeVINCodeRequest) SetImageURL(v string) *RecognizeVINCodeRequest {
	s.ImageURL = &v
	return s
}

type RecognizeVINCodeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
	ImageType      *int32    `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RecognizeVINCodeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVINCodeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeVINCodeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeVINCodeAdvanceRequest) SetImageType(v int32) *RecognizeVINCodeAdvanceRequest {
	s.ImageType = &v
	return s
}

type RecognizeVINCodeResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RecognizeVINCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RecognizeVINCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVINCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeResponseBody) SetRequestId(v string) *RecognizeVINCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVINCodeResponseBody) SetData(v *RecognizeVINCodeResponseBodyData) *RecognizeVINCodeResponseBody {
	s.Data = v
	return s
}

type RecognizeVINCodeResponseBodyData struct {
	VinCode *string `json:"VinCode,omitempty" xml:"VinCode,omitempty"`
}

func (s RecognizeVINCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVINCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeResponseBodyData) SetVinCode(v string) *RecognizeVINCodeResponseBodyData {
	s.VinCode = &v
	return s
}

type RecognizeVINCodeResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecognizeVINCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVINCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVINCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeResponse) SetHeaders(v map[string]*string) *RecognizeVINCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVINCodeResponse) SetBody(v *RecognizeVINCodeResponseBody) *RecognizeVINCodeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ocr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeDrivingLicenseWithOptions(request *RecognizeDrivingLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeDrivingLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeDrivingLicense"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeDrivingLicense(request *RecognizeDrivingLicenseRequest) (_result *RecognizeDrivingLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.RecognizeDrivingLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeDrivingLicenseAdvance(request *RecognizeDrivingLicenseAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeDrivingLicenseResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeDrivingLicenseReq := &RecognizeDrivingLicenseRequest{}
	openapiutil.Convert(request, recognizeDrivingLicenseReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeDrivingLicenseReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeDrivingLicenseResp, _err := client.RecognizeDrivingLicenseWithOptions(recognizeDrivingLicenseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeDrivingLicenseResp
	return _result, _err
}

func (client *Client) RecognizeChinapassportWithOptions(request *RecognizeChinapassportRequest, runtime *util.RuntimeOptions) (_result *RecognizeChinapassportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeChinapassportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeChinapassport"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeChinapassport(request *RecognizeChinapassportRequest) (_result *RecognizeChinapassportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeChinapassportResponse{}
	_body, _err := client.RecognizeChinapassportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeChinapassportAdvance(request *RecognizeChinapassportAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeChinapassportResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeChinapassportReq := &RecognizeChinapassportRequest{}
	openapiutil.Convert(request, recognizeChinapassportReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeChinapassportReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeChinapassportResp, _err := client.RecognizeChinapassportWithOptions(recognizeChinapassportReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeChinapassportResp
	return _result, _err
}

func (client *Client) TrimDocumentWithOptions(request *TrimDocumentRequest, runtime *util.RuntimeOptions) (_result *TrimDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TrimDocumentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TrimDocument"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TrimDocument(request *TrimDocumentRequest) (_result *TrimDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TrimDocumentResponse{}
	_body, _err := client.TrimDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TrimDocumentAdvance(request *TrimDocumentAdvanceRequest, runtime *util.RuntimeOptions) (_result *TrimDocumentResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	trimDocumentReq := &TrimDocumentRequest{}
	openapiutil.Convert(request, trimDocumentReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.FileURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	trimDocumentReq.FileURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	trimDocumentResp, _err := client.TrimDocumentWithOptions(trimDocumentReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = trimDocumentResp
	return _result, _err
}

func (client *Client) RecognizeTableWithOptions(request *RecognizeTableRequest, runtime *util.RuntimeOptions) (_result *RecognizeTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeTableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTable"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTable(request *RecognizeTableRequest) (_result *RecognizeTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTableResponse{}
	_body, _err := client.RecognizeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTableAdvance(request *RecognizeTableAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTableResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeTableReq := &RecognizeTableRequest{}
	openapiutil.Convert(request, recognizeTableReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeTableReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeTableResp, _err := client.RecognizeTableWithOptions(recognizeTableReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTableResp
	return _result, _err
}

func (client *Client) RecognizeIdentityCardWithOptions(request *RecognizeIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeIdentityCardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeIdentityCard"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeIdentityCard(request *RecognizeIdentityCardRequest) (_result *RecognizeIdentityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeIdentityCardResponse{}
	_body, _err := client.RecognizeIdentityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeIdentityCardAdvance(request *RecognizeIdentityCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeIdentityCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeIdentityCardReq := &RecognizeIdentityCardRequest{}
	openapiutil.Convert(request, recognizeIdentityCardReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeIdentityCardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeIdentityCardResp, _err := client.RecognizeIdentityCardWithOptions(recognizeIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeIdentityCardResp
	return _result, _err
}

func (client *Client) RecognizeBusinessLicenseWithOptions(request *RecognizeBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBusinessLicense"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBusinessLicense(request *RecognizeBusinessLicenseRequest) (_result *RecognizeBusinessLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.RecognizeBusinessLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBusinessLicenseAdvance(request *RecognizeBusinessLicenseAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessLicenseResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeBusinessLicenseReq := &RecognizeBusinessLicenseRequest{}
	openapiutil.Convert(request, recognizeBusinessLicenseReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeBusinessLicenseReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeBusinessLicenseResp, _err := client.RecognizeBusinessLicenseWithOptions(recognizeBusinessLicenseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeBusinessLicenseResp
	return _result, _err
}

func (client *Client) RecognizeBankCardWithOptions(request *RecognizeBankCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBankCard"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBankCard(request *RecognizeBankCardRequest) (_result *RecognizeBankCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.RecognizeBankCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBankCardAdvance(request *RecognizeBankCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeBankCardReq := &RecognizeBankCardRequest{}
	openapiutil.Convert(request, recognizeBankCardReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeBankCardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeBankCardResp, _err := client.RecognizeBankCardWithOptions(recognizeBankCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeBankCardResp
	return _result, _err
}

func (client *Client) RecognizeVerificationcodeWithOptions(request *RecognizeVerificationcodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeVerificationcodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeVerificationcodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeVerificationcode"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVerificationcode(request *RecognizeVerificationcodeRequest) (_result *RecognizeVerificationcodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVerificationcodeResponse{}
	_body, _err := client.RecognizeVerificationcodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVerificationcodeAdvance(request *RecognizeVerificationcodeAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVerificationcodeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVerificationcodeReq := &RecognizeVerificationcodeRequest{}
	openapiutil.Convert(request, recognizeVerificationcodeReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeVerificationcodeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeVerificationcodeResp, _err := client.RecognizeVerificationcodeWithOptions(recognizeVerificationcodeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVerificationcodeResp
	return _result, _err
}

func (client *Client) RecognizeAccountPageWithOptions(request *RecognizeAccountPageRequest, runtime *util.RuntimeOptions) (_result *RecognizeAccountPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeAccountPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeAccountPage"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeAccountPage(request *RecognizeAccountPageRequest) (_result *RecognizeAccountPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeAccountPageResponse{}
	_body, _err := client.RecognizeAccountPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeAccountPageAdvance(request *RecognizeAccountPageAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeAccountPageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeAccountPageReq := &RecognizeAccountPageRequest{}
	openapiutil.Convert(request, recognizeAccountPageReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeAccountPageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeAccountPageResp, _err := client.RecognizeAccountPageWithOptions(recognizeAccountPageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeAccountPageResp
	return _result, _err
}

func (client *Client) RecognizeCharacterWithOptions(request *RecognizeCharacterRequest, runtime *util.RuntimeOptions) (_result *RecognizeCharacterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeCharacterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeCharacter"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeCharacter(request *RecognizeCharacterRequest) (_result *RecognizeCharacterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeCharacterResponse{}
	_body, _err := client.RecognizeCharacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeCharacterAdvance(request *RecognizeCharacterAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeCharacterResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeCharacterReq := &RecognizeCharacterRequest{}
	openapiutil.Convert(request, recognizeCharacterReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeCharacterReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeCharacterResp, _err := client.RecognizeCharacterWithOptions(recognizeCharacterReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeCharacterResp
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAsyncJobResult"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTakeoutOrderWithOptions(request *RecognizeTakeoutOrderRequest, runtime *util.RuntimeOptions) (_result *RecognizeTakeoutOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeTakeoutOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTakeoutOrder"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTakeoutOrder(request *RecognizeTakeoutOrderRequest) (_result *RecognizeTakeoutOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTakeoutOrderResponse{}
	_body, _err := client.RecognizeTakeoutOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTakeoutOrderAdvance(request *RecognizeTakeoutOrderAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTakeoutOrderResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeTakeoutOrderReq := &RecognizeTakeoutOrderRequest{}
	openapiutil.Convert(request, recognizeTakeoutOrderReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeTakeoutOrderReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeTakeoutOrderResp, _err := client.RecognizeTakeoutOrderWithOptions(recognizeTakeoutOrderReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTakeoutOrderResp
	return _result, _err
}

func (client *Client) RecognizeBusinessCardWithOptions(request *RecognizeBusinessCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeBusinessCardResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeBusinessCard"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeBusinessCard(request *RecognizeBusinessCardRequest) (_result *RecognizeBusinessCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeBusinessCardResponse{}
	_body, _err := client.RecognizeBusinessCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeBusinessCardAdvance(request *RecognizeBusinessCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeBusinessCardReq := &RecognizeBusinessCardRequest{}
	openapiutil.Convert(request, recognizeBusinessCardReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeBusinessCardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeBusinessCardResp, _err := client.RecognizeBusinessCardWithOptions(recognizeBusinessCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeBusinessCardResp
	return _result, _err
}

func (client *Client) DetectCardScreenshotWithOptions(request *DetectCardScreenshotRequest, runtime *util.RuntimeOptions) (_result *DetectCardScreenshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectCardScreenshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectCardScreenshot"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectCardScreenshot(request *DetectCardScreenshotRequest) (_result *DetectCardScreenshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectCardScreenshotResponse{}
	_body, _err := client.DetectCardScreenshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectCardScreenshotAdvance(request *DetectCardScreenshotAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectCardScreenshotResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectCardScreenshotReq := &DetectCardScreenshotRequest{}
	openapiutil.Convert(request, detectCardScreenshotReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	detectCardScreenshotReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectCardScreenshotResp, _err := client.DetectCardScreenshotWithOptions(detectCardScreenshotReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectCardScreenshotResp
	return _result, _err
}

func (client *Client) RecognizeDriverLicenseWithOptions(request *RecognizeDriverLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeDriverLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeDriverLicenseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeDriverLicense"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeDriverLicense(request *RecognizeDriverLicenseRequest) (_result *RecognizeDriverLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeDriverLicenseResponse{}
	_body, _err := client.RecognizeDriverLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeDriverLicenseAdvance(request *RecognizeDriverLicenseAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeDriverLicenseResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeDriverLicenseReq := &RecognizeDriverLicenseRequest{}
	openapiutil.Convert(request, recognizeDriverLicenseReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeDriverLicenseReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeDriverLicenseResp, _err := client.RecognizeDriverLicenseWithOptions(recognizeDriverLicenseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeDriverLicenseResp
	return _result, _err
}

func (client *Client) RecognizeLicensePlateWithOptions(request *RecognizeLicensePlateRequest, runtime *util.RuntimeOptions) (_result *RecognizeLicensePlateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeLicensePlateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeLicensePlate"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeLicensePlate(request *RecognizeLicensePlateRequest) (_result *RecognizeLicensePlateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeLicensePlateResponse{}
	_body, _err := client.RecognizeLicensePlateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeLicensePlateAdvance(request *RecognizeLicensePlateAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeLicensePlateResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeLicensePlateReq := &RecognizeLicensePlateRequest{}
	openapiutil.Convert(request, recognizeLicensePlateReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeLicensePlateReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeLicensePlateResp, _err := client.RecognizeLicensePlateWithOptions(recognizeLicensePlateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeLicensePlateResp
	return _result, _err
}

func (client *Client) RecognizeStampWithOptions(request *RecognizeStampRequest, runtime *util.RuntimeOptions) (_result *RecognizeStampResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeStampResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeStamp"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeStamp(request *RecognizeStampRequest) (_result *RecognizeStampResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeStampResponse{}
	_body, _err := client.RecognizeStampWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeStampAdvance(request *RecognizeStampAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeStampResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeStampReq := &RecognizeStampRequest{}
	openapiutil.Convert(request, recognizeStampReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeStampReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeStampResp, _err := client.RecognizeStampWithOptions(recognizeStampReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeStampResp
	return _result, _err
}

func (client *Client) RecognizeTaxiInvoiceWithOptions(request *RecognizeTaxiInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTaxiInvoice"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTaxiInvoice(request *RecognizeTaxiInvoiceRequest) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.RecognizeTaxiInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTaxiInvoiceAdvance(request *RecognizeTaxiInvoiceAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeTaxiInvoiceReq := &RecognizeTaxiInvoiceRequest{}
	openapiutil.Convert(request, recognizeTaxiInvoiceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeTaxiInvoiceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeTaxiInvoiceResp, _err := client.RecognizeTaxiInvoiceWithOptions(recognizeTaxiInvoiceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTaxiInvoiceResp
	return _result, _err
}

func (client *Client) RecognizeVATInvoiceWithOptions(request *RecognizeVATInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVATInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeVATInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeVATInvoice"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVATInvoice(request *RecognizeVATInvoiceRequest) (_result *RecognizeVATInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVATInvoiceResponse{}
	_body, _err := client.RecognizeVATInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVATInvoiceAdvance(request *RecognizeVATInvoiceAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVATInvoiceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVATInvoiceReq := &RecognizeVATInvoiceRequest{}
	openapiutil.Convert(request, recognizeVATInvoiceReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.FileURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeVATInvoiceReq.FileURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeVATInvoiceResp, _err := client.RecognizeVATInvoiceWithOptions(recognizeVATInvoiceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVATInvoiceResp
	return _result, _err
}

func (client *Client) RecognizePassportMRZWithOptions(request *RecognizePassportMRZRequest, runtime *util.RuntimeOptions) (_result *RecognizePassportMRZResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizePassportMRZResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizePassportMRZ"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePassportMRZ(request *RecognizePassportMRZRequest) (_result *RecognizePassportMRZResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePassportMRZResponse{}
	_body, _err := client.RecognizePassportMRZWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizePassportMRZAdvance(request *RecognizePassportMRZAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizePassportMRZResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizePassportMRZReq := &RecognizePassportMRZRequest{}
	openapiutil.Convert(request, recognizePassportMRZReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizePassportMRZReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizePassportMRZResp, _err := client.RecognizePassportMRZWithOptions(recognizePassportMRZReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizePassportMRZResp
	return _result, _err
}

func (client *Client) RecognizeTrainTicketWithOptions(request *RecognizeTrainTicketRequest, runtime *util.RuntimeOptions) (_result *RecognizeTrainTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeTrainTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeTrainTicket"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTrainTicket(request *RecognizeTrainTicketRequest) (_result *RecognizeTrainTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTrainTicketResponse{}
	_body, _err := client.RecognizeTrainTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTrainTicketAdvance(request *RecognizeTrainTicketAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTrainTicketResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeTrainTicketReq := &RecognizeTrainTicketRequest{}
	openapiutil.Convert(request, recognizeTrainTicketReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeTrainTicketReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeTrainTicketResp, _err := client.RecognizeTrainTicketWithOptions(recognizeTrainTicketReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTrainTicketResp
	return _result, _err
}

func (client *Client) RecognizePoiNameWithOptions(request *RecognizePoiNameRequest, runtime *util.RuntimeOptions) (_result *RecognizePoiNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizePoiNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizePoiName"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePoiName(request *RecognizePoiNameRequest) (_result *RecognizePoiNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePoiNameResponse{}
	_body, _err := client.RecognizePoiNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizePoiNameAdvance(request *RecognizePoiNameAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizePoiNameResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizePoiNameReq := &RecognizePoiNameRequest{}
	openapiutil.Convert(request, recognizePoiNameReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizePoiNameReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizePoiNameResp, _err := client.RecognizePoiNameWithOptions(recognizePoiNameReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizePoiNameResp
	return _result, _err
}

func (client *Client) RecognizeQrCodeWithOptions(request *RecognizeQrCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeQrCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeQrCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeQrCode"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeQrCode(request *RecognizeQrCodeRequest) (_result *RecognizeQrCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeQrCodeResponse{}
	_body, _err := client.RecognizeQrCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVINCodeWithOptions(request *RecognizeVINCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeVINCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RecognizeVINCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecognizeVINCode"), tea.String("2019-12-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVINCode(request *RecognizeVINCodeRequest) (_result *RecognizeVINCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVINCodeResponse{}
	_body, _err := client.RecognizeVINCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVINCodeAdvance(request *RecognizeVINCodeAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVINCodeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ocr"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVINCodeReq := &RecognizeVINCodeRequest{}
	openapiutil.Convert(request, recognizeVINCodeReq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	recognizeVINCodeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	recognizeVINCodeResp, _err := client.RecognizeVINCodeWithOptions(recognizeVINCodeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVINCodeResp
	return _result, _err
}
