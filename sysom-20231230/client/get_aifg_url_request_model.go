// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAifgUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAifgUrlRequest
	GetTaskId() *string
}

type GetAifgUrlRequest struct {
	// The task identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9cd7aee5-d365-4d4e-8236-58a110a4ea09
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetAifgUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAifgUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAifgUrlRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAifgUrlRequest) SetTaskId(v string) *GetAifgUrlRequest {
	s.TaskId = &v
	return s
}

func (s *GetAifgUrlRequest) Validate() error {
	return dara.Validate(s)
}
