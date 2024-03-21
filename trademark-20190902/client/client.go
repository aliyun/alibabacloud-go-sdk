// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BindApplicantRequest struct {
	ApplicantId         *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	AuthorizationOssKey *string `json:"AuthorizationOssKey,omitempty" xml:"AuthorizationOssKey,omitempty"`
	BizId               *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s BindApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s BindApplicantRequest) GoString() string {
	return s.String()
}

func (s *BindApplicantRequest) SetApplicantId(v string) *BindApplicantRequest {
	s.ApplicantId = &v
	return s
}

func (s *BindApplicantRequest) SetAuthorizationOssKey(v string) *BindApplicantRequest {
	s.AuthorizationOssKey = &v
	return s
}

func (s *BindApplicantRequest) SetBizId(v string) *BindApplicantRequest {
	s.BizId = &v
	return s
}

type BindApplicantResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *BindApplicantResponseBody) SetRequestId(v string) *BindApplicantResponseBody {
	s.RequestId = &v
	return s
}

type BindApplicantResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s BindApplicantResponse) GoString() string {
	return s.String()
}

func (s *BindApplicantResponse) SetHeaders(v map[string]*string) *BindApplicantResponse {
	s.Headers = v
	return s
}

func (s *BindApplicantResponse) SetStatusCode(v int32) *BindApplicantResponse {
	s.StatusCode = &v
	return s
}

func (s *BindApplicantResponse) SetBody(v *BindApplicantResponseBody) *BindApplicantResponse {
	s.Body = v
	return s
}

type CancelOrderRequest struct {
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) SetOrderId(v int64) *CancelOrderRequest {
	s.OrderId = &v
	return s
}

type CancelOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBody) SetRequestId(v string) *CancelOrderResponseBody {
	s.RequestId = &v
	return s
}

type CancelOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderResponse) SetHeaders(v map[string]*string) *CancelOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderResponse) SetStatusCode(v int32) *CancelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOrderResponse) SetBody(v *CancelOrderResponseBody) *CancelOrderResponse {
	s.Body = v
	return s
}

type CheckAuthorizationLetterRequest struct {
	ApplicantType  *string `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	Color          *string `json:"Color,omitempty" xml:"Color,omitempty"`
	ContactName    *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber  *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactZipcode *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	OssKey         *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	PersonalType   *string `json:"PersonalType,omitempty" xml:"PersonalType,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckAuthorizationLetterRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthorizationLetterRequest) GoString() string {
	return s.String()
}

func (s *CheckAuthorizationLetterRequest) SetApplicantType(v string) *CheckAuthorizationLetterRequest {
	s.ApplicantType = &v
	return s
}

func (s *CheckAuthorizationLetterRequest) SetColor(v string) *CheckAuthorizationLetterRequest {
	s.Color = &v
	return s
}

func (s *CheckAuthorizationLetterRequest) SetContactName(v string) *CheckAuthorizationLetterRequest {
	s.ContactName = &v
	return s
}

func (s *CheckAuthorizationLetterRequest) SetContactNumber(v string) *CheckAuthorizationLetterRequest {
	s.ContactNumber = &v
	return s
}

func (s *CheckAuthorizationLetterRequest) SetContactZipcode(v string) *CheckAuthorizationLetterRequest {
	s.ContactZipcode = &v
	return s
}

func (s *CheckAuthorizationLetterRequest) SetOssKey(v string) *CheckAuthorizationLetterRequest {
	s.OssKey = &v
	return s
}

func (s *CheckAuthorizationLetterRequest) SetPersonalType(v string) *CheckAuthorizationLetterRequest {
	s.PersonalType = &v
	return s
}

func (s *CheckAuthorizationLetterRequest) SetType(v string) *CheckAuthorizationLetterRequest {
	s.Type = &v
	return s
}

type CheckAuthorizationLetterResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
	Tips        *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s CheckAuthorizationLetterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthorizationLetterResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAuthorizationLetterResponseBody) SetRequestId(v string) *CheckAuthorizationLetterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAuthorizationLetterResponseBody) SetTemplateUrl(v string) *CheckAuthorizationLetterResponseBody {
	s.TemplateUrl = &v
	return s
}

func (s *CheckAuthorizationLetterResponseBody) SetTips(v string) *CheckAuthorizationLetterResponseBody {
	s.Tips = &v
	return s
}

type CheckAuthorizationLetterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAuthorizationLetterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAuthorizationLetterResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAuthorizationLetterResponse) GoString() string {
	return s.String()
}

func (s *CheckAuthorizationLetterResponse) SetHeaders(v map[string]*string) *CheckAuthorizationLetterResponse {
	s.Headers = v
	return s
}

func (s *CheckAuthorizationLetterResponse) SetStatusCode(v int32) *CheckAuthorizationLetterResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAuthorizationLetterResponse) SetBody(v *CheckAuthorizationLetterResponseBody) *CheckAuthorizationLetterResponse {
	s.Body = v
	return s
}

type CheckBizAvailableRequest struct {
	Biz   *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s CheckBizAvailableRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckBizAvailableRequest) GoString() string {
	return s.String()
}

func (s *CheckBizAvailableRequest) SetBiz(v string) *CheckBizAvailableRequest {
	s.Biz = &v
	return s
}

func (s *CheckBizAvailableRequest) SetScene(v string) *CheckBizAvailableRequest {
	s.Scene = &v
	return s
}

type CheckBizAvailableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckBizAvailableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckBizAvailableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBizAvailableResponseBody) SetRequestId(v string) *CheckBizAvailableResponseBody {
	s.RequestId = &v
	return s
}

type CheckBizAvailableResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckBizAvailableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckBizAvailableResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckBizAvailableResponse) GoString() string {
	return s.String()
}

func (s *CheckBizAvailableResponse) SetHeaders(v map[string]*string) *CheckBizAvailableResponse {
	s.Headers = v
	return s
}

func (s *CheckBizAvailableResponse) SetStatusCode(v int32) *CheckBizAvailableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBizAvailableResponse) SetBody(v *CheckBizAvailableResponseBody) *CheckBizAvailableResponse {
	s.Body = v
	return s
}

type CheckDuplicateApplicantRiskRequest struct {
	ApplicantName  *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	EventSceneType *int32  `json:"EventSceneType,omitempty" xml:"EventSceneType,omitempty"`
}

func (s CheckDuplicateApplicantRiskRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateApplicantRiskRequest) GoString() string {
	return s.String()
}

func (s *CheckDuplicateApplicantRiskRequest) SetApplicantName(v string) *CheckDuplicateApplicantRiskRequest {
	s.ApplicantName = &v
	return s
}

func (s *CheckDuplicateApplicantRiskRequest) SetEventSceneType(v int32) *CheckDuplicateApplicantRiskRequest {
	s.EventSceneType = &v
	return s
}

type CheckDuplicateApplicantRiskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckDuplicateApplicantRiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateApplicantRiskResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDuplicateApplicantRiskResponseBody) SetRequestId(v string) *CheckDuplicateApplicantRiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDuplicateApplicantRiskResponseBody) SetResult(v string) *CheckDuplicateApplicantRiskResponseBody {
	s.Result = &v
	return s
}

type CheckDuplicateApplicantRiskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDuplicateApplicantRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDuplicateApplicantRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateApplicantRiskResponse) GoString() string {
	return s.String()
}

func (s *CheckDuplicateApplicantRiskResponse) SetHeaders(v map[string]*string) *CheckDuplicateApplicantRiskResponse {
	s.Headers = v
	return s
}

func (s *CheckDuplicateApplicantRiskResponse) SetStatusCode(v int32) *CheckDuplicateApplicantRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDuplicateApplicantRiskResponse) SetBody(v *CheckDuplicateApplicantRiskResponseBody) *CheckDuplicateApplicantRiskResponse {
	s.Body = v
	return s
}

type CheckDuplicateClassificationRequest struct {
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	EventSceneType *int32  `json:"EventSceneType,omitempty" xml:"EventSceneType,omitempty"`
	TrademarkName  *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
}

func (s CheckDuplicateClassificationRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateClassificationRequest) GoString() string {
	return s.String()
}

func (s *CheckDuplicateClassificationRequest) SetClassification(v string) *CheckDuplicateClassificationRequest {
	s.Classification = &v
	return s
}

func (s *CheckDuplicateClassificationRequest) SetEventSceneType(v int32) *CheckDuplicateClassificationRequest {
	s.EventSceneType = &v
	return s
}

func (s *CheckDuplicateClassificationRequest) SetTrademarkName(v string) *CheckDuplicateClassificationRequest {
	s.TrademarkName = &v
	return s
}

type CheckDuplicateClassificationResponseBody struct {
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CheckDuplicateClassificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDuplicateClassificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateClassificationResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDuplicateClassificationResponseBody) SetCode(v string) *CheckDuplicateClassificationResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDuplicateClassificationResponseBody) SetData(v *CheckDuplicateClassificationResponseBodyData) *CheckDuplicateClassificationResponseBody {
	s.Data = v
	return s
}

func (s *CheckDuplicateClassificationResponseBody) SetMessage(v string) *CheckDuplicateClassificationResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDuplicateClassificationResponseBody) SetRequestId(v string) *CheckDuplicateClassificationResponseBody {
	s.RequestId = &v
	return s
}

type CheckDuplicateClassificationResponseBodyData struct {
	DuplicateSecondaryClassification *CheckDuplicateClassificationResponseBodyDataDuplicateSecondaryClassification `json:"DuplicateSecondaryClassification,omitempty" xml:"DuplicateSecondaryClassification,omitempty" type:"Struct"`
	TrademarkName                    *string                                                                       `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
}

func (s CheckDuplicateClassificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateClassificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckDuplicateClassificationResponseBodyData) SetDuplicateSecondaryClassification(v *CheckDuplicateClassificationResponseBodyDataDuplicateSecondaryClassification) *CheckDuplicateClassificationResponseBodyData {
	s.DuplicateSecondaryClassification = v
	return s
}

func (s *CheckDuplicateClassificationResponseBodyData) SetTrademarkName(v string) *CheckDuplicateClassificationResponseBodyData {
	s.TrademarkName = &v
	return s
}

type CheckDuplicateClassificationResponseBodyDataDuplicateSecondaryClassification struct {
	SecondaryClassification []*string `json:"SecondaryClassification,omitempty" xml:"SecondaryClassification,omitempty" type:"Repeated"`
}

func (s CheckDuplicateClassificationResponseBodyDataDuplicateSecondaryClassification) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateClassificationResponseBodyDataDuplicateSecondaryClassification) GoString() string {
	return s.String()
}

func (s *CheckDuplicateClassificationResponseBodyDataDuplicateSecondaryClassification) SetSecondaryClassification(v []*string) *CheckDuplicateClassificationResponseBodyDataDuplicateSecondaryClassification {
	s.SecondaryClassification = v
	return s
}

type CheckDuplicateClassificationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDuplicateClassificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDuplicateClassificationResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateClassificationResponse) GoString() string {
	return s.String()
}

func (s *CheckDuplicateClassificationResponse) SetHeaders(v map[string]*string) *CheckDuplicateClassificationResponse {
	s.Headers = v
	return s
}

func (s *CheckDuplicateClassificationResponse) SetStatusCode(v int32) *CheckDuplicateClassificationResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDuplicateClassificationResponse) SetBody(v *CheckDuplicateClassificationResponseBody) *CheckDuplicateClassificationResponse {
	s.Body = v
	return s
}

type CheckDuplicateTrademarkRequest struct {
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	EventSceneType *int32  `json:"EventSceneType,omitempty" xml:"EventSceneType,omitempty"`
	MaterialName   *string `json:"MaterialName,omitempty" xml:"MaterialName,omitempty"`
	TrademarkName  *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
}

func (s CheckDuplicateTrademarkRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateTrademarkRequest) GoString() string {
	return s.String()
}

func (s *CheckDuplicateTrademarkRequest) SetClassification(v string) *CheckDuplicateTrademarkRequest {
	s.Classification = &v
	return s
}

func (s *CheckDuplicateTrademarkRequest) SetEventSceneType(v int32) *CheckDuplicateTrademarkRequest {
	s.EventSceneType = &v
	return s
}

func (s *CheckDuplicateTrademarkRequest) SetMaterialName(v string) *CheckDuplicateTrademarkRequest {
	s.MaterialName = &v
	return s
}

func (s *CheckDuplicateTrademarkRequest) SetTrademarkName(v string) *CheckDuplicateTrademarkRequest {
	s.TrademarkName = &v
	return s
}

type CheckDuplicateTrademarkResponseBody struct {
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DuplicateTrademark *string `json:"DuplicateTrademark,omitempty" xml:"DuplicateTrademark,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CheckDuplicateTrademarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateTrademarkResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDuplicateTrademarkResponseBody) SetCode(v string) *CheckDuplicateTrademarkResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDuplicateTrademarkResponseBody) SetDuplicateTrademark(v string) *CheckDuplicateTrademarkResponseBody {
	s.DuplicateTrademark = &v
	return s
}

func (s *CheckDuplicateTrademarkResponseBody) SetMessage(v string) *CheckDuplicateTrademarkResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDuplicateTrademarkResponseBody) SetRequestId(v string) *CheckDuplicateTrademarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDuplicateTrademarkResponseBody) SetType(v string) *CheckDuplicateTrademarkResponseBody {
	s.Type = &v
	return s
}

type CheckDuplicateTrademarkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDuplicateTrademarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDuplicateTrademarkResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDuplicateTrademarkResponse) GoString() string {
	return s.String()
}

func (s *CheckDuplicateTrademarkResponse) SetHeaders(v map[string]*string) *CheckDuplicateTrademarkResponse {
	s.Headers = v
	return s
}

func (s *CheckDuplicateTrademarkResponse) SetStatusCode(v int32) *CheckDuplicateTrademarkResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDuplicateTrademarkResponse) SetBody(v *CheckDuplicateTrademarkResponseBody) *CheckDuplicateTrademarkResponse {
	s.Body = v
	return s
}

type CheckMaterialValidityRequest struct {
	BusinessLicenseOssKey *string `json:"BusinessLicenseOssKey,omitempty" xml:"BusinessLicenseOssKey,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	IdCardName            *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber          *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	IdCardOssKey          *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	MaterialId            *int64  `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	MaterialRegion        *int32  `json:"MaterialRegion,omitempty" xml:"MaterialRegion,omitempty"`
	MaterialType          *int32  `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalType          *int64  `json:"PersonalType,omitempty" xml:"PersonalType,omitempty"`
}

func (s CheckMaterialValidityRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckMaterialValidityRequest) GoString() string {
	return s.String()
}

func (s *CheckMaterialValidityRequest) SetBusinessLicenseOssKey(v string) *CheckMaterialValidityRequest {
	s.BusinessLicenseOssKey = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetCardNumber(v string) *CheckMaterialValidityRequest {
	s.CardNumber = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetIdCardName(v string) *CheckMaterialValidityRequest {
	s.IdCardName = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetIdCardNumber(v string) *CheckMaterialValidityRequest {
	s.IdCardNumber = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetIdCardOssKey(v string) *CheckMaterialValidityRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetMaterialId(v int64) *CheckMaterialValidityRequest {
	s.MaterialId = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetMaterialRegion(v int32) *CheckMaterialValidityRequest {
	s.MaterialRegion = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetMaterialType(v int32) *CheckMaterialValidityRequest {
	s.MaterialType = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetName(v string) *CheckMaterialValidityRequest {
	s.Name = &v
	return s
}

func (s *CheckMaterialValidityRequest) SetPersonalType(v int64) *CheckMaterialValidityRequest {
	s.PersonalType = &v
	return s
}

type CheckMaterialValidityResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMaterialValidityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMaterialValidityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMaterialValidityResponseBody) SetRequestId(v string) *CheckMaterialValidityResponseBody {
	s.RequestId = &v
	return s
}

type CheckMaterialValidityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMaterialValidityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMaterialValidityResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMaterialValidityResponse) GoString() string {
	return s.String()
}

func (s *CheckMaterialValidityResponse) SetHeaders(v map[string]*string) *CheckMaterialValidityResponse {
	s.Headers = v
	return s
}

func (s *CheckMaterialValidityResponse) SetStatusCode(v int32) *CheckMaterialValidityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMaterialValidityResponse) SetBody(v *CheckMaterialValidityResponseBody) *CheckMaterialValidityResponse {
	s.Body = v
	return s
}

type CheckTrademarkNameRequest struct {
	EventSceneType *int32  `json:"EventSceneType,omitempty" xml:"EventSceneType,omitempty"`
	TrademarkName  *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
}

func (s CheckTrademarkNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkNameRequest) GoString() string {
	return s.String()
}

func (s *CheckTrademarkNameRequest) SetEventSceneType(v int32) *CheckTrademarkNameRequest {
	s.EventSceneType = &v
	return s
}

func (s *CheckTrademarkNameRequest) SetTrademarkName(v string) *CheckTrademarkNameRequest {
	s.TrademarkName = &v
	return s
}

type CheckTrademarkNameResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckTrademarkNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTrademarkNameResponseBody) SetMessage(v string) *CheckTrademarkNameResponseBody {
	s.Message = &v
	return s
}

func (s *CheckTrademarkNameResponseBody) SetRequestId(v string) *CheckTrademarkNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTrademarkNameResponseBody) SetResult(v string) *CheckTrademarkNameResponseBody {
	s.Result = &v
	return s
}

type CheckTrademarkNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTrademarkNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTrademarkNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckTrademarkNameResponse) GoString() string {
	return s.String()
}

func (s *CheckTrademarkNameResponse) SetHeaders(v map[string]*string) *CheckTrademarkNameResponse {
	s.Headers = v
	return s
}

func (s *CheckTrademarkNameResponse) SetStatusCode(v int32) *CheckTrademarkNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTrademarkNameResponse) SetBody(v *CheckTrademarkNameResponseBody) *CheckTrademarkNameResponse {
	s.Body = v
	return s
}

type CloseTrademarkApplicationRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s CloseTrademarkApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTrademarkApplicationRequest) GoString() string {
	return s.String()
}

func (s *CloseTrademarkApplicationRequest) SetBizId(v string) *CloseTrademarkApplicationRequest {
	s.BizId = &v
	return s
}

type CloseTrademarkApplicationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseTrademarkApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTrademarkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTrademarkApplicationResponseBody) SetCode(v string) *CloseTrademarkApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTrademarkApplicationResponseBody) SetMessage(v string) *CloseTrademarkApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTrademarkApplicationResponseBody) SetRequestId(v string) *CloseTrademarkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTrademarkApplicationResponseBody) SetSuccess(v bool) *CloseTrademarkApplicationResponseBody {
	s.Success = &v
	return s
}

type CloseTrademarkApplicationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseTrademarkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseTrademarkApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTrademarkApplicationResponse) GoString() string {
	return s.String()
}

func (s *CloseTrademarkApplicationResponse) SetHeaders(v map[string]*string) *CloseTrademarkApplicationResponse {
	s.Headers = v
	return s
}

func (s *CloseTrademarkApplicationResponse) SetStatusCode(v int32) *CloseTrademarkApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseTrademarkApplicationResponse) SetBody(v *CloseTrademarkApplicationResponseBody) *CloseTrademarkApplicationResponse {
	s.Body = v
	return s
}

type CombineAuthorizationLetterRequest struct {
	Address         *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantType   *string `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	ContactName     *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactPhone    *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	ContactPostcode *string `json:"ContactPostcode,omitempty" xml:"ContactPostcode,omitempty"`
	MaterialId      *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	MaterialName    *string `json:"MaterialName,omitempty" xml:"MaterialName,omitempty"`
	Nationality     *string `json:"Nationality,omitempty" xml:"Nationality,omitempty"`
	PersonalType    *string `json:"PersonalType,omitempty" xml:"PersonalType,omitempty"`
	PrincipalName   *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	TmProduceType   *string `json:"TmProduceType,omitempty" xml:"TmProduceType,omitempty"`
	TrademarkName   *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
}

func (s CombineAuthorizationLetterRequest) String() string {
	return tea.Prettify(s)
}

func (s CombineAuthorizationLetterRequest) GoString() string {
	return s.String()
}

