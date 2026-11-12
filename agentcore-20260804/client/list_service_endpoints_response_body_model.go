// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListServiceEndpointsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListServiceEndpointsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListServiceEndpointsResponseBodyItems) *ListServiceEndpointsResponseBody
	GetItems() []*ListServiceEndpointsResponseBodyItems
	SetMaxResults(v int32) *ListServiceEndpointsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListServiceEndpointsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListServiceEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceEndpointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListServiceEndpointsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListServiceEndpointsResponseBody
	GetTotalCount() *int64
}

type ListServiceEndpointsResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of service endpoints.
	Items []*ListServiceEndpointsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page that takes effect in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. Returns an error description if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The cursor for the next page. This parameter is empty if there is no next page.
	//
	// example:
	//
	// djE6YWdlbnRjb3JlLnNlcnZpY2UtZW5kcG9pbnQubGlzdDoyMA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of service endpoints that meet the query conditions.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListServiceEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListServiceEndpointsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListServiceEndpointsResponseBody) GetItems() []*ListServiceEndpointsResponseBodyItems {
	return s.Items
}

func (s *ListServiceEndpointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceEndpointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServiceEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceEndpointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceEndpointsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListServiceEndpointsResponseBody) SetCode(v string) *ListServiceEndpointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetHttpStatusCode(v int32) *ListServiceEndpointsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetItems(v []*ListServiceEndpointsResponseBodyItems) *ListServiceEndpointsResponseBody {
	s.Items = v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetMaxResults(v int32) *ListServiceEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetMessage(v string) *ListServiceEndpointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetNextToken(v string) *ListServiceEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetRequestId(v string) *ListServiceEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetSuccess(v bool) *ListServiceEndpointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetTotalCount(v int64) *ListServiceEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceEndpointsResponseBodyItems struct {
	// The list of access URLs for the service endpoint.
	AccessUrls []*ListServiceEndpointsResponseBodyItemsAccessUrls `json:"accessUrls,omitempty" xml:"accessUrls,omitempty" type:"Repeated"`
	// The authentication configuration of the service endpoint.
	Authentication *ListServiceEndpointsResponseBodyItemsAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// The creation time in UTC format compliant with RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Read-only. This is a shared response field. The list operation does not query RM row by row and returns null. For details, call the GetServiceEndpoint operation to query sandboxPhaseCounts.RUNNING based on the endpoint runtimeKey. Do not replace it with the baseline upper and lower limits, target quantity, or READY status.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2
	CurrentInstances *int32 `json:"currentInstances,omitempty" xml:"currentInstances,omitempty"`
	// The result of combining publicAccessEnabled with the public inbound access upper limit of the workspace. The value is true only when both allow public inbound access. This only indicates the policy target and does not mean that the public endpoint has been published. The actual status is subject to accessUrls and status.
	EffectivePublicAccessEnabled *bool `json:"effectivePublicAccessEnabled,omitempty" xml:"effectivePublicAccessEnabled,omitempty"`
	// Includes the enabled status, basic minSandboxCount and maxSandboxCount, and dynamicPolicy. This does not represent the currently effective limits or the actual number of instances. If the basic configuration lacks dynamicPolicy before the feature is released, it is supplemented with enabled=false and policies=[]. This parameter is null if the endpoint configuration is not saved, and the list does not backfill the version source. Historical basic capacity can exceed 100. Retrieve the details again when you open the edit page.
	//
	// if can be null:
	// true
	ElasticityPolicy *ListServiceEndpointsResponseBodyItemsElasticityPolicy `json:"elasticityPolicy,omitempty" xml:"elasticityPolicy,omitempty" type:"Struct"`
	// The type of the service endpoint. DEFAULT indicates a default endpoint created and maintained by the platform. NAMED indicates a named endpoint explicitly created by a user.
	//
	// example:
	//
	// NAMED
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The name of the service endpoint. The name must be unique within the workspace and 1 to 128 characters in length.
	//
	// example:
	//
	// my-agent-endpoint
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The expected public inbound access configured by the user on the service endpoint. For historical service endpoints that have not been configured, true is returned. This only indicates the configuration expectation. The actual availability of the public IP address is subject to accessUrls and status.
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
	// The reason for the service endpoint status. Returns the specific reason when the status is abnormal.
	//
	// example:
	//
	// ServiceEndpoint.Provider.Unavailable: no provider for the target type
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The target routing configuration of the service endpoint.
	Target *ListServiceEndpointsResponseBodyItemsTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
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

func (s ListServiceEndpointsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItems) GetAccessUrls() []*ListServiceEndpointsResponseBodyItemsAccessUrls {
	return s.AccessUrls
}

func (s *ListServiceEndpointsResponseBodyItems) GetAuthentication() *ListServiceEndpointsResponseBodyItemsAuthentication {
	return s.Authentication
}

func (s *ListServiceEndpointsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListServiceEndpointsResponseBodyItems) GetCurrentInstances() *int32 {
	return s.CurrentInstances
}

func (s *ListServiceEndpointsResponseBodyItems) GetEffectivePublicAccessEnabled() *bool {
	return s.EffectivePublicAccessEnabled
}

func (s *ListServiceEndpointsResponseBodyItems) GetElasticityPolicy() *ListServiceEndpointsResponseBodyItemsElasticityPolicy {
	return s.ElasticityPolicy
}

func (s *ListServiceEndpointsResponseBodyItems) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListServiceEndpointsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListServiceEndpointsResponseBodyItems) GetPublicAccessEnabled() *bool {
	return s.PublicAccessEnabled
}

func (s *ListServiceEndpointsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceEndpointsResponseBodyItems) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *ListServiceEndpointsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListServiceEndpointsResponseBodyItems) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListServiceEndpointsResponseBodyItems) GetTarget() *ListServiceEndpointsResponseBodyItemsTarget {
	return s.Target
}

