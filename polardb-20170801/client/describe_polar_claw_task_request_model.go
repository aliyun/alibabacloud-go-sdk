// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawTaskRequest
	GetApplicationId() *string
	SetTaskId(v string) *DescribePolarClawTaskRequest
	GetTaskId() *string
}

type DescribePolarClawTaskRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The asynchronous task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribePolarClawTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawTaskRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePolarClawTaskRequest) SetApplicationId(v string) *DescribePolarClawTaskRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawTaskRequest) SetTaskId(v string) *DescribePolarClawTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribePolarClawTaskRequest) Validate() error {
	return dara.Validate(s)
}
