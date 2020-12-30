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

type AddIpRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	IpList          *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AddIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
}

func (s *AddIpRequest) SetSourceIp(v string) *AddIpRequest {
	s.SourceIp = &v
	return s
}

func (s *AddIpRequest) SetIpList(v string) *AddIpRequest {
	s.IpList = &v
	return s
}

func (s *AddIpRequest) SetInstanceId(v string) *AddIpRequest {
	s.InstanceId = &v
	return s
}

func (s *AddIpRequest) SetRegionId(v string) *AddIpRequest {
	s.RegionId = &v
	return s
}

func (s *AddIpRequest) SetResourceGroupId(v string) *AddIpRequest {
	s.ResourceGroupId = &v
	return s
}

type AddIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpResponseBody) SetRequestId(v string) *AddIpResponseBody {
	s.RequestId = &v
	return s
}

type AddIpResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponse) GoString() string {
	return s.String()
}

func (s *AddIpResponse) SetHeaders(v map[string]*string) *AddIpResponse {
	s.Headers = v
	return s
}

func (s *AddIpResponse) SetBody(v *AddIpResponseBody) *AddIpResponse {
	s.Body = v
	return s
}

type CheckGrantRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) SetSourceIp(v string) *CheckGrantRequest {
	s.SourceIp = &v
	return s
}

func (s *CheckGrantRequest) SetResourceGroupId(v string) *CheckGrantRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckGrantRequest) SetRegionId(v string) *CheckGrantRequest {
	s.RegionId = &v
	return s
}

type CheckGrantResponseBody struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckGrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGrantResponseBody) SetStatus(v int32) *CheckGrantResponseBody {
	s.Status = &v
	return s
}

func (s *CheckGrantResponseBody) SetRequestId(v string) *CheckGrantResponseBody {
	s.RequestId = &v
	return s
}

type CheckGrantResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckGrantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponse) GoString() string {
	return s.String()
}

func (s *CheckGrantResponse) SetHeaders(v map[string]*string) *CheckGrantResponse {
	s.Headers = v
	return s
}

func (s *CheckGrantResponse) SetBody(v *CheckGrantResponseBody) *CheckGrantResponse {
	s.Body = v
	return s
}

type ConfigSchedruleOnDemandRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandRequest) SetSourceIp(v string) *ConfigSchedruleOnDemandRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetInstanceId(v string) *ConfigSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleName(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleAction(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleSwitch(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoMode(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoBeginTime(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoEndTime(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoEndTime = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetTimeZone(v string) *ConfigSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRegionId(v string) *ConfigSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

type ConfigSchedruleOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandResponseBody) SetRequestId(v string) *ConfigSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type ConfigSchedruleOnDemandResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *ConfigSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *ConfigSchedruleOnDemandResponse) SetBody(v *ConfigSchedruleOnDemandResponseBody) *ConfigSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type CreateSchedruleOnDemandRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandRequest) SetSourceIp(v string) *CreateSchedruleOnDemandRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetInstanceId(v string) *CreateSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleName(v string) *CreateSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleAction(v string) *CreateSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleSwitch(v string) *CreateSchedruleOnDemandRequest {
	s.RuleSwitch = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoMode(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoBeginTime(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoEndTime(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoEndTime = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetTimeZone(v string) *CreateSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRegionId(v string) *CreateSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

type CreateSchedruleOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandResponseBody) SetRequestId(v string) *CreateSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type CreateSchedruleOnDemandResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *CreateSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *CreateSchedruleOnDemandResponse) SetBody(v *CreateSchedruleOnDemandResponseBody) *CreateSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DeleteBlackholeRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBlackholeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeRequest) SetSourceIp(v string) *DeleteBlackholeRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteBlackholeRequest) SetIp(v string) *DeleteBlackholeRequest {
	s.Ip = &v
	return s
}

func (s *DeleteBlackholeRequest) SetInstanceId(v string) *DeleteBlackholeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetResourceGroupId(v string) *DeleteBlackholeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetRegionId(v string) *DeleteBlackholeRequest {
	s.RegionId = &v
	return s
}

type DeleteBlackholeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBlackholeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponseBody) SetRequestId(v string) *DeleteBlackholeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBlackholeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBlackholeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBlackholeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponse) SetHeaders(v map[string]*string) *DeleteBlackholeResponse {
	s.Headers = v
	return s
}

func (s *DeleteBlackholeResponse) SetBody(v *DeleteBlackholeResponseBody) *DeleteBlackholeResponse {
	s.Body = v
	return s
}

type DeleteIpRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	IpList          *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpRequest) SetSourceIp(v string) *DeleteIpRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteIpRequest) SetIpList(v string) *DeleteIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteIpRequest) SetInstanceId(v string) *DeleteIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIpRequest) SetResourceGroupId(v string) *DeleteIpRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteIpRequest) SetRegionId(v string) *DeleteIpRequest {
	s.RegionId = &v
	return s
}

type DeleteIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpResponseBody) SetRequestId(v string) *DeleteIpResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpResponse) SetHeaders(v map[string]*string) *DeleteIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpResponse) SetBody(v *DeleteIpResponseBody) *DeleteIpResponse {
	s.Body = v
	return s
}

type DeleteSchedruleOnDemandRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandRequest) SetSourceIp(v string) *DeleteSchedruleOnDemandRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetInstanceId(v string) *DeleteSchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetRuleName(v string) *DeleteSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetRegionId(v string) *DeleteSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

type DeleteSchedruleOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandResponseBody) SetRequestId(v string) *DeleteSchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSchedruleOnDemandResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandResponse) SetHeaders(v map[string]*string) *DeleteSchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchedruleOnDemandResponse) SetBody(v *DeleteSchedruleOnDemandResponseBody) *DeleteSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DescribeDdosEventRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime       *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventRequest) SetSourceIp(v string) *DescribeDdosEventRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosEventRequest) SetStartTime(v int32) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetEndTime(v int32) *DescribeDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int32) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageNo(v int32) *DescribeDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDdosEventRequest) SetInstanceId(v string) *DescribeDdosEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetIp(v string) *DescribeDdosEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventRequest) SetResourceGroupId(v string) *DescribeDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetRegionId(v string) *DescribeDdosEventRequest {
	s.RegionId = &v
	return s
}

type DescribeDdosEventResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Events    []*DescribeDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	Total     *int64                                 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBody) SetRequestId(v string) *DescribeDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponseBody) SetEvents(v []*DescribeDdosEventResponseBodyEvents) *DescribeDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDdosEventResponseBody) SetTotal(v int64) *DescribeDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosEventResponseBodyEvents struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Mbps      *int32  `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Pps       *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeDdosEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStatus(v string) *DescribeDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetIp(v string) *DescribeDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetPps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

type DescribeDdosEventResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponse) SetHeaders(v map[string]*string) *DescribeDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosEventResponse) SetBody(v *DescribeDdosEventResponseBody) *DescribeDdosEventResponse {
	s.Body = v
	return s
}

type DescribeExcpetionCountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExcpetionCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountRequest) SetSourceIp(v string) *DescribeExcpetionCountRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeExcpetionCountRequest) SetResourceGroupId(v string) *DescribeExcpetionCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExcpetionCountRequest) SetRegionId(v string) *DescribeExcpetionCountRequest {
	s.RegionId = &v
	return s
}

type DescribeExcpetionCountResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExceptionIpCount *int32  `json:"ExceptionIpCount,omitempty" xml:"ExceptionIpCount,omitempty"`
	ExpireTimeCount  *int32  `json:"ExpireTimeCount,omitempty" xml:"ExpireTimeCount,omitempty"`
}

func (s DescribeExcpetionCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponseBody) SetRequestId(v string) *DescribeExcpetionCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetExceptionIpCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExceptionIpCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetExpireTimeCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExpireTimeCount = &v
	return s
}

type DescribeExcpetionCountResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExcpetionCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExcpetionCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponse) SetHeaders(v map[string]*string) *DescribeExcpetionCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeExcpetionCountResponse) SetBody(v *DescribeExcpetionCountResponseBody) *DescribeExcpetionCountResponse {
	s.Body = v
	return s
}

type DescribeInstanceListRequest struct {
	SourceIp        *string                           `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIdList  *string                           `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	Remark          *string                           `json:"Remark,omitempty" xml:"Remark,omitempty"`
	PageNo          *int32                            `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IpVersion       *string                           `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	InstanceType    *string                           `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Ip              *string                           `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Orderby         *string                           `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	Orderdire       *string                           `json:"Orderdire,omitempty" xml:"Orderdire,omitempty"`
	RegionId        *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag             []*DescribeInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequest) SetSourceIp(v string) *DescribeInstanceListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceListRequest) SetResourceGroupId(v string) *DescribeInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceIdList(v string) *DescribeInstanceListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRemark(v string) *DescribeInstanceListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageNo(v int32) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int32) *DescribeInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceListRequest) SetIpVersion(v string) *DescribeInstanceListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceType(v string) *DescribeInstanceListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListRequest) SetIp(v string) *DescribeInstanceListRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderby(v string) *DescribeInstanceListRequest {
	s.Orderby = &v
	return s
}

func (s *DescribeInstanceListRequest) SetOrderdire(v string) *DescribeInstanceListRequest {
	s.Orderdire = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRegionId(v string) *DescribeInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetTag(v []*DescribeInstanceListRequestTag) *DescribeInstanceListRequest {
	s.Tag = v
	return s
}

type DescribeInstanceListRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequestTag) SetKey(v string) *DescribeInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceListRequestTag) SetValue(v string) *DescribeInstanceListRequestTag {
	s.Value = &v
	return s
}

