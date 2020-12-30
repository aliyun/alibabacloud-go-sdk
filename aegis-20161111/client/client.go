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

type CreateInstanceRequest struct {
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Duration          *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle      *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	VersionCode       *int32  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VmNumber          *int32  `json:"VmNumber,omitempty" xml:"VmNumber,omitempty"`
	IsAutoRenew       *bool   `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	AutoRenewDuration *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetDuration(v int32) *CreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetVersionCode(v int32) *CreateInstanceRequest {
	s.VersionCode = &v
	return s
}

func (s *CreateInstanceRequest) SetVmNumber(v int32) *CreateInstanceRequest {
	s.VmNumber = &v
	return s
}

func (s *CreateInstanceRequest) SetIsAutoRenew(v bool) *CreateInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewDuration(v int32) *CreateInstanceRequest {
	s.AutoRenewDuration = &v
	return s
}

type CreateInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetOrderId(v string) *CreateInstanceResponseBody {
	s.OrderId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetSourceIp(v string) *DeleteRuleRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteRuleRequest) SetLang(v string) *DeleteRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRuleRequest) SetId(v string) *DeleteRuleRequest {
	s.Id = &v
	return s
}

type DeleteRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
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

type DescribeInstanceStatisticsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
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
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Vul        *int32  `json:"Vul,omitempty" xml:"Vul,omitempty"`
	Health     *int32  `json:"Health,omitempty" xml:"Health,omitempty"`
	Trojan     *int32  `json:"Trojan,omitempty" xml:"Trojan,omitempty"`
	Suspicious *int32  `json:"Suspicious,omitempty" xml:"Suspicious,omitempty"`
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

func (s *DescribeInstanceStatisticsResponseBodyData) SetTrojan(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Trojan = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetSuspicious(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Suspicious = &v
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
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
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

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetInternetIp(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.InternetIp = &v
	return s
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

type RenewInstanceRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VmNumber     *string `json:"VmNumber,omitempty" xml:"VmNumber,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetVmNumber(v string) *RenewInstanceRequest {
	s.VmNumber = &v
	return s
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

type RenewInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetOrderId(v string) *RenewInstanceResponseBody {
	s.OrderId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type UpgradeInstanceRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VersionCode *int32  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VmNumber    *int32  `json:"VmNumber,omitempty" xml:"VmNumber,omitempty"`
}

func (s UpgradeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) SetOwnerId(v int64) *UpgradeInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetClientToken(v string) *UpgradeInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceId(v string) *UpgradeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetVersionCode(v int32) *UpgradeInstanceRequest {
	s.VersionCode = &v
	return s
}

func (s *UpgradeInstanceRequest) SetVmNumber(v int32) *UpgradeInstanceRequest {
	s.VmNumber = &v
	return s
}

type UpgradeInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s UpgradeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBody) SetRequestId(v string) *UpgradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetOrderId(v string) *UpgradeInstanceResponseBody {
	s.OrderId = &v
	return s
}

type UpgradeInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponse) SetHeaders(v map[string]*string) *UpgradeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceResponse) SetBody(v *UpgradeInstanceResponseBody) *UpgradeInstanceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aegis"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRule"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoDelConfig"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeCheckWarningDetailWithOptions(request *DescribeCheckWarningDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeCheckWarningDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCheckWarningDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCheckWarningDetail"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeConcernNecessityWithOptions(request *DescribeConcernNecessityRequest, runtime *util.RuntimeOptions) (_result *DescribeConcernNecessityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConcernNecessityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConcernNecessity"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeInstanceStatisticsWithOptions(request *DescribeInstanceStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceStatistics"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeStrategyExecDetailWithOptions(request *DescribeStrategyExecDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeStrategyExecDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeStrategyExecDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeStrategyExecDetail"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeStratety"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeVulWhitelistWithOptions(request *DescribeVulWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeVulWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVulWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVulWhitelist"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewInstance"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeInstanceWithOptions(request *UpgradeInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeInstance"), tea.String("2016-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeInstance(request *UpgradeInstanceRequest) (_result *UpgradeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.UpgradeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
