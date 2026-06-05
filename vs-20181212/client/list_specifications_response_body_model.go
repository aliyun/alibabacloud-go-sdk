// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpecificationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListSpecificationsResponseBodyItems) *ListSpecificationsResponseBody
	GetItems() []*ListSpecificationsResponseBodyItems
	SetPageNumber(v int32) *ListSpecificationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSpecificationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSpecificationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListSpecificationsResponseBody
	GetTotalCount() *int64
}

type ListSpecificationsResponseBody struct {
	Items []*ListSpecificationsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSpecificationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpecificationsResponseBody) GetItems() []*ListSpecificationsResponseBodyItems {
	return s.Items
}

func (s *ListSpecificationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSpecificationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSpecificationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSpecificationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSpecificationsResponseBody) SetItems(v []*ListSpecificationsResponseBodyItems) *ListSpecificationsResponseBody {
	s.Items = v
	return s
}

func (s *ListSpecificationsResponseBody) SetPageNumber(v int32) *ListSpecificationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSpecificationsResponseBody) SetPageSize(v int32) *ListSpecificationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSpecificationsResponseBody) SetRequestId(v string) *ListSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpecificationsResponseBody) SetTotalCount(v int64) *ListSpecificationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSpecificationsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSpecificationsResponseBodyItems struct {
	// example:
	//
	// 4090CU
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// example:
	//
	// 8
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 1
	Shard *int32 `json:"Shard,omitempty" xml:"Shard,omitempty"`
	// `RenderingSpec`
	//
	// example:
	//
	// crs.xic.s1
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// example:
	//
	// 90
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s ListSpecificationsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSpecificationsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSpecificationsResponseBodyItems) GetClass() *string {
	return s.Class
}

func (s *ListSpecificationsResponseBodyItems) GetCores() *int32 {
	return s.Cores
}

func (s *ListSpecificationsResponseBodyItems) GetMemory() *int32 {
	return s.Memory
}

func (s *ListSpecificationsResponseBodyItems) GetShard() *int32 {
	return s.Shard
}

func (s *ListSpecificationsResponseBodyItems) GetSpecification() *string {
	return s.Specification
}

func (s *ListSpecificationsResponseBodyItems) GetStorage() *int32 {
	return s.Storage
}

func (s *ListSpecificationsResponseBodyItems) SetClass(v string) *ListSpecificationsResponseBodyItems {
	s.Class = &v
	return s
}

func (s *ListSpecificationsResponseBodyItems) SetCores(v int32) *ListSpecificationsResponseBodyItems {
	s.Cores = &v
	return s
}

func (s *ListSpecificationsResponseBodyItems) SetMemory(v int32) *ListSpecificationsResponseBodyItems {
	s.Memory = &v
	return s
}

func (s *ListSpecificationsResponseBodyItems) SetShard(v int32) *ListSpecificationsResponseBodyItems {
	s.Shard = &v
	return s
}

func (s *ListSpecificationsResponseBodyItems) SetSpecification(v string) *ListSpecificationsResponseBodyItems {
	s.Specification = &v
	return s
}

func (s *ListSpecificationsResponseBodyItems) SetStorage(v int32) *ListSpecificationsResponseBodyItems {
	s.Storage = &v
	return s
}

func (s *ListSpecificationsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
