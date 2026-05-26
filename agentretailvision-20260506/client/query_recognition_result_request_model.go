// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognitionResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderUniqueId(v string) *QueryRecognitionResultRequest
	GetOrderUniqueId() *string
	SetTaskId(v string) *QueryRecognitionResultRequest
	GetTaskId() *string
}

type QueryRecognitionResultRequest struct {
	// example:
	//
	// ORDER_001
	OrderUniqueId *string `json:"OrderUniqueId,omitempty" xml:"OrderUniqueId,omitempty"`
	// example:
	//
	// TASK_001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryRecognitionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognitionResultRequest) GoString() string {
	return s.String()
}

func (s *QueryRecognitionResultRequest) GetOrderUniqueId() *string {
	return s.OrderUniqueId
}

func (s *QueryRecognitionResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryRecognitionResultRequest) SetOrderUniqueId(v string) *QueryRecognitionResultRequest {
	s.OrderUniqueId = &v
	return s
}

func (s *QueryRecognitionResultRequest) SetTaskId(v string) *QueryRecognitionResultRequest {
	s.TaskId = &v
	return s
}

func (s *QueryRecognitionResultRequest) Validate() error {
	return dara.Validate(s)
}
