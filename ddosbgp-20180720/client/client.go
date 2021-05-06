// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeOnDemandInstanceStatusRequest struct {
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" require:"true" type:"Repeated"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusRequest) SetInstanceIdList(v []*string) *DescribeOnDemandInstanceStatusRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeOnDemandInstanceStatusRequest) SetRegionId(v string) *DescribeOnDemandInstanceStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponse struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Instances []*DescribeOnDemandInstanceStatusResponseInstances `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOnDemandInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponse) SetRequestId(v string) *DescribeOnDemandInstanceStatusResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponse) SetInstances(v []*DescribeOnDemandInstanceStatusResponseInstances) *DescribeOnDemandInstanceStatusResponse {
	s.Instances = v
	return s
}

type DescribeOnDemandInstanceStatusResponseInstances struct {
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	Net        *string `json:"Net,omitempty" xml:"Net,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	RegistedAs *string `json:"RegistedAs,omitempty" xml:"RegistedAs,omitempty" require:"true"`
	Desc       *string `json:"Desc,omitempty" xml:"Desc,omitempty" require:"true"`
	Declared   *string `json:"Declared,omitempty" xml:"Declared,omitempty" require:"true"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
}

func (s DescribeOnDemandInstanceStatusResponseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseInstances) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseInstances) SetUserId(v string) *DescribeOnDemandInstanceStatusResponseInstances {
	s.UserId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseInstances) SetNet(v string) *DescribeOnDemandInstanceStatusResponseInstances {
	s.Net = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseInstances) SetInstanceId(v string) *DescribeOnDemandInstanceStatusResponseInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseInstances) SetRegistedAs(v string) *DescribeOnDemandInstanceStatusResponseInstances {
	s.RegistedAs = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseInstances) SetDesc(v string) *DescribeOnDemandInstanceStatusResponseInstances {
	s.Desc = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseInstances) SetDeclared(v string) *DescribeOnDemandInstanceStatusResponseInstances {
	s.Declared = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseInstances) SetMode(v string) *DescribeOnDemandInstanceStatusResponseInstances {
	s.Mode = &v
	return s
}

type SetInstanceModeOnDemandRequest struct {
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" require:"true" type:"Repeated"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetInstanceModeOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandRequest) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandRequest) SetInstanceIdList(v []*string) *SetInstanceModeOnDemandRequest {
	s.InstanceIdList = v
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

type SetInstanceModeOnDemandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetInstanceModeOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s SetInstanceModeOnDemandResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceModeOnDemandResponse) SetRequestId(v string) *SetInstanceModeOnDemandResponse {
	s.RequestId = &v
	return s
}

type QuerySchedruleOnDemandRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QuerySchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandRequest) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandRequest) SetInstanceId(v string) *QuerySchedruleOnDemandRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandRequest) SetRegionId(v string) *QuerySchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

type QuerySchedruleOnDemandResponse struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserId     *string                                     `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	InstanceId *string                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	RuleConfig []*QuerySchedruleOnDemandResponseRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" require:"true" type:"Repeated"`
	RuleStatus []*QuerySchedruleOnDemandResponseRuleStatus `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty" require:"true" type:"Repeated"`
}

