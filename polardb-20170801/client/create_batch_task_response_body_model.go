// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *CreateBatchTaskResponseBody
	GetBatchId() *string
	SetRequestId(v string) *CreateBatchTaskResponseBody
	GetRequestId() *string
}

type CreateBatchTaskResponseBody struct {
	// The release batch ID.
	//
	// example:
	//
	// pcb-xxx
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C61892A4-0850-4516-9E26-44D96C1782DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskResponseBody) GetBatchId() *string {
	return s.BatchId
}

func (s *CreateBatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchTaskResponseBody) SetBatchId(v string) *CreateBatchTaskResponseBody {
	s.BatchId = &v
	return s
}

func (s *CreateBatchTaskResponseBody) SetRequestId(v string) *CreateBatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
