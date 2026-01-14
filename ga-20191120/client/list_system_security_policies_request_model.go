// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemSecurityPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSystemSecurityPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSystemSecurityPoliciesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListSystemSecurityPoliciesRequest
	GetRegionId() *string
}

type ListSystemSecurityPoliciesRequest struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSystemSecurityPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSystemSecurityPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSystemSecurityPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemSecurityPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSystemSecurityPoliciesRequest) SetPageNumber(v int32) *ListSystemSecurityPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSystemSecurityPoliciesRequest) SetPageSize(v int32) *ListSystemSecurityPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSystemSecurityPoliciesRequest) SetRegionId(v string) *ListSystemSecurityPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSystemSecurityPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
