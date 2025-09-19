// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPublishBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *AddPublishBatchResponseBody
	GetBatchId() *int64
	SetRequestId(v string) *AddPublishBatchResponseBody
	GetRequestId() *string
}

type AddPublishBatchResponseBody struct {
	// The ID of the release batch.
	//
	// example:
	//
	// 1
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPublishBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPublishBatchResponseBody) GoString() string {
	return s.String()
}

func (s *AddPublishBatchResponseBody) GetBatchId() *int64 {
	return s.BatchId
}

func (s *AddPublishBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPublishBatchResponseBody) SetBatchId(v int64) *AddPublishBatchResponseBody {
	s.BatchId = &v
	return s
}

func (s *AddPublishBatchResponseBody) SetRequestId(v string) *AddPublishBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPublishBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