type DescribeInstanceListResponseBody struct {
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total        *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
	InstanceList []*DescribeInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBody) SetRequestId(v string) *DescribeInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponseBody) SetTotal(v int64) *DescribeInstanceListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeInstanceListResponseBody) SetInstanceList(v []*DescribeInstanceListResponseBodyInstanceList) *DescribeInstanceListResponseBody {
	s.InstanceList = v
	return s
}

type DescribeInstanceListResponseBodyInstanceList struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	IpType            *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	AutoRenewal       *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Product           *string `json:"Product,omitempty" xml:"Product,omitempty"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType      *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	BlackholdingCount *string `json:"BlackholdingCount,omitempty" xml:"BlackholdingCount,omitempty"`
}

func (s DescribeInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetStatus(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetIpType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetAutoRenewal(v bool) *DescribeInstanceListResponseBodyInstanceList {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetRemark(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetProduct(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetGmtCreate(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceId(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetInstanceType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetBlackholdingCount(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.BlackholdingCount = &v
	return s
}

type DescribeInstanceListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponse) SetHeaders(v map[string]*string) *DescribeInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceListResponse) SetBody(v *DescribeInstanceListResponseBody) *DescribeInstanceListResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceIdList  *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) SetSourceIp(v string) *DescribeInstanceSpecsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetLang(v string) *DescribeInstanceSpecsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIdList(v string) *DescribeInstanceSpecsRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetRegionId(v string) *DescribeInstanceSpecsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSpecsRequest) SetResourceGroupId(v string) *DescribeInstanceSpecsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeInstanceSpecsResponseBody struct {
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	PackConfig                    *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" type:"Struct"`
	Region                        *string                                                   `json:"Region,omitempty" xml:"Region,omitempty"`
	AvailableDeleteBlackholeCount *string                                                   `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty"`
	InstanceId                    *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPackConfig(v *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PackConfig = v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetRegion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.Region = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDeleteBlackholeCount(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig struct {
	PackBasicThre *int32 `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty"`
	BindIpCount   *int32 `json:"BindIpCount,omitempty" xml:"BindIpCount,omitempty"`
	PackAdvThre   *int32 `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty"`
	IpBasicThre   *int32 `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty"`
	IpAdvanceThre *int32 `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty"`
	IpSpec        *int32 `json:"IpSpec,omitempty" xml:"IpSpec,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBindIpCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.BindIpCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackAdvThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpAdvanceThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpSpec(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpSpec = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

type DescribeOnDemandDdosEventRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StartTime       *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOnDemandDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventRequest) SetSourceIp(v string) *DescribeOnDemandDdosEventRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetStartTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetEndTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageSize(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageNo(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetInstanceId(v string) *DescribeOnDemandDdosEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetIp(v string) *DescribeOnDemandDdosEventRequest {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetResourceGroupId(v string) *DescribeOnDemandDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetRegionId(v string) *DescribeOnDemandDdosEventRequest {
	s.RegionId = &v
	return s
}

type DescribeOnDemandDdosEventResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Events    []*DescribeOnDemandDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	Total     *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeOnDemandDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseBody) SetRequestId(v string) *DescribeOnDemandDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetEvents(v []*DescribeOnDemandDdosEventResponseBodyEvents) *DescribeOnDemandDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetTotal(v int64) *DescribeOnDemandDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeOnDemandDdosEventResponseBodyEvents struct {
	EndTime   *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Mbps      *int32  `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Pps       *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
}

func (s DescribeOnDemandDdosEventResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetEndTime(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStatus(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetIp(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetPps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

type DescribeOnDemandDdosEventResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOnDemandDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOnDemandDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponse) SetHeaders(v map[string]*string) *DescribeOnDemandDdosEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetBody(v *DescribeOnDemandDdosEventResponseBody) *DescribeOnDemandDdosEventResponse {
	s.Body = v
	return s
}

type DescribeOnDemandInstanceStatusRequest struct {
	SourceIp       *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s DescribeOnDemandInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusRequest) SetSourceIp(v string) *DescribeOnDemandInstanceStatusRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusRequest) SetRegionId(v string) *DescribeOnDemandInstanceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusRequest) SetInstanceIdList(v []*string) *DescribeOnDemandInstanceStatusRequest {
	s.InstanceIdList = v
	return s
}

type DescribeOnDemandInstanceStatusResponseBody struct {
	Instances []*DescribeOnDemandInstanceStatusResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseBody) SetInstances(v []*DescribeOnDemandInstanceStatusResponseBodyInstances) *DescribeOnDemandInstanceStatusResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBody) SetRequestId(v string) *DescribeOnDemandInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponseBodyInstances struct {
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Declared   *string `json:"Declared,omitempty" xml:"Declared,omitempty"`
	RegistedAs *string `json:"RegistedAs,omitempty" xml:"RegistedAs,omitempty"`
	Net        *string `json:"Net,omitempty" xml:"Net,omitempty"`
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetUserId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.UserId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetMode(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Mode = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetInstanceId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDeclared(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Declared = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetRegistedAs(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.RegistedAs = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetNet(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Net = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDesc(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Desc = &v
	return s
}

type DescribeOnDemandInstanceStatusResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOnDemandInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOnDemandInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeOnDemandInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponse) SetBody(v *DescribeOnDemandInstanceStatusResponseBody) *DescribeOnDemandInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrderBy         *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderDir        *string `json:"OrderDir,omitempty" xml:"OrderDir,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetLang(v string) *DescribeOpEntitiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int32) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderBy(v string) *DescribeOpEntitiesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetOrderDir(v string) *DescribeOpEntitiesRequest {
	s.OrderDir = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetInstanceId(v string) *DescribeOpEntitiesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetRegionId(v string) *DescribeOpEntitiesRequest {
	s.RegionId = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int32) *DescribeOpEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	EntityType   *int32  `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	OpAction     *int32  `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetHeaders(v map[string]*string) *DescribeOpEntitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribePackIpListRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	ProductName     *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePackIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListRequest) GoString() string {
	return s.String()
}

func (s *DescribePackIpListRequest) SetSourceIp(v string) *DescribePackIpListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageNo(v int32) *DescribePackIpListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageSize(v int32) *DescribePackIpListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackIpListRequest) SetInstanceId(v string) *DescribePackIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackIpListRequest) SetIp(v string) *DescribePackIpListRequest {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListRequest) SetProductName(v string) *DescribePackIpListRequest {
	s.ProductName = &v
	return s
}

func (s *DescribePackIpListRequest) SetResourceGroupId(v string) *DescribePackIpListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePackIpListRequest) SetRegionId(v string) *DescribePackIpListRequest {
	s.RegionId = &v
	return s
}

type DescribePackIpListResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpList    []*DescribePackIpListResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	Total     *int32                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePackIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBody) SetRequestId(v string) *DescribePackIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetIpList(v []*DescribePackIpListResponseBodyIpList) *DescribePackIpListResponseBody {
	s.IpList = v
	return s
}

func (s *DescribePackIpListResponseBody) SetTotal(v int32) *DescribePackIpListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetCode(v string) *DescribePackIpListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetSuccess(v bool) *DescribePackIpListResponseBody {
	s.Success = &v
	return s
}

type DescribePackIpListResponseBodyIpList struct {
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribePackIpListResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBodyIpList) SetStatus(v string) *DescribePackIpListResponseBodyIpList {
	s.Status = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRemark(v string) *DescribePackIpListResponseBodyIpList {
	s.Remark = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetProduct(v string) *DescribePackIpListResponseBodyIpList {
	s.Product = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetIp(v string) *DescribePackIpListResponseBodyIpList {
	s.Ip = &v
	return s
}

type DescribePackIpListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePackIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponse) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponse) SetHeaders(v map[string]*string) *DescribePackIpListResponse {
	s.Headers = v
	return s
}

func (s *DescribePackIpListResponse) SetBody(v *DescribePackIpListResponseBody) *DescribePackIpListResponse {
	s.Body = v
	return s
}

type DescribePackPaidTrafficRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribePackPaidTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficRequest) SetSourceIp(v string) *DescribePackPaidTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetInstanceId(v string) *DescribePackPaidTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetCurrentPage(v int32) *DescribePackPaidTrafficRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetPageSize(v int32) *DescribePackPaidTrafficRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetStartTime(v int64) *DescribePackPaidTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetEndTime(v int64) *DescribePackPaidTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetResourceGroupId(v string) *DescribePackPaidTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribePackPaidTrafficResponseBody struct {
	TotalCount       *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PackPaidTraffics []*DescribePackPaidTrafficResponseBodyPackPaidTraffics `json:"PackPaidTraffics,omitempty" xml:"PackPaidTraffics,omitempty" type:"Repeated"`
}

func (s DescribePackPaidTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponseBody) SetTotalCount(v int32) *DescribePackPaidTrafficResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBody) SetRequestId(v string) *DescribePackPaidTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBody) SetPackPaidTraffics(v []*DescribePackPaidTrafficResponseBodyPackPaidTraffics) *DescribePackPaidTrafficResponseBody {
	s.PackPaidTraffics = v
	return s
}

type DescribePackPaidTrafficResponseBodyPackPaidTraffics struct {
	StartTime        *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	BaseBandwidth    *int32   `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty"`
	ElasticBandwidth *int32   `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	TotalCapacity    *float32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	MaxAttack        *float32 `json:"MaxAttack,omitempty" xml:"MaxAttack,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PaidCapacity     *float32 `json:"PaidCapacity,omitempty" xml:"PaidCapacity,omitempty"`
}

func (s DescribePackPaidTrafficResponseBodyPackPaidTraffics) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponseBodyPackPaidTraffics) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetStartTime(v int64) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.StartTime = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetBaseBandwidth(v int32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetElasticBandwidth(v int32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetTotalCapacity(v float32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.TotalCapacity = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetMaxAttack(v float32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.MaxAttack = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetInstanceId(v string) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.InstanceId = &v
	return s
}

func (s *DescribePackPaidTrafficResponseBodyPackPaidTraffics) SetPaidCapacity(v float32) *DescribePackPaidTrafficResponseBodyPackPaidTraffics {
	s.PaidCapacity = &v
	return s
}

type DescribePackPaidTrafficResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePackPaidTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackPaidTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponse) SetHeaders(v map[string]*string) *DescribePackPaidTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribePackPaidTrafficResponse) SetBody(v *DescribePackPaidTrafficResponseBody) *DescribePackPaidTrafficResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetSourceIp(v string) *DescribeRegionsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceGroupId(v string) *DescribeRegionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeResourcePackInstancesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeResourcePackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesRequest) SetSourceIp(v string) *DescribeResourcePackInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetPageSize(v int32) *DescribeResourcePackInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetCurrentPage(v int32) *DescribeResourcePackInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetResourceGroupId(v string) *DescribeResourcePackInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeResourcePackInstancesResponseBody struct {
	TotalCount    *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId     *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourcePacks []*DescribeResourcePackInstancesResponseBodyResourcePacks `json:"ResourcePacks,omitempty" xml:"ResourcePacks,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponseBody) SetTotalCount(v int32) *DescribeResourcePackInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBody) SetRequestId(v string) *DescribeResourcePackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBody) SetResourcePacks(v []*DescribeResourcePackInstancesResponseBodyResourcePacks) *DescribeResourcePackInstancesResponseBody {
	s.ResourcePacks = v
	return s
}

type DescribeResourcePackInstancesResponseBodyResourcePacks struct {
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ResourcePackId *string `json:"ResourcePackId,omitempty" xml:"ResourcePackId,omitempty"`
	CurrCapacity   *int64  `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	InitCapacity   *int64  `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
}

func (s DescribeResourcePackInstancesResponseBodyResourcePacks) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponseBodyResourcePacks) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetEndTime(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetStatus(v string) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.Status = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetStartTime(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetResourcePackId(v string) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.ResourcePackId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetCurrCapacity(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseBodyResourcePacks) SetInitCapacity(v int64) *DescribeResourcePackInstancesResponseBodyResourcePacks {
	s.InitCapacity = &v
	return s
}

type DescribeResourcePackInstancesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourcePackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponse) SetHeaders(v map[string]*string) *DescribeResourcePackInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackInstancesResponse) SetBody(v *DescribeResourcePackInstancesResponseBody) *DescribeResourcePackInstancesResponse {
	s.Body = v
	return s
}

type DescribeResourcePackStatisticsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	DdosRegionId    *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeResourcePackStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsRequest) SetSourceIp(v string) *DescribeResourcePackStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeResourcePackStatisticsRequest) SetDdosRegionId(v string) *DescribeResourcePackStatisticsRequest {
	s.DdosRegionId = &v
	return s
}

func (s *DescribeResourcePackStatisticsRequest) SetInstanceId(v string) *DescribeResourcePackStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourcePackStatisticsRequest) SetResourceGroupId(v string) *DescribeResourcePackStatisticsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeResourcePackStatisticsResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalInitCapacity *int64  `json:"TotalInitCapacity,omitempty" xml:"TotalInitCapacity,omitempty"`
	TotalCurrCapacity *int64  `json:"TotalCurrCapacity,omitempty" xml:"TotalCurrCapacity,omitempty"`
	AvailablePackNum  *int32  `json:"AvailablePackNum,omitempty" xml:"AvailablePackNum,omitempty"`
	TotalUsedCapacity *int64  `json:"TotalUsedCapacity,omitempty" xml:"TotalUsedCapacity,omitempty"`
}

func (s DescribeResourcePackStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsResponseBody) SetRequestId(v string) *DescribeResourcePackStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetTotalInitCapacity(v int64) *DescribeResourcePackStatisticsResponseBody {
	s.TotalInitCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetTotalCurrCapacity(v int64) *DescribeResourcePackStatisticsResponseBody {
	s.TotalCurrCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetAvailablePackNum(v int32) *DescribeResourcePackStatisticsResponseBody {
	s.AvailablePackNum = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponseBody) SetTotalUsedCapacity(v int64) *DescribeResourcePackStatisticsResponseBody {
	s.TotalUsedCapacity = &v
	return s
}

type DescribeResourcePackStatisticsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourcePackStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePackStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsResponse) SetHeaders(v map[string]*string) *DescribeResourcePackStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetBody(v *DescribeResourcePackStatisticsResponseBody) *DescribeResourcePackStatisticsResponse {
	s.Body = v
	return s
}

type DescribeResourcePackUsageRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeResourcePackUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageRequest) SetSourceIp(v string) *DescribeResourcePackUsageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetEndTime(v int64) *DescribeResourcePackUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetStartTime(v int64) *DescribeResourcePackUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetInstanceId(v string) *DescribeResourcePackUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetResourceGroupId(v string) *DescribeResourcePackUsageRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeResourcePackUsageResponseBody struct {
	EndTime    *int64                                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PackUsages []*DescribeResourcePackUsageResponseBodyPackUsages `json:"PackUsages,omitempty" xml:"PackUsages,omitempty" type:"Repeated"`
	StartTime  *int64                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Interval   *int64                                             `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeResourcePackUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponseBody) SetEndTime(v int64) *DescribeResourcePackUsageResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetRequestId(v string) *DescribeResourcePackUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetPackUsages(v []*DescribeResourcePackUsageResponseBodyPackUsages) *DescribeResourcePackUsageResponseBody {
	s.PackUsages = v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetStartTime(v int64) *DescribeResourcePackUsageResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBody) SetInterval(v int64) *DescribeResourcePackUsageResponseBody {
	s.Interval = &v
	return s
}

type DescribeResourcePackUsageResponseBodyPackUsages struct {
	Time    *int64   `json:"Time,omitempty" xml:"Time,omitempty"`
	Traffic *float32 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeResourcePackUsageResponseBodyPackUsages) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponseBodyPackUsages) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponseBodyPackUsages) SetTime(v int64) *DescribeResourcePackUsageResponseBodyPackUsages {
	s.Time = &v
	return s
}

func (s *DescribeResourcePackUsageResponseBodyPackUsages) SetTraffic(v float32) *DescribeResourcePackUsageResponseBodyPackUsages {
	s.Traffic = &v
	return s
}

type DescribeResourcePackUsageResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourcePackUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePackUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponse) SetHeaders(v map[string]*string) *DescribeResourcePackUsageResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetBody(v *DescribeResourcePackUsageResponseBody) *DescribeResourcePackUsageResponse {
	s.Body = v
	return s
}

type DescribeTrafficRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ipnet           *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	StartTime       *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval        *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficRequest) SetSourceIp(v string) *DescribeTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTrafficRequest) SetInstanceId(v string) *DescribeTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTrafficRequest) SetIpnet(v string) *DescribeTrafficRequest {
	s.Ipnet = &v
	return s
}

