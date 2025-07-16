// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstanceWorkerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourceInstanceWorkerRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceInstanceWorkerRequest
	GetPageSize() *int32
	SetWorkerName(v string) *ListResourceInstanceWorkerRequest
	GetWorkerName() *string
}

type ListResourceInstanceWorkerRequest struct {
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListResourceInstanceWorkerRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceInstanceWorkerRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceInstanceWorkerRequest) GetWorkerName() *string {
	return s.WorkerName
}

func (s *ListResourceInstanceWorkerRequest) SetPageNumber(v int32) *ListResourceInstanceWorkerRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetPageSize(v int32) *ListResourceInstanceWorkerRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetWorkerName(v string) *ListResourceInstanceWorkerRequest {
	s.WorkerName = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) Validate() error {
	return dara.Validate(s)
}
