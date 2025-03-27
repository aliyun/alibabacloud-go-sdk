// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AccessPageSetAclRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// FREE_ACCESS
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a-062wec3cwmayw****
	AccessPageId *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	// example:
	//
	// notepad_test
	AccessPageName *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7
	EffectTime *int32 `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// example:
	//
	// Day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s AccessPageSetAclRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessPageSetAclRequest) GoString() string {
	return s.String()
}

func (s *AccessPageSetAclRequest) SetAccessMode(v string) *AccessPageSetAclRequest {
	s.AccessMode = &v
	return s
}

func (s *AccessPageSetAclRequest) SetAccessPageId(v string) *AccessPageSetAclRequest {
	s.AccessPageId = &v
	return s
}

func (s *AccessPageSetAclRequest) SetAccessPageName(v string) *AccessPageSetAclRequest {
	s.AccessPageName = &v
	return s
}

func (s *AccessPageSetAclRequest) SetEffectTime(v int32) *AccessPageSetAclRequest {
	s.EffectTime = &v
	return s
}

func (s *AccessPageSetAclRequest) SetUnit(v string) *AccessPageSetAclRequest {
	s.Unit = &v
	return s
}

type AccessPageSetAclResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// InternalError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E25FC620-6B6F-12D2-A992-AD8727DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AccessPageSetAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AccessPageSetAclResponseBody) GoString() string {
	return s.String()
}

func (s *AccessPageSetAclResponseBody) SetCode(v string) *AccessPageSetAclResponseBody {
	s.Code = &v
	return s
}

func (s *AccessPageSetAclResponseBody) SetMessage(v string) *AccessPageSetAclResponseBody {
	s.Message = &v
	return s
}

func (s *AccessPageSetAclResponseBody) SetRequestId(v string) *AccessPageSetAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccessPageSetAclResponseBody) SetSuccess(v string) *AccessPageSetAclResponseBody {
	s.Success = &v
	return s
}

type AccessPageSetAclResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccessPageSetAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccessPageSetAclResponse) String() string {
	return tea.Prettify(s)
}

func (s AccessPageSetAclResponse) GoString() string {
	return s.String()
}

func (s *AccessPageSetAclResponse) SetHeaders(v map[string]*string) *AccessPageSetAclResponse {
	s.Headers = v
	return s
}

func (s *AccessPageSetAclResponse) SetStatusCode(v int32) *AccessPageSetAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AccessPageSetAclResponse) SetBody(v *AccessPageSetAclResponseBody) *AccessPageSetAclResponse {
	s.Body = v
	return s
}

type ApproveOtaTaskRequest struct {
	// The ID of the delivery group. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the region where the delivery group resides. You can call the [ListRegions](https://help.aliyun.com/document_detail/428500.html) operation to query the list of regions supported by App Streaming.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The type of the OTA update task.
	//
	// Valid values:
	//
	// 	- Fota: update of the system components of Alibaba Cloud Workspace
	//
	// 	- AppUpdate
	//
	// 	- ImageUpdate
	//
	// This parameter is required.
	//
	// example:
	//
	// Fota
	OtaType *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	// The start time of the OTA update task. The time follows the ISO 8601 standard.
	//
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-04T14:36:00+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the OTA update task. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// >  Each successful call to the `ApproveOtaTask` operation causes a value change of this parameter.`` Before you call this operation, call the `ListAppInstanceGroup` operation again to obtain the latest value of this parameter.``
	//
	// This parameter is required.
	//
	// example:
	//
	// ota-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApproveOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskRequest) SetAppInstanceGroupId(v string) *ApproveOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetBizRegionId(v string) *ApproveOtaTaskRequest {
	s.BizRegionId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetOtaType(v string) *ApproveOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetStartTime(v string) *ApproveOtaTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetTaskId(v string) *ApproveOtaTaskRequest {
	s.TaskId = &v
	return s
}

type ApproveOtaTaskResponseBody struct {
	// The execution result. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// OtaTask.Running
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. If the value of `Code` is `success`, this parameter is not returned.
	//
	// example:
	//
	// The task is running and cannot be sumitted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponseBody) SetCode(v string) *ApproveOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetMessage(v string) *ApproveOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetRequestId(v string) *ApproveOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

type ApproveOtaTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponse) SetHeaders(v map[string]*string) *ApproveOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ApproveOtaTaskResponse) SetStatusCode(v int32) *ApproveOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOtaTaskResponse) SetBody(v *ApproveOtaTaskResponseBody) *ApproveOtaTaskResponse {
	s.Body = v
	return s
}

type AskSessionPackagePriceRequest struct {
	ChargeType  *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	MaxSessions *int32  `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	Period      *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit  *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 0
	SessionPackageType *string `json:"SessionPackageType,omitempty" xml:"SessionPackageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// appstreaming.general.entry
	SessionSpec *string `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Windows
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
}

func (s AskSessionPackagePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceRequest) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceRequest) SetChargeType(v string) *AskSessionPackagePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetMaxSessions(v int32) *AskSessionPackagePriceRequest {
	s.MaxSessions = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetPeriod(v int32) *AskSessionPackagePriceRequest {
	s.Period = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetPeriodUnit(v string) *AskSessionPackagePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetRegion(v string) *AskSessionPackagePriceRequest {
	s.Region = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetSessionPackageType(v string) *AskSessionPackagePriceRequest {
	s.SessionPackageType = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetSessionSpec(v string) *AskSessionPackagePriceRequest {
	s.SessionSpec = &v
	return s
}

func (s *AskSessionPackagePriceRequest) SetSessionType(v string) *AskSessionPackagePriceRequest {
	s.SessionType = &v
	return s
}

type AskSessionPackagePriceResponseBody struct {
	Data []*AskSessionPackagePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 2C64D9E5-DFCD-10A5-A911-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AskSessionPackagePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponseBody) SetData(v []*AskSessionPackagePriceResponseBodyData) *AskSessionPackagePriceResponseBody {
	s.Data = v
	return s
}

func (s *AskSessionPackagePriceResponseBody) SetRequestId(v string) *AskSessionPackagePriceResponseBody {
	s.RequestId = &v
	return s
}

type AskSessionPackagePriceResponseBodyData struct {
	Price *AskSessionPackagePriceResponseBodyDataPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s AskSessionPackagePriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponseBodyData) SetPrice(v *AskSessionPackagePriceResponseBodyDataPrice) *AskSessionPackagePriceResponseBodyData {
	s.Price = v
	return s
}

type AskSessionPackagePriceResponseBodyDataPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 0.0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 2000.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 2000.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s AskSessionPackagePriceResponseBodyDataPrice) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponseBodyDataPrice) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetCurrency(v string) *AskSessionPackagePriceResponseBodyDataPrice {
	s.Currency = &v
	return s
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetDiscountPrice(v float32) *AskSessionPackagePriceResponseBodyDataPrice {
	s.DiscountPrice = &v
	return s
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetOriginalPrice(v float32) *AskSessionPackagePriceResponseBodyDataPrice {
	s.OriginalPrice = &v
	return s
}

func (s *AskSessionPackagePriceResponseBodyDataPrice) SetTradePrice(v float32) *AskSessionPackagePriceResponseBodyDataPrice {
	s.TradePrice = &v
	return s
}

type AskSessionPackagePriceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AskSessionPackagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AskSessionPackagePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s AskSessionPackagePriceResponse) GoString() string {
	return s.String()
}

func (s *AskSessionPackagePriceResponse) SetHeaders(v map[string]*string) *AskSessionPackagePriceResponse {
	s.Headers = v
	return s
}

func (s *AskSessionPackagePriceResponse) SetStatusCode(v int32) *AskSessionPackagePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *AskSessionPackagePriceResponse) SetBody(v *AskSessionPackagePriceResponseBody) *AskSessionPackagePriceResponse {
	s.Body = v
	return s
}

type AuthorizeInstanceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId      *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstancePersistentId *string   `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	AuthorizeUserIds        []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType        *string                                `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	UnAuthorizeUserIds []*string                              `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
	UserMeta           *AuthorizeInstanceGroupRequestUserMeta `json:"UserMeta,omitempty" xml:"UserMeta,omitempty" type:"Struct"`
}

func (s AuthorizeInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupRequest) SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAppInstancePersistentId(v string) *AuthorizeInstanceGroupRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetProductType(v string) *AuthorizeInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.UnAuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetUserMeta(v *AuthorizeInstanceGroupRequestUserMeta) *AuthorizeInstanceGroupRequest {
	s.UserMeta = v
	return s
}

type AuthorizeInstanceGroupRequestUserMeta struct {
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AuthorizeInstanceGroupRequestUserMeta) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupRequestUserMeta) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupRequestUserMeta) SetAdDomain(v string) *AuthorizeInstanceGroupRequestUserMeta {
	s.AdDomain = &v
	return s
}

func (s *AuthorizeInstanceGroupRequestUserMeta) SetType(v string) *AuthorizeInstanceGroupRequestUserMeta {
	s.Type = &v
	return s
}

type AuthorizeInstanceGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId      *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstancePersistentId *string   `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	AuthorizeUserIds        []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType        *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	UnAuthorizeUserIds []*string `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
	UserMetaShrink     *string   `json:"UserMeta,omitempty" xml:"UserMeta,omitempty"`
}

func (s AuthorizeInstanceGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAppInstancePersistentId(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupShrinkRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetProductType(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupShrinkRequest {
	s.UnAuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupShrinkRequest) SetUserMetaShrink(v string) *AuthorizeInstanceGroupShrinkRequest {
	s.UserMetaShrink = &v
	return s
}

type AuthorizeInstanceGroupResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupResponseBody) SetRequestId(v string) *AuthorizeInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupResponse) SetHeaders(v map[string]*string) *AuthorizeInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetStatusCode(v int32) *AuthorizeInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetBody(v *AuthorizeInstanceGroupResponseBody) *AuthorizeInstanceGroupResponse {
	s.Body = v
	return s
}

type BuySessionPackageRequest struct {
	AutoPay     *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ChargeType  *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	MaxSessions *int32  `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	Period      *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit  *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// p-xxxxxxxxxxxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SessionPackageName *string `json:"SessionPackageName,omitempty" xml:"SessionPackageName,omitempty"`
	// example:
	//
	// 0
	SessionPackageType *string `json:"SessionPackageType,omitempty" xml:"SessionPackageType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// appstreaming.general.entry
	SessionSpec *string `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Windows
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
}

func (s BuySessionPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s BuySessionPackageRequest) GoString() string {
	return s.String()
}

func (s *BuySessionPackageRequest) SetAutoPay(v bool) *BuySessionPackageRequest {
	s.AutoPay = &v
	return s
}

func (s *BuySessionPackageRequest) SetChargeType(v string) *BuySessionPackageRequest {
	s.ChargeType = &v
	return s
}

func (s *BuySessionPackageRequest) SetMaxSessions(v int32) *BuySessionPackageRequest {
	s.MaxSessions = &v
	return s
}

func (s *BuySessionPackageRequest) SetPeriod(v int32) *BuySessionPackageRequest {
	s.Period = &v
	return s
}

func (s *BuySessionPackageRequest) SetPeriodUnit(v string) *BuySessionPackageRequest {
	s.PeriodUnit = &v
	return s
}

func (s *BuySessionPackageRequest) SetProjectId(v string) *BuySessionPackageRequest {
	s.ProjectId = &v
	return s
}

func (s *BuySessionPackageRequest) SetRegion(v string) *BuySessionPackageRequest {
	s.Region = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionPackageName(v string) *BuySessionPackageRequest {
	s.SessionPackageName = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionPackageType(v string) *BuySessionPackageRequest {
	s.SessionPackageType = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionSpec(v string) *BuySessionPackageRequest {
	s.SessionSpec = &v
	return s
}

func (s *BuySessionPackageRequest) SetSessionType(v string) *BuySessionPackageRequest {
	s.SessionType = &v
	return s
}

type BuySessionPackageResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// There is a missing parameter.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5Fxxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// tp-xxxxxxxxxx
	SessionPackageId *int64 `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	// example:
	//
	// false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuySessionPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuySessionPackageResponseBody) GoString() string {
	return s.String()
}

func (s *BuySessionPackageResponseBody) SetCode(v string) *BuySessionPackageResponseBody {
	s.Code = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetMessage(v string) *BuySessionPackageResponseBody {
	s.Message = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetRequestId(v string) *BuySessionPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetSessionPackageId(v int64) *BuySessionPackageResponseBody {
	s.SessionPackageId = &v
	return s
}

func (s *BuySessionPackageResponseBody) SetSuccess(v string) *BuySessionPackageResponseBody {
	s.Success = &v
	return s
}

type BuySessionPackageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuySessionPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuySessionPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s BuySessionPackageResponse) GoString() string {
	return s.String()
}

func (s *BuySessionPackageResponse) SetHeaders(v map[string]*string) *BuySessionPackageResponse {
	s.Headers = v
	return s
}

func (s *BuySessionPackageResponse) SetStatusCode(v int32) *BuySessionPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *BuySessionPackageResponse) SetBody(v *BuySessionPackageResponseBody) *BuySessionPackageResponse {
	s.Body = v
	return s
}

type CreateAccessPageRequest struct {
	// This parameter is required.
	AccessPageName *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	// example:
	//
	// c-e-06gdesdaxez****
	CloudEnvId *string `json:"CloudEnvId,omitempty" xml:"CloudEnvId,omitempty"`
	// example:
	//
	// 7
	EffectTime *int32 `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// p-065zdecaer07h****
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// Day
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s CreateAccessPageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPageRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessPageRequest) SetAccessPageName(v string) *CreateAccessPageRequest {
	s.AccessPageName = &v
	return s
}

func (s *CreateAccessPageRequest) SetCloudEnvId(v string) *CreateAccessPageRequest {
	s.CloudEnvId = &v
	return s
}

func (s *CreateAccessPageRequest) SetEffectTime(v int32) *CreateAccessPageRequest {
	s.EffectTime = &v
	return s
}

func (s *CreateAccessPageRequest) SetProjectId(v string) *CreateAccessPageRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateAccessPageRequest) SetProjectName(v string) *CreateAccessPageRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateAccessPageRequest) SetUnit(v string) *CreateAccessPageRequest {
	s.Unit = &v
	return s
}

type CreateAccessPageResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a-06xnr5lyp77e7****
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccessPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessPageResponseBody) SetCode(v string) *CreateAccessPageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetData(v string) *CreateAccessPageResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetMessage(v string) *CreateAccessPageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetRequestId(v string) *CreateAccessPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessPageResponseBody) SetSuccess(v string) *CreateAccessPageResponseBody {
	s.Success = &v
	return s
}

type CreateAccessPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessPageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessPageResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessPageResponse) SetHeaders(v map[string]*string) *CreateAccessPageResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessPageResponse) SetStatusCode(v int32) *CreateAccessPageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessPageResponse) SetBody(v *CreateAccessPageResponseBody) *CreateAccessPageResponse {
	s.Body = v
	return s
}

type CreateAppInstanceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId     *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AppPackageType       *string `json:"AppPackageType,omitempty" xml:"AppPackageType,omitempty"`
	AppPolicyId          *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterId  *string                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Network    *CreateAppInstanceGroupRequestNetwork  `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	NodePool   *CreateAppInstanceGroupRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// cag-b2ron*******
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 17440009****
	PromotionId    *string                                      `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RuntimePolicy  *CreateAppInstanceGroupRequestRuntimePolicy  `json:"RuntimePolicy,omitempty" xml:"RuntimePolicy,omitempty" type:"Struct"`
	SecurityPolicy *CreateAppInstanceGroupRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 15
	SessionTimeout   *int32                                         `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StoragePolicy    *CreateAppInstanceGroupRequestStoragePolicy    `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
	SubPayType       *string                                        `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	UserDefinePolicy *CreateAppInstanceGroupRequestUserDefinePolicy `json:"UserDefinePolicy,omitempty" xml:"UserDefinePolicy,omitempty" type:"Struct"`
	UserInfo         *CreateAppInstanceGroupRequestUserInfo         `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	Users            []*string                                      `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	VideoPolicy      *CreateAppInstanceGroupRequestVideoPolicy      `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
}

