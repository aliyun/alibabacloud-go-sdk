// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListOutputFilesRequest
	GetBizId() *string
	SetCollaborationGroupId(v string) *ListOutputFilesRequest
	GetCollaborationGroupId() *string
	SetItemType(v string) *ListOutputFilesRequest
	GetItemType() *string
	SetKeyword(v string) *ListOutputFilesRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListOutputFilesRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListOutputFilesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListOutputFilesRequest
	GetPageSize() *int64
	SetSessionId(v string) *ListOutputFilesRequest
	GetSessionId() *string
	SetSharedOnly(v bool) *ListOutputFilesRequest
	GetSharedOnly() *bool
	SetTags(v []*string) *ListOutputFilesRequest
	GetTags() []*string
	SetTenantId(v string) *ListOutputFilesRequest
	GetTenantId() *string
}

type ListOutputFilesRequest struct {
	// 精确业务编号，来自文件登记或工作流实例关联
	//
	// example:
	//
	// REQ-20261012-001
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 可选协作空间 ID；未传保持当前用户个人产出范围
	//
	// example:
	//
	// cg_123e4567-e89b-42d3-a456-426614174001
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The type of the output item. Valid values: ppt, html, document, picture, slides, video, audio, email, and others.
	//
	// example:
	//
	// ppt
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The keyword for searching. Matches the output title or item name.
	//
	// example:
	//
	// string_value
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The name of the digital employee (operating object). Used to filter results by name.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of items per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 精确会话 ID，覆盖普通会话、技能会话与任务 executionId
	//
	// example:
	//
	// session_123e4567-e89b-42d3-a456-426614174008
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to display only outputs and output items with sharing enabled.
	//
	// example:
	//
	// False
	SharedOnly *bool `json:"sharedOnly,omitempty" xml:"sharedOnly,omitempty"`
	// 全部匹配的标签列表，来自文件登记或实例 variables.tags
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOutputFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesRequest) GoString() string {
	return s.String()
}

func (s *ListOutputFilesRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListOutputFilesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListOutputFilesRequest) GetItemType() *string {
	return s.ItemType
}

func (s *ListOutputFilesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOutputFilesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOutputFilesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListOutputFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOutputFilesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListOutputFilesRequest) GetSharedOnly() *bool {
	return s.SharedOnly
}

func (s *ListOutputFilesRequest) GetTags() []*string {
	return s.Tags
}

func (s *ListOutputFilesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOutputFilesRequest) SetBizId(v string) *ListOutputFilesRequest {
	s.BizId = &v
	return s
}

func (s *ListOutputFilesRequest) SetCollaborationGroupId(v string) *ListOutputFilesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListOutputFilesRequest) SetItemType(v string) *ListOutputFilesRequest {
	s.ItemType = &v
	return s
}

func (s *ListOutputFilesRequest) SetKeyword(v string) *ListOutputFilesRequest {
	s.Keyword = &v
	return s
}

func (s *ListOutputFilesRequest) SetOperatingObjectName(v string) *ListOutputFilesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOutputFilesRequest) SetPage(v int64) *ListOutputFilesRequest {
	s.Page = &v
	return s
}

func (s *ListOutputFilesRequest) SetPageSize(v int64) *ListOutputFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutputFilesRequest) SetSessionId(v string) *ListOutputFilesRequest {
	s.SessionId = &v
	return s
}

func (s *ListOutputFilesRequest) SetSharedOnly(v bool) *ListOutputFilesRequest {
	s.SharedOnly = &v
	return s
}

func (s *ListOutputFilesRequest) SetTags(v []*string) *ListOutputFilesRequest {
	s.Tags = v
	return s
}

func (s *ListOutputFilesRequest) SetTenantId(v string) *ListOutputFilesRequest {
	s.TenantId = &v
	return s
}

func (s *ListOutputFilesRequest) Validate() error {
	return dara.Validate(s)
}