func (s QuerySchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponse) SetRequestId(v string) *QuerySchedruleOnDemandResponse {
	s.RequestId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetUserId(v string) *QuerySchedruleOnDemandResponse {
	s.UserId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetInstanceId(v string) *QuerySchedruleOnDemandResponse {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetRuleConfig(v []*QuerySchedruleOnDemandResponseRuleConfig) *QuerySchedruleOnDemandResponse {
	s.RuleConfig = v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetRuleStatus(v []*QuerySchedruleOnDemandResponseRuleStatus) *QuerySchedruleOnDemandResponse {
	s.RuleStatus = v
	return s
}

type QuerySchedruleOnDemandResponseRuleConfig struct {
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty" require:"true"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty" require:"true"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty" require:"true"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty" require:"true"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty" require:"true"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty" require:"true"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty" require:"true"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty" require:"true"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty" require:"true"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty" require:"true"`
}

func (s QuerySchedruleOnDemandResponseRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseRuleConfig) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleName(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleName = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleConditionCnt(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleConditionCnt = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleUndoBeginTime(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleUndoMode(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleUndoMode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleUndoEndTime(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleUndoEndTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleConditionMbps(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleConditionMbps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleConditionKpps(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleConditionKpps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleSwitch(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleSwitch = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetRuleAction(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.RuleAction = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleConfig) SetTimeZone(v string) *QuerySchedruleOnDemandResponseRuleConfig {
	s.TimeZone = &v
	return s
}

type QuerySchedruleOnDemandResponseRuleStatus struct {
	Net             *string `json:"Net,omitempty" xml:"Net,omitempty" require:"true"`
	RuleSchedStatus *string `json:"RuleSchedStatus,omitempty" xml:"RuleSchedStatus,omitempty" require:"true"`
}

func (s QuerySchedruleOnDemandResponseRuleStatus) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseRuleStatus) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseRuleStatus) SetNet(v string) *QuerySchedruleOnDemandResponseRuleStatus {
	s.Net = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseRuleStatus) SetRuleSchedStatus(v string) *QuerySchedruleOnDemandResponseRuleStatus {
	s.RuleSchedStatus = &v
	return s
}

type DeleteSchedruleOnDemandRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty" require:"true"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandRequest) GoString() string {
	return s.String()
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

type DeleteSchedruleOnDemandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedruleOnDemandResponse) SetRequestId(v string) *DeleteSchedruleOnDemandResponse {
	s.RequestId = &v
	return s
}

type ConfigSchedruleOnDemandRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty" require:"true"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty" require:"true"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty" require:"true"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty" require:"true"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty" require:"true"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty" require:"true"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty" require:"true"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty" require:"true"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty" require:"true"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandRequest) GoString() string {
	return s.String()
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

type ConfigSchedruleOnDemandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ConfigSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *ConfigSchedruleOnDemandResponse) SetRequestId(v string) *ConfigSchedruleOnDemandResponse {
	s.RequestId = &v
	return s
}

type CreateSchedruleOnDemandRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty" require:"true"`
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty" require:"true"`
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty" require:"true"`
	RuleConditionCnt  *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty" require:"true"`
	RuleAction        *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty" require:"true"`
	RuleSwitch        *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty" require:"true"`
	RuleUndoMode      *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty" require:"true"`
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty" require:"true"`
	RuleUndoEndTime   *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	TimeZone          *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty" require:"true"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateSchedruleOnDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandRequest) GoString() string {
	return s.String()
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

type CreateSchedruleOnDemandResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateSchedruleOnDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchedruleOnDemandResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedruleOnDemandResponse) SetRequestId(v string) *CreateSchedruleOnDemandResponse {
	s.RequestId = &v
	return s
}

type GetSlsOpenStatusRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetSlsOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusRequest) SetRegionId(v string) *GetSlsOpenStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetSlsOpenStatusRequest) SetResourceGroupId(v string) *GetSlsOpenStatusRequest {
	s.ResourceGroupId = &v
	return s
}

type GetSlsOpenStatusResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SlsOpenStatus *bool   `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty" require:"true"`
}

func (s GetSlsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponse) SetRequestId(v string) *GetSlsOpenStatusResponse {
	s.RequestId = &v
	return s
}

func (s *GetSlsOpenStatusResponse) SetSlsOpenStatus(v bool) *GetSlsOpenStatusResponse {
	s.SlsOpenStatus = &v
	return s
}

type CheckAccessLogAuthRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckAccessLogAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthRequest) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthRequest) SetRegionId(v string) *CheckAccessLogAuthRequest {
	s.RegionId = &v
	return s
}

func (s *CheckAccessLogAuthRequest) SetResourceGroupId(v string) *CheckAccessLogAuthRequest {
	s.ResourceGroupId = &v
	return s
}

type CheckAccessLogAuthResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AccessLogAuth *bool   `json:"AccessLogAuth,omitempty" xml:"AccessLogAuth,omitempty" require:"true"`
}

func (s CheckAccessLogAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthResponse) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponse) SetRequestId(v string) *CheckAccessLogAuthResponse {
	s.RequestId = &v
	return s
}

func (s *CheckAccessLogAuthResponse) SetAccessLogAuth(v bool) *CheckAccessLogAuthResponse {
	s.AccessLogAuth = &v
	return s
}

type ListOpenedAccessLogInstancesRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	PageNumber      *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOpenedAccessLogInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesRequest) SetResourceGroupId(v string) *ListOpenedAccessLogInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageNumber(v int) *ListOpenedAccessLogInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageSize(v int) *ListOpenedAccessLogInstancesRequest {
	s.PageSize = &v
	return s
}

type ListOpenedAccessLogInstancesResponse struct {
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount      *int                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	SlsConfigStatus []*ListOpenedAccessLogInstancesResponseSlsConfigStatus `json:"SlsConfigStatus,omitempty" xml:"SlsConfigStatus,omitempty" require:"true" type:"Repeated"`
}

func (s ListOpenedAccessLogInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponse) SetRequestId(v string) *ListOpenedAccessLogInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetTotalCount(v int) *ListOpenedAccessLogInstancesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetSlsConfigStatus(v []*ListOpenedAccessLogInstancesResponseSlsConfigStatus) *ListOpenedAccessLogInstancesResponse {
	s.SlsConfigStatus = v
	return s
}

