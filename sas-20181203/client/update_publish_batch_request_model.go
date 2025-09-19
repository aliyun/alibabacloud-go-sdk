// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePublishBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *UpdatePublishBatchRequest
	GetBatchId() *int64
	SetBatchName(v string) *UpdatePublishBatchRequest
	GetBatchName() *string
	SetInterval(v int32) *UpdatePublishBatchRequest
	GetInterval() *int32
	SetOperationBase(v int32) *UpdatePublishBatchRequest
	GetOperationBase() *int32
}

type UpdatePublishBatchRequest struct {
	// The ID of the release batch.
	//
	// example:
	//
	// 52370
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The name of the release batch.
	//
	// example:
	//
	// xxx
	BatchName *string `json:"BatchName,omitempty" xml:"BatchName,omitempty"`
	// The interval between two release batches.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The asset selection dimension. Valid values:
	//
	// 	- **0**: instance
	//
	// 	- **1**: machine group
	//
	// 	- **2**: VPC-based instance ID
	//
	// example:
	//
	// 0
	OperationBase *int32 `json:"OperationBase,omitempty" xml:"OperationBase,omitempty"`
}

func (s UpdatePublishBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePublishBatchRequest) GoString() string {
	return s.String()
}

func (s *UpdatePublishBatchRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *UpdatePublishBatchRequest) GetBatchName() *string {
	return s.BatchName
}

func (s *UpdatePublishBatchRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdatePublishBatchRequest) GetOperationBase() *int32 {
	return s.OperationBase
}

func (s *UpdatePublishBatchRequest) SetBatchId(v int64) *UpdatePublishBatchRequest {
	s.BatchId = &v
	return s
}

func (s *UpdatePublishBatchRequest) SetBatchName(v string) *UpdatePublishBatchRequest {
	s.BatchName = &v
	return s
}

func (s *UpdatePublishBatchRequest) SetInterval(v int32) *UpdatePublishBatchRequest {
	s.Interval = &v
	return s
}

func (s *UpdatePublishBatchRequest) SetOperationBase(v int32) *UpdatePublishBatchRequest {
	s.OperationBase = &v
	return s
}

func (s *UpdatePublishBatchRequest) Validate() error {
	return dara.Validate(s)
}
