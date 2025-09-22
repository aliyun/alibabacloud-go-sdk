// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DeleteDocumentRequest
	GetTaskId() *string
}

type DeleteDocumentRequest struct {
	// example:
	//
	// 74ec62f4f4e74e5882d4086a40f2b9c6
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteDocumentRequest) SetTaskId(v string) *DeleteDocumentRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteDocumentRequest) Validate() error {
	return dara.Validate(s)
}