type ListOpenedAccessLogInstancesResponseSlsConfigStatus struct {
	Enable     *bool   `json:"Enable,omitempty" xml:"Enable,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
}

func (s ListOpenedAccessLogInstancesResponseSlsConfigStatus) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseSlsConfigStatus) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseSlsConfigStatus) SetEnable(v bool) *ListOpenedAccessLogInstancesResponseSlsConfigStatus {
	s.Enable = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseSlsConfigStatus) SetInstanceId(v string) *ListOpenedAccessLogInstancesResponseSlsConfigStatus {
	s.InstanceId = &v
	return s
}

type DescribeOnDemandDdosEventRequest struct {
	StartTime       *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNo          *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
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

func (s *DescribeOnDemandDdosEventRequest) SetStartTime(v int) *DescribeOnDemandDdosEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetEndTime(v int) *DescribeOnDemandDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageSize(v int) *DescribeOnDemandDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageNo(v int) *DescribeOnDemandDdosEventRequest {
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

type DescribeOnDemandDdosEventResponse struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                                     `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Events    []*DescribeOnDemandDdosEventResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOnDemandDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponse) SetRequestId(v string) *DescribeOnDemandDdosEventResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetTotal(v int64) *DescribeOnDemandDdosEventResponse {
	s.Total = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetEvents(v []*DescribeOnDemandDdosEventResponseEvents) *DescribeOnDemandDdosEventResponse {
	s.Events = v
	return s
}

type DescribeOnDemandDdosEventResponseEvents struct {
	StartTime *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Pps       *int    `json:"Pps,omitempty" xml:"Pps,omitempty" require:"true"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	Mbps      *int    `json:"Mbps,omitempty" xml:"Mbps,omitempty" require:"true"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeOnDemandDdosEventResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseEvents) SetStartTime(v int) *DescribeOnDemandDdosEventResponseEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseEvents) SetEndTime(v int) *DescribeOnDemandDdosEventResponseEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseEvents) SetPps(v int) *DescribeOnDemandDdosEventResponseEvents {
	s.Pps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseEvents) SetIp(v string) *DescribeOnDemandDdosEventResponseEvents {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseEvents) SetMbps(v int) *DescribeOnDemandDdosEventResponseEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseEvents) SetStatus(v string) *DescribeOnDemandDdosEventResponseEvents {
	s.Status = &v
	return s
}

type ListTagKeysRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetCurrentPage(v int) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

type ListTagKeysResponse struct {
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	CurrentPage *int                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int                          `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount  *int                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	TagKeys     []*ListTagKeysResponseTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetRequestId(v string) *ListTagKeysResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponse) SetCurrentPage(v int) *ListTagKeysResponse {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponse) SetPageSize(v int) *ListTagKeysResponse {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponse) SetTotalCount(v int) *ListTagKeysResponse {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponse) SetTagKeys(v []*ListTagKeysResponseTagKeys) *ListTagKeysResponse {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseTagKeys struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagCount *int    `json:"TagCount,omitempty" xml:"TagCount,omitempty" require:"true"`
}

func (s ListTagKeysResponseTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseTagKeys) SetTagKey(v string) *ListTagKeysResponseTagKeys {
	s.TagKey = &v
	return s
}

func (s *ListTagKeysResponseTagKeys) SetTagCount(v int) *ListTagKeysResponseTagKeys {
	s.TagCount = &v
	return s
}

type ListTagResourcesRequest struct {
	ResourceGroupId *string                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType    *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId      []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag             []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	NextToken       *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
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

type ListTagResourcesResponse struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken    *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	TagResources *ListTagResourcesResponseTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" require:"true" type:"Struct"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetRequestId(v string) *ListTagResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponse) SetNextToken(v string) *ListTagResourcesResponse {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponse) SetTagResources(v *ListTagResourcesResponseTagResources) *ListTagResourcesResponse {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseTagResources struct {
	TagResource []*ListTagResourcesResponseTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResources) SetTagResource(v []*ListTagResourcesResponseTagResourcesTagResource) *ListTagResourcesResponseTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s ListTagResourcesResponseTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type UntagResourcesRequest struct {
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType    *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId      []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	TagKey          []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All             *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
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

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetRequestId(v string) *UntagResourcesResponse {
	s.RequestId = &v
	return s
}

type TagResourcesRequest struct {
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType    *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId      []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	Tag             []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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

type TagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetRequestId(v string) *TagResourcesResponse {
	s.RequestId = &v
	return s
}

type DescribeExcpetionCountRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExcpetionCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountRequest) SetResourceGroupId(v string) *DescribeExcpetionCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExcpetionCountRequest) SetRegionId(v string) *DescribeExcpetionCountRequest {
	s.RegionId = &v
	return s
}

type DescribeExcpetionCountResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ExceptionIpCount *int    `json:"ExceptionIpCount,omitempty" xml:"ExceptionIpCount,omitempty" require:"true"`
	ExpireTimeCount  *int    `json:"ExpireTimeCount,omitempty" xml:"ExpireTimeCount,omitempty" require:"true"`
}

func (s DescribeExcpetionCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponse) SetRequestId(v string) *DescribeExcpetionCountResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeExcpetionCountResponse) SetExceptionIpCount(v int) *DescribeExcpetionCountResponse {
	s.ExceptionIpCount = &v
	return s
}

func (s *DescribeExcpetionCountResponse) SetExpireTimeCount(v int) *DescribeExcpetionCountResponse {
	s.ExpireTimeCount = &v
	return s
}

type DescribePackIpListRequest struct {
	PageNo          *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
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

func (s *DescribePackIpListRequest) SetPageNo(v int) *DescribePackIpListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePackIpListRequest) SetPageSize(v int) *DescribePackIpListRequest {
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

type DescribePackIpListResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Total     *int                                `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	IpList    []*DescribePackIpListResponseIpList `json:"IpList,omitempty" xml:"IpList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePackIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponse) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponse) SetRequestId(v string) *DescribePackIpListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePackIpListResponse) SetSuccess(v bool) *DescribePackIpListResponse {
	s.Success = &v
	return s
}

func (s *DescribePackIpListResponse) SetCode(v string) *DescribePackIpListResponse {
	s.Code = &v
	return s
}

func (s *DescribePackIpListResponse) SetTotal(v int) *DescribePackIpListResponse {
	s.Total = &v
	return s
}

func (s *DescribePackIpListResponse) SetIpList(v []*DescribePackIpListResponseIpList) *DescribePackIpListResponse {
	s.IpList = v
	return s
}

type DescribePackIpListResponseIpList struct {
	Ip      *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty" require:"true"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty" require:"true"`
}

