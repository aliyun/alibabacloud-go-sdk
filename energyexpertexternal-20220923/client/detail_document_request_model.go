// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DetailDocumentRequest
	GetTaskId() *string
}

type DetailDocumentRequest struct {
	// example:
	//
	// 74ec62f4f4e74e5882d4086a40f2b9c6
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DetailDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DetailDocumentRequest) GoString() string {
	return s.String()
}

func (s *DetailDocumentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DetailDocumentRequest) SetTaskId(v string) *DetailDocumentRequest {
	s.TaskId = &v
	return s
}

func (s *DetailDocumentRequest) Validate() error {
	return dara.Validate(s)
}
