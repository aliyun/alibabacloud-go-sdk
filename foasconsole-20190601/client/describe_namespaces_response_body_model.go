// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaces(v []*DescribeNamespacesResponseBodyNamespaces) *DescribeNamespacesResponseBody
	GetNamespaces() []*DescribeNamespacesResponseBodyNamespaces
	SetPageIndex(v int32) *DescribeNamespacesResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeNamespacesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNamespacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeNamespacesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeNamespacesResponseBody
	GetTotalCount() *int64
	SetTotalPage(v int32) *DescribeNamespacesResponseBody
	GetTotalPage() *int32
}

type DescribeNamespacesResponseBody struct {
	Namespaces []*DescribeNamespacesResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBody) GetNamespaces() []*DescribeNamespacesResponseBodyNamespaces {
	return s.Namespaces
}

func (s *DescribeNamespacesResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeNamespacesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNamespacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeNamespacesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeNamespacesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeNamespacesResponseBody) SetNamespaces(v []*DescribeNamespacesResponseBodyNamespaces) *DescribeNamespacesResponseBody {
	s.Namespaces = v
	return s
}

func (s *DescribeNamespacesResponseBody) SetPageIndex(v int32) *DescribeNamespacesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetPageSize(v int32) *DescribeNamespacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetRequestId(v string) *DescribeNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetSuccess(v bool) *DescribeNamespacesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetTotalCount(v int64) *DescribeNamespacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNamespacesResponseBody) SetTotalPage(v int32) *DescribeNamespacesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeNamespacesResponseBody) Validate() error {
	if s.Namespaces != nil {
		for _, item := range s.Namespaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNamespacesResponseBodyNamespaces struct {
	// example:
	//
	// 1629879567394
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1629879567394
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Ha          *bool  `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// example:
	//
	// ns-1
	Namespace    *string                                               `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	ResourceSpec *DescribeNamespacesResponseBodyNamespacesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	ResourceUsed *DescribeNamespacesResponseBodyNamespacesResourceUsed `json:"ResourceUsed,omitempty" xml:"ResourceUsed,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESS
	Status *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*DescribeNamespacesResponseBodyNamespacesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeNamespacesResponseBodyNamespaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetHa() *bool {
	return s.Ha
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetResourceSpec() *DescribeNamespacesResponseBodyNamespacesResourceSpec {
	return s.ResourceSpec
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetResourceUsed() *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	return s.ResourceUsed
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetStatus() *string {
	return s.Status
}

func (s *DescribeNamespacesResponseBodyNamespaces) GetTags() []*DescribeNamespacesResponseBodyNamespacesTags {
	return s.Tags
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetGmtCreate(v int64) *DescribeNamespacesResponseBodyNamespaces {
	s.GmtCreate = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetGmtModified(v int64) *DescribeNamespacesResponseBodyNamespaces {
	s.GmtModified = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetHa(v bool) *DescribeNamespacesResponseBodyNamespaces {
	s.Ha = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetNamespace(v string) *DescribeNamespacesResponseBodyNamespaces {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetResourceSpec(v *DescribeNamespacesResponseBodyNamespacesResourceSpec) *DescribeNamespacesResponseBodyNamespaces {
	s.ResourceSpec = v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetResourceUsed(v *DescribeNamespacesResponseBodyNamespacesResourceUsed) *DescribeNamespacesResponseBodyNamespaces {
	s.ResourceUsed = v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetStatus(v string) *DescribeNamespacesResponseBodyNamespaces {
	s.Status = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) SetTags(v []*DescribeNamespacesResponseBodyNamespacesTags) *DescribeNamespacesResponseBodyNamespaces {
	s.Tags = v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespaces) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceUsed != nil {
		if err := s.ResourceUsed.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNamespacesResponseBodyNamespacesResourceSpec struct {
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeNamespacesResponseBodyNamespacesResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespacesResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceSpec) SetCpu(v int32) *DescribeNamespacesResponseBodyNamespacesResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceSpec) SetMemoryGB(v int32) *DescribeNamespacesResponseBodyNamespacesResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespacesResponseBodyNamespacesResourceUsed struct {
	// example:
	//
	// 2.5
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 1.6
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 6.6
	MemoryGB *float32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeNamespacesResponseBodyNamespacesResourceUsed) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespacesResourceUsed) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) GetCu() *float32 {
	return s.Cu
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) GetMemoryGB() *float32 {
	return s.MemoryGB
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) SetCpu(v float32) *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	s.Cpu = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) SetCu(v float32) *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	s.Cu = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) SetMemoryGB(v float32) *DescribeNamespacesResponseBodyNamespacesResourceUsed {
	s.MemoryGB = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesResourceUsed) Validate() error {
	return dara.Validate(s)
}

type DescribeNamespacesResponseBodyNamespacesTags struct {
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNamespacesResponseBodyNamespacesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponseBodyNamespacesTags) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponseBodyNamespacesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeNamespacesResponseBodyNamespacesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeNamespacesResponseBodyNamespacesTags) SetKey(v string) *DescribeNamespacesResponseBodyNamespacesTags {
	s.Key = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesTags) SetValue(v string) *DescribeNamespacesResponseBodyNamespacesTags {
	s.Value = &v
	return s
}

func (s *DescribeNamespacesResponseBodyNamespacesTags) Validate() error {
	return dara.Validate(s)
}
