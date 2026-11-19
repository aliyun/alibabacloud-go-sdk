// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateServiceEndpointRequestBody) *UpdateServiceEndpointRequest
	GetBody() *UpdateServiceEndpointRequestBody
	SetClientToken(v string) *UpdateServiceEndpointRequest
	GetClientToken() *string
}

type UpdateServiceEndpointRequest struct {
	// The service endpoint configurations to be modified in this request. Submit the supported modifiable name, bound version, invocation authentication configuration, elastic policy, or public network inbound switch based on the endpoint type. Router and hybrid ingress endpoints cannot be renamed. DEFAULT endpoints do not accept name or authentication configuration updates. When updating the elastic policy, you must specify both updateElasticity=true and clientToken.
	Body *UpdateServiceEndpointRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token. This parameter is required when updateElasticity is set to true and is verified by the server. Reuse the token and the original request for the same retry. Use a new token if the logical content changes. Passing different content with the same token returns InvalidParameter. The token must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), colons (:), and hyphens (-). The scope is the tenant and workspace. The receipt is retained for 7 days from the time of successful acceptance. Replays do not extend the retention period. After expiration, it is verified again as a new submission.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateServiceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequest) GetBody() *UpdateServiceEndpointRequestBody {
	return s.Body
}

func (s *UpdateServiceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceEndpointRequest) SetBody(v *UpdateServiceEndpointRequestBody) *UpdateServiceEndpointRequest {
	s.Body = v
	return s
}

func (s *UpdateServiceEndpointRequest) SetClientToken(v string) *UpdateServiceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceEndpointRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceEndpointRequestBody struct {
	// If not specified, the original binding is retained. If specified, it must be an exact source version that has been successfully deployed for the agent. Otherwise, ServiceEndpointTargetNotReady is returned. When rebinding a DEFAULT service endpoint of the GATEWAY_API type, ServiceEndpointTargetNotReady is also returned if the original gateway runtime instance is unavailable.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// Updates the invocation authentication configuration saved for the service endpoint, specified by type. If not specified, the original configuration is retained. This field is prohibited for DEFAULT endpoints. For AGENT_VERSION endpoints of non-EDGE agents, the actual API key verification is determined by the current inbound authentication configuration of the agent. The management API for updating the service endpoint still uses AccessKey (AK) or Security Token Service (STS) authentication and Resource Access Management (RAM) authorization.
	Authentication *UpdateServiceEndpointRequestBodyAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// Accepted only when updateElasticity is set to true. Omitting subfields indicates that the original values are retained. The basic minSandboxCount and maxSandboxCount must be submitted in pairs. dynamicPolicy.enabled and policies can be updated separately. If policies is provided, it replaces the entire list. To clear the list, you must ensure that the final dynamicPolicy.enabled is set to false. elasticityPolicy and dynamicPolicy do not accept null, and an empty patch is invalid. Only Router service endpoints and DEFAULT service endpoints of the GATEWAY_API_WITH_VPC_INGRESS or GATEWAY_API_WITH_PUBLIC_INGRESS type support elasticity updates. The server merges, verifies, and saves the complete configuration, and does not backfill basic values with the current capacity. The final basic min and max values must be between 0 and 100. If the historical basic capacity exceeds the limit, you must submit valid basic min and max values even if you only modify the dynamic policy. Silent truncation is not allowed.
	ElasticityPolicy *UpdateServiceEndpointRequestBodyElasticityPolicy `json:"elasticityPolicy,omitempty" xml:"elasticityPolicy,omitempty" type:"Struct"`
	// Used to modify existing named endpoints that support renaming. If omitted, the original name is retained. Router and hybrid ingress endpoints do not support renaming. You can only omit this parameter or pass the original name. This field is prohibited for DEFAULT endpoints.
	//
	// example:
	//
	// production
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to enable public network inbound access for the service endpoint. If not specified, the original value is retained. If set to true or false, the value is replaced. Explicitly passing null or a non-Boolean value returns InvalidParameter, and no restore inheritance semantics are provided. Modifying only this field triggers an ingress revision only and does not redeploy the runtime instance. Whether public network access is actually enabled is also subject to the public network inbound limit of the workspace. For more information, see effectivePublicAccessEnabled in the response.
	PublicAccessEnabled *bool `json:"publicAccessEnabled,omitempty" xml:"publicAccessEnabled,omitempty"`
	// Specifies whether to explicitly update the elastic configuration in this request. If omitted or set to false, the elasticityPolicy carried in this request is ignored, and the saved configuration is retained. If set to true, a valid and non-empty elasticityPolicy patch and clientToken must be provided. This parameter is not persisted, returned, or passed through to the resource manager (RM), and no default value is set. Omit this parameter during normal editing and version rebinding.
	UpdateElasticity *bool `json:"updateElasticity,omitempty" xml:"updateElasticity,omitempty"`
}

