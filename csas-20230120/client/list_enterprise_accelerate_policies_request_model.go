// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAcceleratePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListEnterpriseAcceleratePoliciesRequest
	GetCurrentPage() *int32
	SetName(v string) *ListEnterpriseAcceleratePoliciesRequest
	GetName() *string
	SetPageSize(v int32) *ListEnterpriseAcceleratePoliciesRequest
	GetPageSize() *int32
}

type ListEnterpriseAcceleratePoliciesRequest struct {
	// The number of the page to return for a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the enterprise acceleration policy.
	//
	// example:
	//
	// 测试策略
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page for a paged query. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListEnterpriseAcceleratePoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAcceleratePoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAcceleratePoliciesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListEnterpriseAcceleratePoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListEnterpriseAcceleratePoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnterpriseAcceleratePoliciesRequest) SetCurrentPage(v int32) *ListEnterpriseAcceleratePoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesRequest) SetName(v string) *ListEnterpriseAcceleratePoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesRequest) SetPageSize(v int32) *ListEnterpriseAcceleratePoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnterpriseAcceleratePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
