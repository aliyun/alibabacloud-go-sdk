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

type AppUseTimeReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AppUseTimeReportHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportHeaders) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportHeaders) SetCommonHeaders(v map[string]*string) *AppUseTimeReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppUseTimeReportHeaders) SetXAcsAligenieAccessToken(v string) *AppUseTimeReportHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AppUseTimeReportHeaders) SetAuthorization(v string) *AppUseTimeReportHeaders {
	s.Authorization = &v
	return s
}

type AppUseTimeReportRequest struct {
	DeviceInfo *AppUseTimeReportRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *AppUseTimeReportRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *AppUseTimeReportRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s AppUseTimeReportRequest) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportRequest) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequest) SetDeviceInfo(v *AppUseTimeReportRequestDeviceInfo) *AppUseTimeReportRequest {
	s.DeviceInfo = v
	return s
}

func (s *AppUseTimeReportRequest) SetPayload(v *AppUseTimeReportRequestPayload) *AppUseTimeReportRequest {
	s.Payload = v
	return s
}

func (s *AppUseTimeReportRequest) SetUserInfo(v *AppUseTimeReportRequestUserInfo) *AppUseTimeReportRequest {
	s.UserInfo = v
	return s
}

type AppUseTimeReportRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AppUseTimeReportRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequestDeviceInfo) SetEncodeKey(v string) *AppUseTimeReportRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetEncodeType(v string) *AppUseTimeReportRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetId(v string) *AppUseTimeReportRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetIdType(v string) *AppUseTimeReportRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *AppUseTimeReportRequestDeviceInfo) SetOrganizationId(v string) *AppUseTimeReportRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type AppUseTimeReportRequestPayload struct {
	Action       *string `json:"Action,omitempty" xml:"Action,omitempty"`
	IsPrivilege  *int32  `json:"IsPrivilege,omitempty" xml:"IsPrivilege,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StepCode     *string `json:"StepCode,omitempty" xml:"StepCode,omitempty"`
	VipType      *int32  `json:"VipType,omitempty" xml:"VipType,omitempty"`
	OriginUuid   *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s AppUseTimeReportRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportRequestPayload) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequestPayload) SetAction(v string) *AppUseTimeReportRequestPayload {
	s.Action = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetIsPrivilege(v int32) *AppUseTimeReportRequestPayload {
	s.IsPrivilege = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetResourceId(v string) *AppUseTimeReportRequestPayload {
	s.ResourceId = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetResourceType(v int32) *AppUseTimeReportRequestPayload {
	s.ResourceType = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetStepCode(v string) *AppUseTimeReportRequestPayload {
	s.StepCode = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetVipType(v int32) *AppUseTimeReportRequestPayload {
	s.VipType = &v
	return s
}

func (s *AppUseTimeReportRequestPayload) SetOriginUuid(v string) *AppUseTimeReportRequestPayload {
	s.OriginUuid = &v
	return s
}

type AppUseTimeReportRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s AppUseTimeReportRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportRequestUserInfo) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportRequestUserInfo) SetEncodeKey(v string) *AppUseTimeReportRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetEncodeType(v string) *AppUseTimeReportRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetId(v string) *AppUseTimeReportRequestUserInfo {
	s.Id = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetIdType(v string) *AppUseTimeReportRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *AppUseTimeReportRequestUserInfo) SetOrganizationId(v string) *AppUseTimeReportRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type AppUseTimeReportShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s AppUseTimeReportShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportShrinkRequest) SetDeviceInfoShrink(v string) *AppUseTimeReportShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *AppUseTimeReportShrinkRequest) SetPayloadShrink(v string) *AppUseTimeReportShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *AppUseTimeReportShrinkRequest) SetUserInfoShrink(v string) *AppUseTimeReportShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type AppUseTimeReportResponseBody struct {
	RetCode  *int32  `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg   *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *bool   `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
}

func (s AppUseTimeReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportResponseBody) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportResponseBody) SetRetCode(v int32) *AppUseTimeReportResponseBody {
	s.RetCode = &v
	return s
}

func (s *AppUseTimeReportResponseBody) SetRetMsg(v string) *AppUseTimeReportResponseBody {
	s.RetMsg = &v
	return s
}

func (s *AppUseTimeReportResponseBody) SetRetValue(v bool) *AppUseTimeReportResponseBody {
	s.RetValue = &v
	return s
}

type AppUseTimeReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppUseTimeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppUseTimeReportResponse) String() string {
	return tea.Prettify(s)
}

func (s AppUseTimeReportResponse) GoString() string {
	return s.String()
}

func (s *AppUseTimeReportResponse) SetHeaders(v map[string]*string) *AppUseTimeReportResponse {
	s.Headers = v
	return s
}

func (s *AppUseTimeReportResponse) SetStatusCode(v int32) *AppUseTimeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *AppUseTimeReportResponse) SetBody(v *AppUseTimeReportResponseBody) *AppUseTimeReportResponse {
	s.Body = v
	return s
}

type CallBackThirdRightSendPlanHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CallBackThirdRightSendPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s CallBackThirdRightSendPlanHeaders) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanHeaders) SetCommonHeaders(v map[string]*string) *CallBackThirdRightSendPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CallBackThirdRightSendPlanHeaders) SetXAcsAligenieAccessToken(v string) *CallBackThirdRightSendPlanHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CallBackThirdRightSendPlanHeaders) SetAuthorization(v string) *CallBackThirdRightSendPlanHeaders {
	s.Authorization = &v
	return s
}

type CallBackThirdRightSendPlanRequest struct {
	BizGroup      *string                `json:"BizGroup,omitempty" xml:"BizGroup,omitempty"`
	BizType       *string                `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CardType      *int32                 `json:"CardType,omitempty" xml:"CardType,omitempty"`
	ErrorMsg      *string                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	ExtendInfo    map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	GenieOpenId   *string                `json:"GenieOpenId,omitempty" xml:"GenieOpenId,omitempty"`
	ReceiveStatus *int32                 `json:"ReceiveStatus,omitempty" xml:"ReceiveStatus,omitempty"`
	Sn            *string                `json:"Sn,omitempty" xml:"Sn,omitempty"`
	SupplierId    *int64                 `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
}

func (s CallBackThirdRightSendPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CallBackThirdRightSendPlanRequest) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanRequest) SetBizGroup(v string) *CallBackThirdRightSendPlanRequest {
	s.BizGroup = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetBizType(v string) *CallBackThirdRightSendPlanRequest {
	s.BizType = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetCardType(v int32) *CallBackThirdRightSendPlanRequest {
	s.CardType = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetErrorMsg(v string) *CallBackThirdRightSendPlanRequest {
	s.ErrorMsg = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetExtendInfo(v map[string]interface{}) *CallBackThirdRightSendPlanRequest {
	s.ExtendInfo = v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetGenieOpenId(v string) *CallBackThirdRightSendPlanRequest {
	s.GenieOpenId = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetReceiveStatus(v int32) *CallBackThirdRightSendPlanRequest {
	s.ReceiveStatus = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetSn(v string) *CallBackThirdRightSendPlanRequest {
	s.Sn = &v
	return s
}

func (s *CallBackThirdRightSendPlanRequest) SetSupplierId(v int64) *CallBackThirdRightSendPlanRequest {
	s.SupplierId = &v
	return s
}

type CallBackThirdRightSendPlanShrinkRequest struct {
	BizGroup         *string `json:"BizGroup,omitempty" xml:"BizGroup,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CardType         *int32  `json:"CardType,omitempty" xml:"CardType,omitempty"`
	ErrorMsg         *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	ExtendInfoShrink *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	GenieOpenId      *string `json:"GenieOpenId,omitempty" xml:"GenieOpenId,omitempty"`
	ReceiveStatus    *int32  `json:"ReceiveStatus,omitempty" xml:"ReceiveStatus,omitempty"`
	Sn               *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	SupplierId       *int64  `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
}

func (s CallBackThirdRightSendPlanShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CallBackThirdRightSendPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetBizGroup(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.BizGroup = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetBizType(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetCardType(v int32) *CallBackThirdRightSendPlanShrinkRequest {
	s.CardType = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetErrorMsg(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.ErrorMsg = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetExtendInfoShrink(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.ExtendInfoShrink = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetGenieOpenId(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.GenieOpenId = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetReceiveStatus(v int32) *CallBackThirdRightSendPlanShrinkRequest {
	s.ReceiveStatus = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetSn(v string) *CallBackThirdRightSendPlanShrinkRequest {
	s.Sn = &v
	return s
}

func (s *CallBackThirdRightSendPlanShrinkRequest) SetSupplierId(v int64) *CallBackThirdRightSendPlanShrinkRequest {
	s.SupplierId = &v
	return s
}

type CallBackThirdRightSendPlanResponseBody struct {
	RetCode   *string `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg    *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue  *bool   `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CallBackThirdRightSendPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CallBackThirdRightSendPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRetCode(v string) *CallBackThirdRightSendPlanResponseBody {
	s.RetCode = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRetMsg(v string) *CallBackThirdRightSendPlanResponseBody {
	s.RetMsg = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRetValue(v bool) *CallBackThirdRightSendPlanResponseBody {
	s.RetValue = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponseBody) SetRequestId(v string) *CallBackThirdRightSendPlanResponseBody {
	s.RequestId = &v
	return s
}

type CallBackThirdRightSendPlanResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallBackThirdRightSendPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallBackThirdRightSendPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CallBackThirdRightSendPlanResponse) GoString() string {
	return s.String()
}

func (s *CallBackThirdRightSendPlanResponse) SetHeaders(v map[string]*string) *CallBackThirdRightSendPlanResponse {
	s.Headers = v
	return s
}

func (s *CallBackThirdRightSendPlanResponse) SetStatusCode(v int32) *CallBackThirdRightSendPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CallBackThirdRightSendPlanResponse) SetBody(v *CallBackThirdRightSendPlanResponseBody) *CallBackThirdRightSendPlanResponse {
	s.Body = v
	return s
}

type CheckThirdRightSendPlanHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CheckThirdRightSendPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckThirdRightSendPlanHeaders) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanHeaders) SetCommonHeaders(v map[string]*string) *CheckThirdRightSendPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckThirdRightSendPlanHeaders) SetXAcsAligenieAccessToken(v string) *CheckThirdRightSendPlanHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CheckThirdRightSendPlanHeaders) SetAuthorization(v string) *CheckThirdRightSendPlanHeaders {
	s.Authorization = &v
	return s
}

type CheckThirdRightSendPlanRequest struct {
	BizGroup   *string                `json:"BizGroup,omitempty" xml:"BizGroup,omitempty"`
	BizType    *string                `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExtendInfo map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	Sn         *string                `json:"Sn,omitempty" xml:"Sn,omitempty"`
	SupplierId *int64                 `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
}

func (s CheckThirdRightSendPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckThirdRightSendPlanRequest) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanRequest) SetBizGroup(v string) *CheckThirdRightSendPlanRequest {
	s.BizGroup = &v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetBizType(v string) *CheckThirdRightSendPlanRequest {
	s.BizType = &v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetExtendInfo(v map[string]interface{}) *CheckThirdRightSendPlanRequest {
	s.ExtendInfo = v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetSn(v string) *CheckThirdRightSendPlanRequest {
	s.Sn = &v
	return s
}

func (s *CheckThirdRightSendPlanRequest) SetSupplierId(v int64) *CheckThirdRightSendPlanRequest {
	s.SupplierId = &v
	return s
}

type CheckThirdRightSendPlanShrinkRequest struct {
	BizGroup         *string `json:"BizGroup,omitempty" xml:"BizGroup,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExtendInfoShrink *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	Sn               *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	SupplierId       *int64  `json:"SupplierId,omitempty" xml:"SupplierId,omitempty"`
}

func (s CheckThirdRightSendPlanShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckThirdRightSendPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetBizGroup(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.BizGroup = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetBizType(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.BizType = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetExtendInfoShrink(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.ExtendInfoShrink = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetSn(v string) *CheckThirdRightSendPlanShrinkRequest {
	s.Sn = &v
	return s
}

func (s *CheckThirdRightSendPlanShrinkRequest) SetSupplierId(v int64) *CheckThirdRightSendPlanShrinkRequest {
	s.SupplierId = &v
	return s
}

type CheckThirdRightSendPlanResponseBody struct {
	RetCode  *int32                                       `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg   *string                                      `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *CheckThirdRightSendPlanResponseBodyRetValue `json:"RetValue,omitempty" xml:"RetValue,omitempty" type:"Struct"`
}

func (s CheckThirdRightSendPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckThirdRightSendPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanResponseBody) SetRetCode(v int32) *CheckThirdRightSendPlanResponseBody {
	s.RetCode = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBody) SetRetMsg(v string) *CheckThirdRightSendPlanResponseBody {
	s.RetMsg = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBody) SetRetValue(v *CheckThirdRightSendPlanResponseBodyRetValue) *CheckThirdRightSendPlanResponseBody {
	s.RetValue = v
	return s
}

type CheckThirdRightSendPlanResponseBodyRetValue struct {
	ActivateDate      *string                `json:"ActivateDate,omitempty" xml:"ActivateDate,omitempty"`
	CardType          *int32                 `json:"CardType,omitempty" xml:"CardType,omitempty"`
	ChannelCode       *string                `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	ChannelName       *string                `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	ExtendInfo        map[string]interface{} `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	RequestId         *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RightsExpiredDate *string                `json:"RightsExpiredDate,omitempty" xml:"RightsExpiredDate,omitempty"`
}

func (s CheckThirdRightSendPlanResponseBodyRetValue) String() string {
	return tea.Prettify(s)
}

func (s CheckThirdRightSendPlanResponseBodyRetValue) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetActivateDate(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ActivateDate = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetCardType(v int32) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.CardType = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetChannelCode(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ChannelCode = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetChannelName(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ChannelName = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetExtendInfo(v map[string]interface{}) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.ExtendInfo = v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetRequestId(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.RequestId = &v
	return s
}

func (s *CheckThirdRightSendPlanResponseBodyRetValue) SetRightsExpiredDate(v string) *CheckThirdRightSendPlanResponseBodyRetValue {
	s.RightsExpiredDate = &v
	return s
}

type CheckThirdRightSendPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckThirdRightSendPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckThirdRightSendPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckThirdRightSendPlanResponse) GoString() string {
	return s.String()
}

func (s *CheckThirdRightSendPlanResponse) SetHeaders(v map[string]*string) *CheckThirdRightSendPlanResponse {
	s.Headers = v
	return s
}

func (s *CheckThirdRightSendPlanResponse) SetStatusCode(v int32) *CheckThirdRightSendPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckThirdRightSendPlanResponse) SetBody(v *CheckThirdRightSendPlanResponseBody) *CheckThirdRightSendPlanResponse {
	s.Body = v
	return s
}

type CreateReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateReminderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderHeaders) GoString() string {
	return s.String()
}

func (s *CreateReminderHeaders) SetCommonHeaders(v map[string]*string) *CreateReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateReminderHeaders) SetXAcsAligenieAccessToken(v string) *CreateReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateReminderHeaders) SetAuthorization(v string) *CreateReminderHeaders {
	s.Authorization = &v
	return s
}

type CreateReminderRequest struct {
	DeviceInfo *CreateReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *CreateReminderRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *CreateReminderRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateReminderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderRequest) GoString() string {
	return s.String()
}

func (s *CreateReminderRequest) SetDeviceInfo(v *CreateReminderRequestDeviceInfo) *CreateReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *CreateReminderRequest) SetPayload(v *CreateReminderRequestPayload) *CreateReminderRequest {
	s.Payload = v
	return s
}

func (s *CreateReminderRequest) SetUserInfo(v *CreateReminderRequestUserInfo) *CreateReminderRequest {
	s.UserInfo = v
	return s
}

type CreateReminderRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateReminderRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestDeviceInfo) SetEncodeKey(v string) *CreateReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetEncodeType(v string) *CreateReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetId(v string) *CreateReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetIdType(v string) *CreateReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *CreateReminderRequestDeviceInfo) SetOrganizationId(v string) *CreateReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type CreateReminderRequestPayload struct {
	Content        *string                                     `json:"Content,omitempty" xml:"Content,omitempty"`
	IsDebug        *bool                                       `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	RecurrenceRule *CreateReminderRequestPayloadRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
}

func (s CreateReminderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestPayload) SetContent(v string) *CreateReminderRequestPayload {
	s.Content = &v
	return s
}

func (s *CreateReminderRequestPayload) SetIsDebug(v bool) *CreateReminderRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *CreateReminderRequestPayload) SetRecurrenceRule(v *CreateReminderRequestPayloadRecurrenceRule) *CreateReminderRequestPayload {
	s.RecurrenceRule = v
	return s
}

type CreateReminderRequestPayloadRecurrenceRule struct {
	Day           *int32   `json:"Day,omitempty" xml:"Day,omitempty"`
	DaysOfMonth   []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	DaysOfWeek    []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	EndDateTime   *int64   `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	Freq          *string  `json:"Freq,omitempty" xml:"Freq,omitempty"`
	Hour          *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute        *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month         *int32   `json:"Month,omitempty" xml:"Month,omitempty"`
	Second        *int32   `json:"Second,omitempty" xml:"Second,omitempty"`
	StartDateTime *int64   `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
	Year          *int32   `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s CreateReminderRequestPayloadRecurrenceRule) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderRequestPayloadRecurrenceRule) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetDay(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Day = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetDaysOfMonth(v []*int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetDaysOfWeek(v []*int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetEndDateTime(v int64) *CreateReminderRequestPayloadRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetFreq(v string) *CreateReminderRequestPayloadRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetHour(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetMinute(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetMonth(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Month = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetSecond(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Second = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetStartDateTime(v int64) *CreateReminderRequestPayloadRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *CreateReminderRequestPayloadRecurrenceRule) SetYear(v int32) *CreateReminderRequestPayloadRecurrenceRule {
	s.Year = &v
	return s
}

type CreateReminderRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateReminderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateReminderRequestUserInfo) SetEncodeKey(v string) *CreateReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetEncodeType(v string) *CreateReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetId(v string) *CreateReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetIdType(v string) *CreateReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *CreateReminderRequestUserInfo) SetOrganizationId(v string) *CreateReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type CreateReminderShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateReminderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateReminderShrinkRequest) SetDeviceInfoShrink(v string) *CreateReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *CreateReminderShrinkRequest) SetPayloadShrink(v string) *CreateReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *CreateReminderShrinkRequest) SetUserInfoShrink(v string) *CreateReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type CreateReminderResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Model     *int64  `json:"Model,omitempty" xml:"Model,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateReminderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReminderResponseBody) SetErrorCode(v string) *CreateReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateReminderResponseBody) SetErrorMsg(v string) *CreateReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateReminderResponseBody) SetModel(v int64) *CreateReminderResponseBody {
	s.Model = &v
	return s
}

func (s *CreateReminderResponseBody) SetSuccess(v bool) *CreateReminderResponseBody {
	s.Success = &v
	return s
}

type CreateReminderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReminderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReminderResponse) GoString() string {
	return s.String()
}

func (s *CreateReminderResponse) SetHeaders(v map[string]*string) *CreateReminderResponse {
	s.Headers = v
	return s
}

func (s *CreateReminderResponse) SetStatusCode(v int32) *CreateReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReminderResponse) SetBody(v *CreateReminderResponseBody) *CreateReminderResponse {
	s.Body = v
	return s
}

type DeleteReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteReminderHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderHeaders) GoString() string {
	return s.String()
}

func (s *DeleteReminderHeaders) SetCommonHeaders(v map[string]*string) *DeleteReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteReminderHeaders) SetXAcsAligenieAccessToken(v string) *DeleteReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteReminderHeaders) SetAuthorization(v string) *DeleteReminderHeaders {
	s.Authorization = &v
	return s
}

type DeleteReminderRequest struct {
	DeviceInfo *DeleteReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *DeleteReminderRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *DeleteReminderRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s DeleteReminderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderRequest) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequest) SetDeviceInfo(v *DeleteReminderRequestDeviceInfo) *DeleteReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *DeleteReminderRequest) SetPayload(v *DeleteReminderRequestPayload) *DeleteReminderRequest {
	s.Payload = v
	return s
}

func (s *DeleteReminderRequest) SetUserInfo(v *DeleteReminderRequestUserInfo) *DeleteReminderRequest {
	s.UserInfo = v
	return s
}

type DeleteReminderRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteReminderRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequestDeviceInfo) SetEncodeKey(v string) *DeleteReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetEncodeType(v string) *DeleteReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetId(v string) *DeleteReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetIdType(v string) *DeleteReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *DeleteReminderRequestDeviceInfo) SetOrganizationId(v string) *DeleteReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type DeleteReminderRequestPayload struct {
	Id      *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDebug *bool  `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s DeleteReminderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequestPayload) SetId(v int64) *DeleteReminderRequestPayload {
	s.Id = &v
	return s
}

func (s *DeleteReminderRequestPayload) SetIsDebug(v bool) *DeleteReminderRequestPayload {
	s.IsDebug = &v
	return s
}

type DeleteReminderRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s DeleteReminderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *DeleteReminderRequestUserInfo) SetEncodeKey(v string) *DeleteReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetEncodeType(v string) *DeleteReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetId(v string) *DeleteReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetIdType(v string) *DeleteReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *DeleteReminderRequestUserInfo) SetOrganizationId(v string) *DeleteReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type DeleteReminderShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s DeleteReminderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteReminderShrinkRequest) SetDeviceInfoShrink(v string) *DeleteReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *DeleteReminderShrinkRequest) SetPayloadShrink(v string) *DeleteReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *DeleteReminderShrinkRequest) SetUserInfoShrink(v string) *DeleteReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type DeleteReminderResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteReminderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReminderResponseBody) SetErrorCode(v int32) *DeleteReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteReminderResponseBody) SetErrorMsg(v string) *DeleteReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteReminderResponseBody) SetSuccess(v bool) *DeleteReminderResponseBody {
	s.Success = &v
	return s
}

type DeleteReminderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReminderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteReminderResponse) GoString() string {
	return s.String()
}

func (s *DeleteReminderResponse) SetHeaders(v map[string]*string) *DeleteReminderResponse {
	s.Headers = v
	return s
}

func (s *DeleteReminderResponse) SetStatusCode(v int32) *DeleteReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReminderResponse) SetBody(v *DeleteReminderResponseBody) *DeleteReminderResponse {
	s.Body = v
	return s
}

type GetAccountForAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAccountForAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppHeaders) GoString() string {
	return s.String()
}

func (s *GetAccountForAppHeaders) SetCommonHeaders(v map[string]*string) *GetAccountForAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccountForAppHeaders) SetXAcsAligenieAccessToken(v string) *GetAccountForAppHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAccountForAppHeaders) SetAuthorization(v string) *GetAccountForAppHeaders {
	s.Authorization = &v
	return s
}

type GetAccountForAppRequest struct {
	DeviceInfo *GetAccountForAppRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetAccountForAppRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *GetAccountForAppRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetAccountForAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppRequest) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequest) SetDeviceInfo(v *GetAccountForAppRequestDeviceInfo) *GetAccountForAppRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetAccountForAppRequest) SetPayload(v *GetAccountForAppRequestPayload) *GetAccountForAppRequest {
	s.Payload = v
	return s
}

func (s *GetAccountForAppRequest) SetUserInfo(v *GetAccountForAppRequestUserInfo) *GetAccountForAppRequest {
	s.UserInfo = v
	return s
}

type GetAccountForAppRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetAccountForAppRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestDeviceInfo) SetEncodeKey(v string) *GetAccountForAppRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetEncodeType(v string) *GetAccountForAppRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetId(v string) *GetAccountForAppRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetIdType(v string) *GetAccountForAppRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetAccountForAppRequestDeviceInfo) SetOrganizationId(v string) *GetAccountForAppRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetAccountForAppRequestPayload struct {
	Phone      *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s GetAccountForAppRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppRequestPayload) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestPayload) SetPhone(v string) *GetAccountForAppRequestPayload {
	s.Phone = &v
	return s
}

func (s *GetAccountForAppRequestPayload) SetOriginUuid(v string) *GetAccountForAppRequestPayload {
	s.OriginUuid = &v
	return s
}

type GetAccountForAppRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetAccountForAppRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetAccountForAppRequestUserInfo) SetEncodeKey(v string) *GetAccountForAppRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetEncodeType(v string) *GetAccountForAppRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetId(v string) *GetAccountForAppRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetIdType(v string) *GetAccountForAppRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetAccountForAppRequestUserInfo) SetOrganizationId(v string) *GetAccountForAppRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetAccountForAppShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetAccountForAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAccountForAppShrinkRequest) SetDeviceInfoShrink(v string) *GetAccountForAppShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetAccountForAppShrinkRequest) SetPayloadShrink(v string) *GetAccountForAppShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetAccountForAppShrinkRequest) SetUserInfoShrink(v string) *GetAccountForAppShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetAccountForAppResponseBody struct {
	RetCode  *int32                                `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg   *string                               `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *GetAccountForAppResponseBodyRetValue `json:"RetValue,omitempty" xml:"RetValue,omitempty" type:"Struct"`
}

func (s GetAccountForAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountForAppResponseBody) SetRetCode(v int32) *GetAccountForAppResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetAccountForAppResponseBody) SetRetMsg(v string) *GetAccountForAppResponseBody {
	s.RetMsg = &v
	return s
}

func (s *GetAccountForAppResponseBody) SetRetValue(v *GetAccountForAppResponseBodyRetValue) *GetAccountForAppResponseBody {
	s.RetValue = v
	return s
}

type GetAccountForAppResponseBodyRetValue struct {
	IsVip        *bool   `json:"IsVip,omitempty" xml:"IsVip,omitempty"`
	StrVipExpire *string `json:"StrVipExpire,omitempty" xml:"StrVipExpire,omitempty"`
	VipExpireAt  *int64  `json:"VipExpireAt,omitempty" xml:"VipExpireAt,omitempty"`
}

func (s GetAccountForAppResponseBodyRetValue) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppResponseBodyRetValue) GoString() string {
	return s.String()
}

func (s *GetAccountForAppResponseBodyRetValue) SetIsVip(v bool) *GetAccountForAppResponseBodyRetValue {
	s.IsVip = &v
	return s
}

func (s *GetAccountForAppResponseBodyRetValue) SetStrVipExpire(v string) *GetAccountForAppResponseBodyRetValue {
	s.StrVipExpire = &v
	return s
}

func (s *GetAccountForAppResponseBodyRetValue) SetVipExpireAt(v int64) *GetAccountForAppResponseBodyRetValue {
	s.VipExpireAt = &v
	return s
}

type GetAccountForAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountForAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountForAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountForAppResponse) GoString() string {
	return s.String()
}

func (s *GetAccountForAppResponse) SetHeaders(v map[string]*string) *GetAccountForAppResponse {
	s.Headers = v
	return s
}

func (s *GetAccountForAppResponse) SetStatusCode(v int32) *GetAccountForAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountForAppResponse) SetBody(v *GetAccountForAppResponseBody) *GetAccountForAppResponse {
	s.Body = v
	return s
}

type GetBusAppConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetBusAppConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigHeaders) SetCommonHeaders(v map[string]*string) *GetBusAppConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBusAppConfigHeaders) SetXAcsAligenieAccessToken(v string) *GetBusAppConfigHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetBusAppConfigHeaders) SetAuthorization(v string) *GetBusAppConfigHeaders {
	s.Authorization = &v
	return s
}

type GetBusAppConfigRequest struct {
	DeviceInfo *GetBusAppConfigRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetBusAppConfigRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *GetBusAppConfigRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetBusAppConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigRequest) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequest) SetDeviceInfo(v *GetBusAppConfigRequestDeviceInfo) *GetBusAppConfigRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetBusAppConfigRequest) SetPayload(v *GetBusAppConfigRequestPayload) *GetBusAppConfigRequest {
	s.Payload = v
	return s
}

func (s *GetBusAppConfigRequest) SetUserInfo(v *GetBusAppConfigRequestUserInfo) *GetBusAppConfigRequest {
	s.UserInfo = v
	return s
}

type GetBusAppConfigRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetBusAppConfigRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestDeviceInfo) SetEncodeKey(v string) *GetBusAppConfigRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetEncodeType(v string) *GetBusAppConfigRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetId(v string) *GetBusAppConfigRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetIdType(v string) *GetBusAppConfigRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetBusAppConfigRequestDeviceInfo) SetOrganizationId(v string) *GetBusAppConfigRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetBusAppConfigRequestPayload struct {
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
	Phone      *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s GetBusAppConfigRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigRequestPayload) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestPayload) SetOriginUuid(v string) *GetBusAppConfigRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *GetBusAppConfigRequestPayload) SetPhone(v string) *GetBusAppConfigRequestPayload {
	s.Phone = &v
	return s
}

type GetBusAppConfigRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetBusAppConfigRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigRequestUserInfo) SetEncodeKey(v string) *GetBusAppConfigRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetEncodeType(v string) *GetBusAppConfigRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetId(v string) *GetBusAppConfigRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetIdType(v string) *GetBusAppConfigRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetBusAppConfigRequestUserInfo) SetOrganizationId(v string) *GetBusAppConfigRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetBusAppConfigShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetBusAppConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigShrinkRequest) SetDeviceInfoShrink(v string) *GetBusAppConfigShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetBusAppConfigShrinkRequest) SetPayloadShrink(v string) *GetBusAppConfigShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetBusAppConfigShrinkRequest) SetUserInfoShrink(v string) *GetBusAppConfigShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetBusAppConfigResponseBody struct {
	RetCode  *int32                               `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg   *string                              `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *GetBusAppConfigResponseBodyRetValue `json:"RetValue,omitempty" xml:"RetValue,omitempty" type:"Struct"`
}

func (s GetBusAppConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigResponseBody) SetRetCode(v int32) *GetBusAppConfigResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetBusAppConfigResponseBody) SetRetMsg(v string) *GetBusAppConfigResponseBody {
	s.RetMsg = &v
	return s
}

func (s *GetBusAppConfigResponseBody) SetRetValue(v *GetBusAppConfigResponseBodyRetValue) *GetBusAppConfigResponseBody {
	s.RetValue = v
	return s
}

type GetBusAppConfigResponseBodyRetValue struct {
	Cashier        *string `json:"Cashier,omitempty" xml:"Cashier,omitempty"`
	ShoppingBar    *string `json:"ShoppingBar,omitempty" xml:"ShoppingBar,omitempty"`
	ShoppingWindow *string `json:"ShoppingWindow,omitempty" xml:"ShoppingWindow,omitempty"`
	VipLabel       *string `json:"VipLabel,omitempty" xml:"VipLabel,omitempty"`
}

func (s GetBusAppConfigResponseBodyRetValue) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigResponseBodyRetValue) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigResponseBodyRetValue) SetCashier(v string) *GetBusAppConfigResponseBodyRetValue {
	s.Cashier = &v
	return s
}

func (s *GetBusAppConfigResponseBodyRetValue) SetShoppingBar(v string) *GetBusAppConfigResponseBodyRetValue {
	s.ShoppingBar = &v
	return s
}

func (s *GetBusAppConfigResponseBodyRetValue) SetShoppingWindow(v string) *GetBusAppConfigResponseBodyRetValue {
	s.ShoppingWindow = &v
	return s
}

func (s *GetBusAppConfigResponseBodyRetValue) SetVipLabel(v string) *GetBusAppConfigResponseBodyRetValue {
	s.VipLabel = &v
	return s
}

type GetBusAppConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBusAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBusAppConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBusAppConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBusAppConfigResponse) SetHeaders(v map[string]*string) *GetBusAppConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBusAppConfigResponse) SetStatusCode(v int32) *GetBusAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBusAppConfigResponse) SetBody(v *GetBusAppConfigResponseBody) *GetBusAppConfigResponse {
	s.Body = v
	return s
}

type GetPhoneNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetPhoneNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *GetPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPhoneNumberHeaders) SetXAcsAligenieAccessToken(v string) *GetPhoneNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetPhoneNumberHeaders) SetAuthorization(v string) *GetPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

type GetPhoneNumberRequest struct {
	DeviceInfo *GetPhoneNumberRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	UserInfo   *GetPhoneNumberRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberRequest) SetDeviceInfo(v *GetPhoneNumberRequestDeviceInfo) *GetPhoneNumberRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetPhoneNumberRequest) SetUserInfo(v *GetPhoneNumberRequestUserInfo) *GetPhoneNumberRequest {
	s.UserInfo = v
	return s
}

type GetPhoneNumberRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetPhoneNumberRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberRequestDeviceInfo) SetEncodeKey(v string) *GetPhoneNumberRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetEncodeType(v string) *GetPhoneNumberRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetId(v string) *GetPhoneNumberRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetIdType(v string) *GetPhoneNumberRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetPhoneNumberRequestDeviceInfo) SetOrganizationId(v string) *GetPhoneNumberRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetPhoneNumberRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetPhoneNumberRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberRequestUserInfo) SetEncodeKey(v string) *GetPhoneNumberRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetEncodeType(v string) *GetPhoneNumberRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetId(v string) *GetPhoneNumberRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetIdType(v string) *GetPhoneNumberRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetPhoneNumberRequestUserInfo) SetOrganizationId(v string) *GetPhoneNumberRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetPhoneNumberShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetPhoneNumberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberShrinkRequest) SetDeviceInfoShrink(v string) *GetPhoneNumberShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetPhoneNumberShrinkRequest) SetUserInfoShrink(v string) *GetPhoneNumberShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetPhoneNumberResponseBody struct {
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
}

func (s GetPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberResponseBody) SetPhoneNumber(v string) *GetPhoneNumberResponseBody {
	s.PhoneNumber = &v
	return s
}

type GetPhoneNumberResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberResponse) SetHeaders(v map[string]*string) *GetPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberResponse) SetStatusCode(v int32) *GetPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberResponse) SetBody(v *GetPhoneNumberResponseBody) *GetPhoneNumberResponse {
	s.Body = v
	return s
}

type GetReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetReminderHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetReminderHeaders) GoString() string {
	return s.String()
}

func (s *GetReminderHeaders) SetCommonHeaders(v map[string]*string) *GetReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetReminderHeaders) SetXAcsAligenieAccessToken(v string) *GetReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetReminderHeaders) SetAuthorization(v string) *GetReminderHeaders {
	s.Authorization = &v
	return s
}

type GetReminderRequest struct {
	DeviceInfo *GetReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *GetReminderRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *GetReminderRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetReminderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReminderRequest) GoString() string {
	return s.String()
}

func (s *GetReminderRequest) SetDeviceInfo(v *GetReminderRequestDeviceInfo) *GetReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *GetReminderRequest) SetPayload(v *GetReminderRequestPayload) *GetReminderRequest {
	s.Payload = v
	return s
}

func (s *GetReminderRequest) SetUserInfo(v *GetReminderRequestUserInfo) *GetReminderRequest {
	s.UserInfo = v
	return s
}

type GetReminderRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetReminderRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s GetReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *GetReminderRequestDeviceInfo) SetEncodeKey(v string) *GetReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetEncodeType(v string) *GetReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetId(v string) *GetReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetIdType(v string) *GetReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *GetReminderRequestDeviceInfo) SetOrganizationId(v string) *GetReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type GetReminderRequestPayload struct {
	Id      *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDebug *bool  `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s GetReminderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s GetReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *GetReminderRequestPayload) SetId(v int64) *GetReminderRequestPayload {
	s.Id = &v
	return s
}

