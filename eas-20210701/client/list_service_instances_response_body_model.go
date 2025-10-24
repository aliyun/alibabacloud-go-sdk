// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*Instance) *ListServiceInstancesResponseBody
	GetInstances() []*Instance
	SetPageNumber(v int32) *ListServiceInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListServiceInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListServiceInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListServiceInstancesResponseBody
	GetTotalCount() *int32
}

type ListServiceInstancesResponseBody struct {
	// The instances.
	Instances []*Instance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
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
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) GetInstances() []*Instance {
	return s.Instances
}

func (s *ListServiceInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServiceInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceInstancesResponseBody) SetInstances(v []*Instance) *ListServiceInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetPageNumber(v int32) *ListServiceInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetPageSize(v int32) *ListServiceInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int32) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstancesResponseBody) Validate() error {
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
