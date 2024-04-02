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

type AuditItemSubmitRequest struct {
	Data []*AuditItemSubmitRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s AuditItemSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s AuditItemSubmitRequest) GoString() string {
	return s.String()
}

func (s *AuditItemSubmitRequest) SetData(v []*AuditItemSubmitRequestData) *AuditItemSubmitRequest {
	s.Data = v
	return s
}

type AuditItemSubmitRequestData struct {
	CustomResult   *string `json:"CustomResult,omitempty" xml:"CustomResult,omitempty"`
	CustomRiskType *string `json:"CustomRiskType,omitempty" xml:"CustomRiskType,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AuditItemSubmitRequestData) String() string {
	return tea.Prettify(s)
}

func (s AuditItemSubmitRequestData) GoString() string {
	return s.String()
}

func (s *AuditItemSubmitRequestData) SetCustomResult(v string) *AuditItemSubmitRequestData {
	s.CustomResult = &v
	return s
}

func (s *AuditItemSubmitRequestData) SetCustomRiskType(v string) *AuditItemSubmitRequestData {
	s.CustomRiskType = &v
	return s
}

func (s *AuditItemSubmitRequestData) SetId(v int64) *AuditItemSubmitRequestData {
	s.Id = &v
	return s
}

type AuditItemSubmitResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuditItemSubmitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuditItemSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *AuditItemSubmitResponseBody) SetRequestId(v string) *AuditItemSubmitResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuditItemSubmitResponseBody) SetSuccess(v bool) *AuditItemSubmitResponseBody {
	s.Success = &v
	return s
}

type AuditItemSubmitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuditItemSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuditItemSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s AuditItemSubmitResponse) GoString() string {
	return s.String()
}

func (s *AuditItemSubmitResponse) SetHeaders(v map[string]*string) *AuditItemSubmitResponse {
	s.Headers = v
	return s
}

func (s *AuditItemSubmitResponse) SetStatusCode(v int32) *AuditItemSubmitResponse {
	s.StatusCode = &v
	return s
}

func (s *AuditItemSubmitResponse) SetBody(v *AuditItemSubmitResponseBody) *AuditItemSubmitResponse {
	s.Body = v
	return s
}

type CreatCustomOcrTemplateRequest struct {
	ImgUrl        *string `json:"ImgUrl,omitempty" xml:"ImgUrl,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecognizeArea *string `json:"RecognizeArea,omitempty" xml:"RecognizeArea,omitempty"`
	ReferArea     *string `json:"ReferArea,omitempty" xml:"ReferArea,omitempty"`
}

func (s CreatCustomOcrTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatCustomOcrTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreatCustomOcrTemplateRequest) SetImgUrl(v string) *CreatCustomOcrTemplateRequest {
	s.ImgUrl = &v
	return s
}

func (s *CreatCustomOcrTemplateRequest) SetName(v string) *CreatCustomOcrTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreatCustomOcrTemplateRequest) SetRecognizeArea(v string) *CreatCustomOcrTemplateRequest {
	s.RecognizeArea = &v
	return s
}

func (s *CreatCustomOcrTemplateRequest) SetReferArea(v string) *CreatCustomOcrTemplateRequest {
	s.ReferArea = &v
	return s
}

type CreatCustomOcrTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatCustomOcrTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatCustomOcrTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreatCustomOcrTemplateResponseBody) SetRequestId(v string) *CreatCustomOcrTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreatCustomOcrTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatCustomOcrTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatCustomOcrTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatCustomOcrTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreatCustomOcrTemplateResponse) SetHeaders(v map[string]*string) *CreatCustomOcrTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreatCustomOcrTemplateResponse) SetStatusCode(v int32) *CreatCustomOcrTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatCustomOcrTemplateResponse) SetBody(v *CreatCustomOcrTemplateResponseBody) *CreatCustomOcrTemplateResponse {
	s.Body = v
	return s
}

type CreateAuditCallbackRequest struct {
	CallbackSuggestions *string `json:"CallbackSuggestions,omitempty" xml:"CallbackSuggestions,omitempty"`
	CallbackTypes       *string `json:"CallbackTypes,omitempty" xml:"CallbackTypes,omitempty"`
	CryptType           *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Url                 *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateAuditCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuditCallbackRequest) GoString() string {
	return s.String()
}

func (s *CreateAuditCallbackRequest) SetCallbackSuggestions(v string) *CreateAuditCallbackRequest {
	s.CallbackSuggestions = &v
	return s
}

func (s *CreateAuditCallbackRequest) SetCallbackTypes(v string) *CreateAuditCallbackRequest {
	s.CallbackTypes = &v
	return s
}

func (s *CreateAuditCallbackRequest) SetCryptType(v string) *CreateAuditCallbackRequest {
	s.CryptType = &v
	return s
}

func (s *CreateAuditCallbackRequest) SetName(v string) *CreateAuditCallbackRequest {
	s.Name = &v
	return s
}

func (s *CreateAuditCallbackRequest) SetUrl(v string) *CreateAuditCallbackRequest {
	s.Url = &v
	return s
}

type CreateAuditCallbackResponseBody struct {
	CallbackSuggestions []*string `json:"CallbackSuggestions,omitempty" xml:"CallbackSuggestions,omitempty" type:"Repeated"`
	CallbackTypes       []*string `json:"CallbackTypes,omitempty" xml:"CallbackTypes,omitempty" type:"Repeated"`
	CreateTime          *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CryptType           *string   `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	Id                  *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime        *string   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name                *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Seed                *string   `json:"Seed,omitempty" xml:"Seed,omitempty"`
	Url                 *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateAuditCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuditCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuditCallbackResponseBody) SetCallbackSuggestions(v []*string) *CreateAuditCallbackResponseBody {
	s.CallbackSuggestions = v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetCallbackTypes(v []*string) *CreateAuditCallbackResponseBody {
	s.CallbackTypes = v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetCreateTime(v string) *CreateAuditCallbackResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetCryptType(v string) *CreateAuditCallbackResponseBody {
	s.CryptType = &v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetId(v int64) *CreateAuditCallbackResponseBody {
	s.Id = &v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetModifiedTime(v string) *CreateAuditCallbackResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetName(v string) *CreateAuditCallbackResponseBody {
	s.Name = &v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetRequestId(v string) *CreateAuditCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetSeed(v string) *CreateAuditCallbackResponseBody {
	s.Seed = &v
	return s
}

func (s *CreateAuditCallbackResponseBody) SetUrl(v string) *CreateAuditCallbackResponseBody {
	s.Url = &v
	return s
}

type CreateAuditCallbackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuditCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuditCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuditCallbackResponse) GoString() string {
	return s.String()
}

func (s *CreateAuditCallbackResponse) SetHeaders(v map[string]*string) *CreateAuditCallbackResponse {
	s.Headers = v
	return s
}

func (s *CreateAuditCallbackResponse) SetStatusCode(v int32) *CreateAuditCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuditCallbackResponse) SetBody(v *CreateAuditCallbackResponseBody) *CreateAuditCallbackResponse {
	s.Body = v
	return s
}

type CreateBizTypeRequest struct {
	BizTypeImport *string `json:"BizTypeImport,omitempty" xml:"BizTypeImport,omitempty"`
	BizTypeName   *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	CiteTemplate  *bool   `json:"CiteTemplate,omitempty" xml:"CiteTemplate,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IndustryInfo  *string `json:"IndustryInfo,omitempty" xml:"IndustryInfo,omitempty"`
}

func (s CreateBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBizTypeRequest) GoString() string {
	return s.String()
}

func (s *CreateBizTypeRequest) SetBizTypeImport(v string) *CreateBizTypeRequest {
	s.BizTypeImport = &v
	return s
}

func (s *CreateBizTypeRequest) SetBizTypeName(v string) *CreateBizTypeRequest {
	s.BizTypeName = &v
	return s
}

func (s *CreateBizTypeRequest) SetCiteTemplate(v bool) *CreateBizTypeRequest {
	s.CiteTemplate = &v
	return s
}

func (s *CreateBizTypeRequest) SetDescription(v string) *CreateBizTypeRequest {
	s.Description = &v
	return s
}

func (s *CreateBizTypeRequest) SetIndustryInfo(v string) *CreateBizTypeRequest {
	s.IndustryInfo = &v
	return s
}

type CreateBizTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizTypeResponseBody) SetRequestId(v string) *CreateBizTypeResponseBody {
	s.RequestId = &v
	return s
}

type CreateBizTypeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBizTypeResponse) GoString() string {
	return s.String()
}

func (s *CreateBizTypeResponse) SetHeaders(v map[string]*string) *CreateBizTypeResponse {
	s.Headers = v
	return s
}

func (s *CreateBizTypeResponse) SetStatusCode(v int32) *CreateBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBizTypeResponse) SetBody(v *CreateBizTypeResponseBody) *CreateBizTypeResponse {
	s.Body = v
	return s
}

type CreateCdiBagRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	FlowOutSpec   *int32  `json:"FlowOutSpec,omitempty" xml:"FlowOutSpec,omitempty"`
	OrderNum      *int32  `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateCdiBagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCdiBagRequest) GoString() string {
	return s.String()
}

func (s *CreateCdiBagRequest) SetClientToken(v string) *CreateCdiBagRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCdiBagRequest) SetCommodityCode(v string) *CreateCdiBagRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateCdiBagRequest) SetFlowOutSpec(v int32) *CreateCdiBagRequest {
	s.FlowOutSpec = &v
	return s
}

func (s *CreateCdiBagRequest) SetOrderNum(v int32) *CreateCdiBagRequest {
	s.OrderNum = &v
	return s
}

func (s *CreateCdiBagRequest) SetOrderType(v string) *CreateCdiBagRequest {
	s.OrderType = &v
	return s
}

func (s *CreateCdiBagRequest) SetOwnerId(v int64) *CreateCdiBagRequest {
	s.OwnerId = &v
	return s
}

type CreateCdiBagResponseBody struct {
	Code        *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceIds *CreateCdiBagResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	Message     *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId     *string                              `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdiBagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCdiBagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdiBagResponseBody) SetCode(v string) *CreateCdiBagResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCdiBagResponseBody) SetInstanceIds(v *CreateCdiBagResponseBodyInstanceIds) *CreateCdiBagResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateCdiBagResponseBody) SetMessage(v string) *CreateCdiBagResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCdiBagResponseBody) SetOrderId(v string) *CreateCdiBagResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCdiBagResponseBody) SetRequestId(v string) *CreateCdiBagResponseBody {
	s.RequestId = &v
	return s
}

type CreateCdiBagResponseBodyInstanceIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s CreateCdiBagResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s CreateCdiBagResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateCdiBagResponseBodyInstanceIds) SetString_(v []*string) *CreateCdiBagResponseBodyInstanceIds {
	s.String_ = v
	return s
}

type CreateCdiBagResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCdiBagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCdiBagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCdiBagResponse) GoString() string {
	return s.String()
}

func (s *CreateCdiBagResponse) SetHeaders(v map[string]*string) *CreateCdiBagResponse {
	s.Headers = v
	return s
}

func (s *CreateCdiBagResponse) SetStatusCode(v int32) *CreateCdiBagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCdiBagResponse) SetBody(v *CreateCdiBagResponseBody) *CreateCdiBagResponse {
	s.Body = v
	return s
}

type CreateCdiBaseBagRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FlowOutSpec   *int32  `json:"FlowOutSpec,omitempty" xml:"FlowOutSpec,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateCdiBaseBagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCdiBaseBagRequest) GoString() string {
	return s.String()
}

func (s *CreateCdiBaseBagRequest) SetClientToken(v string) *CreateCdiBaseBagRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCdiBaseBagRequest) SetCommodityCode(v string) *CreateCdiBaseBagRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateCdiBaseBagRequest) SetDuration(v int32) *CreateCdiBaseBagRequest {
	s.Duration = &v
	return s
}

func (s *CreateCdiBaseBagRequest) SetFlowOutSpec(v int32) *CreateCdiBaseBagRequest {
	s.FlowOutSpec = &v
	return s
}

func (s *CreateCdiBaseBagRequest) SetOrderType(v string) *CreateCdiBaseBagRequest {
	s.OrderType = &v
	return s
}

func (s *CreateCdiBaseBagRequest) SetOwnerId(v int64) *CreateCdiBaseBagRequest {
	s.OwnerId = &v
	return s
}

type CreateCdiBaseBagResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCdiBaseBagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCdiBaseBagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdiBaseBagResponseBody) SetCode(v string) *CreateCdiBaseBagResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCdiBaseBagResponseBody) SetInstanceId(v string) *CreateCdiBaseBagResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateCdiBaseBagResponseBody) SetMessage(v string) *CreateCdiBaseBagResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCdiBaseBagResponseBody) SetOrderId(v string) *CreateCdiBaseBagResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCdiBaseBagResponseBody) SetRequestId(v string) *CreateCdiBaseBagResponseBody {
	s.RequestId = &v
	return s
}

type CreateCdiBaseBagResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCdiBaseBagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCdiBaseBagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCdiBaseBagResponse) GoString() string {
	return s.String()
}

func (s *CreateCdiBaseBagResponse) SetHeaders(v map[string]*string) *CreateCdiBaseBagResponse {
	s.Headers = v
	return s
}

func (s *CreateCdiBaseBagResponse) SetStatusCode(v int32) *CreateCdiBaseBagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCdiBaseBagResponse) SetBody(v *CreateCdiBaseBagResponseBody) *CreateCdiBaseBagResponse {
	s.Body = v
	return s
}

type CreateImageLibRequest struct {
	BizTypes      *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Enable        *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Scene         *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	ServiceModule *string `json:"ServiceModule,omitempty" xml:"ServiceModule,omitempty"`
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageLibRequest) GoString() string {
	return s.String()
}

func (s *CreateImageLibRequest) SetBizTypes(v string) *CreateImageLibRequest {
	s.BizTypes = &v
	return s
}

func (s *CreateImageLibRequest) SetCategory(v string) *CreateImageLibRequest {
	s.Category = &v
	return s
}

func (s *CreateImageLibRequest) SetEnable(v bool) *CreateImageLibRequest {
	s.Enable = &v
	return s
}

func (s *CreateImageLibRequest) SetName(v string) *CreateImageLibRequest {
	s.Name = &v
	return s
}

func (s *CreateImageLibRequest) SetScene(v string) *CreateImageLibRequest {
	s.Scene = &v
	return s
}

func (s *CreateImageLibRequest) SetServiceModule(v string) *CreateImageLibRequest {
	s.ServiceModule = &v
	return s
}

func (s *CreateImageLibRequest) SetSourceIp(v string) *CreateImageLibRequest {
	s.SourceIp = &v
	return s
}

type CreateImageLibResponseBody struct {
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageLibResponseBody) SetId(v int64) *CreateImageLibResponseBody {
	s.Id = &v
	return s
}

func (s *CreateImageLibResponseBody) SetRequestId(v string) *CreateImageLibResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageLibResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageLibResponse) GoString() string {
	return s.String()
}

func (s *CreateImageLibResponse) SetHeaders(v map[string]*string) *CreateImageLibResponse {
	s.Headers = v
	return s
}

func (s *CreateImageLibResponse) SetStatusCode(v int32) *CreateImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageLibResponse) SetBody(v *CreateImageLibResponseBody) *CreateImageLibResponse {
	s.Body = v
	return s
}

type CreateKeywordRequest struct {
	KeywordLibId *int64  `json:"KeywordLibId,omitempty" xml:"KeywordLibId,omitempty"`
	Keywords     *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeywordRequest) GoString() string {
	return s.String()
}

func (s *CreateKeywordRequest) SetKeywordLibId(v int64) *CreateKeywordRequest {
	s.KeywordLibId = &v
	return s
}

func (s *CreateKeywordRequest) SetKeywords(v string) *CreateKeywordRequest {
	s.Keywords = &v
	return s
}

func (s *CreateKeywordRequest) SetLang(v string) *CreateKeywordRequest {
	s.Lang = &v
	return s
}

func (s *CreateKeywordRequest) SetSourceIp(v string) *CreateKeywordRequest {
	s.SourceIp = &v
	return s
}

type CreateKeywordResponseBody struct {
	InvalidKeywordList []*string                                    `json:"InvalidKeywordList,omitempty" xml:"InvalidKeywordList,omitempty" type:"Repeated"`
	RequestId          *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessCount       *int32                                       `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	ValidKeywordList   []*CreateKeywordResponseBodyValidKeywordList `json:"validKeywordList,omitempty" xml:"validKeywordList,omitempty" type:"Repeated"`
}

func (s CreateKeywordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeywordResponseBody) SetInvalidKeywordList(v []*string) *CreateKeywordResponseBody {
	s.InvalidKeywordList = v
	return s
}

func (s *CreateKeywordResponseBody) SetRequestId(v string) *CreateKeywordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKeywordResponseBody) SetSuccessCount(v int32) *CreateKeywordResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *CreateKeywordResponseBody) SetValidKeywordList(v []*CreateKeywordResponseBodyValidKeywordList) *CreateKeywordResponseBody {
	s.ValidKeywordList = v
	return s
}

type CreateKeywordResponseBodyValidKeywordList struct {
	Id      *int32  `json:"id,omitempty" xml:"id,omitempty"`
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
}

func (s CreateKeywordResponseBodyValidKeywordList) String() string {
	return tea.Prettify(s)
}

func (s CreateKeywordResponseBodyValidKeywordList) GoString() string {
	return s.String()
}

func (s *CreateKeywordResponseBodyValidKeywordList) SetId(v int32) *CreateKeywordResponseBodyValidKeywordList {
	s.Id = &v
	return s
}

func (s *CreateKeywordResponseBodyValidKeywordList) SetKeyword(v string) *CreateKeywordResponseBodyValidKeywordList {
	s.Keyword = &v
	return s
}

type CreateKeywordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeywordResponse) GoString() string {
	return s.String()
}

func (s *CreateKeywordResponse) SetHeaders(v map[string]*string) *CreateKeywordResponse {
	s.Headers = v
	return s
}

func (s *CreateKeywordResponse) SetStatusCode(v int32) *CreateKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeywordResponse) SetBody(v *CreateKeywordResponseBody) *CreateKeywordResponse {
	s.Body = v
	return s
}

type CreateKeywordLibRequest struct {
	BizTypes      *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Enable        *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LibType       *string `json:"LibType,omitempty" xml:"LibType,omitempty"`
	MatchMode     *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceModule *string `json:"ServiceModule,omitempty" xml:"ServiceModule,omitempty"`
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *CreateKeywordLibRequest) SetBizTypes(v string) *CreateKeywordLibRequest {
	s.BizTypes = &v
	return s
}

func (s *CreateKeywordLibRequest) SetCategory(v string) *CreateKeywordLibRequest {
	s.Category = &v
	return s
}

func (s *CreateKeywordLibRequest) SetEnable(v bool) *CreateKeywordLibRequest {
	s.Enable = &v
	return s
}

func (s *CreateKeywordLibRequest) SetLang(v string) *CreateKeywordLibRequest {
	s.Lang = &v
	return s
}

func (s *CreateKeywordLibRequest) SetLanguage(v string) *CreateKeywordLibRequest {
	s.Language = &v
	return s
}

func (s *CreateKeywordLibRequest) SetLibType(v string) *CreateKeywordLibRequest {
	s.LibType = &v
	return s
}

func (s *CreateKeywordLibRequest) SetMatchMode(v string) *CreateKeywordLibRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateKeywordLibRequest) SetName(v string) *CreateKeywordLibRequest {
	s.Name = &v
	return s
}

func (s *CreateKeywordLibRequest) SetResourceType(v string) *CreateKeywordLibRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateKeywordLibRequest) SetServiceModule(v string) *CreateKeywordLibRequest {
	s.ServiceModule = &v
	return s
}

func (s *CreateKeywordLibRequest) SetSourceIp(v string) *CreateKeywordLibRequest {
	s.SourceIp = &v
	return s
}

type CreateKeywordLibResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeywordLibResponseBody) SetId(v string) *CreateKeywordLibResponseBody {
	s.Id = &v
	return s
}

func (s *CreateKeywordLibResponseBody) SetRequestId(v string) *CreateKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

type CreateKeywordLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *CreateKeywordLibResponse) SetHeaders(v map[string]*string) *CreateKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *CreateKeywordLibResponse) SetStatusCode(v int32) *CreateKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeywordLibResponse) SetBody(v *CreateKeywordLibResponseBody) *CreateKeywordLibResponse {
	s.Body = v
	return s
}

type CreateWebSiteInstanceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OrderNum     *int32  `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s CreateWebSiteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWebSiteInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateWebSiteInstanceRequest) SetClientToken(v string) *CreateWebSiteInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWebSiteInstanceRequest) SetDuration(v int32) *CreateWebSiteInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateWebSiteInstanceRequest) SetOrderNum(v int32) *CreateWebSiteInstanceRequest {
	s.OrderNum = &v
	return s
}

func (s *CreateWebSiteInstanceRequest) SetOrderType(v string) *CreateWebSiteInstanceRequest {
	s.OrderType = &v
	return s
}

func (s *CreateWebSiteInstanceRequest) SetOwnerId(v int64) *CreateWebSiteInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateWebSiteInstanceRequest) SetPricingCycle(v string) *CreateWebSiteInstanceRequest {
	s.PricingCycle = &v
	return s
}

type CreateWebSiteInstanceResponseBody struct {
	Code        *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceIds *CreateWebSiteInstanceResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	Message     *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId     *string                                       `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebSiteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWebSiteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebSiteInstanceResponseBody) SetCode(v string) *CreateWebSiteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWebSiteInstanceResponseBody) SetInstanceIds(v *CreateWebSiteInstanceResponseBodyInstanceIds) *CreateWebSiteInstanceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateWebSiteInstanceResponseBody) SetMessage(v string) *CreateWebSiteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWebSiteInstanceResponseBody) SetOrderId(v string) *CreateWebSiteInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateWebSiteInstanceResponseBody) SetRequestId(v string) *CreateWebSiteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateWebSiteInstanceResponseBodyInstanceIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s CreateWebSiteInstanceResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s CreateWebSiteInstanceResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateWebSiteInstanceResponseBodyInstanceIds) SetString_(v []*string) *CreateWebSiteInstanceResponseBodyInstanceIds {
	s.String_ = v
	return s
}

type CreateWebSiteInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebSiteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebSiteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWebSiteInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateWebSiteInstanceResponse) SetHeaders(v map[string]*string) *CreateWebSiteInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateWebSiteInstanceResponse) SetStatusCode(v int32) *CreateWebSiteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebSiteInstanceResponse) SetBody(v *CreateWebSiteInstanceResponseBody) *CreateWebSiteInstanceResponse {
	s.Body = v
	return s
}

type CreateWebsiteIndexPageBaselineRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateWebsiteIndexPageBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWebsiteIndexPageBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreateWebsiteIndexPageBaselineRequest) SetInstanceId(v string) *CreateWebsiteIndexPageBaselineRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWebsiteIndexPageBaselineRequest) SetLang(v string) *CreateWebsiteIndexPageBaselineRequest {
	s.Lang = &v
	return s
}

func (s *CreateWebsiteIndexPageBaselineRequest) SetSourceIp(v string) *CreateWebsiteIndexPageBaselineRequest {
	s.SourceIp = &v
	return s
}

type CreateWebsiteIndexPageBaselineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebsiteIndexPageBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWebsiteIndexPageBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebsiteIndexPageBaselineResponseBody) SetRequestId(v string) *CreateWebsiteIndexPageBaselineResponseBody {
	s.RequestId = &v
	return s
}

type CreateWebsiteIndexPageBaselineResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWebsiteIndexPageBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWebsiteIndexPageBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWebsiteIndexPageBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreateWebsiteIndexPageBaselineResponse) SetHeaders(v map[string]*string) *CreateWebsiteIndexPageBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreateWebsiteIndexPageBaselineResponse) SetStatusCode(v int32) *CreateWebsiteIndexPageBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebsiteIndexPageBaselineResponse) SetBody(v *CreateWebsiteIndexPageBaselineResponseBody) *CreateWebsiteIndexPageBaselineResponse {
	s.Body = v
	return s
}

type DeleteBizTypeRequest struct {
	BizTypeName *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
}

func (s DeleteBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizTypeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizTypeRequest) SetBizTypeName(v string) *DeleteBizTypeRequest {
	s.BizTypeName = &v
	return s
}

type DeleteBizTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizTypeResponseBody) SetRequestId(v string) *DeleteBizTypeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBizTypeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizTypeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizTypeResponse) SetHeaders(v map[string]*string) *DeleteBizTypeResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizTypeResponse) SetStatusCode(v int32) *DeleteBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBizTypeResponse) SetBody(v *DeleteBizTypeResponseBody) *DeleteBizTypeResponse {
	s.Body = v
	return s
}

type DeleteCustomOcrTemplateRequest struct {
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DeleteCustomOcrTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomOcrTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomOcrTemplateRequest) SetIds(v string) *DeleteCustomOcrTemplateRequest {
	s.Ids = &v
	return s
}

type DeleteCustomOcrTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomOcrTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomOcrTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomOcrTemplateResponseBody) SetRequestId(v string) *DeleteCustomOcrTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomOcrTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomOcrTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomOcrTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomOcrTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomOcrTemplateResponse) SetHeaders(v map[string]*string) *DeleteCustomOcrTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomOcrTemplateResponse) SetStatusCode(v int32) *DeleteCustomOcrTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomOcrTemplateResponse) SetBody(v *DeleteCustomOcrTemplateResponseBody) *DeleteCustomOcrTemplateResponse {
	s.Body = v
	return s
}

type DeleteImageFromLibRequest struct {
	Ids      *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteImageFromLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageFromLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageFromLibRequest) SetIds(v string) *DeleteImageFromLibRequest {
	s.Ids = &v
	return s
}

func (s *DeleteImageFromLibRequest) SetSourceIp(v string) *DeleteImageFromLibRequest {
	s.SourceIp = &v
	return s
}

type DeleteImageFromLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageFromLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageFromLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageFromLibResponseBody) SetRequestId(v string) *DeleteImageFromLibResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageFromLibResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageFromLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageFromLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageFromLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageFromLibResponse) SetHeaders(v map[string]*string) *DeleteImageFromLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageFromLibResponse) SetStatusCode(v int32) *DeleteImageFromLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageFromLibResponse) SetBody(v *DeleteImageFromLibResponseBody) *DeleteImageFromLibResponse {
	s.Body = v
	return s
}

type DeleteImageLibRequest struct {
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageLibRequest) SetId(v int32) *DeleteImageLibRequest {
	s.Id = &v
	return s
}

func (s *DeleteImageLibRequest) SetSourceIp(v string) *DeleteImageLibRequest {
	s.SourceIp = &v
	return s
}

type DeleteImageLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageLibResponseBody) SetRequestId(v string) *DeleteImageLibResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageLibResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageLibResponse) SetHeaders(v map[string]*string) *DeleteImageLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageLibResponse) SetStatusCode(v int32) *DeleteImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageLibResponse) SetBody(v *DeleteImageLibResponseBody) *DeleteImageLibResponse {
	s.Body = v
	return s
}

type DeleteKeywordRequest struct {
	Ids          *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	KeywordLibId *string `json:"KeywordLibId,omitempty" xml:"KeywordLibId,omitempty"`
	Keywords     *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeywordRequest) SetIds(v string) *DeleteKeywordRequest {
	s.Ids = &v
	return s
}

func (s *DeleteKeywordRequest) SetKeywordLibId(v string) *DeleteKeywordRequest {
	s.KeywordLibId = &v
	return s
}

func (s *DeleteKeywordRequest) SetKeywords(v string) *DeleteKeywordRequest {
	s.Keywords = &v
	return s
}

func (s *DeleteKeywordRequest) SetLang(v string) *DeleteKeywordRequest {
	s.Lang = &v
	return s
}

func (s *DeleteKeywordRequest) SetSourceIp(v string) *DeleteKeywordRequest {
	s.SourceIp = &v
	return s
}

type DeleteKeywordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeywordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeywordResponseBody) SetRequestId(v string) *DeleteKeywordResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKeywordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeywordResponse) SetHeaders(v map[string]*string) *DeleteKeywordResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeywordResponse) SetStatusCode(v int32) *DeleteKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeywordResponse) SetBody(v *DeleteKeywordResponseBody) *DeleteKeywordResponse {
	s.Body = v
	return s
}

type DeleteKeywordLibRequest struct {
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibRequest) SetId(v int32) *DeleteKeywordLibRequest {
	s.Id = &v
	return s
}

func (s *DeleteKeywordLibRequest) SetLang(v string) *DeleteKeywordLibRequest {
	s.Lang = &v
	return s
}

func (s *DeleteKeywordLibRequest) SetSourceIp(v string) *DeleteKeywordLibRequest {
	s.SourceIp = &v
	return s
}

type DeleteKeywordLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibResponseBody) SetRequestId(v string) *DeleteKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKeywordLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeywordLibResponse) SetHeaders(v map[string]*string) *DeleteKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeywordLibResponse) SetStatusCode(v int32) *DeleteKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeywordLibResponse) SetBody(v *DeleteKeywordLibResponseBody) *DeleteKeywordLibResponse {
	s.Body = v
	return s
}

type DeleteNotificationContactsRequest struct {
	ContactTypes *string `json:"ContactTypes,omitempty" xml:"ContactTypes,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteNotificationContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationContactsRequest) GoString() string {
	return s.String()
}

func (s *DeleteNotificationContactsRequest) SetContactTypes(v string) *DeleteNotificationContactsRequest {
	s.ContactTypes = &v
	return s
}

func (s *DeleteNotificationContactsRequest) SetLang(v string) *DeleteNotificationContactsRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNotificationContactsRequest) SetSourceIp(v string) *DeleteNotificationContactsRequest {
	s.SourceIp = &v
	return s
}

type DeleteNotificationContactsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNotificationContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNotificationContactsResponseBody) SetRequestId(v string) *DeleteNotificationContactsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNotificationContactsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNotificationContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNotificationContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationContactsResponse) GoString() string {
	return s.String()
}

func (s *DeleteNotificationContactsResponse) SetHeaders(v map[string]*string) *DeleteNotificationContactsResponse {
	s.Headers = v
	return s
}

func (s *DeleteNotificationContactsResponse) SetStatusCode(v int32) *DeleteNotificationContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNotificationContactsResponse) SetBody(v *DeleteNotificationContactsResponseBody) *DeleteNotificationContactsResponse {
	s.Body = v
	return s
}

type DescribeAppInfoRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Platform    *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoRequest) SetCurrentPage(v int32) *DescribeAppInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppInfoRequest) SetLang(v string) *DescribeAppInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAppInfoRequest) SetPageSize(v int32) *DescribeAppInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInfoRequest) SetPlatform(v string) *DescribeAppInfoRequest {
	s.Platform = &v
	return s
}

func (s *DescribeAppInfoRequest) SetSourceIp(v string) *DescribeAppInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAppInfoRequest) SetTotalCount(v int32) *DescribeAppInfoRequest {
	s.TotalCount = &v
	return s
}

type DescribeAppInfoResponseBody struct {
	AppInfoList []*DescribeAppInfoResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBody) SetAppInfoList(v []*DescribeAppInfoResponseBodyAppInfoList) *DescribeAppInfoResponseBody {
	s.AppInfoList = v
	return s
}

func (s *DescribeAppInfoResponseBody) SetCurrentPage(v int32) *DescribeAppInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAppInfoResponseBody) SetPageSize(v int32) *DescribeAppInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInfoResponseBody) SetRequestId(v string) *DescribeAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppInfoResponseBody) SetTotalCount(v int32) *DescribeAppInfoResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppInfoResponseBodyAppInfoList struct {
	DebugPackageInfo *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo `json:"DebugPackageInfo,omitempty" xml:"DebugPackageInfo,omitempty" type:"Struct"`
	EndDate          *string                                                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Icon             *string                                                 `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Id               *int64                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PackageInfo      *DescribeAppInfoResponseBodyAppInfoListPackageInfo      `json:"PackageInfo,omitempty" xml:"PackageInfo,omitempty" type:"Struct"`
	PackageName      *string                                                 `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	StartDate        *string                                                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Type             *int32                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppInfoResponseBodyAppInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetDebugPackageInfo(v *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) *DescribeAppInfoResponseBodyAppInfoList {
	s.DebugPackageInfo = v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetEndDate(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.EndDate = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetIcon(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.Icon = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetId(v int64) *DescribeAppInfoResponseBodyAppInfoList {
	s.Id = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetName(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.Name = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetPackageInfo(v *DescribeAppInfoResponseBodyAppInfoListPackageInfo) *DescribeAppInfoResponseBodyAppInfoList {
	s.PackageInfo = v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetPackageName(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.PackageName = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetStartDate(v string) *DescribeAppInfoResponseBodyAppInfoList {
	s.StartDate = &v
	return s
}

func (s *DescribeAppInfoResponseBodyAppInfoList) SetType(v int32) *DescribeAppInfoResponseBodyAppInfoList {
	s.Type = &v
	return s
}

type DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo) SetVersion(v string) *DescribeAppInfoResponseBodyAppInfoListDebugPackageInfo {
	s.Version = &v
	return s
}

type DescribeAppInfoResponseBodyAppInfoListPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppInfoResponseBodyAppInfoListPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponseBodyAppInfoListPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponseBodyAppInfoListPackageInfo) SetVersion(v string) *DescribeAppInfoResponseBodyAppInfoListPackageInfo {
	s.Version = &v
	return s
}

type DescribeAppInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppInfoResponse) SetHeaders(v map[string]*string) *DescribeAppInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppInfoResponse) SetStatusCode(v int32) *DescribeAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppInfoResponse) SetBody(v *DescribeAppInfoResponseBody) *DescribeAppInfoResponse {
	s.Body = v
	return s
}

type DescribeAuditCallbackResponseBody struct {
	Callback  *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CryptType *int32  `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Seed      *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
}

func (s DescribeAuditCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditCallbackResponseBody) SetCallback(v string) *DescribeAuditCallbackResponseBody {
	s.Callback = &v
	return s
}

func (s *DescribeAuditCallbackResponseBody) SetCryptType(v int32) *DescribeAuditCallbackResponseBody {
	s.CryptType = &v
	return s
}

func (s *DescribeAuditCallbackResponseBody) SetRequestId(v string) *DescribeAuditCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditCallbackResponseBody) SetSeed(v string) *DescribeAuditCallbackResponseBody {
	s.Seed = &v
	return s
}

type DescribeAuditCallbackResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditCallbackResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditCallbackResponse) SetHeaders(v map[string]*string) *DescribeAuditCallbackResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditCallbackResponse) SetStatusCode(v int32) *DescribeAuditCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditCallbackResponse) SetBody(v *DescribeAuditCallbackResponseBody) *DescribeAuditCallbackResponse {
	s.Body = v
	return s
}