func (s DescribePackIpListResponseIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseIpList) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseIpList) SetIp(v string) *DescribePackIpListResponseIpList {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListResponseIpList) SetProduct(v string) *DescribePackIpListResponseIpList {
	s.Product = &v
	return s
}

func (s *DescribePackIpListResponseIpList) SetStatus(v string) *DescribePackIpListResponseIpList {
	s.Status = &v
	return s
}

func (s *DescribePackIpListResponseIpList) SetRemark(v string) *DescribePackIpListResponseIpList {
	s.Remark = &v
	return s
}

type DescribeRegionsRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetResourceGroupId(v string) *DescribeRegionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Regions   []*DescribeRegionsResponseRegions `json:"Regions,omitempty" xml:"Regions,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetRequestId(v string) *DescribeRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponse) SetSuccess(v bool) *DescribeRegionsResponse {
	s.Success = &v
	return s
}

func (s *DescribeRegionsResponse) SetCode(v string) *DescribeRegionsResponse {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponse) SetRegions(v []*DescribeRegionsResponseRegions) *DescribeRegionsResponse {
	s.Regions = v
	return s
}

type DescribeRegionsResponseRegions struct {
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty" require:"true"`
	RegionName   *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeRegionsResponseRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegions) SetRegionEnName(v string) *DescribeRegionsResponseRegions {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseRegions) SetRegionName(v string) *DescribeRegionsResponseRegions {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseRegions) SetRegionId(v string) *DescribeRegionsResponseRegions {
	s.RegionId = &v
	return s
}

type ModifyRemarkRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkRequest) GoString() string {
	return s.String()
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

type ModifyRemarkResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyRemarkResponse) SetRequestId(v string) *ModifyRemarkResponse {
	s.RequestId = &v
	return s
}

type DescribeTrafficRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ipnet           *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	StartTime       *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval        *int    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	FlowType        *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
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

func (s *DescribeTrafficRequest) SetStartTime(v int) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetEndTime(v int) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeTrafficRequest) SetFlowType(v string) *DescribeTrafficRequest {
	s.FlowType = &v
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

type DescribeTrafficResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowList  []*DescribeTrafficResponseFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponse) SetRequestId(v string) *DescribeTrafficResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficResponse) SetFlowList(v []*DescribeTrafficResponseFlowList) *DescribeTrafficResponse {
	s.FlowList = v
	return s
}

