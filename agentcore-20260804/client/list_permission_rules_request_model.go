// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPermissionRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPermissionRulesRequest
	GetNextToken() *string
	SetPrincipalId(v string) *ListPermissionRulesRequest
	GetPrincipalId() *string
	SetPrincipalName(v string) *ListPermissionRulesRequest
	GetPrincipalName() *string
	SetPrincipalType(v string) *ListPermissionRulesRequest
	GetPrincipalType() *string
	SetResourceId(v string) *ListPermissionRulesRequest
	GetResourceId() *string
	SetResourceType(v string) *ListPermissionRulesRequest
	GetResourceType() *string
	SetSource(v string) *ListPermissionRulesRequest
	GetSource() *string
	SetTarget(v string) *ListPermissionRulesRequest
	GetTarget() *string
}

type ListPermissionRulesRequest struct {
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token used to retrieve the next page.
	//
	// example:
	//
	// token-example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The principal ID.
	//
	// example:
	//
	// 1234567890123456
	PrincipalId *string `json:"principalId,omitempty" xml:"principalId,omitempty"`
	// The principal name.
	//
	// example:
	//
	// example-user
	PrincipalName *string `json:"principalName,omitempty" xml:"principalName,omitempty"`
	// The principal type.
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"principalType,omitempty" xml:"principalType,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-1234567890
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// AGENT
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The rule source.
	//
	// example:
	//
	// AgentCore
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The target of the rule.
	//
	// example:
	//
	// agent-1234567890
	Target *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s ListPermissionRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPermissionRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionRulesRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListPermissionRulesRequest) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListPermissionRulesRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListPermissionRulesRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListPermissionRulesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPermissionRulesRequest) GetSource() *string {
	return s.Source
}

func (s *ListPermissionRulesRequest) GetTarget() *string {
	return s.Target
}

func (s *ListPermissionRulesRequest) SetMaxResults(v int32) *ListPermissionRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionRulesRequest) SetNextToken(v string) *ListPermissionRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPermissionRulesRequest) SetPrincipalId(v string) *ListPermissionRulesRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListPermissionRulesRequest) SetPrincipalName(v string) *ListPermissionRulesRequest {
	s.PrincipalName = &v
	return s
}

func (s *ListPermissionRulesRequest) SetPrincipalType(v string) *ListPermissionRulesRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListPermissionRulesRequest) SetResourceId(v string) *ListPermissionRulesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListPermissionRulesRequest) SetResourceType(v string) *ListPermissionRulesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionRulesRequest) SetSource(v string) *ListPermissionRulesRequest {
	s.Source = &v
	return s
}

func (s *ListPermissionRulesRequest) SetTarget(v string) *ListPermissionRulesRequest {
	s.Target = &v
	return s
}

func (s *ListPermissionRulesRequest) Validate() error {
	return dara.Validate(s)
}
