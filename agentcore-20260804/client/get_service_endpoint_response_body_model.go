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
	// The response message. An error description is returned if the request fails.
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
	// Indicates whether the request was successful.
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
	// The access URL list of the service endpoint.
	AccessUrls []*GetServiceEndpointResponseBodyDataAccessUrls `json:"accessUrls,omitempty" xml:"accessUrls,omitempty" type:"Repeated"`
	// The authentication configuration of the service endpoint.
	Authentication *GetServiceEndpointResponseBodyDataAuthentication `json:"authentication,omitempty" xml:"authentication,omitempty" type:"Struct"`
	// The creation time in UTC, formatted in RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The service endpoint type. Valid values:
	//
	// - DEFAULT: a default endpoint created and maintained by the platform.
	//
	// - NAMED: a named endpoint explicitly created by the user.
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
	// The region ID where the service endpoint resides.
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
	Target *GetServiceEndpointResponseBodyDataTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// The last modification time in UTC, formatted in RFC 3339.
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

func (s *GetServiceEndpointResponseBodyData) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetServiceEndpointResponseBodyData) GetName() *string {
	return s.Name
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

func (s *GetServiceEndpointResponseBodyData) SetEndpointType(v string) *GetServiceEndpointResponseBodyData {
	s.EndpointType = &v
	return s
}

func (s *GetServiceEndpointResponseBodyData) SetName(v string) *GetServiceEndpointResponseBodyData {
	s.Name = &v
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
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceEndpointResponseBodyDataAccessUrls struct {
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
	// The reason for the access URL status. A specific reason is returned when the status is degraded.
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
	// The authentication method. Valid values:
	//
	// - NONE: no authentication required.
	//
	// - API_KEY: authentication by passing an API key through the x-api-key request header.
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

type GetServiceEndpointResponseBodyDataTarget struct {
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
