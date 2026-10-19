// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListOutputFilesShrinkRequest
	GetBizId() *string
	SetCollaborationGroupId(v string) *ListOutputFilesShrinkRequest
	GetCollaborationGroupId() *string
	SetItemType(v string) *ListOutputFilesShrinkRequest
	GetItemType() *string
	SetKeyword(v string) *ListOutputFilesShrinkRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListOutputFilesShrinkRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListOutputFilesShrinkRequest
	GetPage() *int64
	SetPageSize(v int64) *ListOutputFilesShrinkRequest
	GetPageSize() *int64
	SetSessionId(v string) *ListOutputFilesShrinkRequest
	GetSessionId() *string
	SetSharedOnly(v bool) *ListOutputFilesShrinkRequest
	GetSharedOnly() *bool
	SetTagsShrink(v string) *ListOutputFilesShrinkRequest
	GetTagsShrink() *string
	SetTenantId(v string) *ListOutputFilesShrinkRequest
	GetTenantId() *string
}

type ListOutputFilesShrinkRequest struct {
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
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOutputFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOutputFilesShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListOutputFilesShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListOutputFilesShrinkRequest) GetItemType() *string {
	return s.ItemType
}

func (s *ListOutputFilesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOutputFilesShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOutputFilesShrinkRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListOutputFilesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOutputFilesShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListOutputFilesShrinkRequest) GetSharedOnly() *bool {
	return s.SharedOnly
}

func (s *ListOutputFilesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListOutputFilesShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOutputFilesShrinkRequest) SetBizId(v string) *ListOutputFilesShrinkRequest {
	s.BizId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetCollaborationGroupId(v string) *ListOutputFilesShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetItemType(v string) *ListOutputFilesShrinkRequest {
	s.ItemType = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetKeyword(v string) *ListOutputFilesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetOperatingObjectName(v string) *ListOutputFilesShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetPage(v int64) *ListOutputFilesShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetPageSize(v int64) *ListOutputFilesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetSessionId(v string) *ListOutputFilesShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetSharedOnly(v bool) *ListOutputFilesShrinkRequest {
	s.SharedOnly = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetTagsShrink(v string) *ListOutputFilesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) SetTenantId(v string) *ListOutputFilesShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListOutputFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