type DescribeAuditCallbackListResponseBody struct {
	CallbackList []*DescribeAuditCallbackListResponseBodyCallbackList `json:"CallbackList,omitempty" xml:"CallbackList,omitempty" type:"Repeated"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditCallbackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditCallbackListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditCallbackListResponseBody) SetCallbackList(v []*DescribeAuditCallbackListResponseBodyCallbackList) *DescribeAuditCallbackListResponseBody {
	s.CallbackList = v
	return s
}

func (s *DescribeAuditCallbackListResponseBody) SetRequestId(v string) *DescribeAuditCallbackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditCallbackListResponseBody) SetTotalCount(v int32) *DescribeAuditCallbackListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuditCallbackListResponseBodyCallbackList struct {
	CallbackSuggestions []*string `json:"CallbackSuggestions,omitempty" xml:"CallbackSuggestions,omitempty" type:"Repeated"`
	CallbackTypes       []*string `json:"CallbackTypes,omitempty" xml:"CallbackTypes,omitempty" type:"Repeated"`
	CreateTime          *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CryptType           *string   `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	Id                  *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime        *string   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name                *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Seed                *string   `json:"Seed,omitempty" xml:"Seed,omitempty"`
	Url                 *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeAuditCallbackListResponseBodyCallbackList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditCallbackListResponseBodyCallbackList) GoString() string {
	return s.String()
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetCallbackSuggestions(v []*string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.CallbackSuggestions = v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetCallbackTypes(v []*string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.CallbackTypes = v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetCreateTime(v string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.CreateTime = &v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetCryptType(v string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.CryptType = &v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetId(v int64) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.Id = &v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetModifiedTime(v string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetName(v string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.Name = &v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetSeed(v string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.Seed = &v
	return s
}

func (s *DescribeAuditCallbackListResponseBodyCallbackList) SetUrl(v string) *DescribeAuditCallbackListResponseBodyCallbackList {
	s.Url = &v
	return s
}

type DescribeAuditCallbackListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditCallbackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditCallbackListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditCallbackListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditCallbackListResponse) SetHeaders(v map[string]*string) *DescribeAuditCallbackListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditCallbackListResponse) SetStatusCode(v int32) *DescribeAuditCallbackListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditCallbackListResponse) SetBody(v *DescribeAuditCallbackListResponseBody) *DescribeAuditCallbackListResponse {
	s.Body = v
	return s
}

type DescribeAuditContentRequest struct {
	AuditResult  *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DataId       *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeywordId    *string `json:"KeywordId,omitempty" xml:"KeywordId,omitempty"`
	Label        *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LibType      *string `json:"LibType,omitempty" xml:"LibType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Suggestion   *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount   *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentRequest) SetAuditResult(v string) *DescribeAuditContentRequest {
	s.AuditResult = &v
	return s
}

func (s *DescribeAuditContentRequest) SetBizType(v string) *DescribeAuditContentRequest {
	s.BizType = &v
	return s
}

func (s *DescribeAuditContentRequest) SetCurrentPage(v int32) *DescribeAuditContentRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditContentRequest) SetDataId(v string) *DescribeAuditContentRequest {
	s.DataId = &v
	return s
}

func (s *DescribeAuditContentRequest) SetEndDate(v string) *DescribeAuditContentRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeAuditContentRequest) SetImageId(v string) *DescribeAuditContentRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeAuditContentRequest) SetKeywordId(v string) *DescribeAuditContentRequest {
	s.KeywordId = &v
	return s
}

func (s *DescribeAuditContentRequest) SetLabel(v string) *DescribeAuditContentRequest {
	s.Label = &v
	return s
}

func (s *DescribeAuditContentRequest) SetLang(v string) *DescribeAuditContentRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuditContentRequest) SetLibType(v string) *DescribeAuditContentRequest {
	s.LibType = &v
	return s
}

func (s *DescribeAuditContentRequest) SetPageSize(v int32) *DescribeAuditContentRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditContentRequest) SetResourceType(v string) *DescribeAuditContentRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeAuditContentRequest) SetScene(v string) *DescribeAuditContentRequest {
	s.Scene = &v
	return s
}

func (s *DescribeAuditContentRequest) SetSourceIp(v string) *DescribeAuditContentRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAuditContentRequest) SetStartDate(v string) *DescribeAuditContentRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeAuditContentRequest) SetSuggestion(v string) *DescribeAuditContentRequest {
	s.Suggestion = &v
	return s
}

func (s *DescribeAuditContentRequest) SetTaskId(v string) *DescribeAuditContentRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeAuditContentRequest) SetTotalCount(v int32) *DescribeAuditContentRequest {
	s.TotalCount = &v
	return s
}

type DescribeAuditContentResponseBody struct {
	AuditContentList []*DescribeAuditContentResponseBodyAuditContentList `json:"AuditContentList,omitempty" xml:"AuditContentList,omitempty" type:"Repeated"`
	CurrentPage      *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize         *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentResponseBody) SetAuditContentList(v []*DescribeAuditContentResponseBodyAuditContentList) *DescribeAuditContentResponseBody {
	s.AuditContentList = v
	return s
}

func (s *DescribeAuditContentResponseBody) SetCurrentPage(v int32) *DescribeAuditContentResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditContentResponseBody) SetPageSize(v int32) *DescribeAuditContentResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditContentResponseBody) SetRequestId(v string) *DescribeAuditContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditContentResponseBody) SetTotalCount(v int32) *DescribeAuditContentResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuditContentResponseBodyAuditContentList struct {
	Audit               *int32                                                          `json:"Audit,omitempty" xml:"Audit,omitempty"`
	AuditIllegalReasons []*string                                                       `json:"AuditIllegalReasons,omitempty" xml:"AuditIllegalReasons,omitempty" type:"Repeated"`
	AuditResult         *string                                                         `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	BizType             *string                                                         `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Content             *string                                                         `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId              *string                                                         `json:"DataId,omitempty" xml:"DataId,omitempty"`
	FrameResults        []*DescribeAuditContentResponseBodyAuditContentListFrameResults `json:"FrameResults,omitempty" xml:"FrameResults,omitempty" type:"Repeated"`
	Id                  *int64                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	NewUrl              *string                                                         `json:"NewUrl,omitempty" xml:"NewUrl,omitempty"`
	RequestTime         *string                                                         `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	Results             []*DescribeAuditContentResponseBodyAuditContentListResults      `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	ScanFinishedTime    *string                                                         `json:"ScanFinishedTime,omitempty" xml:"ScanFinishedTime,omitempty"`
	Suggestion          *string                                                         `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TaskId              *string                                                         `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Thumbnail           *string                                                         `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	Url                 *string                                                         `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeAuditContentResponseBodyAuditContentList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentResponseBodyAuditContentList) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetAudit(v int32) *DescribeAuditContentResponseBodyAuditContentList {
	s.Audit = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetAuditIllegalReasons(v []*string) *DescribeAuditContentResponseBodyAuditContentList {
	s.AuditIllegalReasons = v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetAuditResult(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.AuditResult = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetBizType(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.BizType = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetContent(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.Content = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetDataId(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.DataId = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetFrameResults(v []*DescribeAuditContentResponseBodyAuditContentListFrameResults) *DescribeAuditContentResponseBodyAuditContentList {
	s.FrameResults = v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetId(v int64) *DescribeAuditContentResponseBodyAuditContentList {
	s.Id = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetNewUrl(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.NewUrl = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetRequestTime(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.RequestTime = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetResults(v []*DescribeAuditContentResponseBodyAuditContentListResults) *DescribeAuditContentResponseBodyAuditContentList {
	s.Results = v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetScanFinishedTime(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.ScanFinishedTime = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetSuggestion(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.Suggestion = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetTaskId(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.TaskId = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetThumbnail(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.Thumbnail = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentList) SetUrl(v string) *DescribeAuditContentResponseBodyAuditContentList {
	s.Url = &v
	return s
}

type DescribeAuditContentResponseBodyAuditContentListFrameResults struct {
	Label  *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Offset *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeAuditContentResponseBodyAuditContentListFrameResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentResponseBodyAuditContentListFrameResults) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentResponseBodyAuditContentListFrameResults) SetLabel(v string) *DescribeAuditContentResponseBodyAuditContentListFrameResults {
	s.Label = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentListFrameResults) SetOffset(v int32) *DescribeAuditContentResponseBodyAuditContentListFrameResults {
	s.Offset = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentListFrameResults) SetUrl(v string) *DescribeAuditContentResponseBodyAuditContentListFrameResults {
	s.Url = &v
	return s
}

type DescribeAuditContentResponseBodyAuditContentListResults struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DescribeAuditContentResponseBodyAuditContentListResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentResponseBodyAuditContentListResults) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentResponseBodyAuditContentListResults) SetLabel(v string) *DescribeAuditContentResponseBodyAuditContentListResults {
	s.Label = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentListResults) SetScene(v string) *DescribeAuditContentResponseBodyAuditContentListResults {
	s.Scene = &v
	return s
}

func (s *DescribeAuditContentResponseBodyAuditContentListResults) SetSuggestion(v string) *DescribeAuditContentResponseBodyAuditContentListResults {
	s.Suggestion = &v
	return s
}

type DescribeAuditContentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentResponse) SetHeaders(v map[string]*string) *DescribeAuditContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditContentResponse) SetStatusCode(v int32) *DescribeAuditContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditContentResponse) SetBody(v *DescribeAuditContentResponseBody) *DescribeAuditContentResponse {
	s.Body = v
	return s
}

type DescribeAuditContentItemRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount   *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditContentItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentItemRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentItemRequest) SetCurrentPage(v int32) *DescribeAuditContentItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditContentItemRequest) SetLang(v string) *DescribeAuditContentItemRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuditContentItemRequest) SetPageSize(v int32) *DescribeAuditContentItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditContentItemRequest) SetResourceType(v string) *DescribeAuditContentItemRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeAuditContentItemRequest) SetSourceIp(v string) *DescribeAuditContentItemRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAuditContentItemRequest) SetTaskId(v string) *DescribeAuditContentItemRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeAuditContentItemRequest) SetTotalCount(v int32) *DescribeAuditContentItemRequest {
	s.TotalCount = &v
	return s
}

type DescribeAuditContentItemResponseBody struct {
	AuditContentItemList []*DescribeAuditContentItemResponseBodyAuditContentItemList `json:"AuditContentItemList,omitempty" xml:"AuditContentItemList,omitempty" type:"Repeated"`
	CurrentPage          *int32                                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize             *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuditContentItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentItemResponseBody) SetAuditContentItemList(v []*DescribeAuditContentItemResponseBodyAuditContentItemList) *DescribeAuditContentItemResponseBody {
	s.AuditContentItemList = v
	return s
}

func (s *DescribeAuditContentItemResponseBody) SetCurrentPage(v int32) *DescribeAuditContentItemResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAuditContentItemResponseBody) SetPageSize(v int32) *DescribeAuditContentItemResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuditContentItemResponseBody) SetRequestId(v string) *DescribeAuditContentItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditContentItemResponseBody) SetTotalCount(v int32) *DescribeAuditContentItemResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuditContentItemResponseBodyAuditContentItemList struct {
	Audit               *int32    `json:"Audit,omitempty" xml:"Audit,omitempty"`
	AuditIllegalReasons []*string `json:"AuditIllegalReasons,omitempty" xml:"AuditIllegalReasons,omitempty" type:"Repeated"`
	AuditResult         *string   `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	Content             *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	EndTime             *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id                  *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	ParentTaskId        *string   `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	Sn                  *int32    `json:"Sn,omitempty" xml:"Sn,omitempty"`
	StartTime           *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Suggestion          *string   `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DescribeAuditContentItemResponseBodyAuditContentItemList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentItemResponseBodyAuditContentItemList) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetAudit(v int32) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.Audit = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetAuditIllegalReasons(v []*string) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.AuditIllegalReasons = v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetAuditResult(v string) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.AuditResult = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetContent(v string) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.Content = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetEndTime(v string) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.EndTime = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetId(v int64) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.Id = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetParentTaskId(v string) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.ParentTaskId = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetSn(v int32) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.Sn = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetStartTime(v string) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.StartTime = &v
	return s
}

func (s *DescribeAuditContentItemResponseBodyAuditContentItemList) SetSuggestion(v string) *DescribeAuditContentItemResponseBodyAuditContentItemList {
	s.Suggestion = &v
	return s
}

type DescribeAuditContentItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditContentItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditContentItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditContentItemResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditContentItemResponse) SetHeaders(v map[string]*string) *DescribeAuditContentItemResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditContentItemResponse) SetStatusCode(v int32) *DescribeAuditContentItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditContentItemResponse) SetBody(v *DescribeAuditContentItemResponseBody) *DescribeAuditContentItemResponse {
	s.Body = v
	return s
}

type DescribeAuditRangeResponseBody struct {
	AuditRange *DescribeAuditRangeResponseBodyAuditRange `json:"AuditRange,omitempty" xml:"AuditRange,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAuditRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditRangeResponseBody) SetAuditRange(v *DescribeAuditRangeResponseBodyAuditRange) *DescribeAuditRangeResponseBody {
	s.AuditRange = v
	return s
}

func (s *DescribeAuditRangeResponseBody) SetRequestId(v string) *DescribeAuditRangeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAuditRangeResponseBodyAuditRange struct {
	Block  *bool `json:"block,omitempty" xml:"block,omitempty"`
	Pass   *bool `json:"pass,omitempty" xml:"pass,omitempty"`
	Review *bool `json:"review,omitempty" xml:"review,omitempty"`
}

func (s DescribeAuditRangeResponseBodyAuditRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRangeResponseBodyAuditRange) GoString() string {
	return s.String()
}

func (s *DescribeAuditRangeResponseBodyAuditRange) SetBlock(v bool) *DescribeAuditRangeResponseBodyAuditRange {
	s.Block = &v
	return s
}

func (s *DescribeAuditRangeResponseBodyAuditRange) SetPass(v bool) *DescribeAuditRangeResponseBodyAuditRange {
	s.Pass = &v
	return s
}

func (s *DescribeAuditRangeResponseBodyAuditRange) SetReview(v bool) *DescribeAuditRangeResponseBodyAuditRange {
	s.Review = &v
	return s
}

type DescribeAuditRangeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditRangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditRangeResponse) SetHeaders(v map[string]*string) *DescribeAuditRangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditRangeResponse) SetStatusCode(v int32) *DescribeAuditRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditRangeResponse) SetBody(v *DescribeAuditRangeResponseBody) *DescribeAuditRangeResponse {
	s.Body = v
	return s
}

type DescribeAuditSettingRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAuditSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditSettingRequest) SetLang(v string) *DescribeAuditSettingRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuditSettingRequest) SetSourceIp(v string) *DescribeAuditSettingRequest {
	s.SourceIp = &v
	return s
}

type DescribeAuditSettingResponseBody struct {
	AuditRange *DescribeAuditSettingResponseBodyAuditRange `json:"AuditRange,omitempty" xml:"AuditRange,omitempty" type:"Struct"`
	Callback   *string                                     `json:"Callback,omitempty" xml:"Callback,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Seed       *string                                     `json:"Seed,omitempty" xml:"Seed,omitempty"`
}

func (s DescribeAuditSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuditSettingResponseBody) SetAuditRange(v *DescribeAuditSettingResponseBodyAuditRange) *DescribeAuditSettingResponseBody {
	s.AuditRange = v
	return s
}

func (s *DescribeAuditSettingResponseBody) SetCallback(v string) *DescribeAuditSettingResponseBody {
	s.Callback = &v
	return s
}

func (s *DescribeAuditSettingResponseBody) SetRequestId(v string) *DescribeAuditSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuditSettingResponseBody) SetSeed(v string) *DescribeAuditSettingResponseBody {
	s.Seed = &v
	return s
}

type DescribeAuditSettingResponseBodyAuditRange struct {
	Block  *bool `json:"block,omitempty" xml:"block,omitempty"`
	Pass   *bool `json:"pass,omitempty" xml:"pass,omitempty"`
	Review *bool `json:"review,omitempty" xml:"review,omitempty"`
}

func (s DescribeAuditSettingResponseBodyAuditRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditSettingResponseBodyAuditRange) GoString() string {
	return s.String()
}

func (s *DescribeAuditSettingResponseBodyAuditRange) SetBlock(v bool) *DescribeAuditSettingResponseBodyAuditRange {
	s.Block = &v
	return s
}

func (s *DescribeAuditSettingResponseBodyAuditRange) SetPass(v bool) *DescribeAuditSettingResponseBodyAuditRange {
	s.Pass = &v
	return s
}

func (s *DescribeAuditSettingResponseBodyAuditRange) SetReview(v bool) *DescribeAuditSettingResponseBodyAuditRange {
	s.Review = &v
	return s
}

type DescribeAuditSettingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuditSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditSettingResponse) SetHeaders(v map[string]*string) *DescribeAuditSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditSettingResponse) SetStatusCode(v int32) *DescribeAuditSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditSettingResponse) SetBody(v *DescribeAuditSettingResponseBody) *DescribeAuditSettingResponse {
	s.Body = v
	return s
}

type DescribeBizTypeImageLibRequest struct {
	BizTypeName  *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s DescribeBizTypeImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibRequest) SetBizTypeName(v string) *DescribeBizTypeImageLibRequest {
	s.BizTypeName = &v
	return s
}

func (s *DescribeBizTypeImageLibRequest) SetResourceType(v string) *DescribeBizTypeImageLibRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeBizTypeImageLibRequest) SetScene(v string) *DescribeBizTypeImageLibRequest {
	s.Scene = &v
	return s
}

type DescribeBizTypeImageLibResponseBody struct {
	Black     *DescribeBizTypeImageLibResponseBodyBlack  `json:"Black,omitempty" xml:"Black,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Review    *DescribeBizTypeImageLibResponseBodyReview `json:"Review,omitempty" xml:"Review,omitempty" type:"Struct"`
	White     *DescribeBizTypeImageLibResponseBodyWhite  `json:"White,omitempty" xml:"White,omitempty" type:"Struct"`
}

func (s DescribeBizTypeImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBody) SetBlack(v *DescribeBizTypeImageLibResponseBodyBlack) *DescribeBizTypeImageLibResponseBody {
	s.Black = v
	return s
}

func (s *DescribeBizTypeImageLibResponseBody) SetRequestId(v string) *DescribeBizTypeImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizTypeImageLibResponseBody) SetReview(v *DescribeBizTypeImageLibResponseBodyReview) *DescribeBizTypeImageLibResponseBody {
	s.Review = v
	return s
}

func (s *DescribeBizTypeImageLibResponseBody) SetWhite(v *DescribeBizTypeImageLibResponseBodyWhite) *DescribeBizTypeImageLibResponseBody {
	s.White = v
	return s
}

type DescribeBizTypeImageLibResponseBodyBlack struct {
	All      []*DescribeBizTypeImageLibResponseBodyBlackAll      `json:"All,omitempty" xml:"All,omitempty" type:"Repeated"`
	Selected []*DescribeBizTypeImageLibResponseBodyBlackSelected `json:"Selected,omitempty" xml:"Selected,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeImageLibResponseBodyBlack) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyBlack) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyBlack) SetAll(v []*DescribeBizTypeImageLibResponseBodyBlackAll) *DescribeBizTypeImageLibResponseBodyBlack {
	s.All = v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyBlack) SetSelected(v []*DescribeBizTypeImageLibResponseBodyBlackSelected) *DescribeBizTypeImageLibResponseBodyBlack {
	s.Selected = v
	return s
}

type DescribeBizTypeImageLibResponseBodyBlackAll struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeImageLibResponseBodyBlackAll) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyBlackAll) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyBlackAll) SetCode(v string) *DescribeBizTypeImageLibResponseBodyBlackAll {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyBlackAll) SetName(v string) *DescribeBizTypeImageLibResponseBodyBlackAll {
	s.Name = &v
	return s
}

type DescribeBizTypeImageLibResponseBodyBlackSelected struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeImageLibResponseBodyBlackSelected) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyBlackSelected) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyBlackSelected) SetCode(v string) *DescribeBizTypeImageLibResponseBodyBlackSelected {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyBlackSelected) SetName(v string) *DescribeBizTypeImageLibResponseBodyBlackSelected {
	s.Name = &v
	return s
}

type DescribeBizTypeImageLibResponseBodyReview struct {
	All      []*DescribeBizTypeImageLibResponseBodyReviewAll      `json:"All,omitempty" xml:"All,omitempty" type:"Repeated"`
	Selected []*DescribeBizTypeImageLibResponseBodyReviewSelected `json:"Selected,omitempty" xml:"Selected,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeImageLibResponseBodyReview) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyReview) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyReview) SetAll(v []*DescribeBizTypeImageLibResponseBodyReviewAll) *DescribeBizTypeImageLibResponseBodyReview {
	s.All = v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyReview) SetSelected(v []*DescribeBizTypeImageLibResponseBodyReviewSelected) *DescribeBizTypeImageLibResponseBodyReview {
	s.Selected = v
	return s
}

type DescribeBizTypeImageLibResponseBodyReviewAll struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeImageLibResponseBodyReviewAll) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyReviewAll) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyReviewAll) SetCode(v string) *DescribeBizTypeImageLibResponseBodyReviewAll {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyReviewAll) SetName(v string) *DescribeBizTypeImageLibResponseBodyReviewAll {
	s.Name = &v
	return s
}

type DescribeBizTypeImageLibResponseBodyReviewSelected struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeImageLibResponseBodyReviewSelected) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyReviewSelected) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyReviewSelected) SetCode(v string) *DescribeBizTypeImageLibResponseBodyReviewSelected {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyReviewSelected) SetName(v string) *DescribeBizTypeImageLibResponseBodyReviewSelected {
	s.Name = &v
	return s
}

type DescribeBizTypeImageLibResponseBodyWhite struct {
	All      []*DescribeBizTypeImageLibResponseBodyWhiteAll      `json:"All,omitempty" xml:"All,omitempty" type:"Repeated"`
	Selected []*DescribeBizTypeImageLibResponseBodyWhiteSelected `json:"Selected,omitempty" xml:"Selected,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeImageLibResponseBodyWhite) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyWhite) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyWhite) SetAll(v []*DescribeBizTypeImageLibResponseBodyWhiteAll) *DescribeBizTypeImageLibResponseBodyWhite {
	s.All = v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyWhite) SetSelected(v []*DescribeBizTypeImageLibResponseBodyWhiteSelected) *DescribeBizTypeImageLibResponseBodyWhite {
	s.Selected = v
	return s
}

type DescribeBizTypeImageLibResponseBodyWhiteAll struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeImageLibResponseBodyWhiteAll) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyWhiteAll) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyWhiteAll) SetCode(v string) *DescribeBizTypeImageLibResponseBodyWhiteAll {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyWhiteAll) SetName(v string) *DescribeBizTypeImageLibResponseBodyWhiteAll {
	s.Name = &v
	return s
}

type DescribeBizTypeImageLibResponseBodyWhiteSelected struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeImageLibResponseBodyWhiteSelected) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponseBodyWhiteSelected) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponseBodyWhiteSelected) SetCode(v string) *DescribeBizTypeImageLibResponseBodyWhiteSelected {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeImageLibResponseBodyWhiteSelected) SetName(v string) *DescribeBizTypeImageLibResponseBodyWhiteSelected {
	s.Name = &v
	return s
}

type DescribeBizTypeImageLibResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBizTypeImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBizTypeImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeImageLibResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeImageLibResponse) SetHeaders(v map[string]*string) *DescribeBizTypeImageLibResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizTypeImageLibResponse) SetStatusCode(v int32) *DescribeBizTypeImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBizTypeImageLibResponse) SetBody(v *DescribeBizTypeImageLibResponseBody) *DescribeBizTypeImageLibResponse {
	s.Body = v
	return s
}

type DescribeBizTypeSettingRequest struct {
	BizTypeName  *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeBizTypeSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingRequest) SetBizTypeName(v string) *DescribeBizTypeSettingRequest {
	s.BizTypeName = &v
	return s
}

func (s *DescribeBizTypeSettingRequest) SetResourceType(v string) *DescribeBizTypeSettingRequest {
	s.ResourceType = &v
	return s
}

type DescribeBizTypeSettingResponseBody struct {
	Ad        *DescribeBizTypeSettingResponseBodyAd        `json:"Ad,omitempty" xml:"Ad,omitempty" type:"Struct"`
	Antispam  *DescribeBizTypeSettingResponseBodyAntispam  `json:"Antispam,omitempty" xml:"Antispam,omitempty" type:"Struct"`
	Live      *DescribeBizTypeSettingResponseBodyLive      `json:"Live,omitempty" xml:"Live,omitempty" type:"Struct"`
	Porn      *DescribeBizTypeSettingResponseBodyPorn      `json:"Porn,omitempty" xml:"Porn,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Terrorism *DescribeBizTypeSettingResponseBodyTerrorism `json:"Terrorism,omitempty" xml:"Terrorism,omitempty" type:"Struct"`
}

func (s DescribeBizTypeSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingResponseBody) SetAd(v *DescribeBizTypeSettingResponseBodyAd) *DescribeBizTypeSettingResponseBody {
	s.Ad = v
	return s
}

func (s *DescribeBizTypeSettingResponseBody) SetAntispam(v *DescribeBizTypeSettingResponseBodyAntispam) *DescribeBizTypeSettingResponseBody {
	s.Antispam = v
	return s
}

func (s *DescribeBizTypeSettingResponseBody) SetLive(v *DescribeBizTypeSettingResponseBodyLive) *DescribeBizTypeSettingResponseBody {
	s.Live = v
	return s
}

func (s *DescribeBizTypeSettingResponseBody) SetPorn(v *DescribeBizTypeSettingResponseBodyPorn) *DescribeBizTypeSettingResponseBody {
	s.Porn = v
	return s
}

func (s *DescribeBizTypeSettingResponseBody) SetRequestId(v string) *DescribeBizTypeSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizTypeSettingResponseBody) SetTerrorism(v *DescribeBizTypeSettingResponseBodyTerrorism) *DescribeBizTypeSettingResponseBody {
	s.Terrorism = v
	return s
}

type DescribeBizTypeSettingResponseBodyAd struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeSettingResponseBodyAd) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingResponseBodyAd) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingResponseBodyAd) SetCategories(v []*string) *DescribeBizTypeSettingResponseBodyAd {
	s.Categories = v
	return s
}

type DescribeBizTypeSettingResponseBodyAntispam struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeSettingResponseBodyAntispam) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingResponseBodyAntispam) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingResponseBodyAntispam) SetCategories(v []*string) *DescribeBizTypeSettingResponseBodyAntispam {
	s.Categories = v
	return s
}

type DescribeBizTypeSettingResponseBodyLive struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeSettingResponseBodyLive) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingResponseBodyLive) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingResponseBodyLive) SetCategories(v []*string) *DescribeBizTypeSettingResponseBodyLive {
	s.Categories = v
	return s
}

type DescribeBizTypeSettingResponseBodyPorn struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeSettingResponseBodyPorn) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingResponseBodyPorn) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingResponseBodyPorn) SetCategories(v []*string) *DescribeBizTypeSettingResponseBodyPorn {
	s.Categories = v
	return s
}

type DescribeBizTypeSettingResponseBodyTerrorism struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeSettingResponseBodyTerrorism) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingResponseBodyTerrorism) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingResponseBodyTerrorism) SetCategories(v []*string) *DescribeBizTypeSettingResponseBodyTerrorism {
	s.Categories = v
	return s
}

type DescribeBizTypeSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBizTypeSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBizTypeSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeSettingResponse) SetHeaders(v map[string]*string) *DescribeBizTypeSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizTypeSettingResponse) SetStatusCode(v int32) *DescribeBizTypeSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBizTypeSettingResponse) SetBody(v *DescribeBizTypeSettingResponseBody) *DescribeBizTypeSettingResponse {
	s.Body = v
	return s
}

type DescribeBizTypeTextLibRequest struct {
	BizTypeName  *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s DescribeBizTypeTextLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibRequest) SetBizTypeName(v string) *DescribeBizTypeTextLibRequest {
	s.BizTypeName = &v
	return s
}

func (s *DescribeBizTypeTextLibRequest) SetResourceType(v string) *DescribeBizTypeTextLibRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeBizTypeTextLibRequest) SetScene(v string) *DescribeBizTypeTextLibRequest {
	s.Scene = &v
	return s
}

type DescribeBizTypeTextLibResponseBody struct {
	Black     *DescribeBizTypeTextLibResponseBodyBlack  `json:"Black,omitempty" xml:"Black,omitempty" type:"Struct"`
	Ignore    *DescribeBizTypeTextLibResponseBodyIgnore `json:"Ignore,omitempty" xml:"Ignore,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Review    *DescribeBizTypeTextLibResponseBodyReview `json:"Review,omitempty" xml:"Review,omitempty" type:"Struct"`
	White     *DescribeBizTypeTextLibResponseBodyWhite  `json:"White,omitempty" xml:"White,omitempty" type:"Struct"`
}

func (s DescribeBizTypeTextLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBody) SetBlack(v *DescribeBizTypeTextLibResponseBodyBlack) *DescribeBizTypeTextLibResponseBody {
	s.Black = v
	return s
}

func (s *DescribeBizTypeTextLibResponseBody) SetIgnore(v *DescribeBizTypeTextLibResponseBodyIgnore) *DescribeBizTypeTextLibResponseBody {
	s.Ignore = v
	return s
}

func (s *DescribeBizTypeTextLibResponseBody) SetRequestId(v string) *DescribeBizTypeTextLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBody) SetReview(v *DescribeBizTypeTextLibResponseBodyReview) *DescribeBizTypeTextLibResponseBody {
	s.Review = v
	return s
}

func (s *DescribeBizTypeTextLibResponseBody) SetWhite(v *DescribeBizTypeTextLibResponseBodyWhite) *DescribeBizTypeTextLibResponseBody {
	s.White = v
	return s
}

type DescribeBizTypeTextLibResponseBodyBlack struct {
	All      []*DescribeBizTypeTextLibResponseBodyBlackAll      `json:"All,omitempty" xml:"All,omitempty" type:"Repeated"`
	Selected []*DescribeBizTypeTextLibResponseBodyBlackSelected `json:"Selected,omitempty" xml:"Selected,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeTextLibResponseBodyBlack) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyBlack) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyBlack) SetAll(v []*DescribeBizTypeTextLibResponseBodyBlackAll) *DescribeBizTypeTextLibResponseBodyBlack {
	s.All = v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyBlack) SetSelected(v []*DescribeBizTypeTextLibResponseBodyBlackSelected) *DescribeBizTypeTextLibResponseBodyBlack {
	s.Selected = v
	return s
}

type DescribeBizTypeTextLibResponseBodyBlackAll struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyBlackAll) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyBlackAll) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyBlackAll) SetCode(v string) *DescribeBizTypeTextLibResponseBodyBlackAll {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyBlackAll) SetName(v string) *DescribeBizTypeTextLibResponseBodyBlackAll {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponseBodyBlackSelected struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyBlackSelected) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyBlackSelected) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyBlackSelected) SetCode(v string) *DescribeBizTypeTextLibResponseBodyBlackSelected {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyBlackSelected) SetName(v string) *DescribeBizTypeTextLibResponseBodyBlackSelected {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponseBodyIgnore struct {
	All      []*DescribeBizTypeTextLibResponseBodyIgnoreAll      `json:"All,omitempty" xml:"All,omitempty" type:"Repeated"`
	Selected []*DescribeBizTypeTextLibResponseBodyIgnoreSelected `json:"Selected,omitempty" xml:"Selected,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeTextLibResponseBodyIgnore) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyIgnore) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyIgnore) SetAll(v []*DescribeBizTypeTextLibResponseBodyIgnoreAll) *DescribeBizTypeTextLibResponseBodyIgnore {
	s.All = v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyIgnore) SetSelected(v []*DescribeBizTypeTextLibResponseBodyIgnoreSelected) *DescribeBizTypeTextLibResponseBodyIgnore {
	s.Selected = v
	return s
}

type DescribeBizTypeTextLibResponseBodyIgnoreAll struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyIgnoreAll) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyIgnoreAll) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyIgnoreAll) SetCode(v string) *DescribeBizTypeTextLibResponseBodyIgnoreAll {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyIgnoreAll) SetName(v string) *DescribeBizTypeTextLibResponseBodyIgnoreAll {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponseBodyIgnoreSelected struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyIgnoreSelected) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyIgnoreSelected) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyIgnoreSelected) SetCode(v string) *DescribeBizTypeTextLibResponseBodyIgnoreSelected {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyIgnoreSelected) SetName(v string) *DescribeBizTypeTextLibResponseBodyIgnoreSelected {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponseBodyReview struct {
	All      []*DescribeBizTypeTextLibResponseBodyReviewAll      `json:"All,omitempty" xml:"All,omitempty" type:"Repeated"`
	Selected []*DescribeBizTypeTextLibResponseBodyReviewSelected `json:"Selected,omitempty" xml:"Selected,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeTextLibResponseBodyReview) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyReview) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyReview) SetAll(v []*DescribeBizTypeTextLibResponseBodyReviewAll) *DescribeBizTypeTextLibResponseBodyReview {
	s.All = v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyReview) SetSelected(v []*DescribeBizTypeTextLibResponseBodyReviewSelected) *DescribeBizTypeTextLibResponseBodyReview {
	s.Selected = v
	return s
}

type DescribeBizTypeTextLibResponseBodyReviewAll struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyReviewAll) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyReviewAll) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyReviewAll) SetCode(v string) *DescribeBizTypeTextLibResponseBodyReviewAll {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyReviewAll) SetName(v string) *DescribeBizTypeTextLibResponseBodyReviewAll {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponseBodyReviewSelected struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyReviewSelected) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyReviewSelected) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyReviewSelected) SetCode(v string) *DescribeBizTypeTextLibResponseBodyReviewSelected {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyReviewSelected) SetName(v string) *DescribeBizTypeTextLibResponseBodyReviewSelected {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponseBodyWhite struct {
	All      []*DescribeBizTypeTextLibResponseBodyWhiteAll      `json:"All,omitempty" xml:"All,omitempty" type:"Repeated"`
	Selected []*DescribeBizTypeTextLibResponseBodyWhiteSelected `json:"Selected,omitempty" xml:"Selected,omitempty" type:"Repeated"`
}

func (s DescribeBizTypeTextLibResponseBodyWhite) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyWhite) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyWhite) SetAll(v []*DescribeBizTypeTextLibResponseBodyWhiteAll) *DescribeBizTypeTextLibResponseBodyWhite {
	s.All = v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyWhite) SetSelected(v []*DescribeBizTypeTextLibResponseBodyWhiteSelected) *DescribeBizTypeTextLibResponseBodyWhite {
	s.Selected = v
	return s
}

type DescribeBizTypeTextLibResponseBodyWhiteAll struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyWhiteAll) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyWhiteAll) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyWhiteAll) SetCode(v string) *DescribeBizTypeTextLibResponseBodyWhiteAll {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyWhiteAll) SetName(v string) *DescribeBizTypeTextLibResponseBodyWhiteAll {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponseBodyWhiteSelected struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeBizTypeTextLibResponseBodyWhiteSelected) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponseBodyWhiteSelected) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponseBodyWhiteSelected) SetCode(v string) *DescribeBizTypeTextLibResponseBodyWhiteSelected {
	s.Code = &v
	return s
}

func (s *DescribeBizTypeTextLibResponseBodyWhiteSelected) SetName(v string) *DescribeBizTypeTextLibResponseBodyWhiteSelected {
	s.Name = &v
	return s
}

type DescribeBizTypeTextLibResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBizTypeTextLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBizTypeTextLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypeTextLibResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizTypeTextLibResponse) SetHeaders(v map[string]*string) *DescribeBizTypeTextLibResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizTypeTextLibResponse) SetStatusCode(v int32) *DescribeBizTypeTextLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBizTypeTextLibResponse) SetBody(v *DescribeBizTypeTextLibResponseBody) *DescribeBizTypeTextLibResponse {
	s.Body = v
	return s
}

type DescribeBizTypesRequest struct {
	ImportFlag *bool `json:"ImportFlag,omitempty" xml:"ImportFlag,omitempty"`
}

func (s DescribeBizTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBizTypesRequest) SetImportFlag(v bool) *DescribeBizTypesRequest {
	s.ImportFlag = &v
	return s
}

type DescribeBizTypesResponseBody struct {
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBizTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBizTypesResponseBody) SetBizTypeList(v []*string) *DescribeBizTypesResponseBody {
	s.BizTypeList = v
	return s
}

func (s *DescribeBizTypesResponseBody) SetRequestId(v string) *DescribeBizTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBizTypesResponseBody) SetTotalCount(v int32) *DescribeBizTypesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBizTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBizTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBizTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBizTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBizTypesResponse) SetHeaders(v map[string]*string) *DescribeBizTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBizTypesResponse) SetStatusCode(v int32) *DescribeBizTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBizTypesResponse) SetBody(v *DescribeBizTypesResponseBody) *DescribeBizTypesResponse {
	s.Body = v
	return s
}

type DescribeCloudMonitorServicesRequest struct {
	Page     *string `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCloudMonitorServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudMonitorServicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorServicesRequest) SetPage(v string) *DescribeCloudMonitorServicesRequest {
	s.Page = &v
	return s
}

func (s *DescribeCloudMonitorServicesRequest) SetPageSize(v string) *DescribeCloudMonitorServicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudMonitorServicesRequest) SetSourceIp(v string) *DescribeCloudMonitorServicesRequest {
	s.SourceIp = &v
	return s
}

type DescribeCloudMonitorServicesResponseBody struct {
	Items      []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudMonitorServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudMonitorServicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorServicesResponseBody) SetItems(v []*string) *DescribeCloudMonitorServicesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCloudMonitorServicesResponseBody) SetRequestId(v string) *DescribeCloudMonitorServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudMonitorServicesResponseBody) SetTotalCount(v int32) *DescribeCloudMonitorServicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCloudMonitorServicesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudMonitorServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudMonitorServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudMonitorServicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorServicesResponse) SetHeaders(v map[string]*string) *DescribeCloudMonitorServicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudMonitorServicesResponse) SetStatusCode(v int32) *DescribeCloudMonitorServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudMonitorServicesResponse) SetBody(v *DescribeCloudMonitorServicesResponseBody) *DescribeCloudMonitorServicesResponse {
	s.Body = v
	return s
}

type DescribeCustomOcrTemplateRequest struct {
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DescribeCustomOcrTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomOcrTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomOcrTemplateRequest) SetIds(v string) *DescribeCustomOcrTemplateRequest {
	s.Ids = &v
	return s
}

type DescribeCustomOcrTemplateResponseBody struct {
	OcrTemplateList []*DescribeCustomOcrTemplateResponseBodyOcrTemplateList `json:"OcrTemplateList,omitempty" xml:"OcrTemplateList,omitempty" type:"Repeated"`
	RequestId       *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCustomOcrTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomOcrTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomOcrTemplateResponseBody) SetOcrTemplateList(v []*DescribeCustomOcrTemplateResponseBodyOcrTemplateList) *DescribeCustomOcrTemplateResponseBody {
	s.OcrTemplateList = v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBody) SetRequestId(v string) *DescribeCustomOcrTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBody) SetTotalCount(v int32) *DescribeCustomOcrTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCustomOcrTemplateResponseBodyOcrTemplateList struct {
	Id            *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	ImgUrl        *string                                                              `json:"ImgUrl,omitempty" xml:"ImgUrl,omitempty"`
	Name          *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RecognizeArea []*DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea `json:"RecognizeArea,omitempty" xml:"RecognizeArea,omitempty" type:"Repeated"`
	ReferArea     []*DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea     `json:"ReferArea,omitempty" xml:"ReferArea,omitempty" type:"Repeated"`
	Status        *int32                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCustomOcrTemplateResponseBodyOcrTemplateList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomOcrTemplateResponseBodyOcrTemplateList) GoString() string {
	return s.String()
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateList) SetId(v int64) *DescribeCustomOcrTemplateResponseBodyOcrTemplateList {
	s.Id = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateList) SetImgUrl(v string) *DescribeCustomOcrTemplateResponseBodyOcrTemplateList {
	s.ImgUrl = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateList) SetName(v string) *DescribeCustomOcrTemplateResponseBodyOcrTemplateList {
	s.Name = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateList) SetRecognizeArea(v []*DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) *DescribeCustomOcrTemplateResponseBodyOcrTemplateList {
	s.RecognizeArea = v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateList) SetReferArea(v []*DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) *DescribeCustomOcrTemplateResponseBodyOcrTemplateList {
	s.ReferArea = v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateList) SetStatus(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateList {
	s.Status = &v
	return s
}

type DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea struct {
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32  `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32  `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) GoString() string {
	return s.String()
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) SetHeight(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea {
	s.Height = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) SetName(v string) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea {
	s.Name = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) SetWidth(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea {
	s.Width = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) SetX(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea {
	s.X = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea) SetY(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListRecognizeArea {
	s.Y = &v
	return s
}

type DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea struct {
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32  `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32  `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) GoString() string {
	return s.String()
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) SetHeight(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea {
	s.Height = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) SetName(v string) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea {
	s.Name = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) SetWidth(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea {
	s.Width = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) SetX(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea {
	s.X = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea) SetY(v int32) *DescribeCustomOcrTemplateResponseBodyOcrTemplateListReferArea {
	s.Y = &v
	return s
}

type DescribeCustomOcrTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomOcrTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomOcrTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomOcrTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomOcrTemplateResponse) SetHeaders(v map[string]*string) *DescribeCustomOcrTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomOcrTemplateResponse) SetStatusCode(v int32) *DescribeCustomOcrTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomOcrTemplateResponse) SetBody(v *DescribeCustomOcrTemplateResponseBody) *DescribeCustomOcrTemplateResponse {
	s.Body = v
	return s
}

type DescribeImageFromLibRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageLibId  *int32  `json:"ImageLibId,omitempty" xml:"ImageLibId,omitempty"`
	ModelId     *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageFromLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageFromLibRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageFromLibRequest) SetCurrentPage(v int32) *DescribeImageFromLibRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetEndDate(v string) *DescribeImageFromLibRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetId(v int64) *DescribeImageFromLibRequest {
	s.Id = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetImageLibId(v int32) *DescribeImageFromLibRequest {
	s.ImageLibId = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetModelId(v int64) *DescribeImageFromLibRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetPageSize(v int32) *DescribeImageFromLibRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetSourceIp(v string) *DescribeImageFromLibRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetStartDate(v string) *DescribeImageFromLibRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeImageFromLibRequest) SetTotalCount(v int32) *DescribeImageFromLibRequest {
	s.TotalCount = &v
	return s
}

type DescribeImageFromLibResponseBody struct {
	CurrentPage      *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ImageFromLibList []*DescribeImageFromLibResponseBodyImageFromLibList `json:"ImageFromLibList,omitempty" xml:"ImageFromLibList,omitempty" type:"Repeated"`
	PageSize         *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageFromLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageFromLibResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageFromLibResponseBody) SetCurrentPage(v int32) *DescribeImageFromLibResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageFromLibResponseBody) SetImageFromLibList(v []*DescribeImageFromLibResponseBodyImageFromLibList) *DescribeImageFromLibResponseBody {
	s.ImageFromLibList = v
	return s
}

func (s *DescribeImageFromLibResponseBody) SetPageSize(v int32) *DescribeImageFromLibResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageFromLibResponseBody) SetRequestId(v string) *DescribeImageFromLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageFromLibResponseBody) SetTotalCount(v int32) *DescribeImageFromLibResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeImageFromLibResponseBodyImageFromLibList struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Image         *string `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageHitCount *int64  `json:"ImageHitCount,omitempty" xml:"ImageHitCount,omitempty"`
	Thumbnail     *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	VideoHitCount *int64  `json:"VideoHitCount,omitempty" xml:"VideoHitCount,omitempty"`
}