func (s *GetReminderRequestPayload) SetIsDebug(v bool) *GetReminderRequestPayload {
	s.IsDebug = &v
	return s
}

type GetReminderRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s GetReminderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *GetReminderRequestUserInfo) SetEncodeKey(v string) *GetReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetEncodeType(v string) *GetReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetId(v string) *GetReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetIdType(v string) *GetReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *GetReminderRequestUserInfo) SetOrganizationId(v string) *GetReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type GetReminderShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetReminderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetReminderShrinkRequest) SetDeviceInfoShrink(v string) *GetReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *GetReminderShrinkRequest) SetPayloadShrink(v string) *GetReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *GetReminderShrinkRequest) SetUserInfoShrink(v string) *GetReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type GetReminderResponseBody struct {
	ErrorCode *int32                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Model     *GetReminderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetReminderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReminderResponseBody) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBody) SetErrorCode(v int32) *GetReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetReminderResponseBody) SetErrorMsg(v string) *GetReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetReminderResponseBody) SetModel(v *GetReminderResponseBodyModel) *GetReminderResponseBody {
	s.Model = v
	return s
}

func (s *GetReminderResponseBody) SetSuccess(v bool) *GetReminderResponseBody {
	s.Success = &v
	return s
}

type GetReminderResponseBodyModel struct {
	RemindResponses []*GetReminderResponseBodyModelRemindResponses `json:"RemindResponses,omitempty" xml:"RemindResponses,omitempty" type:"Repeated"`
}

func (s GetReminderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s GetReminderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBodyModel) SetRemindResponses(v []*GetReminderResponseBodyModelRemindResponses) *GetReminderResponseBodyModel {
	s.RemindResponses = v
	return s
}

type GetReminderResponseBodyModelRemindResponses struct {
	ActionTopic    *string                                                    `json:"ActionTopic,omitempty" xml:"ActionTopic,omitempty"`
	DayDesc        *string                                                    `json:"DayDesc,omitempty" xml:"DayDesc,omitempty"`
	RecurrenceRule *GetReminderResponseBodyModelRemindResponsesRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
	RemindId       *int64                                                     `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	RemindTime     *string                                                    `json:"RemindTime,omitempty" xml:"RemindTime,omitempty"`
	RepeatCount    *int32                                                     `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	Week           *string                                                    `json:"Week,omitempty" xml:"Week,omitempty"`
}

func (s GetReminderResponseBodyModelRemindResponses) String() string {
	return tea.Prettify(s)
}

func (s GetReminderResponseBodyModelRemindResponses) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBodyModelRemindResponses) SetActionTopic(v string) *GetReminderResponseBodyModelRemindResponses {
	s.ActionTopic = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetDayDesc(v string) *GetReminderResponseBodyModelRemindResponses {
	s.DayDesc = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRecurrenceRule(v *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) *GetReminderResponseBodyModelRemindResponses {
	s.RecurrenceRule = v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRemindId(v int64) *GetReminderResponseBodyModelRemindResponses {
	s.RemindId = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRemindTime(v string) *GetReminderResponseBodyModelRemindResponses {
	s.RemindTime = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetRepeatCount(v int32) *GetReminderResponseBodyModelRemindResponses {
	s.RepeatCount = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponses) SetWeek(v string) *GetReminderResponseBodyModelRemindResponses {
	s.Week = &v
	return s
}

type GetReminderResponseBodyModelRemindResponsesRecurrenceRule struct {
	Day           *int32   `json:"Day,omitempty" xml:"Day,omitempty"`
	DaysOfMonth   []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	DaysOfWeek    []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	EndDateTime   *string  `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	Freq          *string  `json:"Freq,omitempty" xml:"Freq,omitempty"`
	Hour          *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute        *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month         *int32   `json:"Month,omitempty" xml:"Month,omitempty"`
	Second        *int32   `json:"Second,omitempty" xml:"Second,omitempty"`
	StartDateTime *string  `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
	Year          *int32   `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s GetReminderResponseBodyModelRemindResponsesRecurrenceRule) String() string {
	return tea.Prettify(s)
}

func (s GetReminderResponseBodyModelRemindResponsesRecurrenceRule) GoString() string {
	return s.String()
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetDay(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Day = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfMonth(v []*int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfWeek(v []*int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetEndDateTime(v string) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetFreq(v string) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetHour(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetMinute(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetMonth(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Month = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetSecond(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Second = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetStartDateTime(v string) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *GetReminderResponseBodyModelRemindResponsesRecurrenceRule) SetYear(v int32) *GetReminderResponseBodyModelRemindResponsesRecurrenceRule {
	s.Year = &v
	return s
}

type GetReminderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReminderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReminderResponse) GoString() string {
	return s.String()
}

func (s *GetReminderResponse) SetHeaders(v map[string]*string) *GetReminderResponse {
	s.Headers = v
	return s
}

func (s *GetReminderResponse) SetStatusCode(v int32) *GetReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReminderResponse) SetBody(v *GetReminderResponseBody) *GetReminderResponse {
	s.Body = v
	return s
}

type ListRemindersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListRemindersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersHeaders) GoString() string {
	return s.String()
}

func (s *ListRemindersHeaders) SetCommonHeaders(v map[string]*string) *ListRemindersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRemindersHeaders) SetXAcsAligenieAccessToken(v string) *ListRemindersHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListRemindersHeaders) SetAuthorization(v string) *ListRemindersHeaders {
	s.Authorization = &v
	return s
}

type ListRemindersRequest struct {
	DeviceInfo *ListRemindersRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *ListRemindersRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *ListRemindersRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListRemindersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersRequest) GoString() string {
	return s.String()
}

func (s *ListRemindersRequest) SetDeviceInfo(v *ListRemindersRequestDeviceInfo) *ListRemindersRequest {
	s.DeviceInfo = v
	return s
}

func (s *ListRemindersRequest) SetPayload(v *ListRemindersRequestPayload) *ListRemindersRequest {
	s.Payload = v
	return s
}

func (s *ListRemindersRequest) SetUserInfo(v *ListRemindersRequestUserInfo) *ListRemindersRequest {
	s.UserInfo = v
	return s
}

type ListRemindersRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRemindersRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestDeviceInfo) SetEncodeKey(v string) *ListRemindersRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetEncodeType(v string) *ListRemindersRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetId(v string) *ListRemindersRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetIdType(v string) *ListRemindersRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *ListRemindersRequestDeviceInfo) SetOrganizationId(v string) *ListRemindersRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type ListRemindersRequestPayload struct {
	IsDebug *bool `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
}

func (s ListRemindersRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersRequestPayload) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestPayload) SetIsDebug(v bool) *ListRemindersRequestPayload {
	s.IsDebug = &v
	return s
}

type ListRemindersRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ListRemindersRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ListRemindersRequestUserInfo) SetEncodeKey(v string) *ListRemindersRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetEncodeType(v string) *ListRemindersRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetId(v string) *ListRemindersRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetIdType(v string) *ListRemindersRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ListRemindersRequestUserInfo) SetOrganizationId(v string) *ListRemindersRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type ListRemindersShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ListRemindersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRemindersShrinkRequest) SetDeviceInfoShrink(v string) *ListRemindersShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *ListRemindersShrinkRequest) SetPayloadShrink(v string) *ListRemindersShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *ListRemindersShrinkRequest) SetUserInfoShrink(v string) *ListRemindersShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type ListRemindersResponseBody struct {
	ErrorCode *int32                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string                         `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Model     *ListRemindersResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRemindersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBody) SetErrorCode(v int32) *ListRemindersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRemindersResponseBody) SetErrorMsg(v string) *ListRemindersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListRemindersResponseBody) SetModel(v *ListRemindersResponseBodyModel) *ListRemindersResponseBody {
	s.Model = v
	return s
}

func (s *ListRemindersResponseBody) SetSuccess(v bool) *ListRemindersResponseBody {
	s.Success = &v
	return s
}

type ListRemindersResponseBodyModel struct {
	RemindResponses []*ListRemindersResponseBodyModelRemindResponses `json:"RemindResponses,omitempty" xml:"RemindResponses,omitempty" type:"Repeated"`
}

func (s ListRemindersResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBodyModel) SetRemindResponses(v []*ListRemindersResponseBodyModelRemindResponses) *ListRemindersResponseBodyModel {
	s.RemindResponses = v
	return s
}

type ListRemindersResponseBodyModelRemindResponses struct {
	ActionTopic    *string                                                      `json:"ActionTopic,omitempty" xml:"ActionTopic,omitempty"`
	DayDesc        *string                                                      `json:"DayDesc,omitempty" xml:"DayDesc,omitempty"`
	RecurrenceRule *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
	RemindId       *int64                                                       `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	RemindTime     *string                                                      `json:"RemindTime,omitempty" xml:"RemindTime,omitempty"`
	RepeatCount    *int32                                                       `json:"RepeatCount,omitempty" xml:"RepeatCount,omitempty"`
	Week           *string                                                      `json:"Week,omitempty" xml:"Week,omitempty"`
}

