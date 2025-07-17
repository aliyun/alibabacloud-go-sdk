// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUpdateTasksResponseBody
	GetRequestId() *string
	SetSuccessInfo(v map[string]*SuccessInfoValue) *BatchUpdateTasksResponseBody
	GetSuccessInfo() map[string]*SuccessInfoValue
}

type BatchUpdateTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the batch operation, which is in the MAP structure. The task ID serves as a key, and the result serves as a value.
	SuccessInfo map[string]*SuccessInfoValue `json:"SuccessInfo,omitempty" xml:"SuccessInfo,omitempty"`
}

func (s BatchUpdateTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateTasksResponseBody) GetSuccessInfo() map[string]*SuccessInfoValue {
	return s.SuccessInfo
}

func (s *BatchUpdateTasksResponseBody) SetRequestId(v string) *BatchUpdateTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateTasksResponseBody) SetSuccessInfo(v map[string]*SuccessInfoValue) *BatchUpdateTasksResponseBody {
	s.SuccessInfo = v
	return s
}

func (s *BatchUpdateTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