func (s DescribeImageFromLibResponseBodyImageFromLibList) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageFromLibResponseBodyImageFromLibList) GoString() string {
	return s.String()
}

func (s *DescribeImageFromLibResponseBodyImageFromLibList) SetCreateTime(v string) *DescribeImageFromLibResponseBodyImageFromLibList {
	s.CreateTime = &v
	return s
}

func (s *DescribeImageFromLibResponseBodyImageFromLibList) SetId(v int64) *DescribeImageFromLibResponseBodyImageFromLibList {
	s.Id = &v
	return s
}

func (s *DescribeImageFromLibResponseBodyImageFromLibList) SetImage(v string) *DescribeImageFromLibResponseBodyImageFromLibList {
	s.Image = &v
	return s
}

func (s *DescribeImageFromLibResponseBodyImageFromLibList) SetImageHitCount(v int64) *DescribeImageFromLibResponseBodyImageFromLibList {
	s.ImageHitCount = &v
	return s
}

func (s *DescribeImageFromLibResponseBodyImageFromLibList) SetThumbnail(v string) *DescribeImageFromLibResponseBodyImageFromLibList {
	s.Thumbnail = &v
	return s
}

func (s *DescribeImageFromLibResponseBodyImageFromLibList) SetVideoHitCount(v int64) *DescribeImageFromLibResponseBodyImageFromLibList {
	s.VideoHitCount = &v
	return s
}

type DescribeImageFromLibResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageFromLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageFromLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageFromLibResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageFromLibResponse) SetHeaders(v map[string]*string) *DescribeImageFromLibResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageFromLibResponse) SetStatusCode(v int32) *DescribeImageFromLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageFromLibResponse) SetBody(v *DescribeImageFromLibResponseBody) *DescribeImageFromLibResponse {
	s.Body = v
	return s
}

type DescribeImageLibRequest struct {
	ServiceModule *string `json:"ServiceModule,omitempty" xml:"ServiceModule,omitempty"`
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageLibRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageLibRequest) SetServiceModule(v string) *DescribeImageLibRequest {
	s.ServiceModule = &v
	return s
}

func (s *DescribeImageLibRequest) SetSourceIp(v string) *DescribeImageLibRequest {
	s.SourceIp = &v
	return s
}

type DescribeImageLibResponseBody struct {
	ImageLibList []*DescribeImageLibResponseBodyImageLibList `json:"ImageLibList,omitempty" xml:"ImageLibList,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageLibResponseBody) SetImageLibList(v []*DescribeImageLibResponseBodyImageLibList) *DescribeImageLibResponseBody {
	s.ImageLibList = v
	return s
}

func (s *DescribeImageLibResponseBody) SetRequestId(v string) *DescribeImageLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageLibResponseBody) SetTotalCount(v int32) *DescribeImageLibResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeImageLibResponseBodyImageLibList struct {
	BizTypes      []*string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty" type:"Repeated"`
	Category      *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	Code          *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Enable        *string   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Id            *int32    `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageCount    *int32    `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	ModifiedTime  *string   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Scene         *string   `json:"Scene,omitempty" xml:"Scene,omitempty"`
	ServiceModule *string   `json:"ServiceModule,omitempty" xml:"ServiceModule,omitempty"`
	Source        *string   `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeImageLibResponseBodyImageLibList) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageLibResponseBodyImageLibList) GoString() string {
	return s.String()
}

func (s *DescribeImageLibResponseBodyImageLibList) SetBizTypes(v []*string) *DescribeImageLibResponseBodyImageLibList {
	s.BizTypes = v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetCategory(v string) *DescribeImageLibResponseBodyImageLibList {
	s.Category = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetCode(v string) *DescribeImageLibResponseBodyImageLibList {
	s.Code = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetEnable(v string) *DescribeImageLibResponseBodyImageLibList {
	s.Enable = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetId(v int32) *DescribeImageLibResponseBodyImageLibList {
	s.Id = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetImageCount(v int32) *DescribeImageLibResponseBodyImageLibList {
	s.ImageCount = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetModifiedTime(v string) *DescribeImageLibResponseBodyImageLibList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetName(v string) *DescribeImageLibResponseBodyImageLibList {
	s.Name = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetScene(v string) *DescribeImageLibResponseBodyImageLibList {
	s.Scene = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetServiceModule(v string) *DescribeImageLibResponseBodyImageLibList {
	s.ServiceModule = &v
	return s
}

func (s *DescribeImageLibResponseBodyImageLibList) SetSource(v string) *DescribeImageLibResponseBodyImageLibList {
	s.Source = &v
	return s
}

type DescribeImageLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageLibResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageLibResponse) SetHeaders(v map[string]*string) *DescribeImageLibResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageLibResponse) SetStatusCode(v int32) *DescribeImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageLibResponse) SetBody(v *DescribeImageLibResponseBody) *DescribeImageLibResponse {
	s.Body = v
	return s
}

type DescribeImageUploadInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeImageUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageUploadInfoRequest) SetSourceIp(v string) *DescribeImageUploadInfoRequest {
	s.SourceIp = &v
	return s
}

type DescribeImageUploadInfoResponseBody struct {
	Accessid  *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	Expire    *int32  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Folder    *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeImageUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageUploadInfoResponseBody) SetAccessid(v string) *DescribeImageUploadInfoResponseBody {
	s.Accessid = &v
	return s
}

func (s *DescribeImageUploadInfoResponseBody) SetExpire(v int32) *DescribeImageUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeImageUploadInfoResponseBody) SetFolder(v string) *DescribeImageUploadInfoResponseBody {
	s.Folder = &v
	return s
}

func (s *DescribeImageUploadInfoResponseBody) SetHost(v string) *DescribeImageUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeImageUploadInfoResponseBody) SetPolicy(v string) *DescribeImageUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeImageUploadInfoResponseBody) SetRequestId(v string) *DescribeImageUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageUploadInfoResponseBody) SetSignature(v string) *DescribeImageUploadInfoResponseBody {
	s.Signature = &v
	return s
}

type DescribeImageUploadInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageUploadInfoResponse) SetHeaders(v map[string]*string) *DescribeImageUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageUploadInfoResponse) SetStatusCode(v int32) *DescribeImageUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageUploadInfoResponse) SetBody(v *DescribeImageUploadInfoResponseBody) *DescribeImageUploadInfoResponse {
	s.Body = v
	return s
}

type DescribeKeywordRequest struct {
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	KeywordLibId *int32  `json:"KeywordLibId,omitempty" xml:"KeywordLibId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TotalCount   *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeywordRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeywordRequest) SetCurrentPage(v int32) *DescribeKeywordRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeKeywordRequest) SetKeyword(v string) *DescribeKeywordRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeKeywordRequest) SetKeywordLibId(v int32) *DescribeKeywordRequest {
	s.KeywordLibId = &v
	return s
}

func (s *DescribeKeywordRequest) SetLang(v string) *DescribeKeywordRequest {
	s.Lang = &v
	return s
}

func (s *DescribeKeywordRequest) SetPageSize(v int32) *DescribeKeywordRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKeywordRequest) SetSourceIp(v string) *DescribeKeywordRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeKeywordRequest) SetTotalCount(v int32) *DescribeKeywordRequest {
	s.TotalCount = &v
	return s
}

type DescribeKeywordResponseBody struct {
	CurrentPage *int32                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	KeywordList []*DescribeKeywordResponseBodyKeywordList `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeywordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeywordResponseBody) SetCurrentPage(v int32) *DescribeKeywordResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeKeywordResponseBody) SetKeywordList(v []*DescribeKeywordResponseBodyKeywordList) *DescribeKeywordResponseBody {
	s.KeywordList = v
	return s
}

func (s *DescribeKeywordResponseBody) SetPageSize(v int32) *DescribeKeywordResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeKeywordResponseBody) SetRequestId(v string) *DescribeKeywordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeywordResponseBody) SetTotalCount(v int32) *DescribeKeywordResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeKeywordResponseBodyKeywordList struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	HitCount   *int32  `json:"HitCount,omitempty" xml:"HitCount,omitempty"`
	Id         *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s DescribeKeywordResponseBodyKeywordList) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordResponseBodyKeywordList) GoString() string {
	return s.String()
}

func (s *DescribeKeywordResponseBodyKeywordList) SetCreateTime(v string) *DescribeKeywordResponseBodyKeywordList {
	s.CreateTime = &v
	return s
}

func (s *DescribeKeywordResponseBodyKeywordList) SetHitCount(v int32) *DescribeKeywordResponseBodyKeywordList {
	s.HitCount = &v
	return s
}

func (s *DescribeKeywordResponseBodyKeywordList) SetId(v int32) *DescribeKeywordResponseBodyKeywordList {
	s.Id = &v
	return s
}

func (s *DescribeKeywordResponseBodyKeywordList) SetKeyword(v string) *DescribeKeywordResponseBodyKeywordList {
	s.Keyword = &v
	return s
}

type DescribeKeywordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeywordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeywordResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeywordResponse) SetHeaders(v map[string]*string) *DescribeKeywordResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeywordResponse) SetStatusCode(v int32) *DescribeKeywordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeywordResponse) SetBody(v *DescribeKeywordResponseBody) *DescribeKeywordResponse {
	s.Body = v
	return s
}

type DescribeKeywordLibRequest struct {
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ServiceModule *string `json:"ServiceModule,omitempty" xml:"ServiceModule,omitempty"`
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeywordLibRequest) SetLang(v string) *DescribeKeywordLibRequest {
	s.Lang = &v
	return s
}

func (s *DescribeKeywordLibRequest) SetServiceModule(v string) *DescribeKeywordLibRequest {
	s.ServiceModule = &v
	return s
}

func (s *DescribeKeywordLibRequest) SetSourceIp(v string) *DescribeKeywordLibRequest {
	s.SourceIp = &v
	return s
}

type DescribeKeywordLibResponseBody struct {
	KeywordLibList []*DescribeKeywordLibResponseBodyKeywordLibList `json:"KeywordLibList,omitempty" xml:"KeywordLibList,omitempty" type:"Repeated"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKeywordLibResponseBody) SetKeywordLibList(v []*DescribeKeywordLibResponseBodyKeywordLibList) *DescribeKeywordLibResponseBody {
	s.KeywordLibList = v
	return s
}

func (s *DescribeKeywordLibResponseBody) SetRequestId(v string) *DescribeKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKeywordLibResponseBody) SetTotalCount(v int32) *DescribeKeywordLibResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeKeywordLibResponseBodyKeywordLibList struct {
	BizTypes      []*string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty" type:"Repeated"`
	Category      *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	Code          *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Count         *int32    `json:"Count,omitempty" xml:"Count,omitempty"`
	Enable        *bool     `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Id            *int32    `json:"Id,omitempty" xml:"Id,omitempty"`
	Language      *string   `json:"Language,omitempty" xml:"Language,omitempty"`
	LibType       *string   `json:"LibType,omitempty" xml:"LibType,omitempty"`
	MatchMode     *string   `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	ModifiedTime  *string   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceType  *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceModule *string   `json:"ServiceModule,omitempty" xml:"ServiceModule,omitempty"`
	Source        *string   `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeKeywordLibResponseBodyKeywordLibList) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordLibResponseBodyKeywordLibList) GoString() string {
	return s.String()
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetBizTypes(v []*string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.BizTypes = v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetCategory(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Category = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetCode(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Code = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetCount(v int32) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Count = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetEnable(v bool) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Enable = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetId(v int32) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Id = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetLanguage(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Language = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetLibType(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.LibType = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetMatchMode(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.MatchMode = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetModifiedTime(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetName(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Name = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetResourceType(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.ResourceType = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetServiceModule(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.ServiceModule = &v
	return s
}

func (s *DescribeKeywordLibResponseBodyKeywordLibList) SetSource(v string) *DescribeKeywordLibResponseBodyKeywordLibList {
	s.Source = &v
	return s
}

type DescribeKeywordLibResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeywordLibResponse) SetHeaders(v map[string]*string) *DescribeKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeywordLibResponse) SetStatusCode(v int32) *DescribeKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeywordLibResponse) SetBody(v *DescribeKeywordLibResponseBody) *DescribeKeywordLibResponse {
	s.Body = v
	return s
}

type DescribeNotificationSettingRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeNotificationSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotificationSettingRequest) SetLang(v string) *DescribeNotificationSettingRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNotificationSettingRequest) SetSourceIp(v string) *DescribeNotificationSettingRequest {
	s.SourceIp = &v
	return s
}

type DescribeNotificationSettingResponseBody struct {
	Email                   *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone                   *string   `json:"Phone,omitempty" xml:"Phone,omitempty"`
	RealtimeMessageList     []*string `json:"RealtimeMessageList,omitempty" xml:"RealtimeMessageList,omitempty" type:"Repeated"`
	ReminderModeList        []*string `json:"ReminderModeList,omitempty" xml:"ReminderModeList,omitempty" type:"Repeated"`
	RequestId               *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduleMessageTime     *int32    `json:"ScheduleMessageTime,omitempty" xml:"ScheduleMessageTime,omitempty"`
	ScheduleMessageTimeZone *int32    `json:"ScheduleMessageTimeZone,omitempty" xml:"ScheduleMessageTimeZone,omitempty"`
}

func (s DescribeNotificationSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotificationSettingResponseBody) SetEmail(v string) *DescribeNotificationSettingResponseBody {
	s.Email = &v
	return s
}

func (s *DescribeNotificationSettingResponseBody) SetPhone(v string) *DescribeNotificationSettingResponseBody {
	s.Phone = &v
	return s
}

func (s *DescribeNotificationSettingResponseBody) SetRealtimeMessageList(v []*string) *DescribeNotificationSettingResponseBody {
	s.RealtimeMessageList = v
	return s
}

func (s *DescribeNotificationSettingResponseBody) SetReminderModeList(v []*string) *DescribeNotificationSettingResponseBody {
	s.ReminderModeList = v
	return s
}

func (s *DescribeNotificationSettingResponseBody) SetRequestId(v string) *DescribeNotificationSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNotificationSettingResponseBody) SetScheduleMessageTime(v int32) *DescribeNotificationSettingResponseBody {
	s.ScheduleMessageTime = &v
	return s
}

func (s *DescribeNotificationSettingResponseBody) SetScheduleMessageTimeZone(v int32) *DescribeNotificationSettingResponseBody {
	s.ScheduleMessageTimeZone = &v
	return s
}

type DescribeNotificationSettingResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNotificationSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNotificationSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNotificationSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotificationSettingResponse) SetHeaders(v map[string]*string) *DescribeNotificationSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotificationSettingResponse) SetStatusCode(v int32) *DescribeNotificationSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotificationSettingResponse) SetBody(v *DescribeNotificationSettingResponseBody) *DescribeNotificationSettingResponse {
	s.Body = v
	return s
}

type DescribeOpenApiRcpStatsRequest struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeOpenApiRcpStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiRcpStatsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiRcpStatsRequest) SetBizType(v string) *DescribeOpenApiRcpStatsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeOpenApiRcpStatsRequest) SetEndDate(v string) *DescribeOpenApiRcpStatsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeOpenApiRcpStatsRequest) SetResourceType(v string) *DescribeOpenApiRcpStatsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeOpenApiRcpStatsRequest) SetStartDate(v string) *DescribeOpenApiRcpStatsRequest {
	s.StartDate = &v
	return s
}

type DescribeOpenApiRcpStatsResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList   []*DescribeOpenApiRcpStatsResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpenApiRcpStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiRcpStatsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiRcpStatsResponseBody) SetRequestId(v string) *DescribeOpenApiRcpStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBody) SetStatList(v []*DescribeOpenApiRcpStatsResponseBodyStatList) *DescribeOpenApiRcpStatsResponseBody {
	s.StatList = v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBody) SetTotalCount(v int32) *DescribeOpenApiRcpStatsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOpenApiRcpStatsResponseBodyStatList struct {
	BlockCount    *int32  `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	Date          *string `json:"Date,omitempty" xml:"Date,omitempty"`
	PassCount     *int32  `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ReviewCount   *int32  `json:"ReviewCount,omitempty" xml:"ReviewCount,omitempty"`
	TotalCount    *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalDuration *int32  `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
}

func (s DescribeOpenApiRcpStatsResponseBodyStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiRcpStatsResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiRcpStatsResponseBodyStatList) SetBlockCount(v int32) *DescribeOpenApiRcpStatsResponseBodyStatList {
	s.BlockCount = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBodyStatList) SetDate(v string) *DescribeOpenApiRcpStatsResponseBodyStatList {
	s.Date = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBodyStatList) SetPassCount(v int32) *DescribeOpenApiRcpStatsResponseBodyStatList {
	s.PassCount = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBodyStatList) SetResourceType(v string) *DescribeOpenApiRcpStatsResponseBodyStatList {
	s.ResourceType = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBodyStatList) SetReviewCount(v int32) *DescribeOpenApiRcpStatsResponseBodyStatList {
	s.ReviewCount = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBodyStatList) SetTotalCount(v int32) *DescribeOpenApiRcpStatsResponseBodyStatList {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponseBodyStatList) SetTotalDuration(v int32) *DescribeOpenApiRcpStatsResponseBodyStatList {
	s.TotalDuration = &v
	return s
}

type DescribeOpenApiRcpStatsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenApiRcpStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenApiRcpStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiRcpStatsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiRcpStatsResponse) SetHeaders(v map[string]*string) *DescribeOpenApiRcpStatsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenApiRcpStatsResponse) SetStatusCode(v int32) *DescribeOpenApiRcpStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenApiRcpStatsResponse) SetBody(v *DescribeOpenApiRcpStatsResponseBody) *DescribeOpenApiRcpStatsResponse {
	s.Body = v
	return s
}

type DescribeOpenApiUsageRequest struct {
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeOpenApiUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiUsageRequest) SetEndDate(v string) *DescribeOpenApiUsageRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeOpenApiUsageRequest) SetResourceType(v string) *DescribeOpenApiUsageRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeOpenApiUsageRequest) SetSourceIp(v string) *DescribeOpenApiUsageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpenApiUsageRequest) SetStartDate(v string) *DescribeOpenApiUsageRequest {
	s.StartDate = &v
	return s
}

type DescribeOpenApiUsageResponseBody struct {
	OpenApiUsageList []*DescribeOpenApiUsageResponseBodyOpenApiUsageList `json:"OpenApiUsageList,omitempty" xml:"OpenApiUsageList,omitempty" type:"Repeated"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpenApiUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiUsageResponseBody) SetOpenApiUsageList(v []*DescribeOpenApiUsageResponseBodyOpenApiUsageList) *DescribeOpenApiUsageResponseBody {
	s.OpenApiUsageList = v
	return s
}

func (s *DescribeOpenApiUsageResponseBody) SetRequestId(v string) *DescribeOpenApiUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBody) SetTotalCount(v int32) *DescribeOpenApiUsageResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOpenApiUsageResponseBodyOpenApiUsageList struct {
	BlockCount      *int32  `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	BlockDuration   *int32  `json:"BlockDuration,omitempty" xml:"BlockDuration,omitempty"`
	Date            *string `json:"Date,omitempty" xml:"Date,omitempty"`
	InnerTotalCount *int32  `json:"InnerTotalCount,omitempty" xml:"InnerTotalCount,omitempty"`
	OuterTotalCount *int32  `json:"OuterTotalCount,omitempty" xml:"OuterTotalCount,omitempty"`
	PassCount       *int32  `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	PassDuration    *int32  `json:"PassDuration,omitempty" xml:"PassDuration,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ReviewCount     *int32  `json:"ReviewCount,omitempty" xml:"ReviewCount,omitempty"`
	ReviewDuration  *int32  `json:"ReviewDuration,omitempty" xml:"ReviewDuration,omitempty"`
	Scene           *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TotalCount      *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalDuration   *int32  `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
}

func (s DescribeOpenApiUsageResponseBodyOpenApiUsageList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiUsageResponseBodyOpenApiUsageList) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetBlockCount(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.BlockCount = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetBlockDuration(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.BlockDuration = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetDate(v string) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.Date = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetInnerTotalCount(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.InnerTotalCount = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetOuterTotalCount(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.OuterTotalCount = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetPassCount(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.PassCount = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetPassDuration(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.PassDuration = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetResourceType(v string) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.ResourceType = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetReviewCount(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.ReviewCount = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetReviewDuration(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.ReviewDuration = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetScene(v string) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.Scene = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetTotalCount(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpenApiUsageResponseBodyOpenApiUsageList) SetTotalDuration(v int32) *DescribeOpenApiUsageResponseBodyOpenApiUsageList {
	s.TotalDuration = &v
	return s
}

type DescribeOpenApiUsageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenApiUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenApiUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpenApiUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiUsageResponse) SetHeaders(v map[string]*string) *DescribeOpenApiUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenApiUsageResponse) SetStatusCode(v int32) *DescribeOpenApiUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenApiUsageResponse) SetBody(v *DescribeOpenApiUsageResponseBody) *DescribeOpenApiUsageResponse {
	s.Body = v
	return s
}

type DescribeOssCallbackSettingResponseBody struct {
	AuditCallback           *bool     `json:"AuditCallback,omitempty" xml:"AuditCallback,omitempty"`
	CallbackSeed            *string   `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	CallbackUrl             *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	RequestId               *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanCallback            *bool     `json:"ScanCallback,omitempty" xml:"ScanCallback,omitempty"`
	ScanCallbackSuggestions []*string `json:"ScanCallbackSuggestions,omitempty" xml:"ScanCallbackSuggestions,omitempty" type:"Repeated"`
	ServiceModules          []*string `json:"ServiceModules,omitempty" xml:"ServiceModules,omitempty" type:"Repeated"`
}

func (s DescribeOssCallbackSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssCallbackSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssCallbackSettingResponseBody) SetAuditCallback(v bool) *DescribeOssCallbackSettingResponseBody {
	s.AuditCallback = &v
	return s
}

func (s *DescribeOssCallbackSettingResponseBody) SetCallbackSeed(v string) *DescribeOssCallbackSettingResponseBody {
	s.CallbackSeed = &v
	return s
}

func (s *DescribeOssCallbackSettingResponseBody) SetCallbackUrl(v string) *DescribeOssCallbackSettingResponseBody {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeOssCallbackSettingResponseBody) SetRequestId(v string) *DescribeOssCallbackSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssCallbackSettingResponseBody) SetScanCallback(v bool) *DescribeOssCallbackSettingResponseBody {
	s.ScanCallback = &v
	return s
}

func (s *DescribeOssCallbackSettingResponseBody) SetScanCallbackSuggestions(v []*string) *DescribeOssCallbackSettingResponseBody {
	s.ScanCallbackSuggestions = v
	return s
}

func (s *DescribeOssCallbackSettingResponseBody) SetServiceModules(v []*string) *DescribeOssCallbackSettingResponseBody {
	s.ServiceModules = v
	return s
}

type DescribeOssCallbackSettingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssCallbackSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssCallbackSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssCallbackSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssCallbackSettingResponse) SetHeaders(v map[string]*string) *DescribeOssCallbackSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssCallbackSettingResponse) SetStatusCode(v int32) *DescribeOssCallbackSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssCallbackSettingResponse) SetBody(v *DescribeOssCallbackSettingResponseBody) *DescribeOssCallbackSettingResponse {
	s.Body = v
	return s
}

type DescribeOssIncrementCheckSettingRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeOssIncrementCheckSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingRequest) SetLang(v string) *DescribeOssIncrementCheckSettingRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingRequest) SetSourceIp(v string) *DescribeOssIncrementCheckSettingRequest {
	s.SourceIp = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBody struct {
	AudioAntispamFreezeConfig      *DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig      `json:"AudioAntispamFreezeConfig,omitempty" xml:"AudioAntispamFreezeConfig,omitempty" type:"Struct"`
	AudioAutoFreezeOpened          *bool                                                                       `json:"AudioAutoFreezeOpened,omitempty" xml:"AudioAutoFreezeOpened,omitempty"`
	AudioMaxSize                   *int32                                                                      `json:"AudioMaxSize,omitempty" xml:"AudioMaxSize,omitempty"`
	AudioScanLimit                 *int64                                                                      `json:"AudioScanLimit,omitempty" xml:"AudioScanLimit,omitempty"`
	AudioSceneList                 []*string                                                                   `json:"AudioSceneList,omitempty" xml:"AudioSceneList,omitempty" type:"Repeated"`
	AutoFreezeType                 *string                                                                     `json:"AutoFreezeType,omitempty" xml:"AutoFreezeType,omitempty"`
	BizType                        *string                                                                     `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizTypeTemplate                *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate                `json:"BizTypeTemplate,omitempty" xml:"BizTypeTemplate,omitempty" type:"Struct"`
	BucketConfigList               []*DescribeOssIncrementCheckSettingResponseBodyBucketConfigList             `json:"BucketConfigList,omitempty" xml:"BucketConfigList,omitempty" type:"Repeated"`
	CallbackId                     *string                                                                     `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	CallbackName                   *string                                                                     `json:"CallbackName,omitempty" xml:"CallbackName,omitempty"`
	EndTime                        *string                                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ImageAdFreezeConfig            *DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig            `json:"ImageAdFreezeConfig,omitempty" xml:"ImageAdFreezeConfig,omitempty" type:"Struct"`
	ImageAutoFreeze                *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze                `json:"ImageAutoFreeze,omitempty" xml:"ImageAutoFreeze,omitempty" type:"Struct"`
	ImageAutoFreezeOpened          *bool                                                                       `json:"ImageAutoFreezeOpened,omitempty" xml:"ImageAutoFreezeOpened,omitempty"`
	ImageEnableLimit               *bool                                                                       `json:"ImageEnableLimit,omitempty" xml:"ImageEnableLimit,omitempty"`
	ImageLiveFreezeConfig          *DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig          `json:"ImageLiveFreezeConfig,omitempty" xml:"ImageLiveFreezeConfig,omitempty" type:"Struct"`
	ImagePornFreezeConfig          *DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig          `json:"ImagePornFreezeConfig,omitempty" xml:"ImagePornFreezeConfig,omitempty" type:"Struct"`
	ImageScanLimit                 *int64                                                                      `json:"ImageScanLimit,omitempty" xml:"ImageScanLimit,omitempty"`
	ImageSceneList                 []*string                                                                   `json:"ImageSceneList,omitempty" xml:"ImageSceneList,omitempty" type:"Repeated"`
	ImageTerrorismFreezeConfig     *DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig     `json:"ImageTerrorismFreezeConfig,omitempty" xml:"ImageTerrorismFreezeConfig,omitempty" type:"Struct"`
	RequestId                      *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanImageNoFileType            *bool                                                                       `json:"ScanImageNoFileType,omitempty" xml:"ScanImageNoFileType,omitempty"`
	StartTime                      *string                                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VideoAdFreezeConfig            *DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig            `json:"VideoAdFreezeConfig,omitempty" xml:"VideoAdFreezeConfig,omitempty" type:"Struct"`
	VideoAutoFreezeOpened          *bool                                                                       `json:"VideoAutoFreezeOpened,omitempty" xml:"VideoAutoFreezeOpened,omitempty"`
	VideoAutoFreezeSceneList       []*string                                                                   `json:"VideoAutoFreezeSceneList,omitempty" xml:"VideoAutoFreezeSceneList,omitempty" type:"Repeated"`
	VideoFrameInterval             *int32                                                                      `json:"VideoFrameInterval,omitempty" xml:"VideoFrameInterval,omitempty"`
	VideoLiveFreezeConfig          *DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig          `json:"VideoLiveFreezeConfig,omitempty" xml:"VideoLiveFreezeConfig,omitempty" type:"Struct"`
	VideoMaxFrames                 *int32                                                                      `json:"VideoMaxFrames,omitempty" xml:"VideoMaxFrames,omitempty"`
	VideoMaxSize                   *int32                                                                      `json:"VideoMaxSize,omitempty" xml:"VideoMaxSize,omitempty"`
	VideoPornFreezeConfig          *DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig          `json:"VideoPornFreezeConfig,omitempty" xml:"VideoPornFreezeConfig,omitempty" type:"Struct"`
	VideoScanLimit                 *int64                                                                      `json:"VideoScanLimit,omitempty" xml:"VideoScanLimit,omitempty"`
	VideoSceneList                 []*string                                                                   `json:"VideoSceneList,omitempty" xml:"VideoSceneList,omitempty" type:"Repeated"`
	VideoTerrorismFreezeConfig     *DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig     `json:"VideoTerrorismFreezeConfig,omitempty" xml:"VideoTerrorismFreezeConfig,omitempty" type:"Struct"`
	VideoVoiceAntispamFreezeConfig *DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig `json:"VideoVoiceAntispamFreezeConfig,omitempty" xml:"VideoVoiceAntispamFreezeConfig,omitempty" type:"Struct"`
}

func (s DescribeOssIncrementCheckSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetAudioAntispamFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.AudioAntispamFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetAudioAutoFreezeOpened(v bool) *DescribeOssIncrementCheckSettingResponseBody {
	s.AudioAutoFreezeOpened = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetAudioMaxSize(v int32) *DescribeOssIncrementCheckSettingResponseBody {
	s.AudioMaxSize = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetAudioScanLimit(v int64) *DescribeOssIncrementCheckSettingResponseBody {
	s.AudioScanLimit = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetAudioSceneList(v []*string) *DescribeOssIncrementCheckSettingResponseBody {
	s.AudioSceneList = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetAutoFreezeType(v string) *DescribeOssIncrementCheckSettingResponseBody {
	s.AutoFreezeType = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetBizType(v string) *DescribeOssIncrementCheckSettingResponseBody {
	s.BizType = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetBizTypeTemplate(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) *DescribeOssIncrementCheckSettingResponseBody {
	s.BizTypeTemplate = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetBucketConfigList(v []*DescribeOssIncrementCheckSettingResponseBodyBucketConfigList) *DescribeOssIncrementCheckSettingResponseBody {
	s.BucketConfigList = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetCallbackId(v string) *DescribeOssIncrementCheckSettingResponseBody {
	s.CallbackId = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetCallbackName(v string) *DescribeOssIncrementCheckSettingResponseBody {
	s.CallbackName = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetEndTime(v string) *DescribeOssIncrementCheckSettingResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageAdFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageAdFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageAutoFreeze(v *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageAutoFreeze = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageAutoFreezeOpened(v bool) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageAutoFreezeOpened = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageEnableLimit(v bool) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageEnableLimit = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageLiveFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageLiveFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImagePornFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImagePornFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageScanLimit(v int64) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageScanLimit = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageSceneList(v []*string) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageSceneList = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetImageTerrorismFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.ImageTerrorismFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetRequestId(v string) *DescribeOssIncrementCheckSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetScanImageNoFileType(v bool) *DescribeOssIncrementCheckSettingResponseBody {
	s.ScanImageNoFileType = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetStartTime(v string) *DescribeOssIncrementCheckSettingResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoAdFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoAdFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoAutoFreezeOpened(v bool) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoAutoFreezeOpened = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoAutoFreezeSceneList(v []*string) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoAutoFreezeSceneList = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoFrameInterval(v int32) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoFrameInterval = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoLiveFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoLiveFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoMaxFrames(v int32) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoMaxFrames = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoMaxSize(v int32) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoMaxSize = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoPornFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoPornFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoScanLimit(v int64) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoScanLimit = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoSceneList(v []*string) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoSceneList = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoTerrorismFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoTerrorismFreezeConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBody) SetVideoVoiceAntispamFreezeConfig(v *DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig) *DescribeOssIncrementCheckSettingResponseBody {
	s.VideoVoiceAntispamFreezeConfig = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyAudioAntispamFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate struct {
	BizType        *string                                                                 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Description    *string                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageConfig    *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig `json:"ImageConfig,omitempty" xml:"ImageConfig,omitempty" type:"Struct"`
	IncludeChannel *int32                                                                  `json:"IncludeChannel,omitempty" xml:"IncludeChannel,omitempty"`
	Name           *string                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	VideoConfig    *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig `json:"VideoConfig,omitempty" xml:"VideoConfig,omitempty" type:"Struct"`
	VoiceConfig    *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfig `json:"VoiceConfig,omitempty" xml:"VoiceConfig,omitempty" type:"Struct"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) SetBizType(v string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate {
	s.BizType = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) SetDescription(v string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate {
	s.Description = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) SetImageConfig(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate {
	s.ImageConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) SetIncludeChannel(v int32) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate {
	s.IncludeChannel = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) SetName(v string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate {
	s.Name = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) SetVideoConfig(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate {
	s.VideoConfig = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate) SetVoiceConfig(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfig) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplate {
	s.VoiceConfig = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig struct {
	Ad        *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigAd        `json:"Ad,omitempty" xml:"Ad,omitempty" type:"Struct"`
	Live      *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigLive      `json:"Live,omitempty" xml:"Live,omitempty" type:"Struct"`
	Porn      *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigPorn      `json:"Porn,omitempty" xml:"Porn,omitempty" type:"Struct"`
	Terrorism *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigTerrorism `json:"Terrorism,omitempty" xml:"Terrorism,omitempty" type:"Struct"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig) SetAd(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigAd) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig {
	s.Ad = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig) SetLive(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigLive) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig {
	s.Live = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig) SetPorn(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigPorn) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig {
	s.Porn = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig) SetTerrorism(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigTerrorism) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfig {
	s.Terrorism = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigAd struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigAd) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigAd) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigAd) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigAd {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigLive struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigLive) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigLive) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigLive) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigLive {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigPorn struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigPorn) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigPorn) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigPorn) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigPorn {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigTerrorism struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigTerrorism) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigTerrorism) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigTerrorism) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateImageConfigTerrorism {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig struct {
	Ad        *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigAd        `json:"Ad,omitempty" xml:"Ad,omitempty" type:"Struct"`
	Live      *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigLive      `json:"Live,omitempty" xml:"Live,omitempty" type:"Struct"`
	Porn      *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigPorn      `json:"Porn,omitempty" xml:"Porn,omitempty" type:"Struct"`
	Terrorism *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigTerrorism `json:"Terrorism,omitempty" xml:"Terrorism,omitempty" type:"Struct"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig) SetAd(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigAd) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig {
	s.Ad = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig) SetLive(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigLive) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig {
	s.Live = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig) SetPorn(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigPorn) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig {
	s.Porn = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig) SetTerrorism(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigTerrorism) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfig {
	s.Terrorism = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigAd struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigAd) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigAd) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigAd) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigAd {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigLive struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigLive) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigLive) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigLive) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigLive {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigPorn struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigPorn) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigPorn) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigPorn) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigPorn {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigTerrorism struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigTerrorism) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigTerrorism) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigTerrorism) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVideoConfigTerrorism {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfig struct {
	Antispam *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfigAntispam `json:"Antispam,omitempty" xml:"Antispam,omitempty" type:"Struct"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfig) SetAntispam(v *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfigAntispam) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfig {
	s.Antispam = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfigAntispam struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfigAntispam) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfigAntispam) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfigAntispam) SetCategories(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBizTypeTemplateVoiceConfigAntispam {
	s.Categories = v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyBucketConfigList struct {
	Bucket   *string   `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Prefixes []*string `json:"Prefixes,omitempty" xml:"Prefixes,omitempty" type:"Repeated"`
	Selected *bool     `json:"Selected,omitempty" xml:"Selected,omitempty"`
	Type     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyBucketConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyBucketConfigList) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList) SetBucket(v string) *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList {
	s.Bucket = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList) SetPrefixes(v []*string) *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList {
	s.Prefixes = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList) SetSelected(v bool) *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList {
	s.Selected = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyBucketConfigList {
	s.Type = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyImageAdFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze struct {
	Ad        *string `json:"Ad,omitempty" xml:"Ad,omitempty"`
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Live      *string `json:"Live,omitempty" xml:"Live,omitempty"`
	Porn      *string `json:"Porn,omitempty" xml:"Porn,omitempty"`
	Terrorism *string `json:"Terrorism,omitempty" xml:"Terrorism,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) SetAd(v string) *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze {
	s.Ad = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) SetEnabled(v bool) *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze {
	s.Enabled = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) SetLive(v string) *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze {
	s.Live = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) SetPorn(v string) *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze {
	s.Porn = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze) SetTerrorism(v string) *DescribeOssIncrementCheckSettingResponseBodyImageAutoFreeze {
	s.Terrorism = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyImageLiveFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyImagePornFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyImageTerrorismFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoAdFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoLiveFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoPornFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoTerrorismFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig) SetType(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig {
	s.Type = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig) SetValue(v string) *DescribeOssIncrementCheckSettingResponseBodyVideoVoiceAntispamFreezeConfig {
	s.Value = &v
	return s
}

type DescribeOssIncrementCheckSettingResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssIncrementCheckSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssIncrementCheckSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementCheckSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementCheckSettingResponse) SetHeaders(v map[string]*string) *DescribeOssIncrementCheckSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponse) SetStatusCode(v int32) *DescribeOssIncrementCheckSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssIncrementCheckSettingResponse) SetBody(v *DescribeOssIncrementCheckSettingResponseBody) *DescribeOssIncrementCheckSettingResponse {
	s.Body = v
	return s
}

type DescribeOssIncrementOverviewRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeOssIncrementOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementOverviewRequest) SetLang(v string) *DescribeOssIncrementOverviewRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssIncrementOverviewRequest) SetSourceIp(v string) *DescribeOssIncrementOverviewRequest {
	s.SourceIp = &v
	return s
}

type DescribeOssIncrementOverviewResponseBody struct {
	AdUnhandleCount            *int32  `json:"AdUnhandleCount,omitempty" xml:"AdUnhandleCount,omitempty"`
	AudioCount                 *int32  `json:"AudioCount,omitempty" xml:"AudioCount,omitempty"`
	ImageCount                 *int32  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	LiveUnhandleCount          *int32  `json:"LiveUnhandleCount,omitempty" xml:"LiveUnhandleCount,omitempty"`
	PornUnhandleCount          *int32  `json:"PornUnhandleCount,omitempty" xml:"PornUnhandleCount,omitempty"`
	RequestId                  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TerrorismUnhandleCount     *int32  `json:"TerrorismUnhandleCount,omitempty" xml:"TerrorismUnhandleCount,omitempty"`
	VideoCount                 *int32  `json:"VideoCount,omitempty" xml:"VideoCount,omitempty"`
	VideoFrameCount            *int32  `json:"VideoFrameCount,omitempty" xml:"VideoFrameCount,omitempty"`
	VoiceAntispamUnhandleCount *int32  `json:"VoiceAntispamUnhandleCount,omitempty" xml:"VoiceAntispamUnhandleCount,omitempty"`
}

func (s DescribeOssIncrementOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementOverviewResponseBody) SetAdUnhandleCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.AdUnhandleCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetAudioCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.AudioCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetImageCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.ImageCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetLiveUnhandleCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.LiveUnhandleCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetPornUnhandleCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.PornUnhandleCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetRequestId(v string) *DescribeOssIncrementOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetTerrorismUnhandleCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.TerrorismUnhandleCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetVideoCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.VideoCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetVideoFrameCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.VideoFrameCount = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponseBody) SetVoiceAntispamUnhandleCount(v int32) *DescribeOssIncrementOverviewResponseBody {
	s.VoiceAntispamUnhandleCount = &v
	return s
}

type DescribeOssIncrementOverviewResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssIncrementOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssIncrementOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementOverviewResponse) SetHeaders(v map[string]*string) *DescribeOssIncrementOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssIncrementOverviewResponse) SetStatusCode(v int32) *DescribeOssIncrementOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssIncrementOverviewResponse) SetBody(v *DescribeOssIncrementOverviewResponseBody) *DescribeOssIncrementOverviewResponse {
	s.Body = v
	return s
}

