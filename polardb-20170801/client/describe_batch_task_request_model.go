// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *DescribeBatchTaskRequest
	GetBatchId() *string
}

type DescribeBatchTaskRequest struct {
	// The ID of the batch task.
	//
	// example:
	//
	// pcb-xxx
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s DescribeBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeBatchTaskRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *DescribeBatchTaskRequest) SetBatchId(v string) *DescribeBatchTaskRequest {
	s.BatchId = &v
	return s
}

func (s *DescribeBatchTaskRequest) Validate() error {
	return dara.Validate(s)
}