func (s *DescribeTrafficRequest) SetIp(v string) *DescribeTrafficRequest {
	s.Ip = &v
	return s
}

func (s *DescribeTrafficRequest) SetStartTime(v int32) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetEndTime(v int32) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int32) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeTrafficRequest) SetResourceGroupId(v string) *DescribeTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTrafficRequest) SetRegionId(v string) *DescribeTrafficRequest {
	s.RegionId = &v
	return s
}

type DescribeTrafficResponseBody struct {
	FlowList  []*DescribeTrafficResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBody) SetFlowList(v []*DescribeTrafficResponseBodyFlowList) *DescribeTrafficResponseBody {
	s.FlowList = v
	return s
}

func (s *DescribeTrafficResponseBody) SetRequestId(v string) *DescribeTrafficResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTrafficResponseBodyFlowList struct {
	Time      *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
	FlowType  *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	AttackPps *int64  `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pps       *int32  `json:"Pps,omitempty" xml:"Pps,omitempty"`
	Kbps      *int32  `json:"Kbps,omitempty" xml:"Kbps,omitempty"`
	AttackBps *int64  `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
}

func (s DescribeTrafficResponseBodyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBodyFlowList) SetTime(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Time = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetFlowType(v string) *DescribeTrafficResponseBodyFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackPps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetName(v string) *DescribeTrafficResponseBodyFlowList {
	s.Name = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetPps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Pps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetKbps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Kbps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackBps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackBps = &v
	return s
}

type DescribeTrafficResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponse) SetHeaders(v map[string]*string) *DescribeTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficResponse) SetBody(v *DescribeTrafficResponseBody) *DescribeTrafficResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TagOwnerUid     *string `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
	TagOwnerBid     *string `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetSourceIp(v string) *ListTagKeysRequest {
	s.SourceIp = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetTagOwnerUid(v string) *ListTagKeysRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *ListTagKeysRequest) SetTagOwnerBid(v string) *ListTagKeysRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetScope(v string) *ListTagKeysRequest {
	s.Scope = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

type ListTagKeysResponseBody struct {
	TotalCount  *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TagKeys     []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	SourceIp        *string                       `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TagOwnerUid     *string                       `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
	TagOwnerBid     *string                       `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	ResourceType    *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope           *string                       `json:"Scope,omitempty" xml:"Scope,omitempty"`
	NextToken       *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId      []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag             []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetSourceIp(v string) *ListTagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetTagOwnerUid(v string) *ListTagResourcesRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *ListTagResourcesRequest) SetTagOwnerBid(v string) *ListTagResourcesRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetScope(v string) *ListTagResourcesRequest {
	s.Scope = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyRemarkRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyRemarkRequest) SetSourceIp(v string) *ModifyRemarkRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyRemarkRequest) SetLang(v string) *ModifyRemarkRequest {
	s.Lang = &v
	return s
}

func (s *ModifyRemarkRequest) SetInstanceId(v string) *ModifyRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRemark(v string) *ModifyRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyRemarkRequest) SetResourceGroupId(v string) *ModifyRemarkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyRemarkRequest) SetRegionId(v string) *ModifyRemarkRequest {
	s.RegionId = &v
	return s
}

type ModifyRemarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponseBody) SetRequestId(v string) *ModifyRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ModifyRemarkResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponse) SetHeaders(v map[string]*string) *ModifyRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyRemarkResponse) SetBody(v *ModifyRemarkResponseBody) *ModifyRemarkResponse {
	s.Body = v
	return s
}

type QuerySchedruleOnDemandRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QuerySchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandRequest) SetSourceIp(v string) *QuerySchedruleOnDemandRequest {
	s.SourceIp = &v
	return s
}

func (s *QuerySchedruleOnDemandRequest) SetInstanceId(v string) *QuerySchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandRequest) SetRegionId(v string) *QuerySchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

type QuerySchedruleOnDemandResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string                                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RuleStatus []*QuerySchedruleOnDemandResponseBodyRuleStatus `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty" type:"Repeated"`
	RuleConfig []*QuerySchedruleOnDemandResponseBodyRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" type:"Repeated"`
}

func (s QuerySchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBody) SetRequestId(v string) *QuerySchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetInstanceId(v string) *QuerySchedruleOnDemandResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetUserId(v string) *QuerySchedruleOnDemandResponseBody {
	s.UserId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleStatus(v []*QuerySchedruleOnDemandResponseBodyRuleStatus) *QuerySchedruleOnDemandResponseBody {
	s.RuleStatus = v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleConfig(v []*QuerySchedruleOnDemandResponseBodyRuleConfig) *QuerySchedruleOnDemandResponseBody {
	s.RuleConfig = v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleStatus struct {
	RuleSchedStatus *string `json:"RuleSchedStatus,omitempty" xml:"RuleSchedStatus,omitempty"`
	Net             *string `json:"Net,omitempty" xml:"Net,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetRuleSchedStatus(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.RuleSchedStatus = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetNet(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.Net = &v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleConfig struct {
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleSwitch(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleSwitch = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionMbps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionMbps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetTimeZone(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.TimeZone = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleAction(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleAction = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionKpps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionKpps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoMode(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoMode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoBeginTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionCnt(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionCnt = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoEndTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoEndTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleName(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleName = &v
	return s
}

type QuerySchedruleOnDemandResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponse) SetHeaders(v map[string]*string) *QuerySchedruleOnDemandResponse {
	s.Headers = v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetBody(v *QuerySchedruleOnDemandResponseBody) *QuerySchedruleOnDemandResponse {
	s.Body = v
	return s
}

type SetInstanceModeOnDemandRequest struct {
	SourceIp       *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s SetInstanceModeOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandRequest) SetSourceIp(v string) *SetInstanceModeOnDemandRequest {
	s.SourceIp = &v
	return s
}

func (s *SetInstanceModeOnDemandRequest) SetMode(v string) *SetInstanceModeOnDemandRequest {
	s.Mode = &v
	return s
}

func (s *SetInstanceModeOnDemandRequest) SetRegionId(v string) *SetInstanceModeOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *SetInstanceModeOnDemandRequest) SetInstanceIdList(v []*string) *SetInstanceModeOnDemandRequest {
	s.InstanceIdList = v
	return s
}

type SetInstanceModeOnDemandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceModeOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandResponseBody) SetRequestId(v string) *SetInstanceModeOnDemandResponseBody {
	s.RequestId = &v
	return s
}

type SetInstanceModeOnDemandResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetInstanceModeOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetInstanceModeOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandResponse) SetHeaders(v map[string]*string) *SetInstanceModeOnDemandResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceModeOnDemandResponse) SetBody(v *SetInstanceModeOnDemandResponseBody) *SetInstanceModeOnDemandResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	SourceIp        *string                   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TagOwnerUid     *string                   `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
	TagOwnerBid     *string                   `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	ResourceType    *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope           *string                   `json:"Scope,omitempty" xml:"Scope,omitempty"`
	ResourceId      []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag             []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetSourceIp(v string) *TagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetTagOwnerUid(v string) *TagResourcesRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *TagResourcesRequest) SetTagOwnerBid(v string) *TagResourcesRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetScope(v string) *TagResourcesRequest {
	s.Scope = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	SourceIp        *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TagOwnerUid     *string   `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
	TagOwnerBid     *string   `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	ResourceType    *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All             *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId      []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey          []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetSourceIp(v string) *UntagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetTagOwnerUid(v string) *UntagResourcesRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *UntagResourcesRequest) SetTagOwnerBid(v string) *UntagResourcesRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
		"cn-qingdao":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-beijing":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("ddosbgp.aliyuncs.com"),
		"cn-huhehaote":          tea.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen":           tea.String("ddosbgp.aliyuncs.com"),
		"ap-northeast-1":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("ddosbgp.aliyuncs.com"),
		"eu-central-1":          tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("ddosbgp.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddosbgp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddIpWithOptions(request *AddIpRequest, runtime *util.RuntimeOptions) (_result *AddIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddIp"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIp(request *AddIpRequest) (_result *AddIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpResponse{}
	_body, _err := client.AddIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *util.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckGrant"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckGrant(request *CheckGrantRequest) (_result *CheckGrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckGrantResponse{}
	_body, _err := client.CheckGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigSchedruleOnDemandWithOptions(request *ConfigSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *ConfigSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigSchedruleOnDemand"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigSchedruleOnDemand(request *ConfigSchedruleOnDemandRequest) (_result *ConfigSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.ConfigSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSchedruleOnDemandWithOptions(request *CreateSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *CreateSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSchedruleOnDemand"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSchedruleOnDemand(request *CreateSchedruleOnDemandRequest) (_result *CreateSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.CreateSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *util.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBlackhole"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBlackhole(request *DeleteBlackholeRequest) (_result *DeleteBlackholeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DeleteBlackholeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *util.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteIp"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIp(request *DeleteIpRequest) (_result *DeleteIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpResponse{}
	_body, _err := client.DeleteIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSchedruleOnDemandWithOptions(request *DeleteSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *DeleteSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSchedruleOnDemand"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSchedruleOnDemand(request *DeleteSchedruleOnDemandRequest) (_result *DeleteSchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.DeleteSchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDdosEvent"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosEvent(request *DescribeDdosEventRequest) (_result *DescribeDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DescribeDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExcpetionCountWithOptions(request *DescribeExcpetionCountRequest, runtime *util.RuntimeOptions) (_result *DescribeExcpetionCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExcpetionCount"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExcpetionCount(request *DescribeExcpetionCountRequest) (_result *DescribeExcpetionCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.DescribeExcpetionCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceList"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (_result *DescribeInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DescribeInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceSpecs"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOnDemandDdosEventWithOptions(request *DescribeOnDemandDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOnDemandDdosEvent"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOnDemandDdosEvent(request *DescribeOnDemandDdosEventRequest) (_result *DescribeOnDemandDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.DescribeOnDemandDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOnDemandInstanceStatusWithOptions(request *DescribeOnDemandInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOnDemandInstanceStatus"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOnDemandInstanceStatus(request *DescribeOnDemandInstanceStatusRequest) (_result *DescribeOnDemandInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.DescribeOnDemandInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOpEntities"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackIpListWithOptions(request *DescribePackIpListRequest, runtime *util.RuntimeOptions) (_result *DescribePackIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePackIpList"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackIpList(request *DescribePackIpListRequest) (_result *DescribePackIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.DescribePackIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackPaidTrafficWithOptions(request *DescribePackPaidTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribePackPaidTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePackPaidTrafficResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePackPaidTraffic"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackPaidTraffic(request *DescribePackPaidTrafficRequest) (_result *DescribePackPaidTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackPaidTrafficResponse{}
	_body, _err := client.DescribePackPaidTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackInstancesWithOptions(request *DescribeResourcePackInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeResourcePackInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeResourcePackInstances"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackInstances(request *DescribeResourcePackInstancesRequest) (_result *DescribeResourcePackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackInstancesResponse{}
	_body, _err := client.DescribeResourcePackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackStatisticsWithOptions(request *DescribeResourcePackStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeResourcePackStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeResourcePackStatistics"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackStatistics(request *DescribeResourcePackStatisticsRequest) (_result *DescribeResourcePackStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackStatisticsResponse{}
	_body, _err := client.DescribeResourcePackStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackUsageWithOptions(request *DescribeResourcePackUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeResourcePackUsageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeResourcePackUsage"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackUsage(request *DescribeResourcePackUsageRequest) (_result *DescribeResourcePackUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackUsageResponse{}
	_body, _err := client.DescribeResourcePackUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTraffic"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraffic(request *DescribeTrafficRequest) (_result *DescribeTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DescribeTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRemarkWithOptions(request *ModifyRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyRemark"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRemark(request *ModifyRemarkRequest) (_result *ModifyRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.ModifyRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySchedruleOnDemandWithOptions(request *QuerySchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *QuerySchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySchedruleOnDemand"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySchedruleOnDemand(request *QuerySchedruleOnDemandRequest) (_result *QuerySchedruleOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.QuerySchedruleOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetInstanceModeOnDemandWithOptions(request *SetInstanceModeOnDemandRequest, runtime *util.RuntimeOptions) (_result *SetInstanceModeOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetInstanceModeOnDemand"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetInstanceModeOnDemand(request *SetInstanceModeOnDemandRequest) (_result *SetInstanceModeOnDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.SetInstanceModeOnDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2018-07-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
