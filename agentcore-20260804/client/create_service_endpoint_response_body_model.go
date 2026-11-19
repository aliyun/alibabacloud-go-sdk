// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceEndpointResponseBody
	GetCode() *string
	SetData(v *CreateServiceEndpointResponseBodyData) *CreateServiceEndpointResponseBody
	GetData() *CreateServiceEndpointResponseBodyData
	SetHttpStatusCode(v int32) *CreateServiceEndpointResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateServiceEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceEndpointResponseBody
	GetSuccess() *bool
}

type CreateServiceEndpointResponseBody struct {
	// The response code of the request. The value is SUCCESS if the request is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The service endpoint information returned after a successful creation, including the binding target, saved authentication and elasticity configurations, and current status. The endpoint may still be in the creating state.
	Data *CreateServiceEndpointResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the request. The value is 200 if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message of the request. The value is success if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request, used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5E8F3C2A-1B4D-4F6E-9A7C-0D2B8E6F4A1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates that the creation request is accepted, but it does not mean that the endpoint deployment is complete. To confirm the deployment result, check data.status and data.accessUrls.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateServiceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceEndpointResponseBody) GetData() *CreateServiceEndpointResponseBodyData {
	return s.Data
}

func (s *CreateServiceEndpointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateServiceEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceEndpointResponseBody) SetCode(v string) *CreateServiceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetData(v *CreateServiceEndpointResponseBodyData) *CreateServiceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetHttpStatusCode(v int32) *CreateServiceEndpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetMessage(v string) *CreateServiceEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetRequestId(v string) *CreateServiceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) SetSuccess(v bool) *CreateServiceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceEndpointResponseBodyData struct {
	// The list of access URLs for the service endpoint.
	AccessUrls []*CreateServiceEndpointResponseBodyDataAccessUrls `json:"accessUrls,omitempty" xml:"accessUrls,omitempty" type:"Repeated"`
	// The authentication configuration of the service endpoint.
	Authentication *CreateServiceEndpointResponseBodyDataAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// The creation time in UTC format compliant with RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The current number of instances. This parameter is read-only. The Get operation queries RM sandboxPhaseCounts.RUNNING based on the endpoint runtimeKey. The Create and Update operations return null. The value is a non-negative integer. If the statistics are unavailable, null is returned. This parameter is not replaced by the base upper and lower limits, target quantity, or READY status. Unavailable statistics do not block submission.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2
	CurrentInstances *int32 `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
	// The combined result of publicAccessEnabled and the workspace-level public access limit. The value is true only when both allow public access. This value only indicates the policy target and does not mean that the public endpoint has been fully provisioned. For the actual status, refer to accessUrls and status.
	EffectivePublicAccessEnabled *bool `json:"effectivePublicAccessEnabled,omitempty" xml:"effectivePublicAccessEnabled,omitempty"`
	// The configuration used for edit backfill, which is not the currently effective upper and lower limits or the actual number of instances. The Get operation preferentially returns the saved configuration. If no configuration is saved, an endpoint that supports elasticity backfills the base values from the expected runtime frozen configuration or the bound version configuration when the runtime is not yet established, without requiring a successful publication. The value is null when the source is temporarily unavailable or no applicable configuration exists, but the status and statusReason are still returned. The Create and Update operations return only the saved configuration, and the value is null if no configuration is saved. The historical base capacity may exceed 100, and the read value is not truncated.
	//
	// if can be null:
	// true
	ElasticityPolicy *CreateServiceEndpointResponseBodyDataElasticityPolicy `json:"elasticityPolicy,omitempty" xml:"elasticityPolicy,omitempty" type:"Struct"`
	// The type of the service endpoint. Valid values:
	//
	// - DEFAULT: The default endpoint created and maintained by the platform.
	//
	// - NAMED: The named endpoint explicitly created by the user.
	//
	// example:
	//
	// NAMED
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The name of the service endpoint. The name must be unique within the workspace and be 1 to 128 characters in length.
	//
	// example:
	//
	// my-agent-endpoint
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The expected public access setting configured by the user on the service endpoint. For historical service endpoints that have not been configured, this parameter returns true. This value only indicates the configuration expectation. To check whether the actual public endpoint is available, refer to accessUrls and status.
	PublicAccessEnabled *bool `json:"publicAccessEnabled,omitempty" xml:"publicAccessEnabled,omitempty"`
	// The ID of the region where the service endpoint resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the service endpoint.
	//
	// example:
	//
	// se-123456
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
	// The status of the service endpoint. Valid values: CREATING, READY, UPDATING, DEGRADED, DISABLED, and DELETING.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the service endpoint status. The specific reason is returned when the status is abnormal.
	//
	// example:
	//
	// ServiceEndpoint.Provider.Unavailable: no provider for the target type
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The target routing configuration of the service endpoint.
	Target *CreateServiceEndpointResponseBodyDataTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// The time of the last modification in UTC format compliant with RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateServiceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyData) GetAccessUrls() []*CreateServiceEndpointResponseBodyDataAccessUrls {
	return s.AccessUrls
}

func (s *CreateServiceEndpointResponseBodyData) GetAuthentication() *CreateServiceEndpointResponseBodyDataAuthentication {
	return s.Authentication
}

func (s *CreateServiceEndpointResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateServiceEndpointResponseBodyData) GetCurrentInstances() *int32 {
	return s.CurrentInstances
}

func (s *CreateServiceEndpointResponseBodyData) GetEffectivePublicAccessEnabled() *bool {
	return s.EffectivePublicAccessEnabled
}

func (s *CreateServiceEndpointResponseBodyData) GetElasticityPolicy() *CreateServiceEndpointResponseBodyDataElasticityPolicy {
	return s.ElasticityPolicy
}

func (s *CreateServiceEndpointResponseBodyData) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateServiceEndpointResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateServiceEndpointResponseBodyData) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *CreateServiceEndpointResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceEndpointResponseBodyData) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *CreateServiceEndpointResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceEndpointResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CreateServiceEndpointResponseBodyData) GetTarget() *CreateServiceEndpointResponseBodyDataTarget {
	return s.Target
}

func (s *CreateServiceEndpointResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateServiceEndpointResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateServiceEndpointResponseBodyData) SetAccessUrls(v []*CreateServiceEndpointResponseBodyDataAccessUrls) *CreateServiceEndpointResponseBodyData {
	s.AccessUrls = v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetAuthentication(v *CreateServiceEndpointResponseBodyDataAuthentication) *CreateServiceEndpointResponseBodyData {
	s.Authentication = v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetCreatedAt(v string) *CreateServiceEndpointResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetCurrentInstances(v int32) *CreateServiceEndpointResponseBodyData {
	s.CurrentInstances = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetEffectivePublicAccessEnabled(v bool) *CreateServiceEndpointResponseBodyData {
	s.EffectivePublicAccessEnabled = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetElasticityPolicy(v *CreateServiceEndpointResponseBodyDataElasticityPolicy) *CreateServiceEndpointResponseBodyData {
	s.ElasticityPolicy = v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetEndpointType(v string) *CreateServiceEndpointResponseBodyData {
	s.EndpointType = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetName(v string) *CreateServiceEndpointResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetPublicAccessEnabled(v bool) *CreateServiceEndpointResponseBodyData {
	s.PublicAccessEnabled = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetRegionId(v string) *CreateServiceEndpointResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetServiceEndpointId(v string) *CreateServiceEndpointResponseBodyData {
	s.ServiceEndpointId = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetStatus(v string) *CreateServiceEndpointResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetStatusReason(v string) *CreateServiceEndpointResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetTarget(v *CreateServiceEndpointResponseBodyDataTarget) *CreateServiceEndpointResponseBodyData {
	s.Target = v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetUpdatedAt(v string) *CreateServiceEndpointResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) SetWorkspaceId(v string) *CreateServiceEndpointResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyData) Validate() error {
	if s.AccessUrls != nil {
		for _, item := range s.AccessUrls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateServiceEndpointResponseBodyDataAccessUrls struct {
	// The type of the access URL. Valid values: INTERNET and VPC.
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
	// The status of the access URL. Valid values: CREATING, READY, and DEGRADED.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the access URL status. The specific reason is returned when the status is abnormal.
	//
	// example:
	//
	// ServiceEndpoint.Provider.Unavailable: no provider for the target type
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The access URL.
	//
	// example:
	//
	// https://endpoint.example.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateServiceEndpointResponseBodyDataAccessUrls) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyDataAccessUrls) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) GetAccessType() *string {
	return s.AccessType
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) GetStatus() *string {
	return s.Status
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) GetUrl() *string {
	return s.Url
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) SetAccessType(v string) *CreateServiceEndpointResponseBodyDataAccessUrls {
	s.AccessType = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) SetStatus(v string) *CreateServiceEndpointResponseBodyDataAccessUrls {
	s.Status = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) SetStatusReason(v string) *CreateServiceEndpointResponseBodyDataAccessUrls {
	s.StatusReason = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) SetUrl(v string) *CreateServiceEndpointResponseBodyDataAccessUrls {
	s.Url = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataAccessUrls) Validate() error {
	return dara.Validate(s)
}

type CreateServiceEndpointResponseBodyDataAuthentication struct {
	// The authentication method. Valid values:
	//
	// - NONE: No authentication is required.
	//
	// - API_KEY: Authentication is performed by passing an API key in the x-api-key request header.
	//
	// example:
	//
	// API_KEY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateServiceEndpointResponseBodyDataAuthentication) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyDataAuthentication) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyDataAuthentication) GetType() *string {
	return s.Type
}

func (s *CreateServiceEndpointResponseBodyDataAuthentication) SetType(v string) *CreateServiceEndpointResponseBodyDataAuthentication {
	s.Type = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataAuthentication) Validate() error {
	return dara.Validate(s)
}

type CreateServiceEndpointResponseBodyDataElasticityPolicy struct {
	// The saved dynamic configuration. If the base configuration before the requirement goes online lacks dynamicPolicy, it is supplemented with enabled=false and policies=[].
	DynamicPolicy *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy `json:"dynamicPolicy,omitempty" xml:"dynamicPolicy,omitempty" type:"Struct"`
	// Indicates whether the elasticity policy is enabled. This parameter is independent of dynamicPolicy.enabled. If this parameter is omitted during an update, the original value is retained.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The saved or backfilled base maximum number of instances, which is not less than minSandboxCount. The historical value may exceed 100, and the read value is not truncated. The base value is used when dynamic scaling is disabled or no candidates are available.
	//
	// example:
	//
	// 5
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The saved or backfilled base minimum number of instances, which is a non-negative integer and not greater than maxSandboxCount. The historical value may exceed 100, and the read value is not truncated. A newly written complete configuration must still meet the requirement of 0 to 100.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
}

func (s CreateServiceEndpointResponseBodyDataElasticityPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyDataElasticityPolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) GetDynamicPolicy() *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	return s.DynamicPolicy
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) SetDynamicPolicy(v *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) *CreateServiceEndpointResponseBodyDataElasticityPolicy {
	s.DynamicPolicy = v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) SetEnabled(v bool) *CreateServiceEndpointResponseBodyDataElasticityPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) SetMaxSandboxCount(v int32) *CreateServiceEndpointResponseBodyDataElasticityPolicy {
	s.MaxSandboxCount = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) SetMinSandboxCount(v int32) *CreateServiceEndpointResponseBodyDataElasticityPolicy {
	s.MinSandboxCount = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicy) Validate() error {
	if s.DynamicPolicy != nil {
		if err := s.DynamicPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy struct {
	// Specifies whether to enable dynamic policies. Required during creation. If omitted during an update, the original value is retained. Disabling it retains existing policies and restores the basic upper and lower limits. Re-enabling it triggers immediate calculation. When finally enabled, policies must not be empty.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Required during creation. If omitted during an update, the original list is retained. If passed, the entire list is replaced and not merged by ID. To clear the list, you must also ensure that the final enabled value is false. Each item must be complete. Conditional requirements and cross-field constraints are verified by the server. The display order does not indicate priority.
	Policies []*CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Repeated"`
}