type DescribeOssIncrementStatsRequest struct {
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeOssIncrementStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementStatsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementStatsRequest) SetEndDate(v string) *DescribeOssIncrementStatsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeOssIncrementStatsRequest) SetLang(v string) *DescribeOssIncrementStatsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssIncrementStatsRequest) SetResourceType(v string) *DescribeOssIncrementStatsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeOssIncrementStatsRequest) SetScene(v string) *DescribeOssIncrementStatsRequest {
	s.Scene = &v
	return s
}

func (s *DescribeOssIncrementStatsRequest) SetSourceIp(v string) *DescribeOssIncrementStatsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOssIncrementStatsRequest) SetStartDate(v string) *DescribeOssIncrementStatsRequest {
	s.StartDate = &v
	return s
}

type DescribeOssIncrementStatsResponseBody struct {
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList   []*DescribeOssIncrementStatsResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOssIncrementStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementStatsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementStatsResponseBody) SetRequestId(v string) *DescribeOssIncrementStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssIncrementStatsResponseBody) SetStatList(v []*DescribeOssIncrementStatsResponseBodyStatList) *DescribeOssIncrementStatsResponseBody {
	s.StatList = v
	return s
}

func (s *DescribeOssIncrementStatsResponseBody) SetTotalCount(v int32) *DescribeOssIncrementStatsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOssIncrementStatsResponseBodyStatList struct {
	BlockCount   *int32  `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	Date         *string `json:"Date,omitempty" xml:"Date,omitempty"`
	PassCount    *int32  `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ReviewCount  *int32  `json:"ReviewCount,omitempty" xml:"ReviewCount,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	TotalCount   *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOssIncrementStatsResponseBodyStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementStatsResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementStatsResponseBodyStatList) SetBlockCount(v int32) *DescribeOssIncrementStatsResponseBodyStatList {
	s.BlockCount = &v
	return s
}

func (s *DescribeOssIncrementStatsResponseBodyStatList) SetDate(v string) *DescribeOssIncrementStatsResponseBodyStatList {
	s.Date = &v
	return s
}

func (s *DescribeOssIncrementStatsResponseBodyStatList) SetPassCount(v int32) *DescribeOssIncrementStatsResponseBodyStatList {
	s.PassCount = &v
	return s
}

func (s *DescribeOssIncrementStatsResponseBodyStatList) SetResourceType(v string) *DescribeOssIncrementStatsResponseBodyStatList {
	s.ResourceType = &v
	return s
}

func (s *DescribeOssIncrementStatsResponseBodyStatList) SetReviewCount(v int32) *DescribeOssIncrementStatsResponseBodyStatList {
	s.ReviewCount = &v
	return s
}

func (s *DescribeOssIncrementStatsResponseBodyStatList) SetScene(v string) *DescribeOssIncrementStatsResponseBodyStatList {
	s.Scene = &v
	return s
}

func (s *DescribeOssIncrementStatsResponseBodyStatList) SetTotalCount(v int32) *DescribeOssIncrementStatsResponseBodyStatList {
	s.TotalCount = &v
	return s
}

type DescribeOssIncrementStatsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssIncrementStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssIncrementStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssIncrementStatsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssIncrementStatsResponse) SetHeaders(v map[string]*string) *DescribeOssIncrementStatsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssIncrementStatsResponse) SetStatusCode(v int32) *DescribeOssIncrementStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssIncrementStatsResponse) SetBody(v *DescribeOssIncrementStatsResponseBody) *DescribeOssIncrementStatsResponse {
	s.Body = v
	return s
}

type DescribeOssResultItemsRequest struct {
	Bucket       *string  `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	CurrentPage  *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndDate      *string  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang         *string  `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MaxScore     *float32 `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	MinScore     *float32 `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
	Object       *string  `json:"Object,omitempty" xml:"Object,omitempty"`
	PageSize     *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryId      *string  `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	ResourceType *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string  `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceIp     *string  `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate    *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Stock        *bool    `json:"Stock,omitempty" xml:"Stock,omitempty"`
	StockTaskId  *int64   `json:"StockTaskId,omitempty" xml:"StockTaskId,omitempty"`
	Suggestion   *string  `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TotalCount   *int32   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOssResultItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssResultItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssResultItemsRequest) SetBucket(v string) *DescribeOssResultItemsRequest {
	s.Bucket = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetCurrentPage(v int32) *DescribeOssResultItemsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetEndDate(v string) *DescribeOssResultItemsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetLang(v string) *DescribeOssResultItemsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetMaxScore(v float32) *DescribeOssResultItemsRequest {
	s.MaxScore = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetMinScore(v float32) *DescribeOssResultItemsRequest {
	s.MinScore = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetObject(v string) *DescribeOssResultItemsRequest {
	s.Object = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetPageSize(v int32) *DescribeOssResultItemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetQueryId(v string) *DescribeOssResultItemsRequest {
	s.QueryId = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetResourceType(v string) *DescribeOssResultItemsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetScene(v string) *DescribeOssResultItemsRequest {
	s.Scene = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetSourceIp(v string) *DescribeOssResultItemsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetStartDate(v string) *DescribeOssResultItemsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetStock(v bool) *DescribeOssResultItemsRequest {
	s.Stock = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetStockTaskId(v int64) *DescribeOssResultItemsRequest {
	s.StockTaskId = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetSuggestion(v string) *DescribeOssResultItemsRequest {
	s.Suggestion = &v
	return s
}

func (s *DescribeOssResultItemsRequest) SetTotalCount(v int32) *DescribeOssResultItemsRequest {
	s.TotalCount = &v
	return s
}

type DescribeOssResultItemsResponseBody struct {
	CurrentPage    *int32                                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize       *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryId        *string                                             `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanResultList []*DescribeOssResultItemsResponseBodyScanResultList `json:"ScanResultList,omitempty" xml:"ScanResultList,omitempty" type:"Repeated"`
	TotalCount     *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOssResultItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssResultItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssResultItemsResponseBody) SetCurrentPage(v int32) *DescribeOssResultItemsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOssResultItemsResponseBody) SetPageSize(v int32) *DescribeOssResultItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOssResultItemsResponseBody) SetQueryId(v string) *DescribeOssResultItemsResponseBody {
	s.QueryId = &v
	return s
}

func (s *DescribeOssResultItemsResponseBody) SetRequestId(v string) *DescribeOssResultItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssResultItemsResponseBody) SetScanResultList(v []*DescribeOssResultItemsResponseBodyScanResultList) *DescribeOssResultItemsResponseBody {
	s.ScanResultList = v
	return s
}

func (s *DescribeOssResultItemsResponseBody) SetTotalCount(v int32) *DescribeOssResultItemsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOssResultItemsResponseBodyScanResultList struct {
	Bucket                      *string                                                                        `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Content                     *string                                                                        `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime                  *string                                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataId                      *string                                                                        `json:"DataId,omitempty" xml:"DataId,omitempty"`
	FrameResults                []*DescribeOssResultItemsResponseBodyScanResultListFrameResults                `json:"FrameResults,omitempty" xml:"FrameResults,omitempty" type:"Repeated"`
	HandleStatus                *int32                                                                         `json:"HandleStatus,omitempty" xml:"HandleStatus,omitempty"`
	Id                          *int64                                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	NewUrl                      *string                                                                        `json:"NewUrl,omitempty" xml:"NewUrl,omitempty"`
	Object                      *string                                                                        `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestTime                 *string                                                                        `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	ResourceStatus              *int32                                                                         `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ScanFinishedTime            *string                                                                        `json:"ScanFinishedTime,omitempty" xml:"ScanFinishedTime,omitempty"`
	Score                       *string                                                                        `json:"Score,omitempty" xml:"Score,omitempty"`
	Suggestion                  *string                                                                        `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TaskId                      *string                                                                        `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Thumbnail                   *string                                                                        `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	VoiceSegmentAntispamResults []*DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults `json:"VoiceSegmentAntispamResults,omitempty" xml:"VoiceSegmentAntispamResults,omitempty" type:"Repeated"`
}

func (s DescribeOssResultItemsResponseBodyScanResultList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssResultItemsResponseBodyScanResultList) GoString() string {
	return s.String()
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetBucket(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.Bucket = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetContent(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.Content = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetCreateTime(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.CreateTime = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetDataId(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.DataId = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetFrameResults(v []*DescribeOssResultItemsResponseBodyScanResultListFrameResults) *DescribeOssResultItemsResponseBodyScanResultList {
	s.FrameResults = v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetHandleStatus(v int32) *DescribeOssResultItemsResponseBodyScanResultList {
	s.HandleStatus = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetId(v int64) *DescribeOssResultItemsResponseBodyScanResultList {
	s.Id = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetNewUrl(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.NewUrl = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetObject(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.Object = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetRequestTime(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.RequestTime = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetResourceStatus(v int32) *DescribeOssResultItemsResponseBodyScanResultList {
	s.ResourceStatus = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetScanFinishedTime(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.ScanFinishedTime = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetScore(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.Score = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetSuggestion(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.Suggestion = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetTaskId(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.TaskId = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetThumbnail(v string) *DescribeOssResultItemsResponseBodyScanResultList {
	s.Thumbnail = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultList) SetVoiceSegmentAntispamResults(v []*DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults) *DescribeOssResultItemsResponseBodyScanResultList {
	s.VoiceSegmentAntispamResults = v
	return s
}

type DescribeOssResultItemsResponseBodyScanResultListFrameResults struct {
	Offset    *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeOssResultItemsResponseBodyScanResultListFrameResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssResultItemsResponseBodyScanResultListFrameResults) GoString() string {
	return s.String()
}

func (s *DescribeOssResultItemsResponseBodyScanResultListFrameResults) SetOffset(v int32) *DescribeOssResultItemsResponseBodyScanResultListFrameResults {
	s.Offset = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultListFrameResults) SetThumbnail(v string) *DescribeOssResultItemsResponseBodyScanResultListFrameResults {
	s.Thumbnail = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultListFrameResults) SetUrl(v string) *DescribeOssResultItemsResponseBodyScanResultListFrameResults {
	s.Url = &v
	return s
}

type DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults) GoString() string {
	return s.String()
}

func (s *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults) SetEndTime(v int32) *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults {
	s.EndTime = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults) SetLabel(v string) *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults {
	s.Label = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults) SetStartTime(v int32) *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults {
	s.StartTime = &v
	return s
}

func (s *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults) SetText(v string) *DescribeOssResultItemsResponseBodyScanResultListVoiceSegmentAntispamResults {
	s.Text = &v
	return s
}

type DescribeOssResultItemsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssResultItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssResultItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssResultItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssResultItemsResponse) SetHeaders(v map[string]*string) *DescribeOssResultItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssResultItemsResponse) SetStatusCode(v int32) *DescribeOssResultItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssResultItemsResponse) SetBody(v *DescribeOssResultItemsResponseBody) *DescribeOssResultItemsResponse {
	s.Body = v
	return s
}

type DescribeOssStockStatusRequest struct {
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StockTaskId *int64  `json:"StockTaskId,omitempty" xml:"StockTaskId,omitempty"`
}

func (s DescribeOssStockStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssStockStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssStockStatusRequest) SetLang(v string) *DescribeOssStockStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssStockStatusRequest) SetSourceIp(v string) *DescribeOssStockStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOssStockStatusRequest) SetStockTaskId(v int64) *DescribeOssStockStatusRequest {
	s.StockTaskId = &v
	return s
}

type DescribeOssStockStatusResponseBody struct {
	AudioAntispamCount      *int32                                          `json:"AudioAntispamCount,omitempty" xml:"AudioAntispamCount,omitempty"`
	AudioTotalCount         *int32                                          `json:"AudioTotalCount,omitempty" xml:"AudioTotalCount,omitempty"`
	BucketList              []*DescribeOssStockStatusResponseBodyBucketList `json:"BucketList,omitempty" xml:"BucketList,omitempty" type:"Repeated"`
	FinishedTime            *string                                         `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	ImageAdCount            *int32                                          `json:"ImageAdCount,omitempty" xml:"ImageAdCount,omitempty"`
	ImageLiveCount          *int32                                          `json:"ImageLiveCount,omitempty" xml:"ImageLiveCount,omitempty"`
	ImagePornCount          *int32                                          `json:"ImagePornCount,omitempty" xml:"ImagePornCount,omitempty"`
	ImageTerrorismCount     *int32                                          `json:"ImageTerrorismCount,omitempty" xml:"ImageTerrorismCount,omitempty"`
	ImageTotalCount         *int32                                          `json:"ImageTotalCount,omitempty" xml:"ImageTotalCount,omitempty"`
	RequestId               *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResouceTypeList         []*string                                       `json:"ResouceTypeList,omitempty" xml:"ResouceTypeList,omitempty" type:"Repeated"`
	SceneList               []*string                                       `json:"SceneList,omitempty" xml:"SceneList,omitempty" type:"Repeated"`
	StockStatus             *int32                                          `json:"StockStatus,omitempty" xml:"StockStatus,omitempty"`
	VideoAdCount            *int32                                          `json:"VideoAdCount,omitempty" xml:"VideoAdCount,omitempty"`
	VideoLiveCount          *int32                                          `json:"VideoLiveCount,omitempty" xml:"VideoLiveCount,omitempty"`
	VideoPornCount          *int32                                          `json:"VideoPornCount,omitempty" xml:"VideoPornCount,omitempty"`
	VideoTerrorismCount     *int32                                          `json:"VideoTerrorismCount,omitempty" xml:"VideoTerrorismCount,omitempty"`
	VideoTotalCount         *int32                                          `json:"VideoTotalCount,omitempty" xml:"VideoTotalCount,omitempty"`
	VideoVoiceAntispamCount *int32                                          `json:"VideoVoiceAntispamCount,omitempty" xml:"VideoVoiceAntispamCount,omitempty"`
}

func (s DescribeOssStockStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssStockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssStockStatusResponseBody) SetAudioAntispamCount(v int32) *DescribeOssStockStatusResponseBody {
	s.AudioAntispamCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetAudioTotalCount(v int32) *DescribeOssStockStatusResponseBody {
	s.AudioTotalCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetBucketList(v []*DescribeOssStockStatusResponseBodyBucketList) *DescribeOssStockStatusResponseBody {
	s.BucketList = v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetFinishedTime(v string) *DescribeOssStockStatusResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetImageAdCount(v int32) *DescribeOssStockStatusResponseBody {
	s.ImageAdCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetImageLiveCount(v int32) *DescribeOssStockStatusResponseBody {
	s.ImageLiveCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetImagePornCount(v int32) *DescribeOssStockStatusResponseBody {
	s.ImagePornCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetImageTerrorismCount(v int32) *DescribeOssStockStatusResponseBody {
	s.ImageTerrorismCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetImageTotalCount(v int32) *DescribeOssStockStatusResponseBody {
	s.ImageTotalCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetRequestId(v string) *DescribeOssStockStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetResouceTypeList(v []*string) *DescribeOssStockStatusResponseBody {
	s.ResouceTypeList = v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetSceneList(v []*string) *DescribeOssStockStatusResponseBody {
	s.SceneList = v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetStockStatus(v int32) *DescribeOssStockStatusResponseBody {
	s.StockStatus = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetVideoAdCount(v int32) *DescribeOssStockStatusResponseBody {
	s.VideoAdCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetVideoLiveCount(v int32) *DescribeOssStockStatusResponseBody {
	s.VideoLiveCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetVideoPornCount(v int32) *DescribeOssStockStatusResponseBody {
	s.VideoPornCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetVideoTerrorismCount(v int32) *DescribeOssStockStatusResponseBody {
	s.VideoTerrorismCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetVideoTotalCount(v int32) *DescribeOssStockStatusResponseBody {
	s.VideoTotalCount = &v
	return s
}

func (s *DescribeOssStockStatusResponseBody) SetVideoVoiceAntispamCount(v int32) *DescribeOssStockStatusResponseBody {
	s.VideoVoiceAntispamCount = &v
	return s
}

type DescribeOssStockStatusResponseBodyBucketList struct {
	Bucket   *string   `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Prefixes []*string `json:"Prefixes,omitempty" xml:"Prefixes,omitempty" type:"Repeated"`
	Selected *bool     `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeOssStockStatusResponseBodyBucketList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssStockStatusResponseBodyBucketList) GoString() string {
	return s.String()
}

func (s *DescribeOssStockStatusResponseBodyBucketList) SetBucket(v string) *DescribeOssStockStatusResponseBodyBucketList {
	s.Bucket = &v
	return s
}

func (s *DescribeOssStockStatusResponseBodyBucketList) SetPrefixes(v []*string) *DescribeOssStockStatusResponseBodyBucketList {
	s.Prefixes = v
	return s
}

func (s *DescribeOssStockStatusResponseBodyBucketList) SetSelected(v bool) *DescribeOssStockStatusResponseBodyBucketList {
	s.Selected = &v
	return s
}

type DescribeOssStockStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssStockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssStockStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssStockStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssStockStatusResponse) SetHeaders(v map[string]*string) *DescribeOssStockStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssStockStatusResponse) SetStatusCode(v int32) *DescribeOssStockStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssStockStatusResponse) SetBody(v *DescribeOssStockStatusResponseBody) *DescribeOssStockStatusResponse {
	s.Body = v
	return s
}

type DescribeSdkUrlRequest struct {
	Debug    *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSdkUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlRequest) SetDebug(v bool) *DescribeSdkUrlRequest {
	s.Debug = &v
	return s
}

func (s *DescribeSdkUrlRequest) SetId(v int64) *DescribeSdkUrlRequest {
	s.Id = &v
	return s
}

func (s *DescribeSdkUrlRequest) SetLang(v string) *DescribeSdkUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSdkUrlRequest) SetSourceIp(v string) *DescribeSdkUrlRequest {
	s.SourceIp = &v
	return s
}

type DescribeSdkUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkUrl    *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
}

func (s DescribeSdkUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlResponseBody) SetRequestId(v string) *DescribeSdkUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSdkUrlResponseBody) SetSdkUrl(v string) *DescribeSdkUrlResponseBody {
	s.SdkUrl = &v
	return s
}

type DescribeSdkUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSdkUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSdkUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSdkUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdkUrlResponse) SetHeaders(v map[string]*string) *DescribeSdkUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdkUrlResponse) SetStatusCode(v int32) *DescribeSdkUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSdkUrlResponse) SetBody(v *DescribeSdkUrlResponseBody) *DescribeSdkUrlResponse {
	s.Body = v
	return s
}

type DescribeUpdatePackageResultRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeUpdatePackageResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultRequest) SetLang(v string) *DescribeUpdatePackageResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUpdatePackageResultRequest) SetSourceIp(v string) *DescribeUpdatePackageResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUpdatePackageResultRequest) SetTaskId(v string) *DescribeUpdatePackageResultRequest {
	s.TaskId = &v
	return s
}

type DescribeUpdatePackageResultResponseBody struct {
	AppInfo   *DescribeUpdatePackageResultResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBody) SetAppInfo(v *DescribeUpdatePackageResultResponseBodyAppInfo) *DescribeUpdatePackageResultResponseBody {
	s.AppInfo = v
	return s
}

func (s *DescribeUpdatePackageResultResponseBody) SetRequestId(v string) *DescribeUpdatePackageResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUpdatePackageResultResponseBodyAppInfo struct {
	DebugPackageInfo *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo `json:"DebugPackageInfo,omitempty" xml:"DebugPackageInfo,omitempty" type:"Struct"`
	EndDate          *string                                                         `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Icon             *string                                                         `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Id               *int64                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PackageInfo      *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo      `json:"PackageInfo,omitempty" xml:"PackageInfo,omitempty" type:"Struct"`
	PackageName      *string                                                         `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	StartDate        *string                                                         `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Type             *int32                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBodyAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetDebugPackageInfo(v *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.DebugPackageInfo = v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetEndDate(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.EndDate = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetIcon(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Icon = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetId(v int64) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Id = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetName(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Name = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetPackageInfo(v *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.PackageInfo = v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetPackageName(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.PackageName = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetStartDate(v string) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.StartDate = &v
	return s
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfo) SetType(v int32) *DescribeUpdatePackageResultResponseBodyAppInfo {
	s.Type = &v
	return s
}

type DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo) SetVersion(v string) *DescribeUpdatePackageResultResponseBodyAppInfoDebugPackageInfo {
	s.Version = &v
	return s
}

type DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo) SetVersion(v string) *DescribeUpdatePackageResultResponseBodyAppInfoPackageInfo {
	s.Version = &v
	return s
}

type DescribeUpdatePackageResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpdatePackageResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpdatePackageResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdatePackageResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpdatePackageResultResponse) SetHeaders(v map[string]*string) *DescribeUpdatePackageResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpdatePackageResultResponse) SetStatusCode(v int32) *DescribeUpdatePackageResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpdatePackageResultResponse) SetBody(v *DescribeUpdatePackageResultResponseBody) *DescribeUpdatePackageResultResponse {
	s.Body = v
	return s
}

type DescribeUploadInfoRequest struct {
	Biz      *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUploadInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUploadInfoRequest) SetBiz(v string) *DescribeUploadInfoRequest {
	s.Biz = &v
	return s
}

func (s *DescribeUploadInfoRequest) SetLang(v string) *DescribeUploadInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUploadInfoRequest) SetSourceIp(v string) *DescribeUploadInfoRequest {
	s.SourceIp = &v
	return s
}

type DescribeUploadInfoResponseBody struct {
	Accessid  *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	Expire    *int32  `json:"Expire,omitempty" xml:"Expire,omitempty"`
	Folder    *string `json:"Folder,omitempty" xml:"Folder,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s DescribeUploadInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUploadInfoResponseBody) SetAccessid(v string) *DescribeUploadInfoResponseBody {
	s.Accessid = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetExpire(v int32) *DescribeUploadInfoResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetFolder(v string) *DescribeUploadInfoResponseBody {
	s.Folder = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetHost(v string) *DescribeUploadInfoResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetPolicy(v string) *DescribeUploadInfoResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetRequestId(v string) *DescribeUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUploadInfoResponseBody) SetSignature(v string) *DescribeUploadInfoResponseBody {
	s.Signature = &v
	return s
}

type DescribeUploadInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUploadInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUploadInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUploadInfoResponse) SetHeaders(v map[string]*string) *DescribeUploadInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUploadInfoResponse) SetStatusCode(v int32) *DescribeUploadInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUploadInfoResponse) SetBody(v *DescribeUploadInfoResponseBody) *DescribeUploadInfoResponse {
	s.Body = v
	return s
}

type DescribeUsageBillRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Day         *string `json:"Day,omitempty" xml:"Day,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUsageBillRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageBillRequest) SetCurrentPage(v int32) *DescribeUsageBillRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUsageBillRequest) SetDay(v string) *DescribeUsageBillRequest {
	s.Day = &v
	return s
}

func (s *DescribeUsageBillRequest) SetPageSize(v int32) *DescribeUsageBillRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUsageBillRequest) SetTotalCount(v int32) *DescribeUsageBillRequest {
	s.TotalCount = &v
	return s
}

func (s *DescribeUsageBillRequest) SetType(v string) *DescribeUsageBillRequest {
	s.Type = &v
	return s
}

type DescribeUsageBillResponseBody struct {
	BillList    []*DescribeUsageBillResponseBodyBillList `json:"BillList,omitempty" xml:"BillList,omitempty" type:"Repeated"`
	CurrentPage *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUsageBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageBillResponseBody) SetBillList(v []*DescribeUsageBillResponseBodyBillList) *DescribeUsageBillResponseBody {
	s.BillList = v
	return s
}

func (s *DescribeUsageBillResponseBody) SetCurrentPage(v int32) *DescribeUsageBillResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUsageBillResponseBody) SetPageSize(v int32) *DescribeUsageBillResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUsageBillResponseBody) SetRequestId(v string) *DescribeUsageBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageBillResponseBody) SetTotalCount(v int32) *DescribeUsageBillResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeUsageBillResponseBodyBillList struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ConfirmCount *int64  `json:"ConfirmCount,omitempty" xml:"ConfirmCount,omitempty"`
	Day          *string `json:"Day,omitempty" xml:"Day,omitempty"`
	FreeCount    *int64  `json:"FreeCount,omitempty" xml:"FreeCount,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ReviewCount  *int64  `json:"ReviewCount,omitempty" xml:"ReviewCount,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SubUid       *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	TotalCount   *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUsageBillResponseBodyBillList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageBillResponseBodyBillList) GoString() string {
	return s.String()
}

func (s *DescribeUsageBillResponseBodyBillList) SetBizType(v string) *DescribeUsageBillResponseBodyBillList {
	s.BizType = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetConfirmCount(v int64) *DescribeUsageBillResponseBodyBillList {
	s.ConfirmCount = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetDay(v string) *DescribeUsageBillResponseBodyBillList {
	s.Day = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetFreeCount(v int64) *DescribeUsageBillResponseBodyBillList {
	s.FreeCount = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetRegion(v string) *DescribeUsageBillResponseBodyBillList {
	s.Region = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetReviewCount(v int64) *DescribeUsageBillResponseBodyBillList {
	s.ReviewCount = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetScene(v string) *DescribeUsageBillResponseBodyBillList {
	s.Scene = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetSubUid(v string) *DescribeUsageBillResponseBodyBillList {
	s.SubUid = &v
	return s
}

func (s *DescribeUsageBillResponseBodyBillList) SetTotalCount(v int64) *DescribeUsageBillResponseBodyBillList {
	s.TotalCount = &v
	return s
}

type DescribeUsageBillResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageBillResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageBillResponse) SetHeaders(v map[string]*string) *DescribeUsageBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageBillResponse) SetStatusCode(v int32) *DescribeUsageBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageBillResponse) SetBody(v *DescribeUsageBillResponseBody) *DescribeUsageBillResponse {
	s.Body = v
	return s
}

type DescribeUserBizTypesRequest struct {
	Customized *bool `json:"Customized,omitempty" xml:"Customized,omitempty"`
}

func (s DescribeUserBizTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBizTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBizTypesRequest) SetCustomized(v bool) *DescribeUserBizTypesRequest {
	s.Customized = &v
	return s
}

type DescribeUserBizTypesResponseBody struct {
	BizTypeList       []*DescribeUserBizTypesResponseBodyBizTypeList       `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	BizTypeListImport []*DescribeUserBizTypesResponseBodyBizTypeListImport `json:"BizTypeListImport,omitempty" xml:"BizTypeListImport,omitempty" type:"Repeated"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserBizTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBizTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBizTypesResponseBody) SetBizTypeList(v []*DescribeUserBizTypesResponseBodyBizTypeList) *DescribeUserBizTypesResponseBody {
	s.BizTypeList = v
	return s
}

func (s *DescribeUserBizTypesResponseBody) SetBizTypeListImport(v []*DescribeUserBizTypesResponseBodyBizTypeListImport) *DescribeUserBizTypesResponseBody {
	s.BizTypeListImport = v
	return s
}

func (s *DescribeUserBizTypesResponseBody) SetRequestId(v string) *DescribeUserBizTypesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUserBizTypesResponseBodyBizTypeList struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CiteTemplate  *bool   `json:"CiteTemplate,omitempty" xml:"CiteTemplate,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Gray          *bool   `json:"Gray,omitempty" xml:"Gray,omitempty"`
	IndustryInfo  *string `json:"IndustryInfo,omitempty" xml:"IndustryInfo,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceBizType *string `json:"SourceBizType,omitempty" xml:"SourceBizType,omitempty"`
}

func (s DescribeUserBizTypesResponseBodyBizTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBizTypesResponseBodyBizTypeList) GoString() string {
	return s.String()
}

func (s *DescribeUserBizTypesResponseBodyBizTypeList) SetBizType(v string) *DescribeUserBizTypesResponseBodyBizTypeList {
	s.BizType = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeList) SetCiteTemplate(v bool) *DescribeUserBizTypesResponseBodyBizTypeList {
	s.CiteTemplate = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeList) SetDescription(v string) *DescribeUserBizTypesResponseBodyBizTypeList {
	s.Description = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeList) SetGray(v bool) *DescribeUserBizTypesResponseBodyBizTypeList {
	s.Gray = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeList) SetIndustryInfo(v string) *DescribeUserBizTypesResponseBodyBizTypeList {
	s.IndustryInfo = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeList) SetSource(v string) *DescribeUserBizTypesResponseBodyBizTypeList {
	s.Source = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeList) SetSourceBizType(v string) *DescribeUserBizTypesResponseBodyBizTypeList {
	s.SourceBizType = &v
	return s
}

type DescribeUserBizTypesResponseBodyBizTypeListImport struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CiteTemplate  *bool   `json:"CiteTemplate,omitempty" xml:"CiteTemplate,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Gray          *bool   `json:"Gray,omitempty" xml:"Gray,omitempty"`
	IndustryInfo  *string `json:"IndustryInfo,omitempty" xml:"IndustryInfo,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceBizType *string `json:"SourceBizType,omitempty" xml:"SourceBizType,omitempty"`
}

func (s DescribeUserBizTypesResponseBodyBizTypeListImport) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBizTypesResponseBodyBizTypeListImport) GoString() string {
	return s.String()
}

func (s *DescribeUserBizTypesResponseBodyBizTypeListImport) SetBizType(v string) *DescribeUserBizTypesResponseBodyBizTypeListImport {
	s.BizType = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeListImport) SetCiteTemplate(v bool) *DescribeUserBizTypesResponseBodyBizTypeListImport {
	s.CiteTemplate = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeListImport) SetDescription(v string) *DescribeUserBizTypesResponseBodyBizTypeListImport {
	s.Description = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeListImport) SetGray(v bool) *DescribeUserBizTypesResponseBodyBizTypeListImport {
	s.Gray = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeListImport) SetIndustryInfo(v string) *DescribeUserBizTypesResponseBodyBizTypeListImport {
	s.IndustryInfo = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeListImport) SetSource(v string) *DescribeUserBizTypesResponseBodyBizTypeListImport {
	s.Source = &v
	return s
}

func (s *DescribeUserBizTypesResponseBodyBizTypeListImport) SetSourceBizType(v string) *DescribeUserBizTypesResponseBodyBizTypeListImport {
	s.SourceBizType = &v
	return s
}

type DescribeUserBizTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBizTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBizTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBizTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBizTypesResponse) SetHeaders(v map[string]*string) *DescribeUserBizTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBizTypesResponse) SetStatusCode(v int32) *DescribeUserBizTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBizTypesResponse) SetBody(v *DescribeUserBizTypesResponseBody) *DescribeUserBizTypesResponse {
	s.Body = v
	return s
}

type DescribeUserStatusRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusRequest) SetLang(v string) *DescribeUserStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserStatusRequest) SetSourceIp(v string) *DescribeUserStatusRequest {
	s.SourceIp = &v
	return s
}

type DescribeUserStatusResponseBody struct {
	AgreeChannel      *int32  `json:"AgreeChannel,omitempty" xml:"AgreeChannel,omitempty"`
	Buyed             *bool   `json:"Buyed,omitempty" xml:"Buyed,omitempty"`
	InDept            *bool   `json:"InDept,omitempty" xml:"InDept,omitempty"`
	OpenApiBeginTime  *string `json:"OpenApiBeginTime,omitempty" xml:"OpenApiBeginTime,omitempty"`
	OpenApiUsed       *bool   `json:"OpenApiUsed,omitempty" xml:"OpenApiUsed,omitempty"`
	OssCheckStatus    *string `json:"OssCheckStatus,omitempty" xml:"OssCheckStatus,omitempty"`
	OssVideoSizeLimit *int32  `json:"OssVideoSizeLimit,omitempty" xml:"OssVideoSizeLimit,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Uid               *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponseBody) SetAgreeChannel(v int32) *DescribeUserStatusResponseBody {
	s.AgreeChannel = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetBuyed(v bool) *DescribeUserStatusResponseBody {
	s.Buyed = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetInDept(v bool) *DescribeUserStatusResponseBody {
	s.InDept = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetOpenApiBeginTime(v string) *DescribeUserStatusResponseBody {
	s.OpenApiBeginTime = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetOpenApiUsed(v bool) *DescribeUserStatusResponseBody {
	s.OpenApiUsed = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetOssCheckStatus(v string) *DescribeUserStatusResponseBody {
	s.OssCheckStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetOssVideoSizeLimit(v int32) *DescribeUserStatusResponseBody {
	s.OssVideoSizeLimit = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetRequestId(v string) *DescribeUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetUid(v string) *DescribeUserStatusResponseBody {
	s.Uid = &v
	return s
}

type DescribeUserStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponse) SetHeaders(v map[string]*string) *DescribeUserStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserStatusResponse) SetStatusCode(v int32) *DescribeUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserStatusResponse) SetBody(v *DescribeUserStatusResponseBody) *DescribeUserStatusResponse {
	s.Body = v
	return s
}

type DescribeViewContentRequest struct {
	AuditResult  *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DataId       *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	KeywordId    *string `json:"KeywordId,omitempty" xml:"KeywordId,omitempty"`
	Label        *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LibType      *string `json:"LibType,omitempty" xml:"LibType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Suggestion   *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TotalCount   *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeViewContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeViewContentRequest) SetAuditResult(v string) *DescribeViewContentRequest {
	s.AuditResult = &v
	return s
}

func (s *DescribeViewContentRequest) SetBizType(v string) *DescribeViewContentRequest {
	s.BizType = &v
	return s
}

func (s *DescribeViewContentRequest) SetCurrentPage(v int32) *DescribeViewContentRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeViewContentRequest) SetDataId(v string) *DescribeViewContentRequest {
	s.DataId = &v
	return s
}

func (s *DescribeViewContentRequest) SetEndDate(v string) *DescribeViewContentRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeViewContentRequest) SetImageId(v string) *DescribeViewContentRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeViewContentRequest) SetKeyword(v string) *DescribeViewContentRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeViewContentRequest) SetKeywordId(v string) *DescribeViewContentRequest {
	s.KeywordId = &v
	return s
}

func (s *DescribeViewContentRequest) SetLabel(v string) *DescribeViewContentRequest {
	s.Label = &v
	return s
}

func (s *DescribeViewContentRequest) SetLibType(v string) *DescribeViewContentRequest {
	s.LibType = &v
	return s
}

func (s *DescribeViewContentRequest) SetPageSize(v int32) *DescribeViewContentRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeViewContentRequest) SetResourceType(v string) *DescribeViewContentRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeViewContentRequest) SetScene(v string) *DescribeViewContentRequest {
	s.Scene = &v
	return s
}

func (s *DescribeViewContentRequest) SetStartDate(v string) *DescribeViewContentRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeViewContentRequest) SetSuggestion(v string) *DescribeViewContentRequest {
	s.Suggestion = &v
	return s
}

func (s *DescribeViewContentRequest) SetTaskId(v string) *DescribeViewContentRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeViewContentRequest) SetTotalCount(v int32) *DescribeViewContentRequest {
	s.TotalCount = &v
	return s
}

type DescribeViewContentResponseBody struct {
	CurrentPage     *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ViewContentList []*DescribeViewContentResponseBodyViewContentList `json:"ViewContentList,omitempty" xml:"ViewContentList,omitempty" type:"Repeated"`
}

func (s DescribeViewContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBody) SetCurrentPage(v int32) *DescribeViewContentResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeViewContentResponseBody) SetPageSize(v int32) *DescribeViewContentResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeViewContentResponseBody) SetRequestId(v string) *DescribeViewContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeViewContentResponseBody) SetTotalCount(v int32) *DescribeViewContentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeViewContentResponseBody) SetViewContentList(v []*DescribeViewContentResponseBodyViewContentList) *DescribeViewContentResponseBody {
	s.ViewContentList = v
	return s
}

