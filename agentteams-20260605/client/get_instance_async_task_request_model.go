// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceAsyncTaskRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *GetInstanceAsyncTaskRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetInstanceAsyncTaskRequest
	GetNextToken() *string
	SetStatus(v string) *GetInstanceAsyncTaskRequest
	GetStatus() *string
	SetTaskCode(v string) *GetInstanceAsyncTaskRequest
	GetTaskCode() *string
}

type GetInstanceAsyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// at-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// LIFECYCLE_MAGIC_PAY_ORDER_CALLBACK_CREATE
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
}

func (s GetInstanceAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceAsyncTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceAsyncTaskRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetInstanceAsyncTaskRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetInstanceAsyncTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceAsyncTaskRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *GetInstanceAsyncTaskRequest) SetInstanceId(v string) *GetInstanceAsyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceAsyncTaskRequest) SetMaxResults(v int32) *GetInstanceAsyncTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *GetInstanceAsyncTaskRequest) SetNextToken(v string) *GetInstanceAsyncTaskRequest {
	s.NextToken = &v
	return s
}

func (s *GetInstanceAsyncTaskRequest) SetStatus(v string) *GetInstanceAsyncTaskRequest {
	s.Status = &v
	return s
}

func (s *GetInstanceAsyncTaskRequest) SetTaskCode(v string) *GetInstanceAsyncTaskRequest {
	s.TaskCode = &v
	return s
}

func (s *GetInstanceAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