func (s *CombineAuthorizationLetterRequest) SetAddress(v string) *CombineAuthorizationLetterRequest {
	s.Address = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetApplicantType(v string) *CombineAuthorizationLetterRequest {
	s.ApplicantType = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetContactName(v string) *CombineAuthorizationLetterRequest {
	s.ContactName = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetContactPhone(v string) *CombineAuthorizationLetterRequest {
	s.ContactPhone = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetContactPostcode(v string) *CombineAuthorizationLetterRequest {
	s.ContactPostcode = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetMaterialId(v string) *CombineAuthorizationLetterRequest {
	s.MaterialId = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetMaterialName(v string) *CombineAuthorizationLetterRequest {
	s.MaterialName = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetNationality(v string) *CombineAuthorizationLetterRequest {
	s.Nationality = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetPersonalType(v string) *CombineAuthorizationLetterRequest {
	s.PersonalType = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetPrincipalName(v int32) *CombineAuthorizationLetterRequest {
	s.PrincipalName = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetTmProduceType(v string) *CombineAuthorizationLetterRequest {
	s.TmProduceType = &v
	return s
}

func (s *CombineAuthorizationLetterRequest) SetTrademarkName(v string) *CombineAuthorizationLetterRequest {
	s.TrademarkName = &v
	return s
}

type CombineAuthorizationLetterResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateCombineUrl *string `json:"TemplateCombineUrl,omitempty" xml:"TemplateCombineUrl,omitempty"`
}

func (s CombineAuthorizationLetterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CombineAuthorizationLetterResponseBody) GoString() string {
	return s.String()
}

func (s *CombineAuthorizationLetterResponseBody) SetRequestId(v string) *CombineAuthorizationLetterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CombineAuthorizationLetterResponseBody) SetTemplateCombineUrl(v string) *CombineAuthorizationLetterResponseBody {
	s.TemplateCombineUrl = &v
	return s
}

type CombineAuthorizationLetterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CombineAuthorizationLetterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CombineAuthorizationLetterResponse) String() string {
	return tea.Prettify(s)
}

func (s CombineAuthorizationLetterResponse) GoString() string {
	return s.String()
}

func (s *CombineAuthorizationLetterResponse) SetHeaders(v map[string]*string) *CombineAuthorizationLetterResponse {
	s.Headers = v
	return s
}

func (s *CombineAuthorizationLetterResponse) SetStatusCode(v int32) *CombineAuthorizationLetterResponse {
	s.StatusCode = &v
	return s
}

func (s *CombineAuthorizationLetterResponse) SetBody(v *CombineAuthorizationLetterResponseBody) *CombineAuthorizationLetterResponse {
	s.Body = v
	return s
}

type ComplementTrademarkApplicationRequest struct {
	AgreementId         *string `json:"AgreementId,omitempty" xml:"AgreementId,omitempty"`
	AuthorizationOssKey *string `json:"AuthorizationOssKey,omitempty" xml:"AuthorizationOssKey,omitempty"`
	BizId               *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IsBlackIcon         *bool   `json:"IsBlackIcon,omitempty" xml:"IsBlackIcon,omitempty"`
	MaterialId          *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	OrderData           *string `json:"OrderData,omitempty" xml:"OrderData,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TrademarkComment    *string `json:"TrademarkComment,omitempty" xml:"TrademarkComment,omitempty"`
	TrademarkIconOssKey *string `json:"TrademarkIconOssKey,omitempty" xml:"TrademarkIconOssKey,omitempty"`
	TrademarkName       *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNameType   *string `json:"TrademarkNameType,omitempty" xml:"TrademarkNameType,omitempty"`
	TrademarkType       *int32  `json:"TrademarkType,omitempty" xml:"TrademarkType,omitempty"`
}

func (s ComplementTrademarkApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ComplementTrademarkApplicationRequest) GoString() string {
	return s.String()
}

func (s *ComplementTrademarkApplicationRequest) SetAgreementId(v string) *ComplementTrademarkApplicationRequest {
	s.AgreementId = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetAuthorizationOssKey(v string) *ComplementTrademarkApplicationRequest {
	s.AuthorizationOssKey = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetBizId(v string) *ComplementTrademarkApplicationRequest {
	s.BizId = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetIsBlackIcon(v bool) *ComplementTrademarkApplicationRequest {
	s.IsBlackIcon = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetMaterialId(v string) *ComplementTrademarkApplicationRequest {
	s.MaterialId = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetOrderData(v string) *ComplementTrademarkApplicationRequest {
	s.OrderData = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetSource(v string) *ComplementTrademarkApplicationRequest {
	s.Source = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetTrademarkComment(v string) *ComplementTrademarkApplicationRequest {
	s.TrademarkComment = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetTrademarkIconOssKey(v string) *ComplementTrademarkApplicationRequest {
	s.TrademarkIconOssKey = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetTrademarkName(v string) *ComplementTrademarkApplicationRequest {
	s.TrademarkName = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetTrademarkNameType(v string) *ComplementTrademarkApplicationRequest {
	s.TrademarkNameType = &v
	return s
}

func (s *ComplementTrademarkApplicationRequest) SetTrademarkType(v int32) *ComplementTrademarkApplicationRequest {
	s.TrademarkType = &v
	return s
}

type ComplementTrademarkApplicationResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ComplementTrademarkApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ComplementTrademarkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ComplementTrademarkApplicationResponseBody) SetErrorCode(v string) *ComplementTrademarkApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ComplementTrademarkApplicationResponseBody) SetErrorMessage(v string) *ComplementTrademarkApplicationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ComplementTrademarkApplicationResponseBody) SetRequestId(v string) *ComplementTrademarkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ComplementTrademarkApplicationResponseBody) SetSuccess(v bool) *ComplementTrademarkApplicationResponseBody {
	s.Success = &v
	return s
}

type ComplementTrademarkApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComplementTrademarkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ComplementTrademarkApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ComplementTrademarkApplicationResponse) GoString() string {
	return s.String()
}

func (s *ComplementTrademarkApplicationResponse) SetHeaders(v map[string]*string) *ComplementTrademarkApplicationResponse {
	s.Headers = v
	return s
}

func (s *ComplementTrademarkApplicationResponse) SetStatusCode(v int32) *ComplementTrademarkApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ComplementTrademarkApplicationResponse) SetBody(v *ComplementTrademarkApplicationResponseBody) *ComplementTrademarkApplicationResponse {
	s.Body = v
	return s
}

type ConfirmExpertSolutionRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note  *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s ConfirmExpertSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmExpertSolutionRequest) GoString() string {
	return s.String()
}

func (s *ConfirmExpertSolutionRequest) SetBizId(v string) *ConfirmExpertSolutionRequest {
	s.BizId = &v
	return s
}

func (s *ConfirmExpertSolutionRequest) SetNote(v string) *ConfirmExpertSolutionRequest {
	s.Note = &v
	return s
}

type ConfirmExpertSolutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmExpertSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmExpertSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmExpertSolutionResponseBody) SetRequestId(v string) *ConfirmExpertSolutionResponseBody {
	s.RequestId = &v
	return s
}

type ConfirmExpertSolutionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmExpertSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmExpertSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmExpertSolutionResponse) GoString() string {
	return s.String()
}

func (s *ConfirmExpertSolutionResponse) SetHeaders(v map[string]*string) *ConfirmExpertSolutionResponse {
	s.Headers = v
	return s
}

func (s *ConfirmExpertSolutionResponse) SetStatusCode(v int32) *ConfirmExpertSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmExpertSolutionResponse) SetBody(v *ConfirmExpertSolutionResponseBody) *ConfirmExpertSolutionResponse {
	s.Body = v
	return s
}

type CreateApplicantRequest struct {
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantName         *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantRegion       *int32  `json:"ApplicantRegion,omitempty" xml:"ApplicantRegion,omitempty"`
	ApplicantType         *int32  `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	AuthorizationOssKey   *string `json:"AuthorizationOssKey,omitempty" xml:"AuthorizationOssKey,omitempty"`
	BusinessLicenceOssKey *string `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ContactAddress        *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactCity           *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactCounty         *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
	ContactDistrict       *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactEmail          *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName           *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber         *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactProvince       *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactZipcode        *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Country               *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EAddress              *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	EName                 *string `json:"EName,omitempty" xml:"EName,omitempty"`
	IdCardName            *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber          *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	IdCardOssKey          *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	LegalNoticeOssKey     *string `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	PassportOssKey        *string `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	PersonalType          *int64  `json:"PersonalType,omitempty" xml:"PersonalType,omitempty"`
	PrincipalName         *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Province              *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s CreateApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicantRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicantRequest) SetAddress(v string) *CreateApplicantRequest {
	s.Address = &v
	return s
}

func (s *CreateApplicantRequest) SetApplicantName(v string) *CreateApplicantRequest {
	s.ApplicantName = &v
	return s
}

func (s *CreateApplicantRequest) SetApplicantRegion(v int32) *CreateApplicantRequest {
	s.ApplicantRegion = &v
	return s
}

func (s *CreateApplicantRequest) SetApplicantType(v int32) *CreateApplicantRequest {
	s.ApplicantType = &v
	return s
}

func (s *CreateApplicantRequest) SetAuthorizationOssKey(v string) *CreateApplicantRequest {
	s.AuthorizationOssKey = &v
	return s
}

func (s *CreateApplicantRequest) SetBusinessLicenceOssKey(v string) *CreateApplicantRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *CreateApplicantRequest) SetCardNumber(v string) *CreateApplicantRequest {
	s.CardNumber = &v
	return s
}

func (s *CreateApplicantRequest) SetContactAddress(v string) *CreateApplicantRequest {
	s.ContactAddress = &v
	return s
}

func (s *CreateApplicantRequest) SetContactCity(v string) *CreateApplicantRequest {
	s.ContactCity = &v
	return s
}

func (s *CreateApplicantRequest) SetContactCounty(v string) *CreateApplicantRequest {
	s.ContactCounty = &v
	return s
}

func (s *CreateApplicantRequest) SetContactDistrict(v string) *CreateApplicantRequest {
	s.ContactDistrict = &v
	return s
}

func (s *CreateApplicantRequest) SetContactEmail(v string) *CreateApplicantRequest {
	s.ContactEmail = &v
	return s
}

func (s *CreateApplicantRequest) SetContactName(v string) *CreateApplicantRequest {
	s.ContactName = &v
	return s
}

func (s *CreateApplicantRequest) SetContactNumber(v string) *CreateApplicantRequest {
	s.ContactNumber = &v
	return s
}

func (s *CreateApplicantRequest) SetContactProvince(v string) *CreateApplicantRequest {
	s.ContactProvince = &v
	return s
}

func (s *CreateApplicantRequest) SetContactZipcode(v string) *CreateApplicantRequest {
	s.ContactZipcode = &v
	return s
}

func (s *CreateApplicantRequest) SetCountry(v string) *CreateApplicantRequest {
	s.Country = &v
	return s
}

func (s *CreateApplicantRequest) SetEAddress(v string) *CreateApplicantRequest {
	s.EAddress = &v
	return s
}

func (s *CreateApplicantRequest) SetEName(v string) *CreateApplicantRequest {
	s.EName = &v
	return s
}

func (s *CreateApplicantRequest) SetIdCardName(v string) *CreateApplicantRequest {
	s.IdCardName = &v
	return s
}

func (s *CreateApplicantRequest) SetIdCardNumber(v string) *CreateApplicantRequest {
	s.IdCardNumber = &v
	return s
}

func (s *CreateApplicantRequest) SetIdCardOssKey(v string) *CreateApplicantRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *CreateApplicantRequest) SetLegalNoticeOssKey(v string) *CreateApplicantRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *CreateApplicantRequest) SetPassportOssKey(v string) *CreateApplicantRequest {
	s.PassportOssKey = &v
	return s
}

func (s *CreateApplicantRequest) SetPersonalType(v int64) *CreateApplicantRequest {
	s.PersonalType = &v
	return s
}

func (s *CreateApplicantRequest) SetPrincipalName(v int32) *CreateApplicantRequest {
	s.PrincipalName = &v
	return s
}

func (s *CreateApplicantRequest) SetProvince(v string) *CreateApplicantRequest {
	s.Province = &v
	return s
}

type CreateApplicantResponseBody struct {
	ApplicantId *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicantResponseBody) SetApplicantId(v string) *CreateApplicantResponseBody {
	s.ApplicantId = &v
	return s
}

func (s *CreateApplicantResponseBody) SetRequestId(v string) *CreateApplicantResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicantResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicantResponse) SetHeaders(v map[string]*string) *CreateApplicantResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicantResponse) SetStatusCode(v int32) *CreateApplicantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicantResponse) SetBody(v *CreateApplicantResponseBody) *CreateApplicantResponse {
	s.Body = v
	return s
}

type CreateCommodityOrderRequest struct {
	AutoPay       *bool                  `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BizType       *string                `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ChargeType    *string                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken   *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityCode *string                `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Components    map[string]interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	Duration      *int32                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId    *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderParams   map[string]interface{} `json:"OrderParams,omitempty" xml:"OrderParams,omitempty"`
	OrderType     *string                `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PricingCycle  *string                `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Quantity      *int32                 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SpecCode      *string                `json:"SpecCode,omitempty" xml:"SpecCode,omitempty"`
	UserId        *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateCommodityOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCommodityOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateCommodityOrderRequest) SetAutoPay(v bool) *CreateCommodityOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetBizType(v string) *CreateCommodityOrderRequest {
	s.BizType = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetChargeType(v string) *CreateCommodityOrderRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetClientToken(v string) *CreateCommodityOrderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetCommodityCode(v string) *CreateCommodityOrderRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetComponents(v map[string]interface{}) *CreateCommodityOrderRequest {
	s.Components = v
	return s
}

func (s *CreateCommodityOrderRequest) SetDuration(v int32) *CreateCommodityOrderRequest {
	s.Duration = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetInstanceId(v string) *CreateCommodityOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetOrderParams(v map[string]interface{}) *CreateCommodityOrderRequest {
	s.OrderParams = v
	return s
}

func (s *CreateCommodityOrderRequest) SetOrderType(v string) *CreateCommodityOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetPricingCycle(v string) *CreateCommodityOrderRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetQuantity(v int32) *CreateCommodityOrderRequest {
	s.Quantity = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetSpecCode(v string) *CreateCommodityOrderRequest {
	s.SpecCode = &v
	return s
}

func (s *CreateCommodityOrderRequest) SetUserId(v string) *CreateCommodityOrderRequest {
	s.UserId = &v
	return s
}

type CreateCommodityOrderShrinkRequest struct {
	AutoPay           *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BizType           *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ChargeType        *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityCode     *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ComponentsShrink  *string `json:"Components,omitempty" xml:"Components,omitempty"`
	Duration          *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderParamsShrink *string `json:"OrderParams,omitempty" xml:"OrderParams,omitempty"`
	OrderType         *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PricingCycle      *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Quantity          *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SpecCode          *string `json:"SpecCode,omitempty" xml:"SpecCode,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateCommodityOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCommodityOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCommodityOrderShrinkRequest) SetAutoPay(v bool) *CreateCommodityOrderShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetBizType(v string) *CreateCommodityOrderShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetChargeType(v string) *CreateCommodityOrderShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetClientToken(v string) *CreateCommodityOrderShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetCommodityCode(v string) *CreateCommodityOrderShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetComponentsShrink(v string) *CreateCommodityOrderShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetDuration(v int32) *CreateCommodityOrderShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetInstanceId(v string) *CreateCommodityOrderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetOrderParamsShrink(v string) *CreateCommodityOrderShrinkRequest {
	s.OrderParamsShrink = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetOrderType(v string) *CreateCommodityOrderShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetPricingCycle(v string) *CreateCommodityOrderShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetQuantity(v int32) *CreateCommodityOrderShrinkRequest {
	s.Quantity = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetSpecCode(v string) *CreateCommodityOrderShrinkRequest {
	s.SpecCode = &v
	return s
}

func (s *CreateCommodityOrderShrinkRequest) SetUserId(v string) *CreateCommodityOrderShrinkRequest {
	s.UserId = &v
	return s
}

type CreateCommodityOrderResponseBody struct {
	Data      []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCommodityOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCommodityOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommodityOrderResponseBody) SetData(v []*int64) *CreateCommodityOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateCommodityOrderResponseBody) SetRequestId(v string) *CreateCommodityOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommodityOrderResponseBody) SetSuccess(v bool) *CreateCommodityOrderResponseBody {
	s.Success = &v
	return s
}

type CreateCommodityOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommodityOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommodityOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCommodityOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateCommodityOrderResponse) SetHeaders(v map[string]*string) *CreateCommodityOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateCommodityOrderResponse) SetStatusCode(v int32) *CreateCommodityOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommodityOrderResponse) SetBody(v *CreateCommodityOrderResponseBody) *CreateCommodityOrderResponse {
	s.Body = v
	return s
}

type CreateOrderRequest struct {
	AgreementId         *string `json:"AgreementId,omitempty" xml:"AgreementId,omitempty"`
	ApplicantId         *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicationType     *int32  `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	AuthorizationOssKey *string `json:"AuthorizationOssKey,omitempty" xml:"AuthorizationOssKey,omitempty"`
	AutoPay             *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BlackAndWhiteIcon   *string `json:"BlackAndWhiteIcon,omitempty" xml:"BlackAndWhiteIcon,omitempty"`
	Channel             *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Classifications     *string `json:"Classifications,omitempty" xml:"Classifications,omitempty"`
	IdempotentId        *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	LegalNoticeKey      *string `json:"LegalNoticeKey,omitempty" xml:"LegalNoticeKey,omitempty"`
	PayType             *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PaymentCallback     *string `json:"PaymentCallback,omitempty" xml:"PaymentCallback,omitempty"`
	PrincipalName       *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TrademarkComment    *string `json:"TrademarkComment,omitempty" xml:"TrademarkComment,omitempty"`
	TrademarkIcon       *string `json:"TrademarkIcon,omitempty" xml:"TrademarkIcon,omitempty"`
	TrademarkName       *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNameType   *string `json:"TrademarkNameType,omitempty" xml:"TrademarkNameType,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetAgreementId(v string) *CreateOrderRequest {
	s.AgreementId = &v
	return s
}

func (s *CreateOrderRequest) SetApplicantId(v string) *CreateOrderRequest {
	s.ApplicantId = &v
	return s
}

func (s *CreateOrderRequest) SetApplicationType(v int32) *CreateOrderRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateOrderRequest) SetAuthorizationOssKey(v string) *CreateOrderRequest {
	s.AuthorizationOssKey = &v
	return s
}

func (s *CreateOrderRequest) SetAutoPay(v bool) *CreateOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderRequest) SetBlackAndWhiteIcon(v string) *CreateOrderRequest {
	s.BlackAndWhiteIcon = &v
	return s
}

func (s *CreateOrderRequest) SetChannel(v string) *CreateOrderRequest {
	s.Channel = &v
	return s
}

func (s *CreateOrderRequest) SetClassifications(v string) *CreateOrderRequest {
	s.Classifications = &v
	return s
}

func (s *CreateOrderRequest) SetIdempotentId(v string) *CreateOrderRequest {
	s.IdempotentId = &v
	return s
}

func (s *CreateOrderRequest) SetLegalNoticeKey(v string) *CreateOrderRequest {
	s.LegalNoticeKey = &v
	return s
}

func (s *CreateOrderRequest) SetPayType(v string) *CreateOrderRequest {
	s.PayType = &v
	return s
}

func (s *CreateOrderRequest) SetPaymentCallback(v string) *CreateOrderRequest {
	s.PaymentCallback = &v
	return s
}

func (s *CreateOrderRequest) SetPrincipalName(v int32) *CreateOrderRequest {
	s.PrincipalName = &v
	return s
}

func (s *CreateOrderRequest) SetSource(v string) *CreateOrderRequest {
	s.Source = &v
	return s
}

func (s *CreateOrderRequest) SetTrademarkComment(v string) *CreateOrderRequest {
	s.TrademarkComment = &v
	return s
}

func (s *CreateOrderRequest) SetTrademarkIcon(v string) *CreateOrderRequest {
	s.TrademarkIcon = &v
	return s
}

func (s *CreateOrderRequest) SetTrademarkName(v string) *CreateOrderRequest {
	s.TrademarkName = &v
	return s
}

func (s *CreateOrderRequest) SetTrademarkNameType(v string) *CreateOrderRequest {
	s.TrademarkNameType = &v
	return s
}

type CreateOrderResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderIds  map[string]interface{} `json:"OrderIds,omitempty" xml:"OrderIds,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetMessage(v string) *CreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrderResponseBody) SetOrderIds(v map[string]interface{}) *CreateOrderResponseBody {
	s.OrderIds = v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetSuccess(v bool) *CreateOrderResponseBody {
	s.Success = &v
	return s
}

type CreateOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetStatusCode(v int32) *CreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type CreateTrademarkApplicationRequest struct {
	AgreementId         *string `json:"AgreementId,omitempty" xml:"AgreementId,omitempty"`
	ApplicantId         *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicationType     *int32  `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	AuthorizationOssKey *string `json:"AuthorizationOssKey,omitempty" xml:"AuthorizationOssKey,omitempty"`
	AutoPay             *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BlackAndWhiteIcon   *string `json:"BlackAndWhiteIcon,omitempty" xml:"BlackAndWhiteIcon,omitempty"`
	Channel             *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Classifications     *string `json:"Classifications,omitempty" xml:"Classifications,omitempty"`
	IdempotentId        *string `json:"IdempotentId,omitempty" xml:"IdempotentId,omitempty"`
	LegalNoticeKey      *string `json:"LegalNoticeKey,omitempty" xml:"LegalNoticeKey,omitempty"`
	PrincipalName       *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TrademarkComment    *string `json:"TrademarkComment,omitempty" xml:"TrademarkComment,omitempty"`
	TrademarkIcon       *string `json:"TrademarkIcon,omitempty" xml:"TrademarkIcon,omitempty"`
	TrademarkName       *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNameType   *string `json:"TrademarkNameType,omitempty" xml:"TrademarkNameType,omitempty"`
}

func (s CreateTrademarkApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrademarkApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateTrademarkApplicationRequest) SetAgreementId(v string) *CreateTrademarkApplicationRequest {
	s.AgreementId = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetApplicantId(v string) *CreateTrademarkApplicationRequest {
	s.ApplicantId = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetApplicationType(v int32) *CreateTrademarkApplicationRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetAuthorizationOssKey(v string) *CreateTrademarkApplicationRequest {
	s.AuthorizationOssKey = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetAutoPay(v bool) *CreateTrademarkApplicationRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetBlackAndWhiteIcon(v string) *CreateTrademarkApplicationRequest {
	s.BlackAndWhiteIcon = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetChannel(v string) *CreateTrademarkApplicationRequest {
	s.Channel = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetClassifications(v string) *CreateTrademarkApplicationRequest {
	s.Classifications = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetIdempotentId(v string) *CreateTrademarkApplicationRequest {
	s.IdempotentId = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetLegalNoticeKey(v string) *CreateTrademarkApplicationRequest {
	s.LegalNoticeKey = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetPrincipalName(v int32) *CreateTrademarkApplicationRequest {
	s.PrincipalName = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetSource(v string) *CreateTrademarkApplicationRequest {
	s.Source = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetTrademarkComment(v string) *CreateTrademarkApplicationRequest {
	s.TrademarkComment = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetTrademarkIcon(v string) *CreateTrademarkApplicationRequest {
	s.TrademarkIcon = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetTrademarkName(v string) *CreateTrademarkApplicationRequest {
	s.TrademarkName = &v
	return s
}

func (s *CreateTrademarkApplicationRequest) SetTrademarkNameType(v string) *CreateTrademarkApplicationRequest {
	s.TrademarkNameType = &v
	return s
}

type CreateTrademarkApplicationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTrademarkApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrademarkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrademarkApplicationResponseBody) SetMessage(v string) *CreateTrademarkApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTrademarkApplicationResponseBody) SetOrderId(v int64) *CreateTrademarkApplicationResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateTrademarkApplicationResponseBody) SetRequestId(v string) *CreateTrademarkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrademarkApplicationResponseBody) SetSuccess(v bool) *CreateTrademarkApplicationResponseBody {
	s.Success = &v
	return s
}

type CreateTrademarkApplicationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrademarkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrademarkApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrademarkApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateTrademarkApplicationResponse) SetHeaders(v map[string]*string) *CreateTrademarkApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateTrademarkApplicationResponse) SetStatusCode(v int32) *CreateTrademarkApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrademarkApplicationResponse) SetBody(v *CreateTrademarkApplicationResponseBody) *CreateTrademarkApplicationResponse {
	s.Body = v
	return s
}

type DeleteSearchConditionRequest struct {
	ConditionId *int64  `json:"ConditionId,omitempty" xml:"ConditionId,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Umid        *string `json:"Umid,omitempty" xml:"Umid,omitempty"`
}

func (s DeleteSearchConditionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSearchConditionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSearchConditionRequest) SetConditionId(v int64) *DeleteSearchConditionRequest {
	s.ConditionId = &v
	return s
}

func (s *DeleteSearchConditionRequest) SetSessionId(v string) *DeleteSearchConditionRequest {
	s.SessionId = &v
	return s
}

func (s *DeleteSearchConditionRequest) SetUmid(v string) *DeleteSearchConditionRequest {
	s.Umid = &v
	return s
}

type DeleteSearchConditionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSearchConditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSearchConditionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSearchConditionResponseBody) SetCode(v string) *DeleteSearchConditionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSearchConditionResponseBody) SetMessage(v string) *DeleteSearchConditionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSearchConditionResponseBody) SetRequestId(v string) *DeleteSearchConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSearchConditionResponseBody) SetSuccess(v bool) *DeleteSearchConditionResponseBody {
	s.Success = &v
	return s
}

type DeleteSearchConditionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSearchConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSearchConditionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSearchConditionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSearchConditionResponse) SetHeaders(v map[string]*string) *DeleteSearchConditionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSearchConditionResponse) SetStatusCode(v int32) *DeleteSearchConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSearchConditionResponse) SetBody(v *DeleteSearchConditionResponseBody) *DeleteSearchConditionResponse {
	s.Body = v
	return s
}

type DescribeAdminTrademarkApplicationRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeAdminTrademarkApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdminTrademarkApplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdminTrademarkApplicationRequest) SetBizId(v string) *DescribeAdminTrademarkApplicationRequest {
	s.BizId = &v
	return s
}

type DescribeAdminTrademarkApplicationResponseBody struct {
	AcceptUrl            *string                                                              `json:"AcceptUrl,omitempty" xml:"AcceptUrl,omitempty"`
	Applicant            *DescribeAdminTrademarkApplicationResponseBodyApplicant              `json:"Applicant,omitempty" xml:"Applicant,omitempty" type:"Struct"`
	ApplicantId          *int64                                                               `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicationStatus    *int32                                                               `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	ApplicationType      *int32                                                               `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	AuthorizationUrl     *string                                                              `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BizId                *string                                                              `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BlackAndWhiteIconUrl *string                                                              `json:"BlackAndWhiteIconUrl,omitempty" xml:"BlackAndWhiteIconUrl,omitempty"`
	CreateTime           *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtendInfo           map[string]interface{}                                               `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FirstClassification  *DescribeAdminTrademarkApplicationResponseBodyFirstClassification    `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
	JudgeResultUrls      []*string                                                            `json:"JudgeResultUrls,omitempty" xml:"JudgeResultUrls,omitempty" type:"Repeated"`
	Note                 *string                                                              `json:"Note,omitempty" xml:"Note,omitempty"`
	OrderId              *string                                                              `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderPrice           *float32                                                             `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	PrincipalName        *int32                                                               `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	ReceiptUrls          []*string                                                            `json:"ReceiptUrls,omitempty" xml:"ReceiptUrls,omitempty" type:"Repeated"`
	RecvUserLogistics    *string                                                              `json:"RecvUserLogistics,omitempty" xml:"RecvUserLogistics,omitempty"`
	RequestId            *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SendSbjLogistics     *string                                                              `json:"SendSbjLogistics,omitempty" xml:"SendSbjLogistics,omitempty"`
	SendUserLogistics    *string                                                              `json:"SendUserLogistics,omitempty" xml:"SendUserLogistics,omitempty"`
	ServicePrice         *float32                                                             `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	Supplements          []*DescribeAdminTrademarkApplicationResponseBodySupplements          `json:"Supplements,omitempty" xml:"Supplements,omitempty" type:"Repeated"`
	ThirdClassifications []*DescribeAdminTrademarkApplicationResponseBodyThirdClassifications `json:"ThirdClassifications,omitempty" xml:"ThirdClassifications,omitempty" type:"Repeated"`
	TotalPrice           *float32                                                             `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	TrademarkIcon        *string                                                              `json:"TrademarkIcon,omitempty" xml:"TrademarkIcon,omitempty"`
	TrademarkName        *string                                                              `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNameType    *int32                                                               `json:"TrademarkNameType,omitempty" xml:"TrademarkNameType,omitempty"`
	TrademarkNumber      *string                                                              `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UpdateTime           *int64                                                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId               *string                                                              `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeAdminTrademarkApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdminTrademarkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetAcceptUrl(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.AcceptUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetApplicant(v *DescribeAdminTrademarkApplicationResponseBodyApplicant) *DescribeAdminTrademarkApplicationResponseBody {
	s.Applicant = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetApplicantId(v int64) *DescribeAdminTrademarkApplicationResponseBody {
	s.ApplicantId = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetApplicationStatus(v int32) *DescribeAdminTrademarkApplicationResponseBody {
	s.ApplicationStatus = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetApplicationType(v int32) *DescribeAdminTrademarkApplicationResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetAuthorizationUrl(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.AuthorizationUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetBizId(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.BizId = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetBlackAndWhiteIconUrl(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.BlackAndWhiteIconUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetCreateTime(v int64) *DescribeAdminTrademarkApplicationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetExtendInfo(v map[string]interface{}) *DescribeAdminTrademarkApplicationResponseBody {
	s.ExtendInfo = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetFirstClassification(v *DescribeAdminTrademarkApplicationResponseBodyFirstClassification) *DescribeAdminTrademarkApplicationResponseBody {
	s.FirstClassification = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetJudgeResultUrls(v []*string) *DescribeAdminTrademarkApplicationResponseBody {
	s.JudgeResultUrls = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetNote(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.Note = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetOrderId(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetOrderPrice(v float32) *DescribeAdminTrademarkApplicationResponseBody {
	s.OrderPrice = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetPrincipalName(v int32) *DescribeAdminTrademarkApplicationResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetReceiptUrls(v []*string) *DescribeAdminTrademarkApplicationResponseBody {
	s.ReceiptUrls = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetRecvUserLogistics(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.RecvUserLogistics = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetRequestId(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetSendSbjLogistics(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.SendSbjLogistics = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetSendUserLogistics(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.SendUserLogistics = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetServicePrice(v float32) *DescribeAdminTrademarkApplicationResponseBody {
	s.ServicePrice = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetSupplements(v []*DescribeAdminTrademarkApplicationResponseBodySupplements) *DescribeAdminTrademarkApplicationResponseBody {
	s.Supplements = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetThirdClassifications(v []*DescribeAdminTrademarkApplicationResponseBodyThirdClassifications) *DescribeAdminTrademarkApplicationResponseBody {
	s.ThirdClassifications = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetTotalPrice(v float32) *DescribeAdminTrademarkApplicationResponseBody {
	s.TotalPrice = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetTrademarkIcon(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.TrademarkIcon = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetTrademarkName(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.TrademarkName = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetTrademarkNameType(v int32) *DescribeAdminTrademarkApplicationResponseBody {
	s.TrademarkNameType = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetTrademarkNumber(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.TrademarkNumber = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetUpdateTime(v int64) *DescribeAdminTrademarkApplicationResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBody) SetUserId(v string) *DescribeAdminTrademarkApplicationResponseBody {
	s.UserId = &v
	return s
}

type DescribeAdminTrademarkApplicationResponseBodyApplicant struct {
	Address            *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantName      *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantRegion    *int32  `json:"ApplicantRegion,omitempty" xml:"ApplicantRegion,omitempty"`
	ApplicantType      *int32  `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	AuditStatus        *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuthorizationUrl   *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BusinessLicenceUrl *string `json:"BusinessLicenceUrl,omitempty" xml:"BusinessLicenceUrl,omitempty"`
	CardNumber         *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ContactAddress     *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactEmail       *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName        *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber      *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactZipcode     *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Country            *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EAddress           *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	EName              *string `json:"EName,omitempty" xml:"EName,omitempty"`
	IdCardUrl          *string `json:"IdCardUrl,omitempty" xml:"IdCardUrl,omitempty"`
	LegalNoticeUrl     *string `json:"LegalNoticeUrl,omitempty" xml:"LegalNoticeUrl,omitempty"`
	PassportUrl        *string `json:"PassportUrl,omitempty" xml:"PassportUrl,omitempty"`
	PrincipalName      *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Province           *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeAdminTrademarkApplicationResponseBodyApplicant) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdminTrademarkApplicationResponseBodyApplicant) GoString() string {
	return s.String()
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetAddress(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.Address = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetApplicantName(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ApplicantName = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetApplicantRegion(v int32) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ApplicantRegion = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetApplicantType(v int32) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ApplicantType = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetAuditStatus(v int32) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.AuditStatus = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetAuthorizationUrl(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.AuthorizationUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetBusinessLicenceUrl(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.BusinessLicenceUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetCardNumber(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.CardNumber = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetContactAddress(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ContactAddress = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetContactEmail(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ContactEmail = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetContactName(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ContactName = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetContactNumber(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ContactNumber = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetContactZipcode(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.ContactZipcode = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetCountry(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.Country = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetEAddress(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.EAddress = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetEName(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.EName = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetIdCardUrl(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.IdCardUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetLegalNoticeUrl(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.LegalNoticeUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetPassportUrl(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.PassportUrl = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetPrincipalName(v int32) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.PrincipalName = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyApplicant) SetProvince(v string) *DescribeAdminTrademarkApplicationResponseBodyApplicant {
	s.Province = &v
	return s
}

type DescribeAdminTrademarkApplicationResponseBodyFirstClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s DescribeAdminTrademarkApplicationResponseBodyFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdminTrademarkApplicationResponseBodyFirstClassification) GoString() string {
	return s.String()
}

func (s *DescribeAdminTrademarkApplicationResponseBodyFirstClassification) SetClassificationCode(v string) *DescribeAdminTrademarkApplicationResponseBodyFirstClassification {
	s.ClassificationCode = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyFirstClassification) SetClassificationName(v string) *DescribeAdminTrademarkApplicationResponseBodyFirstClassification {
	s.ClassificationName = &v
	return s
}

type DescribeAdminTrademarkApplicationResponseBodySupplements struct {
	AcceptExpirationDate *int64    `json:"AcceptExpirationDate,omitempty" xml:"AcceptExpirationDate,omitempty"`
	AcceptTime           *int64    `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	ApplicationType      *int32    `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Content              *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	OfficialFile         *string   `json:"OfficialFile,omitempty" xml:"OfficialFile,omitempty"`
	OperateTime          *int64    `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OrderId              *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	SbjExpirationDate    *int64    `json:"SbjExpirationDate,omitempty" xml:"SbjExpirationDate,omitempty"`
	SendTime             *int64    `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	SerialNumber         *string   `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SupplementId         *int64    `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	SupplementStatus     *int32    `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TrademarkNumber      *string   `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UserFiles            []*string `json:"UserFiles,omitempty" xml:"UserFiles,omitempty" type:"Repeated"`
}

func (s DescribeAdminTrademarkApplicationResponseBodySupplements) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdminTrademarkApplicationResponseBodySupplements) GoString() string {
	return s.String()
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetAcceptExpirationDate(v int64) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.AcceptExpirationDate = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetAcceptTime(v int64) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.AcceptTime = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetApplicationType(v int32) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.ApplicationType = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetContent(v string) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.Content = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetOfficialFile(v string) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.OfficialFile = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetOperateTime(v int64) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.OperateTime = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetOrderId(v string) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.OrderId = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetSbjExpirationDate(v int64) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.SbjExpirationDate = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetSendTime(v int64) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.SendTime = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetSerialNumber(v string) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.SerialNumber = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetSupplementId(v int64) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.SupplementId = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetSupplementStatus(v int32) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.SupplementStatus = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetTrademarkNumber(v string) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.TrademarkNumber = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodySupplements) SetUserFiles(v []*string) *DescribeAdminTrademarkApplicationResponseBodySupplements {
	s.UserFiles = v
	return s
}

type DescribeAdminTrademarkApplicationResponseBodyThirdClassifications struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s DescribeAdminTrademarkApplicationResponseBodyThirdClassifications) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdminTrademarkApplicationResponseBodyThirdClassifications) GoString() string {
	return s.String()
}

func (s *DescribeAdminTrademarkApplicationResponseBodyThirdClassifications) SetClassificationCode(v string) *DescribeAdminTrademarkApplicationResponseBodyThirdClassifications {
	s.ClassificationCode = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponseBodyThirdClassifications) SetClassificationName(v string) *DescribeAdminTrademarkApplicationResponseBodyThirdClassifications {
	s.ClassificationName = &v
	return s
}

type DescribeAdminTrademarkApplicationResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdminTrademarkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdminTrademarkApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAdminTrademarkApplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdminTrademarkApplicationResponse) SetHeaders(v map[string]*string) *DescribeAdminTrademarkApplicationResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponse) SetStatusCode(v int32) *DescribeAdminTrademarkApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdminTrademarkApplicationResponse) SetBody(v *DescribeAdminTrademarkApplicationResponseBody) *DescribeAdminTrademarkApplicationResponse {
	s.Body = v
	return s
}

type DescribeApplicantRequest struct {
	ApplicantId *int64 `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
}

func (s DescribeApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicantRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicantRequest) SetApplicantId(v int64) *DescribeApplicantRequest {
	s.ApplicantId = &v
	return s
}

type DescribeApplicantResponseBody struct {
	Address                  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantId              *int64  `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicantName            *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantRegion          *int32  `json:"ApplicantRegion,omitempty" xml:"ApplicantRegion,omitempty"`
	ApplicantType            *int32  `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	ApplicantVersion         *string `json:"ApplicantVersion,omitempty" xml:"ApplicantVersion,omitempty"`
	AuditStatus              *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuthorizationAuditStatus *int32  `json:"AuthorizationAuditStatus,omitempty" xml:"AuthorizationAuditStatus,omitempty"`
	AuthorizationUrl         *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BusinessLicenceUrl       *string `json:"BusinessLicenceUrl,omitempty" xml:"BusinessLicenceUrl,omitempty"`
	CardNumber               *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ContactAddress           *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactCity              *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactCounty            *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
	ContactDistrict          *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactEmail             *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName              *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber            *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactProvince          *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactZipcode           *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Country                  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EAddress                 *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	EName                    *string `json:"EName,omitempty" xml:"EName,omitempty"`
	IdCardName               *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber             *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	IdCardUrl                *string `json:"IdCardUrl,omitempty" xml:"IdCardUrl,omitempty"`
	LegalNoticeUrl           *string `json:"LegalNoticeUrl,omitempty" xml:"LegalNoticeUrl,omitempty"`
	Note                     *string `json:"Note,omitempty" xml:"Note,omitempty"`
	PassportUrl              *string `json:"PassportUrl,omitempty" xml:"PassportUrl,omitempty"`
	PersonalType             *int64  `json:"PersonalType,omitempty" xml:"PersonalType,omitempty"`
	PrincipalName            *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Province                 *string `json:"Province,omitempty" xml:"Province,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ValidDate                *int64  `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
}

func (s DescribeApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicantResponseBody) SetAddress(v string) *DescribeApplicantResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetApplicantId(v int64) *DescribeApplicantResponseBody {
	s.ApplicantId = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetApplicantName(v string) *DescribeApplicantResponseBody {
	s.ApplicantName = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetApplicantRegion(v int32) *DescribeApplicantResponseBody {
	s.ApplicantRegion = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetApplicantType(v int32) *DescribeApplicantResponseBody {
	s.ApplicantType = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetApplicantVersion(v string) *DescribeApplicantResponseBody {
	s.ApplicantVersion = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetAuditStatus(v int32) *DescribeApplicantResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetAuthorizationAuditStatus(v int32) *DescribeApplicantResponseBody {
	s.AuthorizationAuditStatus = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetAuthorizationUrl(v string) *DescribeApplicantResponseBody {
	s.AuthorizationUrl = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetBusinessLicenceUrl(v string) *DescribeApplicantResponseBody {
	s.BusinessLicenceUrl = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetCardNumber(v string) *DescribeApplicantResponseBody {
	s.CardNumber = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactAddress(v string) *DescribeApplicantResponseBody {
	s.ContactAddress = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactCity(v string) *DescribeApplicantResponseBody {
	s.ContactCity = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactCounty(v string) *DescribeApplicantResponseBody {
	s.ContactCounty = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactDistrict(v string) *DescribeApplicantResponseBody {
	s.ContactDistrict = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactEmail(v string) *DescribeApplicantResponseBody {
	s.ContactEmail = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactName(v string) *DescribeApplicantResponseBody {
	s.ContactName = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactNumber(v string) *DescribeApplicantResponseBody {
	s.ContactNumber = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactProvince(v string) *DescribeApplicantResponseBody {
	s.ContactProvince = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetContactZipcode(v string) *DescribeApplicantResponseBody {
	s.ContactZipcode = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetCountry(v string) *DescribeApplicantResponseBody {
	s.Country = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetEAddress(v string) *DescribeApplicantResponseBody {
	s.EAddress = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetEName(v string) *DescribeApplicantResponseBody {
	s.EName = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetIdCardName(v string) *DescribeApplicantResponseBody {
	s.IdCardName = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetIdCardNumber(v string) *DescribeApplicantResponseBody {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetIdCardUrl(v string) *DescribeApplicantResponseBody {
	s.IdCardUrl = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetLegalNoticeUrl(v string) *DescribeApplicantResponseBody {
	s.LegalNoticeUrl = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetNote(v string) *DescribeApplicantResponseBody {
	s.Note = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetPassportUrl(v string) *DescribeApplicantResponseBody {
	s.PassportUrl = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetPersonalType(v int64) *DescribeApplicantResponseBody {
	s.PersonalType = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetPrincipalName(v int32) *DescribeApplicantResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetProvince(v string) *DescribeApplicantResponseBody {
	s.Province = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetRequestId(v string) *DescribeApplicantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicantResponseBody) SetValidDate(v int64) *DescribeApplicantResponseBody {
	s.ValidDate = &v
	return s
}

type DescribeApplicantResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicantResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicantResponse) SetHeaders(v map[string]*string) *DescribeApplicantResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicantResponse) SetStatusCode(v int32) *DescribeApplicantResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicantResponse) SetBody(v *DescribeApplicantResponseBody) *DescribeApplicantResponse {
	s.Body = v
	return s
}

type DescribePartnerTrademarkApplicationRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribePartnerTrademarkApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerTrademarkApplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribePartnerTrademarkApplicationRequest) SetBizId(v string) *DescribePartnerTrademarkApplicationRequest {
	s.BizId = &v
	return s
}

type DescribePartnerTrademarkApplicationResponseBody struct {
	AcceptUrl            *string                                                                `json:"AcceptUrl,omitempty" xml:"AcceptUrl,omitempty"`
	Applicant            *DescribePartnerTrademarkApplicationResponseBodyApplicant              `json:"Applicant,omitempty" xml:"Applicant,omitempty" type:"Struct"`
	ApplicantId          *int64                                                                 `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicationStatus    *int32                                                                 `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	ApplicationType      *int32                                                                 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	AuthorizationUrl     *string                                                                `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BizId                *string                                                                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BlackAndWhiteIconUrl *string                                                                `json:"BlackAndWhiteIconUrl,omitempty" xml:"BlackAndWhiteIconUrl,omitempty"`
	CreateTime           *int64                                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtendInfo           map[string]interface{}                                                 `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FirstClassification  *DescribePartnerTrademarkApplicationResponseBodyFirstClassification    `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
	JudgeResultUrls      []*string                                                              `json:"JudgeResultUrls,omitempty" xml:"JudgeResultUrls,omitempty" type:"Repeated"`
	Note                 *string                                                                `json:"Note,omitempty" xml:"Note,omitempty"`
	OrderId              *string                                                                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderPrice           *float32                                                               `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	PrincipalName        *int32                                                                 `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	ReceiptUrls          []*string                                                              `json:"ReceiptUrls,omitempty" xml:"ReceiptUrls,omitempty" type:"Repeated"`
	RecvUserLogistics    *string                                                                `json:"RecvUserLogistics,omitempty" xml:"RecvUserLogistics,omitempty"`
	RequestId            *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SendSbjLogistics     *string                                                                `json:"SendSbjLogistics,omitempty" xml:"SendSbjLogistics,omitempty"`
	SendUserLogistics    *string                                                                `json:"SendUserLogistics,omitempty" xml:"SendUserLogistics,omitempty"`
	ServicePrice         *float32                                                               `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	Supplements          []*DescribePartnerTrademarkApplicationResponseBodySupplements          `json:"Supplements,omitempty" xml:"Supplements,omitempty" type:"Repeated"`
	ThirdClassifications []*DescribePartnerTrademarkApplicationResponseBodyThirdClassifications `json:"ThirdClassifications,omitempty" xml:"ThirdClassifications,omitempty" type:"Repeated"`
	TotalPrice           *float32                                                               `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	TrademarkIcon        *string                                                                `json:"TrademarkIcon,omitempty" xml:"TrademarkIcon,omitempty"`
	TrademarkName        *string                                                                `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNameType    *int32                                                                 `json:"TrademarkNameType,omitempty" xml:"TrademarkNameType,omitempty"`
	TrademarkNumber      *string                                                                `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UpdateTime           *int64                                                                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribePartnerTrademarkApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerTrademarkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetAcceptUrl(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.AcceptUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetApplicant(v *DescribePartnerTrademarkApplicationResponseBodyApplicant) *DescribePartnerTrademarkApplicationResponseBody {
	s.Applicant = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetApplicantId(v int64) *DescribePartnerTrademarkApplicationResponseBody {
	s.ApplicantId = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetApplicationStatus(v int32) *DescribePartnerTrademarkApplicationResponseBody {
	s.ApplicationStatus = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetApplicationType(v int32) *DescribePartnerTrademarkApplicationResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetAuthorizationUrl(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.AuthorizationUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetBizId(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.BizId = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetBlackAndWhiteIconUrl(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.BlackAndWhiteIconUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetCreateTime(v int64) *DescribePartnerTrademarkApplicationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetExtendInfo(v map[string]interface{}) *DescribePartnerTrademarkApplicationResponseBody {
	s.ExtendInfo = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetFirstClassification(v *DescribePartnerTrademarkApplicationResponseBodyFirstClassification) *DescribePartnerTrademarkApplicationResponseBody {
	s.FirstClassification = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetJudgeResultUrls(v []*string) *DescribePartnerTrademarkApplicationResponseBody {
	s.JudgeResultUrls = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetNote(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.Note = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetOrderId(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetOrderPrice(v float32) *DescribePartnerTrademarkApplicationResponseBody {
	s.OrderPrice = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetPrincipalName(v int32) *DescribePartnerTrademarkApplicationResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetReceiptUrls(v []*string) *DescribePartnerTrademarkApplicationResponseBody {
	s.ReceiptUrls = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetRecvUserLogistics(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.RecvUserLogistics = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetRequestId(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetSendSbjLogistics(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.SendSbjLogistics = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetSendUserLogistics(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.SendUserLogistics = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetServicePrice(v float32) *DescribePartnerTrademarkApplicationResponseBody {
	s.ServicePrice = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetSupplements(v []*DescribePartnerTrademarkApplicationResponseBodySupplements) *DescribePartnerTrademarkApplicationResponseBody {
	s.Supplements = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetThirdClassifications(v []*DescribePartnerTrademarkApplicationResponseBodyThirdClassifications) *DescribePartnerTrademarkApplicationResponseBody {
	s.ThirdClassifications = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetTotalPrice(v float32) *DescribePartnerTrademarkApplicationResponseBody {
	s.TotalPrice = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetTrademarkIcon(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.TrademarkIcon = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetTrademarkName(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.TrademarkName = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetTrademarkNameType(v int32) *DescribePartnerTrademarkApplicationResponseBody {
	s.TrademarkNameType = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetTrademarkNumber(v string) *DescribePartnerTrademarkApplicationResponseBody {
	s.TrademarkNumber = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBody) SetUpdateTime(v int64) *DescribePartnerTrademarkApplicationResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribePartnerTrademarkApplicationResponseBodyApplicant struct {
	Address            *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantName      *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantRegion    *int32  `json:"ApplicantRegion,omitempty" xml:"ApplicantRegion,omitempty"`
	ApplicantType      *int32  `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	AuditStatus        *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuthorizationUrl   *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BusinessLicenceUrl *string `json:"BusinessLicenceUrl,omitempty" xml:"BusinessLicenceUrl,omitempty"`
	CardNumber         *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ContactAddress     *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactEmail       *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName        *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber      *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactZipcode     *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Country            *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EAddress           *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	EName              *string `json:"EName,omitempty" xml:"EName,omitempty"`
	IdCardUrl          *string `json:"IdCardUrl,omitempty" xml:"IdCardUrl,omitempty"`
	LegalNoticeUrl     *string `json:"LegalNoticeUrl,omitempty" xml:"LegalNoticeUrl,omitempty"`
	PassportUrl        *string `json:"PassportUrl,omitempty" xml:"PassportUrl,omitempty"`
	PrincipalName      *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Province           *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribePartnerTrademarkApplicationResponseBodyApplicant) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerTrademarkApplicationResponseBodyApplicant) GoString() string {
	return s.String()
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetAddress(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.Address = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetApplicantName(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ApplicantName = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetApplicantRegion(v int32) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ApplicantRegion = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetApplicantType(v int32) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ApplicantType = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetAuditStatus(v int32) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.AuditStatus = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetAuthorizationUrl(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.AuthorizationUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetBusinessLicenceUrl(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.BusinessLicenceUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetCardNumber(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.CardNumber = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetContactAddress(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ContactAddress = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetContactEmail(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ContactEmail = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetContactName(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ContactName = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetContactNumber(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ContactNumber = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetContactZipcode(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.ContactZipcode = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetCountry(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.Country = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetEAddress(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.EAddress = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetEName(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.EName = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetIdCardUrl(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.IdCardUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetLegalNoticeUrl(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.LegalNoticeUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetPassportUrl(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.PassportUrl = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetPrincipalName(v int32) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.PrincipalName = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyApplicant) SetProvince(v string) *DescribePartnerTrademarkApplicationResponseBodyApplicant {
	s.Province = &v
	return s
}

type DescribePartnerTrademarkApplicationResponseBodyFirstClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s DescribePartnerTrademarkApplicationResponseBodyFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerTrademarkApplicationResponseBodyFirstClassification) GoString() string {
	return s.String()
}

func (s *DescribePartnerTrademarkApplicationResponseBodyFirstClassification) SetClassificationCode(v string) *DescribePartnerTrademarkApplicationResponseBodyFirstClassification {
	s.ClassificationCode = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyFirstClassification) SetClassificationName(v string) *DescribePartnerTrademarkApplicationResponseBodyFirstClassification {
	s.ClassificationName = &v
	return s
}

type DescribePartnerTrademarkApplicationResponseBodySupplements struct {
	AcceptExpirationDate *int64    `json:"AcceptExpirationDate,omitempty" xml:"AcceptExpirationDate,omitempty"`
	AcceptTime           *int64    `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	ApplicationType      *int32    `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Content              *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	OfficialFile         *string   `json:"OfficialFile,omitempty" xml:"OfficialFile,omitempty"`
	OperateTime          *int64    `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OrderId              *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	SbjExpirationDate    *int64    `json:"SbjExpirationDate,omitempty" xml:"SbjExpirationDate,omitempty"`
	SendTime             *int64    `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	SerialNumber         *string   `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SupplementId         *int64    `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	SupplementStatus     *int32    `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TrademarkNumber      *string   `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UserFiles            []*string `json:"UserFiles,omitempty" xml:"UserFiles,omitempty" type:"Repeated"`
}

func (s DescribePartnerTrademarkApplicationResponseBodySupplements) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerTrademarkApplicationResponseBodySupplements) GoString() string {
	return s.String()
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetAcceptExpirationDate(v int64) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.AcceptExpirationDate = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetAcceptTime(v int64) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.AcceptTime = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetApplicationType(v int32) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.ApplicationType = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetContent(v string) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.Content = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetOfficialFile(v string) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.OfficialFile = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetOperateTime(v int64) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.OperateTime = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetOrderId(v string) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.OrderId = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetSbjExpirationDate(v int64) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.SbjExpirationDate = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetSendTime(v int64) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.SendTime = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetSerialNumber(v string) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.SerialNumber = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetSupplementId(v int64) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.SupplementId = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetSupplementStatus(v int32) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.SupplementStatus = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetTrademarkNumber(v string) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.TrademarkNumber = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodySupplements) SetUserFiles(v []*string) *DescribePartnerTrademarkApplicationResponseBodySupplements {
	s.UserFiles = v
	return s
}

type DescribePartnerTrademarkApplicationResponseBodyThirdClassifications struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s DescribePartnerTrademarkApplicationResponseBodyThirdClassifications) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerTrademarkApplicationResponseBodyThirdClassifications) GoString() string {
	return s.String()
}

func (s *DescribePartnerTrademarkApplicationResponseBodyThirdClassifications) SetClassificationCode(v string) *DescribePartnerTrademarkApplicationResponseBodyThirdClassifications {
	s.ClassificationCode = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponseBodyThirdClassifications) SetClassificationName(v string) *DescribePartnerTrademarkApplicationResponseBodyThirdClassifications {
	s.ClassificationName = &v
	return s
}

type DescribePartnerTrademarkApplicationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePartnerTrademarkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePartnerTrademarkApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePartnerTrademarkApplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribePartnerTrademarkApplicationResponse) SetHeaders(v map[string]*string) *DescribePartnerTrademarkApplicationResponse {
	s.Headers = v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponse) SetStatusCode(v int32) *DescribePartnerTrademarkApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePartnerTrademarkApplicationResponse) SetBody(v *DescribePartnerTrademarkApplicationResponseBody) *DescribePartnerTrademarkApplicationResponse {
	s.Body = v
	return s
}

type DescribeQualificationStatusRequest struct {
	TbUid *string `json:"TbUid,omitempty" xml:"TbUid,omitempty"`
}

func (s DescribeQualificationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualificationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualificationStatusRequest) SetTbUid(v string) *DescribeQualificationStatusRequest {
	s.TbUid = &v
	return s
}

type DescribeQualificationStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeQualificationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualificationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualificationStatusResponseBody) SetCode(v string) *DescribeQualificationStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeQualificationStatusResponseBody) SetMessage(v string) *DescribeQualificationStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeQualificationStatusResponseBody) SetRequestId(v string) *DescribeQualificationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQualificationStatusResponseBody) SetStatus(v string) *DescribeQualificationStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeQualificationStatusResponseBody) SetSuccess(v bool) *DescribeQualificationStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeQualificationStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualificationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualificationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualificationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualificationStatusResponse) SetHeaders(v map[string]*string) *DescribeQualificationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualificationStatusResponse) SetStatusCode(v int32) *DescribeQualificationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualificationStatusResponse) SetBody(v *DescribeQualificationStatusResponseBody) *DescribeQualificationStatusResponse {
	s.Body = v
	return s
}

type DescribeSupplementRequest struct {
	SupplementId *int64 `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
}

func (s DescribeSupplementRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupplementRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupplementRequest) SetSupplementId(v int64) *DescribeSupplementRequest {
	s.SupplementId = &v
	return s
}

type DescribeSupplementResponseBody struct {
	AcceptExpirationDate *int64                                   `json:"AcceptExpirationDate,omitempty" xml:"AcceptExpirationDate,omitempty"`
	AcceptTime           *int64                                   `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	ApplicationType      *int32                                   `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Content              *string                                  `json:"Content,omitempty" xml:"Content,omitempty"`
	OfficialFile         *string                                  `json:"OfficialFile,omitempty" xml:"OfficialFile,omitempty"`
	OperateTime          *int64                                   `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	RequestId            *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SbjExpirationDate    *int64                                   `json:"SbjExpirationDate,omitempty" xml:"SbjExpirationDate,omitempty"`
	SendTime             *int64                                   `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	SerialNumber         *string                                  `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SupplementId         *int64                                   `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	SupplementStatus     *int32                                   `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TrademarkNumber      *string                                  `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UserFiles            *DescribeSupplementResponseBodyUserFiles `json:"UserFiles,omitempty" xml:"UserFiles,omitempty" type:"Struct"`
}

func (s DescribeSupplementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupplementResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupplementResponseBody) SetAcceptExpirationDate(v int64) *DescribeSupplementResponseBody {
	s.AcceptExpirationDate = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetAcceptTime(v int64) *DescribeSupplementResponseBody {
	s.AcceptTime = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetApplicationType(v int32) *DescribeSupplementResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetContent(v string) *DescribeSupplementResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetOfficialFile(v string) *DescribeSupplementResponseBody {
	s.OfficialFile = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetOperateTime(v int64) *DescribeSupplementResponseBody {
	s.OperateTime = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetRequestId(v string) *DescribeSupplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetSbjExpirationDate(v int64) *DescribeSupplementResponseBody {
	s.SbjExpirationDate = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetSendTime(v int64) *DescribeSupplementResponseBody {
	s.SendTime = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetSerialNumber(v string) *DescribeSupplementResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetSupplementId(v int64) *DescribeSupplementResponseBody {
	s.SupplementId = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetSupplementStatus(v int32) *DescribeSupplementResponseBody {
	s.SupplementStatus = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetTrademarkNumber(v string) *DescribeSupplementResponseBody {
	s.TrademarkNumber = &v
	return s
}

func (s *DescribeSupplementResponseBody) SetUserFiles(v *DescribeSupplementResponseBodyUserFiles) *DescribeSupplementResponseBody {
	s.UserFiles = v
	return s
}

type DescribeSupplementResponseBodyUserFiles struct {
	UserFile []*string `json:"UserFile,omitempty" xml:"UserFile,omitempty" type:"Repeated"`
}

func (s DescribeSupplementResponseBodyUserFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupplementResponseBodyUserFiles) GoString() string {
	return s.String()
}

func (s *DescribeSupplementResponseBodyUserFiles) SetUserFile(v []*string) *DescribeSupplementResponseBodyUserFiles {
	s.UserFile = v
	return s
}

type DescribeSupplementResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupplementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupplementResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSupplementResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupplementResponse) SetHeaders(v map[string]*string) *DescribeSupplementResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupplementResponse) SetStatusCode(v int32) *DescribeSupplementResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupplementResponse) SetBody(v *DescribeSupplementResponseBody) *DescribeSupplementResponse {
	s.Body = v
	return s
}

type DescribeTrademarkApplicationRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s DescribeTrademarkApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationRequest) SetBizId(v string) *DescribeTrademarkApplicationRequest {
	s.BizId = &v
	return s
}

type DescribeTrademarkApplicationResponseBody struct {
	AcceptUrl            *string                                                       `json:"AcceptUrl,omitempty" xml:"AcceptUrl,omitempty"`
	AgreementId          *string                                                       `json:"AgreementId,omitempty" xml:"AgreementId,omitempty"`
	Applicant            *DescribeTrademarkApplicationResponseBodyApplicant            `json:"Applicant,omitempty" xml:"Applicant,omitempty" type:"Struct"`
	ApplicantId          *int64                                                        `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicationStatus    *int32                                                        `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	ApplicationType      *int32                                                        `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	AuthorizationUrl     *string                                                       `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BizId                *string                                                       `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BlackAndWhiteIconUrl *string                                                       `json:"BlackAndWhiteIconUrl,omitempty" xml:"BlackAndWhiteIconUrl,omitempty"`
	CreateTime           *int64                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtendInfo           map[string]interface{}                                        `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FirstClassification  *DescribeTrademarkApplicationResponseBodyFirstClassification  `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
	Flags                *DescribeTrademarkApplicationResponseBodyFlags                `json:"Flags,omitempty" xml:"Flags,omitempty" type:"Struct"`
	JudgeResultUrls      *DescribeTrademarkApplicationResponseBodyJudgeResultUrls      `json:"JudgeResultUrls,omitempty" xml:"JudgeResultUrls,omitempty" type:"Struct"`
	Note                 *string                                                       `json:"Note,omitempty" xml:"Note,omitempty"`
	OrderId              *string                                                       `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderPrice           *float32                                                      `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	PrincipalName        *int32                                                        `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	ReceiptUrls          *DescribeTrademarkApplicationResponseBodyReceiptUrls          `json:"ReceiptUrls,omitempty" xml:"ReceiptUrls,omitempty" type:"Struct"`
	RecvUserLogistics    *string                                                       `json:"RecvUserLogistics,omitempty" xml:"RecvUserLogistics,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SendSbjLogistics     *string                                                       `json:"SendSbjLogistics,omitempty" xml:"SendSbjLogistics,omitempty"`
	SendUserLogistics    *string                                                       `json:"SendUserLogistics,omitempty" xml:"SendUserLogistics,omitempty"`
	ServicePrice         *float32                                                      `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	Supplements          *DescribeTrademarkApplicationResponseBodySupplements          `json:"Supplements,omitempty" xml:"Supplements,omitempty" type:"Struct"`
	ThirdClassifications *DescribeTrademarkApplicationResponseBodyThirdClassifications `json:"ThirdClassifications,omitempty" xml:"ThirdClassifications,omitempty" type:"Struct"`
	TotalPrice           *float32                                                      `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	TrademarkIcon        *string                                                       `json:"TrademarkIcon,omitempty" xml:"TrademarkIcon,omitempty"`
	TrademarkName        *string                                                       `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNameType    *int32                                                        `json:"TrademarkNameType,omitempty" xml:"TrademarkNameType,omitempty"`
	TrademarkNumber      *string                                                       `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UpdateTime           *int64                                                        `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTrademarkApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBody) SetAcceptUrl(v string) *DescribeTrademarkApplicationResponseBody {
	s.AcceptUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetAgreementId(v string) *DescribeTrademarkApplicationResponseBody {
	s.AgreementId = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetApplicant(v *DescribeTrademarkApplicationResponseBodyApplicant) *DescribeTrademarkApplicationResponseBody {
	s.Applicant = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetApplicantId(v int64) *DescribeTrademarkApplicationResponseBody {
	s.ApplicantId = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetApplicationStatus(v int32) *DescribeTrademarkApplicationResponseBody {
	s.ApplicationStatus = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetApplicationType(v int32) *DescribeTrademarkApplicationResponseBody {
	s.ApplicationType = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetAuthorizationUrl(v string) *DescribeTrademarkApplicationResponseBody {
	s.AuthorizationUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetBizId(v string) *DescribeTrademarkApplicationResponseBody {
	s.BizId = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetBlackAndWhiteIconUrl(v string) *DescribeTrademarkApplicationResponseBody {
	s.BlackAndWhiteIconUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetCreateTime(v int64) *DescribeTrademarkApplicationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetExtendInfo(v map[string]interface{}) *DescribeTrademarkApplicationResponseBody {
	s.ExtendInfo = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetFirstClassification(v *DescribeTrademarkApplicationResponseBodyFirstClassification) *DescribeTrademarkApplicationResponseBody {
	s.FirstClassification = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetFlags(v *DescribeTrademarkApplicationResponseBodyFlags) *DescribeTrademarkApplicationResponseBody {
	s.Flags = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetJudgeResultUrls(v *DescribeTrademarkApplicationResponseBodyJudgeResultUrls) *DescribeTrademarkApplicationResponseBody {
	s.JudgeResultUrls = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetNote(v string) *DescribeTrademarkApplicationResponseBody {
	s.Note = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetOrderId(v string) *DescribeTrademarkApplicationResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetOrderPrice(v float32) *DescribeTrademarkApplicationResponseBody {
	s.OrderPrice = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetPrincipalName(v int32) *DescribeTrademarkApplicationResponseBody {
	s.PrincipalName = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetReceiptUrls(v *DescribeTrademarkApplicationResponseBodyReceiptUrls) *DescribeTrademarkApplicationResponseBody {
	s.ReceiptUrls = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetRecvUserLogistics(v string) *DescribeTrademarkApplicationResponseBody {
	s.RecvUserLogistics = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetRequestId(v string) *DescribeTrademarkApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetSendSbjLogistics(v string) *DescribeTrademarkApplicationResponseBody {
	s.SendSbjLogistics = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetSendUserLogistics(v string) *DescribeTrademarkApplicationResponseBody {
	s.SendUserLogistics = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetServicePrice(v float32) *DescribeTrademarkApplicationResponseBody {
	s.ServicePrice = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetSupplements(v *DescribeTrademarkApplicationResponseBodySupplements) *DescribeTrademarkApplicationResponseBody {
	s.Supplements = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetThirdClassifications(v *DescribeTrademarkApplicationResponseBodyThirdClassifications) *DescribeTrademarkApplicationResponseBody {
	s.ThirdClassifications = v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetTotalPrice(v float32) *DescribeTrademarkApplicationResponseBody {
	s.TotalPrice = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetTrademarkIcon(v string) *DescribeTrademarkApplicationResponseBody {
	s.TrademarkIcon = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetTrademarkName(v string) *DescribeTrademarkApplicationResponseBody {
	s.TrademarkName = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetTrademarkNameType(v int32) *DescribeTrademarkApplicationResponseBody {
	s.TrademarkNameType = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetTrademarkNumber(v string) *DescribeTrademarkApplicationResponseBody {
	s.TrademarkNumber = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBody) SetUpdateTime(v int64) *DescribeTrademarkApplicationResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeTrademarkApplicationResponseBodyApplicant struct {
	Address            *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantName      *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantRegion    *int32  `json:"ApplicantRegion,omitempty" xml:"ApplicantRegion,omitempty"`
	ApplicantType      *int32  `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	AuditStatus        *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuthorizationUrl   *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BusinessLicenceUrl *string `json:"BusinessLicenceUrl,omitempty" xml:"BusinessLicenceUrl,omitempty"`
	CardNumber         *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ContactAddress     *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactEmail       *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName        *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber      *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactZipcode     *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	Country            *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EAddress           *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	EName              *string `json:"EName,omitempty" xml:"EName,omitempty"`
	IdCardName         *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber       *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	IdCardUrl          *string `json:"IdCardUrl,omitempty" xml:"IdCardUrl,omitempty"`
	LegalNoticeUrl     *string `json:"LegalNoticeUrl,omitempty" xml:"LegalNoticeUrl,omitempty"`
	PassportUrl        *string `json:"PassportUrl,omitempty" xml:"PassportUrl,omitempty"`
	PersonalType       *int64  `json:"PersonalType,omitempty" xml:"PersonalType,omitempty"`
	PrincipalName      *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Province           *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeTrademarkApplicationResponseBodyApplicant) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodyApplicant) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetAddress(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.Address = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetApplicantName(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ApplicantName = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetApplicantRegion(v int32) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ApplicantRegion = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetApplicantType(v int32) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ApplicantType = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetAuditStatus(v int32) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.AuditStatus = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetAuthorizationUrl(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.AuthorizationUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetBusinessLicenceUrl(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.BusinessLicenceUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetCardNumber(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.CardNumber = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetContactAddress(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ContactAddress = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetContactEmail(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ContactEmail = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetContactName(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ContactName = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetContactNumber(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ContactNumber = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetContactZipcode(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.ContactZipcode = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetCountry(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.Country = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetEAddress(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.EAddress = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetEName(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.EName = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetIdCardName(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.IdCardName = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetIdCardNumber(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.IdCardNumber = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetIdCardUrl(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.IdCardUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetLegalNoticeUrl(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.LegalNoticeUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetPassportUrl(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.PassportUrl = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetPersonalType(v int64) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.PersonalType = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetPrincipalName(v int32) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.PrincipalName = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyApplicant) SetProvince(v string) *DescribeTrademarkApplicationResponseBodyApplicant {
	s.Province = &v
	return s
}

type DescribeTrademarkApplicationResponseBodyFirstClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s DescribeTrademarkApplicationResponseBodyFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodyFirstClassification) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodyFirstClassification) SetClassificationCode(v string) *DescribeTrademarkApplicationResponseBodyFirstClassification {
	s.ClassificationCode = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyFirstClassification) SetClassificationName(v string) *DescribeTrademarkApplicationResponseBodyFirstClassification {
	s.ClassificationName = &v
	return s
}

type DescribeTrademarkApplicationResponseBodyFlags struct {
	Flag []*int32 `json:"Flag,omitempty" xml:"Flag,omitempty" type:"Repeated"`
}

func (s DescribeTrademarkApplicationResponseBodyFlags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodyFlags) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodyFlags) SetFlag(v []*int32) *DescribeTrademarkApplicationResponseBodyFlags {
	s.Flag = v
	return s
}

type DescribeTrademarkApplicationResponseBodyJudgeResultUrls struct {
	JudgeResultUrl []*string `json:"JudgeResultUrl,omitempty" xml:"JudgeResultUrl,omitempty" type:"Repeated"`
}

func (s DescribeTrademarkApplicationResponseBodyJudgeResultUrls) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodyJudgeResultUrls) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodyJudgeResultUrls) SetJudgeResultUrl(v []*string) *DescribeTrademarkApplicationResponseBodyJudgeResultUrls {
	s.JudgeResultUrl = v
	return s
}

type DescribeTrademarkApplicationResponseBodyReceiptUrls struct {
	ReceiptUrl []*string `json:"ReceiptUrl,omitempty" xml:"ReceiptUrl,omitempty" type:"Repeated"`
}

func (s DescribeTrademarkApplicationResponseBodyReceiptUrls) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodyReceiptUrls) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodyReceiptUrls) SetReceiptUrl(v []*string) *DescribeTrademarkApplicationResponseBodyReceiptUrls {
	s.ReceiptUrl = v
	return s
}

type DescribeTrademarkApplicationResponseBodySupplements struct {
	Supplement []*DescribeTrademarkApplicationResponseBodySupplementsSupplement `json:"Supplement,omitempty" xml:"Supplement,omitempty" type:"Repeated"`
}

func (s DescribeTrademarkApplicationResponseBodySupplements) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodySupplements) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodySupplements) SetSupplement(v []*DescribeTrademarkApplicationResponseBodySupplementsSupplement) *DescribeTrademarkApplicationResponseBodySupplements {
	s.Supplement = v
	return s
}

type DescribeTrademarkApplicationResponseBodySupplementsSupplement struct {
	AcceptExpirationDate *int64                                                                  `json:"AcceptExpirationDate,omitempty" xml:"AcceptExpirationDate,omitempty"`
	AcceptTime           *int64                                                                  `json:"AcceptTime,omitempty" xml:"AcceptTime,omitempty"`
	ApplicationType      *int32                                                                  `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Content              *string                                                                 `json:"Content,omitempty" xml:"Content,omitempty"`
	OfficialFile         *string                                                                 `json:"OfficialFile,omitempty" xml:"OfficialFile,omitempty"`
	OperateTime          *int64                                                                  `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OrderId              *string                                                                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	SbjExpirationDate    *int64                                                                  `json:"SbjExpirationDate,omitempty" xml:"SbjExpirationDate,omitempty"`
	SendTime             *int64                                                                  `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	SerialNumber         *string                                                                 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SupplementId         *int64                                                                  `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	SupplementStatus     *int32                                                                  `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TrademarkNumber      *string                                                                 `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UserFiles            *DescribeTrademarkApplicationResponseBodySupplementsSupplementUserFiles `json:"UserFiles,omitempty" xml:"UserFiles,omitempty" type:"Struct"`
}

func (s DescribeTrademarkApplicationResponseBodySupplementsSupplement) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodySupplementsSupplement) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetAcceptExpirationDate(v int64) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.AcceptExpirationDate = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetAcceptTime(v int64) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.AcceptTime = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetApplicationType(v int32) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.ApplicationType = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetContent(v string) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.Content = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetOfficialFile(v string) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.OfficialFile = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetOperateTime(v int64) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.OperateTime = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetOrderId(v string) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.OrderId = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetSbjExpirationDate(v int64) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.SbjExpirationDate = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetSendTime(v int64) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.SendTime = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetSerialNumber(v string) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.SerialNumber = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetSupplementId(v int64) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.SupplementId = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetSupplementStatus(v int32) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.SupplementStatus = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetTrademarkNumber(v string) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.TrademarkNumber = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplement) SetUserFiles(v *DescribeTrademarkApplicationResponseBodySupplementsSupplementUserFiles) *DescribeTrademarkApplicationResponseBodySupplementsSupplement {
	s.UserFiles = v
	return s
}

type DescribeTrademarkApplicationResponseBodySupplementsSupplementUserFiles struct {
	UserFile []*string `json:"UserFile,omitempty" xml:"UserFile,omitempty" type:"Repeated"`
}

func (s DescribeTrademarkApplicationResponseBodySupplementsSupplementUserFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodySupplementsSupplementUserFiles) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodySupplementsSupplementUserFiles) SetUserFile(v []*string) *DescribeTrademarkApplicationResponseBodySupplementsSupplementUserFiles {
	s.UserFile = v
	return s
}

type DescribeTrademarkApplicationResponseBodyThirdClassifications struct {
	ThirdClassification []*DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Repeated"`
}

func (s DescribeTrademarkApplicationResponseBodyThirdClassifications) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodyThirdClassifications) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodyThirdClassifications) SetThirdClassification(v []*DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification) *DescribeTrademarkApplicationResponseBodyThirdClassifications {
	s.ThirdClassification = v
	return s
}

type DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification) SetClassificationCode(v string) *DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification {
	s.ClassificationCode = &v
	return s
}

func (s *DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification) SetClassificationName(v string) *DescribeTrademarkApplicationResponseBodyThirdClassificationsThirdClassification {
	s.ClassificationName = &v
	return s
}

type DescribeTrademarkApplicationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrademarkApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrademarkApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkApplicationResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkApplicationResponse) SetHeaders(v map[string]*string) *DescribeTrademarkApplicationResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrademarkApplicationResponse) SetStatusCode(v int32) *DescribeTrademarkApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrademarkApplicationResponse) SetBody(v *DescribeTrademarkApplicationResponseBody) *DescribeTrademarkApplicationResponse {
	s.Body = v
	return s
}

type DescribeTrademarkDetailForInnerRequest struct {
	Uid  *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Umid *string `json:"Umid,omitempty" xml:"Umid,omitempty"`
}

func (s DescribeTrademarkDetailForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkDetailForInnerRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkDetailForInnerRequest) SetUid(v string) *DescribeTrademarkDetailForInnerRequest {
	s.Uid = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerRequest) SetUmid(v string) *DescribeTrademarkDetailForInnerRequest {
	s.Umid = &v
	return s
}

type DescribeTrademarkDetailForInnerResponseBody struct {
	Agency                    *string                                                  `json:"Agency,omitempty" xml:"Agency,omitempty"`
	ApplyDate                 *string                                                  `json:"ApplyDate,omitempty" xml:"ApplyDate,omitempty"`
	Classification            *string                                                  `json:"Classification,omitempty" xml:"Classification,omitempty"`
	ExclusiveDateLimit        *string                                                  `json:"ExclusiveDateLimit,omitempty" xml:"ExclusiveDateLimit,omitempty"`
	FlowList                  []*DescribeTrademarkDetailForInnerResponseBodyFlowList   `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	Image                     *string                                                  `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageElement              *string                                                  `json:"ImageElement,omitempty" xml:"ImageElement,omitempty"`
	IntlRegDate               *string                                                  `json:"IntlRegDate,omitempty" xml:"IntlRegDate,omitempty"`
	LastProcedureStatus       *string                                                  `json:"LastProcedureStatus,omitempty" xml:"LastProcedureStatus,omitempty"`
	Name                      *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	NoticeList                []*DescribeTrademarkDetailForInnerResponseBodyNoticeList `json:"NoticeList,omitempty" xml:"NoticeList,omitempty" type:"Repeated"`
	OwnerAddress              *string                                                  `json:"OwnerAddress,omitempty" xml:"OwnerAddress,omitempty"`
	OwnerEnAddress            *string                                                  `json:"OwnerEnAddress,omitempty" xml:"OwnerEnAddress,omitempty"`
	OwnerEnName               *string                                                  `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	OwnerName                 *string                                                  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	PreAnnDate                *string                                                  `json:"PreAnnDate,omitempty" xml:"PreAnnDate,omitempty"`
	PreAnnNum                 *string                                                  `json:"PreAnnNum,omitempty" xml:"PreAnnNum,omitempty"`
	PriorityDate              *string                                                  `json:"PriorityDate,omitempty" xml:"PriorityDate,omitempty"`
	Product                   *string                                                  `json:"Product,omitempty" xml:"Product,omitempty"`
	ProductDescription        *string                                                  `json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	RegAnnDate                *string                                                  `json:"RegAnnDate,omitempty" xml:"RegAnnDate,omitempty"`
	RegAnnNum                 *int32                                                   `json:"RegAnnNum,omitempty" xml:"RegAnnNum,omitempty"`
	RegistrationNumber        *string                                                  `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	RegistrationType          *string                                                  `json:"RegistrationType,omitempty" xml:"RegistrationType,omitempty"`
	RequestId                 *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Share                     *string                                                  `json:"Share,omitempty" xml:"Share,omitempty"`
	SimilarGroup              *string                                                  `json:"SimilarGroup,omitempty" xml:"SimilarGroup,omitempty"`
	Status                    *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubsequentDesignationDate *string                                                  `json:"SubsequentDesignationDate,omitempty" xml:"SubsequentDesignationDate,omitempty"`
	Uid                       *string                                                  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeTrademarkDetailForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkDetailForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetAgency(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Agency = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetApplyDate(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.ApplyDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetClassification(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Classification = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetExclusiveDateLimit(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.ExclusiveDateLimit = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetFlowList(v []*DescribeTrademarkDetailForInnerResponseBodyFlowList) *DescribeTrademarkDetailForInnerResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetImage(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Image = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetImageElement(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.ImageElement = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetIntlRegDate(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.IntlRegDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetLastProcedureStatus(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.LastProcedureStatus = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetName(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetNoticeList(v []*DescribeTrademarkDetailForInnerResponseBodyNoticeList) *DescribeTrademarkDetailForInnerResponseBody {
	s.NoticeList = v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetOwnerAddress(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.OwnerAddress = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetOwnerEnAddress(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.OwnerEnAddress = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetOwnerEnName(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.OwnerEnName = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetOwnerName(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.OwnerName = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetPreAnnDate(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.PreAnnDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetPreAnnNum(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.PreAnnNum = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetPriorityDate(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.PriorityDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetProduct(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Product = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetProductDescription(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.ProductDescription = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetRegAnnDate(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.RegAnnDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetRegAnnNum(v int32) *DescribeTrademarkDetailForInnerResponseBody {
	s.RegAnnNum = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetRegistrationNumber(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.RegistrationNumber = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetRegistrationType(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.RegistrationType = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetRequestId(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetShare(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Share = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetSimilarGroup(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.SimilarGroup = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetStatus(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetSubsequentDesignationDate(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.SubsequentDesignationDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBody) SetUid(v string) *DescribeTrademarkDetailForInnerResponseBody {
	s.Uid = &v
	return s
}

type DescribeTrademarkDetailForInnerResponseBodyFlowList struct {
	Date               *string `json:"Date,omitempty" xml:"Date,omitempty"`
	ProcedureCode      *string `json:"ProcedureCode,omitempty" xml:"ProcedureCode,omitempty"`
	ProcedureDate      *string `json:"ProcedureDate,omitempty" xml:"ProcedureDate,omitempty"`
	ProcedureName      *string `json:"ProcedureName,omitempty" xml:"ProcedureName,omitempty"`
	ProcedureResult    *string `json:"ProcedureResult,omitempty" xml:"ProcedureResult,omitempty"`
	ProcedureStep      *string `json:"ProcedureStep,omitempty" xml:"ProcedureStep,omitempty"`
	RegistrationNumber *string `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
}

func (s DescribeTrademarkDetailForInnerResponseBodyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkDetailForInnerResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkDetailForInnerResponseBodyFlowList) SetDate(v string) *DescribeTrademarkDetailForInnerResponseBodyFlowList {
	s.Date = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyFlowList) SetProcedureCode(v string) *DescribeTrademarkDetailForInnerResponseBodyFlowList {
	s.ProcedureCode = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyFlowList) SetProcedureDate(v string) *DescribeTrademarkDetailForInnerResponseBodyFlowList {
	s.ProcedureDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyFlowList) SetProcedureName(v string) *DescribeTrademarkDetailForInnerResponseBodyFlowList {
	s.ProcedureName = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyFlowList) SetProcedureResult(v string) *DescribeTrademarkDetailForInnerResponseBodyFlowList {
	s.ProcedureResult = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyFlowList) SetProcedureStep(v string) *DescribeTrademarkDetailForInnerResponseBodyFlowList {
	s.ProcedureStep = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyFlowList) SetRegistrationNumber(v string) *DescribeTrademarkDetailForInnerResponseBodyFlowList {
	s.RegistrationNumber = &v
	return s
}

type DescribeTrademarkDetailForInnerResponseBodyNoticeList struct {
	AnnDate            *string `json:"AnnDate,omitempty" xml:"AnnDate,omitempty"`
	AnnId              *string `json:"AnnId,omitempty" xml:"AnnId,omitempty"`
	AnnNum             *string `json:"AnnNum,omitempty" xml:"AnnNum,omitempty"`
	AnnTypeCode        *string `json:"AnnTypeCode,omitempty" xml:"AnnTypeCode,omitempty"`
	AnnTypeName        *string `json:"AnnTypeName,omitempty" xml:"AnnTypeName,omitempty"`
	Applicant          *string `json:"Applicant,omitempty" xml:"Applicant,omitempty"`
	Date               *string `json:"Date,omitempty" xml:"Date,omitempty"`
	ImageUrl           *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	OriginalImageUrl   *string `json:"OriginalImageUrl,omitempty" xml:"OriginalImageUrl,omitempty"`
	PageNo             *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	ReactNum           *string `json:"ReactNum,omitempty" xml:"ReactNum,omitempty"`
	RegistrationNumber *string `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	TrademarkName      *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
}

func (s DescribeTrademarkDetailForInnerResponseBodyNoticeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkDetailForInnerResponseBodyNoticeList) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetAnnDate(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.AnnDate = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetAnnId(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.AnnId = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetAnnNum(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.AnnNum = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetAnnTypeCode(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.AnnTypeCode = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetAnnTypeName(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.AnnTypeName = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetApplicant(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.Applicant = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetDate(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.Date = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetImageUrl(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.ImageUrl = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetOriginalImageUrl(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.OriginalImageUrl = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetPageNo(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.PageNo = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetReactNum(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.ReactNum = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetRegistrationNumber(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.RegistrationNumber = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponseBodyNoticeList) SetTrademarkName(v string) *DescribeTrademarkDetailForInnerResponseBodyNoticeList {
	s.TrademarkName = &v
	return s
}

type DescribeTrademarkDetailForInnerResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrademarkDetailForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrademarkDetailForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrademarkDetailForInnerResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrademarkDetailForInnerResponse) SetHeaders(v map[string]*string) *DescribeTrademarkDetailForInnerResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponse) SetStatusCode(v int32) *DescribeTrademarkDetailForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrademarkDetailForInnerResponse) SetBody(v *DescribeTrademarkDetailForInnerResponseBody) *DescribeTrademarkDetailForInnerResponse {
	s.Body = v
	return s
}

type GenerateUploadFilePolicyRequest struct {
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GenerateUploadFilePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyRequest) SetFileType(v string) *GenerateUploadFilePolicyRequest {
	s.FileType = &v
	return s
}

type GenerateUploadFilePolicyResponseBody struct {
	// OSSAccessKeyId
	AccessId      *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	ExpireTime    *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	FileDir       *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	// host
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateUploadFilePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponseBody) SetAccessId(v string) *GenerateUploadFilePolicyResponseBody {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetEncodedPolicy(v string) *GenerateUploadFilePolicyResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetExpireTime(v int64) *GenerateUploadFilePolicyResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetFileDir(v string) *GenerateUploadFilePolicyResponseBody {
	s.FileDir = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetHost(v string) *GenerateUploadFilePolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetRequestId(v string) *GenerateUploadFilePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadFilePolicyResponseBody) SetSignature(v string) *GenerateUploadFilePolicyResponseBody {
	s.Signature = &v
	return s
}

type GenerateUploadFilePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadFilePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadFilePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyResponse) SetHeaders(v map[string]*string) *GenerateUploadFilePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadFilePolicyResponse) SetStatusCode(v int32) *GenerateUploadFilePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadFilePolicyResponse) SetBody(v *GenerateUploadFilePolicyResponseBody) *GenerateUploadFilePolicyResponse {
	s.Body = v
	return s
}

type GetAlipayUrlRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlipayUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlipayUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlRequest) SetBizType(v string) *GetAlipayUrlRequest {
	s.BizType = &v
	return s
}

func (s *GetAlipayUrlRequest) SetOrderId(v int64) *GetAlipayUrlRequest {
	s.OrderId = &v
	return s
}

func (s *GetAlipayUrlRequest) SetReturnUrl(v string) *GetAlipayUrlRequest {
	s.ReturnUrl = &v
	return s
}

func (s *GetAlipayUrlRequest) SetType(v string) *GetAlipayUrlRequest {
	s.Type = &v
	return s
}

type GetAlipayUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlipayUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlipayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponseBody) SetData(v string) *GetAlipayUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetAlipayUrlResponseBody) SetRequestId(v string) *GetAlipayUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetAlipayUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlipayUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlipayUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlipayUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAlipayUrlResponse) SetHeaders(v map[string]*string) *GetAlipayUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAlipayUrlResponse) SetStatusCode(v int32) *GetAlipayUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlipayUrlResponse) SetBody(v *GetAlipayUrlResponseBody) *GetAlipayUrlResponse {
	s.Body = v
	return s
}

type GetOrderConfirmUrlRequest struct {
	Items        []*GetOrderConfirmUrlRequestItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	OutTraceCode *string                           `json:"OutTraceCode,omitempty" xml:"OutTraceCode,omitempty"`
	OutTraceType *string                           `json:"OutTraceType,omitempty" xml:"OutTraceType,omitempty"`
}

func (s GetOrderConfirmUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrderConfirmUrlRequest) GoString() string {
	return s.String()
}

func (s *GetOrderConfirmUrlRequest) SetItems(v []*GetOrderConfirmUrlRequestItems) *GetOrderConfirmUrlRequest {
	s.Items = v
	return s
}

func (s *GetOrderConfirmUrlRequest) SetOutTraceCode(v string) *GetOrderConfirmUrlRequest {
	s.OutTraceCode = &v
	return s
}

func (s *GetOrderConfirmUrlRequest) SetOutTraceType(v string) *GetOrderConfirmUrlRequest {
	s.OutTraceType = &v
	return s
}

type GetOrderConfirmUrlRequestItems struct {
	ItemCode *string `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	Quantity *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s GetOrderConfirmUrlRequestItems) String() string {
	return tea.Prettify(s)
}

func (s GetOrderConfirmUrlRequestItems) GoString() string {
	return s.String()
}

func (s *GetOrderConfirmUrlRequestItems) SetItemCode(v string) *GetOrderConfirmUrlRequestItems {
	s.ItemCode = &v
	return s
}

func (s *GetOrderConfirmUrlRequestItems) SetQuantity(v int64) *GetOrderConfirmUrlRequestItems {
	s.Quantity = &v
	return s
}

type GetOrderConfirmUrlResponseBody struct {
	ConfirmUrl   *string `json:"ConfirmUrl,omitempty" xml:"ConfirmUrl,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderConfirmUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrderConfirmUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderConfirmUrlResponseBody) SetConfirmUrl(v string) *GetOrderConfirmUrlResponseBody {
	s.ConfirmUrl = &v
	return s
}

func (s *GetOrderConfirmUrlResponseBody) SetErrorCode(v string) *GetOrderConfirmUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrderConfirmUrlResponseBody) SetErrorMessage(v string) *GetOrderConfirmUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrderConfirmUrlResponseBody) SetRequestId(v string) *GetOrderConfirmUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderConfirmUrlResponseBody) SetSuccess(v bool) *GetOrderConfirmUrlResponseBody {
	s.Success = &v
	return s
}

type GetOrderConfirmUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrderConfirmUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderConfirmUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderConfirmUrlResponse) GoString() string {
	return s.String()
}

func (s *GetOrderConfirmUrlResponse) SetHeaders(v map[string]*string) *GetOrderConfirmUrlResponse {
	s.Headers = v
	return s
}

func (s *GetOrderConfirmUrlResponse) SetStatusCode(v int32) *GetOrderConfirmUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderConfirmUrlResponse) SetBody(v *GetOrderConfirmUrlResponseBody) *GetOrderConfirmUrlResponse {
	s.Body = v
	return s
}

type GetStsByTaobaoUidRequest struct {
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	TbUid     *string `json:"TbUid,omitempty" xml:"TbUid,omitempty"`
}

func (s GetStsByTaobaoUidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStsByTaobaoUidRequest) GoString() string {
	return s.String()
}

func (s *GetStsByTaobaoUidRequest) SetAliyunUid(v string) *GetStsByTaobaoUidRequest {
	s.AliyunUid = &v
	return s
}

func (s *GetStsByTaobaoUidRequest) SetTbUid(v string) *GetStsByTaobaoUidRequest {
	s.TbUid = &v
	return s
}

type GetStsByTaobaoUidResponseBody struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStsByTaobaoUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStsByTaobaoUidResponseBody) GoString() string {
	return s.String()
}

func (s *GetStsByTaobaoUidResponseBody) SetAccessKeyId(v string) *GetStsByTaobaoUidResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *GetStsByTaobaoUidResponseBody) SetAccessKeySecret(v string) *GetStsByTaobaoUidResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *GetStsByTaobaoUidResponseBody) SetErrorCode(v string) *GetStsByTaobaoUidResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStsByTaobaoUidResponseBody) SetErrorMessage(v string) *GetStsByTaobaoUidResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStsByTaobaoUidResponseBody) SetExpiration(v string) *GetStsByTaobaoUidResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetStsByTaobaoUidResponseBody) SetRequestId(v string) *GetStsByTaobaoUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStsByTaobaoUidResponseBody) SetSecurityToken(v string) *GetStsByTaobaoUidResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *GetStsByTaobaoUidResponseBody) SetSuccess(v bool) *GetStsByTaobaoUidResponseBody {
	s.Success = &v
	return s
}

type GetStsByTaobaoUidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStsByTaobaoUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStsByTaobaoUidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStsByTaobaoUidResponse) GoString() string {
	return s.String()
}

func (s *GetStsByTaobaoUidResponse) SetHeaders(v map[string]*string) *GetStsByTaobaoUidResponse {
	s.Headers = v
	return s
}

func (s *GetStsByTaobaoUidResponse) SetStatusCode(v int32) *GetStsByTaobaoUidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStsByTaobaoUidResponse) SetBody(v *GetStsByTaobaoUidResponseBody) *GetStsByTaobaoUidResponse {
	s.Body = v
	return s
}

type ListAdminTrademarkApplicationLogsRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s ListAdminTrademarkApplicationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationLogsRequest) SetBizId(v string) *ListAdminTrademarkApplicationLogsRequest {
	s.BizId = &v
	return s
}

type ListAdminTrademarkApplicationLogsResponseBody struct {
	RequestId                *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrademarkApplicationLogs []*ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs `json:"TrademarkApplicationLogs,omitempty" xml:"TrademarkApplicationLogs,omitempty" type:"Repeated"`
}

func (s ListAdminTrademarkApplicationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationLogsResponseBody) SetRequestId(v string) *ListAdminTrademarkApplicationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdminTrademarkApplicationLogsResponseBody) SetTrademarkApplicationLogs(v []*ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) *ListAdminTrademarkApplicationLogsResponseBody {
	s.TrademarkApplicationLogs = v
	return s
}

type ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizStatus   *int32  `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OperateTime *int64  `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OperateType *int32  `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) SetBizId(v string) *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs {
	s.BizId = &v
	return s
}

func (s *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) SetBizStatus(v int32) *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs {
	s.BizStatus = &v
	return s
}

func (s *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) SetNote(v string) *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs {
	s.Note = &v
	return s
}

func (s *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) SetOperateTime(v int64) *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs {
	s.OperateTime = &v
	return s
}

func (s *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) SetOperateType(v int32) *ListAdminTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs {
	s.OperateType = &v
	return s
}

type ListAdminTrademarkApplicationLogsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdminTrademarkApplicationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdminTrademarkApplicationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationLogsResponse) SetHeaders(v map[string]*string) *ListAdminTrademarkApplicationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListAdminTrademarkApplicationLogsResponse) SetStatusCode(v int32) *ListAdminTrademarkApplicationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdminTrademarkApplicationLogsResponse) SetBody(v *ListAdminTrademarkApplicationLogsResponseBody) *ListAdminTrademarkApplicationLogsResponse {
	s.Body = v
	return s
}

type ListAdminTrademarkApplicationsRequest struct {
	ApplicantName     *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicationStatus *int32  `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	ApplicationType   *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	BizId             *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortOrder         *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	SupplementStatus  *int32  `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TrademarkName     *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNumber   *string `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAdminTrademarkApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationsRequest) SetApplicantName(v string) *ListAdminTrademarkApplicationsRequest {
	s.ApplicantName = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetApplicationStatus(v int32) *ListAdminTrademarkApplicationsRequest {
	s.ApplicationStatus = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetApplicationType(v string) *ListAdminTrademarkApplicationsRequest {
	s.ApplicationType = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetBizId(v string) *ListAdminTrademarkApplicationsRequest {
	s.BizId = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetOrderId(v string) *ListAdminTrademarkApplicationsRequest {
	s.OrderId = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetPageNumber(v int32) *ListAdminTrademarkApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetPageSize(v int32) *ListAdminTrademarkApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetSortOrder(v string) *ListAdminTrademarkApplicationsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetSupplementStatus(v int32) *ListAdminTrademarkApplicationsRequest {
	s.SupplementStatus = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetTrademarkName(v string) *ListAdminTrademarkApplicationsRequest {
	s.TrademarkName = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetTrademarkNumber(v string) *ListAdminTrademarkApplicationsRequest {
	s.TrademarkNumber = &v
	return s
}

func (s *ListAdminTrademarkApplicationsRequest) SetUserId(v string) *ListAdminTrademarkApplicationsRequest {
	s.UserId = &v
	return s
}

type ListAdminTrademarkApplicationsResponseBody struct {
	PageNumber            *int32                                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrademarkApplications []*ListAdminTrademarkApplicationsResponseBodyTrademarkApplications `json:"TrademarkApplications,omitempty" xml:"TrademarkApplications,omitempty" type:"Repeated"`
}

func (s ListAdminTrademarkApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationsResponseBody) SetPageNumber(v int32) *ListAdminTrademarkApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBody) SetPageSize(v int32) *ListAdminTrademarkApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBody) SetRequestId(v string) *ListAdminTrademarkApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBody) SetTotalCount(v int32) *ListAdminTrademarkApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBody) SetTrademarkApplications(v []*ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) *ListAdminTrademarkApplicationsResponseBody {
	s.TrademarkApplications = v
	return s
}

type ListAdminTrademarkApplicationsResponseBodyTrademarkApplications struct {
	ApplicantId         *int64                                                                                `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicantName       *string                                                                               `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicationStatus   *int32                                                                                `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	ApplicationType     *int32                                                                                `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	AuthorizationUrl    *string                                                                               `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BizId               *string                                                                               `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CreateTime          *int64                                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FirstClassification *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification   `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
	Flags               []*string                                                                             `json:"Flags,omitempty" xml:"Flags,omitempty" type:"Repeated"`
	Note                *string                                                                               `json:"Note,omitempty" xml:"Note,omitempty"`
	OrderId             *string                                                                               `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderPrice          *float32                                                                              `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	PrincipalName       *int32                                                                                `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	ServicePrice        *float32                                                                              `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	SupplementId        *int64                                                                                `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	SupplementStatus    *int32                                                                                `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	SystemVersion       *string                                                                               `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	ThirdClassification []*ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Repeated"`
	TotalPrice          *float32                                                                              `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	TrademarkIcon       *string                                                                               `json:"TrademarkIcon,omitempty" xml:"TrademarkIcon,omitempty"`
	TrademarkName       *string                                                                               `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNumber     *string                                                                               `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UpdateTime          *int64                                                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId              *string                                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetApplicantId(v int64) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.ApplicantId = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetApplicantName(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.ApplicantName = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetApplicationStatus(v int32) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.ApplicationStatus = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetApplicationType(v int32) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.ApplicationType = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetAuthorizationUrl(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.AuthorizationUrl = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetBizId(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.BizId = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetCreateTime(v int64) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.CreateTime = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetFirstClassification(v *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.FirstClassification = v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetFlags(v []*string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.Flags = v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetNote(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.Note = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetOrderId(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.OrderId = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetOrderPrice(v float32) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.OrderPrice = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetPrincipalName(v int32) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.PrincipalName = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetServicePrice(v float32) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.ServicePrice = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetSupplementId(v int64) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.SupplementId = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetSupplementStatus(v int32) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.SupplementStatus = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetSystemVersion(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.SystemVersion = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetThirdClassification(v []*ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.ThirdClassification = v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetTotalPrice(v float32) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.TotalPrice = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetTrademarkIcon(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.TrademarkIcon = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetTrademarkName(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.TrademarkName = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetTrademarkNumber(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.TrademarkNumber = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetUpdateTime(v int64) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.UpdateTime = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications) SetUserId(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplications {
	s.UserId = &v
	return s
}

type ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification) SetClassificationCode(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification {
	s.ClassificationCode = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification) SetClassificationName(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsFirstClassification {
	s.ClassificationName = &v
	return s
}

type ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification) SetClassificationCode(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification {
	s.ClassificationCode = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification) SetClassificationName(v string) *ListAdminTrademarkApplicationsResponseBodyTrademarkApplicationsThirdClassification {
	s.ClassificationName = &v
	return s
}

type ListAdminTrademarkApplicationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAdminTrademarkApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAdminTrademarkApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAdminTrademarkApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListAdminTrademarkApplicationsResponse) SetHeaders(v map[string]*string) *ListAdminTrademarkApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListAdminTrademarkApplicationsResponse) SetStatusCode(v int32) *ListAdminTrademarkApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAdminTrademarkApplicationsResponse) SetBody(v *ListAdminTrademarkApplicationsResponseBody) *ListAdminTrademarkApplicationsResponse {
	s.Body = v
	return s
}

type ListApplicantsRequest struct {
	ApplicantName    *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantRegion  *int32  `json:"ApplicantRegion,omitempty" xml:"ApplicantRegion,omitempty"`
	ApplicantType    *int32  `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	ApplicantVersion *string `json:"ApplicantVersion,omitempty" xml:"ApplicantVersion,omitempty"`
	AuditStatus      *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	CardNumber       *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrincipalName    *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	SystemVersion    *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
}

func (s ListApplicantsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicantsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicantsRequest) SetApplicantName(v string) *ListApplicantsRequest {
	s.ApplicantName = &v
	return s
}

func (s *ListApplicantsRequest) SetApplicantRegion(v int32) *ListApplicantsRequest {
	s.ApplicantRegion = &v
	return s
}

func (s *ListApplicantsRequest) SetApplicantType(v int32) *ListApplicantsRequest {
	s.ApplicantType = &v
	return s
}

func (s *ListApplicantsRequest) SetApplicantVersion(v string) *ListApplicantsRequest {
	s.ApplicantVersion = &v
	return s
}

func (s *ListApplicantsRequest) SetAuditStatus(v int32) *ListApplicantsRequest {
	s.AuditStatus = &v
	return s
}

func (s *ListApplicantsRequest) SetCardNumber(v string) *ListApplicantsRequest {
	s.CardNumber = &v
	return s
}

func (s *ListApplicantsRequest) SetPageNumber(v int32) *ListApplicantsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicantsRequest) SetPageSize(v int32) *ListApplicantsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicantsRequest) SetPrincipalName(v int32) *ListApplicantsRequest {
	s.PrincipalName = &v
	return s
}

func (s *ListApplicantsRequest) SetSystemVersion(v string) *ListApplicantsRequest {
	s.SystemVersion = &v
	return s
}

type ListApplicantsResponseBody struct {
	Applicants *ListApplicantsResponseBodyApplicants `json:"Applicants,omitempty" xml:"Applicants,omitempty" type:"Struct"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicantsResponseBody) SetApplicants(v *ListApplicantsResponseBodyApplicants) *ListApplicantsResponseBody {
	s.Applicants = v
	return s
}

func (s *ListApplicantsResponseBody) SetPageNumber(v int32) *ListApplicantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicantsResponseBody) SetPageSize(v int32) *ListApplicantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicantsResponseBody) SetRequestId(v string) *ListApplicantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicantsResponseBody) SetTotalCount(v int32) *ListApplicantsResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicantsResponseBodyApplicants struct {
	Applicant []*ListApplicantsResponseBodyApplicantsApplicant `json:"Applicant,omitempty" xml:"Applicant,omitempty" type:"Repeated"`
}

func (s ListApplicantsResponseBodyApplicants) String() string {
	return tea.Prettify(s)
}

func (s ListApplicantsResponseBodyApplicants) GoString() string {
	return s.String()
}

func (s *ListApplicantsResponseBodyApplicants) SetApplicant(v []*ListApplicantsResponseBodyApplicantsApplicant) *ListApplicantsResponseBodyApplicants {
	s.Applicant = v
	return s
}

type ListApplicantsResponseBodyApplicantsApplicant struct {
	ApplicantId              *int64  `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicantName            *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantRegion          *int32  `json:"ApplicantRegion,omitempty" xml:"ApplicantRegion,omitempty"`
	ApplicantType            *int32  `json:"ApplicantType,omitempty" xml:"ApplicantType,omitempty"`
	ApplicantVersion         *string `json:"ApplicantVersion,omitempty" xml:"ApplicantVersion,omitempty"`
	AuditStatus              *int32  `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	AuthorizationAuditStatus *int32  `json:"AuthorizationAuditStatus,omitempty" xml:"AuthorizationAuditStatus,omitempty"`
	AuthorizationUrl         *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	CardNumber               *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ContactName              *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	PrincipalName            *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	SystemVersion            *string `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	ValidDate                *int64  `json:"ValidDate,omitempty" xml:"ValidDate,omitempty"`
}

func (s ListApplicantsResponseBodyApplicantsApplicant) String() string {
	return tea.Prettify(s)
}

func (s ListApplicantsResponseBodyApplicantsApplicant) GoString() string {
	return s.String()
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetApplicantId(v int64) *ListApplicantsResponseBodyApplicantsApplicant {
	s.ApplicantId = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetApplicantName(v string) *ListApplicantsResponseBodyApplicantsApplicant {
	s.ApplicantName = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetApplicantRegion(v int32) *ListApplicantsResponseBodyApplicantsApplicant {
	s.ApplicantRegion = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetApplicantType(v int32) *ListApplicantsResponseBodyApplicantsApplicant {
	s.ApplicantType = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetApplicantVersion(v string) *ListApplicantsResponseBodyApplicantsApplicant {
	s.ApplicantVersion = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetAuditStatus(v int32) *ListApplicantsResponseBodyApplicantsApplicant {
	s.AuditStatus = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetAuthorizationAuditStatus(v int32) *ListApplicantsResponseBodyApplicantsApplicant {
	s.AuthorizationAuditStatus = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetAuthorizationUrl(v string) *ListApplicantsResponseBodyApplicantsApplicant {
	s.AuthorizationUrl = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetCardNumber(v string) *ListApplicantsResponseBodyApplicantsApplicant {
	s.CardNumber = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetContactName(v string) *ListApplicantsResponseBodyApplicantsApplicant {
	s.ContactName = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetPrincipalName(v int32) *ListApplicantsResponseBodyApplicantsApplicant {
	s.PrincipalName = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetSystemVersion(v string) *ListApplicantsResponseBodyApplicantsApplicant {
	s.SystemVersion = &v
	return s
}

func (s *ListApplicantsResponseBodyApplicantsApplicant) SetValidDate(v int64) *ListApplicantsResponseBodyApplicantsApplicant {
	s.ValidDate = &v
	return s
}

type ListApplicantsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicantsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicantsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicantsResponse) SetHeaders(v map[string]*string) *ListApplicantsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicantsResponse) SetStatusCode(v int32) *ListApplicantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicantsResponse) SetBody(v *ListApplicantsResponseBody) *ListApplicantsResponse {
	s.Body = v
	return s
}

type ListAreasRequest struct {
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ParentCode *string `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
}

func (s ListAreasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAreasRequest) GoString() string {
	return s.String()
}

func (s *ListAreasRequest) SetBizType(v string) *ListAreasRequest {
	s.BizType = &v
	return s
}

func (s *ListAreasRequest) SetParentCode(v string) *ListAreasRequest {
	s.ParentCode = &v
	return s
}

type ListAreasResponseBody struct {
	Datas     []*ListAreasResponseBodyDatas `json:"Datas,omitempty" xml:"Datas,omitempty" type:"Repeated"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAreasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAreasResponseBody) GoString() string {
	return s.String()
}

func (s *ListAreasResponseBody) SetDatas(v []*ListAreasResponseBodyDatas) *ListAreasResponseBody {
	s.Datas = v
	return s
}

func (s *ListAreasResponseBody) SetRequestId(v string) *ListAreasResponseBody {
	s.RequestId = &v
	return s
}

type ListAreasResponseBodyDatas struct {
	Code       *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Name       *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCode *string                              `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
	SortNum    *int32                               `json:"SortNum,omitempty" xml:"SortNum,omitempty"`
	SubArea    []*ListAreasResponseBodyDatasSubArea `json:"SubArea,omitempty" xml:"SubArea,omitempty" type:"Repeated"`
}

func (s ListAreasResponseBodyDatas) String() string {
	return tea.Prettify(s)
}

func (s ListAreasResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *ListAreasResponseBodyDatas) SetCode(v string) *ListAreasResponseBodyDatas {
	s.Code = &v
	return s
}

func (s *ListAreasResponseBodyDatas) SetName(v string) *ListAreasResponseBodyDatas {
	s.Name = &v
	return s
}

func (s *ListAreasResponseBodyDatas) SetParentCode(v string) *ListAreasResponseBodyDatas {
	s.ParentCode = &v
	return s
}

func (s *ListAreasResponseBodyDatas) SetSortNum(v int32) *ListAreasResponseBodyDatas {
	s.SortNum = &v
	return s
}

func (s *ListAreasResponseBodyDatas) SetSubArea(v []*ListAreasResponseBodyDatasSubArea) *ListAreasResponseBodyDatas {
	s.SubArea = v
	return s
}

type ListAreasResponseBodyDatasSubArea struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentCode *string `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
	SortNum    *int32  `json:"SortNum,omitempty" xml:"SortNum,omitempty"`
}

func (s ListAreasResponseBodyDatasSubArea) String() string {
	return tea.Prettify(s)
}

func (s ListAreasResponseBodyDatasSubArea) GoString() string {
	return s.String()
}

func (s *ListAreasResponseBodyDatasSubArea) SetCode(v string) *ListAreasResponseBodyDatasSubArea {
	s.Code = &v
	return s
}

func (s *ListAreasResponseBodyDatasSubArea) SetName(v string) *ListAreasResponseBodyDatasSubArea {
	s.Name = &v
	return s
}

func (s *ListAreasResponseBodyDatasSubArea) SetParentCode(v string) *ListAreasResponseBodyDatasSubArea {
	s.ParentCode = &v
	return s
}

func (s *ListAreasResponseBodyDatasSubArea) SetSortNum(v int32) *ListAreasResponseBodyDatasSubArea {
	s.SortNum = &v
	return s
}

type ListAreasResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAreasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAreasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAreasResponse) GoString() string {
	return s.String()
}

func (s *ListAreasResponse) SetHeaders(v map[string]*string) *ListAreasResponse {
	s.Headers = v
	return s
}

func (s *ListAreasResponse) SetStatusCode(v int32) *ListAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAreasResponse) SetBody(v *ListAreasResponseBody) *ListAreasResponse {
	s.Body = v
	return s
}

type ListClassificationConditionsRequest struct {
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Type    *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListClassificationConditionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationConditionsRequest) GoString() string {
	return s.String()
}

func (s *ListClassificationConditionsRequest) SetTagName(v string) *ListClassificationConditionsRequest {
	s.TagName = &v
	return s
}

func (s *ListClassificationConditionsRequest) SetType(v int64) *ListClassificationConditionsRequest {
	s.Type = &v
	return s
}

type ListClassificationConditionsResponseBody struct {
	Data      []*ListClassificationConditionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClassificationConditionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationConditionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClassificationConditionsResponseBody) SetData(v []*ListClassificationConditionsResponseBodyData) *ListClassificationConditionsResponseBody {
	s.Data = v
	return s
}

func (s *ListClassificationConditionsResponseBody) SetRequestId(v string) *ListClassificationConditionsResponseBody {
	s.RequestId = &v
	return s
}

type ListClassificationConditionsResponseBodyData struct {
	ConditionContent *string `json:"ConditionContent,omitempty" xml:"ConditionContent,omitempty"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	SessionId        *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TagName          *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Type             *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
	Umid             *string `json:"Umid,omitempty" xml:"Umid,omitempty"`
	UpdateTime       *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId           *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListClassificationConditionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationConditionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClassificationConditionsResponseBodyData) SetConditionContent(v string) *ListClassificationConditionsResponseBodyData {
	s.ConditionContent = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetCreateTime(v int64) *ListClassificationConditionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetId(v int64) *ListClassificationConditionsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetSessionId(v string) *ListClassificationConditionsResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetTagName(v string) *ListClassificationConditionsResponseBodyData {
	s.TagName = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetType(v int64) *ListClassificationConditionsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetUmid(v string) *ListClassificationConditionsResponseBodyData {
	s.Umid = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetUpdateTime(v int64) *ListClassificationConditionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListClassificationConditionsResponseBodyData) SetUserId(v int64) *ListClassificationConditionsResponseBodyData {
	s.UserId = &v
	return s
}

type ListClassificationConditionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClassificationConditionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClassificationConditionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationConditionsResponse) GoString() string {
	return s.String()
}

func (s *ListClassificationConditionsResponse) SetHeaders(v map[string]*string) *ListClassificationConditionsResponse {
	s.Headers = v
	return s
}

func (s *ListClassificationConditionsResponse) SetStatusCode(v int32) *ListClassificationConditionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClassificationConditionsResponse) SetBody(v *ListClassificationConditionsResponseBody) *ListClassificationConditionsResponse {
	s.Body = v
	return s
}

type ListClassificationsRequest struct {
	ParentCode *string `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
}

func (s ListClassificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationsRequest) GoString() string {
	return s.String()
}

func (s *ListClassificationsRequest) SetParentCode(v string) *ListClassificationsRequest {
	s.ParentCode = &v
	return s
}

type ListClassificationsResponseBody struct {
	Classifications *ListClassificationsResponseBodyClassifications `json:"Classifications,omitempty" xml:"Classifications,omitempty" type:"Struct"`
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClassificationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClassificationsResponseBody) SetClassifications(v *ListClassificationsResponseBodyClassifications) *ListClassificationsResponseBody {
	s.Classifications = v
	return s
}

func (s *ListClassificationsResponseBody) SetRequestId(v string) *ListClassificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClassificationsResponseBody) SetTotalCount(v int32) *ListClassificationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListClassificationsResponseBodyClassifications struct {
	Classification []*ListClassificationsResponseBodyClassificationsClassification `json:"Classification,omitempty" xml:"Classification,omitempty" type:"Repeated"`
}

func (s ListClassificationsResponseBodyClassifications) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationsResponseBodyClassifications) GoString() string {
	return s.String()
}

func (s *ListClassificationsResponseBodyClassifications) SetClassification(v []*ListClassificationsResponseBodyClassificationsClassification) *ListClassificationsResponseBodyClassifications {
	s.Classification = v
	return s
}

type ListClassificationsResponseBodyClassificationsClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Level              *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	OfficialCode       *string `json:"OfficialCode,omitempty" xml:"OfficialCode,omitempty"`
	ParentCode         *string `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
}

func (s ListClassificationsResponseBodyClassificationsClassification) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationsResponseBodyClassificationsClassification) GoString() string {
	return s.String()
}

func (s *ListClassificationsResponseBodyClassificationsClassification) SetClassificationCode(v string) *ListClassificationsResponseBodyClassificationsClassification {
	s.ClassificationCode = &v
	return s
}

func (s *ListClassificationsResponseBodyClassificationsClassification) SetClassificationName(v string) *ListClassificationsResponseBodyClassificationsClassification {
	s.ClassificationName = &v
	return s
}

func (s *ListClassificationsResponseBodyClassificationsClassification) SetId(v int64) *ListClassificationsResponseBodyClassificationsClassification {
	s.Id = &v
	return s
}

func (s *ListClassificationsResponseBodyClassificationsClassification) SetLevel(v int32) *ListClassificationsResponseBodyClassificationsClassification {
	s.Level = &v
	return s
}

func (s *ListClassificationsResponseBodyClassificationsClassification) SetOfficialCode(v string) *ListClassificationsResponseBodyClassificationsClassification {
	s.OfficialCode = &v
	return s
}

func (s *ListClassificationsResponseBodyClassificationsClassification) SetParentCode(v string) *ListClassificationsResponseBodyClassificationsClassification {
	s.ParentCode = &v
	return s
}

type ListClassificationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClassificationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClassificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClassificationsResponse) GoString() string {
	return s.String()
}

func (s *ListClassificationsResponse) SetHeaders(v map[string]*string) *ListClassificationsResponse {
	s.Headers = v
	return s
}

func (s *ListClassificationsResponse) SetStatusCode(v int32) *ListClassificationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClassificationsResponse) SetBody(v *ListClassificationsResponseBody) *ListClassificationsResponse {
	s.Body = v
	return s
}

type ListTrademarkApplicationLogsRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s ListTrademarkApplicationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationLogsRequest) SetBizId(v string) *ListTrademarkApplicationLogsRequest {
	s.BizId = &v
	return s
}

type ListTrademarkApplicationLogsResponseBody struct {
	RequestId                *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrademarkApplicationLogs *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs `json:"TrademarkApplicationLogs,omitempty" xml:"TrademarkApplicationLogs,omitempty" type:"Struct"`
}

func (s ListTrademarkApplicationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationLogsResponseBody) SetRequestId(v string) *ListTrademarkApplicationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrademarkApplicationLogsResponseBody) SetTrademarkApplicationLogs(v *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) *ListTrademarkApplicationLogsResponseBody {
	s.TrademarkApplicationLogs = v
	return s
}

type ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs struct {
	TrademarkApplicationLog []*ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog `json:"TrademarkApplicationLog,omitempty" xml:"TrademarkApplicationLog,omitempty" type:"Repeated"`
}

func (s ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs) SetTrademarkApplicationLog(v []*ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogs {
	s.TrademarkApplicationLog = v
	return s
}

type ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizStatus   *int32  `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OperateTime *int64  `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	OperateType *int32  `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) SetBizId(v string) *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog {
	s.BizId = &v
	return s
}

func (s *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) SetBizStatus(v int32) *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog {
	s.BizStatus = &v
	return s
}

func (s *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) SetNote(v string) *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog {
	s.Note = &v
	return s
}

func (s *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) SetOperateTime(v int64) *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog {
	s.OperateTime = &v
	return s
}

func (s *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog) SetOperateType(v int32) *ListTrademarkApplicationLogsResponseBodyTrademarkApplicationLogsTrademarkApplicationLog {
	s.OperateType = &v
	return s
}

type ListTrademarkApplicationLogsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrademarkApplicationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrademarkApplicationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationLogsResponse) SetHeaders(v map[string]*string) *ListTrademarkApplicationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListTrademarkApplicationLogsResponse) SetStatusCode(v int32) *ListTrademarkApplicationLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrademarkApplicationLogsResponse) SetBody(v *ListTrademarkApplicationLogsResponseBody) *ListTrademarkApplicationLogsResponse {
	s.Body = v
	return s
}

type ListTrademarkApplicationsRequest struct {
	ApplicantName             *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicationStatus         *int32  `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	ApplicationType           *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	BizId                     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CreateTimeLeft            *int64  `json:"CreateTimeLeft,omitempty" xml:"CreateTimeLeft,omitempty"`
	CreateTimeRight           *int64  `json:"CreateTimeRight,omitempty" xml:"CreateTimeRight,omitempty"`
	Flag                      *int32  `json:"Flag,omitempty" xml:"Flag,omitempty"`
	OrderId                   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber                *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType               *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	QueryVoucherOrderDoneFlag *bool   `json:"QueryVoucherOrderDoneFlag,omitempty" xml:"QueryVoucherOrderDoneFlag,omitempty"`
	QueryVoucherOrderFlag     *bool   `json:"QueryVoucherOrderFlag,omitempty" xml:"QueryVoucherOrderFlag,omitempty"`
	SortFiled                 *string `json:"SortFiled,omitempty" xml:"SortFiled,omitempty"`
	SortOrder                 *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	SupplementStatus          *int32  `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	TrademarkName             *string `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNumber           *string `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
}

func (s ListTrademarkApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsRequest) SetApplicantName(v string) *ListTrademarkApplicationsRequest {
	s.ApplicantName = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetApplicationStatus(v int32) *ListTrademarkApplicationsRequest {
	s.ApplicationStatus = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetApplicationType(v string) *ListTrademarkApplicationsRequest {
	s.ApplicationType = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetBizId(v string) *ListTrademarkApplicationsRequest {
	s.BizId = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetCreateTimeLeft(v int64) *ListTrademarkApplicationsRequest {
	s.CreateTimeLeft = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetCreateTimeRight(v int64) *ListTrademarkApplicationsRequest {
	s.CreateTimeRight = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetFlag(v int32) *ListTrademarkApplicationsRequest {
	s.Flag = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetOrderId(v string) *ListTrademarkApplicationsRequest {
	s.OrderId = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetPageNumber(v int32) *ListTrademarkApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetPageSize(v int32) *ListTrademarkApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetProductType(v int32) *ListTrademarkApplicationsRequest {
	s.ProductType = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetQueryVoucherOrderDoneFlag(v bool) *ListTrademarkApplicationsRequest {
	s.QueryVoucherOrderDoneFlag = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetQueryVoucherOrderFlag(v bool) *ListTrademarkApplicationsRequest {
	s.QueryVoucherOrderFlag = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetSortFiled(v string) *ListTrademarkApplicationsRequest {
	s.SortFiled = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetSortOrder(v string) *ListTrademarkApplicationsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetSupplementStatus(v int32) *ListTrademarkApplicationsRequest {
	s.SupplementStatus = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetTrademarkName(v string) *ListTrademarkApplicationsRequest {
	s.TrademarkName = &v
	return s
}

func (s *ListTrademarkApplicationsRequest) SetTrademarkNumber(v string) *ListTrademarkApplicationsRequest {
	s.TrademarkNumber = &v
	return s
}

type ListTrademarkApplicationsResponseBody struct {
	PageNumber            *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrademarkApplications *ListTrademarkApplicationsResponseBodyTrademarkApplications `json:"TrademarkApplications,omitempty" xml:"TrademarkApplications,omitempty" type:"Struct"`
}

func (s ListTrademarkApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponseBody) SetPageNumber(v int32) *ListTrademarkApplicationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBody) SetPageSize(v int32) *ListTrademarkApplicationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBody) SetRequestId(v string) *ListTrademarkApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBody) SetTotalCount(v int32) *ListTrademarkApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBody) SetTrademarkApplications(v *ListTrademarkApplicationsResponseBodyTrademarkApplications) *ListTrademarkApplicationsResponseBody {
	s.TrademarkApplications = v
	return s
}

type ListTrademarkApplicationsResponseBodyTrademarkApplications struct {
	TrademarkApplication []*ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication `json:"TrademarkApplication,omitempty" xml:"TrademarkApplication,omitempty" type:"Repeated"`
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplications) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplications) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplications) SetTrademarkApplication(v []*ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) *ListTrademarkApplicationsResponseBodyTrademarkApplications {
	s.TrademarkApplication = v
	return s
}

type ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication struct {
	AgreementId         *string                                                                                            `json:"AgreementId,omitempty" xml:"AgreementId,omitempty"`
	ApplicantId         *int64                                                                                             `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicantName       *string                                                                                            `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicationStatus   *int32                                                                                             `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	ApplicationType     *int32                                                                                             `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	AuthorizationUrl    *string                                                                                            `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	BizId               *string                                                                                            `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CreateTime          *int64                                                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FirstClassification *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification `json:"FirstClassification,omitempty" xml:"FirstClassification,omitempty" type:"Struct"`
	Flags               *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFlags               `json:"Flags,omitempty" xml:"Flags,omitempty" type:"Struct"`
	Note                *string                                                                                            `json:"Note,omitempty" xml:"Note,omitempty"`
	OrderId             *string                                                                                            `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderPrice          *float32                                                                                           `json:"OrderPrice,omitempty" xml:"OrderPrice,omitempty"`
	PrincipalName       *int32                                                                                             `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	ServicePrice        *float32                                                                                           `json:"ServicePrice,omitempty" xml:"ServicePrice,omitempty"`
	SupplementId        *int64                                                                                             `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	SupplementStatus    *int32                                                                                             `json:"SupplementStatus,omitempty" xml:"SupplementStatus,omitempty"`
	SystemVersion       *string                                                                                            `json:"SystemVersion,omitempty" xml:"SystemVersion,omitempty"`
	ThirdClassification *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Struct"`
	TotalPrice          *float32                                                                                           `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
	TrademarkIcon       *string                                                                                            `json:"TrademarkIcon,omitempty" xml:"TrademarkIcon,omitempty"`
	TrademarkName       *string                                                                                            `json:"TrademarkName,omitempty" xml:"TrademarkName,omitempty"`
	TrademarkNumber     *string                                                                                            `json:"TrademarkNumber,omitempty" xml:"TrademarkNumber,omitempty"`
	UpdateTime          *int64                                                                                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId              *string                                                                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetAgreementId(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.AgreementId = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetApplicantId(v int64) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.ApplicantId = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetApplicantName(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.ApplicantName = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetApplicationStatus(v int32) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.ApplicationStatus = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetApplicationType(v int32) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.ApplicationType = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetAuthorizationUrl(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.AuthorizationUrl = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetBizId(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.BizId = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetCreateTime(v int64) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.CreateTime = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetFirstClassification(v *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.FirstClassification = v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetFlags(v *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFlags) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.Flags = v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetNote(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.Note = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetOrderId(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.OrderId = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetOrderPrice(v float32) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.OrderPrice = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetPrincipalName(v int32) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.PrincipalName = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetServicePrice(v float32) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.ServicePrice = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetSupplementId(v int64) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.SupplementId = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetSupplementStatus(v int32) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.SupplementStatus = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetSystemVersion(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.SystemVersion = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetThirdClassification(v *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassification) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.ThirdClassification = v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetTotalPrice(v float32) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.TotalPrice = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetTrademarkIcon(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.TrademarkIcon = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetTrademarkName(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.TrademarkName = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetTrademarkNumber(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.TrademarkNumber = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetUpdateTime(v int64) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.UpdateTime = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication) SetUserId(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplication {
	s.UserId = &v
	return s
}

type ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification) SetClassificationCode(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification {
	s.ClassificationCode = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification) SetClassificationName(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFirstClassification {
	s.ClassificationName = &v
	return s
}

type ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFlags struct {
	Flags []*string `json:"Flags,omitempty" xml:"Flags,omitempty" type:"Repeated"`
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFlags) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFlags) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFlags) SetFlags(v []*string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationFlags {
	s.Flags = v
	return s
}

type ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassification struct {
	ThirdClassification []*ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification `json:"ThirdClassification,omitempty" xml:"ThirdClassification,omitempty" type:"Repeated"`
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassification) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassification) SetThirdClassification(v []*ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassification {
	s.ThirdClassification = v
	return s
}

type ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification struct {
	ClassificationCode *string `json:"ClassificationCode,omitempty" xml:"ClassificationCode,omitempty"`
	ClassificationName *string `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification) SetClassificationCode(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification {
	s.ClassificationCode = &v
	return s
}

func (s *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification) SetClassificationName(v string) *ListTrademarkApplicationsResponseBodyTrademarkApplicationsTrademarkApplicationThirdClassificationThirdClassification {
	s.ClassificationName = &v
	return s
}

type ListTrademarkApplicationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrademarkApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrademarkApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListTrademarkApplicationsResponse) SetHeaders(v map[string]*string) *ListTrademarkApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListTrademarkApplicationsResponse) SetStatusCode(v int32) *ListTrademarkApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrademarkApplicationsResponse) SetBody(v *ListTrademarkApplicationsResponseBody) *ListTrademarkApplicationsResponse {
	s.Body = v
	return s
}

type ListTrademarkSearchForInnerRequest struct {
	ApplyBeginTime   *string `json:"ApplyBeginTime,omitempty" xml:"ApplyBeginTime,omitempty"`
	ApplyEndTime     *string `json:"ApplyEndTime,omitempty" xml:"ApplyEndTime,omitempty"`
	Classification   *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	IfPrecision      *bool   `json:"IfPrecision,omitempty" xml:"IfPrecision,omitempty"`
	Keyword          *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Product          *string `json:"Product,omitempty" xml:"Product,omitempty"`
	SearchPreference *string `json:"SearchPreference,omitempty" xml:"SearchPreference,omitempty"`
	SearchType       *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Umid             *string `json:"Umid,omitempty" xml:"Umid,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListTrademarkSearchForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkSearchForInnerRequest) GoString() string {
	return s.String()
}

func (s *ListTrademarkSearchForInnerRequest) SetApplyBeginTime(v string) *ListTrademarkSearchForInnerRequest {
	s.ApplyBeginTime = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetApplyEndTime(v string) *ListTrademarkSearchForInnerRequest {
	s.ApplyEndTime = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetClassification(v string) *ListTrademarkSearchForInnerRequest {
	s.Classification = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetIfPrecision(v bool) *ListTrademarkSearchForInnerRequest {
	s.IfPrecision = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetKeyword(v string) *ListTrademarkSearchForInnerRequest {
	s.Keyword = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetPageNumber(v int32) *ListTrademarkSearchForInnerRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetPageSize(v int32) *ListTrademarkSearchForInnerRequest {
	s.PageSize = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetProduct(v string) *ListTrademarkSearchForInnerRequest {
	s.Product = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetSearchPreference(v string) *ListTrademarkSearchForInnerRequest {
	s.SearchPreference = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetSearchType(v string) *ListTrademarkSearchForInnerRequest {
	s.SearchType = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetStatus(v string) *ListTrademarkSearchForInnerRequest {
	s.Status = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetUmid(v string) *ListTrademarkSearchForInnerRequest {
	s.Umid = &v
	return s
}

func (s *ListTrademarkSearchForInnerRequest) SetUserId(v string) *ListTrademarkSearchForInnerRequest {
	s.UserId = &v
	return s
}

type ListTrademarkSearchForInnerResponseBody struct {
	PageNumber              *string                                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                *string                                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Products                []*string                                                         `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	RequestId               *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount              *string                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrademarkSearchContents []*ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents `json:"TrademarkSearchContents,omitempty" xml:"TrademarkSearchContents,omitempty" type:"Repeated"`
}

func (s ListTrademarkSearchForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkSearchForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrademarkSearchForInnerResponseBody) SetPageNumber(v string) *ListTrademarkSearchForInnerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBody) SetPageSize(v string) *ListTrademarkSearchForInnerResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBody) SetProducts(v []*string) *ListTrademarkSearchForInnerResponseBody {
	s.Products = v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBody) SetRequestId(v string) *ListTrademarkSearchForInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBody) SetTotalCount(v string) *ListTrademarkSearchForInnerResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBody) SetTrademarkSearchContents(v []*ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) *ListTrademarkSearchForInnerResponseBody {
	s.TrademarkSearchContents = v
	return s
}

type ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents struct {
	ApplyDate             *string   `json:"ApplyDate,omitempty" xml:"ApplyDate,omitempty"`
	Classification        *string   `json:"Classification,omitempty" xml:"Classification,omitempty"`
	ExclusiveDateLimit    *string   `json:"ExclusiveDateLimit,omitempty" xml:"ExclusiveDateLimit,omitempty"`
	Id                    *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Image                 *string   `json:"Image,omitempty" xml:"Image,omitempty"`
	LastProcedureStatus   *string   `json:"LastProcedureStatus,omitempty" xml:"LastProcedureStatus,omitempty"`
	Name                  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NameCharSection       *string   `json:"NameCharSection,omitempty" xml:"NameCharSection,omitempty"`
	NameOrigin            *string   `json:"NameOrigin,omitempty" xml:"NameOrigin,omitempty"`
	NameSimplifiedChinese *string   `json:"NameSimplifiedChinese,omitempty" xml:"NameSimplifiedChinese,omitempty"`
	NameSort              *string   `json:"NameSort,omitempty" xml:"NameSort,omitempty"`
	OnSale                *string   `json:"OnSale,omitempty" xml:"OnSale,omitempty"`
	OwnerAddress          *string   `json:"OwnerAddress,omitempty" xml:"OwnerAddress,omitempty"`
	OwnerEnAddress        *string   `json:"OwnerEnAddress,omitempty" xml:"OwnerEnAddress,omitempty"`
	OwnerEnName           *string   `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	OwnerName             *string   `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	PreAnnDate            *string   `json:"PreAnnDate,omitempty" xml:"PreAnnDate,omitempty"`
	PreAnnNum             *string   `json:"PreAnnNum,omitempty" xml:"PreAnnNum,omitempty"`
	Product               *string   `json:"Product,omitempty" xml:"Product,omitempty"`
	ProductDel            []*string `json:"ProductDel,omitempty" xml:"ProductDel,omitempty" type:"Repeated"`
	ProductDescription    *string   `json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	RegAnnNum             *string   `json:"RegAnnNum,omitempty" xml:"RegAnnNum,omitempty"`
	RegistrationNumber    *string   `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	Share                 *string   `json:"Share,omitempty" xml:"Share,omitempty"`
	SimilarGroupDel       []*string `json:"SimilarGroupDel,omitempty" xml:"SimilarGroupDel,omitempty" type:"Repeated"`
	Uid                   *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
	WellKnow              *string   `json:"WellKnow,omitempty" xml:"WellKnow,omitempty"`
}

func (s ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) GoString() string {
	return s.String()
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetApplyDate(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.ApplyDate = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetClassification(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.Classification = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetExclusiveDateLimit(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.ExclusiveDateLimit = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetId(v int64) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.Id = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetImage(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.Image = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetLastProcedureStatus(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.LastProcedureStatus = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetName(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.Name = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetNameCharSection(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.NameCharSection = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetNameOrigin(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.NameOrigin = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetNameSimplifiedChinese(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.NameSimplifiedChinese = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetNameSort(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.NameSort = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetOnSale(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.OnSale = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetOwnerAddress(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.OwnerAddress = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetOwnerEnAddress(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.OwnerEnAddress = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetOwnerEnName(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.OwnerEnName = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetOwnerName(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.OwnerName = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetPreAnnDate(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.PreAnnDate = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetPreAnnNum(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.PreAnnNum = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetProduct(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.Product = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetProductDel(v []*string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.ProductDel = v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetProductDescription(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.ProductDescription = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetRegAnnNum(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.RegAnnNum = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetRegistrationNumber(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.RegistrationNumber = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetShare(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.Share = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetSimilarGroupDel(v []*string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.SimilarGroupDel = v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetUid(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.Uid = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents) SetWellKnow(v string) *ListTrademarkSearchForInnerResponseBodyTrademarkSearchContents {
	s.WellKnow = &v
	return s
}

type ListTrademarkSearchForInnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrademarkSearchForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrademarkSearchForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTrademarkSearchForInnerResponse) GoString() string {
	return s.String()
}

func (s *ListTrademarkSearchForInnerResponse) SetHeaders(v map[string]*string) *ListTrademarkSearchForInnerResponse {
	s.Headers = v
	return s
}

func (s *ListTrademarkSearchForInnerResponse) SetStatusCode(v int32) *ListTrademarkSearchForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrademarkSearchForInnerResponse) SetBody(v *ListTrademarkSearchForInnerResponseBody) *ListTrademarkSearchForInnerResponse {
	s.Body = v
	return s
}

type PutMeasureDataRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PutMeasureDataRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureDataRequest) SetBizType(v string) *PutMeasureDataRequest {
	s.BizType = &v
	return s
}

func (s *PutMeasureDataRequest) SetData(v string) *PutMeasureDataRequest {
	s.Data = &v
	return s
}

func (s *PutMeasureDataRequest) SetDataType(v string) *PutMeasureDataRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureDataRequest) SetEndTime(v string) *PutMeasureDataRequest {
	s.EndTime = &v
	return s
}

func (s *PutMeasureDataRequest) SetStartTime(v string) *PutMeasureDataRequest {
	s.StartTime = &v
	return s
}

type PutMeasureDataResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutMeasureDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponseBody) SetData(v bool) *PutMeasureDataResponseBody {
	s.Data = &v
	return s
}

func (s *PutMeasureDataResponseBody) SetRequestId(v string) *PutMeasureDataResponseBody {
	s.RequestId = &v
	return s
}

type PutMeasureDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMeasureDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMeasureDataResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureDataResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureDataResponse) SetHeaders(v map[string]*string) *PutMeasureDataResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureDataResponse) SetStatusCode(v int32) *PutMeasureDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureDataResponse) SetBody(v *PutMeasureDataResponseBody) *PutMeasureDataResponse {
	s.Body = v
	return s
}

type PutMeasureReadyFlagRequest struct {
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReadyFlag *string `json:"ReadyFlag,omitempty" xml:"ReadyFlag,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PutMeasureReadyFlagRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureReadyFlagRequest) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagRequest) SetBizType(v string) *PutMeasureReadyFlagRequest {
	s.BizType = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetDataType(v string) *PutMeasureReadyFlagRequest {
	s.DataType = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetEndTime(v string) *PutMeasureReadyFlagRequest {
	s.EndTime = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetReadyFlag(v string) *PutMeasureReadyFlagRequest {
	s.ReadyFlag = &v
	return s
}

func (s *PutMeasureReadyFlagRequest) SetStartTime(v string) *PutMeasureReadyFlagRequest {
	s.StartTime = &v
	return s
}

type PutMeasureReadyFlagResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutMeasureReadyFlagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureReadyFlagResponseBody) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagResponseBody) SetData(v bool) *PutMeasureReadyFlagResponseBody {
	s.Data = &v
	return s
}

func (s *PutMeasureReadyFlagResponseBody) SetRequestId(v string) *PutMeasureReadyFlagResponseBody {
	s.RequestId = &v
	return s
}

type PutMeasureReadyFlagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutMeasureReadyFlagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutMeasureReadyFlagResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMeasureReadyFlagResponse) GoString() string {
	return s.String()
}

func (s *PutMeasureReadyFlagResponse) SetHeaders(v map[string]*string) *PutMeasureReadyFlagResponse {
	s.Headers = v
	return s
}

func (s *PutMeasureReadyFlagResponse) SetStatusCode(v int32) *PutMeasureReadyFlagResponse {
	s.StatusCode = &v
	return s
}

func (s *PutMeasureReadyFlagResponse) SetBody(v *PutMeasureReadyFlagResponseBody) *PutMeasureReadyFlagResponse {
	s.Body = v
	return s
}

type QueryActivityItemsRequest struct {
	ActivityId *int32  `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FloorIndex *int32  `json:"FloorIndex,omitempty" xml:"FloorIndex,omitempty"`
	Mock       *bool   `json:"Mock,omitempty" xml:"Mock,omitempty"`
	PageIndex  *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Refresh    *bool   `json:"Refresh,omitempty" xml:"Refresh,omitempty"`
	UmId       *string `json:"UmId,omitempty" xml:"UmId,omitempty"`
}

func (s QueryActivityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsRequest) SetActivityId(v int32) *QueryActivityItemsRequest {
	s.ActivityId = &v
	return s
}

func (s *QueryActivityItemsRequest) SetExtendInfo(v string) *QueryActivityItemsRequest {
	s.ExtendInfo = &v
	return s
}

func (s *QueryActivityItemsRequest) SetFloorIndex(v int32) *QueryActivityItemsRequest {
	s.FloorIndex = &v
	return s
}

func (s *QueryActivityItemsRequest) SetMock(v bool) *QueryActivityItemsRequest {
	s.Mock = &v
	return s
}

func (s *QueryActivityItemsRequest) SetPageIndex(v int32) *QueryActivityItemsRequest {
	s.PageIndex = &v
	return s
}

func (s *QueryActivityItemsRequest) SetPageSize(v int32) *QueryActivityItemsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryActivityItemsRequest) SetRefresh(v bool) *QueryActivityItemsRequest {
	s.Refresh = &v
	return s
}

func (s *QueryActivityItemsRequest) SetUmId(v string) *QueryActivityItemsRequest {
	s.UmId = &v
	return s
}

type QueryActivityItemsResponseBody struct {
	Module *QueryActivityItemsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
}

func (s QueryActivityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBody) SetModule(v *QueryActivityItemsResponseBodyModule) *QueryActivityItemsResponseBody {
	s.Module = v
	return s
}

type QueryActivityItemsResponseBodyModule struct {
	Data              *string                                                `json:"Data,omitempty" xml:"Data,omitempty"`
	FloorDisplayInfos *QueryActivityItemsResponseBodyModuleFloorDisplayInfos `json:"FloorDisplayInfos,omitempty" xml:"FloorDisplayInfos,omitempty" type:"Struct"`
	FloorItems        *string                                                `json:"FloorItems,omitempty" xml:"FloorItems,omitempty"`
}

func (s QueryActivityItemsResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyModule) SetData(v string) *QueryActivityItemsResponseBodyModule {
	s.Data = &v
	return s
}

func (s *QueryActivityItemsResponseBodyModule) SetFloorDisplayInfos(v *QueryActivityItemsResponseBodyModuleFloorDisplayInfos) *QueryActivityItemsResponseBodyModule {
	s.FloorDisplayInfos = v
	return s
}

func (s *QueryActivityItemsResponseBodyModule) SetFloorItems(v string) *QueryActivityItemsResponseBodyModule {
	s.FloorItems = &v
	return s
}

type QueryActivityItemsResponseBodyModuleFloorDisplayInfos struct {
	Floor []*QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor `json:"floor,omitempty" xml:"floor,omitempty" type:"Repeated"`
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfos) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfos) SetFloor(v []*QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) *QueryActivityItemsResponseBodyModuleFloorDisplayInfos {
	s.Floor = v
	return s
}

type QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor struct {
	Icon      *string                                                              `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Index     *int32                                                               `json:"Index,omitempty" xml:"Index,omitempty"`
	Name      *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	SubTitles *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitles `json:"SubTitles,omitempty" xml:"SubTitles,omitempty" type:"Struct"`
	Title     *string                                                              `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) SetIcon(v string) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor {
	s.Icon = &v
	return s
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) SetIndex(v int32) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor {
	s.Index = &v
	return s
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) SetName(v string) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor {
	s.Name = &v
	return s
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) SetSubTitles(v *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitles) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor {
	s.SubTitles = v
	return s
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor) SetTitle(v string) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloor {
	s.Title = &v
	return s
}

type QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitles struct {
	SubFloor []*QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor `json:"subFloor,omitempty" xml:"subFloor,omitempty" type:"Repeated"`
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitles) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitles) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitles) SetSubFloor(v []*QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitles {
	s.SubFloor = v
	return s
}

type QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor struct {
	Icon  *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor) SetIcon(v string) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor {
	s.Icon = &v
	return s
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor) SetName(v string) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor {
	s.Name = &v
	return s
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor) SetTitle(v string) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor {
	s.Title = &v
	return s
}

func (s *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor) SetValue(v string) *QueryActivityItemsResponseBodyModuleFloorDisplayInfosFloorSubTitlesSubFloor {
	s.Value = &v
	return s
}

type QueryActivityItemsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryActivityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryActivityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponse) SetHeaders(v map[string]*string) *QueryActivityItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryActivityItemsResponse) SetStatusCode(v int32) *QueryActivityItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryActivityItemsResponse) SetBody(v *QueryActivityItemsResponseBody) *QueryActivityItemsResponse {
	s.Body = v
	return s
}

type QueryAliyunUidRequest struct {
	TbUid *string `json:"TbUid,omitempty" xml:"TbUid,omitempty"`
}

func (s QueryAliyunUidRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAliyunUidRequest) GoString() string {
	return s.String()
}

func (s *QueryAliyunUidRequest) SetTbUid(v string) *QueryAliyunUidRequest {
	s.TbUid = &v
	return s
}

type QueryAliyunUidResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryAliyunUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAliyunUidResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAliyunUidResponseBody) SetCode(v string) *QueryAliyunUidResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAliyunUidResponseBody) SetMessage(v string) *QueryAliyunUidResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAliyunUidResponseBody) SetRequestId(v string) *QueryAliyunUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAliyunUidResponseBody) SetSuccess(v bool) *QueryAliyunUidResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAliyunUidResponseBody) SetUid(v string) *QueryAliyunUidResponseBody {
	s.Uid = &v
	return s
}

type QueryAliyunUidResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAliyunUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAliyunUidResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAliyunUidResponse) GoString() string {
	return s.String()
}

func (s *QueryAliyunUidResponse) SetHeaders(v map[string]*string) *QueryAliyunUidResponse {
	s.Headers = v
	return s
}

func (s *QueryAliyunUidResponse) SetStatusCode(v int32) *QueryAliyunUidResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAliyunUidResponse) SetBody(v *QueryAliyunUidResponseBody) *QueryAliyunUidResponse {
	s.Body = v
	return s
}

type QueryDetailItemRequest struct {
	DetailConvertType *string `json:"DetailConvertType,omitempty" xml:"DetailConvertType,omitempty"`
	DetailId          *string `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
	DetailType        *string `json:"DetailType,omitempty" xml:"DetailType,omitempty"`
	Mock              *bool   `json:"Mock,omitempty" xml:"Mock,omitempty"`
}

func (s QueryDetailItemRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDetailItemRequest) GoString() string {
	return s.String()
}

func (s *QueryDetailItemRequest) SetDetailConvertType(v string) *QueryDetailItemRequest {
	s.DetailConvertType = &v
	return s
}

func (s *QueryDetailItemRequest) SetDetailId(v string) *QueryDetailItemRequest {
	s.DetailId = &v
	return s
}

func (s *QueryDetailItemRequest) SetDetailType(v string) *QueryDetailItemRequest {
	s.DetailType = &v
	return s
}

func (s *QueryDetailItemRequest) SetMock(v bool) *QueryDetailItemRequest {
	s.Mock = &v
	return s
}

type QueryDetailItemResponseBody struct {
	Module *QueryDetailItemResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
}

func (s QueryDetailItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDetailItemResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDetailItemResponseBody) SetModule(v *QueryDetailItemResponseBodyModule) *QueryDetailItemResponseBody {
	s.Module = v
	return s
}

type QueryDetailItemResponseBodyModule struct {
	Attributes        *QueryDetailItemResponseBodyModuleAttributes   `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Struct"`
	DetailDescription *string                                        `json:"DetailDescription,omitempty" xml:"DetailDescription,omitempty"`
	DetailObjectJson  *string                                        `json:"DetailObjectJson,omitempty" xml:"DetailObjectJson,omitempty"`
	DetailPagePicUrl  *string                                        `json:"DetailPagePicUrl,omitempty" xml:"DetailPagePicUrl,omitempty"`
	DetailPicUrl      *QueryDetailItemResponseBodyModuleDetailPicUrl `json:"DetailPicUrl,omitempty" xml:"DetailPicUrl,omitempty" type:"Struct"`
	Label             *string                                        `json:"Label,omitempty" xml:"Label,omitempty"`
	MainPicUrl        *string                                        `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	Title             *string                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	Type              *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDetailItemResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s QueryDetailItemResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryDetailItemResponseBodyModule) SetAttributes(v *QueryDetailItemResponseBodyModuleAttributes) *QueryDetailItemResponseBodyModule {
	s.Attributes = v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetDetailDescription(v string) *QueryDetailItemResponseBodyModule {
	s.DetailDescription = &v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetDetailObjectJson(v string) *QueryDetailItemResponseBodyModule {
	s.DetailObjectJson = &v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetDetailPagePicUrl(v string) *QueryDetailItemResponseBodyModule {
	s.DetailPagePicUrl = &v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetDetailPicUrl(v *QueryDetailItemResponseBodyModuleDetailPicUrl) *QueryDetailItemResponseBodyModule {
	s.DetailPicUrl = v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetLabel(v string) *QueryDetailItemResponseBodyModule {
	s.Label = &v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetMainPicUrl(v string) *QueryDetailItemResponseBodyModule {
	s.MainPicUrl = &v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetTitle(v string) *QueryDetailItemResponseBodyModule {
	s.Title = &v
	return s
}

func (s *QueryDetailItemResponseBodyModule) SetType(v string) *QueryDetailItemResponseBodyModule {
	s.Type = &v
	return s
}

type QueryDetailItemResponseBodyModuleAttributes struct {
	Attribute []*QueryDetailItemResponseBodyModuleAttributesAttribute `json:"attribute,omitempty" xml:"attribute,omitempty" type:"Repeated"`
}

func (s QueryDetailItemResponseBodyModuleAttributes) String() string {
	return tea.Prettify(s)
}

func (s QueryDetailItemResponseBodyModuleAttributes) GoString() string {
	return s.String()
}

func (s *QueryDetailItemResponseBodyModuleAttributes) SetAttribute(v []*QueryDetailItemResponseBodyModuleAttributesAttribute) *QueryDetailItemResponseBodyModuleAttributes {
	s.Attribute = v
	return s
}

type QueryDetailItemResponseBodyModuleAttributesAttribute struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryDetailItemResponseBodyModuleAttributesAttribute) String() string {
	return tea.Prettify(s)
}

func (s QueryDetailItemResponseBodyModuleAttributesAttribute) GoString() string {
	return s.String()
}

func (s *QueryDetailItemResponseBodyModuleAttributesAttribute) SetName(v string) *QueryDetailItemResponseBodyModuleAttributesAttribute {
	s.Name = &v
	return s
}

func (s *QueryDetailItemResponseBodyModuleAttributesAttribute) SetTitle(v string) *QueryDetailItemResponseBodyModuleAttributesAttribute {
	s.Title = &v
	return s
}

func (s *QueryDetailItemResponseBodyModuleAttributesAttribute) SetValue(v string) *QueryDetailItemResponseBodyModuleAttributesAttribute {
	s.Value = &v
	return s
}

type QueryDetailItemResponseBodyModuleDetailPicUrl struct {
	PicUlr []*string `json:"picUlr,omitempty" xml:"picUlr,omitempty" type:"Repeated"`
}

func (s QueryDetailItemResponseBodyModuleDetailPicUrl) String() string {
	return tea.Prettify(s)
}

func (s QueryDetailItemResponseBodyModuleDetailPicUrl) GoString() string {
	return s.String()
}

func (s *QueryDetailItemResponseBodyModuleDetailPicUrl) SetPicUlr(v []*string) *QueryDetailItemResponseBodyModuleDetailPicUrl {
	s.PicUlr = v
	return s
}

type QueryDetailItemResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDetailItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDetailItemResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDetailItemResponse) GoString() string {
	return s.String()
}

func (s *QueryDetailItemResponse) SetHeaders(v map[string]*string) *QueryDetailItemResponse {
	s.Headers = v
	return s
}

func (s *QueryDetailItemResponse) SetStatusCode(v int32) *QueryDetailItemResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDetailItemResponse) SetBody(v *QueryDetailItemResponseBody) *QueryDetailItemResponse {
	s.Body = v
	return s
}

type QueryRemainResourcesRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s QueryRemainResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRemainResourcesRequest) GoString() string {
	return s.String()
}

func (s *QueryRemainResourcesRequest) SetBizType(v string) *QueryRemainResourcesRequest {
	s.BizType = &v
	return s
}

type QueryRemainResourcesResponseBody struct {
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRemainResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRemainResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRemainResourcesResponseBody) SetData(v int64) *QueryRemainResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRemainResourcesResponseBody) SetRequestId(v string) *QueryRemainResourcesResponseBody {
	s.RequestId = &v
	return s
}

type QueryRemainResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRemainResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRemainResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRemainResourcesResponse) GoString() string {
	return s.String()
}

func (s *QueryRemainResourcesResponse) SetHeaders(v map[string]*string) *QueryRemainResourcesResponse {
	s.Headers = v
	return s
}

func (s *QueryRemainResourcesResponse) SetStatusCode(v int32) *QueryRemainResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRemainResourcesResponse) SetBody(v *QueryRemainResourcesResponseBody) *QueryRemainResourcesResponse {
	s.Body = v
	return s
}

type RefuseSupplementRequest struct {
	SupplementId *int64 `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
}

func (s RefuseSupplementRequest) String() string {
	return tea.Prettify(s)
}

func (s RefuseSupplementRequest) GoString() string {
	return s.String()
}

func (s *RefuseSupplementRequest) SetSupplementId(v int64) *RefuseSupplementRequest {
	s.SupplementId = &v
	return s
}

type RefuseSupplementResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefuseSupplementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefuseSupplementResponseBody) GoString() string {
	return s.String()
}

func (s *RefuseSupplementResponseBody) SetCode(v string) *RefuseSupplementResponseBody {
	s.Code = &v
	return s
}

func (s *RefuseSupplementResponseBody) SetMessage(v string) *RefuseSupplementResponseBody {
	s.Message = &v
	return s
}

func (s *RefuseSupplementResponseBody) SetRequestId(v string) *RefuseSupplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefuseSupplementResponseBody) SetSuccess(v bool) *RefuseSupplementResponseBody {
	s.Success = &v
	return s
}

type RefuseSupplementResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefuseSupplementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefuseSupplementResponse) String() string {
	return tea.Prettify(s)
}

func (s RefuseSupplementResponse) GoString() string {
	return s.String()
}

func (s *RefuseSupplementResponse) SetHeaders(v map[string]*string) *RefuseSupplementResponse {
	s.Headers = v
	return s
}

func (s *RefuseSupplementResponse) SetStatusCode(v int32) *RefuseSupplementResponse {
	s.StatusCode = &v
	return s
}

func (s *RefuseSupplementResponse) SetBody(v *RefuseSupplementResponseBody) *RefuseSupplementResponse {
	s.Body = v
	return s
}

type RejectExpertSolutionRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Note  *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s RejectExpertSolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectExpertSolutionRequest) GoString() string {
	return s.String()
}

func (s *RejectExpertSolutionRequest) SetBizId(v string) *RejectExpertSolutionRequest {
	s.BizId = &v
	return s
}

func (s *RejectExpertSolutionRequest) SetNote(v string) *RejectExpertSolutionRequest {
	s.Note = &v
	return s
}

type RejectExpertSolutionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectExpertSolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectExpertSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *RejectExpertSolutionResponseBody) SetRequestId(v string) *RejectExpertSolutionResponseBody {
	s.RequestId = &v
	return s
}

type RejectExpertSolutionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectExpertSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectExpertSolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectExpertSolutionResponse) GoString() string {
	return s.String()
}

func (s *RejectExpertSolutionResponse) SetHeaders(v map[string]*string) *RejectExpertSolutionResponse {
	s.Headers = v
	return s
}

func (s *RejectExpertSolutionResponse) SetStatusCode(v int32) *RejectExpertSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectExpertSolutionResponse) SetBody(v *RejectExpertSolutionResponseBody) *RejectExpertSolutionResponse {
	s.Body = v
	return s
}

type RemoveApplicantRequest struct {
	ApplicantId *int64 `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
}

func (s RemoveApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveApplicantRequest) GoString() string {
	return s.String()
}

func (s *RemoveApplicantRequest) SetApplicantId(v int64) *RemoveApplicantRequest {
	s.ApplicantId = &v
	return s
}

type RemoveApplicantResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApplicantResponseBody) SetCode(v string) *RemoveApplicantResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveApplicantResponseBody) SetMessage(v string) *RemoveApplicantResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveApplicantResponseBody) SetRequestId(v string) *RemoveApplicantResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveApplicantResponseBody) SetSuccess(v bool) *RemoveApplicantResponseBody {
	s.Success = &v
	return s
}

type RemoveApplicantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveApplicantResponse) GoString() string {
	return s.String()
}

func (s *RemoveApplicantResponse) SetHeaders(v map[string]*string) *RemoveApplicantResponse {
	s.Headers = v
	return s
}

func (s *RemoveApplicantResponse) SetStatusCode(v int32) *RemoveApplicantResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApplicantResponse) SetBody(v *RemoveApplicantResponseBody) *RemoveApplicantResponse {
	s.Body = v
	return s
}

type SaveSearchConditionRequest struct {
	ConditionContent *string `json:"ConditionContent,omitempty" xml:"ConditionContent,omitempty"`
	SessionId        *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TagName          *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Umid             *string `json:"Umid,omitempty" xml:"Umid,omitempty"`
}

func (s SaveSearchConditionRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSearchConditionRequest) GoString() string {
	return s.String()
}

func (s *SaveSearchConditionRequest) SetConditionContent(v string) *SaveSearchConditionRequest {
	s.ConditionContent = &v
	return s
}

func (s *SaveSearchConditionRequest) SetSessionId(v string) *SaveSearchConditionRequest {
	s.SessionId = &v
	return s
}

func (s *SaveSearchConditionRequest) SetTagName(v string) *SaveSearchConditionRequest {
	s.TagName = &v
	return s
}

func (s *SaveSearchConditionRequest) SetType(v int32) *SaveSearchConditionRequest {
	s.Type = &v
	return s
}

func (s *SaveSearchConditionRequest) SetUmid(v string) *SaveSearchConditionRequest {
	s.Umid = &v
	return s
}

type SaveSearchConditionResponseBody struct {
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ConditionContent *string `json:"ConditionContent,omitempty" xml:"ConditionContent,omitempty"`
	ConditionId      *int64  `json:"ConditionId,omitempty" xml:"ConditionId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionId        *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TagName          *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Umid             *string `json:"Umid,omitempty" xml:"Umid,omitempty"`
	UserId           *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SaveSearchConditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSearchConditionResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSearchConditionResponseBody) SetCode(v string) *SaveSearchConditionResponseBody {
	s.Code = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetConditionContent(v string) *SaveSearchConditionResponseBody {
	s.ConditionContent = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetConditionId(v int64) *SaveSearchConditionResponseBody {
	s.ConditionId = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetMessage(v string) *SaveSearchConditionResponseBody {
	s.Message = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetRequestId(v string) *SaveSearchConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetSessionId(v string) *SaveSearchConditionResponseBody {
	s.SessionId = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetSuccess(v bool) *SaveSearchConditionResponseBody {
	s.Success = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetTagName(v string) *SaveSearchConditionResponseBody {
	s.TagName = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetType(v int32) *SaveSearchConditionResponseBody {
	s.Type = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetUmid(v string) *SaveSearchConditionResponseBody {
	s.Umid = &v
	return s
}

func (s *SaveSearchConditionResponseBody) SetUserId(v int64) *SaveSearchConditionResponseBody {
	s.UserId = &v
	return s
}

type SaveSearchConditionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSearchConditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSearchConditionResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSearchConditionResponse) GoString() string {
	return s.String()
}

func (s *SaveSearchConditionResponse) SetHeaders(v map[string]*string) *SaveSearchConditionResponse {
	s.Headers = v
	return s
}

func (s *SaveSearchConditionResponse) SetStatusCode(v int32) *SaveSearchConditionResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSearchConditionResponse) SetBody(v *SaveSearchConditionResponseBody) *SaveSearchConditionResponse {
	s.Body = v
	return s
}

type SaveTemporaryApplicantRequest struct {
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantId           *int64  `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	BusinessLicenceOssKey *string `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	City                  *string `json:"City,omitempty" xml:"City,omitempty"`
	CompleteApplicant     *bool   `json:"CompleteApplicant,omitempty" xml:"CompleteApplicant,omitempty"`
	ContactAddress        *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactCity           *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactCounty         *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
	ContactDistrict       *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactEmail          *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName           *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber         *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactProvince       *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactZipCode        *string `json:"ContactZipCode,omitempty" xml:"ContactZipCode,omitempty"`
	Country               *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EAddress              *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	EName                 *string `json:"EName,omitempty" xml:"EName,omitempty"`
	IdCardOssKey          *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	LegalNoticeOssKey     *string `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	LoaOssKey             *string `json:"LoaOssKey,omitempty" xml:"LoaOssKey,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PassportOssKey        *string `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	PrincipalName         *int32  `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	Province              *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Region                *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Town                  *string `json:"Town,omitempty" xml:"Town,omitempty"`
	Type                  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SaveTemporaryApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTemporaryApplicantRequest) GoString() string {
	return s.String()
}

func (s *SaveTemporaryApplicantRequest) SetAddress(v string) *SaveTemporaryApplicantRequest {
	s.Address = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetApplicantId(v int64) *SaveTemporaryApplicantRequest {
	s.ApplicantId = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetBusinessLicenceOssKey(v string) *SaveTemporaryApplicantRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetCardNumber(v string) *SaveTemporaryApplicantRequest {
	s.CardNumber = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetCity(v string) *SaveTemporaryApplicantRequest {
	s.City = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetCompleteApplicant(v bool) *SaveTemporaryApplicantRequest {
	s.CompleteApplicant = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactAddress(v string) *SaveTemporaryApplicantRequest {
	s.ContactAddress = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactCity(v string) *SaveTemporaryApplicantRequest {
	s.ContactCity = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactCounty(v string) *SaveTemporaryApplicantRequest {
	s.ContactCounty = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactDistrict(v string) *SaveTemporaryApplicantRequest {
	s.ContactDistrict = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactEmail(v string) *SaveTemporaryApplicantRequest {
	s.ContactEmail = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactName(v string) *SaveTemporaryApplicantRequest {
	s.ContactName = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactNumber(v string) *SaveTemporaryApplicantRequest {
	s.ContactNumber = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactProvince(v string) *SaveTemporaryApplicantRequest {
	s.ContactProvince = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetContactZipCode(v string) *SaveTemporaryApplicantRequest {
	s.ContactZipCode = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetCountry(v string) *SaveTemporaryApplicantRequest {
	s.Country = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetEAddress(v string) *SaveTemporaryApplicantRequest {
	s.EAddress = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetEName(v string) *SaveTemporaryApplicantRequest {
	s.EName = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetIdCardOssKey(v string) *SaveTemporaryApplicantRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetLegalNoticeOssKey(v string) *SaveTemporaryApplicantRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetLoaOssKey(v string) *SaveTemporaryApplicantRequest {
	s.LoaOssKey = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetName(v string) *SaveTemporaryApplicantRequest {
	s.Name = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetPassportOssKey(v string) *SaveTemporaryApplicantRequest {
	s.PassportOssKey = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetPrincipalName(v int32) *SaveTemporaryApplicantRequest {
	s.PrincipalName = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetProvince(v string) *SaveTemporaryApplicantRequest {
	s.Province = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetRegion(v string) *SaveTemporaryApplicantRequest {
	s.Region = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetTown(v string) *SaveTemporaryApplicantRequest {
	s.Town = &v
	return s
}

func (s *SaveTemporaryApplicantRequest) SetType(v string) *SaveTemporaryApplicantRequest {
	s.Type = &v
	return s
}

type SaveTemporaryApplicantResponseBody struct {
	ApplicantId  *int64  `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveTemporaryApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTemporaryApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTemporaryApplicantResponseBody) SetApplicantId(v int64) *SaveTemporaryApplicantResponseBody {
	s.ApplicantId = &v
	return s
}

func (s *SaveTemporaryApplicantResponseBody) SetErrorCode(v string) *SaveTemporaryApplicantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SaveTemporaryApplicantResponseBody) SetErrorMessage(v string) *SaveTemporaryApplicantResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SaveTemporaryApplicantResponseBody) SetRequestId(v string) *SaveTemporaryApplicantResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTemporaryApplicantResponseBody) SetSuccess(v bool) *SaveTemporaryApplicantResponseBody {
	s.Success = &v
	return s
}

type SaveTemporaryApplicantResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveTemporaryApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveTemporaryApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTemporaryApplicantResponse) GoString() string {
	return s.String()
}

func (s *SaveTemporaryApplicantResponse) SetHeaders(v map[string]*string) *SaveTemporaryApplicantResponse {
	s.Headers = v
	return s
}

func (s *SaveTemporaryApplicantResponse) SetStatusCode(v int32) *SaveTemporaryApplicantResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTemporaryApplicantResponse) SetBody(v *SaveTemporaryApplicantResponseBody) *SaveTemporaryApplicantResponse {
	s.Body = v
	return s
}

type SearchItemsRequest struct {
	ExcludedTags        *string `json:"ExcludedTags,omitempty" xml:"ExcludedTags,omitempty"`
	ExcludedUids        *string `json:"ExcludedUids,omitempty" xml:"ExcludedUids,omitempty"`
	FeedsType           *bool   `json:"FeedsType,omitempty" xml:"FeedsType,omitempty"`
	IntCls              *string `json:"IntCls,omitempty" xml:"IntCls,omitempty"`
	Keywords            *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Mock                *bool   `json:"Mock,omitempty" xml:"Mock,omitempty"`
	PageIndex           *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PriceLeft           *string `json:"PriceLeft,omitempty" xml:"PriceLeft,omitempty"`
	PriceRight          *string `json:"PriceRight,omitempty" xml:"PriceRight,omitempty"`
	Products            *string `json:"Products,omitempty" xml:"Products,omitempty"`
	RegisterNumber      *string `json:"RegisterNumber,omitempty" xml:"RegisterNumber,omitempty"`
	Sort                *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SortType            *int32  `json:"SortType,omitempty" xml:"SortType,omitempty"`
	Tags                *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TrademarkNameLength *int32  `json:"TrademarkNameLength,omitempty" xml:"TrademarkNameLength,omitempty"`
	TrademarkNameType   *string `json:"TrademarkNameType,omitempty" xml:"TrademarkNameType,omitempty"`
	UmId                *string `json:"UmId,omitempty" xml:"UmId,omitempty"`
}

func (s SearchItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchItemsRequest) GoString() string {
	return s.String()
}

func (s *SearchItemsRequest) SetExcludedTags(v string) *SearchItemsRequest {
	s.ExcludedTags = &v
	return s
}

func (s *SearchItemsRequest) SetExcludedUids(v string) *SearchItemsRequest {
	s.ExcludedUids = &v
	return s
}

func (s *SearchItemsRequest) SetFeedsType(v bool) *SearchItemsRequest {
	s.FeedsType = &v
	return s
}

func (s *SearchItemsRequest) SetIntCls(v string) *SearchItemsRequest {
	s.IntCls = &v
	return s
}

func (s *SearchItemsRequest) SetKeywords(v string) *SearchItemsRequest {
	s.Keywords = &v
	return s
}

func (s *SearchItemsRequest) SetMock(v bool) *SearchItemsRequest {
	s.Mock = &v
	return s
}

func (s *SearchItemsRequest) SetPageIndex(v int32) *SearchItemsRequest {
	s.PageIndex = &v
	return s
}

func (s *SearchItemsRequest) SetPageSize(v int32) *SearchItemsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchItemsRequest) SetPriceLeft(v string) *SearchItemsRequest {
	s.PriceLeft = &v
	return s
}

func (s *SearchItemsRequest) SetPriceRight(v string) *SearchItemsRequest {
	s.PriceRight = &v
	return s
}

func (s *SearchItemsRequest) SetProducts(v string) *SearchItemsRequest {
	s.Products = &v
	return s
}

func (s *SearchItemsRequest) SetRegisterNumber(v string) *SearchItemsRequest {
	s.RegisterNumber = &v
	return s
}

func (s *SearchItemsRequest) SetSort(v string) *SearchItemsRequest {
	s.Sort = &v
	return s
}

func (s *SearchItemsRequest) SetSortType(v int32) *SearchItemsRequest {
	s.SortType = &v
	return s
}

func (s *SearchItemsRequest) SetTags(v string) *SearchItemsRequest {
	s.Tags = &v
	return s
}

func (s *SearchItemsRequest) SetTrademarkNameLength(v int32) *SearchItemsRequest {
	s.TrademarkNameLength = &v
	return s
}

func (s *SearchItemsRequest) SetTrademarkNameType(v string) *SearchItemsRequest {
	s.TrademarkNameType = &v
	return s
}

func (s *SearchItemsRequest) SetUmId(v string) *SearchItemsRequest {
	s.UmId = &v
	return s
}

type SearchItemsResponseBody struct {
	Module *SearchItemsResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
}

func (s SearchItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchItemsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchItemsResponseBody) SetModule(v *SearchItemsResponseBodyModule) *SearchItemsResponseBody {
	s.Module = v
	return s
}

type SearchItemsResponseBodyModule struct {
	CurrentPageNum *int32                             `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Date           *SearchItemsResponseBodyModuleDate `json:"Date,omitempty" xml:"Date,omitempty" type:"Struct"`
	PageSize       *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                             `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                             `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s SearchItemsResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s SearchItemsResponseBodyModule) GoString() string {
	return s.String()
}

func (s *SearchItemsResponseBodyModule) SetCurrentPageNum(v int32) *SearchItemsResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *SearchItemsResponseBodyModule) SetDate(v *SearchItemsResponseBodyModuleDate) *SearchItemsResponseBodyModule {
	s.Date = v
	return s
}

func (s *SearchItemsResponseBodyModule) SetPageSize(v int32) *SearchItemsResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *SearchItemsResponseBodyModule) SetTotalItemNum(v int32) *SearchItemsResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *SearchItemsResponseBodyModule) SetTotalPageNum(v int32) *SearchItemsResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

type SearchItemsResponseBodyModuleDate struct {
	Item []*SearchItemsResponseBodyModuleDateItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s SearchItemsResponseBodyModuleDate) String() string {
	return tea.Prettify(s)
}

func (s SearchItemsResponseBodyModuleDate) GoString() string {
	return s.String()
}

func (s *SearchItemsResponseBodyModuleDate) SetItem(v []*SearchItemsResponseBodyModuleDateItem) *SearchItemsResponseBodyModuleDate {
	s.Item = v
	return s
}

type SearchItemsResponseBodyModuleDateItem struct {
	DetailViewObjectSourceDatum *string `json:"DetailViewObjectSourceDatum,omitempty" xml:"DetailViewObjectSourceDatum,omitempty"`
	DetailViewObjectSourceType  *string `json:"DetailViewObjectSourceType,omitempty" xml:"DetailViewObjectSourceType,omitempty"`
}

func (s SearchItemsResponseBodyModuleDateItem) String() string {
	return tea.Prettify(s)
}

func (s SearchItemsResponseBodyModuleDateItem) GoString() string {
	return s.String()
}

func (s *SearchItemsResponseBodyModuleDateItem) SetDetailViewObjectSourceDatum(v string) *SearchItemsResponseBodyModuleDateItem {
	s.DetailViewObjectSourceDatum = &v
	return s
}

func (s *SearchItemsResponseBodyModuleDateItem) SetDetailViewObjectSourceType(v string) *SearchItemsResponseBodyModuleDateItem {
	s.DetailViewObjectSourceType = &v
	return s
}

type SearchItemsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchItemsResponse) GoString() string {
	return s.String()
}

func (s *SearchItemsResponse) SetHeaders(v map[string]*string) *SearchItemsResponse {
	s.Headers = v
	return s
}

func (s *SearchItemsResponse) SetStatusCode(v int32) *SearchItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchItemsResponse) SetBody(v *SearchItemsResponseBody) *SearchItemsResponse {
	s.Body = v
	return s
}

type SearchSimilarityRequest struct {
	Classifications map[string]interface{}                `json:"Classifications,omitempty" xml:"Classifications,omitempty"`
	Limit           *int32                                `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NameUriList     []*SearchSimilarityRequestNameUriList `json:"NameUriList,omitempty" xml:"NameUriList,omitempty" type:"Repeated"`
	SearchType      *string                               `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	ShowDetail      *bool                                 `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty"`
	SimilarGroups   map[string]interface{}                `json:"SimilarGroups,omitempty" xml:"SimilarGroups,omitempty"`
	Sorter          *string                               `json:"Sorter,omitempty" xml:"Sorter,omitempty"`
	Umid            *string                               `json:"Umid,omitempty" xml:"Umid,omitempty"`
}

func (s SearchSimilarityRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityRequest) GoString() string {
	return s.String()
}

func (s *SearchSimilarityRequest) SetClassifications(v map[string]interface{}) *SearchSimilarityRequest {
	s.Classifications = v
	return s
}

func (s *SearchSimilarityRequest) SetLimit(v int32) *SearchSimilarityRequest {
	s.Limit = &v
	return s
}

func (s *SearchSimilarityRequest) SetNameUriList(v []*SearchSimilarityRequestNameUriList) *SearchSimilarityRequest {
	s.NameUriList = v
	return s
}

func (s *SearchSimilarityRequest) SetSearchType(v string) *SearchSimilarityRequest {
	s.SearchType = &v
	return s
}

func (s *SearchSimilarityRequest) SetShowDetail(v bool) *SearchSimilarityRequest {
	s.ShowDetail = &v
	return s
}

func (s *SearchSimilarityRequest) SetSimilarGroups(v map[string]interface{}) *SearchSimilarityRequest {
	s.SimilarGroups = v
	return s
}

func (s *SearchSimilarityRequest) SetSorter(v string) *SearchSimilarityRequest {
	s.Sorter = &v
	return s
}

func (s *SearchSimilarityRequest) SetUmid(v string) *SearchSimilarityRequest {
	s.Umid = &v
	return s
}

type SearchSimilarityRequestNameUriList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Uri  *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SearchSimilarityRequestNameUriList) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityRequestNameUriList) GoString() string {
	return s.String()
}

func (s *SearchSimilarityRequestNameUriList) SetName(v string) *SearchSimilarityRequestNameUriList {
	s.Name = &v
	return s
}

func (s *SearchSimilarityRequestNameUriList) SetUri(v string) *SearchSimilarityRequestNameUriList {
	s.Uri = &v
	return s
}

type SearchSimilarityShrinkRequest struct {
	ClassificationsShrink *string                                     `json:"Classifications,omitempty" xml:"Classifications,omitempty"`
	Limit                 *int32                                      `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NameUriList           []*SearchSimilarityShrinkRequestNameUriList `json:"NameUriList,omitempty" xml:"NameUriList,omitempty" type:"Repeated"`
	SearchType            *string                                     `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	ShowDetail            *bool                                       `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty"`
	SimilarGroupsShrink   *string                                     `json:"SimilarGroups,omitempty" xml:"SimilarGroups,omitempty"`
	Sorter                *string                                     `json:"Sorter,omitempty" xml:"Sorter,omitempty"`
	Umid                  *string                                     `json:"Umid,omitempty" xml:"Umid,omitempty"`
}

func (s SearchSimilarityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchSimilarityShrinkRequest) SetClassificationsShrink(v string) *SearchSimilarityShrinkRequest {
	s.ClassificationsShrink = &v
	return s
}

func (s *SearchSimilarityShrinkRequest) SetLimit(v int32) *SearchSimilarityShrinkRequest {
	s.Limit = &v
	return s
}

func (s *SearchSimilarityShrinkRequest) SetNameUriList(v []*SearchSimilarityShrinkRequestNameUriList) *SearchSimilarityShrinkRequest {
	s.NameUriList = v
	return s
}

func (s *SearchSimilarityShrinkRequest) SetSearchType(v string) *SearchSimilarityShrinkRequest {
	s.SearchType = &v
	return s
}

func (s *SearchSimilarityShrinkRequest) SetShowDetail(v bool) *SearchSimilarityShrinkRequest {
	s.ShowDetail = &v
	return s
}

func (s *SearchSimilarityShrinkRequest) SetSimilarGroupsShrink(v string) *SearchSimilarityShrinkRequest {
	s.SimilarGroupsShrink = &v
	return s
}

func (s *SearchSimilarityShrinkRequest) SetSorter(v string) *SearchSimilarityShrinkRequest {
	s.Sorter = &v
	return s
}

func (s *SearchSimilarityShrinkRequest) SetUmid(v string) *SearchSimilarityShrinkRequest {
	s.Umid = &v
	return s
}

type SearchSimilarityShrinkRequestNameUriList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Uri  *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SearchSimilarityShrinkRequestNameUriList) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityShrinkRequestNameUriList) GoString() string {
	return s.String()
}

func (s *SearchSimilarityShrinkRequestNameUriList) SetName(v string) *SearchSimilarityShrinkRequestNameUriList {
	s.Name = &v
	return s
}

func (s *SearchSimilarityShrinkRequestNameUriList) SetUri(v string) *SearchSimilarityShrinkRequestNameUriList {
	s.Uri = &v
	return s
}

type SearchSimilarityResponseBody struct {
	DataList  []*SearchSimilarityResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchSimilarityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityResponseBody) GoString() string {
	return s.String()
}

func (s *SearchSimilarityResponseBody) SetDataList(v []*SearchSimilarityResponseBodyDataList) *SearchSimilarityResponseBody {
	s.DataList = v
	return s
}

func (s *SearchSimilarityResponseBody) SetRequestId(v string) *SearchSimilarityResponseBody {
	s.RequestId = &v
	return s
}

type SearchSimilarityResponseBodyDataList struct {
	ClassificationSimilarityList []*SearchSimilarityResponseBodyDataListClassificationSimilarityList `json:"ClassificationSimilarityList,omitempty" xml:"ClassificationSimilarityList,omitempty" type:"Repeated"`
	Name                         *string                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Uri                          *string                                                             `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SearchSimilarityResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *SearchSimilarityResponseBodyDataList) SetClassificationSimilarityList(v []*SearchSimilarityResponseBodyDataListClassificationSimilarityList) *SearchSimilarityResponseBodyDataList {
	s.ClassificationSimilarityList = v
	return s
}

func (s *SearchSimilarityResponseBodyDataList) SetName(v string) *SearchSimilarityResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataList) SetUri(v string) *SearchSimilarityResponseBodyDataList {
	s.Uri = &v
	return s
}

type SearchSimilarityResponseBodyDataListClassificationSimilarityList struct {
	Classification     *int32                                                                              `json:"Classification,omitempty" xml:"Classification,omitempty"`
	ClassificationName *string                                                                             `json:"ClassificationName,omitempty" xml:"ClassificationName,omitempty"`
	Rate               *int32                                                                              `json:"Rate,omitempty" xml:"Rate,omitempty"`
	SimilarGroupList   []*SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList `json:"SimilarGroupList,omitempty" xml:"SimilarGroupList,omitempty" type:"Repeated"`
}

func (s SearchSimilarityResponseBodyDataListClassificationSimilarityList) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityResponseBodyDataListClassificationSimilarityList) GoString() string {
	return s.String()
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityList) SetClassification(v int32) *SearchSimilarityResponseBodyDataListClassificationSimilarityList {
	s.Classification = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityList) SetClassificationName(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityList {
	s.ClassificationName = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityList) SetRate(v int32) *SearchSimilarityResponseBodyDataListClassificationSimilarityList {
	s.Rate = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityList) SetSimilarGroupList(v []*SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList) *SearchSimilarityResponseBodyDataListClassificationSimilarityList {
	s.SimilarGroupList = v
	return s
}

type SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList struct {
	DetailList       []*SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
	Rate             *int32                                                                                        `json:"Rate,omitempty" xml:"Rate,omitempty"`
	SimilarGroup     *string                                                                                       `json:"SimilarGroup,omitempty" xml:"SimilarGroup,omitempty"`
	SimilarGroupName *string                                                                                       `json:"SimilarGroupName,omitempty" xml:"SimilarGroupName,omitempty"`
}

func (s SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList) GoString() string {
	return s.String()
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList) SetDetailList(v []*SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList {
	s.DetailList = v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList) SetRate(v int32) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList {
	s.Rate = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList) SetSimilarGroup(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList {
	s.SimilarGroup = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList) SetSimilarGroupName(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupList {
	s.SimilarGroupName = &v
	return s
}

type SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Rate     *int32  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	TmNumber *string `json:"TmNumber,omitempty" xml:"TmNumber,omitempty"`
	Uid      *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Uri      *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) GoString() string {
	return s.String()
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) SetCode(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList {
	s.Code = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) SetName(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList {
	s.Name = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) SetRate(v int32) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList {
	s.Rate = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) SetTmNumber(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList {
	s.TmNumber = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) SetUid(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList {
	s.Uid = &v
	return s
}

func (s *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList) SetUri(v string) *SearchSimilarityResponseBodyDataListClassificationSimilarityListSimilarGroupListDetailList {
	s.Uri = &v
	return s
}

type SearchSimilarityResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchSimilarityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchSimilarityResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityResponse) GoString() string {
	return s.String()
}

func (s *SearchSimilarityResponse) SetHeaders(v map[string]*string) *SearchSimilarityResponse {
	s.Headers = v
	return s
}

func (s *SearchSimilarityResponse) SetStatusCode(v int32) *SearchSimilarityResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchSimilarityResponse) SetBody(v *SearchSimilarityResponseBody) *SearchSimilarityResponse {
	s.Body = v
	return s
}

type SearchSimilarityListRequest struct {
	Classifications   map[string]interface{} `json:"Classifications,omitempty" xml:"Classifications,omitempty"`
	Name              *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderId           *string                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber        *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SimilarGroups     map[string]interface{} `json:"SimilarGroups,omitempty" xml:"SimilarGroups,omitempty"`
	Status            *int32                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessSearchType *string                `json:"SuccessSearchType,omitempty" xml:"SuccessSearchType,omitempty"`
	Umid              *string                `json:"Umid,omitempty" xml:"Umid,omitempty"`
	Uri               *string                `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SearchSimilarityListRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityListRequest) GoString() string {
	return s.String()
}

func (s *SearchSimilarityListRequest) SetClassifications(v map[string]interface{}) *SearchSimilarityListRequest {
	s.Classifications = v
	return s
}

func (s *SearchSimilarityListRequest) SetName(v string) *SearchSimilarityListRequest {
	s.Name = &v
	return s
}

func (s *SearchSimilarityListRequest) SetOrderId(v string) *SearchSimilarityListRequest {
	s.OrderId = &v
	return s
}

func (s *SearchSimilarityListRequest) SetPageNumber(v int32) *SearchSimilarityListRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchSimilarityListRequest) SetPageSize(v int32) *SearchSimilarityListRequest {
	s.PageSize = &v
	return s
}

func (s *SearchSimilarityListRequest) SetSimilarGroups(v map[string]interface{}) *SearchSimilarityListRequest {
	s.SimilarGroups = v
	return s
}

func (s *SearchSimilarityListRequest) SetStatus(v int32) *SearchSimilarityListRequest {
	s.Status = &v
	return s
}

func (s *SearchSimilarityListRequest) SetSuccessSearchType(v string) *SearchSimilarityListRequest {
	s.SuccessSearchType = &v
	return s
}

func (s *SearchSimilarityListRequest) SetUmid(v string) *SearchSimilarityListRequest {
	s.Umid = &v
	return s
}

func (s *SearchSimilarityListRequest) SetUri(v string) *SearchSimilarityListRequest {
	s.Uri = &v
	return s
}

type SearchSimilarityListShrinkRequest struct {
	ClassificationsShrink *string `json:"Classifications,omitempty" xml:"Classifications,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderId               *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SimilarGroupsShrink   *string `json:"SimilarGroups,omitempty" xml:"SimilarGroups,omitempty"`
	Status                *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessSearchType     *string `json:"SuccessSearchType,omitempty" xml:"SuccessSearchType,omitempty"`
	Umid                  *string `json:"Umid,omitempty" xml:"Umid,omitempty"`
	Uri                   *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s SearchSimilarityListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityListShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchSimilarityListShrinkRequest) SetClassificationsShrink(v string) *SearchSimilarityListShrinkRequest {
	s.ClassificationsShrink = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetName(v string) *SearchSimilarityListShrinkRequest {
	s.Name = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetOrderId(v string) *SearchSimilarityListShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetPageNumber(v int32) *SearchSimilarityListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetPageSize(v int32) *SearchSimilarityListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetSimilarGroupsShrink(v string) *SearchSimilarityListShrinkRequest {
	s.SimilarGroupsShrink = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetStatus(v int32) *SearchSimilarityListShrinkRequest {
	s.Status = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetSuccessSearchType(v string) *SearchSimilarityListShrinkRequest {
	s.SuccessSearchType = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetUmid(v string) *SearchSimilarityListShrinkRequest {
	s.Umid = &v
	return s
}

func (s *SearchSimilarityListShrinkRequest) SetUri(v string) *SearchSimilarityListShrinkRequest {
	s.Uri = &v
	return s
}

type SearchSimilarityListResponseBody struct {
	Data       []*SearchSimilarityListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchSimilarityListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchSimilarityListResponseBody) SetData(v []*SearchSimilarityListResponseBodyData) *SearchSimilarityListResponseBody {
	s.Data = v
	return s
}

func (s *SearchSimilarityListResponseBody) SetPageNumber(v int32) *SearchSimilarityListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchSimilarityListResponseBody) SetPageSize(v int32) *SearchSimilarityListResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchSimilarityListResponseBody) SetRequestId(v string) *SearchSimilarityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchSimilarityListResponseBody) SetTotalCount(v int32) *SearchSimilarityListResponseBody {
	s.TotalCount = &v
	return s
}

type SearchSimilarityListResponseBodyData struct {
	ApplyDate           *string `json:"ApplyDate,omitempty" xml:"ApplyDate,omitempty"`
	Classification      *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	ExclusiveDateLimit  *string `json:"ExclusiveDateLimit,omitempty" xml:"ExclusiveDateLimit,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Image               *string `json:"Image,omitempty" xml:"Image,omitempty"`
	LastProcedureStatus *string `json:"LastProcedureStatus,omitempty" xml:"LastProcedureStatus,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OnSale              *int32  `json:"OnSale,omitempty" xml:"OnSale,omitempty"`
	OwnerAddress        *string `json:"OwnerAddress,omitempty" xml:"OwnerAddress,omitempty"`
	OwnerEnAddress      *string `json:"OwnerEnAddress,omitempty" xml:"OwnerEnAddress,omitempty"`
	OwnerEnName         *string `json:"OwnerEnName,omitempty" xml:"OwnerEnName,omitempty"`
	OwnerName           *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	PreAnnDate          *string `json:"PreAnnDate,omitempty" xml:"PreAnnDate,omitempty"`
	PreAnnNum           *string `json:"PreAnnNum,omitempty" xml:"PreAnnNum,omitempty"`
	Product             *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ProductDesc         *string `json:"ProductDesc,omitempty" xml:"ProductDesc,omitempty"`
	RegAnnDate          *string `json:"RegAnnDate,omitempty" xml:"RegAnnDate,omitempty"`
	RegAnnNum           *string `json:"RegAnnNum,omitempty" xml:"RegAnnNum,omitempty"`
	RegistrationNumber  *string `json:"RegistrationNumber,omitempty" xml:"RegistrationNumber,omitempty"`
	Share               *string `json:"Share,omitempty" xml:"Share,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Uid                 *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s SearchSimilarityListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityListResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchSimilarityListResponseBodyData) SetApplyDate(v string) *SearchSimilarityListResponseBodyData {
	s.ApplyDate = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetClassification(v string) *SearchSimilarityListResponseBodyData {
	s.Classification = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetExclusiveDateLimit(v string) *SearchSimilarityListResponseBodyData {
	s.ExclusiveDateLimit = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetId(v int64) *SearchSimilarityListResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetImage(v string) *SearchSimilarityListResponseBodyData {
	s.Image = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetLastProcedureStatus(v string) *SearchSimilarityListResponseBodyData {
	s.LastProcedureStatus = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetName(v string) *SearchSimilarityListResponseBodyData {
	s.Name = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetOnSale(v int32) *SearchSimilarityListResponseBodyData {
	s.OnSale = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetOwnerAddress(v string) *SearchSimilarityListResponseBodyData {
	s.OwnerAddress = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetOwnerEnAddress(v string) *SearchSimilarityListResponseBodyData {
	s.OwnerEnAddress = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetOwnerEnName(v string) *SearchSimilarityListResponseBodyData {
	s.OwnerEnName = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetOwnerName(v string) *SearchSimilarityListResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetPreAnnDate(v string) *SearchSimilarityListResponseBodyData {
	s.PreAnnDate = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetPreAnnNum(v string) *SearchSimilarityListResponseBodyData {
	s.PreAnnNum = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetProduct(v string) *SearchSimilarityListResponseBodyData {
	s.Product = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetProductDesc(v string) *SearchSimilarityListResponseBodyData {
	s.ProductDesc = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetRegAnnDate(v string) *SearchSimilarityListResponseBodyData {
	s.RegAnnDate = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetRegAnnNum(v string) *SearchSimilarityListResponseBodyData {
	s.RegAnnNum = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetRegistrationNumber(v string) *SearchSimilarityListResponseBodyData {
	s.RegistrationNumber = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetShare(v string) *SearchSimilarityListResponseBodyData {
	s.Share = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetStatus(v string) *SearchSimilarityListResponseBodyData {
	s.Status = &v
	return s
}

func (s *SearchSimilarityListResponseBodyData) SetUid(v string) *SearchSimilarityListResponseBodyData {
	s.Uid = &v
	return s
}

type SearchSimilarityListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchSimilarityListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchSimilarityListResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarityListResponse) GoString() string {
	return s.String()
}

func (s *SearchSimilarityListResponse) SetHeaders(v map[string]*string) *SearchSimilarityListResponse {
	s.Headers = v
	return s
}

func (s *SearchSimilarityListResponse) SetStatusCode(v int32) *SearchSimilarityListResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchSimilarityListResponse) SetBody(v *SearchSimilarityListResponseBody) *SearchSimilarityListResponse {
	s.Body = v
	return s
}

type SendMessageToUserRequest struct {
	ReceiverNickName *string                `json:"ReceiverNickName,omitempty" xml:"ReceiverNickName,omitempty"`
	SenderNickName   *string                `json:"SenderNickName,omitempty" xml:"SenderNickName,omitempty"`
	TemplateData     map[string]interface{} `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	TemplateId       *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SendMessageToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToUserRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToUserRequest) SetReceiverNickName(v string) *SendMessageToUserRequest {
	s.ReceiverNickName = &v
	return s
}

func (s *SendMessageToUserRequest) SetSenderNickName(v string) *SendMessageToUserRequest {
	s.SenderNickName = &v
	return s
}

func (s *SendMessageToUserRequest) SetTemplateData(v map[string]interface{}) *SendMessageToUserRequest {
	s.TemplateData = v
	return s
}

func (s *SendMessageToUserRequest) SetTemplateId(v string) *SendMessageToUserRequest {
	s.TemplateId = &v
	return s
}

type SendMessageToUserShrinkRequest struct {
	ReceiverNickName   *string `json:"ReceiverNickName,omitempty" xml:"ReceiverNickName,omitempty"`
	SenderNickName     *string `json:"SenderNickName,omitempty" xml:"SenderNickName,omitempty"`
	TemplateDataShrink *string `json:"TemplateData,omitempty" xml:"TemplateData,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SendMessageToUserShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendMessageToUserShrinkRequest) SetReceiverNickName(v string) *SendMessageToUserShrinkRequest {
	s.ReceiverNickName = &v
	return s
}

func (s *SendMessageToUserShrinkRequest) SetSenderNickName(v string) *SendMessageToUserShrinkRequest {
	s.SenderNickName = &v
	return s
}

func (s *SendMessageToUserShrinkRequest) SetTemplateDataShrink(v string) *SendMessageToUserShrinkRequest {
	s.TemplateDataShrink = &v
	return s
}

func (s *SendMessageToUserShrinkRequest) SetTemplateId(v string) *SendMessageToUserShrinkRequest {
	s.TemplateId = &v
	return s
}

type SendMessageToUserResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMessageToUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToUserResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageToUserResponseBody) SetErrorCode(v string) *SendMessageToUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SendMessageToUserResponseBody) SetErrorMessage(v string) *SendMessageToUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SendMessageToUserResponseBody) SetRequestId(v string) *SendMessageToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageToUserResponseBody) SetSuccess(v bool) *SendMessageToUserResponseBody {
	s.Success = &v
	return s
}

type SendMessageToUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageToUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageToUserResponse) GoString() string {
	return s.String()
}

func (s *SendMessageToUserResponse) SetHeaders(v map[string]*string) *SendMessageToUserResponse {
	s.Headers = v
	return s
}

func (s *SendMessageToUserResponse) SetStatusCode(v int32) *SendMessageToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageToUserResponse) SetBody(v *SendMessageToUserResponseBody) *SendMessageToUserResponse {
	s.Body = v
	return s
}

type SubmitSupplementRequest struct {
	Content      *string                `json:"Content,omitempty" xml:"Content,omitempty"`
	SupplementId *int64                 `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	UserFiles    map[string]interface{} `json:"UserFiles,omitempty" xml:"UserFiles,omitempty"`
}

func (s SubmitSupplementRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementRequest) GoString() string {
	return s.String()
}

func (s *SubmitSupplementRequest) SetContent(v string) *SubmitSupplementRequest {
	s.Content = &v
	return s
}

func (s *SubmitSupplementRequest) SetSupplementId(v int64) *SubmitSupplementRequest {
	s.SupplementId = &v
	return s
}

func (s *SubmitSupplementRequest) SetUserFiles(v map[string]interface{}) *SubmitSupplementRequest {
	s.UserFiles = v
	return s
}

type SubmitSupplementShrinkRequest struct {
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SupplementId    *int64  `json:"SupplementId,omitempty" xml:"SupplementId,omitempty"`
	UserFilesShrink *string `json:"UserFiles,omitempty" xml:"UserFiles,omitempty"`
}

func (s SubmitSupplementShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSupplementShrinkRequest) SetContent(v string) *SubmitSupplementShrinkRequest {
	s.Content = &v
	return s
}

func (s *SubmitSupplementShrinkRequest) SetSupplementId(v int64) *SubmitSupplementShrinkRequest {
	s.SupplementId = &v
	return s
}

func (s *SubmitSupplementShrinkRequest) SetUserFilesShrink(v string) *SubmitSupplementShrinkRequest {
	s.UserFilesShrink = &v
	return s
}

type SubmitSupplementResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitSupplementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSupplementResponseBody) SetCode(v string) *SubmitSupplementResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitSupplementResponseBody) SetMessage(v string) *SubmitSupplementResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSupplementResponseBody) SetRequestId(v string) *SubmitSupplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSupplementResponseBody) SetSuccess(v bool) *SubmitSupplementResponseBody {
	s.Success = &v
	return s
}

type SubmitSupplementResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSupplementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSupplementResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSupplementResponse) GoString() string {
	return s.String()
}

func (s *SubmitSupplementResponse) SetHeaders(v map[string]*string) *SubmitSupplementResponse {
	s.Headers = v
	return s
}

func (s *SubmitSupplementResponse) SetStatusCode(v int32) *SubmitSupplementResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSupplementResponse) SetBody(v *SubmitSupplementResponseBody) *SubmitSupplementResponse {
	s.Body = v
	return s
}

type UpdateApplicantRequest struct {
	Address               *string `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicantId           *int64  `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	ApplicantName         *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	AuthorizationOssKey   *string `json:"AuthorizationOssKey,omitempty" xml:"AuthorizationOssKey,omitempty"`
	BusinessLicenceOssKey *string `json:"BusinessLicenceOssKey,omitempty" xml:"BusinessLicenceOssKey,omitempty"`
	CardNumber            *string `json:"CardNumber,omitempty" xml:"CardNumber,omitempty"`
	ContactAddress        *string `json:"ContactAddress,omitempty" xml:"ContactAddress,omitempty"`
	ContactCity           *string `json:"ContactCity,omitempty" xml:"ContactCity,omitempty"`
	ContactCounty         *string `json:"ContactCounty,omitempty" xml:"ContactCounty,omitempty"`
	ContactDistrict       *string `json:"ContactDistrict,omitempty" xml:"ContactDistrict,omitempty"`
	ContactEmail          *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	ContactName           *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactNumber         *string `json:"ContactNumber,omitempty" xml:"ContactNumber,omitempty"`
	ContactProvince       *string `json:"ContactProvince,omitempty" xml:"ContactProvince,omitempty"`
	ContactZipcode        *string `json:"ContactZipcode,omitempty" xml:"ContactZipcode,omitempty"`
	EAddress              *string `json:"EAddress,omitempty" xml:"EAddress,omitempty"`
	EName                 *string `json:"EName,omitempty" xml:"EName,omitempty"`
	IdCardName            *string `json:"IdCardName,omitempty" xml:"IdCardName,omitempty"`
	IdCardNumber          *string `json:"IdCardNumber,omitempty" xml:"IdCardNumber,omitempty"`
	IdCardOssKey          *string `json:"IdCardOssKey,omitempty" xml:"IdCardOssKey,omitempty"`
	LegalNoticeOssKey     *string `json:"LegalNoticeOssKey,omitempty" xml:"LegalNoticeOssKey,omitempty"`
	PassportOssKey        *string `json:"PassportOssKey,omitempty" xml:"PassportOssKey,omitempty"`
	PersonalType          *int64  `json:"PersonalType,omitempty" xml:"PersonalType,omitempty"`
	Province              *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s UpdateApplicantRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicantRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicantRequest) SetAddress(v string) *UpdateApplicantRequest {
	s.Address = &v
	return s
}

func (s *UpdateApplicantRequest) SetApplicantId(v int64) *UpdateApplicantRequest {
	s.ApplicantId = &v
	return s
}

func (s *UpdateApplicantRequest) SetApplicantName(v string) *UpdateApplicantRequest {
	s.ApplicantName = &v
	return s
}

func (s *UpdateApplicantRequest) SetAuthorizationOssKey(v string) *UpdateApplicantRequest {
	s.AuthorizationOssKey = &v
	return s
}

func (s *UpdateApplicantRequest) SetBusinessLicenceOssKey(v string) *UpdateApplicantRequest {
	s.BusinessLicenceOssKey = &v
	return s
}

func (s *UpdateApplicantRequest) SetCardNumber(v string) *UpdateApplicantRequest {
	s.CardNumber = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactAddress(v string) *UpdateApplicantRequest {
	s.ContactAddress = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactCity(v string) *UpdateApplicantRequest {
	s.ContactCity = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactCounty(v string) *UpdateApplicantRequest {
	s.ContactCounty = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactDistrict(v string) *UpdateApplicantRequest {
	s.ContactDistrict = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactEmail(v string) *UpdateApplicantRequest {
	s.ContactEmail = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactName(v string) *UpdateApplicantRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactNumber(v string) *UpdateApplicantRequest {
	s.ContactNumber = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactProvince(v string) *UpdateApplicantRequest {
	s.ContactProvince = &v
	return s
}

func (s *UpdateApplicantRequest) SetContactZipcode(v string) *UpdateApplicantRequest {
	s.ContactZipcode = &v
	return s
}

func (s *UpdateApplicantRequest) SetEAddress(v string) *UpdateApplicantRequest {
	s.EAddress = &v
	return s
}

func (s *UpdateApplicantRequest) SetEName(v string) *UpdateApplicantRequest {
	s.EName = &v
	return s
}

func (s *UpdateApplicantRequest) SetIdCardName(v string) *UpdateApplicantRequest {
	s.IdCardName = &v
	return s
}

func (s *UpdateApplicantRequest) SetIdCardNumber(v string) *UpdateApplicantRequest {
	s.IdCardNumber = &v
	return s
}

func (s *UpdateApplicantRequest) SetIdCardOssKey(v string) *UpdateApplicantRequest {
	s.IdCardOssKey = &v
	return s
}

func (s *UpdateApplicantRequest) SetLegalNoticeOssKey(v string) *UpdateApplicantRequest {
	s.LegalNoticeOssKey = &v
	return s
}

func (s *UpdateApplicantRequest) SetPassportOssKey(v string) *UpdateApplicantRequest {
	s.PassportOssKey = &v
	return s
}

func (s *UpdateApplicantRequest) SetPersonalType(v int64) *UpdateApplicantRequest {
	s.PersonalType = &v
	return s
}

func (s *UpdateApplicantRequest) SetProvince(v string) *UpdateApplicantRequest {
	s.Province = &v
	return s
}

type UpdateApplicantResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateApplicantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicantResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicantResponseBody) SetRequestId(v string) *UpdateApplicantResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicantResponseBody) SetSuccess(v bool) *UpdateApplicantResponseBody {
	s.Success = &v
	return s
}

type UpdateApplicantResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicantResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicantResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicantResponse) SetHeaders(v map[string]*string) *UpdateApplicantResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicantResponse) SetStatusCode(v int32) *UpdateApplicantResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicantResponse) SetBody(v *UpdateApplicantResponseBody) *UpdateApplicantResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("trademark"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BindApplicantWithOptions(request *BindApplicantRequest, runtime *util.RuntimeOptions) (_result *BindApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantId)) {
		query["ApplicantId"] = request.ApplicantId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationOssKey)) {
		query["AuthorizationOssKey"] = request.AuthorizationOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindApplicant"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindApplicant(request *BindApplicantRequest) (_result *BindApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindApplicantResponse{}
	_body, _err := client.BindApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelOrderWithOptions(request *CancelOrderRequest, runtime *util.RuntimeOptions) (_result *CancelOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOrder"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOrder(request *CancelOrderRequest) (_result *CancelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelOrderResponse{}
	_body, _err := client.CancelOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAuthorizationLetterWithOptions(request *CheckAuthorizationLetterRequest, runtime *util.RuntimeOptions) (_result *CheckAuthorizationLetterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantType)) {
		query["ApplicantType"] = request.ApplicantType
	}

	if !tea.BoolValue(util.IsUnset(request.Color)) {
		query["Color"] = request.Color
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactNumber)) {
		query["ContactNumber"] = request.ContactNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ContactZipcode)) {
		query["ContactZipcode"] = request.ContactZipcode
	}

	if !tea.BoolValue(util.IsUnset(request.OssKey)) {
		query["OssKey"] = request.OssKey
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalType)) {
		query["PersonalType"] = request.PersonalType
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAuthorizationLetter"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAuthorizationLetterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAuthorizationLetter(request *CheckAuthorizationLetterRequest) (_result *CheckAuthorizationLetterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAuthorizationLetterResponse{}
	_body, _err := client.CheckAuthorizationLetterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckBizAvailableWithOptions(request *CheckBizAvailableRequest, runtime *util.RuntimeOptions) (_result *CheckBizAvailableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Biz)) {
		query["Biz"] = request.Biz
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckBizAvailable"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckBizAvailableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckBizAvailable(request *CheckBizAvailableRequest) (_result *CheckBizAvailableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckBizAvailableResponse{}
	_body, _err := client.CheckBizAvailableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDuplicateApplicantRiskWithOptions(request *CheckDuplicateApplicantRiskRequest, runtime *util.RuntimeOptions) (_result *CheckDuplicateApplicantRiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantName)) {
		query["ApplicantName"] = request.ApplicantName
	}

	if !tea.BoolValue(util.IsUnset(request.EventSceneType)) {
		query["EventSceneType"] = request.EventSceneType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDuplicateApplicantRisk"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDuplicateApplicantRiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDuplicateApplicantRisk(request *CheckDuplicateApplicantRiskRequest) (_result *CheckDuplicateApplicantRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDuplicateApplicantRiskResponse{}
	_body, _err := client.CheckDuplicateApplicantRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDuplicateClassificationWithOptions(request *CheckDuplicateClassificationRequest, runtime *util.RuntimeOptions) (_result *CheckDuplicateClassificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Classification)) {
		query["Classification"] = request.Classification
	}

	if !tea.BoolValue(util.IsUnset(request.EventSceneType)) {
		query["EventSceneType"] = request.EventSceneType
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDuplicateClassification"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDuplicateClassificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDuplicateClassification(request *CheckDuplicateClassificationRequest) (_result *CheckDuplicateClassificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDuplicateClassificationResponse{}
	_body, _err := client.CheckDuplicateClassificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDuplicateTrademarkWithOptions(request *CheckDuplicateTrademarkRequest, runtime *util.RuntimeOptions) (_result *CheckDuplicateTrademarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Classification)) {
		query["Classification"] = request.Classification
	}

	if !tea.BoolValue(util.IsUnset(request.EventSceneType)) {
		query["EventSceneType"] = request.EventSceneType
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialName)) {
		query["MaterialName"] = request.MaterialName
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDuplicateTrademark"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDuplicateTrademarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDuplicateTrademark(request *CheckDuplicateTrademarkRequest) (_result *CheckDuplicateTrademarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDuplicateTrademarkResponse{}
	_body, _err := client.CheckDuplicateTrademarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMaterialValidityWithOptions(request *CheckMaterialValidityRequest, runtime *util.RuntimeOptions) (_result *CheckMaterialValidityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessLicenseOssKey)) {
		query["BusinessLicenseOssKey"] = request.BusinessLicenseOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardNumber)) {
		query["CardNumber"] = request.CardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardName)) {
		query["IdCardName"] = request.IdCardName
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNumber)) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardOssKey)) {
		query["IdCardOssKey"] = request.IdCardOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialId)) {
		query["MaterialId"] = request.MaterialId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialRegion)) {
		query["MaterialRegion"] = request.MaterialRegion
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialType)) {
		query["MaterialType"] = request.MaterialType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalType)) {
		query["PersonalType"] = request.PersonalType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckMaterialValidity"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckMaterialValidityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMaterialValidity(request *CheckMaterialValidityRequest) (_result *CheckMaterialValidityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMaterialValidityResponse{}
	_body, _err := client.CheckMaterialValidityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckTrademarkNameWithOptions(request *CheckTrademarkNameRequest, runtime *util.RuntimeOptions) (_result *CheckTrademarkNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventSceneType)) {
		query["EventSceneType"] = request.EventSceneType
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckTrademarkName"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckTrademarkNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckTrademarkName(request *CheckTrademarkNameRequest) (_result *CheckTrademarkNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckTrademarkNameResponse{}
	_body, _err := client.CheckTrademarkNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseTrademarkApplicationWithOptions(request *CloseTrademarkApplicationRequest, runtime *util.RuntimeOptions) (_result *CloseTrademarkApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseTrademarkApplication"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseTrademarkApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTrademarkApplication(request *CloseTrademarkApplicationRequest) (_result *CloseTrademarkApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseTrademarkApplicationResponse{}
	_body, _err := client.CloseTrademarkApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CombineAuthorizationLetterWithOptions(request *CombineAuthorizationLetterRequest, runtime *util.RuntimeOptions) (_result *CombineAuthorizationLetterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantType)) {
		query["ApplicantType"] = request.ApplicantType
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPhone)) {
		query["ContactPhone"] = request.ContactPhone
	}

	if !tea.BoolValue(util.IsUnset(request.ContactPostcode)) {
		query["ContactPostcode"] = request.ContactPostcode
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialId)) {
		query["MaterialId"] = request.MaterialId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialName)) {
		query["MaterialName"] = request.MaterialName
	}

	if !tea.BoolValue(util.IsUnset(request.Nationality)) {
		query["Nationality"] = request.Nationality
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalType)) {
		query["PersonalType"] = request.PersonalType
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.TmProduceType)) {
		query["TmProduceType"] = request.TmProduceType
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CombineAuthorizationLetter"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CombineAuthorizationLetterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CombineAuthorizationLetter(request *CombineAuthorizationLetterRequest) (_result *CombineAuthorizationLetterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CombineAuthorizationLetterResponse{}
	_body, _err := client.CombineAuthorizationLetterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ComplementTrademarkApplicationWithOptions(request *ComplementTrademarkApplicationRequest, runtime *util.RuntimeOptions) (_result *ComplementTrademarkApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgreementId)) {
		query["AgreementId"] = request.AgreementId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationOssKey)) {
		query["AuthorizationOssKey"] = request.AuthorizationOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.IsBlackIcon)) {
		query["IsBlackIcon"] = request.IsBlackIcon
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialId)) {
		query["MaterialId"] = request.MaterialId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderData)) {
		query["OrderData"] = request.OrderData
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkComment)) {
		query["TrademarkComment"] = request.TrademarkComment
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkIconOssKey)) {
		query["TrademarkIconOssKey"] = request.TrademarkIconOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkNameType)) {
		query["TrademarkNameType"] = request.TrademarkNameType
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkType)) {
		query["TrademarkType"] = request.TrademarkType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ComplementTrademarkApplication"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ComplementTrademarkApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ComplementTrademarkApplication(request *ComplementTrademarkApplicationRequest) (_result *ComplementTrademarkApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ComplementTrademarkApplicationResponse{}
	_body, _err := client.ComplementTrademarkApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmExpertSolutionWithOptions(request *ConfirmExpertSolutionRequest, runtime *util.RuntimeOptions) (_result *ConfirmExpertSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmExpertSolution"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmExpertSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmExpertSolution(request *ConfirmExpertSolutionRequest) (_result *ConfirmExpertSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmExpertSolutionResponse{}
	_body, _err := client.ConfirmExpertSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicantWithOptions(request *CreateApplicantRequest, runtime *util.RuntimeOptions) (_result *CreateApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantName)) {
		query["ApplicantName"] = request.ApplicantName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantRegion)) {
		query["ApplicantRegion"] = request.ApplicantRegion
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantType)) {
		query["ApplicantType"] = request.ApplicantType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationOssKey)) {
		query["AuthorizationOssKey"] = request.AuthorizationOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessLicenceOssKey)) {
		query["BusinessLicenceOssKey"] = request.BusinessLicenceOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardNumber)) {
		query["CardNumber"] = request.CardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ContactAddress)) {
		query["ContactAddress"] = request.ContactAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ContactCity)) {
		query["ContactCity"] = request.ContactCity
	}

	if !tea.BoolValue(util.IsUnset(request.ContactCounty)) {
		query["ContactCounty"] = request.ContactCounty
	}

	if !tea.BoolValue(util.IsUnset(request.ContactDistrict)) {
		query["ContactDistrict"] = request.ContactDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.ContactEmail)) {
		query["ContactEmail"] = request.ContactEmail
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactNumber)) {
		query["ContactNumber"] = request.ContactNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ContactProvince)) {
		query["ContactProvince"] = request.ContactProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ContactZipcode)) {
		query["ContactZipcode"] = request.ContactZipcode
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.EAddress)) {
		query["EAddress"] = request.EAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EName)) {
		query["EName"] = request.EName
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardName)) {
		query["IdCardName"] = request.IdCardName
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNumber)) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardOssKey)) {
		query["IdCardOssKey"] = request.IdCardOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.LegalNoticeOssKey)) {
		query["LegalNoticeOssKey"] = request.LegalNoticeOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.PassportOssKey)) {
		query["PassportOssKey"] = request.PassportOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalType)) {
		query["PersonalType"] = request.PersonalType
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplicant"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplicant(request *CreateApplicantRequest) (_result *CreateApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicantResponse{}
	_body, _err := client.CreateApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCommodityOrderWithOptions(tmpReq *CreateCommodityOrderRequest, runtime *util.RuntimeOptions) (_result *CreateCommodityOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCommodityOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Components)) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, tea.String("Components"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OrderParams)) {
		request.OrderParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderParams, tea.String("OrderParams"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentsShrink)) {
		query["Components"] = request.ComponentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderParamsShrink)) {
		query["OrderParams"] = request.OrderParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.Quantity)) {
		query["Quantity"] = request.Quantity
	}

	if !tea.BoolValue(util.IsUnset(request.SpecCode)) {
		query["SpecCode"] = request.SpecCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCommodityOrder"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCommodityOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCommodityOrder(request *CreateCommodityOrderRequest) (_result *CreateCommodityOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCommodityOrderResponse{}
	_body, _err := client.CreateCommodityOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgreementId)) {
		query["AgreementId"] = request.AgreementId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantId)) {
		query["ApplicantId"] = request.ApplicantId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationOssKey)) {
		query["AuthorizationOssKey"] = request.AuthorizationOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.BlackAndWhiteIcon)) {
		query["BlackAndWhiteIcon"] = request.BlackAndWhiteIcon
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Classifications)) {
		query["Classifications"] = request.Classifications
	}

	if !tea.BoolValue(util.IsUnset(request.IdempotentId)) {
		query["IdempotentId"] = request.IdempotentId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalNoticeKey)) {
		query["LegalNoticeKey"] = request.LegalNoticeKey
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentCallback)) {
		query["PaymentCallback"] = request.PaymentCallback
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkComment)) {
		query["TrademarkComment"] = request.TrademarkComment
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkIcon)) {
		query["TrademarkIcon"] = request.TrademarkIcon
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkNameType)) {
		query["TrademarkNameType"] = request.TrademarkNameType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrder"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrademarkApplicationWithOptions(request *CreateTrademarkApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateTrademarkApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgreementId)) {
		query["AgreementId"] = request.AgreementId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantId)) {
		query["ApplicantId"] = request.ApplicantId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationOssKey)) {
		query["AuthorizationOssKey"] = request.AuthorizationOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.BlackAndWhiteIcon)) {
		query["BlackAndWhiteIcon"] = request.BlackAndWhiteIcon
	}

	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["Channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Classifications)) {
		query["Classifications"] = request.Classifications
	}

	if !tea.BoolValue(util.IsUnset(request.IdempotentId)) {
		query["IdempotentId"] = request.IdempotentId
	}

	if !tea.BoolValue(util.IsUnset(request.LegalNoticeKey)) {
		query["LegalNoticeKey"] = request.LegalNoticeKey
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkComment)) {
		query["TrademarkComment"] = request.TrademarkComment
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkIcon)) {
		query["TrademarkIcon"] = request.TrademarkIcon
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkNameType)) {
		query["TrademarkNameType"] = request.TrademarkNameType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrademarkApplication"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrademarkApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrademarkApplication(request *CreateTrademarkApplicationRequest) (_result *CreateTrademarkApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrademarkApplicationResponse{}
	_body, _err := client.CreateTrademarkApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSearchConditionWithOptions(request *DeleteSearchConditionRequest, runtime *util.RuntimeOptions) (_result *DeleteSearchConditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConditionId)) {
		query["ConditionId"] = request.ConditionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Umid)) {
		query["Umid"] = request.Umid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSearchCondition"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSearchConditionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSearchCondition(request *DeleteSearchConditionRequest) (_result *DeleteSearchConditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSearchConditionResponse{}
	_body, _err := client.DeleteSearchConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAdminTrademarkApplicationWithOptions(request *DescribeAdminTrademarkApplicationRequest, runtime *util.RuntimeOptions) (_result *DescribeAdminTrademarkApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAdminTrademarkApplication"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAdminTrademarkApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAdminTrademarkApplication(request *DescribeAdminTrademarkApplicationRequest) (_result *DescribeAdminTrademarkApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAdminTrademarkApplicationResponse{}
	_body, _err := client.DescribeAdminTrademarkApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApplicantWithOptions(request *DescribeApplicantRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantId)) {
		query["ApplicantId"] = request.ApplicantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicant"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApplicant(request *DescribeApplicantRequest) (_result *DescribeApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicantResponse{}
	_body, _err := client.DescribeApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePartnerTrademarkApplicationWithOptions(request *DescribePartnerTrademarkApplicationRequest, runtime *util.RuntimeOptions) (_result *DescribePartnerTrademarkApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePartnerTrademarkApplication"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePartnerTrademarkApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePartnerTrademarkApplication(request *DescribePartnerTrademarkApplicationRequest) (_result *DescribePartnerTrademarkApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePartnerTrademarkApplicationResponse{}
	_body, _err := client.DescribePartnerTrademarkApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQualificationStatusWithOptions(request *DescribeQualificationStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeQualificationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TbUid)) {
		query["TbUid"] = request.TbUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualificationStatus"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualificationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQualificationStatus(request *DescribeQualificationStatusRequest) (_result *DescribeQualificationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeQualificationStatusResponse{}
	_body, _err := client.DescribeQualificationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribeSupplementRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSupplementResponse
 */
func (client *Client) DescribeSupplementWithOptions(request *DescribeSupplementRequest, runtime *util.RuntimeOptions) (_result *DescribeSupplementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SupplementId)) {
		query["SupplementId"] = request.SupplementId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSupplement"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSupplementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ****
 *
 * @param request DescribeSupplementRequest
 * @return DescribeSupplementResponse
 */
func (client *Client) DescribeSupplement(request *DescribeSupplementRequest) (_result *DescribeSupplementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSupplementResponse{}
	_body, _err := client.DescribeSupplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrademarkApplicationWithOptions(request *DescribeTrademarkApplicationRequest, runtime *util.RuntimeOptions) (_result *DescribeTrademarkApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrademarkApplication"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrademarkApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrademarkApplication(request *DescribeTrademarkApplicationRequest) (_result *DescribeTrademarkApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrademarkApplicationResponse{}
	_body, _err := client.DescribeTrademarkApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrademarkDetailForInnerWithOptions(request *DescribeTrademarkDetailForInnerRequest, runtime *util.RuntimeOptions) (_result *DescribeTrademarkDetailForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.Umid)) {
		query["Umid"] = request.Umid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrademarkDetailForInner"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrademarkDetailForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrademarkDetailForInner(request *DescribeTrademarkDetailForInnerRequest) (_result *DescribeTrademarkDetailForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrademarkDetailForInnerResponse{}
	_body, _err := client.DescribeTrademarkDetailForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateUploadFilePolicyWithOptions(request *GenerateUploadFilePolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadFilePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUploadFilePolicy"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateUploadFilePolicy(request *GenerateUploadFilePolicyRequest) (_result *GenerateUploadFilePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.GenerateUploadFilePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlipayUrlWithOptions(request *GetAlipayUrlRequest, runtime *util.RuntimeOptions) (_result *GetAlipayUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlipayUrl"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlipayUrl(request *GetAlipayUrlRequest) (_result *GetAlipayUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.GetAlipayUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrderConfirmUrlWithOptions(request *GetOrderConfirmUrlRequest, runtime *util.RuntimeOptions) (_result *GetOrderConfirmUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Items)) {
		query["Items"] = request.Items
	}

	if !tea.BoolValue(util.IsUnset(request.OutTraceCode)) {
		query["OutTraceCode"] = request.OutTraceCode
	}

	if !tea.BoolValue(util.IsUnset(request.OutTraceType)) {
		query["OutTraceType"] = request.OutTraceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrderConfirmUrl"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrderConfirmUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrderConfirmUrl(request *GetOrderConfirmUrlRequest) (_result *GetOrderConfirmUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrderConfirmUrlResponse{}
	_body, _err := client.GetOrderConfirmUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStsByTaobaoUidWithOptions(request *GetStsByTaobaoUidRequest, runtime *util.RuntimeOptions) (_result *GetStsByTaobaoUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUid)) {
		query["AliyunUid"] = request.AliyunUid
	}

	if !tea.BoolValue(util.IsUnset(request.TbUid)) {
		query["TbUid"] = request.TbUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStsByTaobaoUid"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStsByTaobaoUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStsByTaobaoUid(request *GetStsByTaobaoUidRequest) (_result *GetStsByTaobaoUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStsByTaobaoUidResponse{}
	_body, _err := client.GetStsByTaobaoUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAdminTrademarkApplicationLogsWithOptions(request *ListAdminTrademarkApplicationLogsRequest, runtime *util.RuntimeOptions) (_result *ListAdminTrademarkApplicationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAdminTrademarkApplicationLogs"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAdminTrademarkApplicationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAdminTrademarkApplicationLogs(request *ListAdminTrademarkApplicationLogsRequest) (_result *ListAdminTrademarkApplicationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAdminTrademarkApplicationLogsResponse{}
	_body, _err := client.ListAdminTrademarkApplicationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAdminTrademarkApplicationsWithOptions(request *ListAdminTrademarkApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListAdminTrademarkApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantName)) {
		query["ApplicantName"] = request.ApplicantName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationStatus)) {
		query["ApplicationStatus"] = request.ApplicationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.SupplementStatus)) {
		query["SupplementStatus"] = request.SupplementStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkNumber)) {
		query["TrademarkNumber"] = request.TrademarkNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAdminTrademarkApplications"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAdminTrademarkApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAdminTrademarkApplications(request *ListAdminTrademarkApplicationsRequest) (_result *ListAdminTrademarkApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAdminTrademarkApplicationsResponse{}
	_body, _err := client.ListAdminTrademarkApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicantsWithOptions(request *ListApplicantsRequest, runtime *util.RuntimeOptions) (_result *ListApplicantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantName)) {
		query["ApplicantName"] = request.ApplicantName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantRegion)) {
		query["ApplicantRegion"] = request.ApplicantRegion
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantType)) {
		query["ApplicantType"] = request.ApplicantType
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantVersion)) {
		query["ApplicantVersion"] = request.ApplicantVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CardNumber)) {
		query["CardNumber"] = request.CardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.SystemVersion)) {
		query["SystemVersion"] = request.SystemVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicants"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicants(request *ListApplicantsRequest) (_result *ListApplicantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicantsResponse{}
	_body, _err := client.ListApplicantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAreasWithOptions(request *ListAreasRequest, runtime *util.RuntimeOptions) (_result *ListAreasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCode)) {
		query["ParentCode"] = request.ParentCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAreas"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAreasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAreas(request *ListAreasRequest) (_result *ListAreasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAreasResponse{}
	_body, _err := client.ListAreasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClassificationConditionsWithOptions(request *ListClassificationConditionsRequest, runtime *util.RuntimeOptions) (_result *ListClassificationConditionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		query["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClassificationConditions"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClassificationConditionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClassificationConditions(request *ListClassificationConditionsRequest) (_result *ListClassificationConditionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClassificationConditionsResponse{}
	_body, _err := client.ListClassificationConditionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClassificationsWithOptions(request *ListClassificationsRequest, runtime *util.RuntimeOptions) (_result *ListClassificationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParentCode)) {
		query["ParentCode"] = request.ParentCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClassifications"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClassificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClassifications(request *ListClassificationsRequest) (_result *ListClassificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClassificationsResponse{}
	_body, _err := client.ListClassificationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTrademarkApplicationLogsWithOptions(request *ListTrademarkApplicationLogsRequest, runtime *util.RuntimeOptions) (_result *ListTrademarkApplicationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrademarkApplicationLogs"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrademarkApplicationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTrademarkApplicationLogs(request *ListTrademarkApplicationLogsRequest) (_result *ListTrademarkApplicationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrademarkApplicationLogsResponse{}
	_body, _err := client.ListTrademarkApplicationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTrademarkApplicationsWithOptions(request *ListTrademarkApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListTrademarkApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantName)) {
		query["ApplicantName"] = request.ApplicantName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationStatus)) {
		query["ApplicationStatus"] = request.ApplicationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeLeft)) {
		query["CreateTimeLeft"] = request.CreateTimeLeft
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeRight)) {
		query["CreateTimeRight"] = request.CreateTimeRight
	}

	if !tea.BoolValue(util.IsUnset(request.Flag)) {
		query["Flag"] = request.Flag
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryVoucherOrderDoneFlag)) {
		query["QueryVoucherOrderDoneFlag"] = request.QueryVoucherOrderDoneFlag
	}

	if !tea.BoolValue(util.IsUnset(request.QueryVoucherOrderFlag)) {
		query["QueryVoucherOrderFlag"] = request.QueryVoucherOrderFlag
	}

	if !tea.BoolValue(util.IsUnset(request.SortFiled)) {
		query["SortFiled"] = request.SortFiled
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.SupplementStatus)) {
		query["SupplementStatus"] = request.SupplementStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkName)) {
		query["TrademarkName"] = request.TrademarkName
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkNumber)) {
		query["TrademarkNumber"] = request.TrademarkNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrademarkApplications"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrademarkApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTrademarkApplications(request *ListTrademarkApplicationsRequest) (_result *ListTrademarkApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrademarkApplicationsResponse{}
	_body, _err := client.ListTrademarkApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTrademarkSearchForInnerWithOptions(request *ListTrademarkSearchForInnerRequest, runtime *util.RuntimeOptions) (_result *ListTrademarkSearchForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyBeginTime)) {
		query["ApplyBeginTime"] = request.ApplyBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyEndTime)) {
		query["ApplyEndTime"] = request.ApplyEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Classification)) {
		query["Classification"] = request.Classification
	}

	if !tea.BoolValue(util.IsUnset(request.IfPrecision)) {
		query["IfPrecision"] = request.IfPrecision
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.SearchPreference)) {
		query["SearchPreference"] = request.SearchPreference
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		query["SearchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Umid)) {
		query["Umid"] = request.Umid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTrademarkSearchForInner"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTrademarkSearchForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTrademarkSearchForInner(request *ListTrademarkSearchForInnerRequest) (_result *ListTrademarkSearchForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTrademarkSearchForInnerResponse{}
	_body, _err := client.ListTrademarkSearchForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMeasureDataWithOptions(request *PutMeasureDataRequest, runtime *util.RuntimeOptions) (_result *PutMeasureDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		body["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMeasureData"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMeasureData(request *PutMeasureDataRequest) (_result *PutMeasureDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.PutMeasureDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMeasureReadyFlagWithOptions(request *PutMeasureReadyFlagRequest, runtime *util.RuntimeOptions) (_result *PutMeasureReadyFlagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReadyFlag)) {
		query["ReadyFlag"] = request.ReadyFlag
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutMeasureReadyFlag"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutMeasureReadyFlagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMeasureReadyFlag(request *PutMeasureReadyFlagRequest) (_result *PutMeasureReadyFlagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMeasureReadyFlagResponse{}
	_body, _err := client.PutMeasureReadyFlagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryActivityItemsWithOptions(request *QueryActivityItemsRequest, runtime *util.RuntimeOptions) (_result *QueryActivityItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		query["ActivityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendInfo)) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !tea.BoolValue(util.IsUnset(request.FloorIndex)) {
		query["FloorIndex"] = request.FloorIndex
	}

	if !tea.BoolValue(util.IsUnset(request.Mock)) {
		query["Mock"] = request.Mock
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Refresh)) {
		query["Refresh"] = request.Refresh
	}

	if !tea.BoolValue(util.IsUnset(request.UmId)) {
		query["UmId"] = request.UmId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryActivityItems"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryActivityItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryActivityItems(request *QueryActivityItemsRequest) (_result *QueryActivityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryActivityItemsResponse{}
	_body, _err := client.QueryActivityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAliyunUidWithOptions(request *QueryAliyunUidRequest, runtime *util.RuntimeOptions) (_result *QueryAliyunUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TbUid)) {
		query["TbUid"] = request.TbUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAliyunUid"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAliyunUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAliyunUid(request *QueryAliyunUidRequest) (_result *QueryAliyunUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAliyunUidResponse{}
	_body, _err := client.QueryAliyunUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDetailItemWithOptions(request *QueryDetailItemRequest, runtime *util.RuntimeOptions) (_result *QueryDetailItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DetailConvertType)) {
		query["DetailConvertType"] = request.DetailConvertType
	}

	if !tea.BoolValue(util.IsUnset(request.DetailId)) {
		query["DetailId"] = request.DetailId
	}

	if !tea.BoolValue(util.IsUnset(request.DetailType)) {
		query["DetailType"] = request.DetailType
	}

	if !tea.BoolValue(util.IsUnset(request.Mock)) {
		query["Mock"] = request.Mock
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDetailItem"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDetailItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDetailItem(request *QueryDetailItemRequest) (_result *QueryDetailItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDetailItemResponse{}
	_body, _err := client.QueryDetailItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRemainResourcesWithOptions(request *QueryRemainResourcesRequest, runtime *util.RuntimeOptions) (_result *QueryRemainResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRemainResources"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRemainResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRemainResources(request *QueryRemainResourcesRequest) (_result *QueryRemainResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRemainResourcesResponse{}
	_body, _err := client.QueryRemainResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefuseSupplementWithOptions(request *RefuseSupplementRequest, runtime *util.RuntimeOptions) (_result *RefuseSupplementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SupplementId)) {
		query["SupplementId"] = request.SupplementId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefuseSupplement"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefuseSupplementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefuseSupplement(request *RefuseSupplementRequest) (_result *RefuseSupplementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefuseSupplementResponse{}
	_body, _err := client.RefuseSupplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RejectExpertSolutionWithOptions(request *RejectExpertSolutionRequest, runtime *util.RuntimeOptions) (_result *RejectExpertSolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		query["Note"] = request.Note
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectExpertSolution"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectExpertSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectExpertSolution(request *RejectExpertSolutionRequest) (_result *RejectExpertSolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectExpertSolutionResponse{}
	_body, _err := client.RejectExpertSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveApplicantWithOptions(request *RemoveApplicantRequest, runtime *util.RuntimeOptions) (_result *RemoveApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicantId)) {
		query["ApplicantId"] = request.ApplicantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveApplicant"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveApplicant(request *RemoveApplicantRequest) (_result *RemoveApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveApplicantResponse{}
	_body, _err := client.RemoveApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSearchConditionWithOptions(request *SaveSearchConditionRequest, runtime *util.RuntimeOptions) (_result *SaveSearchConditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConditionContent)) {
		body["ConditionContent"] = request.ConditionContent
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		body["TagName"] = request.TagName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Umid)) {
		body["Umid"] = request.Umid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSearchCondition"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSearchConditionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSearchCondition(request *SaveSearchConditionRequest) (_result *SaveSearchConditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveSearchConditionResponse{}
	_body, _err := client.SaveSearchConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTemporaryApplicantWithOptions(request *SaveTemporaryApplicantRequest, runtime *util.RuntimeOptions) (_result *SaveTemporaryApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantId)) {
		query["ApplicantId"] = request.ApplicantId
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessLicenceOssKey)) {
		query["BusinessLicenceOssKey"] = request.BusinessLicenceOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardNumber)) {
		query["CardNumber"] = request.CardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.City)) {
		query["City"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.CompleteApplicant)) {
		query["CompleteApplicant"] = request.CompleteApplicant
	}

	if !tea.BoolValue(util.IsUnset(request.ContactAddress)) {
		query["ContactAddress"] = request.ContactAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ContactCity)) {
		query["ContactCity"] = request.ContactCity
	}

	if !tea.BoolValue(util.IsUnset(request.ContactCounty)) {
		query["ContactCounty"] = request.ContactCounty
	}

	if !tea.BoolValue(util.IsUnset(request.ContactDistrict)) {
		query["ContactDistrict"] = request.ContactDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.ContactEmail)) {
		query["ContactEmail"] = request.ContactEmail
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactNumber)) {
		query["ContactNumber"] = request.ContactNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ContactProvince)) {
		query["ContactProvince"] = request.ContactProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ContactZipCode)) {
		query["ContactZipCode"] = request.ContactZipCode
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		query["Country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.EAddress)) {
		query["EAddress"] = request.EAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EName)) {
		query["EName"] = request.EName
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardOssKey)) {
		query["IdCardOssKey"] = request.IdCardOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.LegalNoticeOssKey)) {
		query["LegalNoticeOssKey"] = request.LegalNoticeOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.LoaOssKey)) {
		query["LoaOssKey"] = request.LoaOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PassportOssKey)) {
		query["PassportOssKey"] = request.PassportOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalName)) {
		query["PrincipalName"] = request.PrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Town)) {
		query["Town"] = request.Town
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTemporaryApplicant"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTemporaryApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTemporaryApplicant(request *SaveTemporaryApplicantRequest) (_result *SaveTemporaryApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTemporaryApplicantResponse{}
	_body, _err := client.SaveTemporaryApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchItemsWithOptions(request *SearchItemsRequest, runtime *util.RuntimeOptions) (_result *SearchItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludedTags)) {
		query["ExcludedTags"] = request.ExcludedTags
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludedUids)) {
		query["ExcludedUids"] = request.ExcludedUids
	}

	if !tea.BoolValue(util.IsUnset(request.FeedsType)) {
		query["FeedsType"] = request.FeedsType
	}

	if !tea.BoolValue(util.IsUnset(request.IntCls)) {
		query["IntCls"] = request.IntCls
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Mock)) {
		query["Mock"] = request.Mock
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PriceLeft)) {
		query["PriceLeft"] = request.PriceLeft
	}

	if !tea.BoolValue(util.IsUnset(request.PriceRight)) {
		query["PriceRight"] = request.PriceRight
	}

	if !tea.BoolValue(util.IsUnset(request.Products)) {
		query["Products"] = request.Products
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterNumber)) {
		query["RegisterNumber"] = request.RegisterNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkNameLength)) {
		query["TrademarkNameLength"] = request.TrademarkNameLength
	}

	if !tea.BoolValue(util.IsUnset(request.TrademarkNameType)) {
		query["TrademarkNameType"] = request.TrademarkNameType
	}

	if !tea.BoolValue(util.IsUnset(request.UmId)) {
		query["UmId"] = request.UmId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchItems"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchItems(request *SearchItemsRequest) (_result *SearchItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchItemsResponse{}
	_body, _err := client.SearchItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchSimilarityWithOptions(tmpReq *SearchSimilarityRequest, runtime *util.RuntimeOptions) (_result *SearchSimilarityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchSimilarityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Classifications)) {
		request.ClassificationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Classifications, tea.String("Classifications"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SimilarGroups)) {
		request.SimilarGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SimilarGroups, tea.String("SimilarGroups"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassificationsShrink)) {
		query["Classifications"] = request.ClassificationsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NameUriList)) {
		query["NameUriList"] = request.NameUriList
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		query["SearchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDetail)) {
		query["ShowDetail"] = request.ShowDetail
	}

	if !tea.BoolValue(util.IsUnset(request.SimilarGroupsShrink)) {
		query["SimilarGroups"] = request.SimilarGroupsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Sorter)) {
		query["Sorter"] = request.Sorter
	}

	if !tea.BoolValue(util.IsUnset(request.Umid)) {
		query["Umid"] = request.Umid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchSimilarity"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchSimilarityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchSimilarity(request *SearchSimilarityRequest) (_result *SearchSimilarityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchSimilarityResponse{}
	_body, _err := client.SearchSimilarityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchSimilarityListWithOptions(tmpReq *SearchSimilarityListRequest, runtime *util.RuntimeOptions) (_result *SearchSimilarityListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchSimilarityListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Classifications)) {
		request.ClassificationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Classifications, tea.String("Classifications"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SimilarGroups)) {
		request.SimilarGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SimilarGroups, tea.String("SimilarGroups"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClassificationsShrink)) {
		query["Classifications"] = request.ClassificationsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SimilarGroupsShrink)) {
		query["SimilarGroups"] = request.SimilarGroupsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SuccessSearchType)) {
		query["SuccessSearchType"] = request.SuccessSearchType
	}

	if !tea.BoolValue(util.IsUnset(request.Umid)) {
		query["Umid"] = request.Umid
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		query["Uri"] = request.Uri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchSimilarityList"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchSimilarityListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchSimilarityList(request *SearchSimilarityListRequest) (_result *SearchSimilarityListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchSimilarityListResponse{}
	_body, _err := client.SearchSimilarityListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageToUserWithOptions(tmpReq *SendMessageToUserRequest, runtime *util.RuntimeOptions) (_result *SendMessageToUserResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendMessageToUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TemplateData)) {
		request.TemplateDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateData, tea.String("TemplateData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReceiverNickName)) {
		query["ReceiverNickName"] = request.ReceiverNickName
	}

	if !tea.BoolValue(util.IsUnset(request.SenderNickName)) {
		query["SenderNickName"] = request.SenderNickName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateDataShrink)) {
		query["TemplateData"] = request.TemplateDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessageToUser"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessageToUser(request *SendMessageToUserRequest) (_result *SendMessageToUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageToUserResponse{}
	_body, _err := client.SendMessageToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ***
 *
 * @param tmpReq SubmitSupplementRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SubmitSupplementResponse
 */
func (client *Client) SubmitSupplementWithOptions(tmpReq *SubmitSupplementRequest, runtime *util.RuntimeOptions) (_result *SubmitSupplementResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitSupplementShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserFiles)) {
		request.UserFilesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserFiles, tea.String("UserFiles"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.SupplementId)) {
		query["SupplementId"] = request.SupplementId
	}

	if !tea.BoolValue(util.IsUnset(request.UserFilesShrink)) {
		query["UserFiles"] = request.UserFilesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSupplement"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSupplementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ***
 *
 * @param request SubmitSupplementRequest
 * @return SubmitSupplementResponse
 */
func (client *Client) SubmitSupplement(request *SubmitSupplementRequest) (_result *SubmitSupplementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSupplementResponse{}
	_body, _err := client.SubmitSupplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicantWithOptions(request *UpdateApplicantRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		query["Address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantId)) {
		query["ApplicantId"] = request.ApplicantId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicantName)) {
		query["ApplicantName"] = request.ApplicantName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationOssKey)) {
		query["AuthorizationOssKey"] = request.AuthorizationOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessLicenceOssKey)) {
		query["BusinessLicenceOssKey"] = request.BusinessLicenceOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardNumber)) {
		query["CardNumber"] = request.CardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ContactAddress)) {
		query["ContactAddress"] = request.ContactAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ContactCity)) {
		query["ContactCity"] = request.ContactCity
	}

	if !tea.BoolValue(util.IsUnset(request.ContactCounty)) {
		query["ContactCounty"] = request.ContactCounty
	}

	if !tea.BoolValue(util.IsUnset(request.ContactDistrict)) {
		query["ContactDistrict"] = request.ContactDistrict
	}

	if !tea.BoolValue(util.IsUnset(request.ContactEmail)) {
		query["ContactEmail"] = request.ContactEmail
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactNumber)) {
		query["ContactNumber"] = request.ContactNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ContactProvince)) {
		query["ContactProvince"] = request.ContactProvince
	}

	if !tea.BoolValue(util.IsUnset(request.ContactZipcode)) {
		query["ContactZipcode"] = request.ContactZipcode
	}

	if !tea.BoolValue(util.IsUnset(request.EAddress)) {
		query["EAddress"] = request.EAddress
	}

	if !tea.BoolValue(util.IsUnset(request.EName)) {
		query["EName"] = request.EName
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardName)) {
		query["IdCardName"] = request.IdCardName
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardNumber)) {
		query["IdCardNumber"] = request.IdCardNumber
	}

	if !tea.BoolValue(util.IsUnset(request.IdCardOssKey)) {
		query["IdCardOssKey"] = request.IdCardOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.LegalNoticeOssKey)) {
		query["LegalNoticeOssKey"] = request.LegalNoticeOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.PassportOssKey)) {
		query["PassportOssKey"] = request.PassportOssKey
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalType)) {
		query["PersonalType"] = request.PersonalType
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		query["Province"] = request.Province
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicant"),
		Version:     tea.String("2019-09-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplicant(request *UpdateApplicantRequest) (_result *UpdateApplicantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicantResponse{}
	_body, _err := client.UpdateApplicantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
