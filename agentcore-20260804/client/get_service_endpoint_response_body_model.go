// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceEndpointResponseBody
	GetCode() *string
	SetData(v *GetServiceEndpointResponseBodyData) *GetServiceEndpointResponseBody
	GetData() *GetServiceEndpointResponseBodyData
	SetHttpStatusCode(v int32) *GetServiceEndpointResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetServiceEndpointResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceEndpointResponseBody
	GetSuccess() *bool
}

type GetServiceEndpointResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The service endpoint details.
	Data *GetServiceEndpointResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. Returns an error description if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetServiceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceEndpointResponseBody) GetData() *GetServiceEndpointResponseBodyData {
	return s.Data
}

func (s *GetServiceEndpointResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetServiceEndpointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceEndpointResponseBody) SetCode(v string) *GetServiceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceEndpointResponseBody) SetData(v *GetServiceEndpointResponseBodyData) *GetServiceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceEndpointResponseBody) SetHttpStatusCode(v int32) *GetServiceEndpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetServiceEndpointResponseBody) SetMessage(v string) *GetServiceEndpointResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceEndpointResponseBody) SetRequestId(v string) *GetServiceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceEndpointResponseBody) SetSuccess(v bool) *GetServiceEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceEndpointResponseBodyData struct {
	// The list of access URLs for the service endpoint.
	AccessUrls []*GetServiceEndpointResponseBodyDataAccessUrls `json:"accessUrls,omitempty" xml:"accessUrls,omitempty" type:"Repeated"`
	// The authentication configuration of the service endpoint.
	Authentication *GetServiceEndpointResponseBodyDataAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// The creation time in UTC format based on RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Read-only. For Get operations, it queries RM sandboxPhaseCounts.RUNNING based on the Endpoint runtimeKey. For Create and Update operations, it returns null. The value is a non-negative integer. If the statistics are unavailable, it returns null. It is not replaced by the baseline upper and lower limits, target quantity, or READY status. Unavailable statistics do not block submission.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2
	CurrentInstances *int32 `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
	// Read-only and returned only in detail queries. The console reads this from autoScalingStatus.effectiveLimits in the same RM detail request that queries the current instance count. It represents the upper and lower limits currently applied by the resource manager (RM), rather than the saved base configuration or the actual instance count. This value is null if it is not applicable, the running object does not exist, the identity does not match, the query fails, or the RM does not return a valid range. Configuration values are not used as a fallback, and this field is not used for edit backfill or submission.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"minSandboxCount":8,"maxSandboxCount":20}
	EffectiveCapacity *GetServiceEndpointResponseBodyDataEffectiveCapacity `json:"effectiveCapacity,omitempty" xml:"effectiveCapacity,omitempty" type:"Struct"`
	// The result of combining publicAccessEnabled with the public network inbound access limit of the workspace. The value is true only when both allow public access. This only indicates the policy target and does not mean that the public network entry has been published. The actual status is subject to accessUrls and status.
	EffectivePublicAccessEnabled *bool `json:"effectivePublicAccessEnabled,omitempty" xml:"effectivePublicAccessEnabled,omitempty"`
	// Used for edit backfill. This is not the currently effective upper and lower limits or the actual instance count. The Get operation preferentially returns the saved configuration. If no configuration is saved, an endpoint that supports elasticity backfills the base values from the frozen configuration of the expected runtime or the bound version configuration when the runtime is not yet established. Successful publication is not required. This value is null when the source is temporarily unavailable or no applicable configuration exists, but the status and statusReason are still returned. The Create and Update operations return only the saved configuration, and return null if no configuration is saved. The historical base capacity can exceed 100, and the read value is not truncated.
	//
	// if can be null:
	// true
	ElasticityPolicy *GetServiceEndpointResponseBodyDataElasticityPolicy `json:"elasticityPolicy,omitempty" xml:"elasticityPolicy,omitempty" type:"Struct"`
	// The type of the service endpoint. DEFAULT indicates a default endpoint created and maintained by the platform. NAMED indicates a named endpoint explicitly created by a user.
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
	// The expected public network inbound access configured by the user on the service endpoint. For historical service endpoints where this has not been configured, true is returned. This only indicates the configuration expectation. The actual availability of the public IP address is subject to accessUrls and status.
	PublicAccessEnabled *bool `json:"publicAccessEnabled,omitempty" xml:"publicAccessEnabled,omitempty"`
	// The ID of the region where the service endpoint resides.
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
	// The status of the service endpoint. Valid values: CREATING, READY, UPDATING, DEGRADED, DISABLED, and DELETING.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the service endpoint status. Returns the specific reason when the status is abnormal.
	//
	// example:
	//
	// ServiceEndpoint.Provider.Unavailable: no provider for the target type
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The target routing configuration of the service endpoint.
	Target *GetServiceEndpointResponseBodyDataTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// The time of the last modification in UTC format based on RFC 3339.
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

func (s GetServiceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyData) GetAccessUrls() []*GetServiceEndpointResponseBodyDataAccessUrls {
	return s.AccessUrls
}

func (s *GetServiceEndpointResponseBodyData) GetAuthentication() *GetServiceEndpointResponseBodyDataAuthentication {
	return s.Authentication
}

func (s *GetServiceEndpointResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetServiceEndpointResponseBodyData) GetCurrentInstances() *int32 {
	return s.CurrentInstances
}

func (s *GetServiceEndpointResponseBodyData) GetEffectiveCapacity() *GetServiceEndpointResponseBodyDataEffectiveCapacity {
	return s.EffectiveCapacity
}

func (s *GetServiceEndpointResponseBodyData) GetEffectivePublicAccessEnabled() *bool {
	return s.EffectivePublicAccessEnabled
}

func (s *GetServiceEndpointResponseBodyData) GetElasticityPolicy() *GetServiceEndpointResponseBodyDataElasticityPolicy {
	return s.ElasticityPolicy
}

func (s *GetServiceEndpointResponseBodyData) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetServiceEndpointResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetServiceEndpointResponseBodyData) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *GetServiceEndpointResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceEndpointResponseBodyData) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *GetServiceEndpointResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetServiceEndpointResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetServiceEndpointResponseBodyData) GetTarget() *GetServiceEndpointResponseBodyDataTarget {
	return s.Target
}

func (s *GetServiceEndpointResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetServiceEndpointResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetServiceEndpointResponseBodyData) SetAccessUrls(v []*GetServiceEndpointResponseBodyDataAccessUrls) *GetServiceEndpointResponseBodyData {
	s.AccessUrls = v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetAuthentication(v *GetServiceEndpointResponseBodyDataAuthentication) *GetServiceEndpointResponseBodyData {
	s.Authentication = v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetCreatedAt(v string) *GetServiceEndpointResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetCurrentInstances(v int32) *GetServiceEndpointResponseBodyData {
	s.CurrentInstances = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetEffectiveCapacity(v *GetServiceEndpointResponseBodyDataEffectiveCapacity) *GetServiceEndpointResponseBodyData {
	s.EffectiveCapacity = v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetEffectivePublicAccessEnabled(v bool) *GetServiceEndpointResponseBodyData {
	s.EffectivePublicAccessEnabled = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetElasticityPolicy(v *GetServiceEndpointResponseBodyDataElasticityPolicy) *GetServiceEndpointResponseBodyData {
	s.ElasticityPolicy = v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetEndpointType(v string) *GetServiceEndpointResponseBodyData {
	s.EndpointType = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetName(v string) *GetServiceEndpointResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetPublicAccessEnabled(v bool) *GetServiceEndpointResponseBodyData {
	s.PublicAccessEnabled = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetRegionId(v string) *GetServiceEndpointResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetServiceEndpointId(v string) *GetServiceEndpointResponseBodyData {
	s.ServiceEndpointId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetStatus(v string) *GetServiceEndpointResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetStatusReason(v string) *GetServiceEndpointResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetTarget(v *GetServiceEndpointResponseBodyDataTarget) *GetServiceEndpointResponseBodyData {
	s.Target = v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetUpdatedAt(v string) *GetServiceEndpointResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetWorkspaceId(v string) *GetServiceEndpointResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) Validate() error {
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
	if s.EffectiveCapacity != nil {
		if err := s.EffectiveCapacity.Validate(); err != nil {
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

type GetServiceEndpointResponseBodyDataAccessUrls struct {
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
	// The reason for the access URL status. Returns the specific reason when the status is abnormal.
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

func (s GetServiceEndpointResponseBodyDataAccessUrls) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataAccessUrls) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) GetAccessType() *string {
	return s.AccessType
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) GetStatus() *string {
	return s.Status
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) GetUrl() *string {
	return s.Url
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) SetAccessType(v string) *GetServiceEndpointResponseBodyDataAccessUrls {
	s.AccessType = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) SetStatus(v string) *GetServiceEndpointResponseBodyDataAccessUrls {
	s.Status = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) SetStatusReason(v string) *GetServiceEndpointResponseBodyDataAccessUrls {
	s.StatusReason = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) SetUrl(v string) *GetServiceEndpointResponseBodyDataAccessUrls {
	s.Url = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataAccessUrls) Validate() error {
	return dara.Validate(s)
}

type GetServiceEndpointResponseBodyDataAuthentication struct {
	// The authentication method. NONE indicates that no authentication is required. API_KEY indicates that authentication is performed by passing an API key in the x-api-key request header.
	//
	// example:
	//
	// API_KEY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetServiceEndpointResponseBodyDataAuthentication) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataAuthentication) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataAuthentication) GetType() *string {
	return s.Type
}

func (s *GetServiceEndpointResponseBodyDataAuthentication) SetType(v string) *GetServiceEndpointResponseBodyDataAuthentication {
	s.Type = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataAuthentication) Validate() error {
	return dara.Validate(s)
}

type GetServiceEndpointResponseBodyDataEffectiveCapacity struct {
	// Read-only. Maps to autoScalingStatus.effectiveLimits.maxSandboxes in the RM. When effectiveCapacity is not null, this value must be returned in pairs with the minimum value. The value must be non-negative and no less than the minimum value. Zero is a valid value. The read value is not truncated at the upper limit of 100 configured for writing.
	//
	// example:
	//
	// 20
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// Read-only. Maps to autoScalingStatus.effectiveLimits.minSandboxes in the RM. When effectiveCapacity is not null, this value must be returned in pairs with the maximum value. The value must be non-negative and no greater than the maximum value. The read value is not truncated at the upper limit of 100 configured for writing.
	//
	// example:
	//
	// 8
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
}

func (s GetServiceEndpointResponseBodyDataEffectiveCapacity) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataEffectiveCapacity) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataEffectiveCapacity) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *GetServiceEndpointResponseBodyDataEffectiveCapacity) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *GetServiceEndpointResponseBodyDataEffectiveCapacity) SetMaxSandboxCount(v int32) *GetServiceEndpointResponseBodyDataEffectiveCapacity {
	s.MaxSandboxCount = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataEffectiveCapacity) SetMinSandboxCount(v int32) *GetServiceEndpointResponseBodyDataEffectiveCapacity {
	s.MinSandboxCount = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataEffectiveCapacity) Validate() error {
	return dara.Validate(s)
}

type GetServiceEndpointResponseBodyDataElasticityPolicy struct {
	// Returns the saved dynamic configuration. If the base configuration before the requirement goes online lacks dynamicPolicy, it is supplemented with enabled=false and policies=[].
	DynamicPolicy *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy `json:"dynamicPolicy,omitempty" xml:"dynamicPolicy,omitempty" type:"Struct"`
	// Independent of dynamicPolicy.enabled. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The saved or backfilled base maximum instance count. The value must be no less than minSandboxCount. Historical values can exceed 100, and the read value is not truncated. The base value is used when dynamic scaling is disabled or no candidates are available.
	//
	// example:
	//
	// 5
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The saved or backfilled base minimum instance count. The value must be non-negative and no greater than maxSandboxCount. Historical values can exceed 100, and the read value is not truncated. A newly written complete configuration must still meet the 0 to 100 range.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicy) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) GetDynamicPolicy() *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	return s.DynamicPolicy
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) SetDynamicPolicy(v *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) *GetServiceEndpointResponseBodyDataElasticityPolicy {
	s.DynamicPolicy = v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) SetEnabled(v bool) *GetServiceEndpointResponseBodyDataElasticityPolicy {
	s.Enabled = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) SetMaxSandboxCount(v int32) *GetServiceEndpointResponseBodyDataElasticityPolicy {
	s.MaxSandboxCount = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) SetMinSandboxCount(v int32) *GetServiceEndpointResponseBodyDataElasticityPolicy {
	s.MinSandboxCount = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicy) Validate() error {
	if s.DynamicPolicy != nil {
		if err := s.DynamicPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy struct {
	// Required during creation. If omitted during an update, the original value is retained. Disabling this feature retains existing policies and restores the base upper and lower limits. Re-enabling it triggers immediate calculation. When finally enabled, policies must not be empty.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Required during creation. If omitted during an update, the original list is retained. If passed, the entire list is replaced rather than merged by ID. To clear the list, you must also ensure that the final enabled value is false. Each item must be complete. Conditional requirements and cross-field constraints are validated by the server. The display order does not indicate priority.
	Policies []*GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Repeated"`
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) GetPolicies() []*GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	return s.Policies
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) SetEnabled(v bool) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	s.Enabled = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) SetPolicies(v []*GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy {
	s.Policies = v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicy) Validate() error {
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

type GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies struct {
	// Required only when validity is set to range. The format is the same as startTime, with seconds fixed at 00. The end time must be later than the startTime. The interval does not include the end time. If validity is set to forever, the start and end times must be omitted. Valid dates and cross-field constraints are validated by the console.
	//
	// example:
	//
	// 2026-12-01T00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Required. Must be unique and stable within the endpoint. Generated by the frontend when added, and retained during editing and retries.
	//
	// example:
	//
	// policy-weekday
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Required. 5 indicates regular priority, and 10 indicates high priority. New policies default to regular priority. Must be explicitly provided during submission, as the API does not set a default value. When the final dynamicPolicy.enabled is true, the validity periods of high-priority policies must not overlap, except that end-to-end connection is allowed, and long-term high-priority policies can overlap with other high-priority policies. Overlapping drafts can be saved when disabled, but the final configuration is validated when re-enabled. The validity of other fields and trigger conflicts within the policy are always validated. The priority must be explicitly provided.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Required. daily means every day, weekly means every week, and monthly means every month. Only submit fields corresponding to the specified cycle.
	//
	// example:
	//
	// weekly
	Repeat *string `json:"repeat,omitempty" xml:"repeat,omitempty"`
	// Required only when validity is set to range. The format is YYYY-MM-DDTHH:mm:ss in the policy time zone, without a time zone offset. The seconds are fixed at 00. The interval includes the start time. When editing, existing past start times are retained, and updates or retries are not rejected even if the start time is earlier than the current time.
	//
	// example:
	//
	// 2026-11-01T00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Required and cannot be empty. Each item must be complete. The day of the week or day of the month is configured independently for each item. Within the same policy, daily policies conflict if they have the same time. Weekly and monthly policies conflict only if they have the same time and their date sets intersect. This is validated by the console.
	TimePoints []*GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints `json:"timePoints,omitempty" xml:"timePoints,omitempty" type:"Repeated"`
	// Required. Must be a valid IANA time zone. All dates and times in this policy are interpreted in this time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// Required. Valid values: forever for long-term validity, and range for short-term validity. After expiration, the policy and its trigger results are excluded from the current calculation, and the remaining policies are recalculated. If no candidate policy is available, the baseline value is restored. Saved configurations are not automatically deleted.
	//
	// example:
	//
	// range
	Validity *string `json:"validity,omitempty" xml:"validity,omitempty"`
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetEndTime() *string {
	return s.EndTime
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetId() *string {
	return s.Id
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetRepeat() *string {
	return s.Repeat
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetStartTime() *string {
	return s.StartTime
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetTimePoints() []*GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	return s.TimePoints
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) GetValidity() *string {
	return s.Validity
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetEndTime(v string) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.EndTime = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetId(v string) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Id = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetPriority(v int32) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Priority = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetRepeat(v string) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Repeat = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetStartTime(v string) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.StartTime = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetTimePoints(v []*GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.TimePoints = v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetTimeZone(v string) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.TimeZone = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) SetValidity(v string) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies {
	s.Validity = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPolicies) Validate() error {
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

type GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints struct {
	// Required. Valid values: 1 to 100.
	//
	// example:
	//
	// 5
	MaxInstances *int32 `json:"maxInstances,omitempty" xml:"maxInstances,omitempty"`
	// Required. Valid values: 0 to 100. The value cannot be greater than the maxInstances value in the same item.
	//
	// example:
	//
	// 1
	MinInstances *int32 `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	// Required only for each trigger point of the monthly cycle. The list must be non-empty and contain unique integers from 1 to 31. Each row is selected independently. If a selected day does not exist in a specific month, the trigger does not fire and is not postponed to the end of the month. This field must be omitted for other cycles.
	MonthDays []*int32 `json:"monthDays,omitempty" xml:"monthDays,omitempty" type:"Repeated"`
	// Required. The format is HH:mm:ss, interpreted in the time zone of the associated policy. The precision is down to the minute, and the seconds are fixed at 00.
	//
	// example:
	//
	// 09:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// Required only for each trigger point of weekly policies. Cannot be empty and must be deduplicated. Valid values are integers from 1 to 7, where 1 represents Monday and 7 represents Sunday. Each row is selected independently. Must be omitted for other cycles.
	Weekdays []*int32 `json:"weekdays,omitempty" xml:"weekdays,omitempty" type:"Repeated"`
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMaxInstances() *int32 {
	return s.MaxInstances
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMonthDays() []*int32 {
	return s.MonthDays
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetTime() *string {
	return s.Time
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) GetWeekdays() []*int32 {
	return s.Weekdays
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMaxInstances(v int32) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MaxInstances = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMinInstances(v int32) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MinInstances = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMonthDays(v []*int32) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MonthDays = v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetTime(v string) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Time = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) SetWeekdays(v []*int32) *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Weekdays = v
	return s
}

func (s *GetServiceEndpointResponseBodyDataElasticityPolicyDynamicPolicyPoliciesTimePoints) Validate() error {
	return dara.Validate(s)
}

type GetServiceEndpointResponseBodyDataTarget struct {
	// The ID of the target agent. Returned when the target type is AGENT_VERSION.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The version number of the target agent. Returned when the target type is AGENT_VERSION.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// The type of the collaboration component. Returned when the target type is TEAM_COLLABORATION.
	//
	// example:
	//
	// ELEMENT_WEB
	CollaborationComponent *string `json:"collaborationComponent,omitempty" xml:"collaborationComponent,omitempty"`
	// The ID of the workspace resource binding associated with the target collaboration component. Returned when the target type is TEAM_COLLABORATION.
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

func (s GetServiceEndpointResponseBodyDataTarget) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponseBodyDataTarget) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponseBodyDataTarget) GetAgentId() *string {
	return s.AgentId
}

func (s *GetServiceEndpointResponseBodyDataTarget) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *GetServiceEndpointResponseBodyDataTarget) GetCollaborationComponent() *string {
	return s.CollaborationComponent
}

func (s *GetServiceEndpointResponseBodyDataTarget) GetResourceBindingId() *string {
	return s.ResourceBindingId
}

func (s *GetServiceEndpointResponseBodyDataTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *GetServiceEndpointResponseBodyDataTarget) SetAgentId(v string) *GetServiceEndpointResponseBodyDataTarget {
	s.AgentId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataTarget) SetAgentVersion(v string) *GetServiceEndpointResponseBodyDataTarget {
	s.AgentVersion = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataTarget) SetCollaborationComponent(v string) *GetServiceEndpointResponseBodyDataTarget {
	s.CollaborationComponent = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataTarget) SetResourceBindingId(v string) *GetServiceEndpointResponseBodyDataTarget {
	s.ResourceBindingId = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataTarget) SetTargetType(v string) *GetServiceEndpointResponseBodyDataTarget {
	s.TargetType = &v
	return s
}

func (s *GetServiceEndpointResponseBodyDataTarget) Validate() error {
	return dara.Validate(s)
}
