// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListCollaborationGroupMembersRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListCollaborationGroupMembersRequest
	GetKeyword() *string
	SetPage(v int64) *ListCollaborationGroupMembersRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupMembersRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListCollaborationGroupMembersRequest
	GetTenantId() *string
}

type ListCollaborationGroupMembersRequest struct {
	// 协作空间ID，必填，长度1至64。调用者必须是该租户下该空间的有效成员。可从ListCollaborationGroups返回的items中获取。
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 成员显示名称搜索关键词，可选，最长256字符；忽略大小写，先筛选后分页。不传时查询全部有效成员。
	//
	// example:
	//
	// 销售
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 页码，从1开始；请求省略时默认为1。
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页记录数，取值1至100；请求省略时默认为20。
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 业务租户ID，可选。必须属于当前调用身份的可访问租户；无权访问时拒绝请求。未传时按现有鉴权规则解析租户，多租户用户无法唯一确定时需显式指定。
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListCollaborationGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupMembersRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListCollaborationGroupMembersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCollaborationGroupMembersRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupMembersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupMembersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListCollaborationGroupMembersRequest) SetCollaborationGroupId(v string) *ListCollaborationGroupMembersRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetKeyword(v string) *ListCollaborationGroupMembersRequest {
	s.Keyword = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetPage(v int64) *ListCollaborationGroupMembersRequest {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetPageSize(v int64) *ListCollaborationGroupMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetTenantId(v string) *ListCollaborationGroupMembersRequest {
	s.TenantId = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
