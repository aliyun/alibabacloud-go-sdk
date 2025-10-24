// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ResourceInstance) *ListResourceInstancesResponseBody
	GetInstances() []*ResourceInstance
	SetPageNumber(v int32) *ListResourceInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourceInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListResourceInstancesResponseBody
	GetTotalCount() *int32
}

type ListResourceInstancesResponseBody struct {
	// The instances.
	Instances []*ResourceInstance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
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
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponseBody) GetInstances() []*ResourceInstance {
	return s.Instances
}

func (s *ListResourceInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceInstancesResponseBody) SetInstances(v []*ResourceInstance) *ListResourceInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListResourceInstancesResponseBody) SetPageNumber(v int32) *ListResourceInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetPageSize(v int32) *ListResourceInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetRequestId(v string) *ListResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetTotalCount(v int32) *ListResourceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
