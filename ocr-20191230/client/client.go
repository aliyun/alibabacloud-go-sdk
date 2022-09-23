// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *DetectCardScreenshotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectCardScreenshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotResponseBody) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotResponseBody) SetData(v *DetectCardScreenshotResponseBodyData) *DetectCardScreenshotResponseBody {
	s.Data = v
	return s
}

func (s *DetectCardScreenshotResponseBody) SetRequestId(v string) *DetectCardScreenshotResponseBody {
	s.RequestId = &v
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
	ScreenScore     *float32 `json:"ScreenScore,omitempty" xml:"ScreenScore,omitempty"`
	ScreenThreshold *float32 `json:"ScreenThreshold,omitempty" xml:"ScreenThreshold,omitempty"`
}

func (s DetectCardScreenshotResponseBodyDataSpoofResultResultMap) String() string {
	return tea.Prettify(s)
}

func (s DetectCardScreenshotResponseBodyDataSpoofResultResultMap) GoString() string {
	return s.String()
}

func (s *DetectCardScreenshotResponseBodyDataSpoofResultResultMap) SetScreenScore(v float32) *DetectCardScreenshotResponseBodyDataSpoofResultResultMap {
	s.ScreenScore = &v
	return s
}

func (s *DetectCardScreenshotResponseBodyDataSpoofResultResultMap) SetScreenThreshold(v float32) *DetectCardScreenshotResponseBodyDataSpoofResultResultMap {
	s.ScreenThreshold = &v
	return s
}

type DetectCardScreenshotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectCardScreenshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetectCardScreenshotResponse) SetStatusCode(v int32) *DetectCardScreenshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectCardScreenshotResponse) SetBody(v *DetectCardScreenshotResponseBody) *DetectCardScreenshotResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type RecognizeAccountPageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeAccountPageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageRequest) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageRequest) SetImageURL(v string) *RecognizeAccountPageRequest {
	s.ImageURL = &v
	return s
}

type RecognizeAccountPageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeAccountPageResponseBody struct {
	Data      *RecognizeAccountPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeAccountPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBody) SetData(v *RecognizeAccountPageResponseBodyData) *RecognizeAccountPageResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeAccountPageResponseBody) SetRequestId(v string) *RecognizeAccountPageResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeAccountPageResponseBodyData struct {
	Angle               *float32                                                   `json:"Angle,omitempty" xml:"Angle,omitempty"`
	BirthDate           *string                                                    `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	BirthPlace          *string                                                    `json:"BirthPlace,omitempty" xml:"BirthPlace,omitempty"`
	Gender              *string                                                    `json:"Gender,omitempty" xml:"Gender,omitempty"`
	IDNumber            *string                                                    `json:"IDNumber,omitempty" xml:"IDNumber,omitempty"`
	InvalidStampAreas   []*RecognizeAccountPageResponseBodyDataInvalidStampAreas   `json:"InvalidStampAreas,omitempty" xml:"InvalidStampAreas,omitempty" type:"Repeated"`
	Name                *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality         *string                                                    `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	NativePlace         *string                                                    `json:"NativePlace,omitempty" xml:"NativePlace,omitempty"`
	OtherStampAreas     []*RecognizeAccountPageResponseBodyDataOtherStampAreas     `json:"OtherStampAreas,omitempty" xml:"OtherStampAreas,omitempty" type:"Repeated"`
	RegisterStampAreas  []*RecognizeAccountPageResponseBodyDataRegisterStampAreas  `json:"RegisterStampAreas,omitempty" xml:"RegisterStampAreas,omitempty" type:"Repeated"`
	Relation            *string                                                    `json:"Relation,omitempty" xml:"Relation,omitempty"`
	TitleArea           *RecognizeAccountPageResponseBodyDataTitleArea             `json:"TitleArea,omitempty" xml:"TitleArea,omitempty" type:"Struct"`
	UndertakeStampAreas []*RecognizeAccountPageResponseBodyDataUndertakeStampAreas `json:"UndertakeStampAreas,omitempty" xml:"UndertakeStampAreas,omitempty" type:"Repeated"`
}