func (s CreateAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppPackageType(v string) *CreateAppInstanceGroupRequest {
	s.AppPackageType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppPolicyId(v string) *CreateAppInstanceGroupRequest {
	s.AppPolicyId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAuthMode(v string) *CreateAppInstanceGroupRequest {
	s.AuthMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoPay(v bool) *CreateAppInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetBizRegionId(v string) *CreateAppInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeType(v string) *CreateAppInstanceGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetClusterId(v string) *CreateAppInstanceGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetNetwork(v *CreateAppInstanceGroupRequestNetwork) *CreateAppInstanceGroupRequest {
	s.Network = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetNodePool(v *CreateAppInstanceGroupRequestNodePool) *CreateAppInstanceGroupRequest {
	s.NodePool = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriod(v int32) *CreateAppInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPreOpenAppId(v string) *CreateAppInstanceGroupRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetProductType(v string) *CreateAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPromotionId(v string) *CreateAppInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetRuntimePolicy(v *CreateAppInstanceGroupRequestRuntimePolicy) *CreateAppInstanceGroupRequest {
	s.RuntimePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSecurityPolicy(v *CreateAppInstanceGroupRequestSecurityPolicy) *CreateAppInstanceGroupRequest {
	s.SecurityPolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetStoragePolicy(v *CreateAppInstanceGroupRequestStoragePolicy) *CreateAppInstanceGroupRequest {
	s.StoragePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSubPayType(v string) *CreateAppInstanceGroupRequest {
	s.SubPayType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUserDefinePolicy(v *CreateAppInstanceGroupRequestUserDefinePolicy) *CreateAppInstanceGroupRequest {
	s.UserDefinePolicy = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUserInfo(v *CreateAppInstanceGroupRequestUserInfo) *CreateAppInstanceGroupRequest {
	s.UserInfo = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUsers(v []*string) *CreateAppInstanceGroupRequest {
	s.Users = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetVideoPolicy(v *CreateAppInstanceGroupRequestVideoPolicy) *CreateAppInstanceGroupRequest {
	s.VideoPolicy = v
	return s
}

type CreateAppInstanceGroupRequestNetwork struct {
	DomainRules []*CreateAppInstanceGroupRequestNetworkDomainRules `json:"DomainRules,omitempty" xml:"DomainRules,omitempty" type:"Repeated"`
	// example:
	//
	// 60
	IpExpireMinutes *int32                                        `json:"IpExpireMinutes,omitempty" xml:"IpExpireMinutes,omitempty"`
	OfficeSiteId    *string                                       `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Routes          []*CreateAppInstanceGroupRequestNetworkRoutes `json:"Routes,omitempty" xml:"Routes,omitempty" type:"Repeated"`
	// example:
	//
	// Shared
	StrategyType *string   `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	VSwitchIds   []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequestNetwork) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetwork) SetDomainRules(v []*CreateAppInstanceGroupRequestNetworkDomainRules) *CreateAppInstanceGroupRequestNetwork {
	s.DomainRules = v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetIpExpireMinutes(v int32) *CreateAppInstanceGroupRequestNetwork {
	s.IpExpireMinutes = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetOfficeSiteId(v string) *CreateAppInstanceGroupRequestNetwork {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetRoutes(v []*CreateAppInstanceGroupRequestNetworkRoutes) *CreateAppInstanceGroupRequestNetwork {
	s.Routes = v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetStrategyType(v string) *CreateAppInstanceGroupRequestNetwork {
	s.StrategyType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetwork) SetVSwitchIds(v []*string) *CreateAppInstanceGroupRequestNetwork {
	s.VSwitchIds = v
	return s
}

type CreateAppInstanceGroupRequestNetworkDomainRules struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s CreateAppInstanceGroupRequestNetworkDomainRules) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetworkDomainRules) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) SetDomain(v string) *CreateAppInstanceGroupRequestNetworkDomainRules {
	s.Domain = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkDomainRules) SetPolicy(v string) *CreateAppInstanceGroupRequestNetworkDomainRules {
	s.Policy = &v
	return s
}

type CreateAppInstanceGroupRequestNetworkRoutes struct {
	// example:
	//
	// 139.196.XX.XX/32
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// example:
	//
	// Shared
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateAppInstanceGroupRequestNetworkRoutes) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNetworkRoutes) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) SetDestination(v string) *CreateAppInstanceGroupRequestNetworkRoutes {
	s.Destination = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNetworkRoutes) SetMode(v string) *CreateAppInstanceGroupRequestNetworkRoutes {
	s.Mode = &v
	return s
}

type CreateAppInstanceGroupRequestNodePool struct {
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// example:
	//
	// 10
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// example:
	//
	// 2
	NodeCapacity        *int32                                                      `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodeInstanceType    *string                                                     `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	RecurrenceSchedules []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// example:
	//
	// 2022-09-08
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// example:
	//
	// 2022-08-01
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// example:
	//
	// false
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s CreateAppInstanceGroupRequestNodePool) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePool) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePool) SetMaxIdleAppInstanceAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetMaxScalingAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeCapacity(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeInstanceType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetRecurrenceSchedules(v []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) *CreateAppInstanceGroupRequestNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingDownAfterIdleMinutes(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingStep(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingStep = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingUsageThreshold(v string) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyDisableDate(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyEnableDate(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetWarmUp(v bool) *CreateAppInstanceGroupRequestNodePool {
	s.WarmUp = &v
	return s
}

type CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules struct {
	// example:
	//
	// weekly
	RecurrenceType   *string                                                                 `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32                                                                `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	TimerPeriods     []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules) SetTimerPeriods(v []*CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods struct {
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 15:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 12:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *CreateAppInstanceGroupRequestNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type CreateAppInstanceGroupRequestRuntimePolicy struct {
	DebugMode        *string `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	PerSessionPerApp *bool   `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// DYNAMIC
	PersistentAppInstanceScheduleMode *string `json:"PersistentAppInstanceScheduleMode,omitempty" xml:"PersistentAppInstanceScheduleMode,omitempty"`
	SessionPreOpen                    *string `json:"SessionPreOpen,omitempty" xml:"SessionPreOpen,omitempty"`
	// 会话类型。
	//
	// example:
	//
	// NORMAL
	SessionType               *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	SessionUserGenerationMode *string `json:"SessionUserGenerationMode,omitempty" xml:"SessionUserGenerationMode,omitempty"`
}

func (s CreateAppInstanceGroupRequestRuntimePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestRuntimePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetDebugMode(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.DebugMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetPerSessionPerApp(v bool) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.PerSessionPerApp = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetPersistentAppInstanceScheduleMode(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.PersistentAppInstanceScheduleMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetSessionPreOpen(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.SessionPreOpen = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetSessionType(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.SessionType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestRuntimePolicy) SetSessionUserGenerationMode(v string) *CreateAppInstanceGroupRequestRuntimePolicy {
	s.SessionUserGenerationMode = &v
	return s
}

type CreateAppInstanceGroupRequestSecurityPolicy struct {
	// example:
	//
	// true
	ResetAfterUnbind *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
}

func (s CreateAppInstanceGroupRequestSecurityPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) SetResetAfterUnbind(v bool) *CreateAppInstanceGroupRequestSecurityPolicy {
	s.ResetAfterUnbind = &v
	return s
}

func (s *CreateAppInstanceGroupRequestSecurityPolicy) SetSkipUserAuthCheck(v bool) *CreateAppInstanceGroupRequestSecurityPolicy {
	s.SkipUserAuthCheck = &v
	return s
}

type CreateAppInstanceGroupRequestStoragePolicy struct {
	StorageTypeList []*string                                              `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
	UserProfile     *CreateAppInstanceGroupRequestStoragePolicyUserProfile `json:"UserProfile,omitempty" xml:"UserProfile,omitempty" type:"Struct"`
}

func (s CreateAppInstanceGroupRequestStoragePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) SetStorageTypeList(v []*string) *CreateAppInstanceGroupRequestStoragePolicy {
	s.StorageTypeList = v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicy) SetUserProfile(v *CreateAppInstanceGroupRequestStoragePolicyUserProfile) *CreateAppInstanceGroupRequestStoragePolicy {
	s.UserProfile = v
	return s
}

type CreateAppInstanceGroupRequestStoragePolicyUserProfile struct {
	RemoteStoragePath *string `json:"RemoteStoragePath,omitempty" xml:"RemoteStoragePath,omitempty"`
	RemoteStorageType *string `json:"RemoteStorageType,omitempty" xml:"RemoteStorageType,omitempty"`
	UserProfileSwitch *bool   `json:"UserProfileSwitch,omitempty" xml:"UserProfileSwitch,omitempty"`
}

func (s CreateAppInstanceGroupRequestStoragePolicyUserProfile) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestStoragePolicyUserProfile) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) SetRemoteStoragePath(v string) *CreateAppInstanceGroupRequestStoragePolicyUserProfile {
	s.RemoteStoragePath = &v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) SetRemoteStorageType(v string) *CreateAppInstanceGroupRequestStoragePolicyUserProfile {
	s.RemoteStorageType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestStoragePolicyUserProfile) SetUserProfileSwitch(v bool) *CreateAppInstanceGroupRequestStoragePolicyUserProfile {
	s.UserProfileSwitch = &v
	return s
}

type CreateAppInstanceGroupRequestUserDefinePolicy struct {
	CustomConfig *string `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
}

func (s CreateAppInstanceGroupRequestUserDefinePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestUserDefinePolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestUserDefinePolicy) SetCustomConfig(v string) *CreateAppInstanceGroupRequestUserDefinePolicy {
	s.CustomConfig = &v
	return s
}

type CreateAppInstanceGroupRequestUserInfo struct {
	// example:
	//
	// Simple
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppInstanceGroupRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestUserInfo) SetType(v string) *CreateAppInstanceGroupRequestUserInfo {
	s.Type = &v
	return s
}

type CreateAppInstanceGroupRequestVideoPolicy struct {
	FrameRate                  *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	SessionResolutionHeight    *int32  `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	SessionResolutionWidth     *int32  `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	StreamingMode              *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	TerminalResolutionAdaptive *bool   `json:"TerminalResolutionAdaptive,omitempty" xml:"TerminalResolutionAdaptive,omitempty"`
	Webrtc                     *bool   `json:"Webrtc,omitempty" xml:"Webrtc,omitempty"`
}

func (s CreateAppInstanceGroupRequestVideoPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestVideoPolicy) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetFrameRate(v int32) *CreateAppInstanceGroupRequestVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetSessionResolutionHeight(v int32) *CreateAppInstanceGroupRequestVideoPolicy {
	s.SessionResolutionHeight = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetSessionResolutionWidth(v int32) *CreateAppInstanceGroupRequestVideoPolicy {
	s.SessionResolutionWidth = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetStreamingMode(v string) *CreateAppInstanceGroupRequestVideoPolicy {
	s.StreamingMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetTerminalResolutionAdaptive(v bool) *CreateAppInstanceGroupRequestVideoPolicy {
	s.TerminalResolutionAdaptive = &v
	return s
}

func (s *CreateAppInstanceGroupRequestVideoPolicy) SetWebrtc(v bool) *CreateAppInstanceGroupRequestVideoPolicy {
	s.Webrtc = &v
	return s
}

type CreateAppInstanceGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId     *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AppPackageType       *string `json:"AppPackageType,omitempty" xml:"AppPackageType,omitempty"`
	AppPolicyId          *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// App
	AuthMode *string `json:"AuthMode,omitempty" xml:"AuthMode,omitempty"`
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	NetworkShrink  *string `json:"Network,omitempty" xml:"Network,omitempty"`
	NodePoolShrink *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// cag-b2ron*******
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 17440009****
	PromotionId          *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	RuntimePolicyShrink  *string `json:"RuntimePolicy,omitempty" xml:"RuntimePolicy,omitempty"`
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15
	SessionTimeout         *int32    `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	StoragePolicyShrink    *string   `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
	SubPayType             *string   `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	UserDefinePolicyShrink *string   `json:"UserDefinePolicy,omitempty" xml:"UserDefinePolicy,omitempty"`
	UserInfoShrink         *string   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	Users                  []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	VideoPolicyShrink      *string   `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty"`
}

func (s CreateAppInstanceGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppPackageType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppPackageType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppPolicyId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppPolicyId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAuthMode(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AuthMode = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoPay(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetBizRegionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetClusterId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetNetworkShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetNodePoolShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriod(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPreOpenAppId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetProductType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPromotionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetRuntimePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.RuntimePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSecurityPolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetStoragePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.StoragePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSubPayType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.SubPayType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUserDefinePolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.UserDefinePolicyShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUserInfoShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUsers(v []*string) *CreateAppInstanceGroupShrinkRequest {
	s.Users = v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetVideoPolicyShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.VideoPolicyShrink = &v
	return s
}

type CreateAppInstanceGroupResponseBody struct {
	AppInstanceGroupModel *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel `json:"AppInstanceGroupModel,omitempty" xml:"AppInstanceGroupModel,omitempty" type:"Struct"`
	RequestId             *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponseBody) SetAppInstanceGroupModel(v *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) *CreateAppInstanceGroupResponseBody {
	s.AppInstanceGroupModel = v
	return s
}

func (s *CreateAppInstanceGroupResponseBody) SetRequestId(v string) *CreateAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppInstanceGroupResponseBodyAppInstanceGroupModel struct {
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// rg-ew7va2g1wl3vm****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// example:
	//
	// 12345****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetAppInstanceGroupId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetNodePoolId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.NodePoolId = &v
	return s
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetOrderId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.OrderId = &v
	return s
}

type CreateAppInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetStatusCode(v int32) *CreateAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetBody(v *CreateAppInstanceGroupResponseBody) *CreateAppInstanceGroupResponse {
	s.Body = v
	return s
}

type CreateImageFromAppInstanceGroupRequest struct {
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_name
	AppCenterImageName *string `json:"AppCenterImageName,omitempty" xml:"AppCenterImageName,omitempty"`
	// The ID of the delivery group. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s CreateImageFromAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupRequest) SetAppCenterImageName(v string) *CreateImageFromAppInstanceGroupRequest {
	s.AppCenterImageName = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *CreateImageFromAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupRequest) SetProductType(v string) *CreateImageFromAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type CreateImageFromAppInstanceGroupResponseBody struct {
	// The image ID.
	//
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageFromAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupResponseBody) SetImageId(v string) *CreateImageFromAppInstanceGroupResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponseBody) SetRequestId(v string) *CreateImageFromAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageFromAppInstanceGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageFromAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageFromAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageFromAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateImageFromAppInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateImageFromAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponse) SetStatusCode(v int32) *CreateImageFromAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageFromAppInstanceGroupResponse) SetBody(v *CreateImageFromAppInstanceGroupResponseBody) *CreateImageFromAppInstanceGroupResponse {
	s.Body = v
	return s
}

type DeleteAccessPageRequest struct {
	// example:
	//
	// a-075nu7bcqim2wvxli
	AccessPageId *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
}

func (s DeleteAccessPageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPageRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessPageRequest) SetAccessPageId(v string) *DeleteAccessPageRequest {
	s.AccessPageId = &v
	return s
}

type DeleteAccessPageResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAccessPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessPageResponseBody) SetCode(v string) *DeleteAccessPageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAccessPageResponseBody) SetMessage(v string) *DeleteAccessPageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAccessPageResponseBody) SetRequestId(v string) *DeleteAccessPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessPageResponseBody) SetSuccess(v string) *DeleteAccessPageResponseBody {
	s.Success = &v
	return s
}

type DeleteAccessPageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessPageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessPageResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessPageResponse) SetHeaders(v map[string]*string) *DeleteAccessPageResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessPageResponse) SetStatusCode(v int32) *DeleteAccessPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessPageResponse) SetBody(v *DeleteAccessPageResponseBody) *DeleteAccessPageResponse {
	s.Body = v
	return s
}

type DeleteAppInstanceGroupRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DeleteAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *DeleteAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DeleteAppInstanceGroupRequest) SetProductType(v string) *DeleteAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type DeleteAppInstanceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupResponseBody) SetRequestId(v string) *DeleteAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceGroupResponse) SetHeaders(v map[string]*string) *DeleteAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstanceGroupResponse) SetStatusCode(v int32) *DeleteAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstanceGroupResponse) SetBody(v *DeleteAppInstanceGroupResponseBody) *DeleteAppInstanceGroupResponse {
	s.Body = v
	return s
}

type DeleteAppInstancesRequest struct {
	// The ID of the delivery group. You can call the [listAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The IDs of application instances.
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	AppInstanceIds []*string `json:"AppInstanceIds,omitempty" xml:"AppInstanceIds,omitempty" type:"Repeated"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DeleteAppInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesRequest) SetAppInstanceGroupId(v string) *DeleteAppInstancesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DeleteAppInstancesRequest) SetAppInstanceIds(v []*string) *DeleteAppInstancesRequest {
	s.AppInstanceIds = v
	return s
}

func (s *DeleteAppInstancesRequest) SetProductType(v string) *DeleteAppInstancesRequest {
	s.ProductType = &v
	return s
}

type DeleteAppInstancesResponseBody struct {
	// The data returned.
	DeleteAppInstanceModels []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels `json:"DeleteAppInstanceModels,omitempty" xml:"DeleteAppInstanceModels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponseBody) SetDeleteAppInstanceModels(v []*DeleteAppInstancesResponseBodyDeleteAppInstanceModels) *DeleteAppInstancesResponseBody {
	s.DeleteAppInstanceModels = v
	return s
}

func (s *DeleteAppInstancesResponseBody) SetRequestId(v string) *DeleteAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppInstancesResponseBodyDeleteAppInstanceModels struct {
	// The ID of the application instance.
	//
	// example:
	//
	// ai-gbuea*****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Specifies whether the application instance is deleted.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppInstancesResponseBodyDeleteAppInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesResponseBodyDeleteAppInstanceModels) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetAppInstanceId(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.AppInstanceId = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetCode(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Code = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetMessage(v string) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Message = &v
	return s
}

func (s *DeleteAppInstancesResponseBodyDeleteAppInstanceModels) SetSuccess(v bool) *DeleteAppInstancesResponseBodyDeleteAppInstanceModels {
	s.Success = &v
	return s
}

type DeleteAppInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInstancesResponse) SetHeaders(v map[string]*string) *DeleteAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInstancesResponse) SetStatusCode(v int32) *DeleteAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppInstancesResponse) SetBody(v *DeleteAppInstancesResponseBody) *DeleteAppInstancesResponse {
	s.Body = v
	return s
}

type GetAccessPageSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a-062wec3cwmayw****
	AccessPageId *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8141B1A674D48ACB8E5D2D6CE53FDB2F3CF8710A5F8F78578D5254BC6F******
	AccessPageToken *string `json:"AccessPageToken,omitempty" xml:"AccessPageToken,omitempty"`
	// example:
	//
	// Banca******
	ExternalUserId *string `json:"ExternalUserId,omitempty" xml:"ExternalUserId,omitempty"`
}

func (s GetAccessPageSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionRequest) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionRequest) SetAccessPageId(v string) *GetAccessPageSessionRequest {
	s.AccessPageId = &v
	return s
}

func (s *GetAccessPageSessionRequest) SetAccessPageToken(v string) *GetAccessPageSessionRequest {
	s.AccessPageToken = &v
	return s
}

func (s *GetAccessPageSessionRequest) SetExternalUserId(v string) *GetAccessPageSessionRequest {
	s.ExternalUserId = &v
	return s
}

type GetAccessPageSessionResponseBody struct {
	// example:
	//
	// 200
	Code    *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetAccessPageSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccessPageSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionResponseBody) SetCode(v string) *GetAccessPageSessionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetData(v *GetAccessPageSessionResponseBodyData) *GetAccessPageSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetMessage(v string) *GetAccessPageSessionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetRequestId(v string) *GetAccessPageSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessPageSessionResponseBody) SetSuccess(v string) *GetAccessPageSessionResponseBody {
	s.Success = &v
	return s
}

type GetAccessPageSessionResponseBodyData struct {
	// example:
	//
	// DQpbRGVza3RvcF0NCkZvcmNlVGxzVHlwZT0xDQpHV1Rva2VuUGFydDE9MDBzQU5DTGVsZ0RqMnAyMGpZdUNkRDMrNTlLekpzUTRXNElPWVdjWGIwZ2QrUkNyais0ZGM3WUJGM1NBdVFJWWl2ejhaWDlvakh1cDJ4c0Vpc3lrQ1I4RVEzeDhIYXdCb2pRRDJReklaQTZIbU52VjdrRlNlWkxQTXV3Y0hNTEZTTkxKaDNOY1BtU0tVYTNqWVVUMFUyKzNqWDBRMXdLM29ZQnJaOFVoL2RWY0xpem5tWEhHTmE0THVGRCtrajloSWlFT0w3b2w4OHY1cjBPelpHZnh4aXlVbk1yQURnRHhoK1F6K3UvdXYxaFYyQ3UwVlduTHJ3cDNRM3hYWWtORW81N00xYWZ4cTJBdWplVmx3aGVxOHd4dDl2Y3NGVGMxNFpPL2hudWdoeGJkaz0NCkdXVG9rZW5QYXJ0Mj0NCkFzcEF1dGhUb2tlbj0wME5LTnlLZXE3UGpzNzNzeURrdEs2NUlaenpYb05WNWxnY3BjVTJMd2NKUEoyMk53WmlrSGNaaDdNVGk2azlLazNyeHFCd2FHZTlmc2NUbFljUDJ0MEVldXRha2paUklMM0x3a0RsT1BZZTZxYmtsay9STHFrajg4dWg0ZFZNTHE2QkF0WVRSeW9ZTC9WUHhaVWxyZEZOQT09DQpBZXNLZXlWZXJzaW9uPTE3MDEyMzk1ODc4NzcNCkdhdGV3YXlEb21haW5BZGRyZXNzPWd3LWFwcC1jbi1oYW5nemhvdS1pLWFwcC10ZXN0MC13aW4ud3V5aW5nLmFsaXl1bmNzLmNvbQ0KR2F0ZXdheUFkZHJlc3M9NDcuOTkuMjIzLjE0OQ0KR2F0ZXdheVBvcnQ9NTkxMg0KR2F0ZXdheVdzc1BvcnQ9ODAwOA0KU3RyZWFtVHlwZT0wDQpSdGNHd1Rva2VuPVAwMFJkc013dVdQVkt4MGtRYUNkSlY3MUc2OC9iaWFhSEJwVn******
	ConnectTicket *string `json:"ConnectTicket,omitempty" xml:"ConnectTicket,omitempty"`
	// flow ID
	//
	// example:
	//
	// e4fa0b4c18c5437a8b1746e7c228172e
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetAccessPageSessionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionResponseBodyData) SetConnectTicket(v string) *GetAccessPageSessionResponseBodyData {
	s.ConnectTicket = &v
	return s
}

func (s *GetAccessPageSessionResponseBodyData) SetFlowId(v string) *GetAccessPageSessionResponseBodyData {
	s.FlowId = &v
	return s
}

type GetAccessPageSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessPageSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessPageSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessPageSessionResponse) GoString() string {
	return s.String()
}

func (s *GetAccessPageSessionResponse) SetHeaders(v map[string]*string) *GetAccessPageSessionResponse {
	s.Headers = v
	return s
}

func (s *GetAccessPageSessionResponse) SetStatusCode(v int32) *GetAccessPageSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessPageSessionResponse) SetBody(v *GetAccessPageSessionResponseBody) *GetAccessPageSessionResponse {
	s.Body = v
	return s
}

type GetAppInstanceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *GetAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetAppInstanceGroupRequest) SetProductType(v string) *GetAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type GetAppInstanceGroupResponseBody struct {
	// AppInstanceGroupModels
	AppInstanceGroupModels *GetAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Struct"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) *GetAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *GetAppInstanceGroupResponseBody) SetRequestId(v string) *GetAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// example:
	//
	// OfficeApps
	AppCenterImageName *string `json:"AppCenterImageName,omitempty" xml:"AppCenterImageName,omitempty"`
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId   *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// example:
	//
	// __dynamic__
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// example:
	//
	// test001
	AppInstanceTypeName *string `json:"AppInstanceTypeName,omitempty" xml:"AppInstanceTypeName,omitempty"`
	// example:
	//
	// pg-g3k5wa2ms2****
	AppPolicyId *string                                                      `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	Apps        []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 2022-04-27T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// 2022-04-26T15:06:16.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 10
	MaxAmount *int32 `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	// example:
	//
	// 1
	MinAmount    *int32                                                           `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	NodePool     []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	OfficeSiteId *string                                                          `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// Windows
	OsType  *string                                                       `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OtaInfo *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 20
	ReserveAmountRatio *string `json:"ReserveAmountRatio,omitempty" xml:"ReserveAmountRatio,omitempty"`
	// example:
	//
	// 5
	ReserveMaxAmount *int32 `json:"ReserveMaxAmount,omitempty" xml:"ReserveMaxAmount,omitempty"`
	// example:
	//
	// 1
	ReserveMinAmount *int32 `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
	// example:
	//
	// AVAILABLE
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// example:
	//
	// 10
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// example:
	//
	// 15
	SessionTimeout *string `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	SessionType    *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
	// example:
	//
	// spec-8o18t8uc31qib0****
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceTypeName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceTypeName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeResourceMode(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeResourceMode = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMaxAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MaxAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMinAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MinAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetNodePool(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.NodePool = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOfficeSiteId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OfficeSiteId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOsType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOtaInfo(v *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OtaInfo = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetProductType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ProductType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveAmountRatio(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveAmountRatio = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMaxAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMaxAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMinAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMinAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceStatus(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceStatus = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingDownAfterIdleMinutes(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingStep(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingStep = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingUsageThreshold(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionTimeout(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionTimeout = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSkipUserAuthCheck(v bool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSpecId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SpecId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// example:
	//
	// ca-i87mycyn419nu****
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersion     *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool struct {
	// example:
	//
	// 2
	Amount                   *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// example:
	//
	// 8
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// example:
	//
	// rg-g6922kced36hx****
	NodePoolId   *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
	// example:
	//
	// 1
	NodeUsed            *int32                                                                              `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
	RecurrenceSchedules []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// example:
	//
	// 8
	ScalingNodeAmount *int32 `json:"ScalingNodeAmount,omitempty" xml:"ScalingNodeAmount,omitempty"`
	// example:
	//
	// 4
	ScalingNodeUsed *int32 `json:"ScalingNodeUsed,omitempty" xml:"ScalingNodeUsed,omitempty"`
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// example:
	//
	// 2022-09-08
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// example:
	//
	// 2022-08-01
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// example:
	//
	// false
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxIdleAppInstanceAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxScalingAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeCapacity(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeInstanceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodePoolId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodePoolId = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeTypeName(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeTypeName = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeUsed(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeUsed = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetRecurrenceSchedules(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingDownAfterIdleMinutes(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeAmount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeUsed(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeUsed = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingStep(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingStep = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingUsageThreshold(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyDisableDate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyEnableDate(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetWarmUp(v bool) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.WarmUp = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	// example:
	//
	// Weekly
	RecurrenceType   *string                                                                                         `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32                                                                                        `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	TimerPeriods     []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetTimerPeriods(v []*GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
	// example:
	//
	// 5
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 11:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 09:30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo struct {
	// example:
	//
	// 0.0.1-D-20220630.11****
	NewOtaVersion *string `json:"NewOtaVersion,omitempty" xml:"NewOtaVersion,omitempty"`
	// example:
	//
	// 0.0.1-D-20220615.11****
	OtaVersion *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	// example:
	//
	// ota-e49929gv8acz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetNewOtaVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.NewOtaVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetOtaVersion(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.OtaVersion = &v
	return s
}

func (s *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetTaskId(v string) *GetAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.TaskId = &v
	return s
}

type GetAppInstanceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAppInstanceGroupResponse) SetHeaders(v map[string]*string) *GetAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAppInstanceGroupResponse) SetStatusCode(v int32) *GetAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInstanceGroupResponse) SetBody(v *GetAppInstanceGroupResponseBody) *GetAppInstanceGroupResponse {
	s.Body = v
	return s
}

type GetConnectionTicketRequest struct {
	// The application ID.
	//
	// >  This parameter is required for the first call to this operation and optional for subsequent calls to the operation.
	//
	// example:
	//
	// ca-e4s0puhmwi7v****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery groups.
	//
	// > 	- If you configure this parameter, the system assigns application instances only among the specified authorized delivery groups.
	//
	// > 	- This parameter is required if you configure `AppInstanceId` or `AppInstancePersistentId`.
	AppInstanceGroupIdList []*string `json:"AppInstanceGroupIdList,omitempty" xml:"AppInstanceGroupIdList,omitempty" type:"Repeated"`
	// The ID of the application instance.
	//
	// > 	- If you configure this parameter, the system attempts to assign only the specified application instance.
	//
	// > 	- If you configure this parameter, you must also configure `AppInstanceGroupIdList` and the number of delivery groups specified by `AppInstanceGroupIdList` must be 1.
	//
	// example:
	//
	// ai-1rznfnrvsa99d****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the persistent session.
	//
	// example:
	//
	// p-0bxls9m3cl7s****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The parameters that are configured to start the application. For information about how to obtain these parameters, see [Obtain parameters configured to install and start an application](https://help.aliyun.com/document_detail/426045.html).
	//
	// example:
	//
	// /q /n
	AppStartParam *string `json:"AppStartParam,omitempty" xml:"AppStartParam,omitempty"`
	// The application version. If you configure this parameter, only an application of the specified version is started. If you do not configure this parameter, an application of a random authorized version is started.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The region ID.
	//
	// >  If you configure this parameter, the system assigns application instances only among the delivery groups that reside in the specified region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The ID of the convenience account.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The product type.
	//
	// Valid values:
	//
	// 	- CloudApp: App Streaming
	//
	// 	- AndroidCloud: Cloud Phone
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The task ID.
	//
	// >  This parameter is required for calls other than the first call to this operation. You can use this parameter to query the task status and connection credential.
	//
	// example:
	//
	// 28778acb-a469-4bc0-8e0f****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetAppId(v string) *GetConnectionTicketRequest {
	s.AppId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceGroupIdList(v []*string) *GetConnectionTicketRequest {
	s.AppInstanceGroupIdList = v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceId(v string) *GetConnectionTicketRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstancePersistentId(v string) *GetConnectionTicketRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppStartParam(v string) *GetConnectionTicketRequest {
	s.AppStartParam = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppVersion(v string) *GetConnectionTicketRequest {
	s.AppVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetBizRegionId(v string) *GetConnectionTicketRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetProductType(v string) *GetConnectionTicketRequest {
	s.ProductType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-53fvrq1oan****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// ai-7ybdeiyoeh5e****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the persistent session.
	//
	// example:
	//
	// p-0bxls9m3cl7s****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The operating system.
	//
	// Valid value:
	//
	// 	- Windows: the Windows operating system
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// f3d1b31c-605e-4d04-a896****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status.
	//
	// Valid values:
	//
	// 	- Finished: The task is complete.
	//
	// 	- Failed: The task failed.
	//
	// 	- Running: The task is being executed.
	//
	// example:
	//
	// Running
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 148871678899****
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The credential that is used to connect to App Streaming.
	//
	// >  This parameter is displayed for calls other than the first call to this operation.
	//
	// example:
	//
	// DQpbRGVza3RvcF0NCkZvcmNlVGxzVHlwZT0xDQpHV1Rva2VuUGFydDE9MDAva09ROW1FUTU3dU****
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetBizRegionId(v string) *GetConnectionTicketResponseBody {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetOsType(v string) *GetConnectionTicketResponseBody {
	s.OsType = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTenantId(v int64) *GetConnectionTicketResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetStatusCode(v int32) *GetConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type GetDebugAppInstanceRequest struct {
	// The ID of the delivery group. You can call the `listAppInstanceGroup` operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetDebugAppInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDebugAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceRequest) SetAppInstanceGroupId(v string) *GetDebugAppInstanceRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetDebugAppInstanceRequest) SetProductType(v string) *GetDebugAppInstanceRequest {
	s.ProductType = &v
	return s
}

type GetDebugAppInstanceResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// __DEBUG_APP
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// ai-7ybdeiyoeh5e****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the application version.
	//
	// example:
	//
	// 1.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The authorization code. This authorization code is valid for 3 minutes and can be used only once, regardless of whether the authentication succeeds. If multiple authentication codes are generated for a user, only the latest authentication code takes effect.
	//
	// example:
	//
	// e4e169bea1cc48e8afac53**********
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// __debug__
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDebugAppInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDebugAppInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceResponseBody) SetAppId(v string) *GetDebugAppInstanceResponseBody {
	s.AppId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppInstanceGroupId(v string) *GetDebugAppInstanceResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppInstanceId(v string) *GetDebugAppInstanceResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAppVersion(v string) *GetDebugAppInstanceResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetAuthCode(v string) *GetDebugAppInstanceResponseBody {
	s.AuthCode = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetRequestId(v string) *GetDebugAppInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDebugAppInstanceResponseBody) SetUserId(v string) *GetDebugAppInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetDebugAppInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDebugAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDebugAppInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDebugAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetDebugAppInstanceResponse) SetHeaders(v map[string]*string) *GetDebugAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetDebugAppInstanceResponse) SetStatusCode(v int32) *GetDebugAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDebugAppInstanceResponse) SetBody(v *GetDebugAppInstanceResponseBody) *GetDebugAppInstanceResponse {
	s.Body = v
	return s
}

type GetOtaTaskByTaskIdRequest struct {
	// The ID of the OTA update task. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ota-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetOtaTaskByTaskIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdRequest) SetTaskId(v string) *GetOtaTaskByTaskIdRequest {
	s.TaskId = &v
	return s
}

type GetOtaTaskByTaskIdResponseBody struct {
	// The error code.
	//
	// example:
	//
	// OtaTask.Running
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The task is running and cannot be sumitted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The OTA version.
	//
	// example:
	//
	// 0.0.1-R-20220708.110604
	OtaVersion *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	// The version description.
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution time of the OTA update task. The time follows the ISO 8601 standard.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-04T14:36:00+08:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s GetOtaTaskByTaskIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponseBody) SetCode(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetMessage(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetOtaVersion(v string) *GetOtaTaskByTaskIdResponseBody {
	s.OtaVersion = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetReleaseNote(v string) *GetOtaTaskByTaskIdResponseBody {
	s.ReleaseNote = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetRequestId(v string) *GetOtaTaskByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetTaskStartTime(v string) *GetOtaTaskByTaskIdResponseBody {
	s.TaskStartTime = &v
	return s
}

type GetOtaTaskByTaskIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOtaTaskByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOtaTaskByTaskIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponse) SetHeaders(v map[string]*string) *GetOtaTaskByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetStatusCode(v int32) *GetOtaTaskByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetBody(v *GetOtaTaskByTaskIdResponseBody) *GetOtaTaskByTaskIdResponse {
	s.Body = v
	return s
}

type GetResourcePriceRequest struct {
	// The number of resources to purchase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The type ID of the sessions that you purchase. You can call the `ListAppInstanceType` operation to obtain the ID.
	//
	// You must specify one of AppInstanceType and NodeInstanceType. If you specify both of the parameters, the value of NodeInstanceType takes effect.
	//
	// example:
	//
	// appstreaming.general
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai).
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the resource type that you purchase. You can call the [ListNodeInstanceType](https://help.aliyun.com/document_detail/428502.html) to obtain the ID.
	//
	// You must specify one of AppInstanceType and NodeInstanceType. If you specify both of the parameters, the value of NodeInstanceType takes effect.
	//
	// Valid values:
	//
	// 	- appstreaming.vgpu.8c16g.4g: WUYING - Graphics - 8 vCPUs, 16 GiB Memory, 4 GiB GPU Memory
	//
	// 	- appstreaming.general.8c16g: WUYING - General - 8 vCPUs, 16 GiB Memory
	//
	// 	- appstreaming.general.4c8g: WUYING - General - 4 vCPUs, 8 GiB Memory
	//
	// 	- appstreaming.vgpu.14c93g.12g: WUYING - Graphics - 14 vCPUs, 93 GiB Memory, 12 GiB GPU Memory.
	//
	// 	- appstreaming.vgpu.8c31g.16g: WUYING - Graphics - 8 vCPUs, 31 GiB Memory, 16 GiB GPU Memory
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The subscription duration of resources. This parameter must be configured together with `PeriodUnit`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. This parameter must be configured together with `Period`. The following items describe valid values for the combinations of `Period` and `PeriodUnit`:
	//
	// 	- 1 Week
	//
	// 	- 1 Month
	//
	// 	- 2 Month
	//
	// 	- 3 Month
	//
	// 	- 6 Month
	//
	// 	- 1 Year
	//
	// 	- 2 Year
	//
	// 	- 3 Year
	//
	// >  The value of this parameter is case-insensitive. For example, `Week` is valid and `week` is invalid. If you specify a value combination other than the preceding combinations, such as `2 Week`, the operation can still be called. However, an error occurs when you place the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetResourcePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePriceRequest) SetAmount(v int64) *GetResourcePriceRequest {
	s.Amount = &v
	return s
}

func (s *GetResourcePriceRequest) SetAppInstanceType(v string) *GetResourcePriceRequest {
	s.AppInstanceType = &v
	return s
}

func (s *GetResourcePriceRequest) SetBizRegionId(v string) *GetResourcePriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetResourcePriceRequest) SetChargeType(v string) *GetResourcePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *GetResourcePriceRequest) SetNodeInstanceType(v string) *GetResourcePriceRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriod(v int64) *GetResourcePriceRequest {
	s.Period = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriodUnit(v string) *GetResourcePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetResourcePriceRequest) SetProductType(v string) *GetResourcePriceRequest {
	s.ProductType = &v
	return s
}

type GetResourcePriceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The price objects.
	//
	// This parameter is returned only if a value is specified for AppInstanceType.
	PriceList []*GetResourcePriceResponseBodyPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The price object.
	//
	// This parameter is returned only if a value is specified for NodeInstanceType.
	PriceModel *GetResourcePriceResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourcePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBody) SetCode(v string) *GetResourcePriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetMessage(v string) *GetResourcePriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetPriceList(v []*GetResourcePriceResponseBodyPriceList) *GetResourcePriceResponseBody {
	s.PriceList = v
	return s
}

func (s *GetResourcePriceResponseBody) SetPriceModel(v *GetResourcePriceResponseBodyPriceModel) *GetResourcePriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *GetResourcePriceResponseBody) SetRequestId(v string) *GetResourcePriceResponseBody {
	s.RequestId = &v
	return s
}

type GetResourcePriceResponseBodyPriceList struct {
	// The price details.
	Price *GetResourcePriceResponseBodyPriceListPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The price type.
	//
	// Valid values:
	//
	// 	- Connected: in use
	//
	// 	- Standby: pending for use.
	//
	// example:
	//
	// Standby
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	// The price calculation rules.
	Rules []*GetResourcePriceResponseBodyPriceListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourcePriceResponseBodyPriceList) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceList) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceList) SetPrice(v *GetResourcePriceResponseBodyPriceListPrice) *GetResourcePriceResponseBodyPriceList {
	s.Price = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceList) SetPriceType(v string) *GetResourcePriceResponseBodyPriceList {
	s.PriceType = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceList) SetRules(v []*GetResourcePriceResponseBodyPriceListRules) *GetResourcePriceResponseBodyPriceList {
	s.Rules = v
	return s
}

type GetResourcePriceResponseBodyPriceListPrice struct {
	// The currency type.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 999.0
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 6700
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The coupon metadata.
	Promotions []*GetResourcePriceResponseBodyPriceListPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 5278.0
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListPrice) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListPrice) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetCurrency(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.Currency = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetDiscountPrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetOriginalPrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetPromotions(v []*GetResourcePriceResponseBodyPriceListPricePromotions) *GetResourcePriceResponseBodyPriceListPrice {
	s.Promotions = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetTradePrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.TradePrice = &v
	return s
}

type GetResourcePriceResponseBodyPriceListPricePromotions struct {
	// The coupon code.
	//
	// example:
	//
	// coupon****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The coupon description.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 1847709****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The coupon name.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether the coupon was used.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetOptionCode(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionDesc(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionId(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionName(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetSelected(v bool) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.Selected = &v
	return s
}

type GetResourcePriceResponseBodyPriceListRules struct {
	// The description of the price calculation rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the price calculation rule.
	//
	// example:
	//
	// 260904273633****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListRules) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListRules) SetDescription(v string) *GetResourcePriceResponseBodyPriceListRules {
	s.Description = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListRules) SetRuleId(v int64) *GetResourcePriceResponseBodyPriceListRules {
	s.RuleId = &v
	return s
}

type GetResourcePriceResponseBodyPriceModel struct {
	// The price details.
	Price *GetResourcePriceResponseBodyPriceModelPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The price calculation rules.
	Rules []*GetResourcePriceResponseBodyPriceModelRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourcePriceResponseBodyPriceModel) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModel) SetPrice(v *GetResourcePriceResponseBodyPriceModelPrice) *GetResourcePriceResponseBodyPriceModel {
	s.Price = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModel) SetRules(v []*GetResourcePriceResponseBodyPriceModelRules) *GetResourcePriceResponseBodyPriceModel {
	s.Rules = v
	return s
}

type GetResourcePriceResponseBodyPriceModelPrice struct {
	// The currency type.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 1.00
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 11.00
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The coupon metadata.
	Promotions []*GetResourcePriceResponseBodyPriceModelPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 10.00
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPrice) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPrice) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetCurrency(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Currency = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetDiscountPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetOriginalPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetPromotions(v []*GetResourcePriceResponseBodyPriceModelPricePromotions) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Promotions = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetTradePrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.TradePrice = &v
	return s
}

type GetResourcePriceResponseBodyPriceModelPricePromotions struct {
	// The coupon code.
	//
	// example:
	//
	// coupon****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The coupon description.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The coupon name.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether the coupon was used.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetOptionCode(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionDesc(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionId(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionName(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetSelected(v bool) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.Selected = &v
	return s
}

type GetResourcePriceResponseBodyPriceModelRules struct {
	// The description of the price calculation rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the price calculation rule.
	//
	// example:
	//
	// 102002100393****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelRules) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetDescription(v string) *GetResourcePriceResponseBodyPriceModelRules {
	s.Description = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetRuleId(v int64) *GetResourcePriceResponseBodyPriceModelRules {
	s.RuleId = &v
	return s
}

type GetResourcePriceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponse) SetHeaders(v map[string]*string) *GetResourcePriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePriceResponse) SetStatusCode(v int32) *GetResourcePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePriceResponse) SetBody(v *GetResourcePriceResponseBody) *GetResourcePriceResponse {
	s.Body = v
	return s
}

type GetResourceRenewPriceRequest struct {
	// The ID of the delivery group. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The subscription duration of resources. This parameter must be configured together with `PeriodUnit`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. This parameter must be configured together with `Period`. The following items describe valid values for the combinations of `Period` and `PeriodUnit`:
	//
	// 	- 1 Week
	//
	// 	- 1 Month
	//
	// 	- 2 Month
	//
	// 	- 3 Month
	//
	// 	- 6 Month
	//
	// 	- 1 Year
	//
	// 	- 2 Year
	//
	// 	- 3 Year
	//
	// >  The value of this parameter is case-insensitive. For example, `Week` is valid and `week` is invalid. If you specify a value combination other than the preceding combinations, such as `2 Week`, the operation can still be called. However, an error occurs when you place the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetResourceRenewPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceRequest) SetAppInstanceGroupId(v string) *GetResourceRenewPriceRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetPeriod(v int64) *GetResourceRenewPriceRequest {
	s.Period = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetPeriodUnit(v string) *GetResourceRenewPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetResourceRenewPriceRequest) SetProductType(v string) *GetResourceRenewPriceRequest {
	s.ProductType = &v
	return s
}

type GetResourceRenewPriceResponseBody struct {
	// The price object.
	Data *GetResourceRenewPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceRenewPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBody) SetData(v *GetResourceRenewPriceResponseBodyData) *GetResourceRenewPriceResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceRenewPriceResponseBody) SetRequestId(v string) *GetResourceRenewPriceResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceRenewPriceResponseBodyData struct {
	// The price details.
	Price *GetResourceRenewPriceResponseBodyDataPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The price calculation rules.
	Rules []*GetResourceRenewPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourceRenewPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyData) SetPrice(v *GetResourceRenewPriceResponseBodyDataPrice) *GetResourceRenewPriceResponseBodyData {
	s.Price = v
	return s
}

func (s *GetResourceRenewPriceResponseBodyData) SetRules(v []*GetResourceRenewPriceResponseBodyDataRules) *GetResourceRenewPriceResponseBodyData {
	s.Rules = v
	return s
}

type GetResourceRenewPriceResponseBodyDataPrice struct {
	// The currency type.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 1
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 11
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The coupon description.
	Promotions []*GetResourceRenewPriceResponseBodyDataPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 10
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataPrice) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataPrice) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetCurrency(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.Currency = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetDiscountPrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetOriginalPrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetPromotions(v []*GetResourceRenewPriceResponseBodyDataPricePromotions) *GetResourceRenewPriceResponseBodyDataPrice {
	s.Promotions = v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetTradePrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.TradePrice = &v
	return s
}

type GetResourceRenewPriceResponseBodyDataPricePromotions struct {
	// The coupon code.
	//
	// example:
	//
	// coupon****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The coupon description.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 139965*****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The coupon name.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether the coupon was used.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetOptionCode(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionDesc(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionId(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionName(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetSelected(v bool) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.Selected = &v
	return s
}

type GetResourceRenewPriceResponseBodyDataRules struct {
	// The description of the price calculation rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the price calculation rule.
	//
	// example:
	//
	// 20002****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataRules) SetDescription(v string) *GetResourceRenewPriceResponseBodyDataRules {
	s.Description = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataRules) SetRuleId(v int64) *GetResourceRenewPriceResponseBodyDataRules {
	s.RuleId = &v
	return s
}

type GetResourceRenewPriceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceRenewPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceRenewPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRenewPriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponse) SetHeaders(v map[string]*string) *GetResourceRenewPriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceRenewPriceResponse) SetStatusCode(v int32) *GetResourceRenewPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceRenewPriceResponse) SetBody(v *GetResourceRenewPriceResponseBody) *GetResourceRenewPriceResponse {
	s.Body = v
	return s
}

type ListAccessPagesRequest struct {
	// example:
	//
	// a-062wec3cwmayw****
	AccessPageId   *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	AccessPageName *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// p-065zdecaer07h****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// ASC
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListAccessPagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesRequest) GoString() string {
	return s.String()
}

func (s *ListAccessPagesRequest) SetAccessPageId(v string) *ListAccessPagesRequest {
	s.AccessPageId = &v
	return s
}

func (s *ListAccessPagesRequest) SetAccessPageName(v string) *ListAccessPagesRequest {
	s.AccessPageName = &v
	return s
}

func (s *ListAccessPagesRequest) SetPageNumber(v int32) *ListAccessPagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAccessPagesRequest) SetPageSize(v int32) *ListAccessPagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAccessPagesRequest) SetProjectId(v string) *ListAccessPagesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListAccessPagesRequest) SetSortType(v string) *ListAccessPagesRequest {
	s.SortType = &v
	return s
}

type ListAccessPagesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Count *string                            `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*ListAccessPagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InternalError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AF8361BD-5ECB-139A-B019-2E0350CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAccessPagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessPagesResponseBody) SetCode(v string) *ListAccessPagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetCount(v string) *ListAccessPagesResponseBody {
	s.Count = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetData(v []*ListAccessPagesResponseBodyData) *ListAccessPagesResponseBody {
	s.Data = v
	return s
}

func (s *ListAccessPagesResponseBody) SetMessage(v string) *ListAccessPagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetPageNumber(v string) *ListAccessPagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetPageSize(v string) *ListAccessPagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetRequestId(v string) *ListAccessPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessPagesResponseBody) SetSuccess(v string) *ListAccessPagesResponseBody {
	s.Success = &v
	return s
}

type ListAccessPagesResponseBodyData struct {
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// example:
	//
	// a-062wec3cwmayw****
	AccessPageId *string `json:"AccessPageId,omitempty" xml:"AccessPageId,omitempty"`
	// example:
	//
	// notepad_test
	AccessPageName *string `json:"AccessPageName,omitempty" xml:"AccessPageName,omitempty"`
	// example:
	//
	// 1
	AccessPageState *string `json:"AccessPageState,omitempty" xml:"AccessPageState,omitempty"`
	// example:
	//
	// id=a-062wec3cwmayw****&token=9E9A62937B0E41F4AEFE5EC9B238156CCDFB682954003AEE940A05FB2568****
	AccessUrl *string `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
	// example:
	//
	// c-05to6wm3w5d53****
	ContentId   *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	ContentName *string `json:"ContentName,omitempty" xml:"ContentName,omitempty"`
	// example:
	//
	// 168
	EffectTime *int32 `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// example:
	//
	// 2023-11-16T08:48:15.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// p-062wec3cwmayu****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// notepad_demo
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// Hour
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// example:
	//
	// 2023-11-23T08:48:15.000+00:00
	UrlExpireTime *string `json:"UrlExpireTime,omitempty" xml:"UrlExpireTime,omitempty"`
}

func (s ListAccessPagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAccessPagesResponseBodyData) SetAccessMode(v string) *ListAccessPagesResponseBodyData {
	s.AccessMode = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessPageId(v string) *ListAccessPagesResponseBodyData {
	s.AccessPageId = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessPageName(v string) *ListAccessPagesResponseBodyData {
	s.AccessPageName = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessPageState(v string) *ListAccessPagesResponseBodyData {
	s.AccessPageState = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetAccessUrl(v string) *ListAccessPagesResponseBodyData {
	s.AccessUrl = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetContentId(v string) *ListAccessPagesResponseBodyData {
	s.ContentId = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetContentName(v string) *ListAccessPagesResponseBodyData {
	s.ContentName = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetEffectTime(v int32) *ListAccessPagesResponseBodyData {
	s.EffectTime = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetGmtCreate(v string) *ListAccessPagesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetProjectId(v string) *ListAccessPagesResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetProjectName(v string) *ListAccessPagesResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetUnit(v string) *ListAccessPagesResponseBodyData {
	s.Unit = &v
	return s
}

func (s *ListAccessPagesResponseBodyData) SetUrlExpireTime(v string) *ListAccessPagesResponseBodyData {
	s.UrlExpireTime = &v
	return s
}

type ListAccessPagesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessPagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessPagesResponse) GoString() string {
	return s.String()
}

func (s *ListAccessPagesResponse) SetHeaders(v map[string]*string) *ListAccessPagesResponse {
	s.Headers = v
	return s
}

func (s *ListAccessPagesResponse) SetStatusCode(v int32) *ListAccessPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessPagesResponse) SetBody(v *ListAccessPagesResponseBody) *ListAccessPagesResponse {
	s.Body = v
	return s
}

type ListAppInstanceGroupRequest struct {
	// The image ID of the app. You can obtain the ID from the Images page in the App Streaming console.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery groups to query. Fuzzy match is used for queries. For example, if you set this parameter to `Office App`, all delivery groups whose names contain `Office App` are queried, such as `My Office Apps` and `Office App A`.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The ID of the resource specification that you purchase. You can call the [ListNodeInstanceType](~~ListNodeInstanceType~~) operation to obtain the ID.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// example:
	//
	// cn-hongkong+dir-643067****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value cannot be greater than `100`.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Deprecated
	//
	// The region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the delivery groups.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupRequest) SetAppCenterImageId(v string) *ListAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetBizRegionId(v string) *ListAppInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetNodeInstanceType(v string) *ListAppInstanceGroupRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetOfficeSiteId(v string) *ListAppInstanceGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageNumber(v int32) *ListAppInstanceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageSize(v int32) *ListAppInstanceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetProductType(v string) *ListAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetRegionId(v string) *ListAppInstanceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetStatus(v []*string) *ListAppInstanceGroupRequest {
	s.Status = v
	return s
}

type ListAppInstanceGroupResponseBody struct {
	// The delivery groups.
	AppInstanceGroupModels []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels) *ListAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageNumber(v int32) *ListAppInstanceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageSize(v int32) *ListAppInstanceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetRequestId(v string) *ListAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetTotalCount(v int32) *ListAppInstanceGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	// The number of subscription resources. Minimum value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The image ID of the app.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The resource type of the delivery group.
	//
	// example:
	//
	// __dynamic__
	AppInstanceType *string `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-g3k5wa2ms2****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// example:
	//
	// false
	AppPolicyImageCheck *bool `json:"AppPolicyImageCheck,omitempty" xml:"AppPolicyImageCheck,omitempty"`
	// example:
	//
	// CENTER
	AppPolicyVersion *string `json:"AppPolicyVersion,omitempty" xml:"AppPolicyVersion,omitempty"`
	// The apps.
	Apps []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// The sales mode.
	//
	// Valid values:
	//
	// 	- AppInstance: by session
	//
	// 	- Node: by resource
	//
	// example:
	//
	// Node
	ChargeResourceMode *string `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the delivery group expires.
	//
	// example:
	//
	// 2022-04-27T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The time when the delivery group was created.
	//
	// example:
	//
	// 2022-04-26T15:06:16.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The maximum number of instances. Minimum value: 1.
	//
	// example:
	//
	// 10
	MaxAmount *int32 `json:"MaxAmount,omitempty" xml:"MaxAmount,omitempty"`
	// The minimum number of instances. Minimum value: 1.
	//
	// example:
	//
	// 1
	MinAmount *int32 `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	// The resource groups.
	NodePool []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing+dir-172301****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The type of the operating system.
	//
	// Valid value:
	//
	// 	- Windows
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The information about the over-the-air (OTA) update task.
	OtaInfo *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The percentage of reserved instances. The value indicates the percentage of unused sessions in the delivery group. Valid values: 0 to 99.
	//
	// example:
	//
	// 20
	ReserveAmountRatio *string `json:"ReserveAmountRatio,omitempty" xml:"ReserveAmountRatio,omitempty"`
	// The maximum number of reserved instances. The value indicates the maximum number of unused sessions in the delivery group. Minimum value: 1.
	//
	// example:
	//
	// 5
	ReserveMaxAmount *int32 `json:"ReserveMaxAmount,omitempty" xml:"ReserveMaxAmount,omitempty"`
	// The minimum number of reserved instances. The value indicates the minimum number of unused sessions in the delivery group. Minimum value: 1.
	//
	// example:
	//
	// 1
	ReserveMinAmount *int32 `json:"ReserveMinAmount,omitempty" xml:"ReserveMinAmount,omitempty"`
	// The resource status.
	//
	// Valid values:
	//
	// 	- AVAILABLE
	//
	// 	- RELEASED
	//
	// 	- EXPIRED_IN_7_DAYS
	//
	// 	- UNAVAILABLE
	//
	// 	- UPGRADING
	//
	// 	- CREATING
	//
	// example:
	//
	// AVAILABLE
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The resource tags.
	ResourceTags []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Repeated"`
	// The duration for which no session is connected. Unit: minutes. If no session is connected in the resources after the specified duration elapses, auto scale-in is triggered. Minimum value: 0.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The number of sessions that are created each time the delivery group is scaled out. Minimum value: 1.
	//
	// example:
	//
	// 10
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper limit of session usage. If the session usage exceeds the specified upper limit, auto scale-out is triggered. The session usage rate is calculated by using the following formula: Session usage rate = Number of sessions in use/Total number of sessions × 100%. Valid values: 0 to 99.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The duration for which sessions are retained after disconnection. Unit: minutes. After an end user disconnects from a session, the session is closed only after the specified duration elapses. If you want to permanently retain sessions, set this parameter to `-1`. Valid values:-1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *string `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// Indicates whether user permission verification is skipped.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: This is the default value.
	//
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
	// The specification ID that uniquely corresponds to the ID of the delivery group.
	//
	// example:
	//
	// spec-8o18t8uc31qib0****
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// The status of the delivery group.
	//
	// Valid values:
	//
	// 	- PUBLISHED: The delivery group is published.
	//
	// 	- FAILED: The delivery group failed to be published.
	//
	// 	- MAINTAIN_FAILED: The delivery group failed to be updated.
	//
	// 	- EXPIRED: The delivery group is expired.
	//
	// 	- MAINTAINING: The delivery group is being updated.
	//
	// 	- CEASED: The delivery group has overdue payments.
	//
	// 	- EXPIRED_RECYCLING: The delivery group is expired and being recycled.
	//
	// 	- DEPLOYING: The delivery group is being published.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyImageCheck(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyImageCheck = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppPolicyVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppPolicyVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeResourceMode(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeResourceMode = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMaxAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MaxAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetMinAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.MinAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetNodePool(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.NodePool = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOfficeSiteId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OfficeSiteId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOsType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOtaInfo(v *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OtaInfo = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetProductType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveAmountRatio(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveAmountRatio = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMaxAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMaxAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetReserveMinAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ReserveMinAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceStatus(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceStatus = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetResourceTags(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ResourceTags = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingDownAfterIdleMinutes(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingStep(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingStep = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetScalingUsageThreshold(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionTimeout(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionTimeout = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSkipUserAuthCheck(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSpecId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SpecId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	// The app icon.
	//
	// example:
	//
	// https://app-center-icon-****.png
	AppIcon *string `json:"AppIcon,omitempty" xml:"AppIcon,omitempty"`
	// The app ID.
	//
	// example:
	//
	// ca-i87mycyn419nu****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The app name.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The app version.
	//
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The name of the app version.
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppIcon(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppIcon = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppVersionName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppVersionName = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool struct {
	// The number of resources purchased when the delivery group was created.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The maximum number of idle sessions. After you specify a value for this parameter, auto scale-out is triggered only if the number of idle sessions in the delivery group is smaller than the specified value and the session usage exceeds the value specified for `ScalingUsageThreshold`. Otherwise, the system determines that idle sessions in the delivery group are sufficient and does not perform auto scale-out.`` You can use this parameter to flexibly manage auto scaling and reduce costs.
	//
	// example:
	//
	// 3
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// The maximum number of resources that can be created for scale-out.
	//
	// example:
	//
	// 8
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// The total number of subscription resources.
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The maximum number of sessions that can be connected to a resource at the same time. If a resource connects to a large number of sessions at the same time, user experience can be compromised. The value range varies based on the resource specification. The following items describe the value ranges of different resource types:
	//
	// 	- appstreaming.general.4c8g: 1 to 2
	//
	// 	- appstreaming.general.8c16g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c16g.4g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c31g.16g: 1 to 4
	//
	// 	- appstreaming.vgpu.14c93g.12g: 1 to 6
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The ID of the resource specification that you purchase.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-g6922kced36hx****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	// The name of the resource specification.
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
	// The number of subscription resources that are in use.
	//
	// example:
	//
	// 1
	NodeUsed *int32 `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
	// The intervals at which the scaling policy is executed.
	RecurrenceSchedules []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// The duration for which no session is connected. Unit: minutes. If no session is connected in the resources after the specified duration elapses, auto scale-in is triggered. Default value: 5.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The total number of scalable resources.
	//
	// example:
	//
	// 8
	ScalingNodeAmount *int32 `json:"ScalingNodeAmount,omitempty" xml:"ScalingNodeAmount,omitempty"`
	// The number of scalable resources that are in use.
	//
	// example:
	//
	// 4
	ScalingNodeUsed *int32 `json:"ScalingNodeUsed,omitempty" xml:"ScalingNodeUsed,omitempty"`
	// The number of resources that are created each time resources are scaled out. Valid values: 1 to 10.
	//
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper limit of session usage. If the session usage exceeds the specified upper limit, auto scale-out is triggered. The session usage is calculated by using the following formula: `Session usage = Number of current sessions/(Total number of resources × Number of concurrent sessions) × 100%`.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The expiration date of the scaling policy. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2022-09-08
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// The effective date of the scaling policy. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2022-08-01
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// The type of the scaling policy.
	//
	// >  `NODE_SCALING_BY_USAGE` is returned for this parameter only if ChargeType is set to `PrePaid`. `NODE_SCALING_BY_SCHEDULE` is returned for this parameter only if ChargeType is set to `PostPaid`.
	//
	// Valid values:
	//
	// 	- NODE_FIXED: No scalable resources are used.
	//
	// 	- NODE_SCALING_BY_SCHEDULE: Scheduled scaling is used.
	//
	// 	- NODE_SCALING_BY_USAGE: Resources are scaled based on usage.
	//
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// Indicates whether resource prefetch is enabled.
	//
	// example:
	//
	// false
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxIdleAppInstanceAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxScalingAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeCapacity(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodePoolId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodePoolId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeTypeName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeTypeName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetRecurrenceSchedules(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingDownAfterIdleMinutes(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingStep(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingStep = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingUsageThreshold(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyDisableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyEnableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetWarmUp(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.WarmUp = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	// The schedule type of the scaling policy. This parameter must be configured together with `RecurrenceValues`.``
	//
	// Valid value:
	//
	// 	- weekly: The scaling policy is executed on specific days each week.
	//
	// example:
	//
	// Weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The days of each week on which the scaling policy is executed.
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The time periods during which the scaling policy can be executed.
	TimerPeriods []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetTimerPeriods(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
	// The number of destination resources.
	//
	// example:
	//
	// 5
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The time when the scaling policy ends. Format: HH:mm.
	//
	// example:
	//
	// 11:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the scaling policy starts. Format: HH:mm.
	//
	// example:
	//
	// 09:30
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo struct {
	// The new OTA version. A null value indicates that no new version is available.
	//
	// example:
	//
	// 0.0.1-D-20220630.11****
	NewOtaVersion *string `json:"NewOtaVersion,omitempty" xml:"NewOtaVersion,omitempty"`
	// The current OTA version.
	//
	// example:
	//
	// 0.0.1-D-20220615.11****
	OtaVersion *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	// The ID of the OTA update task.
	//
	// example:
	//
	// ota-e49929gv8acz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetNewOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.NewOtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.OtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetTaskId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.TaskId = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag type. Valid values: Custom System
	//
	// example:
	//
	// Custom
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag value.
	//
	// example:
	//
	// design
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) SetKey(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags {
	s.Key = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) SetScope(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags {
	s.Scope = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags) SetValue(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsResourceTags {
	s.Value = &v
	return s
}

type ListAppInstanceGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponse) SetHeaders(v map[string]*string) *ListAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstanceGroupResponse) SetStatusCode(v int32) *ListAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstanceGroupResponse) SetBody(v *ListAppInstanceGroupResponseBody) *ListAppInstanceGroupResponse {
	s.Body = v
	return s
}

type ListAppInstancesRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-4p5f8tj16yb8b****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// ai-azn3kmwruh1vl****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The IDs of the application instances. Up to 100 IDs can be specified.
	//
	// if can be null:
	// false
	AppInstanceIdList []*string `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to query the information about deleted app instances. If you set this parameter to true, you must configure AppInstanceIdList. Otherwise, a parameter error is reported.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IncludeDeleted *bool `json:"IncludeDeleted,omitempty" xml:"IncludeDeleted,omitempty"`
	// The page number. Default value: `1`. We recommend that you specify this parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value cannot be greater than `100`. Default value: `20`. We recommend that you specify this parameter.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the application instances.
	//
	// if can be null:
	// false
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The user IDs. You can specify up to 100 IDs.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s ListAppInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstancesRequest) SetAppInstanceGroupId(v string) *ListAppInstancesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstancesRequest) SetAppInstanceId(v string) *ListAppInstancesRequest {
	s.AppInstanceId = &v
	return s
}

func (s *ListAppInstancesRequest) SetAppInstanceIdList(v []*string) *ListAppInstancesRequest {
	s.AppInstanceIdList = v
	return s
}

func (s *ListAppInstancesRequest) SetIncludeDeleted(v bool) *ListAppInstancesRequest {
	s.IncludeDeleted = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageNumber(v int32) *ListAppInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageSize(v int32) *ListAppInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesRequest) SetStatus(v []*string) *ListAppInstancesRequest {
	s.Status = v
	return s
}

func (s *ListAppInstancesRequest) SetUserIdList(v []*string) *ListAppInstancesRequest {
	s.UserIdList = v
	return s
}

type ListAppInstancesResponseBody struct {
	// The app instances.
	AppInstanceModels []*ListAppInstancesResponseBodyAppInstanceModels `json:"AppInstanceModels,omitempty" xml:"AppInstanceModels,omitempty" type:"Repeated"`
	// The page number of the returned page. We recommend that you configure this parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. The value cannot be greater than `100`. We recommend that you configure this parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 18
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBody) SetAppInstanceModels(v []*ListAppInstancesResponseBodyAppInstanceModels) *ListAppInstancesResponseBody {
	s.AppInstanceModels = v
	return s
}

func (s *ListAppInstancesResponseBody) SetPageNumber(v int32) *ListAppInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetPageSize(v int32) *ListAppInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetRequestId(v string) *ListAppInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstancesResponseBody) SetTotalCount(v int32) *ListAppInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppInstancesResponseBodyAppInstanceModels struct {
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-dk8p95irqfst9****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the application instance.
	//
	// example:
	//
	// ai-8dl7dzchklmka****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The information about the binding between the application instance and end users.
	BindInfo *ListAppInstancesResponseBodyAppInstanceModelsBindInfo `json:"BindInfo,omitempty" xml:"BindInfo,omitempty" type:"Struct"`
	// The billing method of the app instance. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the app instance belongs is set to Node.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the application instance was created.
	//
	// example:
	//
	// 2023-03-07T20:29:19.000+08:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the application instance was updated.
	//
	// example:
	//
	// 2023-03-07T20:29:19.000+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The public IP address associated with the primary NIC. This value is returned only if `StrategyType` is set to `Mixed`.
	//
	// example:
	//
	// 10.13.13.211
	MainEthPublicIp    *string `json:"MainEthPublicIp,omitempty" xml:"MainEthPublicIp,omitempty"`
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// The ID of the node on which the app instance runs.
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the app instance belongs is set to Node.
	//
	// example:
	//
	// i-bp13********
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The session status. This parameter is returned only if the application instance is in the `RUNNING` state.
	//
	// Valid values:
	//
	// 	- disconnect: disconnected
	//
	// 	- connect: connected
	//
	// example:
	//
	// connect
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// The status of the application instance.
	//
	// example:
	//
	// BOUND
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInstancesResponseBodyAppInstanceModels) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponseBodyAppInstanceModels) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetAppInstanceGroupId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetAppInstanceId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.AppInstanceId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetBindInfo(v *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) *ListAppInstancesResponseBodyAppInstanceModels {
	s.BindInfo = v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetChargeType(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.ChargeType = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetGmtCreate(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetGmtModified(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.GmtModified = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetMainEthPublicIp(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.MainEthPublicIp = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetNetworkInterfaceIp(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetNodeId(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.NodeId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetSessionStatus(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.SessionStatus = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModels) SetStatus(v string) *ListAppInstancesResponseBodyAppInstanceModels {
	s.Status = &v
	return s
}

type ListAppInstancesResponseBodyAppInstanceModelsBindInfo struct {
	// The ID of the end user that is bound to the application instance.
	//
	// example:
	//
	// app.test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The use duration of the application instance. Unit: seconds.
	//
	// example:
	//
	// 2000
	UsageDuration *int64 `json:"UsageDuration,omitempty" xml:"UsageDuration,omitempty"`
}

func (s ListAppInstancesResponseBodyAppInstanceModelsBindInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponseBodyAppInstanceModelsBindInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) SetEndUserId(v string) *ListAppInstancesResponseBodyAppInstanceModelsBindInfo {
	s.EndUserId = &v
	return s
}

func (s *ListAppInstancesResponseBodyAppInstanceModelsBindInfo) SetUsageDuration(v int64) *ListAppInstancesResponseBodyAppInstanceModelsBindInfo {
	s.UsageDuration = &v
	return s
}

type ListAppInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstancesResponse) SetHeaders(v map[string]*string) *ListAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstancesResponse) SetStatusCode(v int32) *ListAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstancesResponse) SetBody(v *ListAppInstancesResponseBody) *ListAppInstancesResponse {
	s.Body = v
	return s
}

type ListBindInfoRequest struct {
	AppIdList              []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	AppInstanceGroupIdList []*string `json:"AppInstanceGroupIdList,omitempty" xml:"AppInstanceGroupIdList,omitempty" type:"Repeated"`
	AppInstanceIdList      []*string `json:"AppInstanceIdList,omitempty" xml:"AppInstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
	WyIdList   []*string `json:"WyIdList,omitempty" xml:"WyIdList,omitempty" type:"Repeated"`
}

func (s ListBindInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBindInfoRequest) GoString() string {
	return s.String()
}

func (s *ListBindInfoRequest) SetAppIdList(v []*string) *ListBindInfoRequest {
	s.AppIdList = v
	return s
}

func (s *ListBindInfoRequest) SetAppInstanceGroupIdList(v []*string) *ListBindInfoRequest {
	s.AppInstanceGroupIdList = v
	return s
}

func (s *ListBindInfoRequest) SetAppInstanceIdList(v []*string) *ListBindInfoRequest {
	s.AppInstanceIdList = v
	return s
}

func (s *ListBindInfoRequest) SetPageNumber(v int32) *ListBindInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBindInfoRequest) SetPageSize(v int32) *ListBindInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListBindInfoRequest) SetUserIdList(v []*string) *ListBindInfoRequest {
	s.UserIdList = v
	return s
}

func (s *ListBindInfoRequest) SetWyIdList(v []*string) *ListBindInfoRequest {
	s.WyIdList = v
	return s
}

type ListBindInfoResponseBody struct {
	BindInfoModels []*ListBindInfoResponseBodyBindInfoModels `json:"BindInfoModels,omitempty" xml:"BindInfoModels,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169-D3F8D19C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBindInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBindInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindInfoResponseBody) SetBindInfoModels(v []*ListBindInfoResponseBodyBindInfoModels) *ListBindInfoResponseBody {
	s.BindInfoModels = v
	return s
}

func (s *ListBindInfoResponseBody) SetPageNumber(v int32) *ListBindInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBindInfoResponseBody) SetPageSize(v int32) *ListBindInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBindInfoResponseBody) SetRequestId(v string) *ListBindInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindInfoResponseBody) SetTotalCount(v int32) *ListBindInfoResponseBody {
	s.TotalCount = &v
	return s
}

type ListBindInfoResponseBodyBindInfoModels struct {
	// example:
	//
	// simple
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// ca-fq738or6vd854****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// aig-0abxhr6ce35w8****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// example:
	//
	// ai-83oe276fre4l3****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Alice
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 2ca6f5a93536****
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ListBindInfoResponseBodyBindInfoModels) String() string {
	return tea.Prettify(s)
}

func (s ListBindInfoResponseBodyBindInfoModels) GoString() string {
	return s.String()
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAccountType(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AccountType = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppInstanceGroupId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppInstanceId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppInstanceId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetAppVersion(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.AppVersion = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetProductType(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.ProductType = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetRegionId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.RegionId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetUserId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.UserId = &v
	return s
}

func (s *ListBindInfoResponseBodyBindInfoModels) SetWyId(v string) *ListBindInfoResponseBodyBindInfoModels {
	s.WyId = &v
	return s
}

type ListBindInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBindInfoResponse) GoString() string {
	return s.String()
}

func (s *ListBindInfoResponse) SetHeaders(v map[string]*string) *ListBindInfoResponse {
	s.Headers = v
	return s
}

func (s *ListBindInfoResponse) SetStatusCode(v int32) *ListBindInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindInfoResponse) SetBody(v *ListBindInfoResponseBody) *ListBindInfoResponse {
	s.Body = v
	return s
}

type ListNodeInstanceTypeRequest struct {
	// The ID of the region where the resource resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string  `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	Cpu         *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu         *float32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	GpuMemory   *int32   `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The language that you want to use.
	//
	// Valid values:
	//
	// 	- en-US: English (US)
	//
	// 	- zh-CN: Simplified Chinese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Memory   *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The resource type that you want to query. If you do not configure this parameter, all resource types are returned.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType       *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodeInstanceTypeFamily *string `json:"NodeInstanceTypeFamily,omitempty" xml:"NodeInstanceTypeFamily,omitempty"`
	OrderBy                *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The operating system that is supported.
	//
	// Valid value:
	//
	// 	- Windows: the Windows operating system
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SortType    *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListNodeInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeRequest) SetBizRegionId(v string) *ListNodeInstanceTypeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetCpu(v float32) *ListNodeInstanceTypeRequest {
	s.Cpu = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetGpu(v float32) *ListNodeInstanceTypeRequest {
	s.Gpu = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetGpuMemory(v int32) *ListNodeInstanceTypeRequest {
	s.GpuMemory = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetLanguage(v string) *ListNodeInstanceTypeRequest {
	s.Language = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetMemory(v int32) *ListNodeInstanceTypeRequest {
	s.Memory = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetNodeInstanceType(v string) *ListNodeInstanceTypeRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetNodeInstanceTypeFamily(v string) *ListNodeInstanceTypeRequest {
	s.NodeInstanceTypeFamily = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetOrderBy(v string) *ListNodeInstanceTypeRequest {
	s.OrderBy = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetOsType(v string) *ListNodeInstanceTypeRequest {
	s.OsType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageNumber(v int32) *ListNodeInstanceTypeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageSize(v int32) *ListNodeInstanceTypeRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetProductType(v string) *ListNodeInstanceTypeRequest {
	s.ProductType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetSortType(v string) *ListNodeInstanceTypeRequest {
	s.SortType = &v
	return s
}

type ListNodeInstanceTypeResponseBody struct {
	// The resource types.
	NodeInstanceTypeModels []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels `json:"NodeInstanceTypeModels,omitempty" xml:"NodeInstanceTypeModels,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBody) SetNodeInstanceTypeModels(v []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) *ListNodeInstanceTypeResponseBody {
	s.NodeInstanceTypeModels = v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageNumber(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageSize(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetRequestId(v string) *ListNodeInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetTotalCount(v int32) *ListNodeInstanceTypeResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 4
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 2
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The GPU size. Unit: MB.
	//
	// example:
	//
	// 8192
	GpuMemory *int64 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The maximum number of sessions to which a resource can connect at the same time. If a resource connects to a large number of sessions at the same time, user experience can be compromised. The value range varies based on the resource type. The following items describe the value ranges of different resource types:
	//
	// 	- appstreaming.general.4c8g: 1 to 2
	//
	// 	- appstreaming.general.8c16g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c16g.4g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c31g.16g: 1 to 4
	//
	// 	- appstreaming.vgpu.14c93g.12g: 1 to 6
	//
	// example:
	//
	// 4
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The memory size. Unit: MB.
	//
	// example:
	//
	// 8192
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ID of the resource type.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The resource type family.
	//
	// Valid values:
	//
	// 	- appstreaming.general: WUYING - General
	//
	// 	- appstreaming.vgpu: WUYING - Graphics
	//
	// example:
	//
	// appstreaming.vgpu
	NodeInstanceTypeFamily *string `json:"NodeInstanceTypeFamily,omitempty" xml:"NodeInstanceTypeFamily,omitempty"`
	// The name of the resource type.
	//
	// example:
	//
	// WUYING - General - 4 vCPUs 8 GB Memory
	NodeTypeName *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetCpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Cpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Gpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpuMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.GpuMemory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMaxCapacity(v int32) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.MaxCapacity = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Memory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceType(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceType = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceTypeFamily(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceTypeFamily = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeTypeName(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeTypeName = &v
	return s
}

type ListNodeInstanceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponse) SetHeaders(v map[string]*string) *ListNodeInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetStatusCode(v int32) *ListNodeInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetBody(v *ListNodeInstanceTypeResponseBody) *ListNodeInstanceTypeResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-53fvrq1oanz6c****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetAppInstanceGroupId(v string) *ListNodesRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetProductType(v string) *ListNodesRequest {
	s.ProductType = &v
	return s
}

type ListNodesResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The resource nodes.
	NodeModels []*ListNodesResponseBodyNodeModels `json:"NodeModels,omitempty" xml:"NodeModels,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PerPageSize *int32 `json:"PerPageSize,omitempty" xml:"PerPageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	ToPage *int32 `json:"ToPage,omitempty" xml:"ToPage,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetCount(v int32) *ListNodesResponseBody {
	s.Count = &v
	return s
}

func (s *ListNodesResponseBody) SetNodeModels(v []*ListNodesResponseBodyNodeModels) *ListNodesResponseBody {
	s.NodeModels = v
	return s
}

func (s *ListNodesResponseBody) SetPerPageSize(v int32) *ListNodesResponseBody {
	s.PerPageSize = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetToPage(v int32) *ListNodesResponseBody {
	s.ToPage = &v
	return s
}

type ListNodesResponseBodyNodeModels struct {
	// The billing method of the resource node.
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the resource node belongs is set to Node.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- Prepaid: subscription
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the resource node.
	//
	// >  This parameter is returned only if the ChargeResourceMode parameter of the delivery group to which the resource node belongs is set to Node.
	//
	// example:
	//
	// i-bp13********
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListNodesResponseBodyNodeModels) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodeModels) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodeModels) SetChargeType(v string) *ListNodesResponseBodyNodeModels {
	s.ChargeType = &v
	return s
}

func (s *ListNodesResponseBodyNodeModels) SetNodeId(v string) *ListNodesResponseBodyNodeModels {
	s.NodeId = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListOtaTaskRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-53fvrq1oanz6c****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The type of the OTA update task.
	//
	// Valid values:
	//
	// 	- Fota: update of the system components of Alibaba Cloud Workspace
	//
	// This parameter is required.
	//
	// example:
	//
	// Fota
	OtaType *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	// The page number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOtaTaskRequest) SetAppInstanceGroupId(v string) *ListOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListOtaTaskRequest) SetOtaType(v string) *ListOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageNumber(v int32) *ListOtaTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageSize(v int32) *ListOtaTaskRequest {
	s.PageSize = &v
	return s
}

type ListOtaTaskResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The OTA update tasks.
	TaskList []*ListOtaTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of OTA update tasks.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBody) SetPageNumber(v int32) *ListOtaTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetPageSize(v int32) *ListOtaTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetRequestId(v string) *ListOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetTaskList(v []*ListOtaTaskResponseBodyTaskList) *ListOtaTaskResponseBody {
	s.TaskList = v
	return s
}

func (s *ListOtaTaskResponseBody) SetTotalCount(v int32) *ListOtaTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListOtaTaskResponseBodyTaskList struct {
	// The OTA version.
	//
	// example:
	//
	// 0.0.1-R-20220708.110604
	OtaVersion *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	// The task status.
	//
	// Valid values:
	//
	// 	- FAILED
	//
	// 	- RUNNING
	//
	// 	- TERMINATED
	//
	// 	- PART_FINISHED
	//
	// 	- STANDBY
	//
	// 	- FINISHED
	//
	// example:
	//
	// RUNNING
	TaskDisplayStatus *string `json:"TaskDisplayStatus,omitempty" xml:"TaskDisplayStatus,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ota-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The start time of the OTA update task. The time follows the ISO 8601 standard.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-04T14:36:00+08:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s ListOtaTaskResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBodyTaskList) SetOtaVersion(v string) *ListOtaTaskResponseBodyTaskList {
	s.OtaVersion = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskDisplayStatus(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskDisplayStatus = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskId(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskStartTime(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskStartTime = &v
	return s
}

type ListOtaTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponse) SetHeaders(v map[string]*string) *ListOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOtaTaskResponse) SetStatusCode(v int32) *ListOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOtaTaskResponse) SetBody(v *ListOtaTaskResponseBody) *ListOtaTaskResponse {
	s.Body = v
	return s
}

type ListRegionsRequest struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	BizSource *string `json:"BizSource,omitempty" xml:"BizSource,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) SetBizSource(v string) *ListRegionsRequest {
	s.BizSource = &v
	return s
}

func (s *ListRegionsRequest) SetProductType(v string) *ListRegionsRequest {
	s.ProductType = &v
	return s
}

type ListRegionsResponseBody struct {
	// The region IDs.
	RegionModels []*ListRegionsResponseBodyRegionModels `json:"RegionModels,omitempty" xml:"RegionModels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegionModels(v []*ListRegionsResponseBodyRegionModels) *ListRegionsResponseBody {
	s.RegionModels = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegionModels struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListRegionsResponseBodyRegionModels) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegionModels) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionModels) SetRegionId(v string) *ListRegionsResponseBodyRegionModels {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSessionPackagesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// p-xxxxxxxxxxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// tp-xxxxxxxx
	SessionPackageId   *string `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	SessionPackageName *string `json:"SessionPackageName,omitempty" xml:"SessionPackageName,omitempty"`
	// example:
	//
	// ASC
	SortType  *string  `json:"SortType,omitempty" xml:"SortType,omitempty"`
	StateList []*int32 `json:"StateList,omitempty" xml:"StateList,omitempty" type:"Repeated"`
}

func (s ListSessionPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesRequest) SetPageNumber(v int32) *ListSessionPackagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSessionPackagesRequest) SetPageSize(v int32) *ListSessionPackagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSessionPackagesRequest) SetProjectId(v string) *ListSessionPackagesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListSessionPackagesRequest) SetSessionPackageId(v string) *ListSessionPackagesRequest {
	s.SessionPackageId = &v
	return s
}

func (s *ListSessionPackagesRequest) SetSessionPackageName(v string) *ListSessionPackagesRequest {
	s.SessionPackageName = &v
	return s
}

func (s *ListSessionPackagesRequest) SetSortType(v string) *ListSessionPackagesRequest {
	s.SortType = &v
	return s
}

func (s *ListSessionPackagesRequest) SetStateList(v []*int32) *ListSessionPackagesRequest {
	s.StateList = v
	return s
}

type ListSessionPackagesResponseBody struct {
	Data []*ListSessionPackagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5Fxxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSessionPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponseBody) SetData(v []*ListSessionPackagesResponseBodyData) *ListSessionPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionPackagesResponseBody) SetPageNumber(v int32) *ListSessionPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSessionPackagesResponseBody) SetPageSize(v int32) *ListSessionPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSessionPackagesResponseBody) SetRequestId(v string) *ListSessionPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionPackagesResponseBody) SetTotalCount(v int64) *ListSessionPackagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListSessionPackagesResponseBodyData struct {
	// example:
	//
	// 500
	AvailableHours *int64  `json:"AvailableHours,omitempty" xml:"AvailableHours,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 0
	DeleteStatus *int32 `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	// example:
	//
	// 1701170196000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1701170203000
	GmtModifiedTime *string                                            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InstanceObject  *ListSessionPackagesResponseBodyDataInstanceObject `json:"InstanceObject,omitempty" xml:"InstanceObject,omitempty" type:"Struct"`
	// example:
	//
	// 1000
	MaxHours *int64 `json:"MaxHours,omitempty" xml:"MaxHours,omitempty"`
	// example:
	//
	// 100
	MaxSessions *int64 `json:"MaxSessions,omitempty" xml:"MaxSessions,omitempty"`
	// example:
	//
	// p-xxxxxxxxxxxx
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// tp-xxxxxxxxx
	SessionPackageId   *string `json:"SessionPackageId,omitempty" xml:"SessionPackageId,omitempty"`
	SessionPackageName *string `json:"SessionPackageName,omitempty" xml:"SessionPackageName,omitempty"`
	// example:
	//
	// 0
	SessionPackageTypeId *string `json:"SessionPackageTypeId,omitempty" xml:"SessionPackageTypeId,omitempty"`
	// example:
	//
	// appstreaming.general.entry
	SessionSpec *string `json:"SessionSpec,omitempty" xml:"SessionSpec,omitempty"`
	// example:
	//
	// 60
	SessionUsageRate *int64 `json:"SessionUsageRate,omitempty" xml:"SessionUsageRate,omitempty"`
	// example:
	//
	// 4
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 156xxxxxxxxxx
	UserIdentification *int64 `json:"UserIdentification,omitempty" xml:"UserIdentification,omitempty"`
}

