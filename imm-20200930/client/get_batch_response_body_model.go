// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatch(v *DataIngestion) *GetBatchResponseBody
	GetBatch() *DataIngestion
	SetRequestId(v string) *GetBatchResponseBody
	GetRequestId() *string
}

type GetBatchResponseBody struct {
	// The information about the batch processing task.
	Batch *DataIngestion `json:"Batch,omitempty" xml:"Batch,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchResponseBody) GetBatch() *DataIngestion {
	return s.Batch
}

func (s *GetBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchResponseBody) SetBatch(v *DataIngestion) *GetBatchResponseBody {
	s.Batch = v
	return s
}

func (s *GetBatchResponseBody) SetRequestId(v string) *GetBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchResponseBody) Validate() error {
	if s.Batch != nil {
		if err := s.Batch.Validate(); err != nil {
			return err
		}
	}
	return nil
}
