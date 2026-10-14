// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAifgResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAifgResultRequest
	GetTaskId() *string
}

type GetAifgResultRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9cd7aee5-d365-4d4e-8236-58a110a4ea09
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetAifgResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAifgResultRequest) GoString() string {
	return s.String()
}

func (s *GetAifgResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAifgResultRequest) SetTaskId(v string) *GetAifgResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetAifgResultRequest) Validate() error {
	return dara.Validate(s)
}
