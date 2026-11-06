// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListCollaborationGroupsRequest
	GetKeyword() *string
	SetPage(v int64) *ListCollaborationGroupsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupsRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListCollaborationGroupsRequest
	GetTenantId() *string
}

type ListCollaborationGroupsRequest struct {
	// 空间名称搜索关键词，可选，最长256字符；不传或去除首尾空白后为空时不按名称过滤。
	//
	// example:
	//
	// 工作流
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

func (s ListCollaborationGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCollaborationGroupsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListCollaborationGroupsRequest) SetKeyword(v string) *ListCollaborationGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetPage(v int64) *ListCollaborationGroupsRequest {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetPageSize(v int64) *ListCollaborationGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetTenantId(v string) *ListCollaborationGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *ListCollaborationGroupsRequest) Validate() error {
	return dara.Validate(s)
}
