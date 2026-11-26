// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKnowledgeBaseVisibilityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetAppToken() *string
	SetCascadeMode(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetCascadeMode() *string
	SetOpsUserCode(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetOpsUserId() *string
	SetTargetId(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetTargetId() *string
	SetTargetType(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetTargetType() *string
	SetTenantId(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetTenantId() *string
	SetUserGroupIdsShrink(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetUserGroupIdsShrink() *string
	SetUserIdsShrink(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetUserIdsShrink() *string
	SetVisibilityMode(v string) *SetKnowledgeBaseVisibilityShrinkRequest
	GetVisibilityMode() *string
}

type SetKnowledgeBaseVisibilityShrinkRequest struct {
	// token
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies how to handle subfolders.
	//
	// example:
	//
	// CURRENT_ONLY
	CascadeMode *string `json:"cascadeMode,omitempty" xml:"cascadeMode,omitempty"`
	// The user code.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The target ID, which varies based on the target type.
	//
	// - If the type is period, specify the period ID.
	//
	// - If the type is objective, specify the objective ID.
	//
	// - If the type is kr, specify the keyResult ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb_001
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// The type of the current permission. Valid values:
	//
	// - period: period type
	//
	// - objective: objective type
	//
	// - kr: keyResult type
	//
	// This parameter is required.
	//
	// example:
	//
	// KB
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The ID of the tenant to which the resource belongs. If this parameter is not specified, the current tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The collection of user groups.
	//
	// example:
	//
	// ["group-1"]
	UserGroupIdsShrink *string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty"`
	// The collection of user IDs.
	//
	// example:
	//
	// ["10001"]
	UserIdsShrink *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
	// The visibility scope mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPECIFIED
	VisibilityMode *string `json:"visibilityMode,omitempty" xml:"visibilityMode,omitempty"`
}

func (s SetKnowledgeBaseVisibilityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetKnowledgeBaseVisibilityShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetCascadeMode() *string {
	return s.CascadeMode
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetUserGroupIdsShrink() *string {
	return s.UserGroupIdsShrink
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) GetVisibilityMode() *string {
	return s.VisibilityMode
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetAppToken(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetCascadeMode(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.CascadeMode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetOpsUserCode(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetOpsUserId(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetTargetId(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.TargetId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetTargetType(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetTenantId(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetUserGroupIdsShrink(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetUserIdsShrink(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) SetVisibilityMode(v string) *SetKnowledgeBaseVisibilityShrinkRequest {
	s.VisibilityMode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
