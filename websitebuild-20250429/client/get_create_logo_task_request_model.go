// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateLogoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetCreateLogoTaskRequest
	GetTaskId() *string
}

type GetCreateLogoTaskRequest struct {
	// example:
	//
	// 20051349
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCreateLogoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCreateLogoTaskRequest) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCreateLogoTaskRequest) SetTaskId(v string) *GetCreateLogoTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetCreateLogoTaskRequest) Validate() error {
	return dara.Validate(s)
}
