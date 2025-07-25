// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferDomainResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *TransferDomainResponseBody
	GetTaskId() *int64
}

type TransferDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 112233
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TransferDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferDomainResponseBody) GoString() string {
	return s.String()
}

func (s *TransferDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferDomainResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TransferDomainResponseBody) SetRequestId(v string) *TransferDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferDomainResponseBody) SetTaskId(v int64) *TransferDomainResponseBody {
	s.TaskId = &v
	return s
}

func (s *TransferDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
