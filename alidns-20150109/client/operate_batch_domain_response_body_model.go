// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBatchDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateBatchDomainResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *OperateBatchDomainResponseBody
	GetTaskId() *int64
}

type OperateBatchDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 345345
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OperateBatchDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateBatchDomainResponseBody) GoString() string {
	return s.String()
}

func (s *OperateBatchDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateBatchDomainResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *OperateBatchDomainResponseBody) SetRequestId(v string) *OperateBatchDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateBatchDomainResponseBody) SetTaskId(v int64) *OperateBatchDomainResponseBody {
	s.TaskId = &v
	return s
}

func (s *OperateBatchDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
