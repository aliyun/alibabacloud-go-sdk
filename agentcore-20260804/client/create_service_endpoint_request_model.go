// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateServiceEndpointRequestBody) *CreateServiceEndpointRequest
	GetBody() *CreateServiceEndpointRequestBody
	SetClientToken(v string) *CreateServiceEndpointRequest
	GetClientToken() *string
}

type CreateServiceEndpointRequest struct {
	// The name, routing target, and invocation authentication configuration of the service endpoint to be created, along with the optional initial elasticity configuration and public access switch.
	Body *CreateServiceEndpointRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Required when elasticity initialization is included. It is verified by the server. The same token and original request are reused for the same retry. If the logical content changes, a new token is used. The same token with different content returns InvalidParameter. The token must be 1 to 128 characters in length and can contain letters, digits, and the following special characters: . _ : -. The scope is the tenant and workspace. The receipt is retained for 7 days from the time of successful acceptance. Replays do not extend the retention period. After expiration, it is verified again as a new submission.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateServiceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequest) GetBody() *CreateServiceEndpointRequestBody {
	return s.Body
}

func (s *CreateServiceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceEndpointRequest) SetBody(v *CreateServiceEndpointRequestBody) *CreateServiceEndpointRequest {
	s.Body = v
	return s
}

func (s *CreateServiceEndpointRequest) SetClientToken(v string) *CreateServiceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceEndpointRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceEndpointRequestBody struct {
	// Sets the invocation authentication configuration saved for the service endpoint, specified by type. For AGENT_VERSION endpoints of non-EDGE agents, the actual API key verification is determined by the current inbound authentication configuration of the agent. The management operation for creating service endpoints still uses AccessKey or STS authentication and RAM authorization.
	//
	// This parameter is required.
	Authentication *CreateServiceEndpointRequestBodyAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// Explicitly provided when creating a service endpoint with elasticity capabilities. The capability scope follows existing rules. Null or partial patches are not accepted. It must include enabled, the basic minSandboxCount and maxSandboxCount, and the complete dynamicPolicy. The basic upper and lower limits are both 0 to 100, and min must be less than or equal to max. To disable dynamic scaling without policies, pass enabled=false and policies=[]. When elasticity initialization is included, clientToken must be provided, and updateElasticity is not passed. Conditional requirements and cross-field constraints are verified by the server.
	ElasticityPolicy *CreateServiceEndpointRequestBodyElasticityPolicy `json:"elasticityPolicy,omitempty" xml:"elasticityPolicy,omitempty" type:"Struct"`
	// The name of the named service endpoint to be created. For the Router type, the name must be 1 to 128 characters in length and can contain lowercase letters, digits, underscores, and hyphens. It must start with a letter or a digit, cannot use the reserved name default, and must be unique within the same agent. The names of Router and hybrid ingress endpoints cannot be changed after creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// production
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to enable public network inbound access for the service endpoint. If omitted, it is initialized to true. If true or false is passed, the passed value is saved. An explicit null or non-boolean value returns InvalidParameter. Whether public network access is actually enabled is also subject to the public network inbound access limit of the workspace. For more information, see effectivePublicAccessEnabled in the response.
	PublicAccessEnabled *bool `json:"publicAccessEnabled,omitempty" xml:"publicAccessEnabled,omitempty"`
	// Specifies the agent and its version associated with the service endpoint. The routing targets for collaboration components are maintained internally by the platform.
	//
	// This parameter is required.
	Target *CreateServiceEndpointRequestBodyTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s CreateServiceEndpointRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequestBody) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequestBody) GetAuthentication() *CreateServiceEndpointRequestBodyAuthentication {
	return s.Authentication
}

func (s *CreateServiceEndpointRequestBody) GetElasticityPolicy() *CreateServiceEndpointRequestBodyElasticityPolicy {
	return s.ElasticityPolicy
}

func (s *CreateServiceEndpointRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateServiceEndpointRequestBody) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *CreateServiceEndpointRequestBody) GetTarget() *CreateServiceEndpointRequestBodyTarget {
	return s.Target
}