type DescribeViewContentResponseBodyViewContentList struct {
	BizType          *string                                                       `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Content          *string                                                       `json:"Content,omitempty" xml:"Content,omitempty"`
	DataId           *string                                                       `json:"DataId,omitempty" xml:"DataId,omitempty"`
	FaceResults      []*DescribeViewContentResponseBodyViewContentListFaceResults  `json:"FaceResults,omitempty" xml:"FaceResults,omitempty" type:"Repeated"`
	FrameResults     []*DescribeViewContentResponseBodyViewContentListFrameResults `json:"FrameResults,omitempty" xml:"FrameResults,omitempty" type:"Repeated"`
	Id               *int64                                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	NewUrl           *string                                                       `json:"NewUrl,omitempty" xml:"NewUrl,omitempty"`
	RequestTime      *string                                                       `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	Results          []*DescribeViewContentResponseBodyViewContentListResults      `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	ScanFinishedTime *string                                                       `json:"ScanFinishedTime,omitempty" xml:"ScanFinishedTime,omitempty"`
	ScanResult       *string                                                       `json:"ScanResult,omitempty" xml:"ScanResult,omitempty"`
	Suggestion       *string                                                       `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TaskId           *string                                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Thumbnail        *string                                                       `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	Url              *string                                                       `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentList) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentList) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentList) SetBizType(v string) *DescribeViewContentResponseBodyViewContentList {
	s.BizType = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetContent(v string) *DescribeViewContentResponseBodyViewContentList {
	s.Content = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetDataId(v string) *DescribeViewContentResponseBodyViewContentList {
	s.DataId = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetFaceResults(v []*DescribeViewContentResponseBodyViewContentListFaceResults) *DescribeViewContentResponseBodyViewContentList {
	s.FaceResults = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetFrameResults(v []*DescribeViewContentResponseBodyViewContentListFrameResults) *DescribeViewContentResponseBodyViewContentList {
	s.FrameResults = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetId(v int64) *DescribeViewContentResponseBodyViewContentList {
	s.Id = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetNewUrl(v string) *DescribeViewContentResponseBodyViewContentList {
	s.NewUrl = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetRequestTime(v string) *DescribeViewContentResponseBodyViewContentList {
	s.RequestTime = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetResults(v []*DescribeViewContentResponseBodyViewContentListResults) *DescribeViewContentResponseBodyViewContentList {
	s.Results = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetScanFinishedTime(v string) *DescribeViewContentResponseBodyViewContentList {
	s.ScanFinishedTime = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetScanResult(v string) *DescribeViewContentResponseBodyViewContentList {
	s.ScanResult = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetSuggestion(v string) *DescribeViewContentResponseBodyViewContentList {
	s.Suggestion = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetTaskId(v string) *DescribeViewContentResponseBodyViewContentList {
	s.TaskId = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetThumbnail(v string) *DescribeViewContentResponseBodyViewContentList {
	s.Thumbnail = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentList) SetUrl(v string) *DescribeViewContentResponseBodyViewContentList {
	s.Url = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResults struct {
	Age        *DescribeViewContentResponseBodyViewContentListFaceResultsAge        `json:"Age,omitempty" xml:"Age,omitempty" type:"Struct"`
	Bang       *DescribeViewContentResponseBodyViewContentListFaceResultsBang       `json:"Bang,omitempty" xml:"Bang,omitempty" type:"Struct"`
	Bualified  *bool                                                                `json:"Bualified,omitempty" xml:"Bualified,omitempty"`
	Gender     *DescribeViewContentResponseBodyViewContentListFaceResultsGender     `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	Glasses    *DescribeViewContentResponseBodyViewContentListFaceResultsGlasses    `json:"Glasses,omitempty" xml:"Glasses,omitempty" type:"Struct"`
	Hairstyle  *DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle  `json:"Hairstyle,omitempty" xml:"Hairstyle,omitempty" type:"Struct"`
	Hat        *DescribeViewContentResponseBodyViewContentListFaceResultsHat        `json:"Hat,omitempty" xml:"Hat,omitempty" type:"Struct"`
	Image      *DescribeViewContentResponseBodyViewContentListFaceResultsImage      `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	Location   *DescribeViewContentResponseBodyViewContentListFaceResultsLocation   `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	Mustache   *DescribeViewContentResponseBodyViewContentListFaceResultsMustache   `json:"Mustache,omitempty" xml:"Mustache,omitempty" type:"Struct"`
	Quality    *DescribeViewContentResponseBodyViewContentListFaceResultsQuality    `json:"Quality,omitempty" xml:"Quality,omitempty" type:"Struct"`
	Respirator *DescribeViewContentResponseBodyViewContentListFaceResultsRespirator `json:"Respirator,omitempty" xml:"Respirator,omitempty" type:"Struct"`
	Smile      *DescribeViewContentResponseBodyViewContentListFaceResultsSmile      `json:"Smile,omitempty" xml:"Smile,omitempty" type:"Struct"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResults) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetAge(v *DescribeViewContentResponseBodyViewContentListFaceResultsAge) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Age = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetBang(v *DescribeViewContentResponseBodyViewContentListFaceResultsBang) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Bang = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetBualified(v bool) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Bualified = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetGender(v *DescribeViewContentResponseBodyViewContentListFaceResultsGender) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Gender = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetGlasses(v *DescribeViewContentResponseBodyViewContentListFaceResultsGlasses) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Glasses = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetHairstyle(v *DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Hairstyle = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetHat(v *DescribeViewContentResponseBodyViewContentListFaceResultsHat) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Hat = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetImage(v *DescribeViewContentResponseBodyViewContentListFaceResultsImage) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Image = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetLocation(v *DescribeViewContentResponseBodyViewContentListFaceResultsLocation) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Location = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetMustache(v *DescribeViewContentResponseBodyViewContentListFaceResultsMustache) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Mustache = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetQuality(v *DescribeViewContentResponseBodyViewContentListFaceResultsQuality) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Quality = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetRespirator(v *DescribeViewContentResponseBodyViewContentListFaceResultsRespirator) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Respirator = v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResults) SetSmile(v *DescribeViewContentResponseBodyViewContentListFaceResultsSmile) *DescribeViewContentResponseBodyViewContentListFaceResults {
	s.Smile = v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsAge struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *int32   `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsAge) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsAge) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsAge) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsAge {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsAge) SetValue(v int32) *DescribeViewContentResponseBodyViewContentListFaceResultsAge {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsBang struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsBang) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsBang) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsBang) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsBang {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsBang) SetValue(v string) *DescribeViewContentResponseBodyViewContentListFaceResultsBang {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsGender struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsGender) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsGender) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsGender) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsGender {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsGender) SetValue(v string) *DescribeViewContentResponseBodyViewContentListFaceResultsGender {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsGlasses struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsGlasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsGlasses) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsGlasses) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsGlasses {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsGlasses) SetValue(v string) *DescribeViewContentResponseBodyViewContentListFaceResultsGlasses {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle) SetValue(v string) *DescribeViewContentResponseBodyViewContentListFaceResultsHairstyle {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsHat struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsHat) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsHat) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsHat) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsHat {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsHat) SetValue(v string) *DescribeViewContentResponseBodyViewContentListFaceResultsHat {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsImage struct {
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsImage) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsImage) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsImage) SetHeight(v int32) *DescribeViewContentResponseBodyViewContentListFaceResultsImage {
	s.Height = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsImage) SetWidth(v int32) *DescribeViewContentResponseBodyViewContentListFaceResultsImage {
	s.Width = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsLocation struct {
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsLocation) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsLocation) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsLocation) SetH(v int32) *DescribeViewContentResponseBodyViewContentListFaceResultsLocation {
	s.H = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsLocation) SetW(v int32) *DescribeViewContentResponseBodyViewContentListFaceResultsLocation {
	s.W = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsLocation) SetX(v int32) *DescribeViewContentResponseBodyViewContentListFaceResultsLocation {
	s.X = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsLocation) SetY(v int32) *DescribeViewContentResponseBodyViewContentListFaceResultsLocation {
	s.Y = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsMustache struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsMustache) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsMustache) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsMustache) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsMustache {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsMustache) SetValue(v string) *DescribeViewContentResponseBodyViewContentListFaceResultsMustache {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsQuality struct {
	Blur  *float32 `json:"Blur,omitempty" xml:"Blur,omitempty"`
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	Roll  *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	Yaw   *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsQuality) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsQuality) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsQuality) SetBlur(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsQuality {
	s.Blur = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsQuality) SetPitch(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsQuality {
	s.Pitch = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsQuality) SetRoll(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsQuality {
	s.Roll = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsQuality) SetYaw(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsQuality {
	s.Yaw = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsRespirator struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsRespirator) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsRespirator) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsRespirator) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsRespirator {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsRespirator) SetValue(v string) *DescribeViewContentResponseBodyViewContentListFaceResultsRespirator {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFaceResultsSmile struct {
	Rate  *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsSmile) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFaceResultsSmile) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsSmile) SetRate(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsSmile {
	s.Rate = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFaceResultsSmile) SetValue(v float32) *DescribeViewContentResponseBodyViewContentListFaceResultsSmile {
	s.Value = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListFrameResults struct {
	Label  *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Offset *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListFrameResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListFrameResults) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListFrameResults) SetLabel(v string) *DescribeViewContentResponseBodyViewContentListFrameResults {
	s.Label = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFrameResults) SetOffset(v int32) *DescribeViewContentResponseBodyViewContentListFrameResults {
	s.Offset = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListFrameResults) SetUrl(v string) *DescribeViewContentResponseBodyViewContentListFrameResults {
	s.Url = &v
	return s
}

type DescribeViewContentResponseBodyViewContentListResults struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s DescribeViewContentResponseBodyViewContentListResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponseBodyViewContentListResults) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponseBodyViewContentListResults) SetLabel(v string) *DescribeViewContentResponseBodyViewContentListResults {
	s.Label = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListResults) SetScene(v string) *DescribeViewContentResponseBodyViewContentListResults {
	s.Scene = &v
	return s
}

func (s *DescribeViewContentResponseBodyViewContentListResults) SetSuggestion(v string) *DescribeViewContentResponseBodyViewContentListResults {
	s.Suggestion = &v
	return s
}

type DescribeViewContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeViewContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeViewContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeViewContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeViewContentResponse) SetHeaders(v map[string]*string) *DescribeViewContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeViewContentResponse) SetStatusCode(v int32) *DescribeViewContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeViewContentResponse) SetBody(v *DescribeViewContentResponseBody) *DescribeViewContentResponse {
	s.Body = v
	return s
}

type DescribeWebsiteIndexPageBaselineRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebsiteIndexPageBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteIndexPageBaselineRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteIndexPageBaselineRequest) SetInstanceId(v string) *DescribeWebsiteIndexPageBaselineRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWebsiteIndexPageBaselineRequest) SetLang(v string) *DescribeWebsiteIndexPageBaselineRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteIndexPageBaselineRequest) SetSourceIp(v string) *DescribeWebsiteIndexPageBaselineRequest {
	s.SourceIp = &v
	return s
}

type DescribeWebsiteIndexPageBaselineResponseBody struct {
	BaseLineStatus *string `json:"BaseLineStatus,omitempty" xml:"BaseLineStatus,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshot       *string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
}

func (s DescribeWebsiteIndexPageBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteIndexPageBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteIndexPageBaselineResponseBody) SetBaseLineStatus(v string) *DescribeWebsiteIndexPageBaselineResponseBody {
	s.BaseLineStatus = &v
	return s
}

func (s *DescribeWebsiteIndexPageBaselineResponseBody) SetCreateTime(v string) *DescribeWebsiteIndexPageBaselineResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeWebsiteIndexPageBaselineResponseBody) SetRequestId(v string) *DescribeWebsiteIndexPageBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteIndexPageBaselineResponseBody) SetSnapshot(v string) *DescribeWebsiteIndexPageBaselineResponseBody {
	s.Snapshot = &v
	return s
}

type DescribeWebsiteIndexPageBaselineResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteIndexPageBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteIndexPageBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteIndexPageBaselineResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteIndexPageBaselineResponse) SetHeaders(v map[string]*string) *DescribeWebsiteIndexPageBaselineResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteIndexPageBaselineResponse) SetStatusCode(v int32) *DescribeWebsiteIndexPageBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteIndexPageBaselineResponse) SetBody(v *DescribeWebsiteIndexPageBaselineResponseBody) *DescribeWebsiteIndexPageBaselineResponse {
	s.Body = v
	return s
}

type DescribeWebsiteInstanceRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebsiteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceRequest) SetCurrentPage(v int32) *DescribeWebsiteInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebsiteInstanceRequest) SetInstanceId(v string) *DescribeWebsiteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWebsiteInstanceRequest) SetLang(v string) *DescribeWebsiteInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteInstanceRequest) SetPageSize(v int32) *DescribeWebsiteInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebsiteInstanceRequest) SetSourceIp(v string) *DescribeWebsiteInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebsiteInstanceRequest) SetTotalCount(v int32) *DescribeWebsiteInstanceRequest {
	s.TotalCount = &v
	return s
}

type DescribeWebsiteInstanceResponseBody struct {
	CurrentPage         *int32                                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize            *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WebsiteInstanceList []*DescribeWebsiteInstanceResponseBodyWebsiteInstanceList `json:"WebsiteInstanceList,omitempty" xml:"WebsiteInstanceList,omitempty" type:"Repeated"`
}

func (s DescribeWebsiteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceResponseBody) SetCurrentPage(v int32) *DescribeWebsiteInstanceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBody) SetPageSize(v int32) *DescribeWebsiteInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBody) SetRequestId(v string) *DescribeWebsiteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBody) SetTotalCount(v int32) *DescribeWebsiteInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBody) SetWebsiteInstanceList(v []*DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) *DescribeWebsiteInstanceResponseBody {
	s.WebsiteInstanceList = v
	return s
}

type DescribeWebsiteInstanceResponseBodyWebsiteInstanceList struct {
	BuyTime               *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	Domain                *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExpiredTime           *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	IndexPage             *string `json:"IndexPage,omitempty" xml:"IndexPage,omitempty"`
	IndexPageScanInterval *int32  `json:"IndexPageScanInterval,omitempty" xml:"IndexPageScanInterval,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Protocol              *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	WebsiteScanInterval   *int32  `json:"WebsiteScanInterval,omitempty" xml:"WebsiteScanInterval,omitempty"`
}

func (s DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetBuyTime(v string) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.BuyTime = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetDomain(v string) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.Domain = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetExpiredTime(v string) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetIndexPage(v string) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.IndexPage = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetIndexPageScanInterval(v int32) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.IndexPageScanInterval = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetInstanceId(v string) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetProtocol(v string) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.Protocol = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetStatus(v string) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList) SetWebsiteScanInterval(v int32) *DescribeWebsiteInstanceResponseBodyWebsiteInstanceList {
	s.WebsiteScanInterval = &v
	return s
}

type DescribeWebsiteInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceResponse) SetHeaders(v map[string]*string) *DescribeWebsiteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteInstanceResponse) SetStatusCode(v int32) *DescribeWebsiteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteInstanceResponse) SetBody(v *DescribeWebsiteInstanceResponseBody) *DescribeWebsiteInstanceResponse {
	s.Body = v
	return s
}

type DescribeWebsiteInstanceIdRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebsiteInstanceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceIdRequest) SetLang(v string) *DescribeWebsiteInstanceIdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteInstanceIdRequest) SetSourceIp(v string) *DescribeWebsiteInstanceIdRequest {
	s.SourceIp = &v
	return s
}

type DescribeWebsiteInstanceIdResponseBody struct {
	RequestId             *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WebsiteInstanceIdList []*string `json:"WebsiteInstanceIdList,omitempty" xml:"WebsiteInstanceIdList,omitempty" type:"Repeated"`
}

func (s DescribeWebsiteInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceIdResponseBody) SetRequestId(v string) *DescribeWebsiteInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteInstanceIdResponseBody) SetTotalCount(v int32) *DescribeWebsiteInstanceIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebsiteInstanceIdResponseBody) SetWebsiteInstanceIdList(v []*string) *DescribeWebsiteInstanceIdResponseBody {
	s.WebsiteInstanceIdList = v
	return s
}

type DescribeWebsiteInstanceIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceIdResponse) SetHeaders(v map[string]*string) *DescribeWebsiteInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteInstanceIdResponse) SetStatusCode(v int32) *DescribeWebsiteInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteInstanceIdResponse) SetBody(v *DescribeWebsiteInstanceIdResponseBody) *DescribeWebsiteInstanceIdResponse {
	s.Body = v
	return s
}

type DescribeWebsiteInstanceKeyUrlRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebsiteInstanceKeyUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceKeyUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceKeyUrlRequest) SetInstanceId(v string) *DescribeWebsiteInstanceKeyUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWebsiteInstanceKeyUrlRequest) SetLang(v string) *DescribeWebsiteInstanceKeyUrlRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteInstanceKeyUrlRequest) SetSourceIp(v string) *DescribeWebsiteInstanceKeyUrlRequest {
	s.SourceIp = &v
	return s
}

type DescribeWebsiteInstanceKeyUrlResponseBody struct {
	RequestId                 *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WebsiteInstanceKeyUrlList []*string `json:"WebsiteInstanceKeyUrlList,omitempty" xml:"WebsiteInstanceKeyUrlList,omitempty" type:"Repeated"`
}

func (s DescribeWebsiteInstanceKeyUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceKeyUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceKeyUrlResponseBody) SetRequestId(v string) *DescribeWebsiteInstanceKeyUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteInstanceKeyUrlResponseBody) SetTotalCount(v int32) *DescribeWebsiteInstanceKeyUrlResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebsiteInstanceKeyUrlResponseBody) SetWebsiteInstanceKeyUrlList(v []*string) *DescribeWebsiteInstanceKeyUrlResponseBody {
	s.WebsiteInstanceKeyUrlList = v
	return s
}

type DescribeWebsiteInstanceKeyUrlResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteInstanceKeyUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteInstanceKeyUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteInstanceKeyUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteInstanceKeyUrlResponse) SetHeaders(v map[string]*string) *DescribeWebsiteInstanceKeyUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteInstanceKeyUrlResponse) SetStatusCode(v int32) *DescribeWebsiteInstanceKeyUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteInstanceKeyUrlResponse) SetBody(v *DescribeWebsiteInstanceKeyUrlResponseBody) *DescribeWebsiteInstanceKeyUrlResponse {
	s.Body = v
	return s
}

type DescribeWebsiteScanResultRequest struct {
	CurrentPage      *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Domain           *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HandleStatus     *string `json:"HandleStatus,omitempty" xml:"HandleStatus,omitempty"`
	Label            *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SiteUrl          *string `json:"SiteUrl,omitempty" xml:"SiteUrl,omitempty"`
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SubServiceModule *string `json:"SubServiceModule,omitempty" xml:"SubServiceModule,omitempty"`
	TotalCount       *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebsiteScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultRequest) SetCurrentPage(v int32) *DescribeWebsiteScanResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetDomain(v string) *DescribeWebsiteScanResultRequest {
	s.Domain = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetHandleStatus(v string) *DescribeWebsiteScanResultRequest {
	s.HandleStatus = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetLabel(v string) *DescribeWebsiteScanResultRequest {
	s.Label = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetLang(v string) *DescribeWebsiteScanResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetPageSize(v int32) *DescribeWebsiteScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetSiteUrl(v string) *DescribeWebsiteScanResultRequest {
	s.SiteUrl = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetSourceIp(v string) *DescribeWebsiteScanResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetSubServiceModule(v string) *DescribeWebsiteScanResultRequest {
	s.SubServiceModule = &v
	return s
}

func (s *DescribeWebsiteScanResultRequest) SetTotalCount(v int32) *DescribeWebsiteScanResultRequest {
	s.TotalCount = &v
	return s
}

type DescribeWebsiteScanResultResponseBody struct {
	CurrentPage           *int32                                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize              *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount            *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WebsiteScanResultList []*DescribeWebsiteScanResultResponseBodyWebsiteScanResultList `json:"WebsiteScanResultList,omitempty" xml:"WebsiteScanResultList,omitempty" type:"Repeated"`
}

func (s DescribeWebsiteScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultResponseBody) SetCurrentPage(v int32) *DescribeWebsiteScanResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBody) SetPageSize(v int32) *DescribeWebsiteScanResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBody) SetRequestId(v string) *DescribeWebsiteScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBody) SetTotalCount(v int32) *DescribeWebsiteScanResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBody) SetWebsiteScanResultList(v []*DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) *DescribeWebsiteScanResultResponseBody {
	s.WebsiteScanResultList = v
	return s
}

type DescribeWebsiteScanResultResponseBodyWebsiteScanResultList struct {
	Domain          *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HandleStatus    *int32    `json:"HandleStatus,omitempty" xml:"HandleStatus,omitempty"`
	Id              *int32    `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageRiskCount  *int32    `json:"ImageRiskCount,omitempty" xml:"ImageRiskCount,omitempty"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Labels          []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ScanTime        *string   `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	SourceRiskCount *int32    `json:"SourceRiskCount,omitempty" xml:"SourceRiskCount,omitempty"`
	TaskId          *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TextRiskCount   *int32    `json:"TextRiskCount,omitempty" xml:"TextRiskCount,omitempty"`
	Url             *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetDomain(v string) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.Domain = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetHandleStatus(v int32) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.HandleStatus = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetId(v int32) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.Id = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetImageRiskCount(v int32) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.ImageRiskCount = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetInstanceId(v string) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.InstanceId = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetLabels(v []*string) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.Labels = v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetScanTime(v string) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.ScanTime = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetSourceRiskCount(v int32) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.SourceRiskCount = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetTaskId(v string) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.TaskId = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetTextRiskCount(v int32) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.TextRiskCount = &v
	return s
}

func (s *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList) SetUrl(v string) *DescribeWebsiteScanResultResponseBodyWebsiteScanResultList {
	s.Url = &v
	return s
}

type DescribeWebsiteScanResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultResponse) SetHeaders(v map[string]*string) *DescribeWebsiteScanResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteScanResultResponse) SetStatusCode(v int32) *DescribeWebsiteScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteScanResultResponse) SetBody(v *DescribeWebsiteScanResultResponseBody) *DescribeWebsiteScanResultResponse {
	s.Body = v
	return s
}

type DescribeWebsiteScanResultDetailRequest struct {
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebsiteScanResultDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultDetailRequest) SetId(v int32) *DescribeWebsiteScanResultDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailRequest) SetLang(v string) *DescribeWebsiteScanResultDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailRequest) SetResourceType(v string) *DescribeWebsiteScanResultDetailRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailRequest) SetSourceIp(v string) *DescribeWebsiteScanResultDetailRequest {
	s.SourceIp = &v
	return s
}

type DescribeWebsiteScanResultDetailResponseBody struct {
	Baseline         *string                                                        `json:"Baseline,omitempty" xml:"Baseline,omitempty"`
	Content          *string                                                        `json:"Content,omitempty" xml:"Content,omitempty"`
	HitKeywords      []*string                                                      `json:"HitKeywords,omitempty" xml:"HitKeywords,omitempty" type:"Repeated"`
	ImageScanResults []*DescribeWebsiteScanResultDetailResponseBodyImageScanResults `json:"ImageScanResults,omitempty" xml:"ImageScanResults,omitempty" type:"Repeated"`
	RequestId        *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceType     *string                                                        `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TamperedSource   *string                                                        `json:"TamperedSource,omitempty" xml:"TamperedSource,omitempty"`
}

func (s DescribeWebsiteScanResultDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultDetailResponseBody) SetBaseline(v string) *DescribeWebsiteScanResultDetailResponseBody {
	s.Baseline = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponseBody) SetContent(v string) *DescribeWebsiteScanResultDetailResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponseBody) SetHitKeywords(v []*string) *DescribeWebsiteScanResultDetailResponseBody {
	s.HitKeywords = v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponseBody) SetImageScanResults(v []*DescribeWebsiteScanResultDetailResponseBodyImageScanResults) *DescribeWebsiteScanResultDetailResponseBody {
	s.ImageScanResults = v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponseBody) SetRequestId(v string) *DescribeWebsiteScanResultDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponseBody) SetResourceType(v string) *DescribeWebsiteScanResultDetailResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponseBody) SetTamperedSource(v string) *DescribeWebsiteScanResultDetailResponseBody {
	s.TamperedSource = &v
	return s
}

type DescribeWebsiteScanResultDetailResponseBodyImageScanResults struct {
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Url    *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeWebsiteScanResultDetailResponseBodyImageScanResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultDetailResponseBodyImageScanResults) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultDetailResponseBodyImageScanResults) SetLabels(v []*string) *DescribeWebsiteScanResultDetailResponseBodyImageScanResults {
	s.Labels = v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponseBodyImageScanResults) SetUrl(v string) *DescribeWebsiteScanResultDetailResponseBodyImageScanResults {
	s.Url = &v
	return s
}

type DescribeWebsiteScanResultDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteScanResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteScanResultDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteScanResultDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteScanResultDetailResponse) SetHeaders(v map[string]*string) *DescribeWebsiteScanResultDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponse) SetStatusCode(v int32) *DescribeWebsiteScanResultDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteScanResultDetailResponse) SetBody(v *DescribeWebsiteScanResultDetailResponseBody) *DescribeWebsiteScanResultDetailResponse {
	s.Body = v
	return s
}

type DescribeWebsiteStatRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebsiteStatRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteStatRequest) SetLang(v string) *DescribeWebsiteStatRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteStatRequest) SetSourceIp(v string) *DescribeWebsiteStatRequest {
	s.SourceIp = &v
	return s
}

type DescribeWebsiteStatResponseBody struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WebsiteStatList []*DescribeWebsiteStatResponseBodyWebsiteStatList `json:"WebsiteStatList,omitempty" xml:"WebsiteStatList,omitempty" type:"Repeated"`
}

func (s DescribeWebsiteStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteStatResponseBody) SetRequestId(v string) *DescribeWebsiteStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteStatResponseBody) SetTotalCount(v int32) *DescribeWebsiteStatResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebsiteStatResponseBody) SetWebsiteStatList(v []*DescribeWebsiteStatResponseBodyWebsiteStatList) *DescribeWebsiteStatResponseBody {
	s.WebsiteStatList = v
	return s
}

type DescribeWebsiteStatResponseBodyWebsiteStatList struct {
	InstanceCount    *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	RiskCount        *int32  `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	ScanCount        *int32  `json:"ScanCount,omitempty" xml:"ScanCount,omitempty"`
	SubServiceModule *string `json:"SubServiceModule,omitempty" xml:"SubServiceModule,omitempty"`
}

func (s DescribeWebsiteStatResponseBodyWebsiteStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteStatResponseBodyWebsiteStatList) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteStatResponseBodyWebsiteStatList) SetInstanceCount(v int32) *DescribeWebsiteStatResponseBodyWebsiteStatList {
	s.InstanceCount = &v
	return s
}

func (s *DescribeWebsiteStatResponseBodyWebsiteStatList) SetRiskCount(v int32) *DescribeWebsiteStatResponseBodyWebsiteStatList {
	s.RiskCount = &v
	return s
}

func (s *DescribeWebsiteStatResponseBodyWebsiteStatList) SetScanCount(v int32) *DescribeWebsiteStatResponseBodyWebsiteStatList {
	s.ScanCount = &v
	return s
}

func (s *DescribeWebsiteStatResponseBodyWebsiteStatList) SetSubServiceModule(v string) *DescribeWebsiteStatResponseBodyWebsiteStatList {
	s.SubServiceModule = &v
	return s
}

type DescribeWebsiteStatResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteStatResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteStatResponse) SetHeaders(v map[string]*string) *DescribeWebsiteStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteStatResponse) SetStatusCode(v int32) *DescribeWebsiteStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteStatResponse) SetBody(v *DescribeWebsiteStatResponseBody) *DescribeWebsiteStatResponse {
	s.Body = v
	return s
}

type DescribeWebsiteVerifyInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebsiteVerifyInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteVerifyInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteVerifyInfoRequest) SetInstanceId(v string) *DescribeWebsiteVerifyInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoRequest) SetLang(v string) *DescribeWebsiteVerifyInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoRequest) SetSourceIp(v string) *DescribeWebsiteVerifyInfoRequest {
	s.SourceIp = &v
	return s
}

type DescribeWebsiteVerifyInfoResponseBody struct {
	Cname        *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	HostFile     *string `json:"HostFile,omitempty" xml:"HostFile,omitempty"`
	IndexPage    *string `json:"IndexPage,omitempty" xml:"IndexPage,omitempty"`
	Protocol     *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VerifyMethod *string `json:"VerifyMethod,omitempty" xml:"VerifyMethod,omitempty"`
}

func (s DescribeWebsiteVerifyInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteVerifyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteVerifyInfoResponseBody) SetCname(v string) *DescribeWebsiteVerifyInfoResponseBody {
	s.Cname = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponseBody) SetDomain(v string) *DescribeWebsiteVerifyInfoResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponseBody) SetHostFile(v string) *DescribeWebsiteVerifyInfoResponseBody {
	s.HostFile = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponseBody) SetIndexPage(v string) *DescribeWebsiteVerifyInfoResponseBody {
	s.IndexPage = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponseBody) SetProtocol(v string) *DescribeWebsiteVerifyInfoResponseBody {
	s.Protocol = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponseBody) SetRequestId(v string) *DescribeWebsiteVerifyInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponseBody) SetVerifyMethod(v string) *DescribeWebsiteVerifyInfoResponseBody {
	s.VerifyMethod = &v
	return s
}

type DescribeWebsiteVerifyInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebsiteVerifyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebsiteVerifyInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebsiteVerifyInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebsiteVerifyInfoResponse) SetHeaders(v map[string]*string) *DescribeWebsiteVerifyInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponse) SetStatusCode(v int32) *DescribeWebsiteVerifyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebsiteVerifyInfoResponse) SetBody(v *DescribeWebsiteVerifyInfoResponseBody) *DescribeWebsiteVerifyInfoResponse {
	s.Body = v
	return s
}

type ExportKeywordsRequest struct {
	KeywordLibId *int64 `json:"KeywordLibId,omitempty" xml:"KeywordLibId,omitempty"`
}

func (s ExportKeywordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportKeywordsRequest) GoString() string {
	return s.String()
}

func (s *ExportKeywordsRequest) SetKeywordLibId(v int64) *ExportKeywordsRequest {
	s.KeywordLibId = &v
	return s
}

type ExportKeywordsResponseBody struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportKeywordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportKeywordsResponseBody) SetDownloadUrl(v string) *ExportKeywordsResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *ExportKeywordsResponseBody) SetRequestId(v string) *ExportKeywordsResponseBody {
	s.RequestId = &v
	return s
}

type ExportKeywordsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportKeywordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportKeywordsResponse) GoString() string {
	return s.String()
}

func (s *ExportKeywordsResponse) SetHeaders(v map[string]*string) *ExportKeywordsResponse {
	s.Headers = v
	return s
}

func (s *ExportKeywordsResponse) SetStatusCode(v int32) *ExportKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportKeywordsResponse) SetBody(v *ExportKeywordsResponseBody) *ExportKeywordsResponse {
	s.Body = v
	return s
}

type ExportOpenApiRcpStatsRequest struct {
	BizType      *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ExportOpenApiRcpStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportOpenApiRcpStatsRequest) GoString() string {
	return s.String()
}

func (s *ExportOpenApiRcpStatsRequest) SetBizType(v string) *ExportOpenApiRcpStatsRequest {
	s.BizType = &v
	return s
}

func (s *ExportOpenApiRcpStatsRequest) SetEndDate(v string) *ExportOpenApiRcpStatsRequest {
	s.EndDate = &v
	return s
}

func (s *ExportOpenApiRcpStatsRequest) SetResourceType(v string) *ExportOpenApiRcpStatsRequest {
	s.ResourceType = &v
	return s
}

func (s *ExportOpenApiRcpStatsRequest) SetStartDate(v string) *ExportOpenApiRcpStatsRequest {
	s.StartDate = &v
	return s
}

type ExportOpenApiRcpStatsResponseBody struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportOpenApiRcpStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportOpenApiRcpStatsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportOpenApiRcpStatsResponseBody) SetDownloadUrl(v string) *ExportOpenApiRcpStatsResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *ExportOpenApiRcpStatsResponseBody) SetRequestId(v string) *ExportOpenApiRcpStatsResponseBody {
	s.RequestId = &v
	return s
}

type ExportOpenApiRcpStatsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportOpenApiRcpStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportOpenApiRcpStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportOpenApiRcpStatsResponse) GoString() string {
	return s.String()
}

func (s *ExportOpenApiRcpStatsResponse) SetHeaders(v map[string]*string) *ExportOpenApiRcpStatsResponse {
	s.Headers = v
	return s
}

func (s *ExportOpenApiRcpStatsResponse) SetStatusCode(v int32) *ExportOpenApiRcpStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportOpenApiRcpStatsResponse) SetBody(v *ExportOpenApiRcpStatsResponseBody) *ExportOpenApiRcpStatsResponse {
	s.Body = v
	return s
}

type ExportOssResultRequest struct {
	Bucket       *string  `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	CurrentPage  *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndDate      *string  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Lang         *string  `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MaxScore     *float32 `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	MinScore     *float32 `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
	PageSize     *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType *string  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string  `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceIp     *string  `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate    *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Stock        *bool    `json:"Stock,omitempty" xml:"Stock,omitempty"`
	StockTaskId  *int64   `json:"StockTaskId,omitempty" xml:"StockTaskId,omitempty"`
	Suggestion   *string  `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TotalCount   *int32   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ExportOssResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportOssResultRequest) GoString() string {
	return s.String()
}

func (s *ExportOssResultRequest) SetBucket(v string) *ExportOssResultRequest {
	s.Bucket = &v
	return s
}

func (s *ExportOssResultRequest) SetCurrentPage(v int32) *ExportOssResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *ExportOssResultRequest) SetEndDate(v string) *ExportOssResultRequest {
	s.EndDate = &v
	return s
}

func (s *ExportOssResultRequest) SetLang(v string) *ExportOssResultRequest {
	s.Lang = &v
	return s
}

func (s *ExportOssResultRequest) SetMaxScore(v float32) *ExportOssResultRequest {
	s.MaxScore = &v
	return s
}

func (s *ExportOssResultRequest) SetMinScore(v float32) *ExportOssResultRequest {
	s.MinScore = &v
	return s
}

func (s *ExportOssResultRequest) SetPageSize(v int32) *ExportOssResultRequest {
	s.PageSize = &v
	return s
}

func (s *ExportOssResultRequest) SetResourceType(v string) *ExportOssResultRequest {
	s.ResourceType = &v
	return s
}

func (s *ExportOssResultRequest) SetScene(v string) *ExportOssResultRequest {
	s.Scene = &v
	return s
}

func (s *ExportOssResultRequest) SetSourceIp(v string) *ExportOssResultRequest {
	s.SourceIp = &v
	return s
}

func (s *ExportOssResultRequest) SetStartDate(v string) *ExportOssResultRequest {
	s.StartDate = &v
	return s
}

func (s *ExportOssResultRequest) SetStock(v bool) *ExportOssResultRequest {
	s.Stock = &v
	return s
}

func (s *ExportOssResultRequest) SetStockTaskId(v int64) *ExportOssResultRequest {
	s.StockTaskId = &v
	return s
}

func (s *ExportOssResultRequest) SetSuggestion(v string) *ExportOssResultRequest {
	s.Suggestion = &v
	return s
}

func (s *ExportOssResultRequest) SetTotalCount(v int32) *ExportOssResultRequest {
	s.TotalCount = &v
	return s
}

type ExportOssResultResponseBody struct {
	FileUrl   *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportOssResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportOssResultResponseBody) GoString() string {
	return s.String()
}

func (s *ExportOssResultResponseBody) SetFileUrl(v string) *ExportOssResultResponseBody {
	s.FileUrl = &v
	return s
}

func (s *ExportOssResultResponseBody) SetRequestId(v string) *ExportOssResultResponseBody {
	s.RequestId = &v
	return s
}

type ExportOssResultResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportOssResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportOssResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportOssResultResponse) GoString() string {
	return s.String()
}

func (s *ExportOssResultResponse) SetHeaders(v map[string]*string) *ExportOssResultResponse {
	s.Headers = v
	return s
}

func (s *ExportOssResultResponse) SetStatusCode(v int32) *ExportOssResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportOssResultResponse) SetBody(v *ExportOssResultResponseBody) *ExportOssResultResponse {
	s.Body = v
	return s
}

type GetAuditItemDetailRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAuditItemDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuditItemDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAuditItemDetailRequest) SetTaskId(v string) *GetAuditItemDetailRequest {
	s.TaskId = &v
	return s
}

func (s *GetAuditItemDetailRequest) SetType(v string) *GetAuditItemDetailRequest {
	s.Type = &v
	return s
}

type GetAuditItemDetailResponseBody struct {
	ApiResult   *string `json:"ApiResult,omitempty" xml:"ApiResult,omitempty"`
	ApiRiskType *string `json:"ApiRiskType,omitempty" xml:"ApiRiskType,omitempty"`
	ApiTs       *string `json:"ApiTs,omitempty" xml:"ApiTs,omitempty"`
	NewUrl      *string `json:"NewUrl,omitempty" xml:"NewUrl,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuditItemDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuditItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditItemDetailResponseBody) SetApiResult(v string) *GetAuditItemDetailResponseBody {
	s.ApiResult = &v
	return s
}

func (s *GetAuditItemDetailResponseBody) SetApiRiskType(v string) *GetAuditItemDetailResponseBody {
	s.ApiRiskType = &v
	return s
}

func (s *GetAuditItemDetailResponseBody) SetApiTs(v string) *GetAuditItemDetailResponseBody {
	s.ApiTs = &v
	return s
}

func (s *GetAuditItemDetailResponseBody) SetNewUrl(v string) *GetAuditItemDetailResponseBody {
	s.NewUrl = &v
	return s
}

func (s *GetAuditItemDetailResponseBody) SetRequestId(v string) *GetAuditItemDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetAuditItemDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditItemDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuditItemDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAuditItemDetailResponse) SetHeaders(v map[string]*string) *GetAuditItemDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAuditItemDetailResponse) SetStatusCode(v int32) *GetAuditItemDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditItemDetailResponse) SetBody(v *GetAuditItemDetailResponseBody) *GetAuditItemDetailResponse {
	s.Body = v
	return s
}

type GetAuditItemListRequest struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CurrentPage    *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	CustomResult   *string `json:"CustomResult,omitempty" xml:"CustomResult,omitempty"`
	CustomRiskType *string `json:"CustomRiskType,omitempty" xml:"CustomRiskType,omitempty"`
	DataId         *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	EndDate        *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PageSize       *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RcpResult      *string `json:"RcpResult,omitempty" xml:"RcpResult,omitempty"`
	RcpRiskType    *string `json:"RcpRiskType,omitempty" xml:"RcpRiskType,omitempty"`
	StartDate      *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAuditItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuditItemListRequest) GoString() string {
	return s.String()
}

func (s *GetAuditItemListRequest) SetBizType(v string) *GetAuditItemListRequest {
	s.BizType = &v
	return s
}

func (s *GetAuditItemListRequest) SetCurrentPage(v int64) *GetAuditItemListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAuditItemListRequest) SetCustomResult(v string) *GetAuditItemListRequest {
	s.CustomResult = &v
	return s
}

func (s *GetAuditItemListRequest) SetCustomRiskType(v string) *GetAuditItemListRequest {
	s.CustomRiskType = &v
	return s
}

func (s *GetAuditItemListRequest) SetDataId(v string) *GetAuditItemListRequest {
	s.DataId = &v
	return s
}

func (s *GetAuditItemListRequest) SetEndDate(v string) *GetAuditItemListRequest {
	s.EndDate = &v
	return s
}

func (s *GetAuditItemListRequest) SetPageSize(v int64) *GetAuditItemListRequest {
	s.PageSize = &v
	return s
}

func (s *GetAuditItemListRequest) SetRcpResult(v string) *GetAuditItemListRequest {
	s.RcpResult = &v
	return s
}

func (s *GetAuditItemListRequest) SetRcpRiskType(v string) *GetAuditItemListRequest {
	s.RcpRiskType = &v
	return s
}

func (s *GetAuditItemListRequest) SetStartDate(v string) *GetAuditItemListRequest {
	s.StartDate = &v
	return s
}

func (s *GetAuditItemListRequest) SetTaskId(v string) *GetAuditItemListRequest {
	s.TaskId = &v
	return s
}

func (s *GetAuditItemListRequest) SetType(v string) *GetAuditItemListRequest {
	s.Type = &v
	return s
}

type GetAuditItemListResponseBody struct {
	CurrentPage *int64                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*GetAuditItemListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageSize    *int64                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAuditItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuditItemListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditItemListResponseBody) SetCurrentPage(v int64) *GetAuditItemListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetAuditItemListResponseBody) SetItems(v []*GetAuditItemListResponseBodyItems) *GetAuditItemListResponseBody {
	s.Items = v
	return s
}

func (s *GetAuditItemListResponseBody) SetPageSize(v int64) *GetAuditItemListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAuditItemListResponseBody) SetRequestId(v string) *GetAuditItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditItemListResponseBody) SetTotalCount(v int64) *GetAuditItemListResponseBody {
	s.TotalCount = &v
	return s
}

