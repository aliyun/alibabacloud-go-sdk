// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetSkillRunRequest
	GetAppToken() *string
	SetIncludeLogs(v bool) *GetSkillRunRequest
	GetIncludeLogs() *bool
	SetOpsUserCode(v string) *GetSkillRunRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetSkillRunRequest
	GetOpsUserId() *string
	SetRunId(v string) *GetSkillRunRequest
	GetRunId() *string
	SetTenantId(v string) *GetSkillRunRequest
	GetTenantId() *string
}

type GetSkillRunRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to include execution logs. Default value: false. Enable this only when troubleshooting issues.
	//
	// example:
	//
	// false
	IncludeLogs *bool `json:"includeLogs,omitempty" xml:"includeLogs,omitempty"`
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
	// The evaluation run ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleRunId
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSkillRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRunRequest) GoString() string {
	return s.String()
}

func (s *GetSkillRunRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetSkillRunRequest) GetIncludeLogs() *bool {
	return s.IncludeLogs
}

func (s *GetSkillRunRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetSkillRunRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetSkillRunRequest) GetRunId() *string {
	return s.RunId
}

func (s *GetSkillRunRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSkillRunRequest) SetAppToken(v string) *GetSkillRunRequest {
	s.AppToken = &v
	return s
}

func (s *GetSkillRunRequest) SetIncludeLogs(v bool) *GetSkillRunRequest {
	s.IncludeLogs = &v
	return s
}

func (s *GetSkillRunRequest) SetOpsUserCode(v string) *GetSkillRunRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetSkillRunRequest) SetOpsUserId(v string) *GetSkillRunRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetSkillRunRequest) SetRunId(v string) *GetSkillRunRequest {
	s.RunId = &v
	return s
}

func (s *GetSkillRunRequest) SetTenantId(v string) *GetSkillRunRequest {
	s.TenantId = &v
	return s
}

func (s *GetSkillRunRequest) Validate() error {
	return dara.Validate(s)
}
