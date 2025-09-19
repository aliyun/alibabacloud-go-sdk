// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinarySecurityPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeBinarySecurityPoliciesRequest
	GetCurrentPage() *int32
	SetName(v string) *DescribeBinarySecurityPoliciesRequest
	GetName() *string
	SetPageSize(v int32) *DescribeBinarySecurityPoliciesRequest
	GetPageSize() *int32
	SetResourceOwnerId(v int64) *DescribeBinarySecurityPoliciesRequest
	GetResourceOwnerId() *int64
	SetSourceIp(v string) *DescribeBinarySecurityPoliciesRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeBinarySecurityPoliciesRequest
	GetStatus() *string
}

type DescribeBinarySecurityPoliciesRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// policy-auto-5patxz
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 59.82.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The status of the policy. Valid values:
	//
	// 	- **enabled**
	//
	// 	- **disabled**
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBinarySecurityPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinarySecurityPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBinarySecurityPoliciesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBinarySecurityPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeBinarySecurityPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBinarySecurityPoliciesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBinarySecurityPoliciesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeBinarySecurityPoliciesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeBinarySecurityPoliciesRequest) SetCurrentPage(v int32) *DescribeBinarySecurityPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesRequest) SetName(v string) *DescribeBinarySecurityPoliciesRequest {
	s.Name = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesRequest) SetPageSize(v int32) *DescribeBinarySecurityPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesRequest) SetResourceOwnerId(v int64) *DescribeBinarySecurityPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesRequest) SetSourceIp(v string) *DescribeBinarySecurityPoliciesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesRequest) SetStatus(v string) *DescribeBinarySecurityPoliciesRequest {
	s.Status = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
