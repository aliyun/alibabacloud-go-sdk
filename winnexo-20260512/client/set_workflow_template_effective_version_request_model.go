// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkflowTemplateEffectiveVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetOpsUserId() *string
	SetTemplateId(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetTemplateId() *string
	SetTenantId(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetTenantId() *string
	SetVersion(v string) *SetWorkflowTemplateEffectiveVersionRequest
	GetVersion() *string
}

type SetWorkflowTemplateEffectiveVersionRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl_9401
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version number to which the returned content belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SetWorkflowTemplateEffectiveVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWorkflowTemplateEffectiveVersionRequest) GoString() string {
	return s.String()
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetAppToken(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.AppToken = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetCollaborationGroupId(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetOpsUserCode(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetOpsUserId(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.OpsUserId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetTemplateId(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.TemplateId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetTenantId(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.TenantId = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) SetVersion(v string) *SetWorkflowTemplateEffectiveVersionRequest {
	s.Version = &v
	return s
}

func (s *SetWorkflowTemplateEffectiveVersionRequest) Validate() error {
	return dara.Validate(s)
}
