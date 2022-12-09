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

type CreateInstanceRequest struct {
	AutoRenewDuration *int32  `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Duration          *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IsAutoRenew       *bool   `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PricingCycle      *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	VersionCode       *int32  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VmNumber          *int32  `json:"VmNumber,omitempty" xml:"VmNumber,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoRenewDuration(v int32) *CreateInstanceRequest {
	s.AutoRenewDuration = &v
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

func (s *CreateInstanceRequest) SetIsAutoRenew(v bool) *CreateInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
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

type CreateInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetOrderId(v string) *CreateInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetId(v string) *DeleteRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteRuleRequest) SetLang(v string) *DeleteRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRuleRequest) SetSourceIp(v string) *DeleteRuleRequest {
	s.SourceIp = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRuleResponse) SetStatusCode(v int32) *DeleteRuleResponse {
	s.StatusCode = &v
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
	Days      *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAutoDelConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoDelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoDelConfigResponseBody) SetDays(v int32) *DescribeAutoDelConfigResponseBody {
	s.Days = &v
	return s
}

func (s *DescribeAutoDelConfigResponseBody) SetRequestId(v string) *DescribeAutoDelConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAutoDelConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAutoDelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAutoDelConfigResponse) SetStatusCode(v int32) *DescribeAutoDelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoDelConfigResponse) SetBody(v *DescribeAutoDelConfigResponseBody) *DescribeAutoDelConfigResponse {
	s.Body = v
	return s
}

type DescribeCheckWarningDetailRequest struct {
	CheckWarningId *int64  `json:"CheckWarningId,omitempty" xml:"CheckWarningId,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeCheckWarningDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailRequest) SetCheckWarningId(v int64) *DescribeCheckWarningDetailRequest {
	s.CheckWarningId = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetLang(v string) *DescribeCheckWarningDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningDetailRequest) SetSourceIp(v string) *DescribeCheckWarningDetailRequest {
	s.SourceIp = &v
	return s
}

type DescribeCheckWarningDetailResponseBody struct {
	Advice      *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	CheckId     *int64  `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Item        *string `json:"Item,omitempty" xml:"Item,omitempty"`
	Level       *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Prompt      *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCheckWarningDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckWarningDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningDetailResponseBody) SetAdvice(v string) *DescribeCheckWarningDetailResponseBody {
	s.Advice = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetCheckId(v int64) *DescribeCheckWarningDetailResponseBody {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetDescription(v string) *DescribeCheckWarningDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetItem(v string) *DescribeCheckWarningDetailResponseBody {
	s.Item = &v
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

func (s *DescribeCheckWarningDetailResponseBody) SetRequestId(v string) *DescribeCheckWarningDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningDetailResponseBody) SetType(v string) *DescribeCheckWarningDetailResponseBody {
	s.Type = &v
	return s
}

type DescribeCheckWarningDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCheckWarningDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeCheckWarningDetailResponse) SetStatusCode(v int32) *DescribeCheckWarningDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckWarningDetailResponse) SetBody(v *DescribeCheckWarningDetailResponseBody) *DescribeCheckWarningDetailResponse {
	s.Body = v
	return s
}

type DescribeConcernNecessityRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeConcernNecessityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConcernNecessityRequest) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityRequest) SetLang(v string) *DescribeConcernNecessityRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConcernNecessityRequest) SetSourceIp(v string) *DescribeConcernNecessityRequest {
	s.SourceIp = &v
	return s
}

type DescribeConcernNecessityResponseBody struct {
	ConcernNecessity []*string `json:"ConcernNecessity,omitempty" xml:"ConcernNecessity,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConcernNecessityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConcernNecessityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConcernNecessityResponseBody) SetConcernNecessity(v []*string) *DescribeConcernNecessityResponseBody {
	s.ConcernNecessity = v
	return s
}

