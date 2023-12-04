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

type AddIpRequest struct {
	// The ID of the Anti-DDoS Origin Enterprise instance.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin Enterprise instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IP addresses that you want to add to the Anti-DDoS Origin Enterprise instance. This parameter is a string consisting of JSON arrays. Each element in a JSON array is a JSON struct that includes the following field:
	//
	// *   **ip**: required. The IP address that you want to add. Data type: string.
	//
	//     > The IP address must be the IP address of an asset that belongs to the current Alibaba Cloud account.
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region ID of the Anti-DDoS Origin Enterprise instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin Enterprise instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](~~94485~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AddIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
}

func (s *AddIpRequest) SetInstanceId(v string) *AddIpRequest {
	s.InstanceId = &v
	return s
}

func (s *AddIpRequest) SetIpList(v string) *AddIpRequest {
	s.IpList = &v
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
	// The ID of the request.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddIpResponse) SetStatusCode(v int32) *AddIpResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpResponse) SetBody(v *AddIpResponseBody) *AddIpResponse {
	s.Body = v
	return s
}

type AttachAssetGroupToInstanceRequest struct {
	AssetGroupList []*AttachAssetGroupToInstanceRequestAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
	InstanceId     *string                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceIp       *string                                            `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s AttachAssetGroupToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceRequest) SetAssetGroupList(v []*AttachAssetGroupToInstanceRequestAssetGroupList) *AttachAssetGroupToInstanceRequest {
	s.AssetGroupList = v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) SetInstanceId(v string) *AttachAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) SetRegionId(v string) *AttachAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequest) SetSourceIp(v string) *AttachAssetGroupToInstanceRequest {
	s.SourceIp = &v
	return s
}

type AttachAssetGroupToInstanceRequestAssetGroupList struct {
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AttachAssetGroupToInstanceRequestAssetGroupList) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceRequestAssetGroupList) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetMemberUid(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.MemberUid = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetName(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Name = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetRegion(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Region = &v
	return s
}

func (s *AttachAssetGroupToInstanceRequestAssetGroupList) SetType(v string) *AttachAssetGroupToInstanceRequestAssetGroupList {
	s.Type = &v
	return s
}

type AttachAssetGroupToInstanceShrinkRequest struct {
	AssetGroupListShrink *string `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceIp             *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s AttachAssetGroupToInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetAssetGroupListShrink(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.AssetGroupListShrink = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetInstanceId(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetRegionId(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AttachAssetGroupToInstanceShrinkRequest) SetSourceIp(v string) *AttachAssetGroupToInstanceShrinkRequest {
	s.SourceIp = &v
	return s
}

type AttachAssetGroupToInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAssetGroupToInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceResponseBody) SetRequestId(v string) *AttachAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

type AttachAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachAssetGroupToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *AttachAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachAssetGroupToInstanceResponse) SetStatusCode(v int32) *AttachAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAssetGroupToInstanceResponse) SetBody(v *AttachAssetGroupToInstanceResponseBody) *AttachAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

type CheckAccessLogAuthRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// For more information about the valid values of this parameter, see [Regions and zones](~~188196~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](~~94485~~).
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

type CheckAccessLogAuthResponseBody struct {
	// Indicates whether Anti-DDoS Origin was authorized to access Log Service. Valid values:
	//
	// *   **true**: Anti-DDoS Origin was authorized.
	// *   **false**: Anti-DDoS Origin was not authorized.
	AccessLogAuth *bool `json:"AccessLogAuth,omitempty" xml:"AccessLogAuth,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccessLogAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponseBody) SetAccessLogAuth(v bool) *CheckAccessLogAuthResponseBody {
	s.AccessLogAuth = &v
	return s
}

func (s *CheckAccessLogAuthResponseBody) SetRequestId(v string) *CheckAccessLogAuthResponseBody {
	s.RequestId = &v
	return s
}

type CheckAccessLogAuthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckAccessLogAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckAccessLogAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccessLogAuthResponse) GoString() string {
	return s.String()
}

func (s *CheckAccessLogAuthResponse) SetHeaders(v map[string]*string) *CheckAccessLogAuthResponse {
	s.Headers = v
	return s
}

func (s *CheckAccessLogAuthResponse) SetStatusCode(v int32) *CheckAccessLogAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccessLogAuthResponse) SetBody(v *CheckAccessLogAuthResponseBody) *CheckAccessLogAuthResponse {
	s.Body = v
	return s
}

type CheckGrantRequest struct {
	// Specifies whether to allow Anti-DDoS Origin to check the service-linked role. Valid values:
	//
	// *   **true**
	// *   **false**
	IsSlr *bool `json:"IsSlr,omitempty" xml:"IsSlr,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) SetIsSlr(v bool) *CheckGrantRequest {
	s.IsSlr = &v
	return s
}

func (s *CheckGrantRequest) SetRegionId(v string) *CheckGrantRequest {
	s.RegionId = &v
	return s
}

func (s *CheckGrantRequest) SetResourceGroupId(v string) *CheckGrantRequest {
	s.ResourceGroupId = &v
	return s
}

type CheckGrantResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account. Valid values:
	//
	// *   **1**: Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
	// *   **0**: Anti-DDoS Origin is not authorized to obtain information about the assets within the current Alibaba Cloud account.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckGrantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponseBody) GoString() string {
	return s.String()
}

func (s *CheckGrantResponseBody) SetRequestId(v string) *CheckGrantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckGrantResponseBody) SetStatus(v int32) *CheckGrantResponseBody {
	s.Status = &v
	return s
}

type CheckGrantResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckGrantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CheckGrantResponse) SetStatusCode(v int32) *CheckGrantResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckGrantResponse) SetBody(v *CheckGrantResponseBody) *CheckGrantResponse {
	s.Body = v
	return s
}

