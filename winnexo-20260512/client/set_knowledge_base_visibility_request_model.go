// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKnowledgeBaseVisibilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SetKnowledgeBaseVisibilityRequest
	GetAppToken() *string
	SetCascadeMode(v string) *SetKnowledgeBaseVisibilityRequest
	GetCascadeMode() *string
	SetOpsUserCode(v string) *SetKnowledgeBaseVisibilityRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SetKnowledgeBaseVisibilityRequest
	GetOpsUserId() *string
	SetTargetId(v string) *SetKnowledgeBaseVisibilityRequest
	GetTargetId() *string
	SetTargetType(v string) *SetKnowledgeBaseVisibilityRequest
	GetTargetType() *string
	SetTenantId(v string) *SetKnowledgeBaseVisibilityRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *SetKnowledgeBaseVisibilityRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *SetKnowledgeBaseVisibilityRequest
	GetUserIds() []*string
	SetVisibilityMode(v string) *SetKnowledgeBaseVisibilityRequest
	GetVisibilityMode() *string
}

type SetKnowledgeBaseVisibilityRequest struct {
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
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The collection of user IDs.
	//
	// example:
	//
	// ["10001"]
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// The visibility scope mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPECIFIED
	VisibilityMode *string `json:"visibilityMode,omitempty" xml:"visibilityMode,omitempty"`
}

func (s SetKnowledgeBaseVisibilityRequest) String() string {
	return dara.Prettify(s)
}

func (s SetKnowledgeBaseVisibilityRequest) GoString() string {
	return s.String()
}

func (s *SetKnowledgeBaseVisibilityRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SetKnowledgeBaseVisibilityRequest) GetCascadeMode() *string {
	return s.CascadeMode
}

func (s *SetKnowledgeBaseVisibilityRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SetKnowledgeBaseVisibilityRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SetKnowledgeBaseVisibilityRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *SetKnowledgeBaseVisibilityRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *SetKnowledgeBaseVisibilityRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SetKnowledgeBaseVisibilityRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *SetKnowledgeBaseVisibilityRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *SetKnowledgeBaseVisibilityRequest) GetVisibilityMode() *string {
	return s.VisibilityMode
}

func (s *SetKnowledgeBaseVisibilityRequest) SetAppToken(v string) *SetKnowledgeBaseVisibilityRequest {
	s.AppToken = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetCascadeMode(v string) *SetKnowledgeBaseVisibilityRequest {
	s.CascadeMode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetOpsUserCode(v string) *SetKnowledgeBaseVisibilityRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetOpsUserId(v string) *SetKnowledgeBaseVisibilityRequest {
	s.OpsUserId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetTargetId(v string) *SetKnowledgeBaseVisibilityRequest {
	s.TargetId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetTargetType(v string) *SetKnowledgeBaseVisibilityRequest {
	s.TargetType = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetTenantId(v string) *SetKnowledgeBaseVisibilityRequest {
	s.TenantId = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetUserGroupIds(v []*string) *SetKnowledgeBaseVisibilityRequest {
	s.UserGroupIds = v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetUserIds(v []*string) *SetKnowledgeBaseVisibilityRequest {
	s.UserIds = v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) SetVisibilityMode(v string) *SetKnowledgeBaseVisibilityRequest {
	s.VisibilityMode = &v
	return s
}

func (s *SetKnowledgeBaseVisibilityRequest) Validate() error {
	return dara.Validate(s)
}
