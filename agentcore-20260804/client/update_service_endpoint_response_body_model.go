// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateServiceEndpointResponseBody
	GetCode() *string
	SetData(v *UpdateServiceEndpointResponseBodyData) *UpdateServiceEndpointResponseBody
	GetData() *UpdateServiceEndpointResponseBodyData
	SetHttpStatusCode(v int32) *UpdateServiceEndpointResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateServiceEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateServiceEndpointResponseBody
	GetSuccess() *bool
}

type UpdateServiceEndpointResponseBody struct {
	// The response code of the request. The value is SUCCESS if the request is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The service endpoint information returned after the update is accepted. This includes the binding target, saved authentication and elasticity configurations, and current status. The endpoint may still be updating.
	Data *UpdateServiceEndpointResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The unique identifier of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5E8F3C2A-1B4D-4F6E-9A7C-0D2B8E6F4A1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates that the update request is accepted, but does not mean that the deployment or endpoint change is complete. You must check data.status and data.accessUrls to confirm the result.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateServiceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateServiceEndpointResponseBody) GetData() *UpdateServiceEndpointResponseBodyData {
	return s.Data
}

func (s *UpdateServiceEndpointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateServiceEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateServiceEndpointResponseBody) SetCode(v string) *UpdateServiceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetData(v *UpdateServiceEndpointResponseBodyData) *UpdateServiceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetHttpStatusCode(v int32) *UpdateServiceEndpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetMessage(v string) *UpdateServiceEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetRequestId(v string) *UpdateServiceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) SetSuccess(v bool) *UpdateServiceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateServiceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceEndpointResponseBodyData struct {
	// The list of access URLs for the service endpoint.
	AccessUrls []*UpdateServiceEndpointResponseBodyDataAccessUrls `json:"accessUrls,omitempty" xml:"accessUrls,omitempty" type:"Repeated"`
	// The authentication configuration of the service endpoint.
	Authentication *UpdateServiceEndpointResponseBodyDataAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// The creation time in UTC format compliant with RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Read-only. The Get operation queries RM sandboxPhaseCounts.RUNNING by endpoint runtimeKey. The Create and Update operations return null. The value is a non-negative integer. If the statistics are unavailable, null is returned. The value is not replaced by the base upper and lower limits, target quantity, or READY status. Unavailable statistics do not block submission.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2
	CurrentInstances *int32 `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
	// The effective public access setting derived from combining publicAccessEnabled with the workspace-level public access limit. This parameter returns true only when both settings allow public access. This value only indicates the policy target and does not mean that the public endpoint has been fully published. For the actual status, refer to accessUrls and status.
	EffectivePublicAccessEnabled *bool `json:"effectivePublicAccessEnabled,omitempty" xml:"effectivePublicAccessEnabled,omitempty"`
	// The elasticity configuration used for edit backfill, which is not the current effective upper and lower limits or the actual number of instances. The Get operation preferentially returns the saved configuration. If no configuration is saved, an endpoint that supports elasticity backfills the base values from the expected runtime frozen configuration or the bound version configuration when the runtime is not established. Successful publication is not required. If the source is temporarily unavailable or no applicable configuration exists, the value is null, but the status and statusReason are still returned. The Create and Update operations return only the saved configuration. If no configuration is saved, the value is null. The historical base capacity can exceed 100 and is not truncated when read.
	//
	// if can be null:
	// true
	ElasticityPolicy *UpdateServiceEndpointResponseBodyDataElasticityPolicy `json:"elasticityPolicy,omitempty" xml:"elasticityPolicy,omitempty" type:"Struct"`
	// The service endpoint type. Valid values:
	//
	// - DEFAULT: The default endpoint created and maintained by the platform.
	//
	// - NAMED: The named endpoint explicitly created by the user.
	//
	// example:
	//
	// NAMED
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The service endpoint name. The name must be unique within the workspace and 1 to 128 characters in length.
	//
	// example:
	//
	// my-agent-endpoint
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The expected public access setting configured by the user for the service endpoint. For historical service endpoints that have not been configured, this parameter returns true. This value only indicates the configuration expectation. For the actual availability of the public endpoint, refer to accessUrls and status.
	PublicAccessEnabled *bool `json:"publicAccessEnabled,omitempty" xml:"publicAccessEnabled,omitempty"`
	// The region ID of the service endpoint.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The service endpoint ID.
	//
	// example:
	//
	// se-123456
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
	// The service endpoint status. Valid values: CREATING, READY, UPDATING, DEGRADED, DISABLED, and DELETING.
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
	Target *UpdateServiceEndpointResponseBodyDataTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// The time of the last modification in UTC format compliant with RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateServiceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyData) GetAccessUrls() []*UpdateServiceEndpointResponseBodyDataAccessUrls {
	return s.AccessUrls
}

func (s *UpdateServiceEndpointResponseBodyData) GetAuthentication() *UpdateServiceEndpointResponseBodyDataAuthentication {
	return s.Authentication
}

func (s *UpdateServiceEndpointResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateServiceEndpointResponseBodyData) GetCurrentInstances() *int32 {
	return s.CurrentInstances
}

func (s *UpdateServiceEndpointResponseBodyData) GetEffectivePublicAccessEnabled() *bool {
	return s.EffectivePublicAccessEnabled
}

func (s *UpdateServiceEndpointResponseBodyData) GetElasticityPolicy() *UpdateServiceEndpointResponseBodyDataElasticityPolicy {
	return s.ElasticityPolicy
}

func (s *UpdateServiceEndpointResponseBodyData) GetEndpointType() *string {
	return s.EndpointType
}

func (s *UpdateServiceEndpointResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateServiceEndpointResponseBodyData) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *UpdateServiceEndpointResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceEndpointResponseBodyData) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *UpdateServiceEndpointResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateServiceEndpointResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *UpdateServiceEndpointResponseBodyData) GetTarget() *UpdateServiceEndpointResponseBodyDataTarget {
	return s.Target
}

func (s *UpdateServiceEndpointResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateServiceEndpointResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateServiceEndpointResponseBodyData) SetAccessUrls(v []*UpdateServiceEndpointResponseBodyDataAccessUrls) *UpdateServiceEndpointResponseBodyData {
	s.AccessUrls = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetAuthentication(v *UpdateServiceEndpointResponseBodyDataAuthentication) *UpdateServiceEndpointResponseBodyData {
	s.Authentication = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetCreatedAt(v string) *UpdateServiceEndpointResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetCurrentInstances(v int32) *UpdateServiceEndpointResponseBodyData {
	s.CurrentInstances = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetEffectivePublicAccessEnabled(v bool) *UpdateServiceEndpointResponseBodyData {
	s.EffectivePublicAccessEnabled = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetElasticityPolicy(v *UpdateServiceEndpointResponseBodyDataElasticityPolicy) *UpdateServiceEndpointResponseBodyData {
	s.ElasticityPolicy = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetEndpointType(v string) *UpdateServiceEndpointResponseBodyData {
	s.EndpointType = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetName(v string) *UpdateServiceEndpointResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetPublicAccessEnabled(v bool) *UpdateServiceEndpointResponseBodyData {
	s.PublicAccessEnabled = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetRegionId(v string) *UpdateServiceEndpointResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetServiceEndpointId(v string) *UpdateServiceEndpointResponseBodyData {
	s.ServiceEndpointId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetStatus(v string) *UpdateServiceEndpointResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetStatusReason(v string) *UpdateServiceEndpointResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetTarget(v *UpdateServiceEndpointResponseBodyDataTarget) *UpdateServiceEndpointResponseBodyData {
	s.Target = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetUpdatedAt(v string) *UpdateServiceEndpointResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) SetWorkspaceId(v string) *UpdateServiceEndpointResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyData) Validate() error {
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

type UpdateServiceEndpointResponseBodyDataAccessUrls struct {
	// The access URL type. Valid values: INTERNET and VPC.
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
	// The access URL status. Valid values: CREATING, READY, and DEGRADED.
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

func (s UpdateServiceEndpointResponseBodyDataAccessUrls) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyDataAccessUrls) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) GetAccessType() *string {
	return s.AccessType
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) GetStatus() *string {
	return s.Status
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) GetStatusReason() *string {
	return s.StatusReason
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) GetUrl() *string {
	return s.Url
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) SetAccessType(v string) *UpdateServiceEndpointResponseBodyDataAccessUrls {
	s.AccessType = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) SetStatus(v string) *UpdateServiceEndpointResponseBodyDataAccessUrls {
	s.Status = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) SetStatusReason(v string) *UpdateServiceEndpointResponseBodyDataAccessUrls {
	s.StatusReason = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) SetUrl(v string) *UpdateServiceEndpointResponseBodyDataAccessUrls {
	s.Url = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataAccessUrls) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceEndpointResponseBodyDataAuthentication struct {
	// The authentication method. Valid values:
	//
	// - NONE: No authentication is required.
	//
	// - API_KEY: Authentication is performed by passing the API key in the x-api-key request header.
	//
	// example:
	//
	// API_KEY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateServiceEndpointResponseBodyDataAuthentication) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyDataAuthentication) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyDataAuthentication) GetType() *string {
	return s.Type
}

func (s *UpdateServiceEndpointResponseBodyDataAuthentication) SetType(v string) *UpdateServiceEndpointResponseBodyDataAuthentication {
	s.Type = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataAuthentication) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceEndpointResponseBodyDataElasticityPolicy struct {
	// The saved dynamic configuration. If the base configuration before the requirement is published lacks the dynamicPolicy parameter, it is supplemented with enabled=false and policies=[].
	DynamicPolicy *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy `json:"dynamicPolicy,omitempty" xml:"dynamicPolicy,omitempty" type:"Struct"`
	// Specifies whether to enable the overall automatic elasticity mode. This parameter is independent of dynamicPolicy.enabled. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The saved or backfilled base maximum number of instances, which is not less than minSandboxCount. The historical value can exceed 100 and is not truncated when read. The base value is used when dynamic scaling is disabled or no candidates are available.
	//
	// example:
	//
	// 5
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The saved or backfilled base minimum number of instances, which is a non-negative integer and not greater than maxSandboxCount. The historical value can exceed 100 and is not truncated when read. A newly written complete configuration must still meet the 0 to 100 range.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
}

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) GetDynamicPolicy() *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	return s.DynamicPolicy
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) SetDynamicPolicy(v *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) *UpdateServiceEndpointResponseBodyDataElasticityPolicy {
	s.DynamicPolicy = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) SetEnabled(v bool) *UpdateServiceEndpointResponseBodyDataElasticityPolicy {
	s.Enabled = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) SetMaxSandboxCount(v int32) *UpdateServiceEndpointResponseBodyDataElasticityPolicy {
	s.MaxSandboxCount = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) SetMinSandboxCount(v int32) *UpdateServiceEndpointResponseBodyDataElasticityPolicy {
	s.MinSandboxCount = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicy) Validate() error {
	if s.DynamicPolicy != nil {
		if err := s.DynamicPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy struct {
	// Specifies whether to enable dynamic policies. This parameter is required during creation. If omitted during an update, the original value is retained. Disabling it retains existing policies and restores the basic upper and lower limits. Re-enabling it triggers immediate calculation. When it is finally enabled, policies cannot be empty.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The complete list of dynamic policies. This parameter is required during creation. If omitted during an update, the original list is retained. If provided, it replaces the entire list and is not merged by ID. To clear the list, you must ensure that the final enabled is set to false. Each item must be complete. Conditional requirements and cross-field constraints are verified by the server. The display order does not indicate priority.
	Policies []*UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Repeated"`
}

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GetPolicies() []*UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	return s.Policies
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) SetEnabled(v bool) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	s.Enabled = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) SetPolicies(v []*UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	s.Policies = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) Validate() error {
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

type UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies struct {
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
	TimePoints []*UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints `json:"timePoints,omitempty" xml:"timePoints,omitempty" type:"Repeated"`
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

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetId() *string {
	return s.Id
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetRepeat() *string {
	return s.Repeat
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetTimePoints() []*UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	return s.TimePoints
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetTimeZone() *string {
	return s.TimeZone
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetValidity() *string {
	return s.Validity
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetEndTime(v string) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.EndTime = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetId(v string) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Id = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetPriority(v int32) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Priority = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetRepeat(v string) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Repeat = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetStartTime(v string) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.StartTime = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetTimePoints(v []*UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.TimePoints = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetTimeZone(v string) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.TimeZone = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetValidity(v string) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Validity = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) Validate() error {
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

type UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints struct {
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

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMaxInstances() *int32 {
	return s.MaxInstances
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMonthDays() []*int32 {
	return s.MonthDays
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetTime() *string {
	return s.Time
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetWeekdays() []*int32 {
	return s.Weekdays
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMaxInstances(v int32) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MaxInstances = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMinInstances(v int32) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MinInstances = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMonthDays(v []*int32) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MonthDays = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetTime(v string) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Time = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetWeekdays(v []*int32) *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Weekdays = v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceEndpointResponseBodyDataTarget struct {
	// The target agent ID. This parameter is returned when the target type is AGENT_VERSION.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The target agent version. This parameter is returned when the target type is AGENT_VERSION.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// The collaboration component type. This parameter is returned when the target type is TEAM_COLLABORATION.
	//
	// example:
	//
	// ELEMENT_WEB
	CollaborationComponent *string `json:"collaborationComponent,omitempty" xml:"collaborationComponent,omitempty"`
	// The workspace resource binding ID associated with the target collaboration component. This parameter is returned when the target type is TEAM_COLLABORATION.
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

func (s UpdateServiceEndpointResponseBodyDataTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointResponseBodyDataTarget) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) GetCollaborationComponent() *string {
	return s.CollaborationComponent
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) GetResourceBindingId() *string {
	return s.ResourceBindingId
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) SetAgentId(v string) *UpdateServiceEndpointResponseBodyDataTarget {
	s.AgentId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) SetAgentVersion(v string) *UpdateServiceEndpointResponseBodyDataTarget {
	s.AgentVersion = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) SetCollaborationComponent(v string) *UpdateServiceEndpointResponseBodyDataTarget {
	s.CollaborationComponent = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) SetResourceBindingId(v string) *UpdateServiceEndpointResponseBodyDataTarget {
	s.ResourceBindingId = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) SetTargetType(v string) *UpdateServiceEndpointResponseBodyDataTarget {
	s.TargetType = &v
	return s
}

func (s *UpdateServiceEndpointResponseBodyDataTarget) Validate() error {
	return dara.Validate(s)
}