func (s UpdateServiceEndpointRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequestBody) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UpdateServiceEndpointRequestBody) GetAuthentication() *UpdateServiceEndpointRequestBodyAuthentication {
	return s.Authentication
}

func (s *UpdateServiceEndpointRequestBody) GetElasticityPolicy() *UpdateServiceEndpointRequestBodyElasticityPolicy {
	return s.ElasticityPolicy
}

func (s *UpdateServiceEndpointRequestBody) GetName() *string {
	return s.Name
}

func (s *UpdateServiceEndpointRequestBody) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *UpdateServiceEndpointRequestBody) GetUpdateElasticity() *bool {
	return s.UpdateElasticity
}

func (s *UpdateServiceEndpointRequestBody) SetAgentVersion(v string) *UpdateServiceEndpointRequestBody {
	s.AgentVersion = &v
	return s
}

func (s *UpdateServiceEndpointRequestBody) SetAuthentication(v *UpdateServiceEndpointRequestBodyAuthentication) *UpdateServiceEndpointRequestBody {
	s.Authentication = v
	return s
}

func (s *UpdateServiceEndpointRequestBody) SetElasticityPolicy(v *UpdateServiceEndpointRequestBodyElasticityPolicy) *UpdateServiceEndpointRequestBody {
	s.ElasticityPolicy = v
	return s
}

func (s *UpdateServiceEndpointRequestBody) SetName(v string) *UpdateServiceEndpointRequestBody {
	s.Name = &v
	return s
}

func (s *UpdateServiceEndpointRequestBody) SetPublicAccessEnabled(v bool) *UpdateServiceEndpointRequestBody {
	s.PublicAccessEnabled = &v
	return s
}

func (s *UpdateServiceEndpointRequestBody) SetUpdateElasticity(v bool) *UpdateServiceEndpointRequestBody {
	s.UpdateElasticity = &v
	return s
}

