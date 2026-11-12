// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSandboxSessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSandboxSessionsRequest
	GetNextToken() *string
	SetServiceEndpointId(v string) *ListSandboxSessionsRequest
	GetServiceEndpointId() *string
}

type ListSandboxSessionsRequest struct {
	// The maximum number of entries to return per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. Omit this parameter in the first request. For subsequent requests, use the returned token as is and keep the workspace, agent, instance, and maxResults unchanged. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Must be consistent with the ListSandboxes input parameters used to retrieve the sandboxId. If omitted or left blank, the latest running unit of the agent is queried. If specified, the service endpoint must belong to the agent and the target type must be AGENT_VERSION. Otherwise, Resource.NotFound is returned. If it is not DEFAULT and is not running independently, ServiceEndpointTargetNotReady is returned. If the sandboxId does not belong to the specified running unit, Sandbox.NotFound is returned instead of an empty list.
	//
	// example:
	//
	// se-7f3a92c1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
}

func (s ListSandboxSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListSandboxSessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSandboxSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxSessionsRequest) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *ListSandboxSessionsRequest) SetMaxResults(v int32) *ListSandboxSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSandboxSessionsRequest) SetNextToken(v string) *ListSandboxSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSandboxSessionsRequest) SetServiceEndpointId(v string) *ListSandboxSessionsRequest {
	s.ServiceEndpointId = &v
	return s
}

func (s *ListSandboxSessionsRequest) Validate() error {
	return dara.Validate(s)
}
