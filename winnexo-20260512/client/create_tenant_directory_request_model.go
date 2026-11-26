// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateTenantDirectoryRequest
	GetAppToken() *string
	SetChangeRecordEnabled(v bool) *CreateTenantDirectoryRequest
	GetChangeRecordEnabled() *bool
	SetDescription(v string) *CreateTenantDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreateTenantDirectoryRequest
	GetName() *string
	SetOoVisibilityMode(v string) *CreateTenantDirectoryRequest
	GetOoVisibilityMode() *string
	SetOperatingObjectNames(v []*string) *CreateTenantDirectoryRequest
	GetOperatingObjectNames() []*string
	SetOpsUserCode(v string) *CreateTenantDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateTenantDirectoryRequest
	GetOpsUserId() *string
	SetParentId(v int64) *CreateTenantDirectoryRequest
	GetParentId() *int64
	SetPath(v string) *CreateTenantDirectoryRequest
	GetPath() *string
	SetTenantId(v string) *CreateTenantDirectoryRequest
	GetTenantId() *string
}

type CreateTenantDirectoryRequest struct {
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
	OperatingObjectNames []*string `json:"operatingObjectNames,omitempty" xml:"operatingObjectNames,omitempty" type:"Repeated"`
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

func (s CreateTenantDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantDirectoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateTenantDirectoryRequest) GetChangeRecordEnabled() *bool {
	return s.ChangeRecordEnabled
}

func (s *CreateTenantDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTenantDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateTenantDirectoryRequest) GetOoVisibilityMode() *string {
	return s.OoVisibilityMode
}

func (s *CreateTenantDirectoryRequest) GetOperatingObjectNames() []*string {
	return s.OperatingObjectNames
}

func (s *CreateTenantDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateTenantDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateTenantDirectoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateTenantDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *CreateTenantDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTenantDirectoryRequest) SetAppToken(v string) *CreateTenantDirectoryRequest {
	s.AppToken = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetChangeRecordEnabled(v bool) *CreateTenantDirectoryRequest {
	s.ChangeRecordEnabled = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetDescription(v string) *CreateTenantDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetName(v string) *CreateTenantDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetOoVisibilityMode(v string) *CreateTenantDirectoryRequest {
	s.OoVisibilityMode = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetOperatingObjectNames(v []*string) *CreateTenantDirectoryRequest {
	s.OperatingObjectNames = v
	return s
}

func (s *CreateTenantDirectoryRequest) SetOpsUserCode(v string) *CreateTenantDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetOpsUserId(v string) *CreateTenantDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetParentId(v int64) *CreateTenantDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetPath(v string) *CreateTenantDirectoryRequest {
	s.Path = &v
	return s
}

func (s *CreateTenantDirectoryRequest) SetTenantId(v string) *CreateTenantDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
