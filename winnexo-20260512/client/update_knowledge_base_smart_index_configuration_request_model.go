// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSmartIndexConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetAppToken() *string
	SetCronExpression(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetCronExpression() *string
	SetDirectoryId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetDirectoryId() *string
	SetEnabled(v bool) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetEnabled() *bool
	SetExpectedRevision(v int64) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetExpectedRevision() *int64
	SetOpsUserCode(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetOpsUserId() *string
	SetTenantId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetTenantId() *string
	SetTimezone(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest
	GetTimezone() *string
}

type UpdateKnowledgeBaseSmartIndexConfigurationRequest struct {
	// The platform application credential. Pass this parameter when using the application proxy mode, and specify the operator by using opsUserId or opsUserCode based on the application configuration. Do not use real credentials in logs or public examples.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The automatic execution schedule. Use a five-part cron expression where the minute field must be 0. For example, 0 2 	- 	- 	- indicates execution at 02:00 every day based on the timezone. Omitting this parameter, setting it to null, or using a blank string indicates manual execution and clears the old automatic schedule, but does not disable the smart index. Maximum length: 100 characters.
	//
	// example:
	//
	// 0 2 	- 	- *
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// The business ID of the root directory for the standard enterprise knowledge base of the current tenant. You can obtain this ID from the response of the CreateTenantDirectory operation. You cannot use the internal primary key of the directory, or the IDs of personal spaces, group knowledge bases, subdirectories, or special knowledge bases.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// Specifies whether to enable the LLM Wiki smart index. Valid values:
	//
	// - true: Enabled. Switching from disabled to enabled creates an asynchronous full build task and a bill.
	//
	// - false: Disabled. If artifacts exist, asynchronous cleanup is initiated.
	//
	// You cannot toggle the switch during task execution, and you cannot re-enable it until cleanup is complete.
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The revision obtained from the most recent configuration read. This value must be a non-negative integer. The value is 0 when not configured, and increments by 1 after a successful save. If the version does not match, ERR.SmartIndex.Conflict is returned. You must read the configuration again before saving. This field is not an idempotency key.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ExpectedRevision *int64 `json:"expectedRevision,omitempty" xml:"expectedRevision,omitempty"`
	// The unique code of the proxied user within the tenant. This parameter is used only when the identity mapping mode of the application credential is source_uid. You cannot use this parameter to bypass tenant or knowledge base permission verification.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the proxied WINNEXO platform user. This parameter is used only when the identity mapping mode of the application credential is winnexo_uid. You cannot use this parameter to bypass tenant or knowledge base permission verification.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The ID of the target tenant. This parameter is used for tenant routing in public requests. The server-side authentication chain verifies access permissions based on this parameter. You cannot use this parameter to access unauthorized tenants.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The IANA time zone name used for the execution schedule. Default value: Asia/Shanghai. If this parameter is omitted, set to null, or left as a blank string, the default value is used. The length of a non-empty value must be 1 to 64 characters.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSmartIndexConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetExpectedRevision() *int64 {
	return s.ExpectedRevision
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetAppToken(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetCronExpression(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.CronExpression = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetDirectoryId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetEnabled(v bool) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetExpectedRevision(v int64) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.ExpectedRevision = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetOpsUserCode(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetOpsUserId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetTenantId(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) SetTimezone(v string) *UpdateKnowledgeBaseSmartIndexConfigurationRequest {
	s.Timezone = &v
	return s
}

func (s *UpdateKnowledgeBaseSmartIndexConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