func (s *CreateServiceEndpointRequestBody) SetAuthentication(v *CreateServiceEndpointRequestBodyAuthentication) *CreateServiceEndpointRequestBody {
	s.Authentication = v
	return s
}

func (s *CreateServiceEndpointRequestBody) SetElasticityPolicy(v *CreateServiceEndpointRequestBodyElasticityPolicy) *CreateServiceEndpointRequestBody {
	s.ElasticityPolicy = v
	return s
}

func (s *CreateServiceEndpointRequestBody) SetName(v string) *CreateServiceEndpointRequestBody {
	s.Name = &v
	return s
}

func (s *CreateServiceEndpointRequestBody) SetPublicAccessEnabled(v bool) *CreateServiceEndpointRequestBody {
	s.PublicAccessEnabled = &v
	return s
}

func (s *CreateServiceEndpointRequestBody) SetTarget(v *CreateServiceEndpointRequestBodyTarget) *CreateServiceEndpointRequestBody {
	s.Target = v
	return s
}

func (s *CreateServiceEndpointRequestBody) Validate() error {
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
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceEndpointRequestBodyAuthentication struct {
	// The saved invocation authentication type. NONE indicates that the endpoint API key is not used. API_KEY indicates that the endpoint API key is used. For AGENT_VERSION endpoints of non-EDGE agents, the actual API key verification is determined by the current inbound authentication configuration of the agent and cannot be determined solely based on this field.
	//
	// This parameter is required.
	//
	// example:
	//
	// API_KEY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateServiceEndpointRequestBodyAuthentication) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequestBodyAuthentication) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequestBodyAuthentication) GetType() *string {
	return s.Type
}

func (s *CreateServiceEndpointRequestBodyAuthentication) SetType(v string) *CreateServiceEndpointRequestBodyAuthentication {
	s.Type = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyAuthentication) Validate() error {
	return dara.Validate(s)
}

type CreateServiceEndpointRequestBodyElasticityPolicy struct {
	// Required. It must include enabled and policies, and is verified by the server. It is independent of the overall automatic elasticity switch.
	DynamicPolicy *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy `json:"dynamicPolicy,omitempty" xml:"dynamicPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the entire automatic elasticity mode. It must be provided when elasticityPolicy is passed and is verified by the server.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Specifies whether to enable the entire automatic elasticity mode. It must be provided when elasticityPolicy is passed and is verified by the server.
	//
	// example:
	//
	// 5
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// Specifies whether to enable the entire automatic elasticity mode. It must be provided when elasticityPolicy is passed and is verified by the server.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
}

func (s CreateServiceEndpointRequestBodyElasticityPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequestBodyElasticityPolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) GetDynamicPolicy() *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy {
	return s.DynamicPolicy
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) SetDynamicPolicy(v *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) *CreateServiceEndpointRequestBodyElasticityPolicy {
	s.DynamicPolicy = v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) SetEnabled(v bool) *CreateServiceEndpointRequestBodyElasticityPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) SetMaxSandboxCount(v int32) *CreateServiceEndpointRequestBodyElasticityPolicy {
	s.MaxSandboxCount = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) SetMinSandboxCount(v int32) *CreateServiceEndpointRequestBodyElasticityPolicy {
	s.MinSandboxCount = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicy) Validate() error {
	if s.DynamicPolicy != nil {
		if err := s.DynamicPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy struct {
	// Specifies whether to enable dynamic policies. Required during creation. If omitted during an update, the original value is retained. Disabling it retains existing policies and restores the basic upper and lower limits. Re-enabling it triggers immediate calculation. When finally enabled, policies must not be empty.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Required during creation. If omitted during an update, the original list is retained. If passed, the entire list is replaced and not merged by ID. To clear the list, you must also ensure that the final enabled value is false. Each item must be complete. Conditional requirements and cross-field constraints are verified by the server. The display order does not indicate priority.
	Policies []*CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Repeated"`
}

func (s CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) GetPolicies() []*CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	return s.Policies
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) SetEnabled(v bool) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) SetPolicies(v []*CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy {
	s.Policies = v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicy) Validate() error {
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

type CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies struct {
	// Required only when validity is set to range. The format is the same as startTime, the seconds are fixed at 00, and it must be later than startTime. The interval does not include the end time. For forever, the start and end times must be omitted. Valid dates and cross-field constraints are verified by the console.
	//
	// example:
	//
	// 2026-12-01T00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Required. It must be unique and stable within the endpoint. It is generated by the frontend when added, and retained during editing and retries.
	//
	// example:
	//
	// policy-weekday
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Required. 5 indicates normal, and 10 indicates high priority. New policies default to normal. It must be explicitly provided during submission, and no default value is set for the operation. When the final dynamicPolicy.enabled is true, the validity periods of high-priority policies must not overlap, but end-to-end connection is allowed. Long-term high-priority policies can overlap with other high-priority policies. When disabled, overlapping drafts can be saved, and the final configuration is verified when re-enabled. The validity of other fields and trigger conflicts within the policy are always verified. The priority must be explicitly provided.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Required. daily means every day, weekly means every week, and monthly means every month. Only submit the fields corresponding to the cycle.
	//
	// example:
	//
	// weekly
	Repeat *string `json:"repeat,omitempty" xml:"repeat,omitempty"`
	// Required only when validity is set to range. The format is YYYY-MM-DDTHH:mm:ss in the policy time zone, without a time zone offset, and the seconds are fixed at 00. The interval includes the start time. Editing retains the existing past start time and does not reject updates or retries because it is earlier than the current time.
	//
	// example:
	//
	// 2026-11-01T00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Required and non-empty. Each item must be complete. The day of the week or day of the month is configured independently for each item. Within the same policy, daily policies conflict if the time is the same. Weekly and monthly policies conflict only if the time is the same and the date sets intersect. This is verified by the console.
	TimePoints []*CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints `json:"timePoints,omitempty" xml:"timePoints,omitempty" type:"Repeated"`
	// Required. A valid IANA time zone. All dates and times in this policy are interpreted in this time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// Required. forever indicates long-term validity, and range indicates short-term validity. After expiration, the policy and its trigger results exit the current calculation, and the remaining policies are recalculated. If there are no candidates, the basic values are restored, and the saved configurations are not automatically deleted.
	//
	// example:
	//
	// range
	Validity *string `json:"validity,omitempty" xml:"validity,omitempty"`
}

func (s CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetId() *string {
	return s.Id
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetRepeat() *string {
	return s.Repeat
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetTimePoints() []*CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	return s.TimePoints
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) GetValidity() *string {
	return s.Validity
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetEndTime(v string) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.EndTime = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetId(v string) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Id = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetPriority(v int32) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Priority = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetRepeat(v string) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Repeat = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetStartTime(v string) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.StartTime = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetTimePoints(v []*CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.TimePoints = v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetTimeZone(v string) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.TimeZone = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) SetValidity(v string) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies {
	s.Validity = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPolicies) Validate() error {
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

type CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints struct {
	// Required. The value ranges from 1 to 100.
	//
	// example:
	//
	// 5
	MaxInstances *int32 `json:"maxInstances,omitempty" xml:"maxInstances,omitempty"`
	// Required. The value ranges from 0 to 100 and must not be greater than maxInstances in the same item.
	//
	// example:
	//
	// 1
	MinInstances *int32 `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	// Required only for each trigger point of monthly. It must be non-empty and deduplicated. The values are integers from 1 to 31. Each row is selected independently. If a date does not exist in a month, it is not triggered and is not postponed to the end of the month. It must be omitted for other cycles.
	MonthDays []*int32 `json:"monthDays,omitempty" xml:"monthDays,omitempty" type:"Repeated"`
	// Required. The format is HH:mm:ss. It is interpreted in the time zone of the policy, accurate to the minute, and the seconds are fixed at 00.
	//
	// example:
	//
	// 09:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// Required only for each trigger point of weekly. It must be non-empty and deduplicated. The values are integers from 1 to 7, where 1 indicates Monday and 7 indicates Sunday. Each row is selected independently. It must be omitted for other cycles.
	Weekdays []*int32 `json:"weekdays,omitempty" xml:"weekdays,omitempty" type:"Repeated"`
}

func (s CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMaxInstances() *int32 {
	return s.MaxInstances
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMonthDays() []*int32 {
	return s.MonthDays
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetTime() *string {
	return s.Time
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) GetWeekdays() []*int32 {
	return s.Weekdays
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMaxInstances(v int32) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MaxInstances = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMinInstances(v int32) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MinInstances = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMonthDays(v []*int32) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MonthDays = v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetTime(v string) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Time = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) SetWeekdays(v []*int32) *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Weekdays = v
	return s
}

func (s *CreateServiceEndpointRequestBodyElasticityPolicyDynamicPolicyPoliciesTimePoints) Validate() error {
	return dara.Validate(s)
}

type CreateServiceEndpointRequestBodyTarget struct {
	// The ID of the agent associated with the AGENT_VERSION routing target. The agent is located in the workspace specified in the request.
	//
	// example:
	//
	// agent-0123456789abcdef0123456789abcdef
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The version of the agent associated with the AGENT_VERSION routing target. Use the agentVersion returned by CreateAgentVersion or ListAgentVersions, which is different from the version name versionName.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// The component associated with the TEAM_COLLABORATION routing target. MATRIX_CLIENT indicates the Matrix client ingress, MATRIX_FEDERATION indicates the Matrix federation ingress, and ELEMENT_WEB indicates the Element Web ingress. Service endpoints for collaboration components are created internally by the platform. Do not pass this field when publicly creating an AGENT_VERSION endpoint.
	//
	// example:
	//
	// MATRIX_CLIENT
	CollaborationComponent *string `json:"collaborationComponent,omitempty" xml:"collaborationComponent,omitempty"`
	// The ID of the resource binding associated with the TEAM_COLLABORATION routing target. It is maintained internally by the platform and generated in the format of wrb- followed by a 32-character hexadecimal UUID. Do not pass this field when publicly creating an AGENT_VERSION endpoint.
	//
	// example:
	//
	// wrb-0123456789abcdef0123456789abcdef
	ResourceBindingId *string `json:"resourceBindingId,omitempty" xml:"resourceBindingId,omitempty"`
	// The public create operation only supports AGENT_VERSION, which indicates routing to a specified agent version. TEAM_COLLABORATION indicates a collaboration component. Service endpoints for collaboration components are created internally by the platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// AGENT_VERSION
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s CreateServiceEndpointRequestBodyTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequestBodyTarget) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequestBodyTarget) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateServiceEndpointRequestBodyTarget) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *CreateServiceEndpointRequestBodyTarget) GetCollaborationComponent() *string {
	return s.CollaborationComponent
}

func (s *CreateServiceEndpointRequestBodyTarget) GetResourceBindingId() *string {
	return s.ResourceBindingId
}

func (s *CreateServiceEndpointRequestBodyTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateServiceEndpointRequestBodyTarget) SetAgentId(v string) *CreateServiceEndpointRequestBodyTarget {
	s.AgentId = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyTarget) SetAgentVersion(v string) *CreateServiceEndpointRequestBodyTarget {
	s.AgentVersion = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyTarget) SetCollaborationComponent(v string) *CreateServiceEndpointRequestBodyTarget {
	s.CollaborationComponent = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyTarget) SetResourceBindingId(v string) *CreateServiceEndpointRequestBodyTarget {
	s.ResourceBindingId = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyTarget) SetTargetType(v string) *CreateServiceEndpointRequestBodyTarget {
	s.TargetType = &v
	return s
}

func (s *CreateServiceEndpointRequestBodyTarget) Validate() error {
	return dara.Validate(s)
}
