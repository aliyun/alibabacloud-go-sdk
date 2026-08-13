// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListScheduledTasksRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListScheduledTasksRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListScheduledTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListScheduledTasksRequest
	GetNextToken() *string
	SetPage(v int64) *ListScheduledTasksRequest
	GetPage() *int64
	SetPageSize(v int64) *ListScheduledTasksRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListScheduledTasksRequest
	GetTenantId() *string
}

type ListScheduledTasksRequest struct {
	// 协作群组 ID（如 cg_101）；传入时按群维度返回群任务（调用者需为有效群成员），未传时为个人维度（排除群任务）
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 任务名模糊搜索
	//
	// example:
	//
	// 示例关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 单页最大返回数量（1~100）；传入时优先于 pageSize
	//
	// example:
	//
	// string_value
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 翻页令牌，取上次响应返回的 nextToken；传入时优先于 page，翻页过程中请保持 maxResults 不变
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数（1~100）
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListScheduledTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListScheduledTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListScheduledTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListScheduledTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListScheduledTasksRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListScheduledTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListScheduledTasksRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListScheduledTasksRequest) SetCollaborationGroupId(v string) *ListScheduledTasksRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListScheduledTasksRequest) SetKeyword(v string) *ListScheduledTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListScheduledTasksRequest) SetMaxResults(v int32) *ListScheduledTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListScheduledTasksRequest) SetNextToken(v string) *ListScheduledTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPage(v int64) *ListScheduledTasksRequest {
	s.Page = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPageSize(v int64) *ListScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksRequest) SetTenantId(v string) *ListScheduledTasksRequest {
	s.TenantId = &v
	return s
}

func (s *ListScheduledTasksRequest) Validate() error {
	return dara.Validate(s)
}