func (s RecognizeAccountPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyData) SetAngle(v float32) *RecognizeAccountPageResponseBodyData {
	s.Angle = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetBirthDate(v string) *RecognizeAccountPageResponseBodyData {
	s.BirthDate = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetBirthPlace(v string) *RecognizeAccountPageResponseBodyData {
	s.BirthPlace = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetGender(v string) *RecognizeAccountPageResponseBodyData {
	s.Gender = &v
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

func (s *RecognizeAccountPageResponseBodyData) SetName(v string) *RecognizeAccountPageResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetNationality(v string) *RecognizeAccountPageResponseBodyData {
	s.Nationality = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetNativePlace(v string) *RecognizeAccountPageResponseBodyData {
	s.NativePlace = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetOtherStampAreas(v []*RecognizeAccountPageResponseBodyDataOtherStampAreas) *RecognizeAccountPageResponseBodyData {
	s.OtherStampAreas = v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetRegisterStampAreas(v []*RecognizeAccountPageResponseBodyDataRegisterStampAreas) *RecognizeAccountPageResponseBodyData {
	s.RegisterStampAreas = v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetRelation(v string) *RecognizeAccountPageResponseBodyData {
	s.Relation = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetTitleArea(v *RecognizeAccountPageResponseBodyDataTitleArea) *RecognizeAccountPageResponseBodyData {
	s.TitleArea = v
	return s
}

func (s *RecognizeAccountPageResponseBodyData) SetUndertakeStampAreas(v []*RecognizeAccountPageResponseBodyDataUndertakeStampAreas) *RecognizeAccountPageResponseBodyData {
	s.UndertakeStampAreas = v
	return s
}

type RecognizeAccountPageResponseBodyDataInvalidStampAreas struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataInvalidStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataInvalidStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Left = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataInvalidStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataInvalidStampAreas {
	s.Width = &v
	return s
}

type RecognizeAccountPageResponseBodyDataOtherStampAreas struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataOtherStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataOtherStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Left = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataOtherStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataOtherStampAreas {
	s.Width = &v
	return s
}

type RecognizeAccountPageResponseBodyDataRegisterStampAreas struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataRegisterStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataRegisterStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Left = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataRegisterStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataRegisterStampAreas {
	s.Width = &v
	return s
}

type RecognizeAccountPageResponseBodyDataTitleArea struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataTitleArea) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataTitleArea) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Left = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetTop(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataTitleArea) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataTitleArea {
	s.Width = &v
	return s
}

type RecognizeAccountPageResponseBodyDataUndertakeStampAreas struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeAccountPageResponseBodyDataUndertakeStampAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeAccountPageResponseBodyDataUndertakeStampAreas) GoString() string {
	return s.String()
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetHeight(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Height = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetLeft(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Left = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetTop(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Top = &v
	return s
}

func (s *RecognizeAccountPageResponseBodyDataUndertakeStampAreas) SetWidth(v int32) *RecognizeAccountPageResponseBodyDataUndertakeStampAreas {
	s.Width = &v
	return s
}

type RecognizeAccountPageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeAccountPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeAccountPageResponse) SetStatusCode(v int32) *RecognizeAccountPageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeAccountPageResponse) SetBody(v *RecognizeAccountPageResponseBody) *RecognizeAccountPageResponse {
	s.Body = v
	return s
}

type RecognizeBankCardRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBankCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardRequest) SetImageURL(v string) *RecognizeBankCardRequest {
	s.ImageURL = &v
	return s
}

type RecognizeBankCardAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeBankCardResponseBody struct {
	Data      *RecognizeBankCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBankCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBody) SetData(v *RecognizeBankCardResponseBodyData) *RecognizeBankCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeBankCardResponseBody) SetRequestId(v string) *RecognizeBankCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBankCardResponseBodyData struct {
	BankName   *string `json:"BankName,omitempty" xml:"BankName,omitempty"`
	CardNumber *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ValidDate  *string `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
}

func (s RecognizeBankCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBankCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponseBodyData) SetBankName(v string) *RecognizeBankCardResponseBodyData {
	s.BankName = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) SetCardNumber(v string) *RecognizeBankCardResponseBodyData {
	s.CardNumber = &v
	return s
}

func (s *RecognizeBankCardResponseBodyData) SetValidDate(v string) *RecognizeBankCardResponseBodyData {
	s.ValidDate = &v
	return s
}

type RecognizeBankCardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBankCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBankCardResponse) SetStatusCode(v int32) *RecognizeBankCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankCardResponse) SetBody(v *RecognizeBankCardResponseBody) *RecognizeBankCardResponse {
	s.Body = v
	return s
}

type RecognizeBusinessCardRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBusinessCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardRequest) SetImageURL(v string) *RecognizeBusinessCardRequest {
	s.ImageURL = &v
	return s
}

type RecognizeBusinessCardAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeBusinessCardResponseBody struct {
	Data      *RecognizeBusinessCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBusinessCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardResponseBody) SetData(v *RecognizeBusinessCardResponseBodyData) *RecognizeBusinessCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeBusinessCardResponseBody) SetRequestId(v string) *RecognizeBusinessCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBusinessCardResponseBodyData struct {
	Addresses          []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	CellPhoneNumbers   []*string `json:"CellPhoneNumbers,omitempty" xml:"CellPhoneNumbers,omitempty" type:"Repeated"`
	Companies          []*string `json:"Companies,omitempty" xml:"Companies,omitempty" type:"Repeated"`
	Departments        []*string `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	Emails             []*string `json:"Emails,omitempty" xml:"Emails,omitempty" type:"Repeated"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	OfficePhoneNumbers []*string `json:"OfficePhoneNumbers,omitempty" xml:"OfficePhoneNumbers,omitempty" type:"Repeated"`
	Titles             []*string `json:"Titles,omitempty" xml:"Titles,omitempty" type:"Repeated"`
}

func (s RecognizeBusinessCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessCardResponseBodyData) SetAddresses(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Addresses = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetCellPhoneNumbers(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.CellPhoneNumbers = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetCompanies(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Companies = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetDepartments(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Departments = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetEmails(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Emails = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetName(v string) *RecognizeBusinessCardResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetOfficePhoneNumbers(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.OfficePhoneNumbers = v
	return s
}

func (s *RecognizeBusinessCardResponseBodyData) SetTitles(v []*string) *RecognizeBusinessCardResponseBodyData {
	s.Titles = v
	return s
}

type RecognizeBusinessCardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBusinessCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBusinessCardResponse) SetStatusCode(v int32) *RecognizeBusinessCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBusinessCardResponse) SetBody(v *RecognizeBusinessCardResponseBody) *RecognizeBusinessCardResponse {
	s.Body = v
	return s
}

type RecognizeBusinessLicenseRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeBusinessLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseRequest) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseRequest) SetImageURL(v string) *RecognizeBusinessLicenseRequest {
	s.ImageURL = &v
	return s
}

type RecognizeBusinessLicenseAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeBusinessLicenseResponseBody struct {
	Data      *RecognizeBusinessLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeBusinessLicenseResponseBody) SetData(v *RecognizeBusinessLicenseResponseBodyData) *RecognizeBusinessLicenseResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeBusinessLicenseResponseBody) SetRequestId(v string) *RecognizeBusinessLicenseResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeBusinessLicenseResponseBodyData struct {
	Address        *string                                         `json:"Address,omitempty" xml:"Address,omitempty"`
	Angle          *float32                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Business       *string                                         `json:"Business,omitempty" xml:"Business,omitempty"`
	Capital        *string                                         `json:"Capital,omitempty" xml:"Capital,omitempty"`
	Emblem         *RecognizeBusinessLicenseResponseBodyDataEmblem `json:"Emblem,omitempty" xml:"Emblem,omitempty" type:"Struct"`
	EstablishDate  *string                                         `json:"EstablishDate,omitempty" xml:"EstablishDate,omitempty"`
	LegalPerson    *string                                         `json:"LegalPerson,omitempty" xml:"LegalPerson,omitempty"`
	Name           *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	QRCode         *RecognizeBusinessLicenseResponseBodyDataQRCode `json:"QRCode,omitempty" xml:"QRCode,omitempty" type:"Struct"`
	RegisterNumber *string                                         `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Stamp          *RecognizeBusinessLicenseResponseBodyDataStamp  `json:"Stamp,omitempty" xml:"Stamp,omitempty" type:"Struct"`
	Title          *RecognizeBusinessLicenseResponseBodyDataTitle  `json:"Title,omitempty" xml:"Title,omitempty" type:"Struct"`
	Type           *string                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	ValidPeriod    *string                                         `json:"ValidPeriod,omitempty" xml:"ValidPeriod,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyData) GoString() string {
	return s.String()
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

type RecognizeBusinessLicenseResponseBodyDataEmblem struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataEmblem) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataEmblem) GoString() string {
	return s.String()
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

type RecognizeBusinessLicenseResponseBodyDataQRCode struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataQRCode) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataQRCode) GoString() string {
	return s.String()
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

type RecognizeBusinessLicenseResponseBodyDataStamp struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataStamp) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataStamp) GoString() string {
	return s.String()
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

type RecognizeBusinessLicenseResponseBodyDataTitle struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeBusinessLicenseResponseBodyDataTitle) String() string {
	return tea.Prettify(s)
}

func (s RecognizeBusinessLicenseResponseBodyDataTitle) GoString() string {
	return s.String()
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

type RecognizeBusinessLicenseResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeBusinessLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeBusinessLicenseResponse) SetStatusCode(v int32) *RecognizeBusinessLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBusinessLicenseResponse) SetBody(v *RecognizeBusinessLicenseResponseBody) *RecognizeBusinessLicenseResponse {
	s.Body = v
	return s
}

type RecognizeCharacterRequest struct {
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
	ImageURLObject    io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s *RecognizeCharacterAdvanceRequest) SetMinHeight(v int32) *RecognizeCharacterAdvanceRequest {
	s.MinHeight = &v
	return s
}

func (s *RecognizeCharacterAdvanceRequest) SetOutputProbability(v bool) *RecognizeCharacterAdvanceRequest {
	s.OutputProbability = &v
	return s
}

type RecognizeCharacterResponseBody struct {
	Data      *RecognizeCharacterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeCharacterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBody) SetData(v *RecognizeCharacterResponseBodyData) *RecognizeCharacterResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeCharacterResponseBody) SetRequestId(v string) *RecognizeCharacterResponseBody {
	s.RequestId = &v
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
	Probability    *float32                                                 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	Text           *string                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
	TextRectangles *RecognizeCharacterResponseBodyDataResultsTextRectangles `json:"TextRectangles,omitempty" xml:"TextRectangles,omitempty" type:"Struct"`
}

func (s RecognizeCharacterResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyDataResults) SetProbability(v float32) *RecognizeCharacterResponseBodyDataResults {
	s.Probability = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResults) SetText(v string) *RecognizeCharacterResponseBodyDataResults {
	s.Text = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResults) SetTextRectangles(v *RecognizeCharacterResponseBodyDataResultsTextRectangles) *RecognizeCharacterResponseBodyDataResults {
	s.TextRectangles = v
	return s
}

type RecognizeCharacterResponseBodyDataResultsTextRectangles struct {
	Angle  *int32 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectangles) String() string {
	return tea.Prettify(s)
}

func (s RecognizeCharacterResponseBodyDataResultsTextRectangles) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetAngle(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Angle = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetHeight(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Height = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetLeft(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Left = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetTop(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Top = &v
	return s
}

func (s *RecognizeCharacterResponseBodyDataResultsTextRectangles) SetWidth(v int32) *RecognizeCharacterResponseBodyDataResultsTextRectangles {
	s.Width = &v
	return s
}

type RecognizeCharacterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeCharacterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeCharacterResponse) SetStatusCode(v int32) *RecognizeCharacterResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCharacterResponse) SetBody(v *RecognizeCharacterResponseBody) *RecognizeCharacterResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *RecognizeChinapassportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeChinapassportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinapassportResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeChinapassportResponseBody) SetData(v *RecognizeChinapassportResponseBodyData) *RecognizeChinapassportResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeChinapassportResponseBody) SetRequestId(v string) *RecognizeChinapassportResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeChinapassportResponseBodyData struct {
	Authority      *string `json:"Authority,omitempty" xml:"Authority,omitempty"`
	BirthDate      *string `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	BirthDay       *string `json:"BirthDay,omitempty" xml:"BirthDay,omitempty"`
	BirthPlace     *string `json:"BirthPlace,omitempty" xml:"BirthPlace,omitempty"`
	BirthPlaceRaw  *string `json:"BirthPlaceRaw,omitempty" xml:"BirthPlaceRaw,omitempty"`
	Country        *string `json:"Country,omitempty" xml:"Country,omitempty"`
	ExpiryDate     *string `json:"ExpiryDate,omitempty" xml:"ExpiryDate,omitempty"`
	ExpiryDay      *string `json:"ExpiryDay,omitempty" xml:"ExpiryDay,omitempty"`
	IssueDate      *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	IssuePlace     *string `json:"IssuePlace,omitempty" xml:"IssuePlace,omitempty"`
	IssuePlaceRaw  *string `json:"IssuePlaceRaw,omitempty" xml:"IssuePlaceRaw,omitempty"`
	LineOne        *string `json:"LineOne,omitempty" xml:"LineOne,omitempty"`
	LineZero       *string `json:"LineZero,omitempty" xml:"LineZero,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameChinese    *string `json:"NameChinese,omitempty" xml:"NameChinese,omitempty"`
	NameChineseRaw *string `json:"NameChineseRaw,omitempty" xml:"NameChineseRaw,omitempty"`
	PassportNo     *string `json:"PassportNo,omitempty" xml:"PassportNo,omitempty"`
	PersonId       *string `json:"PersonId,omitempty" xml:"PersonId,omitempty"`
	Sex            *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
	SourceCountry  *string `json:"SourceCountry,omitempty" xml:"SourceCountry,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeChinapassportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeChinapassportResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeChinapassportResponseBodyData) SetAuthority(v string) *RecognizeChinapassportResponseBodyData {
	s.Authority = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthDate(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthDate = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthDay(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthDay = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthPlace(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthPlace = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetBirthPlaceRaw(v string) *RecognizeChinapassportResponseBodyData {
	s.BirthPlaceRaw = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetCountry(v string) *RecognizeChinapassportResponseBodyData {
	s.Country = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetExpiryDate(v string) *RecognizeChinapassportResponseBodyData {
	s.ExpiryDate = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetExpiryDay(v string) *RecognizeChinapassportResponseBodyData {
	s.ExpiryDay = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetIssueDate(v string) *RecognizeChinapassportResponseBodyData {
	s.IssueDate = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetIssuePlace(v string) *RecognizeChinapassportResponseBodyData {
	s.IssuePlace = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetIssuePlaceRaw(v string) *RecognizeChinapassportResponseBodyData {
	s.IssuePlaceRaw = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetLineOne(v string) *RecognizeChinapassportResponseBodyData {
	s.LineOne = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetLineZero(v string) *RecognizeChinapassportResponseBodyData {
	s.LineZero = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetName(v string) *RecognizeChinapassportResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetNameChinese(v string) *RecognizeChinapassportResponseBodyData {
	s.NameChinese = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetNameChineseRaw(v string) *RecognizeChinapassportResponseBodyData {
	s.NameChineseRaw = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetPassportNo(v string) *RecognizeChinapassportResponseBodyData {
	s.PassportNo = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetPersonId(v string) *RecognizeChinapassportResponseBodyData {
	s.PersonId = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetSex(v string) *RecognizeChinapassportResponseBodyData {
	s.Sex = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetSourceCountry(v string) *RecognizeChinapassportResponseBodyData {
	s.SourceCountry = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetSuccess(v bool) *RecognizeChinapassportResponseBodyData {
	s.Success = &v
	return s
}

func (s *RecognizeChinapassportResponseBodyData) SetType(v string) *RecognizeChinapassportResponseBodyData {
	s.Type = &v
	return s
}

type RecognizeChinapassportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeChinapassportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeChinapassportResponse) SetStatusCode(v int32) *RecognizeChinapassportResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeChinapassportResponse) SetBody(v *RecognizeChinapassportResponseBody) *RecognizeChinapassportResponse {
	s.Body = v
	return s
}

type RecognizeDriverLicenseRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Side     *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDriverLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseRequest) GoString() string {
	return s.String()
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s *RecognizeDriverLicenseAdvanceRequest) SetSide(v string) *RecognizeDriverLicenseAdvanceRequest {
	s.Side = &v
	return s
}

type RecognizeDriverLicenseResponseBody struct {
	Data      *RecognizeDriverLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDriverLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBody) SetData(v *RecognizeDriverLicenseResponseBodyData) *RecognizeDriverLicenseResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeDriverLicenseResponseBody) SetRequestId(v string) *RecognizeDriverLicenseResponseBody {
	s.RequestId = &v
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
	CardNumber    *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Record        *string `json:"Record,omitempty" xml:"Record,omitempty"`
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

type RecognizeDriverLicenseResponseBodyDataFaceResult struct {
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EndDate       *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Gender        *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	IssueDate     *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	IssueUnit     *string `json:"IssueUnit,omitempty" xml:"IssueUnit,omitempty"`
	LicenseNumber *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	StartDate     *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	VehicleType   *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
}

func (s RecognizeDriverLicenseResponseBodyDataFaceResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDriverLicenseResponseBodyDataFaceResult) GoString() string {
	return s.String()
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetAddress(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.Address = &v
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

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetStartDate(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.StartDate = &v
	return s
}

func (s *RecognizeDriverLicenseResponseBodyDataFaceResult) SetVehicleType(v string) *RecognizeDriverLicenseResponseBodyDataFaceResult {
	s.VehicleType = &v
	return s
}

type RecognizeDriverLicenseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeDriverLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeDriverLicenseResponse) SetStatusCode(v int32) *RecognizeDriverLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDriverLicenseResponse) SetBody(v *RecognizeDriverLicenseResponseBody) *RecognizeDriverLicenseResponse {
	s.Body = v
	return s
}

type RecognizeDrivingLicenseRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Side     *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeDrivingLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseRequest) GoString() string {
	return s.String()
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s *RecognizeDrivingLicenseAdvanceRequest) SetSide(v string) *RecognizeDrivingLicenseAdvanceRequest {
	s.Side = &v
	return s
}

type RecognizeDrivingLicenseResponseBody struct {
	Data      *RecognizeDrivingLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeDrivingLicenseResponseBody) SetData(v *RecognizeDrivingLicenseResponseBodyData) *RecognizeDrivingLicenseResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeDrivingLicenseResponseBody) SetRequestId(v string) *RecognizeDrivingLicenseResponseBody {
	s.RequestId = &v
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
	ApprovedLoad              *string `json:"ApprovedLoad,omitempty" xml:"ApprovedLoad,omitempty"`
	ApprovedPassengerCapacity *string `json:"ApprovedPassengerCapacity,omitempty" xml:"ApprovedPassengerCapacity,omitempty"`
	EnergyType                *string `json:"EnergyType,omitempty" xml:"EnergyType,omitempty"`
	FileNumber                *string `json:"FileNumber,omitempty" xml:"FileNumber,omitempty"`
	GrossMass                 *string `json:"GrossMass,omitempty" xml:"GrossMass,omitempty"`
	InspectionRecord          *string `json:"InspectionRecord,omitempty" xml:"InspectionRecord,omitempty"`
	OverallDimension          *string `json:"OverallDimension,omitempty" xml:"OverallDimension,omitempty"`
	PlateNumber               *string `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	TractionMass              *string `json:"TractionMass,omitempty" xml:"TractionMass,omitempty"`
	UnladenMass               *string `json:"UnladenMass,omitempty" xml:"UnladenMass,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBodyDataBackResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyDataBackResult) GoString() string {
	return s.String()
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

type RecognizeDrivingLicenseResponseBodyDataFaceResult struct {
	Address      *string `json:"Address,omitempty" xml:"Address,omitempty"`
	EngineNumber *string `json:"EngineNumber,omitempty" xml:"EngineNumber,omitempty"`
	IssueDate    *string `json:"IssueDate,omitempty" xml:"IssueDate,omitempty"`
	Model        *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Owner        *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PlateNumber  *string `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	RegisterDate *string `json:"RegisterDate,omitempty" xml:"RegisterDate,omitempty"`
	UseCharacter *string `json:"UseCharacter,omitempty" xml:"UseCharacter,omitempty"`
	VehicleType  *string `json:"VehicleType,omitempty" xml:"VehicleType,omitempty"`
	Vin          *string `json:"Vin,omitempty" xml:"Vin,omitempty"`
}

func (s RecognizeDrivingLicenseResponseBodyDataFaceResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeDrivingLicenseResponseBodyDataFaceResult) GoString() string {
	return s.String()
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

type RecognizeDrivingLicenseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeDrivingLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeDrivingLicenseResponse) SetStatusCode(v int32) *RecognizeDrivingLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeDrivingLicenseResponse) SetBody(v *RecognizeDrivingLicenseResponseBody) *RecognizeDrivingLicenseResponse {
	s.Body = v
	return s
}

type RecognizeIdentityCardRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Side     *string `json:"Side,omitempty" xml:"Side,omitempty"`
}

func (s RecognizeIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardRequest) GoString() string {
	return s.String()
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s *RecognizeIdentityCardAdvanceRequest) SetSide(v string) *RecognizeIdentityCardAdvanceRequest {
	s.Side = &v
	return s
}

type RecognizeIdentityCardResponseBody struct {
	Data      *RecognizeIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBody) SetData(v *RecognizeIdentityCardResponseBodyData) *RecognizeIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeIdentityCardResponseBody) SetRequestId(v string) *RecognizeIdentityCardResponseBody {
	s.RequestId = &v
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
	Address          *string                                                             `json:"Address,omitempty" xml:"Address,omitempty"`
	BirthDate        *string                                                             `json:"BirthDate,omitempty" xml:"BirthDate,omitempty"`
	CardAreas        []*RecognizeIdentityCardResponseBodyDataFrontResultCardAreas        `json:"CardAreas,omitempty" xml:"CardAreas,omitempty" type:"Repeated"`
	FaceRectVertices []*RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices `json:"FaceRectVertices,omitempty" xml:"FaceRectVertices,omitempty" type:"Repeated"`
	FaceRectangle    *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle      `json:"FaceRectangle,omitempty" xml:"FaceRectangle,omitempty" type:"Struct"`
	Gender           *string                                                             `json:"Gender,omitempty" xml:"Gender,omitempty"`
	IDNumber         *string                                                             `json:"IDNumber,omitempty" xml:"IDNumber,omitempty"`
	Name             *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Nationality      *string                                                             `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResult) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResult) GoString() string {
	return s.String()
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

type RecognizeIdentityCardResponseBodyDataFrontResultCardAreas struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas {
	s.X = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultCardAreas {
	s.Y = &v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices {
	s.X = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectVertices {
	s.Y = &v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle struct {
	Angle  *float32                                                             `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Center *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter `json:"Center,omitempty" xml:"Center,omitempty" type:"Struct"`
	Size   *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize   `json:"Size,omitempty" xml:"Size,omitempty" type:"Struct"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangle) GoString() string {
	return s.String()
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

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) SetX(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter {
	s.X = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter) SetY(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleCenter {
	s.Y = &v
	return s
}

type RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize struct {
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) SetHeight(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize {
	s.Height = &v
	return s
}

func (s *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize) SetWidth(v float32) *RecognizeIdentityCardResponseBodyDataFrontResultFaceRectangleSize {
	s.Width = &v
	return s
}

type RecognizeIdentityCardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeIdentityCardResponse) SetStatusCode(v int32) *RecognizeIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIdentityCardResponse) SetBody(v *RecognizeIdentityCardResponseBody) *RecognizeIdentityCardResponse {
	s.Body = v
	return s
}

type RecognizeIndonesiaIdentityCardRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardRequest) SetImageUrl(v string) *RecognizeIndonesiaIdentityCardRequest {
	s.ImageUrl = &v
	return s
}

type RecognizeIndonesiaIdentityCardAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardAdvanceRequest) SetImageUrlObject(v io.Reader) *RecognizeIndonesiaIdentityCardAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBody struct {
	Data      *RecognizeIndonesiaIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBody) SetData(v *RecognizeIndonesiaIdentityCardResponseBodyData) *RecognizeIndonesiaIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBody) SetRequestId(v string) *RecognizeIndonesiaIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyData struct {
	AddressFifthLine  *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine  `json:"AddressFifthLine,omitempty" xml:"AddressFifthLine,omitempty" type:"Struct"`
	AddressFirstLine  *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine  `json:"AddressFirstLine,omitempty" xml:"AddressFirstLine,omitempty" type:"Struct"`
	AddressFourthLine *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine `json:"AddressFourthLine,omitempty" xml:"AddressFourthLine,omitempty" type:"Struct"`
	AddressSecondLine *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine `json:"AddressSecondLine,omitempty" xml:"AddressSecondLine,omitempty" type:"Struct"`
	AddressThirdLine  *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine  `json:"AddressThirdLine,omitempty" xml:"AddressThirdLine,omitempty" type:"Struct"`
	BirthDate         *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate         `json:"BirthDate,omitempty" xml:"BirthDate,omitempty" type:"Struct"`
	BirthPlace        *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace        `json:"BirthPlace,omitempty" xml:"BirthPlace,omitempty" type:"Struct"`
	CardBox           *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox           `json:"CardBox,omitempty" xml:"CardBox,omitempty" type:"Struct"`
	ExpiryDate        *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate        `json:"ExpiryDate,omitempty" xml:"ExpiryDate,omitempty" type:"Struct"`
	Gender            *RecognizeIndonesiaIdentityCardResponseBodyDataGender            `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	Height            *RecognizeIndonesiaIdentityCardResponseBodyDataHeight            `json:"Height,omitempty" xml:"Height,omitempty" type:"Struct"`
	IdNumber          *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber          `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" type:"Struct"`
	LicenseNumber     *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber     `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty" type:"Struct"`
	MaritalStatus     *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus     `json:"MaritalStatus,omitempty" xml:"MaritalStatus,omitempty" type:"Struct"`
	NameFirstLine     *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine     `json:"NameFirstLine,omitempty" xml:"NameFirstLine,omitempty" type:"Struct"`
	NameSecondLine    *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine    `json:"NameSecondLine,omitempty" xml:"NameSecondLine,omitempty" type:"Struct"`
	Nationality       *RecognizeIndonesiaIdentityCardResponseBodyDataNationality       `json:"Nationality,omitempty" xml:"Nationality,omitempty" type:"Struct"`
	Occupation        *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation        `json:"Occupation,omitempty" xml:"Occupation,omitempty" type:"Struct"`
	PortraitBox       *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox       `json:"PortraitBox,omitempty" xml:"PortraitBox,omitempty" type:"Struct"`
	Province          *RecognizeIndonesiaIdentityCardResponseBodyDataProvince          `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	Religion          *RecognizeIndonesiaIdentityCardResponseBodyDataReligion          `json:"Religion,omitempty" xml:"Religion,omitempty" type:"Struct"`
	Sex               *RecognizeIndonesiaIdentityCardResponseBodyDataSex               `json:"Sex,omitempty" xml:"Sex,omitempty" type:"Struct"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetAddressFifthLine(v *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.AddressFifthLine = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetAddressFirstLine(v *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.AddressFirstLine = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetAddressFourthLine(v *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.AddressFourthLine = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetAddressSecondLine(v *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.AddressSecondLine = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetAddressThirdLine(v *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.AddressThirdLine = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetBirthDate(v *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.BirthDate = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetBirthPlace(v *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.BirthPlace = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetCardBox(v *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.CardBox = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetExpiryDate(v *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.ExpiryDate = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetGender(v *RecognizeIndonesiaIdentityCardResponseBodyDataGender) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.Gender = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetHeight(v *RecognizeIndonesiaIdentityCardResponseBodyDataHeight) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.Height = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetIdNumber(v *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.IdNumber = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetLicenseNumber(v *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.LicenseNumber = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetMaritalStatus(v *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.MaritalStatus = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetNameFirstLine(v *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.NameFirstLine = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetNameSecondLine(v *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.NameSecondLine = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetNationality(v *RecognizeIndonesiaIdentityCardResponseBodyDataNationality) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.Nationality = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetOccupation(v *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.Occupation = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetPortraitBox(v *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.PortraitBox = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetProvince(v *RecognizeIndonesiaIdentityCardResponseBodyDataProvince) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.Province = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetReligion(v *RecognizeIndonesiaIdentityCardResponseBodyDataReligion) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.Religion = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyData) SetSex(v *RecognizeIndonesiaIdentityCardResponseBodyDataSex) *RecognizeIndonesiaIdentityCardResponseBodyData {
	s.Sex = v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine) SetScore(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLine {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine) SetScore(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLine {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine) SetScore(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLine {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine) SetScore(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLine {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine) SetScore(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLine {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDate {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlace {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataBirthPlaceKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataCardBox struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataCardBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataCardBox) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataCardBox {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDate {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataGender struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataGender) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataGender) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataGender) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataGender {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataGender) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataGender {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataGender) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataGender {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataGenderKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataHeight struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataHeight) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataHeight) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataHeight) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataHeight {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataHeight) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataHeight {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataHeight) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataHeight {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataHeightKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumber {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                 `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumber {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataLicenseNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                 `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatus {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataMaritalStatusKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                 `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLine {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataNameFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLine {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataNameSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataNationality struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                               `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNationality) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNationality) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNationality) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataNationality {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNationality) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataNationality {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNationality) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataNationality {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataNationalityKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataOccupation struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataOccupation) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataOccupation) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataOccupation {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataOccupationKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                               `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBox {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataProvince struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataProvince) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataProvince) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataProvince) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataProvince {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataProvince) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataProvince {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataProvince) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataProvince {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataProvinceKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataReligion struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataReligion) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataReligion) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataReligion) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataReligion {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataReligion) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataReligion {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataReligion) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataReligion {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataReligionKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataSex struct {
	KeyPoints []*RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                       `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataSex) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataSex) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataSex) SetKeyPoints(v []*RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints) *RecognizeIndonesiaIdentityCardResponseBodyDataSex {
	s.KeyPoints = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataSex) SetScore(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataSex {
	s.Score = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataSex) SetText(v string) *RecognizeIndonesiaIdentityCardResponseBodyDataSex {
	s.Text = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints) SetX(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints) SetY(v float32) *RecognizeIndonesiaIdentityCardResponseBodyDataSexKeyPoints {
	s.Y = &v
	return s
}

type RecognizeIndonesiaIdentityCardResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeIndonesiaIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeIndonesiaIdentityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIndonesiaIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIndonesiaIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeIndonesiaIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponse) SetStatusCode(v int32) *RecognizeIndonesiaIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIndonesiaIdentityCardResponse) SetBody(v *RecognizeIndonesiaIdentityCardResponseBody) *RecognizeIndonesiaIdentityCardResponse {
	s.Body = v
	return s
}

type RecognizeLicensePlateRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeLicensePlateRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateRequest) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateRequest) SetImageURL(v string) *RecognizeLicensePlateRequest {
	s.ImageURL = &v
	return s
}

type RecognizeLicensePlateAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeLicensePlateResponseBody struct {
	Data      *RecognizeLicensePlateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeLicensePlateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBody) SetData(v *RecognizeLicensePlateResponseBodyData) *RecognizeLicensePlateResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeLicensePlateResponseBody) SetRequestId(v string) *RecognizeLicensePlateResponseBody {
	s.RequestId = &v
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
	Confidence          *float32                                                `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	PlateNumber         *string                                                 `json:"PlateNumber,omitempty" xml:"PlateNumber,omitempty"`
	PlateType           *string                                                 `json:"PlateType,omitempty" xml:"PlateType,omitempty"`
	PlateTypeConfidence *float32                                                `json:"PlateTypeConfidence,omitempty" xml:"PlateTypeConfidence,omitempty"`
	Positions           []*RecognizeLicensePlateResponseBodyDataPlatesPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
	Roi                 *RecognizeLicensePlateResponseBodyDataPlatesRoi         `json:"Roi,omitempty" xml:"Roi,omitempty" type:"Struct"`
}

func (s RecognizeLicensePlateResponseBodyDataPlates) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlates) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetConfidence(v float32) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Confidence = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateNumber(v string) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateNumber = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateType(v string) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateType = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPlateTypeConfidence(v float32) *RecognizeLicensePlateResponseBodyDataPlates {
	s.PlateTypeConfidence = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetPositions(v []*RecognizeLicensePlateResponseBodyDataPlatesPositions) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Positions = v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlates) SetRoi(v *RecognizeLicensePlateResponseBodyDataPlatesRoi) *RecognizeLicensePlateResponseBodyDataPlates {
	s.Roi = v
	return s
}

type RecognizeLicensePlateResponseBodyDataPlatesPositions struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeLicensePlateResponseBodyDataPlatesPositions) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlatesPositions) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesPositions) SetX(v int64) *RecognizeLicensePlateResponseBodyDataPlatesPositions {
	s.X = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesPositions) SetY(v int64) *RecognizeLicensePlateResponseBodyDataPlatesPositions {
	s.Y = &v
	return s
}

type RecognizeLicensePlateResponseBodyDataPlatesRoi struct {
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeLicensePlateResponseBodyDataPlatesRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeLicensePlateResponseBodyDataPlatesRoi) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetH(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.H = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetW(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.W = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetX(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.X = &v
	return s
}

func (s *RecognizeLicensePlateResponseBodyDataPlatesRoi) SetY(v int32) *RecognizeLicensePlateResponseBodyDataPlatesRoi {
	s.Y = &v
	return s
}

type RecognizeLicensePlateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeLicensePlateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeLicensePlateResponse) SetStatusCode(v int32) *RecognizeLicensePlateResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeLicensePlateResponse) SetBody(v *RecognizeLicensePlateResponseBody) *RecognizeLicensePlateResponse {
	s.Body = v
	return s
}

type RecognizeMalaysiaIdentityCardRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardRequest) SetImageUrl(v string) *RecognizeMalaysiaIdentityCardRequest {
	s.ImageUrl = &v
	return s
}

type RecognizeMalaysiaIdentityCardAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardAdvanceRequest) SetImageUrlObject(v io.Reader) *RecognizeMalaysiaIdentityCardAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBody struct {
	Data      *RecognizeMalaysiaIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBody) SetData(v *RecognizeMalaysiaIdentityCardResponseBodyData) *RecognizeMalaysiaIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBody) SetRequestId(v string) *RecognizeMalaysiaIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyData struct {
	AddressFifthLine  *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine  `json:"AddressFifthLine,omitempty" xml:"AddressFifthLine,omitempty" type:"Struct"`
	AddressFirstLine  *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine  `json:"AddressFirstLine,omitempty" xml:"AddressFirstLine,omitempty" type:"Struct"`
	AddressFourthLine *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine `json:"AddressFourthLine,omitempty" xml:"AddressFourthLine,omitempty" type:"Struct"`
	AddressSecondLine *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine `json:"AddressSecondLine,omitempty" xml:"AddressSecondLine,omitempty" type:"Struct"`
	AddressThirdLine  *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine  `json:"AddressThirdLine,omitempty" xml:"AddressThirdLine,omitempty" type:"Struct"`
	CardBox           *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox           `json:"CardBox,omitempty" xml:"CardBox,omitempty" type:"Struct"`
	DriveClass        *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass        `json:"DriveClass,omitempty" xml:"DriveClass,omitempty" type:"Struct"`
	ExpiryDate        *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate        `json:"ExpiryDate,omitempty" xml:"ExpiryDate,omitempty" type:"Struct"`
	IdNumber          *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber          `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" type:"Struct"`
	IssueDate         *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate         `json:"IssueDate,omitempty" xml:"IssueDate,omitempty" type:"Struct"`
	NameFirstLine     *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine     `json:"NameFirstLine,omitempty" xml:"NameFirstLine,omitempty" type:"Struct"`
	NameSecondLine    *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine    `json:"NameSecondLine,omitempty" xml:"NameSecondLine,omitempty" type:"Struct"`
	Nationality       *RecognizeMalaysiaIdentityCardResponseBodyDataNationality       `json:"Nationality,omitempty" xml:"Nationality,omitempty" type:"Struct"`
	PortraitBox       *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox       `json:"PortraitBox,omitempty" xml:"PortraitBox,omitempty" type:"Struct"`
	Sex               *RecognizeMalaysiaIdentityCardResponseBodyDataSex               `json:"Sex,omitempty" xml:"Sex,omitempty" type:"Struct"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetAddressFifthLine(v *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.AddressFifthLine = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetAddressFirstLine(v *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.AddressFirstLine = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetAddressFourthLine(v *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.AddressFourthLine = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetAddressSecondLine(v *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.AddressSecondLine = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetAddressThirdLine(v *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.AddressThirdLine = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetCardBox(v *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.CardBox = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetDriveClass(v *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.DriveClass = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetExpiryDate(v *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.ExpiryDate = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetIdNumber(v *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.IdNumber = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetIssueDate(v *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.IssueDate = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetNameFirstLine(v *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.NameFirstLine = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetNameSecondLine(v *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.NameSecondLine = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetNationality(v *RecognizeMalaysiaIdentityCardResponseBodyDataNationality) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.Nationality = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetPortraitBox(v *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.PortraitBox = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyData) SetSex(v *RecognizeMalaysiaIdentityCardResponseBodyDataSex) *RecognizeMalaysiaIdentityCardResponseBodyData {
	s.Sex = v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine) SetScore(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLine {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFifthLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine) SetScore(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLine {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine) SetScore(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLine {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressFourthLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine) SetScore(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLine {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine) SetScore(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLine {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataAddressThirdLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataCardBox struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataCardBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataCardBox) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataCardBox {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClass {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataDriveClassKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDate {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumber {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDate {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataIssueDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLine {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataNameFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                 `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLine {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataNameSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataNationality struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNationality) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNationality) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNationality) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataNationality {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNationality) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataNationality {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNationality) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataNationality {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataNationalityKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBox {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataSex struct {
	KeyPoints []*RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                      `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                      `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataSex) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataSex) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataSex) SetKeyPoints(v []*RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints) *RecognizeMalaysiaIdentityCardResponseBodyDataSex {
	s.KeyPoints = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataSex) SetScore(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataSex {
	s.Score = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataSex) SetText(v string) *RecognizeMalaysiaIdentityCardResponseBodyDataSex {
	s.Text = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints) SetX(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints) SetY(v float32) *RecognizeMalaysiaIdentityCardResponseBodyDataSexKeyPoints {
	s.Y = &v
	return s
}

type RecognizeMalaysiaIdentityCardResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeMalaysiaIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeMalaysiaIdentityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeMalaysiaIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeMalaysiaIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeMalaysiaIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponse) SetStatusCode(v int32) *RecognizeMalaysiaIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeMalaysiaIdentityCardResponse) SetBody(v *RecognizeMalaysiaIdentityCardResponseBody) *RecognizeMalaysiaIdentityCardResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *RecognizePassportMRZResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePassportMRZResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZResponseBody) SetData(v *RecognizePassportMRZResponseBodyData) *RecognizePassportMRZResponseBody {
	s.Data = v
	return s
}

func (s *RecognizePassportMRZResponseBody) SetRequestId(v string) *RecognizePassportMRZResponseBody {
	s.RequestId = &v
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
	BandBoxes        []*float32 `json:"BandBoxes,omitempty" xml:"BandBoxes,omitempty" type:"Repeated"`
	Content          *string    `json:"Content,omitempty" xml:"Content,omitempty"`
	DetectionScore   *float32   `json:"DetectionScore,omitempty" xml:"DetectionScore,omitempty"`
	Name             *string    `json:"Name,omitempty" xml:"Name,omitempty"`
	RecognitionScore *float32   `json:"RecognitionScore,omitempty" xml:"RecognitionScore,omitempty"`
}

func (s RecognizePassportMRZResponseBodyDataRegions) String() string {
	return tea.Prettify(s)
}

func (s RecognizePassportMRZResponseBodyDataRegions) GoString() string {
	return s.String()
}

func (s *RecognizePassportMRZResponseBodyDataRegions) SetBandBoxes(v []*float32) *RecognizePassportMRZResponseBodyDataRegions {
	s.BandBoxes = v
	return s
}

func (s *RecognizePassportMRZResponseBodyDataRegions) SetContent(v string) *RecognizePassportMRZResponseBodyDataRegions {
	s.Content = &v
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

func (s *RecognizePassportMRZResponseBodyDataRegions) SetRecognitionScore(v float32) *RecognizePassportMRZResponseBodyDataRegions {
	s.RecognitionScore = &v
	return s
}

type RecognizePassportMRZResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizePassportMRZResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizePassportMRZResponse) SetStatusCode(v int32) *RecognizePassportMRZResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePassportMRZResponse) SetBody(v *RecognizePassportMRZResponseBody) *RecognizePassportMRZResponse {
	s.Body = v
	return s
}

type RecognizePdfRequest struct {
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizePdfRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePdfRequest) GoString() string {
	return s.String()
}

func (s *RecognizePdfRequest) SetFileURL(v string) *RecognizePdfRequest {
	s.FileURL = &v
	return s
}

type RecognizePdfAdvanceRequest struct {
	FileURLObject io.Reader `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizePdfAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizePdfAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizePdfAdvanceRequest) SetFileURLObject(v io.Reader) *RecognizePdfAdvanceRequest {
	s.FileURLObject = v
	return s
}

type RecognizePdfResponseBody struct {
	Data      *RecognizePdfResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePdfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePdfResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponseBody) SetData(v *RecognizePdfResponseBodyData) *RecognizePdfResponseBody {
	s.Data = v
	return s
}

func (s *RecognizePdfResponseBody) SetRequestId(v string) *RecognizePdfResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePdfResponseBodyData struct {
	Angle     *int64                                   `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Height    *int64                                   `json:"Height,omitempty" xml:"Height,omitempty"`
	OrgHeight *int64                                   `json:"OrgHeight,omitempty" xml:"OrgHeight,omitempty"`
	OrgWidth  *int64                                   `json:"OrgWidth,omitempty" xml:"OrgWidth,omitempty"`
	PageIndex *int64                                   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	Width     *int64                                   `json:"Width,omitempty" xml:"Width,omitempty"`
	WordsInfo []*RecognizePdfResponseBodyDataWordsInfo `json:"WordsInfo,omitempty" xml:"WordsInfo,omitempty" type:"Repeated"`
}

func (s RecognizePdfResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizePdfResponseBodyData) GoString() string {
	return s.String()
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

type RecognizePdfResponseBodyDataWordsInfo struct {
	Angle     *int64                                            `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Height    *int64                                            `json:"Height,omitempty" xml:"Height,omitempty"`
	Positions []*RecognizePdfResponseBodyDataWordsInfoPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
	Width     *int64                                            `json:"Width,omitempty" xml:"Width,omitempty"`
	Word      *string                                           `json:"Word,omitempty" xml:"Word,omitempty"`
	X         *int64                                            `json:"X,omitempty" xml:"X,omitempty"`
	Y         *int64                                            `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizePdfResponseBodyDataWordsInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizePdfResponseBodyDataWordsInfo) GoString() string {
	return s.String()
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

type RecognizePdfResponseBodyDataWordsInfoPositions struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizePdfResponseBodyDataWordsInfoPositions) String() string {
	return tea.Prettify(s)
}

func (s RecognizePdfResponseBodyDataWordsInfoPositions) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponseBodyDataWordsInfoPositions) SetX(v int64) *RecognizePdfResponseBodyDataWordsInfoPositions {
	s.X = &v
	return s
}

func (s *RecognizePdfResponseBodyDataWordsInfoPositions) SetY(v int64) *RecognizePdfResponseBodyDataWordsInfoPositions {
	s.Y = &v
	return s
}

type RecognizePdfResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizePdfResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizePdfResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizePdfResponse) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponse) SetHeaders(v map[string]*string) *RecognizePdfResponse {
	s.Headers = v
	return s
}

func (s *RecognizePdfResponse) SetStatusCode(v int32) *RecognizePdfResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePdfResponse) SetBody(v *RecognizePdfResponseBody) *RecognizePdfResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *RecognizePoiNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePoiNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBody) SetData(v *RecognizePoiNameResponseBodyData) *RecognizePoiNameResponseBody {
	s.Data = v
	return s
}

func (s *RecognizePoiNameResponseBody) SetRequestId(v string) *RecognizePoiNameResponseBody {
	s.RequestId = &v
	return s
}

type RecognizePoiNameResponseBodyData struct {
	Signboards []*RecognizePoiNameResponseBodyDataSignboards `json:"Signboards,omitempty" xml:"Signboards,omitempty" type:"Repeated"`
	Summary    *RecognizePoiNameResponseBodyDataSummary      `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
}

func (s RecognizePoiNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBodyData) SetSignboards(v []*RecognizePoiNameResponseBodyDataSignboards) *RecognizePoiNameResponseBodyData {
	s.Signboards = v
	return s
}

func (s *RecognizePoiNameResponseBodyData) SetSummary(v *RecognizePoiNameResponseBodyDataSummary) *RecognizePoiNameResponseBodyData {
	s.Summary = v
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
	Label  *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Points []*int32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
	Score  *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Tag    *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizePoiNameResponseBodyDataSignboardsTexts) String() string {
	return tea.Prettify(s)
}

func (s RecognizePoiNameResponseBodyDataSignboardsTexts) GoString() string {
	return s.String()
}

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetLabel(v string) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Label = &v
	return s
}

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetPoints(v []*int32) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Points = v
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

func (s *RecognizePoiNameResponseBodyDataSignboardsTexts) SetType(v string) *RecognizePoiNameResponseBodyDataSignboardsTexts {
	s.Type = &v
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

type RecognizePoiNameResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizePoiNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizePoiNameResponse) SetStatusCode(v int32) *RecognizePoiNameResponse {
	s.StatusCode = &v
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
	Data      *RecognizeQrCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeQrCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBody) SetData(v *RecognizeQrCodeResponseBodyData) *RecognizeQrCodeResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeQrCodeResponseBody) SetRequestId(v string) *RecognizeQrCodeResponseBody {
	s.RequestId = &v
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
	Results  []*RecognizeQrCodeResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	TaskId   *string                                           `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *RecognizeQrCodeResponseBodyDataElements) SetResults(v []*RecognizeQrCodeResponseBodyDataElementsResults) *RecognizeQrCodeResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElements) SetTaskId(v string) *RecognizeQrCodeResponseBodyDataElements {
	s.TaskId = &v
	return s
}

type RecognizeQrCodeResponseBodyDataElementsResults struct {
	Label       *string   `json:"Label,omitempty" xml:"Label,omitempty"`
	QrCodesData []*string `json:"QrCodesData,omitempty" xml:"QrCodesData,omitempty" type:"Repeated"`
	Rate        *float32  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Suggestion  *string   `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s RecognizeQrCodeResponseBodyDataElementsResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQrCodeResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetLabel(v string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetQrCodesData(v []*string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.QrCodesData = v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetRate(v float32) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *RecognizeQrCodeResponseBodyDataElementsResults) SetSuggestion(v string) *RecognizeQrCodeResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

type RecognizeQrCodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeQrCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeQrCodeResponse) SetStatusCode(v int32) *RecognizeQrCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeQrCodeResponse) SetBody(v *RecognizeQrCodeResponseBody) *RecognizeQrCodeResponse {
	s.Body = v
	return s
}

type RecognizeQuotaInvoiceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeQuotaInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceRequest) SetImageURL(v string) *RecognizeQuotaInvoiceRequest {
	s.ImageURL = &v
	return s
}

type RecognizeQuotaInvoiceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeQuotaInvoiceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeQuotaInvoiceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeQuotaInvoiceResponseBody struct {
	Data      *RecognizeQuotaInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBody) SetData(v *RecognizeQuotaInvoiceResponseBodyData) *RecognizeQuotaInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBody) SetRequestId(v string) *RecognizeQuotaInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeQuotaInvoiceResponseBodyData struct {
	Angle         *int64                                                `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Content       *RecognizeQuotaInvoiceResponseBodyDataContent         `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Height        *int64                                                `json:"Height,omitempty" xml:"Height,omitempty"`
	KeyValueInfos []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos `json:"KeyValueInfos,omitempty" xml:"KeyValueInfos,omitempty" type:"Repeated"`
	OrgHeight     *int64                                                `json:"OrgHeight,omitempty" xml:"OrgHeight,omitempty"`
	OrgWidth      *int64                                                `json:"OrgWidth,omitempty" xml:"OrgWidth,omitempty"`
	Width         *int64                                                `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetAngle(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.Angle = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetContent(v *RecognizeQuotaInvoiceResponseBodyDataContent) *RecognizeQuotaInvoiceResponseBodyData {
	s.Content = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetHeight(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetKeyValueInfos(v []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) *RecognizeQuotaInvoiceResponseBodyData {
	s.KeyValueInfos = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetOrgHeight(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.OrgHeight = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetOrgWidth(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.OrgWidth = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyData) SetWidth(v int64) *RecognizeQuotaInvoiceResponseBodyData {
	s.Width = &v
	return s
}

type RecognizeQuotaInvoiceResponseBodyDataContent struct {
	InvoiceAmount  *string `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	InvoiceCode    *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceDetails *string `json:"InvoiceDetails,omitempty" xml:"InvoiceDetails,omitempty"`
	InvoiceNo      *string `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	SumAmount      *string `json:"SumAmount,omitempty" xml:"SumAmount,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceAmount(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceAmount = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceCode(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceCode = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceDetails(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceDetails = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetInvoiceNo(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.InvoiceNo = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataContent) SetSumAmount(v string) *RecognizeQuotaInvoiceResponseBodyDataContent {
	s.SumAmount = &v
	return s
}

type RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos struct {
	Key            *string                                                             `json:"Key,omitempty" xml:"Key,omitempty"`
	Value          *string                                                             `json:"Value,omitempty" xml:"Value,omitempty"`
	ValuePositions []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions `json:"ValuePositions,omitempty" xml:"ValuePositions,omitempty" type:"Repeated"`
	ValueScore     *float32                                                            `json:"ValueScore,omitempty" xml:"ValueScore,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetKey(v string) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.Key = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetValue(v string) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.Value = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetValuePositions(v []*RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.ValuePositions = v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos) SetValueScore(v float32) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfos {
	s.ValueScore = &v
	return s
}

type RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) SetX(v int64) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions {
	s.X = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions) SetY(v int64) *RecognizeQuotaInvoiceResponseBodyDataKeyValueInfosValuePositions {
	s.Y = &v
	return s
}

type RecognizeQuotaInvoiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeQuotaInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeQuotaInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeQuotaInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeQuotaInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeQuotaInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetStatusCode(v int32) *RecognizeQuotaInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeQuotaInvoiceResponse) SetBody(v *RecognizeQuotaInvoiceResponseBody) *RecognizeQuotaInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeRussiaIdentityCardRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeRussiaIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardRequest) SetImageUrl(v string) *RecognizeRussiaIdentityCardRequest {
	s.ImageUrl = &v
	return s
}

type RecognizeRussiaIdentityCardAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeRussiaIdentityCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardAdvanceRequest) SetImageUrlObject(v io.Reader) *RecognizeRussiaIdentityCardAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type RecognizeRussiaIdentityCardResponseBody struct {
	Data      *RecognizeRussiaIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBody) SetData(v *RecognizeRussiaIdentityCardResponseBodyData) *RecognizeRussiaIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBody) SetRequestId(v string) *RecognizeRussiaIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyData struct {
	BirthDate            *RecognizeRussiaIdentityCardResponseBodyDataBirthDate            `json:"BirthDate,omitempty" xml:"BirthDate,omitempty" type:"Struct"`
	BirthPlaceFirstLine  *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine  `json:"BirthPlaceFirstLine,omitempty" xml:"BirthPlaceFirstLine,omitempty" type:"Struct"`
	BirthPlaceSecondLine *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine `json:"BirthPlaceSecondLine,omitempty" xml:"BirthPlaceSecondLine,omitempty" type:"Struct"`
	BirthPlaceThirdLine  *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine  `json:"BirthPlaceThirdLine,omitempty" xml:"BirthPlaceThirdLine,omitempty" type:"Struct"`
	CardBox              *RecognizeRussiaIdentityCardResponseBodyDataCardBox              `json:"CardBox,omitempty" xml:"CardBox,omitempty" type:"Struct"`
	GivenName            *RecognizeRussiaIdentityCardResponseBodyDataGivenName            `json:"GivenName,omitempty" xml:"GivenName,omitempty" type:"Struct"`
	IdNumber             *RecognizeRussiaIdentityCardResponseBodyDataIdNumber             `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" type:"Struct"`
	PaternalName         *RecognizeRussiaIdentityCardResponseBodyDataPaternalName         `json:"PaternalName,omitempty" xml:"PaternalName,omitempty" type:"Struct"`
	PortraitBox          *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox          `json:"PortraitBox,omitempty" xml:"PortraitBox,omitempty" type:"Struct"`
	Sex                  *RecognizeRussiaIdentityCardResponseBodyDataSex                  `json:"Sex,omitempty" xml:"Sex,omitempty" type:"Struct"`
	SurnameFirstLine     *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine     `json:"SurnameFirstLine,omitempty" xml:"SurnameFirstLine,omitempty" type:"Struct"`
	SurnameSecondLine    *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine    `json:"SurnameSecondLine,omitempty" xml:"SurnameSecondLine,omitempty" type:"Struct"`
}

func (s RecognizeRussiaIdentityCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetBirthDate(v *RecognizeRussiaIdentityCardResponseBodyDataBirthDate) *RecognizeRussiaIdentityCardResponseBodyData {
	s.BirthDate = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetBirthPlaceFirstLine(v *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine) *RecognizeRussiaIdentityCardResponseBodyData {
	s.BirthPlaceFirstLine = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetBirthPlaceSecondLine(v *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine) *RecognizeRussiaIdentityCardResponseBodyData {
	s.BirthPlaceSecondLine = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetBirthPlaceThirdLine(v *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine) *RecognizeRussiaIdentityCardResponseBodyData {
	s.BirthPlaceThirdLine = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetCardBox(v *RecognizeRussiaIdentityCardResponseBodyDataCardBox) *RecognizeRussiaIdentityCardResponseBodyData {
	s.CardBox = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetGivenName(v *RecognizeRussiaIdentityCardResponseBodyDataGivenName) *RecognizeRussiaIdentityCardResponseBodyData {
	s.GivenName = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetIdNumber(v *RecognizeRussiaIdentityCardResponseBodyDataIdNumber) *RecognizeRussiaIdentityCardResponseBodyData {
	s.IdNumber = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetPaternalName(v *RecognizeRussiaIdentityCardResponseBodyDataPaternalName) *RecognizeRussiaIdentityCardResponseBodyData {
	s.PaternalName = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetPortraitBox(v *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox) *RecognizeRussiaIdentityCardResponseBodyData {
	s.PortraitBox = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetSex(v *RecognizeRussiaIdentityCardResponseBodyDataSex) *RecognizeRussiaIdentityCardResponseBodyData {
	s.Sex = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetSurnameFirstLine(v *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine) *RecognizeRussiaIdentityCardResponseBodyData {
	s.SurnameFirstLine = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyData) SetSurnameSecondLine(v *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine) *RecognizeRussiaIdentityCardResponseBodyData {
	s.SurnameSecondLine = v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthDate struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthDate) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthDate) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataBirthDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthDate) SetScore(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthDate {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthDate) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataBirthDate {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine) SetScore(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLine {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine) SetScore(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLine {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine) SetScore(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLine {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataBirthPlaceThirdLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataCardBox struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataCardBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataCardBox) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataCardBox) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataCardBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataCardBox) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataCardBox {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataCardBox) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataCardBox {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataGivenName struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataGivenName) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataGivenName) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataGivenName) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataGivenName {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataGivenName) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataGivenName {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataGivenName) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataGivenName {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataGivenNameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataIdNumber struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                         `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataIdNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataIdNumber) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataIdNumber) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataIdNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataIdNumber) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataIdNumber {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataIdNumber) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataIdNumber {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataPaternalName struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPaternalName) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPaternalName) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPaternalName) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataPaternalName {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPaternalName) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataPaternalName {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPaternalName) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataPaternalName {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataPaternalNameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataPortraitBox struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPortraitBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPortraitBox) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataPortraitBox {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataSex struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSex) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSex) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSex) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataSex {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSex) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataSex {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSex) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataSex {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataSexKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                 `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLine {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataSurnameFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine struct {
	KeyPoints []*RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine) SetKeyPoints(v []*RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints) *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine) SetScore(v string) *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine) SetText(v string) *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLine {
	s.Text = &v
	return s
}

type RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints) SetX(v float32) *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints) SetY(v float32) *RecognizeRussiaIdentityCardResponseBodyDataSurnameSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeRussiaIdentityCardResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeRussiaIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeRussiaIdentityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeRussiaIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeRussiaIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeRussiaIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeRussiaIdentityCardResponse) SetStatusCode(v int32) *RecognizeRussiaIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeRussiaIdentityCardResponse) SetBody(v *RecognizeRussiaIdentityCardResponseBody) *RecognizeRussiaIdentityCardResponse {
	s.Body = v
	return s
}

type RecognizeStampRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeStampRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampRequest) GoString() string {
	return s.String()
}

func (s *RecognizeStampRequest) SetImageURL(v string) *RecognizeStampRequest {
	s.ImageURL = &v
	return s
}

type RecognizeStampAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeStampResponseBody struct {
	Data      *RecognizeStampResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeStampResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBody) SetData(v *RecognizeStampResponseBodyData) *RecognizeStampResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeStampResponseBody) SetRequestId(v string) *RecognizeStampResponseBody {
	s.RequestId = &v
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
	GeneralText []*RecognizeStampResponseBodyDataResultsGeneralText `json:"GeneralText,omitempty" xml:"GeneralText,omitempty" type:"Repeated"`
	Roi         *RecognizeStampResponseBodyDataResultsRoi           `json:"Roi,omitempty" xml:"Roi,omitempty" type:"Struct"`
	Text        *RecognizeStampResponseBodyDataResultsText          `json:"Text,omitempty" xml:"Text,omitempty" type:"Struct"`
}

func (s RecognizeStampResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBodyDataResults) SetGeneralText(v []*RecognizeStampResponseBodyDataResultsGeneralText) *RecognizeStampResponseBodyDataResults {
	s.GeneralText = v
	return s
}

func (s *RecognizeStampResponseBodyDataResults) SetRoi(v *RecognizeStampResponseBodyDataResultsRoi) *RecognizeStampResponseBodyDataResults {
	s.Roi = v
	return s
}

func (s *RecognizeStampResponseBodyDataResults) SetText(v *RecognizeStampResponseBodyDataResultsText) *RecognizeStampResponseBodyDataResults {
	s.Text = v
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

type RecognizeStampResponseBodyDataResultsRoi struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeStampResponseBodyDataResultsRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeStampResponseBodyDataResultsRoi) GoString() string {
	return s.String()
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetHeight(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Height = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetLeft(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Left = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetTop(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Top = &v
	return s
}

func (s *RecognizeStampResponseBodyDataResultsRoi) SetWidth(v int32) *RecognizeStampResponseBodyDataResultsRoi {
	s.Width = &v
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

type RecognizeStampResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeStampResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeStampResponse) SetStatusCode(v int32) *RecognizeStampResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeStampResponse) SetBody(v *RecognizeStampResponseBody) *RecognizeStampResponse {
	s.Body = v
	return s
}

type RecognizeTableRequest struct {
	AssureDirection *bool   `json:"AssureDirection,omitempty" xml:"AssureDirection,omitempty"`
	HasLine         *bool   `json:"HasLine,omitempty" xml:"HasLine,omitempty"`
	ImageURL        *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutputFormat    *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	SkipDetection   *bool   `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
	UseFinanceModel *bool   `json:"UseFinanceModel,omitempty" xml:"UseFinanceModel,omitempty"`
}

func (s RecognizeTableRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableRequest) SetAssureDirection(v bool) *RecognizeTableRequest {
	s.AssureDirection = &v
	return s
}

func (s *RecognizeTableRequest) SetHasLine(v bool) *RecognizeTableRequest {
	s.HasLine = &v
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

func (s *RecognizeTableRequest) SetSkipDetection(v bool) *RecognizeTableRequest {
	s.SkipDetection = &v
	return s
}

func (s *RecognizeTableRequest) SetUseFinanceModel(v bool) *RecognizeTableRequest {
	s.UseFinanceModel = &v
	return s
}

type RecognizeTableAdvanceRequest struct {
	AssureDirection *bool     `json:"AssureDirection,omitempty" xml:"AssureDirection,omitempty"`
	HasLine         *bool     `json:"HasLine,omitempty" xml:"HasLine,omitempty"`
	ImageURLObject  io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	OutputFormat    *string   `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	SkipDetection   *bool     `json:"SkipDetection,omitempty" xml:"SkipDetection,omitempty"`
	UseFinanceModel *bool     `json:"UseFinanceModel,omitempty" xml:"UseFinanceModel,omitempty"`
}

func (s RecognizeTableAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTableAdvanceRequest) SetAssureDirection(v bool) *RecognizeTableAdvanceRequest {
	s.AssureDirection = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetHasLine(v bool) *RecognizeTableAdvanceRequest {
	s.HasLine = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTableAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetOutputFormat(v string) *RecognizeTableAdvanceRequest {
	s.OutputFormat = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetSkipDetection(v bool) *RecognizeTableAdvanceRequest {
	s.SkipDetection = &v
	return s
}

func (s *RecognizeTableAdvanceRequest) SetUseFinanceModel(v bool) *RecognizeTableAdvanceRequest {
	s.UseFinanceModel = &v
	return s
}

type RecognizeTableResponseBody struct {
	Data      *RecognizeTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBody) SetData(v *RecognizeTableResponseBodyData) *RecognizeTableResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTableResponseBody) SetRequestId(v string) *RecognizeTableResponseBody {
	s.RequestId = &v
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
	TableRows []*RecognizeTableResponseBodyDataTablesTableRows `json:"TableRows,omitempty" xml:"TableRows,omitempty" type:"Repeated"`
	Tail      []*string                                        `json:"Tail,omitempty" xml:"Tail,omitempty" type:"Repeated"`
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

func (s *RecognizeTableResponseBodyDataTables) SetTableRows(v []*RecognizeTableResponseBodyDataTablesTableRows) *RecognizeTableResponseBodyDataTables {
	s.TableRows = v
	return s
}

func (s *RecognizeTableResponseBodyDataTables) SetTail(v []*string) *RecognizeTableResponseBodyDataTables {
	s.Tail = v
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
	EndColumn   *int32    `json:"EndColumn,omitempty" xml:"EndColumn,omitempty"`
	EndRow      *int32    `json:"EndRow,omitempty" xml:"EndRow,omitempty"`
	Height      *int32    `json:"Height,omitempty" xml:"Height,omitempty"`
	StartColumn *int32    `json:"StartColumn,omitempty" xml:"StartColumn,omitempty"`
	StartRow    *int32    `json:"StartRow,omitempty" xml:"StartRow,omitempty"`
	Texts       []*string `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	Width       *int32    `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeTableResponseBodyDataTablesTableRowsTableColumns) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTableResponseBodyDataTablesTableRowsTableColumns) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetEndColumn(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.EndColumn = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetEndRow(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.EndRow = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetHeight(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Height = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetStartColumn(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.StartColumn = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetStartRow(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.StartRow = &v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetTexts(v []*string) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Texts = v
	return s
}

func (s *RecognizeTableResponseBodyDataTablesTableRowsTableColumns) SetWidth(v int32) *RecognizeTableResponseBodyDataTablesTableRowsTableColumns {
	s.Width = &v
	return s
}

type RecognizeTableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTableResponse) SetStatusCode(v int32) *RecognizeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTableResponse) SetBody(v *RecognizeTableResponseBody) *RecognizeTableResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *RecognizeTakeoutOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTakeoutOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderResponseBody) SetData(v *RecognizeTakeoutOrderResponseBodyData) *RecognizeTakeoutOrderResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTakeoutOrderResponseBody) SetRequestId(v string) *RecognizeTakeoutOrderResponseBody {
	s.RequestId = &v
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
	Boxes []*int32 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RecognizeTakeoutOrderResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTakeoutOrderResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetBoxes(v []*int32) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Boxes = v
	return s
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetName(v string) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Name = &v
	return s
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetScore(v float32) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *RecognizeTakeoutOrderResponseBodyDataElements) SetValue(v string) *RecognizeTakeoutOrderResponseBodyDataElements {
	s.Value = &v
	return s
}

type RecognizeTakeoutOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTakeoutOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTakeoutOrderResponse) SetStatusCode(v int32) *RecognizeTakeoutOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTakeoutOrderResponse) SetBody(v *RecognizeTakeoutOrderResponseBody) *RecognizeTakeoutOrderResponse {
	s.Body = v
	return s
}

type RecognizeTaxiInvoiceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTaxiInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceRequest) SetImageURL(v string) *RecognizeTaxiInvoiceRequest {
	s.ImageURL = &v
	return s
}

type RecognizeTaxiInvoiceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeTaxiInvoiceResponseBody struct {
	Data      *RecognizeTaxiInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBody) SetData(v *RecognizeTaxiInvoiceResponseBodyData) *RecognizeTaxiInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBody) SetRequestId(v string) *RecognizeTaxiInvoiceResponseBody {
	s.RequestId = &v
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
	InvoiceRoi *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi `json:"InvoiceRoi,omitempty" xml:"InvoiceRoi,omitempty" type:"Struct"`
	Items      []*RecognizeTaxiInvoiceResponseBodyDataInvoicesItems    `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RotateType *int32                                                  `json:"RotateType,omitempty" xml:"RotateType,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoices) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoices) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetInvoiceRoi(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.InvoiceRoi = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetItems(v []*RecognizeTaxiInvoiceResponseBodyDataInvoicesItems) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.Items = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoices) SetRotateType(v int32) *RecognizeTaxiInvoiceResponseBodyDataInvoices {
	s.RotateType = &v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi struct {
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetH(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.H = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetW(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.W = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetX(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.X = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi) SetY(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesInvoiceRoi {
	s.Y = &v
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
	Angle  *float32                                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Center *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter `json:"Center,omitempty" xml:"Center,omitempty" type:"Struct"`
	Size   *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize   `json:"Size,omitempty" xml:"Size,omitempty" type:"Struct"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetAngle(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Angle = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetCenter(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Center = v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi) SetSize(v *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoi {
	s.Size = v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) SetX(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter {
	s.X = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter) SetY(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiCenter {
	s.Y = &v
	return s
}

type RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize struct {
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) GoString() string {
	return s.String()
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) SetH(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize {
	s.H = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize) SetW(v float32) *RecognizeTaxiInvoiceResponseBodyDataInvoicesItemsItemRoiSize {
	s.W = &v
	return s
}

type RecognizeTaxiInvoiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTaxiInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTaxiInvoiceResponse) SetStatusCode(v int32) *RecognizeTaxiInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTaxiInvoiceResponse) SetBody(v *RecognizeTaxiInvoiceResponseBody) *RecognizeTaxiInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeTicketInvoiceRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTicketInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceRequest) SetImageURL(v string) *RecognizeTicketInvoiceRequest {
	s.ImageURL = &v
	return s
}

type RecognizeTicketInvoiceAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTicketInvoiceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeTicketInvoiceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeTicketInvoiceResponseBody struct {
	Data      *RecognizeTicketInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBody) SetData(v *RecognizeTicketInvoiceResponseBodyData) *RecognizeTicketInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBody) SetRequestId(v string) *RecognizeTicketInvoiceResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTicketInvoiceResponseBodyData struct {
	Count     *int64                                           `json:"Count,omitempty" xml:"Count,omitempty"`
	Height    *int64                                           `json:"Height,omitempty" xml:"Height,omitempty"`
	OrgHeight *int64                                           `json:"OrgHeight,omitempty" xml:"OrgHeight,omitempty"`
	OrgWidth  *int64                                           `json:"OrgWidth,omitempty" xml:"OrgWidth,omitempty"`
	Results   []*RecognizeTicketInvoiceResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Width     *int64                                           `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetCount(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.Count = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetHeight(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetOrgHeight(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.OrgHeight = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetOrgWidth(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.OrgWidth = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetResults(v []*RecognizeTicketInvoiceResponseBodyDataResults) *RecognizeTicketInvoiceResponseBodyData {
	s.Results = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyData) SetWidth(v int64) *RecognizeTicketInvoiceResponseBodyData {
	s.Width = &v
	return s
}

type RecognizeTicketInvoiceResponseBodyDataResults struct {
	Content        *RecognizeTicketInvoiceResponseBodyDataResultsContent          `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Index          *int64                                                         `json:"Index,omitempty" xml:"Index,omitempty"`
	KeyValueInfos  []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos  `json:"KeyValueInfos,omitempty" xml:"KeyValueInfos,omitempty" type:"Repeated"`
	SliceRectangle []*RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle `json:"SliceRectangle,omitempty" xml:"SliceRectangle,omitempty" type:"Repeated"`
	Type           *string                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetContent(v *RecognizeTicketInvoiceResponseBodyDataResultsContent) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.Content = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetIndex(v int64) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.Index = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetKeyValueInfos(v []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.KeyValueInfos = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetSliceRectangle(v []*RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.SliceRectangle = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResults) SetType(v string) *RecognizeTicketInvoiceResponseBodyDataResults {
	s.Type = &v
	return s
}

type RecognizeTicketInvoiceResponseBodyDataResultsContent struct {
	AntiFakeCode    *string `json:"AntiFakeCode,omitempty" xml:"AntiFakeCode,omitempty"`
	InvoiceCode     *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceDate     *string `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	InvoiceNumber   *string `json:"InvoiceNumber,omitempty" xml:"InvoiceNumber,omitempty"`
	PayeeName       *string `json:"PayeeName,omitempty" xml:"PayeeName,omitempty"`
	PayeeRegisterNo *string `json:"PayeeRegisterNo,omitempty" xml:"PayeeRegisterNo,omitempty"`
	PayerName       *string `json:"PayerName,omitempty" xml:"PayerName,omitempty"`
	PayerRegisterNo *string `json:"PayerRegisterNo,omitempty" xml:"PayerRegisterNo,omitempty"`
	SumAmount       *string `json:"SumAmount,omitempty" xml:"SumAmount,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsContent) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsContent) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetAntiFakeCode(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.AntiFakeCode = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetInvoiceCode(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.InvoiceCode = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetInvoiceDate(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.InvoiceDate = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetInvoiceNumber(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.InvoiceNumber = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayeeName(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayeeName = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayeeRegisterNo(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayeeRegisterNo = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayerName(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayerName = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetPayerRegisterNo(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.PayerRegisterNo = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsContent) SetSumAmount(v string) *RecognizeTicketInvoiceResponseBodyDataResultsContent {
	s.SumAmount = &v
	return s
}

type RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos struct {
	Key            *string                                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value          *string                                                                     `json:"Value,omitempty" xml:"Value,omitempty"`
	ValuePositions []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions `json:"ValuePositions,omitempty" xml:"ValuePositions,omitempty" type:"Repeated"`
	ValueScore     *float32                                                                    `json:"ValueScore,omitempty" xml:"ValueScore,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetKey(v string) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.Key = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetValue(v string) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.Value = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetValuePositions(v []*RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.ValuePositions = v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos) SetValueScore(v float32) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfos {
	s.ValueScore = &v
	return s
}

type RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) SetX(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions {
	s.X = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions) SetY(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsKeyValueInfosValuePositions {
	s.Y = &v
	return s
}

type RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) SetX(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle {
	s.X = &v
	return s
}

func (s *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle) SetY(v int64) *RecognizeTicketInvoiceResponseBodyDataResultsSliceRectangle {
	s.Y = &v
	return s
}

type RecognizeTicketInvoiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTicketInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTicketInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTicketInvoiceResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTicketInvoiceResponse) SetHeaders(v map[string]*string) *RecognizeTicketInvoiceResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTicketInvoiceResponse) SetStatusCode(v int32) *RecognizeTicketInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTicketInvoiceResponse) SetBody(v *RecognizeTicketInvoiceResponseBody) *RecognizeTicketInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeTrainTicketRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeTrainTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketRequest) SetImageURL(v string) *RecognizeTrainTicketRequest {
	s.ImageURL = &v
	return s
}

type RecognizeTrainTicketAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeTrainTicketResponseBody struct {
	Data      *RecognizeTrainTicketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTrainTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponseBody) SetData(v *RecognizeTrainTicketResponseBodyData) *RecognizeTrainTicketResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTrainTicketResponseBody) SetRequestId(v string) *RecognizeTrainTicketResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTrainTicketResponseBodyData struct {
	Date             *string  `json:"Date,omitempty" xml:"Date,omitempty"`
	DepartureStation *string  `json:"DepartureStation,omitempty" xml:"DepartureStation,omitempty"`
	Destination      *string  `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Level            *string  `json:"Level,omitempty" xml:"Level,omitempty"`
	Name             *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Number           *string  `json:"Number,omitempty" xml:"Number,omitempty"`
	Price            *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	Seat             *string  `json:"Seat,omitempty" xml:"Seat,omitempty"`
}

func (s RecognizeTrainTicketResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTrainTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTrainTicketResponseBodyData) SetDate(v string) *RecognizeTrainTicketResponseBodyData {
	s.Date = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetDepartureStation(v string) *RecognizeTrainTicketResponseBodyData {
	s.DepartureStation = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetDestination(v string) *RecognizeTrainTicketResponseBodyData {
	s.Destination = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetLevel(v string) *RecognizeTrainTicketResponseBodyData {
	s.Level = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetName(v string) *RecognizeTrainTicketResponseBodyData {
	s.Name = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetNumber(v string) *RecognizeTrainTicketResponseBodyData {
	s.Number = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetPrice(v float32) *RecognizeTrainTicketResponseBodyData {
	s.Price = &v
	return s
}

func (s *RecognizeTrainTicketResponseBodyData) SetSeat(v string) *RecognizeTrainTicketResponseBodyData {
	s.Seat = &v
	return s
}

type RecognizeTrainTicketResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTrainTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeTrainTicketResponse) SetStatusCode(v int32) *RecognizeTrainTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTrainTicketResponse) SetBody(v *RecognizeTrainTicketResponseBody) *RecognizeTrainTicketResponse {
	s.Body = v
	return s
}

type RecognizeTurkeyIdentityCardRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeTurkeyIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardRequest) SetImageUrl(v string) *RecognizeTurkeyIdentityCardRequest {
	s.ImageUrl = &v
	return s
}

type RecognizeTurkeyIdentityCardAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeTurkeyIdentityCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardAdvanceRequest) SetImageUrlObject(v io.Reader) *RecognizeTurkeyIdentityCardAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type RecognizeTurkeyIdentityCardResponseBody struct {
	Data      *RecognizeTurkeyIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBody) SetData(v *RecognizeTurkeyIdentityCardResponseBodyData) *RecognizeTurkeyIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBody) SetRequestId(v string) *RecognizeTurkeyIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyData struct {
	AuxiliaryTools        *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools        `json:"AuxiliaryTools,omitempty" xml:"AuxiliaryTools,omitempty" type:"Struct"`
	BirthDate             *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate             `json:"BirthDate,omitempty" xml:"BirthDate,omitempty" type:"Struct"`
	BirthPlace            *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace            `json:"BirthPlace,omitempty" xml:"BirthPlace,omitempty" type:"Struct"`
	BloodType             *RecognizeTurkeyIdentityCardResponseBodyDataBloodType             `json:"BloodType,omitempty" xml:"BloodType,omitempty" type:"Struct"`
	CardBox               *RecognizeTurkeyIdentityCardResponseBodyDataCardBox               `json:"CardBox,omitempty" xml:"CardBox,omitempty" type:"Struct"`
	Cilt                  *RecognizeTurkeyIdentityCardResponseBodyDataCilt                  `json:"Cilt,omitempty" xml:"Cilt,omitempty" type:"Struct"`
	DocumentNumber        *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber        `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty" type:"Struct"`
	DriveClass            *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass            `json:"DriveClass,omitempty" xml:"DriveClass,omitempty" type:"Struct"`
	DueDate               *RecognizeTurkeyIdentityCardResponseBodyDataDueDate               `json:"DueDate,omitempty" xml:"DueDate,omitempty" type:"Struct"`
	Duzenleyen            *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen            `json:"Duzenleyen,omitempty" xml:"Duzenleyen,omitempty" type:"Struct"`
	EntryNumber           *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber           `json:"EntryNumber,omitempty" xml:"EntryNumber,omitempty" type:"Struct"`
	ExpiryDate            *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate            `json:"ExpiryDate,omitempty" xml:"ExpiryDate,omitempty" type:"Struct"`
	FatherName            *RecognizeTurkeyIdentityCardResponseBodyDataFatherName            `json:"FatherName,omitempty" xml:"FatherName,omitempty" type:"Struct"`
	ForeignersId          *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId          `json:"ForeignersId,omitempty" xml:"ForeignersId,omitempty" type:"Struct"`
	Gender                *RecognizeTurkeyIdentityCardResponseBodyDataGender                `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	GivenName             *RecognizeTurkeyIdentityCardResponseBodyDataGivenName             `json:"GivenName,omitempty" xml:"GivenName,omitempty" type:"Struct"`
	IdNumber              *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber              `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" type:"Struct"`
	IssueBy               *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy               `json:"IssueBy,omitempty" xml:"IssueBy,omitempty" type:"Struct"`
	IssueCounty           *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty           `json:"IssueCounty,omitempty" xml:"IssueCounty,omitempty" type:"Struct"`
	IssueDate             *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate             `json:"IssueDate,omitempty" xml:"IssueDate,omitempty" type:"Struct"`
	IssuePlace            *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace            `json:"IssuePlace,omitempty" xml:"IssuePlace,omitempty" type:"Struct"`
	Kutuk                 *RecognizeTurkeyIdentityCardResponseBodyDataKutuk                 `json:"Kutuk,omitempty" xml:"Kutuk,omitempty" type:"Struct"`
	LicenseNumber         *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber         `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty" type:"Struct"`
	MaritalStatus         *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus         `json:"MaritalStatus,omitempty" xml:"MaritalStatus,omitempty" type:"Struct"`
	MotherName            *RecognizeTurkeyIdentityCardResponseBodyDataMotherName            `json:"MotherName,omitempty" xml:"MotherName,omitempty" type:"Struct"`
	Name                  *RecognizeTurkeyIdentityCardResponseBodyDataName                  `json:"Name,omitempty" xml:"Name,omitempty" type:"Struct"`
	Nationality           *RecognizeTurkeyIdentityCardResponseBodyDataNationality           `json:"Nationality,omitempty" xml:"Nationality,omitempty" type:"Struct"`
	NeighborhoodVillage   *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage   `json:"NeighborhoodVillage,omitempty" xml:"NeighborhoodVillage,omitempty" type:"Struct"`
	PageNumber            *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" type:"Struct"`
	PassportNumber        *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber        `json:"PassportNumber,omitempty" xml:"PassportNumber,omitempty" type:"Struct"`
	PortraitBox           *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox           `json:"PortraitBox,omitempty" xml:"PortraitBox,omitempty" type:"Struct"`
	Province              *RecognizeTurkeyIdentityCardResponseBodyDataProvince              `json:"Province,omitempty" xml:"Province,omitempty" type:"Struct"`
	ProvinceOfResidence   *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence   `json:"ProvinceOfResidence,omitempty" xml:"ProvinceOfResidence,omitempty" type:"Struct"`
	ReasonOfIssue         *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue         `json:"ReasonOfIssue,omitempty" xml:"ReasonOfIssue,omitempty" type:"Struct"`
	RegisterNumber        *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber        `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty" type:"Struct"`
	Religion              *RecognizeTurkeyIdentityCardResponseBodyDataReligion              `json:"Religion,omitempty" xml:"Religion,omitempty" type:"Struct"`
	Sayfa                 *RecognizeTurkeyIdentityCardResponseBodyDataSayfa                 `json:"Sayfa,omitempty" xml:"Sayfa,omitempty" type:"Struct"`
	Seri                  *RecognizeTurkeyIdentityCardResponseBodyDataSeri                  `json:"Seri,omitempty" xml:"Seri,omitempty" type:"Struct"`
	Sex                   *RecognizeTurkeyIdentityCardResponseBodyDataSex                   `json:"Sex,omitempty" xml:"Sex,omitempty" type:"Struct"`
	Surname               *RecognizeTurkeyIdentityCardResponseBodyDataSurname               `json:"Surname,omitempty" xml:"Surname,omitempty" type:"Struct"`
	TypeOfResidencePermit *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit `json:"TypeOfResidencePermit,omitempty" xml:"TypeOfResidencePermit,omitempty" type:"Struct"`
	ValidUntil            *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil            `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty" type:"Struct"`
	Village               *RecognizeTurkeyIdentityCardResponseBodyDataVillage               `json:"Village,omitempty" xml:"Village,omitempty" type:"Struct"`
	VolumeNumber          *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber          `json:"VolumeNumber,omitempty" xml:"VolumeNumber,omitempty" type:"Struct"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetAuxiliaryTools(v *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.AuxiliaryTools = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetBirthDate(v *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.BirthDate = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetBirthPlace(v *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.BirthPlace = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetBloodType(v *RecognizeTurkeyIdentityCardResponseBodyDataBloodType) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.BloodType = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetCardBox(v *RecognizeTurkeyIdentityCardResponseBodyDataCardBox) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.CardBox = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetCilt(v *RecognizeTurkeyIdentityCardResponseBodyDataCilt) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Cilt = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetDocumentNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.DocumentNumber = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetDriveClass(v *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.DriveClass = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetDueDate(v *RecognizeTurkeyIdentityCardResponseBodyDataDueDate) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.DueDate = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetDuzenleyen(v *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Duzenleyen = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetEntryNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.EntryNumber = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetExpiryDate(v *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.ExpiryDate = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetFatherName(v *RecognizeTurkeyIdentityCardResponseBodyDataFatherName) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.FatherName = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetForeignersId(v *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.ForeignersId = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetGender(v *RecognizeTurkeyIdentityCardResponseBodyDataGender) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Gender = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetGivenName(v *RecognizeTurkeyIdentityCardResponseBodyDataGivenName) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.GivenName = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetIdNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.IdNumber = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetIssueBy(v *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.IssueBy = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetIssueCounty(v *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.IssueCounty = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetIssueDate(v *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.IssueDate = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetIssuePlace(v *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.IssuePlace = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetKutuk(v *RecognizeTurkeyIdentityCardResponseBodyDataKutuk) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Kutuk = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetLicenseNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.LicenseNumber = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetMaritalStatus(v *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.MaritalStatus = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetMotherName(v *RecognizeTurkeyIdentityCardResponseBodyDataMotherName) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.MotherName = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetName(v *RecognizeTurkeyIdentityCardResponseBodyDataName) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Name = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetNationality(v *RecognizeTurkeyIdentityCardResponseBodyDataNationality) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Nationality = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetNeighborhoodVillage(v *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.NeighborhoodVillage = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetPageNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.PageNumber = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetPassportNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.PassportNumber = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetPortraitBox(v *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.PortraitBox = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetProvince(v *RecognizeTurkeyIdentityCardResponseBodyDataProvince) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Province = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetProvinceOfResidence(v *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.ProvinceOfResidence = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetReasonOfIssue(v *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.ReasonOfIssue = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetRegisterNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.RegisterNumber = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetReligion(v *RecognizeTurkeyIdentityCardResponseBodyDataReligion) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Religion = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetSayfa(v *RecognizeTurkeyIdentityCardResponseBodyDataSayfa) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Sayfa = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetSeri(v *RecognizeTurkeyIdentityCardResponseBodyDataSeri) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Seri = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetSex(v *RecognizeTurkeyIdentityCardResponseBodyDataSex) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Sex = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetSurname(v *RecognizeTurkeyIdentityCardResponseBodyDataSurname) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Surname = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetTypeOfResidencePermit(v *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.TypeOfResidencePermit = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetValidUntil(v *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.ValidUntil = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetVillage(v *RecognizeTurkeyIdentityCardResponseBodyDataVillage) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.Village = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyData) SetVolumeNumber(v *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber) *RecognizeTurkeyIdentityCardResponseBodyData {
	s.VolumeNumber = v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                               `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryTools {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataAuxiliaryToolsKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataBirthDate struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthDate) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataBirthDate {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlace {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBirthPlaceKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataBloodType struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBloodType) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBloodType) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBloodType) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataBloodType {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBloodType) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBloodType {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBloodType) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataBloodType {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataBloodTypeKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataCardBox struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCardBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCardBox) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCardBox) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataCardBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCardBox) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataCardBox {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCardBox) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataCardBox {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataCilt struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCilt) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCilt) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCilt) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataCilt {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCilt) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataCilt {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCilt) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataCilt {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataCiltKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                               `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDocumentNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDriveClass struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDriveClass) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDriveClass) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataDriveClass {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDriveClassKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDueDate struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDueDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDueDate) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDueDate) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataDueDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDueDate) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDueDate {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDueDate) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataDueDate {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDueDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyen {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataDuzenleyenKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataEntryNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDate {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataFatherName struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataFatherName) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataFatherName) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataFatherName) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataFatherName {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataFatherName) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataFatherName {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataFatherName) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataFatherName {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataFatherNameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataForeignersId struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataForeignersId) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataForeignersId) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataForeignersId {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataForeignersIdKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataGender struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                      `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                       `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGender) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGender) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGender) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataGender {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGender) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataGender {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGender) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataGender {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataGenderKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataGivenName struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGivenName) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGivenName) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGivenName) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataGivenName {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGivenName) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataGivenName {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGivenName) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataGivenName {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataGivenNameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIdNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                         `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIdNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIdNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataIdNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssueBy struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueBy) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueBy) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataIssueBy {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueByKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataIssueCounty {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueCountyKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssueDate struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueDate) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataIssueDate {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssueDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlace {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataIssuePlaceKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataKutuk struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                     `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                      `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataKutuk) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataKutuk) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataKutuk) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataKutuk {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataKutuk) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataKutuk {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataKutuk) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataKutuk {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataKutukKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataLicenseNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatus {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataMaritalStatusKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataMotherName struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMotherName) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMotherName) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMotherName) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataMotherName {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMotherName) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataMotherName {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMotherName) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataMotherName {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataMotherNameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataName struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataName) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataName) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataName) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataName {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataName) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataName {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataName) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataName {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataNationality struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNationality) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNationality) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNationality) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataNationality {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNationality) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNationality {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNationality) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataNationality {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNationalityKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillage {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataNeighborhoodVillageKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataPageNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPageNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPageNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataPageNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPageNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                               `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPassportNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBox {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataProvince struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                         `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvince) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvince) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvince) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataProvince {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvince) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataProvince {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvince) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataProvince {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataProvinceKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidence {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataProvinceOfResidenceKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssue {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataReasonOfIssueKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                              `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                               `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataRegisterNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataReligion struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                         `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReligion) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReligion) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReligion) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataReligion {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReligion) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataReligion {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReligion) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataReligion {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataReligionKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSayfa struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                     `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                      `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSayfa) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSayfa) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSayfa) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataSayfa {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSayfa) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSayfa {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSayfa) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataSayfa {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSayfaKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSeri struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSeri) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSeri) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSeri) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataSeri {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSeri) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSeri {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSeri) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataSeri {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSeriKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSex struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                   `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                    `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSex) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSex) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSex) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataSex {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSex) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSex {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSex) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataSex {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSexKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSurname struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSurname) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSurname) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSurname) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataSurname {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSurname) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSurname {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSurname) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataSurname {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataSurnameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                     `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                      `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermit {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataTypeOfResidencePermitKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataValidUntil struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataValidUntil) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataValidUntil) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataValidUntil {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataValidUntilKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataVillage struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVillage) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVillage) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVillage) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataVillage {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVillage) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataVillage {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVillage) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataVillage {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataVillageKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber struct {
	KeyPoints []*RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber) SetKeyPoints(v []*RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints) *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber) SetScore(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber {
	s.Score = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber) SetText(v string) *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumber {
	s.Text = &v
	return s
}

type RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints) SetX(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints) SetY(v float32) *RecognizeTurkeyIdentityCardResponseBodyDataVolumeNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeTurkeyIdentityCardResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeTurkeyIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeTurkeyIdentityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeTurkeyIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTurkeyIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeTurkeyIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponse) SetStatusCode(v int32) *RecognizeTurkeyIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTurkeyIdentityCardResponse) SetBody(v *RecognizeTurkeyIdentityCardResponseBody) *RecognizeTurkeyIdentityCardResponse {
	s.Body = v
	return s
}

type RecognizeUkraineIdentityCardRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeUkraineIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardRequest) SetImageUrl(v string) *RecognizeUkraineIdentityCardRequest {
	s.ImageUrl = &v
	return s
}

type RecognizeUkraineIdentityCardAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeUkraineIdentityCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardAdvanceRequest) SetImageUrlObject(v io.Reader) *RecognizeUkraineIdentityCardAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type RecognizeUkraineIdentityCardResponseBody struct {
	Data      *RecognizeUkraineIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBody) SetData(v *RecognizeUkraineIdentityCardResponseBodyData) *RecognizeUkraineIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBody) SetRequestId(v string) *RecognizeUkraineIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyData struct {
	BirthDate      *RecognizeUkraineIdentityCardResponseBodyDataBirthDate      `json:"BirthDate,omitempty" xml:"BirthDate,omitempty" type:"Struct"`
	CardBox        *RecognizeUkraineIdentityCardResponseBodyDataCardBox        `json:"CardBox,omitempty" xml:"CardBox,omitempty" type:"Struct"`
	DocumentNumber *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber `json:"DocumentNumber,omitempty" xml:"DocumentNumber,omitempty" type:"Struct"`
	ExpiryDate     *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate     `json:"ExpiryDate,omitempty" xml:"ExpiryDate,omitempty" type:"Struct"`
	NameEnglish    *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish    `json:"NameEnglish,omitempty" xml:"NameEnglish,omitempty" type:"Struct"`
	NameUkraine    *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine    `json:"NameUkraine,omitempty" xml:"NameUkraine,omitempty" type:"Struct"`
	Nationality    *RecognizeUkraineIdentityCardResponseBodyDataNationality    `json:"Nationality,omitempty" xml:"Nationality,omitempty" type:"Struct"`
	Patronymic     *RecognizeUkraineIdentityCardResponseBodyDataPatronymic     `json:"Patronymic,omitempty" xml:"Patronymic,omitempty" type:"Struct"`
	PortraitBox    *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox    `json:"PortraitBox,omitempty" xml:"PortraitBox,omitempty" type:"Struct"`
	RecordNumber   *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber   `json:"RecordNumber,omitempty" xml:"RecordNumber,omitempty" type:"Struct"`
	Sex            *RecognizeUkraineIdentityCardResponseBodyDataSex            `json:"Sex,omitempty" xml:"Sex,omitempty" type:"Struct"`
	SurnameEnglish *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish `json:"SurnameEnglish,omitempty" xml:"SurnameEnglish,omitempty" type:"Struct"`
	SurnameUkraine *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine `json:"SurnameUkraine,omitempty" xml:"SurnameUkraine,omitempty" type:"Struct"`
}

func (s RecognizeUkraineIdentityCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetBirthDate(v *RecognizeUkraineIdentityCardResponseBodyDataBirthDate) *RecognizeUkraineIdentityCardResponseBodyData {
	s.BirthDate = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetCardBox(v *RecognizeUkraineIdentityCardResponseBodyDataCardBox) *RecognizeUkraineIdentityCardResponseBodyData {
	s.CardBox = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetDocumentNumber(v *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber) *RecognizeUkraineIdentityCardResponseBodyData {
	s.DocumentNumber = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetExpiryDate(v *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate) *RecognizeUkraineIdentityCardResponseBodyData {
	s.ExpiryDate = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetNameEnglish(v *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish) *RecognizeUkraineIdentityCardResponseBodyData {
	s.NameEnglish = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetNameUkraine(v *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine) *RecognizeUkraineIdentityCardResponseBodyData {
	s.NameUkraine = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetNationality(v *RecognizeUkraineIdentityCardResponseBodyDataNationality) *RecognizeUkraineIdentityCardResponseBodyData {
	s.Nationality = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetPatronymic(v *RecognizeUkraineIdentityCardResponseBodyDataPatronymic) *RecognizeUkraineIdentityCardResponseBodyData {
	s.Patronymic = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetPortraitBox(v *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox) *RecognizeUkraineIdentityCardResponseBodyData {
	s.PortraitBox = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetRecordNumber(v *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber) *RecognizeUkraineIdentityCardResponseBodyData {
	s.RecordNumber = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetSex(v *RecognizeUkraineIdentityCardResponseBodyDataSex) *RecognizeUkraineIdentityCardResponseBodyData {
	s.Sex = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetSurnameEnglish(v *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish) *RecognizeUkraineIdentityCardResponseBodyData {
	s.SurnameEnglish = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyData) SetSurnameUkraine(v *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine) *RecognizeUkraineIdentityCardResponseBodyData {
	s.SurnameUkraine = v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataBirthDate struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataBirthDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataBirthDate) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataBirthDate) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataBirthDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataBirthDate) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataBirthDate {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataBirthDate) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataBirthDate {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataCardBox struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                         `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataCardBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataCardBox) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataCardBox) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataCardBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataCardBox) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataCardBox {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataCardBox) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataCardBox {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumber {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataDocumentNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataExpiryDate struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataExpiryDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataExpiryDate) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataExpiryDate {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataNameEnglish struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameEnglish) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameEnglish) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataNameEnglish {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNameEnglishKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataNameUkraine struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameUkraine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameUkraine) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataNameUkraine {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNameUkraineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataNationality struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNationality) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNationality) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNationality) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataNationality {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNationality) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNationality {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNationality) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataNationality {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataNationalityKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataPatronymic struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                           `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPatronymic) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPatronymic) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPatronymic) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataPatronymic {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPatronymic) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataPatronymic {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPatronymic) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataPatronymic {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataPatronymicKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataPortraitBox struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPortraitBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPortraitBox) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataPortraitBox {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataRecordNumber struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                              `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataRecordNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataRecordNumber) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataRecordNumber {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataRecordNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataSex struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                    `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSex) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSex) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSex) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataSex {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSex) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSex {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSex) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataSex {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSexKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglish {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSurnameEnglishKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine struct {
	KeyPoints []*RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine) SetKeyPoints(v []*RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints) *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine) SetScore(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine {
	s.Score = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine) SetText(v string) *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraine {
	s.Text = &v
	return s
}

type RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints) SetX(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints) SetY(v float32) *RecognizeUkraineIdentityCardResponseBodyDataSurnameUkraineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeUkraineIdentityCardResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeUkraineIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeUkraineIdentityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeUkraineIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeUkraineIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeUkraineIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeUkraineIdentityCardResponse) SetStatusCode(v int32) *RecognizeUkraineIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeUkraineIdentityCardResponse) SetBody(v *RecognizeUkraineIdentityCardResponseBody) *RecognizeUkraineIdentityCardResponse {
	s.Body = v
	return s
}

type RecognizeVATInvoiceRequest struct {
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileURL  *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizeVATInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceRequest) SetFileType(v string) *RecognizeVATInvoiceRequest {
	s.FileType = &v
	return s
}

func (s *RecognizeVATInvoiceRequest) SetFileURL(v string) *RecognizeVATInvoiceRequest {
	s.FileURL = &v
	return s
}

type RecognizeVATInvoiceAdvanceRequest struct {
	FileType      *string   `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileURLObject io.Reader `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
}

func (s RecognizeVATInvoiceAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceAdvanceRequest) SetFileType(v string) *RecognizeVATInvoiceAdvanceRequest {
	s.FileType = &v
	return s
}

func (s *RecognizeVATInvoiceAdvanceRequest) SetFileURLObject(v io.Reader) *RecognizeVATInvoiceAdvanceRequest {
	s.FileURLObject = v
	return s
}

type RecognizeVATInvoiceResponseBody struct {
	Data      *RecognizeVATInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVATInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBody) SetData(v *RecognizeVATInvoiceResponseBodyData) *RecognizeVATInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVATInvoiceResponseBody) SetRequestId(v string) *RecognizeVATInvoiceResponseBody {
	s.RequestId = &v
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
	Checkers          []*float32 `json:"Checkers,omitempty" xml:"Checkers,omitempty" type:"Repeated"`
	Clerks            []*float32 `json:"Clerks,omitempty" xml:"Clerks,omitempty" type:"Repeated"`
	InvoiceAmounts    []*float32 `json:"InvoiceAmounts,omitempty" xml:"InvoiceAmounts,omitempty" type:"Repeated"`
	InvoiceCodes      []*float32 `json:"InvoiceCodes,omitempty" xml:"InvoiceCodes,omitempty" type:"Repeated"`
	InvoiceDates      []*float32 `json:"InvoiceDates,omitempty" xml:"InvoiceDates,omitempty" type:"Repeated"`
	InvoiceFakeCodes  []*float32 `json:"InvoiceFakeCodes,omitempty" xml:"InvoiceFakeCodes,omitempty" type:"Repeated"`
	InvoiceNoes       []*float32 `json:"InvoiceNoes,omitempty" xml:"InvoiceNoes,omitempty" type:"Repeated"`
	PayeeAddresses    []*float32 `json:"PayeeAddresses,omitempty" xml:"PayeeAddresses,omitempty" type:"Repeated"`
	PayeeBankNames    []*float32 `json:"PayeeBankNames,omitempty" xml:"PayeeBankNames,omitempty" type:"Repeated"`
	PayeeNames        []*float32 `json:"PayeeNames,omitempty" xml:"PayeeNames,omitempty" type:"Repeated"`
	PayeeRegisterNoes []*float32 `json:"PayeeRegisterNoes,omitempty" xml:"PayeeRegisterNoes,omitempty" type:"Repeated"`
	Payees            []*float32 `json:"Payees,omitempty" xml:"Payees,omitempty" type:"Repeated"`
	PayerAddresses    []*float32 `json:"PayerAddresses,omitempty" xml:"PayerAddresses,omitempty" type:"Repeated"`
	PayerBankNames    []*float32 `json:"PayerBankNames,omitempty" xml:"PayerBankNames,omitempty" type:"Repeated"`
	PayerNames        []*float32 `json:"PayerNames,omitempty" xml:"PayerNames,omitempty" type:"Repeated"`
	PayerRegisterNoes []*float32 `json:"PayerRegisterNoes,omitempty" xml:"PayerRegisterNoes,omitempty" type:"Repeated"`
	SumAmounts        []*float32 `json:"SumAmounts,omitempty" xml:"SumAmounts,omitempty" type:"Repeated"`
	TaxAmounts        []*float32 `json:"TaxAmounts,omitempty" xml:"TaxAmounts,omitempty" type:"Repeated"`
	WithoutTaxAmounts []*float32 `json:"WithoutTaxAmounts,omitempty" xml:"WithoutTaxAmounts,omitempty" type:"Repeated"`
}

func (s RecognizeVATInvoiceResponseBodyDataBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyDataBox) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetCheckers(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Checkers = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetClerks(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Clerks = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceCodes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceCodes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceDates(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceDates = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceFakeCodes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceFakeCodes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceNoes = v
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

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeRegisterNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeRegisterNoes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayees(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Payees = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerAddresses(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerAddresses = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerBankNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerBankNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerRegisterNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerRegisterNoes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetSumAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.SumAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetTaxAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.TaxAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetWithoutTaxAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.WithoutTaxAmounts = v
	return s
}

type RecognizeVATInvoiceResponseBodyDataContent struct {
	AntiFakeCode     *string `json:"AntiFakeCode,omitempty" xml:"AntiFakeCode,omitempty"`
	Checker          *string `json:"Checker,omitempty" xml:"Checker,omitempty"`
	Clerk            *string `json:"Clerk,omitempty" xml:"Clerk,omitempty"`
	InvoiceAmount    *string `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	InvoiceCode      *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	InvoiceDate      *string `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	InvoiceNo        *string `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	Payee            *string `json:"Payee,omitempty" xml:"Payee,omitempty"`
	PayeeAddress     *string `json:"PayeeAddress,omitempty" xml:"PayeeAddress,omitempty"`
	PayeeBankName    *string `json:"PayeeBankName,omitempty" xml:"PayeeBankName,omitempty"`
	PayeeName        *string `json:"PayeeName,omitempty" xml:"PayeeName,omitempty"`
	PayeeRegisterNo  *string `json:"PayeeRegisterNo,omitempty" xml:"PayeeRegisterNo,omitempty"`
	PayerAddress     *string `json:"PayerAddress,omitempty" xml:"PayerAddress,omitempty"`
	PayerBankName    *string `json:"PayerBankName,omitempty" xml:"PayerBankName,omitempty"`
	PayerName        *string `json:"PayerName,omitempty" xml:"PayerName,omitempty"`
	PayerRegisterNo  *string `json:"PayerRegisterNo,omitempty" xml:"PayerRegisterNo,omitempty"`
	SumAmount        *string `json:"SumAmount,omitempty" xml:"SumAmount,omitempty"`
	TaxAmount        *string `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	WithoutTaxAmount *string `json:"WithoutTaxAmount,omitempty" xml:"WithoutTaxAmount,omitempty"`
}

func (s RecognizeVATInvoiceResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetAntiFakeCode(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.AntiFakeCode = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetChecker(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Checker = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetClerk(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Clerk = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceCode(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceCode = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceDate(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceDate = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayee(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Payee = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeAddress(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeAddress = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeBankName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeBankName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeRegisterNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeRegisterNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerAddress(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerAddress = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerBankName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerBankName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerRegisterNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerRegisterNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetSumAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.SumAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetTaxAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.TaxAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetWithoutTaxAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.WithoutTaxAmount = &v
	return s
}

type RecognizeVATInvoiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVATInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeVATInvoiceResponse) SetStatusCode(v int32) *RecognizeVATInvoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVATInvoiceResponse) SetBody(v *RecognizeVATInvoiceResponseBody) *RecognizeVATInvoiceResponse {
	s.Body = v
	return s
}

type RecognizeVINCodeRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVINCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVINCodeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeRequest) SetImageURL(v string) *RecognizeVINCodeRequest {
	s.ImageURL = &v
	return s
}

type RecognizeVINCodeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

type RecognizeVINCodeResponseBody struct {
	Data      *RecognizeVINCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVINCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVINCodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVINCodeResponseBody) SetData(v *RecognizeVINCodeResponseBodyData) *RecognizeVINCodeResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVINCodeResponseBody) SetRequestId(v string) *RecognizeVINCodeResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVINCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeVINCodeResponse) SetStatusCode(v int32) *RecognizeVINCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVINCodeResponse) SetBody(v *RecognizeVINCodeResponseBody) *RecognizeVINCodeResponse {
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
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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
	Data      *RecognizeVerificationcodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVerificationcodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVerificationcodeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVerificationcodeResponseBody) SetData(v *RecognizeVerificationcodeResponseBodyData) *RecognizeVerificationcodeResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVerificationcodeResponseBody) SetRequestId(v string) *RecognizeVerificationcodeResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVerificationcodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RecognizeVerificationcodeResponse) SetStatusCode(v int32) *RecognizeVerificationcodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVerificationcodeResponse) SetBody(v *RecognizeVerificationcodeResponseBody) *RecognizeVerificationcodeResponse {
	s.Body = v
	return s
}

type RecognizeVideoCastCrewListRequest struct {
	Params      []*RecognizeVideoCastCrewListRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RegisterUrl *string                                    `json:"RegisterUrl,omitempty" xml:"RegisterUrl,omitempty"`
	VideoUrl    *string                                    `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListRequest) SetParams(v []*RecognizeVideoCastCrewListRequestParams) *RecognizeVideoCastCrewListRequest {
	s.Params = v
	return s
}

func (s *RecognizeVideoCastCrewListRequest) SetRegisterUrl(v string) *RecognizeVideoCastCrewListRequest {
	s.RegisterUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListRequest {
	s.VideoUrl = &v
	return s
}

type RecognizeVideoCastCrewListRequestParams struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVideoCastCrewListRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListRequestParams) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListRequestParams) SetType(v string) *RecognizeVideoCastCrewListRequestParams {
	s.Type = &v
	return s
}

type RecognizeVideoCastCrewListAdvanceRequest struct {
	Params         []*RecognizeVideoCastCrewListAdvanceRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RegisterUrl    *string                                           `json:"RegisterUrl,omitempty" xml:"RegisterUrl,omitempty"`
	VideoUrlObject io.Reader                                         `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetParams(v []*RecognizeVideoCastCrewListAdvanceRequestParams) *RecognizeVideoCastCrewListAdvanceRequest {
	s.Params = v
	return s
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetRegisterUrl(v string) *RecognizeVideoCastCrewListAdvanceRequest {
	s.RegisterUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetVideoUrlObject(v io.Reader) *RecognizeVideoCastCrewListAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

type RecognizeVideoCastCrewListAdvanceRequestParams struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVideoCastCrewListAdvanceRequestParams) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListAdvanceRequestParams) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListAdvanceRequestParams) SetType(v string) *RecognizeVideoCastCrewListAdvanceRequestParams {
	s.Type = &v
	return s
}

type RecognizeVideoCastCrewListShrinkRequest struct {
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	RegisterUrl  *string `json:"RegisterUrl,omitempty" xml:"RegisterUrl,omitempty"`
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListShrinkRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetParamsShrink(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetRegisterUrl(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.RegisterUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListShrinkRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListShrinkRequest {
	s.VideoUrl = &v
	return s
}

type RecognizeVideoCastCrewListResponseBody struct {
	Data      *RecognizeVideoCastCrewListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBody) SetData(v *RecognizeVideoCastCrewListResponseBodyData) *RecognizeVideoCastCrewListResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBody) SetRequestId(v string) *RecognizeVideoCastCrewListResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyData struct {
	CastResults      []*RecognizeVideoCastCrewListResponseBodyDataCastResults      `json:"CastResults,omitempty" xml:"CastResults,omitempty" type:"Repeated"`
	OcrResults       []*RecognizeVideoCastCrewListResponseBodyDataOcrResults       `json:"OcrResults,omitempty" xml:"OcrResults,omitempty" type:"Repeated"`
	SubtitlesResults []*RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults `json:"SubtitlesResults,omitempty" xml:"SubtitlesResults,omitempty" type:"Repeated"`
	VideoOcrResults  []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults  `json:"VideoOcrResults,omitempty" xml:"VideoOcrResults,omitempty" type:"Repeated"`
}

func (s RecognizeVideoCastCrewListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetCastResults(v []*RecognizeVideoCastCrewListResponseBodyDataCastResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.CastResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetOcrResults(v []*RecognizeVideoCastCrewListResponseBodyDataOcrResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.OcrResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetSubtitlesResults(v []*RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.SubtitlesResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyData) SetVideoOcrResults(v []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) *RecognizeVideoCastCrewListResponseBodyData {
	s.VideoOcrResults = v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataCastResults struct {
	DetailInfo map[string]*string `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	EndTime    *float32           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *float32           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataCastResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataCastResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) SetDetailInfo(v map[string]*string) *RecognizeVideoCastCrewListResponseBodyDataCastResults {
	s.DetailInfo = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) SetEndTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataCastResults {
	s.EndTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataCastResults) SetStartTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataCastResults {
	s.StartTime = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataOcrResults struct {
	DetailInfo []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
	EndTime    *float32                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *float32                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) SetDetailInfo(v []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) *RecognizeVideoCastCrewListResponseBodyDataOcrResults {
	s.DetailInfo = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) SetEndTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResults {
	s.EndTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResults) SetStartTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResults {
	s.StartTime = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo struct {
	Boxes      []*int32                                                                  `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	CharProbs  [][]*float32                                                              `json:"CharProbs,omitempty" xml:"CharProbs,omitempty" type:"Repeated"`
	FrameIndex *int64                                                                    `json:"FrameIndex,omitempty" xml:"FrameIndex,omitempty"`
	Position   []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	Score      *float32                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text       *string                                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
	TextProb   *float32                                                                  `json:"TextProb,omitempty" xml:"TextProb,omitempty"`
	TimeStamp  *float32                                                                  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrackId    *int64                                                                    `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetBoxes(v []*int32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Boxes = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetCharProbs(v [][]*float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.CharProbs = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetFrameIndex(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.FrameIndex = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetPosition(v []*RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Position = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetScore(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Score = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetText(v string) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.Text = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetTextProb(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.TextProb = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetTimeStamp(v float32) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.TimeStamp = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo) SetTrackId(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfo {
	s.TrackId = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) SetX(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition {
	s.X = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition) SetY(v int64) *RecognizeVideoCastCrewListResponseBodyDataOcrResultsDetailInfoPosition {
	s.Y = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults struct {
	SubtitlesAllResults        map[string]*string     `json:"SubtitlesAllResults,omitempty" xml:"SubtitlesAllResults,omitempty"`
	SubtitlesAllResultsUrl     *string                `json:"SubtitlesAllResultsUrl,omitempty" xml:"SubtitlesAllResultsUrl,omitempty"`
	SubtitlesChineseResults    map[string]*string     `json:"SubtitlesChineseResults,omitempty" xml:"SubtitlesChineseResults,omitempty"`
	SubtitlesChineseResultsUrl *string                `json:"SubtitlesChineseResultsUrl,omitempty" xml:"SubtitlesChineseResultsUrl,omitempty"`
	SubtitlesEnglishResults    map[string]interface{} `json:"SubtitlesEnglishResults,omitempty" xml:"SubtitlesEnglishResults,omitempty"`
	SubtitlesEnglishResultsUrl *string                `json:"SubtitlesEnglishResultsUrl,omitempty" xml:"SubtitlesEnglishResultsUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesAllResults(v map[string]*string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesAllResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesAllResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesAllResultsUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesChineseResults(v map[string]*string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesChineseResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesChineseResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesChineseResultsUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesEnglishResults(v map[string]interface{}) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesEnglishResults = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults) SetSubtitlesEnglishResultsUrl(v string) *RecognizeVideoCastCrewListResponseBodyDataSubtitlesResults {
	s.SubtitlesEnglishResultsUrl = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults struct {
	DetailInfo []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
	EndTime    *float32                                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *float32                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) SetDetailInfo(v []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults {
	s.DetailInfo = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) SetEndTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults {
	s.EndTime = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults) SetStartTime(v float32) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResults {
	s.StartTime = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo struct {
	Boxes    []*int64                                                                       `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	Position []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
	Score    *float32                                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text     *string                                                                        `json:"Text,omitempty" xml:"Text,omitempty"`
	TextType *int64                                                                         `json:"TextType,omitempty" xml:"TextType,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetBoxes(v []*int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Boxes = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetPosition(v []*RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Position = v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetScore(v float32) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Score = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetText(v string) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.Text = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo) SetTextType(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfo {
	s.TextType = &v
	return s
}

type RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition struct {
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) SetX(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition {
	s.X = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition) SetY(v int64) *RecognizeVideoCastCrewListResponseBodyDataVideoOcrResultsDetailInfoPosition {
	s.Y = &v
	return s
}

type RecognizeVideoCastCrewListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVideoCastCrewListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVideoCastCrewListResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCastCrewListResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListResponse) SetHeaders(v map[string]*string) *RecognizeVideoCastCrewListResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVideoCastCrewListResponse) SetStatusCode(v int32) *RecognizeVideoCastCrewListResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVideoCastCrewListResponse) SetBody(v *RecognizeVideoCastCrewListResponseBody) *RecognizeVideoCastCrewListResponse {
	s.Body = v
	return s
}

type RecognizeVideoCharacterRequest struct {
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s RecognizeVideoCharacterRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCharacterRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterRequest) SetVideoURL(v string) *RecognizeVideoCharacterRequest {
	s.VideoURL = &v
	return s
}

type RecognizeVideoCharacterResponseBody struct {
	Data      *RecognizeVideoCharacterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVideoCharacterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBody) SetData(v *RecognizeVideoCharacterResponseBodyData) *RecognizeVideoCharacterResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVideoCharacterResponseBody) SetRequestId(v string) *RecognizeVideoCharacterResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVideoCharacterResponseBodyData struct {
	Frames    []*RecognizeVideoCharacterResponseBodyDataFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	Height    *int64                                           `json:"Height,omitempty" xml:"Height,omitempty"`
	InputFile *string                                          `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	Width     *int64                                           `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeVideoCharacterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyData) SetFrames(v []*RecognizeVideoCharacterResponseBodyDataFrames) *RecognizeVideoCharacterResponseBodyData {
	s.Frames = v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyData) SetHeight(v int64) *RecognizeVideoCharacterResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyData) SetInputFile(v string) *RecognizeVideoCharacterResponseBodyData {
	s.InputFile = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyData) SetWidth(v int64) *RecognizeVideoCharacterResponseBodyData {
	s.Width = &v
	return s
}

type RecognizeVideoCharacterResponseBodyDataFrames struct {
	Elements  []*RecognizeVideoCharacterResponseBodyDataFramesElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Timestamp *int64                                                   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s RecognizeVideoCharacterResponseBodyDataFrames) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyDataFrames) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyDataFrames) SetElements(v []*RecognizeVideoCharacterResponseBodyDataFramesElements) *RecognizeVideoCharacterResponseBodyDataFrames {
	s.Elements = v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFrames) SetTimestamp(v int64) *RecognizeVideoCharacterResponseBodyDataFrames {
	s.Timestamp = &v
	return s
}

type RecognizeVideoCharacterResponseBodyDataFramesElements struct {
	Score          *float32                                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	Text           *string                                                                `json:"Text,omitempty" xml:"Text,omitempty"`
	TextRectangles []*RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles `json:"TextRectangles,omitempty" xml:"TextRectangles,omitempty" type:"Repeated"`
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElements) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) SetScore(v float32) *RecognizeVideoCharacterResponseBodyDataFramesElements {
	s.Score = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) SetText(v string) *RecognizeVideoCharacterResponseBodyDataFramesElements {
	s.Text = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) SetTextRectangles(v []*RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) *RecognizeVideoCharacterResponseBodyDataFramesElements {
	s.TextRectangles = v
	return s
}

type RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles struct {
	Angle  *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	Left   *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	Top    *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	Width  *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetAngle(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Angle = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetHeight(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Height = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetLeft(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Left = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetTop(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Top = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetWidth(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Width = &v
	return s
}

type RecognizeVideoCharacterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVideoCharacterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVideoCharacterResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVideoCharacterResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponse) SetHeaders(v map[string]*string) *RecognizeVideoCharacterResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVideoCharacterResponse) SetStatusCode(v int32) *RecognizeVideoCharacterResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVideoCharacterResponse) SetBody(v *RecognizeVideoCharacterResponseBody) *RecognizeVideoCharacterResponse {
	s.Body = v
	return s
}

type RecognizeVietnamIdentityCardRequest struct {
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeVietnamIdentityCardRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardRequest) SetImageUrl(v string) *RecognizeVietnamIdentityCardRequest {
	s.ImageUrl = &v
	return s
}

type RecognizeVietnamIdentityCardAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s RecognizeVietnamIdentityCardAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardAdvanceRequest) SetImageUrlObject(v io.Reader) *RecognizeVietnamIdentityCardAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

type RecognizeVietnamIdentityCardResponseBody struct {
	Data      *RecognizeVietnamIdentityCardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBody) SetData(v *RecognizeVietnamIdentityCardResponseBodyData) *RecognizeVietnamIdentityCardResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBody) SetRequestId(v string) *RecognizeVietnamIdentityCardResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyData struct {
	AddressFirstLine         *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine         `json:"AddressFirstLine,omitempty" xml:"AddressFirstLine,omitempty" type:"Struct"`
	AddressSecondLine        *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine        `json:"AddressSecondLine,omitempty" xml:"AddressSecondLine,omitempty" type:"Struct"`
	BirthDate                *RecognizeVietnamIdentityCardResponseBodyDataBirthDate                `json:"BirthDate,omitempty" xml:"BirthDate,omitempty" type:"Struct"`
	CardBox                  *RecognizeVietnamIdentityCardResponseBodyDataCardBox                  `json:"CardBox,omitempty" xml:"CardBox,omitempty" type:"Struct"`
	DriveClass               *RecognizeVietnamIdentityCardResponseBodyDataDriveClass               `json:"DriveClass,omitempty" xml:"DriveClass,omitempty" type:"Struct"`
	ExpiryDate               *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate               `json:"ExpiryDate,omitempty" xml:"ExpiryDate,omitempty" type:"Struct"`
	FullName                 *RecognizeVietnamIdentityCardResponseBodyDataFullName                 `json:"FullName,omitempty" xml:"FullName,omitempty" type:"Struct"`
	IdNumber                 *RecognizeVietnamIdentityCardResponseBodyDataIdNumber                 `json:"IdNumber,omitempty" xml:"IdNumber,omitempty" type:"Struct"`
	Nationality              *RecognizeVietnamIdentityCardResponseBodyDataNationality              `json:"Nationality,omitempty" xml:"Nationality,omitempty" type:"Struct"`
	OriginPlaceFirstLine     *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine     `json:"OriginPlaceFirstLine,omitempty" xml:"OriginPlaceFirstLine,omitempty" type:"Struct"`
	OriginPlaceSecondLine    *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine    `json:"OriginPlaceSecondLine,omitempty" xml:"OriginPlaceSecondLine,omitempty" type:"Struct"`
	PortraitBox              *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox              `json:"PortraitBox,omitempty" xml:"PortraitBox,omitempty" type:"Struct"`
	ResidencePlaceFirstLine  *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine  `json:"ResidencePlaceFirstLine,omitempty" xml:"ResidencePlaceFirstLine,omitempty" type:"Struct"`
	ResidencePlaceSecondLine *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine `json:"ResidencePlaceSecondLine,omitempty" xml:"ResidencePlaceSecondLine,omitempty" type:"Struct"`
	Sex                      *RecognizeVietnamIdentityCardResponseBodyDataSex                      `json:"Sex,omitempty" xml:"Sex,omitempty" type:"Struct"`
}

func (s RecognizeVietnamIdentityCardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetAddressFirstLine(v *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine) *RecognizeVietnamIdentityCardResponseBodyData {
	s.AddressFirstLine = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetAddressSecondLine(v *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine) *RecognizeVietnamIdentityCardResponseBodyData {
	s.AddressSecondLine = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetBirthDate(v *RecognizeVietnamIdentityCardResponseBodyDataBirthDate) *RecognizeVietnamIdentityCardResponseBodyData {
	s.BirthDate = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetCardBox(v *RecognizeVietnamIdentityCardResponseBodyDataCardBox) *RecognizeVietnamIdentityCardResponseBodyData {
	s.CardBox = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetDriveClass(v *RecognizeVietnamIdentityCardResponseBodyDataDriveClass) *RecognizeVietnamIdentityCardResponseBodyData {
	s.DriveClass = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetExpiryDate(v *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate) *RecognizeVietnamIdentityCardResponseBodyData {
	s.ExpiryDate = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetFullName(v *RecognizeVietnamIdentityCardResponseBodyDataFullName) *RecognizeVietnamIdentityCardResponseBodyData {
	s.FullName = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetIdNumber(v *RecognizeVietnamIdentityCardResponseBodyDataIdNumber) *RecognizeVietnamIdentityCardResponseBodyData {
	s.IdNumber = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetNationality(v *RecognizeVietnamIdentityCardResponseBodyDataNationality) *RecognizeVietnamIdentityCardResponseBodyData {
	s.Nationality = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetOriginPlaceFirstLine(v *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine) *RecognizeVietnamIdentityCardResponseBodyData {
	s.OriginPlaceFirstLine = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetOriginPlaceSecondLine(v *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine) *RecognizeVietnamIdentityCardResponseBodyData {
	s.OriginPlaceSecondLine = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetPortraitBox(v *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox) *RecognizeVietnamIdentityCardResponseBodyData {
	s.PortraitBox = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetResidencePlaceFirstLine(v *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine) *RecognizeVietnamIdentityCardResponseBodyData {
	s.ResidencePlaceFirstLine = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetResidencePlaceSecondLine(v *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine) *RecognizeVietnamIdentityCardResponseBodyData {
	s.ResidencePlaceSecondLine = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyData) SetSex(v *RecognizeVietnamIdentityCardResponseBodyDataSex) *RecognizeVietnamIdentityCardResponseBodyData {
	s.Sex = v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine) SetScore(v float32) *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLine {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataAddressFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                                  `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine) SetScore(v float32) *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLine {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataAddressSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataBirthDate struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *float32                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataBirthDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataBirthDate) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataBirthDate) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataBirthDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataBirthDate) SetScore(v float32) *RecognizeVietnamIdentityCardResponseBodyDataBirthDate {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataBirthDate) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataBirthDate {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataBirthDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataCardBox struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                         `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataCardBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataCardBox) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataCardBox) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataCardBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataCardBox) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataCardBox {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataCardBox) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataCardBox {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataCardBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataDriveClass struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataDriveClass) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataDriveClass) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataDriveClass) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataDriveClass {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataDriveClass) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataDriveClass {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataDriveClass) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataDriveClass {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataDriveClassKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataExpiryDate struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                            `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                            `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataExpiryDate) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataExpiryDate) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataExpiryDate {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataExpiryDateKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataFullName struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataFullName) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataFullName) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataFullName) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataFullName {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataFullName) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataFullName {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataFullName) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataFullName {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataFullNameKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataIdNumber struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataIdNumber) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataIdNumber) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataIdNumber) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataIdNumber {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataIdNumber) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataIdNumber {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataIdNumber) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataIdNumber {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataIdNumberKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataNationality struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataNationality) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataNationality) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataNationality) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataNationality {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataNationality) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataNationality {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataNationality) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataNationality {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataNationalityKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                      `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                      `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLine {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                       `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                       `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLine {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataOriginPlaceSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataPortraitBox struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                             `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataPortraitBox) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataPortraitBox) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataPortraitBox {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataPortraitBoxKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                         `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                         `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLine {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceFirstLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                                          `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                                          `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLine {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataResidencePlaceSecondLineKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataSex struct {
	KeyPoints []*RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	Score     *string                                                     `json:"Score,omitempty" xml:"Score,omitempty"`
	Text      *string                                                     `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataSex) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataSex) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataSex) SetKeyPoints(v []*RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints) *RecognizeVietnamIdentityCardResponseBodyDataSex {
	s.KeyPoints = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataSex) SetScore(v string) *RecognizeVietnamIdentityCardResponseBodyDataSex {
	s.Score = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataSex) SetText(v string) *RecognizeVietnamIdentityCardResponseBodyDataSex {
	s.Text = &v
	return s
}

type RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints struct {
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints) SetX(v float32) *RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints {
	s.X = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints) SetY(v float32) *RecognizeVietnamIdentityCardResponseBodyDataSexKeyPoints {
	s.Y = &v
	return s
}

type RecognizeVietnamIdentityCardResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVietnamIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVietnamIdentityCardResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVietnamIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVietnamIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeVietnamIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVietnamIdentityCardResponse) SetStatusCode(v int32) *RecognizeVietnamIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVietnamIdentityCardResponse) SetBody(v *RecognizeVietnamIdentityCardResponseBody) *RecognizeVietnamIdentityCardResponse {
	s.Body = v
	return s
}

type TrimDocumentRequest struct {
	FileType   *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileURL    *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s TrimDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentRequest) GoString() string {
	return s.String()
}

func (s *TrimDocumentRequest) SetFileType(v string) *TrimDocumentRequest {
	s.FileType = &v
	return s
}

func (s *TrimDocumentRequest) SetFileURL(v string) *TrimDocumentRequest {
	s.FileURL = &v
	return s
}

func (s *TrimDocumentRequest) SetOutputType(v string) *TrimDocumentRequest {
	s.OutputType = &v
	return s
}

type TrimDocumentAdvanceRequest struct {
	FileType      *string   `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FileURLObject io.Reader `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	OutputType    *string   `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s TrimDocumentAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TrimDocumentAdvanceRequest) SetFileType(v string) *TrimDocumentAdvanceRequest {
	s.FileType = &v
	return s
}

func (s *TrimDocumentAdvanceRequest) SetFileURLObject(v io.Reader) *TrimDocumentAdvanceRequest {
	s.FileURLObject = v
	return s
}

func (s *TrimDocumentAdvanceRequest) SetOutputType(v string) *TrimDocumentAdvanceRequest {
	s.OutputType = &v
	return s
}

type TrimDocumentResponseBody struct {
	Data      *TrimDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TrimDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrimDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *TrimDocumentResponseBody) SetData(v *TrimDocumentResponseBodyData) *TrimDocumentResponseBody {
	s.Data = v
	return s
}

func (s *TrimDocumentResponseBody) SetRequestId(v string) *TrimDocumentResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrimDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TrimDocumentResponse) SetStatusCode(v int32) *TrimDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *TrimDocumentResponse) SetBody(v *TrimDocumentResponseBody) *TrimDocumentResponse {
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

func (client *Client) DetectCardScreenshotWithOptions(request *DetectCardScreenshotRequest, runtime *util.RuntimeOptions) (_result *DetectCardScreenshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectCardScreenshot"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectCardScreenshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectCardScreenshotReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectCardScreenshotResp, _err := client.DetectCardScreenshotWithOptions(detectCardScreenshotReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectCardScreenshotResp
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RecognizeAccountPageWithOptions(request *RecognizeAccountPageRequest, runtime *util.RuntimeOptions) (_result *RecognizeAccountPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeAccountPage"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeAccountPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeAccountPageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeAccountPageResp, _err := client.RecognizeAccountPageWithOptions(recognizeAccountPageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeAccountPageResp
	return _result, _err
}

func (client *Client) RecognizeBankCardWithOptions(request *RecognizeBankCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeBankCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBankCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBankCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeBankCardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeBankCardResp, _err := client.RecognizeBankCardWithOptions(recognizeBankCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeBankCardResp
	return _result, _err
}

func (client *Client) RecognizeBusinessCardWithOptions(request *RecognizeBusinessCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBusinessCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBusinessCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeBusinessCardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeBusinessCardResp, _err := client.RecognizeBusinessCardWithOptions(recognizeBusinessCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeBusinessCardResp
	return _result, _err
}

func (client *Client) RecognizeBusinessLicenseWithOptions(request *RecognizeBusinessLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeBusinessLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeBusinessLicense"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeBusinessLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeBusinessLicenseReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeBusinessLicenseResp, _err := client.RecognizeBusinessLicenseWithOptions(recognizeBusinessLicenseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeBusinessLicenseResp
	return _result, _err
}

func (client *Client) RecognizeCharacterWithOptions(request *RecognizeCharacterRequest, runtime *util.RuntimeOptions) (_result *RecognizeCharacterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.MinHeight)) {
		body["MinHeight"] = request.MinHeight
	}

	if !tea.BoolValue(util.IsUnset(request.OutputProbability)) {
		body["OutputProbability"] = request.OutputProbability
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeCharacter"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeCharacterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeCharacterReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeCharacterResp, _err := client.RecognizeCharacterWithOptions(recognizeCharacterReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeCharacterResp
	return _result, _err
}

func (client *Client) RecognizeChinapassportWithOptions(request *RecognizeChinapassportRequest, runtime *util.RuntimeOptions) (_result *RecognizeChinapassportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeChinapassport"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeChinapassportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeChinapassportReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeChinapassportResp, _err := client.RecognizeChinapassportWithOptions(recognizeChinapassportReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeChinapassportResp
	return _result, _err
}

func (client *Client) RecognizeDriverLicenseWithOptions(request *RecognizeDriverLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeDriverLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Side)) {
		body["Side"] = request.Side
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeDriverLicense"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeDriverLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeDriverLicenseReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeDriverLicenseResp, _err := client.RecognizeDriverLicenseWithOptions(recognizeDriverLicenseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeDriverLicenseResp
	return _result, _err
}

func (client *Client) RecognizeDrivingLicenseWithOptions(request *RecognizeDrivingLicenseRequest, runtime *util.RuntimeOptions) (_result *RecognizeDrivingLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Side)) {
		body["Side"] = request.Side
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeDrivingLicense"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeDrivingLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeDrivingLicenseReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeDrivingLicenseResp, _err := client.RecognizeDrivingLicenseWithOptions(recognizeDrivingLicenseReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeDrivingLicenseResp
	return _result, _err
}

func (client *Client) RecognizeIdentityCardWithOptions(request *RecognizeIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Side)) {
		body["Side"] = request.Side
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeIdentityCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeIdentityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeIdentityCardReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeIdentityCardResp, _err := client.RecognizeIdentityCardWithOptions(recognizeIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeIdentityCardResp
	return _result, _err
}

func (client *Client) RecognizeIndonesiaIdentityCardWithOptions(request *RecognizeIndonesiaIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeIndonesiaIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeIndonesiaIdentityCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeIndonesiaIdentityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeIndonesiaIdentityCard(request *RecognizeIndonesiaIdentityCardRequest) (_result *RecognizeIndonesiaIdentityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeIndonesiaIdentityCardResponse{}
	_body, _err := client.RecognizeIndonesiaIdentityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeIndonesiaIdentityCardAdvance(request *RecognizeIndonesiaIdentityCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeIndonesiaIdentityCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeIndonesiaIdentityCardReq := &RecognizeIndonesiaIdentityCardRequest{}
	openapiutil.Convert(request, recognizeIndonesiaIdentityCardReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeIndonesiaIdentityCardReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeIndonesiaIdentityCardResp, _err := client.RecognizeIndonesiaIdentityCardWithOptions(recognizeIndonesiaIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeIndonesiaIdentityCardResp
	return _result, _err
}

func (client *Client) RecognizeLicensePlateWithOptions(request *RecognizeLicensePlateRequest, runtime *util.RuntimeOptions) (_result *RecognizeLicensePlateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeLicensePlate"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeLicensePlateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeLicensePlateReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeLicensePlateResp, _err := client.RecognizeLicensePlateWithOptions(recognizeLicensePlateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeLicensePlateResp
	return _result, _err
}

func (client *Client) RecognizeMalaysiaIdentityCardWithOptions(request *RecognizeMalaysiaIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeMalaysiaIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeMalaysiaIdentityCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeMalaysiaIdentityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeMalaysiaIdentityCard(request *RecognizeMalaysiaIdentityCardRequest) (_result *RecognizeMalaysiaIdentityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeMalaysiaIdentityCardResponse{}
	_body, _err := client.RecognizeMalaysiaIdentityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeMalaysiaIdentityCardAdvance(request *RecognizeMalaysiaIdentityCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeMalaysiaIdentityCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeMalaysiaIdentityCardReq := &RecognizeMalaysiaIdentityCardRequest{}
	openapiutil.Convert(request, recognizeMalaysiaIdentityCardReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeMalaysiaIdentityCardReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeMalaysiaIdentityCardResp, _err := client.RecognizeMalaysiaIdentityCardWithOptions(recognizeMalaysiaIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeMalaysiaIdentityCardResp
	return _result, _err
}

func (client *Client) RecognizePassportMRZWithOptions(request *RecognizePassportMRZRequest, runtime *util.RuntimeOptions) (_result *RecognizePassportMRZResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePassportMRZ"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePassportMRZResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizePassportMRZReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizePassportMRZResp, _err := client.RecognizePassportMRZWithOptions(recognizePassportMRZReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizePassportMRZResp
	return _result, _err
}

func (client *Client) RecognizePdfWithOptions(request *RecognizePdfRequest, runtime *util.RuntimeOptions) (_result *RecognizePdfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileURL)) {
		body["FileURL"] = request.FileURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePdf"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePdfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizePdf(request *RecognizePdfRequest) (_result *RecognizePdfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizePdfResponse{}
	_body, _err := client.RecognizePdfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizePdfAdvance(request *RecognizePdfAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizePdfResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizePdfReq := &RecognizePdfRequest{}
	openapiutil.Convert(request, recognizePdfReq)
	if !tea.BoolValue(util.IsUnset(request.FileURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizePdfReq.FileURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizePdfResp, _err := client.RecognizePdfWithOptions(recognizePdfReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizePdfResp
	return _result, _err
}

func (client *Client) RecognizePoiNameWithOptions(request *RecognizePoiNameRequest, runtime *util.RuntimeOptions) (_result *RecognizePoiNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizePoiName"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizePoiNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizePoiNameReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tasks)) {
		body["Tasks"] = request.Tasks
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeQrCode"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeQrCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) RecognizeQuotaInvoiceWithOptions(request *RecognizeQuotaInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeQuotaInvoice"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeQuotaInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeQuotaInvoice(request *RecognizeQuotaInvoiceRequest) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeQuotaInvoiceResponse{}
	_body, _err := client.RecognizeQuotaInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeQuotaInvoiceAdvance(request *RecognizeQuotaInvoiceAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeQuotaInvoiceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeQuotaInvoiceReq := &RecognizeQuotaInvoiceRequest{}
	openapiutil.Convert(request, recognizeQuotaInvoiceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeQuotaInvoiceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeQuotaInvoiceResp, _err := client.RecognizeQuotaInvoiceWithOptions(recognizeQuotaInvoiceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeQuotaInvoiceResp
	return _result, _err
}

func (client *Client) RecognizeRussiaIdentityCardWithOptions(request *RecognizeRussiaIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeRussiaIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeRussiaIdentityCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeRussiaIdentityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeRussiaIdentityCard(request *RecognizeRussiaIdentityCardRequest) (_result *RecognizeRussiaIdentityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeRussiaIdentityCardResponse{}
	_body, _err := client.RecognizeRussiaIdentityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeRussiaIdentityCardAdvance(request *RecognizeRussiaIdentityCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeRussiaIdentityCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeRussiaIdentityCardReq := &RecognizeRussiaIdentityCardRequest{}
	openapiutil.Convert(request, recognizeRussiaIdentityCardReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeRussiaIdentityCardReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeRussiaIdentityCardResp, _err := client.RecognizeRussiaIdentityCardWithOptions(recognizeRussiaIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeRussiaIdentityCardResp
	return _result, _err
}

func (client *Client) RecognizeStampWithOptions(request *RecognizeStampRequest, runtime *util.RuntimeOptions) (_result *RecognizeStampResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeStamp"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeStampResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeStampReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeStampResp, _err := client.RecognizeStampWithOptions(recognizeStampReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeStampResp
	return _result, _err
}

func (client *Client) RecognizeTableWithOptions(request *RecognizeTableRequest, runtime *util.RuntimeOptions) (_result *RecognizeTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssureDirection)) {
		body["AssureDirection"] = request.AssureDirection
	}

	if !tea.BoolValue(util.IsUnset(request.HasLine)) {
		body["HasLine"] = request.HasLine
	}

	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.OutputFormat)) {
		body["OutputFormat"] = request.OutputFormat
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDetection)) {
		body["SkipDetection"] = request.SkipDetection
	}

	if !tea.BoolValue(util.IsUnset(request.UseFinanceModel)) {
		body["UseFinanceModel"] = request.UseFinanceModel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTable"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeTableReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeTableResp, _err := client.RecognizeTableWithOptions(recognizeTableReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTableResp
	return _result, _err
}

func (client *Client) RecognizeTakeoutOrderWithOptions(request *RecognizeTakeoutOrderRequest, runtime *util.RuntimeOptions) (_result *RecognizeTakeoutOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTakeoutOrder"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTakeoutOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeTakeoutOrderReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeTakeoutOrderResp, _err := client.RecognizeTakeoutOrderWithOptions(recognizeTakeoutOrderReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTakeoutOrderResp
	return _result, _err
}

func (client *Client) RecognizeTaxiInvoiceWithOptions(request *RecognizeTaxiInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTaxiInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTaxiInvoice"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTaxiInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeTaxiInvoiceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeTaxiInvoiceResp, _err := client.RecognizeTaxiInvoiceWithOptions(recognizeTaxiInvoiceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTaxiInvoiceResp
	return _result, _err
}

func (client *Client) RecognizeTicketInvoiceWithOptions(request *RecognizeTicketInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTicketInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTicketInvoice"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTicketInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTicketInvoice(request *RecognizeTicketInvoiceRequest) (_result *RecognizeTicketInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTicketInvoiceResponse{}
	_body, _err := client.RecognizeTicketInvoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTicketInvoiceAdvance(request *RecognizeTicketInvoiceAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTicketInvoiceResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeTicketInvoiceReq := &RecognizeTicketInvoiceRequest{}
	openapiutil.Convert(request, recognizeTicketInvoiceReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeTicketInvoiceReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeTicketInvoiceResp, _err := client.RecognizeTicketInvoiceWithOptions(recognizeTicketInvoiceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTicketInvoiceResp
	return _result, _err
}

func (client *Client) RecognizeTrainTicketWithOptions(request *RecognizeTrainTicketRequest, runtime *util.RuntimeOptions) (_result *RecognizeTrainTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTrainTicket"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTrainTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeTrainTicketReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeTrainTicketResp, _err := client.RecognizeTrainTicketWithOptions(recognizeTrainTicketReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTrainTicketResp
	return _result, _err
}

func (client *Client) RecognizeTurkeyIdentityCardWithOptions(request *RecognizeTurkeyIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeTurkeyIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeTurkeyIdentityCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeTurkeyIdentityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeTurkeyIdentityCard(request *RecognizeTurkeyIdentityCardRequest) (_result *RecognizeTurkeyIdentityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeTurkeyIdentityCardResponse{}
	_body, _err := client.RecognizeTurkeyIdentityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeTurkeyIdentityCardAdvance(request *RecognizeTurkeyIdentityCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeTurkeyIdentityCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeTurkeyIdentityCardReq := &RecognizeTurkeyIdentityCardRequest{}
	openapiutil.Convert(request, recognizeTurkeyIdentityCardReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeTurkeyIdentityCardReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeTurkeyIdentityCardResp, _err := client.RecognizeTurkeyIdentityCardWithOptions(recognizeTurkeyIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeTurkeyIdentityCardResp
	return _result, _err
}

func (client *Client) RecognizeUkraineIdentityCardWithOptions(request *RecognizeUkraineIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeUkraineIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeUkraineIdentityCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeUkraineIdentityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeUkraineIdentityCard(request *RecognizeUkraineIdentityCardRequest) (_result *RecognizeUkraineIdentityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeUkraineIdentityCardResponse{}
	_body, _err := client.RecognizeUkraineIdentityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeUkraineIdentityCardAdvance(request *RecognizeUkraineIdentityCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeUkraineIdentityCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeUkraineIdentityCardReq := &RecognizeUkraineIdentityCardRequest{}
	openapiutil.Convert(request, recognizeUkraineIdentityCardReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeUkraineIdentityCardReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeUkraineIdentityCardResp, _err := client.RecognizeUkraineIdentityCardWithOptions(recognizeUkraineIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeUkraineIdentityCardResp
	return _result, _err
}

func (client *Client) RecognizeVATInvoiceWithOptions(request *RecognizeVATInvoiceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVATInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.FileURL)) {
		body["FileURL"] = request.FileURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVATInvoice"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVATInvoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.FileURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVATInvoiceReq.FileURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVATInvoiceResp, _err := client.RecognizeVATInvoiceWithOptions(recognizeVATInvoiceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVATInvoiceResp
	return _result, _err
}

func (client *Client) RecognizeVINCodeWithOptions(request *RecognizeVINCodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeVINCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		query["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVINCode"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVINCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVINCodeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVINCodeResp, _err := client.RecognizeVINCodeWithOptions(recognizeVINCodeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVINCodeResp
	return _result, _err
}

func (client *Client) RecognizeVerificationcodeWithOptions(request *RecognizeVerificationcodeRequest, runtime *util.RuntimeOptions) (_result *RecognizeVerificationcodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVerificationcode"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVerificationcodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVerificationcodeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVerificationcodeResp, _err := client.RecognizeVerificationcodeWithOptions(recognizeVerificationcodeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVerificationcodeResp
	return _result, _err
}

func (client *Client) RecognizeVideoCastCrewListWithOptions(tmpReq *RecognizeVideoCastCrewListRequest, runtime *util.RuntimeOptions) (_result *RecognizeVideoCastCrewListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RecognizeVideoCastCrewListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Params)) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, tea.String("Params"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParamsShrink)) {
		body["Params"] = request.ParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterUrl)) {
		body["RegisterUrl"] = request.RegisterUrl
	}

	if !tea.BoolValue(util.IsUnset(request.VideoUrl)) {
		body["VideoUrl"] = request.VideoUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVideoCastCrewList"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVideoCastCrewListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVideoCastCrewList(request *RecognizeVideoCastCrewListRequest) (_result *RecognizeVideoCastCrewListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVideoCastCrewListResponse{}
	_body, _err := client.RecognizeVideoCastCrewListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVideoCastCrewListAdvance(request *RecognizeVideoCastCrewListAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVideoCastCrewListResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeVideoCastCrewListReq := &RecognizeVideoCastCrewListRequest{}
	openapiutil.Convert(request, recognizeVideoCastCrewListReq)
	if !tea.BoolValue(util.IsUnset(request.VideoUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVideoCastCrewListReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVideoCastCrewListResp, _err := client.RecognizeVideoCastCrewListWithOptions(recognizeVideoCastCrewListReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVideoCastCrewListResp
	return _result, _err
}

func (client *Client) RecognizeVideoCharacterWithOptions(request *RecognizeVideoCharacterRequest, runtime *util.RuntimeOptions) (_result *RecognizeVideoCharacterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVideoCharacter"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVideoCharacterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVideoCharacter(request *RecognizeVideoCharacterRequest) (_result *RecognizeVideoCharacterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVideoCharacterResponse{}
	_body, _err := client.RecognizeVideoCharacterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVietnamIdentityCardWithOptions(request *RecognizeVietnamIdentityCardRequest, runtime *util.RuntimeOptions) (_result *RecognizeVietnamIdentityCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUrl)) {
		body["ImageUrl"] = request.ImageUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVietnamIdentityCard"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVietnamIdentityCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVietnamIdentityCard(request *RecognizeVietnamIdentityCardRequest) (_result *RecognizeVietnamIdentityCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVietnamIdentityCardResponse{}
	_body, _err := client.RecognizeVietnamIdentityCardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVietnamIdentityCardAdvance(request *RecognizeVietnamIdentityCardAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVietnamIdentityCardResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	recognizeVietnamIdentityCardReq := &RecognizeVietnamIdentityCardRequest{}
	openapiutil.Convert(request, recognizeVietnamIdentityCardReq)
	if !tea.BoolValue(util.IsUnset(request.ImageUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVietnamIdentityCardReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVietnamIdentityCardResp, _err := client.RecognizeVietnamIdentityCardWithOptions(recognizeVietnamIdentityCardReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVietnamIdentityCardResp
	return _result, _err
}

func (client *Client) TrimDocumentWithOptions(request *TrimDocumentRequest, runtime *util.RuntimeOptions) (_result *TrimDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		body["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.FileURL)) {
		body["FileURL"] = request.FileURL
	}

	if !tea.BoolValue(util.IsUnset(request.OutputType)) {
		body["OutputType"] = request.OutputType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TrimDocument"),
		Version:     tea.String("2019-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TrimDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
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
	if !tea.BoolValue(util.IsUnset(request.FileURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		trimDocumentReq.FileURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	trimDocumentResp, _err := client.TrimDocumentWithOptions(trimDocumentReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = trimDocumentResp
	return _result, _err
}
