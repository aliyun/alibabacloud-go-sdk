// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskExecutionInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListTaskExecutionInvocationsRequest
	GetRegionId() *string
	SetStatus(v string) *ListTaskExecutionInvocationsRequest
	GetStatus() *string
	SetTaskExecutionId(v string) *ListTaskExecutionInvocationsRequest
	GetTaskExecutionId() *string
}

type ListTaskExecutionInvocationsRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// exec-a10749813b3xxxxx.t0001
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
}

func (s ListTaskExecutionInvocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionInvocationsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionInvocationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTaskExecutionInvocationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskExecutionInvocationsRequest) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *ListTaskExecutionInvocationsRequest) SetRegionId(v string) *ListTaskExecutionInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTaskExecutionInvocationsRequest) SetStatus(v string) *ListTaskExecutionInvocationsRequest {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionInvocationsRequest) SetTaskExecutionId(v string) *ListTaskExecutionInvocationsRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionInvocationsRequest) Validate() error {
	return dara.Validate(s)
}