func (s CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GetPolicies() []*CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	return s.Policies
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) SetEnabled(v bool) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	s.Enabled = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) SetPolicies(v []*CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	s.Policies = v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) Validate() error {
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

type CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies struct {
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
	TimePoints []*CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints `json:"timePoints,omitempty" xml:"timePoints,omitempty" type:"Repeated"`
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

func (s CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetId() *string {
	return s.Id
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetRepeat() *string {
	return s.Repeat
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetTimePoints() []*CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	return s.TimePoints
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetValidity() *string {
	return s.Validity
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetEndTime(v string) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.EndTime = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetId(v string) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Id = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetPriority(v int32) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Priority = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetRepeat(v string) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Repeat = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetStartTime(v string) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.StartTime = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetTimePoints(v []*CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.TimePoints = v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetTimeZone(v string) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.TimeZone = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetValidity(v string) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Validity = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) Validate() error {
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

type CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints struct {
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

func (s CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMaxInstances() *int32 {
	return s.MaxInstances
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMonthDays() []*int32 {
	return s.MonthDays
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetTime() *string {
	return s.Time
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetWeekdays() []*int32 {
	return s.Weekdays
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMaxInstances(v int32) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MaxInstances = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMinInstances(v int32) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MinInstances = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMonthDays(v []*int32) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MonthDays = v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetTime(v string) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Time = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetWeekdays(v []*int32) *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Weekdays = v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) Validate() error {
	return dara.Validate(s)
}

type CreateServiceEndpointResponseBodyDataTarget struct {
	// The ID of the target agent. This parameter is returned when the target type is AGENT_VERSION.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The version number of the target agent. This parameter is returned when the target type is AGENT_VERSION.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// The type of the collaboration component. This parameter is returned when the target type is TEAM_COLLABORATION.
	//
	// example:
	//
	// ELEMENT_WEB
	CollaborationComponent *string `json:"collaborationComponent,omitempty" xml:"collaborationComponent,omitempty"`
	// The ID of the workspace resource binding associated with the target collaboration component. This parameter is returned when the target type is TEAM_COLLABORATION.
	//
	// example:
	//
	// wrb-123456
	ResourceBindingId *string `json:"resourceBindingId,omitempty" xml:"resourceBindingId,omitempty"`
	// The target type. Valid values: AGENT_VERSION and TEAM_COLLABORATION.
	//
	// example:
	//
	// AGENT_VERSION
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s CreateServiceEndpointResponseBodyDataTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointResponseBodyDataTarget) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointResponseBodyDataTarget) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateServiceEndpointResponseBodyDataTarget) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *CreateServiceEndpointResponseBodyDataTarget) GetCollaborationComponent() *string {
	return s.CollaborationComponent
}

func (s *CreateServiceEndpointResponseBodyDataTarget) GetResourceBindingId() *string {
	return s.ResourceBindingId
}

func (s *CreateServiceEndpointResponseBodyDataTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateServiceEndpointResponseBodyDataTarget) SetAgentId(v string) *CreateServiceEndpointResponseBodyDataTarget {
	s.AgentId = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataTarget) SetAgentVersion(v string) *CreateServiceEndpointResponseBodyDataTarget {
	s.AgentVersion = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataTarget) SetCollaborationComponent(v string) *CreateServiceEndpointResponseBodyDataTarget {
	s.CollaborationComponent = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataTarget) SetResourceBindingId(v string) *CreateServiceEndpointResponseBodyDataTarget {
	s.ResourceBindingId = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataTarget) SetTargetType(v string) *CreateServiceEndpointResponseBodyDataTarget {
	s.TargetType = &v
	return s
}

func (s *CreateServiceEndpointResponseBodyDataTarget) Validate() error {
	return dara.Validate(s)
}