type GetAuditItemListResponseBodyItems struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Create         *string `json:"Create,omitempty" xml:"Create,omitempty"`
	CustomResult   *string `json:"CustomResult,omitempty" xml:"CustomResult,omitempty"`
	CustomRiskType *string `json:"CustomRiskType,omitempty" xml:"CustomRiskType,omitempty"`
	CustomTs       *string `json:"CustomTs,omitempty" xml:"CustomTs,omitempty"`
	DataId         *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Operator       *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RcpResult      *string `json:"RcpResult,omitempty" xml:"RcpResult,omitempty"`
	RcpRiskType    *string `json:"RcpRiskType,omitempty" xml:"RcpRiskType,omitempty"`
	RcpTs          *string `json:"RcpTs,omitempty" xml:"RcpTs,omitempty"`
	SubUid         *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Thumbnail      *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid            *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAuditItemListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s GetAuditItemListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetAuditItemListResponseBodyItems) SetBizType(v string) *GetAuditItemListResponseBodyItems {
	s.BizType = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetContent(v string) *GetAuditItemListResponseBodyItems {
	s.Content = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetCreate(v string) *GetAuditItemListResponseBodyItems {
	s.Create = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetCustomResult(v string) *GetAuditItemListResponseBodyItems {
	s.CustomResult = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetCustomRiskType(v string) *GetAuditItemListResponseBodyItems {
	s.CustomRiskType = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetCustomTs(v string) *GetAuditItemListResponseBodyItems {
	s.CustomTs = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetDataId(v string) *GetAuditItemListResponseBodyItems {
	s.DataId = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetId(v int64) *GetAuditItemListResponseBodyItems {
	s.Id = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetOperator(v string) *GetAuditItemListResponseBodyItems {
	s.Operator = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetRcpResult(v string) *GetAuditItemListResponseBodyItems {
	s.RcpResult = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetRcpRiskType(v string) *GetAuditItemListResponseBodyItems {
	s.RcpRiskType = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetRcpTs(v string) *GetAuditItemListResponseBodyItems {
	s.RcpTs = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetSubUid(v string) *GetAuditItemListResponseBodyItems {
	s.SubUid = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetTaskId(v string) *GetAuditItemListResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetThumbnail(v string) *GetAuditItemListResponseBodyItems {
	s.Thumbnail = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetType(v string) *GetAuditItemListResponseBodyItems {
	s.Type = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetUid(v string) *GetAuditItemListResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *GetAuditItemListResponseBodyItems) SetUrl(v string) *GetAuditItemListResponseBodyItems {
	s.Url = &v
	return s
}

type GetAuditItemListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuditItemListResponse) GoString() string {
	return s.String()
}

func (s *GetAuditItemListResponse) SetHeaders(v map[string]*string) *GetAuditItemListResponse {
	s.Headers = v
	return s
}

func (s *GetAuditItemListResponse) SetStatusCode(v int32) *GetAuditItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditItemListResponse) SetBody(v *GetAuditItemListResponseBody) *GetAuditItemListResponse {
	s.Body = v
	return s
}

type GetAuditUserConfResponseBody struct {
	CustomAudit *bool                `json:"CustomAudit,omitempty" xml:"CustomAudit,omitempty"`
	RcpLabels   map[string][]*string `json:"RcpLabels,omitempty" xml:"RcpLabels,omitempty"`
	RequestId   *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserLabels  map[string][]*string `json:"UserLabels,omitempty" xml:"UserLabels,omitempty"`
}

func (s GetAuditUserConfResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuditUserConfResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditUserConfResponseBody) SetCustomAudit(v bool) *GetAuditUserConfResponseBody {
	s.CustomAudit = &v
	return s
}

func (s *GetAuditUserConfResponseBody) SetRcpLabels(v map[string][]*string) *GetAuditUserConfResponseBody {
	s.RcpLabels = v
	return s
}

func (s *GetAuditUserConfResponseBody) SetRequestId(v string) *GetAuditUserConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditUserConfResponseBody) SetUserLabels(v map[string][]*string) *GetAuditUserConfResponseBody {
	s.UserLabels = v
	return s
}

type GetAuditUserConfResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditUserConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditUserConfResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuditUserConfResponse) GoString() string {
	return s.String()
}

func (s *GetAuditUserConfResponse) SetHeaders(v map[string]*string) *GetAuditUserConfResponse {
	s.Headers = v
	return s
}

func (s *GetAuditUserConfResponse) SetStatusCode(v int32) *GetAuditUserConfResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditUserConfResponse) SetBody(v *GetAuditUserConfResponseBody) *GetAuditUserConfResponse {
	s.Body = v
	return s
}

type ImportKeywordsRequest struct {
	KeywordLibId   *int32  `json:"KeywordLibId,omitempty" xml:"KeywordLibId,omitempty"`
	KeywordsObject *string `json:"KeywordsObject,omitempty" xml:"KeywordsObject,omitempty"`
}

func (s ImportKeywordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeywordsRequest) GoString() string {
	return s.String()
}

func (s *ImportKeywordsRequest) SetKeywordLibId(v int32) *ImportKeywordsRequest {
	s.KeywordLibId = &v
	return s
}

func (s *ImportKeywordsRequest) SetKeywordsObject(v string) *ImportKeywordsRequest {
	s.KeywordsObject = &v
	return s
}

type ImportKeywordsResponseBody struct {
	InvalidKeywordList []*string `json:"InvalidKeywordList,omitempty" xml:"InvalidKeywordList,omitempty" type:"Repeated"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessCount       *int32    `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	ValidKeywordList   []*string `json:"validKeywordList,omitempty" xml:"validKeywordList,omitempty" type:"Repeated"`
}

func (s ImportKeywordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeywordsResponseBody) SetInvalidKeywordList(v []*string) *ImportKeywordsResponseBody {
	s.InvalidKeywordList = v
	return s
}

func (s *ImportKeywordsResponseBody) SetRequestId(v string) *ImportKeywordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportKeywordsResponseBody) SetSuccessCount(v int32) *ImportKeywordsResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *ImportKeywordsResponseBody) SetValidKeywordList(v []*string) *ImportKeywordsResponseBody {
	s.ValidKeywordList = v
	return s
}

type ImportKeywordsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKeywordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeywordsResponse) GoString() string {
	return s.String()
}

func (s *ImportKeywordsResponse) SetHeaders(v map[string]*string) *ImportKeywordsResponse {
	s.Headers = v
	return s
}

func (s *ImportKeywordsResponse) SetStatusCode(v int32) *ImportKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKeywordsResponse) SetBody(v *ImportKeywordsResponseBody) *ImportKeywordsResponse {
	s.Body = v
	return s
}

type MarkAuditContentRequest struct {
	AuditIllegalReasons *string `json:"AuditIllegalReasons,omitempty" xml:"AuditIllegalReasons,omitempty"`
	AuditResult         *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	BizTypes            *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	Ids                 *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s MarkAuditContentRequest) String() string {
	return tea.Prettify(s)
}

func (s MarkAuditContentRequest) GoString() string {
	return s.String()
}

func (s *MarkAuditContentRequest) SetAuditIllegalReasons(v string) *MarkAuditContentRequest {
	s.AuditIllegalReasons = &v
	return s
}

func (s *MarkAuditContentRequest) SetAuditResult(v string) *MarkAuditContentRequest {
	s.AuditResult = &v
	return s
}

func (s *MarkAuditContentRequest) SetBizTypes(v string) *MarkAuditContentRequest {
	s.BizTypes = &v
	return s
}

func (s *MarkAuditContentRequest) SetIds(v string) *MarkAuditContentRequest {
	s.Ids = &v
	return s
}

func (s *MarkAuditContentRequest) SetSourceIp(v string) *MarkAuditContentRequest {
	s.SourceIp = &v
	return s
}

type MarkAuditContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MarkAuditContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MarkAuditContentResponseBody) GoString() string {
	return s.String()
}

func (s *MarkAuditContentResponseBody) SetRequestId(v string) *MarkAuditContentResponseBody {
	s.RequestId = &v
	return s
}

type MarkAuditContentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkAuditContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkAuditContentResponse) String() string {
	return tea.Prettify(s)
}

func (s MarkAuditContentResponse) GoString() string {
	return s.String()
}

func (s *MarkAuditContentResponse) SetHeaders(v map[string]*string) *MarkAuditContentResponse {
	s.Headers = v
	return s
}

func (s *MarkAuditContentResponse) SetStatusCode(v int32) *MarkAuditContentResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkAuditContentResponse) SetBody(v *MarkAuditContentResponseBody) *MarkAuditContentResponse {
	s.Body = v
	return s
}

type MarkAuditContentItemRequest struct {
	AuditIllegalReasons *string `json:"AuditIllegalReasons,omitempty" xml:"AuditIllegalReasons,omitempty"`
	AuditResult         *string `json:"AuditResult,omitempty" xml:"AuditResult,omitempty"`
	BizTypes            *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	Ids                 *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s MarkAuditContentItemRequest) String() string {
	return tea.Prettify(s)
}

func (s MarkAuditContentItemRequest) GoString() string {
	return s.String()
}

func (s *MarkAuditContentItemRequest) SetAuditIllegalReasons(v string) *MarkAuditContentItemRequest {
	s.AuditIllegalReasons = &v
	return s
}

func (s *MarkAuditContentItemRequest) SetAuditResult(v string) *MarkAuditContentItemRequest {
	s.AuditResult = &v
	return s
}

func (s *MarkAuditContentItemRequest) SetBizTypes(v string) *MarkAuditContentItemRequest {
	s.BizTypes = &v
	return s
}

func (s *MarkAuditContentItemRequest) SetIds(v string) *MarkAuditContentItemRequest {
	s.Ids = &v
	return s
}

func (s *MarkAuditContentItemRequest) SetLang(v string) *MarkAuditContentItemRequest {
	s.Lang = &v
	return s
}

func (s *MarkAuditContentItemRequest) SetSourceIp(v string) *MarkAuditContentItemRequest {
	s.SourceIp = &v
	return s
}

type MarkAuditContentItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MarkAuditContentItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MarkAuditContentItemResponseBody) GoString() string {
	return s.String()
}

func (s *MarkAuditContentItemResponseBody) SetRequestId(v string) *MarkAuditContentItemResponseBody {
	s.RequestId = &v
	return s
}

type MarkAuditContentItemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkAuditContentItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkAuditContentItemResponse) String() string {
	return tea.Prettify(s)
}

func (s MarkAuditContentItemResponse) GoString() string {
	return s.String()
}

func (s *MarkAuditContentItemResponse) SetHeaders(v map[string]*string) *MarkAuditContentItemResponse {
	s.Headers = v
	return s
}

func (s *MarkAuditContentItemResponse) SetStatusCode(v int32) *MarkAuditContentItemResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkAuditContentItemResponse) SetBody(v *MarkAuditContentItemResponseBody) *MarkAuditContentItemResponse {
	s.Body = v
	return s
}

type MarkOssResultRequest struct {
	Ids          *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Operation    *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Stock        *bool   `json:"Stock,omitempty" xml:"Stock,omitempty"`
}

func (s MarkOssResultRequest) String() string {
	return tea.Prettify(s)
}

func (s MarkOssResultRequest) GoString() string {
	return s.String()
}

func (s *MarkOssResultRequest) SetIds(v string) *MarkOssResultRequest {
	s.Ids = &v
	return s
}

func (s *MarkOssResultRequest) SetLang(v string) *MarkOssResultRequest {
	s.Lang = &v
	return s
}

func (s *MarkOssResultRequest) SetOperation(v string) *MarkOssResultRequest {
	s.Operation = &v
	return s
}

func (s *MarkOssResultRequest) SetResourceType(v string) *MarkOssResultRequest {
	s.ResourceType = &v
	return s
}

func (s *MarkOssResultRequest) SetScene(v string) *MarkOssResultRequest {
	s.Scene = &v
	return s
}

func (s *MarkOssResultRequest) SetSourceIp(v string) *MarkOssResultRequest {
	s.SourceIp = &v
	return s
}

func (s *MarkOssResultRequest) SetStock(v bool) *MarkOssResultRequest {
	s.Stock = &v
	return s
}

type MarkOssResultResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MarkOssResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MarkOssResultResponseBody) GoString() string {
	return s.String()
}

func (s *MarkOssResultResponseBody) SetRequestId(v string) *MarkOssResultResponseBody {
	s.RequestId = &v
	return s
}

type MarkOssResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkOssResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkOssResultResponse) String() string {
	return tea.Prettify(s)
}

func (s MarkOssResultResponse) GoString() string {
	return s.String()
}

func (s *MarkOssResultResponse) SetHeaders(v map[string]*string) *MarkOssResultResponse {
	s.Headers = v
	return s
}

func (s *MarkOssResultResponse) SetStatusCode(v int32) *MarkOssResultResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkOssResultResponse) SetBody(v *MarkOssResultResponseBody) *MarkOssResultResponse {
	s.Body = v
	return s
}

type MarkWebsiteScanResultRequest struct {
	Ids      *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s MarkWebsiteScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s MarkWebsiteScanResultRequest) GoString() string {
	return s.String()
}

func (s *MarkWebsiteScanResultRequest) SetIds(v string) *MarkWebsiteScanResultRequest {
	s.Ids = &v
	return s
}

func (s *MarkWebsiteScanResultRequest) SetLang(v string) *MarkWebsiteScanResultRequest {
	s.Lang = &v
	return s
}

func (s *MarkWebsiteScanResultRequest) SetSourceIp(v string) *MarkWebsiteScanResultRequest {
	s.SourceIp = &v
	return s
}

type MarkWebsiteScanResultResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MarkWebsiteScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MarkWebsiteScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *MarkWebsiteScanResultResponseBody) SetRequestId(v string) *MarkWebsiteScanResultResponseBody {
	s.RequestId = &v
	return s
}

type MarkWebsiteScanResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MarkWebsiteScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MarkWebsiteScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s MarkWebsiteScanResultResponse) GoString() string {
	return s.String()
}

func (s *MarkWebsiteScanResultResponse) SetHeaders(v map[string]*string) *MarkWebsiteScanResultResponse {
	s.Headers = v
	return s
}

func (s *MarkWebsiteScanResultResponse) SetStatusCode(v int32) *MarkWebsiteScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *MarkWebsiteScanResultResponse) SetBody(v *MarkWebsiteScanResultResponseBody) *MarkWebsiteScanResultResponse {
	s.Body = v
	return s
}

type RefundCdiBagRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RefundCdiBagRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundCdiBagRequest) GoString() string {
	return s.String()
}

func (s *RefundCdiBagRequest) SetInstanceId(v string) *RefundCdiBagRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundCdiBagRequest) SetResourceOwnerId(v string) *RefundCdiBagRequest {
	s.ResourceOwnerId = &v
	return s
}

type RefundCdiBagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundCdiBagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundCdiBagResponseBody) GoString() string {
	return s.String()
}

func (s *RefundCdiBagResponseBody) SetRequestId(v string) *RefundCdiBagResponseBody {
	s.RequestId = &v
	return s
}

type RefundCdiBagResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundCdiBagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundCdiBagResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundCdiBagResponse) GoString() string {
	return s.String()
}

func (s *RefundCdiBagResponse) SetHeaders(v map[string]*string) *RefundCdiBagResponse {
	s.Headers = v
	return s
}

func (s *RefundCdiBagResponse) SetStatusCode(v int32) *RefundCdiBagResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundCdiBagResponse) SetBody(v *RefundCdiBagResponseBody) *RefundCdiBagResponse {
	s.Body = v
	return s
}

type RefundCdiBaseBagRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RefundCdiBaseBagRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundCdiBaseBagRequest) GoString() string {
	return s.String()
}

func (s *RefundCdiBaseBagRequest) SetInstanceId(v string) *RefundCdiBaseBagRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundCdiBaseBagRequest) SetResourceOwnerId(v string) *RefundCdiBaseBagRequest {
	s.ResourceOwnerId = &v
	return s
}

type RefundCdiBaseBagResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundCdiBaseBagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundCdiBaseBagResponseBody) GoString() string {
	return s.String()
}

func (s *RefundCdiBaseBagResponseBody) SetRequestId(v string) *RefundCdiBaseBagResponseBody {
	s.RequestId = &v
	return s
}

type RefundCdiBaseBagResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundCdiBaseBagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundCdiBaseBagResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundCdiBaseBagResponse) GoString() string {
	return s.String()
}

func (s *RefundCdiBaseBagResponse) SetHeaders(v map[string]*string) *RefundCdiBaseBagResponse {
	s.Headers = v
	return s
}

func (s *RefundCdiBaseBagResponse) SetStatusCode(v int32) *RefundCdiBaseBagResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundCdiBaseBagResponse) SetBody(v *RefundCdiBaseBagResponseBody) *RefundCdiBaseBagResponse {
	s.Body = v
	return s
}

type RefundWebSiteInstanceRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RefundWebSiteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundWebSiteInstanceRequest) GoString() string {
	return s.String()
}

func (s *RefundWebSiteInstanceRequest) SetInstanceId(v string) *RefundWebSiteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundWebSiteInstanceRequest) SetResourceOwnerId(v string) *RefundWebSiteInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

type RefundWebSiteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundWebSiteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundWebSiteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RefundWebSiteInstanceResponseBody) SetRequestId(v string) *RefundWebSiteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RefundWebSiteInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundWebSiteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundWebSiteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundWebSiteInstanceResponse) GoString() string {
	return s.String()
}

func (s *RefundWebSiteInstanceResponse) SetHeaders(v map[string]*string) *RefundWebSiteInstanceResponse {
	s.Headers = v
	return s
}

func (s *RefundWebSiteInstanceResponse) SetStatusCode(v int32) *RefundWebSiteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundWebSiteInstanceResponse) SetBody(v *RefundWebSiteInstanceResponseBody) *RefundWebSiteInstanceResponse {
	s.Body = v
	return s
}

type RenewWebSiteInstanceRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderNum      *int32  `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s RenewWebSiteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewWebSiteInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewWebSiteInstanceRequest) SetClientToken(v string) *RenewWebSiteInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewWebSiteInstanceRequest) SetCommodityCode(v string) *RenewWebSiteInstanceRequest {
	s.CommodityCode = &v
	return s
}

func (s *RenewWebSiteInstanceRequest) SetDuration(v int32) *RenewWebSiteInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewWebSiteInstanceRequest) SetInstanceId(v string) *RenewWebSiteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewWebSiteInstanceRequest) SetOrderNum(v int32) *RenewWebSiteInstanceRequest {
	s.OrderNum = &v
	return s
}

func (s *RenewWebSiteInstanceRequest) SetOrderType(v string) *RenewWebSiteInstanceRequest {
	s.OrderType = &v
	return s
}

func (s *RenewWebSiteInstanceRequest) SetOwnerId(v int64) *RenewWebSiteInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewWebSiteInstanceRequest) SetPricingCycle(v string) *RenewWebSiteInstanceRequest {
	s.PricingCycle = &v
	return s
}

type RenewWebSiteInstanceResponseBody struct {
	Code        *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId  *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIds *RenewWebSiteInstanceResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	Message     *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId     *string                                      `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewWebSiteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewWebSiteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewWebSiteInstanceResponseBody) SetCode(v string) *RenewWebSiteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RenewWebSiteInstanceResponseBody) SetInstanceId(v string) *RenewWebSiteInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RenewWebSiteInstanceResponseBody) SetInstanceIds(v *RenewWebSiteInstanceResponseBodyInstanceIds) *RenewWebSiteInstanceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *RenewWebSiteInstanceResponseBody) SetMessage(v string) *RenewWebSiteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RenewWebSiteInstanceResponseBody) SetOrderId(v string) *RenewWebSiteInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewWebSiteInstanceResponseBody) SetRequestId(v string) *RenewWebSiteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RenewWebSiteInstanceResponseBodyInstanceIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s RenewWebSiteInstanceResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s RenewWebSiteInstanceResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *RenewWebSiteInstanceResponseBodyInstanceIds) SetString_(v []*string) *RenewWebSiteInstanceResponseBodyInstanceIds {
	s.String_ = v
	return s
}

type RenewWebSiteInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewWebSiteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewWebSiteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewWebSiteInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewWebSiteInstanceResponse) SetHeaders(v map[string]*string) *RenewWebSiteInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewWebSiteInstanceResponse) SetStatusCode(v int32) *RenewWebSiteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewWebSiteInstanceResponse) SetBody(v *RenewWebSiteInstanceResponseBody) *RenewWebSiteInstanceResponse {
	s.Body = v
	return s
}

type SendVerifyCodeToEmailRequest struct {
	Email    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s SendVerifyCodeToEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeToEmailRequest) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeToEmailRequest) SetEmail(v string) *SendVerifyCodeToEmailRequest {
	s.Email = &v
	return s
}

func (s *SendVerifyCodeToEmailRequest) SetLang(v string) *SendVerifyCodeToEmailRequest {
	s.Lang = &v
	return s
}

func (s *SendVerifyCodeToEmailRequest) SetSourceIp(v string) *SendVerifyCodeToEmailRequest {
	s.SourceIp = &v
	return s
}

type SendVerifyCodeToEmailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerifyCodeToEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeToEmailResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeToEmailResponseBody) SetRequestId(v string) *SendVerifyCodeToEmailResponseBody {
	s.RequestId = &v
	return s
}

type SendVerifyCodeToEmailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerifyCodeToEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerifyCodeToEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeToEmailResponse) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeToEmailResponse) SetHeaders(v map[string]*string) *SendVerifyCodeToEmailResponse {
	s.Headers = v
	return s
}

func (s *SendVerifyCodeToEmailResponse) SetStatusCode(v int32) *SendVerifyCodeToEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerifyCodeToEmailResponse) SetBody(v *SendVerifyCodeToEmailResponseBody) *SendVerifyCodeToEmailResponse {
	s.Body = v
	return s
}

type SendVerifyCodeToPhoneRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Phone    *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s SendVerifyCodeToPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeToPhoneRequest) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeToPhoneRequest) SetLang(v string) *SendVerifyCodeToPhoneRequest {
	s.Lang = &v
	return s
}

func (s *SendVerifyCodeToPhoneRequest) SetPhone(v string) *SendVerifyCodeToPhoneRequest {
	s.Phone = &v
	return s
}

func (s *SendVerifyCodeToPhoneRequest) SetSourceIp(v string) *SendVerifyCodeToPhoneRequest {
	s.SourceIp = &v
	return s
}

type SendVerifyCodeToPhoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerifyCodeToPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeToPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeToPhoneResponseBody) SetRequestId(v string) *SendVerifyCodeToPhoneResponseBody {
	s.RequestId = &v
	return s
}

type SendVerifyCodeToPhoneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerifyCodeToPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerifyCodeToPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerifyCodeToPhoneResponse) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeToPhoneResponse) SetHeaders(v map[string]*string) *SendVerifyCodeToPhoneResponse {
	s.Headers = v
	return s
}

func (s *SendVerifyCodeToPhoneResponse) SetStatusCode(v int32) *SendVerifyCodeToPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerifyCodeToPhoneResponse) SetBody(v *SendVerifyCodeToPhoneResponseBody) *SendVerifyCodeToPhoneResponse {
	s.Body = v
	return s
}

type SendWebsiteFeedbackRequest struct {
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Urls     *string `json:"Urls,omitempty" xml:"Urls,omitempty"`
}

func (s SendWebsiteFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s SendWebsiteFeedbackRequest) GoString() string {
	return s.String()
}

func (s *SendWebsiteFeedbackRequest) SetFeedback(v string) *SendWebsiteFeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *SendWebsiteFeedbackRequest) SetLang(v string) *SendWebsiteFeedbackRequest {
	s.Lang = &v
	return s
}

func (s *SendWebsiteFeedbackRequest) SetSourceIp(v string) *SendWebsiteFeedbackRequest {
	s.SourceIp = &v
	return s
}

func (s *SendWebsiteFeedbackRequest) SetUrls(v string) *SendWebsiteFeedbackRequest {
	s.Urls = &v
	return s
}

type SendWebsiteFeedbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendWebsiteFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendWebsiteFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *SendWebsiteFeedbackResponseBody) SetRequestId(v string) *SendWebsiteFeedbackResponseBody {
	s.RequestId = &v
	return s
}

type SendWebsiteFeedbackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendWebsiteFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendWebsiteFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s SendWebsiteFeedbackResponse) GoString() string {
	return s.String()
}

func (s *SendWebsiteFeedbackResponse) SetHeaders(v map[string]*string) *SendWebsiteFeedbackResponse {
	s.Headers = v
	return s
}

func (s *SendWebsiteFeedbackResponse) SetStatusCode(v int32) *SendWebsiteFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *SendWebsiteFeedbackResponse) SetBody(v *SendWebsiteFeedbackResponseBody) *SendWebsiteFeedbackResponse {
	s.Body = v
	return s
}

type UpdateAppPackageRequest struct {
	Debug      *bool   `json:"Debug,omitempty" xml:"Debug,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UpdateAppPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppPackageRequest) SetDebug(v bool) *UpdateAppPackageRequest {
	s.Debug = &v
	return s
}

func (s *UpdateAppPackageRequest) SetId(v int64) *UpdateAppPackageRequest {
	s.Id = &v
	return s
}

func (s *UpdateAppPackageRequest) SetLang(v string) *UpdateAppPackageRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAppPackageRequest) SetPackageUrl(v string) *UpdateAppPackageRequest {
	s.PackageUrl = &v
	return s
}

func (s *UpdateAppPackageRequest) SetPlatform(v string) *UpdateAppPackageRequest {
	s.Platform = &v
	return s
}

func (s *UpdateAppPackageRequest) SetSourceIp(v string) *UpdateAppPackageRequest {
	s.SourceIp = &v
	return s
}

type UpdateAppPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateAppPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppPackageResponseBody) SetRequestId(v string) *UpdateAppPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppPackageResponseBody) SetTaskId(v string) *UpdateAppPackageResponseBody {
	s.TaskId = &v
	return s
}

type UpdateAppPackageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppPackageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppPackageResponse) SetHeaders(v map[string]*string) *UpdateAppPackageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppPackageResponse) SetStatusCode(v int32) *UpdateAppPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppPackageResponse) SetBody(v *UpdateAppPackageResponseBody) *UpdateAppPackageResponse {
	s.Body = v
	return s
}

type UpdateAuditCallbackRequest struct {
	Callback  *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CryptType *int32  `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	Seed      *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
}

func (s UpdateAuditCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditCallbackRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuditCallbackRequest) SetCallback(v string) *UpdateAuditCallbackRequest {
	s.Callback = &v
	return s
}

func (s *UpdateAuditCallbackRequest) SetCryptType(v int32) *UpdateAuditCallbackRequest {
	s.CryptType = &v
	return s
}

func (s *UpdateAuditCallbackRequest) SetSeed(v string) *UpdateAuditCallbackRequest {
	s.Seed = &v
	return s
}

type UpdateAuditCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuditCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuditCallbackResponseBody) SetRequestId(v string) *UpdateAuditCallbackResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAuditCallbackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuditCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuditCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditCallbackResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuditCallbackResponse) SetHeaders(v map[string]*string) *UpdateAuditCallbackResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuditCallbackResponse) SetStatusCode(v int32) *UpdateAuditCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuditCallbackResponse) SetBody(v *UpdateAuditCallbackResponseBody) *UpdateAuditCallbackResponse {
	s.Body = v
	return s
}

type UpdateAuditRangeRequest struct {
	AuditRange *string `json:"AuditRange,omitempty" xml:"AuditRange,omitempty"`
}

func (s UpdateAuditRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditRangeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuditRangeRequest) SetAuditRange(v string) *UpdateAuditRangeRequest {
	s.AuditRange = &v
	return s
}

type UpdateAuditRangeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuditRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditRangeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuditRangeResponseBody) SetRequestId(v string) *UpdateAuditRangeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAuditRangeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuditRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuditRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditRangeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuditRangeResponse) SetHeaders(v map[string]*string) *UpdateAuditRangeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuditRangeResponse) SetStatusCode(v int32) *UpdateAuditRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuditRangeResponse) SetBody(v *UpdateAuditRangeResponseBody) *UpdateAuditRangeResponse {
	s.Body = v
	return s
}

type UpdateAuditSettingRequest struct {
	AuditRange *string `json:"AuditRange,omitempty" xml:"AuditRange,omitempty"`
	Callback   *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Seed       *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UpdateAuditSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuditSettingRequest) SetAuditRange(v string) *UpdateAuditSettingRequest {
	s.AuditRange = &v
	return s
}

func (s *UpdateAuditSettingRequest) SetCallback(v string) *UpdateAuditSettingRequest {
	s.Callback = &v
	return s
}

func (s *UpdateAuditSettingRequest) SetSeed(v string) *UpdateAuditSettingRequest {
	s.Seed = &v
	return s
}

func (s *UpdateAuditSettingRequest) SetSourceIp(v string) *UpdateAuditSettingRequest {
	s.SourceIp = &v
	return s
}

type UpdateAuditSettingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAuditSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuditSettingResponseBody) SetRequestId(v string) *UpdateAuditSettingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAuditSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuditSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuditSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuditSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuditSettingResponse) SetHeaders(v map[string]*string) *UpdateAuditSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuditSettingResponse) SetStatusCode(v int32) *UpdateAuditSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuditSettingResponse) SetBody(v *UpdateAuditSettingResponseBody) *UpdateAuditSettingResponse {
	s.Body = v
	return s
}

type UpdateBizTypeRequest struct {
	BizTypeName *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeRequest) SetBizTypeName(v string) *UpdateBizTypeRequest {
	s.BizTypeName = &v
	return s
}

func (s *UpdateBizTypeRequest) SetDescription(v string) *UpdateBizTypeRequest {
	s.Description = &v
	return s
}

type UpdateBizTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeResponseBody) SetRequestId(v string) *UpdateBizTypeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBizTypeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeResponse) SetHeaders(v map[string]*string) *UpdateBizTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizTypeResponse) SetStatusCode(v int32) *UpdateBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizTypeResponse) SetBody(v *UpdateBizTypeResponseBody) *UpdateBizTypeResponse {
	s.Body = v
	return s
}

type UpdateBizTypeImageLibRequest struct {
	BizTypeName  *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	Black        *string `json:"Black,omitempty" xml:"Black,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Review       *string `json:"Review,omitempty" xml:"Review,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	White        *string `json:"White,omitempty" xml:"White,omitempty"`
}

func (s UpdateBizTypeImageLibRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeImageLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeImageLibRequest) SetBizTypeName(v string) *UpdateBizTypeImageLibRequest {
	s.BizTypeName = &v
	return s
}

func (s *UpdateBizTypeImageLibRequest) SetBlack(v string) *UpdateBizTypeImageLibRequest {
	s.Black = &v
	return s
}

func (s *UpdateBizTypeImageLibRequest) SetResourceType(v string) *UpdateBizTypeImageLibRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateBizTypeImageLibRequest) SetReview(v string) *UpdateBizTypeImageLibRequest {
	s.Review = &v
	return s
}

func (s *UpdateBizTypeImageLibRequest) SetScene(v string) *UpdateBizTypeImageLibRequest {
	s.Scene = &v
	return s
}

func (s *UpdateBizTypeImageLibRequest) SetWhite(v string) *UpdateBizTypeImageLibRequest {
	s.White = &v
	return s
}

type UpdateBizTypeImageLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBizTypeImageLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeImageLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeImageLibResponseBody) SetRequestId(v string) *UpdateBizTypeImageLibResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBizTypeImageLibResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizTypeImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizTypeImageLibResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeImageLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeImageLibResponse) SetHeaders(v map[string]*string) *UpdateBizTypeImageLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizTypeImageLibResponse) SetStatusCode(v int32) *UpdateBizTypeImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizTypeImageLibResponse) SetBody(v *UpdateBizTypeImageLibResponseBody) *UpdateBizTypeImageLibResponse {
	s.Body = v
	return s
}

type UpdateBizTypeSettingRequest struct {
	Ad           *string `json:"Ad,omitempty" xml:"Ad,omitempty"`
	Antispam     *string `json:"Antispam,omitempty" xml:"Antispam,omitempty"`
	BizTypeName  *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	Live         *string `json:"Live,omitempty" xml:"Live,omitempty"`
	Porn         *string `json:"Porn,omitempty" xml:"Porn,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Terrorism    *string `json:"Terrorism,omitempty" xml:"Terrorism,omitempty"`
}

func (s UpdateBizTypeSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeSettingRequest) SetAd(v string) *UpdateBizTypeSettingRequest {
	s.Ad = &v
	return s
}

func (s *UpdateBizTypeSettingRequest) SetAntispam(v string) *UpdateBizTypeSettingRequest {
	s.Antispam = &v
	return s
}

func (s *UpdateBizTypeSettingRequest) SetBizTypeName(v string) *UpdateBizTypeSettingRequest {
	s.BizTypeName = &v
	return s
}

func (s *UpdateBizTypeSettingRequest) SetLive(v string) *UpdateBizTypeSettingRequest {
	s.Live = &v
	return s
}

func (s *UpdateBizTypeSettingRequest) SetPorn(v string) *UpdateBizTypeSettingRequest {
	s.Porn = &v
	return s
}

func (s *UpdateBizTypeSettingRequest) SetResourceType(v string) *UpdateBizTypeSettingRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateBizTypeSettingRequest) SetTerrorism(v string) *UpdateBizTypeSettingRequest {
	s.Terrorism = &v
	return s
}

type UpdateBizTypeSettingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBizTypeSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeSettingResponseBody) SetRequestId(v string) *UpdateBizTypeSettingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBizTypeSettingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizTypeSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizTypeSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeSettingResponse) SetHeaders(v map[string]*string) *UpdateBizTypeSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizTypeSettingResponse) SetStatusCode(v int32) *UpdateBizTypeSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizTypeSettingResponse) SetBody(v *UpdateBizTypeSettingResponseBody) *UpdateBizTypeSettingResponse {
	s.Body = v
	return s
}

type UpdateBizTypeTextLibRequest struct {
	BizTypeName  *string `json:"BizTypeName,omitempty" xml:"BizTypeName,omitempty"`
	Black        *string `json:"Black,omitempty" xml:"Black,omitempty"`
	Ignore       *string `json:"Ignore,omitempty" xml:"Ignore,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Review       *string `json:"Review,omitempty" xml:"Review,omitempty"`
	Scene        *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	White        *string `json:"White,omitempty" xml:"White,omitempty"`
}

func (s UpdateBizTypeTextLibRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeTextLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeTextLibRequest) SetBizTypeName(v string) *UpdateBizTypeTextLibRequest {
	s.BizTypeName = &v
	return s
}

func (s *UpdateBizTypeTextLibRequest) SetBlack(v string) *UpdateBizTypeTextLibRequest {
	s.Black = &v
	return s
}

func (s *UpdateBizTypeTextLibRequest) SetIgnore(v string) *UpdateBizTypeTextLibRequest {
	s.Ignore = &v
	return s
}

func (s *UpdateBizTypeTextLibRequest) SetResourceType(v string) *UpdateBizTypeTextLibRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateBizTypeTextLibRequest) SetReview(v string) *UpdateBizTypeTextLibRequest {
	s.Review = &v
	return s
}

func (s *UpdateBizTypeTextLibRequest) SetScene(v string) *UpdateBizTypeTextLibRequest {
	s.Scene = &v
	return s
}

func (s *UpdateBizTypeTextLibRequest) SetWhite(v string) *UpdateBizTypeTextLibRequest {
	s.White = &v
	return s
}

type UpdateBizTypeTextLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBizTypeTextLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeTextLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeTextLibResponseBody) SetRequestId(v string) *UpdateBizTypeTextLibResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBizTypeTextLibResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBizTypeTextLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBizTypeTextLibResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizTypeTextLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizTypeTextLibResponse) SetHeaders(v map[string]*string) *UpdateBizTypeTextLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizTypeTextLibResponse) SetStatusCode(v int32) *UpdateBizTypeTextLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizTypeTextLibResponse) SetBody(v *UpdateBizTypeTextLibResponseBody) *UpdateBizTypeTextLibResponse {
	s.Body = v
	return s
}

type UpdateCustomOcrTemplateRequest struct {
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RecognizeArea *string `json:"RecognizeArea,omitempty" xml:"RecognizeArea,omitempty"`
	ReferArea     *string `json:"ReferArea,omitempty" xml:"ReferArea,omitempty"`
}

func (s UpdateCustomOcrTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomOcrTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomOcrTemplateRequest) SetId(v int64) *UpdateCustomOcrTemplateRequest {
	s.Id = &v
	return s
}

func (s *UpdateCustomOcrTemplateRequest) SetName(v string) *UpdateCustomOcrTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomOcrTemplateRequest) SetRecognizeArea(v string) *UpdateCustomOcrTemplateRequest {
	s.RecognizeArea = &v
	return s
}

func (s *UpdateCustomOcrTemplateRequest) SetReferArea(v string) *UpdateCustomOcrTemplateRequest {
	s.ReferArea = &v
	return s
}

type UpdateCustomOcrTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomOcrTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomOcrTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomOcrTemplateResponseBody) SetRequestId(v string) *UpdateCustomOcrTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCustomOcrTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomOcrTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomOcrTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomOcrTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomOcrTemplateResponse) SetHeaders(v map[string]*string) *UpdateCustomOcrTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomOcrTemplateResponse) SetStatusCode(v int32) *UpdateCustomOcrTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomOcrTemplateResponse) SetBody(v *UpdateCustomOcrTemplateResponseBody) *UpdateCustomOcrTemplateResponse {
	s.Body = v
	return s
}

type UpdateKeywordLibRequest struct {
	BizTypes  *string `json:"BizTypes,omitempty" xml:"BizTypes,omitempty"`
	Enable    *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UpdateKeywordLibRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeywordLibRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibRequest) SetBizTypes(v string) *UpdateKeywordLibRequest {
	s.BizTypes = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetEnable(v bool) *UpdateKeywordLibRequest {
	s.Enable = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetId(v int32) *UpdateKeywordLibRequest {
	s.Id = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetLang(v string) *UpdateKeywordLibRequest {
	s.Lang = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetMatchMode(v string) *UpdateKeywordLibRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetName(v string) *UpdateKeywordLibRequest {
	s.Name = &v
	return s
}

func (s *UpdateKeywordLibRequest) SetSourceIp(v string) *UpdateKeywordLibRequest {
	s.SourceIp = &v
	return s
}

type UpdateKeywordLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateKeywordLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibResponseBody) SetRequestId(v string) *UpdateKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

type UpdateKeywordLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKeywordLibResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateKeywordLibResponse) SetHeaders(v map[string]*string) *UpdateKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateKeywordLibResponse) SetStatusCode(v int32) *UpdateKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKeywordLibResponse) SetBody(v *UpdateKeywordLibResponseBody) *UpdateKeywordLibResponse {
	s.Body = v
	return s
}