type ConfigSchedruleOnDemandRequest struct {
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling action. Set the value to **declare**, which indicates that the route is advertised.
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// If the inbound bandwidth or packets consecutively exceed the threshold for the specified number of times, the scheduling rule is triggered and traffic is rerouted to the on-demand instance. The specified number of times is the value of this parameter.
	//
	// >  The threshold of inbound bandwidth is the value of **RuleConditionMbps**. The threshold of inbound packets is the value of **RuleConditionKpps**.
	RuleConditionCnt *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	// The threshold of inbound packets. Unit: Kpps. Minimum value: **10**.
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	// The threshold of inbound bandwidth. Unit: Mbit/s. Minimum value: **100**.
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	// The name of the scheduling rule.
	//
	// The name can contain lowercase letters, digits, hyphens (-), and underscores (\_). The name can be up to 32 characters in length. We recommend that you use the ID of the on-demand instance as the name. Example: `ddosbgp-cn-z2q1qzxb****`.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether the scheduling rule is enabled. Valid values:
	//
	// *   **on**: enabled
	// *   **off**: disabled
	RuleSwitch *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The start time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// If the system detects that DDoS attacks stop, the system no longer reroutes traffic to the on-demand instance from the time you specified. We recommend that you set this parameter to a value that is defined as off-peak hours.
	//
	// >  This parameter takes effect only when the **RuleUndoMode** parameter is set to **auto**.
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	// The end time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	RuleUndoEndTime *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	// The stop method of the scheduling rule. Valid values:
	//
	// *   **auto**: The scheduling rule automatically stops.
	// *   **manual**: The scheduling rule is manually stopped.
	RuleUndoMode *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	// The time zone of the time when the scheduling rule automatically stops. The time zone must be in the `GMT-hh:mm` format.
	//
	// For example, the value `GMT-08:00` indicates that the time zone is UTC+8.
	//
	// >  This parameter takes effect only when the **RuleUndoMode** parameter is set to **auto**.
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
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

func (s *ConfigSchedruleOnDemandRequest) SetRegionId(v string) *ConfigSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleAction(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleName(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetRuleSwitch(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleSwitch = &v
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

func (s *ConfigSchedruleOnDemandRequest) SetRuleUndoMode(v string) *ConfigSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *ConfigSchedruleOnDemandRequest) SetTimeZone(v string) *ConfigSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

type ConfigSchedruleOnDemandResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigSchedruleOnDemandResponse) SetStatusCode(v int32) *ConfigSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSchedruleOnDemandResponse) SetBody(v *ConfigSchedruleOnDemandResponseBody) *ConfigSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type CreateSchedruleOnDemandRequest struct {
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling action. Set the value to **declare**, which indicates that the route is advertised.
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// If the inbound bandwidth or packets consecutively exceed the threshold for the specified number of times, the scheduling rule is triggered and traffic is rerouted to the on-demand instance. The specified number of times is the value of this parameter.
	//
	// >  The threshold of inbound bandwidth is the value of **RuleConditionMbps**. The threshold of inbound packets is the value of **RuleConditionKpps**.
	RuleConditionCnt *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	// The threshold of inbound packets. Unit: Kpps. Minimum value: **10**.
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	// The threshold of inbound bandwidth. Unit: Mbit/s. Minimum value: **100**.
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	// The name of the scheduling rule.
	//
	// The name can contain lowercase letters, digits, hyphens (-), and underscores (\_). The name can be up to 32 characters in length. We recommend that you use the ID of the on-demand instance as the name. Example: `ddosbgp-cn-z2q1qzxb****`.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether the scheduling rule is enabled. Valid values:
	//
	// *   **on**: enabled
	// *   **off**: disabled
	RuleSwitch *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The start time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// If the system detects that DDoS attacks stop, the system no longer reroutes traffic to the on-demand instance from the time you specified. We recommend that you set this parameter to a value that is defined as off-peak hours.
	//
	// >  This parameter takes effect only when the **RuleUndoMode** parameter is set to **auto**.
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	// The end time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	RuleUndoEndTime *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	// The stop method of the scheduling rule. Valid values:
	//
	// *   **auto**: The scheduling rule automatically stops.
	// *   **manual**: The scheduling rule is manually stopped.
	RuleUndoMode *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	// The time zone of the time when the scheduling rule automatically stops. The time zone must be in the `GMT-hh:mm` format.
	//
	// For example, the value `GMT-08:00` indicates that the time zone is UTC+8.
	//
	// >  This parameter takes effect only when the **RuleUndoMode** parameter is set to **auto**.
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
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

func (s *CreateSchedruleOnDemandRequest) SetRegionId(v string) *CreateSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleAction(v string) *CreateSchedruleOnDemandRequest {
	s.RuleAction = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionCnt(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionCnt = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionKpps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionKpps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleConditionMbps(v string) *CreateSchedruleOnDemandRequest {
	s.RuleConditionMbps = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleName(v string) *CreateSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetRuleSwitch(v string) *CreateSchedruleOnDemandRequest {
	s.RuleSwitch = &v
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

func (s *CreateSchedruleOnDemandRequest) SetRuleUndoMode(v string) *CreateSchedruleOnDemandRequest {
	s.RuleUndoMode = &v
	return s
}

func (s *CreateSchedruleOnDemandRequest) SetTimeZone(v string) *CreateSchedruleOnDemandRequest {
	s.TimeZone = &v
	return s
}

type CreateSchedruleOnDemandResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSchedruleOnDemandResponse) SetStatusCode(v int32) *CreateSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchedruleOnDemandResponse) SetBody(v *CreateSchedruleOnDemandResponseBody) *CreateSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DeleteBlackholeRequest struct {
	// The ID of the Anti-DDoS Origin instance.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address for which you want to deactivate blackhole filtering.
	//
	// >  You can call the [DescribePackIpList](~~118701~~) operation to query all the IP addresses that are protected by the Anti-DDoS Origin instance and the protection status of the IP addresses. For example, you can query whether blackhole filtering is triggered for an IP address.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteBlackholeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeRequest) SetInstanceId(v string) *DeleteBlackholeRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetIp(v string) *DeleteBlackholeRequest {
	s.Ip = &v
	return s
}

func (s *DeleteBlackholeRequest) SetRegionId(v string) *DeleteBlackholeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetResourceGroupId(v string) *DeleteBlackholeRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteBlackholeResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBlackholeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteBlackholeResponse) SetStatusCode(v int32) *DeleteBlackholeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBlackholeResponse) SetBody(v *DeleteBlackholeResponseBody) *DeleteBlackholeResponse {
	s.Body = v
	return s
}

type DeleteIpRequest struct {
	// The ID of the Anti-DDoS Origin Enterprise instance.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin Enterprise instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IP addresses that you want to remove from the Anti-DDoS Origin Enterprise instance. This parameter is a string consisting of JSON arrays. Each element in a JSON array is a JSON struct that includes the following field:
	//
	// *   **ip**: required. The IP address that you want to remove. Data type: string.
	//
	//     > The IP addresses that you want to remove must be protected by the Anti-DDoS Origin Enterprise instance.
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region ID of the Anti-DDoS Origin Enterprise instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin Enterprise instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](~~94485~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpRequest) SetInstanceId(v string) *DeleteIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteIpRequest) SetIpList(v string) *DeleteIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteIpRequest) SetRegionId(v string) *DeleteIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpRequest) SetResourceGroupId(v string) *DeleteIpRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteIpResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteIpResponse) SetStatusCode(v int32) *DeleteIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpResponse) SetBody(v *DeleteIpResponseBody) *DeleteIpResponse {
	s.Body = v
	return s
}

type DeleteSchedruleOnDemandRequest struct {
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the scheduling rule that you want to delete.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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

func (s *DeleteSchedruleOnDemandRequest) SetRegionId(v string) *DeleteSchedruleOnDemandRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSchedruleOnDemandRequest) SetRuleName(v string) *DeleteSchedruleOnDemandRequest {
	s.RuleName = &v
	return s
}

type DeleteSchedruleOnDemandResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSchedruleOnDemandResponse) SetStatusCode(v int32) *DeleteSchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchedruleOnDemandResponse) SetBody(v *DeleteSchedruleOnDemandResponseBody) *DeleteSchedruleOnDemandResponse {
	s.Body = v
	return s
}

type DescribeAssetGroupRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupRequest) SetName(v string) *DescribeAssetGroupRequest {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetRegion(v string) *DescribeAssetGroupRequest {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetRegionId(v string) *DescribeAssetGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetSourceIp(v string) *DescribeAssetGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetGroupRequest) SetType(v string) *DescribeAssetGroupRequest {
	s.Type = &v
	return s
}

type DescribeAssetGroupResponseBody struct {
	AssetGroupList []*DescribeAssetGroupResponseBodyAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total          *int64                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAssetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponseBody) SetAssetGroupList(v []*DescribeAssetGroupResponseBodyAssetGroupList) *DescribeAssetGroupResponseBody {
	s.AssetGroupList = v
	return s
}

func (s *DescribeAssetGroupResponseBody) SetRequestId(v string) *DescribeAssetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetGroupResponseBody) SetTotal(v int64) *DescribeAssetGroupResponseBody {
	s.Total = &v
	return s
}

type DescribeAssetGroupResponseBodyAssetGroupList struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupResponseBodyAssetGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupResponseBodyAssetGroupList) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetName(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetRegion(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupResponseBodyAssetGroupList) SetType(v string) *DescribeAssetGroupResponseBodyAssetGroupList {
	s.Type = &v
	return s
}

type DescribeAssetGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAssetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponse) SetHeaders(v map[string]*string) *DescribeAssetGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetGroupResponse) SetStatusCode(v int32) *DescribeAssetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetGroupResponse) SetBody(v *DescribeAssetGroupResponseBody) *DescribeAssetGroupResponse {
	s.Body = v
	return s
}

type DescribeAssetGroupToInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberUid  *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceRequest) SetInstanceId(v string) *DescribeAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetMemberUid(v string) *DescribeAssetGroupToInstanceRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetName(v string) *DescribeAssetGroupToInstanceRequest {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetRegion(v string) *DescribeAssetGroupToInstanceRequest {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetRegionId(v string) *DescribeAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetSourceIp(v string) *DescribeAssetGroupToInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetGroupToInstanceRequest) SetType(v string) *DescribeAssetGroupToInstanceRequest {
	s.Type = &v
	return s
}

type DescribeAssetGroupToInstanceResponseBody struct {
	DataList  []*DescribeAssetGroupToInstanceResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAssetGroupToInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetDataList(v []*DescribeAssetGroupToInstanceResponseBodyDataList) *DescribeAssetGroupToInstanceResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetRequestId(v string) *DescribeAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBody) SetTotal(v int64) *DescribeAssetGroupToInstanceResponseBody {
	s.Total = &v
	return s
}

type DescribeAssetGroupToInstanceResponseBodyDataList struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberUid  *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetGroupToInstanceResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetInstanceId(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetMemberUid(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetName(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetRegion(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Region = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponseBodyDataList) SetType(v string) *DescribeAssetGroupToInstanceResponseBodyDataList {
	s.Type = &v
	return s
}

type DescribeAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAssetGroupToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *DescribeAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetGroupToInstanceResponse) SetStatusCode(v int32) *DescribeAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponse) SetBody(v *DescribeAssetGroupToInstanceResponseBody) *DescribeAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

type DescribeDdosEventRequest struct {
	// The end time of the DDoS attack event to query. This value is a UNIX timestamp. Unit: seconds.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The attacked IP address to query.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start time of the DDoS attack event to query. This value is a UNIX timestamp. Unit: seconds.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventRequest) SetEndTime(v int32) *DescribeDdosEventRequest {
	s.EndTime = &v
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

func (s *DescribeDdosEventRequest) SetPageNo(v int32) *DescribeDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int32) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetRegionId(v string) *DescribeDdosEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetResourceGroupId(v string) *DescribeDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetStartTime(v int32) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

type DescribeDdosEventResponseBody struct {
	// The details about the DDoS attack event.
	Events []*DescribeDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseBody) SetEvents(v []*DescribeDdosEventResponseBodyEvents) *DescribeDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeDdosEventResponseBody) SetRequestId(v string) *DescribeDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponseBody) SetTotal(v int64) *DescribeDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeDdosEventResponseBodyEvents struct {
	// The time when the DDoS attack stopped. This value is a UNIX timestamp. Unit: seconds.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The attacked IP address.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The volume of the request traffic at the start of the DDoS attack. Unit: Mbit/s.
	Mbps *int32 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The number of packets at the start of the DDoS attack. Unit: packets per second (PPS).
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The time when the DDoS attack started. This value is a UNIX timestamp. Unit: seconds.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the DDoS attack event. Valid values:
	//
	// *   **hole_begin**: indicates that blackhole filtering is triggered for the attacked IP address.
	// *   **hole_end**: indicates that blackhole filtering is deactivated for the attacked IP address.
	// *   **defense_begin**: indicates that attack traffic is being scrubbed.
	// *   **defense_end**: indicates that attack traffic is scrubbed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeDdosEventResponseBodyEvents) SetIp(v string) *DescribeDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetPps(v int32) *DescribeDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseBodyEvents) SetStatus(v string) *DescribeDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeDdosEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDdosEventResponse) SetStatusCode(v int32) *DescribeDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosEventResponse) SetBody(v *DescribeDdosEventResponseBody) *DescribeDdosEventResponse {
	s.Body = v
	return s
}

type DescribeExcpetionCountRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeExcpetionCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountRequest) SetRegionId(v string) *DescribeExcpetionCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExcpetionCountRequest) SetResourceGroupId(v string) *DescribeExcpetionCountRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeExcpetionCountResponseBody struct {
	// The number of assets that are in an abnormal state.
	ExceptionIpCount *int32 `json:"ExceptionIpCount,omitempty" xml:"ExceptionIpCount,omitempty"`
	// The number of Anti-DDoS Origin instances that are about to expire.
	ExpireTimeCount *int32 `json:"ExpireTimeCount,omitempty" xml:"ExpireTimeCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExcpetionCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcpetionCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponseBody) SetExceptionIpCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExceptionIpCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetExpireTimeCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExpireTimeCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetRequestId(v string) *DescribeExcpetionCountResponseBody {
	s.RequestId = &v
	return s
}

type DescribeExcpetionCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeExcpetionCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeExcpetionCountResponse) SetStatusCode(v int32) *DescribeExcpetionCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExcpetionCountResponse) SetBody(v *DescribeExcpetionCountResponseBody) *DescribeExcpetionCountResponse {
	s.Body = v
	return s
}

type DescribeInstanceListRequest struct {
	// The IDs of the Anti-DDoS Origin instances to query. Specify the value is in the `["<Instance ID 1>","<Instance ID 2>",……]` format.
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	// The mitigation plan of the Anti-DDoS Origin instance to query. Valid values:
	//
	// *   **0**: the Professional mitigation plan
	// *   **1**: the Enterprise mitigation plan
	InstanceType     *string   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeList []*string `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Repeated"`
	// The IP address of the object that is protected by the Anti-DDoS Origin instance to query.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The protocol type of the IP address asset that is protected by the Anti-DDoS Origin instance to query. Valid values:
	//
	// *   **Ipv4**: IPv4
	// *   **Ipv6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The field that is used to sort the Anti-DDoS Origin instances. Set the value to **expireTime**, which indicates that the instances are sorted based on the expiration time.
	//
	// You can set the **Orderdire** parameter to specify the sorting method.
	Orderby *string `json:"Orderby,omitempty" xml:"Orderby,omitempty"`
	// The sorting method. Valid values:
	//
	// *   **desc**: the descending order. This is the default value.
	// *   **asc**: the ascending order.
	Orderdire *string `json:"Orderdire,omitempty" xml:"Orderdire,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance to query resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the Anti-DDoS Origin instance to query. Fuzzy match is supported.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequest) SetInstanceIdList(v string) *DescribeInstanceListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceType(v string) *DescribeInstanceListRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceTypeList(v []*string) *DescribeInstanceListRequest {
	s.InstanceTypeList = v
	return s
}

func (s *DescribeInstanceListRequest) SetIp(v string) *DescribeInstanceListRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceListRequest) SetIpVersion(v string) *DescribeInstanceListRequest {
	s.IpVersion = &v
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

func (s *DescribeInstanceListRequest) SetPageNo(v int32) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int32) *DescribeInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRegionId(v string) *DescribeInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetRemark(v string) *DescribeInstanceListRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListRequest) SetResourceGroupId(v string) *DescribeInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceListRequest) SetTag(v []*DescribeInstanceListRequestTag) *DescribeInstanceListRequest {
	s.Tag = v
	return s
}

type DescribeInstanceListRequestTag struct {
	// The key of the tag that is added to the Anti-DDoS Origin instance to query.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the Anti-DDoS Origin instance to query.
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
	// The details about the Anti-DDoS Origin instance.
	InstanceList []*DescribeInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of Anti-DDoS Origin instances.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBody) SetInstanceList(v []*DescribeInstanceListResponseBodyInstanceList) *DescribeInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *DescribeInstanceListResponseBody) SetRequestId(v string) *DescribeInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponseBody) SetTotal(v int64) *DescribeInstanceListResponseBody {
	s.Total = &v
	return s
}

type DescribeInstanceListResponseBodyInstanceList struct {
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The number of protected public IP addresses for which blackhole filtering is triggered.
	//
	// >  You can call the [DeleteBlackhole](~~118692~~) operation to deactivate blackhole filtering for a protected IP address.
	BlackholdingCount *string `json:"BlackholdingCount,omitempty" xml:"BlackholdingCount,omitempty"`
	CommodityType     *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	CoverageType      *int32  `json:"CoverageType,omitempty" xml:"CoverageType,omitempty"`
	// The time when the instance expires. This value is a UNIX timestamp. Unit: milliseconds.
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance was purchased. This value is a UNIX timestamp. Unit: milliseconds.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mitigation plan of the instance. Valid values:
	//
	// *   **0**: the Professional mitigation plan
	// *   **1**: the Enterprise mitigation plan
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The protocol type of the IP address asset that is protected by the instance. Valid values:
	//
	// *   **Ipv4**: IPv4
	// *   **Ipv6**: IPv6
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The type of the cloud service that is associated with the Anti-DDoS Origin instance. By default, this parameter is not returned. If the Anti-DDoS Origin instance is created by using a different cloud service, the code of the cloud service is returned.
	//
	// Valid values:
	//
	// *   **gamebox**: The Anti-DDoS Origin instance is created by using Game Security Box.
	// *   **eip**: The Anti-DDoS Origin instance is created by using an elastic IP address (EIP) for which Anti-DDoS (Enhanced Edition) is enabled.
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The remarks of the instance.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the instance. Valid values:
	//
	// *   **1**: normal
	// *   **2**: expired
	// *   **3**: released
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetAutoRenewal(v bool) *DescribeInstanceListResponseBodyInstanceList {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetBlackholdingCount(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.BlackholdingCount = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetCommodityType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.CommodityType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetCoverageType(v int32) *DescribeInstanceListResponseBodyInstanceList {
	s.CoverageType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
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

func (s *DescribeInstanceListResponseBodyInstanceList) SetIpType(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.IpType = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetProduct(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetRemark(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseBodyInstanceList) SetStatus(v string) *DescribeInstanceListResponseBodyInstanceList {
	s.Status = &v
	return s
}

type DescribeInstanceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceListResponse) SetStatusCode(v int32) *DescribeInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceListResponse) SetBody(v *DescribeInstanceListResponseBody) *DescribeInstanceListResponse {
	s.Body = v
	return s
}

type DescribeInstanceSpecsRequest struct {
	// The ID of the Anti-DDoS Origin Enterprise instance. This parameter value is a string consisting of JSON arrays. Each element in a JSON array indicates an instance ID. If you want to query more than one instance, separate instance IDs with commas (,).
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin Enterprise instances in a specific region.
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	// The region ID of the Anti-DDoS Origin Enterprise instance. Default value: **cn-hangzhou**, which indicates the China (Hangzhou) region.
	//
	// >  If your instance does not reside in the China (Hangzhou) region, you must specify this parameter to the region ID of your instance. You can call the [DescribeRegions](~~118703~~) operation to query the regions of cloud assets that are supported by an Anti-DDoS Origin instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin Enterprise instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin Enterprise instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](~~94485~~).
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

type DescribeInstanceSpecsResponseBody struct {
	// The specifications of the Anti-DDoS Origin Enterprise instance, including whether the unlimited protection feature is enabled, and the numbers of times that the unlimited protection feature can be enabled and has been enabled.
	InstanceSpecs []*DescribeInstanceSpecsResponseBodyInstanceSpecs `json:"InstanceSpecs,omitempty" xml:"InstanceSpecs,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBody) SetInstanceSpecs(v []*DescribeInstanceSpecsResponseBodyInstanceSpecs) *DescribeInstanceSpecsResponseBody {
	s.InstanceSpecs = v
	return s
}

