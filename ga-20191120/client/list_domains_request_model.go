// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListDomainsRequest
	GetAcceleratorId() *string
	SetDomain(v string) *ListDomainsRequest
	GetDomain() *string
	SetPageNumber(v int32) *ListDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDomainsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDomainsRequest
	GetRegionId() *string
	SetState(v string) *ListDomainsRequest
	GetState() *string
}

type ListDomainsRequest struct {
	// The ID of the GA instance that you want to query.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The accelerated domain name that you want to query.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ICP filing status of the accelerated domain name that you want to query. Valid values:
	//
	// 	- **illegal:*	- The domain name is illegal.
	//
	// 	- **inactive:*	- The domain name has not completed ICP filing.
	//
	// 	- **active:*	- The domain name has a valid ICP filing.
	//
	// 	- **unknown:*	- The ICP filing status is unknown.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListDomainsRequest) GetDomain() *string {
	return s.Domain
}

func (s *ListDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDomainsRequest) GetState() *string {
	return s.State
}

func (s *ListDomainsRequest) SetAcceleratorId(v string) *ListDomainsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListDomainsRequest) SetDomain(v string) *ListDomainsRequest {
	s.Domain = &v
	return s
}

func (s *ListDomainsRequest) SetPageNumber(v int32) *ListDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsRequest) SetPageSize(v int32) *ListDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainsRequest) SetRegionId(v string) *ListDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDomainsRequest) SetState(v string) *ListDomainsRequest {
	s.State = &v
	return s
}

func (s *ListDomainsRequest) Validate() error {
	return dara.Validate(s)
}