type UpdateNotificationSettingRequest struct {
	Lang                    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RealtimeMessageList     *string `json:"RealtimeMessageList,omitempty" xml:"RealtimeMessageList,omitempty"`
	ReminderModeList        *string `json:"ReminderModeList,omitempty" xml:"ReminderModeList,omitempty"`
	ScheduleMessageTime     *int32  `json:"ScheduleMessageTime,omitempty" xml:"ScheduleMessageTime,omitempty"`
	ScheduleMessageTimeZone *int32  `json:"ScheduleMessageTimeZone,omitempty" xml:"ScheduleMessageTimeZone,omitempty"`
	SourceIp                *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UpdateNotificationSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNotificationSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateNotificationSettingRequest) SetLang(v string) *UpdateNotificationSettingRequest {
	s.Lang = &v
	return s
}

func (s *UpdateNotificationSettingRequest) SetRealtimeMessageList(v string) *UpdateNotificationSettingRequest {
	s.RealtimeMessageList = &v
	return s
}

func (s *UpdateNotificationSettingRequest) SetReminderModeList(v string) *UpdateNotificationSettingRequest {
	s.ReminderModeList = &v
	return s
}

func (s *UpdateNotificationSettingRequest) SetScheduleMessageTime(v int32) *UpdateNotificationSettingRequest {
	s.ScheduleMessageTime = &v
	return s
}

func (s *UpdateNotificationSettingRequest) SetScheduleMessageTimeZone(v int32) *UpdateNotificationSettingRequest {
	s.ScheduleMessageTimeZone = &v
	return s
}

func (s *UpdateNotificationSettingRequest) SetSourceIp(v string) *UpdateNotificationSettingRequest {
	s.SourceIp = &v
	return s
}

type UpdateNotificationSettingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNotificationSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNotificationSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNotificationSettingResponseBody) SetRequestId(v string) *UpdateNotificationSettingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNotificationSettingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNotificationSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNotificationSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNotificationSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateNotificationSettingResponse) SetHeaders(v map[string]*string) *UpdateNotificationSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateNotificationSettingResponse) SetStatusCode(v int32) *UpdateNotificationSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNotificationSettingResponse) SetBody(v *UpdateNotificationSettingResponseBody) *UpdateNotificationSettingResponse {
	s.Body = v
	return s
}

type UpdateOssCallbackSettingRequest struct {
	AuditCallback           *bool   `json:"AuditCallback,omitempty" xml:"AuditCallback,omitempty"`
	CallbackSeed            *string `json:"CallbackSeed,omitempty" xml:"CallbackSeed,omitempty"`
	CallbackUrl             *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ScanCallback            *bool   `json:"ScanCallback,omitempty" xml:"ScanCallback,omitempty"`
	ScanCallbackSuggestions *string `json:"ScanCallbackSuggestions,omitempty" xml:"ScanCallbackSuggestions,omitempty"`
	ServiceModules          *string `json:"ServiceModules,omitempty" xml:"ServiceModules,omitempty"`
}

func (s UpdateOssCallbackSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssCallbackSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssCallbackSettingRequest) SetAuditCallback(v bool) *UpdateOssCallbackSettingRequest {
	s.AuditCallback = &v
	return s
}

func (s *UpdateOssCallbackSettingRequest) SetCallbackSeed(v string) *UpdateOssCallbackSettingRequest {
	s.CallbackSeed = &v
	return s
}

func (s *UpdateOssCallbackSettingRequest) SetCallbackUrl(v string) *UpdateOssCallbackSettingRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateOssCallbackSettingRequest) SetScanCallback(v bool) *UpdateOssCallbackSettingRequest {
	s.ScanCallback = &v
	return s
}

func (s *UpdateOssCallbackSettingRequest) SetScanCallbackSuggestions(v string) *UpdateOssCallbackSettingRequest {
	s.ScanCallbackSuggestions = &v
	return s
}

func (s *UpdateOssCallbackSettingRequest) SetServiceModules(v string) *UpdateOssCallbackSettingRequest {
	s.ServiceModules = &v
	return s
}

type UpdateOssCallbackSettingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOssCallbackSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssCallbackSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssCallbackSettingResponseBody) SetRequestId(v string) *UpdateOssCallbackSettingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOssCallbackSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssCallbackSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssCallbackSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssCallbackSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssCallbackSettingResponse) SetHeaders(v map[string]*string) *UpdateOssCallbackSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssCallbackSettingResponse) SetStatusCode(v int32) *UpdateOssCallbackSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssCallbackSettingResponse) SetBody(v *UpdateOssCallbackSettingResponseBody) *UpdateOssCallbackSettingResponse {
	s.Body = v
	return s
}

type UpdateOssIncrementCheckSettingRequest struct {
	AudioAntispamFreezeConfig      *string `json:"AudioAntispamFreezeConfig,omitempty" xml:"AudioAntispamFreezeConfig,omitempty"`
	AudioAutoFreezeOpened          *bool   `json:"AudioAutoFreezeOpened,omitempty" xml:"AudioAutoFreezeOpened,omitempty"`
	AudioMaxSize                   *int32  `json:"AudioMaxSize,omitempty" xml:"AudioMaxSize,omitempty"`
	AudioScanLimit                 *int64  `json:"AudioScanLimit,omitempty" xml:"AudioScanLimit,omitempty"`
	AudioSceneList                 *string `json:"AudioSceneList,omitempty" xml:"AudioSceneList,omitempty"`
	AutoFreezeType                 *string `json:"AutoFreezeType,omitempty" xml:"AutoFreezeType,omitempty"`
	BizType                        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BucketConfigList               *string `json:"BucketConfigList,omitempty" xml:"BucketConfigList,omitempty"`
	CallbackId                     *string `json:"CallbackId,omitempty" xml:"CallbackId,omitempty"`
	ImageAdFreezeConfig            *string `json:"ImageAdFreezeConfig,omitempty" xml:"ImageAdFreezeConfig,omitempty"`
	ImageAutoFreeze                *string `json:"ImageAutoFreeze,omitempty" xml:"ImageAutoFreeze,omitempty"`
	ImageAutoFreezeOpened          *bool   `json:"ImageAutoFreezeOpened,omitempty" xml:"ImageAutoFreezeOpened,omitempty"`
	ImageLiveFreezeConfig          *string `json:"ImageLiveFreezeConfig,omitempty" xml:"ImageLiveFreezeConfig,omitempty"`
	ImagePornFreezeConfig          *string `json:"ImagePornFreezeConfig,omitempty" xml:"ImagePornFreezeConfig,omitempty"`
	ImageScanLimit                 *string `json:"ImageScanLimit,omitempty" xml:"ImageScanLimit,omitempty"`
	ImageSceneList                 *string `json:"ImageSceneList,omitempty" xml:"ImageSceneList,omitempty"`
	ImageTerrorismFreezeConfig     *string `json:"ImageTerrorismFreezeConfig,omitempty" xml:"ImageTerrorismFreezeConfig,omitempty"`
	Lang                           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ScanImageNoFileType            *bool   `json:"ScanImageNoFileType,omitempty" xml:"ScanImageNoFileType,omitempty"`
	SourceIp                       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VideoAdFreezeConfig            *string `json:"VideoAdFreezeConfig,omitempty" xml:"VideoAdFreezeConfig,omitempty"`
	VideoAutoFreezeOpened          *bool   `json:"VideoAutoFreezeOpened,omitempty" xml:"VideoAutoFreezeOpened,omitempty"`
	VideoAutoFreezeSceneList       *string `json:"VideoAutoFreezeSceneList,omitempty" xml:"VideoAutoFreezeSceneList,omitempty"`
	VideoFrameInterval             *int32  `json:"VideoFrameInterval,omitempty" xml:"VideoFrameInterval,omitempty"`
	VideoLiveFreezeConfig          *string `json:"VideoLiveFreezeConfig,omitempty" xml:"VideoLiveFreezeConfig,omitempty"`
	VideoMaxFrames                 *int32  `json:"VideoMaxFrames,omitempty" xml:"VideoMaxFrames,omitempty"`
	VideoMaxSize                   *int32  `json:"VideoMaxSize,omitempty" xml:"VideoMaxSize,omitempty"`
	VideoPornFreezeConfig          *string `json:"VideoPornFreezeConfig,omitempty" xml:"VideoPornFreezeConfig,omitempty"`
	VideoScanLimit                 *int64  `json:"VideoScanLimit,omitempty" xml:"VideoScanLimit,omitempty"`
	VideoSceneList                 *string `json:"VideoSceneList,omitempty" xml:"VideoSceneList,omitempty"`
	VideoTerrorismFreezeConfig     *string `json:"VideoTerrorismFreezeConfig,omitempty" xml:"VideoTerrorismFreezeConfig,omitempty"`
	VideoVoiceAntispamFreezeConfig *string `json:"VideoVoiceAntispamFreezeConfig,omitempty" xml:"VideoVoiceAntispamFreezeConfig,omitempty"`
}

func (s UpdateOssIncrementCheckSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssIncrementCheckSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssIncrementCheckSettingRequest) SetAudioAntispamFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.AudioAntispamFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetAudioAutoFreezeOpened(v bool) *UpdateOssIncrementCheckSettingRequest {
	s.AudioAutoFreezeOpened = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetAudioMaxSize(v int32) *UpdateOssIncrementCheckSettingRequest {
	s.AudioMaxSize = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetAudioScanLimit(v int64) *UpdateOssIncrementCheckSettingRequest {
	s.AudioScanLimit = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetAudioSceneList(v string) *UpdateOssIncrementCheckSettingRequest {
	s.AudioSceneList = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetAutoFreezeType(v string) *UpdateOssIncrementCheckSettingRequest {
	s.AutoFreezeType = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetBizType(v string) *UpdateOssIncrementCheckSettingRequest {
	s.BizType = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetBucketConfigList(v string) *UpdateOssIncrementCheckSettingRequest {
	s.BucketConfigList = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetCallbackId(v string) *UpdateOssIncrementCheckSettingRequest {
	s.CallbackId = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImageAdFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.ImageAdFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImageAutoFreeze(v string) *UpdateOssIncrementCheckSettingRequest {
	s.ImageAutoFreeze = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImageAutoFreezeOpened(v bool) *UpdateOssIncrementCheckSettingRequest {
	s.ImageAutoFreezeOpened = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImageLiveFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.ImageLiveFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImagePornFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.ImagePornFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImageScanLimit(v string) *UpdateOssIncrementCheckSettingRequest {
	s.ImageScanLimit = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImageSceneList(v string) *UpdateOssIncrementCheckSettingRequest {
	s.ImageSceneList = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetImageTerrorismFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.ImageTerrorismFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetLang(v string) *UpdateOssIncrementCheckSettingRequest {
	s.Lang = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetScanImageNoFileType(v bool) *UpdateOssIncrementCheckSettingRequest {
	s.ScanImageNoFileType = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetSourceIp(v string) *UpdateOssIncrementCheckSettingRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoAdFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.VideoAdFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoAutoFreezeOpened(v bool) *UpdateOssIncrementCheckSettingRequest {
	s.VideoAutoFreezeOpened = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoAutoFreezeSceneList(v string) *UpdateOssIncrementCheckSettingRequest {
	s.VideoAutoFreezeSceneList = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoFrameInterval(v int32) *UpdateOssIncrementCheckSettingRequest {
	s.VideoFrameInterval = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoLiveFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.VideoLiveFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoMaxFrames(v int32) *UpdateOssIncrementCheckSettingRequest {
	s.VideoMaxFrames = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoMaxSize(v int32) *UpdateOssIncrementCheckSettingRequest {
	s.VideoMaxSize = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoPornFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.VideoPornFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoScanLimit(v int64) *UpdateOssIncrementCheckSettingRequest {
	s.VideoScanLimit = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoSceneList(v string) *UpdateOssIncrementCheckSettingRequest {
	s.VideoSceneList = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoTerrorismFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.VideoTerrorismFreezeConfig = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingRequest) SetVideoVoiceAntispamFreezeConfig(v string) *UpdateOssIncrementCheckSettingRequest {
	s.VideoVoiceAntispamFreezeConfig = &v
	return s
}

type UpdateOssIncrementCheckSettingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOssIncrementCheckSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssIncrementCheckSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssIncrementCheckSettingResponseBody) SetRequestId(v string) *UpdateOssIncrementCheckSettingResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOssIncrementCheckSettingResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssIncrementCheckSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssIncrementCheckSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssIncrementCheckSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssIncrementCheckSettingResponse) SetHeaders(v map[string]*string) *UpdateOssIncrementCheckSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssIncrementCheckSettingResponse) SetStatusCode(v int32) *UpdateOssIncrementCheckSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssIncrementCheckSettingResponse) SetBody(v *UpdateOssIncrementCheckSettingResponseBody) *UpdateOssIncrementCheckSettingResponse {
	s.Body = v
	return s
}

type UpdateOssStockStatusRequest struct {
	AudioAutoFreezeSceneList *string `json:"AudioAutoFreezeSceneList,omitempty" xml:"AudioAutoFreezeSceneList,omitempty"`
	AudioMaxSize             *int32  `json:"AudioMaxSize,omitempty" xml:"AudioMaxSize,omitempty"`
	AutoFreezeType           *string `json:"AutoFreezeType,omitempty" xml:"AutoFreezeType,omitempty"`
	BucketConfigList         *string `json:"BucketConfigList,omitempty" xml:"BucketConfigList,omitempty"`
	EndDate                  *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ImageAutoFreeze          *string `json:"ImageAutoFreeze,omitempty" xml:"ImageAutoFreeze,omitempty"`
	Lang                     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Operation                *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	ResourceTypeList         *string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty"`
	SceneList                *string `json:"SceneList,omitempty" xml:"SceneList,omitempty"`
	SourceIp                 *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartDate                *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	VideoAutoFreezeSceneList *string `json:"VideoAutoFreezeSceneList,omitempty" xml:"VideoAutoFreezeSceneList,omitempty"`
	VideoFrameInterval       *int32  `json:"VideoFrameInterval,omitempty" xml:"VideoFrameInterval,omitempty"`
	VideoMaxFrames           *int32  `json:"VideoMaxFrames,omitempty" xml:"VideoMaxFrames,omitempty"`
	VideoMaxSize             *int32  `json:"VideoMaxSize,omitempty" xml:"VideoMaxSize,omitempty"`
}

func (s UpdateOssStockStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssStockStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateOssStockStatusRequest) SetAudioAutoFreezeSceneList(v string) *UpdateOssStockStatusRequest {
	s.AudioAutoFreezeSceneList = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetAudioMaxSize(v int32) *UpdateOssStockStatusRequest {
	s.AudioMaxSize = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetAutoFreezeType(v string) *UpdateOssStockStatusRequest {
	s.AutoFreezeType = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetBucketConfigList(v string) *UpdateOssStockStatusRequest {
	s.BucketConfigList = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetEndDate(v string) *UpdateOssStockStatusRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetImageAutoFreeze(v string) *UpdateOssStockStatusRequest {
	s.ImageAutoFreeze = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetLang(v string) *UpdateOssStockStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetOperation(v string) *UpdateOssStockStatusRequest {
	s.Operation = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetResourceTypeList(v string) *UpdateOssStockStatusRequest {
	s.ResourceTypeList = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetSceneList(v string) *UpdateOssStockStatusRequest {
	s.SceneList = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetSourceIp(v string) *UpdateOssStockStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetStartDate(v string) *UpdateOssStockStatusRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetVideoAutoFreezeSceneList(v string) *UpdateOssStockStatusRequest {
	s.VideoAutoFreezeSceneList = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetVideoFrameInterval(v int32) *UpdateOssStockStatusRequest {
	s.VideoFrameInterval = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetVideoMaxFrames(v int32) *UpdateOssStockStatusRequest {
	s.VideoMaxFrames = &v
	return s
}

func (s *UpdateOssStockStatusRequest) SetVideoMaxSize(v int32) *UpdateOssStockStatusRequest {
	s.VideoMaxSize = &v
	return s
}

type UpdateOssStockStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOssStockStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssStockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssStockStatusResponseBody) SetRequestId(v string) *UpdateOssStockStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOssStockStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOssStockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOssStockStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOssStockStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateOssStockStatusResponse) SetHeaders(v map[string]*string) *UpdateOssStockStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateOssStockStatusResponse) SetStatusCode(v int32) *UpdateOssStockStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOssStockStatusResponse) SetBody(v *UpdateOssStockStatusResponseBody) *UpdateOssStockStatusResponse {
	s.Body = v
	return s
}

type UpdateWebsiteInstanceRequest struct {
	Domain                *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	IndexPage             *string `json:"IndexPage,omitempty" xml:"IndexPage,omitempty"`
	IndexPageScanInterval *int32  `json:"IndexPageScanInterval,omitempty" xml:"IndexPageScanInterval,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SiteProtocol          *string `json:"SiteProtocol,omitempty" xml:"SiteProtocol,omitempty"`
	SourceIp              *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	WebsiteScanInterval   *int32  `json:"WebsiteScanInterval,omitempty" xml:"WebsiteScanInterval,omitempty"`
}

func (s UpdateWebsiteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceRequest) SetDomain(v string) *UpdateWebsiteInstanceRequest {
	s.Domain = &v
	return s
}

func (s *UpdateWebsiteInstanceRequest) SetIndexPage(v string) *UpdateWebsiteInstanceRequest {
	s.IndexPage = &v
	return s
}

func (s *UpdateWebsiteInstanceRequest) SetIndexPageScanInterval(v int32) *UpdateWebsiteInstanceRequest {
	s.IndexPageScanInterval = &v
	return s
}

func (s *UpdateWebsiteInstanceRequest) SetInstanceId(v string) *UpdateWebsiteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWebsiteInstanceRequest) SetLang(v string) *UpdateWebsiteInstanceRequest {
	s.Lang = &v
	return s
}

func (s *UpdateWebsiteInstanceRequest) SetSiteProtocol(v string) *UpdateWebsiteInstanceRequest {
	s.SiteProtocol = &v
	return s
}

func (s *UpdateWebsiteInstanceRequest) SetSourceIp(v string) *UpdateWebsiteInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateWebsiteInstanceRequest) SetWebsiteScanInterval(v int32) *UpdateWebsiteInstanceRequest {
	s.WebsiteScanInterval = &v
	return s
}

type UpdateWebsiteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebsiteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceResponseBody) SetRequestId(v string) *UpdateWebsiteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWebsiteInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebsiteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebsiteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceResponse) SetHeaders(v map[string]*string) *UpdateWebsiteInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebsiteInstanceResponse) SetStatusCode(v int32) *UpdateWebsiteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebsiteInstanceResponse) SetBody(v *UpdateWebsiteInstanceResponseBody) *UpdateWebsiteInstanceResponse {
	s.Body = v
	return s
}

type UpdateWebsiteInstanceKeyUrlRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Urls       *string `json:"Urls,omitempty" xml:"Urls,omitempty"`
}

func (s UpdateWebsiteInstanceKeyUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceKeyUrlRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceKeyUrlRequest) SetInstanceId(v string) *UpdateWebsiteInstanceKeyUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWebsiteInstanceKeyUrlRequest) SetLang(v string) *UpdateWebsiteInstanceKeyUrlRequest {
	s.Lang = &v
	return s
}

func (s *UpdateWebsiteInstanceKeyUrlRequest) SetSourceIp(v string) *UpdateWebsiteInstanceKeyUrlRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateWebsiteInstanceKeyUrlRequest) SetUrls(v string) *UpdateWebsiteInstanceKeyUrlRequest {
	s.Urls = &v
	return s
}

type UpdateWebsiteInstanceKeyUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebsiteInstanceKeyUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceKeyUrlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceKeyUrlResponseBody) SetRequestId(v string) *UpdateWebsiteInstanceKeyUrlResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWebsiteInstanceKeyUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebsiteInstanceKeyUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebsiteInstanceKeyUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceKeyUrlResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceKeyUrlResponse) SetHeaders(v map[string]*string) *UpdateWebsiteInstanceKeyUrlResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebsiteInstanceKeyUrlResponse) SetStatusCode(v int32) *UpdateWebsiteInstanceKeyUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebsiteInstanceKeyUrlResponse) SetBody(v *UpdateWebsiteInstanceKeyUrlResponseBody) *UpdateWebsiteInstanceKeyUrlResponse {
	s.Body = v
	return s
}

type UpdateWebsiteInstanceStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateWebsiteInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceStatusRequest) SetInstanceId(v string) *UpdateWebsiteInstanceStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateWebsiteInstanceStatusRequest) SetLang(v string) *UpdateWebsiteInstanceStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateWebsiteInstanceStatusRequest) SetSourceIp(v string) *UpdateWebsiteInstanceStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateWebsiteInstanceStatusRequest) SetStatus(v string) *UpdateWebsiteInstanceStatusRequest {
	s.Status = &v
	return s
}

type UpdateWebsiteInstanceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebsiteInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceStatusResponseBody) SetRequestId(v string) *UpdateWebsiteInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWebsiteInstanceStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebsiteInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebsiteInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebsiteInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebsiteInstanceStatusResponse) SetHeaders(v map[string]*string) *UpdateWebsiteInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebsiteInstanceStatusResponse) SetStatusCode(v int32) *UpdateWebsiteInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebsiteInstanceStatusResponse) SetBody(v *UpdateWebsiteInstanceStatusResponseBody) *UpdateWebsiteInstanceStatusResponse {
	s.Body = v
	return s
}

type UpgradeCdiBaseBagRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	FlowOutSpec   *int32  `json:"FlowOutSpec,omitempty" xml:"FlowOutSpec,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s UpgradeCdiBaseBagRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeCdiBaseBagRequest) GoString() string {
	return s.String()
}

func (s *UpgradeCdiBaseBagRequest) SetClientToken(v string) *UpgradeCdiBaseBagRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeCdiBaseBagRequest) SetCommodityCode(v string) *UpgradeCdiBaseBagRequest {
	s.CommodityCode = &v
	return s
}

func (s *UpgradeCdiBaseBagRequest) SetFlowOutSpec(v int32) *UpgradeCdiBaseBagRequest {
	s.FlowOutSpec = &v
	return s
}

func (s *UpgradeCdiBaseBagRequest) SetInstanceId(v string) *UpgradeCdiBaseBagRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeCdiBaseBagRequest) SetOrderType(v string) *UpgradeCdiBaseBagRequest {
	s.OrderType = &v
	return s
}

func (s *UpgradeCdiBaseBagRequest) SetOwnerId(v int64) *UpgradeCdiBaseBagRequest {
	s.OwnerId = &v
	return s
}

type UpgradeCdiBaseBagResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeCdiBaseBagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeCdiBaseBagResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeCdiBaseBagResponseBody) SetCode(v string) *UpgradeCdiBaseBagResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeCdiBaseBagResponseBody) SetInstanceId(v string) *UpgradeCdiBaseBagResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpgradeCdiBaseBagResponseBody) SetMessage(v string) *UpgradeCdiBaseBagResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeCdiBaseBagResponseBody) SetOrderId(v string) *UpgradeCdiBaseBagResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeCdiBaseBagResponseBody) SetRequestId(v string) *UpgradeCdiBaseBagResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeCdiBaseBagResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeCdiBaseBagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeCdiBaseBagResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeCdiBaseBagResponse) GoString() string {
	return s.String()
}

func (s *UpgradeCdiBaseBagResponse) SetHeaders(v map[string]*string) *UpgradeCdiBaseBagResponse {
	s.Headers = v
	return s
}

func (s *UpgradeCdiBaseBagResponse) SetStatusCode(v int32) *UpgradeCdiBaseBagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeCdiBaseBagResponse) SetBody(v *UpgradeCdiBaseBagResponseBody) *UpgradeCdiBaseBagResponse {
	s.Body = v
	return s
}

type UploadImageToLibRequest struct {
	ImageLibId *int32  `json:"ImageLibId,omitempty" xml:"ImageLibId,omitempty"`
	Images     *string `json:"Images,omitempty" xml:"Images,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Urls       *string `json:"Urls,omitempty" xml:"Urls,omitempty"`
}

func (s UploadImageToLibRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadImageToLibRequest) GoString() string {
	return s.String()
}

func (s *UploadImageToLibRequest) SetImageLibId(v int32) *UploadImageToLibRequest {
	s.ImageLibId = &v
	return s
}

func (s *UploadImageToLibRequest) SetImages(v string) *UploadImageToLibRequest {
	s.Images = &v
	return s
}

func (s *UploadImageToLibRequest) SetSourceIp(v string) *UploadImageToLibRequest {
	s.SourceIp = &v
	return s
}

func (s *UploadImageToLibRequest) SetUrls(v string) *UploadImageToLibRequest {
	s.Urls = &v
	return s
}

type UploadImageToLibResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadImageToLibResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadImageToLibResponseBody) GoString() string {
	return s.String()
}

func (s *UploadImageToLibResponseBody) SetRequestId(v string) *UploadImageToLibResponseBody {
	s.RequestId = &v
	return s
}

type UploadImageToLibResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadImageToLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadImageToLibResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadImageToLibResponse) GoString() string {
	return s.String()
}

func (s *UploadImageToLibResponse) SetHeaders(v map[string]*string) *UploadImageToLibResponse {
	s.Headers = v
	return s
}

func (s *UploadImageToLibResponse) SetStatusCode(v int32) *UploadImageToLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadImageToLibResponse) SetBody(v *UploadImageToLibResponseBody) *UploadImageToLibResponse {
	s.Body = v
	return s
}

type VerifyCustomOcrTemplateRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TestImgUrl *string `json:"TestImgUrl,omitempty" xml:"TestImgUrl,omitempty"`
}

func (s VerifyCustomOcrTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyCustomOcrTemplateRequest) GoString() string {
	return s.String()
}

func (s *VerifyCustomOcrTemplateRequest) SetId(v int64) *VerifyCustomOcrTemplateRequest {
	s.Id = &v
	return s
}

func (s *VerifyCustomOcrTemplateRequest) SetTestImgUrl(v string) *VerifyCustomOcrTemplateRequest {
	s.TestImgUrl = &v
	return s
}

type VerifyCustomOcrTemplateResponseBody struct {
	ImageUrl      *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	RecognizeInfo *string `json:"RecognizeInfo,omitempty" xml:"RecognizeInfo,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyCustomOcrTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyCustomOcrTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCustomOcrTemplateResponseBody) SetImageUrl(v string) *VerifyCustomOcrTemplateResponseBody {
	s.ImageUrl = &v
	return s
}

func (s *VerifyCustomOcrTemplateResponseBody) SetRecognizeInfo(v string) *VerifyCustomOcrTemplateResponseBody {
	s.RecognizeInfo = &v
	return s
}

func (s *VerifyCustomOcrTemplateResponseBody) SetRequestId(v string) *VerifyCustomOcrTemplateResponseBody {
	s.RequestId = &v
	return s
}

type VerifyCustomOcrTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCustomOcrTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCustomOcrTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyCustomOcrTemplateResponse) GoString() string {
	return s.String()
}

func (s *VerifyCustomOcrTemplateResponse) SetHeaders(v map[string]*string) *VerifyCustomOcrTemplateResponse {
	s.Headers = v
	return s
}

func (s *VerifyCustomOcrTemplateResponse) SetStatusCode(v int32) *VerifyCustomOcrTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCustomOcrTemplateResponse) SetBody(v *VerifyCustomOcrTemplateResponseBody) *VerifyCustomOcrTemplateResponse {
	s.Body = v
	return s
}

type VerifyEmailRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s VerifyEmailRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyEmailRequest) GoString() string {
	return s.String()
}

func (s *VerifyEmailRequest) SetLang(v string) *VerifyEmailRequest {
	s.Lang = &v
	return s
}

func (s *VerifyEmailRequest) SetSourceIp(v string) *VerifyEmailRequest {
	s.SourceIp = &v
	return s
}

func (s *VerifyEmailRequest) SetVerifyCode(v string) *VerifyEmailRequest {
	s.VerifyCode = &v
	return s
}

type VerifyEmailResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyEmailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyEmailResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyEmailResponseBody) SetRequestId(v string) *VerifyEmailResponseBody {
	s.RequestId = &v
	return s
}

type VerifyEmailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyEmailResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyEmailResponse) GoString() string {
	return s.String()
}

func (s *VerifyEmailResponse) SetHeaders(v map[string]*string) *VerifyEmailResponse {
	s.Headers = v
	return s
}

func (s *VerifyEmailResponse) SetStatusCode(v int32) *VerifyEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyEmailResponse) SetBody(v *VerifyEmailResponseBody) *VerifyEmailResponse {
	s.Body = v
	return s
}

type VerifyPhoneRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Phone      *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s VerifyPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneRequest) GoString() string {
	return s.String()
}

func (s *VerifyPhoneRequest) SetLang(v string) *VerifyPhoneRequest {
	s.Lang = &v
	return s
}

func (s *VerifyPhoneRequest) SetPhone(v string) *VerifyPhoneRequest {
	s.Phone = &v
	return s
}

func (s *VerifyPhoneRequest) SetSourceIp(v string) *VerifyPhoneRequest {
	s.SourceIp = &v
	return s
}

func (s *VerifyPhoneRequest) SetVerifyCode(v string) *VerifyPhoneRequest {
	s.VerifyCode = &v
	return s
}

type VerifyPhoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPhoneResponseBody) SetRequestId(v string) *VerifyPhoneResponseBody {
	s.RequestId = &v
	return s
}

type VerifyPhoneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneResponse) GoString() string {
	return s.String()
}

func (s *VerifyPhoneResponse) SetHeaders(v map[string]*string) *VerifyPhoneResponse {
	s.Headers = v
	return s
}

func (s *VerifyPhoneResponse) SetStatusCode(v int32) *VerifyPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPhoneResponse) SetBody(v *VerifyPhoneResponseBody) *VerifyPhoneResponse {
	s.Body = v
	return s
}

type VerifyWebsiteInstanceRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VerifyMethod *string `json:"VerifyMethod,omitempty" xml:"VerifyMethod,omitempty"`
}

func (s VerifyWebsiteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebsiteInstanceRequest) GoString() string {
	return s.String()
}

func (s *VerifyWebsiteInstanceRequest) SetInstanceId(v string) *VerifyWebsiteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *VerifyWebsiteInstanceRequest) SetLang(v string) *VerifyWebsiteInstanceRequest {
	s.Lang = &v
	return s
}

func (s *VerifyWebsiteInstanceRequest) SetSourceIp(v string) *VerifyWebsiteInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *VerifyWebsiteInstanceRequest) SetVerifyMethod(v string) *VerifyWebsiteInstanceRequest {
	s.VerifyMethod = &v
	return s
}

type VerifyWebsiteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyWebsiteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebsiteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWebsiteInstanceResponseBody) SetRequestId(v string) *VerifyWebsiteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type VerifyWebsiteInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyWebsiteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyWebsiteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyWebsiteInstanceResponse) GoString() string {
	return s.String()
}

func (s *VerifyWebsiteInstanceResponse) SetHeaders(v map[string]*string) *VerifyWebsiteInstanceResponse {
	s.Headers = v
	return s
}

