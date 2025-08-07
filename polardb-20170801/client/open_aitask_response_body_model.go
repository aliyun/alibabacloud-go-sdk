// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAITaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenAITaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *OpenAITaskResponseBody
	GetTaskId() *string
}

type OpenAITaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 53879cdb-9a00-428e-acaf-ff4cff******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OpenAITaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenAITaskResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAITaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenAITaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *OpenAITaskResponseBody) SetRequestId(v string) *OpenAITaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenAITaskResponseBody) SetTaskId(v string) *OpenAITaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *OpenAITaskResponseBody) Validate() error {
	return dara.Validate(s)
}
