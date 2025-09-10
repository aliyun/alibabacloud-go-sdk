// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstanceWorkerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListResourceInstanceWorkerRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListResourceInstanceWorkerRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceInstanceWorkerRequest
	GetPageSize() *int32
	SetReady(v bool) *ListResourceInstanceWorkerRequest
	GetReady() *bool
	SetServiceName(v string) *ListResourceInstanceWorkerRequest
	GetServiceName() *string
	SetSort(v string) *ListResourceInstanceWorkerRequest
	GetSort() *string
	SetStatus(v string) *ListResourceInstanceWorkerRequest
	GetStatus() *string
	SetWorkerName(v string) *ListResourceInstanceWorkerRequest
	GetWorkerName() *string
}

type ListResourceInstanceWorkerRequest struct {
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Ready       *bool   `json:"Ready,omitempty" xml:"Ready,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The worker name.
	//
	// example:
	//
	// test-fd95xxxxx-xxxxxx
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
}

func (s ListResourceInstanceWorkerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceInstanceWorkerRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerRequest) GetOrder() *string {
	return s.Order
}

func (s *ListResourceInstanceWorkerRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceInstanceWorkerRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceInstanceWorkerRequest) GetReady() *bool {
	return s.Ready
}

func (s *ListResourceInstanceWorkerRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListResourceInstanceWorkerRequest) GetSort() *string {
	return s.Sort
}

func (s *ListResourceInstanceWorkerRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceInstanceWorkerRequest) GetWorkerName() *string {
	return s.WorkerName
}

func (s *ListResourceInstanceWorkerRequest) SetOrder(v string) *ListResourceInstanceWorkerRequest {
	s.Order = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetPageNumber(v int32) *ListResourceInstanceWorkerRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetPageSize(v int32) *ListResourceInstanceWorkerRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetReady(v bool) *ListResourceInstanceWorkerRequest {
	s.Ready = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetServiceName(v string) *ListResourceInstanceWorkerRequest {
	s.ServiceName = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetSort(v string) *ListResourceInstanceWorkerRequest {
	s.Sort = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetStatus(v string) *ListResourceInstanceWorkerRequest {
	s.Status = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetWorkerName(v string) *ListResourceInstanceWorkerRequest {
	s.WorkerName = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) Validate() error {
	return dara.Validate(s)
}
