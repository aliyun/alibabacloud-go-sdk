// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListGroupDirectoriesRequest
	GetAppToken() *string
	SetDirectoryId(v string) *ListGroupDirectoriesRequest
	GetDirectoryId() *string
	SetGroupId(v string) *ListGroupDirectoriesRequest
	GetGroupId() *string
	SetOpsUserCode(v string) *ListGroupDirectoriesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListGroupDirectoriesRequest
	GetOpsUserId() *string
	SetSortField(v string) *ListGroupDirectoriesRequest
	GetSortField() *string
	SetSortOrder(v string) *ListGroupDirectoriesRequest
	GetSortOrder() *string
	SetTenantId(v string) *ListGroupDirectoriesRequest
	GetTenantId() *string
}

type ListGroupDirectoriesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of a visible directory in the current space. If this parameter is omitted or set to root, the space root is queried. The first query reuses the existing service-initialized internal root directory.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The ID of the collaborative share.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
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
	// The sort field. Valid values: name, gmt_create, and gmt_modified.
	//
	// example:
	//
	// gmt_create
	SortField *string `json:"sortField,omitempty" xml:"sortField,omitempty"`
	// The sort order. Valid values: asc and desc.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGroupDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListGroupDirectoriesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListGroupDirectoriesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListGroupDirectoriesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupDirectoriesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListGroupDirectoriesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListGroupDirectoriesRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListGroupDirectoriesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListGroupDirectoriesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGroupDirectoriesRequest) SetAppToken(v string) *ListGroupDirectoriesRequest {
	s.AppToken = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetDirectoryId(v string) *ListGroupDirectoriesRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetGroupId(v string) *ListGroupDirectoriesRequest {
	s.GroupId = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetOpsUserCode(v string) *ListGroupDirectoriesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetOpsUserId(v string) *ListGroupDirectoriesRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetSortField(v string) *ListGroupDirectoriesRequest {
	s.SortField = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetSortOrder(v string) *ListGroupDirectoriesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListGroupDirectoriesRequest) SetTenantId(v string) *ListGroupDirectoriesRequest {
	s.TenantId = &v
	return s
}

func (s *ListGroupDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