func (s ListRemindersResponseBodyModelRemindResponses) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersResponseBodyModelRemindResponses) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetActionTopic(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.ActionTopic = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetDayDesc(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.DayDesc = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRecurrenceRule(v *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) *ListRemindersResponseBodyModelRemindResponses {
	s.RecurrenceRule = v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRemindId(v int64) *ListRemindersResponseBodyModelRemindResponses {
	s.RemindId = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRemindTime(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.RemindTime = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetRepeatCount(v int32) *ListRemindersResponseBodyModelRemindResponses {
	s.RepeatCount = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponses) SetWeek(v string) *ListRemindersResponseBodyModelRemindResponses {
	s.Week = &v
	return s
}

type ListRemindersResponseBodyModelRemindResponsesRecurrenceRule struct {
	Day           *int32   `json:"Day,omitempty" xml:"Day,omitempty"`
	DaysOfMonth   []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	DaysOfWeek    []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	EndDateTime   *string  `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	Freq          *string  `json:"Freq,omitempty" xml:"Freq,omitempty"`
	Hour          *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute        *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month         *int32   `json:"Month,omitempty" xml:"Month,omitempty"`
	Second        *int32   `json:"Second,omitempty" xml:"Second,omitempty"`
	StartDateTime *string  `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
	Year          *int32   `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) GoString() string {
	return s.String()
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetDay(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Day = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfMonth(v []*int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetDaysOfWeek(v []*int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetEndDateTime(v string) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetFreq(v string) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetHour(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetMinute(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetMonth(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Month = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetSecond(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Second = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetStartDateTime(v string) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule) SetYear(v int32) *ListRemindersResponseBodyModelRemindResponsesRecurrenceRule {
	s.Year = &v
	return s
}

type ListRemindersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRemindersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRemindersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRemindersResponse) GoString() string {
	return s.String()
}

func (s *ListRemindersResponse) SetHeaders(v map[string]*string) *ListRemindersResponse {
	s.Headers = v
	return s
}

func (s *ListRemindersResponse) SetStatusCode(v int32) *ListRemindersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemindersResponse) SetBody(v *ListRemindersResponseBody) *ListRemindersResponse {
	s.Body = v
	return s
}

type PullCashierHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PullCashierHeaders) String() string {
	return tea.Prettify(s)
}

func (s PullCashierHeaders) GoString() string {
	return s.String()
}

func (s *PullCashierHeaders) SetCommonHeaders(v map[string]*string) *PullCashierHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullCashierHeaders) SetXAcsAligenieAccessToken(v string) *PullCashierHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PullCashierHeaders) SetAuthorization(v string) *PullCashierHeaders {
	s.Authorization = &v
	return s
}

type PullCashierRequest struct {
	DeviceInfo *PullCashierRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *PullCashierRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *PullCashierRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s PullCashierRequest) String() string {
	return tea.Prettify(s)
}

func (s PullCashierRequest) GoString() string {
	return s.String()
}

func (s *PullCashierRequest) SetDeviceInfo(v *PullCashierRequestDeviceInfo) *PullCashierRequest {
	s.DeviceInfo = v
	return s
}

func (s *PullCashierRequest) SetPayload(v *PullCashierRequestPayload) *PullCashierRequest {
	s.Payload = v
	return s
}

func (s *PullCashierRequest) SetUserInfo(v *PullCashierRequestUserInfo) *PullCashierRequest {
	s.UserInfo = v
	return s
}

type PullCashierRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PullCashierRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s PullCashierRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *PullCashierRequestDeviceInfo) SetEncodeKey(v string) *PullCashierRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetEncodeType(v string) *PullCashierRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetId(v string) *PullCashierRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetIdType(v string) *PullCashierRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *PullCashierRequestDeviceInfo) SetOrganizationId(v string) *PullCashierRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type PullCashierRequestPayload struct {
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
}

func (s PullCashierRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s PullCashierRequestPayload) GoString() string {
	return s.String()
}

func (s *PullCashierRequestPayload) SetOriginUuid(v string) *PullCashierRequestPayload {
	s.OriginUuid = &v
	return s
}

type PullCashierRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s PullCashierRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s PullCashierRequestUserInfo) GoString() string {
	return s.String()
}

func (s *PullCashierRequestUserInfo) SetEncodeKey(v string) *PullCashierRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetEncodeType(v string) *PullCashierRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetId(v string) *PullCashierRequestUserInfo {
	s.Id = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetIdType(v string) *PullCashierRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *PullCashierRequestUserInfo) SetOrganizationId(v string) *PullCashierRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type PullCashierShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PullCashierShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PullCashierShrinkRequest) GoString() string {
	return s.String()
}

func (s *PullCashierShrinkRequest) SetDeviceInfoShrink(v string) *PullCashierShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PullCashierShrinkRequest) SetPayloadShrink(v string) *PullCashierShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *PullCashierShrinkRequest) SetUserInfoShrink(v string) *PullCashierShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type PullCashierResponseBody struct {
	RetCode  *int32  `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg   *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *bool   `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
}

func (s PullCashierResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullCashierResponseBody) GoString() string {
	return s.String()
}

func (s *PullCashierResponseBody) SetRetCode(v int32) *PullCashierResponseBody {
	s.RetCode = &v
	return s
}

func (s *PullCashierResponseBody) SetRetMsg(v string) *PullCashierResponseBody {
	s.RetMsg = &v
	return s
}

func (s *PullCashierResponseBody) SetRetValue(v bool) *PullCashierResponseBody {
	s.RetValue = &v
	return s
}

type PullCashierResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PullCashierResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PullCashierResponse) String() string {
	return tea.Prettify(s)
}

func (s PullCashierResponse) GoString() string {
	return s.String()
}

func (s *PullCashierResponse) SetHeaders(v map[string]*string) *PullCashierResponse {
	s.Headers = v
	return s
}

func (s *PullCashierResponse) SetStatusCode(v int32) *PullCashierResponse {
	s.StatusCode = &v
	return s
}

func (s *PullCashierResponse) SetBody(v *PullCashierResponseBody) *PullCashierResponse {
	s.Body = v
	return s
}

type PushNotificationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushNotificationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsHeaders) GoString() string {
	return s.String()
}

func (s *PushNotificationsHeaders) SetCommonHeaders(v map[string]*string) *PushNotificationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushNotificationsHeaders) SetXAcsAligenieAccessToken(v string) *PushNotificationsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushNotificationsHeaders) SetAuthorization(v string) *PushNotificationsHeaders {
	s.Authorization = &v
	return s
}

type PushNotificationsRequest struct {
	NotificationUnicastRequest *PushNotificationsRequestNotificationUnicastRequest `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty" type:"Struct"`
	TenantInfo                 *PushNotificationsRequestTenantInfo                 `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
}

func (s PushNotificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequest) SetNotificationUnicastRequest(v *PushNotificationsRequestNotificationUnicastRequest) *PushNotificationsRequest {
	s.NotificationUnicastRequest = v
	return s
}

func (s *PushNotificationsRequest) SetTenantInfo(v *PushNotificationsRequestTenantInfo) *PushNotificationsRequest {
	s.TenantInfo = v
	return s
}

type PushNotificationsRequestNotificationUnicastRequest struct {
	EncodeKey         *string                                                       `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType        *string                                                       `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	IsDebug           *bool                                                         `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	MessageTemplateId *string                                                       `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	OrganizationId    *string                                                       `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	PlaceHolder       map[string]*string                                            `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	SendTarget        *PushNotificationsRequestNotificationUnicastRequestSendTarget `json:"SendTarget,omitempty" xml:"SendTarget,omitempty" type:"Struct"`
}

func (s PushNotificationsRequestNotificationUnicastRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequestNotificationUnicastRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetEncodeKey(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.EncodeKey = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetEncodeType(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.EncodeType = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetIsDebug(v bool) *PushNotificationsRequestNotificationUnicastRequest {
	s.IsDebug = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetMessageTemplateId(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.MessageTemplateId = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetOrganizationId(v string) *PushNotificationsRequestNotificationUnicastRequest {
	s.OrganizationId = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetPlaceHolder(v map[string]*string) *PushNotificationsRequestNotificationUnicastRequest {
	s.PlaceHolder = v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequest) SetSendTarget(v *PushNotificationsRequestNotificationUnicastRequestSendTarget) *PushNotificationsRequestNotificationUnicastRequest {
	s.SendTarget = v
	return s
}

type PushNotificationsRequestNotificationUnicastRequestSendTarget struct {
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
	TargetType     *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s PushNotificationsRequestNotificationUnicastRequestSendTarget) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequestNotificationUnicastRequestSendTarget) GoString() string {
	return s.String()
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) SetTargetIdentity(v string) *PushNotificationsRequestNotificationUnicastRequestSendTarget {
	s.TargetIdentity = &v
	return s
}

func (s *PushNotificationsRequestNotificationUnicastRequestSendTarget) SetTargetType(v string) *PushNotificationsRequestNotificationUnicastRequestSendTarget {
	s.TargetType = &v
	return s
}

type PushNotificationsRequestTenantInfo struct {
}

func (s PushNotificationsRequestTenantInfo) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsRequestTenantInfo) GoString() string {
	return s.String()
}

type PushNotificationsShrinkRequest struct {
	NotificationUnicastRequestShrink *string `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty"`
	TenantInfoShrink                 *string `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty"`
}

func (s PushNotificationsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushNotificationsShrinkRequest) SetNotificationUnicastRequestShrink(v string) *PushNotificationsShrinkRequest {
	s.NotificationUnicastRequestShrink = &v
	return s
}

func (s *PushNotificationsShrinkRequest) SetTenantInfoShrink(v string) *PushNotificationsShrinkRequest {
	s.TenantInfoShrink = &v
	return s
}

type PushNotificationsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PushNotificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s PushNotificationsResponse) GoString() string {
	return s.String()
}

func (s *PushNotificationsResponse) SetHeaders(v map[string]*string) *PushNotificationsResponse {
	s.Headers = v
	return s
}

func (s *PushNotificationsResponse) SetStatusCode(v int32) *PushNotificationsResponse {
	s.StatusCode = &v
	return s
}

type SendNotificationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SendNotificationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsHeaders) GoString() string {
	return s.String()
}

func (s *SendNotificationsHeaders) SetCommonHeaders(v map[string]*string) *SendNotificationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendNotificationsHeaders) SetXAcsAligenieAccessToken(v string) *SendNotificationsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SendNotificationsHeaders) SetAuthorization(v string) *SendNotificationsHeaders {
	s.Authorization = &v
	return s
}

type SendNotificationsRequest struct {
	DeviceInfo                 *SendNotificationsRequestDeviceInfo                 `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	NotificationUnicastRequest *SendNotificationsRequestNotificationUnicastRequest `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty" type:"Struct"`
	TenantInfo                 *SendNotificationsRequestTenantInfo                 `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty" type:"Struct"`
	UserInfo                   *SendNotificationsRequestUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s SendNotificationsRequest) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequest) SetDeviceInfo(v *SendNotificationsRequestDeviceInfo) *SendNotificationsRequest {
	s.DeviceInfo = v
	return s
}

