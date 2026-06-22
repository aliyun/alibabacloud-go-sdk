// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetByBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *ListTargetByBatchRequest
	GetBatchId() *int64
	SetCurrentPage(v int32) *ListTargetByBatchRequest
	GetCurrentPage() *int32
	SetOperationBase(v int32) *ListTargetByBatchRequest
	GetOperationBase() *int32
	SetPageSize(v int32) *ListTargetByBatchRequest
	GetPageSize() *int32
}

type ListTargetByBatchRequest struct {
	// The publish batch ID.
	//
	// example:
	//
	// 1371
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The page number of the current page to display in a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The asset selection dimension. Valid values:
	//
	// - **0**: machine instance
	//
	// - **1**: machine group
	//
	// - **2**: VPC-connected instance ID
	//
	// example:
	//
	// 0
	OperationBase *int32 `json:"OperationBase,omitempty" xml:"OperationBase,omitempty"`
	// The maximum number of entries to display per page in a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTargetByBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTargetByBatchRequest) GoString() string {
	return s.String()
}

func (s *ListTargetByBatchRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ListTargetByBatchRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTargetByBatchRequest) GetOperationBase() *int32 {
	return s.OperationBase
}

func (s *ListTargetByBatchRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTargetByBatchRequest) SetBatchId(v int64) *ListTargetByBatchRequest {
	s.BatchId = &v
	return s
}

func (s *ListTargetByBatchRequest) SetCurrentPage(v int32) *ListTargetByBatchRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTargetByBatchRequest) SetOperationBase(v int32) *ListTargetByBatchRequest {
	s.OperationBase = &v
	return s
}

func (s *ListTargetByBatchRequest) SetPageSize(v int32) *ListTargetByBatchRequest {
	s.PageSize = &v
	return s
}

func (s *ListTargetByBatchRequest) Validate() error {
	return dara.Validate(s)
}
