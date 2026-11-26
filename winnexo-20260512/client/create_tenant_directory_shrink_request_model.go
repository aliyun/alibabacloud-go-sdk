// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantDirectoryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateTenantDirectoryShrinkRequest
	GetAppToken() *string
	SetChangeRecordEnabled(v bool) *CreateTenantDirectoryShrinkRequest
	GetChangeRecordEnabled() *bool
	SetDescription(v string) *CreateTenantDirectoryShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateTenantDirectoryShrinkRequest
	GetName() *string
	SetOoVisibilityMode(v string) *CreateTenantDirectoryShrinkRequest
	GetOoVisibilityMode() *string
	SetOperatingObjectNamesShrink(v string) *CreateTenantDirectoryShrinkRequest
	GetOperatingObjectNamesShrink() *string
	SetOpsUserCode(v string) *CreateTenantDirectoryShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateTenantDirectoryShrinkRequest
	GetOpsUserId() *string
	SetParentId(v int64) *CreateTenantDirectoryShrinkRequest
	GetParentId() *int64
	SetPath(v string) *CreateTenantDirectoryShrinkRequest
	GetPath() *string
	SetTenantId(v string) *CreateTenantDirectoryShrinkRequest
	GetTenantId() *string
}

type CreateTenantDirectoryShrinkRequest struct {
	// The platform application credential. Pass this parameter when you use the application proxy mode, and specify the operator by using opsUserId or opsUserCode based on the application configuration. Do not use real credentials in logs or public examples.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Specifies whether to generate AI version summaries for subsequent content updates. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If this parameter is omitted or set to null, the configuration is not changed. Disabling this feature does not delete existing summaries. This parameter is supported only for standard enterprise knowledge base root directories.
	ChangeRecordEnabled *bool `json:"changeRecordEnabled,omitempty" xml:"changeRecordEnabled,omitempty"`
	// The purpose or content description of the directory. If this parameter is omitted, the value is empty.
	//
	// example:
	//
	// Team product materials and operation guides
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the directory. The name must be 1 to 256 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Product knowledge base
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The digital employee association mode. ALL indicates that the directory is visible to all digital employees, and operatingObjectNames must be omitted, set to null, or set to an empty array. SPECIFIED indicates association by list, and operatingObjectNames must be explicitly provided. An empty array indicates that no employees are associated. If this parameter is omitted or set to null, the association configuration is not changed. This parameter is supported only for standard enterprise knowledge base root directories.
	//
	// example:
	//
	// SPECIFIED
	OoVisibilityMode *string `json:"ooVisibilityMode,omitempty" xml:"ooVisibilityMode,omitempty"`
	// The list of internal names of the specified associated digital employees. This parameter must be used together with ooVisibilityMode. This parameter is required when ooVisibilityMode is set to SPECIFIED. An empty array indicates that all associations are cleared. When ooVisibilityMode is set to ALL, this parameter must be omitted, set to null, or set to an empty array. You can specify up to 100 items. Each item must be 1 to 256 characters in length after leading and trailing spaces are removed. The server verifies that the employees exist and are available in the current tenant, and fully replaces the associations after deduplication.
	OperatingObjectNamesShrink *string `json:"operatingObjectNames,omitempty" xml:"operatingObjectNames,omitempty"`
	// The unique code of the proxied user in the tenant. This parameter is used only when the identity mapping mode of the application credential is source_uid. You cannot use this parameter to bypass tenant or knowledge base permission verification.
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
	// The internal primary key of the parent directory. If this parameter is omitted or set to null, an enterprise knowledge base root directory is created. If this parameter is specified, a subdirectory is created. This parameter must be omitted or set to null when digital employee association or version summary configurations are included.
	//
	// example:
	//
	// 1001
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The path of the directory. If this parameter is omitted or set to null, the server calculates the path based on the parent directory and the name.
	//
	// example:
	//
	// /Product knowledge base
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The ID of the target tenant. This parameter is used for tenant routing in public requests. The server authentication chain verifies access permissions based on this parameter. You cannot use this parameter to access unauthorized tenants.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateTenantDirectoryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantDirectoryShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantDirectoryShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateTenantDirectoryShrinkRequest) GetChangeRecordEnabled() *bool {
	return s.ChangeRecordEnabled
}

func (s *CreateTenantDirectoryShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTenantDirectoryShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateTenantDirectoryShrinkRequest) GetOoVisibilityMode() *string {
	return s.OoVisibilityMode
}

func (s *CreateTenantDirectoryShrinkRequest) GetOperatingObjectNamesShrink() *string {
	return s.OperatingObjectNamesShrink
}

func (s *CreateTenantDirectoryShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateTenantDirectoryShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateTenantDirectoryShrinkRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateTenantDirectoryShrinkRequest) GetPath() *string {
	return s.Path
}

func (s *CreateTenantDirectoryShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTenantDirectoryShrinkRequest) SetAppToken(v string) *CreateTenantDirectoryShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetChangeRecordEnabled(v bool) *CreateTenantDirectoryShrinkRequest {
	s.ChangeRecordEnabled = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetDescription(v string) *CreateTenantDirectoryShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetName(v string) *CreateTenantDirectoryShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetOoVisibilityMode(v string) *CreateTenantDirectoryShrinkRequest {
	s.OoVisibilityMode = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetOperatingObjectNamesShrink(v string) *CreateTenantDirectoryShrinkRequest {
	s.OperatingObjectNamesShrink = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetOpsUserCode(v string) *CreateTenantDirectoryShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetOpsUserId(v string) *CreateTenantDirectoryShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetParentId(v int64) *CreateTenantDirectoryShrinkRequest {
	s.ParentId = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetPath(v string) *CreateTenantDirectoryShrinkRequest {
	s.Path = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) SetTenantId(v string) *CreateTenantDirectoryShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantDirectoryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
