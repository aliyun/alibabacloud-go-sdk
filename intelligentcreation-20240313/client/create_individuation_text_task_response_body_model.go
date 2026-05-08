// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndividuationTextTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateIndividuationTextTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateIndividuationTextTaskResponseBody
	GetTaskId() *string
}

type CreateIndividuationTextTaskResponseBody struct {
	// example:
	//
	// 56AC346B-AF40-5E4F-AFFE-FD8BA5E6FB3A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 837091359375048704
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateIndividuationTextTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIndividuationTextTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndividuationTextTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIndividuationTextTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateIndividuationTextTaskResponseBody) SetRequestId(v string) *CreateIndividuationTextTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndividuationTextTaskResponseBody) SetTaskId(v string) *CreateIndividuationTextTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateIndividuationTextTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