func (s *SendNotificationsRequest) SetNotificationUnicastRequest(v *SendNotificationsRequestNotificationUnicastRequest) *SendNotificationsRequest {
	s.NotificationUnicastRequest = v
	return s
}

func (s *SendNotificationsRequest) SetTenantInfo(v *SendNotificationsRequestTenantInfo) *SendNotificationsRequest {
	s.TenantInfo = v
	return s
}

func (s *SendNotificationsRequest) SetUserInfo(v *SendNotificationsRequestUserInfo) *SendNotificationsRequest {
	s.UserInfo = v
	return s
}

type SendNotificationsRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SendNotificationsRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestDeviceInfo) SetEncodeKey(v string) *SendNotificationsRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetEncodeType(v string) *SendNotificationsRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetId(v string) *SendNotificationsRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetIdType(v string) *SendNotificationsRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *SendNotificationsRequestDeviceInfo) SetOrganizationId(v string) *SendNotificationsRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type SendNotificationsRequestNotificationUnicastRequest struct {
	IsDebug           *bool                                                         `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	MessageTemplateId *string                                                       `json:"MessageTemplateId,omitempty" xml:"MessageTemplateId,omitempty"`
	PlaceHolder       map[string]*string                                            `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	SendTarget        *SendNotificationsRequestNotificationUnicastRequestSendTarget `json:"SendTarget,omitempty" xml:"SendTarget,omitempty" type:"Struct"`
}

func (s SendNotificationsRequestNotificationUnicastRequest) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsRequestNotificationUnicastRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetIsDebug(v bool) *SendNotificationsRequestNotificationUnicastRequest {
	s.IsDebug = &v
	return s
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetMessageTemplateId(v string) *SendNotificationsRequestNotificationUnicastRequest {
	s.MessageTemplateId = &v
	return s
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetPlaceHolder(v map[string]*string) *SendNotificationsRequestNotificationUnicastRequest {
	s.PlaceHolder = v
	return s
}

func (s *SendNotificationsRequestNotificationUnicastRequest) SetSendTarget(v *SendNotificationsRequestNotificationUnicastRequestSendTarget) *SendNotificationsRequestNotificationUnicastRequest {
	s.SendTarget = v
	return s
}

type SendNotificationsRequestNotificationUnicastRequestSendTarget struct {
}

func (s SendNotificationsRequestNotificationUnicastRequestSendTarget) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsRequestNotificationUnicastRequestSendTarget) GoString() string {
	return s.String()
}

type SendNotificationsRequestTenantInfo struct {
}

func (s SendNotificationsRequestTenantInfo) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsRequestTenantInfo) GoString() string {
	return s.String()
}

type SendNotificationsRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s SendNotificationsRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsRequestUserInfo) GoString() string {
	return s.String()
}

func (s *SendNotificationsRequestUserInfo) SetEncodeKey(v string) *SendNotificationsRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetEncodeType(v string) *SendNotificationsRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetId(v string) *SendNotificationsRequestUserInfo {
	s.Id = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetIdType(v string) *SendNotificationsRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *SendNotificationsRequestUserInfo) SetOrganizationId(v string) *SendNotificationsRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type SendNotificationsShrinkRequest struct {
	DeviceInfoShrink                 *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	NotificationUnicastRequestShrink *string `json:"NotificationUnicastRequest,omitempty" xml:"NotificationUnicastRequest,omitempty"`
	TenantInfoShrink                 *string `json:"TenantInfo,omitempty" xml:"TenantInfo,omitempty"`
	UserInfoShrink                   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SendNotificationsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationsShrinkRequest) SetDeviceInfoShrink(v string) *SendNotificationsShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *SendNotificationsShrinkRequest) SetNotificationUnicastRequestShrink(v string) *SendNotificationsShrinkRequest {
	s.NotificationUnicastRequestShrink = &v
	return s
}

func (s *SendNotificationsShrinkRequest) SetTenantInfoShrink(v string) *SendNotificationsShrinkRequest {
	s.TenantInfoShrink = &v
	return s
}

func (s *SendNotificationsShrinkRequest) SetUserInfoShrink(v string) *SendNotificationsShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type SendNotificationsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s SendNotificationsResponse) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationsResponse) GoString() string {
	return s.String()
}

func (s *SendNotificationsResponse) SetHeaders(v map[string]*string) *SendNotificationsResponse {
	s.Headers = v
	return s
}

func (s *SendNotificationsResponse) SetStatusCode(v int32) *SendNotificationsResponse {
	s.StatusCode = &v
	return s
}

type UpdateReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateReminderHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderHeaders) GoString() string {
	return s.String()
}

func (s *UpdateReminderHeaders) SetCommonHeaders(v map[string]*string) *UpdateReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateReminderHeaders) SetXAcsAligenieAccessToken(v string) *UpdateReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateReminderHeaders) SetAuthorization(v string) *UpdateReminderHeaders {
	s.Authorization = &v
	return s
}

type UpdateReminderRequest struct {
	DeviceInfo *UpdateReminderRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *UpdateReminderRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *UpdateReminderRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s UpdateReminderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderRequest) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequest) SetDeviceInfo(v *UpdateReminderRequestDeviceInfo) *UpdateReminderRequest {
	s.DeviceInfo = v
	return s
}

func (s *UpdateReminderRequest) SetPayload(v *UpdateReminderRequestPayload) *UpdateReminderRequest {
	s.Payload = v
	return s
}

func (s *UpdateReminderRequest) SetUserInfo(v *UpdateReminderRequestUserInfo) *UpdateReminderRequest {
	s.UserInfo = v
	return s
}

type UpdateReminderRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateReminderRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestDeviceInfo) SetEncodeKey(v string) *UpdateReminderRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetEncodeType(v string) *UpdateReminderRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetId(v string) *UpdateReminderRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetIdType(v string) *UpdateReminderRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *UpdateReminderRequestDeviceInfo) SetOrganizationId(v string) *UpdateReminderRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type UpdateReminderRequestPayload struct {
	Content        *string                                     `json:"Content,omitempty" xml:"Content,omitempty"`
	Id             *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDebug        *bool                                       `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	RecurrenceRule *UpdateReminderRequestPayloadRecurrenceRule `json:"RecurrenceRule,omitempty" xml:"RecurrenceRule,omitempty" type:"Struct"`
}

func (s UpdateReminderRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderRequestPayload) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestPayload) SetContent(v string) *UpdateReminderRequestPayload {
	s.Content = &v
	return s
}

func (s *UpdateReminderRequestPayload) SetId(v int64) *UpdateReminderRequestPayload {
	s.Id = &v
	return s
}

func (s *UpdateReminderRequestPayload) SetIsDebug(v bool) *UpdateReminderRequestPayload {
	s.IsDebug = &v
	return s
}

func (s *UpdateReminderRequestPayload) SetRecurrenceRule(v *UpdateReminderRequestPayloadRecurrenceRule) *UpdateReminderRequestPayload {
	s.RecurrenceRule = v
	return s
}

type UpdateReminderRequestPayloadRecurrenceRule struct {
	Day           *int32   `json:"Day,omitempty" xml:"Day,omitempty"`
	DaysOfMonth   []*int32 `json:"DaysOfMonth,omitempty" xml:"DaysOfMonth,omitempty" type:"Repeated"`
	DaysOfWeek    []*int32 `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty" type:"Repeated"`
	EndDateTime   *int64   `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	Freq          *string  `json:"Freq,omitempty" xml:"Freq,omitempty"`
	Hour          *int32   `json:"Hour,omitempty" xml:"Hour,omitempty"`
	Minute        *int32   `json:"Minute,omitempty" xml:"Minute,omitempty"`
	Month         *int32   `json:"Month,omitempty" xml:"Month,omitempty"`
	Second        *int32   `json:"Second,omitempty" xml:"Second,omitempty"`
	StartDateTime *int64   `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
	Year          *int32   `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s UpdateReminderRequestPayloadRecurrenceRule) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderRequestPayloadRecurrenceRule) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetDay(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Day = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetDaysOfMonth(v []*int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.DaysOfMonth = v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetDaysOfWeek(v []*int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.DaysOfWeek = v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetEndDateTime(v int64) *UpdateReminderRequestPayloadRecurrenceRule {
	s.EndDateTime = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetFreq(v string) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Freq = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetHour(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Hour = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetMinute(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Minute = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetMonth(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Month = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetSecond(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Second = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetStartDateTime(v int64) *UpdateReminderRequestPayloadRecurrenceRule {
	s.StartDateTime = &v
	return s
}

func (s *UpdateReminderRequestPayloadRecurrenceRule) SetYear(v int32) *UpdateReminderRequestPayloadRecurrenceRule {
	s.Year = &v
	return s
}

type UpdateReminderRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s UpdateReminderRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderRequestUserInfo) GoString() string {
	return s.String()
}

func (s *UpdateReminderRequestUserInfo) SetEncodeKey(v string) *UpdateReminderRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetEncodeType(v string) *UpdateReminderRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetId(v string) *UpdateReminderRequestUserInfo {
	s.Id = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetIdType(v string) *UpdateReminderRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *UpdateReminderRequestUserInfo) SetOrganizationId(v string) *UpdateReminderRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type UpdateReminderShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s UpdateReminderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateReminderShrinkRequest) SetDeviceInfoShrink(v string) *UpdateReminderShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *UpdateReminderShrinkRequest) SetPayloadShrink(v string) *UpdateReminderShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *UpdateReminderShrinkRequest) SetUserInfoShrink(v string) *UpdateReminderShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type UpdateReminderResponseBody struct {
	ErrorCode *int32  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Model     *int64  `json:"Model,omitempty" xml:"Model,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateReminderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReminderResponseBody) SetErrorCode(v int32) *UpdateReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateReminderResponseBody) SetErrorMsg(v string) *UpdateReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateReminderResponseBody) SetModel(v int64) *UpdateReminderResponseBody {
	s.Model = &v
	return s
}

func (s *UpdateReminderResponseBody) SetSuccess(v bool) *UpdateReminderResponseBody {
	s.Success = &v
	return s
}

type UpdateReminderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateReminderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateReminderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateReminderResponse) GoString() string {
	return s.String()
}

func (s *UpdateReminderResponse) SetHeaders(v map[string]*string) *UpdateReminderResponse {
	s.Headers = v
	return s
}

func (s *UpdateReminderResponse) SetStatusCode(v int32) *UpdateReminderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateReminderResponse) SetBody(v *UpdateReminderResponseBody) *UpdateReminderResponse {
	s.Body = v
	return s
}

type VideoAppReportHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s VideoAppReportHeaders) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportHeaders) GoString() string {
	return s.String()
}

func (s *VideoAppReportHeaders) SetCommonHeaders(v map[string]*string) *VideoAppReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *VideoAppReportHeaders) SetXAcsAligenieAccessToken(v string) *VideoAppReportHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *VideoAppReportHeaders) SetAuthorization(v string) *VideoAppReportHeaders {
	s.Authorization = &v
	return s
}

type VideoAppReportRequest struct {
	DeviceInfo *VideoAppReportRequestDeviceInfo `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty" type:"Struct"`
	Payload    *VideoAppReportRequestPayload    `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	UserInfo   *VideoAppReportRequestUserInfo   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s VideoAppReportRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportRequest) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequest) SetDeviceInfo(v *VideoAppReportRequestDeviceInfo) *VideoAppReportRequest {
	s.DeviceInfo = v
	return s
}

func (s *VideoAppReportRequest) SetPayload(v *VideoAppReportRequestPayload) *VideoAppReportRequest {
	s.Payload = v
	return s
}

func (s *VideoAppReportRequest) SetUserInfo(v *VideoAppReportRequestUserInfo) *VideoAppReportRequest {
	s.UserInfo = v
	return s
}

type VideoAppReportRequestDeviceInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s VideoAppReportRequestDeviceInfo) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportRequestDeviceInfo) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequestDeviceInfo) SetEncodeKey(v string) *VideoAppReportRequestDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetEncodeType(v string) *VideoAppReportRequestDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetId(v string) *VideoAppReportRequestDeviceInfo {
	s.Id = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetIdType(v string) *VideoAppReportRequestDeviceInfo {
	s.IdType = &v
	return s
}

func (s *VideoAppReportRequestDeviceInfo) SetOrganizationId(v string) *VideoAppReportRequestDeviceInfo {
	s.OrganizationId = &v
	return s
}

type VideoAppReportRequestPayload struct {
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IsLogin    *bool   `json:"isLogin,omitempty" xml:"isLogin,omitempty"`
	IsVip      *bool   `json:"isVip,omitempty" xml:"isVip,omitempty"`
	LoginNick  *string `json:"loginNick,omitempty" xml:"loginNick,omitempty"`
	OriginUuid *string `json:"originUuid,omitempty" xml:"originUuid,omitempty"`
	Phone      *string `json:"phone,omitempty" xml:"phone,omitempty"`
	PkgName    *string `json:"pkgName,omitempty" xml:"pkgName,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s VideoAppReportRequestPayload) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportRequestPayload) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequestPayload) SetEndTime(v int64) *VideoAppReportRequestPayload {
	s.EndTime = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetIsLogin(v bool) *VideoAppReportRequestPayload {
	s.IsLogin = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetIsVip(v bool) *VideoAppReportRequestPayload {
	s.IsVip = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetLoginNick(v string) *VideoAppReportRequestPayload {
	s.LoginNick = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetOriginUuid(v string) *VideoAppReportRequestPayload {
	s.OriginUuid = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetPhone(v string) *VideoAppReportRequestPayload {
	s.Phone = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetPkgName(v string) *VideoAppReportRequestPayload {
	s.PkgName = &v
	return s
}

func (s *VideoAppReportRequestPayload) SetStartTime(v int64) *VideoAppReportRequestPayload {
	s.StartTime = &v
	return s
}

type VideoAppReportRequestUserInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s VideoAppReportRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportRequestUserInfo) GoString() string {
	return s.String()
}

func (s *VideoAppReportRequestUserInfo) SetEncodeKey(v string) *VideoAppReportRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetEncodeType(v string) *VideoAppReportRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetId(v string) *VideoAppReportRequestUserInfo {
	s.Id = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetIdType(v string) *VideoAppReportRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *VideoAppReportRequestUserInfo) SetOrganizationId(v string) *VideoAppReportRequestUserInfo {
	s.OrganizationId = &v
	return s
}

type VideoAppReportShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink   *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s VideoAppReportShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *VideoAppReportShrinkRequest) SetDeviceInfoShrink(v string) *VideoAppReportShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *VideoAppReportShrinkRequest) SetPayloadShrink(v string) *VideoAppReportShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *VideoAppReportShrinkRequest) SetUserInfoShrink(v string) *VideoAppReportShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

type VideoAppReportResponseBody struct {
	RetCode  *int32  `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg   *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	RetValue *bool   `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
}

func (s VideoAppReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportResponseBody) GoString() string {
	return s.String()
}

func (s *VideoAppReportResponseBody) SetRetCode(v int32) *VideoAppReportResponseBody {
	s.RetCode = &v
	return s
}

func (s *VideoAppReportResponseBody) SetRetMsg(v string) *VideoAppReportResponseBody {
	s.RetMsg = &v
	return s
}

func (s *VideoAppReportResponseBody) SetRetValue(v bool) *VideoAppReportResponseBody {
	s.RetValue = &v
	return s
}

type VideoAppReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoAppReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VideoAppReportResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoAppReportResponse) GoString() string {
	return s.String()
}

func (s *VideoAppReportResponse) SetHeaders(v map[string]*string) *VideoAppReportResponse {
	s.Headers = v
	return s
}

func (s *VideoAppReportResponse) SetStatusCode(v int32) *VideoAppReportResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoAppReportResponse) SetBody(v *VideoAppReportResponseBody) *VideoAppReportResponse {
	s.Body = v
	return s
}

type WakeUpAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s WakeUpAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppHeaders) GoString() string {
	return s.String()
}

func (s *WakeUpAppHeaders) SetCommonHeaders(v map[string]*string) *WakeUpAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WakeUpAppHeaders) SetXAcsAligenieAccessToken(v string) *WakeUpAppHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *WakeUpAppHeaders) SetAuthorization(v string) *WakeUpAppHeaders {
	s.Authorization = &v
	return s
}

type WakeUpAppRequest struct {
	IsDebug    *bool                       `json:"IsDebug,omitempty" xml:"IsDebug,omitempty"`
	Path       *string                     `json:"Path,omitempty" xml:"Path,omitempty"`
	TargetInfo *WakeUpAppRequestTargetInfo `json:"TargetInfo,omitempty" xml:"TargetInfo,omitempty" type:"Struct"`
}

func (s WakeUpAppRequest) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppRequest) GoString() string {
	return s.String()
}

func (s *WakeUpAppRequest) SetIsDebug(v bool) *WakeUpAppRequest {
	s.IsDebug = &v
	return s
}

func (s *WakeUpAppRequest) SetPath(v string) *WakeUpAppRequest {
	s.Path = &v
	return s
}

func (s *WakeUpAppRequest) SetTargetInfo(v *WakeUpAppRequestTargetInfo) *WakeUpAppRequest {
	s.TargetInfo = v
	return s
}

type WakeUpAppRequestTargetInfo struct {
	EncodeKey      *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	EncodeType     *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	TargetIdentity *string `json:"TargetIdentity,omitempty" xml:"TargetIdentity,omitempty"`
	TargetType     *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s WakeUpAppRequestTargetInfo) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppRequestTargetInfo) GoString() string {
	return s.String()
}

