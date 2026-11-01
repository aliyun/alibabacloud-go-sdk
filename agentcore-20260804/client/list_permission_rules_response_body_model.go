// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPermissionRulesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListPermissionRulesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListPermissionRulesResponseBodyItems) *ListPermissionRulesResponseBody
	GetItems() []*ListPermissionRulesResponseBodyItems
	SetMaxResults(v int32) *ListPermissionRulesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListPermissionRulesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListPermissionRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPermissionRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPermissionRulesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListPermissionRulesResponseBody
	GetTotalCount() *int64
}

type ListPermissionRulesResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of permission rules.
	Items []*ListPermissionRulesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries returned in the current request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The request processing result message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The next page token.
	//
	// example:
	//
	// token-example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of rules that meet the specified conditions.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPermissionRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPermissionRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPermissionRulesResponseBody) GetItems() []*ListPermissionRulesResponseBodyItems {
	return s.Items
}

func (s *ListPermissionRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPermissionRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPermissionRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPermissionRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPermissionRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPermissionRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPermissionRulesResponseBody) SetCode(v string) *ListPermissionRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPermissionRulesResponseBody) SetHttpStatusCode(v int32) *ListPermissionRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPermissionRulesResponseBody) SetItems(v []*ListPermissionRulesResponseBodyItems) *ListPermissionRulesResponseBody {
	s.Items = v
	return s
}

func (s *ListPermissionRulesResponseBody) SetMaxResults(v int32) *ListPermissionRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPermissionRulesResponseBody) SetMessage(v string) *ListPermissionRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPermissionRulesResponseBody) SetNextToken(v string) *ListPermissionRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPermissionRulesResponseBody) SetRequestId(v string) *ListPermissionRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionRulesResponseBody) SetSuccess(v bool) *ListPermissionRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPermissionRulesResponseBody) SetTotalCount(v int64) *ListPermissionRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPermissionRulesResponseBody) Validate() error {
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

type ListPermissionRulesResponseBodyItems struct {
	// The creation time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The authorization effect.
	//
	// example:
	//
	// ALLOW
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The rule enforcement mode.
	//
	// example:
	//
	// ENFORCE
	EnforcementMode *string `json:"enforcementMode,omitempty" xml:"enforcementMode,omitempty"`
	// The permission rule ID.
	//
	// example:
	//
	// pr-1234567890
	PermissionRuleId *string `json:"permissionRuleId,omitempty" xml:"permissionRuleId,omitempty"`
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
	// example:
	//
	// agent-1234567890
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
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
	// The rule source details.
	//
	// example:
	//
	// OpenAPI
	SourceDetail *string `json:"sourceDetail,omitempty" xml:"sourceDetail,omitempty"`
	// The target scope.
	//
	// example:
	//
	// SPECIFIED
	TargetScope *string `json:"targetScope,omitempty" xml:"targetScope,omitempty"`
	// The list of rule targets.
	Targets []*string `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
	// The update time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListPermissionRulesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionRulesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListPermissionRulesResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListPermissionRulesResponseBodyItems) GetEffect() *string {
	return s.Effect
}

func (s *ListPermissionRulesResponseBodyItems) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *ListPermissionRulesResponseBodyItems) GetPermissionRuleId() *string {
	return s.PermissionRuleId
}

func (s *ListPermissionRulesResponseBodyItems) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListPermissionRulesResponseBodyItems) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListPermissionRulesResponseBodyItems) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListPermissionRulesResponseBodyItems) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListPermissionRulesResponseBodyItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListPermissionRulesResponseBodyItems) GetSource() *string {
	return s.Source
}

func (s *ListPermissionRulesResponseBodyItems) GetSourceDetail() *string {
	return s.SourceDetail
}

func (s *ListPermissionRulesResponseBodyItems) GetTargetScope() *string {
	return s.TargetScope
}

func (s *ListPermissionRulesResponseBodyItems) GetTargets() []*string {
	return s.Targets
}

func (s *ListPermissionRulesResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListPermissionRulesResponseBodyItems) SetCreatedAt(v string) *ListPermissionRulesResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetEffect(v string) *ListPermissionRulesResponseBodyItems {
	s.Effect = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetEnforcementMode(v string) *ListPermissionRulesResponseBodyItems {
	s.EnforcementMode = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetPermissionRuleId(v string) *ListPermissionRulesResponseBodyItems {
	s.PermissionRuleId = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetPrincipalId(v string) *ListPermissionRulesResponseBodyItems {
	s.PrincipalId = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetPrincipalName(v string) *ListPermissionRulesResponseBodyItems {
	s.PrincipalName = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetPrincipalType(v string) *ListPermissionRulesResponseBodyItems {
	s.PrincipalType = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetResourceId(v string) *ListPermissionRulesResponseBodyItems {
	s.ResourceId = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetResourceType(v string) *ListPermissionRulesResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetSource(v string) *ListPermissionRulesResponseBodyItems {
	s.Source = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetSourceDetail(v string) *ListPermissionRulesResponseBodyItems {
	s.SourceDetail = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetTargetScope(v string) *ListPermissionRulesResponseBodyItems {
	s.TargetScope = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetTargets(v []*string) *ListPermissionRulesResponseBodyItems {
	s.Targets = v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) SetUpdatedAt(v string) *ListPermissionRulesResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListPermissionRulesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
