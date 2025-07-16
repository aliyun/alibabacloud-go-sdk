// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListVirtualResourceResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVirtualResourceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListVirtualResourceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVirtualResourceResponseBody
	GetTotalCount() *int32
	SetVirtualResources(v []*ListVirtualResourceResponseBodyVirtualResources) *ListVirtualResourceResponseBody
	GetVirtualResources() []*ListVirtualResourceResponseBodyVirtualResources
}

type ListVirtualResourceResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The virtual resource groups.
	VirtualResources []*ListVirtualResourceResponseBodyVirtualResources `json:"VirtualResources,omitempty" xml:"VirtualResources,omitempty" type:"Repeated"`
}

func (s ListVirtualResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualResourceResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVirtualResourceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirtualResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirtualResourceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVirtualResourceResponseBody) GetVirtualResources() []*ListVirtualResourceResponseBodyVirtualResources {
	return s.VirtualResources
}

func (s *ListVirtualResourceResponseBody) SetPageNumber(v int32) *ListVirtualResourceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListVirtualResourceResponseBody) SetPageSize(v int32) *ListVirtualResourceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVirtualResourceResponseBody) SetRequestId(v string) *ListVirtualResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualResourceResponseBody) SetTotalCount(v int32) *ListVirtualResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVirtualResourceResponseBody) SetVirtualResources(v []*ListVirtualResourceResponseBodyVirtualResources) *ListVirtualResourceResponseBody {
	s.VirtualResources = v
	return s
}

func (s *ListVirtualResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVirtualResourceResponseBodyVirtualResources struct {
	// The time when the virtual resource group was created.
	//
	// example:
	//
	// 2024-10-16T17:52:49Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of deployed services.
	//
	// example:
	//
	// 1
	ServiceCount *int32 `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
	// The time when the virtual resource group was last updated.
	//
	// example:
	//
	// 2024-10-16T19:52:49Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the virtual resource group.
	//
	// example:
	//
	// eas-vr-npovr28onap1xxxxxx
	VirtualResourceId *string `json:"VirtualResourceId,omitempty" xml:"VirtualResourceId,omitempty"`
	// The name of the virtual resource group.
	//
	// example:
	//
	// MyVirtualResource
	VirtualResourceName *string `json:"VirtualResourceName,omitempty" xml:"VirtualResourceName,omitempty"`
}

func (s ListVirtualResourceResponseBodyVirtualResources) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualResourceResponseBodyVirtualResources) GoString() string {
	return s.String()
}

func (s *ListVirtualResourceResponseBodyVirtualResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListVirtualResourceResponseBodyVirtualResources) GetServiceCount() *int32 {
	return s.ServiceCount
}

func (s *ListVirtualResourceResponseBodyVirtualResources) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListVirtualResourceResponseBodyVirtualResources) GetVirtualResourceId() *string {
	return s.VirtualResourceId
}

func (s *ListVirtualResourceResponseBodyVirtualResources) GetVirtualResourceName() *string {
	return s.VirtualResourceName
}

func (s *ListVirtualResourceResponseBodyVirtualResources) SetCreateTime(v string) *ListVirtualResourceResponseBodyVirtualResources {
	s.CreateTime = &v
	return s
}

func (s *ListVirtualResourceResponseBodyVirtualResources) SetServiceCount(v int32) *ListVirtualResourceResponseBodyVirtualResources {
	s.ServiceCount = &v
	return s
}

func (s *ListVirtualResourceResponseBodyVirtualResources) SetUpdateTime(v string) *ListVirtualResourceResponseBodyVirtualResources {
	s.UpdateTime = &v
	return s
}

func (s *ListVirtualResourceResponseBodyVirtualResources) SetVirtualResourceId(v string) *ListVirtualResourceResponseBodyVirtualResources {
	s.VirtualResourceId = &v
	return s
}

func (s *ListVirtualResourceResponseBodyVirtualResources) SetVirtualResourceName(v string) *ListVirtualResourceResponseBodyVirtualResources {
	s.VirtualResourceName = &v
	return s
}

func (s *ListVirtualResourceResponseBodyVirtualResources) Validate() error {
	return dara.Validate(s)
}