func (s *UpdateServiceEndpointRequestBody) Validate() error {
	if s.Authentication != nil {
		if err := s.Authentication.Validate(); err != nil {
			return err
		}
	}
	if s.ElasticityPolicy != nil {
		if err := s.ElasticityPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceEndpointRequestBodyAuthentication struct {
	// The saved invocation authentication type. Valid values:
	//
	// - NONE: Configured to not use the endpoint API key.
	//
	// - API_KEY: Configured to use the endpoint API key.
	//
	// For AGENT_VERSION endpoints of non-EDGE agents, the actual API key verification is determined by the current inbound authentication configuration of the agent and cannot be determined solely based on this field. You must provide type when passing authentication.
	//
	// This parameter is required.
	//
	// example:
	//
	// API_KEY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateServiceEndpointRequestBodyAuthentication) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequestBodyAuthentication) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequestBodyAuthentication) GetType() *string {
	return s.Type
}

func (s *UpdateServiceEndpointRequestBodyAuthentication) SetType(v string) *UpdateServiceEndpointRequestBodyAuthentication {
	s.Type = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyAuthentication) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceEndpointRequestBodyElasticityPolicy struct {
	// The dynamic policy configuration. It must be fully provided during creation. If omitted during an update, the original configuration is retained. If provided, it is merged by field and does not accept null. The response returns the saved complete configuration. If no historical dynamic configuration exists, it is normalized to enabled=false and policies=[].
	DynamicPolicy *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy `json:"dynamicPolicy,omitempty" xml:"dynamicPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the overall automatic elasticity mode. This parameter is independent of dynamicPolicy.enabled. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The saved basic maximum number of instances. Valid values: 0 to 100. The value cannot be less than minSandboxCount. The basic upper and lower limits are used when dynamic elasticity is disabled or no candidates are available.
	//
	// example:
	//
	// 5
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The saved basic minimum number of instances. Valid values: 0 to 100. The value cannot be greater than maxSandboxCount. It must be submitted in pairs with maxSandboxCount during an update. If both are omitted, the original values are retained. It does not override dynamic results when dynamic candidates are available.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) GetDynamicPolicy() *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy {
	return s.DynamicPolicy
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) SetDynamicPolicy(v *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) *UpdateServiceEndpointRequestBodyElasticityPolicy {
	s.DynamicPolicy = v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) SetEnabled(v bool) *UpdateServiceEndpointRequestBodyElasticityPolicy {
	s.Enabled = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) SetMaxSandboxCount(v int32) *UpdateServiceEndpointRequestBodyElasticityPolicy {
	s.MaxSandboxCount = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) SetMinSandboxCount(v int32) *UpdateServiceEndpointRequestBodyElasticityPolicy {
	s.MinSandboxCount = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicy) Validate() error {
	if s.DynamicPolicy != nil {
		if err := s.DynamicPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy struct {
	// Specifies whether to enable dynamic policies. This parameter is required during creation. If omitted during an update, the original value is retained. Disabling it retains existing policies and restores the basic upper and lower limits. Re-enabling it triggers immediate calculation. When it is finally enabled, policies cannot be empty.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The complete list of dynamic policies. This parameter is required during creation. If omitted during an update, the original list is retained. If provided, it replaces the entire list and is not merged by ID. To clear the list, you must ensure that the final enabled is set to false. Each item must be complete. Conditional requirements and cross-field constraints are verified by the server. The display order does not indicate priority.
	Policies []*UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Repeated"`
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) GetPolicies() []*UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	return s.Policies
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) SetEnabled(v bool) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy {
	s.Enabled = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) SetPolicies(v []*UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy {
	s.Policies = v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies struct {
	// The end time of the validity period. This parameter is required only when validity is set to range. The format is the same as startTime, and the seconds are fixed at 00. It must be later than startTime. The interval does not include the end time. If validity is set to forever, the start and end times must be omitted. Valid dates and cross-field constraints are verified by the console.
	//
	// example:
	//
	// 2026-12-01T00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The policy identifier. This parameter is required and must be unique and stable within the endpoint. It is generated by the frontend when added, and retained during editing and retries.
	//
	// example:
	//
	// policy-weekday
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The policy priority. This parameter is required. Valid values: 5 (regular) and 10 (high priority). The default selection for a new policy is regular. You must explicitly provide this parameter upon submission, and no default value is set for the API. When the final dynamicPolicy.enabled is set to true, the validity periods of high-priority policies cannot overlap, but end-to-end connection is allowed. Long-term high-priority policies can overlap with other high-priority policies. When disabled, overlapping drafts can be saved, and the final configuration is verified when re-enabled. The validity of other fields and trigger conflicts within the policy are always verified. You must explicitly provide priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The repeat period. This parameter is required. Valid values: daily (every day), weekly (every week), and monthly (every month). Only submit the fields corresponding to the specified period.
	//
	// example:
	//
	// weekly
	Repeat *string `json:"repeat,omitempty" xml:"repeat,omitempty"`
	// The start time of the validity period. This parameter is required only when validity is set to range. Format: YYYY-MM-DDTHH:mm:ss in the policy time zone, without a time zone offset. The seconds are fixed at 00. The interval includes the start time. When editing, existing past start times are retained, and updates or retries are not rejected because the start time is earlier than the current time.
	//
	// example:
	//
	// 2026-11-01T00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The periodic trigger time points. This parameter is required and cannot be empty. Each item must be complete. The day of the week or day of the month is configured independently for each item. Within the same policy, daily policies conflict if the time is the same. weekly and monthly policies conflict only when the time is the same and the date sets intersect. This is verified by the console.
	TimePoints []*UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints `json:"timePoints,omitempty" xml:"timePoints,omitempty" type:"Repeated"`
	// The policy time zone. This parameter is required and must be a valid IANA time zone. All dates and times in this policy are interpreted in this time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// The policy validity period type. This parameter is required. Valid values: forever (long-term validity) and range (short-term validity). After expiration, the policy and its trigger results exit the current calculation, and the remaining policies are recalculated. If no candidates are available, the basic values are restored. The saved configuration is not automatically deleted.
	//
	// example:
	//
	// range
	Validity *string `json:"validity,omitempty" xml:"validity,omitempty"`
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetId() *string {
	return s.Id
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetRepeat() *string {
	return s.Repeat
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetTimePoints() []*UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	return s.TimePoints
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetTimeZone() *string {
	return s.TimeZone
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetValidity() *string {
	return s.Validity
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetEndTime(v string) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.EndTime = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetId(v string) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Id = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetPriority(v int32) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Priority = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetRepeat(v string) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Repeat = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetStartTime(v string) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.StartTime = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetTimePoints(v []*UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.TimePoints = v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetTimeZone(v string) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.TimeZone = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetValidity(v string) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Validity = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) Validate() error {
	if s.TimePoints != nil {
		for _, item := range s.TimePoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints struct {
	// The target maximum number of instances. This parameter is required. Valid values: 1 to 100.
	//
	// example:
	//
	// 5
	MaxInstances *int32 `json:"maxInstances,omitempty" xml:"maxInstances,omitempty"`
	// The target minimum number of instances. This parameter is required. Valid values: 0 to 100. The value cannot be greater than maxInstances in the same item.
	//
	// example:
	//
	// 1
	MinInstances *int32 `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	// The days of the month. This parameter is required only for each trigger point of monthly policies. It cannot be empty and must be deduplicated. Valid values are integers from 1 to 31. Each row is selected independently. If a date does not exist in a specific month, the policy is not triggered and is not postponed to the end of the month. This parameter must be omitted for other periods.
	MonthDays []*int32 `json:"monthDays,omitempty" xml:"monthDays,omitempty" type:"Repeated"`
	// The trigger time. This parameter is required. Format: HH:mm:ss. It is interpreted in the time zone of the associated policy, accurate to the minute, and the seconds are fixed at 00.
	//
	// example:
	//
	// 09:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The days of the week. This parameter is required only for each trigger point of weekly policies. It cannot be empty and must be deduplicated. Valid values are integers from 1 to 7, where 1 represents Monday and 7 represents Sunday. Each row is selected independently. This parameter must be omitted for other periods.
	Weekdays []*int32 `json:"weekdays,omitempty" xml:"weekdays,omitempty" type:"Repeated"`
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMaxInstances() *int32 {
	return s.MaxInstances
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMonthDays() []*int32 {
	return s.MonthDays
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetTime() *string {
	return s.Time
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetWeekdays() []*int32 {
	return s.Weekdays
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMaxInstances(v int32) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MaxInstances = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMinInstances(v int32) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MinInstances = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMonthDays(v []*int32) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MonthDays = v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetTime(v string) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Time = &v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetWeekdays(v []*int32) *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Weekdays = v
	return s
}

func (s *UpdateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) Validate() error {
	return dara.Validate(s)
}
