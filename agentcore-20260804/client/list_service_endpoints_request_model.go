// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListServiceEndpointsRequest
	GetAgentId() *string
	SetAgentVersion(v string) *ListServiceEndpointsRequest
	GetAgentVersion() *string
	SetCollaborationComponent(v string) *ListServiceEndpointsRequest
	GetCollaborationComponent() *string
	SetMaxResults(v int32) *ListServiceEndpointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceEndpointsRequest
	GetNextToken() *string
	SetResourceBindingId(v string) *ListServiceEndpointsRequest
	GetResourceBindingId() *string
	SetStatus(v string) *ListServiceEndpointsRequest
	GetStatus() *string
	SetTargetType(v string) *ListServiceEndpointsRequest
	GetTargetType() *string
}

type ListServiceEndpointsRequest struct {
	// Filters by target agent ID.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// Filters by target agent version number.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// Filters by collaboration component type. Valid values: MATRIX_CLIENT, MATRIX_FEDERATION, ELEMENT_WEB.
	//
	// example:
	//
	// ELEMENT_WEB
	CollaborationComponent *string `json:"collaborationComponent,omitempty" xml:"collaborationComponent,omitempty"`
	// The maximum number of records per page. Valid values: 1 to 100. If this parameter is not specified, 20 records are returned by default.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. Do not specify this parameter for the first request. For subsequent requests, specify the nextToken value returned in the previous response.
	//
	// example:
	//
	// djE6YWdlbnRjb3JlLnNlcnZpY2UtZW5kcG9pbnQubGlzdDoyMA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Filters by the workspace resource binding ID of the target collaboration component.
	//
	// example:
	//
	// wrb-123456
	ResourceBindingId *string `json:"resourceBindingId,omitempty" xml:"resourceBindingId,omitempty"`
	// Filters by service endpoint status. Valid values: CREATING, READY, UPDATING, DEGRADED, DISABLED, DELETING.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Filters by target type. Valid values: AGENT_VERSION, TEAM_COLLABORATION.
	//
	// example:
	//
	// AGENT_VERSION
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s ListServiceEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListServiceEndpointsRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *ListServiceEndpointsRequest) GetCollaborationComponent() *string {
	return s.CollaborationComponent
}

func (s *ListServiceEndpointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceEndpointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceEndpointsRequest) GetResourceBindingId() *string {
	return s.ResourceBindingId
}

func (s *ListServiceEndpointsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListServiceEndpointsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListServiceEndpointsRequest) SetAgentId(v string) *ListServiceEndpointsRequest {
	s.AgentId = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetAgentVersion(v string) *ListServiceEndpointsRequest {
	s.AgentVersion = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetCollaborationComponent(v string) *ListServiceEndpointsRequest {
	s.CollaborationComponent = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetMaxResults(v int32) *ListServiceEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetNextToken(v string) *ListServiceEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetResourceBindingId(v string) *ListServiceEndpointsRequest {
	s.ResourceBindingId = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetStatus(v string) *ListServiceEndpointsRequest {
	s.Status = &v
	return s
}

func (s *ListServiceEndpointsRequest) SetTargetType(v string) *ListServiceEndpointsRequest {
	s.TargetType = &v
	return s
}

func (s *ListServiceEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
