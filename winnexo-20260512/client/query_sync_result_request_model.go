// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *QuerySyncResultRequest
	GetTaskId() *int64
	SetTenantId(v string) *QuerySyncResultRequest
	GetTenantId() *string
}

type QuerySyncResultRequest struct {
	// 同步任务 ID（由 syncOrgStructure 返回）
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QuerySyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncResultRequest) GoString() string {
	return s.String()
}

func (s *QuerySyncResultRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QuerySyncResultRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySyncResultRequest) SetTaskId(v int64) *QuerySyncResultRequest {
	s.TaskId = &v
	return s
}

func (s *QuerySyncResultRequest) SetTenantId(v string) *QuerySyncResultRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySyncResultRequest) Validate() error {
	return dara.Validate(s)
}
