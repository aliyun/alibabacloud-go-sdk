// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAITaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseAITaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CloseAITaskResponseBody
	GetTaskId() *string
}

type CloseAITaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 936C7025-27A5-4CB1-BB31-540E1F0CCA12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task for disabling the PolarDB for AI feature.
	//
	// example:
	//
	// 53879cdb-9a00-428e-acaf-ff4cff******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CloseAITaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseAITaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloseAITaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseAITaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CloseAITaskResponseBody) SetRequestId(v string) *CloseAITaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseAITaskResponseBody) SetTaskId(v string) *CloseAITaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CloseAITaskResponseBody) Validate() error {
	return dara.Validate(s)
}