func (s *DescribeInstanceSpecsResponseBody) SetRequestId(v string) *DescribeInstanceSpecsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecs struct {
	// The number of times that the unlimited protection feature can be enabled.
	AvailableDefenseTimes *int32 `json:"AvailableDefenseTimes,omitempty" xml:"AvailableDefenseTimes,omitempty"`
	// The number of times that blackhole filtering can be deactivated.
	AvailableDeleteBlackholeCount *string `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty"`
	DefenseTimesPercent           *int32  `json:"DefenseTimesPercent,omitempty" xml:"DefenseTimesPercent,omitempty"`
	// The ID of the Anti-DDoS Origin Enterprise instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the unlimited protection feature is enabled. Valid values:
	//
	// *   **0**: The unlimited protection feature is disabled.
	// *   **1**: The unlimited protection feature is enabled.
	IsFullDefenseMode *int32 `json:"IsFullDefenseMode,omitempty" xml:"IsFullDefenseMode,omitempty"`
	// The configurations of the Anti-DDoS Origin Enterprise instance, including the number of protected IP addresses and protection bandwidth.
	PackConfig *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" type:"Struct"`
	// The region ID of the Anti-DDoS Origin Enterprise instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the name of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of times that the unlimited protection feature can be enabled.
	TotalDefenseTimes *int32 `json:"TotalDefenseTimes,omitempty" xml:"TotalDefenseTimes,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecs) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDefenseTimes = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetAvailableDeleteBlackholeCount(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetDefenseTimesPercent(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.DefenseTimesPercent = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetInstanceId(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetIsFullDefenseMode(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.IsFullDefenseMode = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetPackConfig(v *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.PackConfig = v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetRegion(v string) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.Region = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecs) SetTotalDefenseTimes(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecs {
	s.TotalDefenseTimes = &v
	return s
}

type DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig struct {
	// The bandwidth of the package configuration.
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of IP addresses that are protected by the Anti-DDoS Origin Enterprise instance.
	BindIpCount *int32 `json:"BindIpCount,omitempty" xml:"BindIpCount,omitempty"`
	// The burstable bandwidth of each protected IP address. Unit: Gbit/s.
	IpAdvanceThre *int32 `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty"`
	// The basic bandwidth of each protected IP address. Unit: Gbit/s.
	IpBasicThre *int32 `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty"`
	// The number of IP addresses that can be protected by the Anti-DDoS Origin Enterprise instance.
	IpSpec *int32 `json:"IpSpec,omitempty" xml:"IpSpec,omitempty"`
	// The normal clean bandwidth. Unit: Mbit/s.
	NormalBandwidth *int32 `json:"NormalBandwidth,omitempty" xml:"NormalBandwidth,omitempty"`
	// The burstable protection bandwidth of the Anti-DDoS Origin Enterprise instance. Unit: Gbit/s.
	PackAdvThre *int32 `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty"`
	// The basic protection bandwidth of the Anti-DDoS Origin Enterprise instance. Unit: Gbit/s.
	PackBasicThre *int32 `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty"`
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBandwidth(v int64) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetBindIpCount(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.BindIpCount = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpAdvanceThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetIpSpec(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.IpSpec = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetNormalBandwidth(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.NormalBandwidth = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackAdvThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig) SetPackBasicThre(v int32) *DescribeInstanceSpecsResponseBodyInstanceSpecsPackConfig {
	s.PackBasicThre = &v
	return s
}

type DescribeInstanceSpecsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceSpecsResponse) SetStatusCode(v int32) *DescribeInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

type DescribeOnDemandDdosEventRequest struct {
	// The timestamp that specifies the end of the time range to query. Unit: seconds. The timestamp follows the UNIX time format. It is the number of seconds that have elapsed since 00:00:00 Thursday, 1 January 1970.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the on-demand instance to query.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address of the protection target.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. The maximum value is **50**. The default value is **10**.
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The timestamp that specifies the beginning of the time range to query. Unit: seconds. The timestamp follows the UNIX time format. It is the number of seconds that have elapsed since 00:00:00 Thursday, 1 January 1970.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOnDemandDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventRequest) SetEndTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.EndTime = &v
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

func (s *DescribeOnDemandDdosEventRequest) SetPageNo(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetPageSize(v int32) *DescribeOnDemandDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetRegionId(v string) *DescribeOnDemandDdosEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetResourceGroupId(v string) *DescribeOnDemandDdosEventRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOnDemandDdosEventRequest) SetStartTime(v int32) *DescribeOnDemandDdosEventRequest {
	s.StartTime = &v
	return s
}

type DescribeOnDemandDdosEventResponseBody struct {
	// The list of DDoS events and the details of each event.
	Events []*DescribeOnDemandDdosEventResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS events.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeOnDemandDdosEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandDdosEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandDdosEventResponseBody) SetEvents(v []*DescribeOnDemandDdosEventResponseBodyEvents) *DescribeOnDemandDdosEventResponseBody {
	s.Events = v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetRequestId(v string) *DescribeOnDemandDdosEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBody) SetTotal(v int64) *DescribeOnDemandDdosEventResponseBody {
	s.Total = &v
	return s
}

type DescribeOnDemandDdosEventResponseBodyEvents struct {
	// The timestamp that indicates the end time of the attack. Unit: seconds. The timestamp follows the UNIX time format. It is the number of seconds that have elapsed since 00:00:00 Thursday, 1 January 1970.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address of the protection target that encounters the DDoS attack.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The throughput of the DDoS attack. Unit: Mbit/s.
	Mbps *int32 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The packet forwarding rate of the DDoS attack. Unit: packets per second (PPS).
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The timestamp that indicates the start time of the attack. Unit: seconds. The timestamp follows the UNIX time format. It is the number of seconds that have elapsed since 00:00:00 Thursday, 1 January 1970.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the event. Valid values:
	//
	// *   **hole_begin **: indicates that the event is in the blackhole state.
	// *   **hole_end **: indicates that blackhole ends.
	// *   **defense_begin **: indicates that the event is in the cleaning state.
	// *   **defense_end **: indicates that cleaning ends.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetIp(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Ip = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetMbps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetPps(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Pps = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStartTime(v int32) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponseBodyEvents) SetStatus(v string) *DescribeOnDemandDdosEventResponseBodyEvents {
	s.Status = &v
	return s
}

type DescribeOnDemandDdosEventResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOnDemandDdosEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeOnDemandDdosEventResponse) SetStatusCode(v int32) *DescribeOnDemandDdosEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnDemandDdosEventResponse) SetBody(v *DescribeOnDemandDdosEventResponseBody) *DescribeOnDemandDdosEventResponse {
	s.Body = v
	return s
}

type DescribeOnDemandInstanceStatusRequest struct {
	// The IDs of on-demand instances.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type DescribeOnDemandInstanceStatusResponseBody struct {
	// The details of the on-demand instance.
	Instances []*DescribeOnDemandInstanceStatusResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The details of route advertisement for data centers outside the Chinese mainland. This parameter is a JSON string. The following fields are included in the value:
	//
	// *   **region**: The code of the data center outside the Chinese mainland. The value is of the STRING type. For more information, see **Codes of data centers outside the Chinese mainland**.
	// *   **declared**: indicates whether the data center advertised the route. The value is of the STRING type. Valid values: **0** and **1**. The value of 0 indicates that the data center did not advertise the route. The value of 1 indicates that the data center advertised the route.
	Declared *string `json:"Declared,omitempty" xml:"Declared,omitempty"`
	// The description of the on-demand instance.
	//
	// >  The value of this parameter is returned only when the information about multiple on-demand instances is returned. The value of this parameter is not returned because the information about only one on-demand instance is returned.
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the on-demand instance.
	//
	// >  The value of this parameter is returned only when the information about multiple on-demand instances is returned. The value of this parameter is not returned because the information about only one on-demand instance is returned.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mode used to start the on-demand instance. Valid values:
	//
	// *   **manual**: The instance is manually started.
	// *   **netflow-auto**: The instance is automatically started by using NetFlow that monitors network traffic.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The CIDR block of the on-demand instance.
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The number of the autonomous system (AS). Set the value to **0**, which indicates that AS is disabled.
	RegistedAs *string `json:"RegistedAs,omitempty" xml:"RegistedAs,omitempty"`
	// The ID of the Alibaba Cloud account.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceStatusResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDeclared(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Declared = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetDesc(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Desc = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetInstanceId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetMode(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Mode = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetNet(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.Net = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetRegistedAs(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.RegistedAs = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponseBodyInstances) SetUserId(v string) *DescribeOnDemandInstanceStatusResponseBodyInstances {
	s.UserId = &v
	return s
}

type DescribeOnDemandInstanceStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOnDemandInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeOnDemandInstanceStatusResponse) SetStatusCode(v int32) *DescribeOnDemandInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnDemandInstanceStatusResponse) SetBody(v *DescribeOnDemandInstanceStatusResponseBody) *DescribeOnDemandInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeOpEntitiesRequest struct {
	// The operation that you want to perform. Set the value to **DescribeOpEntities**.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The details of the operation log.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number of the returned page.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sort order of operation logs. Valid values:
	//
	// *   **ASC**: the ascending order.
	// *   **DESC**: the descending order.
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	OrderDir *string `json:"OrderDir,omitempty" xml:"OrderDir,omitempty"`
	// The type of the operation object. The value is fixed as **1**, which indicates Anti-DDoS Origin instances.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Alibaba Cloud account that performs the operation.
	//
	// >  If the value is **system**, the operation is performed by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The details about the operation. The value is a string that consists of a JSON struct. The JSON struct contains the following fields:
	//
	// *   **entity**: the operation object. Data type: object. The fields that are included in the value of the **entity** parameter vary based on the value of the **OpAction** parameter. Take note of the following items:
	//
	//     *   If the value of the **OpAction** parameter is **3**, the value of the **entity** parameter consists of the following field:
	//
	//         *   **ips**: the public IP addresses that are protected by the Anti-DDoS Origin instance. Data type: array
	//
	//     *   If the value of the **OpAction** parameter is **4**, the value of the **entity** parameter consists of the following field:
	//
	//         *   **ips**: the public IP addresses that are no longer protected by the Anti-DDoS Origin instance. Data type: array.
	//
	//     *   If the value of the **OpAction** parameter is **5**, the value of the **entity** parameter consists of the following fields:
	//
	//         *   **baseBandwidth**: the basic protection bandwidth. Unit: Gbit/s. Data type: integer.
	//         *   **elasticBandwidth**: the burstable protection bandwidth. Unit: Gbit/s. Data type: integer.
	//         *   **opSource**: the source of the operation. The value is fixed as **1**, indicating that the operation is performed by Anti-DDoS Origin. Data type: integer.
	//
	//     *   If the value of the **OpAction** parameter is **6**, the value of the **entity** parameter consists of the following field:
	//
	//         *   **ips**: the public IP addresses for which to deactivate blackhole filtering. Data type: array.
	//
	//     *   If the value of the **OpAction** parameter is **7**, the **entity** parameter is not returned.
	//
	//     *   If the value of the **OpAction** parameter is **8**, the value of the **entity** parameter consists of the following fields:
	//
	//         *   **baseBandwidth**: the basic protection bandwidth. Unit: Gbit/s. Data type: integer.
	//         *   **elasticBandwidth**: the burstable protection bandwidth. Unit: Gbit/s. Data type: integer.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sorting method of operation logs. Set the value to **opdate**, which indicates sorting based on the operation time.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int32) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetInstanceId(v string) *DescribeOpEntitiesRequest {
	s.InstanceId = &v
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

func (s *DescribeOpEntitiesRequest) SetPageSize(v int32) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetRegionId(v string) *DescribeOpEntitiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetResourceGroupId(v string) *DescribeOpEntitiesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

type DescribeOpEntitiesResponseBody struct {
	// The ID of the request.
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// The end time. Operation logs that were generated before this time are queried.**** This value is a UNIX timestamp. Unit: milliseconds.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the log was generated. This value is a UNIX timestamp. Unit: milliseconds.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotalCount(v int32) *DescribeOpEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	// Queries the operation logs of an Anti-DDoS Origin instance.
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// All Alibaba Cloud API operations must include common request parameters. For more information about common request parameters, see [Common parameters](~~118841~~).
	//
	// For more information about sample requests, see the **"Examples"** section of this topic.
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// WB01342967
	GmtCreate *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// DescribeOpEntities
	OpAction *int32  `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	OpDesc   *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
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

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOpEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeOpEntitiesResponse) SetStatusCode(v int32) *DescribeOpEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetBody(v *DescribeOpEntitiesResponseBody) *DescribeOpEntitiesResponse {
	s.Body = v
	return s
}

type DescribePackIpListRequest struct {
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The protected IP address to query.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the member.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the cloud asset to which the protected IP address to query belongs. Valid values:
	//
	// *   **ECS**: an Elastic Compute Service (ECS) instance.
	// *   **SLB**: a Classic Load Balancer (CLB) instance, originally called a Server Load Balancer (SLB) instance.
	// *   **EIP**: an elastic IP address (EIP). An Internet-facing Application Load Balancer (ALB) instance uses an EIP. If the IP address belongs to the Internet-facing ALB instance, set this parameter to EIP.
	// *   **WAF**: a Web Application Firewall (WAF) instance.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribePackIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListRequest) GoString() string {
	return s.String()
}

func (s *DescribePackIpListRequest) SetInstanceId(v string) *DescribePackIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackIpListRequest) SetIp(v string) *DescribePackIpListRequest {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListRequest) SetMemberUid(v string) *DescribePackIpListRequest {
	s.MemberUid = &v
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

func (s *DescribePackIpListRequest) SetProductName(v string) *DescribePackIpListRequest {
	s.ProductName = &v
	return s
}

func (s *DescribePackIpListRequest) SetRegionId(v string) *DescribePackIpListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePackIpListRequest) SetResourceGroupId(v string) *DescribePackIpListRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribePackIpListResponseBody struct {
	// The HTTP status code of the request.
	//
	// For more information about status codes, see [Common parameters](~~118841~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IP addresses that are protected by the instance.
	IpList []*DescribePackIpListResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   **true**: The call is successful.
	// *   **false**: The call fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of protected IP addresses.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePackIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBody) SetCode(v string) *DescribePackIpListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetIpList(v []*DescribePackIpListResponseBodyIpList) *DescribePackIpListResponseBody {
	s.IpList = v
	return s
}

func (s *DescribePackIpListResponseBody) SetRequestId(v string) *DescribePackIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetSuccess(v bool) *DescribePackIpListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePackIpListResponseBody) SetTotal(v int32) *DescribePackIpListResponseBody {
	s.Total = &v
	return s
}

type DescribePackIpListResponseBodyIpList struct {
	// The IP address.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the member.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The type of the cloud asset to which the IP address belongs. Valid values:
	//
	// *   **ECS**: an ECS instance.
	// *   **SLB**: a CLB instance, originally called an SLB instance.
	// *   **EIP**: an EIP. If the IP address belongs to an ALB instance, the value EIP is returned.
	// *   **WAF**: a WAF instance.
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region to which the protected IP address belongs.
	//
	// >  If the protected IP address is in the same region as the instance, this parameter is not returned.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The description of the cloud asset to which the IP address belongs. The asset can be an ECS instance or an SLB instance.
	//
	// >  If no descriptions are provided for the asset, this parameter is not returned.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the IP address. Valid values:
	//
	// *   **normal**: The IP address is in the normal state, which indicates that the IP address is not under attack.
	// *   **hole_begin**: Blackhole filtering is triggered for the IP address.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePackIpListResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackIpListResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *DescribePackIpListResponseBodyIpList) SetIp(v string) *DescribePackIpListResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetMemberUid(v string) *DescribePackIpListResponseBodyIpList {
	s.MemberUid = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetProduct(v string) *DescribePackIpListResponseBodyIpList {
	s.Product = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRegion(v string) *DescribePackIpListResponseBodyIpList {
	s.Region = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetRemark(v string) *DescribePackIpListResponseBodyIpList {
	s.Remark = &v
	return s
}

func (s *DescribePackIpListResponseBodyIpList) SetStatus(v string) *DescribePackIpListResponseBodyIpList {
	s.Status = &v
	return s
}

type DescribePackIpListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePackIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribePackIpListResponse) SetStatusCode(v int32) *DescribePackIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackIpListResponse) SetBody(v *DescribePackIpListResponseBody) *DescribePackIpListResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The region ID to query. The default value is **cn-hangzhou**, which indicates that the regions of cloud assets that are supported by an Anti-DDoS Origin instance in the China (Hangzhou) region are queried.
	//
	// For more information about the IDs of other regions, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](~~94485~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceGroupId(v string) *DescribeRegionsRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The HTTP status code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about regions of the cloud assets that are supported by the Anti-DDoS Origin instance. The information includes region IDs and names.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   **true**: The request is successful.
	// *   **false**: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The English name of the region where the cloud assets reside.
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region where the cloud assets reside.
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
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

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeTrafficRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// If you do not specify this parameter, the current system time is used as the end time.
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the traffic statistics to query. Valid values:
	//
	// *   **max**: the peak traffic within the specified interval
	// *   **avg**: the average traffic within the specified interval
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The ID of the Anti-DDoS Origin instance to query.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin instances.
	//
	// If you specify an on-demand instance, you must configure the **Interval** parameter.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The interval at which the traffic statistics are calculated. Unit: seconds. Default value: **5**.
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The public IP address of the asset to query. If you do not specify this parameter, the traffic statistics of all public IP addresses that are protected by the Anti-DDoS Origin instance are queried.
	//
	// >  The public IP address must be a protected object of the Anti-DDoS Origin instance. You can call the [DescribePackIpList](~~118701~~) operation to query all protected objects of the Anti-DDoS Origin instance.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The Classless Inter-Domain Routing (CIDR) block of the on-demand instance that you want to query.
	Ipnet *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficRequest) SetEndTime(v int32) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetFlowType(v string) *DescribeTrafficRequest {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficRequest) SetInstanceId(v string) *DescribeTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int32) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeTrafficRequest) SetIp(v string) *DescribeTrafficRequest {
	s.Ip = &v
	return s
}

func (s *DescribeTrafficRequest) SetIpnet(v string) *DescribeTrafficRequest {
	s.Ipnet = &v
	return s
}

func (s *DescribeTrafficRequest) SetRegionId(v string) *DescribeTrafficRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTrafficRequest) SetResourceGroupId(v string) *DescribeTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTrafficRequest) SetStartTime(v int32) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeTrafficResponseBody struct {
	// The queried traffic statistics.
	FlowList []*DescribeTrafficResponseBodyFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The bandwidth of attack traffic. Unit: bit/s.
	//
	// >  This parameter is returned only if attack traffic exists.
	AttackBps *int64 `json:"AttackBps,omitempty" xml:"AttackBps,omitempty"`
	// The packet forwarding rate of attack traffic. Unit: packets per second.
	//
	// >  This parameter is returned only if attack traffic exists.
	AttackPps *int64 `json:"AttackPps,omitempty" xml:"AttackPps,omitempty"`
	// The type of the traffic statistics. Valid values:
	//
	// *   **max**: the peak traffic within the specified interval
	// *   **avg**: the average traffic within the specified interval
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The bandwidth of the total traffic. Unit: Kbit/s.
	Kbps *int32 `json:"Kbps,omitempty" xml:"Kbps,omitempty"`
	// The ID of the traffic statistics.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The packet forwarding rate of the total traffic. Unit: packets per second.
	Pps *int32 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The time when the traffic statistics are calculated. This value is a UNIX timestamp. Unit: seconds.
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeTrafficResponseBodyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseBodyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackBps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackBps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetAttackPps(v int64) *DescribeTrafficResponseBodyFlowList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetFlowType(v string) *DescribeTrafficResponseBodyFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseBodyFlowList) SetKbps(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Kbps = &v
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

func (s *DescribeTrafficResponseBodyFlowList) SetTime(v int32) *DescribeTrafficResponseBodyFlowList {
	s.Time = &v
	return s
}

type DescribeTrafficResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTrafficResponse) SetStatusCode(v int32) *DescribeTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficResponse) SetBody(v *DescribeTrafficResponseBody) *DescribeTrafficResponse {
	s.Body = v
	return s
}

type DettachAssetGroupToInstanceRequest struct {
	AssetGroupList []*DettachAssetGroupToInstanceRequestAssetGroupList `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty" type:"Repeated"`
	InstanceId     *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceIp       *string                                             `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DettachAssetGroupToInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceRequest) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceRequest) SetAssetGroupList(v []*DettachAssetGroupToInstanceRequestAssetGroupList) *DettachAssetGroupToInstanceRequest {
	s.AssetGroupList = v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) SetInstanceId(v string) *DettachAssetGroupToInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) SetRegionId(v string) *DettachAssetGroupToInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequest) SetSourceIp(v string) *DettachAssetGroupToInstanceRequest {
	s.SourceIp = &v
	return s
}

type DettachAssetGroupToInstanceRequestAssetGroupList struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DettachAssetGroupToInstanceRequestAssetGroupList) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceRequestAssetGroupList) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetName(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Name = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetRegion(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Region = &v
	return s
}

func (s *DettachAssetGroupToInstanceRequestAssetGroupList) SetType(v string) *DettachAssetGroupToInstanceRequestAssetGroupList {
	s.Type = &v
	return s
}

type DettachAssetGroupToInstanceShrinkRequest struct {
	AssetGroupListShrink *string `json:"AssetGroupList,omitempty" xml:"AssetGroupList,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceIp             *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DettachAssetGroupToInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetAssetGroupListShrink(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.AssetGroupListShrink = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetInstanceId(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetRegionId(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DettachAssetGroupToInstanceShrinkRequest) SetSourceIp(v string) *DettachAssetGroupToInstanceShrinkRequest {
	s.SourceIp = &v
	return s
}

type DettachAssetGroupToInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DettachAssetGroupToInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceResponseBody) SetRequestId(v string) *DettachAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DettachAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DettachAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DettachAssetGroupToInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DettachAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *DettachAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *DettachAssetGroupToInstanceResponse) SetStatusCode(v int32) *DettachAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DettachAssetGroupToInstanceResponse) SetBody(v *DettachAssetGroupToInstanceResponseBody) *DettachAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

type GetSlsOpenStatusRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// For more information about the valid values of this parameter, see [Regions and zones](~~188196~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](~~94485~~).
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

type GetSlsOpenStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Log Service was activated. Valid values:
	//
	// *   **true**: Log Service was activated.
	// *   **false**: Log Service was not activated.
	SlsOpenStatus *bool `json:"SlsOpenStatus,omitempty" xml:"SlsOpenStatus,omitempty"`
}

func (s GetSlsOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponseBody) SetRequestId(v string) *GetSlsOpenStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSlsOpenStatusResponseBody) SetSlsOpenStatus(v bool) *GetSlsOpenStatusResponseBody {
	s.SlsOpenStatus = &v
	return s
}

type GetSlsOpenStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSlsOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponse) SetHeaders(v map[string]*string) *GetSlsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSlsOpenStatusResponse) SetStatusCode(v int32) *GetSlsOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSlsOpenStatusResponse) SetBody(v *GetSlsOpenStatusResponseBody) *GetSlsOpenStatusResponse {
	s.Body = v
	return s
}

type ListOpenedAccessLogInstancesRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management. This parameter is empty by default, which indicates that the Anti-DDoS Origin instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](~~94485~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListOpenedAccessLogInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageNumber(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetPageSize(v int32) *ListOpenedAccessLogInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOpenedAccessLogInstancesRequest) SetResourceGroupId(v string) *ListOpenedAccessLogInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

type ListOpenedAccessLogInstancesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configuration of log analysis for the Anti-DDoS Origin instance.
	SlsConfigStatus []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus `json:"SlsConfigStatus,omitempty" xml:"SlsConfigStatus,omitempty" type:"Repeated"`
	// The number of the Anti-DDoS Origin instances for which log analysis was enabled.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetRequestId(v string) *ListOpenedAccessLogInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetSlsConfigStatus(v []*ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) *ListOpenedAccessLogInstancesResponseBody {
	s.SlsConfigStatus = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBody) SetTotalCount(v int32) *ListOpenedAccessLogInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListOpenedAccessLogInstancesResponseBodySlsConfigStatus struct {
	// Indicates whether log analysis was enabled for the Anti-DDoS Origin instance. Valid values:
	//
	// *   **true**: Log analysis was enabled.
	// *   **false**: Log analysis was disabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The ID of the Anti-DDoS Origin instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetEnable(v bool) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.Enable = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus) SetInstanceId(v string) *ListOpenedAccessLogInstancesResponseBodySlsConfigStatus {
	s.InstanceId = &v
	return s
}

type ListOpenedAccessLogInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOpenedAccessLogInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOpenedAccessLogInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenedAccessLogInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListOpenedAccessLogInstancesResponse) SetHeaders(v map[string]*string) *ListOpenedAccessLogInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetStatusCode(v int32) *ListOpenedAccessLogInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpenedAccessLogInstancesResponse) SetBody(v *ListOpenedAccessLogInstancesResponseBody) *ListOpenedAccessLogInstancesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to **50**. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the resource. Valid value: **INSTANCE**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetCurrentPage(v int32) *ListTagKeysRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tags and the details of each tag.
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of tags.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	// The total number of tag values that correspond to each key.
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The key of each tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The query token. Set the value to the **NextToken** value that is returned in the last call to the ListTagResources operation. Leave this parameter empty the first time you call this operation.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of Anti-DDoS Origin Instances to query.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin instances.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource to query. Set the value to **INSTANCE**, which indicates instances.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags to query.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag to query.
	//
	// >  The **ResourceIds.N** parameter and the key-value pair (Tag.N.Key and Tag.N.Value) cannot be left empty at the same time.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to query.
	//
	// >  The **ResourceIds.N** parameter and the key-value pair (Tag.N.Key and Tag.N.Value) cannot be left empty at the same time.
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
	// The query token that is returned in this call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tags that are added to the Anti-DDoS Origin instance.
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
	// The ID of the Anti-DDoS Origin instance.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. The value is fixed as **INSTANCE**, which indicates instances.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag that is added to the instance.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag that is added to the instance.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyRemarkRequest struct {
	// The ID of the Anti-DDoS Origin instance for which you want to add remarks.
	//
	// >  You can call the [DescribeInstanceList](~~118698~~) operation to query the IDs of all Anti-DDoS Origin instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks for the Anti-DDoS Origin instance.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s *ModifyRemarkRequest) SetRegionId(v string) *ModifyRemarkRequest {
	s.RegionId = &v
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

type ModifyRemarkResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyRemarkResponse) SetStatusCode(v int32) *ModifyRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRemarkResponse) SetBody(v *ModifyRemarkResponseBody) *ModifyRemarkResponse {
	s.Body = v
	return s
}

type QuerySchedruleOnDemandRequest struct {
	// The ID of the on-demand instance.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type QuerySchedruleOnDemandResponseBody struct {
	// The ID of the on-demand instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the scheduling rule.
	RuleConfig []*QuerySchedruleOnDemandResponseBodyRuleConfig `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty" type:"Repeated"`
	// The status of the scheduling rule.
	RuleStatus []*QuerySchedruleOnDemandResponseBodyRuleStatus `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBody) SetInstanceId(v string) *QuerySchedruleOnDemandResponseBody {
	s.InstanceId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRequestId(v string) *QuerySchedruleOnDemandResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleConfig(v []*QuerySchedruleOnDemandResponseBodyRuleConfig) *QuerySchedruleOnDemandResponseBody {
	s.RuleConfig = v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetRuleStatus(v []*QuerySchedruleOnDemandResponseBodyRuleStatus) *QuerySchedruleOnDemandResponseBody {
	s.RuleStatus = v
	return s
}

func (s *QuerySchedruleOnDemandResponseBody) SetUserId(v string) *QuerySchedruleOnDemandResponseBody {
	s.UserId = &v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleConfig struct {
	// The scheduling action. Set the value to **declare**, which indicates that the route is advertised.
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// If the inbound bandwidth or packets consecutively exceed the threshold for the specified number of times, the scheduling rule is triggered and traffic is rerouted to the on-demand instance. The specified number of times is the value of this parameter.
	//
	// >  The threshold of inbound bandwidth is the value of **RuleConditionMbps**. The threshold of inbound packets is the value of **RuleConditionKpps**.
	RuleConditionCnt *string `json:"RuleConditionCnt,omitempty" xml:"RuleConditionCnt,omitempty"`
	// The threshold of inbound packets. Unit: Kpps. Minimum value: **10**.
	RuleConditionKpps *string `json:"RuleConditionKpps,omitempty" xml:"RuleConditionKpps,omitempty"`
	// The threshold of inbound bandwidth. Unit: Mbit/s. Minimum value: **100**.
	RuleConditionMbps *string `json:"RuleConditionMbps,omitempty" xml:"RuleConditionMbps,omitempty"`
	// The name of the scheduling rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Indicates whether the scheduling rule is enabled. Valid values:
	//
	// *   **on**: enabled
	// *   **off**: disabled
	RuleSwitch *string `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The start time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	//
	// If the system detects that DDoS attacks stop, the system no longer reroutes traffic to the on-demand instance from the time you specified. We recommend that you set this parameter to a value that is defined as off-peak hours.
	//
	// >  This parameter takes effect only when the **RuleUndoMode** parameter is set to **auto**.
	RuleUndoBeginTime *string `json:"RuleUndoBeginTime,omitempty" xml:"RuleUndoBeginTime,omitempty"`
	// The end time of the period during which the scheduling rule is automatically stopped. The time must be in the 24-hour clock and in the `hh:mm` format.
	RuleUndoEndTime *string `json:"RuleUndoEndTime,omitempty" xml:"RuleUndoEndTime,omitempty"`
	// The stop method of the scheduling rule. Valid values:
	//
	// *   **auto**: The scheduling rule automatically stops.
	// *   **manual**: The scheduling rule is manually stopped.
	RuleUndoMode *string `json:"RuleUndoMode,omitempty" xml:"RuleUndoMode,omitempty"`
	// The time zone of the time when the scheduling rule automatically stops. The time zone must be in the `GMT-hh:mm` format.
	//
	// For example, the value `GMT-08:00` indicates that the time zone is UTC+8.
	//
	// >  This parameter takes effect only when the **RuleUndoMode** parameter is set to **auto**.
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleConfig) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleAction(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleAction = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionCnt(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionCnt = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionKpps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionKpps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleConditionMbps(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleConditionMbps = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleName(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleName = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleSwitch(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleSwitch = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoBeginTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoBeginTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoEndTime(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoEndTime = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetRuleUndoMode(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.RuleUndoMode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleConfig) SetTimeZone(v string) *QuerySchedruleOnDemandResponseBodyRuleConfig {
	s.TimeZone = &v
	return s
}

type QuerySchedruleOnDemandResponseBodyRuleStatus struct {
	// The CIDR block of the on-demand instance.
	Net *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The scheduling status. Valid values:
	//
	// *   **scheduled**
	// *   **unscheduled**
	RuleSchedStatus *string `json:"RuleSchedStatus,omitempty" xml:"RuleSchedStatus,omitempty"`
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) String() string {
	return tea.Prettify(s)
}

func (s QuerySchedruleOnDemandResponseBodyRuleStatus) GoString() string {
	return s.String()
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetNet(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.Net = &v
	return s
}

func (s *QuerySchedruleOnDemandResponseBodyRuleStatus) SetRuleSchedStatus(v string) *QuerySchedruleOnDemandResponseBodyRuleStatus {
	s.RuleSchedStatus = &v
	return s
}

type QuerySchedruleOnDemandResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySchedruleOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QuerySchedruleOnDemandResponse) SetStatusCode(v int32) *QuerySchedruleOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySchedruleOnDemandResponse) SetBody(v *QuerySchedruleOnDemandResponseBody) *QuerySchedruleOnDemandResponse {
	s.Body = v
	return s
}

type SetInstanceModeOnDemandRequest struct {
	// The IDs of on-demand instances.
	//
	// >  You can call the [DescribeOnDemandInstance](~~152120~~) operation to query the IDs of all on-demand instances.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The scheduling mode of the on-demand instance. Valid values:
	//
	// *   **manual**: manual scheduling
	// *   **netflow-auto**: automatic scheduling
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The region ID of the on-demand instance.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query all regions supported by Anti-DDoS Origin.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type SetInstanceModeOnDemandResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetInstanceModeOnDemandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetInstanceModeOnDemandResponse) SetStatusCode(v int32) *SetInstanceModeOnDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceModeOnDemandResponse) SetBody(v *SetInstanceModeOnDemandResponseBody) *SetInstanceModeOnDemandResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](~~118703~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource to which you want to add tags. Set the value to **INSTANCE**, which indicates instances.
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag to add.
	//
	// >  If the specified key does not exist, a key is created.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add.
	//
	// >  If the specified value does not exist, a value is created.
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
	// The ID of the request.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified Anti-DDoS Origin Enterprise instances.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of the region where the Anti-DDoS Origin Enterprise instances reside.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the specified resource. Set the value to **INSTANCE**.
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIp"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) AttachAssetGroupToInstanceWithOptions(tmpReq *AttachAssetGroupToInstanceRequest, runtime *util.RuntimeOptions) (_result *AttachAssetGroupToInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AttachAssetGroupToInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssetGroupList)) {
		request.AssetGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetGroupList, tea.String("AssetGroupList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetGroupListShrink)) {
		query["AssetGroupList"] = request.AssetGroupListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachAssetGroupToInstance"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachAssetGroupToInstance(request *AttachAssetGroupToInstanceRequest) (_result *AttachAssetGroupToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachAssetGroupToInstanceResponse{}
	_body, _err := client.AttachAssetGroupToInstanceWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAccessLogAuth"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

/**
 * You can call the CheckGrant operation to query whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CheckGrantRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CheckGrantResponse
 */
func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *util.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckGrant"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CheckGrant operation to query whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CheckGrantRequest
 * @return CheckGrantResponse
 */
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleAction)) {
		query["RuleAction"] = request.RuleAction
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionCnt)) {
		query["RuleConditionCnt"] = request.RuleConditionCnt
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionKpps)) {
		query["RuleConditionKpps"] = request.RuleConditionKpps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionMbps)) {
		query["RuleConditionMbps"] = request.RuleConditionMbps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSwitch)) {
		query["RuleSwitch"] = request.RuleSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoBeginTime)) {
		query["RuleUndoBeginTime"] = request.RuleUndoBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoEndTime)) {
		query["RuleUndoEndTime"] = request.RuleUndoEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoMode)) {
		query["RuleUndoMode"] = request.RuleUndoMode
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleAction)) {
		query["RuleAction"] = request.RuleAction
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionCnt)) {
		query["RuleConditionCnt"] = request.RuleConditionCnt
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionKpps)) {
		query["RuleConditionKpps"] = request.RuleConditionKpps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleConditionMbps)) {
		query["RuleConditionMbps"] = request.RuleConditionMbps
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSwitch)) {
		query["RuleSwitch"] = request.RuleSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoBeginTime)) {
		query["RuleUndoBeginTime"] = request.RuleUndoBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoEndTime)) {
		query["RuleUndoEndTime"] = request.RuleUndoEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RuleUndoMode)) {
		query["RuleUndoMode"] = request.RuleUndoMode
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

/**
 * You can call the DeleteBlackhole operation to deactivate blackhole filtering for a protected IP address.
 * Before you call this operation, you can call the [DescribePackIpList](~~118701~~) operation to query the protection status of the IP addresses that are protected by your Anti-DDoS Origin instance. For example, you can query whether blackhole filtering is triggered for an IP address.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteBlackholeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteBlackholeResponse
 */
func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *util.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBlackhole"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteBlackhole operation to deactivate blackhole filtering for a protected IP address.
 * Before you call this operation, you can call the [DescribePackIpList](~~118701~~) operation to query the protection status of the IP addresses that are protected by your Anti-DDoS Origin instance. For example, you can query whether blackhole filtering is triggered for an IP address.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteBlackholeRequest
 * @return DeleteBlackholeResponse
 */
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

/**
 * The Anti-DDoS Origin Enterprise instance no longer protects the IP addresses that are removed.
 *
 * @param request DeleteIpRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteIpResponse
 */
func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *util.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IpList)) {
		query["IpList"] = request.IpList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIp"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The Anti-DDoS Origin Enterprise instance no longer protects the IP addresses that are removed.
 *
 * @param request DeleteIpRequest
 * @return DeleteIpResponse
 */
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeAssetGroupWithOptions(request *DescribeAssetGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssetGroup"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssetGroup(request *DescribeAssetGroupRequest) (_result *DescribeAssetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetGroupResponse{}
	_body, _err := client.DescribeAssetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssetGroupToInstanceWithOptions(request *DescribeAssetGroupToInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetGroupToInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssetGroupToInstance"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssetGroupToInstance(request *DescribeAssetGroupToInstanceRequest) (_result *DescribeAssetGroupToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetGroupToInstanceResponse{}
	_body, _err := client.DescribeAssetGroupToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeDdosEvent operation to query the details about the DDoS attack events that occurred on a specific Anti-DDoS Origin instance by page. The details include the start time, end time, attacked IP address, and status of each event.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDdosEventRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDdosEventResponse
 */
func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDdosEvent"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeDdosEvent operation to query the details about the DDoS attack events that occurred on a specific Anti-DDoS Origin instance by page. The details include the start time, end time, attacked IP address, and status of each event.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDdosEventRequest
 * @return DescribeDdosEventResponse
 */
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

/**
 * ## Usage notes
 * You can call the DescribeExcpetionCount operation to query the number of assets that are in an abnormal state and the number of Anti-DDoS Origin instances that are about to expire in a specific region. For example, if blackhole filtering is triggered for an IP address, the IP address is in an abnormal state. An instance whose remaining validity period is less than seven days is considered as an instance that is about to expire.
 *
 * @param request DescribeExcpetionCountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeExcpetionCountResponse
 */
func (client *Client) DescribeExcpetionCountWithOptions(request *DescribeExcpetionCountRequest, runtime *util.RuntimeOptions) (_result *DescribeExcpetionCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExcpetionCount"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Usage notes
 * You can call the DescribeExcpetionCount operation to query the number of assets that are in an abnormal state and the number of Anti-DDoS Origin instances that are about to expire in a specific region. For example, if blackhole filtering is triggered for an IP address, the IP address is in an abnormal state. An instance whose remaining validity period is less than seven days is considered as an instance that is about to expire.
 *
 * @param request DescribeExcpetionCountRequest
 * @return DescribeExcpetionCountResponse
 */
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

/**
 * You can call the DescribeInstanceList operation to query the details of all Anti-DDoS Origin instances within your Alibaba Cloud account by page. The details include the ID, validity period, and status of each instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeInstanceListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeInstanceListResponse
 */
func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypeList)) {
		query["InstanceTypeList"] = request.InstanceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Orderby)) {
		query["Orderby"] = request.Orderby
	}

	if !tea.BoolValue(util.IsUnset(request.Orderdire)) {
		query["Orderdire"] = request.Orderdire
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeInstanceList operation to query the details of all Anti-DDoS Origin instances within your Alibaba Cloud account by page. The details include the ID, validity period, and status of each instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeInstanceListRequest
 * @return DescribeInstanceListResponse
 */
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceSpecs"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

/**
 * >  Anti-DDoS Origin API operations are available for only Anti-DDoS Origin Enterprise users.
 *
 * @param request DescribeOnDemandDdosEventRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeOnDemandDdosEventResponse
 */
func (client *Client) DescribeOnDemandDdosEventWithOptions(request *DescribeOnDemandDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOnDemandDdosEvent"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOnDemandDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  Anti-DDoS Origin API operations are available for only Anti-DDoS Origin Enterprise users.
 *
 * @param request DescribeOnDemandDdosEventRequest
 * @return DescribeOnDemandDdosEventResponse
 */
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOnDemandInstanceStatus"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOnDemandInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

/**
 * The start time. Operation logs that were generated after this time are queried.**** This value is a UNIX timestamp. Unit: milliseconds.
 *
 * @param request DescribeOpEntitiesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeOpEntitiesResponse
 */
func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDir)) {
		query["OrderDir"] = request.OrderDir
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOpEntities"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The start time. Operation logs that were generated after this time are queried.**** This value is a UNIX timestamp. Unit: milliseconds.
 *
 * @param request DescribeOpEntitiesRequest
 * @return DescribeOpEntitiesResponse
 */
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

/**
 * You can call the DescribePackIpList operation to query the details about each IP address that is protected by a specific Anti-DDoS Origin instance by page. The details include the IP address and the type of the cloud asset to which the IP address belongs. The details also include the status of the IP address, such as whether blackhole filtering is triggered for the IP address.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePackIpListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribePackIpListResponse
 */
func (client *Client) DescribePackIpListWithOptions(request *DescribePackIpListRequest, runtime *util.RuntimeOptions) (_result *DescribePackIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		query["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePackIpList"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribePackIpList operation to query the details about each IP address that is protected by a specific Anti-DDoS Origin instance by page. The details include the IP address and the type of the cloud asset to which the IP address belongs. The details also include the status of the IP address, such as whether blackhole filtering is triggered for the IP address.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePackIpListRequest
 * @return DescribePackIpListResponse
 */
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

/**
 * You can call the DescribeTraffic operation to query traffic statistics of an Anti-DDoS Origin instance within a specific time period.
 * >  When you call this operation, you must configure the **InstanceId** parameter to specify the Anti-DDoS Origin instance whose traffic statistics you want to query.
 * ## Limits
 * You can call this operation once per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeTrafficRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeTrafficResponse
 */
func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FlowType)) {
		query["FlowType"] = request.FlowType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Ipnet)) {
		query["Ipnet"] = request.Ipnet
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTraffic"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeTraffic operation to query traffic statistics of an Anti-DDoS Origin instance within a specific time period.
 * >  When you call this operation, you must configure the **InstanceId** parameter to specify the Anti-DDoS Origin instance whose traffic statistics you want to query.
 * ## Limits
 * You can call this operation once per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeTrafficRequest
 * @return DescribeTrafficResponse
 */
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

func (client *Client) DettachAssetGroupToInstanceWithOptions(tmpReq *DettachAssetGroupToInstanceRequest, runtime *util.RuntimeOptions) (_result *DettachAssetGroupToInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DettachAssetGroupToInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AssetGroupList)) {
		request.AssetGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetGroupList, tea.String("AssetGroupList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetGroupListShrink)) {
		query["AssetGroupList"] = request.AssetGroupListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DettachAssetGroupToInstance"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DettachAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DettachAssetGroupToInstance(request *DettachAssetGroupToInstanceRequest) (_result *DettachAssetGroupToInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DettachAssetGroupToInstanceResponse{}
	_body, _err := client.DettachAssetGroupToInstanceWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSlsOpenStatus"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListOpenedAccessLogInstancesWithOptions(request *ListOpenedAccessLogInstancesRequest, runtime *util.RuntimeOptions) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOpenedAccessLogInstances"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

/**
 * You can call the ListTagResources operation to query the tags that are added to Anti-DDoS Origin instances at a time.
 *
 * @param request ListTagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ListTagResources operation to query the tags that are added to Anti-DDoS Origin instances at a time.
 *
 * @param request ListTagResourcesRequest
 * @return ListTagResourcesResponse
 */
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

/**
 * You can call the ModifyRemark operation to add remarks for a single Anti-DDoS Origin instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyRemarkRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyRemarkResponse
 */
func (client *Client) ModifyRemarkWithOptions(request *ModifyRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRemark"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyRemark operation to add remarks for a single Anti-DDoS Origin instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyRemarkRequest
 * @return ModifyRemarkResponse
 */
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySchedruleOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySchedruleOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdList)) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetInstanceModeOnDemand"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetInstanceModeOnDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

/**
 * You can call the TagResources operation to add tags to Anti-DDoS Origin instances.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request TagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the TagResources operation to add tags to Anti-DDoS Origin instances.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2018-07-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
