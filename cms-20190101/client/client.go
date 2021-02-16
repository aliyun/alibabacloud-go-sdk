// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddTagsRequest struct {
	Tag      []*AddTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	GroupIds []*string            `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s AddTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) SetTag(v []*AddTagsRequestTag) *AddTagsRequest {
	s.Tag = v
	return s
}

func (s *AddTagsRequest) SetGroupIds(v []*string) *AddTagsRequest {
	s.GroupIds = v
	return s
}

type AddTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s AddTagsRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsRequestTag) SetKey(v string) *AddTagsRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsRequestTag) SetValue(v string) *AddTagsRequestTag {
	s.Value = &v
	return s
}

type AddTagsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsResponseBody) SetMessage(v string) *AddTagsResponseBody {
	s.Message = &v
	return s
}

func (s *AddTagsResponseBody) SetRequestId(v string) *AddTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagsResponseBody) SetCode(v string) *AddTagsResponseBody {
	s.Code = &v
	return s
}

func (s *AddTagsResponseBody) SetSuccess(v bool) *AddTagsResponseBody {
	s.Success = &v
	return s
}

type AddTagsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTagsResponse) GoString() string {
	return s.String()
}

func (s *AddTagsResponse) SetHeaders(v map[string]*string) *AddTagsResponse {
	s.Headers = v
	return s
}

func (s *AddTagsResponse) SetBody(v *AddTagsResponseBody) *AddTagsResponse {
	s.Body = v
	return s
}

type ApplyMetricRuleTemplateRequest struct {
	SilenceTime     *int64  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	GroupId         *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	TemplateIds     *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
	EnableStartTime *int64  `json:"EnableStartTime,omitempty" xml:"EnableStartTime,omitempty"`
	EnableEndTime   *int64  `json:"EnableEndTime,omitempty" xml:"EnableEndTime,omitempty"`
	NotifyLevel     *int64  `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	ApplyMode       *string `json:"ApplyMode,omitempty" xml:"ApplyMode,omitempty"`
	Webhook         *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s ApplyMetricRuleTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateRequest) SetSilenceTime(v int64) *ApplyMetricRuleTemplateRequest {
	s.SilenceTime = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetGroupId(v int64) *ApplyMetricRuleTemplateRequest {
	s.GroupId = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetTemplateIds(v string) *ApplyMetricRuleTemplateRequest {
	s.TemplateIds = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetEnableStartTime(v int64) *ApplyMetricRuleTemplateRequest {
	s.EnableStartTime = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetEnableEndTime(v int64) *ApplyMetricRuleTemplateRequest {
	s.EnableEndTime = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetNotifyLevel(v int64) *ApplyMetricRuleTemplateRequest {
	s.NotifyLevel = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetApplyMode(v string) *ApplyMetricRuleTemplateRequest {
	s.ApplyMode = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetWebhook(v string) *ApplyMetricRuleTemplateRequest {
	s.Webhook = &v
	return s
}

type ApplyMetricRuleTemplateResponseBody struct {
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource  *ApplyMetricRuleTemplateResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyMetricRuleTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponseBody) SetMessage(v string) *ApplyMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetRequestId(v string) *ApplyMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetResource(v *ApplyMetricRuleTemplateResponseBodyResource) *ApplyMetricRuleTemplateResponseBody {
	s.Resource = v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetCode(v int32) *ApplyMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBody) SetSuccess(v bool) *ApplyMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

type ApplyMetricRuleTemplateResponseBodyResource struct {
	AlertResults []*ApplyMetricRuleTemplateResponseBodyResourceAlertResults `json:"AlertResults,omitempty" xml:"AlertResults,omitempty" type:"Repeated"`
	GroupId      *int64                                                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ApplyMetricRuleTemplateResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponseBodyResource) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponseBodyResource) SetAlertResults(v []*ApplyMetricRuleTemplateResponseBodyResourceAlertResults) *ApplyMetricRuleTemplateResponseBodyResource {
	s.AlertResults = v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResource) SetGroupId(v int64) *ApplyMetricRuleTemplateResponseBodyResource {
	s.GroupId = &v
	return s
}

type ApplyMetricRuleTemplateResponseBodyResourceAlertResults struct {
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ApplyMetricRuleTemplateResponseBodyResourceAlertResults) String() string {
	return tea.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponseBodyResourceAlertResults) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetGroupId(v int64) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.GroupId = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetSuccess(v bool) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.Success = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetCode(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.Code = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetMessage(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.Message = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetRuleId(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.RuleId = &v
	return s
}

func (s *ApplyMetricRuleTemplateResponseBodyResourceAlertResults) SetRuleName(v string) *ApplyMetricRuleTemplateResponseBodyResourceAlertResults {
	s.RuleName = &v
	return s
}

type ApplyMetricRuleTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyMetricRuleTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *ApplyMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *ApplyMetricRuleTemplateResponse) SetBody(v *ApplyMetricRuleTemplateResponseBody) *ApplyMetricRuleTemplateResponse {
	s.Body = v
	return s
}

type CreateCmsCallNumOrderRequest struct {
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoRenewPeriod *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	AutoPay         *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon   *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	PhoneCount      *string `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
}

func (s CreateCmsCallNumOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsCallNumOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateCmsCallNumOrderRequest) SetPeriod(v int32) *CreateCmsCallNumOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateCmsCallNumOrderRequest) SetPeriodUnit(v string) *CreateCmsCallNumOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCmsCallNumOrderRequest) SetAutoRenewPeriod(v int32) *CreateCmsCallNumOrderRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateCmsCallNumOrderRequest) SetAutoPay(v bool) *CreateCmsCallNumOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCmsCallNumOrderRequest) SetAutoUseCoupon(v bool) *CreateCmsCallNumOrderRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateCmsCallNumOrderRequest) SetPhoneCount(v string) *CreateCmsCallNumOrderRequest {
	s.PhoneCount = &v
	return s
}

type CreateCmsCallNumOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateCmsCallNumOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsCallNumOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCmsCallNumOrderResponseBody) SetRequestId(v string) *CreateCmsCallNumOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCmsCallNumOrderResponseBody) SetOrderId(v string) *CreateCmsCallNumOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateCmsCallNumOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCmsCallNumOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCmsCallNumOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsCallNumOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateCmsCallNumOrderResponse) SetHeaders(v map[string]*string) *CreateCmsCallNumOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateCmsCallNumOrderResponse) SetBody(v *CreateCmsCallNumOrderResponseBody) *CreateCmsCallNumOrderResponse {
	s.Body = v
	return s
}

type CreateCmsOrderRequest struct {
	Period           *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoRenewPeriod  *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	AutoPay          *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon    *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	SuggestType      *string `json:"SuggestType,omitempty" xml:"SuggestType,omitempty"`
	ApiCount         *string `json:"ApiCount,omitempty" xml:"ApiCount,omitempty"`
	SiteOperatorNum  *string `json:"SiteOperatorNum,omitempty" xml:"SiteOperatorNum,omitempty"`
	EventStoreTime   *string `json:"EventStoreTime,omitempty" xml:"EventStoreTime,omitempty"`
	LogMonitorStream *string `json:"LogMonitorStream,omitempty" xml:"LogMonitorStream,omitempty"`
	SiteTaskNum      *string `json:"SiteTaskNum,omitempty" xml:"SiteTaskNum,omitempty"`
	EventStoreNum    *string `json:"EventStoreNum,omitempty" xml:"EventStoreNum,omitempty"`
	SiteEcsNum       *string `json:"SiteEcsNum,omitempty" xml:"SiteEcsNum,omitempty"`
	CustomTimeSeries *string `json:"CustomTimeSeries,omitempty" xml:"CustomTimeSeries,omitempty"`
	SmsCount         *string `json:"SmsCount,omitempty" xml:"SmsCount,omitempty"`
	PhoneCount       *string `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
}

func (s CreateCmsOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateCmsOrderRequest) SetPeriod(v int32) *CreateCmsOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateCmsOrderRequest) SetPeriodUnit(v string) *CreateCmsOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCmsOrderRequest) SetAutoRenewPeriod(v int32) *CreateCmsOrderRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateCmsOrderRequest) SetAutoPay(v bool) *CreateCmsOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCmsOrderRequest) SetAutoUseCoupon(v bool) *CreateCmsOrderRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateCmsOrderRequest) SetPayType(v string) *CreateCmsOrderRequest {
	s.PayType = &v
	return s
}

func (s *CreateCmsOrderRequest) SetSuggestType(v string) *CreateCmsOrderRequest {
	s.SuggestType = &v
	return s
}

func (s *CreateCmsOrderRequest) SetApiCount(v string) *CreateCmsOrderRequest {
	s.ApiCount = &v
	return s
}

func (s *CreateCmsOrderRequest) SetSiteOperatorNum(v string) *CreateCmsOrderRequest {
	s.SiteOperatorNum = &v
	return s
}

func (s *CreateCmsOrderRequest) SetEventStoreTime(v string) *CreateCmsOrderRequest {
	s.EventStoreTime = &v
	return s
}

func (s *CreateCmsOrderRequest) SetLogMonitorStream(v string) *CreateCmsOrderRequest {
	s.LogMonitorStream = &v
	return s
}

func (s *CreateCmsOrderRequest) SetSiteTaskNum(v string) *CreateCmsOrderRequest {
	s.SiteTaskNum = &v
	return s
}

func (s *CreateCmsOrderRequest) SetEventStoreNum(v string) *CreateCmsOrderRequest {
	s.EventStoreNum = &v
	return s
}

func (s *CreateCmsOrderRequest) SetSiteEcsNum(v string) *CreateCmsOrderRequest {
	s.SiteEcsNum = &v
	return s
}

func (s *CreateCmsOrderRequest) SetCustomTimeSeries(v string) *CreateCmsOrderRequest {
	s.CustomTimeSeries = &v
	return s
}

func (s *CreateCmsOrderRequest) SetSmsCount(v string) *CreateCmsOrderRequest {
	s.SmsCount = &v
	return s
}

func (s *CreateCmsOrderRequest) SetPhoneCount(v string) *CreateCmsOrderRequest {
	s.PhoneCount = &v
	return s
}

type CreateCmsOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateCmsOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCmsOrderResponseBody) SetRequestId(v string) *CreateCmsOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCmsOrderResponseBody) SetOrderId(v string) *CreateCmsOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateCmsOrderResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCmsOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCmsOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateCmsOrderResponse) SetHeaders(v map[string]*string) *CreateCmsOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateCmsOrderResponse) SetBody(v *CreateCmsOrderResponseBody) *CreateCmsOrderResponse {
	s.Body = v
	return s
}

type CreateCmsSmspackageOrderRequest struct {
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoRenewPeriod *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	AutoPay         *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon   *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	SmsCount        *string `json:"SmsCount,omitempty" xml:"SmsCount,omitempty"`
}

func (s CreateCmsSmspackageOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsSmspackageOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateCmsSmspackageOrderRequest) SetPeriod(v int32) *CreateCmsSmspackageOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateCmsSmspackageOrderRequest) SetPeriodUnit(v string) *CreateCmsSmspackageOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCmsSmspackageOrderRequest) SetAutoRenewPeriod(v int32) *CreateCmsSmspackageOrderRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateCmsSmspackageOrderRequest) SetAutoPay(v bool) *CreateCmsSmspackageOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCmsSmspackageOrderRequest) SetAutoUseCoupon(v bool) *CreateCmsSmspackageOrderRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateCmsSmspackageOrderRequest) SetSmsCount(v string) *CreateCmsSmspackageOrderRequest {
	s.SmsCount = &v
	return s
}

type CreateCmsSmspackageOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateCmsSmspackageOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsSmspackageOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCmsSmspackageOrderResponseBody) SetRequestId(v string) *CreateCmsSmspackageOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCmsSmspackageOrderResponseBody) SetOrderId(v string) *CreateCmsSmspackageOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateCmsSmspackageOrderResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCmsSmspackageOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCmsSmspackageOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCmsSmspackageOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateCmsSmspackageOrderResponse) SetHeaders(v map[string]*string) *CreateCmsSmspackageOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateCmsSmspackageOrderResponse) SetBody(v *CreateCmsSmspackageOrderResponseBody) *CreateCmsSmspackageOrderResponse {
	s.Body = v
	return s
}

type CreateDynamicTagGroupRequest struct {
	TagKey                     *string                                     `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	EnableSubscribeEvent       *bool                                       `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	EnableInstallAgent         *bool                                       `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	MatchExpressFilterRelation *string                                     `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	MatchExpress               []*CreateDynamicTagGroupRequestMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
	ContactGroupList           []*string                                   `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
	TemplateIdList             []*string                                   `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s CreateDynamicTagGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDynamicTagGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupRequest) SetTagKey(v string) *CreateDynamicTagGroupRequest {
	s.TagKey = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetEnableSubscribeEvent(v bool) *CreateDynamicTagGroupRequest {
	s.EnableSubscribeEvent = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetEnableInstallAgent(v bool) *CreateDynamicTagGroupRequest {
	s.EnableInstallAgent = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetMatchExpressFilterRelation(v string) *CreateDynamicTagGroupRequest {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetMatchExpress(v []*CreateDynamicTagGroupRequestMatchExpress) *CreateDynamicTagGroupRequest {
	s.MatchExpress = v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetContactGroupList(v []*string) *CreateDynamicTagGroupRequest {
	s.ContactGroupList = v
	return s
}

func (s *CreateDynamicTagGroupRequest) SetTemplateIdList(v []*string) *CreateDynamicTagGroupRequest {
	s.TemplateIdList = v
	return s
}

type CreateDynamicTagGroupRequestMatchExpress struct {
	TagValueMatchFunction *string `json:"TagValueMatchFunction,omitempty" xml:"TagValueMatchFunction,omitempty"`
	TagValue              *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateDynamicTagGroupRequestMatchExpress) String() string {
	return tea.Prettify(s)
}

func (s CreateDynamicTagGroupRequestMatchExpress) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupRequestMatchExpress) SetTagValueMatchFunction(v string) *CreateDynamicTagGroupRequestMatchExpress {
	s.TagValueMatchFunction = &v
	return s
}

func (s *CreateDynamicTagGroupRequestMatchExpress) SetTagValue(v string) *CreateDynamicTagGroupRequestMatchExpress {
	s.TagValue = &v
	return s
}

type CreateDynamicTagGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDynamicTagGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDynamicTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupResponseBody) SetMessage(v string) *CreateDynamicTagGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) SetRequestId(v string) *CreateDynamicTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) SetCode(v string) *CreateDynamicTagGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) SetSuccess(v bool) *CreateDynamicTagGroupResponseBody {
	s.Success = &v
	return s
}

type CreateDynamicTagGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDynamicTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDynamicTagGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDynamicTagGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupResponse) SetHeaders(v map[string]*string) *CreateDynamicTagGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateDynamicTagGroupResponse) SetBody(v *CreateDynamicTagGroupResponseBody) *CreateDynamicTagGroupResponse {
	s.Body = v
	return s
}

type CreateGroupMetricRulesRequest struct {
	GroupId          *int64                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupMetricRules []*CreateGroupMetricRulesRequestGroupMetricRules `json:"GroupMetricRules,omitempty" xml:"GroupMetricRules,omitempty" type:"Repeated"`
}

func (s CreateGroupMetricRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequest) SetGroupId(v int64) *CreateGroupMetricRulesRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupMetricRulesRequest) SetGroupMetricRules(v []*CreateGroupMetricRulesRequestGroupMetricRules) *CreateGroupMetricRulesRequest {
	s.GroupMetricRules = v
	return s
}

type CreateGroupMetricRulesRequestGroupMetricRules struct {
	Escalations         *CreateGroupMetricRulesRequestGroupMetricRulesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" require:"true" type:"Struct"`
	MetricName          *string                                                   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	NoEffectiveInterval *string                                                   `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	EffectiveInterval   *string                                                   `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	RuleId              *string                                                   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Dimensions          *string                                                   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	SilenceTime         *int32                                                    `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Webhook             *string                                                   `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Namespace           *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	EmailSubject        *string                                                   `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	Period              *string                                                   `json:"Period,omitempty" xml:"Period,omitempty"`
	RuleName            *string                                                   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Interval            *string                                                   `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Category            *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRules) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRules) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEscalations(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Escalations = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetMetricName(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.MetricName = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNoEffectiveInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.NoEffectiveInterval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEffectiveInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.EffectiveInterval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetRuleId(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.RuleId = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetDimensions(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Dimensions = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetSilenceTime(v int32) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.SilenceTime = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetWebhook(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Webhook = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNamespace(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Namespace = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEmailSubject(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.EmailSubject = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetPeriod(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Period = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetRuleName(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.RuleName = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Interval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetCategory(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Category = &v
	return s
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalations struct {
	Info     *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" require:"true" type:"Struct"`
	Warn     *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" require:"true" type:"Struct"`
	Critical *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" require:"true" type:"Struct"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalations) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetInfo(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Info = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetWarn(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Warn = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetCritical(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Critical = v
	return s
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Statistics = &v
	return s
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical struct {
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

type CreateGroupMetricRulesResponseBody struct {
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources *CreateGroupMetricRulesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Code      *int32                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGroupMetricRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponseBody) SetMessage(v string) *CreateGroupMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetRequestId(v string) *CreateGroupMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetResources(v *CreateGroupMetricRulesResponseBodyResources) *CreateGroupMetricRulesResponseBody {
	s.Resources = v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetCode(v int32) *CreateGroupMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBody) SetSuccess(v bool) *CreateGroupMetricRulesResponseBody {
	s.Success = &v
	return s
}

type CreateGroupMetricRulesResponseBodyResources struct {
	AlertResult []*CreateGroupMetricRulesResponseBodyResourcesAlertResult `json:"AlertResult,omitempty" xml:"AlertResult,omitempty" type:"Repeated"`
}

func (s CreateGroupMetricRulesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponseBodyResources) SetAlertResult(v []*CreateGroupMetricRulesResponseBodyResourcesAlertResult) *CreateGroupMetricRulesResponseBodyResources {
	s.AlertResult = v
	return s
}

type CreateGroupMetricRulesResponseBodyResourcesAlertResult struct {
	Success  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code     *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId   *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateGroupMetricRulesResponseBodyResourcesAlertResult) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesResponseBodyResourcesAlertResult) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetSuccess(v bool) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.Success = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetCode(v int32) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.Code = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetMessage(v string) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.Message = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetRuleName(v string) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.RuleName = &v
	return s
}

func (s *CreateGroupMetricRulesResponseBodyResourcesAlertResult) SetRuleId(v string) *CreateGroupMetricRulesResponseBodyResourcesAlertResult {
	s.RuleId = &v
	return s
}

type CreateGroupMetricRulesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupMetricRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesResponse) SetHeaders(v map[string]*string) *CreateGroupMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupMetricRulesResponse) SetBody(v *CreateGroupMetricRulesResponseBody) *CreateGroupMetricRulesResponse {
	s.Body = v
	return s
}

type CreateGroupMonitoringAgentProcessRequest struct {
	GroupId                    *string                                                 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ProcessName                *string                                                 `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	MatchExpressFilterRelation *string                                                 `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	MatchExpress               []*CreateGroupMonitoringAgentProcessRequestMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
	AlertConfig                []*CreateGroupMonitoringAgentProcessRequestAlertConfig  `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s CreateGroupMonitoringAgentProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetGroupId(v string) *CreateGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetProcessName(v string) *CreateGroupMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetMatchExpressFilterRelation(v string) *CreateGroupMonitoringAgentProcessRequest {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetMatchExpress(v []*CreateGroupMonitoringAgentProcessRequestMatchExpress) *CreateGroupMonitoringAgentProcessRequest {
	s.MatchExpress = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequest) SetAlertConfig(v []*CreateGroupMonitoringAgentProcessRequestAlertConfig) *CreateGroupMonitoringAgentProcessRequest {
	s.AlertConfig = v
	return s
}

type CreateGroupMonitoringAgentProcessRequestMatchExpress struct {
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessRequestMatchExpress) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessRequestMatchExpress) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) SetValue(v string) *CreateGroupMonitoringAgentProcessRequestMatchExpress {
	s.Value = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) SetFunction(v string) *CreateGroupMonitoringAgentProcessRequestMatchExpress {
	s.Function = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestMatchExpress) SetName(v string) *CreateGroupMonitoringAgentProcessRequestMatchExpress {
	s.Name = &v
	return s
}

type CreateGroupMonitoringAgentProcessRequestAlertConfig struct {
	SilenceTime         *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	ComparisonOperator  *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Webhook             *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Times               *string `json:"Times,omitempty" xml:"Times,omitempty"`
	EscalationsLevel    *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	EffectiveInterval   *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	Threshold           *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics          *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessRequestAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetSilenceTime(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetComparisonOperator(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetWebhook(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Webhook = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetTimes(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Times = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetEscalationsLevel(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.EscalationsLevel = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetNoEffectiveInterval(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.NoEffectiveInterval = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetEffectiveInterval(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.EffectiveInterval = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetThreshold(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessRequestAlertConfig) SetStatistics(v string) *CreateGroupMonitoringAgentProcessRequestAlertConfig {
	s.Statistics = &v
	return s
}

type CreateGroupMonitoringAgentProcessResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *CreateGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *CreateGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetCode(v string) *CreateGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *CreateGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

type CreateGroupMonitoringAgentProcessResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGroupMonitoringAgentProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *CreateGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponse) SetBody(v *CreateGroupMonitoringAgentProcessResponseBody) *CreateGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

type CreateHostAvailabilityRequest struct {
	TaskOption                *CreateHostAvailabilityRequestTaskOption                  `json:"TaskOption,omitempty" xml:"TaskOption,omitempty" type:"Struct"`
	AlertConfig               *CreateHostAvailabilityRequestAlertConfig                 `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	GroupId                   *int64                                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	TaskName                  *string                                                   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskScope                 *string                                                   `json:"TaskScope,omitempty" xml:"TaskScope,omitempty"`
	TaskType                  *string                                                   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	AlertConfigEscalationList []*CreateHostAvailabilityRequestAlertConfigEscalationList `json:"AlertConfigEscalationList,omitempty" xml:"AlertConfigEscalationList,omitempty" type:"Repeated"`
	InstanceList              []*string                                                 `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s CreateHostAvailabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequest) SetTaskOption(v *CreateHostAvailabilityRequestTaskOption) *CreateHostAvailabilityRequest {
	s.TaskOption = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetAlertConfig(v *CreateHostAvailabilityRequestAlertConfig) *CreateHostAvailabilityRequest {
	s.AlertConfig = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetGroupId(v int64) *CreateHostAvailabilityRequest {
	s.GroupId = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetTaskName(v string) *CreateHostAvailabilityRequest {
	s.TaskName = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetTaskScope(v string) *CreateHostAvailabilityRequest {
	s.TaskScope = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetTaskType(v string) *CreateHostAvailabilityRequest {
	s.TaskType = &v
	return s
}

func (s *CreateHostAvailabilityRequest) SetAlertConfigEscalationList(v []*CreateHostAvailabilityRequestAlertConfigEscalationList) *CreateHostAvailabilityRequest {
	s.AlertConfigEscalationList = v
	return s
}

func (s *CreateHostAvailabilityRequest) SetInstanceList(v []*string) *CreateHostAvailabilityRequest {
	s.InstanceList = v
	return s
}

type CreateHostAvailabilityRequestTaskOption struct {
	HttpURI                  *string `json:"HttpURI,omitempty" xml:"HttpURI,omitempty"`
	TelnetOrPingHost         *string `json:"TelnetOrPingHost,omitempty" xml:"TelnetOrPingHost,omitempty"`
	HttpResponseCharset      *string `json:"HttpResponseCharset,omitempty" xml:"HttpResponseCharset,omitempty"`
	HttpPostContent          *string `json:"HttpPostContent,omitempty" xml:"HttpPostContent,omitempty"`
	HttpResponseMatchContent *string `json:"HttpResponseMatchContent,omitempty" xml:"HttpResponseMatchContent,omitempty"`
	HttpMethod               *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpNegative             *bool   `json:"HttpNegative,omitempty" xml:"HttpNegative,omitempty"`
}

func (s CreateHostAvailabilityRequestTaskOption) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAvailabilityRequestTaskOption) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpURI(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpURI = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetTelnetOrPingHost(v string) *CreateHostAvailabilityRequestTaskOption {
	s.TelnetOrPingHost = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpResponseCharset(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpResponseCharset = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpPostContent(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpPostContent = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpResponseMatchContent(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpResponseMatchContent = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpMethod(v string) *CreateHostAvailabilityRequestTaskOption {
	s.HttpMethod = &v
	return s
}

func (s *CreateHostAvailabilityRequestTaskOption) SetHttpNegative(v bool) *CreateHostAvailabilityRequestTaskOption {
	s.HttpNegative = &v
	return s
}

type CreateHostAvailabilityRequestAlertConfig struct {
	NotifyType  *int32  `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	StartTime   *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SilenceTime *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	WebHook     *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s CreateHostAvailabilityRequestAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAvailabilityRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetNotifyType(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.NotifyType = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetStartTime(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.StartTime = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetEndTime(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.EndTime = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetSilenceTime(v int32) *CreateHostAvailabilityRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfig) SetWebHook(v string) *CreateHostAvailabilityRequestAlertConfig {
	s.WebHook = &v
	return s
}

type CreateHostAvailabilityRequestAlertConfigEscalationList struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Times      *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Aggregate  *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
}

func (s CreateHostAvailabilityRequestAlertConfigEscalationList) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAvailabilityRequestAlertConfigEscalationList) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetValue(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Value = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetMetricName(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.MetricName = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetTimes(v int32) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Times = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetOperator(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Operator = &v
	return s
}

func (s *CreateHostAvailabilityRequestAlertConfigEscalationList) SetAggregate(v string) *CreateHostAvailabilityRequestAlertConfigEscalationList {
	s.Aggregate = &v
	return s
}

type CreateHostAvailabilityResponseBody struct {
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateHostAvailabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityResponseBody) SetTaskId(v int64) *CreateHostAvailabilityResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetMessage(v string) *CreateHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetRequestId(v string) *CreateHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetCode(v string) *CreateHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetSuccess(v bool) *CreateHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

type CreateHostAvailabilityResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHostAvailabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityResponse) SetHeaders(v map[string]*string) *CreateHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *CreateHostAvailabilityResponse) SetBody(v *CreateHostAvailabilityResponseBody) *CreateHostAvailabilityResponse {
	s.Body = v
	return s
}

type CreateMetricRuleResourcesRequest struct {
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Overwrite *string `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s CreateMetricRuleResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleResourcesRequest) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleResourcesRequest) SetRuleId(v string) *CreateMetricRuleResourcesRequest {
	s.RuleId = &v
	return s
}

func (s *CreateMetricRuleResourcesRequest) SetOverwrite(v string) *CreateMetricRuleResourcesRequest {
	s.Overwrite = &v
	return s
}

func (s *CreateMetricRuleResourcesRequest) SetResources(v string) *CreateMetricRuleResourcesRequest {
	s.Resources = &v
	return s
}

type CreateMetricRuleResourcesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMetricRuleResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleResourcesResponseBody) SetMessage(v string) *CreateMetricRuleResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMetricRuleResourcesResponseBody) SetRequestId(v string) *CreateMetricRuleResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetricRuleResourcesResponseBody) SetCode(v string) *CreateMetricRuleResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMetricRuleResourcesResponseBody) SetSuccess(v bool) *CreateMetricRuleResourcesResponseBody {
	s.Success = &v
	return s
}

type CreateMetricRuleResourcesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMetricRuleResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMetricRuleResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleResourcesResponse) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleResourcesResponse) SetHeaders(v map[string]*string) *CreateMetricRuleResourcesResponse {
	s.Headers = v
	return s
}

func (s *CreateMetricRuleResourcesResponse) SetBody(v *CreateMetricRuleResourcesResponseBody) *CreateMetricRuleResourcesResponse {
	s.Body = v
	return s
}

type CreateMetricRuleTemplateRequest struct {
	Name           *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	AlertTemplates []*CreateMetricRuleTemplateRequestAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Repeated"`
}

func (s CreateMetricRuleTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequest) SetName(v string) *CreateMetricRuleTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateMetricRuleTemplateRequest) SetDescription(v string) *CreateMetricRuleTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateMetricRuleTemplateRequest) SetAlertTemplates(v []*CreateMetricRuleTemplateRequestAlertTemplates) *CreateMetricRuleTemplateRequest {
	s.AlertTemplates = v
	return s
}

type CreateMetricRuleTemplateRequestAlertTemplates struct {
	Escalations *CreateMetricRuleTemplateRequestAlertTemplatesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" require:"true" type:"Struct"`
	MetricName  *string                                                   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Webhook     *string                                                   `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Namespace   *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RuleName    *string                                                   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Period      *int32                                                    `json:"Period,omitempty" xml:"Period,omitempty"`
	Selector    *string                                                   `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Category    *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplates) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplates) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetEscalations(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Escalations = v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetMetricName(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.MetricName = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetWebhook(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Webhook = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetNamespace(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Namespace = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetRuleName(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.RuleName = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetPeriod(v int32) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Period = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetSelector(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Selector = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetCategory(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Category = &v
	return s
}

type CreateMetricRuleTemplateRequestAlertTemplatesEscalations struct {
	Info     *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" require:"true" type:"Struct"`
	Warn     *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" require:"true" type:"Struct"`
	Critical *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" require:"true" type:"Struct"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalations) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalations) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) SetInfo(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) *CreateMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Info = v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) SetWarn(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) *CreateMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Warn = v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) SetCritical(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) *CreateMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Critical = v
	return s
}

type CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetThreshold(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetStatistics(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetComparisonOperator(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetTimes(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Times = &v
	return s
}

type CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetThreshold(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetTimes(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetComparisonOperator(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetStatistics(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Statistics = &v
	return s
}

type CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical struct {
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetTimes(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetThreshold(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetComparisonOperator(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetStatistics(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Statistics = &v
	return s
}

type CreateMetricRuleTemplateResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMetricRuleTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateResponseBody) SetMessage(v string) *CreateMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetRequestId(v string) *CreateMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetId(v int64) *CreateMetricRuleTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetCode(v int32) *CreateMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMetricRuleTemplateResponseBody) SetSuccess(v bool) *CreateMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

type CreateMetricRuleTemplateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMetricRuleTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *CreateMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateMetricRuleTemplateResponse) SetBody(v *CreateMetricRuleTemplateResponseBody) *CreateMetricRuleTemplateResponse {
	s.Body = v
	return s
}

type CreateMonitorAgentProcessRequest struct {
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
}

func (s CreateMonitorAgentProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorAgentProcessRequest) SetProcessName(v string) *CreateMonitorAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateMonitorAgentProcessRequest) SetInstanceId(v string) *CreateMonitorAgentProcessRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMonitorAgentProcessRequest) SetProcessUser(v string) *CreateMonitorAgentProcessRequest {
	s.ProcessUser = &v
	return s
}

type CreateMonitorAgentProcessResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMonitorAgentProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorAgentProcessResponseBody) SetMessage(v string) *CreateMonitorAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetRequestId(v string) *CreateMonitorAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetId(v int64) *CreateMonitorAgentProcessResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetCode(v string) *CreateMonitorAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetSuccess(v bool) *CreateMonitorAgentProcessResponseBody {
	s.Success = &v
	return s
}

type CreateMonitorAgentProcessResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitorAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitorAgentProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorAgentProcessResponse) SetHeaders(v map[string]*string) *CreateMonitorAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorAgentProcessResponse) SetBody(v *CreateMonitorAgentProcessResponseBody) *CreateMonitorAgentProcessResponse {
	s.Body = v
	return s
}

type CreateMonitorGroupRequest struct {
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ServiceId     *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	BindUrl       *string `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Options       *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s CreateMonitorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupRequest) SetType(v string) *CreateMonitorGroupRequest {
	s.Type = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetGroupName(v string) *CreateMonitorGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetServiceId(v int64) *CreateMonitorGroupRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetBindUrl(v string) *CreateMonitorGroupRequest {
	s.BindUrl = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetContactGroups(v string) *CreateMonitorGroupRequest {
	s.ContactGroups = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetOptions(v string) *CreateMonitorGroupRequest {
	s.Options = &v
	return s
}

type CreateMonitorGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s CreateMonitorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupResponseBody) SetMessage(v string) *CreateMonitorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetRequestId(v string) *CreateMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetCode(v int32) *CreateMonitorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetSuccess(v bool) *CreateMonitorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetGroupId(v int64) *CreateMonitorGroupResponseBody {
	s.GroupId = &v
	return s
}

type CreateMonitorGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupResponse) SetBody(v *CreateMonitorGroupResponseBody) *CreateMonitorGroupResponse {
	s.Body = v
	return s
}

type CreateMonitorGroupByResourceGroupIdRequest struct {
	EnableSubscribeEvent *bool     `json:"EnableSubscribeEvent,omitempty" xml:"EnableSubscribeEvent,omitempty"`
	EnableInstallAgent   *bool     `json:"EnableInstallAgent,omitempty" xml:"EnableInstallAgent,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupName    *string   `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ContactGroupList     []*string `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Repeated"`
}

func (s CreateMonitorGroupByResourceGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupByResourceGroupIdRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetEnableSubscribeEvent(v bool) *CreateMonitorGroupByResourceGroupIdRequest {
	s.EnableSubscribeEvent = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetEnableInstallAgent(v bool) *CreateMonitorGroupByResourceGroupIdRequest {
	s.EnableInstallAgent = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetRegionId(v string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetResourceGroupId(v string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetResourceGroupName(v string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdRequest) SetContactGroupList(v []*string) *CreateMonitorGroupByResourceGroupIdRequest {
	s.ContactGroupList = v
	return s
}

type CreateMonitorGroupByResourceGroupIdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMonitorGroupByResourceGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupByResourceGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetMessage(v string) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetRequestId(v string) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetCode(v string) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetSuccess(v bool) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.Success = &v
	return s
}

type CreateMonitorGroupByResourceGroupIdResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitorGroupByResourceGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitorGroupByResourceGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupByResourceGroupIdResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupByResourceGroupIdResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponse) SetBody(v *CreateMonitorGroupByResourceGroupIdResponseBody) *CreateMonitorGroupByResourceGroupIdResponse {
	s.Body = v
	return s
}

type CreateMonitorGroupInstancesRequest struct {
	GroupId   *string                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Instances []*CreateMonitorGroupInstancesRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s CreateMonitorGroupInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesRequest) SetGroupId(v string) *CreateMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) SetInstances(v []*CreateMonitorGroupInstancesRequestInstances) *CreateMonitorGroupInstancesRequest {
	s.Instances = v
	return s
}

type CreateMonitorGroupInstancesRequestInstances struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitorGroupInstancesRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupInstancesRequestInstances) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetInstanceName(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.InstanceName = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetCategory(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.Category = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetInstanceId(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetRegionId(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.RegionId = &v
	return s
}

type CreateMonitorGroupInstancesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMonitorGroupInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesResponseBody) SetMessage(v string) *CreateMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponseBody) SetRequestId(v string) *CreateMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponseBody) SetCode(v int32) *CreateMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupInstancesResponseBody) SetSuccess(v bool) *CreateMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

type CreateMonitorGroupInstancesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitorGroupInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupInstancesResponse) SetBody(v *CreateMonitorGroupInstancesResponseBody) *CreateMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

type CreateMonitorGroupNotifyPolicyRequest struct {
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateMonitorGroupNotifyPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetPolicyType(v string) *CreateMonitorGroupNotifyPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetGroupId(v string) *CreateMonitorGroupNotifyPolicyRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetStartTime(v int64) *CreateMonitorGroupNotifyPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetEndTime(v int64) *CreateMonitorGroupNotifyPolicyRequest {
	s.EndTime = &v
	return s
}

type CreateMonitorGroupNotifyPolicyResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateMonitorGroupNotifyPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetMessage(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetRequestId(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetCode(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetSuccess(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetResult(v int32) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Result = &v
	return s
}

type CreateMonitorGroupNotifyPolicyResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitorGroupNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitorGroupNotifyPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupNotifyPolicyResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponse) SetBody(v *CreateMonitorGroupNotifyPolicyResponseBody) *CreateMonitorGroupNotifyPolicyResponse {
	s.Body = v
	return s
}

type CreateMonitoringAgentProcessRequest struct {
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
}

func (s CreateMonitoringAgentProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitoringAgentProcessRequest) SetProcessName(v string) *CreateMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *CreateMonitoringAgentProcessRequest) SetInstanceId(v string) *CreateMonitoringAgentProcessRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMonitoringAgentProcessRequest) SetProcessUser(v string) *CreateMonitoringAgentProcessRequest {
	s.ProcessUser = &v
	return s
}

type CreateMonitoringAgentProcessResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMonitoringAgentProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitoringAgentProcessResponseBody) SetMessage(v string) *CreateMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetRequestId(v string) *CreateMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetId(v int64) *CreateMonitoringAgentProcessResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetCode(v string) *CreateMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetSuccess(v bool) *CreateMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

type CreateMonitoringAgentProcessResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitoringAgentProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *CreateMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitoringAgentProcessResponse) SetBody(v *CreateMonitoringAgentProcessResponseBody) *CreateMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

type CreateSiteMonitorRequest struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspCities   *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	AlertIds    *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
}

func (s CreateSiteMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorRequest) SetAddress(v string) *CreateSiteMonitorRequest {
	s.Address = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetTaskType(v string) *CreateSiteMonitorRequest {
	s.TaskType = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetTaskName(v string) *CreateSiteMonitorRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetInterval(v string) *CreateSiteMonitorRequest {
	s.Interval = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetIspCities(v string) *CreateSiteMonitorRequest {
	s.IspCities = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetOptionsJson(v string) *CreateSiteMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetAlertIds(v string) *CreateSiteMonitorRequest {
	s.AlertIds = &v
	return s
}

type CreateSiteMonitorResponseBody struct {
	Message          *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data             *CreateSiteMonitorResponseBodyData             `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code             *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateResultList *CreateSiteMonitorResponseBodyCreateResultList `json:"CreateResultList,omitempty" xml:"CreateResultList,omitempty" type:"Struct"`
	Success          *string                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSiteMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBody) SetMessage(v string) *CreateSiteMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetRequestId(v string) *CreateSiteMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetData(v *CreateSiteMonitorResponseBodyData) *CreateSiteMonitorResponseBody {
	s.Data = v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetCode(v string) *CreateSiteMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetCreateResultList(v *CreateSiteMonitorResponseBodyCreateResultList) *CreateSiteMonitorResponseBody {
	s.CreateResultList = v
	return s
}

func (s *CreateSiteMonitorResponseBody) SetSuccess(v string) *CreateSiteMonitorResponseBody {
	s.Success = &v
	return s
}

type CreateSiteMonitorResponseBodyData struct {
	AttachAlertResult *CreateSiteMonitorResponseBodyDataAttachAlertResult `json:"AttachAlertResult,omitempty" xml:"AttachAlertResult,omitempty" type:"Struct"`
}

func (s CreateSiteMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyData) SetAttachAlertResult(v *CreateSiteMonitorResponseBodyDataAttachAlertResult) *CreateSiteMonitorResponseBodyData {
	s.AttachAlertResult = v
	return s
}

type CreateSiteMonitorResponseBodyDataAttachAlertResult struct {
	Contact []*CreateSiteMonitorResponseBodyDataAttachAlertResultContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResult) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResult) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResult) SetContact(v []*CreateSiteMonitorResponseBodyDataAttachAlertResultContact) *CreateSiteMonitorResponseBodyDataAttachAlertResult {
	s.Contact = v
	return s
}

type CreateSiteMonitorResponseBodyDataAttachAlertResultContact struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResultContact) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyDataAttachAlertResultContact) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetRequestId(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.RequestId = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetSuccess(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.Success = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetCode(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.Code = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetMessage(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.Message = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyDataAttachAlertResultContact) SetRuleId(v string) *CreateSiteMonitorResponseBodyDataAttachAlertResultContact {
	s.RuleId = &v
	return s
}

type CreateSiteMonitorResponseBodyCreateResultList struct {
	CreateResultList []*CreateSiteMonitorResponseBodyCreateResultListCreateResultList `json:"CreateResultList,omitempty" xml:"CreateResultList,omitempty" type:"Repeated"`
}

func (s CreateSiteMonitorResponseBodyCreateResultList) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyCreateResultList) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyCreateResultList) SetCreateResultList(v []*CreateSiteMonitorResponseBodyCreateResultListCreateResultList) *CreateSiteMonitorResponseBodyCreateResultList {
	s.CreateResultList = v
	return s
}

type CreateSiteMonitorResponseBodyCreateResultListCreateResultList struct {
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSiteMonitorResponseBodyCreateResultListCreateResultList) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorResponseBodyCreateResultListCreateResultList) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponseBodyCreateResultListCreateResultList) SetTaskName(v string) *CreateSiteMonitorResponseBodyCreateResultListCreateResultList {
	s.TaskName = &v
	return s
}

func (s *CreateSiteMonitorResponseBodyCreateResultListCreateResultList) SetTaskId(v string) *CreateSiteMonitorResponseBodyCreateResultListCreateResultList {
	s.TaskId = &v
	return s
}

type CreateSiteMonitorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSiteMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSiteMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSiteMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorResponse) SetHeaders(v map[string]*string) *CreateSiteMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateSiteMonitorResponse) SetBody(v *CreateSiteMonitorResponseBody) *CreateSiteMonitorResponse {
	s.Body = v
	return s
}

type DeleteContactRequest struct {
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
}

func (s DeleteContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactRequest) SetContactName(v string) *DeleteContactRequest {
	s.ContactName = &v
	return s
}

type DeleteContactResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactResponseBody) SetMessage(v string) *DeleteContactResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactResponseBody) SetRequestId(v string) *DeleteContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactResponseBody) SetCode(v string) *DeleteContactResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactResponseBody) SetSuccess(v bool) *DeleteContactResponseBody {
	s.Success = &v
	return s
}

type DeleteContactResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactResponse) SetHeaders(v map[string]*string) *DeleteContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactResponse) SetBody(v *DeleteContactResponseBody) *DeleteContactResponse {
	s.Body = v
	return s
}

type DeleteContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
}

func (s DeleteContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupRequest) SetContactGroupName(v string) *DeleteContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

type DeleteContactGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupResponseBody) SetMessage(v string) *DeleteContactGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetRequestId(v string) *DeleteContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetCode(v string) *DeleteContactGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetSuccess(v bool) *DeleteContactGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteContactGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupResponse) SetHeaders(v map[string]*string) *DeleteContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactGroupResponse) SetBody(v *DeleteContactGroupResponseBody) *DeleteContactGroupResponse {
	s.Body = v
	return s
}

type DeleteCustomMetricRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Md5        *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	UUID       *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s DeleteCustomMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomMetricRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricRequest) SetGroupId(v string) *DeleteCustomMetricRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetMetricName(v string) *DeleteCustomMetricRequest {
	s.MetricName = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetMd5(v string) *DeleteCustomMetricRequest {
	s.Md5 = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetUUID(v string) *DeleteCustomMetricRequest {
	s.UUID = &v
	return s
}

type DeleteCustomMetricResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteCustomMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricResponseBody) SetMessage(v string) *DeleteCustomMetricResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) SetRequestId(v string) *DeleteCustomMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomMetricResponseBody) SetCode(v string) *DeleteCustomMetricResponseBody {
	s.Code = &v
	return s
}

type DeleteCustomMetricResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCustomMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCustomMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomMetricResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricResponse) SetHeaders(v map[string]*string) *DeleteCustomMetricResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomMetricResponse) SetBody(v *DeleteCustomMetricResponseBody) *DeleteCustomMetricResponse {
	s.Body = v
	return s
}

type DeleteDynamicTagGroupRequest struct {
	DynamicTagRuleId *string `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
}

func (s DeleteDynamicTagGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicTagGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDynamicTagGroupRequest) SetDynamicTagRuleId(v string) *DeleteDynamicTagGroupRequest {
	s.DynamicTagRuleId = &v
	return s
}

type DeleteDynamicTagGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDynamicTagGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDynamicTagGroupResponseBody) SetMessage(v string) *DeleteDynamicTagGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDynamicTagGroupResponseBody) SetRequestId(v string) *DeleteDynamicTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDynamicTagGroupResponseBody) SetCode(v string) *DeleteDynamicTagGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDynamicTagGroupResponseBody) SetSuccess(v bool) *DeleteDynamicTagGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteDynamicTagGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDynamicTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDynamicTagGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicTagGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDynamicTagGroupResponse) SetHeaders(v map[string]*string) *DeleteDynamicTagGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDynamicTagGroupResponse) SetBody(v *DeleteDynamicTagGroupResponseBody) *DeleteDynamicTagGroupResponse {
	s.Body = v
	return s
}

type DeleteEventRulesRequest struct {
	RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DeleteEventRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRulesRequest) SetRuleNames(v []*string) *DeleteEventRulesRequest {
	s.RuleNames = v
	return s
}

type DeleteEventRulesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRulesResponseBody) SetMessage(v string) *DeleteEventRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventRulesResponseBody) SetRequestId(v string) *DeleteEventRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRulesResponseBody) SetCode(v string) *DeleteEventRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRulesResponseBody) SetSuccess(v bool) *DeleteEventRulesResponseBody {
	s.Success = &v
	return s
}

type DeleteEventRulesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEventRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRulesResponse) SetHeaders(v map[string]*string) *DeleteEventRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRulesResponse) SetBody(v *DeleteEventRulesResponseBody) *DeleteEventRulesResponse {
	s.Body = v
	return s
}

type DeleteEventRuleTargetsRequest struct {
	RuleName *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Ids      []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DeleteEventRuleTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleTargetsRequest) SetRuleName(v string) *DeleteEventRuleTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteEventRuleTargetsRequest) SetIds(v []*string) *DeleteEventRuleTargetsRequest {
	s.Ids = v
	return s
}

type DeleteEventRuleTargetsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventRuleTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleTargetsResponseBody) SetMessage(v string) *DeleteEventRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventRuleTargetsResponseBody) SetRequestId(v string) *DeleteEventRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventRuleTargetsResponseBody) SetCode(v string) *DeleteEventRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventRuleTargetsResponseBody) SetSuccess(v bool) *DeleteEventRuleTargetsResponseBody {
	s.Success = &v
	return s
}

type DeleteEventRuleTargetsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEventRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEventRuleTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventRuleTargetsResponse) SetHeaders(v map[string]*string) *DeleteEventRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventRuleTargetsResponse) SetBody(v *DeleteEventRuleTargetsResponseBody) *DeleteEventRuleTargetsResponse {
	s.Body = v
	return s
}

type DeleteExporterOutputRequest struct {
	DestName *string `json:"DestName,omitempty" xml:"DestName,omitempty"`
}

func (s DeleteExporterOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExporterOutputRequest) GoString() string {
	return s.String()
}

func (s *DeleteExporterOutputRequest) SetDestName(v string) *DeleteExporterOutputRequest {
	s.DestName = &v
	return s
}

type DeleteExporterOutputResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExporterOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExporterOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExporterOutputResponseBody) SetMessage(v string) *DeleteExporterOutputResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExporterOutputResponseBody) SetRequestId(v string) *DeleteExporterOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExporterOutputResponseBody) SetCode(v string) *DeleteExporterOutputResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExporterOutputResponseBody) SetSuccess(v bool) *DeleteExporterOutputResponseBody {
	s.Success = &v
	return s
}

type DeleteExporterOutputResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteExporterOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteExporterOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExporterOutputResponse) GoString() string {
	return s.String()
}

func (s *DeleteExporterOutputResponse) SetHeaders(v map[string]*string) *DeleteExporterOutputResponse {
	s.Headers = v
	return s
}

func (s *DeleteExporterOutputResponse) SetBody(v *DeleteExporterOutputResponseBody) *DeleteExporterOutputResponse {
	s.Body = v
	return s
}

type DeleteExporterRuleRequest struct {
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteExporterRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExporterRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteExporterRuleRequest) SetRuleName(v string) *DeleteExporterRuleRequest {
	s.RuleName = &v
	return s
}

type DeleteExporterRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExporterRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExporterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExporterRuleResponseBody) SetMessage(v string) *DeleteExporterRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExporterRuleResponseBody) SetRequestId(v string) *DeleteExporterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExporterRuleResponseBody) SetCode(v string) *DeleteExporterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExporterRuleResponseBody) SetSuccess(v bool) *DeleteExporterRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteExporterRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteExporterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteExporterRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExporterRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteExporterRuleResponse) SetHeaders(v map[string]*string) *DeleteExporterRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteExporterRuleResponse) SetBody(v *DeleteExporterRuleResponseBody) *DeleteExporterRuleResponse {
	s.Body = v
	return s
}

type DeleteGroupMonitoringAgentProcessRequest struct {
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteGroupMonitoringAgentProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupMonitoringAgentProcessRequest) SetGroupId(v string) *DeleteGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessRequest) SetId(v string) *DeleteGroupMonitoringAgentProcessRequest {
	s.Id = &v
	return s
}

type DeleteGroupMonitoringAgentProcessResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGroupMonitoringAgentProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetCode(v string) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

type DeleteGroupMonitoringAgentProcessResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupMonitoringAgentProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *DeleteGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponse) SetBody(v *DeleteGroupMonitoringAgentProcessResponseBody) *DeleteGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

type DeleteHostAvailabilityRequest struct {
	Id []*int `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s DeleteHostAvailabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostAvailabilityRequest) SetId(v []*int) *DeleteHostAvailabilityRequest {
	s.Id = v
	return s
}

type DeleteHostAvailabilityResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHostAvailabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostAvailabilityResponseBody) SetMessage(v string) *DeleteHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHostAvailabilityResponseBody) SetRequestId(v string) *DeleteHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHostAvailabilityResponseBody) SetCode(v string) *DeleteHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHostAvailabilityResponseBody) SetSuccess(v bool) *DeleteHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

type DeleteHostAvailabilityResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHostAvailabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostAvailabilityResponse) SetHeaders(v map[string]*string) *DeleteHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostAvailabilityResponse) SetBody(v *DeleteHostAvailabilityResponseBody) *DeleteHostAvailabilityResponse {
	s.Body = v
	return s
}

type DeleteLogMonitorRequest struct {
	LogId *int64 `json:"LogId,omitempty" xml:"LogId,omitempty"`
}

func (s DeleteLogMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogMonitorRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogMonitorRequest) SetLogId(v int64) *DeleteLogMonitorRequest {
	s.LogId = &v
	return s
}

type DeleteLogMonitorResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLogMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogMonitorResponseBody) SetMessage(v string) *DeleteLogMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLogMonitorResponseBody) SetRequestId(v string) *DeleteLogMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogMonitorResponseBody) SetCode(v string) *DeleteLogMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLogMonitorResponseBody) SetSuccess(v bool) *DeleteLogMonitorResponseBody {
	s.Success = &v
	return s
}

type DeleteLogMonitorResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLogMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLogMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogMonitorResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogMonitorResponse) SetHeaders(v map[string]*string) *DeleteLogMonitorResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogMonitorResponse) SetBody(v *DeleteLogMonitorResponseBody) *DeleteLogMonitorResponse {
	s.Body = v
	return s
}

type DeleteMetricRuleResourcesRequest struct {
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DeleteMetricRuleResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleResourcesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleResourcesRequest) SetRuleId(v string) *DeleteMetricRuleResourcesRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMetricRuleResourcesRequest) SetResources(v string) *DeleteMetricRuleResourcesRequest {
	s.Resources = &v
	return s
}

type DeleteMetricRuleResourcesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRuleResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleResourcesResponseBody) SetMessage(v string) *DeleteMetricRuleResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponseBody) SetRequestId(v string) *DeleteMetricRuleResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponseBody) SetCode(v string) *DeleteMetricRuleResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleResourcesResponseBody) SetSuccess(v bool) *DeleteMetricRuleResourcesResponseBody {
	s.Success = &v
	return s
}

type DeleteMetricRuleResourcesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMetricRuleResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMetricRuleResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleResourcesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleResourcesResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleResourcesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleResourcesResponse) SetBody(v *DeleteMetricRuleResourcesResponseBody) *DeleteMetricRuleResourcesResponse {
	s.Body = v
	return s
}

type DeleteMetricRulesRequest struct {
	Id []*string `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s DeleteMetricRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesRequest) SetId(v []*string) *DeleteMetricRulesRequest {
	s.Id = v
	return s
}

type DeleteMetricRulesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesResponseBody) SetMessage(v string) *DeleteMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetRequestId(v string) *DeleteMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetCode(v string) *DeleteMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetSuccess(v bool) *DeleteMetricRulesResponseBody {
	s.Success = &v
	return s
}

type DeleteMetricRulesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMetricRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesResponse) SetHeaders(v map[string]*string) *DeleteMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRulesResponse) SetBody(v *DeleteMetricRulesResponseBody) *DeleteMetricRulesResponse {
	s.Body = v
	return s
}

type DeleteMetricRuleTargetsRequest struct {
	RuleId    *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
}

func (s DeleteMetricRuleTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsRequest) SetRuleId(v string) *DeleteMetricRuleTargetsRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMetricRuleTargetsRequest) SetTargetIds(v []*string) *DeleteMetricRuleTargetsRequest {
	s.TargetIds = v
	return s
}

type DeleteMetricRuleTargetsResponseBody struct {
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	FailIds   *DeleteMetricRuleTargetsResponseBodyFailIds `json:"FailIds,omitempty" xml:"FailIds,omitempty" type:"Struct"`
}

func (s DeleteMetricRuleTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBody) SetMessage(v string) *DeleteMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetRequestId(v string) *DeleteMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetCode(v string) *DeleteMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetSuccess(v bool) *DeleteMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetricRuleTargetsResponseBody) SetFailIds(v *DeleteMetricRuleTargetsResponseBodyFailIds) *DeleteMetricRuleTargetsResponseBody {
	s.FailIds = v
	return s
}

type DeleteMetricRuleTargetsResponseBodyFailIds struct {
	TargetIds *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Struct"`
}

func (s DeleteMetricRuleTargetsResponseBodyFailIds) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBodyFailIds) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIds) SetTargetIds(v *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) *DeleteMetricRuleTargetsResponseBodyFailIds {
	s.TargetIds = v
	return s
}

type DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds struct {
	TargetId []*string `json:"TargetId,omitempty" xml:"TargetId,omitempty" type:"Repeated"`
}

func (s DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds) SetTargetId(v []*string) *DeleteMetricRuleTargetsResponseBodyFailIdsTargetIds {
	s.TargetId = v
	return s
}

type DeleteMetricRuleTargetsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMetricRuleTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleTargetsResponse) SetBody(v *DeleteMetricRuleTargetsResponseBody) *DeleteMetricRuleTargetsResponse {
	s.Body = v
	return s
}

type DeleteMetricRuleTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMetricRuleTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateRequest) SetTemplateId(v string) *DeleteMetricRuleTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteMetricRuleTemplateResponseBody struct {
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource  *DeleteMetricRuleTemplateResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	Code      *int32                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRuleTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateResponseBody) SetMessage(v string) *DeleteMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetRequestId(v string) *DeleteMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetResource(v *DeleteMetricRuleTemplateResponseBodyResource) *DeleteMetricRuleTemplateResponseBody {
	s.Resource = v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetCode(v int32) *DeleteMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleTemplateResponseBody) SetSuccess(v bool) *DeleteMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

type DeleteMetricRuleTemplateResponseBodyResource struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMetricRuleTemplateResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTemplateResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateResponseBodyResource) SetTemplateId(v string) *DeleteMetricRuleTemplateResponseBodyResource {
	s.TemplateId = &v
	return s
}

type DeleteMetricRuleTemplateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMetricRuleTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *DeleteMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteMetricRuleTemplateResponse) SetBody(v *DeleteMetricRuleTemplateResponseBody) *DeleteMetricRuleTemplateResponse {
	s.Body = v
	return s
}

type DeleteMonitorGroupRequest struct {
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteMonitorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupRequest) SetGroupId(v int64) *DeleteMonitorGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteMonitorGroupResponseBody struct {
	Group     *DeleteMonitorGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMonitorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBody) SetGroup(v *DeleteMonitorGroupResponseBodyGroup) *DeleteMonitorGroupResponseBody {
	s.Group = v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetMessage(v string) *DeleteMonitorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetRequestId(v string) *DeleteMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetCode(v int32) *DeleteMonitorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupResponseBody) SetSuccess(v bool) *DeleteMonitorGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteMonitorGroupResponseBodyGroup struct {
	GroupName     *string                                           `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ContactGroups *DeleteMonitorGroupResponseBodyGroupContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
}

func (s DeleteMonitorGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBodyGroup) SetGroupName(v string) *DeleteMonitorGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *DeleteMonitorGroupResponseBodyGroup) SetContactGroups(v *DeleteMonitorGroupResponseBodyGroupContactGroups) *DeleteMonitorGroupResponseBodyGroup {
	s.ContactGroups = v
	return s
}

type DeleteMonitorGroupResponseBodyGroupContactGroups struct {
	ContactGroup []*DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroups) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroups) SetContactGroup(v []*DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) *DeleteMonitorGroupResponseBodyGroupContactGroups {
	s.ContactGroup = v
	return s
}

type DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup) SetName(v string) *DeleteMonitorGroupResponseBodyGroupContactGroupsContactGroup {
	s.Name = &v
	return s
}

type DeleteMonitorGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMonitorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMonitorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupResponse) SetBody(v *DeleteMonitorGroupResponseBody) *DeleteMonitorGroupResponse {
	s.Body = v
	return s
}

type DeleteMonitorGroupDynamicRuleRequest struct {
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DeleteMonitorGroupDynamicRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupDynamicRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupDynamicRuleRequest) SetGroupId(v int64) *DeleteMonitorGroupDynamicRuleRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleRequest) SetCategory(v string) *DeleteMonitorGroupDynamicRuleRequest {
	s.Category = &v
	return s
}

type DeleteMonitorGroupDynamicRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMonitorGroupDynamicRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupDynamicRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetMessage(v string) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetRequestId(v string) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetCode(v int32) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetSuccess(v bool) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteMonitorGroupDynamicRuleResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMonitorGroupDynamicRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMonitorGroupDynamicRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupDynamicRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupDynamicRuleResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupDynamicRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponse) SetBody(v *DeleteMonitorGroupDynamicRuleResponseBody) *DeleteMonitorGroupDynamicRuleResponse {
	s.Body = v
	return s
}

type DeleteMonitorGroupInstancesRequest struct {
	GroupId        *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	Category       *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DeleteMonitorGroupInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupInstancesRequest) SetGroupId(v int64) *DeleteMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMonitorGroupInstancesRequest) SetInstanceIdList(v string) *DeleteMonitorGroupInstancesRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DeleteMonitorGroupInstancesRequest) SetCategory(v string) *DeleteMonitorGroupInstancesRequest {
	s.Category = &v
	return s
}

type DeleteMonitorGroupInstancesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMonitorGroupInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetMessage(v string) *DeleteMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetRequestId(v string) *DeleteMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetCode(v int32) *DeleteMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupInstancesResponseBody) SetSuccess(v bool) *DeleteMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

type DeleteMonitorGroupInstancesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMonitorGroupInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupInstancesResponse) SetBody(v *DeleteMonitorGroupInstancesResponseBody) *DeleteMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

type DeleteMonitorGroupNotifyPolicyRequest struct {
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteMonitorGroupNotifyPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) SetPolicyType(v string) *DeleteMonitorGroupNotifyPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) SetGroupId(v string) *DeleteMonitorGroupNotifyPolicyRequest {
	s.GroupId = &v
	return s
}

type DeleteMonitorGroupNotifyPolicyResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *int32  `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteMonitorGroupNotifyPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetMessage(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetRequestId(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetCode(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetSuccess(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetResult(v int32) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Result = &v
	return s
}

type DeleteMonitorGroupNotifyPolicyResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMonitorGroupNotifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMonitorGroupNotifyPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitorGroupNotifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) SetHeaders(v map[string]*string) *DeleteMonitorGroupNotifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponse) SetBody(v *DeleteMonitorGroupNotifyPolicyResponseBody) *DeleteMonitorGroupNotifyPolicyResponse {
	s.Body = v
	return s
}

type DeleteMonitoringAgentProcessRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessId   *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s DeleteMonitoringAgentProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitoringAgentProcessRequest) SetInstanceId(v string) *DeleteMonitoringAgentProcessRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMonitoringAgentProcessRequest) SetProcessName(v string) *DeleteMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *DeleteMonitoringAgentProcessRequest) SetProcessId(v string) *DeleteMonitoringAgentProcessRequest {
	s.ProcessId = &v
	return s
}

type DeleteMonitoringAgentProcessResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMonitoringAgentProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetMessage(v string) *DeleteMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetRequestId(v string) *DeleteMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetCode(v string) *DeleteMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetSuccess(v bool) *DeleteMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

type DeleteMonitoringAgentProcessResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMonitoringAgentProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *DeleteMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteMonitoringAgentProcessResponse) SetBody(v *DeleteMonitoringAgentProcessResponseBody) *DeleteMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

type DeleteSiteMonitorsRequest struct {
	TaskIds        *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
	IsDeleteAlarms *bool   `json:"IsDeleteAlarms,omitempty" xml:"IsDeleteAlarms,omitempty"`
}

func (s DeleteSiteMonitorsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteMonitorsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsRequest) SetTaskIds(v string) *DeleteSiteMonitorsRequest {
	s.TaskIds = &v
	return s
}

func (s *DeleteSiteMonitorsRequest) SetIsDeleteAlarms(v bool) *DeleteSiteMonitorsRequest {
	s.IsDeleteAlarms = &v
	return s
}

type DeleteSiteMonitorsResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DeleteSiteMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSiteMonitorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsResponseBody) SetMessage(v string) *DeleteSiteMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetRequestId(v string) *DeleteSiteMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetData(v *DeleteSiteMonitorsResponseBodyData) *DeleteSiteMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetCode(v string) *DeleteSiteMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSiteMonitorsResponseBody) SetSuccess(v string) *DeleteSiteMonitorsResponseBody {
	s.Success = &v
	return s
}

type DeleteSiteMonitorsResponseBodyData struct {
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s DeleteSiteMonitorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsResponseBodyData) SetCount(v int32) *DeleteSiteMonitorsResponseBodyData {
	s.Count = &v
	return s
}

type DeleteSiteMonitorsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSiteMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSiteMonitorsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSiteMonitorsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSiteMonitorsResponse) SetHeaders(v map[string]*string) *DeleteSiteMonitorsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSiteMonitorsResponse) SetBody(v *DeleteSiteMonitorsResponseBody) *DeleteSiteMonitorsResponse {
	s.Body = v
	return s
}

type DescribeActiveMetricRuleListRequest struct {
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DescribeActiveMetricRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListRequest) SetProduct(v string) *DescribeActiveMetricRuleListRequest {
	s.Product = &v
	return s
}

type DescribeActiveMetricRuleListResponseBody struct {
	Message    *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AlertList  *DescribeActiveMetricRuleListResponseBodyAlertList  `json:"AlertList,omitempty" xml:"AlertList,omitempty" type:"Struct"`
	Datapoints *DescribeActiveMetricRuleListResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	Code       *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBody) SetMessage(v string) *DescribeActiveMetricRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetRequestId(v string) *DescribeActiveMetricRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetAlertList(v *DescribeActiveMetricRuleListResponseBodyAlertList) *DescribeActiveMetricRuleListResponseBody {
	s.AlertList = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetDatapoints(v *DescribeActiveMetricRuleListResponseBodyDatapoints) *DescribeActiveMetricRuleListResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetCode(v string) *DescribeActiveMetricRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBody) SetSuccess(v bool) *DescribeActiveMetricRuleListResponseBody {
	s.Success = &v
	return s
}

type DescribeActiveMetricRuleListResponseBodyAlertList struct {
	Alert []*DescribeActiveMetricRuleListResponseBodyAlertListAlert `json:"Alert,omitempty" xml:"Alert,omitempty" type:"Repeated"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertList) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertList) SetAlert(v []*DescribeActiveMetricRuleListResponseBodyAlertListAlert) *DescribeActiveMetricRuleListResponseBodyAlertList {
	s.Alert = v
	return s
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlert struct {
	SilenceTime         *string                                                            `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	MetricName          *string                                                            `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Webhook             *string                                                            `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Escalations         *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	ContactGroups       *string                                                            `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Namespace           *string                                                            `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MailSubject         *string                                                            `json:"MailSubject,omitempty" xml:"MailSubject,omitempty"`
	NoEffectiveInterval *string                                                            `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	EffectiveInterval   *string                                                            `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	RuleName            *string                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	AlertState          *string                                                            `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	Period              *string                                                            `json:"Period,omitempty" xml:"Period,omitempty"`
	RuleId              *string                                                            `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Dimensions          *string                                                            `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EnableState         *bool                                                              `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	Resources           *string                                                            `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlert) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlert) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetSilenceTime(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.SilenceTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetMetricName(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.MetricName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetWebhook(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Webhook = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetEscalations(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Escalations = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetContactGroups(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.ContactGroups = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetNamespace(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Namespace = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetMailSubject(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.MailSubject = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetNoEffectiveInterval(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetEffectiveInterval(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetRuleName(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.RuleName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetAlertState(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.AlertState = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetPeriod(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Period = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetRuleId(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.RuleId = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetDimensions(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Dimensions = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetEnableState(v bool) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.EnableState = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlert) SetResources(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlert {
	s.Resources = &v
	return s
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations struct {
	Critical *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) SetCritical(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations {
	s.Critical = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) SetInfo(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations {
	s.Info = v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations) SetWarn(v *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalations {
	s.Warn = v
	return s
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetTimes(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.Times = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsCritical {
	s.Statistics = &v
	return s
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetTimes(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.Times = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsInfo {
	s.Statistics = &v
	return s
}

type DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *string `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetTimes(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.Times = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyAlertListAlertEscalationsWarn {
	s.Statistics = &v
	return s
}

type DescribeActiveMetricRuleListResponseBodyDatapoints struct {
	Alarm []*DescribeActiveMetricRuleListResponseBodyDatapointsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeActiveMetricRuleListResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapoints) SetAlarm(v []*DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) *DescribeActiveMetricRuleListResponseBodyDatapoints {
	s.Alarm = v
	return s
}

type DescribeActiveMetricRuleListResponseBodyDatapointsAlarm struct {
	SilenceTime        *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	EvaluationCount    *string `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Webhook            *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	ContactGroups      *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RuleId             *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Enable             *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetSilenceTime(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetMetricName(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetEvaluationCount(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetWebhook(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Webhook = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetState(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.State = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetContactGroups(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetNamespace(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetPeriod(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetRuleId(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.RuleId = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetRuleName(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.RuleName = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetEndTime(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.EndTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetComparisonOperator(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetStartTime(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetThreshold(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Threshold = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetStatistics(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Statistics = &v
	return s
}

func (s *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm) SetEnable(v string) *DescribeActiveMetricRuleListResponseBodyDatapointsAlarm {
	s.Enable = &v
	return s
}

type DescribeActiveMetricRuleListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeActiveMetricRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeActiveMetricRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeActiveMetricRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListResponse) SetHeaders(v map[string]*string) *DescribeActiveMetricRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeActiveMetricRuleListResponse) SetBody(v *DescribeActiveMetricRuleListResponseBody) *DescribeActiveMetricRuleListResponse {
	s.Body = v
	return s
}

type DescribeAlertHistoryListRequest struct {
	RuleId     *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	Ascending  *bool   `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Page       *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s DescribeAlertHistoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListRequest) SetRuleId(v string) *DescribeAlertHistoryListRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetRuleName(v string) *DescribeAlertHistoryListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetNamespace(v string) *DescribeAlertHistoryListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetMetricName(v string) *DescribeAlertHistoryListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetGroupId(v string) *DescribeAlertHistoryListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetStatus(v string) *DescribeAlertHistoryListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetState(v string) *DescribeAlertHistoryListRequest {
	s.State = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetAscending(v bool) *DescribeAlertHistoryListRequest {
	s.Ascending = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetStartTime(v string) *DescribeAlertHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetEndTime(v string) *DescribeAlertHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetPageSize(v int32) *DescribeAlertHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertHistoryListRequest) SetPage(v int32) *DescribeAlertHistoryListRequest {
	s.Page = &v
	return s
}

type DescribeAlertHistoryListResponseBody struct {
	AlarmHistoryList *DescribeAlertHistoryListResponseBodyAlarmHistoryList `json:"AlarmHistoryList,omitempty" xml:"AlarmHistoryList,omitempty" type:"Struct"`
	Message          *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total            *string                                               `json:"Total,omitempty" xml:"Total,omitempty"`
	Code             *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success          *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertHistoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBody) SetAlarmHistoryList(v *DescribeAlertHistoryListResponseBodyAlarmHistoryList) *DescribeAlertHistoryListResponseBody {
	s.AlarmHistoryList = v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetMessage(v string) *DescribeAlertHistoryListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetRequestId(v string) *DescribeAlertHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetTotal(v string) *DescribeAlertHistoryListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetCode(v string) *DescribeAlertHistoryListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBody) SetSuccess(v bool) *DescribeAlertHistoryListResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryList struct {
	AlarmHistory []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory `json:"AlarmHistory,omitempty" xml:"AlarmHistory,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryList) SetAlarmHistory(v []*DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) *DescribeAlertHistoryListResponseBodyAlarmHistoryList {
	s.AlarmHistory = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory struct {
	Status          *int32                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	MetricName      *string                                                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Contacts        *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts      `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	EvaluationCount *int32                                                                         `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	State           *string                                                                        `json:"State,omitempty" xml:"State,omitempty"`
	ContactGroups   *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	Namespace       *string                                                                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Webhooks        *string                                                                        `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
	RuleId          *string                                                                        `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName        *string                                                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	LastTime        *int64                                                                         `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	Value           *string                                                                        `json:"Value,omitempty" xml:"Value,omitempty"`
	Expression      *string                                                                        `json:"Expression,omitempty" xml:"Expression,omitempty"`
	GroupId         *string                                                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AlertTime       *int64                                                                         `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	InstanceName    *string                                                                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ContactSmses    *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses  `json:"ContactSmses,omitempty" xml:"ContactSmses,omitempty" type:"Struct"`
	Dimensions      *string                                                                        `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	ContactALIIMs   *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs `json:"ContactALIIMs,omitempty" xml:"ContactALIIMs,omitempty" type:"Struct"`
	Level           *string                                                                        `json:"Level,omitempty" xml:"Level,omitempty"`
	ContactMails    *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails  `json:"ContactMails,omitempty" xml:"ContactMails,omitempty" type:"Struct"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetStatus(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Status = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetMetricName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContacts(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Contacts = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetEvaluationCount(v int32) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetState(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.State = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactGroups(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactGroups = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetNamespace(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetWebhooks(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Webhooks = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetRuleId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetRuleName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLastTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.LastTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetValue(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Value = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetExpression(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Expression = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetGroupId(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetAlertTime(v int64) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.AlertTime = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetInstanceName(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactSmses(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactSmses = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetDimensions(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Dimensions = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactALIIMs(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactALIIMs = v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetLevel(v string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.Level = &v
	return s
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory) SetContactMails(v *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistory {
	s.ContactMails = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts) SetContact(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContacts {
	s.Contact = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups) SetContactGroup(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactGroups {
	s.ContactGroup = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses struct {
	ContactSms []*string `json:"ContactSms,omitempty" xml:"ContactSms,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses) SetContactSms(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactSmses {
	s.ContactSms = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs struct {
	ContactALIIM []*string `json:"ContactALIIM,omitempty" xml:"ContactALIIM,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs) SetContactALIIM(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactALIIMs {
	s.ContactALIIM = v
	return s
}

type DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails struct {
	ContactMail []*string `json:"ContactMail,omitempty" xml:"ContactMail,omitempty" type:"Repeated"`
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails) SetContactMail(v []*string) *DescribeAlertHistoryListResponseBodyAlarmHistoryListAlarmHistoryContactMails {
	s.ContactMail = v
	return s
}

type DescribeAlertHistoryListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlertHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertHistoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertHistoryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponse) SetHeaders(v map[string]*string) *DescribeAlertHistoryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertHistoryListResponse) SetBody(v *DescribeAlertHistoryListResponseBody) *DescribeAlertHistoryListResponse {
	s.Body = v
	return s
}

type DescribeAlertLogCountRequest struct {
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey    *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	SendStatus   *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	MetricName   *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	LastMin      *string `json:"LastMin,omitempty" xml:"LastMin,omitempty"`
	GroupBy      *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
}

func (s DescribeAlertLogCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountRequest) SetStartTime(v int64) *DescribeAlertLogCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetEndTime(v int64) *DescribeAlertLogCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetPageNumber(v int32) *DescribeAlertLogCountRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetPageSize(v int32) *DescribeAlertLogCountRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetSearchKey(v string) *DescribeAlertLogCountRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetNamespace(v string) *DescribeAlertLogCountRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetGroupId(v string) *DescribeAlertLogCountRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetProduct(v string) *DescribeAlertLogCountRequest {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetLevel(v string) *DescribeAlertLogCountRequest {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetSendStatus(v string) *DescribeAlertLogCountRequest {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetContactGroup(v string) *DescribeAlertLogCountRequest {
	s.ContactGroup = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetRuleName(v string) *DescribeAlertLogCountRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetMetricName(v string) *DescribeAlertLogCountRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetLastMin(v string) *DescribeAlertLogCountRequest {
	s.LastMin = &v
	return s
}

func (s *DescribeAlertLogCountRequest) SetGroupBy(v string) *DescribeAlertLogCountRequest {
	s.GroupBy = &v
	return s
}

type DescribeAlertLogCountResponseBody struct {
	Message       *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AlertLogCount []*DescribeAlertLogCountResponseBodyAlertLogCount `json:"AlertLogCount,omitempty" xml:"AlertLogCount,omitempty" type:"Repeated"`
	Code          *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success       *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertLogCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponseBody) SetMessage(v string) *DescribeAlertLogCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetRequestId(v string) *DescribeAlertLogCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetAlertLogCount(v []*DescribeAlertLogCountResponseBodyAlertLogCount) *DescribeAlertLogCountResponseBody {
	s.AlertLogCount = v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetCode(v string) *DescribeAlertLogCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogCountResponseBody) SetSuccess(v bool) *DescribeAlertLogCountResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertLogCountResponseBodyAlertLogCount struct {
	Logs  []*DescribeAlertLogCountResponseBodyAlertLogCountLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	Count *int32                                                `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAlertLogCountResponseBodyAlertLogCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogCountResponseBodyAlertLogCount) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCount) SetLogs(v []*DescribeAlertLogCountResponseBodyAlertLogCountLogs) *DescribeAlertLogCountResponseBodyAlertLogCount {
	s.Logs = v
	return s
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCount) SetCount(v int32) *DescribeAlertLogCountResponseBodyAlertLogCount {
	s.Count = &v
	return s
}

type DescribeAlertLogCountResponseBodyAlertLogCountLogs struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeAlertLogCountResponseBodyAlertLogCountLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogCountResponseBodyAlertLogCountLogs) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCountLogs) SetValue(v string) *DescribeAlertLogCountResponseBodyAlertLogCountLogs {
	s.Value = &v
	return s
}

func (s *DescribeAlertLogCountResponseBodyAlertLogCountLogs) SetName(v string) *DescribeAlertLogCountResponseBodyAlertLogCountLogs {
	s.Name = &v
	return s
}

type DescribeAlertLogCountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlertLogCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertLogCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogCountResponse) SetHeaders(v map[string]*string) *DescribeAlertLogCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertLogCountResponse) SetBody(v *DescribeAlertLogCountResponseBody) *DescribeAlertLogCountResponse {
	s.Body = v
	return s
}

type DescribeAlertLogHistogramRequest struct {
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey    *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	SendStatus   *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	MetricName   *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	LastMin      *string `json:"LastMin,omitempty" xml:"LastMin,omitempty"`
	GroupBy      *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
}

func (s DescribeAlertLogHistogramRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramRequest) SetStartTime(v int64) *DescribeAlertLogHistogramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetEndTime(v int64) *DescribeAlertLogHistogramRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetPageNumber(v int32) *DescribeAlertLogHistogramRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetPageSize(v int32) *DescribeAlertLogHistogramRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetSearchKey(v string) *DescribeAlertLogHistogramRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetGroupId(v string) *DescribeAlertLogHistogramRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetProduct(v string) *DescribeAlertLogHistogramRequest {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetNamespace(v string) *DescribeAlertLogHistogramRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetLevel(v string) *DescribeAlertLogHistogramRequest {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetSendStatus(v string) *DescribeAlertLogHistogramRequest {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetContactGroup(v string) *DescribeAlertLogHistogramRequest {
	s.ContactGroup = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetRuleName(v string) *DescribeAlertLogHistogramRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetMetricName(v string) *DescribeAlertLogHistogramRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetLastMin(v string) *DescribeAlertLogHistogramRequest {
	s.LastMin = &v
	return s
}

func (s *DescribeAlertLogHistogramRequest) SetGroupBy(v string) *DescribeAlertLogHistogramRequest {
	s.GroupBy = &v
	return s
}

type DescribeAlertLogHistogramResponseBody struct {
	Message               *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId             *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AlertLogHistogramList []*DescribeAlertLogHistogramResponseBodyAlertLogHistogramList `json:"AlertLogHistogramList,omitempty" xml:"AlertLogHistogramList,omitempty" type:"Repeated"`
	Code                  *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success               *bool                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertLogHistogramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramResponseBody) SetMessage(v string) *DescribeAlertLogHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetRequestId(v string) *DescribeAlertLogHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetAlertLogHistogramList(v []*DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) *DescribeAlertLogHistogramResponseBody {
	s.AlertLogHistogramList = v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetCode(v string) *DescribeAlertLogHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBody) SetSuccess(v bool) *DescribeAlertLogHistogramResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertLogHistogramResponseBodyAlertLogHistogramList struct {
	From  *int64 `json:"From,omitempty" xml:"From,omitempty"`
	To    *int64 `json:"To,omitempty" xml:"To,omitempty"`
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) SetFrom(v int64) *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList {
	s.From = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) SetTo(v int64) *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList {
	s.To = &v
	return s
}

func (s *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList) SetCount(v int32) *DescribeAlertLogHistogramResponseBodyAlertLogHistogramList {
	s.Count = &v
	return s
}

type DescribeAlertLogHistogramResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlertLogHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertLogHistogramResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogHistogramResponse) SetHeaders(v map[string]*string) *DescribeAlertLogHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertLogHistogramResponse) SetBody(v *DescribeAlertLogHistogramResponseBody) *DescribeAlertLogHistogramResponse {
	s.Body = v
	return s
}

type DescribeAlertLogListRequest struct {
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey    *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Namespace    *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	SendStatus   *string `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	MetricName   *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	LastMin      *string `json:"LastMin,omitempty" xml:"LastMin,omitempty"`
	GroupBy      *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
}

func (s DescribeAlertLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListRequest) SetStartTime(v int64) *DescribeAlertLogListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetEndTime(v int64) *DescribeAlertLogListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetPageNumber(v int32) *DescribeAlertLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetPageSize(v int32) *DescribeAlertLogListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetSearchKey(v string) *DescribeAlertLogListRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetGroupId(v string) *DescribeAlertLogListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetNamespace(v string) *DescribeAlertLogListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetProduct(v string) *DescribeAlertLogListRequest {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetLevel(v string) *DescribeAlertLogListRequest {
	s.Level = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetSendStatus(v string) *DescribeAlertLogListRequest {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetContactGroup(v string) *DescribeAlertLogListRequest {
	s.ContactGroup = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetRuleName(v string) *DescribeAlertLogListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetMetricName(v string) *DescribeAlertLogListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetLastMin(v string) *DescribeAlertLogListRequest {
	s.LastMin = &v
	return s
}

func (s *DescribeAlertLogListRequest) SetGroupBy(v string) *DescribeAlertLogListRequest {
	s.GroupBy = &v
	return s
}

type DescribeAlertLogListResponseBody struct {
	AlertLogList []*DescribeAlertLogListResponseBodyAlertLogList `json:"AlertLogList,omitempty" xml:"AlertLogList,omitempty" type:"Repeated"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message      *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize     *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total        *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	Code         *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBody) SetAlertLogList(v []*DescribeAlertLogListResponseBodyAlertLogList) *DescribeAlertLogListResponseBody {
	s.AlertLogList = v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetRequestId(v string) *DescribeAlertLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetMessage(v string) *DescribeAlertLogListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetPageSize(v int32) *DescribeAlertLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetPageNumber(v int32) *DescribeAlertLogListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetTotal(v int32) *DescribeAlertLogListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetCode(v string) *DescribeAlertLogListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertLogListResponseBody) SetSuccess(v bool) *DescribeAlertLogListResponseBody {
	s.Success = &v
	return s
}

type DescribeAlertLogListResponseBodyAlertLogList struct {
	MetricName          *string                                                     `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	EventName           *string                                                     `json:"EventName,omitempty" xml:"EventName,omitempty"`
	ContactALIIWWList   []*string                                                   `json:"ContactALIIWWList,omitempty" xml:"ContactALIIWWList,omitempty" type:"Repeated"`
	Message             *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	LevelChange         *string                                                     `json:"LevelChange,omitempty" xml:"LevelChange,omitempty"`
	RuleId              *string                                                     `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ExtendedInfo        []*DescribeAlertLogListResponseBodyAlertLogListExtendedInfo `json:"ExtendedInfo,omitempty" xml:"ExtendedInfo,omitempty" type:"Repeated"`
	DingdingWebhookList []*string                                                   `json:"DingdingWebhookList,omitempty" xml:"DingdingWebhookList,omitempty" type:"Repeated"`
	InstanceName        *string                                                     `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ContactMailList     []*string                                                   `json:"ContactMailList,omitempty" xml:"ContactMailList,omitempty" type:"Repeated"`
	Dimensions          []*DescribeAlertLogListResponseBodyAlertLogListDimensions   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	ContactSMSList      []*string                                                   `json:"ContactSMSList,omitempty" xml:"ContactSMSList,omitempty" type:"Repeated"`
	SendStatus          *string                                                     `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	ContactOnCallList   []*string                                                   `json:"ContactOnCallList,omitempty" xml:"ContactOnCallList,omitempty" type:"Repeated"`
	Product             *string                                                     `json:"Product,omitempty" xml:"Product,omitempty"`
	ContactGroups       []*string                                                   `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	Namespace           *string                                                     `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Escalation          *DescribeAlertLogListResponseBodyAlertLogListEscalation     `json:"Escalation,omitempty" xml:"Escalation,omitempty" type:"Struct"`
	InstanceId          *string                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ContactDingList     []*string                                                   `json:"ContactDingList,omitempty" xml:"ContactDingList,omitempty" type:"Repeated"`
	RuleName            *string                                                     `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	WebhookList         []*string                                                   `json:"WebhookList,omitempty" xml:"WebhookList,omitempty" type:"Repeated"`
	GroupId             *string                                                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string                                                     `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	AlertTime           *string                                                     `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	Level               *string                                                     `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogList) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetMetricName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.MetricName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetEventName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.EventName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactALIIWWList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactALIIWWList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetMessage(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Message = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetLevelChange(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.LevelChange = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetRuleId(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.RuleId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetExtendedInfo(v []*DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ExtendedInfo = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetDingdingWebhookList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.DingdingWebhookList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetInstanceName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactMailList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactMailList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetDimensions(v []*DescribeAlertLogListResponseBodyAlertLogListDimensions) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Dimensions = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactSMSList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactSMSList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetSendStatus(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.SendStatus = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactOnCallList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactOnCallList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetProduct(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Product = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactGroups(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactGroups = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetNamespace(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Namespace = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetEscalation(v *DescribeAlertLogListResponseBodyAlertLogListEscalation) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Escalation = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetInstanceId(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetContactDingList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.ContactDingList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetRuleName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.RuleName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetWebhookList(v []*string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.WebhookList = v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetGroupId(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.GroupId = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetGroupName(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.GroupName = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetAlertTime(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.AlertTime = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogList) SetLevel(v string) *DescribeAlertLogListResponseBodyAlertLogList {
	s.Level = &v
	return s
}

type DescribeAlertLogListResponseBodyAlertLogListExtendedInfo struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) SetValue(v string) *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo {
	s.Value = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo) SetName(v string) *DescribeAlertLogListResponseBodyAlertLogListExtendedInfo {
	s.Name = &v
	return s
}

type DescribeAlertLogListResponseBodyAlertLogListDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListDimensions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListDimensions) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListDimensions) SetKey(v string) *DescribeAlertLogListResponseBodyAlertLogListDimensions {
	s.Key = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListDimensions) SetValue(v string) *DescribeAlertLogListResponseBodyAlertLogListDimensions {
	s.Value = &v
	return s
}

type DescribeAlertLogListResponseBodyAlertLogListEscalation struct {
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Times      *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeAlertLogListResponseBodyAlertLogListEscalation) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogListResponseBodyAlertLogListEscalation) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) SetExpression(v string) *DescribeAlertLogListResponseBodyAlertLogListEscalation {
	s.Expression = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) SetTimes(v int32) *DescribeAlertLogListResponseBodyAlertLogListEscalation {
	s.Times = &v
	return s
}

func (s *DescribeAlertLogListResponseBodyAlertLogListEscalation) SetLevel(v string) *DescribeAlertLogListResponseBodyAlertLogListEscalation {
	s.Level = &v
	return s
}

type DescribeAlertLogListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlertLogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlertLogListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlertLogListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertLogListResponse) SetHeaders(v map[string]*string) *DescribeAlertLogListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertLogListResponse) SetBody(v *DescribeAlertLogListResponseBody) *DescribeAlertLogListResponse {
	s.Body = v
	return s
}

type DescribeContactGroupListRequest struct {
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeContactGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListRequest) SetPageSize(v int32) *DescribeContactGroupListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContactGroupListRequest) SetPageNumber(v int32) *DescribeContactGroupListRequest {
	s.PageNumber = &v
	return s
}

type DescribeContactGroupListResponseBody struct {
	ContactGroupList *DescribeContactGroupListResponseBodyContactGroupList `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Struct"`
	ContactGroups    *DescribeContactGroupListResponseBodyContactGroups    `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	Message          *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total            *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Code             *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success          *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeContactGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBody) SetContactGroupList(v *DescribeContactGroupListResponseBodyContactGroupList) *DescribeContactGroupListResponseBody {
	s.ContactGroupList = v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetContactGroups(v *DescribeContactGroupListResponseBodyContactGroups) *DescribeContactGroupListResponseBody {
	s.ContactGroups = v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetMessage(v string) *DescribeContactGroupListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetRequestId(v string) *DescribeContactGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetTotal(v int32) *DescribeContactGroupListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetCode(v string) *DescribeContactGroupListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetSuccess(v bool) *DescribeContactGroupListResponseBody {
	s.Success = &v
	return s
}

type DescribeContactGroupListResponseBodyContactGroupList struct {
	ContactGroup []*DescribeContactGroupListResponseBodyContactGroupListContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupListResponseBodyContactGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroupList) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroupList) SetContactGroup(v []*DescribeContactGroupListResponseBodyContactGroupListContactGroup) *DescribeContactGroupListResponseBodyContactGroupList {
	s.ContactGroup = v
	return s
}

type DescribeContactGroupListResponseBodyContactGroupListContactGroup struct {
	Describe            *string                                                                   `json:"Describe,omitempty" xml:"Describe,omitempty"`
	UpdateTime          *int64                                                                    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Contacts            *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	CreateTime          *int64                                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnabledWeeklyReport *bool                                                                     `json:"EnabledWeeklyReport,omitempty" xml:"EnabledWeeklyReport,omitempty"`
	Name                *string                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	EnableSubscribed    *bool                                                                     `json:"EnableSubscribed,omitempty" xml:"EnableSubscribed,omitempty"`
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroup) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetDescribe(v string) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.Describe = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetUpdateTime(v int64) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.UpdateTime = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetContacts(v *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.Contacts = v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetCreateTime(v int64) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetEnabledWeeklyReport(v bool) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.EnabledWeeklyReport = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetName(v string) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.Name = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetEnableSubscribed(v bool) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.EnableSubscribed = &v
	return s
}

type DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) SetContact(v []*string) *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts {
	s.Contact = v
	return s
}

type DescribeContactGroupListResponseBodyContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupListResponseBodyContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroups) SetContactGroup(v []*string) *DescribeContactGroupListResponseBodyContactGroups {
	s.ContactGroup = v
	return s
}

type DescribeContactGroupListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContactGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContactGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponse) SetHeaders(v map[string]*string) *DescribeContactGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactGroupListResponse) SetBody(v *DescribeContactGroupListResponseBody) *DescribeContactGroupListResponse {
	s.Body = v
	return s
}

type DescribeContactListRequest struct {
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ChanelType  *string `json:"ChanelType,omitempty" xml:"ChanelType,omitempty"`
	ChanelValue *string `json:"ChanelValue,omitempty" xml:"ChanelValue,omitempty"`
}

func (s DescribeContactListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactListRequest) SetPageSize(v int32) *DescribeContactListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContactListRequest) SetPageNumber(v int32) *DescribeContactListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeContactListRequest) SetContactName(v string) *DescribeContactListRequest {
	s.ContactName = &v
	return s
}

func (s *DescribeContactListRequest) SetChanelType(v string) *DescribeContactListRequest {
	s.ChanelType = &v
	return s
}

func (s *DescribeContactListRequest) SetChanelValue(v string) *DescribeContactListRequest {
	s.ChanelValue = &v
	return s
}

type DescribeContactListResponseBody struct {
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Contacts  *DescribeContactListResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	Total     *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeContactListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBody) SetMessage(v string) *DescribeContactListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContactListResponseBody) SetRequestId(v string) *DescribeContactListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactListResponseBody) SetContacts(v *DescribeContactListResponseBodyContacts) *DescribeContactListResponseBody {
	s.Contacts = v
	return s
}

func (s *DescribeContactListResponseBody) SetTotal(v int32) *DescribeContactListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeContactListResponseBody) SetCode(v string) *DescribeContactListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContactListResponseBody) SetSuccess(v bool) *DescribeContactListResponseBody {
	s.Success = &v
	return s
}

type DescribeContactListResponseBodyContacts struct {
	Contact []*DescribeContactListResponseBodyContactsContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeContactListResponseBodyContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContacts) SetContact(v []*DescribeContactListResponseBodyContactsContact) *DescribeContactListResponseBodyContacts {
	s.Contact = v
	return s
}

type DescribeContactListResponseBodyContactsContact struct {
	UpdateTime    *int64                                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ChannelsState *DescribeContactListResponseBodyContactsContactChannelsState `json:"ChannelsState,omitempty" xml:"ChannelsState,omitempty" type:"Struct"`
	CreateTime    *int64                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Lang          *string                                                      `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ContactGroups *DescribeContactListResponseBodyContactsContactContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	Channels      *DescribeContactListResponseBodyContactsContactChannels      `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	Name          *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc          *string                                                      `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s DescribeContactListResponseBodyContactsContact) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContact) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContact) SetUpdateTime(v int64) *DescribeContactListResponseBodyContactsContact {
	s.UpdateTime = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetChannelsState(v *DescribeContactListResponseBodyContactsContactChannelsState) *DescribeContactListResponseBodyContactsContact {
	s.ChannelsState = v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetCreateTime(v int64) *DescribeContactListResponseBodyContactsContact {
	s.CreateTime = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetLang(v string) *DescribeContactListResponseBodyContactsContact {
	s.Lang = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetContactGroups(v *DescribeContactListResponseBodyContactsContactContactGroups) *DescribeContactListResponseBodyContactsContact {
	s.ContactGroups = v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetChannels(v *DescribeContactListResponseBodyContactsContactChannels) *DescribeContactListResponseBodyContactsContact {
	s.Channels = v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetName(v string) *DescribeContactListResponseBodyContactsContact {
	s.Name = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContact) SetDesc(v string) *DescribeContactListResponseBodyContactsContact {
	s.Desc = &v
	return s
}

type DescribeContactListResponseBodyContactsContactChannelsState struct {
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
	SMS         *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
	Mail        *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	AliIM       *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
}

func (s DescribeContactListResponseBodyContactsContactChannelsState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContactChannelsState) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetDingWebHook(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.DingWebHook = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetSMS(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.SMS = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetMail(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.Mail = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannelsState) SetAliIM(v string) *DescribeContactListResponseBodyContactsContactChannelsState {
	s.AliIM = &v
	return s
}

type DescribeContactListResponseBodyContactsContactContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeContactListResponseBodyContactsContactContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContactContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContactContactGroups) SetContactGroup(v []*string) *DescribeContactListResponseBodyContactsContactContactGroups {
	s.ContactGroup = v
	return s
}

type DescribeContactListResponseBodyContactsContactChannels struct {
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
	SMS         *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
	Mail        *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	AliIM       *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
}

func (s DescribeContactListResponseBodyContactsContactChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListResponseBodyContactsContactChannels) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetDingWebHook(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.DingWebHook = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetSMS(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.SMS = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetMail(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.Mail = &v
	return s
}

func (s *DescribeContactListResponseBodyContactsContactChannels) SetAliIM(v string) *DescribeContactListResponseBodyContactsContactChannels {
	s.AliIM = &v
	return s
}

type DescribeContactListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContactListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContactListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponse) SetHeaders(v map[string]*string) *DescribeContactListResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactListResponse) SetBody(v *DescribeContactListResponseBody) *DescribeContactListResponse {
	s.Body = v
	return s
}

type DescribeContactListByContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
}

func (s DescribeContactListByContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListByContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupRequest) SetContactGroupName(v string) *DescribeContactListByContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

type DescribeContactListByContactGroupResponseBody struct {
	Message   *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Contacts  *DescribeContactListByContactGroupResponseBodyContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	Code      *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeContactListByContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBody) SetMessage(v string) *DescribeContactListByContactGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetRequestId(v string) *DescribeContactListByContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetContacts(v *DescribeContactListByContactGroupResponseBodyContacts) *DescribeContactListByContactGroupResponseBody {
	s.Contacts = v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetCode(v string) *DescribeContactListByContactGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBody) SetSuccess(v bool) *DescribeContactListByContactGroupResponseBody {
	s.Success = &v
	return s
}

type DescribeContactListByContactGroupResponseBodyContacts struct {
	Contact []*DescribeContactListByContactGroupResponseBodyContactsContact `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeContactListByContactGroupResponseBodyContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBodyContacts) SetContact(v []*DescribeContactListByContactGroupResponseBodyContactsContact) *DescribeContactListByContactGroupResponseBodyContacts {
	s.Contact = v
	return s
}

type DescribeContactListByContactGroupResponseBodyContactsContact struct {
	UpdateTime *int64                                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	CreateTime *int64                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Channels   *DescribeContactListByContactGroupResponseBodyContactsContactChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	Name       *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Desc       *string                                                               `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s DescribeContactListByContactGroupResponseBodyContactsContact) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBodyContactsContact) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetUpdateTime(v int64) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.UpdateTime = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetCreateTime(v int64) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.CreateTime = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetChannels(v *DescribeContactListByContactGroupResponseBodyContactsContactChannels) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.Channels = v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetName(v string) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.Name = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContact) SetDesc(v string) *DescribeContactListByContactGroupResponseBodyContactsContact {
	s.Desc = &v
	return s
}

type DescribeContactListByContactGroupResponseBodyContactsContactChannels struct {
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
	SMS         *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
	Mail        *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	AliIM       *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
}

func (s DescribeContactListByContactGroupResponseBodyContactsContactChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListByContactGroupResponseBodyContactsContactChannels) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetDingWebHook(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.DingWebHook = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetSMS(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.SMS = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetMail(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.Mail = &v
	return s
}

func (s *DescribeContactListByContactGroupResponseBodyContactsContactChannels) SetAliIM(v string) *DescribeContactListByContactGroupResponseBodyContactsContactChannels {
	s.AliIM = &v
	return s
}

type DescribeContactListByContactGroupResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContactListByContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContactListByContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactListByContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactListByContactGroupResponse) SetHeaders(v map[string]*string) *DescribeContactListByContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactListByContactGroupResponse) SetBody(v *DescribeContactListByContactGroupResponseBody) *DescribeContactListByContactGroupResponse {
	s.Body = v
	return s
}

type DescribeCustomEventAttributeRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EventId        *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCustomEventAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeRequest) SetName(v string) *DescribeCustomEventAttributeRequest {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetEventId(v string) *DescribeCustomEventAttributeRequest {
	s.EventId = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetGroupId(v string) *DescribeCustomEventAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetSearchKeywords(v string) *DescribeCustomEventAttributeRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetStartTime(v string) *DescribeCustomEventAttributeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetEndTime(v string) *DescribeCustomEventAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetPageNumber(v int32) *DescribeCustomEventAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomEventAttributeRequest) SetPageSize(v int32) *DescribeCustomEventAttributeRequest {
	s.PageSize = &v
	return s
}

type DescribeCustomEventAttributeResponseBody struct {
	CustomEvents *DescribeCustomEventAttributeResponseBodyCustomEvents `json:"CustomEvents,omitempty" xml:"CustomEvents,omitempty" type:"Struct"`
	Message      *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code         *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *string                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomEventAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponseBody) SetCustomEvents(v *DescribeCustomEventAttributeResponseBodyCustomEvents) *DescribeCustomEventAttributeResponseBody {
	s.CustomEvents = v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetMessage(v string) *DescribeCustomEventAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetRequestId(v string) *DescribeCustomEventAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetCode(v string) *DescribeCustomEventAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBody) SetSuccess(v string) *DescribeCustomEventAttributeResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomEventAttributeResponseBodyCustomEvents struct {
	CustomEvent []*DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent `json:"CustomEvent,omitempty" xml:"CustomEvent,omitempty" type:"Repeated"`
}

func (s DescribeCustomEventAttributeResponseBodyCustomEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventAttributeResponseBodyCustomEvents) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEvents) SetCustomEvent(v []*DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) *DescribeCustomEventAttributeResponseBodyCustomEvents {
	s.CustomEvent = v
	return s
}

type DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent struct {
	Time    *string `json:"Time,omitempty" xml:"Time,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetTime(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Time = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetGroupId(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetName(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetContent(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Content = &v
	return s
}

func (s *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent) SetId(v string) *DescribeCustomEventAttributeResponseBodyCustomEventsCustomEvent {
	s.Id = &v
	return s
}

type DescribeCustomEventAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomEventAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomEventAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponse) SetHeaders(v map[string]*string) *DescribeCustomEventAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomEventAttributeResponse) SetBody(v *DescribeCustomEventAttributeResponseBody) *DescribeCustomEventAttributeResponse {
	s.Body = v
	return s
}

type DescribeCustomEventCountRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EventId        *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeCustomEventCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountRequest) SetName(v string) *DescribeCustomEventCountRequest {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetEventId(v string) *DescribeCustomEventCountRequest {
	s.EventId = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetGroupId(v string) *DescribeCustomEventCountRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetSearchKeywords(v string) *DescribeCustomEventCountRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetStartTime(v string) *DescribeCustomEventCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventCountRequest) SetEndTime(v string) *DescribeCustomEventCountRequest {
	s.EndTime = &v
	return s
}

type DescribeCustomEventCountResponseBody struct {
	CustomEventCounts *DescribeCustomEventCountResponseBodyCustomEventCounts `json:"CustomEventCounts,omitempty" xml:"CustomEventCounts,omitempty" type:"Struct"`
	Message           *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code              *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success           *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomEventCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponseBody) SetCustomEventCounts(v *DescribeCustomEventCountResponseBodyCustomEventCounts) *DescribeCustomEventCountResponseBody {
	s.CustomEventCounts = v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetMessage(v string) *DescribeCustomEventCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetRequestId(v string) *DescribeCustomEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetCode(v string) *DescribeCustomEventCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomEventCountResponseBody) SetSuccess(v bool) *DescribeCustomEventCountResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomEventCountResponseBodyCustomEventCounts struct {
	CustomEventCount []*DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount `json:"CustomEventCount,omitempty" xml:"CustomEventCount,omitempty" type:"Repeated"`
}

func (s DescribeCustomEventCountResponseBodyCustomEventCounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventCountResponseBodyCustomEventCounts) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCounts) SetCustomEventCount(v []*DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) *DescribeCustomEventCountResponseBodyCustomEventCounts {
	s.CustomEventCount = v
	return s
}

type DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount struct {
	Time *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	Num  *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) SetTime(v int64) *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount {
	s.Time = &v
	return s
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) SetNum(v int32) *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount {
	s.Num = &v
	return s
}

func (s *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount) SetName(v string) *DescribeCustomEventCountResponseBodyCustomEventCountsCustomEventCount {
	s.Name = &v
	return s
}

type DescribeCustomEventCountResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomEventCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponse) SetHeaders(v map[string]*string) *DescribeCustomEventCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomEventCountResponse) SetBody(v *DescribeCustomEventCountResponseBody) *DescribeCustomEventCountResponse {
	s.Body = v
	return s
}

type DescribeCustomEventHistogramRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	EventId        *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeCustomEventHistogramRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramRequest) SetName(v string) *DescribeCustomEventHistogramRequest {
	s.Name = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetLevel(v string) *DescribeCustomEventHistogramRequest {
	s.Level = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetEventId(v string) *DescribeCustomEventHistogramRequest {
	s.EventId = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetGroupId(v string) *DescribeCustomEventHistogramRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetSearchKeywords(v string) *DescribeCustomEventHistogramRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetStartTime(v string) *DescribeCustomEventHistogramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventHistogramRequest) SetEndTime(v string) *DescribeCustomEventHistogramRequest {
	s.EndTime = &v
	return s
}

type DescribeCustomEventHistogramResponseBody struct {
	Message         *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code            *string                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	EventHistograms *DescribeCustomEventHistogramResponseBodyEventHistograms `json:"EventHistograms,omitempty" xml:"EventHistograms,omitempty" type:"Struct"`
	Success         *string                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomEventHistogramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponseBody) SetMessage(v string) *DescribeCustomEventHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetRequestId(v string) *DescribeCustomEventHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetCode(v string) *DescribeCustomEventHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetEventHistograms(v *DescribeCustomEventHistogramResponseBodyEventHistograms) *DescribeCustomEventHistogramResponseBody {
	s.EventHistograms = v
	return s
}

func (s *DescribeCustomEventHistogramResponseBody) SetSuccess(v string) *DescribeCustomEventHistogramResponseBody {
	s.Success = &v
	return s
}

type DescribeCustomEventHistogramResponseBodyEventHistograms struct {
	EventHistogram []*DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram `json:"EventHistogram,omitempty" xml:"EventHistogram,omitempty" type:"Repeated"`
}

func (s DescribeCustomEventHistogramResponseBodyEventHistograms) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventHistogramResponseBodyEventHistograms) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistograms) SetEventHistogram(v []*DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) *DescribeCustomEventHistogramResponseBodyEventHistograms {
	s.EventHistogram = v
	return s
}

type DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram struct {
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Count     *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) SetEndTime(v int64) *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram {
	s.EndTime = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) SetStartTime(v int64) *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram {
	s.StartTime = &v
	return s
}

func (s *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram) SetCount(v int64) *DescribeCustomEventHistogramResponseBodyEventHistogramsEventHistogram {
	s.Count = &v
	return s
}

type DescribeCustomEventHistogramResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomEventHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomEventHistogramResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomEventHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventHistogramResponse) SetHeaders(v map[string]*string) *DescribeCustomEventHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomEventHistogramResponse) SetBody(v *DescribeCustomEventHistogramResponseBody) *DescribeCustomEventHistogramResponse {
	s.Body = v
	return s
}

type DescribeCustomMetricListRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Dimension  *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	Md5        *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCustomMetricListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomMetricListRequest) SetGroupId(v string) *DescribeCustomMetricListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetMetricName(v string) *DescribeCustomMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetDimension(v string) *DescribeCustomMetricListRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetMd5(v string) *DescribeCustomMetricListRequest {
	s.Md5 = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetPageNumber(v string) *DescribeCustomMetricListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomMetricListRequest) SetPageSize(v string) *DescribeCustomMetricListRequest {
	s.PageSize = &v
	return s
}

type DescribeCustomMetricListResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeCustomMetricListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomMetricListResponseBody) SetMessage(v string) *DescribeCustomMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomMetricListResponseBody) SetRequestId(v string) *DescribeCustomMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomMetricListResponseBody) SetCode(v string) *DescribeCustomMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomMetricListResponseBody) SetResult(v string) *DescribeCustomMetricListResponseBody {
	s.Result = &v
	return s
}

type DescribeCustomMetricListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCustomMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCustomMetricListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCustomMetricListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomMetricListResponse) SetHeaders(v map[string]*string) *DescribeCustomMetricListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomMetricListResponse) SetBody(v *DescribeCustomMetricListResponseBody) *DescribeCustomMetricListResponse {
	s.Body = v
	return s
}

type DescribeDynamicTagRuleListRequest struct {
	TagKey     *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDynamicTagRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListRequest) SetTagKey(v string) *DescribeDynamicTagRuleListRequest {
	s.TagKey = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) SetPageNumber(v string) *DescribeDynamicTagRuleListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDynamicTagRuleListRequest) SetPageSize(v string) *DescribeDynamicTagRuleListRequest {
	s.PageSize = &v
	return s
}

type DescribeDynamicTagRuleListResponseBody struct {
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message      *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize     *string                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *string                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total        *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	TagGroupList *DescribeDynamicTagRuleListResponseBodyTagGroupList `json:"TagGroupList,omitempty" xml:"TagGroupList,omitempty" type:"Struct"`
	Code         *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDynamicTagRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBody) SetRequestId(v string) *DescribeDynamicTagRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetMessage(v string) *DescribeDynamicTagRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetPageSize(v string) *DescribeDynamicTagRuleListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetPageNumber(v string) *DescribeDynamicTagRuleListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetTotal(v int32) *DescribeDynamicTagRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetTagGroupList(v *DescribeDynamicTagRuleListResponseBodyTagGroupList) *DescribeDynamicTagRuleListResponseBody {
	s.TagGroupList = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetCode(v string) *DescribeDynamicTagRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBody) SetSuccess(v bool) *DescribeDynamicTagRuleListResponseBody {
	s.Success = &v
	return s
}

type DescribeDynamicTagRuleListResponseBodyTagGroupList struct {
	TagGroup []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup `json:"TagGroup,omitempty" xml:"TagGroup,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupList) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupList) SetTagGroup(v []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) *DescribeDynamicTagRuleListResponseBodyTagGroupList {
	s.TagGroup = v
	return s
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup struct {
	Status                     *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	MatchExpress               *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress   `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Struct"`
	TemplateIdList             *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Struct"`
	DynamicTagRuleId           *string                                                                   `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	MatchExpressFilterRelation *string                                                                   `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	RegionId                   *string                                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TagKey                     *string                                                                   `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetStatus(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.Status = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetMatchExpress(v *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.MatchExpress = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetTemplateIdList(v *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.TemplateIdList = v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetDynamicTagRuleId(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetMatchExpressFilterRelation(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetRegionId(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup) SetTagKey(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroup {
	s.TagKey = &v
	return s
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress struct {
	MatchExpress []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress) SetMatchExpress(v []*DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpress {
	s.MatchExpress = v
	return s
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress struct {
	TagValueMatchFunction *string `json:"TagValueMatchFunction,omitempty" xml:"TagValueMatchFunction,omitempty"`
	TagValue              *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) SetTagValueMatchFunction(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress {
	s.TagValueMatchFunction = &v
	return s
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress) SetTagValue(v string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupMatchExpressMatchExpress {
	s.TagValue = &v
	return s
}

type DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList struct {
	TemplateIdList []*string `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList) SetTemplateIdList(v []*string) *DescribeDynamicTagRuleListResponseBodyTagGroupListTagGroupTemplateIdList {
	s.TemplateIdList = v
	return s
}

type DescribeDynamicTagRuleListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDynamicTagRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDynamicTagRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDynamicTagRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDynamicTagRuleListResponse) SetHeaders(v map[string]*string) *DescribeDynamicTagRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDynamicTagRuleListResponse) SetBody(v *DescribeDynamicTagRuleListResponseBody) *DescribeDynamicTagRuleListResponse {
	s.Body = v
	return s
}

type DescribeEventRuleAttributeRequest struct {
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeEventRuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeRequest) SetRuleName(v string) *DescribeEventRuleAttributeRequest {
	s.RuleName = &v
	return s
}

type DescribeEventRuleAttributeResponseBody struct {
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Result    *DescribeEventRuleAttributeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeEventRuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBody) SetMessage(v string) *DescribeEventRuleAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetRequestId(v string) *DescribeEventRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetCode(v string) *DescribeEventRuleAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetSuccess(v bool) *DescribeEventRuleAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBody) SetResult(v *DescribeEventRuleAttributeResponseBodyResult) *DescribeEventRuleAttributeResponseBody {
	s.Result = v
	return s
}

type DescribeEventRuleAttributeResponseBodyResult struct {
	EventType    *string                                                   `json:"EventType,omitempty" xml:"EventType,omitempty"`
	GroupId      *string                                                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Description  *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	State        *string                                                   `json:"State,omitempty" xml:"State,omitempty"`
	Name         *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	EventPattern *DescribeEventRuleAttributeResponseBodyResultEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Struct"`
}

func (s DescribeEventRuleAttributeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetEventType(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.EventType = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetGroupId(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.GroupId = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetDescription(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetState(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.State = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetName(v string) *DescribeEventRuleAttributeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResult) SetEventPattern(v *DescribeEventRuleAttributeResponseBodyResultEventPattern) *DescribeEventRuleAttributeResponseBodyResult {
	s.EventPattern = v
	return s
}

type DescribeEventRuleAttributeResponseBodyResultEventPattern struct {
	StatusList *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Struct"`
	Product    *string                                                             `json:"Product,omitempty" xml:"Product,omitempty"`
	LevelList  *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList  `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Struct"`
	NameList   *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList   `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Struct"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPattern) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPattern) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetStatusList(v *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.StatusList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetProduct(v string) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.Product = &v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetLevelList(v *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.LevelList = v
	return s
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPattern) SetNameList(v *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) *DescribeEventRuleAttributeResponseBodyResultEventPattern {
	s.NameList = v
	return s
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList struct {
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList) SetStatusList(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternStatusList {
	s.StatusList = v
	return s
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList struct {
	LevelList []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList) SetLevelList(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternLevelList {
	s.LevelList = v
	return s
}

type DescribeEventRuleAttributeResponseBodyResultEventPatternNameList struct {
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList) SetNameList(v []*string) *DescribeEventRuleAttributeResponseBodyResultEventPatternNameList {
	s.NameList = v
	return s
}

type DescribeEventRuleAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventRuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventRuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleAttributeResponse) SetHeaders(v map[string]*string) *DescribeEventRuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventRuleAttributeResponse) SetBody(v *DescribeEventRuleAttributeResponseBody) *DescribeEventRuleAttributeResponse {
	s.Body = v
	return s
}

type DescribeEventRuleListRequest struct {
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeEventRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListRequest) SetNamePrefix(v string) *DescribeEventRuleListRequest {
	s.NamePrefix = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetPageNumber(v string) *DescribeEventRuleListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetPageSize(v string) *DescribeEventRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventRuleListRequest) SetGroupId(v string) *DescribeEventRuleListRequest {
	s.GroupId = &v
	return s
}

type DescribeEventRuleListResponseBody struct {
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	EventRules *DescribeEventRuleListResponseBodyEventRules `json:"EventRules,omitempty" xml:"EventRules,omitempty" type:"Struct"`
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEventRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBody) SetMessage(v string) *DescribeEventRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetRequestId(v string) *DescribeEventRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetTotal(v int32) *DescribeEventRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetEventRules(v *DescribeEventRuleListResponseBodyEventRules) *DescribeEventRuleListResponseBody {
	s.EventRules = v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetCode(v string) *DescribeEventRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventRuleListResponseBody) SetSuccess(v bool) *DescribeEventRuleListResponseBody {
	s.Success = &v
	return s
}

type DescribeEventRuleListResponseBodyEventRules struct {
	EventRule []*DescribeEventRuleListResponseBodyEventRulesEventRule `json:"EventRule,omitempty" xml:"EventRule,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRules) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRules) SetEventRule(v []*DescribeEventRuleListResponseBodyEventRulesEventRule) *DescribeEventRuleListResponseBodyEventRules {
	s.EventRule = v
	return s
}

type DescribeEventRuleListResponseBodyEventRulesEventRule struct {
	EventType    *string                                                           `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Description  *string                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string                                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	State        *string                                                           `json:"State,omitempty" xml:"State,omitempty"`
	Name         *string                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	EventPattern *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Struct"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRule) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetEventType(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.EventType = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetDescription(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.Description = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetGroupId(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.GroupId = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetState(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.State = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetName(v string) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.Name = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRule) SetEventPattern(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) *DescribeEventRuleListResponseBodyEventRulesEventRule {
	s.EventPattern = v
	return s
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern struct {
	EventPattern []*DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern) SetEventPattern(v []*DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPattern {
	s.EventPattern = v
	return s
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern struct {
	EventTypeList *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Struct"`
	Product       *string                                                                                    `json:"Product,omitempty" xml:"Product,omitempty"`
	LevelList     *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList     `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Struct"`
	NameList      *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList      `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Struct"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetEventTypeList(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.EventTypeList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetProduct(v string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.Product = &v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetLevelList(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.LevelList = v
	return s
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern) SetNameList(v *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPattern {
	s.NameList = v
	return s
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList struct {
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList) SetEventTypeList(v []*string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternEventTypeList {
	s.EventTypeList = v
	return s
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList struct {
	LevelList []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList) SetLevelList(v []*string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternLevelList {
	s.LevelList = v
	return s
}

type DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList struct {
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList) SetNameList(v []*string) *DescribeEventRuleListResponseBodyEventRulesEventRuleEventPatternEventPatternNameList {
	s.NameList = v
	return s
}

type DescribeEventRuleListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponse) SetHeaders(v map[string]*string) *DescribeEventRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventRuleListResponse) SetBody(v *DescribeEventRuleListResponseBody) *DescribeEventRuleListResponse {
	s.Body = v
	return s
}

type DescribeEventRuleTargetListRequest struct {
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeEventRuleTargetListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListRequest) SetRuleName(v string) *DescribeEventRuleTargetListRequest {
	s.RuleName = &v
	return s
}

type DescribeEventRuleTargetListResponseBody struct {
	Message           *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContactParameters *DescribeEventRuleTargetListResponseBodyContactParameters `json:"ContactParameters,omitempty" xml:"ContactParameters,omitempty" type:"Struct"`
	SlsParameters     *DescribeEventRuleTargetListResponseBodySlsParameters     `json:"SlsParameters,omitempty" xml:"SlsParameters,omitempty" type:"Struct"`
	WebhookParameters *DescribeEventRuleTargetListResponseBodyWebhookParameters `json:"WebhookParameters,omitempty" xml:"WebhookParameters,omitempty" type:"Struct"`
	FcParameters      *DescribeEventRuleTargetListResponseBodyFcParameters      `json:"FcParameters,omitempty" xml:"FcParameters,omitempty" type:"Struct"`
	Code              *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	MnsParameters     *DescribeEventRuleTargetListResponseBodyMnsParameters     `json:"MnsParameters,omitempty" xml:"MnsParameters,omitempty" type:"Struct"`
}

func (s DescribeEventRuleTargetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBody) SetMessage(v string) *DescribeEventRuleTargetListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetRequestId(v string) *DescribeEventRuleTargetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetContactParameters(v *DescribeEventRuleTargetListResponseBodyContactParameters) *DescribeEventRuleTargetListResponseBody {
	s.ContactParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetSlsParameters(v *DescribeEventRuleTargetListResponseBodySlsParameters) *DescribeEventRuleTargetListResponseBody {
	s.SlsParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetWebhookParameters(v *DescribeEventRuleTargetListResponseBodyWebhookParameters) *DescribeEventRuleTargetListResponseBody {
	s.WebhookParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetFcParameters(v *DescribeEventRuleTargetListResponseBodyFcParameters) *DescribeEventRuleTargetListResponseBody {
	s.FcParameters = v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetCode(v string) *DescribeEventRuleTargetListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBody) SetMnsParameters(v *DescribeEventRuleTargetListResponseBodyMnsParameters) *DescribeEventRuleTargetListResponseBody {
	s.MnsParameters = v
	return s
}

type DescribeEventRuleTargetListResponseBodyContactParameters struct {
	ContactParameter []*DescribeEventRuleTargetListResponseBodyContactParametersContactParameter `json:"ContactParameter,omitempty" xml:"ContactParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyContactParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyContactParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyContactParameters) SetContactParameter(v []*DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) *DescribeEventRuleTargetListResponseBodyContactParameters {
	s.ContactParameter = v
	return s
}

type DescribeEventRuleTargetListResponseBodyContactParametersContactParameter struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) SetContactGroupName(v string) *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter {
	s.ContactGroupName = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) SetLevel(v string) *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter {
	s.Level = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyContactParametersContactParameter {
	s.Id = &v
	return s
}

type DescribeEventRuleTargetListResponseBodySlsParameters struct {
	SlsParameter []*DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter `json:"SlsParameter,omitempty" xml:"SlsParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodySlsParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodySlsParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodySlsParameters) SetSlsParameter(v []*DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) *DescribeEventRuleTargetListResponseBodySlsParameters {
	s.SlsParameter = v
	return s
}

type DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter struct {
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Arn      *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetLogStore(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.LogStore = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetRegion(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Region = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetProject(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Project = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetArn(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Arn = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodySlsParametersSlsParameter {
	s.Id = &v
	return s
}

type DescribeEventRuleTargetListResponseBodyWebhookParameters struct {
	WebhookParameter []*DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter `json:"WebhookParameter,omitempty" xml:"WebhookParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParameters) SetWebhookParameter(v []*DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) *DescribeEventRuleTargetListResponseBodyWebhookParameters {
	s.WebhookParameter = v
	return s
}

type DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetProtocol(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Protocol = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetUrl(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Url = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetMethod(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Method = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyWebhookParametersWebhookParameter {
	s.Id = &v
	return s
}

type DescribeEventRuleTargetListResponseBodyFcParameters struct {
	FCParameter []*DescribeEventRuleTargetListResponseBodyFcParametersFCParameter `json:"FCParameter,omitempty" xml:"FCParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyFcParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyFcParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyFcParameters) SetFCParameter(v []*DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) *DescribeEventRuleTargetListResponseBodyFcParameters {
	s.FCParameter = v
	return s
}

type DescribeEventRuleTargetListResponseBodyFcParametersFCParameter struct {
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Arn          *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetFunctionName(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.FunctionName = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetRegion(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.Region = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetServiceName(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.ServiceName = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetArn(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.Arn = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyFcParametersFCParameter {
	s.Id = &v
	return s
}

type DescribeEventRuleTargetListResponseBodyMnsParameters struct {
	MnsParameter []*DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter `json:"MnsParameter,omitempty" xml:"MnsParameter,omitempty" type:"Repeated"`
}

func (s DescribeEventRuleTargetListResponseBodyMnsParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyMnsParameters) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParameters) SetMnsParameter(v []*DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) *DescribeEventRuleTargetListResponseBodyMnsParameters {
	s.MnsParameter = v
	return s
}

type DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Queue  *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	Arn    *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetRegion(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Region = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetQueue(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Queue = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetArn(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Arn = &v
	return s
}

func (s *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter) SetId(v string) *DescribeEventRuleTargetListResponseBodyMnsParametersMnsParameter {
	s.Id = &v
	return s
}

type DescribeEventRuleTargetListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEventRuleTargetListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEventRuleTargetListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEventRuleTargetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponse) SetHeaders(v map[string]*string) *DescribeEventRuleTargetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventRuleTargetListResponse) SetBody(v *DescribeEventRuleTargetListResponseBody) *DescribeEventRuleTargetListResponse {
	s.Body = v
	return s
}

type DescribeExporterOutputListRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeExporterOutputListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterOutputListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListRequest) SetPageNumber(v int32) *DescribeExporterOutputListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterOutputListRequest) SetPageSize(v int32) *DescribeExporterOutputListRequest {
	s.PageSize = &v
	return s
}

type DescribeExporterOutputListResponseBody struct {
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total      *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	Datapoints *DescribeExporterOutputListResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	Code       *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExporterOutputListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterOutputListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBody) SetRequestId(v string) *DescribeExporterOutputListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetMessage(v string) *DescribeExporterOutputListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetPageNumber(v int32) *DescribeExporterOutputListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetTotal(v int32) *DescribeExporterOutputListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetDatapoints(v *DescribeExporterOutputListResponseBodyDatapoints) *DescribeExporterOutputListResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetCode(v string) *DescribeExporterOutputListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExporterOutputListResponseBody) SetSuccess(v bool) *DescribeExporterOutputListResponseBody {
	s.Success = &v
	return s
}

type DescribeExporterOutputListResponseBodyDatapoints struct {
	Datapoint []*DescribeExporterOutputListResponseBodyDatapointsDatapoint `json:"Datapoint,omitempty" xml:"Datapoint,omitempty" type:"Repeated"`
}

func (s DescribeExporterOutputListResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterOutputListResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBodyDatapoints) SetDatapoint(v []*DescribeExporterOutputListResponseBodyDatapointsDatapoint) *DescribeExporterOutputListResponseBodyDatapoints {
	s.Datapoint = v
	return s
}

type DescribeExporterOutputListResponseBodyDatapointsDatapoint struct {
	CreateTime *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConfigJson *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty" type:"Struct"`
	DestName   *string                                                              `json:"DestName,omitempty" xml:"DestName,omitempty"`
	DestType   *string                                                              `json:"DestType,omitempty" xml:"DestType,omitempty"`
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapoint) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetCreateTime(v int64) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.CreateTime = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetConfigJson(v *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.ConfigJson = v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetDestName(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.DestName = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapoint) SetDestType(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapoint {
	s.DestType = &v
	return s
}

type DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson struct {
	As       *string `json:"as,omitempty" xml:"as,omitempty"`
	Ak       *string `json:"ak,omitempty" xml:"ak,omitempty"`
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Project  *string `json:"project,omitempty" xml:"project,omitempty"`
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetAs(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.As = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetAk(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Ak = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetEndpoint(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Endpoint = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetProject(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Project = &v
	return s
}

func (s *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson) SetLogstore(v string) *DescribeExporterOutputListResponseBodyDatapointsDatapointConfigJson {
	s.Logstore = &v
	return s
}

type DescribeExporterOutputListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExporterOutputListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExporterOutputListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterOutputListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListResponse) SetHeaders(v map[string]*string) *DescribeExporterOutputListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExporterOutputListResponse) SetBody(v *DescribeExporterOutputListResponseBody) *DescribeExporterOutputListResponse {
	s.Body = v
	return s
}

type DescribeExporterRuleListRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeExporterRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListRequest) SetPageNumber(v int32) *DescribeExporterRuleListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterRuleListRequest) SetPageSize(v int32) *DescribeExporterRuleListRequest {
	s.PageSize = &v
	return s
}

type DescribeExporterRuleListResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total      *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	Datapoints *DescribeExporterRuleListResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	Code       *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExporterRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBody) SetRequestId(v string) *DescribeExporterRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetMessage(v string) *DescribeExporterRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetPageNumber(v int32) *DescribeExporterRuleListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetTotal(v int32) *DescribeExporterRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetDatapoints(v *DescribeExporterRuleListResponseBodyDatapoints) *DescribeExporterRuleListResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetCode(v string) *DescribeExporterRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetSuccess(v bool) *DescribeExporterRuleListResponseBody {
	s.Success = &v
	return s
}

type DescribeExporterRuleListResponseBodyDatapoints struct {
	Datapoint []*DescribeExporterRuleListResponseBodyDatapointsDatapoint `json:"Datapoint,omitempty" xml:"Datapoint,omitempty" type:"Repeated"`
}

func (s DescribeExporterRuleListResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterRuleListResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBodyDatapoints) SetDatapoint(v []*DescribeExporterRuleListResponseBodyDatapointsDatapoint) *DescribeExporterRuleListResponseBodyDatapoints {
	s.Datapoint = v
	return s
}

type DescribeExporterRuleListResponseBodyDatapointsDatapoint struct {
	MetricName    *string                                                         `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Describe      *string                                                         `json:"Describe,omitempty" xml:"Describe,omitempty"`
	TargetWindows *string                                                         `json:"TargetWindows,omitempty" xml:"TargetWindows,omitempty"`
	CreateTime    *int64                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Enabled       *bool                                                           `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	DstName       *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName `json:"DstName,omitempty" xml:"DstName,omitempty" type:"Struct"`
	Dimension     *string                                                         `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	Namespace     *string                                                         `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RuleName      *string                                                         `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapoint) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetMetricName(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.MetricName = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetDescribe(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Describe = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetTargetWindows(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.TargetWindows = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetCreateTime(v int64) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.CreateTime = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetEnabled(v bool) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Enabled = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetDstName(v *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.DstName = v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetDimension(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Dimension = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetNamespace(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Namespace = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetRuleName(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.RuleName = &v
	return s
}

type DescribeExporterRuleListResponseBodyDatapointsDatapointDstName struct {
	DstName []*string `json:"DstName,omitempty" xml:"DstName,omitempty" type:"Repeated"`
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) SetDstName(v []*string) *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName {
	s.DstName = v
	return s
}

type DescribeExporterRuleListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExporterRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExporterRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExporterRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponse) SetHeaders(v map[string]*string) *DescribeExporterRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExporterRuleListResponse) SetBody(v *DescribeExporterRuleListResponseBody) *DescribeExporterRuleListResponse {
	s.Body = v
	return s
}

type DescribeGroupMonitoringAgentProcessRequest struct {
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetGroupId(v string) *DescribeGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetProcessName(v string) *DescribeGroupMonitoringAgentProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetPageNumber(v int32) *DescribeGroupMonitoringAgentProcessRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessRequest) SetPageSize(v int32) *DescribeGroupMonitoringAgentProcessRequest {
	s.PageSize = &v
	return s
}

type DescribeGroupMonitoringAgentProcessResponseBody struct {
	RequestId  *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total      *string                                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Processes  *DescribeGroupMonitoringAgentProcessResponseBodyProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Struct"`
	Code       *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetPageSize(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetPageNumber(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetTotal(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetProcesses(v *DescribeGroupMonitoringAgentProcessResponseBodyProcesses) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Processes = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetCode(v string) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *DescribeGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcesses struct {
	Process []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess `json:"Process,omitempty" xml:"Process,omitempty" type:"Repeated"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcesses) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcesses) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcesses) SetProcess(v []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) *DescribeGroupMonitoringAgentProcessResponseBodyProcesses {
	s.Process = v
	return s
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess struct {
	ProcessName                *string                                                                      `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	MatchExpress               *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Struct"`
	GroupId                    *string                                                                      `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AlertConfig                *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig  `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	MatchExpressFilterRelation *string                                                                      `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	Id                         *string                                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetProcessName(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.ProcessName = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetMatchExpress(v *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.MatchExpress = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetGroupId(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetAlertConfig(v *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.AlertConfig = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetMatchExpressFilterRelation(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess) SetId(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcess {
	s.Id = &v
	return s
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress struct {
	MatchExpress []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress `json:"MatchExpress,omitempty" xml:"MatchExpress,omitempty" type:"Repeated"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress) SetMatchExpress(v []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpress {
	s.MatchExpress = v
	return s
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress struct {
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) SetValue(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress {
	s.Value = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) SetName(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress {
	s.Name = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress) SetFunction(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessMatchExpressMatchExpress {
	s.Function = &v
	return s
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig struct {
	AlertConfig []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig) SetAlertConfig(v []*DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfig {
	s.AlertConfig = v
	return s
}

type DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig struct {
	ComparisonOperator  *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	SilenceTime         *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Webhook             *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Times               *string `json:"Times,omitempty" xml:"Times,omitempty"`
	EscalationsLevel    *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	EffectiveInterval   *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	Threshold           *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics          *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetComparisonOperator(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetSilenceTime(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetWebhook(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Webhook = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetTimes(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Times = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetEscalationsLevel(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.EscalationsLevel = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetNoEffectiveInterval(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetEffectiveInterval(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetThreshold(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig) SetStatistics(v string) *DescribeGroupMonitoringAgentProcessResponseBodyProcessesProcessAlertConfigAlertConfig {
	s.Statistics = &v
	return s
}

type DescribeGroupMonitoringAgentProcessResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupMonitoringAgentProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *DescribeGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupMonitoringAgentProcessResponse) SetBody(v *DescribeGroupMonitoringAgentProcessResponseBody) *DescribeGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

type DescribeHostAvailabilityListRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	GroupId    *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeHostAvailabilityListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListRequest) SetId(v int64) *DescribeHostAvailabilityListRequest {
	s.Id = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetTaskName(v string) *DescribeHostAvailabilityListRequest {
	s.TaskName = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetPageNumber(v int32) *DescribeHostAvailabilityListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetPageSize(v int32) *DescribeHostAvailabilityListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHostAvailabilityListRequest) SetGroupId(v int64) *DescribeHostAvailabilityListRequest {
	s.GroupId = &v
	return s
}

type DescribeHostAvailabilityListResponseBody struct {
	Message   *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	TaskList  *DescribeHostAvailabilityListResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Struct"`
	Code      *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBody) SetMessage(v string) *DescribeHostAvailabilityListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetRequestId(v string) *DescribeHostAvailabilityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetTotal(v int32) *DescribeHostAvailabilityListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetTaskList(v *DescribeHostAvailabilityListResponseBodyTaskList) *DescribeHostAvailabilityListResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetCode(v string) *DescribeHostAvailabilityListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBody) SetSuccess(v bool) *DescribeHostAvailabilityListResponseBody {
	s.Success = &v
	return s
}

type DescribeHostAvailabilityListResponseBodyTaskList struct {
	NodeTaskConfig []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig `json:"NodeTaskConfig,omitempty" xml:"NodeTaskConfig,omitempty" type:"Repeated"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskList) SetNodeTaskConfig(v []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) *DescribeHostAvailabilityListResponseBodyTaskList {
	s.NodeTaskConfig = v
	return s
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig struct {
	TaskType    *string                                                                    `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	GroupName   *string                                                                    `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId     *int64                                                                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	TaskOption  *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption  `json:"TaskOption,omitempty" xml:"TaskOption,omitempty" type:"Struct"`
	TaskName    *string                                                                    `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Disabled    *bool                                                                      `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	AlertConfig *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	TaskScope   *string                                                                    `json:"TaskScope,omitempty" xml:"TaskScope,omitempty"`
	Instances   *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances   `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	Id          *int64                                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskType(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskType = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetGroupName(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.GroupName = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetGroupId(v int64) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.GroupId = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskOption(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskOption = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskName(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskName = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetDisabled(v bool) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.Disabled = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetAlertConfig(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetTaskScope(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.TaskScope = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetInstances(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.Instances = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig) SetId(v int64) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfig {
	s.Id = &v
	return s
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption struct {
	HttpMethod          *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpURI             *string `json:"HttpURI,omitempty" xml:"HttpURI,omitempty"`
	TelnetOrPingHost    *string `json:"TelnetOrPingHost,omitempty" xml:"TelnetOrPingHost,omitempty"`
	HttpResponseCharset *string `json:"HttpResponseCharset,omitempty" xml:"HttpResponseCharset,omitempty"`
	HttpPostContent     *string `json:"HttpPostContent,omitempty" xml:"HttpPostContent,omitempty"`
	HttpNegative        *bool   `json:"HttpNegative,omitempty" xml:"HttpNegative,omitempty"`
	HttpKeyword         *string `json:"HttpKeyword,omitempty" xml:"HttpKeyword,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpMethod(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpMethod = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpURI(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpURI = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetTelnetOrPingHost(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.TelnetOrPingHost = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpResponseCharset(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpResponseCharset = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpPostContent(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpPostContent = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpNegative(v bool) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpNegative = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption) SetHttpKeyword(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigTaskOption {
	s.HttpKeyword = &v
	return s
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig struct {
	SilenceTime    *int32                                                                                   `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	EndTime        *int32                                                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *int32                                                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	NotifyType     *int32                                                                                   `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	EscalationList *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList `json:"EscalationList,omitempty" xml:"EscalationList,omitempty" type:"Struct"`
	WebHook        *string                                                                                  `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetSilenceTime(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetEndTime(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetStartTime(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetNotifyType(v int32) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.NotifyType = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetEscalationList(v *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.EscalationList = v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig) SetWebHook(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfig {
	s.WebHook = &v
	return s
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList struct {
	EscalationList []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList `json:"escalationList,omitempty" xml:"escalationList,omitempty" type:"Repeated"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList) SetEscalationList(v []*DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationList {
	s.EscalationList = v
	return s
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Times      *string `json:"Times,omitempty" xml:"Times,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Aggregate  *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetValue(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Value = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetMetricName(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.MetricName = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetTimes(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Times = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetOperator(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Operator = &v
	return s
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList) SetAggregate(v string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigAlertConfigEscalationListEscalationList {
	s.Aggregate = &v
	return s
}

type DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances struct {
	Instance []*string `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances) SetInstance(v []*string) *DescribeHostAvailabilityListResponseBodyTaskListNodeTaskConfigInstances {
	s.Instance = v
	return s
}

type DescribeHostAvailabilityListResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHostAvailabilityListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHostAvailabilityListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHostAvailabilityListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostAvailabilityListResponse) SetHeaders(v map[string]*string) *DescribeHostAvailabilityListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostAvailabilityListResponse) SetBody(v *DescribeHostAvailabilityListResponseBody) *DescribeHostAvailabilityListResponse {
	s.Body = v
	return s
}

type DescribeLogMonitorAttributeRequest struct {
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
}

func (s DescribeLogMonitorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeRequest) SetMetricName(v string) *DescribeLogMonitorAttributeRequest {
	s.MetricName = &v
	return s
}

type DescribeLogMonitorAttributeResponseBody struct {
	Message    *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LogMonitor *DescribeLogMonitorAttributeResponseBodyLogMonitor `json:"LogMonitor,omitempty" xml:"LogMonitor,omitempty" type:"Struct"`
	Code       *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBody) SetMessage(v string) *DescribeLogMonitorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetRequestId(v string) *DescribeLogMonitorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetLogMonitor(v *DescribeLogMonitorAttributeResponseBodyLogMonitor) *DescribeLogMonitorAttributeResponseBody {
	s.LogMonitor = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetCode(v string) *DescribeLogMonitorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetSuccess(v bool) *DescribeLogMonitorAttributeResponseBody {
	s.Success = &v
	return s
}

type DescribeLogMonitorAttributeResponseBodyLogMonitor struct {
	ValueFilterRelation *string                                                         `json:"ValueFilterRelation,omitempty" xml:"ValueFilterRelation,omitempty"`
	MetricName          *string                                                         `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	ValueFilter         []*DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty" type:"Repeated"`
	SlsRegionId         *string                                                         `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	SlsLogstore         *string                                                         `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	Aggregates          []*DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates  `json:"Aggregates,omitempty" xml:"Aggregates,omitempty" type:"Repeated"`
	Tumblingwindows     []*string                                                       `json:"Tumblingwindows,omitempty" xml:"Tumblingwindows,omitempty" type:"Repeated"`
	GroupId             *int64                                                          `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Groupbys            []*string                                                       `json:"Groupbys,omitempty" xml:"Groupbys,omitempty" type:"Repeated"`
	LogId               *int64                                                          `json:"LogId,omitempty" xml:"LogId,omitempty"`
	MetricExpress       *string                                                         `json:"MetricExpress,omitempty" xml:"MetricExpress,omitempty"`
	GmtCreate           *int64                                                          `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SlsProject          *string                                                         `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitor) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitor) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetValueFilterRelation(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.ValueFilterRelation = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetMetricName(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.MetricName = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetValueFilter(v []*DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.ValueFilter = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetSlsRegionId(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetSlsLogstore(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetAggregates(v []*DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.Aggregates = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetTumblingwindows(v []*string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.Tumblingwindows = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetGroupId(v int64) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.GroupId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetGroupbys(v []*string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.Groupbys = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetLogId(v int64) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.LogId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetMetricExpress(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.MetricExpress = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetGmtCreate(v int64) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.GmtCreate = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetSlsProject(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.SlsProject = &v
	return s
}

type DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) SetKey(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter {
	s.Key = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) SetValue(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter {
	s.Value = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) SetOperator(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter {
	s.Operator = &v
	return s
}

type DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates struct {
	Max       *string `json:"Max,omitempty" xml:"Max,omitempty"`
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Min       *string `json:"Min,omitempty" xml:"Min,omitempty"`
	Function  *string `json:"Function,omitempty" xml:"Function,omitempty"`
	Alias     *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetMax(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Max = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetFieldName(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.FieldName = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetMin(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Min = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetFunction(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Function = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetAlias(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Alias = &v
	return s
}

type DescribeLogMonitorAttributeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogMonitorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogMonitorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponse) SetHeaders(v map[string]*string) *DescribeLogMonitorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogMonitorAttributeResponse) SetBody(v *DescribeLogMonitorAttributeResponseBody) *DescribeLogMonitorAttributeResponse {
	s.Body = v
	return s
}

type DescribeLogMonitorListRequest struct {
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	GroupId     *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeLogMonitorListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListRequest) SetPageNumber(v int32) *DescribeLogMonitorListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogMonitorListRequest) SetPageSize(v int32) *DescribeLogMonitorListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLogMonitorListRequest) SetSearchValue(v string) *DescribeLogMonitorListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeLogMonitorListRequest) SetGroupId(v int64) *DescribeLogMonitorListRequest {
	s.GroupId = &v
	return s
}

type DescribeLogMonitorListResponseBody struct {
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message        *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize       *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total          *int64                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	LogMonitorList []*DescribeLogMonitorListResponseBodyLogMonitorList `json:"LogMonitorList,omitempty" xml:"LogMonitorList,omitempty" type:"Repeated"`
	Code           *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogMonitorListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponseBody) SetRequestId(v string) *DescribeLogMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetMessage(v string) *DescribeLogMonitorListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetPageSize(v int32) *DescribeLogMonitorListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetPageNumber(v int32) *DescribeLogMonitorListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetTotal(v int64) *DescribeLogMonitorListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetLogMonitorList(v []*DescribeLogMonitorListResponseBodyLogMonitorList) *DescribeLogMonitorListResponseBody {
	s.LogMonitorList = v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetCode(v string) *DescribeLogMonitorListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetSuccess(v bool) *DescribeLogMonitorListResponseBody {
	s.Success = &v
	return s
}

type DescribeLogMonitorListResponseBodyLogMonitorList struct {
	ValueFilterRelation *string                                                        `json:"ValueFilterRelation,omitempty" xml:"ValueFilterRelation,omitempty"`
	SlsLogstore         *string                                                        `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	MetricName          *string                                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	ValueFilter         []*DescribeLogMonitorListResponseBodyLogMonitorListValueFilter `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty" type:"Repeated"`
	GroupId             *int64                                                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	LogId               *int64                                                         `json:"LogId,omitempty" xml:"LogId,omitempty"`
	SlsRegionId         *string                                                        `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	GmtCreate           *int64                                                         `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	SlsProject          *string                                                        `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s DescribeLogMonitorListResponseBodyLogMonitorList) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorListResponseBodyLogMonitorList) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetValueFilterRelation(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.ValueFilterRelation = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetSlsLogstore(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetMetricName(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.MetricName = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetValueFilter(v []*DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.ValueFilter = v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetGroupId(v int64) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.GroupId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetLogId(v int64) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.LogId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetSlsRegionId(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetGmtCreate(v int64) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetSlsProject(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.SlsProject = &v
	return s
}

type DescribeLogMonitorListResponseBodyLogMonitorListValueFilter struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) SetKey(v string) *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter {
	s.Key = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) SetValue(v string) *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter {
	s.Value = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) SetOperator(v string) *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter {
	s.Operator = &v
	return s
}

type DescribeLogMonitorListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogMonitorListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogMonitorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponse) SetHeaders(v map[string]*string) *DescribeLogMonitorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogMonitorListResponse) SetBody(v *DescribeLogMonitorListResponseBody) *DescribeLogMonitorListResponse {
	s.Body = v
	return s
}

type DescribeMetricDataRequest struct {
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Express    *string `json:"Express,omitempty" xml:"Express,omitempty"`
	Length     *string `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s DescribeMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataRequest) SetNamespace(v string) *DescribeMetricDataRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricDataRequest) SetMetricName(v string) *DescribeMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricDataRequest) SetPeriod(v string) *DescribeMetricDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricDataRequest) SetStartTime(v string) *DescribeMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricDataRequest) SetEndTime(v string) *DescribeMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricDataRequest) SetDimensions(v string) *DescribeMetricDataRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricDataRequest) SetExpress(v string) *DescribeMetricDataRequest {
	s.Express = &v
	return s
}

func (s *DescribeMetricDataRequest) SetLength(v string) *DescribeMetricDataRequest {
	s.Length = &v
	return s
}

type DescribeMetricDataResponseBody struct {
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponseBody) SetMessage(v string) *DescribeMetricDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetRequestId(v string) *DescribeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetPeriod(v string) *DescribeMetricDataResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetDatapoints(v string) *DescribeMetricDataResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetCode(v string) *DescribeMetricDataResponseBody {
	s.Code = &v
	return s
}

type DescribeMetricDataResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponse) SetHeaders(v map[string]*string) *DescribeMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricDataResponse) SetBody(v *DescribeMetricDataResponseBody) *DescribeMetricDataResponse {
	s.Body = v
	return s
}

type DescribeMetricLastRequest struct {
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Length     *string `json:"Length,omitempty" xml:"Length,omitempty"`
	Express    *string `json:"Express,omitempty" xml:"Express,omitempty"`
}

func (s DescribeMetricLastRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastRequest) SetNamespace(v string) *DescribeMetricLastRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricLastRequest) SetMetricName(v string) *DescribeMetricLastRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricLastRequest) SetPeriod(v string) *DescribeMetricLastRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricLastRequest) SetStartTime(v string) *DescribeMetricLastRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricLastRequest) SetEndTime(v string) *DescribeMetricLastRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricLastRequest) SetDimensions(v string) *DescribeMetricLastRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricLastRequest) SetNextToken(v string) *DescribeMetricLastRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastRequest) SetLength(v string) *DescribeMetricLastRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricLastRequest) SetExpress(v string) *DescribeMetricLastRequest {
	s.Express = &v
	return s
}

type DescribeMetricLastResponseBody struct {
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricLastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBody) SetNextToken(v string) *DescribeMetricLastResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetRequestId(v string) *DescribeMetricLastResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetMessage(v string) *DescribeMetricLastResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetPeriod(v string) *DescribeMetricLastResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetDatapoints(v string) *DescribeMetricLastResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetCode(v string) *DescribeMetricLastResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetSuccess(v bool) *DescribeMetricLastResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricLastResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricLastResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricLastResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricLastResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponse) SetHeaders(v map[string]*string) *DescribeMetricLastResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricLastResponse) SetBody(v *DescribeMetricLastResponseBody) *DescribeMetricLastResponse {
	s.Body = v
	return s
}

type DescribeMetricListRequest struct {
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Length     *string `json:"Length,omitempty" xml:"Length,omitempty"`
	Express    *string `json:"Express,omitempty" xml:"Express,omitempty"`
}

func (s DescribeMetricListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricListRequest) SetNamespace(v string) *DescribeMetricListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricListRequest) SetMetricName(v string) *DescribeMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricListRequest) SetPeriod(v string) *DescribeMetricListRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricListRequest) SetStartTime(v string) *DescribeMetricListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricListRequest) SetEndTime(v string) *DescribeMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricListRequest) SetDimensions(v string) *DescribeMetricListRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricListRequest) SetNextToken(v string) *DescribeMetricListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricListRequest) SetLength(v string) *DescribeMetricListRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricListRequest) SetExpress(v string) *DescribeMetricListRequest {
	s.Express = &v
	return s
}

type DescribeMetricListResponseBody struct {
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBody) SetNextToken(v string) *DescribeMetricListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetRequestId(v string) *DescribeMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetMessage(v string) *DescribeMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetPeriod(v string) *DescribeMetricListResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetDatapoints(v string) *DescribeMetricListResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetCode(v string) *DescribeMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetSuccess(v bool) *DescribeMetricListResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponse) SetHeaders(v map[string]*string) *DescribeMetricListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricListResponse) SetBody(v *DescribeMetricListResponseBody) *DescribeMetricListResponse {
	s.Body = v
	return s
}

type DescribeMetricMetaListRequest struct {
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Labels     *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMetricMetaListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListRequest) SetNamespace(v string) *DescribeMetricMetaListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetLabels(v string) *DescribeMetricMetaListRequest {
	s.Labels = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetMetricName(v string) *DescribeMetricMetaListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetPageNumber(v int32) *DescribeMetricMetaListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetricMetaListRequest) SetPageSize(v int32) *DescribeMetricMetaListRequest {
	s.PageSize = &v
	return s
}

type DescribeMetricMetaListResponseBody struct {
	TotalCount *string                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  *DescribeMetricMetaListResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricMetaListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBody) SetTotalCount(v string) *DescribeMetricMetaListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetMessage(v string) *DescribeMetricMetaListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetRequestId(v string) *DescribeMetricMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetResources(v *DescribeMetricMetaListResponseBodyResources) *DescribeMetricMetaListResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetCode(v string) *DescribeMetricMetaListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetSuccess(v bool) *DescribeMetricMetaListResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricMetaListResponseBodyResources struct {
	Resource []*DescribeMetricMetaListResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMetricMetaListResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricMetaListResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBodyResources) SetResource(v []*DescribeMetricMetaListResponseBodyResourcesResource) *DescribeMetricMetaListResponseBodyResources {
	s.Resource = v
	return s
}

type DescribeMetricMetaListResponseBodyResourcesResource struct {
	MetricName  *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Unit        *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Dimensions  *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Periods     *string `json:"Periods,omitempty" xml:"Periods,omitempty"`
	Statistics  *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeMetricMetaListResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricMetaListResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetMetricName(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetDescription(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Description = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetLabels(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Labels = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetUnit(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Unit = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetDimensions(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetNamespace(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetPeriods(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Periods = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesResource) SetStatistics(v string) *DescribeMetricMetaListResponseBodyResourcesResource {
	s.Statistics = &v
	return s
}

type DescribeMetricMetaListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricMetaListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricMetaListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponse) SetHeaders(v map[string]*string) *DescribeMetricMetaListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricMetaListResponse) SetBody(v *DescribeMetricMetaListResponseBody) *DescribeMetricMetaListResponse {
	s.Body = v
	return s
}

type DescribeMetricRuleCountRequest struct {
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
}

func (s DescribeMetricRuleCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountRequest) SetNamespace(v string) *DescribeMetricRuleCountRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleCountRequest) SetMetricName(v string) *DescribeMetricRuleCountRequest {
	s.MetricName = &v
	return s
}

type DescribeMetricRuleCountResponseBody struct {
	MetricRuleCount *DescribeMetricRuleCountResponseBodyMetricRuleCount `json:"MetricRuleCount,omitempty" xml:"MetricRuleCount,omitempty" type:"Struct"`
	Message         *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code            *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success         *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricRuleCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountResponseBody) SetMetricRuleCount(v *DescribeMetricRuleCountResponseBodyMetricRuleCount) *DescribeMetricRuleCountResponseBody {
	s.MetricRuleCount = v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetMessage(v string) *DescribeMetricRuleCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetRequestId(v string) *DescribeMetricRuleCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetCode(v string) *DescribeMetricRuleCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBody) SetSuccess(v bool) *DescribeMetricRuleCountResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricRuleCountResponseBodyMetricRuleCount struct {
	Ok      *int32 `json:"Ok,omitempty" xml:"Ok,omitempty"`
	Nodata  *int32 `json:"Nodata,omitempty" xml:"Nodata,omitempty"`
	Disable *int32 `json:"Disable,omitempty" xml:"Disable,omitempty"`
	Total   *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	Alarm   *int32 `json:"Alarm,omitempty" xml:"Alarm,omitempty"`
}

func (s DescribeMetricRuleCountResponseBodyMetricRuleCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleCountResponseBodyMetricRuleCount) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetOk(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Ok = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetNodata(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Nodata = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetDisable(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Disable = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetTotal(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleCountResponseBodyMetricRuleCount) SetAlarm(v int32) *DescribeMetricRuleCountResponseBodyMetricRuleCount {
	s.Alarm = &v
	return s
}

type DescribeMetricRuleCountResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricRuleCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricRuleCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleCountResponse) SetBody(v *DescribeMetricRuleCountResponseBody) *DescribeMetricRuleCountResponse {
	s.Body = v
	return s
}

type DescribeMetricRuleListRequest struct {
	MetricName  *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	EnableState *bool   `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Page        *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AlertState  *string `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	Dimensions  *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RuleIds     *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s DescribeMetricRuleListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListRequest) SetMetricName(v string) *DescribeMetricRuleListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetEnableState(v bool) *DescribeMetricRuleListRequest {
	s.EnableState = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetNamespace(v string) *DescribeMetricRuleListRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetPage(v int32) *DescribeMetricRuleListRequest {
	s.Page = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetPageSize(v int32) *DescribeMetricRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetAlertState(v string) *DescribeMetricRuleListRequest {
	s.AlertState = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetDimensions(v string) *DescribeMetricRuleListRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRuleName(v string) *DescribeMetricRuleListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetGroupId(v string) *DescribeMetricRuleListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleListRequest) SetRuleIds(v string) *DescribeMetricRuleListRequest {
	s.RuleIds = &v
	return s
}

type DescribeMetricRuleListResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *string                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Alarms    *DescribeMetricRuleListResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Struct"`
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricRuleListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBody) SetMessage(v string) *DescribeMetricRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetRequestId(v string) *DescribeMetricRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetTotal(v string) *DescribeMetricRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetAlarms(v *DescribeMetricRuleListResponseBodyAlarms) *DescribeMetricRuleListResponseBody {
	s.Alarms = v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetCode(v int32) *DescribeMetricRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetSuccess(v bool) *DescribeMetricRuleListResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarms struct {
	Alarm []*DescribeMetricRuleListResponseBodyAlarmsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarms) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarms) SetAlarm(v []*DescribeMetricRuleListResponseBodyAlarmsAlarm) *DescribeMetricRuleListResponseBodyAlarms {
	s.Alarm = v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarm struct {
	SilenceTime         *int32                                                    `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	MetricName          *string                                                   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Webhook             *string                                                   `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Escalations         *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	ContactGroups       *string                                                   `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	SourceType          *string                                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Namespace           *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	EffectiveInterval   *string                                                   `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	NoEffectiveInterval *string                                                   `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	MailSubject         *string                                                   `json:"MailSubject,omitempty" xml:"MailSubject,omitempty"`
	RuleName            *string                                                   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	AlertState          *string                                                   `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	RuleId              *string                                                   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Period              *string                                                   `json:"Period,omitempty" xml:"Period,omitempty"`
	GroupName           *string                                                   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId             *string                                                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Dimensions          *string                                                   `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EnableState         *bool                                                     `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	Resources           *string                                                   `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSilenceTime(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMetricName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetWebhook(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Webhook = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEscalations(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Escalations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetContactGroups(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSourceType(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SourceType = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNamespace(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNoEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMailSubject(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MailSubject = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetAlertState(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.AlertState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetPeriod(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetDimensions(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEnableState(v bool) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EnableState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetResources(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Resources = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations struct {
	Critical *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetCritical(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Critical = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetInfo(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Info = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetWarn(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Warn = v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Statistics = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Statistics = &v
	return s
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Statistics = &v
	return s
}

type DescribeMetricRuleListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricRuleListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleListResponse) SetBody(v *DescribeMetricRuleListResponseBody) *DescribeMetricRuleListResponse {
	s.Body = v
	return s
}

type DescribeMetricRuleTargetsRequest struct {
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMetricRuleTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsRequest) SetRuleId(v string) *DescribeMetricRuleTargetsRequest {
	s.RuleId = &v
	return s
}

type DescribeMetricRuleTargetsResponseBody struct {
	Message   *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Targets   *DescribeMetricRuleTargetsResponseBodyTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Struct"`
	Code      *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricRuleTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponseBody) SetMessage(v string) *DescribeMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetRequestId(v string) *DescribeMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetTargets(v *DescribeMetricRuleTargetsResponseBodyTargets) *DescribeMetricRuleTargetsResponseBody {
	s.Targets = v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetCode(v string) *DescribeMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBody) SetSuccess(v bool) *DescribeMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricRuleTargetsResponseBodyTargets struct {
	Target []*DescribeMetricRuleTargetsResponseBodyTargetsTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTargetsResponseBodyTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponseBodyTargets) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponseBodyTargets) SetTarget(v []*DescribeMetricRuleTargetsResponseBodyTargetsTarget) *DescribeMetricRuleTargetsResponseBodyTargets {
	s.Target = v
	return s
}

type DescribeMetricRuleTargetsResponseBodyTargetsTarget struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Arn   *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeMetricRuleTargetsResponseBodyTargetsTarget) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponseBodyTargetsTarget) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) SetId(v string) *DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	s.Id = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) SetArn(v string) *DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	s.Arn = &v
	return s
}

func (s *DescribeMetricRuleTargetsResponseBodyTargetsTarget) SetLevel(v string) *DescribeMetricRuleTargetsResponseBodyTargetsTarget {
	s.Level = &v
	return s
}

type DescribeMetricRuleTargetsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricRuleTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleTargetsResponse) SetBody(v *DescribeMetricRuleTargetsResponseBody) *DescribeMetricRuleTargetsResponse {
	s.Body = v
	return s
}

type DescribeMetricRuleTemplateAttributeRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeRequest) SetName(v string) *DescribeMetricRuleTemplateAttributeRequest {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeRequest) SetTemplateId(v string) *DescribeMetricRuleTemplateAttributeRequest {
	s.TemplateId = &v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBody struct {
	Message   *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource  *DescribeMetricRuleTemplateAttributeResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	Code      *int32                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetMessage(v string) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetRequestId(v string) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetResource(v *DescribeMetricRuleTemplateAttributeResponseBodyResource) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Resource = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetCode(v int32) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetSuccess(v bool) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBodyResource struct {
	Description    *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	AlertTemplates *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Struct"`
	Name           *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	RestVersion    *string                                                                `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	TemplateId     *string                                                                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetDescription(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.Description = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetAlertTemplates(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.AlertTemplates = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetName(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetRestVersion(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.RestVersion = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetTemplateId(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.TemplateId = &v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates struct {
	AlertTemplate []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate `json:"AlertTemplate,omitempty" xml:"AlertTemplate,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) SetAlertTemplate(v []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates {
	s.AlertTemplate = v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate struct {
	MetricName  *string                                                                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Selector    *string                                                                                        `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Webhook     *string                                                                                        `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Escalations *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	Namespace   *string                                                                                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Category    *string                                                                                        `json:"Category,omitempty" xml:"Category,omitempty"`
	RuleName    *string                                                                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetMetricName(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetSelector(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Selector = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetWebhook(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Webhook = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetEscalations(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Escalations = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetNamespace(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetCategory(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Category = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetRuleName(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.RuleName = &v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations struct {
	Critical *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) SetCritical(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations {
	s.Critical = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) SetInfo(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations {
	s.Info = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) SetWarn(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations {
	s.Warn = v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetComparisonOperator(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetTimes(v int32) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetThreshold(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetStatistics(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.Statistics = &v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetComparisonOperator(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetTimes(v int32) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetThreshold(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetStatistics(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.Statistics = &v
	return s
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetComparisonOperator(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetTimes(v int32) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetThreshold(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetStatistics(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.Statistics = &v
	return s
}

type DescribeMetricRuleTemplateAttributeResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricRuleTemplateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricRuleTemplateAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleTemplateAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponse) SetBody(v *DescribeMetricRuleTemplateAttributeResponseBody) *DescribeMetricRuleTemplateAttributeResponse {
	s.Body = v
	return s
}

type DescribeMetricRuleTemplateListRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	TemplateId *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	History    *bool   `json:"History,omitempty" xml:"History,omitempty"`
}

func (s DescribeMetricRuleTemplateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListRequest) SetName(v string) *DescribeMetricRuleTemplateListRequest {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetKeyword(v string) *DescribeMetricRuleTemplateListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetTemplateId(v int64) *DescribeMetricRuleTemplateListRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetPageNumber(v int64) *DescribeMetricRuleTemplateListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetPageSize(v int64) *DescribeMetricRuleTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricRuleTemplateListRequest) SetHistory(v bool) *DescribeMetricRuleTemplateListRequest {
	s.History = &v
	return s
}

type DescribeMetricRuleTemplateListResponseBody struct {
	Message   *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
	Templates *DescribeMetricRuleTemplateListResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
	Code      *int32                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricRuleTemplateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetMessage(v string) *DescribeMetricRuleTemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetRequestId(v string) *DescribeMetricRuleTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetTotal(v int64) *DescribeMetricRuleTemplateListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetTemplates(v *DescribeMetricRuleTemplateListResponseBodyTemplates) *DescribeMetricRuleTemplateListResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetCode(v int32) *DescribeMetricRuleTemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBody) SetSuccess(v bool) *DescribeMetricRuleTemplateListResponseBody {
	s.Success = &v
	return s
}

type DescribeMetricRuleTemplateListResponseBodyTemplates struct {
	Template []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplates) SetTemplate(v []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) *DescribeMetricRuleTemplateListResponseBodyTemplates {
	s.Template = v
	return s
}

type DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate struct {
	ApplyHistories *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories `json:"ApplyHistories,omitempty" xml:"ApplyHistories,omitempty" type:"Struct"`
	Description    *string                                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate      *int64                                                                     `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name           *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	RestVersion    *int64                                                                     `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	GmtModified    *int64                                                                     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TemplateId     *int64                                                                     `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetApplyHistories(v *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.ApplyHistories = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetDescription(v string) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.Description = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetGmtCreate(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetName(v string) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetRestVersion(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.RestVersion = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetGmtModified(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.GmtModified = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate) SetTemplateId(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplate {
	s.TemplateId = &v
	return s
}

type DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories struct {
	ApplyHistory []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory `json:"ApplyHistory,omitempty" xml:"ApplyHistory,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories) SetApplyHistory(v []*DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistories {
	s.ApplyHistory = v
	return s
}

type DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory struct {
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ApplyTime *int64  `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) SetGroupId(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) SetGroupName(v string) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory {
	s.GroupName = &v
	return s
}

func (s *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory) SetApplyTime(v int64) *DescribeMetricRuleTemplateListResponseBodyTemplatesTemplateApplyHistoriesApplyHistory {
	s.ApplyTime = &v
	return s
}

type DescribeMetricRuleTemplateListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricRuleTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricRuleTemplateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricRuleTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateListResponse) SetHeaders(v map[string]*string) *DescribeMetricRuleTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricRuleTemplateListResponse) SetBody(v *DescribeMetricRuleTemplateListResponseBody) *DescribeMetricRuleTemplateListResponse {
	s.Body = v
	return s
}

type DescribeMetricTopRequest struct {
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Orderby    *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	OrderDesc  *string `json:"OrderDesc,omitempty" xml:"OrderDesc,omitempty"`
	Length     *string `json:"Length,omitempty" xml:"Length,omitempty"`
	Express    *string `json:"Express,omitempty" xml:"Express,omitempty"`
}

func (s DescribeMetricTopRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricTopRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopRequest) SetPeriod(v string) *DescribeMetricTopRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricTopRequest) SetNamespace(v string) *DescribeMetricTopRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricTopRequest) SetMetricName(v string) *DescribeMetricTopRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricTopRequest) SetStartTime(v string) *DescribeMetricTopRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricTopRequest) SetEndTime(v string) *DescribeMetricTopRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricTopRequest) SetDimensions(v string) *DescribeMetricTopRequest {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricTopRequest) SetOrderby(v string) *DescribeMetricTopRequest {
	s.Orderby = &v
	return s
}

func (s *DescribeMetricTopRequest) SetOrderDesc(v string) *DescribeMetricTopRequest {
	s.OrderDesc = &v
	return s
}

func (s *DescribeMetricTopRequest) SetLength(v string) *DescribeMetricTopRequest {
	s.Length = &v
	return s
}

func (s *DescribeMetricTopRequest) SetExpress(v string) *DescribeMetricTopRequest {
	s.Express = &v
	return s
}

type DescribeMetricTopResponseBody struct {
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeMetricTopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponseBody) SetMessage(v string) *DescribeMetricTopResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetRequestId(v string) *DescribeMetricTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetPeriod(v string) *DescribeMetricTopResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetDatapoints(v string) *DescribeMetricTopResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetCode(v string) *DescribeMetricTopResponseBody {
	s.Code = &v
	return s
}

type DescribeMetricTopResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricTopResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricTopResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricTopResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponse) SetHeaders(v map[string]*string) *DescribeMetricTopResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricTopResponse) SetBody(v *DescribeMetricTopResponseBody) *DescribeMetricTopResponse {
	s.Body = v
	return s
}

type DescribeMonitorGroupCategoriesRequest struct {
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeMonitorGroupCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesRequest) SetGroupId(v int64) *DescribeMonitorGroupCategoriesRequest {
	s.GroupId = &v
	return s
}

type DescribeMonitorGroupCategoriesResponseBody struct {
	Message                *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId              *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MonitorGroupCategories *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories `json:"MonitorGroupCategories,omitempty" xml:"MonitorGroupCategories,omitempty" type:"Struct"`
	Code                   *int32                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success                *bool                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetMessage(v string) *DescribeMonitorGroupCategoriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetRequestId(v string) *DescribeMonitorGroupCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetMonitorGroupCategories(v *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) *DescribeMonitorGroupCategoriesResponseBody {
	s.MonitorGroupCategories = v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetCode(v int32) *DescribeMonitorGroupCategoriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBody) SetSuccess(v bool) *DescribeMonitorGroupCategoriesResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories struct {
	GroupId              *int64                                                                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MonitorGroupCategory *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory `json:"MonitorGroupCategory,omitempty" xml:"MonitorGroupCategory,omitempty" type:"Struct"`
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) SetGroupId(v int64) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories) SetMonitorGroupCategory(v *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategories {
	s.MonitorGroupCategory = v
	return s
}

type DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory struct {
	CategoryItem []*DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem `json:"CategoryItem,omitempty" xml:"CategoryItem,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory) SetCategoryItem(v []*DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategory {
	s.CategoryItem = v
	return s
}

type DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) SetCategory(v string) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem) SetCount(v int32) *DescribeMonitorGroupCategoriesResponseBodyMonitorGroupCategoriesMonitorGroupCategoryCategoryItem {
	s.Count = &v
	return s
}

type DescribeMonitorGroupCategoriesResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorGroupCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorGroupCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupCategoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponse) SetBody(v *DescribeMonitorGroupCategoriesResponseBody) *DescribeMonitorGroupCategoriesResponse {
	s.Body = v
	return s
}

type DescribeMonitorGroupDynamicRulesRequest struct {
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeMonitorGroupDynamicRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesRequest) SetGroupId(v int64) *DescribeMonitorGroupDynamicRulesRequest {
	s.GroupId = &v
	return s
}

type DescribeMonitorGroupDynamicRulesResponseBody struct {
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource  *DescribeMonitorGroupDynamicRulesResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	Code      *int32                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetMessage(v string) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetRequestId(v string) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetResource(v *DescribeMonitorGroupDynamicRulesResponseBodyResource) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Resource = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetCode(v int32) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetSuccess(v bool) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitorGroupDynamicRulesResponseBodyResource struct {
	Resource []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResource) SetResource(v []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) *DescribeMonitorGroupDynamicRulesResponseBodyResource {
	s.Resource = v
	return s
}

type DescribeMonitorGroupDynamicRulesResponseBodyResourceResource struct {
	FilterRelation *string                                                              `json:"FilterRelation,omitempty" xml:"FilterRelation,omitempty"`
	Filters        *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	Category       *string                                                              `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) SetFilterRelation(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource {
	s.FilterRelation = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) SetFilters(v *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource {
	s.Filters = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) SetCategory(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource {
	s.Category = &v
	return s
}

type DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters struct {
	Filter []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) SetFilter(v []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters {
	s.Filter = v
	return s
}

type DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter struct {
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) SetValue(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter {
	s.Value = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) SetFunction(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter {
	s.Function = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) SetName(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter {
	s.Name = &v
	return s
}

type DescribeMonitorGroupDynamicRulesResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorGroupDynamicRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorGroupDynamicRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupDynamicRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponse) SetBody(v *DescribeMonitorGroupDynamicRulesResponseBody) *DescribeMonitorGroupDynamicRulesResponse {
	s.Body = v
	return s
}

type DescribeMonitorGroupInstanceAttributeRequest struct {
	GroupId     *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total       *bool   `json:"Total,omitempty" xml:"Total,omitempty"`
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetGroupId(v int64) *DescribeMonitorGroupInstanceAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetPageNumber(v int32) *DescribeMonitorGroupInstanceAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetPageSize(v int32) *DescribeMonitorGroupInstanceAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetTotal(v bool) *DescribeMonitorGroupInstanceAttributeRequest {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetCategory(v string) *DescribeMonitorGroupInstanceAttributeRequest {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetKeyword(v string) *DescribeMonitorGroupInstanceAttributeRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeRequest) SetInstanceIds(v string) *DescribeMonitorGroupInstanceAttributeRequest {
	s.InstanceIds = &v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponseBody struct {
	RequestId  *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total      *int32                                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	Resources  *DescribeMonitorGroupInstanceAttributeResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Code       *int32                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetRequestId(v string) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetMessage(v string) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetPageSize(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetPageNumber(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetTotal(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetResources(v *DescribeMonitorGroupInstanceAttributeResponseBodyResources) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetCode(v int32) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBody) SetSuccess(v bool) *DescribeMonitorGroupInstanceAttributeResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResources struct {
	Resource []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResources) SetResource(v []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) *DescribeMonitorGroupInstanceAttributeResponseBodyResources {
	s.Resource = v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource struct {
	InstanceName *string                                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Region       *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Struct"`
	Vpc          *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc    `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Struct"`
	Dimension    *string                                                                   `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	Tags         *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Category     *string                                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	InstanceId   *string                                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkType  *string                                                                   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Desc         *string                                                                   `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetInstanceName(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.InstanceName = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetRegion(v *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Region = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetVpc(v *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Vpc = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetDimension(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Dimension = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetTags(v *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Tags = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetCategory(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetInstanceId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetNetworkType(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.NetworkType = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource) SetDesc(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResource {
	s.Desc = &v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion struct {
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) SetAvailabilityZone(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion {
	s.AvailabilityZone = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion) SetRegionId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceRegion {
	s.RegionId = &v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc struct {
	VswitchInstanceId *string `json:"VswitchInstanceId,omitempty" xml:"VswitchInstanceId,omitempty"`
	VpcInstanceId     *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) SetVswitchInstanceId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc {
	s.VswitchInstanceId = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc) SetVpcInstanceId(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceVpc {
	s.VpcInstanceId = &v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags struct {
	Tag []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags) SetTag(v []*DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTags {
	s.Tag = v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) SetKey(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag) SetValue(v string) *DescribeMonitorGroupInstanceAttributeResponseBodyResourcesResourceTagsTag {
	s.Value = &v
	return s
}

type DescribeMonitorGroupInstanceAttributeResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorGroupInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorGroupInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupInstanceAttributeResponse) SetBody(v *DescribeMonitorGroupInstanceAttributeResponseBody) *DescribeMonitorGroupInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeMonitorGroupInstancesRequest struct {
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	GroupId     *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DescribeMonitorGroupInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesRequest) SetPageSize(v int32) *DescribeMonitorGroupInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetPageNumber(v int32) *DescribeMonitorGroupInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetGroupId(v int64) *DescribeMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetCategory(v string) *DescribeMonitorGroupInstancesRequest {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetKeyword(v string) *DescribeMonitorGroupInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMonitorGroupInstancesRequest) SetInstanceIds(v string) *DescribeMonitorGroupInstancesRequest {
	s.InstanceIds = &v
	return s
}

type DescribeMonitorGroupInstancesResponseBody struct {
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total      *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	Resources  *DescribeMonitorGroupInstancesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Code       *int32                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetRequestId(v string) *DescribeMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetMessage(v string) *DescribeMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetPageSize(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetPageNumber(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetTotal(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetResources(v *DescribeMonitorGroupInstancesResponseBodyResources) *DescribeMonitorGroupInstancesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetCode(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetSuccess(v bool) *DescribeMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitorGroupInstancesResponseBodyResources struct {
	Resource []*DescribeMonitorGroupInstancesResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupInstancesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponseBodyResources) SetResource(v []*DescribeMonitorGroupInstancesResponseBodyResourcesResource) *DescribeMonitorGroupInstancesResponseBodyResources {
	s.Resource = v
	return s
}

type DescribeMonitorGroupInstancesResponseBodyResourcesResource struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitorGroupInstancesResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetInstanceName(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.InstanceName = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetCategory(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetInstanceId(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetId(v int64) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.Id = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetRegionId(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.RegionId = &v
	return s
}

type DescribeMonitorGroupInstancesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorGroupInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupInstancesResponse) SetBody(v *DescribeMonitorGroupInstancesResponseBody) *DescribeMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

type DescribeMonitorGroupNotifyPolicyListRequest struct {
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeMonitorGroupNotifyPolicyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetPolicyType(v string) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.PolicyType = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetPageNumber(v int32) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetPageSize(v int32) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetGroupId(v string) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.GroupId = &v
	return s
}

type DescribeMonitorGroupNotifyPolicyListResponseBody struct {
	NotifyPolicyList *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList `json:"NotifyPolicyList,omitempty" xml:"NotifyPolicyList,omitempty" type:"Struct"`
	Message          *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total            *int32                                                            `json:"Total,omitempty" xml:"Total,omitempty"`
	Code             *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success          *string                                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetNotifyPolicyList(v *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.NotifyPolicyList = v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetMessage(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetRequestId(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetTotal(v int32) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetCode(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBody) SetSuccess(v string) *DescribeMonitorGroupNotifyPolicyListResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList struct {
	NotifyPolicy []*DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy `json:"NotifyPolicy,omitempty" xml:"NotifyPolicy,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList) SetNotifyPolicy(v []*DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyList {
	s.NotifyPolicy = v
	return s
}

type DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetEndTime(v int64) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.EndTime = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetType(v string) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.Type = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetStartTime(v int64) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.StartTime = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetGroupId(v string) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy) SetId(v string) *DescribeMonitorGroupNotifyPolicyListResponseBodyNotifyPolicyListNotifyPolicy {
	s.Id = &v
	return s
}

type DescribeMonitorGroupNotifyPolicyListResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorGroupNotifyPolicyListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorGroupNotifyPolicyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupNotifyPolicyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListResponse) SetBody(v *DescribeMonitorGroupNotifyPolicyListResponseBody) *DescribeMonitorGroupNotifyPolicyListResponse {
	s.Body = v
	return s
}

type DescribeMonitorGroupsRequest struct {
	SelectContactGroups    *bool                              `json:"SelectContactGroups,omitempty" xml:"SelectContactGroups,omitempty"`
	PageNumber             *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Keyword                *string                            `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	InstanceId             *string                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	GroupName              *string                            `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IncludeTemplateHistory *bool                              `json:"IncludeTemplateHistory,omitempty" xml:"IncludeTemplateHistory,omitempty"`
	Type                   *string                            `json:"Type,omitempty" xml:"Type,omitempty"`
	DynamicTagRuleId       *string                            `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	GroupId                *string                            `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ServiceId              *string                            `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ResourceGroupId        *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                    []*DescribeMonitorGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsRequest) SetSelectContactGroups(v bool) *DescribeMonitorGroupsRequest {
	s.SelectContactGroups = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetPageNumber(v int32) *DescribeMonitorGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetPageSize(v int32) *DescribeMonitorGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetKeyword(v string) *DescribeMonitorGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetInstanceId(v string) *DescribeMonitorGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetGroupName(v string) *DescribeMonitorGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetIncludeTemplateHistory(v bool) *DescribeMonitorGroupsRequest {
	s.IncludeTemplateHistory = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetType(v string) *DescribeMonitorGroupsRequest {
	s.Type = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetDynamicTagRuleId(v string) *DescribeMonitorGroupsRequest {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetGroupId(v string) *DescribeMonitorGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetServiceId(v string) *DescribeMonitorGroupsRequest {
	s.ServiceId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetResourceGroupId(v string) *DescribeMonitorGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMonitorGroupsRequest) SetTag(v []*DescribeMonitorGroupsRequestTag) *DescribeMonitorGroupsRequest {
	s.Tag = v
	return s
}

type DescribeMonitorGroupsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsRequestTag) SetKey(v string) *DescribeMonitorGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeMonitorGroupsRequestTag) SetValue(v string) *DescribeMonitorGroupsRequestTag {
	s.Value = &v
	return s
}

type DescribeMonitorGroupsResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total      *int32                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	Resources  *DescribeMonitorGroupsResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Code       *int32                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBody) SetRequestId(v string) *DescribeMonitorGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetMessage(v string) *DescribeMonitorGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetPageSize(v int32) *DescribeMonitorGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetPageNumber(v int32) *DescribeMonitorGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetTotal(v int32) *DescribeMonitorGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetResources(v *DescribeMonitorGroupsResponseBodyResources) *DescribeMonitorGroupsResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetCode(v int32) *DescribeMonitorGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBody) SetSuccess(v bool) *DescribeMonitorGroupsResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitorGroupsResponseBodyResources struct {
	Resource []*DescribeMonitorGroupsResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResources) SetResource(v []*DescribeMonitorGroupsResponseBodyResourcesResource) *DescribeMonitorGroupsResponseBodyResources {
	s.Resource = v
	return s
}

type DescribeMonitorGroupsResponseBodyResourcesResource struct {
	Type                 *string                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	BindUrl              *string                                                          `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	ServiceId            *string                                                          `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ContactGroups        *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	Tags                 *DescribeMonitorGroupsResponseBodyResourcesResourceTags          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	GroupFounderTagKey   *string                                                          `json:"GroupFounderTagKey,omitempty" xml:"GroupFounderTagKey,omitempty"`
	TemplateIds          *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds   `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Struct"`
	GmtModified          *int64                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupFounderTagValue *string                                                          `json:"GroupFounderTagValue,omitempty" xml:"GroupFounderTagValue,omitempty"`
	GroupName            *string                                                          `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId              *int64                                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	DynamicTagRuleId     *string                                                          `json:"DynamicTagRuleId,omitempty" xml:"DynamicTagRuleId,omitempty"`
	GmtCreate            *int64                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetType(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.Type = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetBindUrl(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.BindUrl = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetServiceId(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.ServiceId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetContactGroups(v *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.ContactGroups = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetTags(v *DescribeMonitorGroupsResponseBodyResourcesResourceTags) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.Tags = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupFounderTagKey(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupFounderTagKey = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetTemplateIds(v *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.TemplateIds = v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGmtModified(v int64) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GmtModified = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupFounderTagValue(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupFounderTagValue = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupName(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupName = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGroupId(v int64) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetDynamicTagRuleId(v string) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.DynamicTagRuleId = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResource) SetGmtCreate(v int64) *DescribeMonitorGroupsResponseBodyResourcesResource {
	s.GmtCreate = &v
	return s
}

type DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups struct {
	ContactGroup []*DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups) SetContactGroup(v []*DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroups {
	s.ContactGroup = v
	return s
}

type DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup) SetName(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceContactGroupsContactGroup {
	s.Name = &v
	return s
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTags struct {
	Tag []*DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTags) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTags) SetTag(v []*DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) *DescribeMonitorGroupsResponseBodyResourcesResourceTags {
	s.Tag = v
	return s
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) SetKey(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag) SetValue(v string) *DescribeMonitorGroupsResponseBodyResourcesResourceTagsTag {
	s.Value = &v
	return s
}

type DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds struct {
	TemplateId []*string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds) SetTemplateId(v []*string) *DescribeMonitorGroupsResponseBodyResourcesResourceTemplateIds {
	s.TemplateId = v
	return s
}

type DescribeMonitorGroupsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupsResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupsResponse) SetBody(v *DescribeMonitorGroupsResponseBody) *DescribeMonitorGroupsResponse {
	s.Body = v
	return s
}

type DescribeMonitoringAgentAccessKeyResponseBody struct {
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentAccessKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetSecretKey(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.SecretKey = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetRequestId(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetMessage(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetAccessKey(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.AccessKey = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetCode(v int32) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitoringAgentAccessKeyResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitoringAgentAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitoringAgentAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentAccessKeyResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponse) SetBody(v *DescribeMonitoringAgentAccessKeyResponseBody) *DescribeMonitoringAgentAccessKeyResponse {
	s.Body = v
	return s
}

type DescribeMonitoringAgentConfigResponseBody struct {
	EnableActiveAlert        *string `json:"EnableActiveAlert,omitempty" xml:"EnableActiveAlert,omitempty"`
	AutoInstall              *bool   `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	EnableInstallAgentNewECS *bool   `json:"EnableInstallAgentNewECS,omitempty" xml:"EnableInstallAgentNewECS,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message                  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Code                     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success                  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetEnableActiveAlert(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.EnableActiveAlert = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetAutoInstall(v bool) *DescribeMonitoringAgentConfigResponseBody {
	s.AutoInstall = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetEnableInstallAgentNewECS(v bool) *DescribeMonitoringAgentConfigResponseBody {
	s.EnableInstallAgentNewECS = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetRequestId(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetMessage(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetCode(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentConfigResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitoringAgentConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitoringAgentConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitoringAgentConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentConfigResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentConfigResponse) SetBody(v *DescribeMonitoringAgentConfigResponseBody) *DescribeMonitoringAgentConfigResponse {
	s.Body = v
	return s
}

type DescribeMonitoringAgentHostsRequest struct {
	KeyWord          *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceIds      *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	SerialNumbers    *string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	AliyunHost       *bool   `json:"AliyunHost,omitempty" xml:"AliyunHost,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMonitoringAgentHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsRequest) SetKeyWord(v string) *DescribeMonitoringAgentHostsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetHostName(v string) *DescribeMonitoringAgentHostsRequest {
	s.HostName = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetInstanceIds(v string) *DescribeMonitoringAgentHostsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetSerialNumbers(v string) *DescribeMonitoringAgentHostsRequest {
	s.SerialNumbers = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetPageNumber(v int32) *DescribeMonitoringAgentHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetPageSize(v int32) *DescribeMonitoringAgentHostsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetInstanceRegionId(v string) *DescribeMonitoringAgentHostsRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetAliyunHost(v bool) *DescribeMonitoringAgentHostsRequest {
	s.AliyunHost = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetStatus(v string) *DescribeMonitoringAgentHostsRequest {
	s.Status = &v
	return s
}

type DescribeMonitoringAgentHostsResponseBody struct {
	Hosts      *DescribeMonitoringAgentHostsResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Struct"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageTotal  *int32                                         `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	Total      *int32                                         `json:"Total,omitempty" xml:"Total,omitempty"`
	Code       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetHosts(v *DescribeMonitoringAgentHostsResponseBodyHosts) *DescribeMonitoringAgentHostsResponseBody {
	s.Hosts = v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetRequestId(v string) *DescribeMonitoringAgentHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetMessage(v string) *DescribeMonitoringAgentHostsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetPageSize(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetPageNumber(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetPageTotal(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.PageTotal = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetTotal(v int32) *DescribeMonitoringAgentHostsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetCode(v string) *DescribeMonitoringAgentHostsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentHostsResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitoringAgentHostsResponseBodyHosts struct {
	Host []*DescribeMonitoringAgentHostsResponseBodyHostsHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s DescribeMonitoringAgentHostsResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponseBodyHosts) SetHost(v []*DescribeMonitoringAgentHostsResponseBodyHostsHost) *DescribeMonitoringAgentHostsResponseBodyHosts {
	s.Host = v
	return s
}

type DescribeMonitoringAgentHostsResponseBodyHostsHost struct {
	SerialNumber       *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	NatIp              *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	AliUid             *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkType        *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	IsAliyunHost       *bool   `json:"isAliyunHost,omitempty" xml:"isAliyunHost,omitempty"`
	EipAddress         *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	AgentVersion       *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	EipId              *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	IpGroup            *string `json:"IpGroup,omitempty" xml:"IpGroup,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	OperatingSystem    *string `json:"OperatingSystem,omitempty" xml:"OperatingSystem,omitempty"`
}

func (s DescribeMonitoringAgentHostsResponseBodyHostsHost) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponseBodyHostsHost) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetSerialNumber(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.SerialNumber = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetNatIp(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.NatIp = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetAliUid(v int64) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.AliUid = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetHostName(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.HostName = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetInstanceId(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetNetworkType(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.NetworkType = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetIsAliyunHost(v bool) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.IsAliyunHost = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetEipAddress(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.EipAddress = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetAgentVersion(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.AgentVersion = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetEipId(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.EipId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetIpGroup(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.IpGroup = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetRegion(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.Region = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetInstanceTypeFamily(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeMonitoringAgentHostsResponseBodyHostsHost) SetOperatingSystem(v string) *DescribeMonitoringAgentHostsResponseBodyHostsHost {
	s.OperatingSystem = &v
	return s
}

type DescribeMonitoringAgentHostsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitoringAgentHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitoringAgentHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentHostsResponse) SetBody(v *DescribeMonitoringAgentHostsResponseBody) *DescribeMonitoringAgentHostsResponse {
	s.Body = v
	return s
}

type DescribeMonitoringAgentProcessesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeMonitoringAgentProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesRequest) SetInstanceId(v string) *DescribeMonitoringAgentProcessesRequest {
	s.InstanceId = &v
	return s
}

type DescribeMonitoringAgentProcessesResponseBody struct {
	Message       *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NodeProcesses *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses `json:"NodeProcesses,omitempty" xml:"NodeProcesses,omitempty" type:"Struct"`
	Code          *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success       *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetMessage(v string) *DescribeMonitoringAgentProcessesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetRequestId(v string) *DescribeMonitoringAgentProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetNodeProcesses(v *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) *DescribeMonitoringAgentProcessesResponseBody {
	s.NodeProcesses = v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetCode(v string) *DescribeMonitoringAgentProcessesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentProcessesResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitoringAgentProcessesResponseBodyNodeProcesses struct {
	NodeProcess []*DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess `json:"NodeProcess,omitempty" xml:"NodeProcess,omitempty" type:"Repeated"`
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) SetNodeProcess(v []*DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses {
	s.NodeProcess = v
	return s
}

type DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess struct {
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	ProcessId   *int64  `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Command     *string `json:"Command,omitempty" xml:"Command,omitempty"`
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetProcessName(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessName = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetProcessId(v int64) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetGroupId(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetCommand(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.Command = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetProcessUser(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessUser = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetInstanceId(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.InstanceId = &v
	return s
}

type DescribeMonitoringAgentProcessesResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitoringAgentProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitoringAgentProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentProcessesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponse) SetBody(v *DescribeMonitoringAgentProcessesResponseBody) *DescribeMonitoringAgentProcessesResponse {
	s.Body = v
	return s
}

type DescribeMonitoringAgentStatusesRequest struct {
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DescribeMonitoringAgentStatusesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesRequest) SetInstanceIds(v string) *DescribeMonitoringAgentStatusesRequest {
	s.InstanceIds = &v
	return s
}

type DescribeMonitoringAgentStatusesResponseBody struct {
	Message        *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NodeStatusList *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList `json:"NodeStatusList,omitempty" xml:"NodeStatusList,omitempty" type:"Struct"`
	Code           *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentStatusesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetMessage(v string) *DescribeMonitoringAgentStatusesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetRequestId(v string) *DescribeMonitoringAgentStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetNodeStatusList(v *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) *DescribeMonitoringAgentStatusesResponseBody {
	s.NodeStatusList = v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetCode(v string) *DescribeMonitoringAgentStatusesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentStatusesResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitoringAgentStatusesResponseBodyNodeStatusList struct {
	NodeStatus []*DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty" type:"Repeated"`
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) SetNodeStatus(v []*DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList {
	s.NodeStatus = v
	return s
}

type DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	AutoInstall *bool   `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetStatus(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.Status = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetAutoInstall(v bool) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.AutoInstall = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetInstanceId(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.InstanceId = &v
	return s
}

type DescribeMonitoringAgentStatusesResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitoringAgentStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitoringAgentStatusesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponse) SetHeaders(v map[string]*string) *DescribeMonitoringAgentStatusesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponse) SetBody(v *DescribeMonitoringAgentStatusesResponseBody) *DescribeMonitoringAgentStatusesResponse {
	s.Body = v
	return s
}

type DescribeMonitoringConfigResponseBody struct {
	AutoInstall              *bool   `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	EnableInstallAgentNewECS *bool   `json:"EnableInstallAgentNewECS,omitempty" xml:"EnableInstallAgentNewECS,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message                  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Code                     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success                  *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringConfigResponseBody) SetAutoInstall(v bool) *DescribeMonitoringConfigResponseBody {
	s.AutoInstall = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetEnableInstallAgentNewECS(v bool) *DescribeMonitoringConfigResponseBody {
	s.EnableInstallAgentNewECS = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetRequestId(v string) *DescribeMonitoringConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetMessage(v string) *DescribeMonitoringConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetCode(v string) *DescribeMonitoringConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetSuccess(v bool) *DescribeMonitoringConfigResponseBody {
	s.Success = &v
	return s
}

type DescribeMonitoringConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitoringConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitoringConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitoringConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringConfigResponse) SetHeaders(v map[string]*string) *DescribeMonitoringConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitoringConfigResponse) SetBody(v *DescribeMonitoringConfigResponseBody) *DescribeMonitoringConfigResponse {
	s.Body = v
	return s
}

type DescribeMonitorResourceQuotaAttributeRequest struct {
	ShowUsed *bool `json:"ShowUsed,omitempty" xml:"ShowUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeRequest) SetShowUsed(v bool) *DescribeMonitorResourceQuotaAttributeRequest {
	s.ShowUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBody struct {
	Message       *string                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceQuota *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota `json:"ResourceQuota,omitempty" xml:"ResourceQuota,omitempty" type:"Struct"`
	Code          *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetMessage(v string) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetRequestId(v string) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetResourceQuota(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.ResourceQuota = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBody) SetCode(v string) *DescribeMonitorResourceQuotaAttributeResponseBody {
	s.Code = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota struct {
	Api                      *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi                      `json:"Api,omitempty" xml:"Api,omitempty" type:"Struct"`
	ExpireTime               *string                                                                                 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CustomMonitor            *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor            `json:"CustomMonitor,omitempty" xml:"CustomMonitor,omitempty" type:"Struct"`
	EventMonitor             *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor             `json:"EventMonitor,omitempty" xml:"EventMonitor,omitempty" type:"Struct"`
	InstanceId               *string                                                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SiteMonitorTask          *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask          `json:"SiteMonitorTask,omitempty" xml:"SiteMonitorTask,omitempty" type:"Struct"`
	Phone                    *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone                    `json:"Phone,omitempty" xml:"Phone,omitempty" type:"Struct"`
	SuitInfo                 *string                                                                                 `json:"SuitInfo,omitempty" xml:"SuitInfo,omitempty"`
	SMS                      *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS                      `json:"SMS,omitempty" xml:"SMS,omitempty" type:"Struct"`
	LogMonitor               *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor               `json:"LogMonitor,omitempty" xml:"LogMonitor,omitempty" type:"Struct"`
	SiteMonitorOperatorProbe *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe `json:"SiteMonitorOperatorProbe,omitempty" xml:"SiteMonitorOperatorProbe,omitempty" type:"Struct"`
	SiteMonitorEcsProbe      *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe      `json:"SiteMonitorEcsProbe,omitempty" xml:"SiteMonitorEcsProbe,omitempty" type:"Struct"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetApi(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.Api = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetExpireTime(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.ExpireTime = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetCustomMonitor(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.CustomMonitor = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetEventMonitor(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.EventMonitor = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetInstanceId(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorTask(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorTask = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetPhone(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.Phone = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSuitInfo(v string) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SuitInfo = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSMS(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SMS = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetLogMonitor(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.LogMonitor = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorOperatorProbe(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorOperatorProbe = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota) SetSiteMonitorEcsProbe(v *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuota {
	s.SiteMonitorEcsProbe = v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaApi {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaCustomMonitor {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaEventMonitor {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorTask {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaPhone {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSMS {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaLogMonitor {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorOperatorProbe {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe struct {
	QuotaLimit   *int32 `json:"QuotaLimit,omitempty" xml:"QuotaLimit,omitempty"`
	QuotaPackage *int32 `json:"QuotaPackage,omitempty" xml:"QuotaPackage,omitempty"`
	QuotaUsed    *int32 `json:"QuotaUsed,omitempty" xml:"QuotaUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) SetQuotaLimit(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe {
	s.QuotaLimit = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) SetQuotaPackage(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe {
	s.QuotaPackage = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe) SetQuotaUsed(v int32) *DescribeMonitorResourceQuotaAttributeResponseBodyResourceQuotaSiteMonitorEcsProbe {
	s.QuotaUsed = &v
	return s
}

type DescribeMonitorResourceQuotaAttributeResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMonitorResourceQuotaAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMonitorResourceQuotaAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) SetHeaders(v map[string]*string) *DescribeMonitorResourceQuotaAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeResponse) SetBody(v *DescribeMonitorResourceQuotaAttributeResponseBody) *DescribeMonitorResourceQuotaAttributeResponse {
	s.Body = v
	return s
}

type DescribeProductResourceTagKeyListRequest struct {
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeProductResourceTagKeyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResourceTagKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListRequest) SetNextToken(v string) *DescribeProductResourceTagKeyListRequest {
	s.NextToken = &v
	return s
}

type DescribeProductResourceTagKeyListResponseBody struct {
	Message   *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TagKeys   *DescribeProductResourceTagKeyListResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s DescribeProductResourceTagKeyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResourceTagKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetMessage(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetNextToken(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetRequestId(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetCode(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetSuccess(v bool) *DescribeProductResourceTagKeyListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetTagKeys(v *DescribeProductResourceTagKeyListResponseBodyTagKeys) *DescribeProductResourceTagKeyListResponseBody {
	s.TagKeys = v
	return s
}

type DescribeProductResourceTagKeyListResponseBodyTagKeys struct {
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DescribeProductResourceTagKeyListResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResourceTagKeyListResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListResponseBodyTagKeys) SetTagKey(v []*string) *DescribeProductResourceTagKeyListResponseBodyTagKeys {
	s.TagKey = v
	return s
}

type DescribeProductResourceTagKeyListResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProductResourceTagKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProductResourceTagKeyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductResourceTagKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListResponse) SetHeaders(v map[string]*string) *DescribeProductResourceTagKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductResourceTagKeyListResponse) SetBody(v *DescribeProductResourceTagKeyListResponseBody) *DescribeProductResourceTagKeyListResponse {
	s.Body = v
	return s
}

type DescribeProductsOfActiveMetricRuleResponseBody struct {
	Message                      *string                                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                    *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AllProductInitMetricRuleList *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList `json:"AllProductInitMetricRuleList,omitempty" xml:"AllProductInitMetricRuleList,omitempty" type:"Struct"`
	Datapoints                   *string                                                                     `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	Code                         *int32                                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success                      *bool                                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetMessage(v string) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetRequestId(v string) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetAllProductInitMetricRuleList(v *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.AllProductInitMetricRuleList = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetDatapoints(v string) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetCode(v int32) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetSuccess(v bool) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Success = &v
	return s
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList struct {
	AllProductInitMetricRule []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule `json:"AllProductInitMetricRule,omitempty" xml:"AllProductInitMetricRule,omitempty" type:"Repeated"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) SetAllProductInitMetricRule(v []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList {
	s.AllProductInitMetricRule = v
	return s
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule struct {
	Product             *string                                                                                                                `json:"Product,omitempty" xml:"Product,omitempty"`
	AlertInitConfigList *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList `json:"AlertInitConfigList,omitempty" xml:"AlertInitConfigList,omitempty" type:"Struct"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) SetProduct(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule {
	s.Product = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) SetAlertInitConfigList(v *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule {
	s.AlertInitConfigList = v
	return s
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList struct {
	AlertInitConfig []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig `json:"AlertInitConfig,omitempty" xml:"AlertInitConfig,omitempty" type:"Repeated"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) SetAlertInitConfig(v []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList {
	s.AlertInitConfig = v
	return s
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig struct {
	MetricName      *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	EvaluationCount *string `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Namespace       *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Threshold       *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics      *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Period          *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetMetricName(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.MetricName = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetEvaluationCount(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetNamespace(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Namespace = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetThreshold(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetStatistics(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Statistics = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetPeriod(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Period = &v
	return s
}

type DescribeProductsOfActiveMetricRuleResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProductsOfActiveMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProductsOfActiveMetricRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponse) SetHeaders(v map[string]*string) *DescribeProductsOfActiveMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponse) SetBody(v *DescribeProductsOfActiveMetricRuleResponseBody) *DescribeProductsOfActiveMetricRuleResponse {
	s.Body = v
	return s
}

type DescribeProjectMetaRequest struct {
	Labels     *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeProjectMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaRequest) SetLabels(v string) *DescribeProjectMetaRequest {
	s.Labels = &v
	return s
}

func (s *DescribeProjectMetaRequest) SetPageNumber(v int32) *DescribeProjectMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProjectMetaRequest) SetPageSize(v int32) *DescribeProjectMetaRequest {
	s.PageSize = &v
	return s
}

type DescribeProjectMetaResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *string                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Total      *string                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Resources  *DescribeProjectMetaResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProjectMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBody) SetRequestId(v string) *DescribeProjectMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetMessage(v string) *DescribeProjectMetaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetPageSize(v string) *DescribeProjectMetaResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetPageNumber(v string) *DescribeProjectMetaResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetTotal(v string) *DescribeProjectMetaResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetResources(v *DescribeProjectMetaResponseBodyResources) *DescribeProjectMetaResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetCode(v string) *DescribeProjectMetaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProjectMetaResponseBody) SetSuccess(v bool) *DescribeProjectMetaResponseBody {
	s.Success = &v
	return s
}

type DescribeProjectMetaResponseBodyResources struct {
	Resource []*DescribeProjectMetaResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeProjectMetaResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMetaResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBodyResources) SetResource(v []*DescribeProjectMetaResponseBodyResourcesResource) *DescribeProjectMetaResponseBodyResources {
	s.Resource = v
	return s
}

type DescribeProjectMetaResponseBodyResourcesResource struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeProjectMetaResponseBodyResourcesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMetaResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetDescription(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Description = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetLabels(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Labels = &v
	return s
}

func (s *DescribeProjectMetaResponseBodyResourcesResource) SetNamespace(v string) *DescribeProjectMetaResponseBodyResourcesResource {
	s.Namespace = &v
	return s
}

type DescribeProjectMetaResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProjectMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaResponse) SetHeaders(v map[string]*string) *DescribeProjectMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectMetaResponse) SetBody(v *DescribeProjectMetaResponseBody) *DescribeProjectMetaResponse {
	s.Body = v
	return s
}

type DescribeSiteMonitorAttributeRequest struct {
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	IncludeAlert *bool   `json:"IncludeAlert,omitempty" xml:"IncludeAlert,omitempty"`
}

func (s DescribeSiteMonitorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeRequest) SetTaskId(v string) *DescribeSiteMonitorAttributeRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeRequest) SetIncludeAlert(v bool) *DescribeSiteMonitorAttributeRequest {
	s.IncludeAlert = &v
	return s
}

type DescribeSiteMonitorAttributeResponseBody struct {
	MetricRules  *DescribeSiteMonitorAttributeResponseBodyMetricRules  `json:"MetricRules,omitempty" xml:"MetricRules,omitempty" type:"Struct"`
	Message      *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code         *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	SiteMonitors *DescribeSiteMonitorAttributeResponseBodySiteMonitors `json:"SiteMonitors,omitempty" xml:"SiteMonitors,omitempty" type:"Struct"`
}

func (s DescribeSiteMonitorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetMetricRules(v *DescribeSiteMonitorAttributeResponseBodyMetricRules) *DescribeSiteMonitorAttributeResponseBody {
	s.MetricRules = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetMessage(v string) *DescribeSiteMonitorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetRequestId(v string) *DescribeSiteMonitorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetCode(v string) *DescribeSiteMonitorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetSuccess(v bool) *DescribeSiteMonitorAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBody) SetSiteMonitors(v *DescribeSiteMonitorAttributeResponseBodySiteMonitors) *DescribeSiteMonitorAttributeResponseBody {
	s.SiteMonitors = v
	return s
}

type DescribeSiteMonitorAttributeResponseBodyMetricRules struct {
	MetricRule []*DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule `json:"MetricRule,omitempty" xml:"MetricRule,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRules) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRules) SetMetricRule(v []*DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) *DescribeSiteMonitorAttributeResponseBodyMetricRules {
	s.MetricRule = v
	return s
}

type DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule struct {
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	EvaluationCount    *string `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OkActions          *string `json:"OkActions,omitempty" xml:"OkActions,omitempty"`
	AlarmActions       *string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId             *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Expression         *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Dimensions         *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	StateValue         *string `json:"StateValue,omitempty" xml:"StateValue,omitempty"`
	ActionEnable       *string `json:"ActionEnable,omitempty" xml:"ActionEnable,omitempty"`
	Level              *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetMetricName(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.MetricName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetEvaluationCount(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetNamespace(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Namespace = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetOkActions(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.OkActions = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetAlarmActions(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.AlarmActions = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetPeriod(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Period = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetRuleName(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.RuleName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetRuleId(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.RuleId = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetComparisonOperator(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetExpression(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Expression = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetDimensions(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Dimensions = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetStateValue(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.StateValue = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetActionEnable(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.ActionEnable = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetLevel(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Level = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetThreshold(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Threshold = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule) SetStatistics(v string) *DescribeSiteMonitorAttributeResponseBodyMetricRulesMetricRule {
	s.Statistics = &v
	return s
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitors struct {
	TaskType   *string                                                         `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	OptionJson *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson `json:"OptionJson,omitempty" xml:"OptionJson,omitempty" type:"Struct"`
	Interval   *string                                                         `json:"Interval,omitempty" xml:"Interval,omitempty"`
	TaskState  *string                                                         `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	TaskName   *string                                                         `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Address    *string                                                         `json:"Address,omitempty" xml:"Address,omitempty"`
	IspCities  *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities  `json:"IspCities,omitempty" xml:"IspCities,omitempty" type:"Struct"`
	TaskId     *string                                                         `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitors) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitors) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetOptionJson(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.OptionJson = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetInterval(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.Interval = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskState(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskState = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetAddress(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.Address = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetIspCities(v *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.IspCities = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitors) SetTaskId(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitors {
	s.TaskId = &v
	return s
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson struct {
	Password        *string  `json:"password,omitempty" xml:"password,omitempty"`
	RequestFormat   *string  `json:"request_format,omitempty" xml:"request_format,omitempty"`
	ExpectValue     *string  `json:"expect_value,omitempty" xml:"expect_value,omitempty"`
	ResponseContent *string  `json:"response_content,omitempty" xml:"response_content,omitempty"`
	TimeOut         *int64   `json:"time_out,omitempty" xml:"time_out,omitempty"`
	FailureRate     *float32 `json:"failure_rate,omitempty" xml:"failure_rate,omitempty"`
	Header          *string  `json:"header,omitempty" xml:"header,omitempty"`
	Cookie          *string  `json:"cookie,omitempty" xml:"cookie,omitempty"`
	PingNum         *int32   `json:"ping_num,omitempty" xml:"ping_num,omitempty"`
	Port            *int32   `json:"port,omitempty" xml:"port,omitempty"`
	Authentication  *int32   `json:"authentication,omitempty" xml:"authentication,omitempty"`
	HttpMethod      *string  `json:"http_method,omitempty" xml:"http_method,omitempty"`
	MatchRule       *int32   `json:"match_rule,omitempty" xml:"match_rule,omitempty"`
	RequestContent  *string  `json:"request_content,omitempty" xml:"request_content,omitempty"`
	Username        *string  `json:"username,omitempty" xml:"username,omitempty"`
	Traceroute      *int64   `json:"traceroute,omitempty" xml:"traceroute,omitempty"`
	ResponseFormat  *string  `json:"response_format,omitempty" xml:"response_format,omitempty"`
	DnsType         *string  `json:"dns_type,omitempty" xml:"dns_type,omitempty"`
	DnsServer       *string  `json:"dns_server,omitempty" xml:"dns_server,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPassword(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Password = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetRequestFormat(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.RequestFormat = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetExpectValue(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ExpectValue = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetResponseContent(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ResponseContent = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTimeOut(v int64) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.TimeOut = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetFailureRate(v float32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.FailureRate = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetHeader(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Header = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetCookie(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Cookie = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPingNum(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.PingNum = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetPort(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Port = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetAuthentication(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Authentication = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetHttpMethod(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.HttpMethod = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetMatchRule(v int32) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.MatchRule = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetRequestContent(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.RequestContent = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetUsername(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Username = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetTraceroute(v int64) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.Traceroute = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetResponseFormat(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.ResponseFormat = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDnsType(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DnsType = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson) SetDnsServer(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsOptionJson {
	s.DnsServer = &v
	return s
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities struct {
	IspCity []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity `json:"IspCity,omitempty" xml:"IspCity,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities) SetIspCity(v []*DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCities {
	s.IspCity = v
	return s
}

type DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity struct {
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	IspName  *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetCityName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.CityName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetCity(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.City = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetIspName(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.IspName = &v
	return s
}

func (s *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity) SetIsp(v string) *DescribeSiteMonitorAttributeResponseBodySiteMonitorsIspCitiesIspCity {
	s.Isp = &v
	return s
}

type DescribeSiteMonitorAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSiteMonitorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSiteMonitorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorAttributeResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorAttributeResponse) SetBody(v *DescribeSiteMonitorAttributeResponseBody) *DescribeSiteMonitorAttributeResponse {
	s.Body = v
	return s
}

type DescribeSiteMonitorDataRequest struct {
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Length     *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
}

func (s DescribeSiteMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorDataRequest) SetTaskId(v string) *DescribeSiteMonitorDataRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetType(v string) *DescribeSiteMonitorDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetMetricName(v string) *DescribeSiteMonitorDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetStartTime(v string) *DescribeSiteMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetEndTime(v string) *DescribeSiteMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetPeriod(v string) *DescribeSiteMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetNextToken(v string) *DescribeSiteMonitorDataRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetLength(v int32) *DescribeSiteMonitorDataRequest {
	s.Length = &v
	return s
}

type DescribeSiteMonitorDataResponseBody struct {
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorDataResponseBody) SetNextToken(v string) *DescribeSiteMonitorDataResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetRequestId(v string) *DescribeSiteMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetMessage(v string) *DescribeSiteMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetData(v string) *DescribeSiteMonitorDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetCode(v string) *DescribeSiteMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetSuccess(v string) *DescribeSiteMonitorDataResponseBody {
	s.Success = &v
	return s
}

type DescribeSiteMonitorDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSiteMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSiteMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorDataResponse) SetBody(v *DescribeSiteMonitorDataResponseBody) *DescribeSiteMonitorDataResponse {
	s.Body = v
	return s
}

type DescribeSiteMonitorListRequest struct {
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Page     *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSiteMonitorListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListRequest) SetTaskId(v string) *DescribeSiteMonitorListRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetTaskType(v string) *DescribeSiteMonitorListRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetKeyword(v string) *DescribeSiteMonitorListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetPage(v int32) *DescribeSiteMonitorListRequest {
	s.Page = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetPageSize(v int32) *DescribeSiteMonitorListRequest {
	s.PageSize = &v
	return s
}

type DescribeSiteMonitorListResponseBody struct {
	TotalCount   *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message      *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize     *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Code         *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *string                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	SiteMonitors *DescribeSiteMonitorListResponseBodySiteMonitors `json:"SiteMonitors,omitempty" xml:"SiteMonitors,omitempty" type:"Struct"`
}

func (s DescribeSiteMonitorListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBody) SetTotalCount(v int32) *DescribeSiteMonitorListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetRequestId(v string) *DescribeSiteMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetMessage(v string) *DescribeSiteMonitorListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetPageSize(v int32) *DescribeSiteMonitorListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetPageNumber(v int32) *DescribeSiteMonitorListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetCode(v string) *DescribeSiteMonitorListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetSuccess(v string) *DescribeSiteMonitorListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBody) SetSiteMonitors(v *DescribeSiteMonitorListResponseBodySiteMonitors) *DescribeSiteMonitorListResponseBody {
	s.SiteMonitors = v
	return s
}

type DescribeSiteMonitorListResponseBodySiteMonitors struct {
	SiteMonitor []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor `json:"SiteMonitor,omitempty" xml:"SiteMonitor,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitors) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitors) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitors) SetSiteMonitor(v []*DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) *DescribeSiteMonitorListResponseBodySiteMonitors {
	s.SiteMonitor = v
	return s
}

type DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor struct {
	TaskType    *string                                                                `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UpdateTime  *string                                                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Interval    *string                                                                `json:"Interval,omitempty" xml:"Interval,omitempty"`
	TaskState   *string                                                                `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	OptionsJson *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty" type:"Struct"`
	CreateTime  *string                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TaskName    *string                                                                `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Address     *string                                                                `json:"Address,omitempty" xml:"Address,omitempty"`
	TaskId      *string                                                                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskType(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskType = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetUpdateTime(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetInterval(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.Interval = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskState(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskState = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetOptionsJson(v *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.OptionsJson = v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetCreateTime(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.CreateTime = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskName(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskName = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetAddress(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.Address = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor) SetTaskId(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitor {
	s.TaskId = &v
	return s
}

type DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson struct {
	Password        *string  `json:"password,omitempty" xml:"password,omitempty"`
	RequestFormat   *string  `json:"request_format,omitempty" xml:"request_format,omitempty"`
	ExpectValue     *string  `json:"expect_value,omitempty" xml:"expect_value,omitempty"`
	ResponseContent *string  `json:"response_content,omitempty" xml:"response_content,omitempty"`
	TimeOut         *int64   `json:"time_out,omitempty" xml:"time_out,omitempty"`
	FailureRate     *float32 `json:"failure_rate,omitempty" xml:"failure_rate,omitempty"`
	Header          *string  `json:"header,omitempty" xml:"header,omitempty"`
	Cookie          *string  `json:"cookie,omitempty" xml:"cookie,omitempty"`
	PingNum         *int32   `json:"ping_num,omitempty" xml:"ping_num,omitempty"`
	Port            *int32   `json:"port,omitempty" xml:"port,omitempty"`
	Authentication  *int32   `json:"authentication,omitempty" xml:"authentication,omitempty"`
	HttpMethod      *string  `json:"http_method,omitempty" xml:"http_method,omitempty"`
	MatchRule       *int32   `json:"match_rule,omitempty" xml:"match_rule,omitempty"`
	RequestContent  *string  `json:"request_content,omitempty" xml:"request_content,omitempty"`
	Username        *string  `json:"username,omitempty" xml:"username,omitempty"`
	Traceroute      *int64   `json:"traceroute,omitempty" xml:"traceroute,omitempty"`
	ResponseFormat  *string  `json:"response_format,omitempty" xml:"response_format,omitempty"`
	DnsType         *string  `json:"dns_type,omitempty" xml:"dns_type,omitempty"`
	DnsServer       *string  `json:"dns_server,omitempty" xml:"dns_server,omitempty"`
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetPassword(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Password = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetRequestFormat(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.RequestFormat = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetExpectValue(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.ExpectValue = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetResponseContent(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.ResponseContent = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetTimeOut(v int64) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.TimeOut = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetFailureRate(v float32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.FailureRate = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetHeader(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Header = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetCookie(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Cookie = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetPingNum(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.PingNum = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetPort(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Port = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetAuthentication(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Authentication = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetHttpMethod(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.HttpMethod = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetMatchRule(v int32) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.MatchRule = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetRequestContent(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.RequestContent = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetUsername(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Username = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetTraceroute(v int64) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.Traceroute = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetResponseFormat(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.ResponseFormat = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetDnsType(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.DnsType = &v
	return s
}

func (s *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson) SetDnsServer(v string) *DescribeSiteMonitorListResponseBodySiteMonitorsSiteMonitorOptionsJson {
	s.DnsServer = &v
	return s
}

type DescribeSiteMonitorListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSiteMonitorListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSiteMonitorListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorListResponse) SetBody(v *DescribeSiteMonitorListResponseBody) *DescribeSiteMonitorListResponse {
	s.Body = v
	return s
}

type DescribeSiteMonitorQuotaResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeSiteMonitorQuotaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetMessage(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetRequestId(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetData(v *DescribeSiteMonitorQuotaResponseBodyData) *DescribeSiteMonitorQuotaResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetCode(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBody) SetSuccess(v string) *DescribeSiteMonitorQuotaResponseBody {
	s.Success = &v
	return s
}

type DescribeSiteMonitorQuotaResponseBodyData struct {
	SiteMonitorOperatorQuotaQuota *int32  `json:"SiteMonitorOperatorQuotaQuota,omitempty" xml:"SiteMonitorOperatorQuotaQuota,omitempty"`
	SecondMonitor                 *bool   `json:"SecondMonitor,omitempty" xml:"SecondMonitor,omitempty"`
	SiteMonitorQuotaTaskUsed      *int32  `json:"SiteMonitorQuotaTaskUsed,omitempty" xml:"SiteMonitorQuotaTaskUsed,omitempty"`
	SiteMonitorTaskQuota          *int32  `json:"SiteMonitorTaskQuota,omitempty" xml:"SiteMonitorTaskQuota,omitempty"`
	SiteMonitorVersion            *string `json:"SiteMonitorVersion,omitempty" xml:"SiteMonitorVersion,omitempty"`
	SiteMonitorIdcQuota           *int32  `json:"SiteMonitorIdcQuota,omitempty" xml:"SiteMonitorIdcQuota,omitempty"`
}

func (s DescribeSiteMonitorQuotaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorOperatorQuotaQuota(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorOperatorQuotaQuota = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSecondMonitor(v bool) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SecondMonitor = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorQuotaTaskUsed(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorQuotaTaskUsed = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorTaskQuota(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorTaskQuota = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorVersion(v string) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorVersion = &v
	return s
}

func (s *DescribeSiteMonitorQuotaResponseBodyData) SetSiteMonitorIdcQuota(v int32) *DescribeSiteMonitorQuotaResponseBodyData {
	s.SiteMonitorIdcQuota = &v
	return s
}

type DescribeSiteMonitorQuotaResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSiteMonitorQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSiteMonitorQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorQuotaResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorQuotaResponse) SetBody(v *DescribeSiteMonitorQuotaResponseBody) *DescribeSiteMonitorQuotaResponse {
	s.Body = v
	return s
}

type DescribeSiteMonitorStatisticsRequest struct {
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TimeRange  *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
}

func (s DescribeSiteMonitorStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorStatisticsRequest) SetTaskId(v string) *DescribeSiteMonitorStatisticsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) SetTimeRange(v string) *DescribeSiteMonitorStatisticsRequest {
	s.TimeRange = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) SetStartTime(v string) *DescribeSiteMonitorStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsRequest) SetMetricName(v string) *DescribeSiteMonitorStatisticsRequest {
	s.MetricName = &v
	return s
}

type DescribeSiteMonitorStatisticsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetMessage(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetRequestId(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetData(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetCode(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponseBody) SetSuccess(v string) *DescribeSiteMonitorStatisticsResponseBody {
	s.Success = &v
	return s
}

type DescribeSiteMonitorStatisticsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSiteMonitorStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSiteMonitorStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSiteMonitorStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorStatisticsResponse) SetHeaders(v map[string]*string) *DescribeSiteMonitorStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteMonitorStatisticsResponse) SetBody(v *DescribeSiteMonitorStatisticsResponseBody) *DescribeSiteMonitorStatisticsResponse {
	s.Body = v
	return s
}

type DescribeSystemEventAttributeRequest struct {
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSystemEventAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeRequest) SetProduct(v string) *DescribeSystemEventAttributeRequest {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetEventType(v string) *DescribeSystemEventAttributeRequest {
	s.EventType = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetName(v string) *DescribeSystemEventAttributeRequest {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetLevel(v string) *DescribeSystemEventAttributeRequest {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetStatus(v string) *DescribeSystemEventAttributeRequest {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetGroupId(v string) *DescribeSystemEventAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetSearchKeywords(v string) *DescribeSystemEventAttributeRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetStartTime(v string) *DescribeSystemEventAttributeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetEndTime(v string) *DescribeSystemEventAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetPageNumber(v int32) *DescribeSystemEventAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSystemEventAttributeRequest) SetPageSize(v int32) *DescribeSystemEventAttributeRequest {
	s.PageSize = &v
	return s
}

type DescribeSystemEventAttributeResponseBody struct {
	SystemEvents *DescribeSystemEventAttributeResponseBodySystemEvents `json:"SystemEvents,omitempty" xml:"SystemEvents,omitempty" type:"Struct"`
	Message      *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code         *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *string                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSystemEventAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponseBody) SetSystemEvents(v *DescribeSystemEventAttributeResponseBodySystemEvents) *DescribeSystemEventAttributeResponseBody {
	s.SystemEvents = v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetMessage(v string) *DescribeSystemEventAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetRequestId(v string) *DescribeSystemEventAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetCode(v string) *DescribeSystemEventAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBody) SetSuccess(v string) *DescribeSystemEventAttributeResponseBody {
	s.Success = &v
	return s
}

type DescribeSystemEventAttributeResponseBodySystemEvents struct {
	SystemEvent []*DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent `json:"SystemEvent,omitempty" xml:"SystemEvent,omitempty" type:"Repeated"`
}

func (s DescribeSystemEventAttributeResponseBodySystemEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventAttributeResponseBodySystemEvents) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponseBodySystemEvents) SetSystemEvent(v []*DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) *DescribeSystemEventAttributeResponseBodySystemEvents {
	s.SystemEvent = v
	return s
}

type DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Time         *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetStatus(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetTime(v int64) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Time = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetGroupId(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetProduct(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetInstanceName(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.InstanceName = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetResourceId(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.ResourceId = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetName(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetContent(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Content = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetLevel(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent) SetRegionId(v string) *DescribeSystemEventAttributeResponseBodySystemEventsSystemEvent {
	s.RegionId = &v
	return s
}

type DescribeSystemEventAttributeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSystemEventAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSystemEventAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventAttributeResponse) SetHeaders(v map[string]*string) *DescribeSystemEventAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemEventAttributeResponse) SetBody(v *DescribeSystemEventAttributeResponseBody) *DescribeSystemEventAttributeResponse {
	s.Body = v
	return s
}

type DescribeSystemEventCountRequest struct {
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeSystemEventCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountRequest) SetProduct(v string) *DescribeSystemEventCountRequest {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetEventType(v string) *DescribeSystemEventCountRequest {
	s.EventType = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetName(v string) *DescribeSystemEventCountRequest {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetLevel(v string) *DescribeSystemEventCountRequest {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetStatus(v string) *DescribeSystemEventCountRequest {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetGroupId(v string) *DescribeSystemEventCountRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetSearchKeywords(v string) *DescribeSystemEventCountRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetStartTime(v string) *DescribeSystemEventCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventCountRequest) SetEndTime(v string) *DescribeSystemEventCountRequest {
	s.EndTime = &v
	return s
}

type DescribeSystemEventCountResponseBody struct {
	Message           *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemEventCounts *DescribeSystemEventCountResponseBodySystemEventCounts `json:"SystemEventCounts,omitempty" xml:"SystemEventCounts,omitempty" type:"Struct"`
	Code              *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success           *string                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSystemEventCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponseBody) SetMessage(v string) *DescribeSystemEventCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetRequestId(v string) *DescribeSystemEventCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetSystemEventCounts(v *DescribeSystemEventCountResponseBodySystemEventCounts) *DescribeSystemEventCountResponseBody {
	s.SystemEventCounts = v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetCode(v string) *DescribeSystemEventCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSystemEventCountResponseBody) SetSuccess(v string) *DescribeSystemEventCountResponseBody {
	s.Success = &v
	return s
}

type DescribeSystemEventCountResponseBodySystemEventCounts struct {
	SystemEventCount []*DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount `json:"SystemEventCount,omitempty" xml:"SystemEventCount,omitempty" type:"Repeated"`
}

func (s DescribeSystemEventCountResponseBodySystemEventCounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventCountResponseBodySystemEventCounts) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponseBodySystemEventCounts) SetSystemEventCount(v []*DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) *DescribeSystemEventCountResponseBodySystemEventCounts {
	s.SystemEventCount = v
	return s
}

type DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Time         *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Num          *int64  `json:"Num,omitempty" xml:"Num,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Level        *string `json:"Level,omitempty" xml:"Level,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetStatus(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetTime(v int64) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Time = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetGroupId(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetProduct(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetInstanceName(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.InstanceName = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetNum(v int64) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Num = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetResourceId(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.ResourceId = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetName(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetContent(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Content = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetLevel(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount) SetRegionId(v string) *DescribeSystemEventCountResponseBodySystemEventCountsSystemEventCount {
	s.RegionId = &v
	return s
}

type DescribeSystemEventCountResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSystemEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSystemEventCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventCountResponse) SetHeaders(v map[string]*string) *DescribeSystemEventCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemEventCountResponse) SetBody(v *DescribeSystemEventCountResponseBody) *DescribeSystemEventCountResponse {
	s.Body = v
	return s
}

type DescribeSystemEventHistogramRequest struct {
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	EventType      *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SearchKeywords *string `json:"SearchKeywords,omitempty" xml:"SearchKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeSystemEventHistogramRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventHistogramRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramRequest) SetProduct(v string) *DescribeSystemEventHistogramRequest {
	s.Product = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetEventType(v string) *DescribeSystemEventHistogramRequest {
	s.EventType = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetName(v string) *DescribeSystemEventHistogramRequest {
	s.Name = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetLevel(v string) *DescribeSystemEventHistogramRequest {
	s.Level = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetStatus(v string) *DescribeSystemEventHistogramRequest {
	s.Status = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetGroupId(v string) *DescribeSystemEventHistogramRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetSearchKeywords(v string) *DescribeSystemEventHistogramRequest {
	s.SearchKeywords = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetStartTime(v string) *DescribeSystemEventHistogramRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventHistogramRequest) SetEndTime(v string) *DescribeSystemEventHistogramRequest {
	s.EndTime = &v
	return s
}

type DescribeSystemEventHistogramResponseBody struct {
	Message               *string                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId             *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemEventHistograms *DescribeSystemEventHistogramResponseBodySystemEventHistograms `json:"SystemEventHistograms,omitempty" xml:"SystemEventHistograms,omitempty" type:"Struct"`
	Code                  *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success               *string                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSystemEventHistogramResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventHistogramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponseBody) SetMessage(v string) *DescribeSystemEventHistogramResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetRequestId(v string) *DescribeSystemEventHistogramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetSystemEventHistograms(v *DescribeSystemEventHistogramResponseBodySystemEventHistograms) *DescribeSystemEventHistogramResponseBody {
	s.SystemEventHistograms = v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetCode(v string) *DescribeSystemEventHistogramResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBody) SetSuccess(v string) *DescribeSystemEventHistogramResponseBody {
	s.Success = &v
	return s
}

type DescribeSystemEventHistogramResponseBodySystemEventHistograms struct {
	SystemEventHistogram []*DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram `json:"SystemEventHistogram,omitempty" xml:"SystemEventHistogram,omitempty" type:"Repeated"`
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistograms) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistograms) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistograms) SetSystemEventHistogram(v []*DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) *DescribeSystemEventHistogramResponseBodySystemEventHistograms {
	s.SystemEventHistogram = v
	return s
}

type DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram struct {
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Count     *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) SetEndTime(v int64) *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram {
	s.EndTime = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) SetStartTime(v int64) *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram {
	s.StartTime = &v
	return s
}

func (s *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram) SetCount(v int64) *DescribeSystemEventHistogramResponseBodySystemEventHistogramsSystemEventHistogram {
	s.Count = &v
	return s
}

type DescribeSystemEventHistogramResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSystemEventHistogramResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSystemEventHistogramResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemEventHistogramResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemEventHistogramResponse) SetHeaders(v map[string]*string) *DescribeSystemEventHistogramResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemEventHistogramResponse) SetBody(v *DescribeSystemEventHistogramResponseBody) *DescribeSystemEventHistogramResponse {
	s.Body = v
	return s
}

type DescribeTagKeyListRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeTagKeyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListRequest) SetPageNumber(v int32) *DescribeTagKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagKeyListRequest) SetPageSize(v int32) *DescribeTagKeyListRequest {
	s.PageSize = &v
	return s
}

type DescribeTagKeyListResponseBody struct {
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TagKeys   *DescribeTagKeyListResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s DescribeTagKeyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListResponseBody) SetMessage(v string) *DescribeTagKeyListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetRequestId(v string) *DescribeTagKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetCode(v string) *DescribeTagKeyListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetSuccess(v bool) *DescribeTagKeyListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagKeyListResponseBody) SetTagKeys(v *DescribeTagKeyListResponseBodyTagKeys) *DescribeTagKeyListResponseBody {
	s.TagKeys = v
	return s
}

type DescribeTagKeyListResponseBodyTagKeys struct {
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DescribeTagKeyListResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeyListResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListResponseBodyTagKeys) SetTagKey(v []*string) *DescribeTagKeyListResponseBodyTagKeys {
	s.TagKey = v
	return s
}

type DescribeTagKeyListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTagKeyListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagKeyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagKeyListResponse) SetHeaders(v map[string]*string) *DescribeTagKeyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagKeyListResponse) SetBody(v *DescribeTagKeyListResponseBody) *DescribeTagKeyListResponse {
	s.Body = v
	return s
}

type DescribeTagValueListRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TagKey     *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeTagValueListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagValueListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListRequest) SetPageNumber(v int32) *DescribeTagValueListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagValueListRequest) SetPageSize(v int32) *DescribeTagValueListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagValueListRequest) SetTagKey(v string) *DescribeTagValueListRequest {
	s.TagKey = &v
	return s
}

type DescribeTagValueListResponseBody struct {
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TagValues *DescribeTagValueListResponseBodyTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Struct"`
}

func (s DescribeTagValueListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagValueListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListResponseBody) SetMessage(v string) *DescribeTagValueListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetRequestId(v string) *DescribeTagValueListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetCode(v string) *DescribeTagValueListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetSuccess(v bool) *DescribeTagValueListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagValueListResponseBody) SetTagValues(v *DescribeTagValueListResponseBodyTagValues) *DescribeTagValueListResponseBody {
	s.TagValues = v
	return s
}

type DescribeTagValueListResponseBodyTagValues struct {
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s DescribeTagValueListResponseBodyTagValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagValueListResponseBodyTagValues) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListResponseBodyTagValues) SetTagValue(v []*string) *DescribeTagValueListResponseBodyTagValues {
	s.TagValue = v
	return s
}

type DescribeTagValueListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTagValueListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTagValueListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagValueListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListResponse) SetHeaders(v map[string]*string) *DescribeTagValueListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagValueListResponse) SetBody(v *DescribeTagValueListResponseBody) *DescribeTagValueListResponse {
	s.Body = v
	return s
}

type DescribeUnhealthyHostAvailabilityRequest struct {
	Id []*int `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s DescribeUnhealthyHostAvailabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityRequest) SetId(v []*int) *DescribeUnhealthyHostAvailabilityRequest {
	s.Id = v
	return s
}

type DescribeUnhealthyHostAvailabilityResponseBody struct {
	Message       *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code          *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Success       *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	UnhealthyList *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList `json:"UnhealthyList,omitempty" xml:"UnhealthyList,omitempty" type:"Struct"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetMessage(v string) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetRequestId(v string) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetCode(v string) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetSuccess(v bool) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBody) SetUnhealthyList(v *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) *DescribeUnhealthyHostAvailabilityResponseBody {
	s.UnhealthyList = v
	return s
}

type DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList struct {
	NodeTaskInstance []*DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance `json:"NodeTaskInstance,omitempty" xml:"NodeTaskInstance,omitempty" type:"Repeated"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList) SetNodeTaskInstance(v []*DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyList {
	s.NodeTaskInstance = v
	return s
}

type DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance struct {
	InstanceList *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	Id           *int64                                                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) SetInstanceList(v *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance {
	s.InstanceList = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance) SetId(v int64) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstance {
	s.Id = &v
	return s
}

type DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList) SetString_(v []*string) *DescribeUnhealthyHostAvailabilityResponseBodyUnhealthyListNodeTaskInstanceInstanceList {
	s.String_ = v
	return s
}

type DescribeUnhealthyHostAvailabilityResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUnhealthyHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUnhealthyHostAvailabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityResponse) SetHeaders(v map[string]*string) *DescribeUnhealthyHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityResponse) SetBody(v *DescribeUnhealthyHostAvailabilityResponseBody) *DescribeUnhealthyHostAvailabilityResponse {
	s.Body = v
	return s
}

type DisableActiveMetricRuleRequest struct {
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DisableActiveMetricRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableActiveMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableActiveMetricRuleRequest) SetProduct(v string) *DisableActiveMetricRuleRequest {
	s.Product = &v
	return s
}

type DisableActiveMetricRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableActiveMetricRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableActiveMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableActiveMetricRuleResponseBody) SetMessage(v string) *DisableActiveMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableActiveMetricRuleResponseBody) SetRequestId(v string) *DisableActiveMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableActiveMetricRuleResponseBody) SetCode(v string) *DisableActiveMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableActiveMetricRuleResponseBody) SetSuccess(v bool) *DisableActiveMetricRuleResponseBody {
	s.Success = &v
	return s
}

type DisableActiveMetricRuleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableActiveMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableActiveMetricRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableActiveMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableActiveMetricRuleResponse) SetHeaders(v map[string]*string) *DisableActiveMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableActiveMetricRuleResponse) SetBody(v *DisableActiveMetricRuleResponseBody) *DisableActiveMetricRuleResponse {
	s.Body = v
	return s
}

type DisableEventRulesRequest struct {
	RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DisableEventRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableEventRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableEventRulesRequest) SetRuleNames(v []*string) *DisableEventRulesRequest {
	s.RuleNames = v
	return s
}

type DisableEventRulesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableEventRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableEventRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableEventRulesResponseBody) SetMessage(v string) *DisableEventRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableEventRulesResponseBody) SetRequestId(v string) *DisableEventRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableEventRulesResponseBody) SetCode(v string) *DisableEventRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableEventRulesResponseBody) SetSuccess(v bool) *DisableEventRulesResponseBody {
	s.Success = &v
	return s
}

type DisableEventRulesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableEventRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableEventRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableEventRulesResponse) SetHeaders(v map[string]*string) *DisableEventRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableEventRulesResponse) SetBody(v *DisableEventRulesResponseBody) *DisableEventRulesResponse {
	s.Body = v
	return s
}

type DisableHostAvailabilityRequest struct {
	Id []*int `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s DisableHostAvailabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *DisableHostAvailabilityRequest) SetId(v []*int) *DisableHostAvailabilityRequest {
	s.Id = v
	return s
}

type DisableHostAvailabilityResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableHostAvailabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DisableHostAvailabilityResponseBody) SetMessage(v string) *DisableHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *DisableHostAvailabilityResponseBody) SetRequestId(v string) *DisableHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableHostAvailabilityResponseBody) SetCode(v string) *DisableHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *DisableHostAvailabilityResponseBody) SetSuccess(v bool) *DisableHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

type DisableHostAvailabilityResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableHostAvailabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *DisableHostAvailabilityResponse) SetHeaders(v map[string]*string) *DisableHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *DisableHostAvailabilityResponse) SetBody(v *DisableHostAvailabilityResponseBody) *DisableHostAvailabilityResponse {
	s.Body = v
	return s
}

type DisableMetricRulesRequest struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DisableMetricRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *DisableMetricRulesRequest) SetRuleId(v []*string) *DisableMetricRulesRequest {
	s.RuleId = v
	return s
}

type DisableMetricRulesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableMetricRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMetricRulesResponseBody) SetMessage(v string) *DisableMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableMetricRulesResponseBody) SetRequestId(v string) *DisableMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableMetricRulesResponseBody) SetCode(v string) *DisableMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableMetricRulesResponseBody) SetSuccess(v bool) *DisableMetricRulesResponseBody {
	s.Success = &v
	return s
}

type DisableMetricRulesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableMetricRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableMetricRulesResponse) SetHeaders(v map[string]*string) *DisableMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableMetricRulesResponse) SetBody(v *DisableMetricRulesResponseBody) *DisableMetricRulesResponse {
	s.Body = v
	return s
}

type DisableSiteMonitorsRequest struct {
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DisableSiteMonitorsRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableSiteMonitorsRequest) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsRequest) SetTaskIds(v string) *DisableSiteMonitorsRequest {
	s.TaskIds = &v
	return s
}

type DisableSiteMonitorsResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DisableSiteMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSiteMonitorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSiteMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsResponseBody) SetMessage(v string) *DisableSiteMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetRequestId(v string) *DisableSiteMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetData(v *DisableSiteMonitorsResponseBodyData) *DisableSiteMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetCode(v string) *DisableSiteMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSiteMonitorsResponseBody) SetSuccess(v string) *DisableSiteMonitorsResponseBody {
	s.Success = &v
	return s
}

type DisableSiteMonitorsResponseBodyData struct {
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s DisableSiteMonitorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DisableSiteMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsResponseBodyData) SetCount(v int32) *DisableSiteMonitorsResponseBodyData {
	s.Count = &v
	return s
}

type DisableSiteMonitorsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableSiteMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSiteMonitorsResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSiteMonitorsResponse) GoString() string {
	return s.String()
}

func (s *DisableSiteMonitorsResponse) SetHeaders(v map[string]*string) *DisableSiteMonitorsResponse {
	s.Headers = v
	return s
}

func (s *DisableSiteMonitorsResponse) SetBody(v *DisableSiteMonitorsResponseBody) *DisableSiteMonitorsResponse {
	s.Body = v
	return s
}

type EnableActiveMetricRuleRequest struct {
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s EnableActiveMetricRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableActiveMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableActiveMetricRuleRequest) SetProduct(v string) *EnableActiveMetricRuleRequest {
	s.Product = &v
	return s
}

type EnableActiveMetricRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableActiveMetricRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableActiveMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableActiveMetricRuleResponseBody) SetMessage(v string) *EnableActiveMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableActiveMetricRuleResponseBody) SetRequestId(v string) *EnableActiveMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableActiveMetricRuleResponseBody) SetCode(v string) *EnableActiveMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableActiveMetricRuleResponseBody) SetSuccess(v bool) *EnableActiveMetricRuleResponseBody {
	s.Success = &v
	return s
}

type EnableActiveMetricRuleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableActiveMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableActiveMetricRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableActiveMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableActiveMetricRuleResponse) SetHeaders(v map[string]*string) *EnableActiveMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableActiveMetricRuleResponse) SetBody(v *EnableActiveMetricRuleResponseBody) *EnableActiveMetricRuleResponse {
	s.Body = v
	return s
}

type EnableEventRulesRequest struct {
	RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s EnableEventRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableEventRulesRequest) GoString() string {
	return s.String()
}

func (s *EnableEventRulesRequest) SetRuleNames(v []*string) *EnableEventRulesRequest {
	s.RuleNames = v
	return s
}

type EnableEventRulesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableEventRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableEventRulesResponseBody) GoString() string {
	return s.String()
}

func (s *EnableEventRulesResponseBody) SetMessage(v string) *EnableEventRulesResponseBody {
	s.Message = &v
	return s
}

func (s *EnableEventRulesResponseBody) SetRequestId(v string) *EnableEventRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableEventRulesResponseBody) SetCode(v string) *EnableEventRulesResponseBody {
	s.Code = &v
	return s
}

func (s *EnableEventRulesResponseBody) SetSuccess(v bool) *EnableEventRulesResponseBody {
	s.Success = &v
	return s
}

type EnableEventRulesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableEventRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableEventRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableEventRulesResponse) GoString() string {
	return s.String()
}

func (s *EnableEventRulesResponse) SetHeaders(v map[string]*string) *EnableEventRulesResponse {
	s.Headers = v
	return s
}

func (s *EnableEventRulesResponse) SetBody(v *EnableEventRulesResponseBody) *EnableEventRulesResponse {
	s.Body = v
	return s
}

type EnableHostAvailabilityRequest struct {
	Id []*int `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
}

func (s EnableHostAvailabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *EnableHostAvailabilityRequest) SetId(v []*int) *EnableHostAvailabilityRequest {
	s.Id = v
	return s
}

type EnableHostAvailabilityResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHostAvailabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHostAvailabilityResponseBody) SetMessage(v string) *EnableHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *EnableHostAvailabilityResponseBody) SetRequestId(v string) *EnableHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableHostAvailabilityResponseBody) SetCode(v string) *EnableHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *EnableHostAvailabilityResponseBody) SetSuccess(v bool) *EnableHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

type EnableHostAvailabilityResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableHostAvailabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *EnableHostAvailabilityResponse) SetHeaders(v map[string]*string) *EnableHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *EnableHostAvailabilityResponse) SetBody(v *EnableHostAvailabilityResponseBody) *EnableHostAvailabilityResponse {
	s.Body = v
	return s
}

type EnableMetricRulesRequest struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s EnableMetricRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *EnableMetricRulesRequest) SetRuleId(v []*string) *EnableMetricRulesRequest {
	s.RuleId = v
	return s
}

type EnableMetricRulesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableMetricRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *EnableMetricRulesResponseBody) SetMessage(v string) *EnableMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *EnableMetricRulesResponseBody) SetRequestId(v string) *EnableMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableMetricRulesResponseBody) SetCode(v string) *EnableMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *EnableMetricRulesResponseBody) SetSuccess(v bool) *EnableMetricRulesResponseBody {
	s.Success = &v
	return s
}

type EnableMetricRulesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableMetricRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *EnableMetricRulesResponse) SetHeaders(v map[string]*string) *EnableMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *EnableMetricRulesResponse) SetBody(v *EnableMetricRulesResponseBody) *EnableMetricRulesResponse {
	s.Body = v
	return s
}

type EnableSiteMonitorsRequest struct {
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s EnableSiteMonitorsRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableSiteMonitorsRequest) GoString() string {
	return s.String()
}

func (s *EnableSiteMonitorsRequest) SetTaskIds(v string) *EnableSiteMonitorsRequest {
	s.TaskIds = &v
	return s
}

type EnableSiteMonitorsResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EnableSiteMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableSiteMonitorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSiteMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSiteMonitorsResponseBody) SetMessage(v string) *EnableSiteMonitorsResponseBody {
	s.Message = &v
	return s
}

func (s *EnableSiteMonitorsResponseBody) SetRequestId(v string) *EnableSiteMonitorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSiteMonitorsResponseBody) SetData(v *EnableSiteMonitorsResponseBodyData) *EnableSiteMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *EnableSiteMonitorsResponseBody) SetCode(v string) *EnableSiteMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *EnableSiteMonitorsResponseBody) SetSuccess(v string) *EnableSiteMonitorsResponseBody {
	s.Success = &v
	return s
}

type EnableSiteMonitorsResponseBodyData struct {
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s EnableSiteMonitorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableSiteMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableSiteMonitorsResponseBodyData) SetCount(v int32) *EnableSiteMonitorsResponseBodyData {
	s.Count = &v
	return s
}

type EnableSiteMonitorsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableSiteMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSiteMonitorsResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSiteMonitorsResponse) GoString() string {
	return s.String()
}

func (s *EnableSiteMonitorsResponse) SetHeaders(v map[string]*string) *EnableSiteMonitorsResponse {
	s.Headers = v
	return s
}

func (s *EnableSiteMonitorsResponse) SetBody(v *EnableSiteMonitorsResponseBody) *EnableSiteMonitorsResponse {
	s.Body = v
	return s
}

type InstallMonitoringAgentRequest struct {
	Force       *bool     `json:"Force,omitempty" xml:"Force,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s InstallMonitoringAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallMonitoringAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallMonitoringAgentRequest) SetForce(v bool) *InstallMonitoringAgentRequest {
	s.Force = &v
	return s
}

func (s *InstallMonitoringAgentRequest) SetInstanceIds(v []*string) *InstallMonitoringAgentRequest {
	s.InstanceIds = v
	return s
}

type InstallMonitoringAgentResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallMonitoringAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallMonitoringAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallMonitoringAgentResponseBody) SetMessage(v string) *InstallMonitoringAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InstallMonitoringAgentResponseBody) SetRequestId(v string) *InstallMonitoringAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallMonitoringAgentResponseBody) SetCode(v string) *InstallMonitoringAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InstallMonitoringAgentResponseBody) SetSuccess(v bool) *InstallMonitoringAgentResponseBody {
	s.Success = &v
	return s
}

type InstallMonitoringAgentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallMonitoringAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallMonitoringAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallMonitoringAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallMonitoringAgentResponse) SetHeaders(v map[string]*string) *InstallMonitoringAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallMonitoringAgentResponse) SetBody(v *InstallMonitoringAgentResponseBody) *InstallMonitoringAgentResponse {
	s.Body = v
	return s
}

type ModifyGroupMonitoringAgentProcessRequest struct {
	Id                         *string                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	GroupId                    *string                                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	MatchExpressFilterRelation *string                                                `json:"MatchExpressFilterRelation,omitempty" xml:"MatchExpressFilterRelation,omitempty"`
	AlertConfig                []*ModifyGroupMonitoringAgentProcessRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s ModifyGroupMonitoringAgentProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetId(v string) *ModifyGroupMonitoringAgentProcessRequest {
	s.Id = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetGroupId(v string) *ModifyGroupMonitoringAgentProcessRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetMatchExpressFilterRelation(v string) *ModifyGroupMonitoringAgentProcessRequest {
	s.MatchExpressFilterRelation = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequest) SetAlertConfig(v []*ModifyGroupMonitoringAgentProcessRequestAlertConfig) *ModifyGroupMonitoringAgentProcessRequest {
	s.AlertConfig = v
	return s
}

type ModifyGroupMonitoringAgentProcessRequestAlertConfig struct {
	SilenceTime         *string `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	ComparisonOperator  *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Webhook             *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Times               *string `json:"Times,omitempty" xml:"Times,omitempty"`
	EscalationsLevel    *string `json:"EscalationsLevel,omitempty" xml:"EscalationsLevel,omitempty"`
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	EffectiveInterval   *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	Threshold           *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics          *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s ModifyGroupMonitoringAgentProcessRequestAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetSilenceTime(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetComparisonOperator(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetWebhook(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Webhook = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetTimes(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Times = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetEscalationsLevel(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.EscalationsLevel = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetNoEffectiveInterval(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.NoEffectiveInterval = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetEffectiveInterval(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.EffectiveInterval = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetThreshold(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Threshold = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessRequestAlertConfig) SetStatistics(v string) *ModifyGroupMonitoringAgentProcessRequestAlertConfig {
	s.Statistics = &v
	return s
}

type ModifyGroupMonitoringAgentProcessResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyGroupMonitoringAgentProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetCode(v string) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

type ModifyGroupMonitoringAgentProcessResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyGroupMonitoringAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGroupMonitoringAgentProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessResponse) SetHeaders(v map[string]*string) *ModifyGroupMonitoringAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponse) SetBody(v *ModifyGroupMonitoringAgentProcessResponseBody) *ModifyGroupMonitoringAgentProcessResponse {
	s.Body = v
	return s
}

type ModifyHostAvailabilityRequest struct {
	TaskOption                *ModifyHostAvailabilityRequestTaskOption                  `json:"TaskOption,omitempty" xml:"TaskOption,omitempty" type:"Struct"`
	AlertConfig               *ModifyHostAvailabilityRequestAlertConfig                 `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	GroupId                   *int64                                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Id                        *int64                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskName                  *string                                                   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskScope                 *string                                                   `json:"TaskScope,omitempty" xml:"TaskScope,omitempty"`
	AlertConfigEscalationList []*ModifyHostAvailabilityRequestAlertConfigEscalationList `json:"AlertConfigEscalationList,omitempty" xml:"AlertConfigEscalationList,omitempty" type:"Repeated"`
	InstanceList              []*string                                                 `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s ModifyHostAvailabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequest) SetTaskOption(v *ModifyHostAvailabilityRequestTaskOption) *ModifyHostAvailabilityRequest {
	s.TaskOption = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetAlertConfig(v *ModifyHostAvailabilityRequestAlertConfig) *ModifyHostAvailabilityRequest {
	s.AlertConfig = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetGroupId(v int64) *ModifyHostAvailabilityRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetId(v int64) *ModifyHostAvailabilityRequest {
	s.Id = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetTaskName(v string) *ModifyHostAvailabilityRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetTaskScope(v string) *ModifyHostAvailabilityRequest {
	s.TaskScope = &v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetAlertConfigEscalationList(v []*ModifyHostAvailabilityRequestAlertConfigEscalationList) *ModifyHostAvailabilityRequest {
	s.AlertConfigEscalationList = v
	return s
}

func (s *ModifyHostAvailabilityRequest) SetInstanceList(v []*string) *ModifyHostAvailabilityRequest {
	s.InstanceList = v
	return s
}

type ModifyHostAvailabilityRequestTaskOption struct {
	HttpURI                  *string `json:"HttpURI,omitempty" xml:"HttpURI,omitempty"`
	TelnetOrPingHost         *string `json:"TelnetOrPingHost,omitempty" xml:"TelnetOrPingHost,omitempty"`
	HttpResponseCharset      *string `json:"HttpResponseCharset,omitempty" xml:"HttpResponseCharset,omitempty"`
	HttpPostContent          *string `json:"HttpPostContent,omitempty" xml:"HttpPostContent,omitempty"`
	HttpResponseMatchContent *string `json:"HttpResponseMatchContent,omitempty" xml:"HttpResponseMatchContent,omitempty"`
	HttpMethod               *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpNegative             *bool   `json:"HttpNegative,omitempty" xml:"HttpNegative,omitempty"`
}

func (s ModifyHostAvailabilityRequestTaskOption) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAvailabilityRequestTaskOption) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpURI(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpURI = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetTelnetOrPingHost(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.TelnetOrPingHost = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpResponseCharset(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpResponseCharset = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpPostContent(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpPostContent = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpResponseMatchContent(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpResponseMatchContent = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpMethod(v string) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpMethod = &v
	return s
}

func (s *ModifyHostAvailabilityRequestTaskOption) SetHttpNegative(v bool) *ModifyHostAvailabilityRequestTaskOption {
	s.HttpNegative = &v
	return s
}

type ModifyHostAvailabilityRequestAlertConfig struct {
	NotifyType  *int32  `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	StartTime   *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SilenceTime *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	WebHook     *string `json:"WebHook,omitempty" xml:"WebHook,omitempty"`
}

func (s ModifyHostAvailabilityRequestAlertConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAvailabilityRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetNotifyType(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.NotifyType = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetStartTime(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.StartTime = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetEndTime(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.EndTime = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetSilenceTime(v int32) *ModifyHostAvailabilityRequestAlertConfig {
	s.SilenceTime = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfig) SetWebHook(v string) *ModifyHostAvailabilityRequestAlertConfig {
	s.WebHook = &v
	return s
}

type ModifyHostAvailabilityRequestAlertConfigEscalationList struct {
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Times      *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Aggregate  *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
}

func (s ModifyHostAvailabilityRequestAlertConfigEscalationList) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAvailabilityRequestAlertConfigEscalationList) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetValue(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Value = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetMetricName(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.MetricName = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetTimes(v int32) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Times = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetOperator(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Operator = &v
	return s
}

func (s *ModifyHostAvailabilityRequestAlertConfigEscalationList) SetAggregate(v string) *ModifyHostAvailabilityRequestAlertConfigEscalationList {
	s.Aggregate = &v
	return s
}

type ModifyHostAvailabilityResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyHostAvailabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityResponseBody) SetMessage(v string) *ModifyHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHostAvailabilityResponseBody) SetRequestId(v string) *ModifyHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostAvailabilityResponseBody) SetCode(v string) *ModifyHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHostAvailabilityResponseBody) SetSuccess(v bool) *ModifyHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

type ModifyHostAvailabilityResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostAvailabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityResponse) SetHeaders(v map[string]*string) *ModifyHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostAvailabilityResponse) SetBody(v *ModifyHostAvailabilityResponseBody) *ModifyHostAvailabilityResponse {
	s.Body = v
	return s
}

type ModifyHostInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	HostName   *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
}

func (s ModifyHostInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostInfoRequest) SetInstanceId(v string) *ModifyHostInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostInfoRequest) SetHostName(v string) *ModifyHostInfoRequest {
	s.HostName = &v
	return s
}

type ModifyHostInfoResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyHostInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostInfoResponseBody) SetMessage(v string) *ModifyHostInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHostInfoResponseBody) SetRequestId(v string) *ModifyHostInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostInfoResponseBody) SetCode(v string) *ModifyHostInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHostInfoResponseBody) SetSuccess(v bool) *ModifyHostInfoResponseBody {
	s.Success = &v
	return s
}

type ModifyHostInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHostInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostInfoResponse) SetHeaders(v map[string]*string) *ModifyHostInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostInfoResponse) SetBody(v *ModifyHostInfoResponseBody) *ModifyHostInfoResponse {
	s.Body = v
	return s
}

type ModifyMetricRuleTemplateRequest struct {
	TemplateId     *int64                                           `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Name           *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	RestVersion    *int64                                           `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	AlertTemplates []*ModifyMetricRuleTemplateRequestAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Repeated"`
}

func (s ModifyMetricRuleTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequest) SetTemplateId(v int64) *ModifyMetricRuleTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetName(v string) *ModifyMetricRuleTemplateRequest {
	s.Name = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetDescription(v string) *ModifyMetricRuleTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetRestVersion(v int64) *ModifyMetricRuleTemplateRequest {
	s.RestVersion = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetAlertTemplates(v []*ModifyMetricRuleTemplateRequestAlertTemplates) *ModifyMetricRuleTemplateRequest {
	s.AlertTemplates = v
	return s
}

type ModifyMetricRuleTemplateRequestAlertTemplates struct {
	Escalations *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" require:"true" type:"Struct"`
	MetricName  *string                                                   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Webhook     *string                                                   `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	Namespace   *string                                                   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RuleName    *string                                                   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Period      *int32                                                    `json:"Period,omitempty" xml:"Period,omitempty"`
	Selector    *string                                                   `json:"Selector,omitempty" xml:"Selector,omitempty"`
	Category    *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplates) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplates) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetEscalations(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Escalations = v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetMetricName(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.MetricName = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetWebhook(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Webhook = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetNamespace(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Namespace = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetRuleName(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.RuleName = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetPeriod(v int32) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Period = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetSelector(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Selector = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetCategory(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Category = &v
	return s
}

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalations struct {
	Info     *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" require:"true" type:"Struct"`
	Warn     *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" require:"true" type:"Struct"`
	Critical *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" require:"true" type:"Struct"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) SetInfo(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Info = v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) SetWarn(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Warn = v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) SetCritical(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Critical = v
	return s
}

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetThreshold(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetStatistics(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetComparisonOperator(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetTimes(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Times = &v
	return s
}

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetThreshold(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetTimes(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetComparisonOperator(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetStatistics(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Statistics = &v
	return s
}

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical struct {
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetTimes(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetThreshold(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetComparisonOperator(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetStatistics(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Statistics = &v
	return s
}

type ModifyMetricRuleTemplateResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMetricRuleTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateResponseBody) SetMessage(v string) *ModifyMetricRuleTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponseBody) SetRequestId(v string) *ModifyMetricRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponseBody) SetCode(v int32) *ModifyMetricRuleTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMetricRuleTemplateResponseBody) SetSuccess(v bool) *ModifyMetricRuleTemplateResponseBody {
	s.Success = &v
	return s
}

type ModifyMetricRuleTemplateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMetricRuleTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMetricRuleTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMetricRuleTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateResponse) SetHeaders(v map[string]*string) *ModifyMetricRuleTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyMetricRuleTemplateResponse) SetBody(v *ModifyMetricRuleTemplateResponseBody) *ModifyMetricRuleTemplateResponse {
	s.Body = v
	return s
}

type ModifyMonitorGroupRequest struct {
	BindUrls      *string `json:"BindUrls,omitempty" xml:"BindUrls,omitempty"`
	ServiceId     *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
}

func (s ModifyMonitorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupRequest) SetBindUrls(v string) *ModifyMonitorGroupRequest {
	s.BindUrls = &v
	return s
}

func (s *ModifyMonitorGroupRequest) SetServiceId(v int64) *ModifyMonitorGroupRequest {
	s.ServiceId = &v
	return s
}

func (s *ModifyMonitorGroupRequest) SetGroupId(v string) *ModifyMonitorGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyMonitorGroupRequest) SetGroupName(v string) *ModifyMonitorGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyMonitorGroupRequest) SetContactGroups(v string) *ModifyMonitorGroupRequest {
	s.ContactGroups = &v
	return s
}

type ModifyMonitorGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMonitorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupResponseBody) SetMessage(v string) *ModifyMonitorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMonitorGroupResponseBody) SetRequestId(v string) *ModifyMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMonitorGroupResponseBody) SetCode(v int32) *ModifyMonitorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMonitorGroupResponseBody) SetSuccess(v bool) *ModifyMonitorGroupResponseBody {
	s.Success = &v
	return s
}

type ModifyMonitorGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMonitorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMonitorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMonitorGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupResponse) SetHeaders(v map[string]*string) *ModifyMonitorGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyMonitorGroupResponse) SetBody(v *ModifyMonitorGroupResponseBody) *ModifyMonitorGroupResponse {
	s.Body = v
	return s
}

type ModifyMonitorGroupInstancesRequest struct {
	GroupId   *int64                                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Instances []*ModifyMonitorGroupInstancesRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s ModifyMonitorGroupInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesRequest) SetGroupId(v int64) *ModifyMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequest) SetInstances(v []*ModifyMonitorGroupInstancesRequestInstances) *ModifyMonitorGroupInstancesRequest {
	s.Instances = v
	return s
}

type ModifyMonitorGroupInstancesRequestInstances struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyMonitorGroupInstancesRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s ModifyMonitorGroupInstancesRequestInstances) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetInstanceName(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.InstanceName = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetCategory(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.Category = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetInstanceId(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetRegionId(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.RegionId = &v
	return s
}

type ModifyMonitorGroupInstancesResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMonitorGroupInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetMessage(v string) *ModifyMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetRequestId(v string) *ModifyMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetCode(v int32) *ModifyMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMonitorGroupInstancesResponseBody) SetSuccess(v bool) *ModifyMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

type ModifyMonitorGroupInstancesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMonitorGroupInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMonitorGroupInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMonitorGroupInstancesResponse) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesResponse) SetHeaders(v map[string]*string) *ModifyMonitorGroupInstancesResponse {
	s.Headers = v
	return s
}

func (s *ModifyMonitorGroupInstancesResponse) SetBody(v *ModifyMonitorGroupInstancesResponseBody) *ModifyMonitorGroupInstancesResponse {
	s.Body = v
	return s
}

type ModifySiteMonitorRequest struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName    *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspCities   *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	AlertIds    *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
}

func (s ModifySiteMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySiteMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorRequest) SetAddress(v string) *ModifySiteMonitorRequest {
	s.Address = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetTaskId(v string) *ModifySiteMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetTaskName(v string) *ModifySiteMonitorRequest {
	s.TaskName = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetInterval(v string) *ModifySiteMonitorRequest {
	s.Interval = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetIspCities(v string) *ModifySiteMonitorRequest {
	s.IspCities = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetOptionsJson(v string) *ModifySiteMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetAlertIds(v string) *ModifySiteMonitorRequest {
	s.AlertIds = &v
	return s
}

type ModifySiteMonitorResponseBody struct {
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ModifySiteMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySiteMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySiteMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorResponseBody) SetMessage(v string) *ModifySiteMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetRequestId(v string) *ModifySiteMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetData(v *ModifySiteMonitorResponseBodyData) *ModifySiteMonitorResponseBody {
	s.Data = v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetCode(v string) *ModifySiteMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetSuccess(v string) *ModifySiteMonitorResponseBody {
	s.Success = &v
	return s
}

type ModifySiteMonitorResponseBodyData struct {
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s ModifySiteMonitorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifySiteMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorResponseBodyData) SetCount(v int32) *ModifySiteMonitorResponseBodyData {
	s.Count = &v
	return s
}

type ModifySiteMonitorResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySiteMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySiteMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySiteMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorResponse) SetHeaders(v map[string]*string) *ModifySiteMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifySiteMonitorResponse) SetBody(v *ModifySiteMonitorResponseBody) *ModifySiteMonitorResponse {
	s.Body = v
	return s
}

type OpenCmsServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s OpenCmsServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCmsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCmsServiceResponseBody) SetRequestId(v string) *OpenCmsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenCmsServiceResponseBody) SetOrderId(v string) *OpenCmsServiceResponseBody {
	s.OrderId = &v
	return s
}

type OpenCmsServiceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenCmsServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenCmsServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCmsServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenCmsServiceResponse) SetHeaders(v map[string]*string) *OpenCmsServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenCmsServiceResponse) SetBody(v *OpenCmsServiceResponseBody) *OpenCmsServiceResponse {
	s.Body = v
	return s
}

type PutContactRequest struct {
	Channels    *PutContactRequestChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Struct"`
	ContactName *string                    `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Describe    *string                    `json:"Describe,omitempty" xml:"Describe,omitempty"`
	Lang        *string                    `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s PutContactRequest) String() string {
	return tea.Prettify(s)
}

func (s PutContactRequest) GoString() string {
	return s.String()
}

func (s *PutContactRequest) SetChannels(v *PutContactRequestChannels) *PutContactRequest {
	s.Channels = v
	return s
}

func (s *PutContactRequest) SetContactName(v string) *PutContactRequest {
	s.ContactName = &v
	return s
}

func (s *PutContactRequest) SetDescribe(v string) *PutContactRequest {
	s.Describe = &v
	return s
}

func (s *PutContactRequest) SetLang(v string) *PutContactRequest {
	s.Lang = &v
	return s
}

type PutContactRequestChannels struct {
	SMS         *string `json:"SMS,omitempty" xml:"SMS,omitempty"`
	Mail        *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	AliIM       *string `json:"AliIM,omitempty" xml:"AliIM,omitempty"`
	DingWebHook *string `json:"DingWebHook,omitempty" xml:"DingWebHook,omitempty"`
}

func (s PutContactRequestChannels) String() string {
	return tea.Prettify(s)
}

func (s PutContactRequestChannels) GoString() string {
	return s.String()
}

func (s *PutContactRequestChannels) SetSMS(v string) *PutContactRequestChannels {
	s.SMS = &v
	return s
}

func (s *PutContactRequestChannels) SetMail(v string) *PutContactRequestChannels {
	s.Mail = &v
	return s
}

func (s *PutContactRequestChannels) SetAliIM(v string) *PutContactRequestChannels {
	s.AliIM = &v
	return s
}

func (s *PutContactRequestChannels) SetDingWebHook(v string) *PutContactRequestChannels {
	s.DingWebHook = &v
	return s
}

type PutContactResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutContactResponseBody) GoString() string {
	return s.String()
}

func (s *PutContactResponseBody) SetMessage(v string) *PutContactResponseBody {
	s.Message = &v
	return s
}

func (s *PutContactResponseBody) SetRequestId(v string) *PutContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutContactResponseBody) SetCode(v string) *PutContactResponseBody {
	s.Code = &v
	return s
}

func (s *PutContactResponseBody) SetSuccess(v bool) *PutContactResponseBody {
	s.Success = &v
	return s
}

type PutContactResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutContactResponse) String() string {
	return tea.Prettify(s)
}

func (s PutContactResponse) GoString() string {
	return s.String()
}

func (s *PutContactResponse) SetHeaders(v map[string]*string) *PutContactResponse {
	s.Headers = v
	return s
}

func (s *PutContactResponse) SetBody(v *PutContactResponseBody) *PutContactResponse {
	s.Body = v
	return s
}

type PutContactGroupRequest struct {
	ContactGroupName *string   `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Describe         *string   `json:"Describe,omitempty" xml:"Describe,omitempty"`
	EnableSubscribed *bool     `json:"EnableSubscribed,omitempty" xml:"EnableSubscribed,omitempty"`
	ContactNames     []*string `json:"ContactNames,omitempty" xml:"ContactNames,omitempty" type:"Repeated"`
}

func (s PutContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s PutContactGroupRequest) GoString() string {
	return s.String()
}

func (s *PutContactGroupRequest) SetContactGroupName(v string) *PutContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *PutContactGroupRequest) SetDescribe(v string) *PutContactGroupRequest {
	s.Describe = &v
	return s
}

func (s *PutContactGroupRequest) SetEnableSubscribed(v bool) *PutContactGroupRequest {
	s.EnableSubscribed = &v
	return s
}

func (s *PutContactGroupRequest) SetContactNames(v []*string) *PutContactGroupRequest {
	s.ContactNames = v
	return s
}

type PutContactGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *PutContactGroupResponseBody) SetMessage(v string) *PutContactGroupResponseBody {
	s.Message = &v
	return s
}

func (s *PutContactGroupResponseBody) SetRequestId(v string) *PutContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutContactGroupResponseBody) SetCode(v string) *PutContactGroupResponseBody {
	s.Code = &v
	return s
}

func (s *PutContactGroupResponseBody) SetSuccess(v bool) *PutContactGroupResponseBody {
	s.Success = &v
	return s
}

type PutContactGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s PutContactGroupResponse) GoString() string {
	return s.String()
}

func (s *PutContactGroupResponse) SetHeaders(v map[string]*string) *PutContactGroupResponse {
	s.Headers = v
	return s
}

func (s *PutContactGroupResponse) SetBody(v *PutContactGroupResponseBody) *PutContactGroupResponse {
	s.Body = v
	return s
}

type PutCustomEventRequest struct {
	EventInfo []*PutCustomEventRequestEventInfo `json:"EventInfo,omitempty" xml:"EventInfo,omitempty" type:"Repeated"`
}

func (s PutCustomEventRequest) String() string {
	return tea.Prettify(s)
}

func (s PutCustomEventRequest) GoString() string {
	return s.String()
}

func (s *PutCustomEventRequest) SetEventInfo(v []*PutCustomEventRequestEventInfo) *PutCustomEventRequest {
	s.EventInfo = v
	return s
}

type PutCustomEventRequestEventInfo struct {
	Time      *string `json:"Time,omitempty" xml:"Time,omitempty"`
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s PutCustomEventRequestEventInfo) String() string {
	return tea.Prettify(s)
}

func (s PutCustomEventRequestEventInfo) GoString() string {
	return s.String()
}

func (s *PutCustomEventRequestEventInfo) SetTime(v string) *PutCustomEventRequestEventInfo {
	s.Time = &v
	return s
}

func (s *PutCustomEventRequestEventInfo) SetEventName(v string) *PutCustomEventRequestEventInfo {
	s.EventName = &v
	return s
}

func (s *PutCustomEventRequestEventInfo) SetGroupId(v string) *PutCustomEventRequestEventInfo {
	s.GroupId = &v
	return s
}

func (s *PutCustomEventRequestEventInfo) SetContent(v string) *PutCustomEventRequestEventInfo {
	s.Content = &v
	return s
}

type PutCustomEventResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s PutCustomEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutCustomEventResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomEventResponseBody) SetMessage(v string) *PutCustomEventResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomEventResponseBody) SetRequestId(v string) *PutCustomEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomEventResponseBody) SetCode(v string) *PutCustomEventResponseBody {
	s.Code = &v
	return s
}

type PutCustomEventResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutCustomEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutCustomEventResponse) String() string {
	return tea.Prettify(s)
}

func (s PutCustomEventResponse) GoString() string {
	return s.String()
}

func (s *PutCustomEventResponse) SetHeaders(v map[string]*string) *PutCustomEventResponse {
	s.Headers = v
	return s
}

func (s *PutCustomEventResponse) SetBody(v *PutCustomEventResponseBody) *PutCustomEventResponse {
	s.Body = v
	return s
}

type PutCustomEventRuleRequest struct {
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	EventName         *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	ContactGroups     *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Webhook           *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	Period            *string `json:"Period,omitempty" xml:"Period,omitempty"`
	EmailSubject      *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	Threshold         *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Level             *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutCustomEventRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutCustomEventRuleRequest) GoString() string {
	return s.String()
}

func (s *PutCustomEventRuleRequest) SetGroupId(v string) *PutCustomEventRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetRuleId(v string) *PutCustomEventRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetRuleName(v string) *PutCustomEventRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetEventName(v string) *PutCustomEventRuleRequest {
	s.EventName = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetContactGroups(v string) *PutCustomEventRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetWebhook(v string) *PutCustomEventRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetEffectiveInterval(v string) *PutCustomEventRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetPeriod(v string) *PutCustomEventRuleRequest {
	s.Period = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetEmailSubject(v string) *PutCustomEventRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetThreshold(v string) *PutCustomEventRuleRequest {
	s.Threshold = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetLevel(v string) *PutCustomEventRuleRequest {
	s.Level = &v
	return s
}

type PutCustomEventRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutCustomEventRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutCustomEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomEventRuleResponseBody) SetMessage(v string) *PutCustomEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomEventRuleResponseBody) SetRequestId(v string) *PutCustomEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomEventRuleResponseBody) SetCode(v string) *PutCustomEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutCustomEventRuleResponseBody) SetSuccess(v bool) *PutCustomEventRuleResponseBody {
	s.Success = &v
	return s
}

type PutCustomEventRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutCustomEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutCustomEventRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutCustomEventRuleResponse) GoString() string {
	return s.String()
}

func (s *PutCustomEventRuleResponse) SetHeaders(v map[string]*string) *PutCustomEventRuleResponse {
	s.Headers = v
	return s
}

func (s *PutCustomEventRuleResponse) SetBody(v *PutCustomEventRuleResponseBody) *PutCustomEventRuleResponse {
	s.Body = v
	return s
}

type PutCustomMetricRequest struct {
	MetricList []*PutCustomMetricRequestMetricList `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
}

func (s PutCustomMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricRequest) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRequest) SetMetricList(v []*PutCustomMetricRequestMetricList) *PutCustomMetricRequest {
	s.MetricList = v
	return s
}

type PutCustomMetricRequestMetricList struct {
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Values     *string `json:"Values,omitempty" xml:"Values,omitempty"`
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s PutCustomMetricRequestMetricList) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricRequestMetricList) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRequestMetricList) SetType(v string) *PutCustomMetricRequestMetricList {
	s.Type = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetMetricName(v string) *PutCustomMetricRequestMetricList {
	s.MetricName = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetTime(v string) *PutCustomMetricRequestMetricList {
	s.Time = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetGroupId(v string) *PutCustomMetricRequestMetricList {
	s.GroupId = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetValues(v string) *PutCustomMetricRequestMetricList {
	s.Values = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetDimensions(v string) *PutCustomMetricRequestMetricList {
	s.Dimensions = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetPeriod(v string) *PutCustomMetricRequestMetricList {
	s.Period = &v
	return s
}

type PutCustomMetricResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s PutCustomMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomMetricResponseBody) SetMessage(v string) *PutCustomMetricResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomMetricResponseBody) SetRequestId(v string) *PutCustomMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomMetricResponseBody) SetCode(v string) *PutCustomMetricResponseBody {
	s.Code = &v
	return s
}

type PutCustomMetricResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutCustomMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutCustomMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricResponse) GoString() string {
	return s.String()
}

func (s *PutCustomMetricResponse) SetHeaders(v map[string]*string) *PutCustomMetricResponse {
	s.Headers = v
	return s
}

func (s *PutCustomMetricResponse) SetBody(v *PutCustomMetricResponseBody) *PutCustomMetricResponse {
	s.Body = v
	return s
}

type PutCustomMetricRuleRequest struct {
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RuleId             *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Resources          *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	ContactGroups      *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Webhook            *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	EffectiveInterval  *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	SilenceTime        *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	EmailSubject       *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Level              *string `json:"Level,omitempty" xml:"Level,omitempty"`
	EvaluationCount    *int32  `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
}

func (s PutCustomMetricRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRuleRequest) SetGroupId(v string) *PutCustomMetricRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetRuleId(v string) *PutCustomMetricRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetRuleName(v string) *PutCustomMetricRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetMetricName(v string) *PutCustomMetricRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetResources(v string) *PutCustomMetricRuleRequest {
	s.Resources = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetContactGroups(v string) *PutCustomMetricRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetWebhook(v string) *PutCustomMetricRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetEffectiveInterval(v string) *PutCustomMetricRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetSilenceTime(v int32) *PutCustomMetricRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetPeriod(v string) *PutCustomMetricRuleRequest {
	s.Period = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetEmailSubject(v string) *PutCustomMetricRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetThreshold(v string) *PutCustomMetricRuleRequest {
	s.Threshold = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetLevel(v string) *PutCustomMetricRuleRequest {
	s.Level = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetEvaluationCount(v int32) *PutCustomMetricRuleRequest {
	s.EvaluationCount = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetStatistics(v string) *PutCustomMetricRuleRequest {
	s.Statistics = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetComparisonOperator(v string) *PutCustomMetricRuleRequest {
	s.ComparisonOperator = &v
	return s
}

type PutCustomMetricRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutCustomMetricRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRuleResponseBody) SetMessage(v string) *PutCustomMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomMetricRuleResponseBody) SetRequestId(v string) *PutCustomMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomMetricRuleResponseBody) SetCode(v string) *PutCustomMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutCustomMetricRuleResponseBody) SetSuccess(v bool) *PutCustomMetricRuleResponseBody {
	s.Success = &v
	return s
}

type PutCustomMetricRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutCustomMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutCustomMetricRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutCustomMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRuleResponse) SetHeaders(v map[string]*string) *PutCustomMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *PutCustomMetricRuleResponse) SetBody(v *PutCustomMetricRuleResponseBody) *PutCustomMetricRuleResponse {
	s.Body = v
	return s
}

type PutEventRuleRequest struct {
	RuleName     *string                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	GroupId      *string                            `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	EventType    *string                            `json:"EventType,omitempty" xml:"EventType,omitempty"`
	Description  *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	State        *string                            `json:"State,omitempty" xml:"State,omitempty"`
	EventPattern []*PutEventRuleRequestEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Repeated"`
}

func (s PutEventRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleRequest) GoString() string {
	return s.String()
}

func (s *PutEventRuleRequest) SetRuleName(v string) *PutEventRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutEventRuleRequest) SetGroupId(v string) *PutEventRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutEventRuleRequest) SetEventType(v string) *PutEventRuleRequest {
	s.EventType = &v
	return s
}

func (s *PutEventRuleRequest) SetDescription(v string) *PutEventRuleRequest {
	s.Description = &v
	return s
}

func (s *PutEventRuleRequest) SetState(v string) *PutEventRuleRequest {
	s.State = &v
	return s
}

func (s *PutEventRuleRequest) SetEventPattern(v []*PutEventRuleRequestEventPattern) *PutEventRuleRequest {
	s.EventPattern = v
	return s
}

type PutEventRuleRequestEventPattern struct {
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
	StatusList    []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	Product       *string   `json:"Product,omitempty" xml:"Product,omitempty"`
	LevelList     []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
	NameList      []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
}

func (s PutEventRuleRequestEventPattern) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleRequestEventPattern) GoString() string {
	return s.String()
}

func (s *PutEventRuleRequestEventPattern) SetEventTypeList(v []*string) *PutEventRuleRequestEventPattern {
	s.EventTypeList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetStatusList(v []*string) *PutEventRuleRequestEventPattern {
	s.StatusList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetProduct(v string) *PutEventRuleRequestEventPattern {
	s.Product = &v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetLevelList(v []*string) *PutEventRuleRequestEventPattern {
	s.LevelList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetNameList(v []*string) *PutEventRuleRequestEventPattern {
	s.NameList = v
	return s
}

type PutEventRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutEventRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventRuleResponseBody) SetMessage(v string) *PutEventRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventRuleResponseBody) SetRequestId(v string) *PutEventRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEventRuleResponseBody) SetData(v string) *PutEventRuleResponseBody {
	s.Data = &v
	return s
}

func (s *PutEventRuleResponseBody) SetCode(v string) *PutEventRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventRuleResponseBody) SetSuccess(v bool) *PutEventRuleResponseBody {
	s.Success = &v
	return s
}

type PutEventRuleResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutEventRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleResponse) GoString() string {
	return s.String()
}

func (s *PutEventRuleResponse) SetHeaders(v map[string]*string) *PutEventRuleResponse {
	s.Headers = v
	return s
}

func (s *PutEventRuleResponse) SetBody(v *PutEventRuleResponseBody) *PutEventRuleResponse {
	s.Body = v
	return s
}

type PutEventRuleTargetsRequest struct {
	RuleName          *string                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	FcParameters      []*PutEventRuleTargetsRequestFcParameters      `json:"FcParameters,omitempty" xml:"FcParameters,omitempty" type:"Repeated"`
	ContactParameters []*PutEventRuleTargetsRequestContactParameters `json:"ContactParameters,omitempty" xml:"ContactParameters,omitempty" type:"Repeated"`
	MnsParameters     []*PutEventRuleTargetsRequestMnsParameters     `json:"MnsParameters,omitempty" xml:"MnsParameters,omitempty" type:"Repeated"`
	WebhookParameters []*PutEventRuleTargetsRequestWebhookParameters `json:"WebhookParameters,omitempty" xml:"WebhookParameters,omitempty" type:"Repeated"`
	SlsParameters     []*PutEventRuleTargetsRequestSlsParameters     `json:"SlsParameters,omitempty" xml:"SlsParameters,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequest) SetRuleName(v string) *PutEventRuleTargetsRequest {
	s.RuleName = &v
	return s
}

func (s *PutEventRuleTargetsRequest) SetFcParameters(v []*PutEventRuleTargetsRequestFcParameters) *PutEventRuleTargetsRequest {
	s.FcParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetContactParameters(v []*PutEventRuleTargetsRequestContactParameters) *PutEventRuleTargetsRequest {
	s.ContactParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetMnsParameters(v []*PutEventRuleTargetsRequestMnsParameters) *PutEventRuleTargetsRequest {
	s.MnsParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetWebhookParameters(v []*PutEventRuleTargetsRequestWebhookParameters) *PutEventRuleTargetsRequest {
	s.WebhookParameters = v
	return s
}

func (s *PutEventRuleTargetsRequest) SetSlsParameters(v []*PutEventRuleTargetsRequestSlsParameters) *PutEventRuleTargetsRequest {
	s.SlsParameters = v
	return s
}

type PutEventRuleTargetsRequestFcParameters struct {
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutEventRuleTargetsRequestFcParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsRequestFcParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestFcParameters) SetFunctionName(v string) *PutEventRuleTargetsRequestFcParameters {
	s.FunctionName = &v
	return s
}

func (s *PutEventRuleTargetsRequestFcParameters) SetRegion(v string) *PutEventRuleTargetsRequestFcParameters {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsRequestFcParameters) SetServiceName(v string) *PutEventRuleTargetsRequestFcParameters {
	s.ServiceName = &v
	return s
}

func (s *PutEventRuleTargetsRequestFcParameters) SetId(v string) *PutEventRuleTargetsRequestFcParameters {
	s.Id = &v
	return s
}

type PutEventRuleTargetsRequestContactParameters struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutEventRuleTargetsRequestContactParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsRequestContactParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestContactParameters) SetContactGroupName(v string) *PutEventRuleTargetsRequestContactParameters {
	s.ContactGroupName = &v
	return s
}

func (s *PutEventRuleTargetsRequestContactParameters) SetLevel(v string) *PutEventRuleTargetsRequestContactParameters {
	s.Level = &v
	return s
}

func (s *PutEventRuleTargetsRequestContactParameters) SetId(v string) *PutEventRuleTargetsRequestContactParameters {
	s.Id = &v
	return s
}

type PutEventRuleTargetsRequestMnsParameters struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Queue  *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutEventRuleTargetsRequestMnsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsRequestMnsParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestMnsParameters) SetRegion(v string) *PutEventRuleTargetsRequestMnsParameters {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsRequestMnsParameters) SetQueue(v string) *PutEventRuleTargetsRequestMnsParameters {
	s.Queue = &v
	return s
}

func (s *PutEventRuleTargetsRequestMnsParameters) SetId(v string) *PutEventRuleTargetsRequestMnsParameters {
	s.Id = &v
	return s
}

type PutEventRuleTargetsRequestWebhookParameters struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutEventRuleTargetsRequestWebhookParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsRequestWebhookParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetProtocol(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Protocol = &v
	return s
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetMethod(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Method = &v
	return s
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetUrl(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Url = &v
	return s
}

func (s *PutEventRuleTargetsRequestWebhookParameters) SetId(v string) *PutEventRuleTargetsRequestWebhookParameters {
	s.Id = &v
	return s
}

type PutEventRuleTargetsRequestSlsParameters struct {
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutEventRuleTargetsRequestSlsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsRequestSlsParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetLogStore(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.LogStore = &v
	return s
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetRegion(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetProject(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.Project = &v
	return s
}

func (s *PutEventRuleTargetsRequestSlsParameters) SetId(v string) *PutEventRuleTargetsRequestSlsParameters {
	s.Id = &v
	return s
}

type PutEventRuleTargetsResponseBody struct {
	Message                 *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedMnsParameters     *PutEventRuleTargetsResponseBodyFailedMnsParameters     `json:"FailedMnsParameters,omitempty" xml:"FailedMnsParameters,omitempty" type:"Struct"`
	FailedFcParameters      *PutEventRuleTargetsResponseBodyFailedFcParameters      `json:"FailedFcParameters,omitempty" xml:"FailedFcParameters,omitempty" type:"Struct"`
	FailedParameterCount    *string                                                 `json:"FailedParameterCount,omitempty" xml:"FailedParameterCount,omitempty"`
	FailedContactParameters *PutEventRuleTargetsResponseBodyFailedContactParameters `json:"FailedContactParameters,omitempty" xml:"FailedContactParameters,omitempty" type:"Struct"`
	Code                    *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success                 *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutEventRuleTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBody) SetMessage(v string) *PutEventRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetRequestId(v string) *PutEventRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedMnsParameters(v *PutEventRuleTargetsResponseBodyFailedMnsParameters) *PutEventRuleTargetsResponseBody {
	s.FailedMnsParameters = v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedFcParameters(v *PutEventRuleTargetsResponseBodyFailedFcParameters) *PutEventRuleTargetsResponseBody {
	s.FailedFcParameters = v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedParameterCount(v string) *PutEventRuleTargetsResponseBody {
	s.FailedParameterCount = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetFailedContactParameters(v *PutEventRuleTargetsResponseBodyFailedContactParameters) *PutEventRuleTargetsResponseBody {
	s.FailedContactParameters = v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetCode(v string) *PutEventRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutEventRuleTargetsResponseBody) SetSuccess(v bool) *PutEventRuleTargetsResponseBody {
	s.Success = &v
	return s
}

type PutEventRuleTargetsResponseBodyFailedMnsParameters struct {
	MnsParameter []*PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter `json:"MnsParameter,omitempty" xml:"MnsParameter,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParameters) SetMnsParameter(v []*PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) *PutEventRuleTargetsResponseBodyFailedMnsParameters {
	s.MnsParameter = v
	return s
}

type PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Queue  *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	Id     *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) SetRegion(v string) *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) SetQueue(v string) *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Queue = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter) SetId(v int32) *PutEventRuleTargetsResponseBodyFailedMnsParametersMnsParameter {
	s.Id = &v
	return s
}

type PutEventRuleTargetsResponseBodyFailedFcParameters struct {
	FcParameter []*PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter `json:"FcParameter,omitempty" xml:"FcParameter,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsResponseBodyFailedFcParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedFcParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParameters) SetFcParameter(v []*PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) *PutEventRuleTargetsResponseBodyFailedFcParameters {
	s.FcParameter = v
	return s
}

type PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter struct {
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetFunctionName(v string) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.FunctionName = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetRegion(v string) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.Region = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetServiceName(v string) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.ServiceName = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter) SetId(v int32) *PutEventRuleTargetsResponseBodyFailedFcParametersFcParameter {
	s.Id = &v
	return s
}

type PutEventRuleTargetsResponseBodyFailedContactParameters struct {
	ContactParameter []*PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter `json:"ContactParameter,omitempty" xml:"ContactParameter,omitempty" type:"Repeated"`
}

func (s PutEventRuleTargetsResponseBodyFailedContactParameters) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedContactParameters) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParameters) SetContactParameter(v []*PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) *PutEventRuleTargetsResponseBodyFailedContactParameters {
	s.ContactParameter = v
	return s
}

type PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Id               *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Level            *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) SetContactGroupName(v string) *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter {
	s.ContactGroupName = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) SetId(v int32) *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter {
	s.Id = &v
	return s
}

func (s *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter) SetLevel(v string) *PutEventRuleTargetsResponseBodyFailedContactParametersContactParameter {
	s.Level = &v
	return s
}

type PutEventRuleTargetsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutEventRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutEventRuleTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEventRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutEventRuleTargetsResponse) SetHeaders(v map[string]*string) *PutEventRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutEventRuleTargetsResponse) SetBody(v *PutEventRuleTargetsResponseBody) *PutEventRuleTargetsResponse {
	s.Body = v
	return s
}

type PutExporterOutputRequest struct {
	DestName   *string `json:"DestName,omitempty" xml:"DestName,omitempty"`
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DestType   *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
}

func (s PutExporterOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s PutExporterOutputRequest) GoString() string {
	return s.String()
}

func (s *PutExporterOutputRequest) SetDestName(v string) *PutExporterOutputRequest {
	s.DestName = &v
	return s
}

func (s *PutExporterOutputRequest) SetConfigJson(v string) *PutExporterOutputRequest {
	s.ConfigJson = &v
	return s
}

func (s *PutExporterOutputRequest) SetDesc(v string) *PutExporterOutputRequest {
	s.Desc = &v
	return s
}

func (s *PutExporterOutputRequest) SetDestType(v string) *PutExporterOutputRequest {
	s.DestType = &v
	return s
}

type PutExporterOutputResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutExporterOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutExporterOutputResponseBody) GoString() string {
	return s.String()
}

func (s *PutExporterOutputResponseBody) SetMessage(v string) *PutExporterOutputResponseBody {
	s.Message = &v
	return s
}

func (s *PutExporterOutputResponseBody) SetRequestId(v string) *PutExporterOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutExporterOutputResponseBody) SetCode(v string) *PutExporterOutputResponseBody {
	s.Code = &v
	return s
}

func (s *PutExporterOutputResponseBody) SetSuccess(v bool) *PutExporterOutputResponseBody {
	s.Success = &v
	return s
}

type PutExporterOutputResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutExporterOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutExporterOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s PutExporterOutputResponse) GoString() string {
	return s.String()
}

func (s *PutExporterOutputResponse) SetHeaders(v map[string]*string) *PutExporterOutputResponse {
	s.Headers = v
	return s
}

func (s *PutExporterOutputResponse) SetBody(v *PutExporterOutputResponseBody) *PutExporterOutputResponse {
	s.Body = v
	return s
}

type PutExporterRuleRequest struct {
	RuleName      *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Namespace     *string   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName    *string   `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	TargetWindows *string   `json:"TargetWindows,omitempty" xml:"TargetWindows,omitempty"`
	Describe      *string   `json:"Describe,omitempty" xml:"Describe,omitempty"`
	DstNames      []*string `json:"DstNames,omitempty" xml:"DstNames,omitempty" type:"Repeated"`
}

func (s PutExporterRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutExporterRuleRequest) GoString() string {
	return s.String()
}

func (s *PutExporterRuleRequest) SetRuleName(v string) *PutExporterRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutExporterRuleRequest) SetNamespace(v string) *PutExporterRuleRequest {
	s.Namespace = &v
	return s
}

func (s *PutExporterRuleRequest) SetMetricName(v string) *PutExporterRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutExporterRuleRequest) SetTargetWindows(v string) *PutExporterRuleRequest {
	s.TargetWindows = &v
	return s
}

func (s *PutExporterRuleRequest) SetDescribe(v string) *PutExporterRuleRequest {
	s.Describe = &v
	return s
}

func (s *PutExporterRuleRequest) SetDstNames(v []*string) *PutExporterRuleRequest {
	s.DstNames = v
	return s
}

type PutExporterRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutExporterRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutExporterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutExporterRuleResponseBody) SetMessage(v string) *PutExporterRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutExporterRuleResponseBody) SetRequestId(v string) *PutExporterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutExporterRuleResponseBody) SetCode(v string) *PutExporterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutExporterRuleResponseBody) SetSuccess(v bool) *PutExporterRuleResponseBody {
	s.Success = &v
	return s
}

type PutExporterRuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutExporterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutExporterRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutExporterRuleResponse) GoString() string {
	return s.String()
}

func (s *PutExporterRuleResponse) SetHeaders(v map[string]*string) *PutExporterRuleResponse {
	s.Headers = v
	return s
}

func (s *PutExporterRuleResponse) SetBody(v *PutExporterRuleResponseBody) *PutExporterRuleResponse {
	s.Body = v
	return s
}

type PutGroupMetricRuleRequest struct {
	Escalations         *PutGroupMetricRuleRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	GroupId             *string                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RuleId              *string                               `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Category            *string                               `json:"Category,omitempty" xml:"Category,omitempty"`
	RuleName            *string                               `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Namespace           *string                               `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName          *string                               `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Dimensions          *string                               `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	EffectiveInterval   *string                               `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	NoEffectiveInterval *string                               `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	SilenceTime         *int32                                `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Period              *string                               `json:"Period,omitempty" xml:"Period,omitempty"`
	Interval            *string                               `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Webhook             *string                               `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	EmailSubject        *string                               `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	ContactGroups       *string                               `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
}

func (s PutGroupMetricRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutGroupMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequest) SetEscalations(v *PutGroupMetricRuleRequestEscalations) *PutGroupMetricRuleRequest {
	s.Escalations = v
	return s
}

func (s *PutGroupMetricRuleRequest) SetGroupId(v string) *PutGroupMetricRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetRuleId(v string) *PutGroupMetricRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetCategory(v string) *PutGroupMetricRuleRequest {
	s.Category = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetRuleName(v string) *PutGroupMetricRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetNamespace(v string) *PutGroupMetricRuleRequest {
	s.Namespace = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetMetricName(v string) *PutGroupMetricRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetDimensions(v string) *PutGroupMetricRuleRequest {
	s.Dimensions = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetEffectiveInterval(v string) *PutGroupMetricRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetNoEffectiveInterval(v string) *PutGroupMetricRuleRequest {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetSilenceTime(v int32) *PutGroupMetricRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetPeriod(v string) *PutGroupMetricRuleRequest {
	s.Period = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetInterval(v string) *PutGroupMetricRuleRequest {
	s.Interval = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetWebhook(v string) *PutGroupMetricRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetEmailSubject(v string) *PutGroupMetricRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetContactGroups(v string) *PutGroupMetricRuleRequest {
	s.ContactGroups = &v
	return s
}

type PutGroupMetricRuleRequestEscalations struct {
	Critical *PutGroupMetricRuleRequestEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" require:"true" type:"Struct"`
	Warn     *PutGroupMetricRuleRequestEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" require:"true" type:"Struct"`
	Info     *PutGroupMetricRuleRequestEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" require:"true" type:"Struct"`
}

func (s PutGroupMetricRuleRequestEscalations) String() string {
	return tea.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalations) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalations) SetCritical(v *PutGroupMetricRuleRequestEscalationsCritical) *PutGroupMetricRuleRequestEscalations {
	s.Critical = v
	return s
}

func (s *PutGroupMetricRuleRequestEscalations) SetWarn(v *PutGroupMetricRuleRequestEscalationsWarn) *PutGroupMetricRuleRequestEscalations {
	s.Warn = v
	return s
}

func (s *PutGroupMetricRuleRequestEscalations) SetInfo(v *PutGroupMetricRuleRequestEscalationsInfo) *PutGroupMetricRuleRequestEscalations {
	s.Info = v
	return s
}

type PutGroupMetricRuleRequestEscalationsCritical struct {
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutGroupMetricRuleRequestEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetStatistics(v string) *PutGroupMetricRuleRequestEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetComparisonOperator(v string) *PutGroupMetricRuleRequestEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetThreshold(v string) *PutGroupMetricRuleRequestEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetTimes(v int32) *PutGroupMetricRuleRequestEscalationsCritical {
	s.Times = &v
	return s
}

type PutGroupMetricRuleRequestEscalationsWarn struct {
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutGroupMetricRuleRequestEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetStatistics(v string) *PutGroupMetricRuleRequestEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetComparisonOperator(v string) *PutGroupMetricRuleRequestEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetThreshold(v string) *PutGroupMetricRuleRequestEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetTimes(v int32) *PutGroupMetricRuleRequestEscalationsWarn {
	s.Times = &v
	return s
}

type PutGroupMetricRuleRequestEscalationsInfo struct {
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutGroupMetricRuleRequestEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetStatistics(v string) *PutGroupMetricRuleRequestEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetComparisonOperator(v string) *PutGroupMetricRuleRequestEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetThreshold(v string) *PutGroupMetricRuleRequestEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetTimes(v int32) *PutGroupMetricRuleRequestEscalationsInfo {
	s.Times = &v
	return s
}

type PutGroupMetricRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutGroupMetricRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutGroupMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleResponseBody) SetMessage(v string) *PutGroupMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutGroupMetricRuleResponseBody) SetRequestId(v string) *PutGroupMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutGroupMetricRuleResponseBody) SetCode(v string) *PutGroupMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutGroupMetricRuleResponseBody) SetSuccess(v bool) *PutGroupMetricRuleResponseBody {
	s.Success = &v
	return s
}

type PutGroupMetricRuleResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutGroupMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutGroupMetricRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutGroupMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleResponse) SetHeaders(v map[string]*string) *PutGroupMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *PutGroupMetricRuleResponse) SetBody(v *PutGroupMetricRuleResponseBody) *PutGroupMetricRuleResponse {
	s.Body = v
	return s
}

type PutLogMonitorRequest struct {
	LogId               *string                            `json:"LogId,omitempty" xml:"LogId,omitempty"`
	SlsRegionId         *string                            `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	SlsProject          *string                            `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	SlsLogstore         *string                            `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	MetricName          *string                            `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	MetricExpress       *string                            `json:"MetricExpress,omitempty" xml:"MetricExpress,omitempty"`
	GroupId             *string                            `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ValueFilterRelation *string                            `json:"ValueFilterRelation,omitempty" xml:"ValueFilterRelation,omitempty"`
	Tumblingwindows     *string                            `json:"Tumblingwindows,omitempty" xml:"Tumblingwindows,omitempty"`
	Unit                *string                            `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Aggregates          []*PutLogMonitorRequestAggregates  `json:"Aggregates,omitempty" xml:"Aggregates,omitempty" type:"Repeated"`
	Groupbys            []*PutLogMonitorRequestGroupbys    `json:"Groupbys,omitempty" xml:"Groupbys,omitempty" type:"Repeated"`
	ValueFilter         []*PutLogMonitorRequestValueFilter `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty" type:"Repeated"`
}

func (s PutLogMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s PutLogMonitorRequest) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequest) SetLogId(v string) *PutLogMonitorRequest {
	s.LogId = &v
	return s
}

func (s *PutLogMonitorRequest) SetSlsRegionId(v string) *PutLogMonitorRequest {
	s.SlsRegionId = &v
	return s
}

func (s *PutLogMonitorRequest) SetSlsProject(v string) *PutLogMonitorRequest {
	s.SlsProject = &v
	return s
}

func (s *PutLogMonitorRequest) SetSlsLogstore(v string) *PutLogMonitorRequest {
	s.SlsLogstore = &v
	return s
}

func (s *PutLogMonitorRequest) SetMetricName(v string) *PutLogMonitorRequest {
	s.MetricName = &v
	return s
}

func (s *PutLogMonitorRequest) SetMetricExpress(v string) *PutLogMonitorRequest {
	s.MetricExpress = &v
	return s
}

func (s *PutLogMonitorRequest) SetGroupId(v string) *PutLogMonitorRequest {
	s.GroupId = &v
	return s
}

func (s *PutLogMonitorRequest) SetValueFilterRelation(v string) *PutLogMonitorRequest {
	s.ValueFilterRelation = &v
	return s
}

func (s *PutLogMonitorRequest) SetTumblingwindows(v string) *PutLogMonitorRequest {
	s.Tumblingwindows = &v
	return s
}

func (s *PutLogMonitorRequest) SetUnit(v string) *PutLogMonitorRequest {
	s.Unit = &v
	return s
}

func (s *PutLogMonitorRequest) SetAggregates(v []*PutLogMonitorRequestAggregates) *PutLogMonitorRequest {
	s.Aggregates = v
	return s
}

func (s *PutLogMonitorRequest) SetGroupbys(v []*PutLogMonitorRequestGroupbys) *PutLogMonitorRequest {
	s.Groupbys = v
	return s
}

func (s *PutLogMonitorRequest) SetValueFilter(v []*PutLogMonitorRequestValueFilter) *PutLogMonitorRequest {
	s.ValueFilter = v
	return s
}

type PutLogMonitorRequestAggregates struct {
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Function  *string `json:"Function,omitempty" xml:"Function,omitempty"`
	Alias     *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s PutLogMonitorRequestAggregates) String() string {
	return tea.Prettify(s)
}

func (s PutLogMonitorRequestAggregates) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequestAggregates) SetFieldName(v string) *PutLogMonitorRequestAggregates {
	s.FieldName = &v
	return s
}

func (s *PutLogMonitorRequestAggregates) SetFunction(v string) *PutLogMonitorRequestAggregates {
	s.Function = &v
	return s
}

func (s *PutLogMonitorRequestAggregates) SetAlias(v string) *PutLogMonitorRequestAggregates {
	s.Alias = &v
	return s
}

type PutLogMonitorRequestGroupbys struct {
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Alias     *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s PutLogMonitorRequestGroupbys) String() string {
	return tea.Prettify(s)
}

func (s PutLogMonitorRequestGroupbys) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequestGroupbys) SetFieldName(v string) *PutLogMonitorRequestGroupbys {
	s.FieldName = &v
	return s
}

func (s *PutLogMonitorRequestGroupbys) SetAlias(v string) *PutLogMonitorRequestGroupbys {
	s.Alias = &v
	return s
}

type PutLogMonitorRequestValueFilter struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s PutLogMonitorRequestValueFilter) String() string {
	return tea.Prettify(s)
}

func (s PutLogMonitorRequestValueFilter) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequestValueFilter) SetKey(v string) *PutLogMonitorRequestValueFilter {
	s.Key = &v
	return s
}

func (s *PutLogMonitorRequestValueFilter) SetValue(v string) *PutLogMonitorRequestValueFilter {
	s.Value = &v
	return s
}

func (s *PutLogMonitorRequestValueFilter) SetOperator(v string) *PutLogMonitorRequestValueFilter {
	s.Operator = &v
	return s
}

type PutLogMonitorResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	LogId     *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutLogMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutLogMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *PutLogMonitorResponseBody) SetMessage(v string) *PutLogMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetRequestId(v string) *PutLogMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetCode(v string) *PutLogMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetLogId(v string) *PutLogMonitorResponseBody {
	s.LogId = &v
	return s
}

func (s *PutLogMonitorResponseBody) SetSuccess(v bool) *PutLogMonitorResponseBody {
	s.Success = &v
	return s
}

type PutLogMonitorResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutLogMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutLogMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s PutLogMonitorResponse) GoString() string {
	return s.String()
}

func (s *PutLogMonitorResponse) SetHeaders(v map[string]*string) *PutLogMonitorResponse {
	s.Headers = v
	return s
}

func (s *PutLogMonitorResponse) SetBody(v *PutLogMonitorResponseBody) *PutLogMonitorResponse {
	s.Body = v
	return s
}

type PutMetricRuleTargetsRequest struct {
	RuleId  *string                               `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Targets []*PutMetricRuleTargetsRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s PutMetricRuleTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsRequest) SetRuleId(v string) *PutMetricRuleTargetsRequest {
	s.RuleId = &v
	return s
}

func (s *PutMetricRuleTargetsRequest) SetTargets(v []*PutMetricRuleTargetsRequestTargets) *PutMetricRuleTargetsRequest {
	s.Targets = v
	return s
}

type PutMetricRuleTargetsRequestTargets struct {
	Arn   *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s PutMetricRuleTargetsRequestTargets) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsRequestTargets) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsRequestTargets) SetArn(v string) *PutMetricRuleTargetsRequestTargets {
	s.Arn = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) SetLevel(v string) *PutMetricRuleTargetsRequestTargets {
	s.Level = &v
	return s
}

func (s *PutMetricRuleTargetsRequestTargets) SetId(v string) *PutMetricRuleTargetsRequestTargets {
	s.Id = &v
	return s
}

type PutMetricRuleTargetsResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailData  *PutMetricRuleTargetsResponseBodyFailData `json:"FailData,omitempty" xml:"FailData,omitempty" type:"Struct"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutMetricRuleTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBody) SetMessage(v string) *PutMetricRuleTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetRequestId(v string) *PutMetricRuleTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetFailData(v *PutMetricRuleTargetsResponseBodyFailData) *PutMetricRuleTargetsResponseBody {
	s.FailData = v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetCode(v string) *PutMetricRuleTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBody) SetSuccess(v bool) *PutMetricRuleTargetsResponseBody {
	s.Success = &v
	return s
}

type PutMetricRuleTargetsResponseBodyFailData struct {
	Targets *PutMetricRuleTargetsResponseBodyFailDataTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Struct"`
}

func (s PutMetricRuleTargetsResponseBodyFailData) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailData) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailData) SetTargets(v *PutMetricRuleTargetsResponseBodyFailDataTargets) *PutMetricRuleTargetsResponseBodyFailData {
	s.Targets = v
	return s
}

type PutMetricRuleTargetsResponseBodyFailDataTargets struct {
	Target []*PutMetricRuleTargetsResponseBodyFailDataTargetsTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargets) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargets) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargets) SetTarget(v []*PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) *PutMetricRuleTargetsResponseBodyFailDataTargets {
	s.Target = v
	return s
}

type PutMetricRuleTargetsResponseBodyFailDataTargetsTarget struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Arn   *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetId(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Id = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetArn(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Arn = &v
	return s
}

func (s *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget) SetLevel(v string) *PutMetricRuleTargetsResponseBodyFailDataTargetsTarget {
	s.Level = &v
	return s
}

type PutMetricRuleTargetsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutMetricRuleTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMetricRuleTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMetricRuleTargetsResponse) GoString() string {
	return s.String()
}

func (s *PutMetricRuleTargetsResponse) SetHeaders(v map[string]*string) *PutMetricRuleTargetsResponse {
	s.Headers = v
	return s
}

func (s *PutMetricRuleTargetsResponse) SetBody(v *PutMetricRuleTargetsResponseBody) *PutMetricRuleTargetsResponse {
	s.Body = v
	return s
}

type PutMonitorGroupDynamicRuleRequest struct {
	GroupId    *int64                                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupRules []*PutMonitorGroupDynamicRuleRequestGroupRules `json:"GroupRules,omitempty" xml:"GroupRules,omitempty" type:"Repeated"`
}

func (s PutMonitorGroupDynamicRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleRequest) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleRequest) SetGroupId(v int64) *PutMonitorGroupDynamicRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequest) SetGroupRules(v []*PutMonitorGroupDynamicRuleRequestGroupRules) *PutMonitorGroupDynamicRuleRequest {
	s.GroupRules = v
	return s
}

type PutMonitorGroupDynamicRuleRequestGroupRules struct {
	FilterRelation *string                                               `json:"FilterRelation,omitempty" xml:"FilterRelation,omitempty"`
	Filters        []*PutMonitorGroupDynamicRuleRequestGroupRulesFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	Category       *string                                               `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s PutMonitorGroupDynamicRuleRequestGroupRules) String() string {
	return tea.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleRequestGroupRules) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) SetFilterRelation(v string) *PutMonitorGroupDynamicRuleRequestGroupRules {
	s.FilterRelation = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) SetFilters(v []*PutMonitorGroupDynamicRuleRequestGroupRulesFilters) *PutMonitorGroupDynamicRuleRequestGroupRules {
	s.Filters = v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) SetCategory(v string) *PutMonitorGroupDynamicRuleRequestGroupRules {
	s.Category = &v
	return s
}

type PutMonitorGroupDynamicRuleRequestGroupRulesFilters struct {
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s PutMonitorGroupDynamicRuleRequestGroupRulesFilters) String() string {
	return tea.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleRequestGroupRulesFilters) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) SetValue(v string) *PutMonitorGroupDynamicRuleRequestGroupRulesFilters {
	s.Value = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) SetFunction(v string) *PutMonitorGroupDynamicRuleRequestGroupRulesFilters {
	s.Function = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) SetName(v string) *PutMonitorGroupDynamicRuleRequestGroupRulesFilters {
	s.Name = &v
	return s
}

type PutMonitorGroupDynamicRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutMonitorGroupDynamicRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetMessage(v string) *PutMonitorGroupDynamicRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetRequestId(v string) *PutMonitorGroupDynamicRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetCode(v int32) *PutMonitorGroupDynamicRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponseBody) SetSuccess(v bool) *PutMonitorGroupDynamicRuleResponseBody {
	s.Success = &v
	return s
}

type PutMonitorGroupDynamicRuleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutMonitorGroupDynamicRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMonitorGroupDynamicRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleResponse) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleResponse) SetHeaders(v map[string]*string) *PutMonitorGroupDynamicRuleResponse {
	s.Headers = v
	return s
}

func (s *PutMonitorGroupDynamicRuleResponse) SetBody(v *PutMonitorGroupDynamicRuleResponseBody) *PutMonitorGroupDynamicRuleResponse {
	s.Body = v
	return s
}

type PutMonitoringConfigRequest struct {
	AutoInstall              *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	EnableInstallAgentNewECS *bool `json:"EnableInstallAgentNewECS,omitempty" xml:"EnableInstallAgentNewECS,omitempty"`
}

func (s PutMonitoringConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s PutMonitoringConfigRequest) GoString() string {
	return s.String()
}

func (s *PutMonitoringConfigRequest) SetAutoInstall(v bool) *PutMonitoringConfigRequest {
	s.AutoInstall = &v
	return s
}

func (s *PutMonitoringConfigRequest) SetEnableInstallAgentNewECS(v bool) *PutMonitoringConfigRequest {
	s.EnableInstallAgentNewECS = &v
	return s
}

type PutMonitoringConfigResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutMonitoringConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutMonitoringConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PutMonitoringConfigResponseBody) SetMessage(v string) *PutMonitoringConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PutMonitoringConfigResponseBody) SetRequestId(v string) *PutMonitoringConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutMonitoringConfigResponseBody) SetCode(v int32) *PutMonitoringConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PutMonitoringConfigResponseBody) SetSuccess(v bool) *PutMonitoringConfigResponseBody {
	s.Success = &v
	return s
}

type PutMonitoringConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutMonitoringConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutMonitoringConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s PutMonitoringConfigResponse) GoString() string {
	return s.String()
}

func (s *PutMonitoringConfigResponse) SetHeaders(v map[string]*string) *PutMonitoringConfigResponse {
	s.Headers = v
	return s
}

func (s *PutMonitoringConfigResponse) SetBody(v *PutMonitoringConfigResponseBody) *PutMonitoringConfigResponse {
	s.Body = v
	return s
}

type PutResourceMetricRuleRequest struct {
	Escalations         *PutResourceMetricRuleRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	RuleId              *string                                  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName            *string                                  `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Namespace           *string                                  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	MetricName          *string                                  `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Resources           *string                                  `json:"Resources,omitempty" xml:"Resources,omitempty"`
	ContactGroups       *string                                  `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Webhook             *string                                  `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	EffectiveInterval   *string                                  `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	NoEffectiveInterval *string                                  `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	SilenceTime         *int32                                   `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Period              *string                                  `json:"Period,omitempty" xml:"Period,omitempty"`
	Interval            *string                                  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	EmailSubject        *string                                  `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
}

func (s PutResourceMetricRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequest) SetEscalations(v *PutResourceMetricRuleRequestEscalations) *PutResourceMetricRuleRequest {
	s.Escalations = v
	return s
}

func (s *PutResourceMetricRuleRequest) SetRuleId(v string) *PutResourceMetricRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetRuleName(v string) *PutResourceMetricRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetNamespace(v string) *PutResourceMetricRuleRequest {
	s.Namespace = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetMetricName(v string) *PutResourceMetricRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetResources(v string) *PutResourceMetricRuleRequest {
	s.Resources = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetContactGroups(v string) *PutResourceMetricRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetWebhook(v string) *PutResourceMetricRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetEffectiveInterval(v string) *PutResourceMetricRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetNoEffectiveInterval(v string) *PutResourceMetricRuleRequest {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetSilenceTime(v int32) *PutResourceMetricRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetPeriod(v string) *PutResourceMetricRuleRequest {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetInterval(v string) *PutResourceMetricRuleRequest {
	s.Interval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetEmailSubject(v string) *PutResourceMetricRuleRequest {
	s.EmailSubject = &v
	return s
}

type PutResourceMetricRuleRequestEscalations struct {
	Critical *PutResourceMetricRuleRequestEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" require:"true" type:"Struct"`
	Warn     *PutResourceMetricRuleRequestEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" require:"true" type:"Struct"`
	Info     *PutResourceMetricRuleRequestEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" require:"true" type:"Struct"`
}

func (s PutResourceMetricRuleRequestEscalations) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalations) SetCritical(v *PutResourceMetricRuleRequestEscalationsCritical) *PutResourceMetricRuleRequestEscalations {
	s.Critical = v
	return s
}

func (s *PutResourceMetricRuleRequestEscalations) SetWarn(v *PutResourceMetricRuleRequestEscalationsWarn) *PutResourceMetricRuleRequestEscalations {
	s.Warn = v
	return s
}

func (s *PutResourceMetricRuleRequestEscalations) SetInfo(v *PutResourceMetricRuleRequestEscalationsInfo) *PutResourceMetricRuleRequestEscalations {
	s.Info = v
	return s
}

type PutResourceMetricRuleRequestEscalationsCritical struct {
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Times = &v
	return s
}

type PutResourceMetricRuleRequestEscalationsWarn struct {
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Times = &v
	return s
}

type PutResourceMetricRuleRequestEscalationsInfo struct {
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Times = &v
	return s
}

type PutResourceMetricRuleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutResourceMetricRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleResponseBody) SetMessage(v string) *PutResourceMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetRequestId(v string) *PutResourceMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetCode(v string) *PutResourceMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetSuccess(v bool) *PutResourceMetricRuleResponseBody {
	s.Success = &v
	return s
}

type PutResourceMetricRuleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutResourceMetricRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutResourceMetricRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRuleResponse) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleResponse) SetHeaders(v map[string]*string) *PutResourceMetricRuleResponse {
	s.Headers = v
	return s
}

func (s *PutResourceMetricRuleResponse) SetBody(v *PutResourceMetricRuleResponseBody) *PutResourceMetricRuleResponse {
	s.Body = v
	return s
}

type PutResourceMetricRulesRequest struct {
	Rules []*PutResourceMetricRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s PutResourceMetricRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequest) SetRules(v []*PutResourceMetricRulesRequestRules) *PutResourceMetricRulesRequest {
	s.Rules = v
	return s
}

type PutResourceMetricRulesRequestRules struct {
	Escalations         *PutResourceMetricRulesRequestRulesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" require:"true" type:"Struct"`
	MetricName          *string                                        `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	NoEffectiveInterval *string                                        `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	EffectiveInterval   *string                                        `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	RuleId              *string                                        `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Resources           *string                                        `json:"Resources,omitempty" xml:"Resources,omitempty"`
	SilenceTime         *int32                                         `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	Webhook             *string                                        `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
	ContactGroups       *string                                        `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	Namespace           *string                                        `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	EmailSubject        *string                                        `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	Period              *string                                        `json:"Period,omitempty" xml:"Period,omitempty"`
	RuleName            *string                                        `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Interval            *string                                        `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s PutResourceMetricRulesRequestRules) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesRequestRules) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRules) SetEscalations(v *PutResourceMetricRulesRequestRulesEscalations) *PutResourceMetricRulesRequestRules {
	s.Escalations = v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetMetricName(v string) *PutResourceMetricRulesRequestRules {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNoEffectiveInterval(v string) *PutResourceMetricRulesRequestRules {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetEffectiveInterval(v string) *PutResourceMetricRulesRequestRules {
	s.EffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetRuleId(v string) *PutResourceMetricRulesRequestRules {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetResources(v string) *PutResourceMetricRulesRequestRules {
	s.Resources = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetSilenceTime(v int32) *PutResourceMetricRulesRequestRules {
	s.SilenceTime = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetWebhook(v string) *PutResourceMetricRulesRequestRules {
	s.Webhook = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetContactGroups(v string) *PutResourceMetricRulesRequestRules {
	s.ContactGroups = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNamespace(v string) *PutResourceMetricRulesRequestRules {
	s.Namespace = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetEmailSubject(v string) *PutResourceMetricRulesRequestRules {
	s.EmailSubject = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetPeriod(v string) *PutResourceMetricRulesRequestRules {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetRuleName(v string) *PutResourceMetricRulesRequestRules {
	s.RuleName = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetInterval(v string) *PutResourceMetricRulesRequestRules {
	s.Interval = &v
	return s
}

type PutResourceMetricRulesRequestRulesEscalations struct {
	Info     *PutResourceMetricRulesRequestRulesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" require:"true" type:"Struct"`
	Warn     *PutResourceMetricRulesRequestRulesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" require:"true" type:"Struct"`
	Critical *PutResourceMetricRulesRequestRulesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" require:"true" type:"Struct"`
}

func (s PutResourceMetricRulesRequestRulesEscalations) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetInfo(v *PutResourceMetricRulesRequestRulesEscalationsInfo) *PutResourceMetricRulesRequestRulesEscalations {
	s.Info = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetWarn(v *PutResourceMetricRulesRequestRulesEscalationsWarn) *PutResourceMetricRulesRequestRulesEscalations {
	s.Warn = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetCritical(v *PutResourceMetricRulesRequestRulesEscalationsCritical) *PutResourceMetricRulesRequestRulesEscalations {
	s.Critical = v
	return s
}

type PutResourceMetricRulesRequestRulesEscalationsInfo struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsInfo) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

type PutResourceMetricRulesRequestRulesEscalationsWarn struct {
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsWarn) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Statistics = &v
	return s
}

type PutResourceMetricRulesRequestRulesEscalationsCritical struct {
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsCritical) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

type PutResourceMetricRulesResponseBody struct {
	Message          *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedListResult *PutResourceMetricRulesResponseBodyFailedListResult `json:"FailedListResult,omitempty" xml:"FailedListResult,omitempty" type:"Struct"`
	Code             *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success          *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutResourceMetricRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBody) SetMessage(v string) *PutResourceMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetRequestId(v string) *PutResourceMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetFailedListResult(v *PutResourceMetricRulesResponseBodyFailedListResult) *PutResourceMetricRulesResponseBody {
	s.FailedListResult = v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetCode(v string) *PutResourceMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *PutResourceMetricRulesResponseBody) SetSuccess(v bool) *PutResourceMetricRulesResponseBody {
	s.Success = &v
	return s
}

type PutResourceMetricRulesResponseBodyFailedListResult struct {
	Target []*PutResourceMetricRulesResponseBodyFailedListResultTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s PutResourceMetricRulesResponseBodyFailedListResult) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesResponseBodyFailedListResult) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBodyFailedListResult) SetTarget(v []*PutResourceMetricRulesResponseBodyFailedListResultTarget) *PutResourceMetricRulesResponseBodyFailedListResult {
	s.Target = v
	return s
}

type PutResourceMetricRulesResponseBodyFailedListResultTarget struct {
	Result *PutResourceMetricRulesResponseBodyFailedListResultTargetResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	RuleId *string                                                         `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTarget) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTarget) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTarget) SetResult(v *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) *PutResourceMetricRulesResponseBodyFailedListResultTarget {
	s.Result = v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTarget) SetRuleId(v string) *PutResourceMetricRulesResponseBodyFailedListResultTarget {
	s.RuleId = &v
	return s
}

type PutResourceMetricRulesResponseBodyFailedListResultTargetResult struct {
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTargetResult) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesResponseBodyFailedListResultTargetResult) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) SetSuccess(v bool) *PutResourceMetricRulesResponseBodyFailedListResultTargetResult {
	s.Success = &v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) SetCode(v string) *PutResourceMetricRulesResponseBodyFailedListResultTargetResult {
	s.Code = &v
	return s
}

func (s *PutResourceMetricRulesResponseBodyFailedListResultTargetResult) SetMessage(v string) *PutResourceMetricRulesResponseBodyFailedListResultTargetResult {
	s.Message = &v
	return s
}

type PutResourceMetricRulesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutResourceMetricRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutResourceMetricRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s PutResourceMetricRulesResponse) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesResponse) SetHeaders(v map[string]*string) *PutResourceMetricRulesResponse {
	s.Headers = v
	return s
}

func (s *PutResourceMetricRulesResponse) SetBody(v *PutResourceMetricRulesResponseBody) *PutResourceMetricRulesResponse {
	s.Body = v
	return s
}

type RemoveTagsRequest struct {
	Tag      []*RemoveTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	GroupIds []*string               `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s RemoveTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequest) SetTag(v []*RemoveTagsRequestTag) *RemoveTagsRequest {
	s.Tag = v
	return s
}

func (s *RemoveTagsRequest) SetGroupIds(v []*string) *RemoveTagsRequest {
	s.GroupIds = v
	return s
}

type RemoveTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RemoveTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsRequestTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsRequestTag) SetKey(v string) *RemoveTagsRequestTag {
	s.Key = &v
	return s
}

func (s *RemoveTagsRequestTag) SetValue(v string) *RemoveTagsRequestTag {
	s.Value = &v
	return s
}

type RemoveTagsResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tag       *RemoveTagsResponseBodyTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBody) SetMessage(v string) *RemoveTagsResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveTagsResponseBody) SetRequestId(v string) *RemoveTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTagsResponseBody) SetTag(v *RemoveTagsResponseBodyTag) *RemoveTagsResponseBody {
	s.Tag = v
	return s
}

func (s *RemoveTagsResponseBody) SetCode(v string) *RemoveTagsResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveTagsResponseBody) SetSuccess(v bool) *RemoveTagsResponseBody {
	s.Success = &v
	return s
}

type RemoveTagsResponseBodyTag struct {
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RemoveTagsResponseBodyTag) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponseBodyTag) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponseBodyTag) SetTags(v []*string) *RemoveTagsResponseBodyTag {
	s.Tags = v
	return s
}

type RemoveTagsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTagsResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponse) SetHeaders(v map[string]*string) *RemoveTagsResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagsResponse) SetBody(v *RemoveTagsResponseBody) *RemoveTagsResponse {
	s.Body = v
	return s
}

type SendDryRunSystemEventRequest struct {
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	EventName    *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	EventContent *string `json:"EventContent,omitempty" xml:"EventContent,omitempty"`
}

func (s SendDryRunSystemEventRequest) String() string {
	return tea.Prettify(s)
}

func (s SendDryRunSystemEventRequest) GoString() string {
	return s.String()
}

func (s *SendDryRunSystemEventRequest) SetProduct(v string) *SendDryRunSystemEventRequest {
	s.Product = &v
	return s
}

func (s *SendDryRunSystemEventRequest) SetEventName(v string) *SendDryRunSystemEventRequest {
	s.EventName = &v
	return s
}

func (s *SendDryRunSystemEventRequest) SetGroupId(v string) *SendDryRunSystemEventRequest {
	s.GroupId = &v
	return s
}

func (s *SendDryRunSystemEventRequest) SetEventContent(v string) *SendDryRunSystemEventRequest {
	s.EventContent = &v
	return s
}

type SendDryRunSystemEventResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendDryRunSystemEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendDryRunSystemEventResponseBody) GoString() string {
	return s.String()
}

func (s *SendDryRunSystemEventResponseBody) SetMessage(v string) *SendDryRunSystemEventResponseBody {
	s.Message = &v
	return s
}

func (s *SendDryRunSystemEventResponseBody) SetRequestId(v string) *SendDryRunSystemEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendDryRunSystemEventResponseBody) SetCode(v string) *SendDryRunSystemEventResponseBody {
	s.Code = &v
	return s
}

func (s *SendDryRunSystemEventResponseBody) SetSuccess(v string) *SendDryRunSystemEventResponseBody {
	s.Success = &v
	return s
}

type SendDryRunSystemEventResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendDryRunSystemEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendDryRunSystemEventResponse) String() string {
	return tea.Prettify(s)
}

func (s SendDryRunSystemEventResponse) GoString() string {
	return s.String()
}

func (s *SendDryRunSystemEventResponse) SetHeaders(v map[string]*string) *SendDryRunSystemEventResponse {
	s.Headers = v
	return s
}

func (s *SendDryRunSystemEventResponse) SetBody(v *SendDryRunSystemEventResponseBody) *SendDryRunSystemEventResponse {
	s.Body = v
	return s
}

type UninstallMonitoringAgentRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UninstallMonitoringAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallMonitoringAgentRequest) GoString() string {
	return s.String()
}

func (s *UninstallMonitoringAgentRequest) SetInstanceId(v string) *UninstallMonitoringAgentRequest {
	s.InstanceId = &v
	return s
}

type UninstallMonitoringAgentResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UninstallMonitoringAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallMonitoringAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallMonitoringAgentResponseBody) SetMessage(v string) *UninstallMonitoringAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallMonitoringAgentResponseBody) SetRequestId(v string) *UninstallMonitoringAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallMonitoringAgentResponseBody) SetCode(v string) *UninstallMonitoringAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallMonitoringAgentResponseBody) SetSuccess(v bool) *UninstallMonitoringAgentResponseBody {
	s.Success = &v
	return s
}

type UninstallMonitoringAgentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallMonitoringAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallMonitoringAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallMonitoringAgentResponse) GoString() string {
	return s.String()
}

func (s *UninstallMonitoringAgentResponse) SetHeaders(v map[string]*string) *UninstallMonitoringAgentResponse {
	s.Headers = v
	return s
}

func (s *UninstallMonitoringAgentResponse) SetBody(v *UninstallMonitoringAgentResponseBody) *UninstallMonitoringAgentResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddTagsWithOptions(request *AddTagsRequest, runtime *util.RuntimeOptions) (_result *AddTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddTags"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTags(request *AddTagsRequest) (_result *AddTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTagsResponse{}
	_body, _err := client.AddTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyMetricRuleTemplateWithOptions(request *ApplyMetricRuleTemplateRequest, runtime *util.RuntimeOptions) (_result *ApplyMetricRuleTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyMetricRuleTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyMetricRuleTemplate"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyMetricRuleTemplate(request *ApplyMetricRuleTemplateRequest) (_result *ApplyMetricRuleTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyMetricRuleTemplateResponse{}
	_body, _err := client.ApplyMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCmsCallNumOrderWithOptions(request *CreateCmsCallNumOrderRequest, runtime *util.RuntimeOptions) (_result *CreateCmsCallNumOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCmsCallNumOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCmsCallNumOrder"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCmsCallNumOrder(request *CreateCmsCallNumOrderRequest) (_result *CreateCmsCallNumOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCmsCallNumOrderResponse{}
	_body, _err := client.CreateCmsCallNumOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCmsOrderWithOptions(request *CreateCmsOrderRequest, runtime *util.RuntimeOptions) (_result *CreateCmsOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCmsOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCmsOrder"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCmsOrder(request *CreateCmsOrderRequest) (_result *CreateCmsOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCmsOrderResponse{}
	_body, _err := client.CreateCmsOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCmsSmspackageOrderWithOptions(request *CreateCmsSmspackageOrderRequest, runtime *util.RuntimeOptions) (_result *CreateCmsSmspackageOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCmsSmspackageOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCmsSmspackageOrder"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCmsSmspackageOrder(request *CreateCmsSmspackageOrderRequest) (_result *CreateCmsSmspackageOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCmsSmspackageOrderResponse{}
	_body, _err := client.CreateCmsSmspackageOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDynamicTagGroupWithOptions(request *CreateDynamicTagGroupRequest, runtime *util.RuntimeOptions) (_result *CreateDynamicTagGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDynamicTagGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDynamicTagGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDynamicTagGroup(request *CreateDynamicTagGroupRequest) (_result *CreateDynamicTagGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDynamicTagGroupResponse{}
	_body, _err := client.CreateDynamicTagGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupMetricRulesWithOptions(request *CreateGroupMetricRulesRequest, runtime *util.RuntimeOptions) (_result *CreateGroupMetricRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGroupMetricRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGroupMetricRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupMetricRules(request *CreateGroupMetricRulesRequest) (_result *CreateGroupMetricRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupMetricRulesResponse{}
	_body, _err := client.CreateGroupMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupMonitoringAgentProcessWithOptions(request *CreateGroupMonitoringAgentProcessRequest, runtime *util.RuntimeOptions) (_result *CreateGroupMonitoringAgentProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGroupMonitoringAgentProcess"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGroupMonitoringAgentProcess(request *CreateGroupMonitoringAgentProcessRequest) (_result *CreateGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CreateGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHostAvailabilityWithOptions(request *CreateHostAvailabilityRequest, runtime *util.RuntimeOptions) (_result *CreateHostAvailabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHostAvailabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHostAvailability"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHostAvailability(request *CreateHostAvailabilityRequest) (_result *CreateHostAvailabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHostAvailabilityResponse{}
	_body, _err := client.CreateHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMetricRuleResourcesWithOptions(request *CreateMetricRuleResourcesRequest, runtime *util.RuntimeOptions) (_result *CreateMetricRuleResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMetricRuleResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMetricRuleResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMetricRuleResources(request *CreateMetricRuleResourcesRequest) (_result *CreateMetricRuleResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMetricRuleResourcesResponse{}
	_body, _err := client.CreateMetricRuleResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMetricRuleTemplateWithOptions(request *CreateMetricRuleTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateMetricRuleTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMetricRuleTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMetricRuleTemplate"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMetricRuleTemplate(request *CreateMetricRuleTemplateRequest) (_result *CreateMetricRuleTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMetricRuleTemplateResponse{}
	_body, _err := client.CreateMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitorAgentProcessWithOptions(request *CreateMonitorAgentProcessRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorAgentProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMonitorAgentProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMonitorAgentProcess"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitorAgentProcess(request *CreateMonitorAgentProcessRequest) (_result *CreateMonitorAgentProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorAgentProcessResponse{}
	_body, _err := client.CreateMonitorAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitorGroupWithOptions(request *CreateMonitorGroupRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMonitorGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitorGroup(request *CreateMonitorGroupRequest) (_result *CreateMonitorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CreateMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitorGroupByResourceGroupIdWithOptions(request *CreateMonitorGroupByResourceGroupIdRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorGroupByResourceGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMonitorGroupByResourceGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMonitorGroupByResourceGroupId"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitorGroupByResourceGroupId(request *CreateMonitorGroupByResourceGroupIdRequest) (_result *CreateMonitorGroupByResourceGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorGroupByResourceGroupIdResponse{}
	_body, _err := client.CreateMonitorGroupByResourceGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitorGroupInstancesWithOptions(request *CreateMonitorGroupInstancesRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorGroupInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMonitorGroupInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMonitorGroupInstances"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitorGroupInstances(request *CreateMonitorGroupInstancesRequest) (_result *CreateMonitorGroupInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorGroupInstancesResponse{}
	_body, _err := client.CreateMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitorGroupNotifyPolicyWithOptions(request *CreateMonitorGroupNotifyPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorGroupNotifyPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMonitorGroupNotifyPolicy"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitorGroupNotifyPolicy(request *CreateMonitorGroupNotifyPolicyRequest) (_result *CreateMonitorGroupNotifyPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.CreateMonitorGroupNotifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitoringAgentProcessWithOptions(request *CreateMonitoringAgentProcessRequest, runtime *util.RuntimeOptions) (_result *CreateMonitoringAgentProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMonitoringAgentProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMonitoringAgentProcess"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitoringAgentProcess(request *CreateMonitoringAgentProcessRequest) (_result *CreateMonitoringAgentProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitoringAgentProcessResponse{}
	_body, _err := client.CreateMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSiteMonitorWithOptions(request *CreateSiteMonitorRequest, runtime *util.RuntimeOptions) (_result *CreateSiteMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSiteMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSiteMonitor"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSiteMonitor(request *CreateSiteMonitorRequest) (_result *CreateSiteMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSiteMonitorResponse{}
	_body, _err := client.CreateSiteMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactWithOptions(request *DeleteContactRequest, runtime *util.RuntimeOptions) (_result *DeleteContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteContactResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteContact"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContact(request *DeleteContactRequest) (_result *DeleteContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContactResponse{}
	_body, _err := client.DeleteContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactGroupWithOptions(request *DeleteContactGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteContactGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContactGroup(request *DeleteContactGroupRequest) (_result *DeleteContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.DeleteContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCustomMetricWithOptions(request *DeleteCustomMetricRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCustomMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCustomMetric"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCustomMetric(request *DeleteCustomMetricRequest) (_result *DeleteCustomMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomMetricResponse{}
	_body, _err := client.DeleteCustomMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDynamicTagGroupWithOptions(request *DeleteDynamicTagGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteDynamicTagGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDynamicTagGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDynamicTagGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDynamicTagGroup(request *DeleteDynamicTagGroupRequest) (_result *DeleteDynamicTagGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDynamicTagGroupResponse{}
	_body, _err := client.DeleteDynamicTagGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventRulesWithOptions(request *DeleteEventRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteEventRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEventRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEventRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEventRules(request *DeleteEventRulesRequest) (_result *DeleteEventRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventRulesResponse{}
	_body, _err := client.DeleteEventRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventRuleTargetsWithOptions(request *DeleteEventRuleTargetsRequest, runtime *util.RuntimeOptions) (_result *DeleteEventRuleTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEventRuleTargetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEventRuleTargets"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEventRuleTargets(request *DeleteEventRuleTargetsRequest) (_result *DeleteEventRuleTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventRuleTargetsResponse{}
	_body, _err := client.DeleteEventRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExporterOutputWithOptions(request *DeleteExporterOutputRequest, runtime *util.RuntimeOptions) (_result *DeleteExporterOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteExporterOutputResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteExporterOutput"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExporterOutput(request *DeleteExporterOutputRequest) (_result *DeleteExporterOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExporterOutputResponse{}
	_body, _err := client.DeleteExporterOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExporterRuleWithOptions(request *DeleteExporterRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteExporterRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteExporterRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteExporterRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExporterRule(request *DeleteExporterRuleRequest) (_result *DeleteExporterRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExporterRuleResponse{}
	_body, _err := client.DeleteExporterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupMonitoringAgentProcessWithOptions(request *DeleteGroupMonitoringAgentProcessRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupMonitoringAgentProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGroupMonitoringAgentProcess"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroupMonitoringAgentProcess(request *DeleteGroupMonitoringAgentProcessRequest) (_result *DeleteGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DeleteGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHostAvailabilityWithOptions(request *DeleteHostAvailabilityRequest, runtime *util.RuntimeOptions) (_result *DeleteHostAvailabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHostAvailabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHostAvailability"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHostAvailability(request *DeleteHostAvailabilityRequest) (_result *DeleteHostAvailabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHostAvailabilityResponse{}
	_body, _err := client.DeleteHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLogMonitorWithOptions(request *DeleteLogMonitorRequest, runtime *util.RuntimeOptions) (_result *DeleteLogMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLogMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLogMonitor"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogMonitor(request *DeleteLogMonitorRequest) (_result *DeleteLogMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLogMonitorResponse{}
	_body, _err := client.DeleteLogMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMetricRuleResourcesWithOptions(request *DeleteMetricRuleResourcesRequest, runtime *util.RuntimeOptions) (_result *DeleteMetricRuleResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMetricRuleResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMetricRuleResources"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMetricRuleResources(request *DeleteMetricRuleResourcesRequest) (_result *DeleteMetricRuleResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMetricRuleResourcesResponse{}
	_body, _err := client.DeleteMetricRuleResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMetricRulesWithOptions(request *DeleteMetricRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteMetricRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMetricRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMetricRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMetricRules(request *DeleteMetricRulesRequest) (_result *DeleteMetricRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMetricRulesResponse{}
	_body, _err := client.DeleteMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMetricRuleTargetsWithOptions(request *DeleteMetricRuleTargetsRequest, runtime *util.RuntimeOptions) (_result *DeleteMetricRuleTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMetricRuleTargetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMetricRuleTargets"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMetricRuleTargets(request *DeleteMetricRuleTargetsRequest) (_result *DeleteMetricRuleTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMetricRuleTargetsResponse{}
	_body, _err := client.DeleteMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMetricRuleTemplateWithOptions(request *DeleteMetricRuleTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteMetricRuleTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMetricRuleTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMetricRuleTemplate"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMetricRuleTemplate(request *DeleteMetricRuleTemplateRequest) (_result *DeleteMetricRuleTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMetricRuleTemplateResponse{}
	_body, _err := client.DeleteMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMonitorGroupWithOptions(request *DeleteMonitorGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteMonitorGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMonitorGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMonitorGroup(request *DeleteMonitorGroupRequest) (_result *DeleteMonitorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.DeleteMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMonitorGroupDynamicRuleWithOptions(request *DeleteMonitorGroupDynamicRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteMonitorGroupDynamicRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMonitorGroupDynamicRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMonitorGroupDynamicRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMonitorGroupDynamicRule(request *DeleteMonitorGroupDynamicRuleRequest) (_result *DeleteMonitorGroupDynamicRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMonitorGroupDynamicRuleResponse{}
	_body, _err := client.DeleteMonitorGroupDynamicRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMonitorGroupInstancesWithOptions(request *DeleteMonitorGroupInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteMonitorGroupInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMonitorGroupInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMonitorGroupInstances"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMonitorGroupInstances(request *DeleteMonitorGroupInstancesRequest) (_result *DeleteMonitorGroupInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMonitorGroupInstancesResponse{}
	_body, _err := client.DeleteMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMonitorGroupNotifyPolicyWithOptions(request *DeleteMonitorGroupNotifyPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteMonitorGroupNotifyPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMonitorGroupNotifyPolicy"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMonitorGroupNotifyPolicy(request *DeleteMonitorGroupNotifyPolicyRequest) (_result *DeleteMonitorGroupNotifyPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.DeleteMonitorGroupNotifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMonitoringAgentProcessWithOptions(request *DeleteMonitoringAgentProcessRequest, runtime *util.RuntimeOptions) (_result *DeleteMonitoringAgentProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMonitoringAgentProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMonitoringAgentProcess"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMonitoringAgentProcess(request *DeleteMonitoringAgentProcessRequest) (_result *DeleteMonitoringAgentProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMonitoringAgentProcessResponse{}
	_body, _err := client.DeleteMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSiteMonitorsWithOptions(request *DeleteSiteMonitorsRequest, runtime *util.RuntimeOptions) (_result *DeleteSiteMonitorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSiteMonitorsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSiteMonitors"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSiteMonitors(request *DeleteSiteMonitorsRequest) (_result *DeleteSiteMonitorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSiteMonitorsResponse{}
	_body, _err := client.DeleteSiteMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeActiveMetricRuleListWithOptions(request *DescribeActiveMetricRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeActiveMetricRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeActiveMetricRuleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeActiveMetricRuleList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeActiveMetricRuleList(request *DescribeActiveMetricRuleListRequest) (_result *DescribeActiveMetricRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeActiveMetricRuleListResponse{}
	_body, _err := client.DescribeActiveMetricRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertHistoryListWithOptions(request *DescribeAlertHistoryListRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertHistoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlertHistoryListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlertHistoryList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertHistoryList(request *DescribeAlertHistoryListRequest) (_result *DescribeAlertHistoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertHistoryListResponse{}
	_body, _err := client.DescribeAlertHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertLogCountWithOptions(request *DescribeAlertLogCountRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertLogCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlertLogCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlertLogCount"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertLogCount(request *DescribeAlertLogCountRequest) (_result *DescribeAlertLogCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertLogCountResponse{}
	_body, _err := client.DescribeAlertLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertLogHistogramWithOptions(request *DescribeAlertLogHistogramRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertLogHistogramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlertLogHistogramResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlertLogHistogram"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertLogHistogram(request *DescribeAlertLogHistogramRequest) (_result *DescribeAlertLogHistogramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertLogHistogramResponse{}
	_body, _err := client.DescribeAlertLogHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlertLogListWithOptions(request *DescribeAlertLogListRequest, runtime *util.RuntimeOptions) (_result *DescribeAlertLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlertLogListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlertLogList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlertLogList(request *DescribeAlertLogListRequest) (_result *DescribeAlertLogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlertLogListResponse{}
	_body, _err := client.DescribeAlertLogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContactGroupListWithOptions(request *DescribeContactGroupListRequest, runtime *util.RuntimeOptions) (_result *DescribeContactGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContactGroupListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContactGroupList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContactGroupList(request *DescribeContactGroupListRequest) (_result *DescribeContactGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContactGroupListResponse{}
	_body, _err := client.DescribeContactGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContactListWithOptions(request *DescribeContactListRequest, runtime *util.RuntimeOptions) (_result *DescribeContactListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContactListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContactList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContactList(request *DescribeContactListRequest) (_result *DescribeContactListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContactListResponse{}
	_body, _err := client.DescribeContactListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContactListByContactGroupWithOptions(request *DescribeContactListByContactGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeContactListByContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContactListByContactGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContactListByContactGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContactListByContactGroup(request *DescribeContactListByContactGroupRequest) (_result *DescribeContactListByContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContactListByContactGroupResponse{}
	_body, _err := client.DescribeContactListByContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomEventAttributeWithOptions(request *DescribeCustomEventAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomEventAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCustomEventAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomEventAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomEventAttribute(request *DescribeCustomEventAttributeRequest) (_result *DescribeCustomEventAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomEventAttributeResponse{}
	_body, _err := client.DescribeCustomEventAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomEventCountWithOptions(request *DescribeCustomEventCountRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomEventCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCustomEventCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomEventCount"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomEventCount(request *DescribeCustomEventCountRequest) (_result *DescribeCustomEventCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomEventCountResponse{}
	_body, _err := client.DescribeCustomEventCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomEventHistogramWithOptions(request *DescribeCustomEventHistogramRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomEventHistogramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCustomEventHistogramResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomEventHistogram"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomEventHistogram(request *DescribeCustomEventHistogramRequest) (_result *DescribeCustomEventHistogramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomEventHistogramResponse{}
	_body, _err := client.DescribeCustomEventHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCustomMetricListWithOptions(request *DescribeCustomMetricListRequest, runtime *util.RuntimeOptions) (_result *DescribeCustomMetricListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCustomMetricListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCustomMetricList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCustomMetricList(request *DescribeCustomMetricListRequest) (_result *DescribeCustomMetricListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCustomMetricListResponse{}
	_body, _err := client.DescribeCustomMetricListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDynamicTagRuleListWithOptions(request *DescribeDynamicTagRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeDynamicTagRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDynamicTagRuleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDynamicTagRuleList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDynamicTagRuleList(request *DescribeDynamicTagRuleListRequest) (_result *DescribeDynamicTagRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDynamicTagRuleListResponse{}
	_body, _err := client.DescribeDynamicTagRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventRuleAttributeWithOptions(request *DescribeEventRuleAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeEventRuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventRuleAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEventRuleAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventRuleAttribute(request *DescribeEventRuleAttributeRequest) (_result *DescribeEventRuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventRuleAttributeResponse{}
	_body, _err := client.DescribeEventRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventRuleListWithOptions(request *DescribeEventRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeEventRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventRuleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEventRuleList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventRuleList(request *DescribeEventRuleListRequest) (_result *DescribeEventRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventRuleListResponse{}
	_body, _err := client.DescribeEventRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEventRuleTargetListWithOptions(request *DescribeEventRuleTargetListRequest, runtime *util.RuntimeOptions) (_result *DescribeEventRuleTargetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEventRuleTargetListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEventRuleTargetList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEventRuleTargetList(request *DescribeEventRuleTargetListRequest) (_result *DescribeEventRuleTargetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEventRuleTargetListResponse{}
	_body, _err := client.DescribeEventRuleTargetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExporterOutputListWithOptions(request *DescribeExporterOutputListRequest, runtime *util.RuntimeOptions) (_result *DescribeExporterOutputListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExporterOutputListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExporterOutputList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExporterOutputList(request *DescribeExporterOutputListRequest) (_result *DescribeExporterOutputListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExporterOutputListResponse{}
	_body, _err := client.DescribeExporterOutputListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExporterRuleListWithOptions(request *DescribeExporterRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeExporterRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExporterRuleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExporterRuleList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExporterRuleList(request *DescribeExporterRuleListRequest) (_result *DescribeExporterRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExporterRuleListResponse{}
	_body, _err := client.DescribeExporterRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupMonitoringAgentProcessWithOptions(request *DescribeGroupMonitoringAgentProcessRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupMonitoringAgentProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupMonitoringAgentProcess"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupMonitoringAgentProcess(request *DescribeGroupMonitoringAgentProcessRequest) (_result *DescribeGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DescribeGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHostAvailabilityListWithOptions(request *DescribeHostAvailabilityListRequest, runtime *util.RuntimeOptions) (_result *DescribeHostAvailabilityListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHostAvailabilityListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHostAvailabilityList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHostAvailabilityList(request *DescribeHostAvailabilityListRequest) (_result *DescribeHostAvailabilityListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHostAvailabilityListResponse{}
	_body, _err := client.DescribeHostAvailabilityListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogMonitorAttributeWithOptions(request *DescribeLogMonitorAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeLogMonitorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeLogMonitorAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogMonitorAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogMonitorAttribute(request *DescribeLogMonitorAttributeRequest) (_result *DescribeLogMonitorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogMonitorAttributeResponse{}
	_body, _err := client.DescribeLogMonitorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogMonitorListWithOptions(request *DescribeLogMonitorListRequest, runtime *util.RuntimeOptions) (_result *DescribeLogMonitorListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLogMonitorListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogMonitorList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogMonitorList(request *DescribeLogMonitorListRequest) (_result *DescribeLogMonitorListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogMonitorListResponse{}
	_body, _err := client.DescribeLogMonitorListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricDataWithOptions(request *DescribeMetricDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricData"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricData(request *DescribeMetricDataRequest) (_result *DescribeMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.DescribeMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricLastWithOptions(request *DescribeMetricLastRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricLastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricLast"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricLast(request *DescribeMetricLastRequest) (_result *DescribeMetricLastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.DescribeMetricLastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricListWithOptions(request *DescribeMetricListRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricList(request *DescribeMetricListRequest) (_result *DescribeMetricListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.DescribeMetricListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricMetaListWithOptions(request *DescribeMetricMetaListRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricMetaListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricMetaListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricMetaList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricMetaList(request *DescribeMetricMetaListRequest) (_result *DescribeMetricMetaListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricMetaListResponse{}
	_body, _err := client.DescribeMetricMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricRuleCountWithOptions(request *DescribeMetricRuleCountRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricRuleCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeMetricRuleCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricRuleCount"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricRuleCount(request *DescribeMetricRuleCountRequest) (_result *DescribeMetricRuleCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricRuleCountResponse{}
	_body, _err := client.DescribeMetricRuleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricRuleListWithOptions(request *DescribeMetricRuleListRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricRuleListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricRuleListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricRuleList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricRuleList(request *DescribeMetricRuleListRequest) (_result *DescribeMetricRuleListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricRuleListResponse{}
	_body, _err := client.DescribeMetricRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricRuleTargetsWithOptions(request *DescribeMetricRuleTargetsRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricRuleTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricRuleTargetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricRuleTargets"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricRuleTargets(request *DescribeMetricRuleTargetsRequest) (_result *DescribeMetricRuleTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricRuleTargetsResponse{}
	_body, _err := client.DescribeMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricRuleTemplateAttributeWithOptions(request *DescribeMetricRuleTemplateAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricRuleTemplateAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricRuleTemplateAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricRuleTemplateAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricRuleTemplateAttribute(request *DescribeMetricRuleTemplateAttributeRequest) (_result *DescribeMetricRuleTemplateAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricRuleTemplateAttributeResponse{}
	_body, _err := client.DescribeMetricRuleTemplateAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricRuleTemplateListWithOptions(request *DescribeMetricRuleTemplateListRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricRuleTemplateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricRuleTemplateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricRuleTemplateList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricRuleTemplateList(request *DescribeMetricRuleTemplateListRequest) (_result *DescribeMetricRuleTemplateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricRuleTemplateListResponse{}
	_body, _err := client.DescribeMetricRuleTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricTopWithOptions(request *DescribeMetricTopRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricTopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricTopResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricTop"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricTop(request *DescribeMetricTopRequest) (_result *DescribeMetricTopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricTopResponse{}
	_body, _err := client.DescribeMetricTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorGroupCategoriesWithOptions(request *DescribeMonitorGroupCategoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorGroupCategoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorGroupCategoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorGroupCategories"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorGroupCategories(request *DescribeMonitorGroupCategoriesRequest) (_result *DescribeMonitorGroupCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorGroupCategoriesResponse{}
	_body, _err := client.DescribeMonitorGroupCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorGroupDynamicRulesWithOptions(request *DescribeMonitorGroupDynamicRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorGroupDynamicRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorGroupDynamicRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorGroupDynamicRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorGroupDynamicRules(request *DescribeMonitorGroupDynamicRulesRequest) (_result *DescribeMonitorGroupDynamicRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorGroupDynamicRulesResponse{}
	_body, _err := client.DescribeMonitorGroupDynamicRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorGroupInstanceAttributeWithOptions(request *DescribeMonitorGroupInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorGroupInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorGroupInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorGroupInstanceAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorGroupInstanceAttribute(request *DescribeMonitorGroupInstanceAttributeRequest) (_result *DescribeMonitorGroupInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorGroupInstanceAttributeResponse{}
	_body, _err := client.DescribeMonitorGroupInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorGroupInstancesWithOptions(request *DescribeMonitorGroupInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorGroupInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorGroupInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorGroupInstances"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorGroupInstances(request *DescribeMonitorGroupInstancesRequest) (_result *DescribeMonitorGroupInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorGroupInstancesResponse{}
	_body, _err := client.DescribeMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorGroupNotifyPolicyListWithOptions(request *DescribeMonitorGroupNotifyPolicyListRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorGroupNotifyPolicyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorGroupNotifyPolicyListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorGroupNotifyPolicyList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorGroupNotifyPolicyList(request *DescribeMonitorGroupNotifyPolicyListRequest) (_result *DescribeMonitorGroupNotifyPolicyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorGroupNotifyPolicyListResponse{}
	_body, _err := client.DescribeMonitorGroupNotifyPolicyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorGroupsWithOptions(request *DescribeMonitorGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorGroups"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorGroups(request *DescribeMonitorGroupsRequest) (_result *DescribeMonitorGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorGroupsResponse{}
	_body, _err := client.DescribeMonitorGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentAccessKeyWithOptions(runtime *util.RuntimeOptions) (_result *DescribeMonitoringAgentAccessKeyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeMonitoringAgentAccessKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitoringAgentAccessKey"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentAccessKey() (_result *DescribeMonitoringAgentAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitoringAgentAccessKeyResponse{}
	_body, _err := client.DescribeMonitoringAgentAccessKeyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentConfigWithOptions(runtime *util.RuntimeOptions) (_result *DescribeMonitoringAgentConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeMonitoringAgentConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitoringAgentConfig"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentConfig() (_result *DescribeMonitoringAgentConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitoringAgentConfigResponse{}
	_body, _err := client.DescribeMonitoringAgentConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentHostsWithOptions(request *DescribeMonitoringAgentHostsRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitoringAgentHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitoringAgentHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitoringAgentHosts"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentHosts(request *DescribeMonitoringAgentHostsRequest) (_result *DescribeMonitoringAgentHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitoringAgentHostsResponse{}
	_body, _err := client.DescribeMonitoringAgentHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentProcessesWithOptions(request *DescribeMonitoringAgentProcessesRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitoringAgentProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitoringAgentProcessesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitoringAgentProcesses"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentProcesses(request *DescribeMonitoringAgentProcessesRequest) (_result *DescribeMonitoringAgentProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitoringAgentProcessesResponse{}
	_body, _err := client.DescribeMonitoringAgentProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentStatusesWithOptions(request *DescribeMonitoringAgentStatusesRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitoringAgentStatusesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitoringAgentStatusesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitoringAgentStatuses"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitoringAgentStatuses(request *DescribeMonitoringAgentStatusesRequest) (_result *DescribeMonitoringAgentStatusesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitoringAgentStatusesResponse{}
	_body, _err := client.DescribeMonitoringAgentStatusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitoringConfigWithOptions(runtime *util.RuntimeOptions) (_result *DescribeMonitoringConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeMonitoringConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitoringConfig"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitoringConfig() (_result *DescribeMonitoringConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitoringConfigResponse{}
	_body, _err := client.DescribeMonitoringConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMonitorResourceQuotaAttributeWithOptions(request *DescribeMonitorResourceQuotaAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorResourceQuotaAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMonitorResourceQuotaAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMonitorResourceQuotaAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMonitorResourceQuotaAttribute(request *DescribeMonitorResourceQuotaAttributeRequest) (_result *DescribeMonitorResourceQuotaAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorResourceQuotaAttributeResponse{}
	_body, _err := client.DescribeMonitorResourceQuotaAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProductResourceTagKeyListWithOptions(request *DescribeProductResourceTagKeyListRequest, runtime *util.RuntimeOptions) (_result *DescribeProductResourceTagKeyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProductResourceTagKeyListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProductResourceTagKeyList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProductResourceTagKeyList(request *DescribeProductResourceTagKeyListRequest) (_result *DescribeProductResourceTagKeyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProductResourceTagKeyListResponse{}
	_body, _err := client.DescribeProductResourceTagKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProductsOfActiveMetricRuleWithOptions(runtime *util.RuntimeOptions) (_result *DescribeProductsOfActiveMetricRuleResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeProductsOfActiveMetricRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProductsOfActiveMetricRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProductsOfActiveMetricRule() (_result *DescribeProductsOfActiveMetricRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProductsOfActiveMetricRuleResponse{}
	_body, _err := client.DescribeProductsOfActiveMetricRuleWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectMetaWithOptions(request *DescribeProjectMetaRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProjectMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProjectMeta"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProjectMeta(request *DescribeProjectMetaRequest) (_result *DescribeProjectMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectMetaResponse{}
	_body, _err := client.DescribeProjectMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSiteMonitorAttributeWithOptions(request *DescribeSiteMonitorAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeSiteMonitorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSiteMonitorAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSiteMonitorAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSiteMonitorAttribute(request *DescribeSiteMonitorAttributeRequest) (_result *DescribeSiteMonitorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSiteMonitorAttributeResponse{}
	_body, _err := client.DescribeSiteMonitorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSiteMonitorDataWithOptions(request *DescribeSiteMonitorDataRequest, runtime *util.RuntimeOptions) (_result *DescribeSiteMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSiteMonitorDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSiteMonitorData"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSiteMonitorData(request *DescribeSiteMonitorDataRequest) (_result *DescribeSiteMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSiteMonitorDataResponse{}
	_body, _err := client.DescribeSiteMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSiteMonitorListWithOptions(request *DescribeSiteMonitorListRequest, runtime *util.RuntimeOptions) (_result *DescribeSiteMonitorListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSiteMonitorListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSiteMonitorList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSiteMonitorList(request *DescribeSiteMonitorListRequest) (_result *DescribeSiteMonitorListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSiteMonitorListResponse{}
	_body, _err := client.DescribeSiteMonitorListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSiteMonitorQuotaWithOptions(runtime *util.RuntimeOptions) (_result *DescribeSiteMonitorQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeSiteMonitorQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSiteMonitorQuota"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSiteMonitorQuota() (_result *DescribeSiteMonitorQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSiteMonitorQuotaResponse{}
	_body, _err := client.DescribeSiteMonitorQuotaWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSiteMonitorStatisticsWithOptions(request *DescribeSiteMonitorStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeSiteMonitorStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSiteMonitorStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSiteMonitorStatistics"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSiteMonitorStatistics(request *DescribeSiteMonitorStatisticsRequest) (_result *DescribeSiteMonitorStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSiteMonitorStatisticsResponse{}
	_body, _err := client.DescribeSiteMonitorStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSystemEventAttributeWithOptions(request *DescribeSystemEventAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeSystemEventAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSystemEventAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSystemEventAttribute"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSystemEventAttribute(request *DescribeSystemEventAttributeRequest) (_result *DescribeSystemEventAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSystemEventAttributeResponse{}
	_body, _err := client.DescribeSystemEventAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSystemEventCountWithOptions(request *DescribeSystemEventCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSystemEventCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSystemEventCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSystemEventCount"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSystemEventCount(request *DescribeSystemEventCountRequest) (_result *DescribeSystemEventCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSystemEventCountResponse{}
	_body, _err := client.DescribeSystemEventCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSystemEventHistogramWithOptions(request *DescribeSystemEventHistogramRequest, runtime *util.RuntimeOptions) (_result *DescribeSystemEventHistogramResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSystemEventHistogramResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSystemEventHistogram"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSystemEventHistogram(request *DescribeSystemEventHistogramRequest) (_result *DescribeSystemEventHistogramResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSystemEventHistogramResponse{}
	_body, _err := client.DescribeSystemEventHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagKeyListWithOptions(request *DescribeTagKeyListRequest, runtime *util.RuntimeOptions) (_result *DescribeTagKeyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTagKeyListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTagKeyList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTagKeyList(request *DescribeTagKeyListRequest) (_result *DescribeTagKeyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagKeyListResponse{}
	_body, _err := client.DescribeTagKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagValueListWithOptions(request *DescribeTagValueListRequest, runtime *util.RuntimeOptions) (_result *DescribeTagValueListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTagValueListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTagValueList"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTagValueList(request *DescribeTagValueListRequest) (_result *DescribeTagValueListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagValueListResponse{}
	_body, _err := client.DescribeTagValueListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUnhealthyHostAvailabilityWithOptions(request *DescribeUnhealthyHostAvailabilityRequest, runtime *util.RuntimeOptions) (_result *DescribeUnhealthyHostAvailabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUnhealthyHostAvailabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUnhealthyHostAvailability"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUnhealthyHostAvailability(request *DescribeUnhealthyHostAvailabilityRequest) (_result *DescribeUnhealthyHostAvailabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUnhealthyHostAvailabilityResponse{}
	_body, _err := client.DescribeUnhealthyHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableActiveMetricRuleWithOptions(request *DisableActiveMetricRuleRequest, runtime *util.RuntimeOptions) (_result *DisableActiveMetricRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableActiveMetricRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableActiveMetricRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableActiveMetricRule(request *DisableActiveMetricRuleRequest) (_result *DisableActiveMetricRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableActiveMetricRuleResponse{}
	_body, _err := client.DisableActiveMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableEventRulesWithOptions(request *DisableEventRulesRequest, runtime *util.RuntimeOptions) (_result *DisableEventRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableEventRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableEventRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableEventRules(request *DisableEventRulesRequest) (_result *DisableEventRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableEventRulesResponse{}
	_body, _err := client.DisableEventRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableHostAvailabilityWithOptions(request *DisableHostAvailabilityRequest, runtime *util.RuntimeOptions) (_result *DisableHostAvailabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableHostAvailabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableHostAvailability"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableHostAvailability(request *DisableHostAvailabilityRequest) (_result *DisableHostAvailabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableHostAvailabilityResponse{}
	_body, _err := client.DisableHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableMetricRulesWithOptions(request *DisableMetricRulesRequest, runtime *util.RuntimeOptions) (_result *DisableMetricRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableMetricRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableMetricRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableMetricRules(request *DisableMetricRulesRequest) (_result *DisableMetricRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableMetricRulesResponse{}
	_body, _err := client.DisableMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSiteMonitorsWithOptions(request *DisableSiteMonitorsRequest, runtime *util.RuntimeOptions) (_result *DisableSiteMonitorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableSiteMonitorsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableSiteMonitors"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSiteMonitors(request *DisableSiteMonitorsRequest) (_result *DisableSiteMonitorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableSiteMonitorsResponse{}
	_body, _err := client.DisableSiteMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableActiveMetricRuleWithOptions(request *EnableActiveMetricRuleRequest, runtime *util.RuntimeOptions) (_result *EnableActiveMetricRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableActiveMetricRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableActiveMetricRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableActiveMetricRule(request *EnableActiveMetricRuleRequest) (_result *EnableActiveMetricRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableActiveMetricRuleResponse{}
	_body, _err := client.EnableActiveMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableEventRulesWithOptions(request *EnableEventRulesRequest, runtime *util.RuntimeOptions) (_result *EnableEventRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableEventRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableEventRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableEventRules(request *EnableEventRulesRequest) (_result *EnableEventRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableEventRulesResponse{}
	_body, _err := client.EnableEventRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableHostAvailabilityWithOptions(request *EnableHostAvailabilityRequest, runtime *util.RuntimeOptions) (_result *EnableHostAvailabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableHostAvailabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableHostAvailability"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableHostAvailability(request *EnableHostAvailabilityRequest) (_result *EnableHostAvailabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableHostAvailabilityResponse{}
	_body, _err := client.EnableHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableMetricRulesWithOptions(request *EnableMetricRulesRequest, runtime *util.RuntimeOptions) (_result *EnableMetricRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableMetricRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableMetricRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableMetricRules(request *EnableMetricRulesRequest) (_result *EnableMetricRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableMetricRulesResponse{}
	_body, _err := client.EnableMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSiteMonitorsWithOptions(request *EnableSiteMonitorsRequest, runtime *util.RuntimeOptions) (_result *EnableSiteMonitorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableSiteMonitorsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableSiteMonitors"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSiteMonitors(request *EnableSiteMonitorsRequest) (_result *EnableSiteMonitorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableSiteMonitorsResponse{}
	_body, _err := client.EnableSiteMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallMonitoringAgentWithOptions(request *InstallMonitoringAgentRequest, runtime *util.RuntimeOptions) (_result *InstallMonitoringAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InstallMonitoringAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InstallMonitoringAgent"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallMonitoringAgent(request *InstallMonitoringAgentRequest) (_result *InstallMonitoringAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallMonitoringAgentResponse{}
	_body, _err := client.InstallMonitoringAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGroupMonitoringAgentProcessWithOptions(request *ModifyGroupMonitoringAgentProcessRequest, runtime *util.RuntimeOptions) (_result *ModifyGroupMonitoringAgentProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyGroupMonitoringAgentProcess"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGroupMonitoringAgentProcess(request *ModifyGroupMonitoringAgentProcessRequest) (_result *ModifyGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGroupMonitoringAgentProcessResponse{}
	_body, _err := client.ModifyGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHostAvailabilityWithOptions(request *ModifyHostAvailabilityRequest, runtime *util.RuntimeOptions) (_result *ModifyHostAvailabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHostAvailabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHostAvailability"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHostAvailability(request *ModifyHostAvailabilityRequest) (_result *ModifyHostAvailabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostAvailabilityResponse{}
	_body, _err := client.ModifyHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHostInfoWithOptions(request *ModifyHostInfoRequest, runtime *util.RuntimeOptions) (_result *ModifyHostInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHostInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHostInfo"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHostInfo(request *ModifyHostInfoRequest) (_result *ModifyHostInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostInfoResponse{}
	_body, _err := client.ModifyHostInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMetricRuleTemplateWithOptions(request *ModifyMetricRuleTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyMetricRuleTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMetricRuleTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMetricRuleTemplate"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMetricRuleTemplate(request *ModifyMetricRuleTemplateRequest) (_result *ModifyMetricRuleTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMetricRuleTemplateResponse{}
	_body, _err := client.ModifyMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMonitorGroupWithOptions(request *ModifyMonitorGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyMonitorGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMonitorGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMonitorGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMonitorGroup(request *ModifyMonitorGroupRequest) (_result *ModifyMonitorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMonitorGroupResponse{}
	_body, _err := client.ModifyMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMonitorGroupInstancesWithOptions(request *ModifyMonitorGroupInstancesRequest, runtime *util.RuntimeOptions) (_result *ModifyMonitorGroupInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMonitorGroupInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMonitorGroupInstances"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMonitorGroupInstances(request *ModifyMonitorGroupInstancesRequest) (_result *ModifyMonitorGroupInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMonitorGroupInstancesResponse{}
	_body, _err := client.ModifyMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySiteMonitorWithOptions(request *ModifySiteMonitorRequest, runtime *util.RuntimeOptions) (_result *ModifySiteMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySiteMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySiteMonitor"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySiteMonitor(request *ModifySiteMonitorRequest) (_result *ModifySiteMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySiteMonitorResponse{}
	_body, _err := client.ModifySiteMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenCmsServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenCmsServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &OpenCmsServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OpenCmsService"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenCmsService() (_result *OpenCmsServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCmsServiceResponse{}
	_body, _err := client.OpenCmsServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutContactWithOptions(request *PutContactRequest, runtime *util.RuntimeOptions) (_result *PutContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutContactResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutContact"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutContact(request *PutContactRequest) (_result *PutContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutContactResponse{}
	_body, _err := client.PutContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutContactGroupWithOptions(request *PutContactGroupRequest, runtime *util.RuntimeOptions) (_result *PutContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutContactGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutContactGroup"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutContactGroup(request *PutContactGroupRequest) (_result *PutContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutContactGroupResponse{}
	_body, _err := client.PutContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutCustomEventWithOptions(request *PutCustomEventRequest, runtime *util.RuntimeOptions) (_result *PutCustomEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutCustomEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutCustomEvent"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutCustomEvent(request *PutCustomEventRequest) (_result *PutCustomEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutCustomEventResponse{}
	_body, _err := client.PutCustomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutCustomEventRuleWithOptions(request *PutCustomEventRuleRequest, runtime *util.RuntimeOptions) (_result *PutCustomEventRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutCustomEventRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutCustomEventRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutCustomEventRule(request *PutCustomEventRuleRequest) (_result *PutCustomEventRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutCustomEventRuleResponse{}
	_body, _err := client.PutCustomEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutCustomMetricWithOptions(request *PutCustomMetricRequest, runtime *util.RuntimeOptions) (_result *PutCustomMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutCustomMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutCustomMetric"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutCustomMetric(request *PutCustomMetricRequest) (_result *PutCustomMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutCustomMetricResponse{}
	_body, _err := client.PutCustomMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutCustomMetricRuleWithOptions(request *PutCustomMetricRuleRequest, runtime *util.RuntimeOptions) (_result *PutCustomMetricRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutCustomMetricRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutCustomMetricRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutCustomMetricRule(request *PutCustomMetricRuleRequest) (_result *PutCustomMetricRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutCustomMetricRuleResponse{}
	_body, _err := client.PutCustomMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutEventRuleWithOptions(request *PutEventRuleRequest, runtime *util.RuntimeOptions) (_result *PutEventRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutEventRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutEventRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutEventRule(request *PutEventRuleRequest) (_result *PutEventRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEventRuleResponse{}
	_body, _err := client.PutEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutEventRuleTargetsWithOptions(request *PutEventRuleTargetsRequest, runtime *util.RuntimeOptions) (_result *PutEventRuleTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutEventRuleTargetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutEventRuleTargets"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutEventRuleTargets(request *PutEventRuleTargetsRequest) (_result *PutEventRuleTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEventRuleTargetsResponse{}
	_body, _err := client.PutEventRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutExporterOutputWithOptions(request *PutExporterOutputRequest, runtime *util.RuntimeOptions) (_result *PutExporterOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutExporterOutputResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutExporterOutput"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutExporterOutput(request *PutExporterOutputRequest) (_result *PutExporterOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutExporterOutputResponse{}
	_body, _err := client.PutExporterOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutExporterRuleWithOptions(request *PutExporterRuleRequest, runtime *util.RuntimeOptions) (_result *PutExporterRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutExporterRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutExporterRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutExporterRule(request *PutExporterRuleRequest) (_result *PutExporterRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutExporterRuleResponse{}
	_body, _err := client.PutExporterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutGroupMetricRuleWithOptions(request *PutGroupMetricRuleRequest, runtime *util.RuntimeOptions) (_result *PutGroupMetricRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutGroupMetricRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutGroupMetricRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutGroupMetricRule(request *PutGroupMetricRuleRequest) (_result *PutGroupMetricRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutGroupMetricRuleResponse{}
	_body, _err := client.PutGroupMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutLogMonitorWithOptions(request *PutLogMonitorRequest, runtime *util.RuntimeOptions) (_result *PutLogMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutLogMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutLogMonitor"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutLogMonitor(request *PutLogMonitorRequest) (_result *PutLogMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutLogMonitorResponse{}
	_body, _err := client.PutLogMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMetricRuleTargetsWithOptions(request *PutMetricRuleTargetsRequest, runtime *util.RuntimeOptions) (_result *PutMetricRuleTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutMetricRuleTargetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutMetricRuleTargets"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMetricRuleTargets(request *PutMetricRuleTargetsRequest) (_result *PutMetricRuleTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMetricRuleTargetsResponse{}
	_body, _err := client.PutMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMonitorGroupDynamicRuleWithOptions(request *PutMonitorGroupDynamicRuleRequest, runtime *util.RuntimeOptions) (_result *PutMonitorGroupDynamicRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutMonitorGroupDynamicRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutMonitorGroupDynamicRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMonitorGroupDynamicRule(request *PutMonitorGroupDynamicRuleRequest) (_result *PutMonitorGroupDynamicRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMonitorGroupDynamicRuleResponse{}
	_body, _err := client.PutMonitorGroupDynamicRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutMonitoringConfigWithOptions(request *PutMonitoringConfigRequest, runtime *util.RuntimeOptions) (_result *PutMonitoringConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutMonitoringConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutMonitoringConfig"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutMonitoringConfig(request *PutMonitoringConfigRequest) (_result *PutMonitoringConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutMonitoringConfigResponse{}
	_body, _err := client.PutMonitoringConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutResourceMetricRuleWithOptions(request *PutResourceMetricRuleRequest, runtime *util.RuntimeOptions) (_result *PutResourceMetricRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutResourceMetricRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutResourceMetricRule"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutResourceMetricRule(request *PutResourceMetricRuleRequest) (_result *PutResourceMetricRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutResourceMetricRuleResponse{}
	_body, _err := client.PutResourceMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutResourceMetricRulesWithOptions(request *PutResourceMetricRulesRequest, runtime *util.RuntimeOptions) (_result *PutResourceMetricRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutResourceMetricRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutResourceMetricRules"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutResourceMetricRules(request *PutResourceMetricRulesRequest) (_result *PutResourceMetricRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutResourceMetricRulesResponse{}
	_body, _err := client.PutResourceMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTagsWithOptions(request *RemoveTagsRequest, runtime *util.RuntimeOptions) (_result *RemoveTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveTags"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTags(request *RemoveTagsRequest) (_result *RemoveTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTagsResponse{}
	_body, _err := client.RemoveTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendDryRunSystemEventWithOptions(request *SendDryRunSystemEventRequest, runtime *util.RuntimeOptions) (_result *SendDryRunSystemEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendDryRunSystemEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendDryRunSystemEvent"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendDryRunSystemEvent(request *SendDryRunSystemEventRequest) (_result *SendDryRunSystemEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendDryRunSystemEventResponse{}
	_body, _err := client.SendDryRunSystemEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallMonitoringAgentWithOptions(request *UninstallMonitoringAgentRequest, runtime *util.RuntimeOptions) (_result *UninstallMonitoringAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UninstallMonitoringAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UninstallMonitoringAgent"), tea.String("2019-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallMonitoringAgent(request *UninstallMonitoringAgentRequest) (_result *UninstallMonitoringAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallMonitoringAgentResponse{}
	_body, _err := client.UninstallMonitoringAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