func (s ListSessionPackagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponseBodyData) SetAvailableHours(v int64) *ListSessionPackagesResponseBodyData {
	s.AvailableHours = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetChargeType(v string) *ListSessionPackagesResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetDeleteStatus(v int32) *ListSessionPackagesResponseBodyData {
	s.DeleteStatus = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetGmtCreate(v string) *ListSessionPackagesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetGmtModifiedTime(v string) *ListSessionPackagesResponseBodyData {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetInstanceObject(v *ListSessionPackagesResponseBodyDataInstanceObject) *ListSessionPackagesResponseBodyData {
	s.InstanceObject = v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetMaxHours(v int64) *ListSessionPackagesResponseBodyData {
	s.MaxHours = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetMaxSessions(v int64) *ListSessionPackagesResponseBodyData {
	s.MaxSessions = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetProjectId(v string) *ListSessionPackagesResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetProjectName(v string) *ListSessionPackagesResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetRegion(v string) *ListSessionPackagesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionPackageId(v string) *ListSessionPackagesResponseBodyData {
	s.SessionPackageId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionPackageName(v string) *ListSessionPackagesResponseBodyData {
	s.SessionPackageName = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionPackageTypeId(v string) *ListSessionPackagesResponseBodyData {
	s.SessionPackageTypeId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionSpec(v string) *ListSessionPackagesResponseBodyData {
	s.SessionSpec = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetSessionUsageRate(v int64) *ListSessionPackagesResponseBodyData {
	s.SessionUsageRate = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetState(v int32) *ListSessionPackagesResponseBodyData {
	s.State = &v
	return s
}

func (s *ListSessionPackagesResponseBodyData) SetUserIdentification(v int64) *ListSessionPackagesResponseBodyData {
	s.UserIdentification = &v
	return s
}

type ListSessionPackagesResponseBodyDataInstanceObject struct {
	// example:
	//
	// 2024-05-28T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// tp-xxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// appstreaming.general.entry
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// p-xxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// AppSessionPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 2023-11-28T04:14:07Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 3600000
	TotalTime *int64 `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	// example:
	//
	// 1021
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s ListSessionPackagesResponseBodyDataInstanceObject) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponseBodyDataInstanceObject) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetExpiredTime(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.ExpiredTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetInstanceId(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.InstanceId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetInstanceType(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.InstanceType = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetResourceId(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.ResourceId = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetResourceType(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.ResourceType = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetStartTime(v string) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.StartTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetTotalTime(v int64) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.TotalTime = &v
	return s
}

func (s *ListSessionPackagesResponseBodyDataInstanceObject) SetUsedTime(v int64) *ListSessionPackagesResponseBodyDataInstanceObject {
	s.UsedTime = &v
	return s
}

type ListSessionPackagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSessionPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListSessionPackagesResponse) SetHeaders(v map[string]*string) *ListSessionPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListSessionPackagesResponse) SetStatusCode(v int32) *ListSessionPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionPackagesResponse) SetBody(v *ListSessionPackagesResponseBody) *ListSessionPackagesResponse {
	s.Body = v
	return s
}

type ListTagCloudResourcesRequest struct {
	// The number of entries per page. Maximum value: 1000. Default value: 50.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ptnJAAAAAAAxNzE5OTEwNQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs. You do not need to specify this parameter if you set ResourceType to AliUid.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the cloud resource.
	//
	// Valid values:
	//
	// 	- AppId: app ID.
	//
	// 	- WyId: Alibaba Cloud Workspace user ID.
	//
	// 	- AppInstanceGroupId: delivery group ID.
	//
	// 	- AliUid: tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag type.
	//
	// Valid values:
	//
	// 	- All (default): all tags.
	//
	// 	- Custom: custom tag.
	//
	// 	- System: system tag.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Custom
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListTagCloudResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesRequest) SetMaxResults(v int32) *ListTagCloudResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagCloudResourcesRequest) SetNextToken(v string) *ListTagCloudResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagCloudResourcesRequest) SetResourceIds(v []*string) *ListTagCloudResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagCloudResourcesRequest) SetResourceType(v string) *ListTagCloudResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagCloudResourcesRequest) SetScope(v string) *ListTagCloudResourcesRequest {
	s.Scope = &v
	return s
}

type ListTagCloudResourcesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// example:
	//
	// AAAAAYRHtOLVQzCYj17y+OP7LZRrUJaF4rnBGQkWwMiVHlLZBB1w3Us37CVvhvyM0TXavA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags added to the cloud resources.
	ResourceTags []*ListTagCloudResourcesResponseBodyResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagCloudResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponseBody) SetNextToken(v string) *ListTagCloudResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagCloudResourcesResponseBody) SetRequestId(v string) *ListTagCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagCloudResourcesResponseBody) SetResourceTags(v []*ListTagCloudResourcesResponseBodyResourceTags) *ListTagCloudResourcesResponseBody {
	s.ResourceTags = v
	return s
}

func (s *ListTagCloudResourcesResponseBody) SetTotalCount(v int32) *ListTagCloudResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagCloudResourcesResponseBodyResourceTags struct {
	// The resource ID.
	//
	// example:
	//
	// aig-0001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the cloud resource.
	//
	// Valid values:
	//
	// 	- AppId: app ID.
	//
	// 	- WyId: Alibaba Cloud Workspace user ID.
	//
	// 	- AppInstanceGroupId: delivery group ID.
	//
	// 	- AliUid: tenant ID.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*ListTagCloudResourcesResponseBodyResourceTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagCloudResourcesResponseBodyResourceTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagCloudResourcesResponseBodyResourceTags) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) SetResourceId(v string) *ListTagCloudResourcesResponseBodyResourceTags {
	s.ResourceId = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) SetResourceType(v string) *ListTagCloudResourcesResponseBodyResourceTags {
	s.ResourceType = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTags) SetTags(v []*ListTagCloudResourcesResponseBodyResourceTagsTags) *ListTagCloudResourcesResponseBodyResourceTags {
	s.Tags = v
	return s
}

type ListTagCloudResourcesResponseBodyResourceTagsTags struct {
	// The tag key.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag type.
	//
	// Valid values:
	//
	// 	- Custom: custom tag.
	//
	// 	- System: system tag.
	//
	// example:
	//
	// Custom
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1080p
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagCloudResourcesResponseBodyResourceTagsTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagCloudResourcesResponseBodyResourceTagsTags) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) SetKey(v string) *ListTagCloudResourcesResponseBodyResourceTagsTags {
	s.Key = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) SetScope(v string) *ListTagCloudResourcesResponseBodyResourceTagsTags {
	s.Scope = &v
	return s
}

func (s *ListTagCloudResourcesResponseBodyResourceTagsTags) SetValue(v string) *ListTagCloudResourcesResponseBodyResourceTagsTags {
	s.Value = &v
	return s
}

type ListTagCloudResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagCloudResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponse) SetHeaders(v map[string]*string) *ListTagCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagCloudResourcesResponse) SetStatusCode(v int32) *ListTagCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagCloudResourcesResponse) SetBody(v *ListTagCloudResourcesResponseBody) *ListTagCloudResourcesResponse {
	s.Body = v
	return s
}

type ListTenantConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user configurations.
	TenantConfigModel *ListTenantConfigResponseBodyTenantConfigModel `json:"TenantConfigModel,omitempty" xml:"TenantConfigModel,omitempty" type:"Struct"`
}

func (s ListTenantConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTenantConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponseBody) SetRequestId(v string) *ListTenantConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantConfigResponseBody) SetTenantConfigModel(v *ListTenantConfigResponseBodyTenantConfigModel) *ListTenantConfigResponseBody {
	s.TenantConfigModel = v
	return s
}

type ListTenantConfigResponseBodyTenantConfigModel struct {
	// Indicates whether the resource expiration reminder feature is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AppInstanceGroupExpireRemind *bool `json:"AppInstanceGroupExpireRemind,omitempty" xml:"AppInstanceGroupExpireRemind,omitempty"`
}

func (s ListTenantConfigResponseBodyTenantConfigModel) String() string {
	return tea.Prettify(s)
}

func (s ListTenantConfigResponseBodyTenantConfigModel) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponseBodyTenantConfigModel) SetAppInstanceGroupExpireRemind(v bool) *ListTenantConfigResponseBodyTenantConfigModel {
	s.AppInstanceGroupExpireRemind = &v
	return s
}

type ListTenantConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTenantConfigResponse) GoString() string {
	return s.String()
}

func (s *ListTenantConfigResponse) SetHeaders(v map[string]*string) *ListTenantConfigResponse {
	s.Headers = v
	return s
}

func (s *ListTenantConfigResponse) SetStatusCode(v int32) *ListTenantConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantConfigResponse) SetBody(v *ListTenantConfigResponseBody) *ListTenantConfigResponse {
	s.Body = v
	return s
}

type LogOffAllSessionsInAppInstanceGroupRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *LogOffAllSessionsInAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupRequest) SetProductType(v string) *LogOffAllSessionsInAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

type LogOffAllSessionsInAppInstanceGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetCode(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetMessage(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponseBody) SetRequestId(v string) *LogOffAllSessionsInAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type LogOffAllSessionsInAppInstanceGroupResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogOffAllSessionsInAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogOffAllSessionsInAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s LogOffAllSessionsInAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetHeaders(v map[string]*string) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetStatusCode(v int32) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *LogOffAllSessionsInAppInstanceGroupResponse) SetBody(v *LogOffAllSessionsInAppInstanceGroupResponseBody) *LogOffAllSessionsInAppInstanceGroupResponse {
	s.Body = v
	return s
}

type ModifyAppInstanceGroupAttributeRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The network settings.
	//
	// >  If you want to use this parameter, submit a ticket.
	Network *ModifyAppInstanceGroupAttributeRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The information about the resource group.
	NodePool *ModifyAppInstanceGroupAttributeRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// Specifies whether only one application can be opened in a session.
	//
	// 	- After you enable this feature, the system assigns a session to each application if you open multiple applications in a delivery group. This consumes a larger number of sessions.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	PerSessionPerApp *bool `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// The application ID of the pre-open application. If you set `PreOpenMode` to `SINGLE_APP`, you cannot leave this parameter empty.``
	//
	// example:
	//
	// ca-b2ronxxd****
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// The pre-open mode.
	//
	// Valid values:
	//
	// 	- SINGLE_APP: enables the pre-open mode for a single application.
	//
	// 	- OFF: disables the pre-open mode. This is the default value.
	//
	// example:
	//
	// OFF
	PreOpenMode *string `json:"PreOpenMode,omitempty" xml:"PreOpenMode,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The security policy.
	SecurityPolicy *ModifyAppInstanceGroupAttributeRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The duration for which sessions are retained after disconnection. Unit: minutes. After an end user disconnects from a session, the session is closed only after the specified duration elapses. If you want to permanently retain sessions, set this parameter to `-1`. Valid values:-1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The storage policy.
	StoragePolicy *ModifyAppInstanceGroupAttributeRequestStoragePolicy `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
}

func (s ModifyAppInstanceGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetNetwork(v *ModifyAppInstanceGroupAttributeRequestNetwork) *ModifyAppInstanceGroupAttributeRequest {
	s.Network = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetNodePool(v *ModifyAppInstanceGroupAttributeRequestNodePool) *ModifyAppInstanceGroupAttributeRequest {
	s.NodePool = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPerSessionPerApp(v bool) *ModifyAppInstanceGroupAttributeRequest {
	s.PerSessionPerApp = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.PreOpenMode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetSecurityPolicy(v *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) *ModifyAppInstanceGroupAttributeRequest {
	s.SecurityPolicy = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeRequest {
	s.SessionTimeout = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetStoragePolicy(v *ModifyAppInstanceGroupAttributeRequestStoragePolicy) *ModifyAppInstanceGroupAttributeRequest {
	s.StoragePolicy = v
	return s
}

type ModifyAppInstanceGroupAttributeRequestNetwork struct {
	// The domain name rules.
	DomainRules []*ModifyAppInstanceGroupAttributeRequestNetworkDomainRules `json:"DomainRules,omitempty" xml:"DomainRules,omitempty" type:"Repeated"`
}

func (s ModifyAppInstanceGroupAttributeRequestNetwork) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNetwork) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNetwork) SetDomainRules(v []*ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) *ModifyAppInstanceGroupAttributeRequestNetwork {
	s.DomainRules = v
	return s
}

type ModifyAppInstanceGroupAttributeRequestNetworkDomainRules struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The policy used for the domain name.
	//
	// Valid values:
	//
	// 	- allow
	//
	// 	- block
	//
	// example:
	//
	// block
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) SetDomain(v string) *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules {
	s.Domain = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) SetPolicy(v string) *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules {
	s.Policy = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestNodePool struct {
	// The maximum number of sessions to which a resource can connect at the same time. If a resource connects to a large number of sessions at the same time, user experience can be compromised. The value range varies based on the resource type. The following items describe the value ranges of different resource types:
	//
	// 	- appstreaming.general.4c8g: 1 to 2
	//
	// 	- appstreaming.general.8c16g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c16g.4g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c31g.16g: 1 to 4
	//
	// 	- appstreaming.vgpu.14c93g.12g: 1 to 6
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-ew7va2g1wl3vm****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodeCapacity(v int32) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodePoolId(v string) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodePoolId = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestSecurityPolicy struct {
	// Specifies whether to reset after unbinding from a delivery group.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ResetAfterUnbind *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
	// Specifies whether to skip user permission verification.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: This is the default value.
	//
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestSecurityPolicy) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) SetResetAfterUnbind(v bool) *ModifyAppInstanceGroupAttributeRequestSecurityPolicy {
	s.ResetAfterUnbind = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) SetSkipUserAuthCheck(v bool) *ModifyAppInstanceGroupAttributeRequestSecurityPolicy {
	s.SkipUserAuthCheck = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicy struct {
	// The storage types.
	StorageTypeList []*string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
	// The configurations of user data roaming.
	UserProfile       *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile       `json:"UserProfile,omitempty" xml:"UserProfile,omitempty" type:"Struct"`
	UserProfileFollow *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow `json:"UserProfileFollow,omitempty" xml:"UserProfileFollow,omitempty" type:"Struct"`
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicy) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) SetStorageTypeList(v []*string) *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	s.StorageTypeList = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) SetUserProfile(v *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	s.UserProfile = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) SetUserProfileFollow(v *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	s.UserProfileFollow = v
	return s
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile struct {
	// The ID of the File Storage NAS (NAS) file system used to store user data.
	//
	// example:
	//
	// 06ae94****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Specifies whether user data roaming is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	UserProfileSwitch *bool `json:"UserProfileSwitch,omitempty" xml:"UserProfileSwitch,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) SetFileSystemId(v string) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile {
	s.FileSystemId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) SetUserProfileSwitch(v bool) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile {
	s.UserProfileSwitch = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow struct {
	FileSystemId        *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProfileFollowSwitch *bool   `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) SetFileSystemId(v string) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow {
	s.FileSystemId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) SetProfileFollowSwitch(v bool) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow {
	s.ProfileFollowSwitch = &v
	return s
}

type ModifyAppInstanceGroupAttributeShrinkRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The network settings.
	//
	// >  If you want to use this parameter, submit a ticket.
	NetworkShrink *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The information about the resource group.
	NodePoolShrink *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	// Specifies whether only one application can be opened in a session.
	//
	// 	- After you enable this feature, the system assigns a session to each application if you open multiple applications in a delivery group. This consumes a larger number of sessions.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	PerSessionPerApp *bool `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// The application ID of the pre-open application. If you set `PreOpenMode` to `SINGLE_APP`, you cannot leave this parameter empty.``
	//
	// example:
	//
	// ca-b2ronxxd****
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// The pre-open mode.
	//
	// Valid values:
	//
	// 	- SINGLE_APP: enables the pre-open mode for a single application.
	//
	// 	- OFF: disables the pre-open mode. This is the default value.
	//
	// example:
	//
	// OFF
	PreOpenMode *string `json:"PreOpenMode,omitempty" xml:"PreOpenMode,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The security policy.
	SecurityPolicyShrink *string `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty"`
	// The duration for which sessions are retained after disconnection. Unit: minutes. After an end user disconnects from a session, the session is closed only after the specified duration elapses. If you want to permanently retain sessions, set this parameter to `-1`. Valid values:-1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The storage policy.
	StoragePolicyShrink *string `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetNetworkShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetNodePoolShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPerSessionPerApp(v bool) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PerSessionPerApp = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.PreOpenMode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetSecurityPolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.SecurityPolicyShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetStoragePolicyShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.StoragePolicyShrink = &v
	return s
}

type ModifyAppInstanceGroupAttributeResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetCode(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetMessage(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetRequestId(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppInstanceGroupAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppInstanceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifyAppInstanceGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetStatusCode(v int32) *ModifyAppInstanceGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetBody(v *ModifyAppInstanceGroupAttributeResponseBody) *ModifyAppInstanceGroupAttributeResponse {
	s.Body = v
	return s
}

type ModifyAppPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-ee2znjktwgxu2****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string                            `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	VideoPolicy *ModifyAppPolicyRequestVideoPolicy `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty" type:"Struct"`
}

func (s ModifyAppPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyRequest) SetAppPolicyId(v string) *ModifyAppPolicyRequest {
	s.AppPolicyId = &v
	return s
}

func (s *ModifyAppPolicyRequest) SetProductType(v string) *ModifyAppPolicyRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppPolicyRequest) SetVideoPolicy(v *ModifyAppPolicyRequestVideoPolicy) *ModifyAppPolicyRequest {
	s.VideoPolicy = v
	return s
}

type ModifyAppPolicyRequestVideoPolicy struct {
	// example:
	//
	// 60
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// example:
	//
	// 1080
	SessionResolutionHeight *int32 `json:"SessionResolutionHeight,omitempty" xml:"SessionResolutionHeight,omitempty"`
	// example:
	//
	// 1920
	SessionResolutionWidth *int32 `json:"SessionResolutionWidth,omitempty" xml:"SessionResolutionWidth,omitempty"`
	// example:
	//
	// video
	StreamingMode *string `json:"StreamingMode,omitempty" xml:"StreamingMode,omitempty"`
	// example:
	//
	// false
	TerminalResolutionAdaptive *bool   `json:"TerminalResolutionAdaptive,omitempty" xml:"TerminalResolutionAdaptive,omitempty"`
	VisualQualityStrategy      *string `json:"VisualQualityStrategy,omitempty" xml:"VisualQualityStrategy,omitempty"`
	// example:
	//
	// true
	Webrtc *bool `json:"Webrtc,omitempty" xml:"Webrtc,omitempty"`
}

func (s ModifyAppPolicyRequestVideoPolicy) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppPolicyRequestVideoPolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetFrameRate(v int32) *ModifyAppPolicyRequestVideoPolicy {
	s.FrameRate = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetSessionResolutionHeight(v int32) *ModifyAppPolicyRequestVideoPolicy {
	s.SessionResolutionHeight = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetSessionResolutionWidth(v int32) *ModifyAppPolicyRequestVideoPolicy {
	s.SessionResolutionWidth = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetStreamingMode(v string) *ModifyAppPolicyRequestVideoPolicy {
	s.StreamingMode = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetTerminalResolutionAdaptive(v bool) *ModifyAppPolicyRequestVideoPolicy {
	s.TerminalResolutionAdaptive = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetVisualQualityStrategy(v string) *ModifyAppPolicyRequestVideoPolicy {
	s.VisualQualityStrategy = &v
	return s
}

func (s *ModifyAppPolicyRequestVideoPolicy) SetWebrtc(v bool) *ModifyAppPolicyRequestVideoPolicy {
	s.Webrtc = &v
	return s
}

type ModifyAppPolicyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-ee2znjktwgxu2****
	AppPolicyId *string `json:"AppPolicyId,omitempty" xml:"AppPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType       *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	VideoPolicyShrink *string `json:"VideoPolicy,omitempty" xml:"VideoPolicy,omitempty"`
}

func (s ModifyAppPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyShrinkRequest) SetAppPolicyId(v string) *ModifyAppPolicyShrinkRequest {
	s.AppPolicyId = &v
	return s
}

func (s *ModifyAppPolicyShrinkRequest) SetProductType(v string) *ModifyAppPolicyShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppPolicyShrinkRequest) SetVideoPolicyShrink(v string) *ModifyAppPolicyShrinkRequest {
	s.VideoPolicyShrink = &v
	return s
}

type ModifyAppPolicyResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyResponseBody) SetRequestId(v string) *ModifyAppPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppPolicyResponse) SetHeaders(v map[string]*string) *ModifyAppPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppPolicyResponse) SetStatusCode(v int32) *ModifyAppPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppPolicyResponse) SetBody(v *ModifyAppPolicyResponseBody) *ModifyAppPolicyResponse {
	s.Body = v
	return s
}

type ModifyNodePoolAmountRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The parameters related to the configuration change of the node pool.
	//
	// This parameter is required.
	NodePool *ModifyNodePoolAmountRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAmountRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountRequest) SetAppInstanceGroupId(v string) *ModifyNodePoolAmountRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyNodePoolAmountRequest) SetNodePool(v *ModifyNodePoolAmountRequestNodePool) *ModifyNodePoolAmountRequest {
	s.NodePool = v
	return s
}

func (s *ModifyNodePoolAmountRequest) SetProductType(v string) *ModifyNodePoolAmountRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAmountRequestNodePool struct {
	// The total number of subscription nodes after the change.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The change mode of subscription nodes.
	//
	// Valid value:
	//
	// 	- EXPAND_FROM_POST_PAID_EXPLICIT: changes from specified pay-as-you-go nodes
	//
	// example:
	//
	// EXPAND_FROM_POST_PAID_EXPLICIT
	PrePaidNodeAmountModifyMode *string `json:"PrePaidNodeAmountModifyMode,omitempty" xml:"PrePaidNodeAmountModifyMode,omitempty"`
	// The nodes for which you want to change the billing method.
	PrePaidNodeAmountModifyNodeIds []*string `json:"PrePaidNodeAmountModifyNodeIds,omitempty" xml:"PrePaidNodeAmountModifyNodeIds,omitempty" type:"Repeated"`
}

func (s ModifyNodePoolAmountRequestNodePool) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAmountRequestNodePool) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountRequestNodePool) SetNodeAmount(v int32) *ModifyNodePoolAmountRequestNodePool {
	s.NodeAmount = &v
	return s
}

func (s *ModifyNodePoolAmountRequestNodePool) SetPrePaidNodeAmountModifyMode(v string) *ModifyNodePoolAmountRequestNodePool {
	s.PrePaidNodeAmountModifyMode = &v
	return s
}

func (s *ModifyNodePoolAmountRequestNodePool) SetPrePaidNodeAmountModifyNodeIds(v []*string) *ModifyNodePoolAmountRequestNodePool {
	s.PrePaidNodeAmountModifyNodeIds = v
	return s
}

type ModifyNodePoolAmountShrinkRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The parameters related to the configuration change of the node pool.
	//
	// This parameter is required.
	NodePoolShrink *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAmountShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAmountShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountShrinkRequest) SetAppInstanceGroupId(v string) *ModifyNodePoolAmountShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyNodePoolAmountShrinkRequest) SetNodePoolShrink(v string) *ModifyNodePoolAmountShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *ModifyNodePoolAmountShrinkRequest) SetProductType(v string) *ModifyNodePoolAmountShrinkRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAmountResponseBody struct {
	// The returned data.
	Data *ModifyNodePoolAmountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodePoolAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountResponseBody) SetData(v *ModifyNodePoolAmountResponseBodyData) *ModifyNodePoolAmountResponseBody {
	s.Data = v
	return s
}

func (s *ModifyNodePoolAmountResponseBody) SetRequestId(v string) *ModifyNodePoolAmountResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNodePoolAmountResponseBodyData struct {
	// The order ID.
	//
	// example:
	//
	// 23429322113****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyNodePoolAmountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAmountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountResponseBodyData) SetOrderId(v string) *ModifyNodePoolAmountResponseBodyData {
	s.OrderId = &v
	return s
}

type ModifyNodePoolAmountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodePoolAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodePoolAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAmountResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAmountResponse) SetHeaders(v map[string]*string) *ModifyNodePoolAmountResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodePoolAmountResponse) SetStatusCode(v int32) *ModifyNodePoolAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodePoolAmountResponse) SetBody(v *ModifyNodePoolAmountResponseBody) *ModifyNodePoolAmountResponse {
	s.Body = v
	return s
}

type ModifyNodePoolAttributeRequest struct {
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 2
	NodeCapacity     *int32                                          `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolStrategy *ModifyNodePoolAttributeRequestNodePoolStrategy `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty" type:"Struct"`
	// example:
	//
	// rg-ew7va2g1wl3vm****
	PoolId *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// 产品类型。
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodePoolStrategy(v *ModifyNodePoolAttributeRequestNodePoolStrategy) *ModifyNodePoolAttributeRequest {
	s.NodePoolStrategy = v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetPoolId(v string) *ModifyNodePoolAttributeRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetProductType(v string) *ModifyNodePoolAttributeRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAttributeRequestNodePoolStrategy struct {
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// example:
	//
	// 10
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// 购买资源的数量。取值范围：1~100。
	//
	// >
	//
	// - 若为包年包月资源，则该参数不可修改。
	//
	// - 若为按量付费资源，则当弹性模式（`StrategyType`）为固定数量（`NODE_FIXED`）或自动扩缩容（`NODE_SCALING_BY_USAGE`）时该参数可修改。
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// 策略执行周期列表。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	RecurrenceSchedules []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// 策略失效日期。格式为：yyyy-MM-dd。失效日期与生效日期的间隔必须介于7天到1年之间（含7天和1年）。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	//
	// example:
	//
	// 2023-01-19
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// 策略生效日期。格式为：yyyy-MM-dd。该日期必须大于或等于当前日期。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	//
	// example:
	//
	// 2023-01-05
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	StrategyType       *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// 是否开启资源预热策略。`StrategyType`（弹性模式）设为`NODE_SCALING_BY_SCHEDULE`（定时扩缩容）时，该字段必填。
	//
	// example:
	//
	// false
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetMaxIdleAppInstanceAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetMaxScalingAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.MaxScalingAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetNodeAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.NodeAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetRecurrenceSchedules(v []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.RecurrenceSchedules = v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingDownAfterIdleMinutes(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingStep(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingStep = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingUsageThreshold(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyDisableDate(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyDisableDate = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyEnableDate(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyEnableDate = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyType(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyType = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetWarmUp(v bool) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.WarmUp = &v
	return s
}

type ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules struct {
	// 策略执行周期的类型。必须同时指定`RecurrenceType`和`RecurrenceValues`。
	//
	// example:
	//
	// weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// 策略执行周期的数值列表。
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// 策略执行周期的时间段列表。时间段设置要求：
	//
	// - 最多可添加3个时间段。
	//
	// - 时间段之间不重叠。
	//
	// - 时间段之间的间隔大于或等于5分钟。
	//
	// - 单个时间段的时长大于或等于15分钟。
	//
	// - 所有时间段累计不跨天。
	TimerPeriods []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetRecurrenceType(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetRecurrenceValues(v []*int32) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetTimerPeriods(v []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods struct {
	// 资源数量。
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// 结束时间。格式为HH:mm。
	//
	// example:
	//
	// 15:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 开始时间。格式为HH:mm。
	//
	// example:
	//
	// 12:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type ModifyNodePoolAttributeShrinkRequest struct {
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 2
	NodeCapacity           *int32  `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolStrategyShrink *string `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty"`
	// example:
	//
	// rg-ew7va2g1wl3vm****
	PoolId *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// 产品类型。
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeShrinkRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodePoolStrategyShrink(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.NodePoolStrategyShrink = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetPoolId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetProductType(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAttributeResponseBody struct {
	// example:
	//
	// InvalidParameter.PoolId
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The parameter PoolId is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodePoolAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponseBody) SetCode(v string) *ModifyNodePoolAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetMessage(v string) *ModifyNodePoolAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetRequestId(v string) *ModifyNodePoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNodePoolAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNodePoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNodePoolAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponse) SetHeaders(v map[string]*string) *ModifyNodePoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetStatusCode(v int32) *ModifyNodePoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetBody(v *ModifyNodePoolAttributeResponseBody) *ModifyNodePoolAttributeResponse {
	s.Body = v
	return s
}

type ModifyTenantConfigRequest struct {
	// Specifies whether to enable the resource expiration reminder feature.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AppInstanceGroupExpireRemind *bool `json:"AppInstanceGroupExpireRemind,omitempty" xml:"AppInstanceGroupExpireRemind,omitempty"`
}

func (s ModifyTenantConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigRequest) SetAppInstanceGroupExpireRemind(v bool) *ModifyTenantConfigRequest {
	s.AppInstanceGroupExpireRemind = &v
	return s
}

type ModifyTenantConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigResponseBody) SetRequestId(v string) *ModifyTenantConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTenantConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTenantConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantConfigResponse) SetHeaders(v map[string]*string) *ModifyTenantConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantConfigResponse) SetStatusCode(v int32) *ModifyTenantConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantConfigResponse) SetBody(v *ModifyTenantConfigResponseBody) *ModifyTenantConfigResponse {
	s.Body = v
	return s
}

type PageListAppInstanceGroupUserRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The number of the page to return. We recommend that you configure this parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to be return on each page. The value cannot be greater than `100`. We recommend that you configure this parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s PageListAppInstanceGroupUserRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserRequest) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserRequest) SetAppInstanceGroupId(v string) *PageListAppInstanceGroupUserRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageNumber(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageNumber = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageSize(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageSize = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetProductType(v string) *PageListAppInstanceGroupUserRequest {
	s.ProductType = &v
	return s
}

type PageListAppInstanceGroupUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The users.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s PageListAppInstanceGroupUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserResponseBody) SetRequestId(v string) *PageListAppInstanceGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageListAppInstanceGroupUserResponseBody) SetUsers(v []*string) *PageListAppInstanceGroupUserResponseBody {
	s.Users = v
	return s
}

type PageListAppInstanceGroupUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageListAppInstanceGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageListAppInstanceGroupUserResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserResponse) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserResponse) SetHeaders(v map[string]*string) *PageListAppInstanceGroupUserResponse {
	s.Headers = v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetStatusCode(v int32) *PageListAppInstanceGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetBody(v *PageListAppInstanceGroupUserResponseBody) *PageListAppInstanceGroupUserResponse {
	s.Body = v
	return s
}

type RenewAppInstanceGroupRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: manual payment. This is the default value.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The subscription duration of resources. This parameter must be configured together with `PeriodUnit`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. This parameter must be configured together with `Period`. The following items describe valid values for the combinations of `Period` and `PeriodUnit`:
	//
	// 	- 1 Week
	//
	// 	- 1 Month
	//
	// 	- 2 Month
	//
	// 	- 3 Month
	//
	// 	- 6 Month
	//
	// 	- 1 Year
	//
	// 	- 2 Year
	//
	// 	- 3 Year
	//
	// >  The value of this parameter is case-insensitive. For example, `Week` is valid and `week` is invalid. If you specify a value combination other than the preceding combinations, such as `2 Week`, the operation can still be called. However, an error occurs when you place the order.
	//
	// This parameter is required.
	//
	// example:
	//
	// Week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The promotion ID. You can call the [GetResourcePrice](https://help.aliyun.com/document_detail/428503.html) operation to obtain the ID.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s RenewAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *RenewAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetAutoPay(v bool) *RenewAppInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPeriod(v int32) *RenewAppInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPeriodUnit(v string) *RenewAppInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetProductType(v string) *RenewAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *RenewAppInstanceGroupRequest) SetPromotionId(v string) *RenewAppInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

type RenewAppInstanceGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 123456****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupResponseBody) SetCode(v string) *RenewAppInstanceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetMessage(v string) *RenewAppInstanceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetOrderId(v string) *RenewAppInstanceGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewAppInstanceGroupResponseBody) SetRequestId(v string) *RenewAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type RenewAppInstanceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceGroupResponse) SetHeaders(v map[string]*string) *RenewAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewAppInstanceGroupResponse) SetStatusCode(v int32) *RenewAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppInstanceGroupResponse) SetBody(v *RenewAppInstanceGroupResponseBody) *RenewAppInstanceGroupResponse {
	s.Body = v
	return s
}

type TagCloudResourcesRequest struct {
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	Tags []*TagCloudResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagCloudResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesRequest) SetResourceIds(v []*string) *TagCloudResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagCloudResourcesRequest) SetResourceType(v string) *TagCloudResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagCloudResourcesRequest) SetTags(v []*TagCloudResourcesRequestTags) *TagCloudResourcesRequest {
	s.Tags = v
	return s
}

type TagCloudResourcesRequestTags struct {
	// This parameter is required.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 720p
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagCloudResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagCloudResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesRequestTags) SetKey(v string) *TagCloudResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagCloudResourcesRequestTags) SetValue(v string) *TagCloudResourcesRequestTags {
	s.Value = &v
	return s
}

type TagCloudResourcesResponseBody struct {
	FailedResources []*TagCloudResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagCloudResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponseBody) SetFailedResources(v []*TagCloudResourcesResponseBodyFailedResources) *TagCloudResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *TagCloudResourcesResponseBody) SetRequestId(v string) *TagCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagCloudResourcesResponseBodyFailedResources struct {
	// example:
	//
	// TAG_KEY_DUPLICATED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Duplicate tag keys exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// aig-001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string                                             `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags         []*TagCloudResourcesResponseBodyFailedResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagCloudResourcesResponseBodyFailedResources) String() string {
	return tea.Prettify(s)
}

func (s TagCloudResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetCode(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.Code = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetMessage(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.Message = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetResourceId(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.ResourceId = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetResourceType(v string) *TagCloudResourcesResponseBodyFailedResources {
	s.ResourceType = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResources) SetTags(v []*TagCloudResourcesResponseBodyFailedResourcesTags) *TagCloudResourcesResponseBodyFailedResources {
	s.Tags = v
	return s
}

type TagCloudResourcesResponseBodyFailedResourcesTags struct {
	// example:
	//
	// System/Scheduler/STOP_NEW_USER_CONNECTION
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// System
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagCloudResourcesResponseBodyFailedResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s TagCloudResourcesResponseBodyFailedResourcesTags) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) SetKey(v string) *TagCloudResourcesResponseBodyFailedResourcesTags {
	s.Key = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) SetScope(v string) *TagCloudResourcesResponseBodyFailedResourcesTags {
	s.Scope = &v
	return s
}

func (s *TagCloudResourcesResponseBodyFailedResourcesTags) SetValue(v string) *TagCloudResourcesResponseBodyFailedResourcesTags {
	s.Value = &v
	return s
}

type TagCloudResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagCloudResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagCloudResourcesResponse) SetHeaders(v map[string]*string) *TagCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagCloudResourcesResponse) SetStatusCode(v int32) *TagCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagCloudResourcesResponse) SetBody(v *TagCloudResourcesResponseBody) *TagCloudResourcesResponse {
	s.Body = v
	return s
}

type UnbindRequest struct {
	// The ID of the delivery group. You can call the [GetConnectionTicket](~~GetConnectionTicket~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The session ID. You can call the [GetConnectionTicket](~~GetConnectionTicket~~) operation to obtain the ID.
	//
	// example:
	//
	// ai-d297eyf83g5ni****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The ID of the persistent session. You can call the [GetConnectionTicket](~~GetConnectionTicket~~) operation to obtain the ID.
	//
	// example:
	//
	// p-0bxls9m3cl7s****
	AppInstancePersistentId *string `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s UnbindRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindRequest) GoString() string {
	return s.String()
}

func (s *UnbindRequest) SetAppInstanceGroupId(v string) *UnbindRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceId(v string) *UnbindRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindRequest) SetAppInstancePersistentId(v string) *UnbindRequest {
	s.AppInstancePersistentId = &v
	return s
}

func (s *UnbindRequest) SetEndUserId(v string) *UnbindRequest {
	s.EndUserId = &v
	return s
}

func (s *UnbindRequest) SetProductType(v string) *UnbindRequest {
	s.ProductType = &v
	return s
}

type UnbindResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResponseBody) SetRequestId(v string) *UnbindResponseBody {
	s.RequestId = &v
	return s
}

type UnbindResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponse) GoString() string {
	return s.String()
}

func (s *UnbindResponse) SetHeaders(v map[string]*string) *UnbindResponse {
	s.Headers = v
	return s
}

func (s *UnbindResponse) SetStatusCode(v int32) *UnbindResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindResponse) SetBody(v *UnbindResponseBody) *UnbindResponse {
	s.Body = v
	return s
}

type UntagCloudResourcesRequest struct {
	// The resource IDs. You can specify up to 50 resource IDs. You do not need to specify this parameter if you set ResourceType to AliUid.
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource from which you want to remove tags.
	//
	// Valid values:
	//
	// 	- AppId: app ID.
	//
	// 	- WyId: Alibaba Cloud Workspace user ID.
	//
	// 	- AppInstanceGroupId: delivery group ID.
	//
	// 	- AliUid: tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to remove from the cloud resources. System and custom tags are supported. You can specify up to 10 tags.
	//
	// Valid values for system tags:
	//
	// 	- `System/Scheduler/GRAYSCALE`: canary tags.
	//
	// 	- `System/Scheduler/STOP_NEW_USER_CONNECTION`: tags used to stop new users bound to the delivery group from establishing a connection.
	//
	// This parameter is required.
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s UntagCloudResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesRequest) SetResourceIds(v []*string) *UntagCloudResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagCloudResourcesRequest) SetResourceType(v string) *UntagCloudResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagCloudResourcesRequest) SetTagKeys(v []*string) *UntagCloudResourcesRequest {
	s.TagKeys = v
	return s
}

type UntagCloudResourcesResponseBody struct {
	// The cloud resources whose tags failed to be removed and the corresponding tags.
	FailedResources []*UntagCloudResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E25FC620-6B6F-12D2-A992-AD8727DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagCloudResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponseBody) SetFailedResources(v []*UntagCloudResourcesResponseBodyFailedResources) *UntagCloudResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *UntagCloudResourcesResponseBody) SetRequestId(v string) *UntagCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagCloudResourcesResponseBodyFailedResources struct {
	// The error code.
	//
	// example:
	//
	// UNTAG_RESOURCE_FAILED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to untag resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The resource IDs.
	//
	// example:
	//
	// aig-00000001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the cloud resource.
	//
	// Valid values:
	//
	// 	- AppId: app ID.
	//
	// 	- WyId: Alibaba Cloud Workspace user ID.
	//
	// 	- AppInstanceGroupId: delivery group ID.
	//
	// 	- AliUid: tenant ID.
	//
	// example:
	//
	// AppInstanceGroupId
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that failed to be removed from the cloud resources.
	Tags []*UntagCloudResourcesResponseBodyFailedResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UntagCloudResourcesResponseBodyFailedResources) String() string {
	return tea.Prettify(s)
}

func (s UntagCloudResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetCode(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.Code = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetMessage(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.Message = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetResourceId(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.ResourceId = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetResourceType(v string) *UntagCloudResourcesResponseBodyFailedResources {
	s.ResourceType = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResources) SetTags(v []*UntagCloudResourcesResponseBodyFailedResourcesTags) *UntagCloudResourcesResponseBodyFailedResources {
	s.Tags = v
	return s
}

type UntagCloudResourcesResponseBodyFailedResourcesTags struct {
	// The tag key.
	//
	// example:
	//
	// Resolution
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag type.
	//
	// Valid values:
	//
	// 	- Custom: custom tag.
	//
	// 	- System: system tag.
	//
	// example:
	//
	// Custom
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s UntagCloudResourcesResponseBodyFailedResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s UntagCloudResourcesResponseBodyFailedResourcesTags) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponseBodyFailedResourcesTags) SetKey(v string) *UntagCloudResourcesResponseBodyFailedResourcesTags {
	s.Key = &v
	return s
}

func (s *UntagCloudResourcesResponseBodyFailedResourcesTags) SetScope(v string) *UntagCloudResourcesResponseBodyFailedResourcesTags {
	s.Scope = &v
	return s
}

type UntagCloudResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagCloudResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagCloudResourcesResponse) SetHeaders(v map[string]*string) *UntagCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagCloudResourcesResponse) SetStatusCode(v int32) *UntagCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagCloudResourcesResponse) SetBody(v *UntagCloudResourcesResponseBody) *UntagCloudResourcesResponse {
	s.Body = v
	return s
}

type UpdateAppInstanceGroupImageRequest struct {
	// The image ID of the application. You can obtain the ID from the Images page in the App Streaming console.
	//
	// This parameter is required.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai).
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s UpdateAppInstanceGroupImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppCenterImageId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppInstanceGroupId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetBizRegionId(v string) *UpdateAppInstanceGroupImageRequest {
	s.BizRegionId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetProductType(v string) *UpdateAppInstanceGroupImageRequest {
	s.ProductType = &v
	return s
}

type UpdateAppInstanceGroupImageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppInstanceGroupImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetCode(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetMessage(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetRequestId(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppInstanceGroupImageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppInstanceGroupImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppInstanceGroupImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponse) SetHeaders(v map[string]*string) *UpdateAppInstanceGroupImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetStatusCode(v int32) *UpdateAppInstanceGroupImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetBody(v *UpdateAppInstanceGroupImageResponseBody) *UpdateAppInstanceGroupImageResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 更新访问管理
//
// @param request - AccessPageSetAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccessPageSetAclResponse
func (client *Client) AccessPageSetAclWithOptions(request *AccessPageSetAclRequest, runtime *util.RuntimeOptions) (_result *AccessPageSetAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessMode)) {
		query["AccessMode"] = request.AccessMode
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageName)) {
		query["AccessPageName"] = request.AccessPageName
	}

	if !tea.BoolValue(util.IsUnset(request.EffectTime)) {
		query["EffectTime"] = request.EffectTime
	}

	if !tea.BoolValue(util.IsUnset(request.Unit)) {
		query["Unit"] = request.Unit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AccessPageSetAcl"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AccessPageSetAclResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AccessPageSetAclResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新访问管理
//
// @param request - AccessPageSetAclRequest
//
// @return AccessPageSetAclResponse
func (client *Client) AccessPageSetAcl(request *AccessPageSetAclRequest) (_result *AccessPageSetAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AccessPageSetAclResponse{}
	_body, _err := client.AccessPageSetAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the execution time of an over-the-air (OTA) update task.
//
// @param request - ApproveOtaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveOtaTaskResponse
func (client *Client) ApproveOtaTaskWithOptions(request *ApproveOtaTaskRequest, runtime *util.RuntimeOptions) (_result *ApproveOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OtaType)) {
		body["OtaType"] = request.OtaType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ApproveOtaTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ApproveOtaTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Sets the execution time of an over-the-air (OTA) update task.
//
// @param request - ApproveOtaTaskRequest
//
// @return ApproveOtaTaskResponse
func (client *Client) ApproveOtaTask(request *ApproveOtaTaskRequest) (_result *ApproveOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveOtaTaskResponse{}
	_body, _err := client.ApproveOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话包收费查询
//
// @param request - AskSessionPackagePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AskSessionPackagePriceResponse
func (client *Client) AskSessionPackagePriceWithOptions(request *AskSessionPackagePriceRequest, runtime *util.RuntimeOptions) (_result *AskSessionPackagePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSessions)) {
		query["MaxSessions"] = request.MaxSessions
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageType)) {
		query["SessionPackageType"] = request.SessionPackageType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionSpec)) {
		query["SessionSpec"] = request.SessionSpec
	}

	if !tea.BoolValue(util.IsUnset(request.SessionType)) {
		query["SessionType"] = request.SessionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AskSessionPackagePrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AskSessionPackagePriceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AskSessionPackagePriceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 会话包收费查询
//
// @param request - AskSessionPackagePriceRequest
//
// @return AskSessionPackagePriceResponse
func (client *Client) AskSessionPackagePrice(request *AskSessionPackagePriceRequest) (_result *AskSessionPackagePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AskSessionPackagePriceResponse{}
	_body, _err := client.AskSessionPackagePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权用户
//
// @param tmpReq - AuthorizeInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeInstanceGroupResponse
func (client *Client) AuthorizeInstanceGroupWithOptions(tmpReq *AuthorizeInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *AuthorizeInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AuthorizeInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserMeta)) {
		request.UserMetaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserMeta, tea.String("UserMeta"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstancePersistentId)) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizeUserIds)) {
		bodyFlat["AuthorizeUserIds"] = request.AuthorizeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.UnAuthorizeUserIds)) {
		bodyFlat["UnAuthorizeUserIds"] = request.UnAuthorizeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserMetaShrink)) {
		body["UserMeta"] = request.UserMetaShrink
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AuthorizeInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AuthorizeInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 授权用户
//
// @param request - AuthorizeInstanceGroupRequest
//
// @return AuthorizeInstanceGroupResponse
func (client *Client) AuthorizeInstanceGroup(request *AuthorizeInstanceGroupRequest) (_result *AuthorizeInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeInstanceGroupResponse{}
	_body, _err := client.AuthorizeInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置会话包
//
// @param request - BuySessionPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuySessionPackageResponse
func (client *Client) BuySessionPackageWithOptions(request *BuySessionPackageRequest, runtime *util.RuntimeOptions) (_result *BuySessionPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSessions)) {
		query["MaxSessions"] = request.MaxSessions
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageName)) {
		query["SessionPackageName"] = request.SessionPackageName
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageType)) {
		query["SessionPackageType"] = request.SessionPackageType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionSpec)) {
		query["SessionSpec"] = request.SessionSpec
	}

	if !tea.BoolValue(util.IsUnset(request.SessionType)) {
		query["SessionType"] = request.SessionType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BuySessionPackage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BuySessionPackageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BuySessionPackageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 配置会话包
//
// @param request - BuySessionPackageRequest
//
// @return BuySessionPackageResponse
func (client *Client) BuySessionPackage(request *BuySessionPackageRequest) (_result *BuySessionPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BuySessionPackageResponse{}
	_body, _err := client.BuySessionPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建访问页面
//
// @param request - CreateAccessPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessPageResponse
func (client *Client) CreateAccessPageWithOptions(request *CreateAccessPageRequest, runtime *util.RuntimeOptions) (_result *CreateAccessPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageName)) {
		query["AccessPageName"] = request.AccessPageName
	}

	if !tea.BoolValue(util.IsUnset(request.CloudEnvId)) {
		query["CloudEnvId"] = request.CloudEnvId
	}

	if !tea.BoolValue(util.IsUnset(request.EffectTime)) {
		query["EffectTime"] = request.EffectTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Unit)) {
		query["Unit"] = request.Unit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessPage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAccessPageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAccessPageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建访问页面
//
// @param request - CreateAccessPageRequest
//
// @return CreateAccessPageResponse
func (client *Client) CreateAccessPage(request *CreateAccessPageRequest) (_result *CreateAccessPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessPageResponse{}
	_body, _err := client.CreateAccessPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建云应用交付组
//
// @param tmpReq - CreateAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppInstanceGroupResponse
func (client *Client) CreateAppInstanceGroupWithOptions(tmpReq *CreateAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Network)) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, tea.String("Network"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodePool)) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, tea.String("NodePool"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RuntimePolicy)) {
		request.RuntimePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimePolicy, tea.String("RuntimePolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SecurityPolicy)) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, tea.String("SecurityPolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StoragePolicy)) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, tea.String("StoragePolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserDefinePolicy)) {
		request.UserDefinePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserDefinePolicy, tea.String("UserDefinePolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserInfo)) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInfo, tea.String("UserInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VideoPolicy)) {
		request.VideoPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoPolicy, tea.String("VideoPolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserDefinePolicyShrink)) {
		query["UserDefinePolicy"] = request.UserDefinePolicyShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		body["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		body["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AppPackageType)) {
		body["AppPackageType"] = request.AppPackageType
	}

	if !tea.BoolValue(util.IsUnset(request.AppPolicyId)) {
		body["AppPolicyId"] = request.AppPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthMode)) {
		body["AuthMode"] = request.AuthMode
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		body["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeResourceMode)) {
		body["ChargeResourceMode"] = request.ChargeResourceMode
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkShrink)) {
		body["Network"] = request.NetworkShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolShrink)) {
		body["NodePool"] = request.NodePoolShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PreOpenAppId)) {
		body["PreOpenAppId"] = request.PreOpenAppId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		body["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuntimePolicyShrink)) {
		body["RuntimePolicy"] = request.RuntimePolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyShrink)) {
		body["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SessionTimeout)) {
		body["SessionTimeout"] = request.SessionTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.StoragePolicyShrink)) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SubPayType)) {
		body["SubPayType"] = request.SubPayType
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	if !tea.BoolValue(util.IsUnset(request.VideoPolicyShrink)) {
		body["VideoPolicy"] = request.VideoPolicyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAppInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAppInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建云应用交付组
//
// @param request - CreateAppInstanceGroupRequest
//
// @return CreateAppInstanceGroupResponse
func (client *Client) CreateAppInstanceGroup(request *CreateAppInstanceGroupRequest) (_result *CreateAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppInstanceGroupResponse{}
	_body, _err := client.CreateAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new image by debugging the delivery group.
//
// @param request - CreateImageFromAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageFromAppInstanceGroupResponse
func (client *Client) CreateImageFromAppInstanceGroupWithOptions(request *CreateImageFromAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateImageFromAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageName)) {
		body["AppCenterImageName"] = request.AppCenterImageName
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageFromAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateImageFromAppInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateImageFromAppInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a new image by debugging the delivery group.
//
// @param request - CreateImageFromAppInstanceGroupRequest
//
// @return CreateImageFromAppInstanceGroupResponse
func (client *Client) CreateImageFromAppInstanceGroup(request *CreateImageFromAppInstanceGroupRequest) (_result *CreateImageFromAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageFromAppInstanceGroupResponse{}
	_body, _err := client.CreateImageFromAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除访问页面
//
// @param request - DeleteAccessPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessPageResponse
func (client *Client) DeleteAccessPageWithOptions(request *DeleteAccessPageRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessPage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAccessPageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAccessPageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除访问页面
//
// @param request - DeleteAccessPageRequest
//
// @return DeleteAccessPageResponse
func (client *Client) DeleteAccessPage(request *DeleteAccessPageRequest) (_result *DeleteAccessPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessPageResponse{}
	_body, _err := client.DeleteAccessPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a delivery group that uses the By Resource - Pay-as-you-go billing method.
//
// Description:
//
// >  You cannot call this operation to delete a subscription delivery group.
//
// @param request - DeleteAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInstanceGroupResponse
func (client *Client) DeleteAppInstanceGroupWithOptions(request *DeleteAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAppInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAppInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a delivery group that uses the By Resource - Pay-as-you-go billing method.
//
// Description:
//
// >  You cannot call this operation to delete a subscription delivery group.
//
// @param request - DeleteAppInstanceGroupRequest
//
// @return DeleteAppInstanceGroupResponse
func (client *Client) DeleteAppInstanceGroup(request *DeleteAppInstanceGroupRequest) (_result *DeleteAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppInstanceGroupResponse{}
	_body, _err := client.DeleteAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application instance.
//
// Description:
//
// Only application instances that are in the Initializing or Idle state can be deleted. The operation can be called only by specific customers.
//
// @param request - DeleteAppInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppInstancesResponse
func (client *Client) DeleteAppInstancesWithOptions(request *DeleteAppInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteAppInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceIds)) {
		body["AppInstanceIds"] = request.AppInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppInstances"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAppInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAppInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an application instance.
//
// Description:
//
// Only application instances that are in the Initializing or Idle state can be deleted. The operation can be called only by specific customers.
//
// @param request - DeleteAppInstancesRequest
//
// @return DeleteAppInstancesResponse
func (client *Client) DeleteAppInstances(request *DeleteAppInstancesRequest) (_result *DeleteAppInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppInstancesResponse{}
	_body, _err := client.DeleteAppInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 访客进入访问页面的匿名api
//
// @param request - GetAccessPageSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessPageSessionResponse
func (client *Client) GetAccessPageSessionWithOptions(request *GetAccessPageSessionRequest, runtime *util.RuntimeOptions) (_result *GetAccessPageSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageToken)) {
		query["AccessPageToken"] = request.AccessPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUserId)) {
		query["ExternalUserId"] = request.ExternalUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessPageSession"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAccessPageSessionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAccessPageSessionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 访客进入访问页面的匿名api
//
// @param request - GetAccessPageSessionRequest
//
// @return GetAccessPageSessionResponse
func (client *Client) GetAccessPageSession(request *GetAccessPageSessionRequest) (_result *GetAccessPageSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessPageSessionResponse{}
	_body, _err := client.GetAccessPageSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取交付组详情
//
// @param request - GetAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppInstanceGroupResponse
func (client *Client) GetAppInstanceGroupWithOptions(request *GetAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *GetAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAppInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAppInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取交付组详情
//
// @param request - GetAppInstanceGroupRequest
//
// @return GetAppInstanceGroupResponse
func (client *Client) GetAppInstanceGroup(request *GetAppInstanceGroupRequest) (_result *GetAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppInstanceGroupResponse{}
	_body, _err := client.GetAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the credential that is used to connect to App Streaming.
//
// Description:
//
// You must call this operation at least twice to obtain a connection credential.
//
// The first time you call this operation, the system assigns an application instance to the specified convenience account and then starts the application. In this case, the ID of the started task, which is indicated by `TaskID`, is returned.
//
// In subsequent calls, you must configure `TaskID` to query whether the task is completed. If the value of `TaskStatus` in the response is `Finished`, the connection credential, which is indicated by `Ticket`, is returned.
//
// @param request - GetConnectionTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupIdList)) {
		body["AppInstanceGroupIdList"] = request.AppInstanceGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstancePersistentId)) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !tea.BoolValue(util.IsUnset(request.AppStartParam)) {
		body["AppStartParam"] = request.AppStartParam
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetConnectionTicketResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetConnectionTicketResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the credential that is used to connect to App Streaming.
//
// Description:
//
// You must call this operation at least twice to obtain a connection credential.
//
// The first time you call this operation, the system assigns an application instance to the specified convenience account and then starts the application. In this case, the ID of the started task, which is indicated by `TaskID`, is returned.
//
// In subsequent calls, you must configure `TaskID` to query whether the task is completed. If the value of `TaskStatus` in the response is `Finished`, the connection credential, which is indicated by `Ticket`, is returned.
//
// @param request - GetConnectionTicketRequest
//
// @return GetConnectionTicketResponse
func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information that is used to debug an application instance.
//
// @param request - GetDebugAppInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDebugAppInstanceResponse
func (client *Client) GetDebugAppInstanceWithOptions(request *GetDebugAppInstanceRequest, runtime *util.RuntimeOptions) (_result *GetDebugAppInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDebugAppInstance"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDebugAppInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDebugAppInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries information that is used to debug an application instance.
//
// @param request - GetDebugAppInstanceRequest
//
// @return GetDebugAppInstanceResponse
func (client *Client) GetDebugAppInstance(request *GetDebugAppInstanceRequest) (_result *GetDebugAppInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDebugAppInstanceResponse{}
	_body, _err := client.GetDebugAppInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an over-the-air (OTA) update task, including the available versions and version description.
//
// @param request - GetOtaTaskByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOtaTaskByTaskIdResponse
func (client *Client) GetOtaTaskByTaskIdWithOptions(request *GetOtaTaskByTaskIdRequest, runtime *util.RuntimeOptions) (_result *GetOtaTaskByTaskIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOtaTaskByTaskId"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetOtaTaskByTaskIdResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetOtaTaskByTaskIdResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of an over-the-air (OTA) update task, including the available versions and version description.
//
// @param request - GetOtaTaskByTaskIdRequest
//
// @return GetOtaTaskByTaskIdResponse
func (client *Client) GetOtaTaskByTaskId(request *GetOtaTaskByTaskIdRequest) (_result *GetOtaTaskByTaskIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOtaTaskByTaskIdResponse{}
	_body, _err := client.GetOtaTaskByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource prices.
//
// @param request - GetResourcePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcePriceResponse
func (client *Client) GetResourcePriceWithOptions(request *GetResourcePriceRequest, runtime *util.RuntimeOptions) (_result *GetResourcePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceType)) {
		query["AppInstanceType"] = request.AppInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceType)) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourcePrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourcePriceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourcePriceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries resource prices.
//
// @param request - GetResourcePriceRequest
//
// @return GetResourcePriceResponse
func (client *Client) GetResourcePrice(request *GetResourcePriceRequest) (_result *GetResourcePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourcePriceResponse{}
	_body, _err := client.GetResourcePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the renewal prices of App Streaming resources.
//
// @param request - GetResourceRenewPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceRenewPriceResponse
func (client *Client) GetResourceRenewPriceWithOptions(request *GetResourceRenewPriceRequest, runtime *util.RuntimeOptions) (_result *GetResourceRenewPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceRenewPrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourceRenewPriceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourceRenewPriceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the renewal prices of App Streaming resources.
//
// @param request - GetResourceRenewPriceRequest
//
// @return GetResourceRenewPriceResponse
func (client *Client) GetResourceRenewPrice(request *GetResourceRenewPriceRequest) (_result *GetResourceRenewPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceRenewPriceResponse{}
	_body, _err := client.GetResourceRenewPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 访问页面分页查询
//
// @param request - ListAccessPagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessPagesResponse
func (client *Client) ListAccessPagesWithOptions(request *ListAccessPagesRequest, runtime *util.RuntimeOptions) (_result *ListAccessPagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessPageId)) {
		query["AccessPageId"] = request.AccessPageId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessPageName)) {
		query["AccessPageName"] = request.AccessPageName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessPages"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAccessPagesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAccessPagesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 访问页面分页查询
//
// @param request - ListAccessPagesRequest
//
// @return ListAccessPagesResponse
func (client *Client) ListAccessPages(request *ListAccessPagesRequest) (_result *ListAccessPagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessPagesResponse{}
	_body, _err := client.ListAccessPagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of multiple delivery groups that meet the query conditions.
//
// @param request - ListAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstanceGroupResponse
func (client *Client) ListAppInstanceGroupWithOptions(request *ListAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *ListAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceType)) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAppInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAppInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of multiple delivery groups that meet the query conditions.
//
// @param request - ListAppInstanceGroupRequest
//
// @return ListAppInstanceGroupResponse
func (client *Client) ListAppInstanceGroup(request *ListAppInstanceGroupRequest) (_result *ListAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInstanceGroupResponse{}
	_body, _err := client.ListAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of application instances in a delivery group, including the IDs, status, creation time, update time, session status, and public IP addresses associated with the primary NICs of the instances.
//
// @param request - ListAppInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppInstancesResponse
func (client *Client) ListAppInstancesWithOptions(request *ListAppInstancesRequest, runtime *util.RuntimeOptions) (_result *ListAppInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		query["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDeleted)) {
		query["IncludeDeleted"] = request.IncludeDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		query["UserIdList"] = request.UserIdList
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceIdList)) {
		body["AppInstanceIdList"] = request.AppInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppInstances"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAppInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAppInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the details of application instances in a delivery group, including the IDs, status, creation time, update time, session status, and public IP addresses associated with the primary NICs of the instances.
//
// @param request - ListAppInstancesRequest
//
// @return ListAppInstancesResponse
func (client *Client) ListAppInstances(request *ListAppInstancesRequest) (_result *ListAppInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInstancesResponse{}
	_body, _err := client.ListAppInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询绑定信息，支持分页
//
// @param request - ListBindInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBindInfoResponse
func (client *Client) ListBindInfoWithOptions(request *ListBindInfoRequest, runtime *util.RuntimeOptions) (_result *ListBindInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIdList)) {
		body["AppIdList"] = request.AppIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupIdList)) {
		body["AppInstanceGroupIdList"] = request.AppInstanceGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceIdList)) {
		body["AppInstanceIdList"] = request.AppInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["UserIdList"] = request.UserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.WyIdList)) {
		body["WyIdList"] = request.WyIdList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBindInfo"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListBindInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListBindInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询绑定信息，支持分页
//
// @param request - ListBindInfoRequest
//
// @return ListBindInfoResponse
func (client *Client) ListBindInfo(request *ListBindInfoRequest) (_result *ListBindInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBindInfoResponse{}
	_body, _err := client.ListBindInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource types that are available for purchase when you create a delivery group.
//
// @param request - ListNodeInstanceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeInstanceTypeResponse
func (client *Client) ListNodeInstanceTypeWithOptions(request *ListNodeInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *ListNodeInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Gpu)) {
		query["Gpu"] = request.Gpu
	}

	if !tea.BoolValue(util.IsUnset(request.GpuMemory)) {
		query["GpuMemory"] = request.GpuMemory
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceType)) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceTypeFamily)) {
		query["NodeInstanceTypeFamily"] = request.NodeInstanceTypeFamily
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
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

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeInstanceType"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListNodeInstanceTypeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListNodeInstanceTypeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the resource types that are available for purchase when you create a delivery group.
//
// @param request - ListNodeInstanceTypeRequest
//
// @return ListNodeInstanceTypeResponse
func (client *Client) ListNodeInstanceType(request *ListNodeInstanceTypeRequest) (_result *ListNodeInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeInstanceTypeResponse{}
	_body, _err := client.ListNodeInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource nodes.
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListNodesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListNodesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries resource nodes.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about over-the-air (OTA) update tasks.
//
// @param request - ListOtaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOtaTaskResponse
func (client *Client) ListOtaTaskWithOptions(request *ListOtaTaskRequest, runtime *util.RuntimeOptions) (_result *ListOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OtaType)) {
		body["OtaType"] = request.OtaType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListOtaTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListOtaTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about over-the-air (OTA) update tasks.
//
// @param request - ListOtaTaskRequest
//
// @return ListOtaTaskResponse
func (client *Client) ListOtaTask(request *ListOtaTaskRequest) (_result *ListOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOtaTaskResponse{}
	_body, _err := client.ListOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions that are supported by App Streaming.
//
// Description:
//
// >  All supported regions instead of available regions are returned by this operation. For more information, see [Supported regions](https://help.aliyun.com/document_detail/426036.html).
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizSource)) {
		query["BizSource"] = request.BizSource
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the regions that are supported by App Streaming.
//
// Description:
//
// >  All supported regions instead of available regions are returned by this operation. For more information, see [Supported regions](https://help.aliyun.com/document_detail/426036.html).
//
// @param request - ListRegionsRequest
//
// @return ListRegionsResponse
func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 项目的会话包列表
//
// @param request - ListSessionPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionPackagesResponse
func (client *Client) ListSessionPackagesWithOptions(request *ListSessionPackagesRequest, runtime *util.RuntimeOptions) (_result *ListSessionPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageId)) {
		query["SessionPackageId"] = request.SessionPackageId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionPackageName)) {
		query["SessionPackageName"] = request.SessionPackageName
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.StateList)) {
		query["StateList"] = request.StateList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSessionPackages"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSessionPackagesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSessionPackagesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 项目的会话包列表
//
// @param request - ListSessionPackagesRequest
//
// @return ListSessionPackagesResponse
func (client *Client) ListSessionPackages(request *ListSessionPackagesRequest) (_result *ListSessionPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSessionPackagesResponse{}
	_body, _err := client.ListSessionPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags added to one or more cloud resources.
//
// @param request - ListTagCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagCloudResourcesResponse
func (client *Client) ListTagCloudResourcesWithOptions(request *ListTagCloudResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagCloudResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagCloudResources"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagCloudResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagCloudResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the tags added to one or more cloud resources.
//
// @param request - ListTagCloudResourcesRequest
//
// @return ListTagCloudResourcesResponse
func (client *Client) ListTagCloudResources(request *ListTagCloudResourcesRequest) (_result *ListTagCloudResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagCloudResourcesResponse{}
	_body, _err := client.ListTagCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of the administrator account, such as whether the resource expiration reminder feature is enabled.
//
// @param request - ListTenantConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantConfigResponse
func (client *Client) ListTenantConfigWithOptions(runtime *util.RuntimeOptions) (_result *ListTenantConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListTenantConfig"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTenantConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTenantConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the configurations of the administrator account, such as whether the resource expiration reminder feature is enabled.
//
// @return ListTenantConfigResponse
func (client *Client) ListTenantConfig() (_result *ListTenantConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTenantConfigResponse{}
	_body, _err := client.ListTenantConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Closes all sessions in a pay-as-you-go delivery group for which a scheduled scaling policy is used.
//
// Description:
//
// >  This operation can be called only if you use a pay-as-you-go delivery group for which a scheduled scaling policy is used and if you call the operation at a time other than the scheduled time.
//
// @param request - LogOffAllSessionsInAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogOffAllSessionsInAppInstanceGroupResponse
func (client *Client) LogOffAllSessionsInAppInstanceGroupWithOptions(request *LogOffAllSessionsInAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *LogOffAllSessionsInAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LogOffAllSessionsInAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &LogOffAllSessionsInAppInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &LogOffAllSessionsInAppInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Closes all sessions in a pay-as-you-go delivery group for which a scheduled scaling policy is used.
//
// Description:
//
// >  This operation can be called only if you use a pay-as-you-go delivery group for which a scheduled scaling policy is used and if you call the operation at a time other than the scheduled time.
//
// @param request - LogOffAllSessionsInAppInstanceGroupRequest
//
// @return LogOffAllSessionsInAppInstanceGroupResponse
func (client *Client) LogOffAllSessionsInAppInstanceGroup(request *LogOffAllSessionsInAppInstanceGroupRequest) (_result *LogOffAllSessionsInAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LogOffAllSessionsInAppInstanceGroupResponse{}
	_body, _err := client.LogOffAllSessionsInAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the general policies of a delivery group, including the number of concurrent sessions and the retention period of disconnected sessions.
//
// @param tmpReq - ModifyAppInstanceGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppInstanceGroupAttributeResponse
func (client *Client) ModifyAppInstanceGroupAttributeWithOptions(tmpReq *ModifyAppInstanceGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAppInstanceGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppInstanceGroupAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Network)) {
		request.NetworkShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Network, tea.String("Network"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodePool)) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, tea.String("NodePool"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SecurityPolicy)) {
		request.SecurityPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityPolicy, tea.String("SecurityPolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StoragePolicy)) {
		request.StoragePolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StoragePolicy, tea.String("StoragePolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolShrink)) {
		query["NodePool"] = request.NodePoolShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionTimeout)) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NetworkShrink)) {
		body["Network"] = request.NetworkShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PerSessionPerApp)) {
		body["PerSessionPerApp"] = request.PerSessionPerApp
	}

	if !tea.BoolValue(util.IsUnset(request.PreOpenAppId)) {
		body["PreOpenAppId"] = request.PreOpenAppId
	}

	if !tea.BoolValue(util.IsUnset(request.PreOpenMode)) {
		body["PreOpenMode"] = request.PreOpenMode
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityPolicyShrink)) {
		body["SecurityPolicy"] = request.SecurityPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StoragePolicyShrink)) {
		body["StoragePolicy"] = request.StoragePolicyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppInstanceGroupAttribute"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAppInstanceGroupAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAppInstanceGroupAttributeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the general policies of a delivery group, including the number of concurrent sessions and the retention period of disconnected sessions.
//
// @param request - ModifyAppInstanceGroupAttributeRequest
//
// @return ModifyAppInstanceGroupAttributeResponse
func (client *Client) ModifyAppInstanceGroupAttribute(request *ModifyAppInstanceGroupAttributeRequest) (_result *ModifyAppInstanceGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppInstanceGroupAttributeResponse{}
	_body, _err := client.ModifyAppInstanceGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改策略信息
//
// @param tmpReq - ModifyAppPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppPolicyResponse
func (client *Client) ModifyAppPolicyWithOptions(tmpReq *ModifyAppPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyAppPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.VideoPolicy)) {
		request.VideoPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoPolicy, tea.String("VideoPolicy"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppPolicyId)) {
		query["AppPolicyId"] = request.AppPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.VideoPolicyShrink)) {
		query["VideoPolicy"] = request.VideoPolicyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppPolicy"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAppPolicyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAppPolicyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 修改策略信息
//
// @param request - ModifyAppPolicyRequest
//
// @return ModifyAppPolicyResponse
func (client *Client) ModifyAppPolicy(request *ModifyAppPolicyRequest) (_result *ModifyAppPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppPolicyResponse{}
	_body, _err := client.ModifyAppPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the number of nodes in a subscription delivery group.
//
// @param tmpReq - ModifyNodePoolAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodePoolAmountResponse
func (client *Client) ModifyNodePoolAmountWithOptions(tmpReq *ModifyNodePoolAmountRequest, runtime *util.RuntimeOptions) (_result *ModifyNodePoolAmountResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyNodePoolAmountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodePool)) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePool, tea.String("NodePool"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolShrink)) {
		body["NodePool"] = request.NodePoolShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodePoolAmount"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyNodePoolAmountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyNodePoolAmountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the number of nodes in a subscription delivery group.
//
// @param request - ModifyNodePoolAmountRequest
//
// @return ModifyNodePoolAmountResponse
func (client *Client) ModifyNodePoolAmount(request *ModifyNodePoolAmountRequest) (_result *ModifyNodePoolAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNodePoolAmountResponse{}
	_body, _err := client.ModifyNodePoolAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - ModifyNodePoolAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodePoolAttributeResponse
func (client *Client) ModifyNodePoolAttributeWithOptions(tmpReq *ModifyNodePoolAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyNodePoolAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyNodePoolAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodePoolStrategy)) {
		request.NodePoolStrategyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodePoolStrategy, tea.String("NodePoolStrategy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeCapacity)) {
		body["NodeCapacity"] = request.NodeCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolStrategyShrink)) {
		body["NodePoolStrategy"] = request.NodePoolStrategyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PoolId)) {
		body["PoolId"] = request.PoolId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodePoolAttribute"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyNodePoolAttributeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyNodePoolAttributeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyNodePoolAttributeRequest
//
// @return ModifyNodePoolAttributeResponse
func (client *Client) ModifyNodePoolAttribute(request *ModifyNodePoolAttributeRequest) (_result *ModifyNodePoolAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNodePoolAttributeResponse{}
	_body, _err := client.ModifyNodePoolAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of the administrator account, such as whether to enable the resource expiration reminder feature.
//
// @param request - ModifyTenantConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTenantConfigResponse
func (client *Client) ModifyTenantConfigWithOptions(request *ModifyTenantConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupExpireRemind)) {
		body["AppInstanceGroupExpireRemind"] = request.AppInstanceGroupExpireRemind
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantConfig"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyTenantConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyTenantConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the configurations of the administrator account, such as whether to enable the resource expiration reminder feature.
//
// @param request - ModifyTenantConfigRequest
//
// @return ModifyTenantConfigResponse
func (client *Client) ModifyTenantConfig(request *ModifyTenantConfigRequest) (_result *ModifyTenantConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantConfigResponse{}
	_body, _err := client.ModifyTenantConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the assigned users that are added to a delivery group by page.
//
// @param request - PageListAppInstanceGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageListAppInstanceGroupUserResponse
func (client *Client) PageListAppInstanceGroupUserWithOptions(request *PageListAppInstanceGroupUserRequest, runtime *util.RuntimeOptions) (_result *PageListAppInstanceGroupUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageListAppInstanceGroupUser"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &PageListAppInstanceGroupUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &PageListAppInstanceGroupUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the assigned users that are added to a delivery group by page.
//
// @param request - PageListAppInstanceGroupUserRequest
//
// @return PageListAppInstanceGroupUserResponse
func (client *Client) PageListAppInstanceGroupUser(request *PageListAppInstanceGroupUserRequest) (_result *PageListAppInstanceGroupUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageListAppInstanceGroupUserResponse{}
	_body, _err := client.PageListAppInstanceGroupUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a delivery group.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the [billing methods and prices](https://help.aliyun.com/document_detail/426039.html) of App Streaming.
//
// @param request - RenewAppInstanceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAppInstanceGroupResponse
func (client *Client) RenewAppInstanceGroupWithOptions(request *RenewAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *RenewAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		query["PromotionId"] = request.PromotionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RenewAppInstanceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RenewAppInstanceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Renews a delivery group.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the [billing methods and prices](https://help.aliyun.com/document_detail/426039.html) of App Streaming.
//
// @param request - RenewAppInstanceGroupRequest
//
// @return RenewAppInstanceGroupResponse
func (client *Client) RenewAppInstanceGroup(request *RenewAppInstanceGroupRequest) (_result *RenewAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewAppInstanceGroupResponse{}
	_body, _err := client.RenewAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为云资源创建并绑定标签
//
// @param request - TagCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagCloudResourcesResponse
func (client *Client) TagCloudResourcesWithOptions(request *TagCloudResourcesRequest, runtime *util.RuntimeOptions) (_result *TagCloudResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagCloudResources"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagCloudResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagCloudResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 为云资源创建并绑定标签
//
// @param request - TagCloudResourcesRequest
//
// @return TagCloudResourcesResponse
func (client *Client) TagCloudResources(request *TagCloudResourcesRequest) (_result *TagCloudResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagCloudResourcesResponse{}
	_body, _err := client.TagCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a user and a session.
//
// @param request - UnbindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindResponse
func (client *Client) UnbindWithOptions(request *UnbindRequest, runtime *util.RuntimeOptions) (_result *UnbindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstancePersistentId)) {
		body["AppInstancePersistentId"] = request.AppInstancePersistentId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Unbind"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UnbindResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UnbindResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Unbinds a user and a session.
//
// @param request - UnbindRequest
//
// @return UnbindResponse
func (client *Client) Unbind(request *UnbindRequest) (_result *UnbindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindResponse{}
	_body, _err := client.UnbindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from cloud resources.
//
// @param request - UntagCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagCloudResourcesResponse
func (client *Client) UntagCloudResourcesWithOptions(request *UntagCloudResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagCloudResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeys)) {
		body["TagKeys"] = request.TagKeys
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagCloudResources"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagCloudResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagCloudResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Removes tags from cloud resources.
//
// @param request - UntagCloudResourcesRequest
//
// @return UntagCloudResourcesResponse
func (client *Client) UntagCloudResources(request *UntagCloudResourcesRequest) (_result *UntagCloudResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagCloudResourcesResponse{}
	_body, _err := client.UntagCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the image of a delivery group.
//
// Description:
//
// *
//
// **Warning*	- After the image is updated, the end user session accessing the cloud application will be disconnected. Exercise caution to avoid end user data loss.
//
// >  After the image of the delivery group is updated, the change takes effect on the terminal in approximately 2 minutes.
//
// @param request - UpdateAppInstanceGroupImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppInstanceGroupImageResponse
func (client *Client) UpdateAppInstanceGroupImageWithOptions(request *UpdateAppInstanceGroupImageRequest, runtime *util.RuntimeOptions) (_result *UpdateAppInstanceGroupImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppInstanceGroupImage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAppInstanceGroupImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAppInstanceGroupImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the image of a delivery group.
//
// Description:
//
// *
//
// **Warning*	- After the image is updated, the end user session accessing the cloud application will be disconnected. Exercise caution to avoid end user data loss.
//
// >  After the image of the delivery group is updated, the change takes effect on the terminal in approximately 2 minutes.
//
// @param request - UpdateAppInstanceGroupImageRequest
//
// @return UpdateAppInstanceGroupImageResponse
func (client *Client) UpdateAppInstanceGroupImage(request *UpdateAppInstanceGroupImageRequest) (_result *UpdateAppInstanceGroupImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppInstanceGroupImageResponse{}
	_body, _err := client.UpdateAppInstanceGroupImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
