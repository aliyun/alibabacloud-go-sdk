// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantDirectoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateTenantDirectoryShrinkRequest
	GetAppToken() *string
	SetChangeRecordEnabled(v bool) *UpdateTenantDirectoryShrinkRequest
	GetChangeRecordEnabled() *bool
	SetDescription(v string) *UpdateTenantDirectoryShrinkRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateTenantDirectoryShrinkRequest
	GetDirectoryId() *string
	SetName(v string) *UpdateTenantDirectoryShrinkRequest
	GetName() *string
	SetOoVisibilityMode(v string) *UpdateTenantDirectoryShrinkRequest
	GetOoVisibilityMode() *string
	SetOperatingObjectNamesShrink(v string) *UpdateTenantDirectoryShrinkRequest
	GetOperatingObjectNamesShrink() *string
	SetOpsUserCode(v string) *UpdateTenantDirectoryShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateTenantDirectoryShrinkRequest
	GetOpsUserId() *string
	SetParentId(v int64) *UpdateTenantDirectoryShrinkRequest
	GetParentId() *int64
	SetPath(v string) *UpdateTenantDirectoryShrinkRequest
	GetPath() *string
	SetTenantId(v string) *UpdateTenantDirectoryShrinkRequest
	GetTenantId() *string
}

type UpdateTenantDirectoryShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// 后续内容更新的AI版本摘要开关；false关闭，未传/null保持既有语义，不删除既有摘要
	ChangeRecordEnabled *bool `json:"changeRecordEnabled,omitempty" xml:"changeRecordEnabled,omitempty"`
	// The new directory description. If this parameter is omitted or set to null, the original value is retained.
	//
	// example:
	//
	// Team product materials and operation guides
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The business ID of the enterprise knowledge base directory for the current tenant, which is not the internal primary key. When digital employee associations or version summary configurations are included, the target must be the root directory of a standard enterprise knowledge base.
	//
	// example:
	//
	// kb-example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The new directory name. It can be up to 256 characters in length. If this parameter is omitted or set to null, the original value is retained.
	//
	// example:
	//
	// Product knowledge base
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数字员工关联模式；仅普通企业知识库根可设置。未传/null保持既有语义；SPECIFIED须显式提供名单，ALL仅允许省略或空名单
	//
	// example:
	//
	// SPECIFIED
	OoVisibilityMode *string `json:"ooVisibilityMode,omitempty" xml:"ooVisibilityMode,omitempty"`
	// 数字员工内部名称列表，最多100项，每项非空且最多256字符；SPECIFIED空数组表示不关联任何员工，非空名单按去重后全量替换
	OperatingObjectNamesShrink *string `json:"operatingObjectNames,omitempty" xml:"operatingObjectNames,omitempty"`
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
	// The internal primary key of the new parent folder. If this parameter is omitted or set to null, the original parent folder is retained. If a non-null value is passed, knowledge base management permissions for the target parent folder are required. This parameter cannot be committed at the same time as the digital employee associate or version summary configurations.
	//
	// example:
	//
	// 1001
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The new directory path. If this parameter is omitted or set to null, a new path is not explicitly specified and is processed by the directory service.
	//
	// example:
	//
	// /Product knowledge base
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The ID of the target tenant. It is used for tenant routing in public requests. The server-side authentication chain verifies access permissions. You cannot use this parameter to access unauthorized tenants.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateTenantDirectoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantDirectoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantDirectoryShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateTenantDirectoryShrinkRequest) GetChangeRecordEnabled() *bool {
	return s.ChangeRecordEnabled
}

func (s *UpdateTenantDirectoryShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTenantDirectoryShrinkRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateTenantDirectoryShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTenantDirectoryShrinkRequest) GetOoVisibilityMode() *string {
	return s.OoVisibilityMode
}

func (s *UpdateTenantDirectoryShrinkRequest) GetOperatingObjectNamesShrink() *string {
	return s.OperatingObjectNamesShrink
}

func (s *UpdateTenantDirectoryShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateTenantDirectoryShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateTenantDirectoryShrinkRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateTenantDirectoryShrinkRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateTenantDirectoryShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateTenantDirectoryShrinkRequest) SetAppToken(v string) *UpdateTenantDirectoryShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetChangeRecordEnabled(v bool) *UpdateTenantDirectoryShrinkRequest {
	s.ChangeRecordEnabled = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetDescription(v string) *UpdateTenantDirectoryShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetDirectoryId(v string) *UpdateTenantDirectoryShrinkRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetName(v string) *UpdateTenantDirectoryShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetOoVisibilityMode(v string) *UpdateTenantDirectoryShrinkRequest {
	s.OoVisibilityMode = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetOperatingObjectNamesShrink(v string) *UpdateTenantDirectoryShrinkRequest {
	s.OperatingObjectNamesShrink = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetOpsUserCode(v string) *UpdateTenantDirectoryShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetOpsUserId(v string) *UpdateTenantDirectoryShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetParentId(v int64) *UpdateTenantDirectoryShrinkRequest {
	s.ParentId = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetPath(v string) *UpdateTenantDirectoryShrinkRequest {
	s.Path = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) SetTenantId(v string) *UpdateTenantDirectoryShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateTenantDirectoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