func (s *WakeUpAppRequestTargetInfo) SetEncodeKey(v string) *WakeUpAppRequestTargetInfo {
	s.EncodeKey = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetEncodeType(v string) *WakeUpAppRequestTargetInfo {
	s.EncodeType = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetOrganizationId(v string) *WakeUpAppRequestTargetInfo {
	s.OrganizationId = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetTargetIdentity(v string) *WakeUpAppRequestTargetInfo {
	s.TargetIdentity = &v
	return s
}

func (s *WakeUpAppRequestTargetInfo) SetTargetType(v string) *WakeUpAppRequestTargetInfo {
	s.TargetType = &v
	return s
}

type WakeUpAppResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s WakeUpAppResponse) String() string {
	return tea.Prettify(s)
}

func (s WakeUpAppResponse) GoString() string {
	return s.String()
}

func (s *WakeUpAppResponse) SetHeaders(v map[string]*string) *WakeUpAppResponse {
	s.Headers = v
	return s
}

func (s *WakeUpAppResponse) SetStatusCode(v int32) *WakeUpAppResponse {
	s.StatusCode = &v
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aligenie"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AppUseTimeReportWithOptions(tmpReq *AppUseTimeReportRequest, headers *AppUseTimeReportHeaders, runtime *util.RuntimeOptions) (_result *AppUseTimeReportResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AppUseTimeReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppUseTimeReport"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/vip/use/time/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppUseTimeReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppUseTimeReport(request *AppUseTimeReportRequest) (_result *AppUseTimeReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppUseTimeReportHeaders{}
	_result = &AppUseTimeReportResponse{}
	_body, _err := client.AppUseTimeReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CallBackThirdRightSendPlanWithOptions(tmpReq *CallBackThirdRightSendPlanRequest, headers *CallBackThirdRightSendPlanHeaders, runtime *util.RuntimeOptions) (_result *CallBackThirdRightSendPlanResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CallBackThirdRightSendPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendInfo)) {
		request.ExtendInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendInfo, tea.String("ExtendInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizGroup)) {
		query["BizGroup"] = request.BizGroup
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CardType)) {
		query["CardType"] = request.CardType
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMsg)) {
		query["ErrorMsg"] = request.ErrorMsg
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendInfoShrink)) {
		query["ExtendInfo"] = request.ExtendInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GenieOpenId)) {
		query["GenieOpenId"] = request.GenieOpenId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveStatus)) {
		query["ReceiveStatus"] = request.ReceiveStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["Sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierId)) {
		query["SupplierId"] = request.SupplierId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CallBackThirdRightSendPlan"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/1.0/iap/business/CallBackThirdRightSendPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CallBackThirdRightSendPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CallBackThirdRightSendPlan(request *CallBackThirdRightSendPlanRequest) (_result *CallBackThirdRightSendPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CallBackThirdRightSendPlanHeaders{}
	_result = &CallBackThirdRightSendPlanResponse{}
	_body, _err := client.CallBackThirdRightSendPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckThirdRightSendPlanWithOptions(tmpReq *CheckThirdRightSendPlanRequest, headers *CheckThirdRightSendPlanHeaders, runtime *util.RuntimeOptions) (_result *CheckThirdRightSendPlanResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CheckThirdRightSendPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtendInfo)) {
		request.ExtendInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtendInfo, tea.String("ExtendInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizGroup)) {
		query["BizGroup"] = request.BizGroup
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendInfoShrink)) {
		query["ExtendInfo"] = request.ExtendInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Sn)) {
		query["Sn"] = request.Sn
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierId)) {
		query["SupplierId"] = request.SupplierId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckThirdRightSendPlan"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/business/CheckThirdRightSendPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckThirdRightSendPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckThirdRightSendPlan(request *CheckThirdRightSendPlanRequest) (_result *CheckThirdRightSendPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckThirdRightSendPlanHeaders{}
	_result = &CheckThirdRightSendPlanResponse{}
	_body, _err := client.CheckThirdRightSendPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateReminderWithOptions(tmpReq *CreateReminderRequest, headers *CreateReminderHeaders, runtime *util.RuntimeOptions) (_result *CreateReminderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateReminder"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/reminder/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateReminderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateReminder(request *CreateReminderRequest) (_result *CreateReminderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateReminderHeaders{}
	_result = &CreateReminderResponse{}
	_body, _err := client.CreateReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteReminderWithOptions(tmpReq *DeleteReminderRequest, headers *DeleteReminderHeaders, runtime *util.RuntimeOptions) (_result *DeleteReminderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteReminder"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/reminder/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteReminderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteReminder(request *DeleteReminderRequest) (_result *DeleteReminderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteReminderHeaders{}
	_result = &DeleteReminderResponse{}
	_body, _err := client.DeleteReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountForAppWithOptions(tmpReq *GetAccountForAppRequest, headers *GetAccountForAppHeaders, runtime *util.RuntimeOptions) (_result *GetAccountForAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetAccountForAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountForApp"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/vip/account/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountForAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountForApp(request *GetAccountForAppRequest) (_result *GetAccountForAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAccountForAppHeaders{}
	_result = &GetAccountForAppResponse{}
	_body, _err := client.GetAccountForAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBusAppConfigWithOptions(tmpReq *GetBusAppConfigRequest, headers *GetBusAppConfigHeaders, runtime *util.RuntimeOptions) (_result *GetBusAppConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetBusAppConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBusAppConfig"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/app/config/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBusAppConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBusAppConfig(request *GetBusAppConfigRequest) (_result *GetBusAppConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBusAppConfigHeaders{}
	_result = &GetBusAppConfigResponse{}
	_body, _err := client.GetBusAppConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhoneNumberWithOptions(tmpReq *GetPhoneNumberRequest, headers *GetPhoneNumberHeaders, runtime *util.RuntimeOptions) (_result *GetPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPhoneNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhoneNumber"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/profile/phoneNumber"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhoneNumber(request *GetPhoneNumberRequest) (_result *GetPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPhoneNumberHeaders{}
	_result = &GetPhoneNumberResponse{}
	_body, _err := client.GetPhoneNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetReminderWithOptions(tmpReq *GetReminderRequest, headers *GetReminderHeaders, runtime *util.RuntimeOptions) (_result *GetReminderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReminder"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/reminder/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReminderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetReminder(request *GetReminderRequest) (_result *GetReminderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetReminderHeaders{}
	_result = &GetReminderResponse{}
	_body, _err := client.GetReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRemindersWithOptions(tmpReq *ListRemindersRequest, headers *ListRemindersHeaders, runtime *util.RuntimeOptions) (_result *ListRemindersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListRemindersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListReminders"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/reminder/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRemindersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListReminders(request *ListRemindersRequest) (_result *ListRemindersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRemindersHeaders{}
	_result = &ListRemindersResponse{}
	_body, _err := client.ListRemindersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullCashierWithOptions(tmpReq *PullCashierRequest, headers *PullCashierHeaders, runtime *util.RuntimeOptions) (_result *PullCashierResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PullCashierShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		query["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		query["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		query["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PullCashier"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/pull/cashier/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PullCashierResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullCashier(request *PullCashierRequest) (_result *PullCashierResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PullCashierHeaders{}
	_result = &PullCashierResponse{}
	_body, _err := client.PullCashierWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushNotificationsWithOptions(tmpReq *PushNotificationsRequest, headers *PushNotificationsHeaders, runtime *util.RuntimeOptions) (_result *PushNotificationsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &PushNotificationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NotificationUnicastRequest)) {
		request.NotificationUnicastRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationUnicastRequest, tea.String("NotificationUnicastRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantInfo)) {
		request.TenantInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantInfo, tea.String("TenantInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotificationUnicastRequestShrink)) {
		body["NotificationUnicastRequest"] = request.NotificationUnicastRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantInfoShrink)) {
		body["TenantInfo"] = request.TenantInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushNotifications"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/notifications"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &PushNotificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushNotifications(request *PushNotificationsRequest) (_result *PushNotificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushNotificationsHeaders{}
	_result = &PushNotificationsResponse{}
	_body, _err := client.PushNotificationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendNotificationsWithOptions(tmpReq *SendNotificationsRequest, headers *SendNotificationsHeaders, runtime *util.RuntimeOptions) (_result *SendNotificationsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendNotificationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NotificationUnicastRequest)) {
		request.NotificationUnicastRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationUnicastRequest, tea.String("NotificationUnicastRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TenantInfo)) {
		request.TenantInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TenantInfo, tea.String("TenantInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationUnicastRequestShrink)) {
		body["NotificationUnicastRequest"] = request.NotificationUnicastRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantInfoShrink)) {
		body["TenantInfo"] = request.TenantInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendNotifications"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/general/notifications"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &SendNotificationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendNotifications(request *SendNotificationsRequest) (_result *SendNotificationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendNotificationsHeaders{}
	_result = &SendNotificationsResponse{}
	_body, _err := client.SendNotificationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateReminderWithOptions(tmpReq *UpdateReminderRequest, headers *UpdateReminderHeaders, runtime *util.RuntimeOptions) (_result *UpdateReminderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateReminderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateReminder"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/reminder/update"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateReminderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateReminder(request *UpdateReminderRequest) (_result *UpdateReminderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateReminderHeaders{}
	_result = &UpdateReminderResponse{}
	_body, _err := client.UpdateReminderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoAppReportWithOptions(tmpReq *VideoAppReportRequest, headers *VideoAppReportHeaders, runtime *util.RuntimeOptions) (_result *VideoAppReportResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &VideoAppReportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DeviceInfo)) {
		request.DeviceInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeviceInfo, tea.String("DeviceInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Payload)) {
		request.PayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Payload, tea.String("Payload"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceInfoShrink)) {
		body["DeviceInfo"] = request.DeviceInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadShrink)) {
		body["Payload"] = request.PayloadShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VideoAppReport"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/vip/use/video/report"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoAppReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoAppReport(request *VideoAppReportRequest) (_result *VideoAppReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &VideoAppReportHeaders{}
	_result = &VideoAppReportResponse{}
	_body, _err := client.VideoAppReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WakeUpAppWithOptions(request *WakeUpAppRequest, headers *WakeUpAppHeaders, runtime *util.RuntimeOptions) (_result *WakeUpAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDebug)) {
		body["IsDebug"] = request.IsDebug
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInfo)) {
		body["TargetInfo"] = request.TargetInfo
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAligenieAccessToken)) {
		realHeaders["x-acs-aligenie-access-token"] = util.ToJSONString(headers.XAcsAligenieAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.Authorization)) {
		realHeaders["Authorization"] = util.ToJSONString(headers.Authorization)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("WakeUpApp"),
		Version:     tea.String("iap_1.0"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v1.0/iap/wakeup"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &WakeUpAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WakeUpApp(request *WakeUpAppRequest) (_result *WakeUpAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WakeUpAppHeaders{}
	_result = &WakeUpAppResponse{}
	_body, _err := client.WakeUpAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
