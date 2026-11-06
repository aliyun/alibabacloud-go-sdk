// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePermissionRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePermissionRulesResponseBody
	GetCode() *string
	SetData(v *CreatePermissionRulesResponseBodyData) *CreatePermissionRulesResponseBody
	GetData() *CreatePermissionRulesResponseBodyData
	SetHttpStatusCode(v int32) *CreatePermissionRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePermissionRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePermissionRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePermissionRulesResponseBody
	GetSuccess() *bool
}

type CreatePermissionRulesResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The creation result.
	Data *CreatePermissionRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request processing result message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1a2b3c4d-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreatePermissionRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePermissionRulesResponseBody) GetData() *CreatePermissionRulesResponseBodyData {
	return s.Data
}

func (s *CreatePermissionRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePermissionRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePermissionRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePermissionRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePermissionRulesResponseBody) SetCode(v string) *CreatePermissionRulesResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePermissionRulesResponseBody) SetData(v *CreatePermissionRulesResponseBodyData) *CreatePermissionRulesResponseBody {
	s.Data = v
	return s
}

func (s *CreatePermissionRulesResponseBody) SetHttpStatusCode(v int32) *CreatePermissionRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePermissionRulesResponseBody) SetMessage(v string) *CreatePermissionRulesResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePermissionRulesResponseBody) SetRequestId(v string) *CreatePermissionRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePermissionRulesResponseBody) SetSuccess(v bool) *CreatePermissionRulesResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePermissionRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePermissionRulesResponseBodyData struct {
	// The list of created permission rules.
	Rules []*CreatePermissionRulesResponseBodyDataRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s CreatePermissionRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesResponseBodyData) GetRules() []*CreatePermissionRulesResponseBodyDataRules {
	return s.Rules
}

func (s *CreatePermissionRulesResponseBodyData) SetRules(v []*CreatePermissionRulesResponseBodyDataRules) *CreatePermissionRulesResponseBodyData {
	s.Rules = v
	return s
}

func (s *CreatePermissionRulesResponseBodyData) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePermissionRulesResponseBodyDataRules struct {
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
	// PERMIT
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
	// EXPLICIT
	SourceDetail *string `json:"sourceDetail,omitempty" xml:"sourceDetail,omitempty"`
	// The target scope.
	//
	// example:
	//
	// SPECIFIED
	TargetScope *string `json:"targetScope,omitempty" xml:"targetScope,omitempty"`
	// The list of targets for the rule.
	Targets []*string `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
	// The update time in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreatePermissionRulesResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetEffect() *string {
	return s.Effect
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetPermissionRuleId() *string {
	return s.PermissionRuleId
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetSource() *string {
	return s.Source
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetSourceDetail() *string {
	return s.SourceDetail
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetTargetScope() *string {
	return s.TargetScope
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetTargets() []*string {
	return s.Targets
}

func (s *CreatePermissionRulesResponseBodyDataRules) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetCreatedAt(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.CreatedAt = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetEffect(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.Effect = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetEnforcementMode(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.EnforcementMode = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetPermissionRuleId(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.PermissionRuleId = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetPrincipalId(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.PrincipalId = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetPrincipalName(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.PrincipalName = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetPrincipalType(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.PrincipalType = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetResourceId(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.ResourceId = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetResourceType(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.ResourceType = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetSource(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.Source = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetSourceDetail(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.SourceDetail = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetTargetScope(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.TargetScope = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetTargets(v []*string) *CreatePermissionRulesResponseBodyDataRules {
	s.Targets = v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) SetUpdatedAt(v string) *CreatePermissionRulesResponseBodyDataRules {
	s.UpdatedAt = &v
	return s
}

func (s *CreatePermissionRulesResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}
