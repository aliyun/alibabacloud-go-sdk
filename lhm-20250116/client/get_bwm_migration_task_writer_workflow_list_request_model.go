// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationTaskWriterWorkflowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetBwmMigrationTaskWriterWorkflowListRequest
	GetInstanceId() *string
	SetPageIndex(v int32) *GetBwmMigrationTaskWriterWorkflowListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetBwmMigrationTaskWriterWorkflowListRequest
	GetPageSize() *int32
	SetWorkflowName(v string) *GetBwmMigrationTaskWriterWorkflowListRequest
	GetWorkflowName() *string
}

type GetBwmMigrationTaskWriterWorkflowListRequest struct {
	// The submit instance identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// workflow_name
	WorkflowName *string `json:"workflowName,omitempty" xml:"workflowName,omitempty"`
}

func (s GetBwmMigrationTaskWriterWorkflowListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationTaskWriterWorkflowListRequest) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) SetInstanceId(v string) *GetBwmMigrationTaskWriterWorkflowListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) SetPageIndex(v int32) *GetBwmMigrationTaskWriterWorkflowListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) SetPageSize(v int32) *GetBwmMigrationTaskWriterWorkflowListRequest {
	s.PageSize = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) SetWorkflowName(v string) *GetBwmMigrationTaskWriterWorkflowListRequest {
	s.WorkflowName = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListRequest) Validate() error {
	return dara.Validate(s)
}