func (s *VerifyWebsiteInstanceResponse) SetStatusCode(v int32) *VerifyWebsiteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyWebsiteInstanceResponse) SetBody(v *VerifyWebsiteInstanceResponseBody) *VerifyWebsiteInstanceResponse {
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
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("green.aliyuncs.com"),
		"cn-hongkong":           tea.String("green.aliyuncs.com"),
		"cn-huhehaote":          tea.String("green.aliyuncs.com"),
		"cn-qingdao":            tea.String("green.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("green.aliyuncs.com"),
		"eu-central-1":          tea.String("green.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("green.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("green.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("green.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("green.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("green"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AuditItemSubmitWithOptions(request *AuditItemSubmitRequest, runtime *util.RuntimeOptions) (_result *AuditItemSubmitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuditItemSubmit"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuditItemSubmitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuditItemSubmit(request *AuditItemSubmitRequest) (_result *AuditItemSubmitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuditItemSubmitResponse{}
	_body, _err := client.AuditItemSubmitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatCustomOcrTemplateWithOptions(request *CreatCustomOcrTemplateRequest, runtime *util.RuntimeOptions) (_result *CreatCustomOcrTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImgUrl)) {
		query["ImgUrl"] = request.ImgUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RecognizeArea)) {
		query["RecognizeArea"] = request.RecognizeArea
	}

	if !tea.BoolValue(util.IsUnset(request.ReferArea)) {
		query["ReferArea"] = request.ReferArea
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatCustomOcrTemplate"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatCustomOcrTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatCustomOcrTemplate(request *CreatCustomOcrTemplateRequest) (_result *CreatCustomOcrTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatCustomOcrTemplateResponse{}
	_body, _err := client.CreatCustomOcrTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAuditCallbackWithOptions(request *CreateAuditCallbackRequest, runtime *util.RuntimeOptions) (_result *CreateAuditCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackSuggestions)) {
		query["CallbackSuggestions"] = request.CallbackSuggestions
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackTypes)) {
		query["CallbackTypes"] = request.CallbackTypes
	}

	if !tea.BoolValue(util.IsUnset(request.CryptType)) {
		query["CryptType"] = request.CryptType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuditCallback"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAuditCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAuditCallback(request *CreateAuditCallbackRequest) (_result *CreateAuditCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuditCallbackResponse{}
	_body, _err := client.CreateAuditCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBizTypeWithOptions(request *CreateBizTypeRequest, runtime *util.RuntimeOptions) (_result *CreateBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeImport)) {
		query["BizTypeImport"] = request.BizTypeImport
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.CiteTemplate)) {
		query["CiteTemplate"] = request.CiteTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IndustryInfo)) {
		query["IndustryInfo"] = request.IndustryInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBizType"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBizTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBizType(request *CreateBizTypeRequest) (_result *CreateBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBizTypeResponse{}
	_body, _err := client.CreateBizTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCdiBagWithOptions(request *CreateCdiBagRequest, runtime *util.RuntimeOptions) (_result *CreateCdiBagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.FlowOutSpec)) {
		query["FlowOutSpec"] = request.FlowOutSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		query["OrderNum"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCdiBag"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCdiBagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCdiBag(request *CreateCdiBagRequest) (_result *CreateCdiBagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCdiBagResponse{}
	_body, _err := client.CreateCdiBagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCdiBaseBagWithOptions(request *CreateCdiBaseBagRequest, runtime *util.RuntimeOptions) (_result *CreateCdiBaseBagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.FlowOutSpec)) {
		query["FlowOutSpec"] = request.FlowOutSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCdiBaseBag"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCdiBaseBagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCdiBaseBag(request *CreateCdiBaseBagRequest) (_result *CreateCdiBaseBagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCdiBaseBagResponse{}
	_body, _err := client.CreateCdiBaseBagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageLibWithOptions(request *CreateImageLibRequest, runtime *util.RuntimeOptions) (_result *CreateImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		query["BizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceModule)) {
		query["ServiceModule"] = request.ServiceModule
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageLib(request *CreateImageLibRequest) (_result *CreateImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageLibResponse{}
	_body, _err := client.CreateImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKeywordWithOptions(request *CreateKeywordRequest, runtime *util.RuntimeOptions) (_result *CreateKeywordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeywordLibId)) {
		query["KeywordLibId"] = request.KeywordLibId
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKeyword"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateKeywordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKeyword(request *CreateKeywordRequest) (_result *CreateKeywordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeywordResponse{}
	_body, _err := client.CreateKeywordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKeywordLibWithOptions(request *CreateKeywordLibRequest, runtime *util.RuntimeOptions) (_result *CreateKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		query["BizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.LibType)) {
		query["LibType"] = request.LibType
	}

	if !tea.BoolValue(util.IsUnset(request.MatchMode)) {
		query["MatchMode"] = request.MatchMode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceModule)) {
		query["ServiceModule"] = request.ServiceModule
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKeywordLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKeywordLib(request *CreateKeywordLibRequest) (_result *CreateKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeywordLibResponse{}
	_body, _err := client.CreateKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWebSiteInstanceWithOptions(request *CreateWebSiteInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateWebSiteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		query["OrderNum"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWebSiteInstance"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWebSiteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWebSiteInstance(request *CreateWebSiteInstanceRequest) (_result *CreateWebSiteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWebSiteInstanceResponse{}
	_body, _err := client.CreateWebSiteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWebsiteIndexPageBaselineWithOptions(request *CreateWebsiteIndexPageBaselineRequest, runtime *util.RuntimeOptions) (_result *CreateWebsiteIndexPageBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWebsiteIndexPageBaseline"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWebsiteIndexPageBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWebsiteIndexPageBaseline(request *CreateWebsiteIndexPageBaselineRequest) (_result *CreateWebsiteIndexPageBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWebsiteIndexPageBaselineResponse{}
	_body, _err := client.CreateWebsiteIndexPageBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBizTypeWithOptions(request *DeleteBizTypeRequest, runtime *util.RuntimeOptions) (_result *DeleteBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBizType"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBizTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBizType(request *DeleteBizTypeRequest) (_result *DeleteBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBizTypeResponse{}
	_body, _err := client.DeleteBizTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomOcrTemplateWithOptions(request *DeleteCustomOcrTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomOcrTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomOcrTemplate"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomOcrTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomOcrTemplate(request *DeleteCustomOcrTemplateRequest) (_result *DeleteCustomOcrTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomOcrTemplateResponse{}
	_body, _err := client.DeleteCustomOcrTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageFromLibWithOptions(request *DeleteImageFromLibRequest, runtime *util.RuntimeOptions) (_result *DeleteImageFromLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImageFromLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageFromLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImageFromLib(request *DeleteImageFromLibRequest) (_result *DeleteImageFromLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageFromLibResponse{}
	_body, _err := client.DeleteImageFromLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageLibWithOptions(request *DeleteImageLibRequest, runtime *util.RuntimeOptions) (_result *DeleteImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImageLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImageLib(request *DeleteImageLibRequest) (_result *DeleteImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageLibResponse{}
	_body, _err := client.DeleteImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKeywordWithOptions(request *DeleteKeywordRequest, runtime *util.RuntimeOptions) (_result *DeleteKeywordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordLibId)) {
		query["KeywordLibId"] = request.KeywordLibId
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeyword"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKeywordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKeyword(request *DeleteKeywordRequest) (_result *DeleteKeywordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeywordResponse{}
	_body, _err := client.DeleteKeywordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKeywordLibWithOptions(request *DeleteKeywordLibRequest, runtime *util.RuntimeOptions) (_result *DeleteKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeywordLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKeywordLib(request *DeleteKeywordLibRequest) (_result *DeleteKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeywordLibResponse{}
	_body, _err := client.DeleteKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNotificationContactsWithOptions(request *DeleteNotificationContactsRequest, runtime *util.RuntimeOptions) (_result *DeleteNotificationContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactTypes)) {
		query["ContactTypes"] = request.ContactTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNotificationContacts"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNotificationContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNotificationContacts(request *DeleteNotificationContactsRequest) (_result *DeleteNotificationContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNotificationContactsResponse{}
	_body, _err := client.DeleteNotificationContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppInfoWithOptions(request *DescribeAppInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppInfo"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppInfo(request *DescribeAppInfoRequest) (_result *DescribeAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppInfoResponse{}
	_body, _err := client.DescribeAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditCallbackWithOptions(runtime *util.RuntimeOptions) (_result *DescribeAuditCallbackResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditCallback"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditCallback() (_result *DescribeAuditCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditCallbackResponse{}
	_body, _err := client.DescribeAuditCallbackWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditCallbackListWithOptions(runtime *util.RuntimeOptions) (_result *DescribeAuditCallbackListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditCallbackList"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditCallbackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditCallbackList() (_result *DescribeAuditCallbackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditCallbackListResponse{}
	_body, _err := client.DescribeAuditCallbackListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditContentWithOptions(request *DescribeAuditContentRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditResult)) {
		query["AuditResult"] = request.AuditResult
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordId)) {
		query["KeywordId"] = request.KeywordId
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LibType)) {
		query["LibType"] = request.LibType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Suggestion)) {
		query["Suggestion"] = request.Suggestion
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditContent"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditContent(request *DescribeAuditContentRequest) (_result *DescribeAuditContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditContentResponse{}
	_body, _err := client.DescribeAuditContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditContentItemWithOptions(request *DescribeAuditContentItemRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditContentItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditContentItem"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditContentItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditContentItem(request *DescribeAuditContentItemRequest) (_result *DescribeAuditContentItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditContentItemResponse{}
	_body, _err := client.DescribeAuditContentItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditRangeWithOptions(runtime *util.RuntimeOptions) (_result *DescribeAuditRangeResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditRange"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditRange() (_result *DescribeAuditRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditRangeResponse{}
	_body, _err := client.DescribeAuditRangeWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuditSettingWithOptions(request *DescribeAuditSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeAuditSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuditSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuditSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuditSetting(request *DescribeAuditSettingRequest) (_result *DescribeAuditSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuditSettingResponse{}
	_body, _err := client.DescribeAuditSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBizTypeImageLibWithOptions(request *DescribeBizTypeImageLibRequest, runtime *util.RuntimeOptions) (_result *DescribeBizTypeImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBizTypeImageLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBizTypeImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBizTypeImageLib(request *DescribeBizTypeImageLibRequest) (_result *DescribeBizTypeImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBizTypeImageLibResponse{}
	_body, _err := client.DescribeBizTypeImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBizTypeSettingWithOptions(request *DescribeBizTypeSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeBizTypeSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBizTypeSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBizTypeSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBizTypeSetting(request *DescribeBizTypeSettingRequest) (_result *DescribeBizTypeSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBizTypeSettingResponse{}
	_body, _err := client.DescribeBizTypeSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBizTypeTextLibWithOptions(request *DescribeBizTypeTextLibRequest, runtime *util.RuntimeOptions) (_result *DescribeBizTypeTextLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBizTypeTextLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBizTypeTextLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBizTypeTextLib(request *DescribeBizTypeTextLibRequest) (_result *DescribeBizTypeTextLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBizTypeTextLibResponse{}
	_body, _err := client.DescribeBizTypeTextLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBizTypesWithOptions(request *DescribeBizTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeBizTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportFlag)) {
		query["ImportFlag"] = request.ImportFlag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBizTypes"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBizTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBizTypes(request *DescribeBizTypesRequest) (_result *DescribeBizTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBizTypesResponse{}
	_body, _err := client.DescribeBizTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudMonitorServicesWithOptions(request *DescribeCloudMonitorServicesRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudMonitorServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudMonitorServices"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudMonitorServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudMonitorServices(request *DescribeCloudMonitorServicesRequest) (_result *DescribeCloudMonitorServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudMonitorServicesResponse{}
	_body, _err := client.DescribeCloudMonitorServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomOcrTemplateWithOptions(request *DescribeCustomOcrTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomOcrTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCustomOcrTemplate"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCustomOcrTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomOcrTemplate(request *DescribeCustomOcrTemplateRequest) (_result *DescribeCustomOcrTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomOcrTemplateResponse{}
	_body, _err := client.DescribeCustomOcrTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageFromLibWithOptions(request *DescribeImageFromLibRequest, runtime *util.RuntimeOptions) (_result *DescribeImageFromLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ImageLibId)) {
		query["ImageLibId"] = request.ImageLibId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageFromLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageFromLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageFromLib(request *DescribeImageFromLibRequest) (_result *DescribeImageFromLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageFromLibResponse{}
	_body, _err := client.DescribeImageFromLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageLibWithOptions(request *DescribeImageLibRequest, runtime *util.RuntimeOptions) (_result *DescribeImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ServiceModule)) {
		query["ServiceModule"] = request.ServiceModule
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageLib(request *DescribeImageLibRequest) (_result *DescribeImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageLibResponse{}
	_body, _err := client.DescribeImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageUploadInfoWithOptions(request *DescribeImageUploadInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeImageUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageUploadInfo"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageUploadInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageUploadInfo(request *DescribeImageUploadInfoRequest) (_result *DescribeImageUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageUploadInfoResponse{}
	_body, _err := client.DescribeImageUploadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKeywordWithOptions(request *DescribeKeywordRequest, runtime *util.RuntimeOptions) (_result *DescribeKeywordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordLibId)) {
		query["KeywordLibId"] = request.KeywordLibId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKeyword"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeKeywordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKeyword(request *DescribeKeywordRequest) (_result *DescribeKeywordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeywordResponse{}
	_body, _err := client.DescribeKeywordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKeywordLibWithOptions(request *DescribeKeywordLibRequest, runtime *util.RuntimeOptions) (_result *DescribeKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceModule)) {
		query["ServiceModule"] = request.ServiceModule
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeKeywordLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKeywordLib(request *DescribeKeywordLibRequest) (_result *DescribeKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeywordLibResponse{}
	_body, _err := client.DescribeKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNotificationSettingWithOptions(request *DescribeNotificationSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeNotificationSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNotificationSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNotificationSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNotificationSetting(request *DescribeNotificationSettingRequest) (_result *DescribeNotificationSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNotificationSettingResponse{}
	_body, _err := client.DescribeNotificationSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpenApiRcpStatsWithOptions(request *DescribeOpenApiRcpStatsRequest, runtime *util.RuntimeOptions) (_result *DescribeOpenApiRcpStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpenApiRcpStats"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOpenApiRcpStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpenApiRcpStats(request *DescribeOpenApiRcpStatsRequest) (_result *DescribeOpenApiRcpStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpenApiRcpStatsResponse{}
	_body, _err := client.DescribeOpenApiRcpStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpenApiUsageWithOptions(request *DescribeOpenApiUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeOpenApiUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpenApiUsage"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOpenApiUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpenApiUsage(request *DescribeOpenApiUsageRequest) (_result *DescribeOpenApiUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpenApiUsageResponse{}
	_body, _err := client.DescribeOpenApiUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssCallbackSettingWithOptions(runtime *util.RuntimeOptions) (_result *DescribeOssCallbackSettingResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssCallbackSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssCallbackSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssCallbackSetting() (_result *DescribeOssCallbackSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssCallbackSettingResponse{}
	_body, _err := client.DescribeOssCallbackSettingWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssIncrementCheckSettingWithOptions(request *DescribeOssIncrementCheckSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeOssIncrementCheckSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssIncrementCheckSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssIncrementCheckSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssIncrementCheckSetting(request *DescribeOssIncrementCheckSettingRequest) (_result *DescribeOssIncrementCheckSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssIncrementCheckSettingResponse{}
	_body, _err := client.DescribeOssIncrementCheckSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssIncrementOverviewWithOptions(request *DescribeOssIncrementOverviewRequest, runtime *util.RuntimeOptions) (_result *DescribeOssIncrementOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssIncrementOverview"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssIncrementOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssIncrementOverview(request *DescribeOssIncrementOverviewRequest) (_result *DescribeOssIncrementOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssIncrementOverviewResponse{}
	_body, _err := client.DescribeOssIncrementOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssIncrementStatsWithOptions(request *DescribeOssIncrementStatsRequest, runtime *util.RuntimeOptions) (_result *DescribeOssIncrementStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssIncrementStats"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssIncrementStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssIncrementStats(request *DescribeOssIncrementStatsRequest) (_result *DescribeOssIncrementStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssIncrementStatsResponse{}
	_body, _err := client.DescribeOssIncrementStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssResultItemsWithOptions(request *DescribeOssResultItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeOssResultItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScore)) {
		query["MaxScore"] = request.MaxScore
	}

	if !tea.BoolValue(util.IsUnset(request.MinScore)) {
		query["MinScore"] = request.MinScore
	}

	if !tea.BoolValue(util.IsUnset(request.Object)) {
		query["Object"] = request.Object
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Stock)) {
		query["Stock"] = request.Stock
	}

	if !tea.BoolValue(util.IsUnset(request.StockTaskId)) {
		query["StockTaskId"] = request.StockTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Suggestion)) {
		query["Suggestion"] = request.Suggestion
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssResultItems"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssResultItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssResultItems(request *DescribeOssResultItemsRequest) (_result *DescribeOssResultItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssResultItemsResponse{}
	_body, _err := client.DescribeOssResultItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssStockStatusWithOptions(request *DescribeOssStockStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeOssStockStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StockTaskId)) {
		query["StockTaskId"] = request.StockTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOssStockStatus"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOssStockStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssStockStatus(request *DescribeOssStockStatusRequest) (_result *DescribeOssStockStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssStockStatusResponse{}
	_body, _err := client.DescribeOssStockStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSdkUrlWithOptions(request *DescribeSdkUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeSdkUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Debug)) {
		query["Debug"] = request.Debug
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSdkUrl"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSdkUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSdkUrl(request *DescribeSdkUrlRequest) (_result *DescribeSdkUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSdkUrlResponse{}
	_body, _err := client.DescribeSdkUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpdatePackageResultWithOptions(request *DescribeUpdatePackageResultRequest, runtime *util.RuntimeOptions) (_result *DescribeUpdatePackageResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUpdatePackageResult"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUpdatePackageResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpdatePackageResult(request *DescribeUpdatePackageResultRequest) (_result *DescribeUpdatePackageResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpdatePackageResultResponse{}
	_body, _err := client.DescribeUpdatePackageResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUploadInfoWithOptions(request *DescribeUploadInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeUploadInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Biz)) {
		query["Biz"] = request.Biz
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUploadInfo"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUploadInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest) (_result *DescribeUploadInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUploadInfoResponse{}
	_body, _err := client.DescribeUploadInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUsageBillWithOptions(request *DescribeUsageBillRequest, runtime *util.RuntimeOptions) (_result *DescribeUsageBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Day)) {
		query["Day"] = request.Day
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageBill"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUsageBill(request *DescribeUsageBillRequest) (_result *DescribeUsageBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsageBillResponse{}
	_body, _err := client.DescribeUsageBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserBizTypesWithOptions(request *DescribeUserBizTypesRequest, runtime *util.RuntimeOptions) (_result *DescribeUserBizTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Customized)) {
		query["Customized"] = request.Customized
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserBizTypes"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserBizTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserBizTypes(request *DescribeUserBizTypesRequest) (_result *DescribeUserBizTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserBizTypesResponse{}
	_body, _err := client.DescribeUserBizTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserStatusWithOptions(request *DescribeUserStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserStatus"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (_result *DescribeUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DescribeUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeViewContentWithOptions(request *DescribeViewContentRequest, runtime *util.RuntimeOptions) (_result *DescribeViewContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditResult)) {
		query["AuditResult"] = request.AuditResult
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		query["DataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordId)) {
		query["KeywordId"] = request.KeywordId
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.LibType)) {
		query["LibType"] = request.LibType
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Suggestion)) {
		query["Suggestion"] = request.Suggestion
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeViewContent"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeViewContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeViewContent(request *DescribeViewContentRequest) (_result *DescribeViewContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeViewContentResponse{}
	_body, _err := client.DescribeViewContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteIndexPageBaselineWithOptions(request *DescribeWebsiteIndexPageBaselineRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteIndexPageBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteIndexPageBaseline"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteIndexPageBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteIndexPageBaseline(request *DescribeWebsiteIndexPageBaselineRequest) (_result *DescribeWebsiteIndexPageBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteIndexPageBaselineResponse{}
	_body, _err := client.DescribeWebsiteIndexPageBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteInstanceWithOptions(request *DescribeWebsiteInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteInstance"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteInstance(request *DescribeWebsiteInstanceRequest) (_result *DescribeWebsiteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteInstanceResponse{}
	_body, _err := client.DescribeWebsiteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteInstanceIdWithOptions(request *DescribeWebsiteInstanceIdRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteInstanceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteInstanceId"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteInstanceId(request *DescribeWebsiteInstanceIdRequest) (_result *DescribeWebsiteInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteInstanceIdResponse{}
	_body, _err := client.DescribeWebsiteInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteInstanceKeyUrlWithOptions(request *DescribeWebsiteInstanceKeyUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteInstanceKeyUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteInstanceKeyUrl"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteInstanceKeyUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteInstanceKeyUrl(request *DescribeWebsiteInstanceKeyUrlRequest) (_result *DescribeWebsiteInstanceKeyUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteInstanceKeyUrlResponse{}
	_body, _err := client.DescribeWebsiteInstanceKeyUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteScanResultWithOptions(request *DescribeWebsiteScanResultRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteScanResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.HandleStatus)) {
		query["HandleStatus"] = request.HandleStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SiteUrl)) {
		query["SiteUrl"] = request.SiteUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SubServiceModule)) {
		query["SubServiceModule"] = request.SubServiceModule
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteScanResult"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteScanResult(request *DescribeWebsiteScanResultRequest) (_result *DescribeWebsiteScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteScanResultResponse{}
	_body, _err := client.DescribeWebsiteScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteScanResultDetailWithOptions(request *DescribeWebsiteScanResultDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteScanResultDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteScanResultDetail"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteScanResultDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteScanResultDetail(request *DescribeWebsiteScanResultDetailRequest) (_result *DescribeWebsiteScanResultDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteScanResultDetailResponse{}
	_body, _err := client.DescribeWebsiteScanResultDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteStatWithOptions(request *DescribeWebsiteStatRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteStat"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteStat(request *DescribeWebsiteStatRequest) (_result *DescribeWebsiteStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteStatResponse{}
	_body, _err := client.DescribeWebsiteStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebsiteVerifyInfoWithOptions(request *DescribeWebsiteVerifyInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeWebsiteVerifyInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebsiteVerifyInfo"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebsiteVerifyInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebsiteVerifyInfo(request *DescribeWebsiteVerifyInfoRequest) (_result *DescribeWebsiteVerifyInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebsiteVerifyInfoResponse{}
	_body, _err := client.DescribeWebsiteVerifyInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportKeywordsWithOptions(request *ExportKeywordsRequest, runtime *util.RuntimeOptions) (_result *ExportKeywordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeywordLibId)) {
		query["KeywordLibId"] = request.KeywordLibId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportKeywords"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportKeywordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportKeywords(request *ExportKeywordsRequest) (_result *ExportKeywordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportKeywordsResponse{}
	_body, _err := client.ExportKeywordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportOpenApiRcpStatsWithOptions(request *ExportOpenApiRcpStatsRequest, runtime *util.RuntimeOptions) (_result *ExportOpenApiRcpStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportOpenApiRcpStats"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportOpenApiRcpStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportOpenApiRcpStats(request *ExportOpenApiRcpStatsRequest) (_result *ExportOpenApiRcpStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportOpenApiRcpStatsResponse{}
	_body, _err := client.ExportOpenApiRcpStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportOssResultWithOptions(request *ExportOssResultRequest, runtime *util.RuntimeOptions) (_result *ExportOssResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Bucket)) {
		query["Bucket"] = request.Bucket
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MaxScore)) {
		query["MaxScore"] = request.MaxScore
	}

	if !tea.BoolValue(util.IsUnset(request.MinScore)) {
		query["MinScore"] = request.MinScore
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Stock)) {
		query["Stock"] = request.Stock
	}

	if !tea.BoolValue(util.IsUnset(request.StockTaskId)) {
		query["StockTaskId"] = request.StockTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Suggestion)) {
		query["Suggestion"] = request.Suggestion
	}

	if !tea.BoolValue(util.IsUnset(request.TotalCount)) {
		query["TotalCount"] = request.TotalCount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportOssResult"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportOssResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportOssResult(request *ExportOssResultRequest) (_result *ExportOssResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportOssResultResponse{}
	_body, _err := client.ExportOssResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuditItemDetailWithOptions(request *GetAuditItemDetailRequest, runtime *util.RuntimeOptions) (_result *GetAuditItemDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuditItemDetail"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuditItemDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuditItemDetail(request *GetAuditItemDetailRequest) (_result *GetAuditItemDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuditItemDetailResponse{}
	_body, _err := client.GetAuditItemDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuditItemListWithOptions(request *GetAuditItemListRequest, runtime *util.RuntimeOptions) (_result *GetAuditItemListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuditItemList"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuditItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuditItemList(request *GetAuditItemListRequest) (_result *GetAuditItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuditItemListResponse{}
	_body, _err := client.GetAuditItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuditUserConfWithOptions(runtime *util.RuntimeOptions) (_result *GetAuditUserConfResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAuditUserConf"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuditUserConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuditUserConf() (_result *GetAuditUserConfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuditUserConfResponse{}
	_body, _err := client.GetAuditUserConfWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportKeywordsWithOptions(request *ImportKeywordsRequest, runtime *util.RuntimeOptions) (_result *ImportKeywordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeywordLibId)) {
		query["KeywordLibId"] = request.KeywordLibId
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsObject)) {
		query["KeywordsObject"] = request.KeywordsObject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportKeywords"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportKeywordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportKeywords(request *ImportKeywordsRequest) (_result *ImportKeywordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeywordsResponse{}
	_body, _err := client.ImportKeywordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MarkAuditContentWithOptions(request *MarkAuditContentRequest, runtime *util.RuntimeOptions) (_result *MarkAuditContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditIllegalReasons)) {
		query["AuditIllegalReasons"] = request.AuditIllegalReasons
	}

	if !tea.BoolValue(util.IsUnset(request.AuditResult)) {
		query["AuditResult"] = request.AuditResult
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		query["BizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MarkAuditContent"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MarkAuditContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MarkAuditContent(request *MarkAuditContentRequest) (_result *MarkAuditContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MarkAuditContentResponse{}
	_body, _err := client.MarkAuditContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MarkAuditContentItemWithOptions(request *MarkAuditContentItemRequest, runtime *util.RuntimeOptions) (_result *MarkAuditContentItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditIllegalReasons)) {
		query["AuditIllegalReasons"] = request.AuditIllegalReasons
	}

	if !tea.BoolValue(util.IsUnset(request.AuditResult)) {
		query["AuditResult"] = request.AuditResult
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		query["BizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MarkAuditContentItem"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MarkAuditContentItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MarkAuditContentItem(request *MarkAuditContentItemRequest) (_result *MarkAuditContentItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MarkAuditContentItemResponse{}
	_body, _err := client.MarkAuditContentItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MarkOssResultWithOptions(request *MarkOssResultRequest, runtime *util.RuntimeOptions) (_result *MarkOssResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Stock)) {
		query["Stock"] = request.Stock
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MarkOssResult"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MarkOssResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MarkOssResult(request *MarkOssResultRequest) (_result *MarkOssResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MarkOssResultResponse{}
	_body, _err := client.MarkOssResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MarkWebsiteScanResultWithOptions(request *MarkWebsiteScanResultRequest, runtime *util.RuntimeOptions) (_result *MarkWebsiteScanResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		query["Ids"] = request.Ids
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MarkWebsiteScanResult"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MarkWebsiteScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MarkWebsiteScanResult(request *MarkWebsiteScanResultRequest) (_result *MarkWebsiteScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MarkWebsiteScanResultResponse{}
	_body, _err := client.MarkWebsiteScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundCdiBagWithOptions(request *RefundCdiBagRequest, runtime *util.RuntimeOptions) (_result *RefundCdiBagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefundCdiBag"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefundCdiBagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundCdiBag(request *RefundCdiBagRequest) (_result *RefundCdiBagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundCdiBagResponse{}
	_body, _err := client.RefundCdiBagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundCdiBaseBagWithOptions(request *RefundCdiBaseBagRequest, runtime *util.RuntimeOptions) (_result *RefundCdiBaseBagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefundCdiBaseBag"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefundCdiBaseBagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundCdiBaseBag(request *RefundCdiBaseBagRequest) (_result *RefundCdiBaseBagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundCdiBaseBagResponse{}
	_body, _err := client.RefundCdiBaseBagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundWebSiteInstanceWithOptions(request *RefundWebSiteInstanceRequest, runtime *util.RuntimeOptions) (_result *RefundWebSiteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefundWebSiteInstance"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefundWebSiteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundWebSiteInstance(request *RefundWebSiteInstanceRequest) (_result *RefundWebSiteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundWebSiteInstanceResponse{}
	_body, _err := client.RefundWebSiteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewWebSiteInstanceWithOptions(request *RenewWebSiteInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewWebSiteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		query["OrderNum"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewWebSiteInstance"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewWebSiteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewWebSiteInstance(request *RenewWebSiteInstanceRequest) (_result *RenewWebSiteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewWebSiteInstanceResponse{}
	_body, _err := client.RenewWebSiteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVerifyCodeToEmailWithOptions(request *SendVerifyCodeToEmailRequest, runtime *util.RuntimeOptions) (_result *SendVerifyCodeToEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerifyCodeToEmail"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerifyCodeToEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVerifyCodeToEmail(request *SendVerifyCodeToEmailRequest) (_result *SendVerifyCodeToEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerifyCodeToEmailResponse{}
	_body, _err := client.SendVerifyCodeToEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVerifyCodeToPhoneWithOptions(request *SendVerifyCodeToPhoneRequest, runtime *util.RuntimeOptions) (_result *SendVerifyCodeToPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerifyCodeToPhone"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerifyCodeToPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVerifyCodeToPhone(request *SendVerifyCodeToPhoneRequest) (_result *SendVerifyCodeToPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerifyCodeToPhoneResponse{}
	_body, _err := client.SendVerifyCodeToPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendWebsiteFeedbackWithOptions(request *SendWebsiteFeedbackRequest, runtime *util.RuntimeOptions) (_result *SendWebsiteFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		query["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Urls)) {
		query["Urls"] = request.Urls
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendWebsiteFeedback"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendWebsiteFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendWebsiteFeedback(request *SendWebsiteFeedbackRequest) (_result *SendWebsiteFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendWebsiteFeedbackResponse{}
	_body, _err := client.SendWebsiteFeedbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppPackageWithOptions(request *UpdateAppPackageRequest, runtime *util.RuntimeOptions) (_result *UpdateAppPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Debug)) {
		query["Debug"] = request.Debug
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PackageUrl)) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppPackage"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppPackage(request *UpdateAppPackageRequest) (_result *UpdateAppPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppPackageResponse{}
	_body, _err := client.UpdateAppPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAuditCallbackWithOptions(request *UpdateAuditCallbackRequest, runtime *util.RuntimeOptions) (_result *UpdateAuditCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.CryptType)) {
		query["CryptType"] = request.CryptType
	}

	if !tea.BoolValue(util.IsUnset(request.Seed)) {
		query["Seed"] = request.Seed
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAuditCallback"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuditCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAuditCallback(request *UpdateAuditCallbackRequest) (_result *UpdateAuditCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAuditCallbackResponse{}
	_body, _err := client.UpdateAuditCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAuditRangeWithOptions(request *UpdateAuditRangeRequest, runtime *util.RuntimeOptions) (_result *UpdateAuditRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditRange)) {
		query["AuditRange"] = request.AuditRange
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAuditRange"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuditRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAuditRange(request *UpdateAuditRangeRequest) (_result *UpdateAuditRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAuditRangeResponse{}
	_body, _err := client.UpdateAuditRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAuditSettingWithOptions(request *UpdateAuditSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateAuditSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditRange)) {
		query["AuditRange"] = request.AuditRange
	}

	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.Seed)) {
		query["Seed"] = request.Seed
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAuditSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuditSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAuditSetting(request *UpdateAuditSettingRequest) (_result *UpdateAuditSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAuditSettingResponse{}
	_body, _err := client.UpdateAuditSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBizTypeWithOptions(request *UpdateBizTypeRequest, runtime *util.RuntimeOptions) (_result *UpdateBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBizType"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBizTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBizType(request *UpdateBizTypeRequest) (_result *UpdateBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBizTypeResponse{}
	_body, _err := client.UpdateBizTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBizTypeImageLibWithOptions(request *UpdateBizTypeImageLibRequest, runtime *util.RuntimeOptions) (_result *UpdateBizTypeImageLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.Black)) {
		query["Black"] = request.Black
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Review)) {
		query["Review"] = request.Review
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.White)) {
		query["White"] = request.White
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBizTypeImageLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBizTypeImageLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBizTypeImageLib(request *UpdateBizTypeImageLibRequest) (_result *UpdateBizTypeImageLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBizTypeImageLibResponse{}
	_body, _err := client.UpdateBizTypeImageLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBizTypeSettingWithOptions(request *UpdateBizTypeSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateBizTypeSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ad)) {
		query["Ad"] = request.Ad
	}

	if !tea.BoolValue(util.IsUnset(request.Antispam)) {
		query["Antispam"] = request.Antispam
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.Live)) {
		query["Live"] = request.Live
	}

	if !tea.BoolValue(util.IsUnset(request.Porn)) {
		query["Porn"] = request.Porn
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Terrorism)) {
		query["Terrorism"] = request.Terrorism
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBizTypeSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBizTypeSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBizTypeSetting(request *UpdateBizTypeSettingRequest) (_result *UpdateBizTypeSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBizTypeSettingResponse{}
	_body, _err := client.UpdateBizTypeSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBizTypeTextLibWithOptions(request *UpdateBizTypeTextLibRequest, runtime *util.RuntimeOptions) (_result *UpdateBizTypeTextLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypeName)) {
		query["BizTypeName"] = request.BizTypeName
	}

	if !tea.BoolValue(util.IsUnset(request.Black)) {
		query["Black"] = request.Black
	}

	if !tea.BoolValue(util.IsUnset(request.Ignore)) {
		query["Ignore"] = request.Ignore
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Review)) {
		query["Review"] = request.Review
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		query["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.White)) {
		query["White"] = request.White
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBizTypeTextLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBizTypeTextLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBizTypeTextLib(request *UpdateBizTypeTextLibRequest) (_result *UpdateBizTypeTextLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateBizTypeTextLibResponse{}
	_body, _err := client.UpdateBizTypeTextLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomOcrTemplateWithOptions(request *UpdateCustomOcrTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomOcrTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RecognizeArea)) {
		query["RecognizeArea"] = request.RecognizeArea
	}

	if !tea.BoolValue(util.IsUnset(request.ReferArea)) {
		query["ReferArea"] = request.ReferArea
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomOcrTemplate"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomOcrTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomOcrTemplate(request *UpdateCustomOcrTemplateRequest) (_result *UpdateCustomOcrTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomOcrTemplateResponse{}
	_body, _err := client.UpdateCustomOcrTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKeywordLibWithOptions(request *UpdateKeywordLibRequest, runtime *util.RuntimeOptions) (_result *UpdateKeywordLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		query["BizTypes"] = request.BizTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MatchMode)) {
		query["MatchMode"] = request.MatchMode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateKeywordLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateKeywordLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKeywordLib(request *UpdateKeywordLibRequest) (_result *UpdateKeywordLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKeywordLibResponse{}
	_body, _err := client.UpdateKeywordLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateNotificationSettingWithOptions(request *UpdateNotificationSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateNotificationSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RealtimeMessageList)) {
		query["RealtimeMessageList"] = request.RealtimeMessageList
	}

	if !tea.BoolValue(util.IsUnset(request.ReminderModeList)) {
		query["ReminderModeList"] = request.ReminderModeList
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleMessageTime)) {
		query["ScheduleMessageTime"] = request.ScheduleMessageTime
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleMessageTimeZone)) {
		query["ScheduleMessageTimeZone"] = request.ScheduleMessageTimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNotificationSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNotificationSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateNotificationSetting(request *UpdateNotificationSettingRequest) (_result *UpdateNotificationSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNotificationSettingResponse{}
	_body, _err := client.UpdateNotificationSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOssCallbackSettingWithOptions(request *UpdateOssCallbackSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateOssCallbackSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditCallback)) {
		query["AuditCallback"] = request.AuditCallback
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackSeed)) {
		query["CallbackSeed"] = request.CallbackSeed
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ScanCallback)) {
		query["ScanCallback"] = request.ScanCallback
	}

	if !tea.BoolValue(util.IsUnset(request.ScanCallbackSuggestions)) {
		query["ScanCallbackSuggestions"] = request.ScanCallbackSuggestions
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceModules)) {
		query["ServiceModules"] = request.ServiceModules
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOssCallbackSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOssCallbackSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOssCallbackSetting(request *UpdateOssCallbackSettingRequest) (_result *UpdateOssCallbackSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOssCallbackSettingResponse{}
	_body, _err := client.UpdateOssCallbackSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOssIncrementCheckSettingWithOptions(request *UpdateOssIncrementCheckSettingRequest, runtime *util.RuntimeOptions) (_result *UpdateOssIncrementCheckSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AudioAntispamFreezeConfig)) {
		query["AudioAntispamFreezeConfig"] = request.AudioAntispamFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.AudioAutoFreezeOpened)) {
		query["AudioAutoFreezeOpened"] = request.AudioAutoFreezeOpened
	}

	if !tea.BoolValue(util.IsUnset(request.AudioMaxSize)) {
		query["AudioMaxSize"] = request.AudioMaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.AudioScanLimit)) {
		query["AudioScanLimit"] = request.AudioScanLimit
	}

	if !tea.BoolValue(util.IsUnset(request.AudioSceneList)) {
		query["AudioSceneList"] = request.AudioSceneList
	}

	if !tea.BoolValue(util.IsUnset(request.AutoFreezeType)) {
		query["AutoFreezeType"] = request.AutoFreezeType
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.BucketConfigList)) {
		query["BucketConfigList"] = request.BucketConfigList
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackId)) {
		query["CallbackId"] = request.CallbackId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAdFreezeConfig)) {
		query["ImageAdFreezeConfig"] = request.ImageAdFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAutoFreeze)) {
		query["ImageAutoFreeze"] = request.ImageAutoFreeze
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAutoFreezeOpened)) {
		query["ImageAutoFreezeOpened"] = request.ImageAutoFreezeOpened
	}

	if !tea.BoolValue(util.IsUnset(request.ImageLiveFreezeConfig)) {
		query["ImageLiveFreezeConfig"] = request.ImageLiveFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ImagePornFreezeConfig)) {
		query["ImagePornFreezeConfig"] = request.ImagePornFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ImageScanLimit)) {
		query["ImageScanLimit"] = request.ImageScanLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ImageSceneList)) {
		query["ImageSceneList"] = request.ImageSceneList
	}

	if !tea.BoolValue(util.IsUnset(request.ImageTerrorismFreezeConfig)) {
		query["ImageTerrorismFreezeConfig"] = request.ImageTerrorismFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ScanImageNoFileType)) {
		query["ScanImageNoFileType"] = request.ScanImageNoFileType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.VideoAdFreezeConfig)) {
		query["VideoAdFreezeConfig"] = request.VideoAdFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.VideoAutoFreezeOpened)) {
		query["VideoAutoFreezeOpened"] = request.VideoAutoFreezeOpened
	}

	if !tea.BoolValue(util.IsUnset(request.VideoAutoFreezeSceneList)) {
		query["VideoAutoFreezeSceneList"] = request.VideoAutoFreezeSceneList
	}

	if !tea.BoolValue(util.IsUnset(request.VideoFrameInterval)) {
		query["VideoFrameInterval"] = request.VideoFrameInterval
	}

	if !tea.BoolValue(util.IsUnset(request.VideoLiveFreezeConfig)) {
		query["VideoLiveFreezeConfig"] = request.VideoLiveFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.VideoMaxFrames)) {
		query["VideoMaxFrames"] = request.VideoMaxFrames
	}

	if !tea.BoolValue(util.IsUnset(request.VideoMaxSize)) {
		query["VideoMaxSize"] = request.VideoMaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.VideoPornFreezeConfig)) {
		query["VideoPornFreezeConfig"] = request.VideoPornFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.VideoScanLimit)) {
		query["VideoScanLimit"] = request.VideoScanLimit
	}

	if !tea.BoolValue(util.IsUnset(request.VideoSceneList)) {
		query["VideoSceneList"] = request.VideoSceneList
	}

	if !tea.BoolValue(util.IsUnset(request.VideoTerrorismFreezeConfig)) {
		query["VideoTerrorismFreezeConfig"] = request.VideoTerrorismFreezeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.VideoVoiceAntispamFreezeConfig)) {
		query["VideoVoiceAntispamFreezeConfig"] = request.VideoVoiceAntispamFreezeConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOssIncrementCheckSetting"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOssIncrementCheckSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOssIncrementCheckSetting(request *UpdateOssIncrementCheckSettingRequest) (_result *UpdateOssIncrementCheckSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOssIncrementCheckSettingResponse{}
	_body, _err := client.UpdateOssIncrementCheckSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateOssStockStatusWithOptions(request *UpdateOssStockStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateOssStockStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AudioAutoFreezeSceneList)) {
		query["AudioAutoFreezeSceneList"] = request.AudioAutoFreezeSceneList
	}

	if !tea.BoolValue(util.IsUnset(request.AudioMaxSize)) {
		query["AudioMaxSize"] = request.AudioMaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.AutoFreezeType)) {
		query["AutoFreezeType"] = request.AutoFreezeType
	}

	if !tea.BoolValue(util.IsUnset(request.BucketConfigList)) {
		query["BucketConfigList"] = request.BucketConfigList
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAutoFreeze)) {
		query["ImageAutoFreeze"] = request.ImageAutoFreeze
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypeList)) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.SceneList)) {
		query["SceneList"] = request.SceneList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoAutoFreezeSceneList)) {
		query["VideoAutoFreezeSceneList"] = request.VideoAutoFreezeSceneList
	}

	if !tea.BoolValue(util.IsUnset(request.VideoFrameInterval)) {
		query["VideoFrameInterval"] = request.VideoFrameInterval
	}

	if !tea.BoolValue(util.IsUnset(request.VideoMaxFrames)) {
		query["VideoMaxFrames"] = request.VideoMaxFrames
	}

	if !tea.BoolValue(util.IsUnset(request.VideoMaxSize)) {
		query["VideoMaxSize"] = request.VideoMaxSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOssStockStatus"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOssStockStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateOssStockStatus(request *UpdateOssStockStatusRequest) (_result *UpdateOssStockStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOssStockStatusResponse{}
	_body, _err := client.UpdateOssStockStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWebsiteInstanceWithOptions(request *UpdateWebsiteInstanceRequest, runtime *util.RuntimeOptions) (_result *UpdateWebsiteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.IndexPage)) {
		query["IndexPage"] = request.IndexPage
	}

	if !tea.BoolValue(util.IsUnset(request.IndexPageScanInterval)) {
		query["IndexPageScanInterval"] = request.IndexPageScanInterval
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SiteProtocol)) {
		query["SiteProtocol"] = request.SiteProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.WebsiteScanInterval)) {
		query["WebsiteScanInterval"] = request.WebsiteScanInterval
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWebsiteInstance"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWebsiteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWebsiteInstance(request *UpdateWebsiteInstanceRequest) (_result *UpdateWebsiteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWebsiteInstanceResponse{}
	_body, _err := client.UpdateWebsiteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWebsiteInstanceKeyUrlWithOptions(request *UpdateWebsiteInstanceKeyUrlRequest, runtime *util.RuntimeOptions) (_result *UpdateWebsiteInstanceKeyUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Urls)) {
		query["Urls"] = request.Urls
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWebsiteInstanceKeyUrl"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWebsiteInstanceKeyUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWebsiteInstanceKeyUrl(request *UpdateWebsiteInstanceKeyUrlRequest) (_result *UpdateWebsiteInstanceKeyUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWebsiteInstanceKeyUrlResponse{}
	_body, _err := client.UpdateWebsiteInstanceKeyUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWebsiteInstanceStatusWithOptions(request *UpdateWebsiteInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateWebsiteInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWebsiteInstanceStatus"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWebsiteInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWebsiteInstanceStatus(request *UpdateWebsiteInstanceStatusRequest) (_result *UpdateWebsiteInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWebsiteInstanceStatusResponse{}
	_body, _err := client.UpdateWebsiteInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeCdiBaseBagWithOptions(request *UpgradeCdiBaseBagRequest, runtime *util.RuntimeOptions) (_result *UpgradeCdiBaseBagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.FlowOutSpec)) {
		query["FlowOutSpec"] = request.FlowOutSpec
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeCdiBaseBag"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeCdiBaseBagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeCdiBaseBag(request *UpgradeCdiBaseBagRequest) (_result *UpgradeCdiBaseBagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeCdiBaseBagResponse{}
	_body, _err := client.UpgradeCdiBaseBagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadImageToLibWithOptions(request *UploadImageToLibRequest, runtime *util.RuntimeOptions) (_result *UploadImageToLibResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageLibId)) {
		query["ImageLibId"] = request.ImageLibId
	}

	if !tea.BoolValue(util.IsUnset(request.Images)) {
		query["Images"] = request.Images
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Urls)) {
		query["Urls"] = request.Urls
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadImageToLib"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadImageToLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadImageToLib(request *UploadImageToLibRequest) (_result *UploadImageToLibResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadImageToLibResponse{}
	_body, _err := client.UploadImageToLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyCustomOcrTemplateWithOptions(request *VerifyCustomOcrTemplateRequest, runtime *util.RuntimeOptions) (_result *VerifyCustomOcrTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.TestImgUrl)) {
		query["TestImgUrl"] = request.TestImgUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyCustomOcrTemplate"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyCustomOcrTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyCustomOcrTemplate(request *VerifyCustomOcrTemplateRequest) (_result *VerifyCustomOcrTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyCustomOcrTemplateResponse{}
	_body, _err := client.VerifyCustomOcrTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyEmailWithOptions(request *VerifyEmailRequest, runtime *util.RuntimeOptions) (_result *VerifyEmailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyEmail"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyEmailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyEmail(request *VerifyEmailRequest) (_result *VerifyEmailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyEmailResponse{}
	_body, _err := client.VerifyEmailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyPhoneWithOptions(request *VerifyPhoneRequest, runtime *util.RuntimeOptions) (_result *VerifyPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyPhone"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyPhone(request *VerifyPhoneRequest) (_result *VerifyPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyPhoneResponse{}
	_body, _err := client.VerifyPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyWebsiteInstanceWithOptions(request *VerifyWebsiteInstanceRequest, runtime *util.RuntimeOptions) (_result *VerifyWebsiteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyMethod)) {
		query["VerifyMethod"] = request.VerifyMethod
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyWebsiteInstance"),
		Version:     tea.String("2017-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyWebsiteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyWebsiteInstance(request *VerifyWebsiteInstanceRequest) (_result *VerifyWebsiteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyWebsiteInstanceResponse{}
	_body, _err := client.VerifyWebsiteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