func (s *DescribeConcernNecessityResponseBody) SetRequestId(v string) *DescribeConcernNecessityResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConcernNecessityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeConcernNecessityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeConcernNecessityResponse) SetStatusCode(v int32) *DescribeConcernNecessityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConcernNecessityResponse) SetBody(v *DescribeConcernNecessityResponseBody) *DescribeConcernNecessityResponse {
	s.Body = v
	return s
}

type DescribeInstanceStatisticsRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsRequest) SetLang(v string) *DescribeInstanceStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetSourceIp(v string) *DescribeInstanceStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceStatisticsRequest) SetUuid(v string) *DescribeInstanceStatisticsRequest {
	s.Uuid = &v
	return s
}

type DescribeInstanceStatisticsResponseBody struct {
	Data      []*DescribeInstanceStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatisticsResponseBody) SetData(v []*DescribeInstanceStatisticsResponseBodyData) *DescribeInstanceStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceStatisticsResponseBody) SetRequestId(v string) *DescribeInstanceStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceStatisticsResponseBodyData struct {
	Account    *int32  `json:"Account,omitempty" xml:"Account,omitempty"`
	Health     *int32  `json:"Health,omitempty" xml:"Health,omitempty"`
	Suspicious *int32  `json:"Suspicious,omitempty" xml:"Suspicious,omitempty"`
	Trojan     *int32  `json:"Trojan,omitempty" xml:"Trojan,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Vul        *int32  `json:"Vul,omitempty" xml:"Vul,omitempty"`
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

func (s *DescribeInstanceStatisticsResponseBodyData) SetHealth(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Health = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetSuspicious(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Suspicious = &v
	return s
}

func (s *DescribeInstanceStatisticsResponseBodyData) SetTrojan(v int32) *DescribeInstanceStatisticsResponseBodyData {
	s.Trojan = &v
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

type DescribeInstanceStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceStatisticsResponse) SetStatusCode(v int32) *DescribeInstanceStatisticsResponse {
	s.StatusCode = &v
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
	FailCount      *int32                                                 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	FailedEcsList  []*DescribeStrategyExecDetailResponseBodyFailedEcsList `json:"FailedEcsList,omitempty" xml:"FailedEcsList,omitempty" type:"Repeated"`
	InProcessCount *int32                                                 `json:"InProcessCount,omitempty" xml:"InProcessCount,omitempty"`
	Percent        *string                                                `json:"Percent,omitempty" xml:"Percent,omitempty"`
	RequestId      *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Source         *string                                                `json:"Source,omitempty" xml:"Source,omitempty"`
	StartTime      *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SuccessCount   *int32                                                 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
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

func (s *DescribeStrategyExecDetailResponseBody) SetFailCount(v int32) *DescribeStrategyExecDetailResponseBody {
	s.FailCount = &v
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

func (s *DescribeStrategyExecDetailResponseBody) SetPercent(v string) *DescribeStrategyExecDetailResponseBody {
	s.Percent = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetRequestId(v string) *DescribeStrategyExecDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBody) SetSource(v string) *DescribeStrategyExecDetailResponseBody {
	s.Source = &v
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

type DescribeStrategyExecDetailResponseBodyFailedEcsList struct {
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InternetIp   *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IntranetIp   *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeStrategyExecDetailResponseBodyFailedEcsList) String() string {
	return tea.Prettify(s)
}

func (s DescribeStrategyExecDetailResponseBodyFailedEcsList) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetIP(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.IP = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetInstanceName(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.InstanceName = &v
	return s
}

func (s *DescribeStrategyExecDetailResponseBodyFailedEcsList) SetInternetIp(v string) *DescribeStrategyExecDetailResponseBodyFailedEcsList {
	s.InternetIp = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeStrategyExecDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeStrategyExecDetailResponse) SetStatusCode(v int32) *DescribeStrategyExecDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStrategyExecDetailResponse) SetBody(v *DescribeStrategyExecDetailResponseBody) *DescribeStrategyExecDetailResponse {
	s.Body = v
	return s
}

type DescribeVulWhitelistRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeVulWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistRequest) SetCurrentPage(v int32) *DescribeVulWhitelistRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulWhitelistRequest) SetPageSize(v int32) *DescribeVulWhitelistRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulWhitelistRequest) SetSourceIp(v string) *DescribeVulWhitelistRequest {
	s.SourceIp = &v
	return s
}

type DescribeVulWhitelistResponseBody struct {
	CurrentPage   *int32                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VulWhitelists []*DescribeVulWhitelistResponseBodyVulWhitelists `json:"VulWhitelists,omitempty" xml:"VulWhitelists,omitempty" type:"Repeated"`
}

func (s DescribeVulWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulWhitelistResponseBody) SetCurrentPage(v int32) *DescribeVulWhitelistResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetPageSize(v int32) *DescribeVulWhitelistResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetRequestId(v string) *DescribeVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetTotalCount(v int32) *DescribeVulWhitelistResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulWhitelistResponseBody) SetVulWhitelists(v []*DescribeVulWhitelistResponseBodyVulWhitelists) *DescribeVulWhitelistResponseBody {
	s.VulWhitelists = v
	return s
}

type DescribeVulWhitelistResponseBodyVulWhitelists struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulWhitelistResponseBodyVulWhitelists) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulWhitelistResponseBodyVulWhitelists) GoString() string {
	return s.String()
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

func (s *DescribeVulWhitelistResponseBodyVulWhitelists) SetType(v string) *DescribeVulWhitelistResponseBodyVulWhitelists {
	s.Type = &v
	return s
}

type DescribeVulWhitelistResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVulWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVulWhitelistResponse) SetStatusCode(v int32) *DescribeVulWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulWhitelistResponse) SetBody(v *DescribeVulWhitelistResponseBody) *DescribeVulWhitelistResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	VmNumber     *string `json:"VmNumber,omitempty" xml:"VmNumber,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstanceRequest) SetPricingCycle(v string) *RenewInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewInstanceRequest) SetVmNumber(v string) *RenewInstanceRequest {
	s.VmNumber = &v
	return s
}

type RenewInstanceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetOrderId(v string) *RenewInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type UpgradeInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	VersionCode *int32  `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	VmNumber    *int32  `json:"VmNumber,omitempty" xml:"VmNumber,omitempty"`
}

func (s UpgradeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) SetClientToken(v string) *UpgradeInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceId(v string) *UpgradeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetOwnerId(v int64) *UpgradeInstanceRequest {
	s.OwnerId = &v
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
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBody) SetOrderId(v string) *UpgradeInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeInstanceResponseBody) SetRequestId(v string) *UpgradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpgradeInstanceResponse) SetStatusCode(v int32) *UpgradeInstanceResponse {
	s.StatusCode = &v
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoRenew)) {
		query["IsAutoRenew"] = request.IsAutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.VersionCode)) {
		query["VersionCode"] = request.VersionCode
	}

	if !tea.BoolValue(util.IsUnset(request.VmNumber)) {
		query["VmNumber"] = request.VmNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Action:      tea.String("DeleteRule"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoDelConfig"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoDelConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckWarningId)) {
		query["CheckWarningId"] = request.CheckWarningId
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
		Action:      tea.String("DescribeCheckWarningDetail"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCheckWarningDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
		Action:      tea.String("DescribeConcernNecessity"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConcernNecessityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceStatistics"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StrategyId)) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStrategyExecDetail"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStrategyExecDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeVulWhitelistWithOptions(request *DescribeVulWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeVulWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
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
		Action:      tea.String("DescribeVulWhitelist"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVulWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.VmNumber)) {
		query["VmNumber"] = request.VmNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionCode)) {
		query["VersionCode"] = request.VersionCode
	}

	if !tea.BoolValue(util.IsUnset(request.VmNumber)) {
		query["VmNumber"] = request.VmNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeInstance"),
		Version:     tea.String("2016-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
