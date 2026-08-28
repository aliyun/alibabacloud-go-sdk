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
	// The maximum number of records per page that takes effect for this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. This value is empty if no more pages exist.
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
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of service endpoints that match the query conditions.
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
	// The creation time in UTC, formatted according to RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The service endpoint type. DEFAULT indicates a default endpoint created and maintained by the platform. NAMED indicates a named endpoint explicitly created by the user.
	//
	// example:
	//
	// NAMED
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The service endpoint name. The name is unique within the workspace and is 1 to 128 characters in length.
	//
	// example:
	//
	// my-agent-endpoint
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region ID where the service endpoint is located.
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
	// The service endpoint status. Valid values: CREATING, READY, UPDATING, DEGRADED, DISABLED, DELETING.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the service endpoint status. A specific reason is returned when the status is abnormal.
	//
	// example:
	//
	// ServiceEndpoint.Provider.Unavailable: no provider for the target type
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The target routing configuration of the service endpoint.
	Target *ListServiceEndpointsResponseBodyItemsTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// The last modification time in UTC, formatted according to RFC 3339.
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

func (s *ListServiceEndpointsResponseBodyItems) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ListServiceEndpointsResponseBodyItems) GetName() *string {
	return s.Name
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

func (s *ListServiceEndpointsResponseBodyItems) SetEndpointType(v string) *ListServiceEndpointsResponseBodyItems {
	s.EndpointType = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetName(v string) *ListServiceEndpointsResponseBodyItems {
	s.Name = &v
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
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceEndpointsResponseBodyItemsAccessUrls struct {
	// The access URL type. Valid values: INTERNET, VPC.
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"accessType,omitempty" xml:"accessType,omitempty"`
	// The access URL status. Valid values: CREATING, READY, DEGRADED.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the access URL status. A specific reason is returned when the status is abnormal.
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
	// The authentication method. NONE indicates that no authentication is required. API_KEY indicates that authentication is performed by passing an API key through the x-api-key request header.
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

type ListServiceEndpointsResponseBodyItemsTarget struct {
	// The target agent ID. This parameter is returned when the target type is AGENT_VERSION.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The target agent version number. This parameter is returned when the target type is AGENT_VERSION.
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
	// The target type. Valid values: AGENT_VERSION, TEAM_COLLABORATION.
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
