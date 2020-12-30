// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddVpcHoneyPotRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcSwitchId *string `json:"VpcSwitchId,omitempty" xml:"VpcSwitchId,omitempty"`
}

func (s AddVpcHoneyPotRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVpcHoneyPotRequest) GoString() string {
	return s.String()
}

func (s *AddVpcHoneyPotRequest) SetSourceIp(v string) *AddVpcHoneyPotRequest {
	s.SourceIp = &v
	return s
}

func (s *AddVpcHoneyPotRequest) SetVpcId(v string) *AddVpcHoneyPotRequest {
	s.VpcId = &v
	return s
}

func (s *AddVpcHoneyPotRequest) SetVpcSwitchId(v string) *AddVpcHoneyPotRequest {
	s.VpcSwitchId = &v
	return s
}

type AddVpcHoneyPotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVpcHoneyPotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVpcHoneyPotResponseBody) GoString() string {
	return s.String()
}

func (s *AddVpcHoneyPotResponseBody) SetRequestId(v string) *AddVpcHoneyPotResponseBody {
	s.RequestId = &v
	return s
}

type AddVpcHoneyPotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVpcHoneyPotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVpcHoneyPotResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVpcHoneyPotResponse) GoString() string {
	return s.String()
}

func (s *AddVpcHoneyPotResponse) SetHeaders(v map[string]*string) *AddVpcHoneyPotResponse {
	s.Headers = v
	return s
}

func (s *AddVpcHoneyPotResponse) SetBody(v *AddVpcHoneyPotResponseBody) *AddVpcHoneyPotResponse {
	s.Body = v
	return s
}

type CreateAntiBruteForceRuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Span            *int32    `json:"Span,omitempty" xml:"Span,omitempty"`
	FailCount       *int32    `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	ForbiddenTime   *int32    `json:"ForbiddenTime,omitempty" xml:"ForbiddenTime,omitempty"`
	EnableSmartRule *bool     `json:"EnableSmartRule,omitempty" xml:"EnableSmartRule,omitempty"`
	DefaultRule     *bool     `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	UuidList        []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s CreateAntiBruteForceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAntiBruteForceRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleRequest) SetSourceIp(v string) *CreateAntiBruteForceRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetResourceOwnerId(v int64) *CreateAntiBruteForceRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetName(v string) *CreateAntiBruteForceRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetSpan(v int32) *CreateAntiBruteForceRuleRequest {
	s.Span = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetFailCount(v int32) *CreateAntiBruteForceRuleRequest {
	s.FailCount = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetForbiddenTime(v int32) *CreateAntiBruteForceRuleRequest {
	s.ForbiddenTime = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetEnableSmartRule(v bool) *CreateAntiBruteForceRuleRequest {
	s.EnableSmartRule = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetDefaultRule(v bool) *CreateAntiBruteForceRuleRequest {
	s.DefaultRule = &v
	return s
}

func (s *CreateAntiBruteForceRuleRequest) SetUuidList(v []*string) *CreateAntiBruteForceRuleRequest {
	s.UuidList = v
	return s
}

type CreateAntiBruteForceRuleResponseBody struct {
	RequestId                *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateAntiBruteForceRule *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule `json:"CreateAntiBruteForceRule,omitempty" xml:"CreateAntiBruteForceRule,omitempty" type:"Struct"`
}

func (s CreateAntiBruteForceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAntiBruteForceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleResponseBody) SetRequestId(v string) *CreateAntiBruteForceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAntiBruteForceRuleResponseBody) SetCreateAntiBruteForceRule(v *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) *CreateAntiBruteForceRuleResponseBody {
	s.CreateAntiBruteForceRule = v
	return s
}

type CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule struct {
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) String() string {
	return tea.Prettify(s)
}

func (s CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule) SetRuleId(v int64) *CreateAntiBruteForceRuleResponseBodyCreateAntiBruteForceRule {
	s.RuleId = &v
	return s
}

type CreateAntiBruteForceRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAntiBruteForceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAntiBruteForceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAntiBruteForceRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleResponse) SetHeaders(v map[string]*string) *CreateAntiBruteForceRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAntiBruteForceRuleResponse) SetBody(v *CreateAntiBruteForceRuleResponseBody) *CreateAntiBruteForceRuleResponse {
	s.Body = v
	return s
}

type CreateOrUpdateAssetGroupRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuids     *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s CreateOrUpdateAssetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAssetGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAssetGroupRequest) SetSourceIp(v string) *CreateOrUpdateAssetGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateOrUpdateAssetGroupRequest) SetUuids(v string) *CreateOrUpdateAssetGroupRequest {
	s.Uuids = &v
	return s
}

func (s *CreateOrUpdateAssetGroupRequest) SetGroupName(v string) *CreateOrUpdateAssetGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateOrUpdateAssetGroupRequest) SetGroupId(v int64) *CreateOrUpdateAssetGroupRequest {
	s.GroupId = &v
	return s
}

type CreateOrUpdateAssetGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateAssetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAssetGroupResponseBody) SetRequestId(v string) *CreateOrUpdateAssetGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrUpdateAssetGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrUpdateAssetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateAssetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAssetGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAssetGroupResponse) SetHeaders(v map[string]*string) *CreateOrUpdateAssetGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateAssetGroupResponse) SetBody(v *CreateOrUpdateAssetGroupResponseBody) *CreateOrUpdateAssetGroupResponse {
	s.Body = v
	return s
}

type CreateSasOrderRequest struct {
	Period              *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit          *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoRenewPeriod     *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	AutoPay             *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon       *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	Spec                *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	InstanceCount       *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	SasSlsStorage       *string `json:"SasSlsStorage,omitempty" xml:"SasSlsStorage,omitempty"`
	SasAntiRansomware   *string `json:"SasAntiRansomware,omitempty" xml:"SasAntiRansomware,omitempty"`
	SasWebguardBoolean  *string `json:"SasWebguardBoolean,omitempty" xml:"SasWebguardBoolean,omitempty"`
	SasSc               *string `json:"SasSc,omitempty" xml:"SasSc,omitempty"`
	SasProductService   *string `json:"SasProductService,omitempty" xml:"SasProductService,omitempty"`
	SasWebguardOrderNum *string `json:"SasWebguardOrderNum,omitempty" xml:"SasWebguardOrderNum,omitempty"`
}

func (s CreateSasOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSasOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateSasOrderRequest) SetPeriod(v int32) *CreateSasOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateSasOrderRequest) SetPeriodUnit(v string) *CreateSasOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateSasOrderRequest) SetAutoRenewPeriod(v int32) *CreateSasOrderRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateSasOrderRequest) SetAutoPay(v bool) *CreateSasOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateSasOrderRequest) SetAutoUseCoupon(v bool) *CreateSasOrderRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreateSasOrderRequest) SetSpec(v string) *CreateSasOrderRequest {
	s.Spec = &v
	return s
}

func (s *CreateSasOrderRequest) SetInstanceCount(v string) *CreateSasOrderRequest {
	s.InstanceCount = &v
	return s
}

func (s *CreateSasOrderRequest) SetSasSlsStorage(v string) *CreateSasOrderRequest {
	s.SasSlsStorage = &v
	return s
}

func (s *CreateSasOrderRequest) SetSasAntiRansomware(v string) *CreateSasOrderRequest {
	s.SasAntiRansomware = &v
	return s
}

func (s *CreateSasOrderRequest) SetSasWebguardBoolean(v string) *CreateSasOrderRequest {
	s.SasWebguardBoolean = &v
	return s
}

func (s *CreateSasOrderRequest) SetSasSc(v string) *CreateSasOrderRequest {
	s.SasSc = &v
	return s
}

func (s *CreateSasOrderRequest) SetSasProductService(v string) *CreateSasOrderRequest {
	s.SasProductService = &v
	return s
}

func (s *CreateSasOrderRequest) SetSasWebguardOrderNum(v string) *CreateSasOrderRequest {
	s.SasWebguardOrderNum = &v
	return s
}

type CreateSasOrderResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateSasOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSasOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSasOrderResponseBody) SetRequestId(v string) *CreateSasOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSasOrderResponseBody) SetOrderId(v string) *CreateSasOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateSasOrderResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSasOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSasOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSasOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateSasOrderResponse) SetHeaders(v map[string]*string) *CreateSasOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateSasOrderResponse) SetBody(v *CreateSasOrderResponseBody) *CreateSasOrderResponse {
	s.Body = v
	return s
}

type CreateSimilarSecurityEventsQueryTaskRequest struct {
	SourceIp                 *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityEventId          *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	SimilarEventScenarioCode *string `json:"SimilarEventScenarioCode,omitempty" xml:"SimilarEventScenarioCode,omitempty"`
}

func (s CreateSimilarSecurityEventsQueryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSourceIp(v string) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetResourceOwnerId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSecurityEventId(v int64) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SecurityEventId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskRequest) SetSimilarEventScenarioCode(v string) *CreateSimilarSecurityEventsQueryTaskRequest {
	s.SimilarEventScenarioCode = &v
	return s
}

type CreateSimilarSecurityEventsQueryTaskResponseBody struct {
	RequestId                                    *string                                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateSimilarSecurityEventsQueryTaskResponse *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse `json:"CreateSimilarSecurityEventsQueryTaskResponse,omitempty" xml:"CreateSimilarSecurityEventsQueryTaskResponse,omitempty" type:"Struct"`
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBody) SetRequestId(v string) *CreateSimilarSecurityEventsQueryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBody) SetCreateSimilarSecurityEventsQueryTaskResponse(v *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) *CreateSimilarSecurityEventsQueryTaskResponseBody {
	s.CreateSimilarSecurityEventsQueryTaskResponse = v
	return s
}

type CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) SetStatus(v string) *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse {
	s.Status = &v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse) SetTaskId(v int64) *CreateSimilarSecurityEventsQueryTaskResponseBodyCreateSimilarSecurityEventsQueryTaskResponse {
	s.TaskId = &v
	return s
}

type CreateSimilarSecurityEventsQueryTaskResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSimilarSecurityEventsQueryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSimilarSecurityEventsQueryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarSecurityEventsQueryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) SetHeaders(v map[string]*string) *CreateSimilarSecurityEventsQueryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSimilarSecurityEventsQueryTaskResponse) SetBody(v *CreateSimilarSecurityEventsQueryTaskResponseBody) *CreateSimilarSecurityEventsQueryTaskResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	GroupId  *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetSourceIp(v string) *DeleteGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteGroupRequest) SetGroupId(v int64) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteLoginBaseConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Config   *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Target   *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s DeleteLoginBaseConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginBaseConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoginBaseConfigRequest) SetSourceIp(v string) *DeleteLoginBaseConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteLoginBaseConfigRequest) SetType(v string) *DeleteLoginBaseConfigRequest {
	s.Type = &v
	return s
}

func (s *DeleteLoginBaseConfigRequest) SetConfig(v string) *DeleteLoginBaseConfigRequest {
	s.Config = &v
	return s
}

func (s *DeleteLoginBaseConfigRequest) SetTarget(v string) *DeleteLoginBaseConfigRequest {
	s.Target = &v
	return s
}

type DeleteLoginBaseConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoginBaseConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginBaseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoginBaseConfigResponseBody) SetRequestId(v string) *DeleteLoginBaseConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoginBaseConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLoginBaseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLoginBaseConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginBaseConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoginBaseConfigResponse) SetHeaders(v map[string]*string) *DeleteLoginBaseConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoginBaseConfigResponse) SetBody(v *DeleteLoginBaseConfigResponseBody) *DeleteLoginBaseConfigResponse {
	s.Body = v
	return s
}

type DeleteTagWithUuidRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TagName  *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s DeleteTagWithUuidRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagWithUuidRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagWithUuidRequest) SetSourceIp(v string) *DeleteTagWithUuidRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteTagWithUuidRequest) SetTagName(v string) *DeleteTagWithUuidRequest {
	s.TagName = &v
	return s
}

func (s *DeleteTagWithUuidRequest) SetUuidList(v string) *DeleteTagWithUuidRequest {
	s.UuidList = &v
	return s
}

type DeleteTagWithUuidResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagWithUuidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagWithUuidResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagWithUuidResponseBody) SetRequestId(v string) *DeleteTagWithUuidResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagWithUuidResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTagWithUuidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTagWithUuidResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagWithUuidResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagWithUuidResponse) SetHeaders(v map[string]*string) *DeleteTagWithUuidResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagWithUuidResponse) SetBody(v *DeleteTagWithUuidResponseBody) *DeleteTagWithUuidResponse {
	s.Body = v
	return s
}

type DeleteVpcHoneyPotRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VpcId    *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeleteVpcHoneyPotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcHoneyPotRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcHoneyPotRequest) SetSourceIp(v string) *DeleteVpcHoneyPotRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteVpcHoneyPotRequest) SetVpcId(v string) *DeleteVpcHoneyPotRequest {
	s.VpcId = &v
	return s
}

type DeleteVpcHoneyPotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcHoneyPotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcHoneyPotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcHoneyPotResponseBody) SetRequestId(v string) *DeleteVpcHoneyPotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcHoneyPotResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVpcHoneyPotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcHoneyPotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcHoneyPotResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcHoneyPotResponse) SetHeaders(v map[string]*string) *DeleteVpcHoneyPotResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcHoneyPotResponse) SetBody(v *DeleteVpcHoneyPotResponseBody) *DeleteVpcHoneyPotResponse {
	s.Body = v
	return s
}

type DescribeAccesskeyLeakListRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
	StartTs     *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeAccesskeyLeakListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccesskeyLeakListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListRequest) SetSourceIp(v string) *DescribeAccesskeyLeakListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetStatus(v string) *DescribeAccesskeyLeakListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetQuery(v string) *DescribeAccesskeyLeakListRequest {
	s.Query = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetStartTs(v int64) *DescribeAccesskeyLeakListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetPageSize(v int32) *DescribeAccesskeyLeakListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetCurrentPage(v int32) *DescribeAccesskeyLeakListRequest {
	s.CurrentPage = &v
	return s
}

type DescribeAccesskeyLeakListResponseBody struct {
	AkLeakCount       *int32                                                    `json:"AkLeakCount,omitempty" xml:"AkLeakCount,omitempty"`
	AccessKeyLeakList []*DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList `json:"AccessKeyLeakList,omitempty" xml:"AccessKeyLeakList,omitempty" type:"Repeated"`
	TotalCount        *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize          *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int32                                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	GmtLast           *int64                                                    `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
}

func (s DescribeAccesskeyLeakListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccesskeyLeakListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListResponseBody) SetAkLeakCount(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.AkLeakCount = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetAccessKeyLeakList(v []*DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) *DescribeAccesskeyLeakListResponseBody {
	s.AccessKeyLeakList = v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetTotalCount(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetRequestId(v string) *DescribeAccesskeyLeakListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetPageSize(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetCurrentPage(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetGmtLast(v int64) *DescribeAccesskeyLeakListResponseBody {
	s.GmtLast = &v
	return s
}

type DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList struct {
	DealTime    *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserType    *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	AccesskeyId *string `json:"AccesskeyId,omitempty" xml:"AccesskeyId,omitempty"`
	AliUserName *string `json:"AliUserName,omitempty" xml:"AliUserName,omitempty"`
	DealType    *string `json:"DealType,omitempty" xml:"DealType,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Asset       *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetDealTime(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.DealTime = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetStatus(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Status = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetType(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Type = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetUserType(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.UserType = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetAccesskeyId(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.AccesskeyId = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetAliUserName(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.AliUserName = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetDealType(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.DealType = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetUrl(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Url = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetGmtModified(v int64) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.GmtModified = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetAsset(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Asset = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetId(v int64) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Id = &v
	return s
}

type DescribeAccesskeyLeakListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccesskeyLeakListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccesskeyLeakListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccesskeyLeakListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListResponse) SetHeaders(v map[string]*string) *DescribeAccesskeyLeakListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccesskeyLeakListResponse) SetBody(v *DescribeAccesskeyLeakListResponseBody) *DescribeAccesskeyLeakListResponse {
	s.Body = v
	return s
}

type DescribeAffectedMaliciousFileImagesRequest struct {
	SourceIp       *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	MaliciousMd5   *string   `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	CurrentPage    *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize       *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Lang           *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RepoRegionId   *string   `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	RepoInstanceId *string   `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	RepoId         *string   `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepoName       *string   `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespace  *string   `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	ImageTag       *string   `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ImageDigest    *string   `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	ImageLayer     *string   `json:"ImageLayer,omitempty" xml:"ImageLayer,omitempty"`
	Uuids          []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s DescribeAffectedMaliciousFileImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetSourceIp(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetMaliciousMd5(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.MaliciousMd5 = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetCurrentPage(v int32) *DescribeAffectedMaliciousFileImagesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetPageSize(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetLang(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoRegionId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoInstanceId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoId(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoName(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetRepoNamespace(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetImageTag(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetImageDigest(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetImageLayer(v string) *DescribeAffectedMaliciousFileImagesRequest {
	s.ImageLayer = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesRequest) SetUuids(v []*string) *DescribeAffectedMaliciousFileImagesRequest {
	s.Uuids = v
	return s
}

type DescribeAffectedMaliciousFileImagesResponseBody struct {
	AffectedMaliciousFileImagesResponse []*DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse `json:"AffectedMaliciousFileImagesResponse,omitempty" xml:"AffectedMaliciousFileImagesResponse,omitempty" type:"Repeated"`
	PageInfo                            *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo                              `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId                           *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) SetAffectedMaliciousFileImagesResponse(v []*DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) *DescribeAffectedMaliciousFileImagesResponseBody {
	s.AffectedMaliciousFileImagesResponse = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) SetPageInfo(v *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) *DescribeAffectedMaliciousFileImagesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBody) SetRequestId(v string) *DescribeAffectedMaliciousFileImagesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse struct {
	Status                *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Digest                *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	LatestVerifyTimestamp *int64  `json:"LatestVerifyTimestamp,omitempty" xml:"LatestVerifyTimestamp,omitempty"`
	RepoInstanceId        *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	Namespace             *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RepoRegionId          *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	Tag                   *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	ImageUuid             *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	MaliciousMd5          *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	FirstScanTimestamp    *int64  `json:"FirstScanTimestamp,omitempty" xml:"FirstScanTimestamp,omitempty"`
	FilePath              *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	RepoId                *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	Layer                 *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	LatestScanTimestamp   *int64  `json:"LatestScanTimestamp,omitempty" xml:"LatestScanTimestamp,omitempty"`
	RepoName              *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	Level                 *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetStatus(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Status = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetDigest(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Digest = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLatestVerifyTimestamp(v int64) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.LatestVerifyTimestamp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoInstanceId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetNamespace(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Namespace = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoRegionId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetTag(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Tag = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetImageUuid(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.ImageUuid = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetMaliciousMd5(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.MaliciousMd5 = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetFirstScanTimestamp(v int64) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.FirstScanTimestamp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetFilePath(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.FilePath = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoId(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoId = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLayer(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Layer = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLatestScanTimestamp(v int64) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.LatestScanTimestamp = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetRepoName(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.RepoName = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse) SetLevel(v string) *DescribeAffectedMaliciousFileImagesResponseBodyAffectedMaliciousFileImagesResponse {
	s.Level = &v
	return s
}

type DescribeAffectedMaliciousFileImagesResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetPageSize(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo) SetCount(v int32) *DescribeAffectedMaliciousFileImagesResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeAffectedMaliciousFileImagesResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAffectedMaliciousFileImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAffectedMaliciousFileImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAffectedMaliciousFileImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAffectedMaliciousFileImagesResponse) SetHeaders(v map[string]*string) *DescribeAffectedMaliciousFileImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAffectedMaliciousFileImagesResponse) SetBody(v *DescribeAffectedMaliciousFileImagesResponseBody) *DescribeAffectedMaliciousFileImagesResponse {
	s.Body = v
	return s
}

type DescribeAlarmEventDetailRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	From            *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s DescribeAlarmEventDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailRequest) SetSourceIp(v string) *DescribeAlarmEventDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAlarmEventDetailRequest) SetLang(v string) *DescribeAlarmEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAlarmEventDetailRequest) SetAlarmUniqueInfo(v string) *DescribeAlarmEventDetailRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventDetailRequest) SetFrom(v string) *DescribeAlarmEventDetailRequest {
	s.From = &v
	return s
}

type DescribeAlarmEventDetailResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeAlarmEventDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeAlarmEventDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBody) SetRequestId(v string) *DescribeAlarmEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBody) SetData(v *DescribeAlarmEventDetailResponseBodyData) *DescribeAlarmEventDetailResponseBody {
	s.Data = v
	return s
}

type DescribeAlarmEventDetailResponseBodyData struct {
	Type                *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	InternetIp          *string                                                 `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	K8sClusterName      *string                                                 `json:"K8sClusterName,omitempty" xml:"K8sClusterName,omitempty"`
	ContainerImageId    *string                                                 `json:"ContainerImageId,omitempty" xml:"ContainerImageId,omitempty"`
	AlarmEventDesc      *string                                                 `json:"AlarmEventDesc,omitempty" xml:"AlarmEventDesc,omitempty"`
	AlarmUniqueInfo     *string                                                 `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	CauseDetails        []*DescribeAlarmEventDetailResponseBodyDataCauseDetails `json:"CauseDetails,omitempty" xml:"CauseDetails,omitempty" type:"Repeated"`
	CanCancelFault      *bool                                                   `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	AppName             *string                                                 `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CanBeDealOnLine     *bool                                                   `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	ContainerImageName  *string                                                 `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	K8sClusterId        *string                                                 `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ContainHwMode       *bool                                                   `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	K8sNodeId           *string                                                 `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	InstanceName        *string                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Solution            *string                                                 `json:"Solution,omitempty" xml:"Solution,omitempty"`
	DataSource          *string                                                 `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	IntranetIp          *string                                                 `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	EndTime             *int64                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AlarmEventAliasName *string                                                 `json:"AlarmEventAliasName,omitempty" xml:"AlarmEventAliasName,omitempty"`
	StartTime           *int64                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Uuid                *string                                                 `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	K8sPodName          *string                                                 `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	ContainerId         *string                                                 `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	K8sNamespace        *string                                                 `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	K8sNodeName         *string                                                 `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	Level               *string                                                 `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeAlarmEventDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetType(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetInternetIp(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sClusterName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sClusterName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainerImageId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainerImageId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAlarmEventDesc(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AlarmEventDesc = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAlarmUniqueInfo(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetCauseDetails(v []*DescribeAlarmEventDetailResponseBodyDataCauseDetails) *DescribeAlarmEventDetailResponseBodyData {
	s.CauseDetails = v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetCanCancelFault(v bool) *DescribeAlarmEventDetailResponseBodyData {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAppName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetCanBeDealOnLine(v bool) *DescribeAlarmEventDetailResponseBodyData {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainerImageName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainerImageName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sClusterId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainHwMode(v bool) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainHwMode = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sNodeId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetInstanceName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetSolution(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Solution = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetDataSource(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.DataSource = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetIntranetIp(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetEndTime(v int64) *DescribeAlarmEventDetailResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetAlarmEventAliasName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.AlarmEventAliasName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetStartTime(v int64) *DescribeAlarmEventDetailResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetUuid(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sPodName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sPodName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetContainerId(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.ContainerId = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sNamespace(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetK8sNodeName(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyData) SetLevel(v string) *DescribeAlarmEventDetailResponseBodyData {
	s.Level = &v
	return s
}

type DescribeAlarmEventDetailResponseBodyDataCauseDetails struct {
	Key   *string                                                      `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetails) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetails) SetKey(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetails {
	s.Key = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetails) SetValue(v []*DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) *DescribeAlarmEventDetailResponseBodyDataCauseDetails {
	s.Value = v
	return s
}

type DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) SetType(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue {
	s.Type = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) SetValue(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue {
	s.Value = &v
	return s
}

func (s *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue) SetName(v string) *DescribeAlarmEventDetailResponseBodyDataCauseDetailsValue {
	s.Name = &v
	return s
}

type DescribeAlarmEventDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlarmEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlarmEventDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailResponse) SetHeaders(v map[string]*string) *DescribeAlarmEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmEventDetailResponse) SetBody(v *DescribeAlarmEventDetailResponseBody) *DescribeAlarmEventDetailResponse {
	s.Body = v
	return s
}

type DescribeAlarmEventListRequest struct {
	SourceIp             *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                 *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Dealed               *string   `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	From                 *string   `json:"From,omitempty" xml:"From,omitempty"`
	Levels               *string   `json:"Levels,omitempty" xml:"Levels,omitempty"`
	Remark               *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	GroupId              *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	AlarmEventName       *string   `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	AlarmEventType       *string   `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	CurrentPage          *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize             *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClusterId            *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerFieldName   *string   `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	ContainerFieldValue  *string   `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	TargetType           *string   `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	OperateErrorCodeList []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
}

func (s DescribeAlarmEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventListRequest) SetSourceIp(v string) *DescribeAlarmEventListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetLang(v string) *DescribeAlarmEventListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetDealed(v string) *DescribeAlarmEventListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetFrom(v string) *DescribeAlarmEventListRequest {
	s.From = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetLevels(v string) *DescribeAlarmEventListRequest {
	s.Levels = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetRemark(v string) *DescribeAlarmEventListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetGroupId(v string) *DescribeAlarmEventListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetAlarmEventName(v string) *DescribeAlarmEventListRequest {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetAlarmEventType(v string) *DescribeAlarmEventListRequest {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetCurrentPage(v int32) *DescribeAlarmEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetPageSize(v string) *DescribeAlarmEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetClusterId(v string) *DescribeAlarmEventListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetContainerFieldName(v string) *DescribeAlarmEventListRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetContainerFieldValue(v string) *DescribeAlarmEventListRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetTargetType(v string) *DescribeAlarmEventListRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeAlarmEventListRequest) SetOperateErrorCodeList(v []*string) *DescribeAlarmEventListRequest {
	s.OperateErrorCodeList = v
	return s
}

type DescribeAlarmEventListResponseBody struct {
	PageInfo   *DescribeAlarmEventListResponseBodyPageInfo     `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuspEvents []*DescribeAlarmEventListResponseBodySuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" type:"Repeated"`
}

func (s DescribeAlarmEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventListResponseBody) SetPageInfo(v *DescribeAlarmEventListResponseBodyPageInfo) *DescribeAlarmEventListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAlarmEventListResponseBody) SetRequestId(v string) *DescribeAlarmEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmEventListResponseBody) SetSuspEvents(v []*DescribeAlarmEventListResponseBodySuspEvents) *DescribeAlarmEventListResponseBody {
	s.SuspEvents = v
	return s
}

type DescribeAlarmEventListResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAlarmEventListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAlarmEventListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodyPageInfo) SetPageSize(v int32) *DescribeAlarmEventListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAlarmEventListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodyPageInfo) SetCount(v int32) *DescribeAlarmEventListResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeAlarmEventListResponseBodySuspEvents struct {
	Dealed                 *bool   `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Stages                 *string `json:"Stages,omitempty" xml:"Stages,omitempty"`
	InternetIp             *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	SuspiciousEventCount   *int32  `json:"SuspiciousEventCount,omitempty" xml:"SuspiciousEventCount,omitempty"`
	K8sClusterName         *string `json:"K8sClusterName,omitempty" xml:"K8sClusterName,omitempty"`
	ContainerImageId       *string `json:"ContainerImageId,omitempty" xml:"ContainerImageId,omitempty"`
	GmtModified            *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	AlarmEventNameOriginal *string `json:"AlarmEventNameOriginal,omitempty" xml:"AlarmEventNameOriginal,omitempty"`
	AlarmUniqueInfo        *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	CanCancelFault         *bool   `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	AppName                *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	SecurityEventIds       *string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty"`
	CanBeDealOnLine        *bool   `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	ContainerImageName     *string `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	K8sClusterId           *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ContainHwMode          *bool   `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	InstanceName           *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	K8sNodeId              *string `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	SaleVersion            *string `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	OperateErrorCode       *string `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	Solution               *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	HasTraceInfo           *bool   `json:"HasTraceInfo,omitempty" xml:"HasTraceInfo,omitempty"`
	OperateTime            *int64  `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	DataSource             *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntranetIp             *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	EndTime                *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Uuid                   *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	StartTime              *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	K8sPodName             *string `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	ContainerId            *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	AlarmEventType         *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	K8sNamespace           *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	K8sNodeName            *string `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	AlarmEventName         *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	Level                  *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeAlarmEventListResponseBodySuspEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventListResponseBodySuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetDealed(v bool) *DescribeAlarmEventListResponseBodySuspEvents {
	s.Dealed = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetStages(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.Stages = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetInternetIp(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.InternetIp = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetSuspiciousEventCount(v int32) *DescribeAlarmEventListResponseBodySuspEvents {
	s.SuspiciousEventCount = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetK8sClusterName(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.K8sClusterName = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetContainerImageId(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.ContainerImageId = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetGmtModified(v int64) *DescribeAlarmEventListResponseBodySuspEvents {
	s.GmtModified = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetAlarmEventNameOriginal(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.AlarmEventNameOriginal = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetAlarmUniqueInfo(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetCanCancelFault(v bool) *DescribeAlarmEventListResponseBodySuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetAppName(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.AppName = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetSecurityEventIds(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.SecurityEventIds = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetCanBeDealOnLine(v bool) *DescribeAlarmEventListResponseBodySuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetContainerImageName(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.ContainerImageName = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetK8sClusterId(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetDescription(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.Description = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetContainHwMode(v bool) *DescribeAlarmEventListResponseBodySuspEvents {
	s.ContainHwMode = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetInstanceName(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.InstanceName = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetK8sNodeId(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetSaleVersion(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.SaleVersion = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetOperateErrorCode(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetSolution(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.Solution = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetHasTraceInfo(v bool) *DescribeAlarmEventListResponseBodySuspEvents {
	s.HasTraceInfo = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetOperateTime(v int64) *DescribeAlarmEventListResponseBodySuspEvents {
	s.OperateTime = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetDataSource(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetInstanceId(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.InstanceId = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetIntranetIp(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetEndTime(v int64) *DescribeAlarmEventListResponseBodySuspEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetUuid(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.Uuid = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetStartTime(v int64) *DescribeAlarmEventListResponseBodySuspEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetK8sPodName(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.K8sPodName = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetContainerId(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.ContainerId = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetAlarmEventType(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetK8sNamespace(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetK8sNodeName(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetAlarmEventName(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeAlarmEventListResponseBodySuspEvents) SetLevel(v string) *DescribeAlarmEventListResponseBodySuspEvents {
	s.Level = &v
	return s
}

type DescribeAlarmEventListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAlarmEventListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAlarmEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAlarmEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventListResponse) SetHeaders(v map[string]*string) *DescribeAlarmEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlarmEventListResponse) SetBody(v *DescribeAlarmEventListResponseBody) *DescribeAlarmEventListResponse {
	s.Body = v
	return s
}

type DescribeAllEntityRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAllEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEntityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllEntityRequest) SetSourceIp(v string) *DescribeAllEntityRequest {
	s.SourceIp = &v
	return s
}

type DescribeAllEntityResponseBody struct {
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EntityList []*DescribeAllEntityResponseBodyEntityList `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
}

func (s DescribeAllEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllEntityResponseBody) SetRequestId(v string) *DescribeAllEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllEntityResponseBody) SetEntityList(v []*DescribeAllEntityResponseBodyEntityList) *DescribeAllEntityResponseBody {
	s.EntityList = v
	return s
}

type DescribeAllEntityResponseBodyEntityList struct {
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	GroupId      *int32  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Os           *string `json:"Os,omitempty" xml:"Os,omitempty"`
	IntranetIp   *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
}

func (s DescribeAllEntityResponseBodyEntityList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEntityResponseBodyEntityList) GoString() string {
	return s.String()
}

func (s *DescribeAllEntityResponseBodyEntityList) SetUuid(v string) *DescribeAllEntityResponseBodyEntityList {
	s.Uuid = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetGroupId(v int32) *DescribeAllEntityResponseBodyEntityList {
	s.GroupId = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetInternetIp(v string) *DescribeAllEntityResponseBodyEntityList {
	s.InternetIp = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetInstanceName(v string) *DescribeAllEntityResponseBodyEntityList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetIp(v string) *DescribeAllEntityResponseBodyEntityList {
	s.Ip = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetOs(v string) *DescribeAllEntityResponseBodyEntityList {
	s.Os = &v
	return s
}

func (s *DescribeAllEntityResponseBodyEntityList) SetIntranetIp(v string) *DescribeAllEntityResponseBodyEntityList {
	s.IntranetIp = &v
	return s
}

type DescribeAllEntityResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEntityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllEntityResponse) SetHeaders(v map[string]*string) *DescribeAllEntityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllEntityResponse) SetBody(v *DescribeAllEntityResponseBody) *DescribeAllEntityResponse {
	s.Body = v
	return s
}

type DescribeAllGroupsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAllGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsRequest) SetSourceIp(v string) *DescribeAllGroupsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAllGroupsRequest) SetLang(v string) *DescribeAllGroupsRequest {
	s.Lang = &v
	return s
}

type DescribeAllGroupsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Groups    []*DescribeAllGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Count     *int32                                 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAllGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsResponseBody) SetRequestId(v string) *DescribeAllGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllGroupsResponseBody) SetGroups(v []*DescribeAllGroupsResponseBodyGroups) *DescribeAllGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeAllGroupsResponseBody) SetCount(v int32) *DescribeAllGroupsResponseBody {
	s.Count = &v
	return s
}

type DescribeAllGroupsResponseBodyGroups struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId   *int32  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupFlag *int32  `json:"GroupFlag,omitempty" xml:"GroupFlag,omitempty"`
}

func (s DescribeAllGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsResponseBodyGroups) SetGroupName(v string) *DescribeAllGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeAllGroupsResponseBodyGroups) SetGroupId(v int32) *DescribeAllGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeAllGroupsResponseBodyGroups) SetGroupFlag(v int32) *DescribeAllGroupsResponseBodyGroups {
	s.GroupFlag = &v
	return s
}

type DescribeAllGroupsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsResponse) SetHeaders(v map[string]*string) *DescribeAllGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllGroupsResponse) SetBody(v *DescribeAllGroupsResponseBody) *DescribeAllGroupsResponse {
	s.Body = v
	return s
}

type DescribeAllRegionsStatisticsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	From        *string `json:"From,omitempty" xml:"From,omitempty"`
	GroupId     *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusList  *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	WebGroupId  *int64  `json:"WebGroupId,omitempty" xml:"WebGroupId,omitempty"`
	RuleType    *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Action1     *int32  `json:"Action1,omitempty" xml:"Action1,omitempty"`
	Flow        *int32  `json:"Flow,omitempty" xml:"Flow,omitempty"`
	SaleId      *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	Dealed      *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecureToken *string `json:"SecureToken,omitempty" xml:"SecureToken,omitempty"`
	AllRegion   *bool   `json:"AllRegion,omitempty" xml:"AllRegion,omitempty"`
}

func (s DescribeAllRegionsStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllRegionsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllRegionsStatisticsRequest) SetSourceIp(v string) *DescribeAllRegionsStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetFrom(v string) *DescribeAllRegionsStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetGroupId(v int64) *DescribeAllRegionsStatisticsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetRemark(v string) *DescribeAllRegionsStatisticsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetType(v string) *DescribeAllRegionsStatisticsRequest {
	s.Type = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetUuid(v string) *DescribeAllRegionsStatisticsRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetStatus(v int32) *DescribeAllRegionsStatisticsRequest {
	s.Status = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetStatusList(v string) *DescribeAllRegionsStatisticsRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetStartTime(v string) *DescribeAllRegionsStatisticsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetEndTime(v string) *DescribeAllRegionsStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetWebGroupId(v int64) *DescribeAllRegionsStatisticsRequest {
	s.WebGroupId = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetRuleType(v int32) *DescribeAllRegionsStatisticsRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetAction1(v int32) *DescribeAllRegionsStatisticsRequest {
	s.Action1 = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetFlow(v int32) *DescribeAllRegionsStatisticsRequest {
	s.Flow = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetSaleId(v string) *DescribeAllRegionsStatisticsRequest {
	s.SaleId = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetDealed(v string) *DescribeAllRegionsStatisticsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetTag(v string) *DescribeAllRegionsStatisticsRequest {
	s.Tag = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetCurrentPage(v int32) *DescribeAllRegionsStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetPageSize(v int32) *DescribeAllRegionsStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetSecureToken(v string) *DescribeAllRegionsStatisticsRequest {
	s.SecureToken = &v
	return s
}

func (s *DescribeAllRegionsStatisticsRequest) SetAllRegion(v bool) *DescribeAllRegionsStatisticsRequest {
	s.AllRegion = &v
	return s
}

type DescribeAllRegionsStatisticsResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeAllRegionsStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeAllRegionsStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllRegionsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllRegionsStatisticsResponseBody) SetRequestId(v string) *DescribeAllRegionsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllRegionsStatisticsResponseBody) SetData(v *DescribeAllRegionsStatisticsResponseBodyData) *DescribeAllRegionsStatisticsResponseBody {
	s.Data = v
	return s
}

type DescribeAllRegionsStatisticsResponseBodyData struct {
	Account       *int32 `json:"Account,omitempty" xml:"Account,omitempty"`
	Vul           *int32 `json:"Vul,omitempty" xml:"Vul,omitempty"`
	Health        *int32 `json:"Health,omitempty" xml:"Health,omitempty"`
	Trojan        *int32 `json:"Trojan,omitempty" xml:"Trojan,omitempty"`
	NewSuspicious *int32 `json:"NewSuspicious,omitempty" xml:"NewSuspicious,omitempty"`
	Suspicious    *int32 `json:"Suspicious,omitempty" xml:"Suspicious,omitempty"`
}

func (s DescribeAllRegionsStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllRegionsStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAllRegionsStatisticsResponseBodyData) SetAccount(v int32) *DescribeAllRegionsStatisticsResponseBodyData {
	s.Account = &v
	return s
}

func (s *DescribeAllRegionsStatisticsResponseBodyData) SetVul(v int32) *DescribeAllRegionsStatisticsResponseBodyData {
	s.Vul = &v
	return s
}

func (s *DescribeAllRegionsStatisticsResponseBodyData) SetHealth(v int32) *DescribeAllRegionsStatisticsResponseBodyData {
	s.Health = &v
	return s
}

func (s *DescribeAllRegionsStatisticsResponseBodyData) SetTrojan(v int32) *DescribeAllRegionsStatisticsResponseBodyData {
	s.Trojan = &v
	return s
}

func (s *DescribeAllRegionsStatisticsResponseBodyData) SetNewSuspicious(v int32) *DescribeAllRegionsStatisticsResponseBodyData {
	s.NewSuspicious = &v
	return s
}

func (s *DescribeAllRegionsStatisticsResponseBodyData) SetSuspicious(v int32) *DescribeAllRegionsStatisticsResponseBodyData {
	s.Suspicious = &v
	return s
}

type DescribeAllRegionsStatisticsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllRegionsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllRegionsStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllRegionsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllRegionsStatisticsResponse) SetHeaders(v map[string]*string) *DescribeAllRegionsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllRegionsStatisticsResponse) SetBody(v *DescribeAllRegionsStatisticsResponseBody) *DescribeAllRegionsStatisticsResponse {
	s.Body = v
	return s
}

type DescribeAntiBruteForceRulesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAntiBruteForceRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntiBruteForceRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesRequest) SetSourceIp(v string) *DescribeAntiBruteForceRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAntiBruteForceRulesRequest) SetResourceOwnerId(v int64) *DescribeAntiBruteForceRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAntiBruteForceRulesResponseBody struct {
	PageInfo  *DescribeAntiBruteForceRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeAntiBruteForceRulesResponseBodyRules  `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeAntiBruteForceRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponseBody) SetPageInfo(v *DescribeAntiBruteForceRulesResponseBodyPageInfo) *DescribeAntiBruteForceRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBody) SetRequestId(v string) *DescribeAntiBruteForceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBody) SetRules(v []*DescribeAntiBruteForceRulesResponseBodyRules) *DescribeAntiBruteForceRulesResponseBody {
	s.Rules = v
	return s
}

type DescribeAntiBruteForceRulesResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeAntiBruteForceRulesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetPageSize(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyPageInfo) SetCount(v int32) *DescribeAntiBruteForceRulesResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeAntiBruteForceRulesResponseBodyRules struct {
	UuidList        []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
	MachineCount    *int32    `json:"MachineCount,omitempty" xml:"MachineCount,omitempty"`
	EnableSmartRule *bool     `json:"EnableSmartRule,omitempty" xml:"EnableSmartRule,omitempty"`
	FailCount       *int32    `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	ForbiddenTime   *int32    `json:"ForbiddenTime,omitempty" xml:"ForbiddenTime,omitempty"`
	Span            *int32    `json:"Span,omitempty" xml:"Span,omitempty"`
	DefaultRule     *bool     `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Id              *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeAntiBruteForceRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetUuidList(v []*string) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.UuidList = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetMachineCount(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.MachineCount = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetEnableSmartRule(v bool) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.EnableSmartRule = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetFailCount(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.FailCount = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetForbiddenTime(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.ForbiddenTime = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetSpan(v int32) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.Span = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetDefaultRule(v bool) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.DefaultRule = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetName(v string) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *DescribeAntiBruteForceRulesResponseBodyRules) SetId(v int64) *DescribeAntiBruteForceRulesResponseBodyRules {
	s.Id = &v
	return s
}

type DescribeAntiBruteForceRulesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAntiBruteForceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAntiBruteForceRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAntiBruteForceRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAntiBruteForceRulesResponse) SetHeaders(v map[string]*string) *DescribeAntiBruteForceRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAntiBruteForceRulesResponse) SetBody(v *DescribeAntiBruteForceRulesResponseBody) *DescribeAntiBruteForceRulesResponse {
	s.Body = v
	return s
}

type DescribeAssetDetailByUuidRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAssetDetailByUuidRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidRequest) SetSourceIp(v string) *DescribeAssetDetailByUuidRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidRequest) SetUuid(v string) *DescribeAssetDetailByUuidRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeAssetDetailByUuidRequest) SetLang(v string) *DescribeAssetDetailByUuidRequest {
	s.Lang = &v
	return s
}

type DescribeAssetDetailByUuidResponseBody struct {
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AssetDetail *DescribeAssetDetailByUuidResponseBodyAssetDetail `json:"AssetDetail,omitempty" xml:"AssetDetail,omitempty" type:"Struct"`
}

func (s DescribeAssetDetailByUuidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidResponseBody) SetRequestId(v string) *DescribeAssetDetailByUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBody) SetAssetDetail(v *DescribeAssetDetailByUuidResponseBodyAssetDetail) *DescribeAssetDetailByUuidResponseBody {
	s.AssetDetail = v
	return s
}

type DescribeAssetDetailByUuidResponseBodyAssetDetail struct {
	CpuInfo        *string   `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	InternetIp     *string   `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	OsDetail       *string   `json:"OsDetail,omitempty" xml:"OsDetail,omitempty"`
	Kernel         *string   `json:"Kernel,omitempty" xml:"Kernel,omitempty"`
	CreateTime     *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OsName         *string   `json:"OsName,omitempty" xml:"OsName,omitempty"`
	Tag            *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	ClientStatus   *string   `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	Mem            *int32    `json:"Mem,omitempty" xml:"Mem,omitempty"`
	VpcInstanceId  *string   `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	InstanceName   *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Region         *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	IpList         []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	GroupTrace     *string   `json:"GroupTrace,omitempty" xml:"GroupTrace,omitempty"`
	DiskInfoList   []*string `json:"DiskInfoList,omitempty" xml:"DiskInfoList,omitempty" type:"Repeated"`
	HostName       *string   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Ip             *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	MacList        []*string `json:"MacList,omitempty" xml:"MacList,omitempty" type:"Repeated"`
	Os             *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	InstanceId     *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AssetType      *string   `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	IntranetIp     *string   `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	SysInfo        *string   `json:"SysInfo,omitempty" xml:"SysInfo,omitempty"`
	Uuid           *string   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Cpu            *int32    `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	RegionName     *string   `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	InstanceStatus *string   `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
}

func (s DescribeAssetDetailByUuidResponseBodyAssetDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidResponseBodyAssetDetail) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetCpuInfo(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.CpuInfo = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetInternetIp(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.InternetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetOsDetail(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.OsDetail = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetKernel(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Kernel = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetCreateTime(v int64) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.CreateTime = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetOsName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.OsName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetTag(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Tag = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetClientStatus(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.ClientStatus = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetMem(v int32) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Mem = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetVpcInstanceId(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetInstanceName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.InstanceName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetRegion(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Region = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetIpList(v []*string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.IpList = v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetGroupTrace(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.GroupTrace = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetDiskInfoList(v []*string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.DiskInfoList = v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetHostName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.HostName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetIp(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Ip = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetMacList(v []*string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.MacList = v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetOs(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Os = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetInstanceId(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetAssetType(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.AssetType = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetIntranetIp(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetSysInfo(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.SysInfo = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetUuid(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Uuid = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetCpu(v int32) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetRegionName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.RegionName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetInstanceStatus(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.InstanceStatus = &v
	return s
}

type DescribeAssetDetailByUuidResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAssetDetailByUuidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssetDetailByUuidResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidResponse) SetHeaders(v map[string]*string) *DescribeAssetDetailByUuidResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetDetailByUuidResponse) SetBody(v *DescribeAssetDetailByUuidResponseBody) *DescribeAssetDetailByUuidResponse {
	s.Body = v
	return s
}

type DescribeAssetDetailByUuidsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuids    *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeAssetDetailByUuidsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsRequest) SetSourceIp(v string) *DescribeAssetDetailByUuidsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidsRequest) SetUuids(v string) *DescribeAssetDetailByUuidsRequest {
	s.Uuids = &v
	return s
}

type DescribeAssetDetailByUuidsResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AssetList []*DescribeAssetDetailByUuidsResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
}

func (s DescribeAssetDetailByUuidsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsResponseBody) SetRequestId(v string) *DescribeAssetDetailByUuidsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBody) SetAssetList(v []*DescribeAssetDetailByUuidsResponseBodyAssetList) *DescribeAssetDetailByUuidsResponseBody {
	s.AssetList = v
	return s
}

type DescribeAssetDetailByUuidsResponseBodyAssetList struct {
	InternetIp    *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	OsName        *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Os            *string `json:"Os,omitempty" xml:"Os,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ClientStatus  *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	AssetType     *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	IntranetIp    *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	RegionName    *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeAssetDetailByUuidsResponseBodyAssetList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidsResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetInternetIp(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.InternetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetOsName(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.OsName = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetIp(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Ip = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetOs(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Os = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetInstanceId(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetClientStatus(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.ClientStatus = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetVpcInstanceId(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetAssetType(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.AssetType = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetIntranetIp(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetRegionId(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetUuid(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Uuid = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetRegionName(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.RegionName = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetInstanceName(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetRegion(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Region = &v
	return s
}

type DescribeAssetDetailByUuidsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAssetDetailByUuidsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssetDetailByUuidsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetDetailByUuidsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsResponse) SetHeaders(v map[string]*string) *DescribeAssetDetailByUuidsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetDetailByUuidsResponse) SetBody(v *DescribeAssetDetailByUuidsResponseBody) *DescribeAssetDetailByUuidsResponse {
	s.Body = v
	return s
}

type DescribeAutoDelConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAutoDelConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoDelConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoDelConfigRequest) SetSourceIp(v string) *DescribeAutoDelConfigRequest {
	s.SourceIp = &v
	return s
}

type DescribeAutoDelConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Days      *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s DescribeAutoDelConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoDelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoDelConfigResponseBody) SetRequestId(v string) *DescribeAutoDelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoDelConfigResponseBody) SetDays(v int32) *DescribeAutoDelConfigResponseBody {
	s.Days = &v
	return s
}

type DescribeAutoDelConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoDelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoDelConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoDelConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoDelConfigResponse) SetHeaders(v map[string]*string) *DescribeAutoDelConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoDelConfigResponse) SetBody(v *DescribeAutoDelConfigResponseBody) *DescribeAutoDelConfigResponse {
	s.Body = v
	return s
}

type DescribeBruteForceSummaryRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBruteForceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBruteForceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryRequest) SetSourceIp(v string) *DescribeBruteForceSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBruteForceSummaryRequest) SetResourceOwnerId(v int64) *DescribeBruteForceSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeBruteForceSummaryResponseBody struct {
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BruteForceSummary *DescribeBruteForceSummaryResponseBodyBruteForceSummary `json:"BruteForceSummary,omitempty" xml:"BruteForceSummary,omitempty" type:"Struct"`
}

func (s DescribeBruteForceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBruteForceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryResponseBody) SetRequestId(v string) *DescribeBruteForceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBody) SetBruteForceSummary(v *DescribeBruteForceSummaryResponseBodyBruteForceSummary) *DescribeBruteForceSummaryResponseBody {
	s.BruteForceSummary = v
	return s
}

type DescribeBruteForceSummaryResponseBodyBruteForceSummary struct {
	AllStrategyCount *int32 `json:"AllStrategyCount,omitempty" xml:"AllStrategyCount,omitempty"`
	EffectiveCount   *int32 `json:"EffectiveCount,omitempty" xml:"EffectiveCount,omitempty"`
}

func (s DescribeBruteForceSummaryResponseBodyBruteForceSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeBruteForceSummaryResponseBodyBruteForceSummary) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetAllStrategyCount(v int32) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.AllStrategyCount = &v
	return s
}

func (s *DescribeBruteForceSummaryResponseBodyBruteForceSummary) SetEffectiveCount(v int32) *DescribeBruteForceSummaryResponseBodyBruteForceSummary {
	s.EffectiveCount = &v
	return s
}

type DescribeBruteForceSummaryResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBruteForceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBruteForceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBruteForceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeBruteForceSummaryResponse) SetHeaders(v map[string]*string) *DescribeBruteForceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeBruteForceSummaryResponse) SetBody(v *DescribeBruteForceSummaryResponseBody) *DescribeBruteForceSummaryResponse {
	s.Body = v
	return s
}

type DescribeCheckEcsWarningsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s DescribeCheckEcsWarningsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckEcsWarningsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckEcsWarningsRequest) SetSourceIp(v string) *DescribeCheckEcsWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckEcsWarningsRequest) SetUuidList(v string) *DescribeCheckEcsWarningsRequest {
	s.UuidList = &v
	return s
}

type DescribeCheckEcsWarningsResponseBody struct {
	SasVersion        *string `json:"SasVersion,omitempty" xml:"SasVersion,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CanTry            *string `json:"CanTry,omitempty" xml:"CanTry,omitempty"`
	WeakPasswordCount *string `json:"WeakPasswordCount,omitempty" xml:"WeakPasswordCount,omitempty"`
}

func (s DescribeCheckEcsWarningsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckEcsWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckEcsWarningsResponseBody) SetSasVersion(v string) *DescribeCheckEcsWarningsResponseBody {
	s.SasVersion = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponseBody) SetRequestId(v string) *DescribeCheckEcsWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponseBody) SetCanTry(v string) *DescribeCheckEcsWarningsResponseBody {
	s.CanTry = &v
	return s
}

func (s *DescribeCheckEcsWarningsResponseBody) SetWeakPasswordCount(v string) *DescribeCheckEcsWarningsResponseBody {
	s.WeakPasswordCount = &v
	return s
}

type DescribeCheckEcsWarningsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCheckEcsWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCheckEcsWarningsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckEcsWarningsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckEcsWarningsResponse) SetHeaders(v map[string]*string) *DescribeCheckEcsWarningsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckEcsWarningsResponse) SetBody(v *DescribeCheckEcsWarningsResponseBody) *DescribeCheckEcsWarningsResponse {
	s.Body = v
	return s
}

type DescribeCheckWarningDetailRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	CheckWarningId *int64  `json:"CheckWarningId,omitempty" xml:"CheckWarningId,omitempty"`
}

func (s DescribeCheckWarningDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailRequest) SetSourceIp(v string) *DescribeCheckWarningDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetLang(v string) *DescribeCheckWarningDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetCheckWarningId(v int64) *DescribeCheckWarningDetailRequest {
	s.CheckWarningId = &v
	return s
}

type DescribeCheckWarningDetailResponseBody struct {
	Item        *string `json:"Item,omitempty" xml:"Item,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CheckId     *int64  `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Level       *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Prompt      *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Advice      *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
}

func (s DescribeCheckWarningDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailResponseBody) SetItem(v string) *DescribeCheckWarningDetailResponseBody {
	s.Item = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetType(v string) *DescribeCheckWarningDetailResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetDescription(v string) *DescribeCheckWarningDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetRequestId(v string) *DescribeCheckWarningDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetCheckId(v int64) *DescribeCheckWarningDetailResponseBody {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetLevel(v string) *DescribeCheckWarningDetailResponseBody {
	s.Level = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetPrompt(v string) *DescribeCheckWarningDetailResponseBody {
	s.Prompt = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetAdvice(v string) *DescribeCheckWarningDetailResponseBody {
	s.Advice = &v
	return s
}

type DescribeCheckWarningDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCheckWarningDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCheckWarningDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningDetailResponse) SetBody(v *DescribeCheckWarningDetailResponseBody) *DescribeCheckWarningDetailResponse {
	s.Body = v
	return s
}

type DescribeCheckWarningsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	RiskId      *int64  `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeCheckWarningsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsRequest) SetSourceIp(v string) *DescribeCheckWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetLang(v string) *DescribeCheckWarningsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetUuid(v string) *DescribeCheckWarningsRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetRiskId(v int64) *DescribeCheckWarningsRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetPageSize(v int32) *DescribeCheckWarningsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetCurrentPage(v int32) *DescribeCheckWarningsRequest {
	s.CurrentPage = &v
	return s
}

type DescribeCheckWarningsResponseBody struct {
	TotalCount    *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	CheckWarnings []*DescribeCheckWarningsResponseBodyCheckWarnings `json:"CheckWarnings,omitempty" xml:"CheckWarnings,omitempty" type:"Repeated"`
	PageSize      *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage   *int32                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Count         *int32                                            `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeCheckWarningsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsResponseBody) SetTotalCount(v int32) *DescribeCheckWarningsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetCheckWarnings(v []*DescribeCheckWarningsResponseBodyCheckWarnings) *DescribeCheckWarningsResponseBody {
	s.CheckWarnings = v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetPageSize(v int32) *DescribeCheckWarningsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetRequestId(v string) *DescribeCheckWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetCurrentPage(v int32) *DescribeCheckWarningsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningsResponseBody) SetCount(v int32) *DescribeCheckWarningsResponseBody {
	s.Count = &v
	return s
}

type DescribeCheckWarningsResponseBodyCheckWarnings struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	CheckWarningId *int64  `json:"CheckWarningId,omitempty" xml:"CheckWarningId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuid           *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Item           *string `json:"Item,omitempty" xml:"Item,omitempty"`
	CheckId        *int64  `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeCheckWarningsResponseBodyCheckWarnings) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningsResponseBodyCheckWarnings) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetStatus(v int32) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Status = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetCheckWarningId(v int64) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.CheckWarningId = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetType(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Type = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetUuid(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Uuid = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetItem(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Item = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetCheckId(v int64) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningsResponseBodyCheckWarnings) SetLevel(v string) *DescribeCheckWarningsResponseBodyCheckWarnings {
	s.Level = &v
	return s
}

type DescribeCheckWarningsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCheckWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCheckWarningsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningsResponse) SetBody(v *DescribeCheckWarningsResponseBody) *DescribeCheckWarningsResponse {
	s.Body = v
	return s
}

type DescribeCheckWarningSummaryRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TypeName    *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RiskStatus  *int32  `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	RiskName    *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	StrategyId  *int64  `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Uuids       *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeCheckWarningSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryRequest) SetSourceIp(v string) *DescribeCheckWarningSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetLang(v string) *DescribeCheckWarningSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetTypeName(v string) *DescribeCheckWarningSummaryRequest {
	s.TypeName = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetStatus(v string) *DescribeCheckWarningSummaryRequest {
	s.Status = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetRiskStatus(v int32) *DescribeCheckWarningSummaryRequest {
	s.RiskStatus = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetRiskName(v string) *DescribeCheckWarningSummaryRequest {
	s.RiskName = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetStrategyId(v int64) *DescribeCheckWarningSummaryRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetUuids(v string) *DescribeCheckWarningSummaryRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetPageSize(v int32) *DescribeCheckWarningSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetCurrentPage(v int32) *DescribeCheckWarningSummaryRequest {
	s.CurrentPage = &v
	return s
}

type DescribeCheckWarningSummaryResponseBody struct {
	TotalCount      *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage     *int32                                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Count           *int32                                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	WarningSummarys []*DescribeCheckWarningSummaryResponseBodyWarningSummarys `json:"WarningSummarys,omitempty" xml:"WarningSummarys,omitempty" type:"Repeated"`
}

func (s DescribeCheckWarningSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryResponseBody) SetTotalCount(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetPageSize(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetRequestId(v string) *DescribeCheckWarningSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetCurrentPage(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetCount(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetWarningSummarys(v []*DescribeCheckWarningSummaryResponseBodyWarningSummarys) *DescribeCheckWarningSummaryResponseBody {
	s.WarningSummarys = v
	return s
}

type DescribeCheckWarningSummaryResponseBodyWarningSummarys struct {
	LowWarningCount     *int32  `json:"LowWarningCount,omitempty" xml:"LowWarningCount,omitempty"`
	CheckCount          *int32  `json:"CheckCount,omitempty" xml:"CheckCount,omitempty"`
	MediumWarningCount  *int32  `json:"MediumWarningCount,omitempty" xml:"MediumWarningCount,omitempty"`
	LastFoundTime       *string `json:"LastFoundTime,omitempty" xml:"LastFoundTime,omitempty"`
	RiskId              *int64  `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	SubTypeAlias        *string `json:"SubTypeAlias,omitempty" xml:"SubTypeAlias,omitempty"`
	WarningMachineCount *int32  `json:"WarningMachineCount,omitempty" xml:"WarningMachineCount,omitempty"`
	HighWarningCount    *int32  `json:"HighWarningCount,omitempty" xml:"HighWarningCount,omitempty"`
	TypeAlias           *string `json:"TypeAlias,omitempty" xml:"TypeAlias,omitempty"`
	RiskName            *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	Level               *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeCheckWarningSummaryResponseBodyWarningSummarys) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningSummaryResponseBodyWarningSummarys) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetLowWarningCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.LowWarningCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetCheckCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.CheckCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetMediumWarningCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.MediumWarningCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetLastFoundTime(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.LastFoundTime = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetRiskId(v int64) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.RiskId = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetSubTypeAlias(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.SubTypeAlias = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetWarningMachineCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.WarningMachineCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetHighWarningCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.HighWarningCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetTypeAlias(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.TypeAlias = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetRiskName(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.RiskName = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetLevel(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.Level = &v
	return s
}

type DescribeCheckWarningSummaryResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCheckWarningSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCheckWarningSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningSummaryResponse) SetBody(v *DescribeCheckWarningSummaryResponseBody) *DescribeCheckWarningSummaryResponse {
	s.Body = v
	return s
}

type DescribeCloudCenterInstancesRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Criteria     *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	LogicalExp   *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	NoPage       *bool   `json:"NoPage,omitempty" xml:"NoPage,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage  *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Importance   *int32  `json:"Importance,omitempty" xml:"Importance,omitempty"`
}

func (s DescribeCloudCenterInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudCenterInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesRequest) SetSourceIp(v string) *DescribeCloudCenterInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetRegionId(v string) *DescribeCloudCenterInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetCriteria(v string) *DescribeCloudCenterInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetMachineTypes(v string) *DescribeCloudCenterInstancesRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetLogicalExp(v string) *DescribeCloudCenterInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetNoPage(v bool) *DescribeCloudCenterInstancesRequest {
	s.NoPage = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetPageSize(v int32) *DescribeCloudCenterInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetCurrentPage(v int32) *DescribeCloudCenterInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudCenterInstancesRequest) SetImportance(v int32) *DescribeCloudCenterInstancesRequest {
	s.Importance = &v
	return s
}

type DescribeCloudCenterInstancesResponseBody struct {
	Instances []*DescribeCloudCenterInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageInfo  *DescribeCloudCenterInstancesResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBody) SetInstances(v []*DescribeCloudCenterInstancesResponseBodyInstances) *DescribeCloudCenterInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetPageInfo(v *DescribeCloudCenterInstancesResponseBodyPageInfo) *DescribeCloudCenterInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetRequestId(v string) *DescribeCloudCenterInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBody) SetSuccess(v bool) *DescribeCloudCenterInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeCloudCenterInstancesResponseBodyInstances struct {
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InternetIp       *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	OsName           *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	Tag              *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	ClientStatus     *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	VpcInstanceId    *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	Flag             *int32  `json:"Flag,omitempty" xml:"Flag,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PodCount         *int32  `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	VulCount         *int32  `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	HcStatus         *string `json:"HcStatus,omitempty" xml:"HcStatus,omitempty"`
	CreatedTime      *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RiskStatus       *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	VulStatus        *string `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
	AlarmStatus      *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	HealthCheckCount *int32  `json:"HealthCheckCount,omitempty" xml:"HealthCheckCount,omitempty"`
	Importance       *int32  `json:"Importance,omitempty" xml:"Importance,omitempty"`
	Ip               *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Os               *string `json:"Os,omitempty" xml:"Os,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SafeEventCount   *int32  `json:"SafeEventCount,omitempty" xml:"SafeEventCount,omitempty"`
	AssetType        *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	IntranetIp       *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Uuid             *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionName       *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	ClusterName      *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ExposedStatus    *int32  `json:"ExposedStatus,omitempty" xml:"ExposedStatus,omitempty"`
	RiskCount        *string `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	ClientVersion    *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetInternetIp(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetOsName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.OsName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetTag(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Tag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClientStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClientStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVpcInstanceId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetFlag(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Flag = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRegion(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Region = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetPodCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.PodCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVulCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VulCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetHcStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.HcStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetCreatedTime(v int64) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.CreatedTime = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClusterId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClusterId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRiskStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RiskStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetVulStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.VulStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAlarmStatus(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AlarmStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetHealthCheckCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.HealthCheckCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetImportance(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Importance = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetIp(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Ip = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetOs(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Os = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetSafeEventCount(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.SafeEventCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetAssetType(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetIntranetIp(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRegionId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetUuid(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.Uuid = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetGroupId(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.GroupId = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRegionName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RegionName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClusterName(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClusterName = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetExposedStatus(v int32) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ExposedStatus = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetRiskCount(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.RiskCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyInstances) SetClientVersion(v string) *DescribeCloudCenterInstancesResponseBodyInstances {
	s.ClientVersion = &v
	return s
}

type DescribeCloudCenterInstancesResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeCloudCenterInstancesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeCloudCenterInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeCloudCenterInstancesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCloudCenterInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudCenterInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponse) SetHeaders(v map[string]*string) *DescribeCloudCenterInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudCenterInstancesResponse) SetBody(v *DescribeCloudCenterInstancesResponseBody) *DescribeCloudCenterInstancesResponse {
	s.Body = v
	return s
}

type DescribeCloudProductFieldStatisticsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCloudProductFieldStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudProductFieldStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudProductFieldStatisticsRequest) SetSourceIp(v string) *DescribeCloudProductFieldStatisticsRequest {
	s.SourceIp = &v
	return s
}

type DescribeCloudProductFieldStatisticsResponseBody struct {
	RequestId     *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupedFields *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields `json:"GroupedFields,omitempty" xml:"GroupedFields,omitempty" type:"Struct"`
}

func (s DescribeCloudProductFieldStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudProductFieldStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudProductFieldStatisticsResponseBody) SetRequestId(v string) *DescribeCloudProductFieldStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBody) SetGroupedFields(v *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) *DescribeCloudProductFieldStatisticsResponseBody {
	s.GroupedFields = v
	return s
}

type DescribeCloudProductFieldStatisticsResponseBodyGroupedFields struct {
	CategoryCount     *string `json:"CategoryCount,omitempty" xml:"CategoryCount,omitempty"`
	InstanceCount     *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	RiskInstanceCount *int32  `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
}

func (s DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) GoString() string {
	return s.String()
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) SetCategoryCount(v string) *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields {
	s.CategoryCount = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) SetInstanceCount(v int32) *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields {
	s.InstanceCount = &v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields) SetRiskInstanceCount(v int32) *DescribeCloudProductFieldStatisticsResponseBodyGroupedFields {
	s.RiskInstanceCount = &v
	return s
}

type DescribeCloudProductFieldStatisticsResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCloudProductFieldStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCloudProductFieldStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudProductFieldStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudProductFieldStatisticsResponse) SetHeaders(v map[string]*string) *DescribeCloudProductFieldStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudProductFieldStatisticsResponse) SetBody(v *DescribeCloudProductFieldStatisticsResponseBody) *DescribeCloudProductFieldStatisticsResponse {
	s.Body = v
	return s
}

type DescribeConcernNecessityRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeConcernNecessityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConcernNecessityRequest) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityRequest) SetSourceIp(v string) *DescribeConcernNecessityRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeConcernNecessityRequest) SetLang(v string) *DescribeConcernNecessityRequest {
	s.Lang = &v
	return s
}

type DescribeConcernNecessityResponseBody struct {
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConcernNecessity []*string `json:"ConcernNecessity,omitempty" xml:"ConcernNecessity,omitempty" type:"Repeated"`
}

func (s DescribeConcernNecessityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConcernNecessityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityResponseBody) SetRequestId(v string) *DescribeConcernNecessityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConcernNecessityResponseBody) SetConcernNecessity(v []*string) *DescribeConcernNecessityResponseBody {
	s.ConcernNecessity = v
	return s
}

type DescribeConcernNecessityResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConcernNecessityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConcernNecessityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConcernNecessityResponse) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityResponse) SetHeaders(v map[string]*string) *DescribeConcernNecessityResponse {
	s.Headers = v
	return s
}

func (s *DescribeConcernNecessityResponse) SetBody(v *DescribeConcernNecessityResponseBody) *DescribeConcernNecessityResponse {
	s.Body = v
	return s
}

type DescribeContainerStatisticsRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeContainerStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsRequest) SetSourceIp(v string) *DescribeContainerStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeContainerStatisticsRequest) SetClusterId(v string) *DescribeContainerStatisticsRequest {
	s.ClusterId = &v
	return s
}

type DescribeContainerStatisticsResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeContainerStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeContainerStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsResponseBody) SetRequestId(v string) *DescribeContainerStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBody) SetData(v *DescribeContainerStatisticsResponseBodyData) *DescribeContainerStatisticsResponseBody {
	s.Data = v
	return s
}

type DescribeContainerStatisticsResponseBodyData struct {
	TotalNode            *int32 `json:"TotalNode,omitempty" xml:"TotalNode,omitempty"`
	RemindAlarmCount     *int32 `json:"RemindAlarmCount,omitempty" xml:"RemindAlarmCount,omitempty"`
	TotalAlarmCount      *int32 `json:"TotalAlarmCount,omitempty" xml:"TotalAlarmCount,omitempty"`
	SuspiciousAlarmCount *int32 `json:"SuspiciousAlarmCount,omitempty" xml:"SuspiciousAlarmCount,omitempty"`
	SeriousAlarmCount    *int32 `json:"SeriousAlarmCount,omitempty" xml:"SeriousAlarmCount,omitempty"`
	HasRiskNode          *int32 `json:"hasRiskNode,omitempty" xml:"hasRiskNode,omitempty"`
}

func (s DescribeContainerStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsResponseBodyData) SetTotalNode(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.TotalNode = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetRemindAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.RemindAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetTotalAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.TotalAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetSuspiciousAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.SuspiciousAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetSeriousAlarmCount(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.SeriousAlarmCount = &v
	return s
}

func (s *DescribeContainerStatisticsResponseBodyData) SetHasRiskNode(v int32) *DescribeContainerStatisticsResponseBodyData {
	s.HasRiskNode = &v
	return s
}

type DescribeContainerStatisticsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsResponse) SetHeaders(v map[string]*string) *DescribeContainerStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerStatisticsResponse) SetBody(v *DescribeContainerStatisticsResponseBody) *DescribeContainerStatisticsResponse {
	s.Body = v
	return s
}

type DescribeCriteriaRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCriteriaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaRequest) SetSourceIp(v string) *DescribeCriteriaRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCriteriaRequest) SetMachineTypes(v string) *DescribeCriteriaRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeCriteriaRequest) SetValue(v string) *DescribeCriteriaRequest {
	s.Value = &v
	return s
}

type DescribeCriteriaResponseBody struct {
	CriteriaList []*DescribeCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCriteriaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaResponseBody) SetCriteriaList(v []*DescribeCriteriaResponseBodyCriteriaList) *DescribeCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeCriteriaResponseBody) SetRequestId(v string) *DescribeCriteriaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCriteriaResponseBodyCriteriaList struct {
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCriteriaResponseBodyCriteriaList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

type DescribeCriteriaResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCriteriaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaResponse) SetHeaders(v map[string]*string) *DescribeCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeCriteriaResponse) SetBody(v *DescribeCriteriaResponseBody) *DescribeCriteriaResponse {
	s.Body = v
	return s
}

type DescribeDialogMessagesRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDialogMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogMessagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDialogMessagesRequest) SetSourceIp(v string) *DescribeDialogMessagesRequest {
	s.SourceIp = &v
	return s
}

type DescribeDialogMessagesResponseBody struct {
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DialogList []*DescribeDialogMessagesResponseBodyDialogList `json:"DialogList,omitempty" xml:"DialogList,omitempty" type:"Repeated"`
}

func (s DescribeDialogMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDialogMessagesResponseBody) SetTotalCount(v int32) *DescribeDialogMessagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDialogMessagesResponseBody) SetRequestId(v string) *DescribeDialogMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDialogMessagesResponseBody) SetDialogList(v []*DescribeDialogMessagesResponseBodyDialogList) *DescribeDialogMessagesResponseBody {
	s.DialogList = v
	return s
}

type DescribeDialogMessagesResponseBodyDialogList struct {
	DialogKey *string `json:"DialogKey,omitempty" xml:"DialogKey,omitempty"`
	Params    *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ID        *int64  `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s DescribeDialogMessagesResponseBodyDialogList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogMessagesResponseBodyDialogList) GoString() string {
	return s.String()
}

func (s *DescribeDialogMessagesResponseBodyDialogList) SetDialogKey(v string) *DescribeDialogMessagesResponseBodyDialogList {
	s.DialogKey = &v
	return s
}

func (s *DescribeDialogMessagesResponseBodyDialogList) SetParams(v string) *DescribeDialogMessagesResponseBodyDialogList {
	s.Params = &v
	return s
}

func (s *DescribeDialogMessagesResponseBodyDialogList) SetID(v int64) *DescribeDialogMessagesResponseBodyDialogList {
	s.ID = &v
	return s
}

type DescribeDialogMessagesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDialogMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDialogMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDialogMessagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDialogMessagesResponse) SetHeaders(v map[string]*string) *DescribeDialogMessagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDialogMessagesResponse) SetBody(v *DescribeDialogMessagesResponseBody) *DescribeDialogMessagesResponse {
	s.Body = v
	return s
}

type DescribeDingTalkRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RuleActionName *string `json:"RuleActionName,omitempty" xml:"RuleActionName,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeDingTalkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDingTalkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkRequest) SetSourceIp(v string) *DescribeDingTalkRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDingTalkRequest) SetRuleActionName(v string) *DescribeDingTalkRequest {
	s.RuleActionName = &v
	return s
}

func (s *DescribeDingTalkRequest) SetPageSize(v int32) *DescribeDingTalkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDingTalkRequest) SetCurrentPage(v int32) *DescribeDingTalkRequest {
	s.CurrentPage = &v
	return s
}

type DescribeDingTalkResponseBody struct {
	PageInfo   *DescribeDingTalkResponseBodyPageInfo     `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ActionList []*DescribeDingTalkResponseBodyActionList `json:"ActionList,omitempty" xml:"ActionList,omitempty" type:"Repeated"`
}

func (s DescribeDingTalkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDingTalkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponseBody) SetPageInfo(v *DescribeDingTalkResponseBodyPageInfo) *DescribeDingTalkResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeDingTalkResponseBody) SetRequestId(v string) *DescribeDingTalkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDingTalkResponseBody) SetActionList(v []*DescribeDingTalkResponseBodyActionList) *DescribeDingTalkResponseBody {
	s.ActionList = v
	return s
}

type DescribeDingTalkResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDingTalkResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDingTalkResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeDingTalkResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDingTalkResponseBodyPageInfo) SetPageSize(v int32) *DescribeDingTalkResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDingTalkResponseBodyPageInfo) SetTotalCount(v int32) *DescribeDingTalkResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeDingTalkResponseBodyActionList struct {
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ConfigList   *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ActionName   *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
	AliUid       *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DingTalkLang *string `json:"DingTalkLang,omitempty" xml:"DingTalkLang,omitempty"`
	IntervalTime *int32  `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	GmtModified  *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupIdList  *string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeDingTalkResponseBodyActionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDingTalkResponseBodyActionList) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponseBodyActionList) SetStatus(v int32) *DescribeDingTalkResponseBodyActionList {
	s.Status = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetConfigList(v string) *DescribeDingTalkResponseBodyActionList {
	s.ConfigList = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetGmtCreate(v int64) *DescribeDingTalkResponseBodyActionList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetActionName(v string) *DescribeDingTalkResponseBodyActionList {
	s.ActionName = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetUrl(v string) *DescribeDingTalkResponseBodyActionList {
	s.Url = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetAliUid(v int64) *DescribeDingTalkResponseBodyActionList {
	s.AliUid = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetDingTalkLang(v string) *DescribeDingTalkResponseBodyActionList {
	s.DingTalkLang = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetIntervalTime(v int32) *DescribeDingTalkResponseBodyActionList {
	s.IntervalTime = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetGmtModified(v int64) *DescribeDingTalkResponseBodyActionList {
	s.GmtModified = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetGroupIdList(v string) *DescribeDingTalkResponseBodyActionList {
	s.GroupIdList = &v
	return s
}

func (s *DescribeDingTalkResponseBodyActionList) SetId(v int32) *DescribeDingTalkResponseBodyActionList {
	s.Id = &v
	return s
}

type DescribeDingTalkResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDingTalkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDingTalkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDingTalkResponse) GoString() string {
	return s.String()
}

func (s *DescribeDingTalkResponse) SetHeaders(v map[string]*string) *DescribeDingTalkResponse {
	s.Headers = v
	return s
}

func (s *DescribeDingTalkResponse) SetBody(v *DescribeDingTalkResponseBody) *DescribeDingTalkResponse {
	s.Body = v
	return s
}

type DescribeDomainCountRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCountRequest) SetSourceIp(v string) *DescribeDomainCountRequest {
	s.SourceIp = &v
	return s
}

type DescribeDomainCountResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalDomainsCount *int32  `json:"TotalDomainsCount,omitempty" xml:"TotalDomainsCount,omitempty"`
	RootDomainsCount  *int32  `json:"RootDomainsCount,omitempty" xml:"RootDomainsCount,omitempty"`
}

func (s DescribeDomainCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCountResponseBody) SetRequestId(v string) *DescribeDomainCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCountResponseBody) SetTotalDomainsCount(v int32) *DescribeDomainCountResponseBody {
	s.TotalDomainsCount = &v
	return s
}

func (s *DescribeDomainCountResponseBody) SetRootDomainsCount(v int32) *DescribeDomainCountResponseBody {
	s.RootDomainsCount = &v
	return s
}

type DescribeDomainCountResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCountResponse) SetHeaders(v map[string]*string) *DescribeDomainCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCountResponse) SetBody(v *DescribeDomainCountResponseBody) *DescribeDomainCountResponse {
	s.Body = v
	return s
}

type DescribeDomainDetailRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailRequest) SetSourceIp(v string) *DescribeDomainDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetDomainName(v string) *DescribeDomainDetailRequest {
	s.DomainName = &v
	return s
}

type DescribeDomainDetailResponseBody struct {
	DomainDetailItems []*DescribeDomainDetailResponseBodyDomainDetailItems `json:"DomainDetailItems,omitempty" xml:"DomainDetailItems,omitempty" type:"Repeated"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RootDomain        *string                                              `json:"RootDomain,omitempty" xml:"RootDomain,omitempty"`
	Domain            *string                                              `json:"Domain,omitempty" xml:"Domain,omitempty"`
	VulCount          *int32                                               `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	AlarmCount        *int32                                               `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
}

func (s DescribeDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBody) SetDomainDetailItems(v []*DescribeDomainDetailResponseBodyDomainDetailItems) *DescribeDomainDetailResponseBody {
	s.DomainDetailItems = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRequestId(v string) *DescribeDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRootDomain(v string) *DescribeDomainDetailResponseBody {
	s.RootDomain = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetDomain(v string) *DescribeDomainDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetVulCount(v int32) *DescribeDomainDetailResponseBody {
	s.VulCount = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetAlarmCount(v int32) *DescribeDomainDetailResponseBody {
	s.AlarmCount = &v
	return s
}

type DescribeDomainDetailResponseBodyDomainDetailItems struct {
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntranetIp   *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	AssetType    *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
}

func (s DescribeDomainDetailResponseBodyDomainDetailItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyDomainDetailItems) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetUuid(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.Uuid = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetInternetIp(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.InternetIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetInstanceName(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetInstanceId(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetIntranetIp(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.IntranetIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyDomainDetailItems) SetAssetType(v string) *DescribeDomainDetailResponseBodyDomainDetailItems {
	s.AssetType = &v
	return s
}

type DescribeDomainDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainDetailResponse) SetBody(v *DescribeDomainDetailResponseBody) *DescribeDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeDomainListRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	FuzzyDomain *string `json:"FuzzyDomain,omitempty" xml:"FuzzyDomain,omitempty"`
	DomainType  *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainListRequest) SetSourceIp(v string) *DescribeDomainListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainListRequest) SetFuzzyDomain(v string) *DescribeDomainListRequest {
	s.FuzzyDomain = &v
	return s
}

func (s *DescribeDomainListRequest) SetDomainType(v string) *DescribeDomainListRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainListRequest) SetPageSize(v int32) *DescribeDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainListRequest) SetCurrentPage(v int32) *DescribeDomainListRequest {
	s.CurrentPage = &v
	return s
}

type DescribeDomainListResponseBody struct {
	PageInfo               *DescribeDomainListResponseBodyPageInfo                 `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId              *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DomainListResponseList []*DescribeDomainListResponseBodyDomainListResponseList `json:"DomainListResponseList,omitempty" xml:"DomainListResponseList,omitempty" type:"Repeated"`
}

func (s DescribeDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponseBody) SetPageInfo(v *DescribeDomainListResponseBodyPageInfo) *DescribeDomainListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeDomainListResponseBody) SetRequestId(v string) *DescribeDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainListResponseBody) SetDomainListResponseList(v []*DescribeDomainListResponseBodyDomainListResponseList) *DescribeDomainListResponseBody {
	s.DomainListResponseList = v
	return s
}

type DescribeDomainListResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDomainListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDomainListResponseBodyPageInfo) SetPageSize(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainListResponseBodyPageInfo) SetCount(v int32) *DescribeDomainListResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeDomainListResponseBodyDomainListResponseList struct {
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s DescribeDomainListResponseBodyDomainListResponseList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListResponseBodyDomainListResponseList) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponseBodyDomainListResponseList) SetDomain(v string) *DescribeDomainListResponseBodyDomainListResponseList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainListResponseBodyDomainListResponseList) SetIpList(v string) *DescribeDomainListResponseBodyDomainListResponseList {
	s.IpList = &v
	return s
}

type DescribeDomainListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainListResponse) SetHeaders(v map[string]*string) *DescribeDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainListResponse) SetBody(v *DescribeDomainListResponseBody) *DescribeDomainListResponse {
	s.Body = v
	return s
}

type DescribeEmgVulGroupRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeEmgVulGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmgVulGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulGroupRequest) SetSourceIp(v string) *DescribeEmgVulGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeEmgVulGroupRequest) SetLang(v string) *DescribeEmgVulGroupRequest {
	s.Lang = &v
	return s
}

type DescribeEmgVulGroupResponseBody struct {
	TotalCount      *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EmgVulGroupList []*DescribeEmgVulGroupResponseBodyEmgVulGroupList `json:"EmgVulGroupList,omitempty" xml:"EmgVulGroupList,omitempty" type:"Repeated"`
}

func (s DescribeEmgVulGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmgVulGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulGroupResponseBody) SetTotalCount(v int32) *DescribeEmgVulGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEmgVulGroupResponseBody) SetRequestId(v string) *DescribeEmgVulGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEmgVulGroupResponseBody) SetEmgVulGroupList(v []*DescribeEmgVulGroupResponseBodyEmgVulGroupList) *DescribeEmgVulGroupResponseBody {
	s.EmgVulGroupList = v
	return s
}

type DescribeEmgVulGroupResponseBodyEmgVulGroupList struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtPublish   *int64  `json:"GmtPublish,omitempty" xml:"GmtPublish,omitempty"`
	PendingCount *int32  `json:"PendingCount,omitempty" xml:"PendingCount,omitempty"`
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeEmgVulGroupResponseBodyEmgVulGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmgVulGroupResponseBodyEmgVulGroupList) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulGroupResponseBodyEmgVulGroupList) SetType(v string) *DescribeEmgVulGroupResponseBodyEmgVulGroupList {
	s.Type = &v
	return s
}

func (s *DescribeEmgVulGroupResponseBodyEmgVulGroupList) SetDescription(v string) *DescribeEmgVulGroupResponseBodyEmgVulGroupList {
	s.Description = &v
	return s
}

func (s *DescribeEmgVulGroupResponseBodyEmgVulGroupList) SetGmtPublish(v int64) *DescribeEmgVulGroupResponseBodyEmgVulGroupList {
	s.GmtPublish = &v
	return s
}

func (s *DescribeEmgVulGroupResponseBodyEmgVulGroupList) SetPendingCount(v int32) *DescribeEmgVulGroupResponseBodyEmgVulGroupList {
	s.PendingCount = &v
	return s
}

func (s *DescribeEmgVulGroupResponseBodyEmgVulGroupList) SetAliasName(v string) *DescribeEmgVulGroupResponseBodyEmgVulGroupList {
	s.AliasName = &v
	return s
}

func (s *DescribeEmgVulGroupResponseBodyEmgVulGroupList) SetName(v string) *DescribeEmgVulGroupResponseBodyEmgVulGroupList {
	s.Name = &v
	return s
}

type DescribeEmgVulGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEmgVulGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEmgVulGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEmgVulGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulGroupResponse) SetHeaders(v map[string]*string) *DescribeEmgVulGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmgVulGroupResponse) SetBody(v *DescribeEmgVulGroupResponseBody) *DescribeEmgVulGroupResponse {
	s.Body = v
	return s
}

type DescribeExportInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ExportId *int64  `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
}

func (s DescribeExportInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeExportInfoRequest) SetSourceIp(v string) *DescribeExportInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeExportInfoRequest) SetExportId(v int64) *DescribeExportInfoRequest {
	s.ExportId = &v
	return s
}

type DescribeExportInfoResponseBody struct {
	Progress     *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	TotalCount   *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	CurrentCount *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Link         *string `json:"Link,omitempty" xml:"Link,omitempty"`
}

func (s DescribeExportInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExportInfoResponseBody) SetProgress(v int32) *DescribeExportInfoResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetTotalCount(v int32) *DescribeExportInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetRequestId(v string) *DescribeExportInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetMessage(v string) *DescribeExportInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetFileName(v string) *DescribeExportInfoResponseBody {
	s.FileName = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetExportStatus(v string) *DescribeExportInfoResponseBody {
	s.ExportStatus = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetCurrentCount(v int32) *DescribeExportInfoResponseBody {
	s.CurrentCount = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetId(v int64) *DescribeExportInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeExportInfoResponseBody) SetLink(v string) *DescribeExportInfoResponseBody {
	s.Link = &v
	return s
}

type DescribeExportInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExportInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExportInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExportInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportInfoResponse) SetHeaders(v map[string]*string) *DescribeExportInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportInfoResponse) SetBody(v *DescribeExportInfoResponseBody) *DescribeExportInfoResponse {
	s.Body = v
	return s
}

type DescribeExposedInstanceCriteriaRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExposedInstanceCriteriaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaRequest) SetSourceIp(v string) *DescribeExposedInstanceCriteriaRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaRequest) SetValue(v string) *DescribeExposedInstanceCriteriaRequest {
	s.Value = &v
	return s
}

type DescribeExposedInstanceCriteriaResponseBody struct {
	CriteriaList []*DescribeExposedInstanceCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExposedInstanceCriteriaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaResponseBody) SetCriteriaList(v []*DescribeExposedInstanceCriteriaResponseBodyCriteriaList) *DescribeExposedInstanceCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBody) SetRequestId(v string) *DescribeExposedInstanceCriteriaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeExposedInstanceCriteriaResponseBodyCriteriaList struct {
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeExposedInstanceCriteriaResponseBodyCriteriaList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeExposedInstanceCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeExposedInstanceCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeExposedInstanceCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

type DescribeExposedInstanceCriteriaResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExposedInstanceCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExposedInstanceCriteriaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceCriteriaResponse) SetHeaders(v map[string]*string) *DescribeExposedInstanceCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedInstanceCriteriaResponse) SetBody(v *DescribeExposedInstanceCriteriaResponseBody) *DescribeExposedInstanceCriteriaResponse {
	s.Body = v
	return s
}

type DescribeExposedInstanceDetailRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeExposedInstanceDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailRequest) SetSourceIp(v string) *DescribeExposedInstanceDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeExposedInstanceDetailRequest) SetUuid(v string) *DescribeExposedInstanceDetailRequest {
	s.Uuid = &v
	return s
}

type DescribeExposedInstanceDetailResponseBody struct {
	RequestId     *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExposedChains []*DescribeExposedInstanceDetailResponseBodyExposedChains `json:"ExposedChains,omitempty" xml:"ExposedChains,omitempty" type:"Repeated"`
}

func (s DescribeExposedInstanceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBody) SetRequestId(v string) *DescribeExposedInstanceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBody) SetExposedChains(v []*DescribeExposedInstanceDetailResponseBodyExposedChains) *DescribeExposedInstanceDetailResponseBody {
	s.ExposedChains = v
	return s
}

type DescribeExposedInstanceDetailResponseBodyExposedChains struct {
	ExposureIp        *string                                                              `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	InternetIp        *string                                                              `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	NntfVulCount      *int32                                                               `json:"NntfVulCount,omitempty" xml:"NntfVulCount,omitempty"`
	AllVulList        []*DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList  `json:"AllVulList,omitempty" xml:"AllVulList,omitempty" type:"Repeated"`
	InstanceId        *string                                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ExposureType      *string                                                              `json:"ExposureType,omitempty" xml:"ExposureType,omitempty"`
	IntranetIp        *string                                                              `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	ExposureTypeId    *string                                                              `json:"ExposureTypeId,omitempty" xml:"ExposureTypeId,omitempty"`
	RegionId          *string                                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AsapVulCount      *int32                                                               `json:"AsapVulCount,omitempty" xml:"AsapVulCount,omitempty"`
	Uuid              *string                                                              `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	RealVulList       []*DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList `json:"RealVulList,omitempty" xml:"RealVulList,omitempty" type:"Repeated"`
	ExposurePort      *string                                                              `json:"ExposurePort,omitempty" xml:"ExposurePort,omitempty"`
	InstanceName      *string                                                              `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	LaterVulCount     *int32                                                               `json:"LaterVulCount,omitempty" xml:"LaterVulCount,omitempty"`
	ExposureComponent *string                                                              `json:"ExposureComponent,omitempty" xml:"ExposureComponent,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChains) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChains) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureIp(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetInternetIp(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.InternetIp = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetNntfVulCount(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.NntfVulCount = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetAllVulList(v []*DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.AllVulList = v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetInstanceId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.InstanceId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureType(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureType = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetIntranetIp(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.IntranetIp = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureTypeId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureTypeId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetRegionId(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.RegionId = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetAsapVulCount(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.AsapVulCount = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetUuid(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetRealVulList(v []*DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.RealVulList = v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposurePort(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposurePort = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetInstanceName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.InstanceName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetLaterVulCount(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.LaterVulCount = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChains) SetExposureComponent(v string) *DescribeExposedInstanceDetailResponseBodyExposedChains {
	s.ExposureComponent = &v
	return s
}

type DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList struct {
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Uuid      *int32  `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetType(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Type = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetNecessity(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Necessity = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetUuid(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetAliasName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.AliasName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList) SetName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsAllVulList {
	s.Name = &v
	return s
}

type DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList struct {
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Necessity *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Uuid      *int32  `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetType(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Type = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetNecessity(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Necessity = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetUuid(v int32) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetAliasName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.AliasName = &v
	return s
}

func (s *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList) SetName(v string) *DescribeExposedInstanceDetailResponseBodyExposedChainsRealVulList {
	s.Name = &v
	return s
}

type DescribeExposedInstanceDetailResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExposedInstanceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExposedInstanceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceDetailResponse) SetHeaders(v map[string]*string) *DescribeExposedInstanceDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedInstanceDetailResponse) SetBody(v *DescribeExposedInstanceDetailResponseBody) *DescribeExposedInstanceDetailResponse {
	s.Body = v
	return s
}

type DescribeExposedInstanceListRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	From              *string `json:"From,omitempty" xml:"From,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	GroupId           *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	VulStatus         *bool   `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
	ExposureComponent *string `json:"ExposureComponent,omitempty" xml:"ExposureComponent,omitempty"`
	ExposureType      *string `json:"ExposureType,omitempty" xml:"ExposureType,omitempty"`
	ExposurePort      *string `json:"ExposurePort,omitempty" xml:"ExposurePort,omitempty"`
	ExposureIp        *string `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DescribeExposedInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListRequest) SetSourceIp(v string) *DescribeExposedInstanceListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetFrom(v string) *DescribeExposedInstanceListRequest {
	s.From = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetPageSize(v int32) *DescribeExposedInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetCurrentPage(v int32) *DescribeExposedInstanceListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetUuid(v string) *DescribeExposedInstanceListRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetGroupId(v int64) *DescribeExposedInstanceListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetVulStatus(v bool) *DescribeExposedInstanceListRequest {
	s.VulStatus = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposureComponent(v string) *DescribeExposedInstanceListRequest {
	s.ExposureComponent = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposureType(v string) *DescribeExposedInstanceListRequest {
	s.ExposureType = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposurePort(v string) *DescribeExposedInstanceListRequest {
	s.ExposurePort = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposureIp(v string) *DescribeExposedInstanceListRequest {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetInstanceId(v string) *DescribeExposedInstanceListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetInstanceName(v string) *DescribeExposedInstanceListRequest {
	s.InstanceName = &v
	return s
}

type DescribeExposedInstanceListResponseBody struct {
	PageInfo         *DescribeExposedInstanceListResponseBodyPageInfo           `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId        *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExposedInstances []*DescribeExposedInstanceListResponseBodyExposedInstances `json:"ExposedInstances,omitempty" xml:"ExposedInstances,omitempty" type:"Repeated"`
	Success          *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExposedInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponseBody) SetPageInfo(v *DescribeExposedInstanceListResponseBodyPageInfo) *DescribeExposedInstanceListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeExposedInstanceListResponseBody) SetRequestId(v string) *DescribeExposedInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBody) SetExposedInstances(v []*DescribeExposedInstanceListResponseBodyExposedInstances) *DescribeExposedInstanceListResponseBody {
	s.ExposedInstances = v
	return s
}

func (s *DescribeExposedInstanceListResponseBody) SetSuccess(v bool) *DescribeExposedInstanceListResponseBody {
	s.Success = &v
	return s
}

type DescribeExposedInstanceListResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeExposedInstanceListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetPageSize(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyPageInfo) SetCount(v int32) *DescribeExposedInstanceListResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeExposedInstanceListResponseBodyExposedInstances struct {
	TotalVulCount     *int32  `json:"TotalVulCount,omitempty" xml:"TotalVulCount,omitempty"`
	ExposureIp        *string `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	InternetIp        *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	NntfVulCount      *int32  `json:"NntfVulCount,omitempty" xml:"NntfVulCount,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ExposureType      *string `json:"ExposureType,omitempty" xml:"ExposureType,omitempty"`
	IntranetIp        *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	ExposureTypeId    *string `json:"ExposureTypeId,omitempty" xml:"ExposureTypeId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AsapVulCount      *int32  `json:"AsapVulCount,omitempty" xml:"AsapVulCount,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	ExposurePort      *string `json:"ExposurePort,omitempty" xml:"ExposurePort,omitempty"`
	GroupName         *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId           *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceName      *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ExposureComponent *string `json:"ExposureComponent,omitempty" xml:"ExposureComponent,omitempty"`
	LaterVulCount     *int32  `json:"LaterVulCount,omitempty" xml:"LaterVulCount,omitempty"`
}

func (s DescribeExposedInstanceListResponseBodyExposedInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceListResponseBodyExposedInstances) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetTotalVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.TotalVulCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureIp(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetInternetIp(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetNntfVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.NntfVulCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetInstanceId(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureType(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureType = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetIntranetIp(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureTypeId(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureTypeId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetRegionId(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetAsapVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.AsapVulCount = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetUuid(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.Uuid = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposurePort(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposurePort = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetGroupName(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.GroupName = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetGroupId(v int64) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.GroupId = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetInstanceName(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetExposureComponent(v string) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.ExposureComponent = &v
	return s
}

func (s *DescribeExposedInstanceListResponseBodyExposedInstances) SetLaterVulCount(v int32) *DescribeExposedInstanceListResponseBodyExposedInstances {
	s.LaterVulCount = &v
	return s
}

type DescribeExposedInstanceListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExposedInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExposedInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListResponse) SetHeaders(v map[string]*string) *DescribeExposedInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedInstanceListResponse) SetBody(v *DescribeExposedInstanceListResponseBody) *DescribeExposedInstanceListResponse {
	s.Body = v
	return s
}

type DescribeExposedStatisticsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeExposedStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsRequest) SetSourceIp(v string) *DescribeExposedStatisticsRequest {
	s.SourceIp = &v
	return s
}

type DescribeExposedStatisticsResponseBody struct {
	ExposedPortCount      *int32  `json:"ExposedPortCount,omitempty" xml:"ExposedPortCount,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExposedLaterVulCount  *int32  `json:"ExposedLaterVulCount,omitempty" xml:"ExposedLaterVulCount,omitempty"`
	ExposedInstanceCount  *int32  `json:"ExposedInstanceCount,omitempty" xml:"ExposedInstanceCount,omitempty"`
	GatewayAssetCount     *int32  `json:"GatewayAssetCount,omitempty" xml:"GatewayAssetCount,omitempty"`
	ExposedComponentCount *int32  `json:"ExposedComponentCount,omitempty" xml:"ExposedComponentCount,omitempty"`
	ExposedNntfVulCount   *int32  `json:"ExposedNntfVulCount,omitempty" xml:"ExposedNntfVulCount,omitempty"`
	ExposedIpCount        *int32  `json:"ExposedIpCount,omitempty" xml:"ExposedIpCount,omitempty"`
	ExposedAsapVulCount   *int32  `json:"ExposedAsapVulCount,omitempty" xml:"ExposedAsapVulCount,omitempty"`
}

func (s DescribeExposedStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedPortCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedPortCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetRequestId(v string) *DescribeExposedStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedLaterVulCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedLaterVulCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedInstanceCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedInstanceCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetGatewayAssetCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.GatewayAssetCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedComponentCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedComponentCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedNntfVulCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedNntfVulCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedIpCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedIpCount = &v
	return s
}

func (s *DescribeExposedStatisticsResponseBody) SetExposedAsapVulCount(v int32) *DescribeExposedStatisticsResponseBody {
	s.ExposedAsapVulCount = &v
	return s
}

type DescribeExposedStatisticsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExposedStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExposedStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExposedStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeExposedStatisticsResponse) SetHeaders(v map[string]*string) *DescribeExposedStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeExposedStatisticsResponse) SetBody(v *DescribeExposedStatisticsResponseBody) *DescribeExposedStatisticsResponse {
	s.Body = v
	return s
}

type DescribeFieldStatisticsRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
}

func (s DescribeFieldStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsRequest) SetSourceIp(v string) *DescribeFieldStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeFieldStatisticsRequest) SetRegionId(v string) *DescribeFieldStatisticsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFieldStatisticsRequest) SetMachineTypes(v string) *DescribeFieldStatisticsRequest {
	s.MachineTypes = &v
	return s
}

type DescribeFieldStatisticsResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GroupedFields *DescribeFieldStatisticsResponseBodyGroupedFields `json:"GroupedFields,omitempty" xml:"GroupedFields,omitempty" type:"Struct"`
}

func (s DescribeFieldStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsResponseBody) SetRequestId(v string) *DescribeFieldStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBody) SetGroupedFields(v *DescribeFieldStatisticsResponseBodyGroupedFields) *DescribeFieldStatisticsResponseBody {
	s.GroupedFields = v
	return s
}

type DescribeFieldStatisticsResponseBodyGroupedFields struct {
	OfflineInstanceCount     *int32 `json:"OfflineInstanceCount,omitempty" xml:"OfflineInstanceCount,omitempty"`
	RegionCount              *int32 `json:"RegionCount,omitempty" xml:"RegionCount,omitempty"`
	NewInstanceCount         *int32 `json:"NewInstanceCount,omitempty" xml:"NewInstanceCount,omitempty"`
	GroupCount               *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	ExposedInstanceCount     *int32 `json:"ExposedInstanceCount,omitempty" xml:"ExposedInstanceCount,omitempty"`
	GeneralAssetCount        *int32 `json:"GeneralAssetCount,omitempty" xml:"GeneralAssetCount,omitempty"`
	UnprotectedInstanceCount *int32 `json:"UnprotectedInstanceCount,omitempty" xml:"UnprotectedInstanceCount,omitempty"`
	ImportantAssetCount      *int32 `json:"ImportantAssetCount,omitempty" xml:"ImportantAssetCount,omitempty"`
	TestAssetCount           *int32 `json:"TestAssetCount,omitempty" xml:"TestAssetCount,omitempty"`
	VpcCount                 *int32 `json:"VpcCount,omitempty" xml:"VpcCount,omitempty"`
	InstanceCount            *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	NotRunningStatusCount    *int32 `json:"NotRunningStatusCount,omitempty" xml:"NotRunningStatusCount,omitempty"`
	RiskInstanceCount        *int32 `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
}

func (s DescribeFieldStatisticsResponseBodyGroupedFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldStatisticsResponseBodyGroupedFields) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetOfflineInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.OfflineInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetRegionCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.RegionCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNewInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NewInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetGroupCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.GroupCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetExposedInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.ExposedInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetGeneralAssetCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.GeneralAssetCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetUnprotectedInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.UnprotectedInstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetImportantAssetCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.ImportantAssetCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetTestAssetCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.TestAssetCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetVpcCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.VpcCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.InstanceCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetNotRunningStatusCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.NotRunningStatusCount = &v
	return s
}

func (s *DescribeFieldStatisticsResponseBodyGroupedFields) SetRiskInstanceCount(v int32) *DescribeFieldStatisticsResponseBodyGroupedFields {
	s.RiskInstanceCount = &v
	return s
}

type DescribeFieldStatisticsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFieldStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFieldStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFieldStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldStatisticsResponse) SetHeaders(v map[string]*string) *DescribeFieldStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldStatisticsResponse) SetBody(v *DescribeFieldStatisticsResponseBody) *DescribeFieldStatisticsResponse {
	s.Body = v
	return s
}

type DescribeGraph4InvestigationOnlineRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	VertexId    *string `json:"VertexId,omitempty" xml:"VertexId,omitempty"`
	AnomalyUuid *string `json:"AnomalyUuid,omitempty" xml:"AnomalyUuid,omitempty"`
	AnomalyId   *string `json:"AnomalyId,omitempty" xml:"AnomalyId,omitempty"`
	PathLength  *int32  `json:"PathLength,omitempty" xml:"PathLength,omitempty"`
	Direction   *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineRequest) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetSourceIp(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetLang(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetNamespace(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetVertexId(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.VertexId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetAnomalyUuid(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.AnomalyUuid = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetAnomalyId(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.AnomalyId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetPathLength(v int32) *DescribeGraph4InvestigationOnlineRequest {
	s.PathLength = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineRequest) SetDirection(v string) *DescribeGraph4InvestigationOnlineRequest {
	s.Direction = &v
	return s
}

type DescribeGraph4InvestigationOnlineResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeGraph4InvestigationOnlineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeGraph4InvestigationOnlineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponseBody) SetRequestId(v string) *DescribeGraph4InvestigationOnlineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBody) SetData(v *DescribeGraph4InvestigationOnlineResponseBodyData) *DescribeGraph4InvestigationOnlineResponseBody {
	s.Data = v
	return s
}

type DescribeGraph4InvestigationOnlineResponseBodyData struct {
	VertexList       []*DescribeGraph4InvestigationOnlineResponseBodyDataVertexList       `json:"VertexList,omitempty" xml:"VertexList,omitempty" type:"Repeated"`
	EntityTypeList   []*DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList   `json:"EntityTypeList,omitempty" xml:"EntityTypeList,omitempty" type:"Repeated"`
	RelationTypeList []*DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList `json:"RelationTypeList,omitempty" xml:"RelationTypeList,omitempty" type:"Repeated"`
	EdgeList         []*DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList         `json:"EdgeList,omitempty" xml:"EdgeList,omitempty" type:"Repeated"`
}

func (s DescribeGraph4InvestigationOnlineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyData) SetVertexList(v []*DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) *DescribeGraph4InvestigationOnlineResponseBodyData {
	s.VertexList = v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyData) SetEntityTypeList(v []*DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) *DescribeGraph4InvestigationOnlineResponseBodyData {
	s.EntityTypeList = v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyData) SetRelationTypeList(v []*DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList) *DescribeGraph4InvestigationOnlineResponseBodyData {
	s.RelationTypeList = v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyData) SetEdgeList(v []*DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) *DescribeGraph4InvestigationOnlineResponseBodyData {
	s.EdgeList = v
	return s
}

type DescribeGraph4InvestigationOnlineResponseBodyDataVertexList struct {
	Type         *string                                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuid         *string                                                                    `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Time         *string                                                                    `json:"Time,omitempty" xml:"Time,omitempty"`
	Aliuid       *string                                                                    `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	NeighborList []*DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList `json:"NeighborList,omitempty" xml:"NeighborList,omitempty" type:"Repeated"`
	Position     *string                                                                    `json:"Position,omitempty" xml:"Position,omitempty"`
	PositionId   *string                                                                    `json:"PositionId,omitempty" xml:"PositionId,omitempty"`
	Name         *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *string                                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	Properties   *string                                                                    `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetType(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Type = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetUuid(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Uuid = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetTime(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Time = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetAliuid(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Aliuid = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetNeighborList(v []*DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.NeighborList = v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetPosition(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Position = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetPositionId(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.PositionId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetName(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Name = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetId(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Id = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList) SetProperties(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexList {
	s.Properties = &v
	return s
}

type DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList struct {
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	HasMore *bool   `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	Count   *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList) SetType(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList {
	s.Type = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList) SetHasMore(v bool) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList {
	s.HasMore = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList) SetCount(v int32) *DescribeGraph4InvestigationOnlineResponseBodyDataVertexListNeighborList {
	s.Count = &v
	return s
}

type DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList struct {
	DisplayColor *string `json:"DisplayColor,omitempty" xml:"DisplayColor,omitempty"`
	DisplayIcon  *string `json:"DisplayIcon,omitempty" xml:"DisplayIcon,omitempty"`
	DisplayOrder *string `json:"DisplayOrder,omitempty" xml:"DisplayOrder,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) SetDisplayColor(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList {
	s.DisplayColor = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) SetDisplayIcon(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList {
	s.DisplayIcon = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) SetDisplayOrder(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList {
	s.DisplayOrder = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) SetName(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList {
	s.Name = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList) SetId(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEntityTypeList {
	s.Id = &v
	return s
}

type DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList struct {
	ShowType *string `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Directed *int32  `json:"Directed,omitempty" xml:"Directed,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList) SetShowType(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList {
	s.ShowType = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList) SetName(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList {
	s.Name = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList) SetId(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList {
	s.Id = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList) SetDirected(v int32) *DescribeGraph4InvestigationOnlineResponseBodyDataRelationTypeList {
	s.Directed = &v
	return s
}

type DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList struct {
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Time      *string `json:"Time,omitempty" xml:"Time,omitempty"`
	EndId     *string `json:"EndId,omitempty" xml:"EndId,omitempty"`
	StartType *string `json:"StartType,omitempty" xml:"StartType,omitempty"`
	EndType   *string `json:"EndType,omitempty" xml:"EndType,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	StartId   *string `json:"StartId,omitempty" xml:"StartId,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetType(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.Type = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetTime(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.Time = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetEndId(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.EndId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetStartType(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.StartType = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetEndType(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.EndType = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetName(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.Name = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetStartId(v string) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.StartId = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList) SetId(v int32) *DescribeGraph4InvestigationOnlineResponseBodyDataEdgeList {
	s.Id = &v
	return s
}

type DescribeGraph4InvestigationOnlineResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGraph4InvestigationOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGraph4InvestigationOnlineResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponse) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponse) SetHeaders(v map[string]*string) *DescribeGraph4InvestigationOnlineResponse {
	s.Headers = v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponse) SetBody(v *DescribeGraph4InvestigationOnlineResponseBody) *DescribeGraph4InvestigationOnlineResponse {
	s.Body = v
	return s
}

type DescribeGroupedContainerInstancesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Criteria    *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	LogicalExp  *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	GroupField  *string `json:"GroupField,omitempty" xml:"GroupField,omitempty"`
	FieldValue  *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeGroupedContainerInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedContainerInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesRequest) SetSourceIp(v string) *DescribeGroupedContainerInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetCriteria(v string) *DescribeGroupedContainerInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetLogicalExp(v string) *DescribeGroupedContainerInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetGroupField(v string) *DescribeGroupedContainerInstancesRequest {
	s.GroupField = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetFieldValue(v string) *DescribeGroupedContainerInstancesRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetPageSize(v int32) *DescribeGroupedContainerInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedContainerInstancesRequest) SetCurrentPage(v int32) *DescribeGroupedContainerInstancesRequest {
	s.CurrentPage = &v
	return s
}

type DescribeGroupedContainerInstancesResponseBody struct {
	PageInfo                     *DescribeGroupedContainerInstancesResponseBodyPageInfo                       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	GroupedContainerInstanceList []*DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList `json:"GroupedContainerInstanceList,omitempty" xml:"GroupedContainerInstanceList,omitempty" type:"Repeated"`
	RequestId                    *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupedContainerInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponseBody) SetPageInfo(v *DescribeGroupedContainerInstancesResponseBodyPageInfo) *DescribeGroupedContainerInstancesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBody) SetGroupedContainerInstanceList(v []*DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) *DescribeGroupedContainerInstancesResponseBody {
	s.GroupedContainerInstanceList = v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBody) SetRequestId(v string) *DescribeGroupedContainerInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGroupedContainerInstancesResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeGroupedContainerInstancesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetPageSize(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyPageInfo) SetCount(v int32) *DescribeGroupedContainerInstancesResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList struct {
	HostIp            *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	RiskLevel         *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	RiskStatus        *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	Pod               *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CusterState       *string `json:"CusterState,omitempty" xml:"CusterState,omitempty"`
	Namespace         *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	AppName           *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InstanceCount     *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	ClusterType       *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ClusterName       *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	PodIp             *string `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	VulCount          *int32  `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
	AlarmCount        *int32  `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	RiskInstanceCount *int32  `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetHostIp(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.HostIp = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRiskLevel(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRiskStatus(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RiskStatus = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetPod(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.Pod = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetCreateTime(v int64) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.CreateTime = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetCusterState(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.CusterState = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetNamespace(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.Namespace = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetInstanceId(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRegionId(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetAppName(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.AppName = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetInstanceCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.InstanceCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetClusterType(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ClusterType = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetClusterName(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ClusterName = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetPodIp(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.PodIp = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetVulCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.VulCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetAlarmCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.AlarmCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetRiskInstanceCount(v int32) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.RiskInstanceCount = &v
	return s
}

func (s *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList) SetClusterId(v string) *DescribeGroupedContainerInstancesResponseBodyGroupedContainerInstanceList {
	s.ClusterId = &v
	return s
}

type DescribeGroupedContainerInstancesResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupedContainerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupedContainerInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedContainerInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedContainerInstancesResponse) SetHeaders(v map[string]*string) *DescribeGroupedContainerInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedContainerInstancesResponse) SetBody(v *DescribeGroupedContainerInstancesResponseBody) *DescribeGroupedContainerInstancesResponse {
	s.Body = v
	return s
}

type DescribeGroupedMaliciousFilesRequest struct {
	SourceIp           *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang               *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Levels             *string   `json:"Levels,omitempty" xml:"Levels,omitempty"`
	FuzzyMaliciousName *string   `json:"FuzzyMaliciousName,omitempty" xml:"FuzzyMaliciousName,omitempty"`
	CurrentPage        *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize           *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RepoRegionId       *string   `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	RepoInstanceId     *string   `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	RepoId             *string   `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepoName           *string   `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespace      *string   `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	ImageTag           *string   `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ImageDigest        *string   `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	ImageLayer         *string   `json:"ImageLayer,omitempty" xml:"ImageLayer,omitempty"`
	Uuids              []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s DescribeGroupedMaliciousFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesRequest) SetSourceIp(v string) *DescribeGroupedMaliciousFilesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetLang(v string) *DescribeGroupedMaliciousFilesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetLevels(v string) *DescribeGroupedMaliciousFilesRequest {
	s.Levels = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetFuzzyMaliciousName(v string) *DescribeGroupedMaliciousFilesRequest {
	s.FuzzyMaliciousName = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetCurrentPage(v int32) *DescribeGroupedMaliciousFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetPageSize(v string) *DescribeGroupedMaliciousFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoRegionId(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoInstanceId(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoId(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoName(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetRepoNamespace(v string) *DescribeGroupedMaliciousFilesRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetImageTag(v string) *DescribeGroupedMaliciousFilesRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetImageDigest(v string) *DescribeGroupedMaliciousFilesRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetImageLayer(v string) *DescribeGroupedMaliciousFilesRequest {
	s.ImageLayer = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesRequest) SetUuids(v []*string) *DescribeGroupedMaliciousFilesRequest {
	s.Uuids = v
	return s
}

type DescribeGroupedMaliciousFilesResponseBody struct {
	PageInfo                     *DescribeGroupedMaliciousFilesResponseBodyPageInfo                       `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	GroupedMaliciousFileResponse []*DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse `json:"GroupedMaliciousFileResponse,omitempty" xml:"GroupedMaliciousFileResponse,omitempty" type:"Repeated"`
	RequestId                    *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupedMaliciousFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponseBody) SetPageInfo(v *DescribeGroupedMaliciousFilesResponseBodyPageInfo) *DescribeGroupedMaliciousFilesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBody) SetGroupedMaliciousFileResponse(v []*DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) *DescribeGroupedMaliciousFilesResponseBody {
	s.GroupedMaliciousFileResponse = v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBody) SetRequestId(v string) *DescribeGroupedMaliciousFilesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGroupedMaliciousFilesResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeGroupedMaliciousFilesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetPageSize(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyPageInfo) SetCount(v int32) *DescribeGroupedMaliciousFilesResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse struct {
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	ImageCount          *int64  `json:"ImageCount,omitempty" xml:"ImageCount,omitempty"`
	LatestScanTimestamp *int64  `json:"LatestScanTimestamp,omitempty" xml:"LatestScanTimestamp,omitempty"`
	MaliciousName       *string `json:"MaliciousName,omitempty" xml:"MaliciousName,omitempty"`
	MaliciousMd5        *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	FirstScanTimestamp  *int64  `json:"FirstScanTimestamp,omitempty" xml:"FirstScanTimestamp,omitempty"`
	Level               *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetStatus(v int32) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.Status = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetImageCount(v int64) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.ImageCount = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetLatestScanTimestamp(v int64) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.LatestScanTimestamp = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetMaliciousName(v string) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.MaliciousName = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetMaliciousMd5(v string) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.MaliciousMd5 = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetFirstScanTimestamp(v int64) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.FirstScanTimestamp = &v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse) SetLevel(v string) *DescribeGroupedMaliciousFilesResponseBodyGroupedMaliciousFileResponse {
	s.Level = &v
	return s
}

type DescribeGroupedMaliciousFilesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupedMaliciousFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupedMaliciousFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedMaliciousFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedMaliciousFilesResponse) SetHeaders(v map[string]*string) *DescribeGroupedMaliciousFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedMaliciousFilesResponse) SetBody(v *DescribeGroupedMaliciousFilesResponseBody) *DescribeGroupedMaliciousFilesResponse {
	s.Body = v
	return s
}

type DescribeGroupedTagsRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
}

func (s DescribeGroupedTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsRequest) SetSourceIp(v string) *DescribeGroupedTagsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGroupedTagsRequest) SetMachineTypes(v string) *DescribeGroupedTagsRequest {
	s.MachineTypes = &v
	return s
}

type DescribeGroupedTagsResponseBody struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	GroupedFileds  []*DescribeGroupedTagsResponseBodyGroupedFileds `json:"GroupedFileds,omitempty" xml:"GroupedFileds,omitempty" type:"Repeated"`
	Count          *int32                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	Success        *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGroupedTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsResponseBody) SetRequestId(v string) *DescribeGroupedTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetHttpStatusCode(v int32) *DescribeGroupedTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetGroupedFileds(v []*DescribeGroupedTagsResponseBodyGroupedFileds) *DescribeGroupedTagsResponseBody {
	s.GroupedFileds = v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetCount(v int32) *DescribeGroupedTagsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeGroupedTagsResponseBody) SetSuccess(v bool) *DescribeGroupedTagsResponseBody {
	s.Success = &v
	return s
}

type DescribeGroupedTagsResponseBodyGroupedFileds struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	TagId *int32  `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DescribeGroupedTagsResponseBodyGroupedFileds) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedTagsResponseBodyGroupedFileds) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) SetName(v string) *DescribeGroupedTagsResponseBodyGroupedFileds {
	s.Name = &v
	return s
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) SetCount(v string) *DescribeGroupedTagsResponseBodyGroupedFileds {
	s.Count = &v
	return s
}

func (s *DescribeGroupedTagsResponseBodyGroupedFileds) SetTagId(v int32) *DescribeGroupedTagsResponseBodyGroupedFileds {
	s.TagId = &v
	return s
}

type DescribeGroupedTagsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupedTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupedTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedTagsResponse) SetHeaders(v map[string]*string) *DescribeGroupedTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedTagsResponse) SetBody(v *DescribeGroupedTagsResponseBody) *DescribeGroupedTagsResponse {
	s.Body = v
	return s
}

type DescribeGroupedVulRequest struct {
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuids               *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	AliasName           *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Necessity           *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Dealed              *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	CurrentPage         *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList          *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ContainerFieldName  *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MinScore            *int32  `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
}

func (s DescribeGroupedVulRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulRequest) SetSourceIp(v string) *DescribeGroupedVulRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetLang(v string) *DescribeGroupedVulRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetType(v string) *DescribeGroupedVulRequest {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetUuids(v string) *DescribeGroupedVulRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetAliasName(v string) *DescribeGroupedVulRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetNecessity(v string) *DescribeGroupedVulRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetDealed(v string) *DescribeGroupedVulRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetCurrentPage(v int32) *DescribeGroupedVulRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetPageSize(v int32) *DescribeGroupedVulRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetStatusList(v string) *DescribeGroupedVulRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetGroupId(v string) *DescribeGroupedVulRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetContainerFieldName(v string) *DescribeGroupedVulRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetContainerFieldValue(v string) *DescribeGroupedVulRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetTargetType(v string) *DescribeGroupedVulRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetClusterId(v string) *DescribeGroupedVulRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGroupedVulRequest) SetMinScore(v int32) *DescribeGroupedVulRequest {
	s.MinScore = &v
	return s
}

type DescribeGroupedVulResponseBody struct {
	TotalCount      *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize        *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	GroupedVulItems []*DescribeGroupedVulResponseBodyGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" type:"Repeated"`
}

func (s DescribeGroupedVulResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBody) SetTotalCount(v int32) *DescribeGroupedVulResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetRequestId(v string) *DescribeGroupedVulResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetPageSize(v int32) *DescribeGroupedVulResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetCurrentPage(v int32) *DescribeGroupedVulResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetGroupedVulItems(v []*DescribeGroupedVulResponseBodyGroupedVulItems) *DescribeGroupedVulResponseBody {
	s.GroupedVulItems = v
	return s
}

type DescribeGroupedVulResponseBodyGroupedVulItems struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	NntfCount    *int32  `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	HandledCount *int32  `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	GmtLast      *int64  `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	LaterCount   *int32  `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AsapCount    *int32  `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetType(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetNntfCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.NntfCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetHandledCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.HandledCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetGmtLast(v int64) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.GmtLast = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetTags(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Tags = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetLaterCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.LaterCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAliasName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAsapCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AsapCount = &v
	return s
}

type DescribeGroupedVulResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGroupedVulResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGroupedVulResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupedVulResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponse) SetHeaders(v map[string]*string) *DescribeGroupedVulResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedVulResponse) SetBody(v *DescribeGroupedVulResponseBody) *DescribeGroupedVulResponse {
	s.Body = v
	return s
}

type DescribeHoneyPotAuthRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeHoneyPotAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHoneyPotAuthRequest) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotAuthRequest) SetSourceIp(v string) *DescribeHoneyPotAuthRequest {
	s.SourceIp = &v
	return s
}

type DescribeHoneyPotAuthResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HoneyPotAuthCount *int64  `json:"HoneyPotAuthCount,omitempty" xml:"HoneyPotAuthCount,omitempty"`
	HoneyPotCount     *int32  `json:"HoneyPotCount,omitempty" xml:"HoneyPotCount,omitempty"`
}

func (s DescribeHoneyPotAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHoneyPotAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotAuthResponseBody) SetRequestId(v string) *DescribeHoneyPotAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHoneyPotAuthResponseBody) SetHoneyPotAuthCount(v int64) *DescribeHoneyPotAuthResponseBody {
	s.HoneyPotAuthCount = &v
	return s
}

func (s *DescribeHoneyPotAuthResponseBody) SetHoneyPotCount(v int32) *DescribeHoneyPotAuthResponseBody {
	s.HoneyPotCount = &v
	return s
}

type DescribeHoneyPotAuthResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHoneyPotAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHoneyPotAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHoneyPotAuthResponse) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotAuthResponse) SetHeaders(v map[string]*string) *DescribeHoneyPotAuthResponse {
	s.Headers = v
	return s
}

func (s *DescribeHoneyPotAuthResponse) SetBody(v *DescribeHoneyPotAuthResponseBody) *DescribeHoneyPotAuthResponse {
	s.Body = v
	return s
}

type DescribeHoneyPotSuspStatisticsRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	From              *string `json:"From,omitempty" xml:"From,omitempty"`
	StatisticsKeyType *string `json:"StatisticsKeyType,omitempty" xml:"StatisticsKeyType,omitempty"`
	StatisticsDays    *int32  `json:"StatisticsDays,omitempty" xml:"StatisticsDays,omitempty"`
}

func (s DescribeHoneyPotSuspStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetSourceIp(v string) *DescribeHoneyPotSuspStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetFrom(v string) *DescribeHoneyPotSuspStatisticsRequest {
	s.From = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetStatisticsKeyType(v string) *DescribeHoneyPotSuspStatisticsRequest {
	s.StatisticsKeyType = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsRequest) SetStatisticsDays(v int32) *DescribeHoneyPotSuspStatisticsRequest {
	s.StatisticsDays = &v
	return s
}

type DescribeHoneyPotSuspStatisticsResponseBody struct {
	RequestId                      *string                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuspHoneyPotStatisticsResponse []*DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse `json:"SuspHoneyPotStatisticsResponse,omitempty" xml:"SuspHoneyPotStatisticsResponse,omitempty" type:"Repeated"`
}

func (s DescribeHoneyPotSuspStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsResponseBody) SetRequestId(v string) *DescribeHoneyPotSuspStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBody) SetSuspHoneyPotStatisticsResponse(v []*DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) *DescribeHoneyPotSuspStatisticsResponseBody {
	s.SuspHoneyPotStatisticsResponse = v
	return s
}

type DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse struct {
	VpcName      *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetVpcName(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.VpcName = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetType(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.Type = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetVpcId(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.VpcId = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetInstanceName(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.InstanceName = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetInstanceId(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.InstanceId = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetName(v string) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.Name = &v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse) SetCount(v int32) *DescribeHoneyPotSuspStatisticsResponseBodySuspHoneyPotStatisticsResponse {
	s.Count = &v
	return s
}

type DescribeHoneyPotSuspStatisticsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHoneyPotSuspStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeHoneyPotSuspStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHoneyPotSuspStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotSuspStatisticsResponse) SetHeaders(v map[string]*string) *DescribeHoneyPotSuspStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHoneyPotSuspStatisticsResponse) SetBody(v *DescribeHoneyPotSuspStatisticsResponseBody) *DescribeHoneyPotSuspStatisticsResponse {
	s.Body = v
	return s
}

type DescribeImageGroupedVulListRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	CveId          *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	Uuids          *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AliasName      *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	PatchId        *int64  `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	LastTsStart    *int64  `json:"LastTsStart,omitempty" xml:"LastTsStart,omitempty"`
	LastTsEnd      *int64  `json:"LastTsEnd,omitempty" xml:"LastTsEnd,omitempty"`
	StatusList     *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	OrderBy        *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Direction      *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Necessity      *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Dealed         *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	CreateTsStart  *int64  `json:"CreateTsStart,omitempty" xml:"CreateTsStart,omitempty"`
	CreateTsEnd    *int64  `json:"CreateTsEnd,omitempty" xml:"CreateTsEnd,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Remark         *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SearchTags     *string `json:"SearchTags,omitempty" xml:"SearchTags,omitempty"`
	RepoRegionId   *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	RepoId         *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepoName       *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespace  *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	ImageTag       *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ImageDigest    *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	ImageLayer     *string `json:"ImageLayer,omitempty" xml:"ImageLayer,omitempty"`
}

func (s DescribeImageGroupedVulListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGroupedVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListRequest) SetSourceIp(v string) *DescribeImageGroupedVulListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetType(v string) *DescribeImageGroupedVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetGroupId(v string) *DescribeImageGroupedVulListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetCveId(v string) *DescribeImageGroupedVulListRequest {
	s.CveId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetUuids(v string) *DescribeImageGroupedVulListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetName(v string) *DescribeImageGroupedVulListRequest {
	s.Name = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetAliasName(v string) *DescribeImageGroupedVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetPatchId(v int64) *DescribeImageGroupedVulListRequest {
	s.PatchId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetLevel(v string) *DescribeImageGroupedVulListRequest {
	s.Level = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetLastTsStart(v int64) *DescribeImageGroupedVulListRequest {
	s.LastTsStart = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetLastTsEnd(v int64) *DescribeImageGroupedVulListRequest {
	s.LastTsEnd = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetStatusList(v string) *DescribeImageGroupedVulListRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetOrderBy(v string) *DescribeImageGroupedVulListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetDirection(v string) *DescribeImageGroupedVulListRequest {
	s.Direction = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetNecessity(v string) *DescribeImageGroupedVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetDealed(v string) *DescribeImageGroupedVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetCreateTsStart(v int64) *DescribeImageGroupedVulListRequest {
	s.CreateTsStart = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetCreateTsEnd(v int64) *DescribeImageGroupedVulListRequest {
	s.CreateTsEnd = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetCurrentPage(v int32) *DescribeImageGroupedVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetPageSize(v int32) *DescribeImageGroupedVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRemark(v string) *DescribeImageGroupedVulListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetSearchTags(v string) *DescribeImageGroupedVulListRequest {
	s.SearchTags = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoRegionId(v string) *DescribeImageGroupedVulListRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoInstanceId(v string) *DescribeImageGroupedVulListRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoId(v string) *DescribeImageGroupedVulListRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoName(v string) *DescribeImageGroupedVulListRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetRepoNamespace(v string) *DescribeImageGroupedVulListRequest {
	s.RepoNamespace = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetImageTag(v string) *DescribeImageGroupedVulListRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetImageDigest(v string) *DescribeImageGroupedVulListRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeImageGroupedVulListRequest) SetImageLayer(v string) *DescribeImageGroupedVulListRequest {
	s.ImageLayer = &v
	return s
}

type DescribeImageGroupedVulListResponseBody struct {
	TotalCount      *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize        *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32                                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	GroupedVulItems []*DescribeImageGroupedVulListResponseBodyGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" type:"Repeated"`
}

func (s DescribeImageGroupedVulListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGroupedVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListResponseBody) SetTotalCount(v int32) *DescribeImageGroupedVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetRequestId(v string) *DescribeImageGroupedVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetPageSize(v int32) *DescribeImageGroupedVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetCurrentPage(v int32) *DescribeImageGroupedVulListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetGroupedVulItems(v []*DescribeImageGroupedVulListResponseBodyGroupedVulItems) *DescribeImageGroupedVulListResponseBody {
	s.GroupedVulItems = v
	return s
}

type DescribeImageGroupedVulListResponseBodyGroupedVulItems struct {
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	NntfCount    *int32  `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	GmtLast      *int64  `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	LastScanTime *int64  `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	Tags         *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	LaterCount   *int32  `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	AliasName    *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AsapCount    *int32  `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
}

func (s DescribeImageGroupedVulListResponseBodyGroupedVulItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGroupedVulListResponseBodyGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetStatus(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Status = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetType(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Type = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetNntfCount(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.NntfCount = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetGmtLast(v int64) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.GmtLast = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetLastScanTime(v int64) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetTags(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Tags = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetLaterCount(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.LaterCount = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetAliasName(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetName(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetAsapCount(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.AsapCount = &v
	return s
}

type DescribeImageGroupedVulListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageGroupedVulListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageGroupedVulListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGroupedVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListResponse) SetHeaders(v map[string]*string) *DescribeImageGroupedVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageGroupedVulListResponse) SetBody(v *DescribeImageGroupedVulListResponseBody) *DescribeImageGroupedVulListResponse {
	s.Body = v
	return s
}

type DescribeImageStatisticsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeImageStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageStatisticsRequest) SetSourceIp(v string) *DescribeImageStatisticsRequest {
	s.SourceIp = &v
	return s
}

type DescribeImageStatisticsResponseBody struct {
	InstanceCount     *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskInstanceCount *int32  `json:"RiskInstanceCount,omitempty" xml:"RiskInstanceCount,omitempty"`
}

func (s DescribeImageStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageStatisticsResponseBody) SetInstanceCount(v int32) *DescribeImageStatisticsResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *DescribeImageStatisticsResponseBody) SetRequestId(v string) *DescribeImageStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageStatisticsResponseBody) SetRiskInstanceCount(v int32) *DescribeImageStatisticsResponseBody {
	s.RiskInstanceCount = &v
	return s
}

type DescribeImageStatisticsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageStatisticsResponse) SetHeaders(v map[string]*string) *DescribeImageStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageStatisticsResponse) SetBody(v *DescribeImageStatisticsResponseBody) *DescribeImageStatisticsResponse {
	s.Body = v
	return s
}

type DescribeImageVulListRequest struct {
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ids                 *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Remark              *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuids               *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AliasName           *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Level               *string `json:"Level,omitempty" xml:"Level,omitempty"`
	StatusList          *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	Necessity           *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Dealed              *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	BatchName           *string `json:"BatchName,omitempty" xml:"BatchName,omitempty"`
	Resource            *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	CreateTsStart       *int64  `json:"CreateTsStart,omitempty" xml:"CreateTsStart,omitempty"`
	CreateTsEnd         *int64  `json:"CreateTsEnd,omitempty" xml:"CreateTsEnd,omitempty"`
	CurrentPage         *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ModifyTsStart       *int64  `json:"ModifyTsStart,omitempty" xml:"ModifyTsStart,omitempty"`
	ModifyTsEnd         *int64  `json:"ModifyTsEnd,omitempty" xml:"ModifyTsEnd,omitempty"`
	CveId               *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	RepoName            *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RepoId              *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	Tag                 *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Digest              *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	ContainerFieldName  *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeImageVulListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListRequest) SetSourceIp(v string) *DescribeImageVulListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeImageVulListRequest) SetLang(v string) *DescribeImageVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageVulListRequest) SetIds(v string) *DescribeImageVulListRequest {
	s.Ids = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRemark(v string) *DescribeImageVulListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeImageVulListRequest) SetGroupId(v string) *DescribeImageVulListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetType(v string) *DescribeImageVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeImageVulListRequest) SetUuids(v string) *DescribeImageVulListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeImageVulListRequest) SetName(v string) *DescribeImageVulListRequest {
	s.Name = &v
	return s
}

func (s *DescribeImageVulListRequest) SetAliasName(v string) *DescribeImageVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeImageVulListRequest) SetLevel(v string) *DescribeImageVulListRequest {
	s.Level = &v
	return s
}

func (s *DescribeImageVulListRequest) SetStatusList(v string) *DescribeImageVulListRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeImageVulListRequest) SetNecessity(v string) *DescribeImageVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeImageVulListRequest) SetDealed(v string) *DescribeImageVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeImageVulListRequest) SetBatchName(v string) *DescribeImageVulListRequest {
	s.BatchName = &v
	return s
}

func (s *DescribeImageVulListRequest) SetResource(v string) *DescribeImageVulListRequest {
	s.Resource = &v
	return s
}

func (s *DescribeImageVulListRequest) SetCreateTsStart(v int64) *DescribeImageVulListRequest {
	s.CreateTsStart = &v
	return s
}

func (s *DescribeImageVulListRequest) SetCreateTsEnd(v int64) *DescribeImageVulListRequest {
	s.CreateTsEnd = &v
	return s
}

func (s *DescribeImageVulListRequest) SetCurrentPage(v int32) *DescribeImageVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageVulListRequest) SetPageSize(v int32) *DescribeImageVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageVulListRequest) SetModifyTsStart(v int64) *DescribeImageVulListRequest {
	s.ModifyTsStart = &v
	return s
}

func (s *DescribeImageVulListRequest) SetModifyTsEnd(v int64) *DescribeImageVulListRequest {
	s.ModifyTsEnd = &v
	return s
}

func (s *DescribeImageVulListRequest) SetCveId(v string) *DescribeImageVulListRequest {
	s.CveId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRepoName(v string) *DescribeImageVulListRequest {
	s.RepoName = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRegionId(v string) *DescribeImageVulListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetInstanceId(v string) *DescribeImageVulListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetRepoId(v string) *DescribeImageVulListRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeImageVulListRequest) SetTag(v string) *DescribeImageVulListRequest {
	s.Tag = &v
	return s
}

func (s *DescribeImageVulListRequest) SetDigest(v string) *DescribeImageVulListRequest {
	s.Digest = &v
	return s
}

func (s *DescribeImageVulListRequest) SetContainerFieldName(v string) *DescribeImageVulListRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeImageVulListRequest) SetContainerFieldValue(v string) *DescribeImageVulListRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeImageVulListRequest) SetTargetType(v string) *DescribeImageVulListRequest {
	s.TargetType = &v
	return s
}

type DescribeImageVulListResponseBody struct {
	TotalCount  *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VulRecords  []*DescribeImageVulListResponseBodyVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" type:"Repeated"`
	CurrentPage *int32                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeImageVulListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBody) SetTotalCount(v int32) *DescribeImageVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageVulListResponseBody) SetRequestId(v string) *DescribeImageVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageVulListResponseBody) SetPageSize(v int32) *DescribeImageVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageVulListResponseBody) SetVulRecords(v []*DescribeImageVulListResponseBodyVulRecords) *DescribeImageVulListResponseBody {
	s.VulRecords = v
	return s
}

func (s *DescribeImageVulListResponseBody) SetCurrentPage(v int32) *DescribeImageVulListResponseBody {
	s.CurrentPage = &v
	return s
}

type DescribeImageVulListResponseBodyVulRecords struct {
	Status            *int32                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Type              *string                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	ModifyTs          *int64                                                       `json:"ModifyTs,omitempty" xml:"ModifyTs,omitempty"`
	ImageDigest       *string                                                      `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	Layers            []*string                                                    `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	PrimaryId         *int64                                                       `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	Tag               *string                                                      `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Related           *string                                                      `json:"Related,omitempty" xml:"Related,omitempty"`
	LastTs            *int64                                                       `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	FirstTs           *int64                                                       `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	Necessity         *string                                                      `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Uuid              *string                                                      `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	AliasName         *string                                                      `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name              *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	ExtendContentJson *DescribeImageVulListResponseBodyVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" type:"Struct"`
}

func (s DescribeImageVulListResponseBodyVulRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageVulListResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetStatus(v int32) *DescribeImageVulListResponseBodyVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetType(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Type = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetModifyTs(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetImageDigest(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.ImageDigest = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetLayers(v []*string) *DescribeImageVulListResponseBodyVulRecords {
	s.Layers = v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetPrimaryId(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.PrimaryId = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetTag(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetRelated(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetLastTs(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetFirstTs(v int64) *DescribeImageVulListResponseBodyVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetNecessity(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetUuid(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Uuid = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetAliasName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetName(v string) *DescribeImageVulListResponseBodyVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecords) SetExtendContentJson(v *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) *DescribeImageVulListResponseBodyVulRecords {
	s.ExtendContentJson = v
	return s
}

type DescribeImageVulListResponseBodyVulRecordsExtendContentJson struct {
	OsRelease     *string                                                                     `json:"OsRelease,omitempty" xml:"OsRelease,omitempty"`
	RpmEntityList []*DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" type:"Repeated"`
	Os            *string                                                                     `json:"Os,omitempty" xml:"Os,omitempty"`
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) SetOsRelease(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJson {
	s.OsRelease = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) *DescribeImageVulListResponseBodyVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJson) SetOs(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJson {
	s.Os = &v
	return s
}

type DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList struct {
	Layer       *string `json:"Layer,omitempty" xml:"Layer,omitempty"`
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateCmd   *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty"`
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetLayer(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Layer = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetVersion(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Version = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeImageVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

type DescribeImageVulListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageVulListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageVulListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponse) SetHeaders(v map[string]*string) *DescribeImageVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageVulListResponse) SetBody(v *DescribeImageVulListResponseBody) *DescribeImageVulListResponse {
	s.Body = v
	return s
}

type DescribeInstallCaptchaRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
}

func (s DescribeInstallCaptchaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstallCaptchaRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstallCaptchaRequest) SetSourceIp(v string) *DescribeInstallCaptchaRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstallCaptchaRequest) SetLang(v string) *DescribeInstallCaptchaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstallCaptchaRequest) SetDeadline(v string) *DescribeInstallCaptchaRequest {
	s.Deadline = &v
	return s
}

type DescribeInstallCaptchaResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Deadline    *int32  `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	CaptchaCode *string `json:"CaptchaCode,omitempty" xml:"CaptchaCode,omitempty"`
}

func (s DescribeInstallCaptchaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstallCaptchaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstallCaptchaResponseBody) SetRequestId(v string) *DescribeInstallCaptchaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstallCaptchaResponseBody) SetDeadline(v int32) *DescribeInstallCaptchaResponseBody {
	s.Deadline = &v
	return s
}

func (s *DescribeInstallCaptchaResponseBody) SetCaptchaCode(v string) *DescribeInstallCaptchaResponseBody {
	s.CaptchaCode = &v
	return s
}

type DescribeInstallCaptchaResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstallCaptchaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstallCaptchaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstallCaptchaResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstallCaptchaResponse) SetHeaders(v map[string]*string) *DescribeInstallCaptchaResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstallCaptchaResponse) SetBody(v *DescribeInstallCaptchaResponseBody) *DescribeInstallCaptchaResponse {
	s.Body = v
	return s
}

type DescribeInstanceAntiBruteForceRulesRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UuidList        []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAntiBruteForceRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetSourceIp(v string) *DescribeInstanceAntiBruteForceRulesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetResourceOwnerId(v int64) *DescribeInstanceAntiBruteForceRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesRequest) SetUuidList(v []*string) *DescribeInstanceAntiBruteForceRulesRequest {
	s.UuidList = v
	return s
}

type DescribeInstanceAntiBruteForceRulesResponseBody struct {
	PageInfo  *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeInstanceAntiBruteForceRulesResponseBodyRules  `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAntiBruteForceRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) SetPageInfo(v *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) *DescribeInstanceAntiBruteForceRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) SetRequestId(v string) *DescribeInstanceAntiBruteForceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBody) SetRules(v []*DescribeInstanceAntiBruteForceRulesResponseBodyRules) *DescribeInstanceAntiBruteForceRulesResponseBody {
	s.Rules = v
	return s
}

type DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetPageSize(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo) SetCount(v int32) *DescribeInstanceAntiBruteForceRulesResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeInstanceAntiBruteForceRulesResponseBodyRules struct {
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) SetUuid(v string) *DescribeInstanceAntiBruteForceRulesResponseBodyRules {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) SetName(v string) *DescribeInstanceAntiBruteForceRulesResponseBodyRules {
	s.Name = &v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponseBodyRules) SetId(v int64) *DescribeInstanceAntiBruteForceRulesResponseBodyRules {
	s.Id = &v
	return s
}

type DescribeInstanceAntiBruteForceRulesResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAntiBruteForceRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAntiBruteForceRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAntiBruteForceRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) SetHeaders(v map[string]*string) *DescribeInstanceAntiBruteForceRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAntiBruteForceRulesResponse) SetBody(v *DescribeInstanceAntiBruteForceRulesResponseBody) *DescribeInstanceAntiBruteForceRulesResponse {
	s.Body = v
	return s
}

type DescribeInstanceStatisticsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	From     *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetLang(v string) *DescribeInstanceStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetUuid(v string) *DescribeInstanceStatisticsRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetFrom(v string) *DescribeInstanceStatisticsRequest {
	s.From = &v
	return s
}

type DescribeInstanceStatisticsResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*DescribeInstanceStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeInstanceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBody) SetRequestId(v string) *DescribeInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBody) SetData(v []*DescribeInstanceStatisticsResponseBodyData) *DescribeInstanceStatisticsResponseBody {
	s.Data = v
	return s
}

type DescribeInstanceStatisticsResponseBodyData struct {
	Account    *int32  `json:"Account,omitempty" xml:"Account,omitempty"`
	AppNum     *int32  `json:"AppNum,omitempty" xml:"AppNum,omitempty"`
	ScaNum     *int32  `json:"ScaNum,omitempty" xml:"ScaNum,omitempty"`
	Trojan     *int32  `json:"Trojan,omitempty" xml:"Trojan,omitempty"`
	EmgNum     *int32  `json:"EmgNum,omitempty" xml:"EmgNum,omitempty"`
	CveNum     *int32  `json:"CveNum,omitempty" xml:"CveNum,omitempty"`
	Suspicious *int32  `json:"Suspicious,omitempty" xml:"Suspicious,omitempty"`
	CmsNum     *int32  `json:"CmsNum,omitempty" xml:"CmsNum,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Vul        *int32  `json:"Vul,omitempty" xml:"Vul,omitempty"`
	Health     *int32  `json:"Health,omitempty" xml:"Health,omitempty"`
	SysNum     *int32  `json:"SysNum,omitempty" xml:"SysNum,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetAccount(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Account = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetAppNum(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.AppNum = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetScaNum(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.ScaNum = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetTrojan(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Trojan = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetEmgNum(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.EmgNum = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetCveNum(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.CveNum = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetSuspicious(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Suspicious = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetCmsNum(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.CmsNum = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetUuid(v string) *DescribeInstanceStatisticsResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetVul(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Vul = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetHealth(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Health = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetSysNum(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.SysNum = &v
	return s
}

type DescribeInstanceStatisticsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponse) SetHeaders(v map[string]*string) *DescribeInstanceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceStatisticsResponse) SetBody(v *DescribeInstanceStatisticsResponseBody) *DescribeInstanceStatisticsResponse {
	s.Body = v
	return s
}

type DescribeIpInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Field    *string `json:"Field,omitempty" xml:"Field,omitempty"`
}

func (s DescribeIpInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoRequest) SetSourceIp(v string) *DescribeIpInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIpInfoRequest) SetIp(v string) *DescribeIpInfoRequest {
	s.Ip = &v
	return s
}

func (s *DescribeIpInfoRequest) SetField(v string) *DescribeIpInfoRequest {
	s.Field = &v
	return s
}

type DescribeIpInfoResponseBody struct {
	Country                 *string `json:"country,omitempty" xml:"country,omitempty"`
	MaliciousScore          *string `json:"malicious_score,omitempty" xml:"malicious_score,omitempty"`
	IsProxy7d               *string `json:"is_proxy_7d,omitempty" xml:"is_proxy_7d,omitempty"`
	DayCnt30dWebAttack      *string `json:"day_cnt_30d_web_attack,omitempty" xml:"day_cnt_30d_web_attack,omitempty"`
	IsWebAttack             *string `json:"is_web_attack,omitempty" xml:"is_web_attack,omitempty"`
	Isp                     *string `json:"isp,omitempty" xml:"isp,omitempty"`
	GmtLastC2               *string `json:"gmt_last_c2,omitempty" xml:"gmt_last_c2,omitempty"`
	IsNat1d                 *string `json:"is_nat_1d,omitempty" xml:"is_nat_1d,omitempty"`
	GmtLastMaliciousSource  *string `json:"gmt_last_malicious_source,omitempty" xml:"gmt_last_malicious_source,omitempty"`
	Province                *string `json:"province,omitempty" xml:"province,omitempty"`
	ProxyDayTrace           *string `json:"proxy_day_trace,omitempty" xml:"proxy_day_trace,omitempty"`
	GmtLastWebAttack        *string `json:"gmt_last_web_attack,omitempty" xml:"gmt_last_web_attack,omitempty"`
	IsWebAttack7d           *string `json:"is_web_attack_7d,omitempty" xml:"is_web_attack_7d,omitempty"`
	MiningPoolDayTrace      *string `json:"mining_pool_day_trace,omitempty" xml:"mining_pool_day_trace,omitempty"`
	IsC2                    *string `json:"is_c2,omitempty" xml:"is_c2,omitempty"`
	GmtLastMiningPool       *string `json:"gmt_last_mining_pool,omitempty" xml:"gmt_last_mining_pool,omitempty"`
	IsMaliciousSource1d     *string `json:"is_malicious_source_1d,omitempty" xml:"is_malicious_source_1d,omitempty"`
	IsMiningPool1d          *string `json:"is_mining_pool_1d,omitempty" xml:"is_mining_pool_1d,omitempty"`
	IsTor1d                 *string `json:"is_tor_1d,omitempty" xml:"is_tor_1d,omitempty"`
	IsNat                   *string `json:"is_nat,omitempty" xml:"is_nat,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ip                      *string `json:"ip,omitempty" xml:"ip,omitempty"`
	IsC27d                  *string `json:"is_c2_7d,omitempty" xml:"is_c2_7d,omitempty"`
	IsMaliciousLogin1d      *string `json:"is_malicious_login_1d,omitempty" xml:"is_malicious_login_1d,omitempty"`
	Tags                    *string `json:"tags,omitempty" xml:"tags,omitempty"`
	MaliciousSourceDayTrace *string `json:"malicious_source_day_trace,omitempty" xml:"malicious_source_day_trace,omitempty"`
	IsMiningPool            *string `json:"is_mining_pool,omitempty" xml:"is_mining_pool,omitempty"`
	GmtFirstMiningPool      *string `json:"gmt_first_mining_pool,omitempty" xml:"gmt_first_mining_pool,omitempty"`
	IsIdc                   *string `json:"is_idc,omitempty" xml:"is_idc,omitempty"`
	IsMaliciousLogin        *string `json:"is_malicious_login,omitempty" xml:"is_malicious_login,omitempty"`
	Rdns                    *string `json:"rdns,omitempty" xml:"rdns,omitempty"`
	City                    *string `json:"city,omitempty" xml:"city,omitempty"`
	IsTor7d                 *string `json:"is_tor_7d,omitempty" xml:"is_tor_7d,omitempty"`
	GmtLastNat              *string `json:"gmt_last_nat,omitempty" xml:"gmt_last_nat,omitempty"`
	GmtLastMaliciousLogin   *string `json:"gmt_last_malicious_login,omitempty" xml:"gmt_last_malicious_login,omitempty"`
	IsNat7d                 *string `json:"is_nat_7d,omitempty" xml:"is_nat_7d,omitempty"`
	IsProxy1d               *string `json:"is_proxy_1d,omitempty" xml:"is_proxy_1d,omitempty"`
	Geo                     *string `json:"geo,omitempty" xml:"geo,omitempty"`
	GmtFirstC2              *string `json:"gmt_first_c2,omitempty" xml:"gmt_first_c2,omitempty"`
	IsMaliciousSource7d     *string `json:"is_malicious_source_7d,omitempty" xml:"is_malicious_source_7d,omitempty"`
	IsMiningPool7d          *string `json:"is_mining_pool_7d,omitempty" xml:"is_mining_pool_7d,omitempty"`
	DayCnt7dTor             *string `json:"day_cnt_7d_tor,omitempty" xml:"day_cnt_7d_tor,omitempty"`
	IsWebAttack30d          *string `json:"is_web_attack_30d,omitempty" xml:"is_web_attack_30d,omitempty"`
	GmtLastProxy            *string `json:"gmt_last_proxy,omitempty" xml:"gmt_last_proxy,omitempty"`
	NatDayTrace             *string `json:"nat_day_trace,omitempty" xml:"nat_day_trace,omitempty"`
	TotalDayCntWebAttack    *string `json:"total_day_cnt_web_attack,omitempty" xml:"total_day_cnt_web_attack,omitempty"`
	IsWebAttack1d           *string `json:"is_web_attack_1d,omitempty" xml:"is_web_attack_1d,omitempty"`
	MaliciousLoginDayTrace  *string `json:"malicious_login_day_trace,omitempty" xml:"malicious_login_day_trace,omitempty"`
	DayCnt30dTor            *string `json:"day_cnt_30d_tor,omitempty" xml:"day_cnt_30d_tor,omitempty"`
	TotalDayCntTor          *string `json:"total_day_cnt_tor,omitempty" xml:"total_day_cnt_tor,omitempty"`
	IsProxy                 *string `json:"is_proxy,omitempty" xml:"is_proxy,omitempty"`
	C2DayTrace              *string `json:"c2_day_trace,omitempty" xml:"c2_day_trace,omitempty"`
	IsC21d                  *string `json:"is_c2_1d,omitempty" xml:"is_c2_1d,omitempty"`
	IsMaliciousSource30d    *string `json:"is_malicious_source_30d,omitempty" xml:"is_malicious_source_30d,omitempty"`
	DayCnt7dWebAttack       *string `json:"day_cnt_7d_web_attack,omitempty" xml:"day_cnt_7d_web_attack,omitempty"`
	IdcName                 *string `json:"idc_name,omitempty" xml:"idc_name,omitempty"`
	TorDayTrace             *string `json:"tor_day_trace,omitempty" xml:"tor_day_trace,omitempty"`
	IsTor                   *string `json:"is_tor,omitempty" xml:"is_tor,omitempty"`
	IsMaliciousLogin7d      *string `json:"is_malicious_login_7d,omitempty" xml:"is_malicious_login_7d,omitempty"`
	GmtLastTor              *string `json:"gmt_last_tor,omitempty" xml:"gmt_last_tor,omitempty"`
	IsMaliciousSource       *string `json:"is_malicious_source,omitempty" xml:"is_malicious_source,omitempty"`
	WebAttackDayTrace       *string `json:"web_attack_day_trace,omitempty" xml:"web_attack_day_trace,omitempty"`
}

func (s DescribeIpInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoResponseBody) SetCountry(v string) *DescribeIpInfoResponseBody {
	s.Country = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetMaliciousScore(v string) *DescribeIpInfoResponseBody {
	s.MaliciousScore = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsProxy7d(v string) *DescribeIpInfoResponseBody {
	s.IsProxy7d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetDayCnt30dWebAttack(v string) *DescribeIpInfoResponseBody {
	s.DayCnt30dWebAttack = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsWebAttack(v string) *DescribeIpInfoResponseBody {
	s.IsWebAttack = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsp(v string) *DescribeIpInfoResponseBody {
	s.Isp = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastC2(v string) *DescribeIpInfoResponseBody {
	s.GmtLastC2 = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsNat1d(v string) *DescribeIpInfoResponseBody {
	s.IsNat1d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastMaliciousSource(v string) *DescribeIpInfoResponseBody {
	s.GmtLastMaliciousSource = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetProvince(v string) *DescribeIpInfoResponseBody {
	s.Province = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetProxyDayTrace(v string) *DescribeIpInfoResponseBody {
	s.ProxyDayTrace = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastWebAttack(v string) *DescribeIpInfoResponseBody {
	s.GmtLastWebAttack = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsWebAttack7d(v string) *DescribeIpInfoResponseBody {
	s.IsWebAttack7d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetMiningPoolDayTrace(v string) *DescribeIpInfoResponseBody {
	s.MiningPoolDayTrace = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsC2(v string) *DescribeIpInfoResponseBody {
	s.IsC2 = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastMiningPool(v string) *DescribeIpInfoResponseBody {
	s.GmtLastMiningPool = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMaliciousSource1d(v string) *DescribeIpInfoResponseBody {
	s.IsMaliciousSource1d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMiningPool1d(v string) *DescribeIpInfoResponseBody {
	s.IsMiningPool1d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsTor1d(v string) *DescribeIpInfoResponseBody {
	s.IsTor1d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsNat(v string) *DescribeIpInfoResponseBody {
	s.IsNat = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRequestId(v string) *DescribeIpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIp(v string) *DescribeIpInfoResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsC27d(v string) *DescribeIpInfoResponseBody {
	s.IsC27d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMaliciousLogin1d(v string) *DescribeIpInfoResponseBody {
	s.IsMaliciousLogin1d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetTags(v string) *DescribeIpInfoResponseBody {
	s.Tags = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetMaliciousSourceDayTrace(v string) *DescribeIpInfoResponseBody {
	s.MaliciousSourceDayTrace = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMiningPool(v string) *DescribeIpInfoResponseBody {
	s.IsMiningPool = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtFirstMiningPool(v string) *DescribeIpInfoResponseBody {
	s.GmtFirstMiningPool = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsIdc(v string) *DescribeIpInfoResponseBody {
	s.IsIdc = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMaliciousLogin(v string) *DescribeIpInfoResponseBody {
	s.IsMaliciousLogin = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRdns(v string) *DescribeIpInfoResponseBody {
	s.Rdns = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetCity(v string) *DescribeIpInfoResponseBody {
	s.City = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsTor7d(v string) *DescribeIpInfoResponseBody {
	s.IsTor7d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastNat(v string) *DescribeIpInfoResponseBody {
	s.GmtLastNat = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastMaliciousLogin(v string) *DescribeIpInfoResponseBody {
	s.GmtLastMaliciousLogin = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsNat7d(v string) *DescribeIpInfoResponseBody {
	s.IsNat7d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsProxy1d(v string) *DescribeIpInfoResponseBody {
	s.IsProxy1d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGeo(v string) *DescribeIpInfoResponseBody {
	s.Geo = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtFirstC2(v string) *DescribeIpInfoResponseBody {
	s.GmtFirstC2 = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMaliciousSource7d(v string) *DescribeIpInfoResponseBody {
	s.IsMaliciousSource7d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMiningPool7d(v string) *DescribeIpInfoResponseBody {
	s.IsMiningPool7d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetDayCnt7dTor(v string) *DescribeIpInfoResponseBody {
	s.DayCnt7dTor = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsWebAttack30d(v string) *DescribeIpInfoResponseBody {
	s.IsWebAttack30d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastProxy(v string) *DescribeIpInfoResponseBody {
	s.GmtLastProxy = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetNatDayTrace(v string) *DescribeIpInfoResponseBody {
	s.NatDayTrace = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetTotalDayCntWebAttack(v string) *DescribeIpInfoResponseBody {
	s.TotalDayCntWebAttack = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsWebAttack1d(v string) *DescribeIpInfoResponseBody {
	s.IsWebAttack1d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetMaliciousLoginDayTrace(v string) *DescribeIpInfoResponseBody {
	s.MaliciousLoginDayTrace = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetDayCnt30dTor(v string) *DescribeIpInfoResponseBody {
	s.DayCnt30dTor = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetTotalDayCntTor(v string) *DescribeIpInfoResponseBody {
	s.TotalDayCntTor = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsProxy(v string) *DescribeIpInfoResponseBody {
	s.IsProxy = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetC2DayTrace(v string) *DescribeIpInfoResponseBody {
	s.C2DayTrace = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsC21d(v string) *DescribeIpInfoResponseBody {
	s.IsC21d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMaliciousSource30d(v string) *DescribeIpInfoResponseBody {
	s.IsMaliciousSource30d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetDayCnt7dWebAttack(v string) *DescribeIpInfoResponseBody {
	s.DayCnt7dWebAttack = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIdcName(v string) *DescribeIpInfoResponseBody {
	s.IdcName = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetTorDayTrace(v string) *DescribeIpInfoResponseBody {
	s.TorDayTrace = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsTor(v string) *DescribeIpInfoResponseBody {
	s.IsTor = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMaliciousLogin7d(v string) *DescribeIpInfoResponseBody {
	s.IsMaliciousLogin7d = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetGmtLastTor(v string) *DescribeIpInfoResponseBody {
	s.GmtLastTor = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIsMaliciousSource(v string) *DescribeIpInfoResponseBody {
	s.IsMaliciousSource = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetWebAttackDayTrace(v string) *DescribeIpInfoResponseBody {
	s.WebAttackDayTrace = &v
	return s
}

type DescribeIpInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoResponse) SetHeaders(v map[string]*string) *DescribeIpInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpInfoResponse) SetBody(v *DescribeIpInfoResponseBody) *DescribeIpInfoResponse {
	s.Body = v
	return s
}

type DescribeModuleConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeModuleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModuleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigRequest) SetSourceIp(v string) *DescribeModuleConfigRequest {
	s.SourceIp = &v
	return s
}

type DescribeModuleConfigResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode   *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Count            *int32                                              `json:"Count,omitempty" xml:"Count,omitempty"`
	ModuleConfigList []*DescribeModuleConfigResponseBodyModuleConfigList `json:"ModuleConfigList,omitempty" xml:"ModuleConfigList,omitempty" type:"Repeated"`
	Success          *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeModuleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponseBody) SetRequestId(v string) *DescribeModuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetHttpStatusCode(v int32) *DescribeModuleConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetCount(v int32) *DescribeModuleConfigResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetModuleConfigList(v []*DescribeModuleConfigResponseBodyModuleConfigList) *DescribeModuleConfigResponseBody {
	s.ModuleConfigList = v
	return s
}

func (s *DescribeModuleConfigResponseBody) SetSuccess(v bool) *DescribeModuleConfigResponseBody {
	s.Success = &v
	return s
}

type DescribeModuleConfigResponseBodyModuleConfigList struct {
	ModuleName *string                                                  `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ConfigName *string                                                  `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	Items      []*DescribeModuleConfigResponseBodyModuleConfigListItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeModuleConfigResponseBodyModuleConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeModuleConfigResponseBodyModuleConfigList) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) SetModuleName(v string) *DescribeModuleConfigResponseBodyModuleConfigList {
	s.ModuleName = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) SetConfigName(v string) *DescribeModuleConfigResponseBodyModuleConfigList {
	s.ConfigName = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigList) SetItems(v []*DescribeModuleConfigResponseBodyModuleConfigListItems) *DescribeModuleConfigResponseBodyModuleConfigList {
	s.Items = v
	return s
}

type DescribeModuleConfigResponseBodyModuleConfigListItems struct {
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	GroupId      *int32  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeModuleConfigResponseBodyModuleConfigListItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeModuleConfigResponseBodyModuleConfigListItems) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetUuid(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.Uuid = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetGroupId(v int32) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.GroupId = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetInstanceName(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetRegion(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.Region = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetIp(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.Ip = &v
	return s
}

func (s *DescribeModuleConfigResponseBodyModuleConfigListItems) SetInstanceId(v string) *DescribeModuleConfigResponseBodyModuleConfigListItems {
	s.InstanceId = &v
	return s
}

type DescribeModuleConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeModuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeModuleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModuleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponse) SetHeaders(v map[string]*string) *DescribeModuleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeModuleConfigResponse) SetBody(v *DescribeModuleConfigResponseBody) *DescribeModuleConfigResponse {
	s.Body = v
	return s
}

type DescribeNoticeConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeNoticeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNoticeConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeNoticeConfigRequest) SetSourceIp(v string) *DescribeNoticeConfigRequest {
	s.SourceIp = &v
	return s
}

type DescribeNoticeConfigResponseBody struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NoticeConfigList []*DescribeNoticeConfigResponseBodyNoticeConfigList `json:"NoticeConfigList,omitempty" xml:"NoticeConfigList,omitempty" type:"Repeated"`
}

func (s DescribeNoticeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNoticeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNoticeConfigResponseBody) SetRequestId(v string) *DescribeNoticeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNoticeConfigResponseBody) SetNoticeConfigList(v []*DescribeNoticeConfigResponseBodyNoticeConfigList) *DescribeNoticeConfigResponseBody {
	s.NoticeConfigList = v
	return s
}

type DescribeNoticeConfigResponseBodyNoticeConfigList struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Route       *int32  `json:"Route,omitempty" xml:"Route,omitempty"`
	TimeLimit   *int32  `json:"TimeLimit,omitempty" xml:"TimeLimit,omitempty"`
	AliUid      *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Project     *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeNoticeConfigResponseBodyNoticeConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeNoticeConfigResponseBodyNoticeConfigList) GoString() string {
	return s.String()
}

func (s *DescribeNoticeConfigResponseBodyNoticeConfigList) SetCurrentPage(v int32) *DescribeNoticeConfigResponseBodyNoticeConfigList {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNoticeConfigResponseBodyNoticeConfigList) SetRoute(v int32) *DescribeNoticeConfigResponseBodyNoticeConfigList {
	s.Route = &v
	return s
}

func (s *DescribeNoticeConfigResponseBodyNoticeConfigList) SetTimeLimit(v int32) *DescribeNoticeConfigResponseBodyNoticeConfigList {
	s.TimeLimit = &v
	return s
}

func (s *DescribeNoticeConfigResponseBodyNoticeConfigList) SetAliUid(v int64) *DescribeNoticeConfigResponseBodyNoticeConfigList {
	s.AliUid = &v
	return s
}

func (s *DescribeNoticeConfigResponseBodyNoticeConfigList) SetProject(v string) *DescribeNoticeConfigResponseBodyNoticeConfigList {
	s.Project = &v
	return s
}

type DescribeNoticeConfigResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNoticeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNoticeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNoticeConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeNoticeConfigResponse) SetHeaders(v map[string]*string) *DescribeNoticeConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeNoticeConfigResponse) SetBody(v *DescribeNoticeConfigResponseBody) *DescribeNoticeConfigResponse {
	s.Body = v
	return s
}

type DescribePropertyCountRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s DescribePropertyCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCountRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyCountRequest) SetSourceIp(v string) *DescribePropertyCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyCountRequest) SetType(v string) *DescribePropertyCountRequest {
	s.Type = &v
	return s
}

func (s *DescribePropertyCountRequest) SetUuidList(v string) *DescribePropertyCountRequest {
	s.UuidList = &v
	return s
}

type DescribePropertyCountResponseBody struct {
	Sca       *int32  `json:"Sca,omitempty" xml:"Sca,omitempty"`
	User      *int32  `json:"User,omitempty" xml:"User,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Process   *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	Software  *int32  `json:"Software,omitempty" xml:"Software,omitempty"`
	AutoRun   *int32  `json:"AutoRun,omitempty" xml:"AutoRun,omitempty"`
	Cron      *int32  `json:"Cron,omitempty" xml:"Cron,omitempty"`
}

func (s DescribePropertyCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyCountResponseBody) SetSca(v int32) *DescribePropertyCountResponseBody {
	s.Sca = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetUser(v int32) *DescribePropertyCountResponseBody {
	s.User = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetRequestId(v string) *DescribePropertyCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetPort(v int32) *DescribePropertyCountResponseBody {
	s.Port = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetProcess(v int32) *DescribePropertyCountResponseBody {
	s.Process = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetSoftware(v int32) *DescribePropertyCountResponseBody {
	s.Software = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetAutoRun(v int32) *DescribePropertyCountResponseBody {
	s.AutoRun = &v
	return s
}

func (s *DescribePropertyCountResponseBody) SetCron(v int32) *DescribePropertyCountResponseBody {
	s.Cron = &v
	return s
}

type DescribePropertyCountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCountResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyCountResponse) SetHeaders(v map[string]*string) *DescribePropertyCountResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyCountResponse) SetBody(v *DescribePropertyCountResponseBody) *DescribePropertyCountResponse {
	s.Body = v
	return s
}

type DescribePropertyCronDetailRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	User        *string `json:"User,omitempty" xml:"User,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyCronDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCronDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailRequest) SetSourceIp(v string) *DescribePropertyCronDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetRemark(v string) *DescribePropertyCronDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetSource(v string) *DescribePropertyCronDetailRequest {
	s.Source = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetUser(v string) *DescribePropertyCronDetailRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetUuid(v string) *DescribePropertyCronDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetCurrentPage(v int32) *DescribePropertyCronDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyCronDetailRequest) SetPageSize(v int32) *DescribePropertyCronDetailRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyCronDetailResponseBody struct {
	Propertys []*DescribePropertyCronDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	PageInfo  *DescribePropertyCronDetailResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyCronDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCronDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponseBody) SetPropertys(v []*DescribePropertyCronDetailResponseBodyPropertys) *DescribePropertyCronDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyCronDetailResponseBody) SetPageInfo(v *DescribePropertyCronDetailResponseBodyPageInfo) *DescribePropertyCronDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyCronDetailResponseBody) SetRequestId(v string) *DescribePropertyCronDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribePropertyCronDetailResponseBodyPropertys struct {
	Create          *string `json:"Create,omitempty" xml:"Create,omitempty"`
	InternetIp      *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Cmd             *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	IntranetIp      *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Period          *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribePropertyCronDetailResponseBodyPropertys) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCronDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetCreate(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Create = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetUser(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.User = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetSource(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Source = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetCmd(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Cmd = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetPeriod(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Period = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetMd5(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Md5 = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyCronDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

type DescribePropertyCronDetailResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyCronDetailResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCronDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyCronDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyCronDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyCronDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyCronDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyCronDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyCronDetailResponse) SetBody(v *DescribePropertyCronDetailResponseBody) *DescribePropertyCronDetailResponse {
	s.Body = v
	return s
}

type DescribePropertyPortDetailRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Port        *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ProcName    *string `json:"ProcName,omitempty" xml:"ProcName,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyPortDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailRequest) SetSourceIp(v string) *DescribePropertyPortDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetRemark(v string) *DescribePropertyPortDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetPort(v string) *DescribePropertyPortDetailRequest {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetProcName(v string) *DescribePropertyPortDetailRequest {
	s.ProcName = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetUuid(v string) *DescribePropertyPortDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetCurrentPage(v int32) *DescribePropertyPortDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetPageSize(v int32) *DescribePropertyPortDetailRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyPortDetailResponseBody struct {
	Propertys []*DescribePropertyPortDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	PageInfo  *DescribePropertyPortDetailResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyPortDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponseBody) SetPropertys(v []*DescribePropertyPortDetailResponseBodyPropertys) *DescribePropertyPortDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyPortDetailResponseBody) SetPageInfo(v *DescribePropertyPortDetailResponseBodyPageInfo) *DescribePropertyPortDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyPortDetailResponseBody) SetRequestId(v string) *DescribePropertyPortDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribePropertyPortDetailResponseBodyPropertys struct {
	Create          *string `json:"Create,omitempty" xml:"Create,omitempty"`
	InternetIp      *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	BindIp          *string `json:"BindIp,omitempty" xml:"BindIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProcName        *string `json:"ProcName,omitempty" xml:"ProcName,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	IntranetIp      *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Proto           *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
}

func (s DescribePropertyPortDetailResponseBodyPropertys) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetCreate(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Create = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetBindIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.BindIp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetProcName(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.ProcName = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetPort(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyPortDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPropertys) SetProto(v string) *DescribePropertyPortDetailResponseBodyPropertys {
	s.Proto = &v
	return s
}

type DescribePropertyPortDetailResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyPortDetailResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyPortDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyPortDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyPortDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyPortDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyPortDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyPortDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyPortDetailResponse) SetBody(v *DescribePropertyPortDetailResponseBody) *DescribePropertyPortDetailResponse {
	s.Body = v
	return s
}

type DescribePropertyPortItemRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ForceFlush  *bool   `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	Port        *string `json:"Port,omitempty" xml:"Port,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyPortItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemRequest) SetSourceIp(v string) *DescribePropertyPortItemRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyPortItemRequest) SetForceFlush(v bool) *DescribePropertyPortItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyPortItemRequest) SetPort(v string) *DescribePropertyPortItemRequest {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortItemRequest) SetCurrentPage(v int32) *DescribePropertyPortItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortItemRequest) SetPageSize(v int32) *DescribePropertyPortItemRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyPortItemResponseBody struct {
	PageInfo      *DescribePropertyPortItemResponseBodyPageInfo        `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PropertyItems []*DescribePropertyPortItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
}

func (s DescribePropertyPortItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponseBody) SetPageInfo(v *DescribePropertyPortItemResponseBodyPageInfo) *DescribePropertyPortItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyPortItemResponseBody) SetRequestId(v string) *DescribePropertyPortItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyPortItemResponseBody) SetPropertyItems(v []*DescribePropertyPortItemResponseBodyPropertyItems) *DescribePropertyPortItemResponseBody {
	s.PropertyItems = v
	return s
}

type DescribePropertyPortItemResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyPortItemResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyPortItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyPortItemResponseBodyPropertyItems struct {
	Port  *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
}

func (s DescribePropertyPortItemResponseBodyPropertyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) SetPort(v string) *DescribePropertyPortItemResponseBodyPropertyItems {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyPortItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

func (s *DescribePropertyPortItemResponseBodyPropertyItems) SetProto(v string) *DescribePropertyPortItemResponseBodyPropertyItems {
	s.Proto = &v
	return s
}

type DescribePropertyPortItemResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyPortItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyPortItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyPortItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponse) SetHeaders(v map[string]*string) *DescribePropertyPortItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyPortItemResponse) SetBody(v *DescribePropertyPortItemResponseBody) *DescribePropertyPortItemResponse {
	s.Body = v
	return s
}

type DescribePropertyProcDetailRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	User        *string `json:"User,omitempty" xml:"User,omitempty"`
	Cmdline     *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyProcDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailRequest) SetSourceIp(v string) *DescribePropertyProcDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetRemark(v string) *DescribePropertyProcDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetName(v string) *DescribePropertyProcDetailRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetUser(v string) *DescribePropertyProcDetailRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetCmdline(v string) *DescribePropertyProcDetailRequest {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetUuid(v string) *DescribePropertyProcDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetCurrentPage(v int32) *DescribePropertyProcDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetPageSize(v int32) *DescribePropertyProcDetailRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyProcDetailResponseBody struct {
	Propertys []*DescribePropertyProcDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	PageInfo  *DescribePropertyProcDetailResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyProcDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponseBody) SetPropertys(v []*DescribePropertyProcDetailResponseBodyPropertys) *DescribePropertyProcDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyProcDetailResponseBody) SetPageInfo(v *DescribePropertyProcDetailResponseBodyPageInfo) *DescribePropertyProcDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyProcDetailResponseBody) SetRequestId(v string) *DescribePropertyProcDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribePropertyProcDetailResponseBodyPropertys struct {
	Create          *string `json:"Create,omitempty" xml:"Create,omitempty"`
	InternetIp      *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Pid             *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Cmdline         *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	IntranetIp      *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	EuidName        *string `json:"EuidName,omitempty" xml:"EuidName,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Pname           *string `json:"Pname,omitempty" xml:"Pname,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribePropertyProcDetailResponseBodyPropertys) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetCreate(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Create = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetPid(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Pid = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetUser(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.User = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetCmdline(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetEuidName(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.EuidName = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetStartTime(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.StartTime = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetPname(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Pname = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetPath(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Path = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetMd5(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Md5 = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetName(v string) *DescribePropertyProcDetailResponseBodyPropertys {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyProcDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

type DescribePropertyProcDetailResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyProcDetailResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyProcDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyProcDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyProcDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyProcDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyProcDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyProcDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyProcDetailResponse) SetBody(v *DescribePropertyProcDetailResponseBody) *DescribePropertyProcDetailResponse {
	s.Body = v
	return s
}

type DescribePropertyProcItemRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ForceFlush  *bool   `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyProcItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemRequest) SetSourceIp(v string) *DescribePropertyProcItemRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyProcItemRequest) SetForceFlush(v bool) *DescribePropertyProcItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyProcItemRequest) SetName(v string) *DescribePropertyProcItemRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcItemRequest) SetCurrentPage(v int32) *DescribePropertyProcItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcItemRequest) SetPageSize(v int32) *DescribePropertyProcItemRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyProcItemResponseBody struct {
	PageInfo      *DescribePropertyProcItemResponseBodyPageInfo        `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PropertyItems []*DescribePropertyProcItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
}

func (s DescribePropertyProcItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponseBody) SetPageInfo(v *DescribePropertyProcItemResponseBodyPageInfo) *DescribePropertyProcItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyProcItemResponseBody) SetRequestId(v string) *DescribePropertyProcItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyProcItemResponseBody) SetPropertyItems(v []*DescribePropertyProcItemResponseBodyPropertyItems) *DescribePropertyProcItemResponseBody {
	s.PropertyItems = v
	return s
}

type DescribePropertyProcItemResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyProcItemResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyProcItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyProcItemResponseBodyPropertyItems struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyProcItemResponseBodyPropertyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponseBodyPropertyItems) SetName(v string) *DescribePropertyProcItemResponseBodyPropertyItems {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyProcItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

type DescribePropertyProcItemResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyProcItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyProcItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyProcItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemResponse) SetHeaders(v map[string]*string) *DescribePropertyProcItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyProcItemResponse) SetBody(v *DescribePropertyProcItemResponseBody) *DescribePropertyProcItemResponse {
	s.Body = v
	return s
}

type DescribePropertyScaDetailRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SearchItem    *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	SearchInfo    *string `json:"SearchInfo,omitempty" xml:"SearchInfo,omitempty"`
	ScaName       *string `json:"ScaName,omitempty" xml:"ScaName,omitempty"`
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	SearchItemSub *string `json:"SearchItemSub,omitempty" xml:"SearchItemSub,omitempty"`
	SearchInfoSub *string `json:"SearchInfoSub,omitempty" xml:"SearchInfoSub,omitempty"`
	Remark        *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Name          *int64  `json:"Name,omitempty" xml:"Name,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyScaDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyScaDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailRequest) SetSourceIp(v string) *DescribePropertyScaDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetLang(v string) *DescribePropertyScaDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchItem(v string) *DescribePropertyScaDetailRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchInfo(v string) *DescribePropertyScaDetailRequest {
	s.SearchInfo = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetScaName(v string) *DescribePropertyScaDetailRequest {
	s.ScaName = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetBizType(v string) *DescribePropertyScaDetailRequest {
	s.BizType = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchItemSub(v string) *DescribePropertyScaDetailRequest {
	s.SearchItemSub = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchInfoSub(v string) *DescribePropertyScaDetailRequest {
	s.SearchInfoSub = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetRemark(v string) *DescribePropertyScaDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetName(v int64) *DescribePropertyScaDetailRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetUuid(v string) *DescribePropertyScaDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetCurrentPage(v int32) *DescribePropertyScaDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetPageSize(v int32) *DescribePropertyScaDetailRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyScaDetailResponseBody struct {
	Propertys []*DescribePropertyScaDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	PageInfo  *DescribePropertyScaDetailResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyScaDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyScaDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponseBody) SetPropertys(v []*DescribePropertyScaDetailResponseBodyPropertys) *DescribePropertyScaDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyScaDetailResponseBody) SetPageInfo(v *DescribePropertyScaDetailResponseBodyPageInfo) *DescribePropertyScaDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyScaDetailResponseBody) SetRequestId(v string) *DescribePropertyScaDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribePropertyScaDetailResponseBodyPropertys struct {
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	BizTypeDispaly  *string `json:"BizTypeDispaly,omitempty" xml:"BizTypeDispaly,omitempty"`
	ProcessStarted  *int64  `json:"ProcessStarted,omitempty" xml:"ProcessStarted,omitempty"`
	InternetIp      *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	ConfigPath      *string `json:"ConfigPath,omitempty" xml:"ConfigPath,omitempty"`
	Pid             *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Cmdline         *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ListenIp        *string `json:"ListenIp,omitempty" xml:"ListenIp,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ListenStatus    *string `json:"ListenStatus,omitempty" xml:"ListenStatus,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Create          *string `json:"Create,omitempty" xml:"Create,omitempty"`
	ProcessUser     *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	WebPath         *string `json:"WebPath,omitempty" xml:"WebPath,omitempty"`
	Ppid            *string `json:"Ppid,omitempty" xml:"Ppid,omitempty"`
	IntranetIp      *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	ListenProtocol  *string `json:"ListenProtocol,omitempty" xml:"ListenProtocol,omitempty"`
	ImageName       *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ContainerName   *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Proof           *string `json:"Proof,omitempty" xml:"Proof,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribePropertyScaDetailResponseBodyPropertys) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyScaDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetType(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Type = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetBizTypeDispaly(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.BizTypeDispaly = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetProcessStarted(v int64) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ProcessStarted = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetConfigPath(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ConfigPath = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPid(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Pid = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPort(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Port = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetCmdline(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetBizType(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.BizType = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetListenIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ListenIp = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetVersion(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Version = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetListenStatus(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ListenStatus = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Name = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetCreate(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Create = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetProcessUser(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ProcessUser = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetWebPath(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.WebPath = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPpid(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Ppid = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetListenProtocol(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ListenProtocol = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetImageName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ImageName = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetPath(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Path = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetContainerName(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.ContainerName = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetProof(v string) *DescribePropertyScaDetailResponseBodyPropertys {
	s.Proof = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyScaDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

type DescribePropertyScaDetailResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyScaDetailResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyScaDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyScaDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyScaDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyScaDetailResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyScaDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyScaDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyScaDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyScaDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyScaDetailResponse) SetBody(v *DescribePropertyScaDetailResponseBody) *DescribePropertyScaDetailResponse {
	s.Body = v
	return s
}

type DescribePropertySoftwareDetailRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertySoftwareDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailRequest) SetSourceIp(v string) *DescribePropertySoftwareDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetRemark(v string) *DescribePropertySoftwareDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetName(v string) *DescribePropertySoftwareDetailRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetPath(v string) *DescribePropertySoftwareDetailRequest {
	s.Path = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetSoftwareVersion(v string) *DescribePropertySoftwareDetailRequest {
	s.SoftwareVersion = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetUuid(v string) *DescribePropertySoftwareDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetCurrentPage(v int32) *DescribePropertySoftwareDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareDetailRequest) SetPageSize(v int32) *DescribePropertySoftwareDetailRequest {
	s.PageSize = &v
	return s
}

type DescribePropertySoftwareDetailResponseBody struct {
	Propertys []*DescribePropertySoftwareDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	PageInfo  *DescribePropertySoftwareDetailResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertySoftwareDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponseBody) SetPropertys(v []*DescribePropertySoftwareDetailResponseBodyPropertys) *DescribePropertySoftwareDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBody) SetPageInfo(v *DescribePropertySoftwareDetailResponseBodyPageInfo) *DescribePropertySoftwareDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBody) SetRequestId(v string) *DescribePropertySoftwareDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribePropertySoftwareDetailResponseBodyPropertys struct {
	Create          *int64  `json:"Create,omitempty" xml:"Create,omitempty"`
	InternetIp      *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntranetIp      *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	InstallTime     *string `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribePropertySoftwareDetailResponseBodyPropertys) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetCreate(v int64) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Create = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetIp(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInstallTime(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InstallTime = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetVersion(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Version = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetPath(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Path = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetName(v string) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertySoftwareDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

type DescribePropertySoftwareDetailResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertySoftwareDetailResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertySoftwareDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertySoftwareDetailResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertySoftwareDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertySoftwareDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponse) SetHeaders(v map[string]*string) *DescribePropertySoftwareDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertySoftwareDetailResponse) SetBody(v *DescribePropertySoftwareDetailResponseBody) *DescribePropertySoftwareDetailResponse {
	s.Body = v
	return s
}

type DescribePropertySoftwareItemRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ForceFlush  *bool   `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertySoftwareItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemRequest) SetSourceIp(v string) *DescribePropertySoftwareItemRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) SetForceFlush(v bool) *DescribePropertySoftwareItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) SetName(v string) *DescribePropertySoftwareItemRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) SetCurrentPage(v int32) *DescribePropertySoftwareItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) SetPageSize(v int32) *DescribePropertySoftwareItemRequest {
	s.PageSize = &v
	return s
}

type DescribePropertySoftwareItemResponseBody struct {
	PageInfo      *DescribePropertySoftwareItemResponseBodyPageInfo        `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId     *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PropertyItems []*DescribePropertySoftwareItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
}

func (s DescribePropertySoftwareItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponseBody) SetPageInfo(v *DescribePropertySoftwareItemResponseBodyPageInfo) *DescribePropertySoftwareItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertySoftwareItemResponseBody) SetRequestId(v string) *DescribePropertySoftwareItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBody) SetPropertyItems(v []*DescribePropertySoftwareItemResponseBodyPropertyItems) *DescribePropertySoftwareItemResponseBody {
	s.PropertyItems = v
	return s
}

type DescribePropertySoftwareItemResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertySoftwareItemResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertySoftwareItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertySoftwareItemResponseBodyPropertyItems struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertySoftwareItemResponseBodyPropertyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponseBodyPropertyItems) SetName(v string) *DescribePropertySoftwareItemResponseBodyPropertyItems {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertySoftwareItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

type DescribePropertySoftwareItemResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertySoftwareItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertySoftwareItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertySoftwareItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponse) SetHeaders(v map[string]*string) *DescribePropertySoftwareItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertySoftwareItemResponse) SetBody(v *DescribePropertySoftwareItemResponseBody) *DescribePropertySoftwareItemResponse {
	s.Body = v
	return s
}

type DescribePropertyUsageNewestRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePropertyUsageNewestRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUsageNewestRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageNewestRequest) SetSourceIp(v string) *DescribePropertyUsageNewestRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyUsageNewestRequest) SetType(v string) *DescribePropertyUsageNewestRequest {
	s.Type = &v
	return s
}

type DescribePropertyUsageNewestResponseBody struct {
	NewestStatisticItems []*DescribePropertyUsageNewestResponseBodyNewestStatisticItems `json:"NewestStatisticItems,omitempty" xml:"NewestStatisticItems,omitempty" type:"Repeated"`
	Type                 *string                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	RequestId            *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ItemCount            *int32                                                         `json:"ItemCount,omitempty" xml:"ItemCount,omitempty"`
}

func (s DescribePropertyUsageNewestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUsageNewestResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageNewestResponseBody) SetNewestStatisticItems(v []*DescribePropertyUsageNewestResponseBodyNewestStatisticItems) *DescribePropertyUsageNewestResponseBody {
	s.NewestStatisticItems = v
	return s
}

func (s *DescribePropertyUsageNewestResponseBody) SetType(v string) *DescribePropertyUsageNewestResponseBody {
	s.Type = &v
	return s
}

func (s *DescribePropertyUsageNewestResponseBody) SetRequestId(v string) *DescribePropertyUsageNewestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyUsageNewestResponseBody) SetItemCount(v int32) *DescribePropertyUsageNewestResponseBody {
	s.ItemCount = &v
	return s
}

type DescribePropertyUsageNewestResponseBodyNewestStatisticItems struct {
	Create *int64  `json:"Create,omitempty" xml:"Create,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePropertyUsageNewestResponseBodyNewestStatisticItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUsageNewestResponseBodyNewestStatisticItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageNewestResponseBodyNewestStatisticItems) SetCreate(v int64) *DescribePropertyUsageNewestResponseBodyNewestStatisticItems {
	s.Create = &v
	return s
}

func (s *DescribePropertyUsageNewestResponseBodyNewestStatisticItems) SetName(v string) *DescribePropertyUsageNewestResponseBodyNewestStatisticItems {
	s.Name = &v
	return s
}

type DescribePropertyUsageNewestResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyUsageNewestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyUsageNewestResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUsageNewestResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyUsageNewestResponse) SetHeaders(v map[string]*string) *DescribePropertyUsageNewestResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyUsageNewestResponse) SetBody(v *DescribePropertyUsageNewestResponseBody) *DescribePropertyUsageNewestResponse {
	s.Body = v
	return s
}

type DescribePropertyUserDetailRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	User        *string `json:"User,omitempty" xml:"User,omitempty"`
	IsRoot      *string `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	Uuid        *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyUserDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailRequest) SetSourceIp(v string) *DescribePropertyUserDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetRemark(v string) *DescribePropertyUserDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetUser(v string) *DescribePropertyUserDetailRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetIsRoot(v string) *DescribePropertyUserDetailRequest {
	s.IsRoot = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetUuid(v string) *DescribePropertyUserDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetCurrentPage(v int32) *DescribePropertyUserDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserDetailRequest) SetPageSize(v int32) *DescribePropertyUserDetailRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyUserDetailResponseBody struct {
	Propertys []*DescribePropertyUserDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	PageInfo  *DescribePropertyUserDetailResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyUserDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponseBody) SetPropertys(v []*DescribePropertyUserDetailResponseBodyPropertys) *DescribePropertyUserDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyUserDetailResponseBody) SetPageInfo(v *DescribePropertyUserDetailResponseBodyPageInfo) *DescribePropertyUserDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyUserDetailResponseBody) SetRequestId(v string) *DescribePropertyUserDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribePropertyUserDetailResponseBodyPropertys struct {
	Status                 *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	LastLoginIp            *string   `json:"LastLoginIp,omitempty" xml:"LastLoginIp,omitempty"`
	Create                 *string   `json:"Create,omitempty" xml:"Create,omitempty"`
	InternetIp             *string   `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IsRoot                 *string   `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	LastLoginTime          *string   `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	User                   *string   `json:"User,omitempty" xml:"User,omitempty"`
	Ip                     *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PasswordExpirationDate *string   `json:"PasswordExpirationDate,omitempty" xml:"PasswordExpirationDate,omitempty"`
	IntranetIp             *string   `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	GroupNames             []*string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty" type:"Repeated"`
	Uuid                   *string   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	LastLoginTimestamp     *int64    `json:"LastLoginTimestamp,omitempty" xml:"LastLoginTimestamp,omitempty"`
	InstanceName           *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	AccountsExpirationDate *string   `json:"AccountsExpirationDate,omitempty" xml:"AccountsExpirationDate,omitempty"`
	CreateTimestamp        *int64    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
}

func (s DescribePropertyUserDetailResponseBodyPropertys) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetStatus(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.Status = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetLastLoginIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.LastLoginIp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetCreate(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.Create = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIsRoot(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IsRoot = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetLastLoginTime(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.LastLoginTime = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetUser(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.User = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetPasswordExpirationDate(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.PasswordExpirationDate = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetGroupNames(v []*string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.GroupNames = v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetLastLoginTimestamp(v int64) *DescribePropertyUserDetailResponseBodyPropertys {
	s.LastLoginTimestamp = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetAccountsExpirationDate(v string) *DescribePropertyUserDetailResponseBodyPropertys {
	s.AccountsExpirationDate = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyUserDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

type DescribePropertyUserDetailResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyUserDetailResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyUserDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyUserDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyUserDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyUserDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyUserDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyUserDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyUserDetailResponse) SetBody(v *DescribePropertyUserDetailResponseBody) *DescribePropertyUserDetailResponse {
	s.Body = v
	return s
}

type DescribePropertyUserItemRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ForceFlush  *bool   `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	User        *string `json:"User,omitempty" xml:"User,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyUserItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemRequest) SetSourceIp(v string) *DescribePropertyUserItemRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePropertyUserItemRequest) SetForceFlush(v bool) *DescribePropertyUserItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyUserItemRequest) SetUser(v string) *DescribePropertyUserItemRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyUserItemRequest) SetCurrentPage(v int32) *DescribePropertyUserItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserItemRequest) SetPageSize(v int32) *DescribePropertyUserItemRequest {
	s.PageSize = &v
	return s
}

type DescribePropertyUserItemResponseBody struct {
	PageInfo      *DescribePropertyUserItemResponseBodyPageInfo        `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PropertyItems []*DescribePropertyUserItemResponseBodyPropertyItems `json:"PropertyItems,omitempty" xml:"PropertyItems,omitempty" type:"Repeated"`
}

func (s DescribePropertyUserItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponseBody) SetPageInfo(v *DescribePropertyUserItemResponseBodyPageInfo) *DescribePropertyUserItemResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyUserItemResponseBody) SetRequestId(v string) *DescribePropertyUserItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyUserItemResponseBody) SetPropertyItems(v []*DescribePropertyUserItemResponseBodyPropertyItems) *DescribePropertyUserItemResponseBody {
	s.PropertyItems = v
	return s
}

type DescribePropertyUserItemResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyUserItemResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserItemResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPageInfo) SetCount(v int32) *DescribePropertyUserItemResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribePropertyUserItemResponseBodyPropertyItems struct {
	User  *string `json:"User,omitempty" xml:"User,omitempty"`
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePropertyUserItemResponseBodyPropertyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserItemResponseBodyPropertyItems) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponseBodyPropertyItems) SetUser(v string) *DescribePropertyUserItemResponseBodyPropertyItems {
	s.User = &v
	return s
}

func (s *DescribePropertyUserItemResponseBodyPropertyItems) SetCount(v int32) *DescribePropertyUserItemResponseBodyPropertyItems {
	s.Count = &v
	return s
}

type DescribePropertyUserItemResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePropertyUserItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePropertyUserItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePropertyUserItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserItemResponse) SetHeaders(v map[string]*string) *DescribePropertyUserItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyUserItemResponse) SetBody(v *DescribePropertyUserItemResponseBody) *DescribePropertyUserItemResponse {
	s.Body = v
	return s
}

type DescribeRiskCheckItemResultRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ItemId          *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeRiskCheckItemResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckItemResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultRequest) SetSourceIp(v string) *DescribeRiskCheckItemResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetResourceOwnerId(v int64) *DescribeRiskCheckItemResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetLang(v string) *DescribeRiskCheckItemResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetItemId(v int64) *DescribeRiskCheckItemResultRequest {
	s.ItemId = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetCurrentPage(v int32) *DescribeRiskCheckItemResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetPageSize(v int32) *DescribeRiskCheckItemResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckItemResultRequest) SetInstanceId(v string) *DescribeRiskCheckItemResultRequest {
	s.InstanceId = &v
	return s
}

type DescribeRiskCheckItemResultResponseBody struct {
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageContentResource *DescribeRiskCheckItemResultResponseBodyPageContentResource `json:"PageContentResource,omitempty" xml:"PageContentResource,omitempty" type:"Struct"`
}

func (s DescribeRiskCheckItemResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckItemResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultResponseBody) SetRequestId(v string) *DescribeRiskCheckItemResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBody) SetPageContentResource(v *DescribeRiskCheckItemResultResponseBodyPageContentResource) *DescribeRiskCheckItemResultResponseBody {
	s.PageContentResource = v
	return s
}

type DescribeRiskCheckItemResultResponseBodyPageContentResource struct {
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ContentResource *string `json:"ContentResource,omitempty" xml:"ContentResource,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount      *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Count           *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRiskCheckItemResultResponseBodyPageContentResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckItemResultResponseBodyPageContentResource) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetCurrentPage(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetContentResource(v string) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.ContentResource = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetPageSize(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetTotalCount(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.TotalCount = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetPageCount(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.PageCount = &v
	return s
}

func (s *DescribeRiskCheckItemResultResponseBodyPageContentResource) SetCount(v int32) *DescribeRiskCheckItemResultResponseBodyPageContentResource {
	s.Count = &v
	return s
}

type DescribeRiskCheckItemResultResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRiskCheckItemResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRiskCheckItemResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckItemResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckItemResultResponse) SetHeaders(v map[string]*string) *DescribeRiskCheckItemResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskCheckItemResultResponse) SetBody(v *DescribeRiskCheckItemResultResponseBody) *DescribeRiskCheckItemResultResponse {
	s.Body = v
	return s
}

type DescribeRiskCheckResultRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	GroupId         *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	CurrentPage     *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	RiskLevel       *string   `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Status          *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	AssetType       *string   `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueryFlag       *string   `json:"QueryFlag,omitempty" xml:"QueryFlag,omitempty"`
	ItemIds         []*string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Repeated"`
}

func (s DescribeRiskCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultRequest) SetSourceIp(v string) *DescribeRiskCheckResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetResourceOwnerId(v int64) *DescribeRiskCheckResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetLang(v string) *DescribeRiskCheckResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetGroupId(v int64) *DescribeRiskCheckResultRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetCurrentPage(v int32) *DescribeRiskCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetRiskLevel(v string) *DescribeRiskCheckResultRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetStatus(v string) *DescribeRiskCheckResultRequest {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetAssetType(v string) *DescribeRiskCheckResultRequest {
	s.AssetType = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetName(v string) *DescribeRiskCheckResultRequest {
	s.Name = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetPageSize(v int32) *DescribeRiskCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetInstanceId(v string) *DescribeRiskCheckResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetQueryFlag(v string) *DescribeRiskCheckResultRequest {
	s.QueryFlag = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetItemIds(v []*string) *DescribeRiskCheckResultRequest {
	s.ItemIds = v
	return s
}

type DescribeRiskCheckResultResponseBody struct {
	TotalCount  *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageCount   *int32                                     `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	CurrentPage *int32                                     `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*DescribeRiskCheckResultResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Count       *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRiskCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponseBody) SetTotalCount(v int32) *DescribeRiskCheckResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetRequestId(v string) *DescribeRiskCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetPageSize(v int32) *DescribeRiskCheckResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetPageCount(v int32) *DescribeRiskCheckResultResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetCurrentPage(v int32) *DescribeRiskCheckResultResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetList(v []*DescribeRiskCheckResultResponseBodyList) *DescribeRiskCheckResultResponseBody {
	s.List = v
	return s
}

func (s *DescribeRiskCheckResultResponseBody) SetCount(v int32) *DescribeRiskCheckResultResponseBody {
	s.Count = &v
	return s
}

type DescribeRiskCheckResultResponseBodyList struct {
	Type              *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	Status            *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	RiskLevel         *string                                                     `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Sort              *int32                                                      `json:"Sort,omitempty" xml:"Sort,omitempty"`
	RepairStatus      *string                                                     `json:"RepairStatus,omitempty" xml:"RepairStatus,omitempty"`
	RemainingTime     *int32                                                      `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty"`
	ItemId            *int64                                                      `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	StartStatus       *string                                                     `json:"StartStatus,omitempty" xml:"StartStatus,omitempty"`
	AffectedCount     *int32                                                      `json:"AffectedCount,omitempty" xml:"AffectedCount,omitempty"`
	RiskAssertType    *string                                                     `json:"RiskAssertType,omitempty" xml:"RiskAssertType,omitempty"`
	RiskItemResources []*DescribeRiskCheckResultResponseBodyListRiskItemResources `json:"RiskItemResources,omitempty" xml:"RiskItemResources,omitempty" type:"Repeated"`
	Title             *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	TaskId            *int64                                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CheckTime         *int64                                                      `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
}

func (s DescribeRiskCheckResultResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckResultResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponseBodyList) SetType(v string) *DescribeRiskCheckResultResponseBodyList {
	s.Type = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetStatus(v string) *DescribeRiskCheckResultResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRiskLevel(v string) *DescribeRiskCheckResultResponseBodyList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetSort(v int32) *DescribeRiskCheckResultResponseBodyList {
	s.Sort = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRepairStatus(v string) *DescribeRiskCheckResultResponseBodyList {
	s.RepairStatus = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRemainingTime(v int32) *DescribeRiskCheckResultResponseBodyList {
	s.RemainingTime = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetItemId(v int64) *DescribeRiskCheckResultResponseBodyList {
	s.ItemId = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetStartStatus(v string) *DescribeRiskCheckResultResponseBodyList {
	s.StartStatus = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetAffectedCount(v int32) *DescribeRiskCheckResultResponseBodyList {
	s.AffectedCount = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRiskAssertType(v string) *DescribeRiskCheckResultResponseBodyList {
	s.RiskAssertType = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetRiskItemResources(v []*DescribeRiskCheckResultResponseBodyListRiskItemResources) *DescribeRiskCheckResultResponseBodyList {
	s.RiskItemResources = v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetTitle(v string) *DescribeRiskCheckResultResponseBodyList {
	s.Title = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetTaskId(v int64) *DescribeRiskCheckResultResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyList) SetCheckTime(v int64) *DescribeRiskCheckResultResponseBodyList {
	s.CheckTime = &v
	return s
}

type DescribeRiskCheckResultResponseBodyListRiskItemResources struct {
	ContentResource *string `json:"ContentResource,omitempty" xml:"ContentResource,omitempty"`
	ResourceName    *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s DescribeRiskCheckResultResponseBodyListRiskItemResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckResultResponseBodyListRiskItemResources) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponseBodyListRiskItemResources) SetContentResource(v string) *DescribeRiskCheckResultResponseBodyListRiskItemResources {
	s.ContentResource = &v
	return s
}

func (s *DescribeRiskCheckResultResponseBodyListRiskItemResources) SetResourceName(v string) *DescribeRiskCheckResultResponseBodyListRiskItemResources {
	s.ResourceName = &v
	return s
}

type DescribeRiskCheckResultResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRiskCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRiskCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultResponse) SetHeaders(v map[string]*string) *DescribeRiskCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskCheckResultResponse) SetBody(v *DescribeRiskCheckResultResponseBody) *DescribeRiskCheckResultResponse {
	s.Body = v
	return s
}

type DescribeRiskCheckSummaryRequest struct {
	SourceIp                   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId            *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang                       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s DescribeRiskCheckSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryRequest) SetSourceIp(v string) *DescribeRiskCheckSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskCheckSummaryRequest) SetResourceOwnerId(v int64) *DescribeRiskCheckSummaryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskCheckSummaryRequest) SetLang(v string) *DescribeRiskCheckSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskCheckSummaryRequest) SetResourceDirectoryAccountId(v string) *DescribeRiskCheckSummaryRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

type DescribeRiskCheckSummaryResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskCheckSummary *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary `json:"RiskCheckSummary,omitempty" xml:"RiskCheckSummary,omitempty" type:"Struct"`
}

func (s DescribeRiskCheckSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBody) SetRequestId(v string) *DescribeRiskCheckSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBody) SetRiskCheckSummary(v *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) *DescribeRiskCheckSummaryResponseBody {
	s.RiskCheckSummary = v
	return s
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummary struct {
	ItemCount          *int32                                                                `json:"ItemCount,omitempty" xml:"ItemCount,omitempty"`
	Groups             []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups         `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	AffectedAssetCount *int32                                                                `json:"AffectedAssetCount,omitempty" xml:"AffectedAssetCount,omitempty"`
	DisabledRiskCount  *int32                                                                `json:"DisabledRiskCount,omitempty" xml:"DisabledRiskCount,omitempty"`
	RiskLevelCount     []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount `json:"RiskLevelCount,omitempty" xml:"RiskLevelCount,omitempty" type:"Repeated"`
	RiskCount          *int32                                                                `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	RiskRate           *float32                                                              `json:"RiskRate,omitempty" xml:"RiskRate,omitempty"`
	PreviousCount      *int32                                                                `json:"PreviousCount,omitempty" xml:"PreviousCount,omitempty"`
	PreviousTime       *int64                                                                `json:"PreviousTime,omitempty" xml:"PreviousTime,omitempty"`
	EnabledRiskCount   *int32                                                                `json:"EnabledRiskCount,omitempty" xml:"EnabledRiskCount,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetItemCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.ItemCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetGroups(v []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.Groups = v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetAffectedAssetCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.AffectedAssetCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetDisabledRiskCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.DisabledRiskCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetRiskLevelCount(v []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.RiskLevelCount = v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetRiskCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.RiskCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetRiskRate(v float32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.RiskRate = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetPreviousCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.PreviousCount = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetPreviousTime(v int64) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.PreviousTime = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary) SetEnabledRiskCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummary {
	s.EnabledRiskCount = &v
	return s
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups struct {
	Status        *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Sort          *int32                                                                     `json:"Sort,omitempty" xml:"Sort,omitempty"`
	RemainingTime *int32                                                                     `json:"RemainingTime,omitempty" xml:"RemainingTime,omitempty"`
	CountByStatus []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus `json:"CountByStatus,omitempty" xml:"CountByStatus,omitempty" type:"Repeated"`
	Title         *string                                                                    `json:"Title,omitempty" xml:"Title,omitempty"`
	Id            *int64                                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetStatus(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetSort(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Sort = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetRemainingTime(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.RemainingTime = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetCountByStatus(v []*DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.CountByStatus = v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetTitle(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Title = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups) SetId(v int64) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroups {
	s.Id = &v
	return s
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Count  *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) SetStatus(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus) SetCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryGroupsCountByStatus {
	s.Count = &v
	return s
}

type DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) SetKey(v string) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount {
	s.Key = &v
	return s
}

func (s *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount) SetCount(v int32) *DescribeRiskCheckSummaryResponseBodyRiskCheckSummaryRiskLevelCount {
	s.Count = &v
	return s
}

type DescribeRiskCheckSummaryResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRiskCheckSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRiskCheckSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskCheckSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckSummaryResponse) SetHeaders(v map[string]*string) *DescribeRiskCheckSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskCheckSummaryResponse) SetBody(v *DescribeRiskCheckSummaryResponseBody) *DescribeRiskCheckSummaryResponse {
	s.Body = v
	return s
}

type DescribeRiskItemTypeRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeRiskItemTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskItemTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskItemTypeRequest) SetSourceIp(v string) *DescribeRiskItemTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskItemTypeRequest) SetResourceOwnerId(v int64) *DescribeRiskItemTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskItemTypeRequest) SetLang(v string) *DescribeRiskItemTypeRequest {
	s.Lang = &v
	return s
}

type DescribeRiskItemTypeResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	List      []*DescribeRiskItemTypeResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeRiskItemTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskItemTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskItemTypeResponseBody) SetRequestId(v string) *DescribeRiskItemTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskItemTypeResponseBody) SetList(v []*DescribeRiskItemTypeResponseBodyList) *DescribeRiskItemTypeResponseBody {
	s.List = v
	return s
}

type DescribeRiskItemTypeResponseBodyList struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeRiskItemTypeResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskItemTypeResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeRiskItemTypeResponseBodyList) SetTitle(v string) *DescribeRiskItemTypeResponseBodyList {
	s.Title = &v
	return s
}

func (s *DescribeRiskItemTypeResponseBodyList) SetId(v int64) *DescribeRiskItemTypeResponseBodyList {
	s.Id = &v
	return s
}

type DescribeRiskItemTypeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRiskItemTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRiskItemTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskItemTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskItemTypeResponse) SetHeaders(v map[string]*string) *DescribeRiskItemTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskItemTypeResponse) SetBody(v *DescribeRiskItemTypeResponseBody) *DescribeRiskItemTypeResponse {
	s.Body = v
	return s
}

type DescribeRiskListCheckResultRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	CurrentPage     *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeRiskListCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskListCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultRequest) SetSourceIp(v string) *DescribeRiskListCheckResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetResourceOwnerId(v int64) *DescribeRiskListCheckResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetLang(v string) *DescribeRiskListCheckResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetCurrentPage(v int32) *DescribeRiskListCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetPageSize(v int32) *DescribeRiskListCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskListCheckResultRequest) SetInstanceIds(v []*string) *DescribeRiskListCheckResultRequest {
	s.InstanceIds = v
	return s
}

type DescribeRiskListCheckResultResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	List      []*DescribeRiskListCheckResultResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeRiskListCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskListCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultResponseBody) SetRequestId(v string) *DescribeRiskListCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskListCheckResultResponseBody) SetList(v []*DescribeRiskListCheckResultResponseBodyList) *DescribeRiskListCheckResultResponseBody {
	s.List = v
	return s
}

type DescribeRiskListCheckResultResponseBodyList struct {
	RiskCount  *int64  `json:"riskCount,omitempty" xml:"riskCount,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeRiskListCheckResultResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskListCheckResultResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultResponseBodyList) SetRiskCount(v int64) *DescribeRiskListCheckResultResponseBodyList {
	s.RiskCount = &v
	return s
}

func (s *DescribeRiskListCheckResultResponseBodyList) SetInstanceId(v string) *DescribeRiskListCheckResultResponseBodyList {
	s.InstanceId = &v
	return s
}

type DescribeRiskListCheckResultResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRiskListCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRiskListCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskListCheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskListCheckResultResponse) SetHeaders(v map[string]*string) *DescribeRiskListCheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskListCheckResultResponse) SetBody(v *DescribeRiskListCheckResultResponseBody) *DescribeRiskListCheckResultResponse {
	s.Body = v
	return s
}

type DescribeSasAssetStatisticsColumnRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSasAssetStatisticsColumnRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSasAssetStatisticsColumnRequest) GoString() string {
	return s.String()
}

func (s *DescribeSasAssetStatisticsColumnRequest) SetSourceIp(v string) *DescribeSasAssetStatisticsColumnRequest {
	s.SourceIp = &v
	return s
}

type DescribeSasAssetStatisticsColumnResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatisticsColumn *string `json:"StatisticsColumn,omitempty" xml:"StatisticsColumn,omitempty"`
}

func (s DescribeSasAssetStatisticsColumnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSasAssetStatisticsColumnResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSasAssetStatisticsColumnResponseBody) SetRequestId(v string) *DescribeSasAssetStatisticsColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSasAssetStatisticsColumnResponseBody) SetStatisticsColumn(v string) *DescribeSasAssetStatisticsColumnResponseBody {
	s.StatisticsColumn = &v
	return s
}

type DescribeSasAssetStatisticsColumnResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSasAssetStatisticsColumnResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSasAssetStatisticsColumnResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSasAssetStatisticsColumnResponse) GoString() string {
	return s.String()
}

func (s *DescribeSasAssetStatisticsColumnResponse) SetHeaders(v map[string]*string) *DescribeSasAssetStatisticsColumnResponse {
	s.Headers = v
	return s
}

func (s *DescribeSasAssetStatisticsColumnResponse) SetBody(v *DescribeSasAssetStatisticsColumnResponseBody) *DescribeSasAssetStatisticsColumnResponse {
	s.Body = v
	return s
}

type DescribeSearchConditionRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSearchConditionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSearchConditionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSearchConditionRequest) SetSourceIp(v string) *DescribeSearchConditionRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSearchConditionRequest) SetLang(v string) *DescribeSearchConditionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSearchConditionRequest) SetType(v string) *DescribeSearchConditionRequest {
	s.Type = &v
	return s
}

type DescribeSearchConditionResponseBody struct {
	ConditionList []*DescribeSearchConditionResponseBodyConditionList `json:"ConditionList,omitempty" xml:"ConditionList,omitempty" type:"Repeated"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSearchConditionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSearchConditionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSearchConditionResponseBody) SetConditionList(v []*DescribeSearchConditionResponseBodyConditionList) *DescribeSearchConditionResponseBody {
	s.ConditionList = v
	return s
}

func (s *DescribeSearchConditionResponseBody) SetRequestId(v string) *DescribeSearchConditionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSearchConditionResponseBodyConditionList struct {
	ConditionType    *string `json:"ConditionType,omitempty" xml:"ConditionType,omitempty"`
	NameKey          *string `json:"NameKey,omitempty" xml:"NameKey,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	FilterConditions *string `json:"FilterConditions,omitempty" xml:"FilterConditions,omitempty"`
}

func (s DescribeSearchConditionResponseBodyConditionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSearchConditionResponseBodyConditionList) GoString() string {
	return s.String()
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetConditionType(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.ConditionType = &v
	return s
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetNameKey(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.NameKey = &v
	return s
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetName(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.Name = &v
	return s
}

func (s *DescribeSearchConditionResponseBodyConditionList) SetFilterConditions(v string) *DescribeSearchConditionResponseBodyConditionList {
	s.FilterConditions = &v
	return s
}

type DescribeSearchConditionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSearchConditionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSearchConditionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSearchConditionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSearchConditionResponse) SetHeaders(v map[string]*string) *DescribeSearchConditionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSearchConditionResponse) SetBody(v *DescribeSearchConditionResponseBody) *DescribeSearchConditionResponse {
	s.Body = v
	return s
}

type DescribeSecureSuggestionRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeSecureSuggestionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecureSuggestionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionRequest) SetSourceIp(v string) *DescribeSecureSuggestionRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecureSuggestionRequest) SetLang(v string) *DescribeSecureSuggestionRequest {
	s.Lang = &v
	return s
}

type DescribeSecureSuggestionResponseBody struct {
	Suggestions []*DescribeSecureSuggestionResponseBodySuggestions `json:"Suggestions,omitempty" xml:"Suggestions,omitempty" type:"Repeated"`
	TotalCount  *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecureSuggestionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBody) SetSuggestions(v []*DescribeSecureSuggestionResponseBodySuggestions) *DescribeSecureSuggestionResponseBody {
	s.Suggestions = v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetTotalCount(v int32) *DescribeSecureSuggestionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetRequestId(v string) *DescribeSecureSuggestionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSecureSuggestionResponseBodySuggestions struct {
	SuggestType *string                                                  `json:"SuggestType,omitempty" xml:"SuggestType,omitempty"`
	Points      *int32                                                   `json:"Points,omitempty" xml:"Points,omitempty"`
	Detail      []*DescribeSecureSuggestionResponseBodySuggestionsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s DescribeSecureSuggestionResponseBodySuggestions) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBodySuggestions) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetSuggestType(v string) *DescribeSecureSuggestionResponseBodySuggestions {
	s.SuggestType = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetPoints(v int32) *DescribeSecureSuggestionResponseBodySuggestions {
	s.Points = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetDetail(v []*DescribeSecureSuggestionResponseBodySuggestionsDetail) *DescribeSecureSuggestionResponseBodySuggestions {
	s.Detail = v
	return s
}

type DescribeSecureSuggestionResponseBodySuggestionsDetail struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	SubType     *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s DescribeSecureSuggestionResponseBodySuggestionsDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBodySuggestionsDetail) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetDescription(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.Description = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetTitle(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.Title = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetSubType(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.SubType = &v
	return s
}

type DescribeSecureSuggestionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecureSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecureSuggestionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecureSuggestionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponse) SetHeaders(v map[string]*string) *DescribeSecureSuggestionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecureSuggestionResponse) SetBody(v *DescribeSecureSuggestionResponseBody) *DescribeSecureSuggestionResponse {
	s.Body = v
	return s
}

type DescribeSecurityCheckScheduleConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeSecurityCheckScheduleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigRequest) SetSourceIp(v string) *DescribeSecurityCheckScheduleConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigRequest) SetResourceOwnerId(v int64) *DescribeSecurityCheckScheduleConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigRequest) SetLang(v string) *DescribeSecurityCheckScheduleConfigRequest {
	s.Lang = &v
	return s
}

type DescribeSecurityCheckScheduleConfigResponseBody struct {
	RiskCheckJobConfig *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig `json:"RiskCheckJobConfig,omitempty" xml:"RiskCheckJobConfig,omitempty" type:"Struct"`
	RequestId          *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityCheckScheduleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigResponseBody) SetRiskCheckJobConfig(v *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) *DescribeSecurityCheckScheduleConfigResponseBody {
	s.RiskCheckJobConfig = v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBody) SetRequestId(v string) *DescribeSecurityCheckScheduleConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig struct {
	EndTime    *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DaysOfWeek *string `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty"`
}

func (s DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) SetEndTime(v int32) *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig {
	s.EndTime = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) SetStartTime(v int32) *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig {
	s.StartTime = &v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig) SetDaysOfWeek(v string) *DescribeSecurityCheckScheduleConfigResponseBodyRiskCheckJobConfig {
	s.DaysOfWeek = &v
	return s
}

type DescribeSecurityCheckScheduleConfigResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityCheckScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityCheckScheduleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityCheckScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityCheckScheduleConfigResponse) SetHeaders(v map[string]*string) *DescribeSecurityCheckScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityCheckScheduleConfigResponse) SetBody(v *DescribeSecurityCheckScheduleConfigResponseBody) *DescribeSecurityCheckScheduleConfigResponse {
	s.Body = v
	return s
}

type DescribeSecurityEventOperationsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s DescribeSecurityEventOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsRequest) SetSourceIp(v string) *DescribeSecurityEventOperationsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetResourceOwnerId(v int64) *DescribeSecurityEventOperationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityEventOperationsRequest) SetSecurityEventId(v int64) *DescribeSecurityEventOperationsRequest {
	s.SecurityEventId = &v
	return s
}

type DescribeSecurityEventOperationsResponseBody struct {
	RequestId                       *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEventOperationsResponse []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse `json:"SecurityEventOperationsResponse,omitempty" xml:"SecurityEventOperationsResponse,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBody) SetSecurityEventOperationsResponse(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) *DescribeSecurityEventOperationsResponseBody {
	s.SecurityEventOperationsResponse = v
	return s
}

type DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse struct {
	OperationParams  *string                                                                                       `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	OperationCode    *string                                                                                       `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	MarkField        []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField        `json:"MarkField,omitempty" xml:"MarkField,omitempty" type:"Repeated"`
	UserCanOperate   *bool                                                                                         `json:"UserCanOperate,omitempty" xml:"UserCanOperate,omitempty"`
	MarkFieldsSource []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource `json:"MarkFieldsSource,omitempty" xml:"MarkFieldsSource,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetOperationParams(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.OperationParams = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetOperationCode(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.OperationCode = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetMarkField(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.MarkField = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetUserCanOperate(v bool) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.UserCanOperate = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse) SetMarkFieldsSource(v []*DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponse {
	s.MarkFieldsSource = v
	return s
}

type DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField struct {
	MarkMisType      *string   `json:"MarkMisType,omitempty" xml:"MarkMisType,omitempty"`
	MarkMisValue     *string   `json:"MarkMisValue,omitempty" xml:"MarkMisValue,omitempty"`
	SupportedMisType []*string `json:"SupportedMisType,omitempty" xml:"SupportedMisType,omitempty" type:"Repeated"`
	FiledName        *string   `json:"FiledName,omitempty" xml:"FiledName,omitempty"`
	FiledAliasName   *string   `json:"FiledAliasName,omitempty" xml:"FiledAliasName,omitempty"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetMarkMisType(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.MarkMisType = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetMarkMisValue(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.MarkMisValue = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetSupportedMisType(v []*string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.SupportedMisType = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetFiledName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.FiledName = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField) SetFiledAliasName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkField {
	s.FiledAliasName = &v
	return s
}

type DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource struct {
	MarkMisValue     *string   `json:"MarkMisValue,omitempty" xml:"MarkMisValue,omitempty"`
	SupportedMisType []*string `json:"SupportedMisType,omitempty" xml:"SupportedMisType,omitempty" type:"Repeated"`
	FiledName        *string   `json:"FiledName,omitempty" xml:"FiledName,omitempty"`
	FiledAliasName   *string   `json:"FiledAliasName,omitempty" xml:"FiledAliasName,omitempty"`
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetMarkMisValue(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.MarkMisValue = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetSupportedMisType(v []*string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.SupportedMisType = v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetFiledName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.FiledName = &v
	return s
}

func (s *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource) SetFiledAliasName(v string) *DescribeSecurityEventOperationsResponseBodySecurityEventOperationsResponseMarkFieldsSource {
	s.FiledAliasName = &v
	return s
}

type DescribeSecurityEventOperationsResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityEventOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationsResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationsResponse) SetBody(v *DescribeSecurityEventOperationsResponseBody) *DescribeSecurityEventOperationsResponse {
	s.Body = v
	return s
}

type DescribeSecurityEventOperationStatusRequest struct {
	SourceIp         *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId  *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId           *int64    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSourceIp(v string) *DescribeSecurityEventOperationStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetResourceOwnerId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetTaskId(v int64) *DescribeSecurityEventOperationStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusRequest) SetSecurityEventIds(v []*string) *DescribeSecurityEventOperationStatusRequest {
	s.SecurityEventIds = v
	return s
}

type DescribeSecurityEventOperationStatusResponseBody struct {
	SecurityEventOperationStatusResponse *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse `json:"SecurityEventOperationStatusResponse,omitempty" xml:"SecurityEventOperationStatusResponse,omitempty" type:"Struct"`
	RequestId                            *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetSecurityEventOperationStatusResponse(v *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) *DescribeSecurityEventOperationStatusResponseBody {
	s.SecurityEventOperationStatusResponse = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBody) SetRequestId(v string) *DescribeSecurityEventOperationStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse struct {
	TaskStatus                     *string                                                                                                               `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	SecurityEventOperationStatuses []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses `json:"SecurityEventOperationStatuses,omitempty" xml:"SecurityEventOperationStatuses,omitempty" type:"Repeated"`
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) SetTaskStatus(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse {
	s.TaskStatus = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse) SetSecurityEventOperationStatuses(v []*DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponse {
	s.SecurityEventOperationStatuses = v
	return s
}

type DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SecurityEventId *string `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	ErrorCode       *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetStatus(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.Status = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetSecurityEventId(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses) SetErrorCode(v string) *DescribeSecurityEventOperationStatusResponseBodySecurityEventOperationStatusResponseSecurityEventOperationStatuses {
	s.ErrorCode = &v
	return s
}

type DescribeSecurityEventOperationStatusResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityEventOperationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityEventOperationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) SetBody(v *DescribeSecurityEventOperationStatusResponseBody) *DescribeSecurityEventOperationStatusResponse {
	s.Body = v
	return s
}

type DescribeSecurityStatInfoRequest struct {
	SourceIp                   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s DescribeSecurityStatInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityStatInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoRequest) SetSourceIp(v string) *DescribeSecurityStatInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecurityStatInfoRequest) SetLang(v string) *DescribeSecurityStatInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecurityStatInfoRequest) SetResourceDirectoryAccountId(v string) *DescribeSecurityStatInfoRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

type DescribeSecurityStatInfoResponseBody struct {
	SecurityEvent *DescribeSecurityStatInfoResponseBodySecurityEvent `json:"SecurityEvent,omitempty" xml:"SecurityEvent,omitempty" type:"Struct"`
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HealthCheck   *DescribeSecurityStatInfoResponseBodyHealthCheck   `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	Vulnerability *DescribeSecurityStatInfoResponseBodyVulnerability `json:"Vulnerability,omitempty" xml:"Vulnerability,omitempty" type:"Struct"`
	AttackEvent   *DescribeSecurityStatInfoResponseBodyAttackEvent   `json:"AttackEvent,omitempty" xml:"AttackEvent,omitempty" type:"Struct"`
	Success       *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSecurityStatInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBody) SetSecurityEvent(v *DescribeSecurityStatInfoResponseBodySecurityEvent) *DescribeSecurityStatInfoResponseBody {
	s.SecurityEvent = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetRequestId(v string) *DescribeSecurityStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetHealthCheck(v *DescribeSecurityStatInfoResponseBodyHealthCheck) *DescribeSecurityStatInfoResponseBody {
	s.HealthCheck = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetVulnerability(v *DescribeSecurityStatInfoResponseBodyVulnerability) *DescribeSecurityStatInfoResponseBody {
	s.Vulnerability = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetAttackEvent(v *DescribeSecurityStatInfoResponseBodyAttackEvent) *DescribeSecurityStatInfoResponseBody {
	s.AttackEvent = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetSuccess(v bool) *DescribeSecurityStatInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeSecurityStatInfoResponseBodySecurityEvent struct {
	SuspiciousCount *int32    `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
	RemindCount     *int32    `json:"RemindCount,omitempty" xml:"RemindCount,omitempty"`
	ValueArray      []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
	TimeArray       []*string `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Repeated"`
	RemindList      []*string `json:"RemindList,omitempty" xml:"RemindList,omitempty" type:"Repeated"`
	LevelsOn        []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	SeriousCount    *int32    `json:"SeriousCount,omitempty" xml:"SeriousCount,omitempty"`
	TotalCount      *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DateArray       []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	SuspiciousList  []*string `json:"SuspiciousList,omitempty" xml:"SuspiciousList,omitempty" type:"Repeated"`
	SeriousList     []*string `json:"SeriousList,omitempty" xml:"SeriousList,omitempty" type:"Repeated"`
}

func (s DescribeSecurityStatInfoResponseBodySecurityEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodySecurityEvent) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSuspiciousCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SuspiciousCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetRemindCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.RemindCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetTimeArray(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.TimeArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetRemindList(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.RemindList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetLevelsOn(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.LevelsOn = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSeriousCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SeriousCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.DateArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSuspiciousList(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SuspiciousList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSeriousList(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SeriousList = v
	return s
}

type DescribeSecurityStatInfoResponseBodyHealthCheck struct {
	HighCount   *int32    `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	LowCount    *int32    `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	ValueArray  []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
	TimeArray   []*string `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Repeated"`
	LevelsOn    []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	LowList     []*string `json:"LowList,omitempty" xml:"LowList,omitempty" type:"Repeated"`
	MediumList  []*string `json:"MediumList,omitempty" xml:"MediumList,omitempty" type:"Repeated"`
	TotalCount  *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MediumCount *int32    `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	DateArray   []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	HighList    []*string `json:"HighList,omitempty" xml:"HighList,omitempty" type:"Repeated"`
}

func (s DescribeSecurityStatInfoResponseBodyHealthCheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodyHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetHighCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.HighCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetLowCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.LowCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetTimeArray(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.TimeArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetLevelsOn(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.LevelsOn = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetLowList(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.LowList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetMediumList(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.MediumList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetMediumCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.MediumCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.DateArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetHighList(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.HighList = v
	return s
}

type DescribeSecurityStatInfoResponseBodyVulnerability struct {
	NntfCount  *int32    `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	NntfList   []*string `json:"NntfList,omitempty" xml:"NntfList,omitempty" type:"Repeated"`
	AsapList   []*string `json:"AsapList,omitempty" xml:"AsapList,omitempty" type:"Repeated"`
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
	TimeArray  []*string `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Repeated"`
	LevelsOn   []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	LaterList  []*string `json:"LaterList,omitempty" xml:"LaterList,omitempty" type:"Repeated"`
	LaterCount *int32    `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DateArray  []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	AsapCount  *int32    `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
}

func (s DescribeSecurityStatInfoResponseBodyVulnerability) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodyVulnerability) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetNntfCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.NntfCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetNntfList(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.NntfList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetAsapList(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.AsapList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetTimeArray(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.TimeArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetLevelsOn(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.LevelsOn = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetLaterList(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.LaterList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetLaterCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.LaterCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.DateArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetAsapCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.AsapCount = &v
	return s
}

type DescribeSecurityStatInfoResponseBodyAttackEvent struct {
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DateArray  []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
}

func (s DescribeSecurityStatInfoResponseBodyAttackEvent) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodyAttackEvent) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodyAttackEvent {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodyAttackEvent {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodyAttackEvent {
	s.DateArray = v
	return s
}

type DescribeSecurityStatInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityStatInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponse) SetHeaders(v map[string]*string) *DescribeSecurityStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityStatInfoResponse) SetBody(v *DescribeSecurityStatInfoResponseBody) *DescribeSecurityStatInfoResponse {
	s.Body = v
	return s
}

type DescribeSimilarEventScenariosRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
}

func (s DescribeSimilarEventScenariosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarEventScenariosRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosRequest) SetSourceIp(v string) *DescribeSimilarEventScenariosRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSimilarEventScenariosRequest) SetResourceOwnerId(v int64) *DescribeSimilarEventScenariosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSimilarEventScenariosRequest) SetSecurityEventId(v int64) *DescribeSimilarEventScenariosRequest {
	s.SecurityEventId = &v
	return s
}

type DescribeSimilarEventScenariosResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scenarios []*DescribeSimilarEventScenariosResponseBodyScenarios `json:"Scenarios,omitempty" xml:"Scenarios,omitempty" type:"Repeated"`
}

func (s DescribeSimilarEventScenariosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarEventScenariosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosResponseBody) SetRequestId(v string) *DescribeSimilarEventScenariosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimilarEventScenariosResponseBody) SetScenarios(v []*DescribeSimilarEventScenariosResponseBodyScenarios) *DescribeSimilarEventScenariosResponseBody {
	s.Scenarios = v
	return s
}

type DescribeSimilarEventScenariosResponseBodyScenarios struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeSimilarEventScenariosResponseBodyScenarios) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarEventScenariosResponseBodyScenarios) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosResponseBodyScenarios) SetCode(v string) *DescribeSimilarEventScenariosResponseBodyScenarios {
	s.Code = &v
	return s
}

type DescribeSimilarEventScenariosResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSimilarEventScenariosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSimilarEventScenariosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarEventScenariosResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimilarEventScenariosResponse) SetHeaders(v map[string]*string) *DescribeSimilarEventScenariosResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimilarEventScenariosResponse) SetBody(v *DescribeSimilarEventScenariosResponseBody) *DescribeSimilarEventScenariosResponse {
	s.Body = v
	return s
}

type DescribeSimilarSecurityEventsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSimilarSecurityEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsRequest) SetSourceIp(v string) *DescribeSimilarSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetResourceOwnerId(v int64) *DescribeSimilarSecurityEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetLang(v string) *DescribeSimilarSecurityEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetTaskId(v int64) *DescribeSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetCurrentPage(v int32) *DescribeSimilarSecurityEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSimilarSecurityEventsRequest) SetPageSize(v int32) *DescribeSimilarSecurityEventsRequest {
	s.PageSize = &v
	return s
}

type DescribeSimilarSecurityEventsResponseBody struct {
	PageInfo               *DescribeSimilarSecurityEventsResponseBodyPageInfo                 `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId              *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityEventsResponse []*DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse `json:"SecurityEventsResponse,omitempty" xml:"SecurityEventsResponse,omitempty" type:"Repeated"`
}

func (s DescribeSimilarSecurityEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponseBody) SetPageInfo(v *DescribeSimilarSecurityEventsResponseBodyPageInfo) *DescribeSimilarSecurityEventsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBody) SetRequestId(v string) *DescribeSimilarSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBody) SetSecurityEventsResponse(v []*DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) *DescribeSimilarSecurityEventsResponseBody {
	s.SecurityEventsResponse = v
	return s
}

type DescribeSimilarSecurityEventsResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSimilarSecurityEventsResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetPageSize(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetCount(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse struct {
	LastTime        *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	EventName       *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	EventType       *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	SecurityEventId *int64  `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	OccurrenceTime  *int64  `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
}

func (s DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetLastTime(v int64) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.LastTime = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetUuid(v string) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.Uuid = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetEventName(v string) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.EventName = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetEventType(v string) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.EventType = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetSecurityEventId(v int64) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetOccurrenceTime(v int64) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.OccurrenceTime = &v
	return s
}

type DescribeSimilarSecurityEventsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSimilarSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSimilarSecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponse) SetHeaders(v map[string]*string) *DescribeSimilarSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimilarSecurityEventsResponse) SetBody(v *DescribeSimilarSecurityEventsResponseBody) *DescribeSimilarSecurityEventsResponse {
	s.Body = v
	return s
}

type DescribeStrategyExecDetailRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StrategyId *int32  `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeStrategyExecDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStrategyExecDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailRequest) SetSourceIp(v string) *DescribeStrategyExecDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeStrategyExecDetailRequest) SetStrategyId(v int32) *DescribeStrategyExecDetailRequest {
	s.StrategyId = &v
	return s
}

type DescribeStrategyExecDetailResponseBody struct {
	EndTime        *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Percent        *string                                                `json:"Percent,omitempty" xml:"Percent,omitempty"`
	FailCount      *int32                                                 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	StartTime      *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SuccessCount   *int32                                                 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	Source         *string                                                `json:"Source,omitempty" xml:"Source,omitempty"`
	FailedEcsList  []*DescribeStrategyExecDetailResponseBodyFailedEcsList `json:"FailedEcsList,omitempty" xml:"FailedEcsList,omitempty" type:"Repeated"`
	InProcessCount *int32                                                 `json:"InProcessCount,omitempty" xml:"InProcessCount,omitempty"`
}

func (s DescribeStrategyExecDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStrategyExecDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailResponseBody) SetEndTime(v string) *DescribeStrategyExecDetailResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetRequestId(v string) *DescribeStrategyExecDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetPercent(v string) *DescribeStrategyExecDetailResponseBody {
	s.Percent = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetFailCount(v int32) *DescribeStrategyExecDetailResponseBody {
	s.FailCount = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetStartTime(v string) *DescribeStrategyExecDetailResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetSuccessCount(v int32) *DescribeStrategyExecDetailResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetSource(v string) *DescribeStrategyExecDetailResponseBody {
	s.Source = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetFailedEcsList(v []*DescribeStrategyExecDetailResponseBodyFailedEcsList) *DescribeStrategyExecDetailResponseBody {
	s.FailedEcsList = v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetInProcessCount(v int32) *DescribeStrategyExecDetailResponseBody {
	s.InProcessCount = &v
	return s
}

type DescribeStrategyExecDetailResponseBodyFailedEcsList struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	IntranetIp   *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeStrategyExecDetailResponseBodyFailedEcsList) String() string {
	return tea.Prettify(s)
}

func (s DescribeStrategyExecDetailResponseBodyFailedEcsList) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetInstanceName(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.InstanceName = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetIP(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.IP = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetIntranetIp(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetReason(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.Reason = &v
	return s
}

type DescribeStrategyExecDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStrategyExecDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStrategyExecDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStrategyExecDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailResponse) SetHeaders(v map[string]*string) *DescribeStrategyExecDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeStrategyExecDetailResponse) SetBody(v *DescribeStrategyExecDetailResponseBody) *DescribeStrategyExecDetailResponse {
	s.Body = v
	return s
}

type DescribeStratetyRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StrategyIds *string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty"`
}

func (s DescribeStratetyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeStratetyRequest) GoString() string {
	return s.String()
}

func (s *DescribeStratetyRequest) SetSourceIp(v string) *DescribeStratetyRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeStratetyRequest) SetLang(v string) *DescribeStratetyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeStratetyRequest) SetStrategyIds(v string) *DescribeStratetyRequest {
	s.StrategyIds = &v
	return s
}

type DescribeStratetyResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Strategies []*DescribeStratetyResponseBodyStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Repeated"`
}

func (s DescribeStratetyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStratetyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStratetyResponseBody) SetRequestId(v string) *DescribeStratetyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStratetyResponseBody) SetStrategies(v []*DescribeStratetyResponseBodyStrategies) *DescribeStratetyResponseBody {
	s.Strategies = v
	return s
}

type DescribeStratetyResponseBodyStrategies struct {
	ExecStatus     *int32                                                 `json:"ExecStatus,omitempty" xml:"ExecStatus,omitempty"`
	Type           *int32                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	ConfigTargets  []*DescribeStratetyResponseBodyStrategiesConfigTargets `json:"ConfigTargets,omitempty" xml:"ConfigTargets,omitempty" type:"Repeated"`
	CycleStartTime *int32                                                 `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	EcsCount       *int32                                                 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	ProcessRate    *int32                                                 `json:"ProcessRate,omitempty" xml:"ProcessRate,omitempty"`
	CycleDays      *int32                                                 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty"`
	RiskCount      *int32                                                 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	Name           *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PassRate       *int32                                                 `json:"PassRate,omitempty" xml:"PassRate,omitempty"`
	Id             *int32                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeStratetyResponseBodyStrategies) String() string {
	return tea.Prettify(s)
}

func (s DescribeStratetyResponseBodyStrategies) GoString() string {
	return s.String()
}

func (s *DescribeStratetyResponseBodyStrategies) SetExecStatus(v int32) *DescribeStratetyResponseBodyStrategies {
	s.ExecStatus = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetType(v int32) *DescribeStratetyResponseBodyStrategies {
	s.Type = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetConfigTargets(v []*DescribeStratetyResponseBodyStrategiesConfigTargets) *DescribeStratetyResponseBodyStrategies {
	s.ConfigTargets = v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetCycleStartTime(v int32) *DescribeStratetyResponseBodyStrategies {
	s.CycleStartTime = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetEcsCount(v int32) *DescribeStratetyResponseBodyStrategies {
	s.EcsCount = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetProcessRate(v int32) *DescribeStratetyResponseBodyStrategies {
	s.ProcessRate = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetCycleDays(v int32) *DescribeStratetyResponseBodyStrategies {
	s.CycleDays = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetRiskCount(v int32) *DescribeStratetyResponseBodyStrategies {
	s.RiskCount = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetName(v string) *DescribeStratetyResponseBodyStrategies {
	s.Name = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetPassRate(v int32) *DescribeStratetyResponseBodyStrategies {
	s.PassRate = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategies) SetId(v int32) *DescribeStratetyResponseBodyStrategies {
	s.Id = &v
	return s
}

type DescribeStratetyResponseBodyStrategiesConfigTargets struct {
	Flag       *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	Target     *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeStratetyResponseBodyStrategiesConfigTargets) String() string {
	return tea.Prettify(s)
}

func (s DescribeStratetyResponseBodyStrategiesConfigTargets) GoString() string {
	return s.String()
}

func (s *DescribeStratetyResponseBodyStrategiesConfigTargets) SetFlag(v string) *DescribeStratetyResponseBodyStrategiesConfigTargets {
	s.Flag = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategiesConfigTargets) SetTarget(v string) *DescribeStratetyResponseBodyStrategiesConfigTargets {
	s.Target = &v
	return s
}

func (s *DescribeStratetyResponseBodyStrategiesConfigTargets) SetTargetType(v string) *DescribeStratetyResponseBodyStrategiesConfigTargets {
	s.TargetType = &v
	return s
}

type DescribeStratetyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeStratetyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStratetyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStratetyResponse) GoString() string {
	return s.String()
}

func (s *DescribeStratetyResponse) SetHeaders(v map[string]*string) *DescribeStratetyResponse {
	s.Headers = v
	return s
}

func (s *DescribeStratetyResponse) SetBody(v *DescribeStratetyResponseBody) *DescribeStratetyResponse {
	s.Body = v
	return s
}

type DescribeSummaryInfoRequest struct {
	SourceIp                   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s DescribeSummaryInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSummaryInfoRequest) SetSourceIp(v string) *DescribeSummaryInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSummaryInfoRequest) SetLang(v string) *DescribeSummaryInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSummaryInfoRequest) SetResourceDirectoryAccountId(v string) *DescribeSummaryInfoRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

type DescribeSummaryInfoResponseBody struct {
	AegisClientOnlineCount  *int32  `json:"AegisClientOnlineCount,omitempty" xml:"AegisClientOnlineCount,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AegisClientOfflineCount *int32  `json:"AegisClientOfflineCount,omitempty" xml:"AegisClientOfflineCount,omitempty"`
	SecurityScore           *int32  `json:"SecurityScore,omitempty" xml:"SecurityScore,omitempty"`
	Success                 *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSummaryInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSummaryInfoResponseBody) SetAegisClientOnlineCount(v int32) *DescribeSummaryInfoResponseBody {
	s.AegisClientOnlineCount = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetRequestId(v string) *DescribeSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetAegisClientOfflineCount(v int32) *DescribeSummaryInfoResponseBody {
	s.AegisClientOfflineCount = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetSecurityScore(v int32) *DescribeSummaryInfoResponseBody {
	s.SecurityScore = &v
	return s
}

func (s *DescribeSummaryInfoResponseBody) SetSuccess(v bool) *DescribeSummaryInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeSummaryInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSummaryInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSummaryInfoResponse) SetHeaders(v map[string]*string) *DescribeSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSummaryInfoResponse) SetBody(v *DescribeSummaryInfoResponseBody) *DescribeSummaryInfoResponse {
	s.Body = v
	return s
}

type DescribeSuspEventDetailRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SuspiciousEventId *int32  `json:"SuspiciousEventId,omitempty" xml:"SuspiciousEventId,omitempty"`
	From              *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s DescribeSuspEventDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailRequest) SetSourceIp(v string) *DescribeSuspEventDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) SetLang(v string) *DescribeSuspEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) SetSuspiciousEventId(v int32) *DescribeSuspEventDetailRequest {
	s.SuspiciousEventId = &v
	return s
}

func (s *DescribeSuspEventDetailRequest) SetFrom(v string) *DescribeSuspEventDetailRequest {
	s.From = &v
	return s
}

type DescribeSuspEventDetailResponseBody struct {
	EventDesc        *string                                       `json:"EventDesc,omitempty" xml:"EventDesc,omitempty"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EventTypeDesc    *string                                       `json:"EventTypeDesc,omitempty" xml:"EventTypeDesc,omitempty"`
	OperateErrorCode *string                                       `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	EventStatus      *string                                       `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	EventName        *string                                       `json:"EventName,omitempty" xml:"EventName,omitempty"`
	SaleVersion      *string                                       `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	IntranetIp       *string                                       `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	DataSource       *string                                       `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	InstanceName     *string                                       `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OperateMsg       *string                                       `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty"`
	CanBeDealOnLine  *bool                                         `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	Uuid             *string                                       `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Details          []*DescribeSuspEventDetailResponseBodyDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	InternetIp       *string                                       `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Level            *string                                       `json:"Level,omitempty" xml:"Level,omitempty"`
	Id               *int32                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	LastTime         *string                                       `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	SasId            *string                                       `json:"SasId,omitempty" xml:"SasId,omitempty"`
}

func (s DescribeSuspEventDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailResponseBody) SetEventDesc(v string) *DescribeSuspEventDetailResponseBody {
	s.EventDesc = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetRequestId(v string) *DescribeSuspEventDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetEventTypeDesc(v string) *DescribeSuspEventDetailResponseBody {
	s.EventTypeDesc = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetOperateErrorCode(v string) *DescribeSuspEventDetailResponseBody {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetEventStatus(v string) *DescribeSuspEventDetailResponseBody {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetEventName(v string) *DescribeSuspEventDetailResponseBody {
	s.EventName = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetSaleVersion(v string) *DescribeSuspEventDetailResponseBody {
	s.SaleVersion = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetIntranetIp(v string) *DescribeSuspEventDetailResponseBody {
	s.IntranetIp = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetDataSource(v string) *DescribeSuspEventDetailResponseBody {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetInstanceName(v string) *DescribeSuspEventDetailResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetOperateMsg(v string) *DescribeSuspEventDetailResponseBody {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetCanBeDealOnLine(v bool) *DescribeSuspEventDetailResponseBody {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetUuid(v string) *DescribeSuspEventDetailResponseBody {
	s.Uuid = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetDetails(v []*DescribeSuspEventDetailResponseBodyDetails) *DescribeSuspEventDetailResponseBody {
	s.Details = v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetInternetIp(v string) *DescribeSuspEventDetailResponseBody {
	s.InternetIp = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetLevel(v string) *DescribeSuspEventDetailResponseBody {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetId(v int32) *DescribeSuspEventDetailResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetLastTime(v string) *DescribeSuspEventDetailResponseBody {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBody) SetSasId(v string) *DescribeSuspEventDetailResponseBody {
	s.SasId = &v
	return s
}

type DescribeSuspEventDetailResponseBodyDetails struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
}

func (s DescribeSuspEventDetailResponseBodyDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventDetailResponseBodyDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailResponseBodyDetails) SetType(v string) *DescribeSuspEventDetailResponseBodyDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBodyDetails) SetValue(v string) *DescribeSuspEventDetailResponseBodyDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBodyDetails) SetName(v string) *DescribeSuspEventDetailResponseBodyDetails {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventDetailResponseBodyDetails) SetInfoType(v string) *DescribeSuspEventDetailResponseBodyDetails {
	s.InfoType = &v
	return s
}

type DescribeSuspEventDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventDetailResponse) SetHeaders(v map[string]*string) *DescribeSuspEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventDetailResponse) SetBody(v *DescribeSuspEventDetailResponseBody) *DescribeSuspEventDetailResponse {
	s.Body = v
	return s
}

type DescribeSuspEventQuaraFilesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	QuaraTag    *string `json:"QuaraTag,omitempty" xml:"QuaraTag,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	From        *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s DescribeSuspEventQuaraFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesRequest) SetSourceIp(v string) *DescribeSuspEventQuaraFilesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetPageSize(v string) *DescribeSuspEventQuaraFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetStatus(v string) *DescribeSuspEventQuaraFilesRequest {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetQuaraTag(v string) *DescribeSuspEventQuaraFilesRequest {
	s.QuaraTag = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetCurrentPage(v string) *DescribeSuspEventQuaraFilesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesRequest) SetFrom(v string) *DescribeSuspEventQuaraFilesRequest {
	s.From = &v
	return s
}

type DescribeSuspEventQuaraFilesResponseBody struct {
	TotalCount  *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                               `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	QuaraFiles  []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles `json:"QuaraFiles,omitempty" xml:"QuaraFiles,omitempty" type:"Repeated"`
	Count       *int32                                               `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetTotalCount(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetPageSize(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetRequestId(v string) *DescribeSuspEventQuaraFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetCurrentPage(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetQuaraFiles(v []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) *DescribeSuspEventQuaraFilesResponseBody {
	s.QuaraFiles = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetCount(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.Count = &v
	return s
}

type DescribeSuspEventQuaraFilesResponseBodyQuaraFiles struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	EventName    *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	EventType    *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Md5          *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime   *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetStatus(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetInternetIp(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.InternetIp = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetIp(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Ip = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetTag(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Tag = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetUuid(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Uuid = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventType(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventType = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetInstanceName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.InstanceName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetPath(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Path = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetMd5(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Md5 = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetId(v int32) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetModifyTime(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.ModifyTime = &v
	return s
}

type DescribeSuspEventQuaraFilesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventQuaraFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventQuaraFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponse) SetHeaders(v map[string]*string) *DescribeSuspEventQuaraFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetBody(v *DescribeSuspEventQuaraFilesResponseBody) *DescribeSuspEventQuaraFilesResponse {
	s.Body = v
	return s
}

type DescribeSuspEventsRequest struct {
	SourceIp             *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Dealed               *string   `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	Name                 *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Levels               *string   `json:"Levels,omitempty" xml:"Levels,omitempty"`
	ParentEventTypes     *string   `json:"ParentEventTypes,omitempty" xml:"ParentEventTypes,omitempty"`
	Remark               *string   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	PageSize             *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage          *string   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Lang                 *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AlarmUniqueInfo      *string   `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	UniqueInfo           *string   `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	From                 *string   `json:"From,omitempty" xml:"From,omitempty"`
	Source               *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	GroupId              *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Uuids                *string   `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	ClusterId            *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerFieldName   *string   `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	ContainerFieldValue  *string   `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	TargetType           *string   `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	OperateErrorCodeList []*string `json:"OperateErrorCodeList,omitempty" xml:"OperateErrorCodeList,omitempty" type:"Repeated"`
}

func (s DescribeSuspEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsRequest) SetSourceIp(v string) *DescribeSuspEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetDealed(v string) *DescribeSuspEventsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetName(v string) *DescribeSuspEventsRequest {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLevels(v string) *DescribeSuspEventsRequest {
	s.Levels = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetParentEventTypes(v string) *DescribeSuspEventsRequest {
	s.ParentEventTypes = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetRemark(v string) *DescribeSuspEventsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetStatus(v string) *DescribeSuspEventsRequest {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetPageSize(v string) *DescribeSuspEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetCurrentPage(v string) *DescribeSuspEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetLang(v string) *DescribeSuspEventsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetAlarmUniqueInfo(v string) *DescribeSuspEventsRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetUniqueInfo(v string) *DescribeSuspEventsRequest {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetFrom(v string) *DescribeSuspEventsRequest {
	s.From = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetSource(v string) *DescribeSuspEventsRequest {
	s.Source = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetGroupId(v int64) *DescribeSuspEventsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetUuids(v string) *DescribeSuspEventsRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetClusterId(v string) *DescribeSuspEventsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetContainerFieldName(v string) *DescribeSuspEventsRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetContainerFieldValue(v string) *DescribeSuspEventsRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetTargetType(v string) *DescribeSuspEventsRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeSuspEventsRequest) SetOperateErrorCodeList(v []*string) *DescribeSuspEventsRequest {
	s.OperateErrorCodeList = v
	return s
}

type DescribeSuspEventsResponseBody struct {
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Count       *int32                                      `json:"Count,omitempty" xml:"Count,omitempty"`
	SuspEvents  []*DescribeSuspEventsResponseBodySuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" type:"Repeated"`
}

func (s DescribeSuspEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBody) SetTotalCount(v int32) *DescribeSuspEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetPageSize(v int32) *DescribeSuspEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetRequestId(v string) *DescribeSuspEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetCurrentPage(v int32) *DescribeSuspEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetCount(v int32) *DescribeSuspEventsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetSuspEvents(v []*DescribeSuspEventsResponseBodySuspEvents) *DescribeSuspEventsResponseBody {
	s.SuspEvents = v
	return s
}

type DescribeSuspEventsResponseBodySuspEvents struct {
	Stages                *string                                            `json:"Stages,omitempty" xml:"Stages,omitempty"`
	InternetIp            *string                                            `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	K8sClusterName        *string                                            `json:"K8sClusterName,omitempty" xml:"K8sClusterName,omitempty"`
	ContainerImageId      *string                                            `json:"ContainerImageId,omitempty" xml:"ContainerImageId,omitempty"`
	LastTimeStamp         *int64                                             `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	OccurrenceTime        *string                                            `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	AlarmUniqueInfo       *string                                            `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	Desc                  *string                                            `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Details               []*DescribeSuspEventsResponseBodySuspEventsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	CanCancelFault        *bool                                              `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	AlarmEventNameDisplay *string                                            `json:"AlarmEventNameDisplay,omitempty" xml:"AlarmEventNameDisplay,omitempty"`
	AppName               *string                                            `json:"AppName,omitempty" xml:"AppName,omitempty"`
	SecurityEventIds      *string                                            `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty"`
	CanBeDealOnLine       *bool                                              `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	MarkMisRules          *string                                            `json:"MarkMisRules,omitempty" xml:"MarkMisRules,omitempty"`
	ContainerImageName    *string                                            `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	K8sClusterId          *string                                            `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ContainHwMode         *bool                                              `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	InstanceName          *string                                            `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	K8sNodeId             *string                                            `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	EventStatus           *int32                                             `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	SaleVersion           *string                                            `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	OperateErrorCode      *string                                            `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	Name                  *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	HasTraceInfo          *bool                                              `json:"HasTraceInfo,omitempty" xml:"HasTraceInfo,omitempty"`
	DataSource            *string                                            `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	OperateTime           *int64                                             `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	EventSubType          *string                                            `json:"EventSubType,omitempty" xml:"EventSubType,omitempty"`
	Advanced              *bool                                              `json:"Advanced,omitempty" xml:"Advanced,omitempty"`
	OccurrenceTimeStamp   *int64                                             `json:"OccurrenceTimeStamp,omitempty" xml:"OccurrenceTimeStamp,omitempty"`
	AlarmEventTypeDisplay *string                                            `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty"`
	IntranetIp            *string                                            `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	LastTime              *string                                            `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	OperateMsg            *string                                            `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty"`
	Uuid                  *string                                            `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	K8sPodName            *string                                            `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	ContainerId           *string                                            `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	AlarmEventType        *string                                            `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	K8sNamespace          *string                                            `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	AutoBreaking          *bool                                              `json:"AutoBreaking,omitempty" xml:"AutoBreaking,omitempty"`
	K8sNodeName           *string                                            `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	AlarmEventName        *string                                            `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	UniqueInfo            *string                                            `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	Level                 *string                                            `json:"Level,omitempty" xml:"Level,omitempty"`
	Id                    *int64                                             `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetStages(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Stages = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetInternetIp(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.InternetIp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sClusterName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sClusterName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainerImageId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainerImageId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLastTimeStamp(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.LastTimeStamp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOccurrenceTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesc(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Desc = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDetails(v []*DescribeSuspEventsResponseBodySuspEventsDetails) *DescribeSuspEventsResponseBodySuspEvents {
	s.Details = v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanCancelFault(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventNameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAppName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AppName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetSecurityEventIds(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.SecurityEventIds = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanBeDealOnLine(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetMarkMisRules(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.MarkMisRules = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainerImageName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainerImageName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sClusterId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainHwMode(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainHwMode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetInstanceName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.InstanceName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sNodeId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventStatus(v int32) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetSaleVersion(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.SaleVersion = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateErrorCode(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetHasTraceInfo(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.HasTraceInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDataSource(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateTime(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventSubType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventSubType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAdvanced(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.Advanced = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOccurrenceTimeStamp(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.OccurrenceTimeStamp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventTypeDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetIntranetIp(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.IntranetIp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLastTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateMsg(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetUuid(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Uuid = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sPodName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sPodName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainerId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainerId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sNamespace(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAutoBreaking(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.AutoBreaking = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sNodeName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLevel(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetId(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.Id = &v
	return s
}

type DescribeSuspEventsResponseBodySuspEventsDetails struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NameDisplay  *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	InfoType     *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetType(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValue(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetName(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.NameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetInfoType(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.InfoType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValueDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.ValueDisplay = &v
	return s
}

type DescribeSuspEventsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSuspEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSuspEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSuspEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponse) SetHeaders(v map[string]*string) *DescribeSuspEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventsResponse) SetBody(v *DescribeSuspEventsResponseBody) *DescribeSuspEventsResponse {
	s.Body = v
	return s
}

type DescribeUserBaselineAuthorizationRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserBaselineAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationRequest) SetSourceIp(v string) *DescribeUserBaselineAuthorizationRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationRequest) SetResourceOwnerId(v int64) *DescribeUserBaselineAuthorizationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationRequest) SetLang(v string) *DescribeUserBaselineAuthorizationRequest {
	s.Lang = &v
	return s
}

type DescribeUserBaselineAuthorizationResponseBody struct {
	RequestId                 *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserBaselineAuthorization *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization `json:"UserBaselineAuthorization,omitempty" xml:"UserBaselineAuthorization,omitempty" type:"Struct"`
}

func (s DescribeUserBaselineAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationResponseBody) SetRequestId(v string) *DescribeUserBaselineAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponseBody) SetUserBaselineAuthorization(v *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) *DescribeUserBaselineAuthorizationResponseBody {
	s.UserBaselineAuthorization = v
	return s
}

type DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization struct {
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization) SetStatus(v int32) *DescribeUserBaselineAuthorizationResponseBodyUserBaselineAuthorization {
	s.Status = &v
	return s
}

type DescribeUserBaselineAuthorizationResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserBaselineAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserBaselineAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserBaselineAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBaselineAuthorizationResponse) SetHeaders(v map[string]*string) *DescribeUserBaselineAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBaselineAuthorizationResponse) SetBody(v *DescribeUserBaselineAuthorizationResponseBody) *DescribeUserBaselineAuthorizationResponse {
	s.Body = v
	return s
}

type DescribeUserLayoutAuthorizationRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserLayoutAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLayoutAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserLayoutAuthorizationRequest) SetSourceIp(v string) *DescribeUserLayoutAuthorizationRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserLayoutAuthorizationRequest) SetResourceOwnerId(v int64) *DescribeUserLayoutAuthorizationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserLayoutAuthorizationRequest) SetLang(v string) *DescribeUserLayoutAuthorizationRequest {
	s.Lang = &v
	return s
}

type DescribeUserLayoutAuthorizationResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Authorized *bool   `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
}

func (s DescribeUserLayoutAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLayoutAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLayoutAuthorizationResponseBody) SetRequestId(v string) *DescribeUserLayoutAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserLayoutAuthorizationResponseBody) SetAuthorized(v bool) *DescribeUserLayoutAuthorizationResponseBody {
	s.Authorized = &v
	return s
}

type DescribeUserLayoutAuthorizationResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserLayoutAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserLayoutAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserLayoutAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLayoutAuthorizationResponse) SetHeaders(v map[string]*string) *DescribeUserLayoutAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLayoutAuthorizationResponse) SetBody(v *DescribeUserLayoutAuthorizationResponseBody) *DescribeUserLayoutAuthorizationResponse {
	s.Body = v
	return s
}

type DescribeVersionConfigRequest struct {
	SourceIp                   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceDirectoryAccountId *string `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s DescribeVersionConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigRequest) SetSourceIp(v string) *DescribeVersionConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVersionConfigRequest) SetResourceDirectoryAccountId(v string) *DescribeVersionConfigRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

type DescribeVersionConfigResponseBody struct {
	ImageScanCapacity     *int64  `json:"ImageScanCapacity,omitempty" xml:"ImageScanCapacity,omitempty"`
	AppWhiteListAuthCount *int64  `json:"AppWhiteListAuthCount,omitempty" xml:"AppWhiteListAuthCount,omitempty"`
	SasLog                *int32  `json:"SasLog,omitempty" xml:"SasLog,omitempty"`
	Version               *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	LastTrailEndTime      *int64  `json:"LastTrailEndTime,omitempty" xml:"LastTrailEndTime,omitempty"`
	WebLockAuthCount      *int64  `json:"WebLockAuthCount,omitempty" xml:"WebLockAuthCount,omitempty"`
	SlsCapacity           *int64  `json:"SlsCapacity,omitempty" xml:"SlsCapacity,omitempty"`
	UserDefinedAlarms     *int32  `json:"UserDefinedAlarms,omitempty" xml:"UserDefinedAlarms,omitempty"`
	WebLock               *int32  `json:"WebLock,omitempty" xml:"WebLock,omitempty"`
	IsOverBalance         *bool   `json:"IsOverBalance,omitempty" xml:"IsOverBalance,omitempty"`
	VmCores               *int32  `json:"VmCores,omitempty" xml:"VmCores,omitempty"`
	HoneypotCapacity      *int64  `json:"HoneypotCapacity,omitempty" xml:"HoneypotCapacity,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AssetLevel            *int32  `json:"AssetLevel,omitempty" xml:"AssetLevel,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SasScreen             *int32  `json:"SasScreen,omitempty" xml:"SasScreen,omitempty"`
	ReleaseTime           *int64  `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	IsTrialVersion        *int32  `json:"IsTrialVersion,omitempty" xml:"IsTrialVersion,omitempty"`
	AppWhiteList          *int32  `json:"AppWhiteList,omitempty" xml:"AppWhiteList,omitempty"`
}

func (s DescribeVersionConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponseBody) SetImageScanCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.ImageScanCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAppWhiteListAuthCount(v int64) *DescribeVersionConfigResponseBody {
	s.AppWhiteListAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasLog(v int32) *DescribeVersionConfigResponseBody {
	s.SasLog = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetVersion(v int32) *DescribeVersionConfigResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetLastTrailEndTime(v int64) *DescribeVersionConfigResponseBody {
	s.LastTrailEndTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetWebLockAuthCount(v int64) *DescribeVersionConfigResponseBody {
	s.WebLockAuthCount = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSlsCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.SlsCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetUserDefinedAlarms(v int32) *DescribeVersionConfigResponseBody {
	s.UserDefinedAlarms = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetWebLock(v int32) *DescribeVersionConfigResponseBody {
	s.WebLock = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIsOverBalance(v bool) *DescribeVersionConfigResponseBody {
	s.IsOverBalance = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetVmCores(v int32) *DescribeVersionConfigResponseBody {
	s.VmCores = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetHoneypotCapacity(v int64) *DescribeVersionConfigResponseBody {
	s.HoneypotCapacity = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetRequestId(v string) *DescribeVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAssetLevel(v int32) *DescribeVersionConfigResponseBody {
	s.AssetLevel = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetInstanceId(v string) *DescribeVersionConfigResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetSasScreen(v int32) *DescribeVersionConfigResponseBody {
	s.SasScreen = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetReleaseTime(v int64) *DescribeVersionConfigResponseBody {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetIsTrialVersion(v int32) *DescribeVersionConfigResponseBody {
	s.IsTrialVersion = &v
	return s
}

func (s *DescribeVersionConfigResponseBody) SetAppWhiteList(v int32) *DescribeVersionConfigResponseBody {
	s.AppWhiteList = &v
	return s
}

type DescribeVersionConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVersionConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponse) SetHeaders(v map[string]*string) *DescribeVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVersionConfigResponse) SetBody(v *DescribeVersionConfigResponseBody) *DescribeVersionConfigResponse {
	s.Body = v
	return s
}

type DescribeVolDingdingMessageRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeVolDingdingMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVolDingdingMessageRequest) GoString() string {
	return s.String()
}

func (s *DescribeVolDingdingMessageRequest) SetSourceIp(v string) *DescribeVolDingdingMessageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVolDingdingMessageRequest) SetLang(v string) *DescribeVolDingdingMessageRequest {
	s.Lang = &v
	return s
}

type DescribeVolDingdingMessageResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DingdingUrl *string `json:"DingdingUrl,omitempty" xml:"DingdingUrl,omitempty"`
}

func (s DescribeVolDingdingMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVolDingdingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVolDingdingMessageResponseBody) SetRequestId(v string) *DescribeVolDingdingMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVolDingdingMessageResponseBody) SetDingdingUrl(v string) *DescribeVolDingdingMessageResponseBody {
	s.DingdingUrl = &v
	return s
}

type DescribeVolDingdingMessageResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVolDingdingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVolDingdingMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVolDingdingMessageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVolDingdingMessageResponse) SetHeaders(v map[string]*string) *DescribeVolDingdingMessageResponse {
	s.Headers = v
	return s
}

func (s *DescribeVolDingdingMessageResponse) SetBody(v *DescribeVolDingdingMessageResponseBody) *DescribeVolDingdingMessageResponse {
	s.Body = v
	return s
}

type DescribeVpcHoneyPotCriteriaRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeVpcHoneyPotCriteriaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaRequest) SetSourceIp(v string) *DescribeVpcHoneyPotCriteriaRequest {
	s.SourceIp = &v
	return s
}

type DescribeVpcHoneyPotCriteriaResponseBody struct {
	CriteriaList []*DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList `json:"CriteriaList,omitempty" xml:"CriteriaList,omitempty" type:"Repeated"`
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpcHoneyPotCriteriaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaResponseBody) SetCriteriaList(v []*DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) *DescribeVpcHoneyPotCriteriaResponseBody {
	s.CriteriaList = v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBody) SetRequestId(v string) *DescribeVpcHoneyPotCriteriaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList struct {
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) SetType(v string) *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList {
	s.Type = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) SetValues(v string) *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList {
	s.Values = &v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList) SetName(v string) *DescribeVpcHoneyPotCriteriaResponseBodyCriteriaList {
	s.Name = &v
	return s
}

type DescribeVpcHoneyPotCriteriaResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVpcHoneyPotCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcHoneyPotCriteriaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotCriteriaResponse) SetHeaders(v map[string]*string) *DescribeVpcHoneyPotCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcHoneyPotCriteriaResponse) SetBody(v *DescribeVpcHoneyPotCriteriaResponseBody) *DescribeVpcHoneyPotCriteriaResponse {
	s.Body = v
	return s
}

type DescribeVpcHoneyPotListRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName           *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	VpcRegionId       *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	HoneyPotExistence *bool   `json:"HoneyPotExistence,omitempty" xml:"HoneyPotExistence,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage       *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeVpcHoneyPotListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListRequest) SetSourceIp(v string) *DescribeVpcHoneyPotListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetVpcId(v string) *DescribeVpcHoneyPotListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetVpcName(v string) *DescribeVpcHoneyPotListRequest {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetVpcRegionId(v string) *DescribeVpcHoneyPotListRequest {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetHoneyPotExistence(v bool) *DescribeVpcHoneyPotListRequest {
	s.HoneyPotExistence = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetPageSize(v int32) *DescribeVpcHoneyPotListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcHoneyPotListRequest) SetCurrentPage(v int32) *DescribeVpcHoneyPotListRequest {
	s.CurrentPage = &v
	return s
}

type DescribeVpcHoneyPotListResponseBody struct {
	VpcHoneyPotDTOList []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList `json:"VpcHoneyPotDTOList,omitempty" xml:"VpcHoneyPotDTOList,omitempty" type:"Repeated"`
	PageInfo           *DescribeVpcHoneyPotListResponseBodyPageInfo             `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId          *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpcHoneyPotListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponseBody) SetVpcHoneyPotDTOList(v []*DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) *DescribeVpcHoneyPotListResponseBody {
	s.VpcHoneyPotDTOList = v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBody) SetPageInfo(v *DescribeVpcHoneyPotListResponseBodyPageInfo) *DescribeVpcHoneyPotListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBody) SetRequestId(v string) *DescribeVpcHoneyPotListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList struct {
	VpcId                     *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcStatus                 *string   `json:"VpcStatus,omitempty" xml:"VpcStatus,omitempty"`
	CreateTime                *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	HoneyPotExistence         *bool     `json:"HoneyPotExistence,omitempty" xml:"HoneyPotExistence,omitempty"`
	VpcRegionId               *string   `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
	HoneyPotInstanceStatus    *string   `json:"HoneyPotInstanceStatus,omitempty" xml:"HoneyPotInstanceStatus,omitempty"`
	HoneyPotVSwitchId         *string   `json:"HoneyPotVSwitchId,omitempty" xml:"HoneyPotVSwitchId,omitempty"`
	VpcSwitchIdList           []*string `json:"VpcSwitchIdList,omitempty" xml:"VpcSwitchIdList,omitempty" type:"Repeated"`
	VpcName                   *string   `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	HoneyPotEniInstanceId     *string   `json:"HoneyPotEniInstanceId,omitempty" xml:"HoneyPotEniInstanceId,omitempty"`
	CidrBlock                 *string   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	HoneyPotEcsInstanceStatus *string   `json:"HoneyPotEcsInstanceStatus,omitempty" xml:"HoneyPotEcsInstanceStatus,omitempty"`
}

func (s DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcStatus(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcStatus = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetCreateTime(v int64) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotExistence(v bool) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotExistence = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcRegionId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcRegionId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotInstanceStatus(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotInstanceStatus = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotVSwitchId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotVSwitchId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcSwitchIdList(v []*string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcSwitchIdList = v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetVpcName(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotEniInstanceId(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotEniInstanceId = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetCidrBlock(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList) SetHoneyPotEcsInstanceStatus(v string) *DescribeVpcHoneyPotListResponseBodyVpcHoneyPotDTOList {
	s.HoneyPotEcsInstanceStatus = &v
	return s
}

type DescribeVpcHoneyPotListResponseBodyPageInfo struct {
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Count       *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeVpcHoneyPotListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetPageSize(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcHoneyPotListResponseBodyPageInfo) SetCount(v int32) *DescribeVpcHoneyPotListResponseBodyPageInfo {
	s.Count = &v
	return s
}

type DescribeVpcHoneyPotListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVpcHoneyPotListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcHoneyPotListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcHoneyPotListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcHoneyPotListResponse) SetHeaders(v map[string]*string) *DescribeVpcHoneyPotListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcHoneyPotListResponse) SetBody(v *DescribeVpcHoneyPotListResponseBody) *DescribeVpcHoneyPotListResponse {
	s.Body = v
	return s
}

type DescribeVpcListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeVpcListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcListRequest) SetSourceIp(v string) *DescribeVpcListRequest {
	s.SourceIp = &v
	return s
}

type DescribeVpcListResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcList   []*DescribeVpcListResponseBodyVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
	Count     *int32                                `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeVpcListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcListResponseBody) SetRequestId(v string) *DescribeVpcListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcListResponseBody) SetVpcList(v []*DescribeVpcListResponseBodyVpcList) *DescribeVpcListResponseBody {
	s.VpcList = v
	return s
}

func (s *DescribeVpcListResponseBody) SetCount(v int32) *DescribeVpcListResponseBody {
	s.Count = &v
	return s
}

type DescribeVpcListResponseBodyVpcList struct {
	InstanceDesc *string `json:"InstanceDesc,omitempty" xml:"InstanceDesc,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	EcsCount     *int32  `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVpcListResponseBodyVpcList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListResponseBodyVpcList) GoString() string {
	return s.String()
}

func (s *DescribeVpcListResponseBodyVpcList) SetInstanceDesc(v string) *DescribeVpcListResponseBodyVpcList {
	s.InstanceDesc = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetInstanceName(v string) *DescribeVpcListResponseBodyVpcList {
	s.InstanceName = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetEcsCount(v int32) *DescribeVpcListResponseBodyVpcList {
	s.EcsCount = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetInstanceId(v string) *DescribeVpcListResponseBodyVpcList {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetRegionId(v string) *DescribeVpcListResponseBodyVpcList {
	s.RegionId = &v
	return s
}

type DescribeVpcListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVpcListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcListResponse) SetHeaders(v map[string]*string) *DescribeVpcListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcListResponse) SetBody(v *DescribeVpcListResponseBody) *DescribeVpcListResponse {
	s.Body = v
	return s
}

type DescribeVulDetailsRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
}

func (s DescribeVulDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsRequest) SetSourceIp(v string) *DescribeVulDetailsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetLang(v string) *DescribeVulDetailsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetType(v string) *DescribeVulDetailsRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetName(v string) *DescribeVulDetailsRequest {
	s.Name = &v
	return s
}

func (s *DescribeVulDetailsRequest) SetAliasName(v string) *DescribeVulDetailsRequest {
	s.AliasName = &v
	return s
}

type DescribeVulDetailsResponseBody struct {
	Cves      []*DescribeVulDetailsResponseBodyCves `json:"Cves,omitempty" xml:"Cves,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVulDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBody) SetCves(v []*DescribeVulDetailsResponseBodyCves) *DescribeVulDetailsResponseBody {
	s.Cves = v
	return s
}

func (s *DescribeVulDetailsResponseBody) SetRequestId(v string) *DescribeVulDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVulDetailsResponseBodyCves struct {
	Complexity        *string                                        `json:"Complexity,omitempty" xml:"Complexity,omitempty"`
	Summary           *string                                        `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Product           *string                                        `json:"Product,omitempty" xml:"Product,omitempty"`
	PocCreateTime     *int64                                         `json:"PocCreateTime,omitempty" xml:"PocCreateTime,omitempty"`
	CveId             *string                                        `json:"CveId,omitempty" xml:"CveId,omitempty"`
	CnvdId            *string                                        `json:"CnvdId,omitempty" xml:"CnvdId,omitempty"`
	Reference         *string                                        `json:"Reference,omitempty" xml:"Reference,omitempty"`
	CvssScore         *string                                        `json:"CvssScore,omitempty" xml:"CvssScore,omitempty"`
	Vendor            *string                                        `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	PocDisclosureTime *int64                                         `json:"PocDisclosureTime,omitempty" xml:"PocDisclosureTime,omitempty"`
	CvssVector        *string                                        `json:"CvssVector,omitempty" xml:"CvssVector,omitempty"`
	Classify          *string                                        `json:"Classify,omitempty" xml:"Classify,omitempty"`
	VulLevel          *string                                        `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
	ReleaseTime       *int64                                         `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	Title             *string                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	Solution          *string                                        `json:"Solution,omitempty" xml:"Solution,omitempty"`
	Content           *string                                        `json:"Content,omitempty" xml:"Content,omitempty"`
	Poc               *string                                        `json:"Poc,omitempty" xml:"Poc,omitempty"`
	Classifys         []*DescribeVulDetailsResponseBodyCvesClassifys `json:"Classifys,omitempty" xml:"Classifys,omitempty" type:"Repeated"`
}

func (s DescribeVulDetailsResponseBodyCves) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseBodyCves) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBodyCves) SetComplexity(v string) *DescribeVulDetailsResponseBodyCves {
	s.Complexity = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetSummary(v string) *DescribeVulDetailsResponseBodyCves {
	s.Summary = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetProduct(v string) *DescribeVulDetailsResponseBodyCves {
	s.Product = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetPocCreateTime(v int64) *DescribeVulDetailsResponseBodyCves {
	s.PocCreateTime = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCveId(v string) *DescribeVulDetailsResponseBodyCves {
	s.CveId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCnvdId(v string) *DescribeVulDetailsResponseBodyCves {
	s.CnvdId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetReference(v string) *DescribeVulDetailsResponseBodyCves {
	s.Reference = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCvssScore(v string) *DescribeVulDetailsResponseBodyCves {
	s.CvssScore = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetVendor(v string) *DescribeVulDetailsResponseBodyCves {
	s.Vendor = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetPocDisclosureTime(v int64) *DescribeVulDetailsResponseBodyCves {
	s.PocDisclosureTime = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCvssVector(v string) *DescribeVulDetailsResponseBodyCves {
	s.CvssVector = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetClassify(v string) *DescribeVulDetailsResponseBodyCves {
	s.Classify = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetVulLevel(v string) *DescribeVulDetailsResponseBodyCves {
	s.VulLevel = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetReleaseTime(v int64) *DescribeVulDetailsResponseBodyCves {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetTitle(v string) *DescribeVulDetailsResponseBodyCves {
	s.Title = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetSolution(v string) *DescribeVulDetailsResponseBodyCves {
	s.Solution = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetContent(v string) *DescribeVulDetailsResponseBodyCves {
	s.Content = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetPoc(v string) *DescribeVulDetailsResponseBodyCves {
	s.Poc = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetClassifys(v []*DescribeVulDetailsResponseBodyCvesClassifys) *DescribeVulDetailsResponseBodyCves {
	s.Classifys = v
	return s
}

type DescribeVulDetailsResponseBodyCvesClassifys struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Classify     *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	DemoVideoUrl *string `json:"DemoVideoUrl,omitempty" xml:"DemoVideoUrl,omitempty"`
}

func (s DescribeVulDetailsResponseBodyCvesClassifys) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponseBodyCvesClassifys) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) SetDescription(v string) *DescribeVulDetailsResponseBodyCvesClassifys {
	s.Description = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) SetClassify(v string) *DescribeVulDetailsResponseBodyCvesClassifys {
	s.Classify = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) SetDemoVideoUrl(v string) *DescribeVulDetailsResponseBodyCvesClassifys {
	s.DemoVideoUrl = &v
	return s
}

type DescribeVulDetailsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponse) SetHeaders(v map[string]*string) *DescribeVulDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulDetailsResponse) SetBody(v *DescribeVulDetailsResponseBody) *DescribeVulDetailsResponse {
	s.Body = v
	return s
}

type DescribeVulListRequest struct {
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Ids                 *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	Remark              *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuids               *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AliasName           *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Level               *string `json:"Level,omitempty" xml:"Level,omitempty"`
	StatusList          *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	Necessity           *string `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	Dealed              *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	BatchName           *string `json:"BatchName,omitempty" xml:"BatchName,omitempty"`
	Resource            *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	CreateTsStart       *int64  `json:"CreateTsStart,omitempty" xml:"CreateTsStart,omitempty"`
	CreateTsEnd         *int64  `json:"CreateTsEnd,omitempty" xml:"CreateTsEnd,omitempty"`
	CurrentPage         *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ModifyTsStart       *int64  `json:"ModifyTsStart,omitempty" xml:"ModifyTsStart,omitempty"`
	ModifyTsEnd         *int64  `json:"ModifyTsEnd,omitempty" xml:"ModifyTsEnd,omitempty"`
	AttachTypes         *string `json:"AttachTypes,omitempty" xml:"AttachTypes,omitempty"`
	ContainerFieldName  *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MinScore            *int32  `json:"MinScore,omitempty" xml:"MinScore,omitempty"`
	VpcInstanceIds      *string `json:"VpcInstanceIds,omitempty" xml:"VpcInstanceIds,omitempty"`
}

func (s DescribeVulListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulListRequest) SetSourceIp(v string) *DescribeVulListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVulListRequest) SetLang(v string) *DescribeVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulListRequest) SetIds(v string) *DescribeVulListRequest {
	s.Ids = &v
	return s
}

func (s *DescribeVulListRequest) SetRemark(v string) *DescribeVulListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeVulListRequest) SetGroupId(v string) *DescribeVulListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeVulListRequest) SetType(v string) *DescribeVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeVulListRequest) SetUuids(v string) *DescribeVulListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeVulListRequest) SetName(v string) *DescribeVulListRequest {
	s.Name = &v
	return s
}

func (s *DescribeVulListRequest) SetAliasName(v string) *DescribeVulListRequest {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListRequest) SetLevel(v string) *DescribeVulListRequest {
	s.Level = &v
	return s
}

func (s *DescribeVulListRequest) SetStatusList(v string) *DescribeVulListRequest {
	s.StatusList = &v
	return s
}

func (s *DescribeVulListRequest) SetNecessity(v string) *DescribeVulListRequest {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListRequest) SetDealed(v string) *DescribeVulListRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeVulListRequest) SetBatchName(v string) *DescribeVulListRequest {
	s.BatchName = &v
	return s
}

func (s *DescribeVulListRequest) SetResource(v string) *DescribeVulListRequest {
	s.Resource = &v
	return s
}

func (s *DescribeVulListRequest) SetCreateTsStart(v int64) *DescribeVulListRequest {
	s.CreateTsStart = &v
	return s
}

func (s *DescribeVulListRequest) SetCreateTsEnd(v int64) *DescribeVulListRequest {
	s.CreateTsEnd = &v
	return s
}

func (s *DescribeVulListRequest) SetCurrentPage(v int32) *DescribeVulListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListRequest) SetPageSize(v int32) *DescribeVulListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListRequest) SetModifyTsStart(v int64) *DescribeVulListRequest {
	s.ModifyTsStart = &v
	return s
}

func (s *DescribeVulListRequest) SetModifyTsEnd(v int64) *DescribeVulListRequest {
	s.ModifyTsEnd = &v
	return s
}

func (s *DescribeVulListRequest) SetAttachTypes(v string) *DescribeVulListRequest {
	s.AttachTypes = &v
	return s
}

func (s *DescribeVulListRequest) SetContainerFieldName(v string) *DescribeVulListRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeVulListRequest) SetContainerFieldValue(v string) *DescribeVulListRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeVulListRequest) SetTargetType(v string) *DescribeVulListRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeVulListRequest) SetClusterId(v string) *DescribeVulListRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeVulListRequest) SetMinScore(v int32) *DescribeVulListRequest {
	s.MinScore = &v
	return s
}

func (s *DescribeVulListRequest) SetVpcInstanceIds(v string) *DescribeVulListRequest {
	s.VpcInstanceIds = &v
	return s
}

type DescribeVulListResponseBody struct {
	TotalCount  *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VulRecords  []*DescribeVulListResponseBodyVulRecords `json:"VulRecords,omitempty" xml:"VulRecords,omitempty" type:"Repeated"`
	CurrentPage *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeVulListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBody) SetTotalCount(v int32) *DescribeVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulListResponseBody) SetRequestId(v string) *DescribeVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulListResponseBody) SetPageSize(v int32) *DescribeVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListResponseBody) SetVulRecords(v []*DescribeVulListResponseBodyVulRecords) *DescribeVulListResponseBody {
	s.VulRecords = v
	return s
}

func (s *DescribeVulListResponseBody) SetCurrentPage(v int32) *DescribeVulListResponseBody {
	s.CurrentPage = &v
	return s
}

type DescribeVulListResponseBodyVulRecords struct {
	Status             *int32                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	ModifyTs           *int64                                                  `json:"ModifyTs,omitempty" xml:"ModifyTs,omitempty"`
	InternetIp         *string                                                 `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	ContainerImageId   *string                                                 `json:"ContainerImageId,omitempty" xml:"ContainerImageId,omitempty"`
	PrimaryId          *int64                                                  `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	Tag                *string                                                 `json:"Tag,omitempty" xml:"Tag,omitempty"`
	K8sClusterId       *string                                                 `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ContainerImageName *string                                                 `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	K8sNodeId          *string                                                 `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	InstanceName       *string                                                 `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Online             *bool                                                   `json:"Online,omitempty" xml:"Online,omitempty"`
	ContainerInnerPath *string                                                 `json:"ContainerInnerPath,omitempty" xml:"ContainerInnerPath,omitempty"`
	OsVersion          *string                                                 `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	Name               *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	ExtendContentJson  *DescribeVulListResponseBodyVulRecordsExtendContentJson `json:"ExtendContentJson,omitempty" xml:"ExtendContentJson,omitempty" type:"Struct"`
	ResultCode         *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	InstanceId         *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Related            *string                                                 `json:"Related,omitempty" xml:"Related,omitempty"`
	IntranetIp         *string                                                 `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	LastTs             *int64                                                  `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	FirstTs            *int64                                                  `json:"FirstTs,omitempty" xml:"FirstTs,omitempty"`
	Necessity          *string                                                 `json:"Necessity,omitempty" xml:"Necessity,omitempty"`
	RepairTs           *int64                                                  `json:"RepairTs,omitempty" xml:"RepairTs,omitempty"`
	Uuid               *string                                                 `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	K8sPodName         *string                                                 `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	ContainerId        *string                                                 `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	GroupId            *int32                                                  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ResultMessage      *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	K8sNamespace       *string                                                 `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	AliasName          *string                                                 `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	K8sNodeName        *string                                                 `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	ContainerName      *string                                                 `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecords) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecords) SetStatus(v int32) *DescribeVulListResponseBodyVulRecords {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetType(v string) *DescribeVulListResponseBodyVulRecords {
	s.Type = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetModifyTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.ModifyTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetInternetIp(v string) *DescribeVulListResponseBodyVulRecords {
	s.InternetIp = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetContainerImageId(v string) *DescribeVulListResponseBodyVulRecords {
	s.ContainerImageId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetPrimaryId(v int64) *DescribeVulListResponseBodyVulRecords {
	s.PrimaryId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetTag(v string) *DescribeVulListResponseBodyVulRecords {
	s.Tag = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sClusterId(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetContainerImageName(v string) *DescribeVulListResponseBodyVulRecords {
	s.ContainerImageName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sNodeId(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetInstanceName(v string) *DescribeVulListResponseBodyVulRecords {
	s.InstanceName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOnline(v bool) *DescribeVulListResponseBodyVulRecords {
	s.Online = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetContainerInnerPath(v string) *DescribeVulListResponseBodyVulRecords {
	s.ContainerInnerPath = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetOsVersion(v string) *DescribeVulListResponseBodyVulRecords {
	s.OsVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetName(v string) *DescribeVulListResponseBodyVulRecords {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetExtendContentJson(v *DescribeVulListResponseBodyVulRecordsExtendContentJson) *DescribeVulListResponseBodyVulRecords {
	s.ExtendContentJson = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultCode(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultCode = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetInstanceId(v string) *DescribeVulListResponseBodyVulRecords {
	s.InstanceId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRelated(v string) *DescribeVulListResponseBodyVulRecords {
	s.Related = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetIntranetIp(v string) *DescribeVulListResponseBodyVulRecords {
	s.IntranetIp = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetLastTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.LastTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetFirstTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.FirstTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetNecessity(v string) *DescribeVulListResponseBodyVulRecords {
	s.Necessity = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetRepairTs(v int64) *DescribeVulListResponseBodyVulRecords {
	s.RepairTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetUuid(v string) *DescribeVulListResponseBodyVulRecords {
	s.Uuid = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sPodName(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sPodName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetContainerId(v string) *DescribeVulListResponseBodyVulRecords {
	s.ContainerId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetGroupId(v int32) *DescribeVulListResponseBodyVulRecords {
	s.GroupId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetResultMessage(v string) *DescribeVulListResponseBodyVulRecords {
	s.ResultMessage = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sNamespace(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetAliasName(v string) *DescribeVulListResponseBodyVulRecords {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetK8sNodeName(v string) *DescribeVulListResponseBodyVulRecords {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecords) SetContainerName(v string) *DescribeVulListResponseBodyVulRecords {
	s.ContainerName = &v
	return s
}

type DescribeVulListResponseBodyVulRecordsExtendContentJson struct {
	Status        *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	CveList       []*string                                                              `json:"cveList,omitempty" xml:"cveList,omitempty" type:"Repeated"`
	PrimaryId     *int64                                                                 `json:"PrimaryId,omitempty" xml:"PrimaryId,omitempty"`
	Tag           *string                                                                `json:"Tag,omitempty" xml:"Tag,omitempty"`
	OsRelease     *string                                                                `json:"OsRelease,omitempty" xml:"OsRelease,omitempty"`
	RpmEntityList []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList `json:"RpmEntityList,omitempty" xml:"RpmEntityList,omitempty" type:"Repeated"`
	Ip            *string                                                                `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Os            *string                                                                `json:"Os,omitempty" xml:"Os,omitempty"`
	LastTs        *int64                                                                 `json:"LastTs,omitempty" xml:"LastTs,omitempty"`
	Necessity     *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity       `json:"Necessity,omitempty" xml:"Necessity,omitempty" type:"Struct"`
	AliasName     *string                                                                `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	AbsolutePath  *string                                                                `json:"AbsolutePath,omitempty" xml:"AbsolutePath,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJson) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetStatus(v int32) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetCveList(v []*string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.CveList = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetPrimaryId(v int64) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.PrimaryId = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetTag(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Tag = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetOsRelease(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.OsRelease = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetRpmEntityList(v []*DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.RpmEntityList = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetIp(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Ip = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetOs(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Os = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetLastTs(v int64) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.LastTs = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetNecessity(v *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.Necessity = v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetAliasName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.AliasName = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJson) SetAbsolutePath(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJson {
	s.AbsolutePath = &v
	return s
}

type DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList struct {
	FullVersion *string `json:"FullVersion,omitempty" xml:"FullVersion,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	MatchDetail *string `json:"MatchDetail,omitempty" xml:"MatchDetail,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateCmd   *string `json:"UpdateCmd,omitempty" xml:"UpdateCmd,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetFullVersion(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.FullVersion = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetVersion(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Version = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetMatchDetail(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.MatchDetail = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetPath(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Path = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetName(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.Name = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList) SetUpdateCmd(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonRpmEntityList {
	s.UpdateCmd = &v
	return s
}

type DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity struct {
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeFactor       *string `json:"Time_factor,omitempty" xml:"Time_factor,omitempty"`
	EnviromentFactor *string `json:"Enviroment_factor,omitempty" xml:"Enviroment_factor,omitempty"`
	IsCalc           *string `json:"Is_calc,omitempty" xml:"Is_calc,omitempty"`
	TotalScore       *string `json:"Total_score,omitempty" xml:"Total_score,omitempty"`
	CvssFactor       *string `json:"Cvss_factor,omitempty" xml:"Cvss_factor,omitempty"`
	AssetsFactor     *string `json:"Assets_factor,omitempty" xml:"Assets_factor,omitempty"`
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetStatus(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.Status = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetTimeFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.TimeFactor = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetEnviromentFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.EnviromentFactor = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetIsCalc(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.IsCalc = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetTotalScore(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.TotalScore = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetCvssFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.CvssFactor = &v
	return s
}

func (s *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity) SetAssetsFactor(v string) *DescribeVulListResponseBodyVulRecordsExtendContentJsonNecessity {
	s.AssetsFactor = &v
	return s
}

type DescribeVulListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulListResponse) SetHeaders(v map[string]*string) *DescribeVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulListResponse) SetBody(v *DescribeVulListResponseBody) *DescribeVulListResponse {
	s.Body = v
	return s
}

type DescribeVulWhitelistRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVulWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistRequest) SetSourceIp(v string) *DescribeVulWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVulWhitelistRequest) SetCurrentPage(v int32) *DescribeVulWhitelistRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulWhitelistRequest) SetPageSize(v int32) *DescribeVulWhitelistRequest {
	s.PageSize = &v
	return s
}

type DescribeVulWhitelistResponseBody struct {
	TotalCount    *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize      *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VulWhitelists []*DescribeVulWhitelistResponseBodyVulWhitelists `json:"VulWhitelists,omitempty" xml:"VulWhitelists,omitempty" type:"Repeated"`
	CurrentPage   *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeVulWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistResponseBody) SetTotalCount(v int32) *DescribeVulWhitelistResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetRequestId(v string) *DescribeVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetPageSize(v int32) *DescribeVulWhitelistResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetVulWhitelists(v []*DescribeVulWhitelistResponseBodyVulWhitelists) *DescribeVulWhitelistResponseBody {
	s.VulWhitelists = v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetCurrentPage(v int32) *DescribeVulWhitelistResponseBody {
	s.CurrentPage = &v
	return s
}

type DescribeVulWhitelistResponseBodyVulWhitelists struct {
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeVulWhitelistResponseBodyVulWhitelists) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulWhitelistResponseBodyVulWhitelists) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetType(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Type = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetAliasName(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.AliasName = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetName(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Name = &v
	return s
}

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetReason(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Reason = &v
	return s
}

type DescribeVulWhitelistResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVulWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistResponse) SetHeaders(v map[string]*string) *DescribeVulWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulWhitelistResponse) SetBody(v *DescribeVulWhitelistResponseBody) *DescribeVulWhitelistResponse {
	s.Body = v
	return s
}

type DescribeWarningMachinesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	Uuids       *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	RiskId      *int64  `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	StrategyId  *int64  `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s DescribeWarningMachinesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWarningMachinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesRequest) SetSourceIp(v string) *DescribeWarningMachinesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetLang(v string) *DescribeWarningMachinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetMachineName(v string) *DescribeWarningMachinesRequest {
	s.MachineName = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetUuids(v string) *DescribeWarningMachinesRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetRiskId(v int64) *DescribeWarningMachinesRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetStrategyId(v int64) *DescribeWarningMachinesRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetPageSize(v int32) *DescribeWarningMachinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWarningMachinesRequest) SetCurrentPage(v int32) *DescribeWarningMachinesRequest {
	s.CurrentPage = &v
	return s
}

type DescribeWarningMachinesResponseBody struct {
	TotalCount      *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage     *int32                                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	WarningMachines []*DescribeWarningMachinesResponseBodyWarningMachines `json:"WarningMachines,omitempty" xml:"WarningMachines,omitempty" type:"Repeated"`
	Count           *int32                                                `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeWarningMachinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWarningMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesResponseBody) SetTotalCount(v int32) *DescribeWarningMachinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetPageSize(v int32) *DescribeWarningMachinesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetRequestId(v string) *DescribeWarningMachinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetCurrentPage(v int32) *DescribeWarningMachinesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetWarningMachines(v []*DescribeWarningMachinesResponseBodyWarningMachines) *DescribeWarningMachinesResponseBody {
	s.WarningMachines = v
	return s
}

func (s *DescribeWarningMachinesResponseBody) SetCount(v int32) *DescribeWarningMachinesResponseBody {
	s.Count = &v
	return s
}

type DescribeWarningMachinesResponseBodyWarningMachines struct {
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	LowWarningCount    *int32  `json:"LowWarningCount,omitempty" xml:"LowWarningCount,omitempty"`
	Uuid               *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	MediumWarningCount *int32  `json:"MediumWarningCount,omitempty" xml:"MediumWarningCount,omitempty"`
	PassCount          *int32  `json:"PassCount,omitempty" xml:"PassCount,omitempty"`
	InternetIp         *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	HighWarningCount   *int32  `json:"HighWarningCount,omitempty" xml:"HighWarningCount,omitempty"`
	IntranetIp         *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeWarningMachinesResponseBodyWarningMachines) String() string {
	return tea.Prettify(s)
}

func (s DescribeWarningMachinesResponseBodyWarningMachines) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetStatus(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.Status = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetLowWarningCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.LowWarningCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetUuid(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.Uuid = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetMediumWarningCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.MediumWarningCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetPassCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.PassCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetInternetIp(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.InternetIp = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetInstanceName(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.InstanceName = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetInstanceId(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.InstanceId = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetHighWarningCount(v int32) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.HighWarningCount = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetIntranetIp(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.IntranetIp = &v
	return s
}

func (s *DescribeWarningMachinesResponseBodyWarningMachines) SetRegionId(v string) *DescribeWarningMachinesResponseBodyWarningMachines {
	s.RegionId = &v
	return s
}

type DescribeWarningMachinesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWarningMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWarningMachinesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWarningMachinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWarningMachinesResponse) SetHeaders(v map[string]*string) *DescribeWarningMachinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWarningMachinesResponse) SetBody(v *DescribeWarningMachinesResponseBody) *DescribeWarningMachinesResponse {
	s.Body = v
	return s
}

type DescribeWebLockBindListRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebLockBindListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockBindListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListRequest) SetSourceIp(v string) *DescribeWebLockBindListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetLang(v string) *DescribeWebLockBindListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetRemark(v string) *DescribeWebLockBindListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetStatus(v string) *DescribeWebLockBindListRequest {
	s.Status = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetCurrentPage(v int32) *DescribeWebLockBindListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockBindListRequest) SetPageSize(v int32) *DescribeWebLockBindListRequest {
	s.PageSize = &v
	return s
}

type DescribeWebLockBindListResponseBody struct {
	TotalCount  *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32                                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	BindList    []*DescribeWebLockBindListResponseBodyBindList `json:"BindList,omitempty" xml:"BindList,omitempty" type:"Repeated"`
}

func (s DescribeWebLockBindListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockBindListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListResponseBody) SetTotalCount(v int32) *DescribeWebLockBindListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetRequestId(v string) *DescribeWebLockBindListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetPageSize(v int32) *DescribeWebLockBindListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetCurrentPage(v int32) *DescribeWebLockBindListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetBindList(v []*DescribeWebLockBindListResponseBodyBindList) *DescribeWebLockBindListResponseBody {
	s.BindList = v
	return s
}

type DescribeWebLockBindListResponseBodyBindList struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent       *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ServiceDetail *string `json:"ServiceDetail,omitempty" xml:"ServiceDetail,omitempty"`
	InternetIp    *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	Os            *string `json:"Os,omitempty" xml:"Os,omitempty"`
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	IntranetIp    *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	AuditCount    *string `json:"AuditCount,omitempty" xml:"AuditCount,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	ServiceCode   *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	InstanceName  *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	DirCount      *string `json:"DirCount,omitempty" xml:"DirCount,omitempty"`
	BlockCount    *string `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
}

func (s DescribeWebLockBindListResponseBodyBindList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockBindListResponseBodyBindList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetStatus(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.Status = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetPercent(v int32) *DescribeWebLockBindListResponseBodyBindList {
	s.Percent = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetServiceDetail(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.ServiceDetail = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetInternetIp(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.InternetIp = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetOs(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.Os = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetServiceStatus(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetIntranetIp(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetAuditCount(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.AuditCount = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetUuid(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetServiceCode(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.ServiceCode = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetInstanceName(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.InstanceName = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetDirCount(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.DirCount = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetBlockCount(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.BlockCount = &v
	return s
}

type DescribeWebLockBindListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebLockBindListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebLockBindListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockBindListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListResponse) SetHeaders(v map[string]*string) *DescribeWebLockBindListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockBindListResponse) SetBody(v *DescribeWebLockBindListResponseBody) *DescribeWebLockBindListResponse {
	s.Body = v
	return s
}

type DescribeWebLockConfigListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWebLockConfigListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockConfigListRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListRequest) SetSourceIp(v string) *DescribeWebLockConfigListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebLockConfigListRequest) SetLang(v string) *DescribeWebLockConfigListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWebLockConfigListRequest) SetUuid(v string) *DescribeWebLockConfigListRequest {
	s.Uuid = &v
	return s
}

type DescribeWebLockConfigListResponseBody struct {
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConfigList []*DescribeWebLockConfigListResponseBodyConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s DescribeWebLockConfigListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockConfigListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListResponseBody) SetTotalCount(v int32) *DescribeWebLockConfigListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBody) SetRequestId(v string) *DescribeWebLockConfigListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBody) SetConfigList(v []*DescribeWebLockConfigListResponseBodyConfigList) *DescribeWebLockConfigListResponseBody {
	s.ConfigList = v
	return s
}

type DescribeWebLockConfigListResponseBodyConfigList struct {
	ExclusiveDir      *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	DefenceMode       *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	InclusiveFile     *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Dir               *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	ExclusiveFile     *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	LocalBackupDir    *string `json:"LocalBackupDir,omitempty" xml:"LocalBackupDir,omitempty"`
}

func (s DescribeWebLockConfigListResponseBodyConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockConfigListResponseBodyConfigList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetExclusiveDir(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.ExclusiveDir = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetUuid(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetInclusiveFileType(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.InclusiveFileType = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetDefenceMode(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.DefenceMode = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetExclusiveFileType(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.ExclusiveFileType = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetInclusiveFile(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.InclusiveFile = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetMode(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Mode = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetDir(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Dir = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetExclusiveFile(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.ExclusiveFile = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetId(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.Id = &v
	return s
}

func (s *DescribeWebLockConfigListResponseBodyConfigList) SetLocalBackupDir(v string) *DescribeWebLockConfigListResponseBodyConfigList {
	s.LocalBackupDir = &v
	return s
}

type DescribeWebLockConfigListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeWebLockConfigListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebLockConfigListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebLockConfigListResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockConfigListResponse) SetHeaders(v map[string]*string) *DescribeWebLockConfigListResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockConfigListResponse) SetBody(v *DescribeWebLockConfigListResponseBody) *DescribeWebLockConfigListResponse {
	s.Body = v
	return s
}

type ExportRecordRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Params     *string `json:"Params,omitempty" xml:"Params,omitempty"`
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
}

func (s ExportRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportRecordRequest) GoString() string {
	return s.String()
}

func (s *ExportRecordRequest) SetSourceIp(v string) *ExportRecordRequest {
	s.SourceIp = &v
	return s
}

func (s *ExportRecordRequest) SetLang(v string) *ExportRecordRequest {
	s.Lang = &v
	return s
}

func (s *ExportRecordRequest) SetParams(v string) *ExportRecordRequest {
	s.Params = &v
	return s
}

func (s *ExportRecordRequest) SetExportType(v string) *ExportRecordRequest {
	s.ExportType = &v
	return s
}

type ExportRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ExportRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ExportRecordResponseBody) SetRequestId(v string) *ExportRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportRecordResponseBody) SetFileName(v string) *ExportRecordResponseBody {
	s.FileName = &v
	return s
}

func (s *ExportRecordResponseBody) SetId(v int64) *ExportRecordResponseBody {
	s.Id = &v
	return s
}

type ExportRecordResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExportRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExportRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportRecordResponse) GoString() string {
	return s.String()
}

func (s *ExportRecordResponse) SetHeaders(v map[string]*string) *ExportRecordResponse {
	s.Headers = v
	return s
}

func (s *ExportRecordResponse) SetBody(v *ExportRecordResponseBody) *ExportRecordResponse {
	s.Body = v
	return s
}

type FixCheckWarningsRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RiskId      *int64  `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	CheckParams *string `json:"CheckParams,omitempty" xml:"CheckParams,omitempty"`
	Uuids       *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s FixCheckWarningsRequest) String() string {
	return tea.Prettify(s)
}

func (s FixCheckWarningsRequest) GoString() string {
	return s.String()
}

func (s *FixCheckWarningsRequest) SetSourceIp(v string) *FixCheckWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *FixCheckWarningsRequest) SetLang(v string) *FixCheckWarningsRequest {
	s.Lang = &v
	return s
}

func (s *FixCheckWarningsRequest) SetRiskId(v int64) *FixCheckWarningsRequest {
	s.RiskId = &v
	return s
}

func (s *FixCheckWarningsRequest) SetCheckParams(v string) *FixCheckWarningsRequest {
	s.CheckParams = &v
	return s
}

func (s *FixCheckWarningsRequest) SetUuids(v string) *FixCheckWarningsRequest {
	s.Uuids = &v
	return s
}

type FixCheckWarningsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BatchId   *int64  `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s FixCheckWarningsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FixCheckWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *FixCheckWarningsResponseBody) SetRequestId(v string) *FixCheckWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *FixCheckWarningsResponseBody) SetBatchId(v int64) *FixCheckWarningsResponseBody {
	s.BatchId = &v
	return s
}

type FixCheckWarningsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FixCheckWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FixCheckWarningsResponse) String() string {
	return tea.Prettify(s)
}

func (s FixCheckWarningsResponse) GoString() string {
	return s.String()
}

func (s *FixCheckWarningsResponse) SetHeaders(v map[string]*string) *FixCheckWarningsResponse {
	s.Headers = v
	return s
}

func (s *FixCheckWarningsResponse) SetBody(v *FixCheckWarningsResponseBody) *FixCheckWarningsResponse {
	s.Body = v
	return s
}

type GetIncIOCsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Date     *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetIncIOCsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIncIOCsRequest) GoString() string {
	return s.String()
}

func (s *GetIncIOCsRequest) SetSourceIp(v string) *GetIncIOCsRequest {
	s.SourceIp = &v
	return s
}

func (s *GetIncIOCsRequest) SetType(v string) *GetIncIOCsRequest {
	s.Type = &v
	return s
}

func (s *GetIncIOCsRequest) SetDate(v string) *GetIncIOCsRequest {
	s.Date = &v
	return s
}

type GetIncIOCsResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIncIOCsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIncIOCsResponseBody) GoString() string {
	return s.String()
}

func (s *GetIncIOCsResponseBody) SetData(v string) *GetIncIOCsResponseBody {
	s.Data = &v
	return s
}

func (s *GetIncIOCsResponseBody) SetRequestId(v string) *GetIncIOCsResponseBody {
	s.RequestId = &v
	return s
}

type GetIncIOCsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIncIOCsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIncIOCsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIncIOCsResponse) GoString() string {
	return s.String()
}

func (s *GetIncIOCsResponse) SetHeaders(v map[string]*string) *GetIncIOCsResponse {
	s.Headers = v
	return s
}

func (s *GetIncIOCsResponse) SetBody(v *GetIncIOCsResponseBody) *GetIncIOCsResponse {
	s.Body = v
	return s
}

type GetIOCsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Date     *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetIOCsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIOCsRequest) GoString() string {
	return s.String()
}

func (s *GetIOCsRequest) SetSourceIp(v string) *GetIOCsRequest {
	s.SourceIp = &v
	return s
}

func (s *GetIOCsRequest) SetType(v string) *GetIOCsRequest {
	s.Type = &v
	return s
}

func (s *GetIOCsRequest) SetDate(v string) *GetIOCsRequest {
	s.Date = &v
	return s
}

type GetIOCsResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIOCsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIOCsResponseBody) GoString() string {
	return s.String()
}

func (s *GetIOCsResponseBody) SetData(v string) *GetIOCsResponseBody {
	s.Data = &v
	return s
}

func (s *GetIOCsResponseBody) SetRequestId(v string) *GetIOCsResponseBody {
	s.RequestId = &v
	return s
}

type GetIOCsResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIOCsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIOCsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIOCsResponse) GoString() string {
	return s.String()
}

func (s *GetIOCsResponse) SetHeaders(v map[string]*string) *GetIOCsResponse {
	s.Headers = v
	return s
}

func (s *GetIOCsResponse) SetBody(v *GetIOCsResponseBody) *GetIOCsResponse {
	s.Body = v
	return s
}

type HandleSecurityEventsRequest struct {
	SourceIp         *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId  *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OperationCode    *string   `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	OperationParams  *string   `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	MarkMissParam    *string   `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
	SecurityEventIds []*string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty" type:"Repeated"`
}

func (s HandleSecurityEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsRequest) SetSourceIp(v string) *HandleSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetResourceOwnerId(v int64) *HandleSecurityEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationCode(v string) *HandleSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetOperationParams(v string) *HandleSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetMarkMissParam(v string) *HandleSecurityEventsRequest {
	s.MarkMissParam = &v
	return s
}

func (s *HandleSecurityEventsRequest) SetSecurityEventIds(v []*string) *HandleSecurityEventsRequest {
	s.SecurityEventIds = v
	return s
}

type HandleSecurityEventsResponseBody struct {
	HandleSecurityEventsResponse *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse `json:"HandleSecurityEventsResponse,omitempty" xml:"HandleSecurityEventsResponse,omitempty" type:"Struct"`
	RequestId                    *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleSecurityEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponseBody) SetHandleSecurityEventsResponse(v *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) *HandleSecurityEventsResponseBody {
	s.HandleSecurityEventsResponse = v
	return s
}

func (s *HandleSecurityEventsResponseBody) SetRequestId(v string) *HandleSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

type HandleSecurityEventsResponseBodyHandleSecurityEventsResponse struct {
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse) SetTaskId(v int64) *HandleSecurityEventsResponseBodyHandleSecurityEventsResponse {
	s.TaskId = &v
	return s
}

type HandleSecurityEventsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HandleSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandleSecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponse) SetHeaders(v map[string]*string) *HandleSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *HandleSecurityEventsResponse) SetBody(v *HandleSecurityEventsResponseBody) *HandleSecurityEventsResponse {
	s.Body = v
	return s
}

type HandleSimilarSecurityEventsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	OperationCode   *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
	OperationParams *string `json:"OperationParams,omitempty" xml:"OperationParams,omitempty"`
	MarkMissParam   *string `json:"MarkMissParam,omitempty" xml:"MarkMissParam,omitempty"`
}

func (s HandleSimilarSecurityEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s HandleSimilarSecurityEventsRequest) GoString() string {
	return s.String()
}

func (s *HandleSimilarSecurityEventsRequest) SetSourceIp(v string) *HandleSimilarSecurityEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetResourceOwnerId(v int64) *HandleSimilarSecurityEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetTaskId(v int64) *HandleSimilarSecurityEventsRequest {
	s.TaskId = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetOperationCode(v string) *HandleSimilarSecurityEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetOperationParams(v string) *HandleSimilarSecurityEventsRequest {
	s.OperationParams = &v
	return s
}

func (s *HandleSimilarSecurityEventsRequest) SetMarkMissParam(v string) *HandleSimilarSecurityEventsRequest {
	s.MarkMissParam = &v
	return s
}

type HandleSimilarSecurityEventsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleSimilarSecurityEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HandleSimilarSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *HandleSimilarSecurityEventsResponseBody) SetRequestId(v string) *HandleSimilarSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

type HandleSimilarSecurityEventsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HandleSimilarSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HandleSimilarSecurityEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s HandleSimilarSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSimilarSecurityEventsResponse) SetHeaders(v map[string]*string) *HandleSimilarSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *HandleSimilarSecurityEventsResponse) SetBody(v *HandleSimilarSecurityEventsResponseBody) *HandleSimilarSecurityEventsResponse {
	s.Body = v
	return s
}

type IgnoreHcCheckWarningsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	CheckWarningIds *string `json:"CheckWarningIds,omitempty" xml:"CheckWarningIds,omitempty"`
	CheckIds        *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	RiskId          *string `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	Type            *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
	Reason          *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s IgnoreHcCheckWarningsRequest) String() string {
	return tea.Prettify(s)
}

func (s IgnoreHcCheckWarningsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreHcCheckWarningsRequest) SetSourceIp(v string) *IgnoreHcCheckWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetCheckWarningIds(v string) *IgnoreHcCheckWarningsRequest {
	s.CheckWarningIds = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetCheckIds(v string) *IgnoreHcCheckWarningsRequest {
	s.CheckIds = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetRiskId(v string) *IgnoreHcCheckWarningsRequest {
	s.RiskId = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetType(v int64) *IgnoreHcCheckWarningsRequest {
	s.Type = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetReason(v string) *IgnoreHcCheckWarningsRequest {
	s.Reason = &v
	return s
}

type IgnoreHcCheckWarningsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreHcCheckWarningsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IgnoreHcCheckWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreHcCheckWarningsResponseBody) SetRequestId(v string) *IgnoreHcCheckWarningsResponseBody {
	s.RequestId = &v
	return s
}

type IgnoreHcCheckWarningsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IgnoreHcCheckWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IgnoreHcCheckWarningsResponse) String() string {
	return tea.Prettify(s)
}

func (s IgnoreHcCheckWarningsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreHcCheckWarningsResponse) SetHeaders(v map[string]*string) *IgnoreHcCheckWarningsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreHcCheckWarningsResponse) SetBody(v *IgnoreHcCheckWarningsResponseBody) *IgnoreHcCheckWarningsResponse {
	s.Body = v
	return s
}

type ModifyAntiBruteForceRuleRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Id              *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Span            *int32    `json:"Span,omitempty" xml:"Span,omitempty"`
	FailCount       *int32    `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	ForbiddenTime   *int32    `json:"ForbiddenTime,omitempty" xml:"ForbiddenTime,omitempty"`
	EnableSmartRule *bool     `json:"EnableSmartRule,omitempty" xml:"EnableSmartRule,omitempty"`
	DefaultRule     *bool     `json:"DefaultRule,omitempty" xml:"DefaultRule,omitempty"`
	UuidList        []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s ModifyAntiBruteForceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAntiBruteForceRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleRequest) SetSourceIp(v string) *ModifyAntiBruteForceRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetResourceOwnerId(v int64) *ModifyAntiBruteForceRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetId(v int64) *ModifyAntiBruteForceRuleRequest {
	s.Id = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetName(v string) *ModifyAntiBruteForceRuleRequest {
	s.Name = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetSpan(v int32) *ModifyAntiBruteForceRuleRequest {
	s.Span = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetFailCount(v int32) *ModifyAntiBruteForceRuleRequest {
	s.FailCount = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetForbiddenTime(v int32) *ModifyAntiBruteForceRuleRequest {
	s.ForbiddenTime = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetEnableSmartRule(v bool) *ModifyAntiBruteForceRuleRequest {
	s.EnableSmartRule = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetDefaultRule(v bool) *ModifyAntiBruteForceRuleRequest {
	s.DefaultRule = &v
	return s
}

func (s *ModifyAntiBruteForceRuleRequest) SetUuidList(v []*string) *ModifyAntiBruteForceRuleRequest {
	s.UuidList = v
	return s
}

type ModifyAntiBruteForceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAntiBruteForceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAntiBruteForceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleResponseBody) SetRequestId(v string) *ModifyAntiBruteForceRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAntiBruteForceRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAntiBruteForceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAntiBruteForceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAntiBruteForceRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleResponse) SetHeaders(v map[string]*string) *ModifyAntiBruteForceRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAntiBruteForceRuleResponse) SetBody(v *ModifyAntiBruteForceRuleResponseBody) *ModifyAntiBruteForceRuleResponse {
	s.Body = v
	return s
}

type ModifyCreateVulWhitelistRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifyCreateVulWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCreateVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyCreateVulWhitelistRequest) SetSourceIp(v string) *ModifyCreateVulWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyCreateVulWhitelistRequest) SetWhitelist(v string) *ModifyCreateVulWhitelistRequest {
	s.Whitelist = &v
	return s
}

func (s *ModifyCreateVulWhitelistRequest) SetReason(v string) *ModifyCreateVulWhitelistRequest {
	s.Reason = &v
	return s
}

type ModifyCreateVulWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCreateVulWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCreateVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCreateVulWhitelistResponseBody) SetRequestId(v string) *ModifyCreateVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ModifyCreateVulWhitelistResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCreateVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCreateVulWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCreateVulWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyCreateVulWhitelistResponse) SetHeaders(v map[string]*string) *ModifyCreateVulWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyCreateVulWhitelistResponse) SetBody(v *ModifyCreateVulWhitelistResponseBody) *ModifyCreateVulWhitelistResponse {
	s.Body = v
	return s
}

type ModifyEmgVulSubmitRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserAgreement *string `json:"UserAgreement,omitempty" xml:"UserAgreement,omitempty"`
}

func (s ModifyEmgVulSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyEmgVulSubmitRequest) GoString() string {
	return s.String()
}

func (s *ModifyEmgVulSubmitRequest) SetSourceIp(v string) *ModifyEmgVulSubmitRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetLang(v string) *ModifyEmgVulSubmitRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetName(v string) *ModifyEmgVulSubmitRequest {
	s.Name = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetUserAgreement(v string) *ModifyEmgVulSubmitRequest {
	s.UserAgreement = &v
	return s
}

type ModifyEmgVulSubmitResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEmgVulSubmitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyEmgVulSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEmgVulSubmitResponseBody) SetRequestId(v string) *ModifyEmgVulSubmitResponseBody {
	s.RequestId = &v
	return s
}

type ModifyEmgVulSubmitResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyEmgVulSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyEmgVulSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyEmgVulSubmitResponse) GoString() string {
	return s.String()
}

func (s *ModifyEmgVulSubmitResponse) SetHeaders(v map[string]*string) *ModifyEmgVulSubmitResponse {
	s.Headers = v
	return s
}

func (s *ModifyEmgVulSubmitResponse) SetBody(v *ModifyEmgVulSubmitResponseBody) *ModifyEmgVulSubmitResponse {
	s.Body = v
	return s
}

type ModifyGroupPropertyRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Data     *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ModifyGroupPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupPropertyRequest) SetSourceIp(v string) *ModifyGroupPropertyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyGroupPropertyRequest) SetData(v string) *ModifyGroupPropertyRequest {
	s.Data = &v
	return s
}

type ModifyGroupPropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGroupPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupPropertyResponseBody) SetRequestId(v string) *ModifyGroupPropertyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyGroupPropertyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyGroupPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGroupPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupPropertyResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupPropertyResponse) SetHeaders(v map[string]*string) *ModifyGroupPropertyResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupPropertyResponse) SetBody(v *ModifyGroupPropertyResponseBody) *ModifyGroupPropertyResponse {
	s.Body = v
	return s
}

type ModifyInstanceAntiBruteForceRuleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	NewRuleId       *int64  `json:"NewRuleId,omitempty" xml:"NewRuleId,omitempty"`
}

func (s ModifyInstanceAntiBruteForceRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAntiBruteForceRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetSourceIp(v string) *ModifyInstanceAntiBruteForceRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetResourceOwnerId(v int64) *ModifyInstanceAntiBruteForceRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetUuid(v string) *ModifyInstanceAntiBruteForceRuleRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleRequest) SetNewRuleId(v int64) *ModifyInstanceAntiBruteForceRuleRequest {
	s.NewRuleId = &v
	return s
}

type ModifyInstanceAntiBruteForceRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAntiBruteForceRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAntiBruteForceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAntiBruteForceRuleResponseBody) SetRequestId(v string) *ModifyInstanceAntiBruteForceRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAntiBruteForceRuleResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAntiBruteForceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAntiBruteForceRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAntiBruteForceRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) SetHeaders(v map[string]*string) *ModifyInstanceAntiBruteForceRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) SetBody(v *ModifyInstanceAntiBruteForceRuleResponseBody) *ModifyInstanceAntiBruteForceRuleResponse {
	s.Body = v
	return s
}

type ModifyLoginBaseConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Config   *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Target   *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyLoginBaseConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoginBaseConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoginBaseConfigRequest) SetSourceIp(v string) *ModifyLoginBaseConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyLoginBaseConfigRequest) SetType(v string) *ModifyLoginBaseConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyLoginBaseConfigRequest) SetConfig(v string) *ModifyLoginBaseConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyLoginBaseConfigRequest) SetTarget(v string) *ModifyLoginBaseConfigRequest {
	s.Target = &v
	return s
}

type ModifyLoginBaseConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoginBaseConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoginBaseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoginBaseConfigResponseBody) SetRequestId(v string) *ModifyLoginBaseConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLoginBaseConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLoginBaseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLoginBaseConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoginBaseConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoginBaseConfigResponse) SetHeaders(v map[string]*string) *ModifyLoginBaseConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoginBaseConfigResponse) SetBody(v *ModifyLoginBaseConfigResponseBody) *ModifyLoginBaseConfigResponse {
	s.Body = v
	return s
}

type ModifyLoginSwitchConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Item     *string `json:"Item,omitempty" xml:"Item,omitempty"`
	Status   *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyLoginSwitchConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoginSwitchConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLoginSwitchConfigRequest) SetSourceIp(v string) *ModifyLoginSwitchConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyLoginSwitchConfigRequest) SetItem(v string) *ModifyLoginSwitchConfigRequest {
	s.Item = &v
	return s
}

func (s *ModifyLoginSwitchConfigRequest) SetStatus(v int32) *ModifyLoginSwitchConfigRequest {
	s.Status = &v
	return s
}

type ModifyLoginSwitchConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoginSwitchConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoginSwitchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoginSwitchConfigResponseBody) SetRequestId(v string) *ModifyLoginSwitchConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLoginSwitchConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLoginSwitchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLoginSwitchConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLoginSwitchConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoginSwitchConfigResponse) SetHeaders(v map[string]*string) *ModifyLoginSwitchConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoginSwitchConfigResponse) SetBody(v *ModifyLoginSwitchConfigResponseBody) *ModifyLoginSwitchConfigResponse {
	s.Body = v
	return s
}

type ModifyNoticeConfigRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Route     *int32  `json:"Route,omitempty" xml:"Route,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TimeLimit *int32  `json:"TimeLimit,omitempty" xml:"TimeLimit,omitempty"`
}

func (s ModifyNoticeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNoticeConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyNoticeConfigRequest) SetSourceIp(v string) *ModifyNoticeConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetRoute(v int32) *ModifyNoticeConfigRequest {
	s.Route = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetProject(v string) *ModifyNoticeConfigRequest {
	s.Project = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetTimeLimit(v int32) *ModifyNoticeConfigRequest {
	s.TimeLimit = &v
	return s
}

type ModifyNoticeConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNoticeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNoticeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNoticeConfigResponseBody) SetRequestId(v string) *ModifyNoticeConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNoticeConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyNoticeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNoticeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNoticeConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyNoticeConfigResponse) SetHeaders(v map[string]*string) *ModifyNoticeConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyNoticeConfigResponse) SetBody(v *ModifyNoticeConfigResponseBody) *ModifyNoticeConfigResponse {
	s.Body = v
	return s
}

type ModifyOperateVulRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Info        *string `json:"Info,omitempty" xml:"Info,omitempty"`
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ModifyOperateVulRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulRequest) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequest) SetSourceIp(v string) *ModifyOperateVulRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyOperateVulRequest) SetInfo(v string) *ModifyOperateVulRequest {
	s.Info = &v
	return s
}

func (s *ModifyOperateVulRequest) SetOperateType(v string) *ModifyOperateVulRequest {
	s.OperateType = &v
	return s
}

func (s *ModifyOperateVulRequest) SetType(v string) *ModifyOperateVulRequest {
	s.Type = &v
	return s
}

func (s *ModifyOperateVulRequest) SetReason(v string) *ModifyOperateVulRequest {
	s.Reason = &v
	return s
}

type ModifyOperateVulResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOperateVulResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponseBody) SetRequestId(v string) *ModifyOperateVulResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOperateVulResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyOperateVulResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyOperateVulResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOperateVulResponse) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulResponse) SetHeaders(v map[string]*string) *ModifyOperateVulResponse {
	s.Headers = v
	return s
}

func (s *ModifyOperateVulResponse) SetBody(v *ModifyOperateVulResponseBody) *ModifyOperateVulResponse {
	s.Body = v
	return s
}

type ModifyPushAllTaskRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuids    *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	Tasks    *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s ModifyPushAllTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPushAllTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskRequest) SetSourceIp(v string) *ModifyPushAllTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyPushAllTaskRequest) SetUuids(v string) *ModifyPushAllTaskRequest {
	s.Uuids = &v
	return s
}

func (s *ModifyPushAllTaskRequest) SetTasks(v string) *ModifyPushAllTaskRequest {
	s.Tasks = &v
	return s
}

type ModifyPushAllTaskResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PushTaskRsp *ModifyPushAllTaskResponseBodyPushTaskRsp `json:"PushTaskRsp,omitempty" xml:"PushTaskRsp,omitempty" type:"Struct"`
}

func (s ModifyPushAllTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPushAllTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponseBody) SetRequestId(v string) *ModifyPushAllTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPushAllTaskResponseBody) SetPushTaskRsp(v *ModifyPushAllTaskResponseBodyPushTaskRsp) *ModifyPushAllTaskResponseBody {
	s.PushTaskRsp = v
	return s
}

type ModifyPushAllTaskResponseBodyPushTaskRsp struct {
	PushTaskResultList []*ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList `json:"PushTaskResultList,omitempty" xml:"PushTaskResultList,omitempty" type:"Repeated"`
}

func (s ModifyPushAllTaskResponseBodyPushTaskRsp) String() string {
	return tea.Prettify(s)
}

func (s ModifyPushAllTaskResponseBodyPushTaskRsp) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRsp) SetPushTaskResultList(v []*ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) *ModifyPushAllTaskResponseBodyPushTaskRsp {
	s.PushTaskResultList = v
	return s
}

type ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList struct {
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	GroupId      *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Online       *bool   `json:"Online,omitempty" xml:"Online,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	OsVersion    *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) String() string {
	return tea.Prettify(s)
}

func (s ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetUuid(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Uuid = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetGroupId(v int64) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.GroupId = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetSuccess(v bool) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Success = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetRegion(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Region = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetInstanceName(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.InstanceName = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetOnline(v bool) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Online = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetMessage(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Message = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetIp(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.Ip = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetOsVersion(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.OsVersion = &v
	return s
}

func (s *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList) SetInstanceId(v string) *ModifyPushAllTaskResponseBodyPushTaskRspPushTaskResultList {
	s.InstanceId = &v
	return s
}

type ModifyPushAllTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPushAllTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPushAllTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPushAllTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponse) SetHeaders(v map[string]*string) *ModifyPushAllTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyPushAllTaskResponse) SetBody(v *ModifyPushAllTaskResponseBody) *ModifyPushAllTaskResponse {
	s.Body = v
	return s
}

type ModifyRiskCheckStatusRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ItemId          *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	TaskId          *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyRiskCheckStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRiskCheckStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyRiskCheckStatusRequest) SetSourceIp(v string) *ModifyRiskCheckStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyRiskCheckStatusRequest) SetResourceOwnerId(v int64) *ModifyRiskCheckStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRiskCheckStatusRequest) SetLang(v string) *ModifyRiskCheckStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRiskCheckStatusRequest) SetItemId(v int64) *ModifyRiskCheckStatusRequest {
	s.ItemId = &v
	return s
}

func (s *ModifyRiskCheckStatusRequest) SetTaskId(v int64) *ModifyRiskCheckStatusRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyRiskCheckStatusRequest) SetStatus(v string) *ModifyRiskCheckStatusRequest {
	s.Status = &v
	return s
}

type ModifyRiskCheckStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRiskCheckStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRiskCheckStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRiskCheckStatusResponseBody) SetRequestId(v string) *ModifyRiskCheckStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRiskCheckStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRiskCheckStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRiskCheckStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRiskCheckStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyRiskCheckStatusResponse) SetHeaders(v map[string]*string) *ModifyRiskCheckStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyRiskCheckStatusResponse) SetBody(v *ModifyRiskCheckStatusResponseBody) *ModifyRiskCheckStatusResponse {
	s.Body = v
	return s
}

type ModifyRiskSingleResultStatusRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Status          *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *int64    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Ids             []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s ModifyRiskSingleResultStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRiskSingleResultStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyRiskSingleResultStatusRequest) SetSourceIp(v string) *ModifyRiskSingleResultStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyRiskSingleResultStatusRequest) SetResourceOwnerId(v int64) *ModifyRiskSingleResultStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRiskSingleResultStatusRequest) SetLang(v string) *ModifyRiskSingleResultStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRiskSingleResultStatusRequest) SetStatus(v string) *ModifyRiskSingleResultStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyRiskSingleResultStatusRequest) SetTaskId(v int64) *ModifyRiskSingleResultStatusRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyRiskSingleResultStatusRequest) SetIds(v []*string) *ModifyRiskSingleResultStatusRequest {
	s.Ids = v
	return s
}

type ModifyRiskSingleResultStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRiskSingleResultStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRiskSingleResultStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRiskSingleResultStatusResponseBody) SetRequestId(v string) *ModifyRiskSingleResultStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRiskSingleResultStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRiskSingleResultStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRiskSingleResultStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRiskSingleResultStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyRiskSingleResultStatusResponse) SetHeaders(v map[string]*string) *ModifyRiskSingleResultStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyRiskSingleResultStatusResponse) SetBody(v *ModifyRiskSingleResultStatusResponseBody) *ModifyRiskSingleResultStatusResponse {
	s.Body = v
	return s
}

type ModifySecurityCheckScheduleConfigRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	DaysOfWeek      *string `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty"`
	StartTime       *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ModifySecurityCheckScheduleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityCheckScheduleConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetSourceIp(v string) *ModifySecurityCheckScheduleConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetResourceOwnerId(v int64) *ModifySecurityCheckScheduleConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetLang(v string) *ModifySecurityCheckScheduleConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetDaysOfWeek(v string) *ModifySecurityCheckScheduleConfigRequest {
	s.DaysOfWeek = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetStartTime(v int32) *ModifySecurityCheckScheduleConfigRequest {
	s.StartTime = &v
	return s
}

func (s *ModifySecurityCheckScheduleConfigRequest) SetEndTime(v int32) *ModifySecurityCheckScheduleConfigRequest {
	s.EndTime = &v
	return s
}

type ModifySecurityCheckScheduleConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityCheckScheduleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityCheckScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityCheckScheduleConfigResponseBody) SetRequestId(v string) *ModifySecurityCheckScheduleConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifySecurityCheckScheduleConfigResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySecurityCheckScheduleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityCheckScheduleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityCheckScheduleConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityCheckScheduleConfigResponse) SetHeaders(v map[string]*string) *ModifySecurityCheckScheduleConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityCheckScheduleConfigResponse) SetBody(v *ModifySecurityCheckScheduleConfigResponseBody) *ModifySecurityCheckScheduleConfigResponse {
	s.Body = v
	return s
}

type ModifyStartVulScanRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Types    *string `json:"Types,omitempty" xml:"Types,omitempty"`
	Uuids    *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ModifyStartVulScanRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyStartVulScanRequest) GoString() string {
	return s.String()
}

func (s *ModifyStartVulScanRequest) SetSourceIp(v string) *ModifyStartVulScanRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyStartVulScanRequest) SetTypes(v string) *ModifyStartVulScanRequest {
	s.Types = &v
	return s
}

func (s *ModifyStartVulScanRequest) SetUuids(v string) *ModifyStartVulScanRequest {
	s.Uuids = &v
	return s
}

type ModifyStartVulScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStartVulScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyStartVulScanResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStartVulScanResponseBody) SetRequestId(v string) *ModifyStartVulScanResponseBody {
	s.RequestId = &v
	return s
}

type ModifyStartVulScanResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyStartVulScanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyStartVulScanResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyStartVulScanResponse) GoString() string {
	return s.String()
}

func (s *ModifyStartVulScanResponse) SetHeaders(v map[string]*string) *ModifyStartVulScanResponse {
	s.Headers = v
	return s
}

func (s *ModifyStartVulScanResponse) SetBody(v *ModifyStartVulScanResponseBody) *ModifyStartVulScanResponse {
	s.Body = v
	return s
}

type ModifyTagWithUuidRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	UuidList     *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
	TagList      *string `json:"TagList,omitempty" xml:"TagList,omitempty"`
	TagId        *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
}

func (s ModifyTagWithUuidRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTagWithUuidRequest) GoString() string {
	return s.String()
}

func (s *ModifyTagWithUuidRequest) SetSourceIp(v string) *ModifyTagWithUuidRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetUuidList(v string) *ModifyTagWithUuidRequest {
	s.UuidList = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetTagList(v string) *ModifyTagWithUuidRequest {
	s.TagList = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetTagId(v string) *ModifyTagWithUuidRequest {
	s.TagId = &v
	return s
}

func (s *ModifyTagWithUuidRequest) SetMachineTypes(v string) *ModifyTagWithUuidRequest {
	s.MachineTypes = &v
	return s
}

type ModifyTagWithUuidResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTagWithUuidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTagWithUuidResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTagWithUuidResponseBody) SetRequestId(v string) *ModifyTagWithUuidResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTagWithUuidResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTagWithUuidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTagWithUuidResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTagWithUuidResponse) GoString() string {
	return s.String()
}

func (s *ModifyTagWithUuidResponse) SetHeaders(v map[string]*string) *ModifyTagWithUuidResponse {
	s.Headers = v
	return s
}

func (s *ModifyTagWithUuidResponse) SetBody(v *ModifyTagWithUuidResponseBody) *ModifyTagWithUuidResponse {
	s.Body = v
	return s
}

type ModifyVpcHoneyPotRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VpcId          *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	HoneyPotAction *string `json:"HoneyPotAction,omitempty" xml:"HoneyPotAction,omitempty"`
}

func (s ModifyVpcHoneyPotRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcHoneyPotRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcHoneyPotRequest) SetSourceIp(v string) *ModifyVpcHoneyPotRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVpcHoneyPotRequest) SetVpcId(v string) *ModifyVpcHoneyPotRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyVpcHoneyPotRequest) SetHoneyPotAction(v string) *ModifyVpcHoneyPotRequest {
	s.HoneyPotAction = &v
	return s
}

type ModifyVpcHoneyPotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcHoneyPotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcHoneyPotResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcHoneyPotResponseBody) SetRequestId(v string) *ModifyVpcHoneyPotResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcHoneyPotResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyVpcHoneyPotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVpcHoneyPotResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcHoneyPotResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcHoneyPotResponse) SetHeaders(v map[string]*string) *ModifyVpcHoneyPotResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcHoneyPotResponse) SetBody(v *ModifyVpcHoneyPotResponseBody) *ModifyVpcHoneyPotResponse {
	s.Body = v
	return s
}

type ModifyVulTargetConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Config   *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ModifyVulTargetConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVulTargetConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetConfigRequest) SetSourceIp(v string) *ModifyVulTargetConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVulTargetConfigRequest) SetType(v string) *ModifyVulTargetConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyVulTargetConfigRequest) SetUuid(v string) *ModifyVulTargetConfigRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyVulTargetConfigRequest) SetConfig(v string) *ModifyVulTargetConfigRequest {
	s.Config = &v
	return s
}

type ModifyVulTargetConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVulTargetConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVulTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetConfigResponseBody) SetRequestId(v string) *ModifyVulTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVulTargetConfigResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyVulTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVulTargetConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVulTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetConfigResponse) SetHeaders(v map[string]*string) *ModifyVulTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyVulTargetConfigResponse) SetBody(v *ModifyVulTargetConfigResponseBody) *ModifyVulTargetConfigResponse {
	s.Body = v
	return s
}

type ModifyWebLockCreateConfigRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Dir               *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	ExclusiveDir      *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	LocalBackupDir    *string `json:"LocalBackupDir,omitempty" xml:"LocalBackupDir,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	ExclusiveFile     *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	InclusiveFile     *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
	DefenceMode       *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
}

func (s ModifyWebLockCreateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockCreateConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockCreateConfigRequest) SetSourceIp(v string) *ModifyWebLockCreateConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetLang(v string) *ModifyWebLockCreateConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetUuid(v string) *ModifyWebLockCreateConfigRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetDir(v string) *ModifyWebLockCreateConfigRequest {
	s.Dir = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetExclusiveDir(v string) *ModifyWebLockCreateConfigRequest {
	s.ExclusiveDir = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetExclusiveFileType(v string) *ModifyWebLockCreateConfigRequest {
	s.ExclusiveFileType = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetLocalBackupDir(v string) *ModifyWebLockCreateConfigRequest {
	s.LocalBackupDir = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetMode(v string) *ModifyWebLockCreateConfigRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetInclusiveFileType(v string) *ModifyWebLockCreateConfigRequest {
	s.InclusiveFileType = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetExclusiveFile(v string) *ModifyWebLockCreateConfigRequest {
	s.ExclusiveFile = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetInclusiveFile(v string) *ModifyWebLockCreateConfigRequest {
	s.InclusiveFile = &v
	return s
}

func (s *ModifyWebLockCreateConfigRequest) SetDefenceMode(v string) *ModifyWebLockCreateConfigRequest {
	s.DefenceMode = &v
	return s
}

type ModifyWebLockCreateConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockCreateConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockCreateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockCreateConfigResponseBody) SetRequestId(v string) *ModifyWebLockCreateConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebLockCreateConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebLockCreateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebLockCreateConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockCreateConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockCreateConfigResponse) SetHeaders(v map[string]*string) *ModifyWebLockCreateConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockCreateConfigResponse) SetBody(v *ModifyWebLockCreateConfigResponseBody) *ModifyWebLockCreateConfigResponse {
	s.Body = v
	return s
}

type ModifyWebLockDeleteConfigRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockDeleteConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockDeleteConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockDeleteConfigRequest) SetSourceIp(v string) *ModifyWebLockDeleteConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockDeleteConfigRequest) SetLang(v string) *ModifyWebLockDeleteConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWebLockDeleteConfigRequest) SetId(v int32) *ModifyWebLockDeleteConfigRequest {
	s.Id = &v
	return s
}

func (s *ModifyWebLockDeleteConfigRequest) SetUuid(v string) *ModifyWebLockDeleteConfigRequest {
	s.Uuid = &v
	return s
}

type ModifyWebLockDeleteConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockDeleteConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockDeleteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockDeleteConfigResponseBody) SetRequestId(v string) *ModifyWebLockDeleteConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebLockDeleteConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebLockDeleteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebLockDeleteConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockDeleteConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockDeleteConfigResponse) SetHeaders(v map[string]*string) *ModifyWebLockDeleteConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockDeleteConfigResponse) SetBody(v *ModifyWebLockDeleteConfigResponseBody) *ModifyWebLockDeleteConfigResponse {
	s.Body = v
	return s
}

type ModifyWebLockStartRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	LocalBackupDir    *string `json:"LocalBackupDir,omitempty" xml:"LocalBackupDir,omitempty"`
	ExclusiveFile     *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	Dir               *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	ExclusiveDir      *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	DefenceMode       *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
}

func (s ModifyWebLockStartRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockStartRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStartRequest) SetSourceIp(v string) *ModifyWebLockStartRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetMode(v string) *ModifyWebLockStartRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetLocalBackupDir(v string) *ModifyWebLockStartRequest {
	s.LocalBackupDir = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetExclusiveFile(v string) *ModifyWebLockStartRequest {
	s.ExclusiveFile = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetDir(v string) *ModifyWebLockStartRequest {
	s.Dir = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetInclusiveFileType(v string) *ModifyWebLockStartRequest {
	s.InclusiveFileType = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetUuid(v string) *ModifyWebLockStartRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetExclusiveFileType(v string) *ModifyWebLockStartRequest {
	s.ExclusiveFileType = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetExclusiveDir(v string) *ModifyWebLockStartRequest {
	s.ExclusiveDir = &v
	return s
}

func (s *ModifyWebLockStartRequest) SetDefenceMode(v string) *ModifyWebLockStartRequest {
	s.DefenceMode = &v
	return s
}

type ModifyWebLockStartResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockStartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockStartResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStartResponseBody) SetRequestId(v string) *ModifyWebLockStartResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebLockStartResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebLockStartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebLockStartResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockStartResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStartResponse) SetHeaders(v map[string]*string) *ModifyWebLockStartResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockStartResponse) SetBody(v *ModifyWebLockStartResponseBody) *ModifyWebLockStartResponse {
	s.Body = v
	return s
}

type ModifyWebLockUnbindRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockUnbindRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockUnbindRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUnbindRequest) SetSourceIp(v string) *ModifyWebLockUnbindRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockUnbindRequest) SetUuid(v string) *ModifyWebLockUnbindRequest {
	s.Uuid = &v
	return s
}

type ModifyWebLockUnbindResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockUnbindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockUnbindResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUnbindResponseBody) SetRequestId(v string) *ModifyWebLockUnbindResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebLockUnbindResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebLockUnbindResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebLockUnbindResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockUnbindResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUnbindResponse) SetHeaders(v map[string]*string) *ModifyWebLockUnbindResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockUnbindResponse) SetBody(v *ModifyWebLockUnbindResponseBody) *ModifyWebLockUnbindResponse {
	s.Body = v
	return s
}

type ModifyWebLockUpdateConfigRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id                *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Uuid              *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Dir               *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	ExclusiveDir      *string `json:"ExclusiveDir,omitempty" xml:"ExclusiveDir,omitempty"`
	ExclusiveFileType *string `json:"ExclusiveFileType,omitempty" xml:"ExclusiveFileType,omitempty"`
	LocalBackupDir    *string `json:"LocalBackupDir,omitempty" xml:"LocalBackupDir,omitempty"`
	Mode              *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	InclusiveFileType *string `json:"InclusiveFileType,omitempty" xml:"InclusiveFileType,omitempty"`
	ExclusiveFile     *string `json:"ExclusiveFile,omitempty" xml:"ExclusiveFile,omitempty"`
	InclusiveFile     *string `json:"InclusiveFile,omitempty" xml:"InclusiveFile,omitempty"`
	DefenceMode       *string `json:"DefenceMode,omitempty" xml:"DefenceMode,omitempty"`
}

func (s ModifyWebLockUpdateConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockUpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUpdateConfigRequest) SetSourceIp(v string) *ModifyWebLockUpdateConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetLang(v string) *ModifyWebLockUpdateConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetId(v int32) *ModifyWebLockUpdateConfigRequest {
	s.Id = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetUuid(v string) *ModifyWebLockUpdateConfigRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetDir(v string) *ModifyWebLockUpdateConfigRequest {
	s.Dir = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetExclusiveDir(v string) *ModifyWebLockUpdateConfigRequest {
	s.ExclusiveDir = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetExclusiveFileType(v string) *ModifyWebLockUpdateConfigRequest {
	s.ExclusiveFileType = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetLocalBackupDir(v string) *ModifyWebLockUpdateConfigRequest {
	s.LocalBackupDir = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetMode(v string) *ModifyWebLockUpdateConfigRequest {
	s.Mode = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetInclusiveFileType(v string) *ModifyWebLockUpdateConfigRequest {
	s.InclusiveFileType = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetExclusiveFile(v string) *ModifyWebLockUpdateConfigRequest {
	s.ExclusiveFile = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetInclusiveFile(v string) *ModifyWebLockUpdateConfigRequest {
	s.InclusiveFile = &v
	return s
}

func (s *ModifyWebLockUpdateConfigRequest) SetDefenceMode(v string) *ModifyWebLockUpdateConfigRequest {
	s.DefenceMode = &v
	return s
}

type ModifyWebLockUpdateConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockUpdateConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockUpdateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUpdateConfigResponseBody) SetRequestId(v string) *ModifyWebLockUpdateConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyWebLockUpdateConfigResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyWebLockUpdateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyWebLockUpdateConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyWebLockUpdateConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUpdateConfigResponse) SetHeaders(v map[string]*string) *ModifyWebLockUpdateConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebLockUpdateConfigResponse) SetBody(v *ModifyWebLockUpdateConfigResponseBody) *ModifyWebLockUpdateConfigResponse {
	s.Body = v
	return s
}

type OperateSuspiciousTargetConfigRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	TargetType       *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TargetOperations *string `json:"TargetOperations,omitempty" xml:"TargetOperations,omitempty"`
}

func (s OperateSuspiciousTargetConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateSuspiciousTargetConfigRequest) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousTargetConfigRequest) SetSourceIp(v string) *OperateSuspiciousTargetConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetLang(v string) *OperateSuspiciousTargetConfigRequest {
	s.Lang = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetType(v string) *OperateSuspiciousTargetConfigRequest {
	s.Type = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetTargetType(v string) *OperateSuspiciousTargetConfigRequest {
	s.TargetType = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetTargetOperations(v string) *OperateSuspiciousTargetConfigRequest {
	s.TargetOperations = &v
	return s
}

type OperateSuspiciousTargetConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateSuspiciousTargetConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateSuspiciousTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousTargetConfigResponseBody) SetRequestId(v string) *OperateSuspiciousTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

type OperateSuspiciousTargetConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperateSuspiciousTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperateSuspiciousTargetConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateSuspiciousTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousTargetConfigResponse) SetHeaders(v map[string]*string) *OperateSuspiciousTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *OperateSuspiciousTargetConfigResponse) SetBody(v *OperateSuspiciousTargetConfigResponseBody) *OperateSuspiciousTargetConfigResponse {
	s.Body = v
	return s
}

type OperationSuspEventsRequest struct {
	SourceIp           *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SuspiciousEventIds *string `json:"SuspiciousEventIds,omitempty" xml:"SuspiciousEventIds,omitempty"`
	Operation          *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	SubOperation       *string `json:"SubOperation,omitempty" xml:"SubOperation,omitempty"`
	From               *string `json:"From,omitempty" xml:"From,omitempty"`
	WarnType           *string `json:"WarnType,omitempty" xml:"WarnType,omitempty"`
}

func (s OperationSuspEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s OperationSuspEventsRequest) GoString() string {
	return s.String()
}

func (s *OperationSuspEventsRequest) SetSourceIp(v string) *OperationSuspEventsRequest {
	s.SourceIp = &v
	return s
}

func (s *OperationSuspEventsRequest) SetSuspiciousEventIds(v string) *OperationSuspEventsRequest {
	s.SuspiciousEventIds = &v
	return s
}

func (s *OperationSuspEventsRequest) SetOperation(v string) *OperationSuspEventsRequest {
	s.Operation = &v
	return s
}

func (s *OperationSuspEventsRequest) SetSubOperation(v string) *OperationSuspEventsRequest {
	s.SubOperation = &v
	return s
}

func (s *OperationSuspEventsRequest) SetFrom(v string) *OperationSuspEventsRequest {
	s.From = &v
	return s
}

func (s *OperationSuspEventsRequest) SetWarnType(v string) *OperationSuspEventsRequest {
	s.WarnType = &v
	return s
}

type OperationSuspEventsResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessCode *string `json:"AccessCode,omitempty" xml:"AccessCode,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperationSuspEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperationSuspEventsResponseBody) GoString() string {
	return s.String()
}

func (s *OperationSuspEventsResponseBody) SetRequestId(v string) *OperationSuspEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperationSuspEventsResponseBody) SetAccessCode(v string) *OperationSuspEventsResponseBody {
	s.AccessCode = &v
	return s
}

func (s *OperationSuspEventsResponseBody) SetSuccess(v bool) *OperationSuspEventsResponseBody {
	s.Success = &v
	return s
}

type OperationSuspEventsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OperationSuspEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OperationSuspEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s OperationSuspEventsResponse) GoString() string {
	return s.String()
}

func (s *OperationSuspEventsResponse) SetHeaders(v map[string]*string) *OperationSuspEventsResponse {
	s.Headers = v
	return s
}

func (s *OperationSuspEventsResponse) SetBody(v *OperationSuspEventsResponseBody) *OperationSuspEventsResponse {
	s.Body = v
	return s
}

type PauseClientRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuids    *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s PauseClientRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseClientRequest) GoString() string {
	return s.String()
}

func (s *PauseClientRequest) SetSourceIp(v string) *PauseClientRequest {
	s.SourceIp = &v
	return s
}

func (s *PauseClientRequest) SetUuids(v string) *PauseClientRequest {
	s.Uuids = &v
	return s
}

func (s *PauseClientRequest) SetValue(v string) *PauseClientRequest {
	s.Value = &v
	return s
}

func (s *PauseClientRequest) SetAppName(v string) *PauseClientRequest {
	s.AppName = &v
	return s
}

type PauseClientResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseClientResponseBody) GoString() string {
	return s.String()
}

func (s *PauseClientResponseBody) SetRequestId(v string) *PauseClientResponseBody {
	s.RequestId = &v
	return s
}

type PauseClientResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PauseClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PauseClientResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseClientResponse) GoString() string {
	return s.String()
}

func (s *PauseClientResponse) SetHeaders(v map[string]*string) *PauseClientResponse {
	s.Headers = v
	return s
}

func (s *PauseClientResponse) SetBody(v *PauseClientResponseBody) *PauseClientResponse {
	s.Body = v
	return s
}

type RefreshContainerAssetsRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
}

func (s RefreshContainerAssetsRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshContainerAssetsRequest) GoString() string {
	return s.String()
}

func (s *RefreshContainerAssetsRequest) SetSourceIp(v string) *RefreshContainerAssetsRequest {
	s.SourceIp = &v
	return s
}

func (s *RefreshContainerAssetsRequest) SetAssetType(v string) *RefreshContainerAssetsRequest {
	s.AssetType = &v
	return s
}

type RefreshContainerAssetsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshContainerAssetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshContainerAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshContainerAssetsResponseBody) SetRequestId(v string) *RefreshContainerAssetsResponseBody {
	s.RequestId = &v
	return s
}

type RefreshContainerAssetsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshContainerAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshContainerAssetsResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshContainerAssetsResponse) GoString() string {
	return s.String()
}

func (s *RefreshContainerAssetsResponse) SetHeaders(v map[string]*string) *RefreshContainerAssetsResponse {
	s.Headers = v
	return s
}

func (s *RefreshContainerAssetsResponse) SetBody(v *RefreshContainerAssetsResponseBody) *RefreshContainerAssetsResponse {
	s.Body = v
	return s
}

type RollbackSuspEventQuaraFileRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	QuaraFileId *int32  `json:"QuaraFileId,omitempty" xml:"QuaraFileId,omitempty"`
	From        *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s RollbackSuspEventQuaraFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileRequest) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileRequest) SetSourceIp(v string) *RollbackSuspEventQuaraFileRequest {
	s.SourceIp = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetQuaraFileId(v int32) *RollbackSuspEventQuaraFileRequest {
	s.QuaraFileId = &v
	return s
}

func (s *RollbackSuspEventQuaraFileRequest) SetFrom(v string) *RollbackSuspEventQuaraFileRequest {
	s.From = &v
	return s
}

type RollbackSuspEventQuaraFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackSuspEventQuaraFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponseBody) SetRequestId(v string) *RollbackSuspEventQuaraFileResponseBody {
	s.RequestId = &v
	return s
}

type RollbackSuspEventQuaraFileResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollbackSuspEventQuaraFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackSuspEventQuaraFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackSuspEventQuaraFileResponse) GoString() string {
	return s.String()
}

func (s *RollbackSuspEventQuaraFileResponse) SetHeaders(v map[string]*string) *RollbackSuspEventQuaraFileResponse {
	s.Headers = v
	return s
}

func (s *RollbackSuspEventQuaraFileResponse) SetBody(v *RollbackSuspEventQuaraFileResponseBody) *RollbackSuspEventQuaraFileResponse {
	s.Body = v
	return s
}

type SasInstallCodeRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s SasInstallCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SasInstallCodeRequest) GoString() string {
	return s.String()
}

func (s *SasInstallCodeRequest) SetSourceIp(v string) *SasInstallCodeRequest {
	s.SourceIp = &v
	return s
}

type SasInstallCodeResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SasInstallCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SasInstallCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SasInstallCodeResponseBody) SetData(v string) *SasInstallCodeResponseBody {
	s.Data = &v
	return s
}

func (s *SasInstallCodeResponseBody) SetRequestId(v string) *SasInstallCodeResponseBody {
	s.RequestId = &v
	return s
}

type SasInstallCodeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SasInstallCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SasInstallCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SasInstallCodeResponse) GoString() string {
	return s.String()
}

func (s *SasInstallCodeResponse) SetHeaders(v map[string]*string) *SasInstallCodeResponse {
	s.Headers = v
	return s
}

func (s *SasInstallCodeResponse) SetBody(v *SasInstallCodeResponseBody) *SasInstallCodeResponse {
	s.Body = v
	return s
}

type StartBaselineSecurityCheckRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Lang            *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Type            *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	ItemIds         []*int    `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Repeated"`
	Assets          []*string `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
}

func (s StartBaselineSecurityCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s StartBaselineSecurityCheckRequest) GoString() string {
	return s.String()
}

func (s *StartBaselineSecurityCheckRequest) SetSourceIp(v string) *StartBaselineSecurityCheckRequest {
	s.SourceIp = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetResourceOwnerId(v int64) *StartBaselineSecurityCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetLang(v string) *StartBaselineSecurityCheckRequest {
	s.Lang = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetType(v string) *StartBaselineSecurityCheckRequest {
	s.Type = &v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetItemIds(v []*int) *StartBaselineSecurityCheckRequest {
	s.ItemIds = v
	return s
}

func (s *StartBaselineSecurityCheckRequest) SetAssets(v []*string) *StartBaselineSecurityCheckRequest {
	s.Assets = v
	return s
}

type StartBaselineSecurityCheckResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartBaselineSecurityCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartBaselineSecurityCheckResponseBody) GoString() string {
	return s.String()
}

func (s *StartBaselineSecurityCheckResponseBody) SetRequestId(v string) *StartBaselineSecurityCheckResponseBody {
	s.RequestId = &v
	return s
}

type StartBaselineSecurityCheckResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartBaselineSecurityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartBaselineSecurityCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBaselineSecurityCheckResponse) GoString() string {
	return s.String()
}

func (s *StartBaselineSecurityCheckResponse) SetHeaders(v map[string]*string) *StartBaselineSecurityCheckResponse {
	s.Headers = v
	return s
}

func (s *StartBaselineSecurityCheckResponse) SetBody(v *StartBaselineSecurityCheckResponseBody) *StartBaselineSecurityCheckResponse {
	s.Body = v
	return s
}

type StartImageVulScanRequest struct {
	SourceIp       *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RepoRegionId   *string   `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	RepoInstanceId *string   `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	RepoId         *string   `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepName        *string   `json:"RepName,omitempty" xml:"RepName,omitempty"`
	RepoNamespace  *string   `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	ImageTag       *string   `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ImageDigest    *string   `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	ImageLayer     *string   `json:"ImageLayer,omitempty" xml:"ImageLayer,omitempty"`
	RegistryTypes  []*string `json:"RegistryTypes,omitempty" xml:"RegistryTypes,omitempty" type:"Repeated"`
}

func (s StartImageVulScanRequest) String() string {
	return tea.Prettify(s)
}

func (s StartImageVulScanRequest) GoString() string {
	return s.String()
}

func (s *StartImageVulScanRequest) SetSourceIp(v string) *StartImageVulScanRequest {
	s.SourceIp = &v
	return s
}

func (s *StartImageVulScanRequest) SetLang(v string) *StartImageVulScanRequest {
	s.Lang = &v
	return s
}

func (s *StartImageVulScanRequest) SetRepoRegionId(v string) *StartImageVulScanRequest {
	s.RepoRegionId = &v
	return s
}

func (s *StartImageVulScanRequest) SetRepoInstanceId(v string) *StartImageVulScanRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *StartImageVulScanRequest) SetRepoId(v string) *StartImageVulScanRequest {
	s.RepoId = &v
	return s
}

func (s *StartImageVulScanRequest) SetRepName(v string) *StartImageVulScanRequest {
	s.RepName = &v
	return s
}

func (s *StartImageVulScanRequest) SetRepoNamespace(v string) *StartImageVulScanRequest {
	s.RepoNamespace = &v
	return s
}

func (s *StartImageVulScanRequest) SetImageTag(v string) *StartImageVulScanRequest {
	s.ImageTag = &v
	return s
}

func (s *StartImageVulScanRequest) SetImageDigest(v string) *StartImageVulScanRequest {
	s.ImageDigest = &v
	return s
}

func (s *StartImageVulScanRequest) SetImageLayer(v string) *StartImageVulScanRequest {
	s.ImageLayer = &v
	return s
}

func (s *StartImageVulScanRequest) SetRegistryTypes(v []*string) *StartImageVulScanRequest {
	s.RegistryTypes = v
	return s
}

type StartImageVulScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartImageVulScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartImageVulScanResponseBody) GoString() string {
	return s.String()
}

func (s *StartImageVulScanResponseBody) SetRequestId(v string) *StartImageVulScanResponseBody {
	s.RequestId = &v
	return s
}

type StartImageVulScanResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartImageVulScanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartImageVulScanResponse) String() string {
	return tea.Prettify(s)
}

func (s StartImageVulScanResponse) GoString() string {
	return s.String()
}

func (s *StartImageVulScanResponse) SetHeaders(v map[string]*string) *StartImageVulScanResponse {
	s.Headers = v
	return s
}

func (s *StartImageVulScanResponse) SetBody(v *StartImageVulScanResponseBody) *StartImageVulScanResponse {
	s.Body = v
	return s
}

type ValidateHcWarningsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RiskIds  *string `json:"RiskIds,omitempty" xml:"RiskIds,omitempty"`
	Uuids    *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ValidateHcWarningsRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateHcWarningsRequest) GoString() string {
	return s.String()
}

func (s *ValidateHcWarningsRequest) SetSourceIp(v string) *ValidateHcWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *ValidateHcWarningsRequest) SetRiskIds(v string) *ValidateHcWarningsRequest {
	s.RiskIds = &v
	return s
}

func (s *ValidateHcWarningsRequest) SetUuids(v string) *ValidateHcWarningsRequest {
	s.Uuids = &v
	return s
}

type ValidateHcWarningsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateHcWarningsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateHcWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateHcWarningsResponseBody) SetRequestId(v string) *ValidateHcWarningsResponseBody {
	s.RequestId = &v
	return s
}

type ValidateHcWarningsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateHcWarningsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateHcWarningsResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateHcWarningsResponse) GoString() string {
	return s.String()
}

func (s *ValidateHcWarningsResponse) SetHeaders(v map[string]*string) *ValidateHcWarningsResponse {
	s.Headers = v
	return s
}

func (s *ValidateHcWarningsResponse) SetBody(v *ValidateHcWarningsResponseBody) *ValidateHcWarningsResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":         tea.String("tds.aliyuncs.com"),
		"ap-southeast-3":      tea.String("tds.ap-southeast-3.aliyuncs.com"),
		"cn-shanghai-et2-b01": tea.String("tds.cn-shanghai-et2-b01.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddVpcHoneyPotWithOptions(request *AddVpcHoneyPotRequest, runtime *util.RuntimeOptions) (_result *AddVpcHoneyPotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddVpcHoneyPotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddVpcHoneyPot"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVpcHoneyPot(request *AddVpcHoneyPotRequest) (_result *AddVpcHoneyPotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVpcHoneyPotResponse{}
	_body, _err := client.AddVpcHoneyPotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAntiBruteForceRuleWithOptions(request *CreateAntiBruteForceRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAntiBruteForceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAntiBruteForceRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAntiBruteForceRule"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAntiBruteForceRule(request *CreateAntiBruteForceRuleRequest) (_result *CreateAntiBruteForceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAntiBruteForceRuleResponse{}
	_body, _err := client.CreateAntiBruteForceRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateAssetGroupWithOptions(request *CreateOrUpdateAssetGroupRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateAssetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOrUpdateAssetGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOrUpdateAssetGroup"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateAssetGroup(request *CreateOrUpdateAssetGroupRequest) (_result *CreateOrUpdateAssetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateAssetGroupResponse{}
	_body, _err := client.CreateOrUpdateAssetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSasOrderWithOptions(request *CreateSasOrderRequest, runtime *util.RuntimeOptions) (_result *CreateSasOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSasOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSasOrder"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSasOrder(request *CreateSasOrderRequest) (_result *CreateSasOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSasOrderResponse{}
	_body, _err := client.CreateSasOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSimilarSecurityEventsQueryTaskWithOptions(request *CreateSimilarSecurityEventsQueryTaskRequest, runtime *util.RuntimeOptions) (_result *CreateSimilarSecurityEventsQueryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSimilarSecurityEventsQueryTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSimilarSecurityEventsQueryTask"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSimilarSecurityEventsQueryTask(request *CreateSimilarSecurityEventsQueryTaskRequest) (_result *CreateSimilarSecurityEventsQueryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSimilarSecurityEventsQueryTaskResponse{}
	_body, _err := client.CreateSimilarSecurityEventsQueryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGroup"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoginBaseConfigWithOptions(request *DeleteLoginBaseConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLoginBaseConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLoginBaseConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLoginBaseConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLoginBaseConfig(request *DeleteLoginBaseConfigRequest) (_result *DeleteLoginBaseConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoginBaseConfigResponse{}
	_body, _err := client.DeleteLoginBaseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagWithUuidWithOptions(request *DeleteTagWithUuidRequest, runtime *util.RuntimeOptions) (_result *DeleteTagWithUuidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTagWithUuidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTagWithUuid"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTagWithUuid(request *DeleteTagWithUuidRequest) (_result *DeleteTagWithUuidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagWithUuidResponse{}
	_body, _err := client.DeleteTagWithUuidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpcHoneyPotWithOptions(request *DeleteVpcHoneyPotRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcHoneyPotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVpcHoneyPotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVpcHoneyPot"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpcHoneyPot(request *DeleteVpcHoneyPotRequest) (_result *DeleteVpcHoneyPotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcHoneyPotResponse{}
	_body, _err := client.DeleteVpcHoneyPotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccesskeyLeakListWithOptions(request *DescribeAccesskeyLeakListRequest, runtime *util.RuntimeOptions) (_result *DescribeAccesskeyLeakListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccesskeyLeakListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccesskeyLeakList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccesskeyLeakList(request *DescribeAccesskeyLeakListRequest) (_result *DescribeAccesskeyLeakListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccesskeyLeakListResponse{}
	_body, _err := client.DescribeAccesskeyLeakListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAffectedMaliciousFileImagesWithOptions(request *DescribeAffectedMaliciousFileImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeAffectedMaliciousFileImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAffectedMaliciousFileImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAffectedMaliciousFileImages"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAffectedMaliciousFileImages(request *DescribeAffectedMaliciousFileImagesRequest) (_result *DescribeAffectedMaliciousFileImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAffectedMaliciousFileImagesResponse{}
	_body, _err := client.DescribeAffectedMaliciousFileImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlarmEventDetailWithOptions(request *DescribeAlarmEventDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmEventDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlarmEventDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlarmEventDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlarmEventDetail(request *DescribeAlarmEventDetailRequest) (_result *DescribeAlarmEventDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmEventDetailResponse{}
	_body, _err := client.DescribeAlarmEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAlarmEventListWithOptions(request *DescribeAlarmEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeAlarmEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAlarmEventListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAlarmEventList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAlarmEventList(request *DescribeAlarmEventListRequest) (_result *DescribeAlarmEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAlarmEventListResponse{}
	_body, _err := client.DescribeAlarmEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllEntityWithOptions(request *DescribeAllEntityRequest, runtime *util.RuntimeOptions) (_result *DescribeAllEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllEntity"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllEntity(request *DescribeAllEntityRequest) (_result *DescribeAllEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllEntityResponse{}
	_body, _err := client.DescribeAllEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllGroupsWithOptions(request *DescribeAllGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeAllGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllGroups"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllGroups(request *DescribeAllGroupsRequest) (_result *DescribeAllGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllGroupsResponse{}
	_body, _err := client.DescribeAllGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllRegionsStatisticsWithOptions(request *DescribeAllRegionsStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeAllRegionsStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllRegionsStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllRegionsStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllRegionsStatistics(request *DescribeAllRegionsStatisticsRequest) (_result *DescribeAllRegionsStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllRegionsStatisticsResponse{}
	_body, _err := client.DescribeAllRegionsStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAntiBruteForceRulesWithOptions(request *DescribeAntiBruteForceRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeAntiBruteForceRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAntiBruteForceRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAntiBruteForceRules"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAntiBruteForceRules(request *DescribeAntiBruteForceRulesRequest) (_result *DescribeAntiBruteForceRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAntiBruteForceRulesResponse{}
	_body, _err := client.DescribeAntiBruteForceRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssetDetailByUuidWithOptions(request *DescribeAssetDetailByUuidRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetDetailByUuidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAssetDetailByUuidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAssetDetailByUuid"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssetDetailByUuid(request *DescribeAssetDetailByUuidRequest) (_result *DescribeAssetDetailByUuidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetDetailByUuidResponse{}
	_body, _err := client.DescribeAssetDetailByUuidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssetDetailByUuidsWithOptions(request *DescribeAssetDetailByUuidsRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetDetailByUuidsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAssetDetailByUuidsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAssetDetailByUuids"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssetDetailByUuids(request *DescribeAssetDetailByUuidsRequest) (_result *DescribeAssetDetailByUuidsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetDetailByUuidsResponse{}
	_body, _err := client.DescribeAssetDetailByUuidsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoDelConfigWithOptions(request *DescribeAutoDelConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoDelConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoDelConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoDelConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoDelConfig(request *DescribeAutoDelConfigRequest) (_result *DescribeAutoDelConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoDelConfigResponse{}
	_body, _err := client.DescribeAutoDelConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBruteForceSummaryWithOptions(request *DescribeBruteForceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeBruteForceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBruteForceSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBruteForceSummary"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBruteForceSummary(request *DescribeBruteForceSummaryRequest) (_result *DescribeBruteForceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBruteForceSummaryResponse{}
	_body, _err := client.DescribeBruteForceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCheckEcsWarningsWithOptions(request *DescribeCheckEcsWarningsRequest, runtime *util.RuntimeOptions) (_result *DescribeCheckEcsWarningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCheckEcsWarningsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCheckEcsWarnings"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCheckEcsWarnings(request *DescribeCheckEcsWarningsRequest) (_result *DescribeCheckEcsWarningsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCheckEcsWarningsResponse{}
	_body, _err := client.DescribeCheckEcsWarningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCheckWarningDetailWithOptions(request *DescribeCheckWarningDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCheckWarningDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCheckWarningDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCheckWarningDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCheckWarningDetail(request *DescribeCheckWarningDetailRequest) (_result *DescribeCheckWarningDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCheckWarningDetailResponse{}
	_body, _err := client.DescribeCheckWarningDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCheckWarningsWithOptions(request *DescribeCheckWarningsRequest, runtime *util.RuntimeOptions) (_result *DescribeCheckWarningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCheckWarningsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCheckWarnings"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCheckWarnings(request *DescribeCheckWarningsRequest) (_result *DescribeCheckWarningsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCheckWarningsResponse{}
	_body, _err := client.DescribeCheckWarningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCheckWarningSummaryWithOptions(request *DescribeCheckWarningSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeCheckWarningSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCheckWarningSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCheckWarningSummary"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCheckWarningSummary(request *DescribeCheckWarningSummaryRequest) (_result *DescribeCheckWarningSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCheckWarningSummaryResponse{}
	_body, _err := client.DescribeCheckWarningSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudCenterInstancesWithOptions(request *DescribeCloudCenterInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudCenterInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCloudCenterInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCloudCenterInstances"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudCenterInstances(request *DescribeCloudCenterInstancesRequest) (_result *DescribeCloudCenterInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudCenterInstancesResponse{}
	_body, _err := client.DescribeCloudCenterInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCloudProductFieldStatisticsWithOptions(request *DescribeCloudProductFieldStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudProductFieldStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCloudProductFieldStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCloudProductFieldStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCloudProductFieldStatistics(request *DescribeCloudProductFieldStatisticsRequest) (_result *DescribeCloudProductFieldStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudProductFieldStatisticsResponse{}
	_body, _err := client.DescribeCloudProductFieldStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConcernNecessityWithOptions(request *DescribeConcernNecessityRequest, runtime *util.RuntimeOptions) (_result *DescribeConcernNecessityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConcernNecessityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConcernNecessity"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConcernNecessity(request *DescribeConcernNecessityRequest) (_result *DescribeConcernNecessityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConcernNecessityResponse{}
	_body, _err := client.DescribeConcernNecessityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerStatisticsWithOptions(request *DescribeContainerStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeContainerStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContainerStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerStatistics(request *DescribeContainerStatisticsRequest) (_result *DescribeContainerStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerStatisticsResponse{}
	_body, _err := client.DescribeContainerStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCriteriaWithOptions(request *DescribeCriteriaRequest, runtime *util.RuntimeOptions) (_result *DescribeCriteriaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCriteriaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCriteria"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCriteria(request *DescribeCriteriaRequest) (_result *DescribeCriteriaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCriteriaResponse{}
	_body, _err := client.DescribeCriteriaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDialogMessagesWithOptions(request *DescribeDialogMessagesRequest, runtime *util.RuntimeOptions) (_result *DescribeDialogMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDialogMessagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDialogMessages"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDialogMessages(request *DescribeDialogMessagesRequest) (_result *DescribeDialogMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDialogMessagesResponse{}
	_body, _err := client.DescribeDialogMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDingTalkWithOptions(request *DescribeDingTalkRequest, runtime *util.RuntimeOptions) (_result *DescribeDingTalkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDingTalkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDingTalk"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDingTalk(request *DescribeDingTalkRequest) (_result *DescribeDingTalkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDingTalkResponse{}
	_body, _err := client.DescribeDingTalkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainCountWithOptions(request *DescribeDomainCountRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainCount"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainCount(request *DescribeDomainCountRequest) (_result *DescribeDomainCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainCountResponse{}
	_body, _err := client.DescribeDomainCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainDetailWithOptions(request *DescribeDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainDetail(request *DescribeDomainDetailRequest) (_result *DescribeDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.DescribeDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainListWithOptions(request *DescribeDomainListRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDomainListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDomainList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainList(request *DescribeDomainListRequest) (_result *DescribeDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainListResponse{}
	_body, _err := client.DescribeDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEmgVulGroupWithOptions(request *DescribeEmgVulGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeEmgVulGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEmgVulGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEmgVulGroup"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEmgVulGroup(request *DescribeEmgVulGroupRequest) (_result *DescribeEmgVulGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEmgVulGroupResponse{}
	_body, _err := client.DescribeEmgVulGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExportInfoWithOptions(request *DescribeExportInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeExportInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExportInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExportInfo"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExportInfo(request *DescribeExportInfoRequest) (_result *DescribeExportInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExportInfoResponse{}
	_body, _err := client.DescribeExportInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExposedInstanceCriteriaWithOptions(request *DescribeExposedInstanceCriteriaRequest, runtime *util.RuntimeOptions) (_result *DescribeExposedInstanceCriteriaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExposedInstanceCriteriaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExposedInstanceCriteria"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExposedInstanceCriteria(request *DescribeExposedInstanceCriteriaRequest) (_result *DescribeExposedInstanceCriteriaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExposedInstanceCriteriaResponse{}
	_body, _err := client.DescribeExposedInstanceCriteriaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExposedInstanceDetailWithOptions(request *DescribeExposedInstanceDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeExposedInstanceDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExposedInstanceDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExposedInstanceDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExposedInstanceDetail(request *DescribeExposedInstanceDetailRequest) (_result *DescribeExposedInstanceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExposedInstanceDetailResponse{}
	_body, _err := client.DescribeExposedInstanceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExposedInstanceListWithOptions(request *DescribeExposedInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeExposedInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExposedInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExposedInstanceList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExposedInstanceList(request *DescribeExposedInstanceListRequest) (_result *DescribeExposedInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExposedInstanceListResponse{}
	_body, _err := client.DescribeExposedInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExposedStatisticsWithOptions(request *DescribeExposedStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeExposedStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExposedStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExposedStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExposedStatistics(request *DescribeExposedStatisticsRequest) (_result *DescribeExposedStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExposedStatisticsResponse{}
	_body, _err := client.DescribeExposedStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFieldStatisticsWithOptions(request *DescribeFieldStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeFieldStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFieldStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFieldStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFieldStatistics(request *DescribeFieldStatisticsRequest) (_result *DescribeFieldStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFieldStatisticsResponse{}
	_body, _err := client.DescribeFieldStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGraph4InvestigationOnlineWithOptions(request *DescribeGraph4InvestigationOnlineRequest, runtime *util.RuntimeOptions) (_result *DescribeGraph4InvestigationOnlineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGraph4InvestigationOnlineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGraph4InvestigationOnline"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGraph4InvestigationOnline(request *DescribeGraph4InvestigationOnlineRequest) (_result *DescribeGraph4InvestigationOnlineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGraph4InvestigationOnlineResponse{}
	_body, _err := client.DescribeGraph4InvestigationOnlineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupedContainerInstancesWithOptions(request *DescribeGroupedContainerInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupedContainerInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupedContainerInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupedContainerInstances"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupedContainerInstances(request *DescribeGroupedContainerInstancesRequest) (_result *DescribeGroupedContainerInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupedContainerInstancesResponse{}
	_body, _err := client.DescribeGroupedContainerInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupedMaliciousFilesWithOptions(request *DescribeGroupedMaliciousFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupedMaliciousFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupedMaliciousFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupedMaliciousFiles"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupedMaliciousFiles(request *DescribeGroupedMaliciousFilesRequest) (_result *DescribeGroupedMaliciousFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupedMaliciousFilesResponse{}
	_body, _err := client.DescribeGroupedMaliciousFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupedTagsWithOptions(request *DescribeGroupedTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupedTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupedTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupedTags"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupedTags(request *DescribeGroupedTagsRequest) (_result *DescribeGroupedTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupedTagsResponse{}
	_body, _err := client.DescribeGroupedTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGroupedVulWithOptions(request *DescribeGroupedVulRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupedVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGroupedVul"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGroupedVul(request *DescribeGroupedVulRequest) (_result *DescribeGroupedVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupedVulResponse{}
	_body, _err := client.DescribeGroupedVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHoneyPotAuthWithOptions(request *DescribeHoneyPotAuthRequest, runtime *util.RuntimeOptions) (_result *DescribeHoneyPotAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHoneyPotAuthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHoneyPotAuth"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHoneyPotAuth(request *DescribeHoneyPotAuthRequest) (_result *DescribeHoneyPotAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHoneyPotAuthResponse{}
	_body, _err := client.DescribeHoneyPotAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHoneyPotSuspStatisticsWithOptions(request *DescribeHoneyPotSuspStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeHoneyPotSuspStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeHoneyPotSuspStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeHoneyPotSuspStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHoneyPotSuspStatistics(request *DescribeHoneyPotSuspStatisticsRequest) (_result *DescribeHoneyPotSuspStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHoneyPotSuspStatisticsResponse{}
	_body, _err := client.DescribeHoneyPotSuspStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageGroupedVulListWithOptions(request *DescribeImageGroupedVulListRequest, runtime *util.RuntimeOptions) (_result *DescribeImageGroupedVulListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImageGroupedVulListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImageGroupedVulList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageGroupedVulList(request *DescribeImageGroupedVulListRequest) (_result *DescribeImageGroupedVulListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageGroupedVulListResponse{}
	_body, _err := client.DescribeImageGroupedVulListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageStatisticsWithOptions(request *DescribeImageStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeImageStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImageStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImageStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageStatistics(request *DescribeImageStatisticsRequest) (_result *DescribeImageStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageStatisticsResponse{}
	_body, _err := client.DescribeImageStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageVulListWithOptions(request *DescribeImageVulListRequest, runtime *util.RuntimeOptions) (_result *DescribeImageVulListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeImageVulListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImageVulList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageVulList(request *DescribeImageVulListRequest) (_result *DescribeImageVulListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageVulListResponse{}
	_body, _err := client.DescribeImageVulListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstallCaptchaWithOptions(request *DescribeInstallCaptchaRequest, runtime *util.RuntimeOptions) (_result *DescribeInstallCaptchaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstallCaptchaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstallCaptcha"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstallCaptcha(request *DescribeInstallCaptchaRequest) (_result *DescribeInstallCaptchaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstallCaptchaResponse{}
	_body, _err := client.DescribeInstallCaptchaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAntiBruteForceRulesWithOptions(request *DescribeInstanceAntiBruteForceRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAntiBruteForceRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAntiBruteForceRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAntiBruteForceRules"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAntiBruteForceRules(request *DescribeInstanceAntiBruteForceRulesRequest) (_result *DescribeInstanceAntiBruteForceRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAntiBruteForceRulesResponse{}
	_body, _err := client.DescribeInstanceAntiBruteForceRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceStatisticsWithOptions(request *DescribeInstanceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStatistics"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceStatistics(request *DescribeInstanceStatisticsRequest) (_result *DescribeInstanceStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DescribeInstanceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpInfoWithOptions(request *DescribeIpInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeIpInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIpInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIpInfo"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpInfo(request *DescribeIpInfoRequest) (_result *DescribeIpInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpInfoResponse{}
	_body, _err := client.DescribeIpInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModuleConfigWithOptions(request *DescribeModuleConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeModuleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeModuleConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeModuleConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModuleConfig(request *DescribeModuleConfigRequest) (_result *DescribeModuleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModuleConfigResponse{}
	_body, _err := client.DescribeModuleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNoticeConfigWithOptions(request *DescribeNoticeConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeNoticeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeNoticeConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNoticeConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNoticeConfig(request *DescribeNoticeConfigRequest) (_result *DescribeNoticeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNoticeConfigResponse{}
	_body, _err := client.DescribeNoticeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyCountWithOptions(request *DescribePropertyCountRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyCount"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyCount(request *DescribePropertyCountRequest) (_result *DescribePropertyCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyCountResponse{}
	_body, _err := client.DescribePropertyCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyCronDetailWithOptions(request *DescribePropertyCronDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyCronDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyCronDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyCronDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyCronDetail(request *DescribePropertyCronDetailRequest) (_result *DescribePropertyCronDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyCronDetailResponse{}
	_body, _err := client.DescribePropertyCronDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyPortDetailWithOptions(request *DescribePropertyPortDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyPortDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyPortDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyPortDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyPortDetail(request *DescribePropertyPortDetailRequest) (_result *DescribePropertyPortDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyPortDetailResponse{}
	_body, _err := client.DescribePropertyPortDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyPortItemWithOptions(request *DescribePropertyPortItemRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyPortItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyPortItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyPortItem"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyPortItem(request *DescribePropertyPortItemRequest) (_result *DescribePropertyPortItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyPortItemResponse{}
	_body, _err := client.DescribePropertyPortItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyProcDetailWithOptions(request *DescribePropertyProcDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyProcDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyProcDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyProcDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyProcDetail(request *DescribePropertyProcDetailRequest) (_result *DescribePropertyProcDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyProcDetailResponse{}
	_body, _err := client.DescribePropertyProcDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyProcItemWithOptions(request *DescribePropertyProcItemRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyProcItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyProcItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyProcItem"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyProcItem(request *DescribePropertyProcItemRequest) (_result *DescribePropertyProcItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyProcItemResponse{}
	_body, _err := client.DescribePropertyProcItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyScaDetailWithOptions(request *DescribePropertyScaDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyScaDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyScaDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyScaDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyScaDetail(request *DescribePropertyScaDetailRequest) (_result *DescribePropertyScaDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyScaDetailResponse{}
	_body, _err := client.DescribePropertyScaDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertySoftwareDetailWithOptions(request *DescribePropertySoftwareDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePropertySoftwareDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertySoftwareDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertySoftwareDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertySoftwareDetail(request *DescribePropertySoftwareDetailRequest) (_result *DescribePropertySoftwareDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertySoftwareDetailResponse{}
	_body, _err := client.DescribePropertySoftwareDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertySoftwareItemWithOptions(request *DescribePropertySoftwareItemRequest, runtime *util.RuntimeOptions) (_result *DescribePropertySoftwareItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertySoftwareItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertySoftwareItem"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertySoftwareItem(request *DescribePropertySoftwareItemRequest) (_result *DescribePropertySoftwareItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertySoftwareItemResponse{}
	_body, _err := client.DescribePropertySoftwareItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyUsageNewestWithOptions(request *DescribePropertyUsageNewestRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyUsageNewestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyUsageNewestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyUsageNewest"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyUsageNewest(request *DescribePropertyUsageNewestRequest) (_result *DescribePropertyUsageNewestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyUsageNewestResponse{}
	_body, _err := client.DescribePropertyUsageNewestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyUserDetailWithOptions(request *DescribePropertyUserDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyUserDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyUserDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyUserDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyUserDetail(request *DescribePropertyUserDetailRequest) (_result *DescribePropertyUserDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyUserDetailResponse{}
	_body, _err := client.DescribePropertyUserDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePropertyUserItemWithOptions(request *DescribePropertyUserItemRequest, runtime *util.RuntimeOptions) (_result *DescribePropertyUserItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePropertyUserItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePropertyUserItem"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePropertyUserItem(request *DescribePropertyUserItemRequest) (_result *DescribePropertyUserItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePropertyUserItemResponse{}
	_body, _err := client.DescribePropertyUserItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRiskCheckItemResultWithOptions(request *DescribeRiskCheckItemResultRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskCheckItemResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRiskCheckItemResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRiskCheckItemResult"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRiskCheckItemResult(request *DescribeRiskCheckItemResultRequest) (_result *DescribeRiskCheckItemResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskCheckItemResultResponse{}
	_body, _err := client.DescribeRiskCheckItemResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRiskCheckResultWithOptions(request *DescribeRiskCheckResultRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRiskCheckResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRiskCheckResult"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRiskCheckResult(request *DescribeRiskCheckResultRequest) (_result *DescribeRiskCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskCheckResultResponse{}
	_body, _err := client.DescribeRiskCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRiskCheckSummaryWithOptions(request *DescribeRiskCheckSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskCheckSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRiskCheckSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRiskCheckSummary"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRiskCheckSummary(request *DescribeRiskCheckSummaryRequest) (_result *DescribeRiskCheckSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskCheckSummaryResponse{}
	_body, _err := client.DescribeRiskCheckSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRiskItemTypeWithOptions(request *DescribeRiskItemTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskItemTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRiskItemTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRiskItemType"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRiskItemType(request *DescribeRiskItemTypeRequest) (_result *DescribeRiskItemTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskItemTypeResponse{}
	_body, _err := client.DescribeRiskItemTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRiskListCheckResultWithOptions(request *DescribeRiskListCheckResultRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskListCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRiskListCheckResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRiskListCheckResult"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRiskListCheckResult(request *DescribeRiskListCheckResultRequest) (_result *DescribeRiskListCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskListCheckResultResponse{}
	_body, _err := client.DescribeRiskListCheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSasAssetStatisticsColumnWithOptions(request *DescribeSasAssetStatisticsColumnRequest, runtime *util.RuntimeOptions) (_result *DescribeSasAssetStatisticsColumnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSasAssetStatisticsColumnResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSasAssetStatisticsColumn"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSasAssetStatisticsColumn(request *DescribeSasAssetStatisticsColumnRequest) (_result *DescribeSasAssetStatisticsColumnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSasAssetStatisticsColumnResponse{}
	_body, _err := client.DescribeSasAssetStatisticsColumnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSearchConditionWithOptions(request *DescribeSearchConditionRequest, runtime *util.RuntimeOptions) (_result *DescribeSearchConditionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSearchConditionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSearchCondition"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSearchCondition(request *DescribeSearchConditionRequest) (_result *DescribeSearchConditionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSearchConditionResponse{}
	_body, _err := client.DescribeSearchConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecureSuggestionWithOptions(request *DescribeSecureSuggestionRequest, runtime *util.RuntimeOptions) (_result *DescribeSecureSuggestionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecureSuggestionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecureSuggestion"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecureSuggestion(request *DescribeSecureSuggestionRequest) (_result *DescribeSecureSuggestionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecureSuggestionResponse{}
	_body, _err := client.DescribeSecureSuggestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityCheckScheduleConfigWithOptions(request *DescribeSecurityCheckScheduleConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityCheckScheduleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityCheckScheduleConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityCheckScheduleConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityCheckScheduleConfig(request *DescribeSecurityCheckScheduleConfigRequest) (_result *DescribeSecurityCheckScheduleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityCheckScheduleConfigResponse{}
	_body, _err := client.DescribeSecurityCheckScheduleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationsWithOptions(request *DescribeSecurityEventOperationsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityEventOperations"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperations(request *DescribeSecurityEventOperationsRequest) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DescribeSecurityEventOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatusWithOptions(request *DescribeSecurityEventOperationStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityEventOperationStatus"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityEventOperationStatus(request *DescribeSecurityEventOperationStatusRequest) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.DescribeSecurityEventOperationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityStatInfoWithOptions(request *DescribeSecurityStatInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityStatInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityStatInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityStatInfo"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityStatInfo(request *DescribeSecurityStatInfoRequest) (_result *DescribeSecurityStatInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityStatInfoResponse{}
	_body, _err := client.DescribeSecurityStatInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSimilarEventScenariosWithOptions(request *DescribeSimilarEventScenariosRequest, runtime *util.RuntimeOptions) (_result *DescribeSimilarEventScenariosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSimilarEventScenariosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSimilarEventScenarios"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSimilarEventScenarios(request *DescribeSimilarEventScenariosRequest) (_result *DescribeSimilarEventScenariosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSimilarEventScenariosResponse{}
	_body, _err := client.DescribeSimilarEventScenariosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSimilarSecurityEventsWithOptions(request *DescribeSimilarSecurityEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeSimilarSecurityEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSimilarSecurityEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSimilarSecurityEvents"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSimilarSecurityEvents(request *DescribeSimilarSecurityEventsRequest) (_result *DescribeSimilarSecurityEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSimilarSecurityEventsResponse{}
	_body, _err := client.DescribeSimilarSecurityEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStrategyExecDetailWithOptions(request *DescribeStrategyExecDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeStrategyExecDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStrategyExecDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStrategyExecDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStrategyExecDetail(request *DescribeStrategyExecDetailRequest) (_result *DescribeStrategyExecDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStrategyExecDetailResponse{}
	_body, _err := client.DescribeStrategyExecDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStratetyWithOptions(request *DescribeStratetyRequest, runtime *util.RuntimeOptions) (_result *DescribeStratetyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStratetyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStratety"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStratety(request *DescribeStratetyRequest) (_result *DescribeStratetyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeStratetyResponse{}
	_body, _err := client.DescribeStratetyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSummaryInfoWithOptions(request *DescribeSummaryInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeSummaryInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSummaryInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSummaryInfo"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSummaryInfo(request *DescribeSummaryInfoRequest) (_result *DescribeSummaryInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSummaryInfoResponse{}
	_body, _err := client.DescribeSummaryInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventDetailWithOptions(request *DescribeSuspEventDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSuspEventDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSuspEventDetail"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventDetail(request *DescribeSuspEventDetailRequest) (_result *DescribeSuspEventDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventDetailResponse{}
	_body, _err := client.DescribeSuspEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFilesWithOptions(request *DescribeSuspEventQuaraFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSuspEventQuaraFiles"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEventQuaraFiles(request *DescribeSuspEventQuaraFilesRequest) (_result *DescribeSuspEventQuaraFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventQuaraFilesResponse{}
	_body, _err := client.DescribeSuspEventQuaraFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSuspEventsWithOptions(request *DescribeSuspEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSuspEvents"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSuspEvents(request *DescribeSuspEventsRequest) (_result *DescribeSuspEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DescribeSuspEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserBaselineAuthorizationWithOptions(request *DescribeUserBaselineAuthorizationRequest, runtime *util.RuntimeOptions) (_result *DescribeUserBaselineAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserBaselineAuthorizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserBaselineAuthorization"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserBaselineAuthorization(request *DescribeUserBaselineAuthorizationRequest) (_result *DescribeUserBaselineAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserBaselineAuthorizationResponse{}
	_body, _err := client.DescribeUserBaselineAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserLayoutAuthorizationWithOptions(request *DescribeUserLayoutAuthorizationRequest, runtime *util.RuntimeOptions) (_result *DescribeUserLayoutAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserLayoutAuthorizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserLayoutAuthorization"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserLayoutAuthorization(request *DescribeUserLayoutAuthorizationRequest) (_result *DescribeUserLayoutAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserLayoutAuthorizationResponse{}
	_body, _err := client.DescribeUserLayoutAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVersionConfigWithOptions(request *DescribeVersionConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeVersionConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVersionConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVersionConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVersionConfig(request *DescribeVersionConfigRequest) (_result *DescribeVersionConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVersionConfigResponse{}
	_body, _err := client.DescribeVersionConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVolDingdingMessageWithOptions(request *DescribeVolDingdingMessageRequest, runtime *util.RuntimeOptions) (_result *DescribeVolDingdingMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVolDingdingMessageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVolDingdingMessage"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVolDingdingMessage(request *DescribeVolDingdingMessageRequest) (_result *DescribeVolDingdingMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVolDingdingMessageResponse{}
	_body, _err := client.DescribeVolDingdingMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcHoneyPotCriteriaWithOptions(request *DescribeVpcHoneyPotCriteriaRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcHoneyPotCriteriaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVpcHoneyPotCriteriaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVpcHoneyPotCriteria"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcHoneyPotCriteria(request *DescribeVpcHoneyPotCriteriaRequest) (_result *DescribeVpcHoneyPotCriteriaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcHoneyPotCriteriaResponse{}
	_body, _err := client.DescribeVpcHoneyPotCriteriaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcHoneyPotListWithOptions(request *DescribeVpcHoneyPotListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcHoneyPotListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVpcHoneyPotListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVpcHoneyPotList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcHoneyPotList(request *DescribeVpcHoneyPotListRequest) (_result *DescribeVpcHoneyPotListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcHoneyPotListResponse{}
	_body, _err := client.DescribeVpcHoneyPotListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcListWithOptions(request *DescribeVpcListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVpcListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVpcList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcList(request *DescribeVpcListRequest) (_result *DescribeVpcListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcListResponse{}
	_body, _err := client.DescribeVpcListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulDetailsWithOptions(request *DescribeVulDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeVulDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulDetails"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulDetails(request *DescribeVulDetailsRequest) (_result *DescribeVulDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulDetailsResponse{}
	_body, _err := client.DescribeVulDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulListWithOptions(request *DescribeVulListRequest, runtime *util.RuntimeOptions) (_result *DescribeVulListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulList(request *DescribeVulListRequest) (_result *DescribeVulListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulListResponse{}
	_body, _err := client.DescribeVulListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulWhitelistWithOptions(request *DescribeVulWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeVulWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulWhitelist"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulWhitelist(request *DescribeVulWhitelistRequest) (_result *DescribeVulWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulWhitelistResponse{}
	_body, _err := client.DescribeVulWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWarningMachinesWithOptions(request *DescribeWarningMachinesRequest, runtime *util.RuntimeOptions) (_result *DescribeWarningMachinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWarningMachinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWarningMachines"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWarningMachines(request *DescribeWarningMachinesRequest) (_result *DescribeWarningMachinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWarningMachinesResponse{}
	_body, _err := client.DescribeWarningMachinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebLockBindListWithOptions(request *DescribeWebLockBindListRequest, runtime *util.RuntimeOptions) (_result *DescribeWebLockBindListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebLockBindListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebLockBindList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebLockBindList(request *DescribeWebLockBindListRequest) (_result *DescribeWebLockBindListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebLockBindListResponse{}
	_body, _err := client.DescribeWebLockBindListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebLockConfigListWithOptions(request *DescribeWebLockConfigListRequest, runtime *util.RuntimeOptions) (_result *DescribeWebLockConfigListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeWebLockConfigListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeWebLockConfigList"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebLockConfigList(request *DescribeWebLockConfigListRequest) (_result *DescribeWebLockConfigListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebLockConfigListResponse{}
	_body, _err := client.DescribeWebLockConfigListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportRecordWithOptions(request *ExportRecordRequest, runtime *util.RuntimeOptions) (_result *ExportRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExportRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExportRecord"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportRecord(request *ExportRecordRequest) (_result *ExportRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportRecordResponse{}
	_body, _err := client.ExportRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FixCheckWarningsWithOptions(request *FixCheckWarningsRequest, runtime *util.RuntimeOptions) (_result *FixCheckWarningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FixCheckWarningsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FixCheckWarnings"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FixCheckWarnings(request *FixCheckWarningsRequest) (_result *FixCheckWarningsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FixCheckWarningsResponse{}
	_body, _err := client.FixCheckWarningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIncIOCsWithOptions(request *GetIncIOCsRequest, runtime *util.RuntimeOptions) (_result *GetIncIOCsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIncIOCsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIncIOCs"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIncIOCs(request *GetIncIOCsRequest) (_result *GetIncIOCsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIncIOCsResponse{}
	_body, _err := client.GetIncIOCsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIOCsWithOptions(request *GetIOCsRequest, runtime *util.RuntimeOptions) (_result *GetIOCsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetIOCsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIOCs"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIOCs(request *GetIOCsRequest) (_result *GetIOCsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIOCsResponse{}
	_body, _err := client.GetIOCsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandleSecurityEventsWithOptions(request *HandleSecurityEventsRequest, runtime *util.RuntimeOptions) (_result *HandleSecurityEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HandleSecurityEvents"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandleSecurityEvents(request *HandleSecurityEventsRequest) (_result *HandleSecurityEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.HandleSecurityEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HandleSimilarSecurityEventsWithOptions(request *HandleSimilarSecurityEventsRequest, runtime *util.RuntimeOptions) (_result *HandleSimilarSecurityEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HandleSimilarSecurityEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HandleSimilarSecurityEvents"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HandleSimilarSecurityEvents(request *HandleSimilarSecurityEventsRequest) (_result *HandleSimilarSecurityEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HandleSimilarSecurityEventsResponse{}
	_body, _err := client.HandleSimilarSecurityEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IgnoreHcCheckWarningsWithOptions(request *IgnoreHcCheckWarningsRequest, runtime *util.RuntimeOptions) (_result *IgnoreHcCheckWarningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IgnoreHcCheckWarningsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IgnoreHcCheckWarnings"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IgnoreHcCheckWarnings(request *IgnoreHcCheckWarningsRequest) (_result *IgnoreHcCheckWarningsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IgnoreHcCheckWarningsResponse{}
	_body, _err := client.IgnoreHcCheckWarningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAntiBruteForceRuleWithOptions(request *ModifyAntiBruteForceRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyAntiBruteForceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAntiBruteForceRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAntiBruteForceRule"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAntiBruteForceRule(request *ModifyAntiBruteForceRuleRequest) (_result *ModifyAntiBruteForceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAntiBruteForceRuleResponse{}
	_body, _err := client.ModifyAntiBruteForceRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCreateVulWhitelistWithOptions(request *ModifyCreateVulWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyCreateVulWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCreateVulWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCreateVulWhitelist"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCreateVulWhitelist(request *ModifyCreateVulWhitelistRequest) (_result *ModifyCreateVulWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCreateVulWhitelistResponse{}
	_body, _err := client.ModifyCreateVulWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyEmgVulSubmitWithOptions(request *ModifyEmgVulSubmitRequest, runtime *util.RuntimeOptions) (_result *ModifyEmgVulSubmitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyEmgVulSubmitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyEmgVulSubmit"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyEmgVulSubmit(request *ModifyEmgVulSubmitRequest) (_result *ModifyEmgVulSubmitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyEmgVulSubmitResponse{}
	_body, _err := client.ModifyEmgVulSubmitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGroupPropertyWithOptions(request *ModifyGroupPropertyRequest, runtime *util.RuntimeOptions) (_result *ModifyGroupPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyGroupPropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyGroupProperty"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGroupProperty(request *ModifyGroupPropertyRequest) (_result *ModifyGroupPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGroupPropertyResponse{}
	_body, _err := client.ModifyGroupPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAntiBruteForceRuleWithOptions(request *ModifyInstanceAntiBruteForceRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAntiBruteForceRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAntiBruteForceRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAntiBruteForceRule"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAntiBruteForceRule(request *ModifyInstanceAntiBruteForceRuleRequest) (_result *ModifyInstanceAntiBruteForceRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAntiBruteForceRuleResponse{}
	_body, _err := client.ModifyInstanceAntiBruteForceRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLoginBaseConfigWithOptions(request *ModifyLoginBaseConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyLoginBaseConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLoginBaseConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLoginBaseConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLoginBaseConfig(request *ModifyLoginBaseConfigRequest) (_result *ModifyLoginBaseConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLoginBaseConfigResponse{}
	_body, _err := client.ModifyLoginBaseConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLoginSwitchConfigWithOptions(request *ModifyLoginSwitchConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyLoginSwitchConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLoginSwitchConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLoginSwitchConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLoginSwitchConfig(request *ModifyLoginSwitchConfigRequest) (_result *ModifyLoginSwitchConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLoginSwitchConfigResponse{}
	_body, _err := client.ModifyLoginSwitchConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNoticeConfigWithOptions(request *ModifyNoticeConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyNoticeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyNoticeConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyNoticeConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNoticeConfig(request *ModifyNoticeConfigRequest) (_result *ModifyNoticeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNoticeConfigResponse{}
	_body, _err := client.ModifyNoticeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOperateVulWithOptions(request *ModifyOperateVulRequest, runtime *util.RuntimeOptions) (_result *ModifyOperateVulResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyOperateVul"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOperateVul(request *ModifyOperateVulRequest) (_result *ModifyOperateVulResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOperateVulResponse{}
	_body, _err := client.ModifyOperateVulWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPushAllTaskWithOptions(request *ModifyPushAllTaskRequest, runtime *util.RuntimeOptions) (_result *ModifyPushAllTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPushAllTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPushAllTask"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPushAllTask(request *ModifyPushAllTaskRequest) (_result *ModifyPushAllTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPushAllTaskResponse{}
	_body, _err := client.ModifyPushAllTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRiskCheckStatusWithOptions(request *ModifyRiskCheckStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyRiskCheckStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRiskCheckStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRiskCheckStatus"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRiskCheckStatus(request *ModifyRiskCheckStatusRequest) (_result *ModifyRiskCheckStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRiskCheckStatusResponse{}
	_body, _err := client.ModifyRiskCheckStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRiskSingleResultStatusWithOptions(request *ModifyRiskSingleResultStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyRiskSingleResultStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRiskSingleResultStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRiskSingleResultStatus"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRiskSingleResultStatus(request *ModifyRiskSingleResultStatusRequest) (_result *ModifyRiskSingleResultStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRiskSingleResultStatusResponse{}
	_body, _err := client.ModifyRiskSingleResultStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityCheckScheduleConfigWithOptions(request *ModifySecurityCheckScheduleConfigRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityCheckScheduleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySecurityCheckScheduleConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySecurityCheckScheduleConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityCheckScheduleConfig(request *ModifySecurityCheckScheduleConfigRequest) (_result *ModifySecurityCheckScheduleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityCheckScheduleConfigResponse{}
	_body, _err := client.ModifySecurityCheckScheduleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyStartVulScanWithOptions(request *ModifyStartVulScanRequest, runtime *util.RuntimeOptions) (_result *ModifyStartVulScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyStartVulScanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyStartVulScan"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyStartVulScan(request *ModifyStartVulScanRequest) (_result *ModifyStartVulScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyStartVulScanResponse{}
	_body, _err := client.ModifyStartVulScanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTagWithUuidWithOptions(request *ModifyTagWithUuidRequest, runtime *util.RuntimeOptions) (_result *ModifyTagWithUuidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyTagWithUuidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyTagWithUuid"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTagWithUuid(request *ModifyTagWithUuidRequest) (_result *ModifyTagWithUuidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTagWithUuidResponse{}
	_body, _err := client.ModifyTagWithUuidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVpcHoneyPotWithOptions(request *ModifyVpcHoneyPotRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcHoneyPotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyVpcHoneyPotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyVpcHoneyPot"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVpcHoneyPot(request *ModifyVpcHoneyPotRequest) (_result *ModifyVpcHoneyPotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcHoneyPotResponse{}
	_body, _err := client.ModifyVpcHoneyPotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVulTargetConfigWithOptions(request *ModifyVulTargetConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyVulTargetConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyVulTargetConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyVulTargetConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVulTargetConfig(request *ModifyVulTargetConfigRequest) (_result *ModifyVulTargetConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVulTargetConfigResponse{}
	_body, _err := client.ModifyVulTargetConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebLockCreateConfigWithOptions(request *ModifyWebLockCreateConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyWebLockCreateConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebLockCreateConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebLockCreateConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebLockCreateConfig(request *ModifyWebLockCreateConfigRequest) (_result *ModifyWebLockCreateConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebLockCreateConfigResponse{}
	_body, _err := client.ModifyWebLockCreateConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebLockDeleteConfigWithOptions(request *ModifyWebLockDeleteConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyWebLockDeleteConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebLockDeleteConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebLockDeleteConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebLockDeleteConfig(request *ModifyWebLockDeleteConfigRequest) (_result *ModifyWebLockDeleteConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebLockDeleteConfigResponse{}
	_body, _err := client.ModifyWebLockDeleteConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebLockStartWithOptions(request *ModifyWebLockStartRequest, runtime *util.RuntimeOptions) (_result *ModifyWebLockStartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebLockStartResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebLockStart"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebLockStart(request *ModifyWebLockStartRequest) (_result *ModifyWebLockStartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebLockStartResponse{}
	_body, _err := client.ModifyWebLockStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebLockUnbindWithOptions(request *ModifyWebLockUnbindRequest, runtime *util.RuntimeOptions) (_result *ModifyWebLockUnbindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebLockUnbindResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebLockUnbind"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebLockUnbind(request *ModifyWebLockUnbindRequest) (_result *ModifyWebLockUnbindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebLockUnbindResponse{}
	_body, _err := client.ModifyWebLockUnbindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyWebLockUpdateConfigWithOptions(request *ModifyWebLockUpdateConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyWebLockUpdateConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyWebLockUpdateConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyWebLockUpdateConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyWebLockUpdateConfig(request *ModifyWebLockUpdateConfigRequest) (_result *ModifyWebLockUpdateConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyWebLockUpdateConfigResponse{}
	_body, _err := client.ModifyWebLockUpdateConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperateSuspiciousTargetConfigWithOptions(request *OperateSuspiciousTargetConfigRequest, runtime *util.RuntimeOptions) (_result *OperateSuspiciousTargetConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OperateSuspiciousTargetConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OperateSuspiciousTargetConfig"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperateSuspiciousTargetConfig(request *OperateSuspiciousTargetConfigRequest) (_result *OperateSuspiciousTargetConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperateSuspiciousTargetConfigResponse{}
	_body, _err := client.OperateSuspiciousTargetConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OperationSuspEventsWithOptions(request *OperationSuspEventsRequest, runtime *util.RuntimeOptions) (_result *OperationSuspEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &OperationSuspEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("OperationSuspEvents"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OperationSuspEvents(request *OperationSuspEventsRequest) (_result *OperationSuspEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OperationSuspEventsResponse{}
	_body, _err := client.OperationSuspEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PauseClientWithOptions(request *PauseClientRequest, runtime *util.RuntimeOptions) (_result *PauseClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PauseClientResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PauseClient"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseClient(request *PauseClientRequest) (_result *PauseClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseClientResponse{}
	_body, _err := client.PauseClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshContainerAssetsWithOptions(request *RefreshContainerAssetsRequest, runtime *util.RuntimeOptions) (_result *RefreshContainerAssetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshContainerAssetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshContainerAssets"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshContainerAssets(request *RefreshContainerAssetsRequest) (_result *RefreshContainerAssetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshContainerAssetsResponse{}
	_body, _err := client.RefreshContainerAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFileWithOptions(request *RollbackSuspEventQuaraFileRequest, runtime *util.RuntimeOptions) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RollbackSuspEventQuaraFile"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackSuspEventQuaraFile(request *RollbackSuspEventQuaraFileRequest) (_result *RollbackSuspEventQuaraFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackSuspEventQuaraFileResponse{}
	_body, _err := client.RollbackSuspEventQuaraFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SasInstallCodeWithOptions(request *SasInstallCodeRequest, runtime *util.RuntimeOptions) (_result *SasInstallCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SasInstallCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SasInstallCode"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SasInstallCode(request *SasInstallCodeRequest) (_result *SasInstallCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SasInstallCodeResponse{}
	_body, _err := client.SasInstallCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartBaselineSecurityCheckWithOptions(request *StartBaselineSecurityCheckRequest, runtime *util.RuntimeOptions) (_result *StartBaselineSecurityCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartBaselineSecurityCheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartBaselineSecurityCheck"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartBaselineSecurityCheck(request *StartBaselineSecurityCheckRequest) (_result *StartBaselineSecurityCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartBaselineSecurityCheckResponse{}
	_body, _err := client.StartBaselineSecurityCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartImageVulScanWithOptions(request *StartImageVulScanRequest, runtime *util.RuntimeOptions) (_result *StartImageVulScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartImageVulScanResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartImageVulScan"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartImageVulScan(request *StartImageVulScanRequest) (_result *StartImageVulScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartImageVulScanResponse{}
	_body, _err := client.StartImageVulScanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateHcWarningsWithOptions(request *ValidateHcWarningsRequest, runtime *util.RuntimeOptions) (_result *ValidateHcWarningsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ValidateHcWarningsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ValidateHcWarnings"), tea.String("2018-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateHcWarnings(request *ValidateHcWarningsRequest) (_result *ValidateHcWarningsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateHcWarningsResponse{}
	_body, _err := client.ValidateHcWarningsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
