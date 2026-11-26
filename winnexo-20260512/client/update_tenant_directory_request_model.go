// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateTenantDirectoryRequest
	GetAppToken() *string
	SetChangeRecordEnabled(v bool) *UpdateTenantDirectoryRequest
	GetChangeRecordEnabled() *bool
	SetDescription(v string) *UpdateTenantDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateTenantDirectoryRequest
	GetDirectoryId() *string
	SetName(v string) *UpdateTenantDirectoryRequest
	GetName() *string
	SetOoVisibilityMode(v string) *UpdateTenantDirectoryRequest
	GetOoVisibilityMode() *string
	SetOperatingObjectNames(v []*string) *UpdateTenantDirectoryRequest
	GetOperatingObjectNames() []*string
	SetOpsUserCode(v string) *UpdateTenantDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateTenantDirectoryRequest
	GetOpsUserId() *string
	SetParentId(v int64) *UpdateTenantDirectoryRequest
	GetParentId() *int64
	SetPath(v string) *UpdateTenantDirectoryRequest
	GetPath() *string
	SetTenantId(v string) *UpdateTenantDirectoryRequest
	GetTenantId() *string
}

type UpdateTenantDirectoryRequest struct {
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
	OperatingObjectNames []*string `json:"operatingObjectNames,omitempty" xml:"operatingObjectNames,omitempty" type:"Repeated"`
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

func (s UpdateTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantDirectoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateTenantDirectoryRequest) GetChangeRecordEnabled() *bool {
	return s.ChangeRecordEnabled
}

func (s *UpdateTenantDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTenantDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateTenantDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTenantDirectoryRequest) GetOoVisibilityMode() *string {
	return s.OoVisibilityMode
}

func (s *UpdateTenantDirectoryRequest) GetOperatingObjectNames() []*string {
	return s.OperatingObjectNames
}

func (s *UpdateTenantDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateTenantDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateTenantDirectoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateTenantDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateTenantDirectoryRequest) SetAppToken(v string) *UpdateTenantDirectoryRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetChangeRecordEnabled(v bool) *UpdateTenantDirectoryRequest {
	s.ChangeRecordEnabled = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetDescription(v string) *UpdateTenantDirectoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetDirectoryId(v string) *UpdateTenantDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetName(v string) *UpdateTenantDirectoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetOoVisibilityMode(v string) *UpdateTenantDirectoryRequest {
	s.OoVisibilityMode = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetOperatingObjectNames(v []*string) *UpdateTenantDirectoryRequest {
	s.OperatingObjectNames = v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetOpsUserCode(v string) *UpdateTenantDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetOpsUserId(v string) *UpdateTenantDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetParentId(v int64) *UpdateTenantDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetPath(v string) *UpdateTenantDirectoryRequest {
	s.Path = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) SetTenantId(v string) *UpdateTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