func (s *ListServiceEndpointsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListServiceEndpointsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListServiceEndpointsResponseBodyItems) SetAccessUrls(v []*ListServiceEndpointsResponseBodyItemsAccessUrls) *ListServiceEndpointsResponseBodyItems {
	s.AccessUrls = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetAuthentication(v *ListServiceEndpointsResponseBodyItemsAuthentication) *ListServiceEndpointsResponseBodyItems {
	s.Authentication = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetCreatedAt(v string) *ListServiceEndpointsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetCurrentInstances(v int32) *ListServiceEndpointsResponseBodyItems {
	s.CurrentInstances = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetEffectivePublicAccessEnabled(v bool) *ListServiceEndpointsResponseBodyItems {
	s.EffectivePublicAccessEnabled = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetElasticityPolicy(v *ListServiceEndpointsResponseBodyItemsElasticityPolicy) *ListServiceEndpointsResponseBodyItems {
	s.ElasticityPolicy = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetEndpointType(v string) *ListServiceEndpointsResponseBodyItems {
	s.EndpointType = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetName(v string) *ListServiceEndpointsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetPublicAccessEnabled(v bool) *ListServiceEndpointsResponseBodyItems {
	s.PublicAccessEnabled = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetRegionId(v string) *ListServiceEndpointsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetServiceEndpointId(v string) *ListServiceEndpointsResponseBodyItems {
	s.ServiceEndpointId = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetStatus(v string) *ListServiceEndpointsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetStatusReason(v string) *ListServiceEndpointsResponseBodyItems {
	s.StatusReason = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetTarget(v *ListServiceEndpointsResponseBodyItemsTarget) *ListServiceEndpointsResponseBodyItems {
	s.Target = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetUpdatedAt(v string) *ListServiceEndpointsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetWorkspaceId(v string) *ListServiceEndpointsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) Validate() error {
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

type ListServiceEndpointsResponseBodyItemsAccessUrls struct {
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

func (s ListServiceEndpointsResponseBodyItemsAccessUrls) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsAccessUrls) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) GetAccessType() *string {
	return s.AccessType
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) GetStatus() *string {
	return s.Status
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) GetUrl() *string {
	return s.Url
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) SetAccessType(v string) *ListServiceEndpointsResponseBodyItemsAccessUrls {
	s.AccessType = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) SetStatus(v string) *ListServiceEndpointsResponseBodyItemsAccessUrls {
	s.Status = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) SetStatusReason(v string) *ListServiceEndpointsResponseBodyItemsAccessUrls {
	s.StatusReason = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) SetUrl(v string) *ListServiceEndpointsResponseBodyItemsAccessUrls {
	s.Url = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsAccessUrls) Validate() error {
	return dara.Validate(s)
}

type ListServiceEndpointsResponseBodyItemsAuthentication struct {
	// The authentication method. NONE indicates that no authentication is required. API_KEY indicates that authentication is performed by passing an API key in the x-api-key request header.
	//
	// example:
	//
	// API_KEY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListServiceEndpointsResponseBodyItemsAuthentication) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsAuthentication) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsAuthentication) GetType() *string {
	return s.Type
}

func (s *ListServiceEndpointsResponseBodyItemsAuthentication) SetType(v string) *ListServiceEndpointsResponseBodyItemsAuthentication {
	s.Type = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsAuthentication) Validate() error {
	return dara.Validate(s)
}

type ListServiceEndpointsResponseBodyItemsElasticityPolicy struct {
	// Returns the saved dynamic configuration. If the basic configuration lacks dynamicPolicy before the feature is released, it is supplemented with enabled=false and policies=[].
	DynamicPolicy *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy `json:"dynamicPolicy,omitempty" xml:"dynamicPolicy,omitempty" type:"Struct"`
	// Independent of dynamicPolicy.enabled. If omitted during an update, the original value is retained.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The saved or backfilled basic maximum number of instances. It must not be less than minSandboxCount. Historical values can exceed 100 and are not truncated when read. The basic value is used when dynamic scaling is disabled or no candidates are available.
	//
	// example:
	//
	// 5
	MaxSandboxCount *int32 `json:"maxSandboxCount,omitempty" xml:"maxSandboxCount,omitempty"`
	// The saved or backfilled basic minimum number of instances. It must be a non-negative integer and not greater than maxSandboxCount. Historical values can exceed 100 and are not truncated when read. However, newly written complete configurations must still meet the 0 to 100 range.
	//
	// example:
	//
	// 1
	MinSandboxCount *int32 `json:"minSandboxCount,omitempty" xml:"minSandboxCount,omitempty"`
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicy) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) GetDynamicPolicy() *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy {
	return s.DynamicPolicy
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) GetMaxSandboxCount() *int32 {
	return s.MaxSandboxCount
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) GetMinSandboxCount() *int32 {
	return s.MinSandboxCount
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) SetDynamicPolicy(v *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) *ListServiceEndpointsResponseBodyItemsElasticityPolicy {
	s.DynamicPolicy = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) SetEnabled(v bool) *ListServiceEndpointsResponseBodyItemsElasticityPolicy {
	s.Enabled = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) SetMaxSandboxCount(v int32) *ListServiceEndpointsResponseBodyItemsElasticityPolicy {
	s.MaxSandboxCount = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) SetMinSandboxCount(v int32) *ListServiceEndpointsResponseBodyItemsElasticityPolicy {
	s.MinSandboxCount = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicy) Validate() error {
	if s.DynamicPolicy != nil {
		if err := s.DynamicPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy struct {
	// Required during creation. If omitted during an update, the original value is retained. Disabling it retains existing policies and restores the basic upper and lower limits. Enabling it again triggers immediate calculation. When finally enabled, policies must not be empty.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Required during creation. If omitted during an update, the original list is retained. If provided, it replaces the entire list and is not merged by ID. To clear the list, you must also ensure that the final enabled status is false. Each item must be complete. Conditional requirements and cross-field constraints are validated by the server. The display order does not indicate priority.
	Policies []*ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Repeated"`
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) GetPolicies() []*ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	return s.Policies
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) SetEnabled(v bool) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy {
	s.Enabled = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) SetPolicies(v []*ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy {
	s.Policies = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicy) Validate() error {
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

type ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies struct {
	// Required only when validity is set to range. The format is the same as startTime, and the seconds are fixed at 00. The value must be later than startTime. The interval does not include the end time. If validity is set to forever, the start and end times must be omitted. Valid dates and cross-field constraints are verified by the console.
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
	// Required. Valid values: 5 for regular and 10 for high priority. New policies default to regular priority. It must be explicitly provided during submission, and no default value is set by the API. When dynamicPolicy.enabled is finally true, the validity periods of high-priority policies must not overlap, except for end-to-end connections. Long-term high-priority policies can overlap with other high-priority policies. Overlapping drafts can be saved when disabled, and the final configuration is validated when re-enabled. The validity of other fields and trigger conflicts within the policy are always validated. The priority must be explicitly provided.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// Required. Valid values: daily for every day, weekly for every week, and monthly for every month. Only submit the fields corresponding to the specified cycle.
	//
	// example:
	//
	// weekly
	Repeat *string `json:"repeat,omitempty" xml:"repeat,omitempty"`
	// Required only when validity is set to range. The format is YYYY-MM-DDTHH:mm:ss in the time zone of the policy, without a time zone offset. The seconds are fixed at 00. The interval includes the start time. When editing, the existing past start time is retained, and updates or retries are not rejected even if it is earlier than the current time.
	//
	// example:
	//
	// 2026-11-01T00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Required and non-empty. Each item must be complete. The day of the week or day of the month is configured independently for each item. Within the same policy, daily policies conflict if they have the same time. Weekly and monthly policies conflict only if they have the same time and their date sets intersect. This is validated by the console.
	TimePoints []*ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints `json:"timePoints,omitempty" xml:"timePoints,omitempty" type:"Repeated"`
	// Required. It must be a valid IANA time zone. All dates and times in this policy are interpreted in this time zone.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// Required. Valid values: forever: The policy is permanently valid. range: The policy is valid for a specific period. After expiration, the policy and its trigger results exit the current calculation, and the remaining policies are recalculated. If no candidates are available, the baseline values are restored, and the saved configurations are not automatically deleted.
	//
	// example:
	//
	// range
	Validity *string `json:"validity,omitempty" xml:"validity,omitempty"`
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetEndTime() *string {
	return s.EndTime
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetId() *string {
	return s.Id
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetPriority() *int32 {
	return s.Priority
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetRepeat() *string {
	return s.Repeat
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetStartTime() *string {
	return s.StartTime
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetTimePoints() []*ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints {
	return s.TimePoints
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) GetValidity() *string {
	return s.Validity
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetEndTime(v string) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.EndTime = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetId(v string) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.Id = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetPriority(v int32) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.Priority = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetRepeat(v string) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.Repeat = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetStartTime(v string) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.StartTime = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetTimePoints(v []*ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.TimePoints = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetTimeZone(v string) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.TimeZone = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) SetValidity(v string) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies {
	s.Validity = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPolicies) Validate() error {
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

type ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints struct {
	// Required. The value ranges from 1 to 100.
	//
	// example:
	//
	// 5
	MaxInstances *int32 `json:"maxInstances,omitempty" xml:"maxInstances,omitempty"`
	// Required. The value ranges from 0 to 100 and must not be greater than the maxInstances value in the same item.
	//
	// example:
	//
	// 1
	MinInstances *int32 `json:"minInstances,omitempty" xml:"minInstances,omitempty"`
	// Required only for each trigger point of the monthly cycle. Must be non-empty and contain unique integers from 1 to 31. Each row is selected independently. If a date does not exist in a specific month, the trigger does not occur and is not postponed to the end of the month. Must be omitted for other cycles.
	MonthDays []*int32 `json:"monthDays,omitempty" xml:"monthDays,omitempty" type:"Repeated"`
	// Required. The format is HH:mm:ss. The time is interpreted based on the time zone of the associated policy. The precision is down to the minute, and the seconds are fixed at 00.
	//
	// example:
	//
	// 09:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// Required only for each trigger point of weekly policies. It must be non-empty and deduplicated. Valid values are integers from 1 to 7, where 1 represents Monday and 7 represents Sunday. Each row is selected independently. It must be omitted for other cycles.
	Weekdays []*int32 `json:"weekdays,omitempty" xml:"weekdays,omitempty" type:"Repeated"`
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMaxInstances() *int32 {
	return s.MaxInstances
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMinInstances() *int32 {
	return s.MinInstances
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) GetMonthDays() []*int32 {
	return s.MonthDays
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) GetTime() *string {
	return s.Time
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) GetWeekdays() []*int32 {
	return s.Weekdays
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMaxInstances(v int32) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MaxInstances = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMinInstances(v int32) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MinInstances = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) SetMonthDays(v []*int32) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.MonthDays = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) SetTime(v string) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Time = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) SetWeekdays(v []*int32) *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints {
	s.Weekdays = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsElasticityPolicyDynamicPolicyPoliciesTimePoints) Validate() error {
	return dara.Validate(s)
}

type ListServiceEndpointsResponseBodyItemsTarget struct {
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
	// The workspace resource binding ID associated with the target collaboration component. Returned when the target type is TEAM_COLLABORATION.
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

func (s ListServiceEndpointsResponseBodyItemsTarget) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsTarget) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) GetAgentId() *string {
	return s.AgentId
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) GetCollaborationComponent() *string {
	return s.CollaborationComponent
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) GetResourceBindingId() *string {
	return s.ResourceBindingId
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) GetTargetType() *string {
	return s.TargetType
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) SetAgentId(v string) *ListServiceEndpointsResponseBodyItemsTarget {
	s.AgentId = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) SetAgentVersion(v string) *ListServiceEndpointsResponseBodyItemsTarget {
	s.AgentVersion = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) SetCollaborationComponent(v string) *ListServiceEndpointsResponseBodyItemsTarget {
	s.CollaborationComponent = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) SetResourceBindingId(v string) *ListServiceEndpointsResponseBodyItemsTarget {
	s.ResourceBindingId = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) SetTargetType(v string) *ListServiceEndpointsResponseBodyItemsTarget {
	s.TargetType = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsTarget) Validate() error {
	return dara.Validate(s)
}
