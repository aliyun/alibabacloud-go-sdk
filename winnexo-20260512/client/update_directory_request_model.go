// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateDirectoryRequest
	GetAppToken() *string
	SetDescription(v string) *UpdateDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateDirectoryRequest
	GetDirectoryId() *string
	SetName(v string) *UpdateDirectoryRequest
	GetName() *string
	SetOpsUserCode(v string) *UpdateDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateDirectoryRequest
	GetOpsUserId() *string
	SetParentId(v int64) *UpdateDirectoryRequest
	GetParentId() *int64
	SetPath(v string) *UpdateDirectoryRequest
	GetPath() *string
	SetTenantId(v string) *UpdateDirectoryRequest
	GetTenantId() *string
}

type UpdateDirectoryRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// The ID of the parent node.
	//
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The path of the node.
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateDirectoryRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *UpdateDirectoryRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateDirectoryRequest) SetAppToken(v string) *UpdateDirectoryRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateDirectoryRequest) SetDescription(v string) *UpdateDirectoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateDirectoryRequest) SetDirectoryId(v string) *UpdateDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDirectoryRequest) SetName(v string) *UpdateDirectoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateDirectoryRequest) SetOpsUserCode(v string) *UpdateDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateDirectoryRequest) SetOpsUserId(v string) *UpdateDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateDirectoryRequest) SetParentId(v int64) *UpdateDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *UpdateDirectoryRequest) SetPath(v string) *UpdateDirectoryRequest {
	s.Path = &v
	return s
}

func (s *UpdateDirectoryRequest) SetTenantId(v string) *UpdateDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
