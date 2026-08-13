// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetScheduledTaskExecutionRecordsRequest
	GetCollaborationGroupId() *string
	SetPage(v int32) *GetScheduledTaskExecutionRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *GetScheduledTaskExecutionRecordsRequest
	GetPageSize() *int32
	SetTenantId(v string) *GetScheduledTaskExecutionRecordsRequest
	GetTenantId() *string
}

type GetScheduledTaskExecutionRecordsRequest struct {
	// 协作群组 ID（如 cg_101）；传入时按群维度返回（调用者需为有效群成员），未传时为个人维度（排除群任务）
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 页码，从1开始
	//
	// example:
	//
	// exampleCollaborationGroupId
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页任务数（1~100）
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScheduledTaskExecutionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetCollaborationGroupId(v string) *GetScheduledTaskExecutionRecordsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetPage(v int32) *GetScheduledTaskExecutionRecordsRequest {
	s.Page = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetPageSize(v int32) *GetScheduledTaskExecutionRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetTenantId(v string) *GetScheduledTaskExecutionRecordsRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