type DescribeTrafficResponseFlowList struct {
	Pps       *int    `json:"Pps,omitempty" xml:"Pps,omitempty" require:"true"`
	FlowType  *string `json:"FlowType,omitempty" xml:"FlowType,omitempty" require:"true"`
	Kbps      *int    `json:"Kbps,omitempty" xml:"Kbps,omitempty" require:"true"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	AttackBps *int64  `json:"AttackBps,omitempty" xml:"AttackBps,omitempty" require:"true"`
	AttackPps *int64  `json:"AttackPps,omitempty" xml:"AttackPps,omitempty" require:"true"`
	Time      *int    `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeTrafficResponseFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseFlowList) SetPps(v int) *DescribeTrafficResponseFlowList {
	s.Pps = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetFlowType(v string) *DescribeTrafficResponseFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetKbps(v int) *DescribeTrafficResponseFlowList {
	s.Kbps = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetName(v string) *DescribeTrafficResponseFlowList {
	s.Name = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetAttackBps(v int64) *DescribeTrafficResponseFlowList {
	s.AttackBps = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetAttackPps(v int64) *DescribeTrafficResponseFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetTime(v int) *DescribeTrafficResponseFlowList {
	s.Time = &v
	return s
}

type DescribeResourcePackUsageRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
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

type DescribeResourcePackUsageResponse struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Interval   *int64                                         `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	StartTime  *int64                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *int64                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PackUsages []*DescribeResourcePackUsageResponsePackUsages `json:"PackUsages,omitempty" xml:"PackUsages,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponse) SetRequestId(v string) *DescribeResourcePackUsageResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetInterval(v int64) *DescribeResourcePackUsageResponse {
	s.Interval = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetStartTime(v int64) *DescribeResourcePackUsageResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetEndTime(v int64) *DescribeResourcePackUsageResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetPackUsages(v []*DescribeResourcePackUsageResponsePackUsages) *DescribeResourcePackUsageResponse {
	s.PackUsages = v
	return s
}

type DescribeResourcePackUsageResponsePackUsages struct {
	Traffic *float32 `json:"Traffic,omitempty" xml:"Traffic,omitempty" require:"true"`
	Time    *int64   `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeResourcePackUsageResponsePackUsages) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponsePackUsages) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponsePackUsages) SetTraffic(v float32) *DescribeResourcePackUsageResponsePackUsages {
	s.Traffic = &v
	return s
}

func (s *DescribeResourcePackUsageResponsePackUsages) SetTime(v int64) *DescribeResourcePackUsageResponsePackUsages {
	s.Time = &v
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

type DescribeResourcePackStatisticsResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AvailablePackNum  *int    `json:"AvailablePackNum,omitempty" xml:"AvailablePackNum,omitempty" require:"true"`
	TotalCurrCapacity *int64  `json:"TotalCurrCapacity,omitempty" xml:"TotalCurrCapacity,omitempty" require:"true"`
	TotalUsedCapacity *int64  `json:"TotalUsedCapacity,omitempty" xml:"TotalUsedCapacity,omitempty" require:"true"`
	TotalInitCapacity *int64  `json:"TotalInitCapacity,omitempty" xml:"TotalInitCapacity,omitempty" require:"true"`
}

func (s DescribeResourcePackStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsResponse) SetRequestId(v string) *DescribeResourcePackStatisticsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetAvailablePackNum(v int) *DescribeResourcePackStatisticsResponse {
	s.AvailablePackNum = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetTotalCurrCapacity(v int64) *DescribeResourcePackStatisticsResponse {
	s.TotalCurrCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetTotalUsedCapacity(v int64) *DescribeResourcePackStatisticsResponse {
	s.TotalUsedCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetTotalInitCapacity(v int64) *DescribeResourcePackStatisticsResponse {
	s.TotalInitCapacity = &v
	return s
}

type DescribeResourcePackInstancesRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
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

func (s *DescribeResourcePackInstancesRequest) SetPageSize(v int) *DescribeResourcePackInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetCurrentPage(v int) *DescribeResourcePackInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetResourceGroupId(v string) *DescribeResourcePackInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeResourcePackInstancesResponse struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount    *int                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	ResourcePacks []*DescribeResourcePackInstancesResponseResourcePacks `json:"ResourcePacks,omitempty" xml:"ResourcePacks,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponse) SetRequestId(v string) *DescribeResourcePackInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponse) SetTotalCount(v int) *DescribeResourcePackInstancesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourcePackInstancesResponse) SetResourcePacks(v []*DescribeResourcePackInstancesResponseResourcePacks) *DescribeResourcePackInstancesResponse {
	s.ResourcePacks = v
	return s
}

type DescribeResourcePackInstancesResponseResourcePacks struct {
	ResourcePackId *string `json:"ResourcePackId,omitempty" xml:"ResourcePackId,omitempty" require:"true"`
	InitCapacity   *int64  `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty" require:"true"`
	CurrCapacity   *int64  `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty" require:"true"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeResourcePackInstancesResponseResourcePacks) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponseResourcePacks) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetResourcePackId(v string) *DescribeResourcePackInstancesResponseResourcePacks {
	s.ResourcePackId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetInitCapacity(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.InitCapacity = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetCurrCapacity(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetStartTime(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetEndTime(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetStatus(v string) *DescribeResourcePackInstancesResponseResourcePacks {
	s.Status = &v
	return s
}

type DescribePackPaidTrafficRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
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

func (s *DescribePackPaidTrafficRequest) SetCurrentPage(v int) *DescribePackPaidTrafficRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetPageSize(v int) *DescribePackPaidTrafficRequest {
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

type DescribePackPaidTrafficResponse struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount       *int                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PackPaidTraffics []*DescribePackPaidTrafficResponsePackPaidTraffics `json:"PackPaidTraffics,omitempty" xml:"PackPaidTraffics,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePackPaidTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponse) SetRequestId(v string) *DescribePackPaidTrafficResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePackPaidTrafficResponse) SetTotalCount(v int) *DescribePackPaidTrafficResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribePackPaidTrafficResponse) SetPackPaidTraffics(v []*DescribePackPaidTrafficResponsePackPaidTraffics) *DescribePackPaidTrafficResponse {
	s.PackPaidTraffics = v
	return s
}

type DescribePackPaidTrafficResponsePackPaidTraffics struct {
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StartTime        *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	BaseBandwidth    *int     `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty" require:"true"`
	ElasticBandwidth *int     `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty" require:"true"`
	PaidCapacity     *float32 `json:"PaidCapacity,omitempty" xml:"PaidCapacity,omitempty" require:"true"`
	TotalCapacity    *float32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty" require:"true"`
	MaxAttack        *float32 `json:"MaxAttack,omitempty" xml:"MaxAttack,omitempty" require:"true"`
}

func (s DescribePackPaidTrafficResponsePackPaidTraffics) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponsePackPaidTraffics) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetInstanceId(v string) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.InstanceId = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetStartTime(v int64) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.StartTime = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetBaseBandwidth(v int) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetElasticBandwidth(v int) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetPaidCapacity(v float32) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.PaidCapacity = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetTotalCapacity(v float32) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.TotalCapacity = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetMaxAttack(v float32) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.MaxAttack = &v
	return s
}

type DescribeOpEntitiesRequest struct {
	CurrentPage     *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
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

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int) *DescribeOpEntitiesRequest {
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

type DescribeOpEntitiesResponse struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	OpEntities []*DescribeOpEntitiesResponseOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetRequestId(v string) *DescribeOpEntitiesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetTotalCount(v int) *DescribeOpEntitiesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetOpEntities(v []*DescribeOpEntitiesResponseOpEntities) *DescribeOpEntitiesResponse {
	s.OpEntities = v
	return s
}

type DescribeOpEntitiesResponseOpEntities struct {
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty" require:"true"`
	EntityType   *int    `json:"EntityType,omitempty" xml:"EntityType,omitempty" require:"true"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty" require:"true"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty" require:"true"`
	OpAction     *int    `json:"OpAction,omitempty" xml:"OpAction,omitempty" require:"true"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponseOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetEntityType(v int) *DescribeOpEntitiesResponseOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseOpEntities {
	s.OpDesc = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpAction(v int) *DescribeOpEntitiesResponseOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseOpEntities {
	s.GmtCreate = &v
	return s
}

type DescribeInstanceSpecsRequest struct {
	InstanceIdList  *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
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

type DescribeInstanceSpecsResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceSpecs []*DescribeInstanceSpecsResponseInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) SetRequestId(v string) *DescribeInstanceSpecsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseInstanceSpecs) *DescribeInstanceSpecsResponse {
	s.InstanceSpecs = v
	return s
}

type DescribeInstanceSpecsResponseInstanceSpecs struct {
	Region                        *string                                               `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	AvailableDeleteBlackholeCount *string                                               `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty" require:"true"`
	InstanceId                    *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	AvailableDefenseTimes         *int                                                  `json:"AvailableDefenseTimes,omitempty" xml:"AvailableDefenseTimes,omitempty" require:"true"`
	TotalDefenseTimes             *int                                                  `json:"TotalDefenseTimes,omitempty" xml:"TotalDefenseTimes,omitempty" require:"true"`
	IsFullDefenseMode             *int                                                  `json:"IsFullDefenseMode,omitempty" xml:"IsFullDefenseMode,omitempty" require:"true"`
	PackConfig                    *DescribeInstanceSpecsResponseInstanceSpecsPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribeInstanceSpecsResponseInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetRegion(v string) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.Region = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetAvailableDeleteBlackholeCount(v string) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetAvailableDefenseTimes(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.AvailableDefenseTimes = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetTotalDefenseTimes(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.TotalDefenseTimes = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetIsFullDefenseMode(v int) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.IsFullDefenseMode = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecs) SetPackConfig(v *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) *DescribeInstanceSpecsResponseInstanceSpecs {
	s.PackConfig = v
	return s
}

type DescribeInstanceSpecsResponseInstanceSpecsPackConfig struct {
	PackAdvThre     *int `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty" require:"true"`
	IpAdvanceThre   *int `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty" require:"true"`
	IpBasicThre     *int `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty" require:"true"`
	PackBasicThre   *int `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty" require:"true"`
	IpSpec          *int `json:"IpSpec,omitempty" xml:"IpSpec,omitempty" require:"true"`
	BindIpCount     *int `json:"BindIpCount,omitempty" xml:"BindIpCount,omitempty" require:"true"`
	NormalBandwidth *int `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty" require:"true"`
}

func (s DescribeInstanceSpecsResponseInstanceSpecsPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseInstanceSpecsPackConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) SetPackAdvThre(v int) *DescribeInstanceSpecsResponseInstanceSpecsPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) SetIpAdvanceThre(v int) *DescribeInstanceSpecsResponseInstanceSpecsPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) SetIpBasicThre(v int) *DescribeInstanceSpecsResponseInstanceSpecsPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) SetPackBasicThre(v int) *DescribeInstanceSpecsResponseInstanceSpecsPackConfig {
	s.PackBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) SetIpSpec(v int) *DescribeInstanceSpecsResponseInstanceSpecsPackConfig {
	s.IpSpec = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) SetBindIpCount(v int) *DescribeInstanceSpecsResponseInstanceSpecsPackConfig {
	s.BindIpCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseInstanceSpecsPackConfig) SetNormalBandwidth(v int) *DescribeInstanceSpecsResponseInstanceSpecsPackConfig {
	s.NormalBandwidth = &v
	return s
}

type DescribeInstanceListRequest struct {
	ResourceGroupId *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceIdList  *string                           `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	Remark          *string                           `json:"Remark,omitempty" xml:"Remark,omitempty"`
	PageNo          *int                              `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize        *int                              `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
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

func (s *DescribeInstanceListRequest) SetPageNo(v int) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int) *DescribeInstanceListRequest {
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

type DescribeInstanceListResponse struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total        *int64                                      `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	InstanceList []*DescribeInstanceListResponseInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponse) SetRequestId(v string) *DescribeInstanceListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponse) SetTotal(v int64) *DescribeInstanceListResponse {
	s.Total = &v
	return s
}

func (s *DescribeInstanceListResponse) SetInstanceList(v []*DescribeInstanceListResponseInstanceList) *DescribeInstanceListResponse {
	s.InstanceList = v
	return s
}

type DescribeInstanceListResponseInstanceList struct {
	ExpireTime        *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Product           *string `json:"Product,omitempty" xml:"Product,omitempty" require:"true"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty" require:"true"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	IpType            *string `json:"IpType,omitempty" xml:"IpType,omitempty" require:"true"`
	AutoRenewal       *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty" require:"true"`
	BlackholdingCount *string `json:"BlackholdingCount,omitempty" xml:"BlackholdingCount,omitempty" require:"true"`
	GmtCreate         *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	InstanceType      *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" require:"true"`
}

func (s DescribeInstanceListResponseInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetInstanceId(v string) *DescribeInstanceListResponseInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetProduct(v string) *DescribeInstanceListResponseInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetRemark(v string) *DescribeInstanceListResponseInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetStatus(v string) *DescribeInstanceListResponseInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetIpType(v string) *DescribeInstanceListResponseInstanceList {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetAutoRenewal(v bool) *DescribeInstanceListResponseInstanceList {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetBlackholdingCount(v string) *DescribeInstanceListResponseInstanceList {
	s.BlackholdingCount = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetGmtCreate(v int64) *DescribeInstanceListResponseInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetInstanceType(v string) *DescribeInstanceListResponseInstanceList {
	s.InstanceType = &v
	return s
}

type DescribeDdosEventRequest struct {
	StartTime       *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNo          *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
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

func (s *DescribeDdosEventRequest) SetStartTime(v int) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetEndTime(v int) *DescribeDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageNo(v int) *DescribeDdosEventRequest {
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

type DescribeDdosEventResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                             `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Events    []*DescribeDdosEventResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponse) SetRequestId(v string) *DescribeDdosEventResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponse) SetTotal(v int64) *DescribeDdosEventResponse {
	s.Total = &v
	return s
}

func (s *DescribeDdosEventResponse) SetEvents(v []*DescribeDdosEventResponseEvents) *DescribeDdosEventResponse {
	s.Events = v
	return s
}

type DescribeDdosEventResponseEvents struct {
	StartTime *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Pps       *int    `json:"Pps,omitempty" xml:"Pps,omitempty" require:"true"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	Mbps      *int    `json:"Mbps,omitempty" xml:"Mbps,omitempty" require:"true"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeDdosEventResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseEvents) SetStartTime(v int) *DescribeDdosEventResponseEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetEndTime(v int) *DescribeDdosEventResponseEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetPps(v int) *DescribeDdosEventResponseEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetIp(v string) *DescribeDdosEventResponseEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetMbps(v int) *DescribeDdosEventResponseEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetStatus(v string) *DescribeDdosEventResponseEvents {
	s.Status = &v
	return s
}

type DeleteIpRequest struct {
	IpList          *string `json:"IpList,omitempty" xml:"IpList,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
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

type DeleteIpResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpResponse) SetRequestId(v string) *DeleteIpResponse {
	s.RequestId = &v
	return s
}

type DeleteBlackholeRequest struct {
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteBlackholeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
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

type DeleteBlackholeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteBlackholeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponse) SetRequestId(v string) *DeleteBlackholeResponse {
	s.RequestId = &v
	return s
}

type CheckGrantRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) SetResourceGroupId(v string) *CheckGrantRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckGrantRequest) SetRegionId(v string) *CheckGrantRequest {
	s.RegionId = &v
	return s
}

type CheckGrantResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Status    *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s CheckGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponse) GoString() string {
	return s.String()
}

func (s *CheckGrantResponse) SetRequestId(v string) *CheckGrantResponse {
	s.RequestId = &v
	return s
}

func (s *CheckGrantResponse) SetStatus(v int) *CheckGrantResponse {
	s.Status = &v
	return s
}

type AddIpRequest struct {
	IpList          *string `json:"IpList,omitempty" xml:"IpList,omitempty" require:"true"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AddIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
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

type AddIpResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponse) GoString() string {
	return s.String()
}

func (s *AddIpResponse) SetRequestId(v string) *AddIpResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) DescribeOnDemandInstanceStatusWithOptions(request *DescribeOnDemandInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOnDemandInstanceStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SetInstanceModeOnDemandWithOptions(request *SetInstanceModeOnDemandRequest, runtime *util.RuntimeOptions) (_result *SetInstanceModeOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.DoRequest(tea.String("SetInstanceModeOnDemand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) QuerySchedruleOnDemandWithOptions(request *QuerySchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *QuerySchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.DoRequest(tea.String("QuerySchedruleOnDemand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteSchedruleOnDemandWithOptions(request *DeleteSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *DeleteSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteSchedruleOnDemand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ConfigSchedruleOnDemandWithOptions(request *ConfigSchedruleOnDemandRequest, runtime *util.RuntimeOptions) (_result *ConfigSchedruleOnDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigSchedruleOnDemand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.DoRequest(tea.String("CreateSchedruleOnDemand"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetSlsOpenStatusWithOptions(request *GetSlsOpenStatusRequest, runtime *util.RuntimeOptions) (_result *GetSlsOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.DoRequest(tea.String("GetSlsOpenStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSlsOpenStatus(request *GetSlsOpenStatusRequest) (_result *GetSlsOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.GetSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAccessLogAuthWithOptions(request *CheckAccessLogAuthRequest, runtime *util.RuntimeOptions) (_result *CheckAccessLogAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.DoRequest(tea.String("CheckAccessLogAuth"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAccessLogAuth(request *CheckAccessLogAuthRequest) (_result *CheckAccessLogAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CheckAccessLogAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOpenedAccessLogInstancesWithOptions(request *ListOpenedAccessLogInstancesRequest, runtime *util.RuntimeOptions) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("ListOpenedAccessLogInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOpenedAccessLogInstances(request *ListOpenedAccessLogInstancesRequest) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.ListOpenedAccessLogInstancesWithOptions(request, runtime)
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
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOnDemandDdosEvent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagKeys"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeExcpetionCountWithOptions(request *DescribeExcpetionCountRequest, runtime *util.RuntimeOptions) (_result *DescribeExcpetionCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeExcpetionCount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribePackIpListWithOptions(request *DescribePackIpListRequest, runtime *util.RuntimeOptions) (_result *DescribePackIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePackIpList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ModifyRemarkWithOptions(request *ModifyRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyRemark"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeTraffic"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeResourcePackUsageWithOptions(request *DescribeResourcePackUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeResourcePackUsageResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeResourcePackUsage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeResourcePackStatisticsWithOptions(request *DescribeResourcePackStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeResourcePackStatisticsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeResourcePackStatistics"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeResourcePackInstancesWithOptions(request *DescribeResourcePackInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeResourcePackInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeResourcePackInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribePackPaidTrafficWithOptions(request *DescribePackPaidTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribePackPaidTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePackPaidTrafficResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePackPaidTraffic"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOpEntities"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceSpecs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDdosEvent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *util.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteIp"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *util.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteBlackhole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *util.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.DoRequest(tea.String("CheckGrant"), tea.String("HTTPS"), tea.String("GET"), tea.String("2018-07-20"), tea.String("AK"), tea.ToMap(request), nil, runtime)
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

func (client *Client) AddIpWithOptions(request *AddIpRequest, runtime *util.RuntimeOptions) (_result *AddIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddIpResponse{}
	_body, _err := client.DoRequest(tea.String("AddIp"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-07-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
