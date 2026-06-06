// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskExecutionStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskCode(v string) *GetTaskExecutionStatisticsRequest
	GetTaskCode() *string
}

type GetTaskExecutionStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// EssayCorrectionTask
	TaskCode *string `json:"taskCode,omitempty" xml:"taskCode,omitempty"`
}

func (s GetTaskExecutionStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskExecutionStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetTaskExecutionStatisticsRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *GetTaskExecutionStatisticsRequest) SetTaskCode(v string) *GetTaskExecutionStatisticsRequest {
	s.TaskCode = &v
	return s
}

func (s *GetTaskExecutionStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
