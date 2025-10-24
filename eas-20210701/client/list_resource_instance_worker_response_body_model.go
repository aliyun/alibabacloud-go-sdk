// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstanceWorkerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourceInstanceWorkerResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceInstanceWorkerResponseBody
	GetPageSize() *int32
	SetPods(v []*ResourceInstanceWorker) *ListResourceInstanceWorkerResponseBody
	GetPods() []*ResourceInstanceWorker
	SetRequestId(v string) *ListResourceInstanceWorkerResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListResourceInstanceWorkerResponseBody
	GetTotalCount() *int32
}

type ListResourceInstanceWorkerResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workers.
	Pods []*ResourceInstanceWorker `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceInstanceWorkerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceInstanceWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceInstanceWorkerResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceInstanceWorkerResponseBody) GetPods() []*ResourceInstanceWorker {
	return s.Pods
}

func (s *ListResourceInstanceWorkerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceInstanceWorkerResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceInstanceWorkerResponseBody) SetPageNumber(v int32) *ListResourceInstanceWorkerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetPageSize(v int32) *ListResourceInstanceWorkerResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetPods(v []*ResourceInstanceWorker) *ListResourceInstanceWorkerResponseBody {
	s.Pods = v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetRequestId(v string) *ListResourceInstanceWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetTotalCount(v int32) *ListResourceInstanceWorkerResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) Validate() error {
	if s.Pods != nil {
		for _, item := range s.Pods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
