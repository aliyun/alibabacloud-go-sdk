// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceDomains(v []*DescribeInstanceDomainsResponseBodyInstanceDomains) *DescribeInstanceDomainsResponseBody
	GetInstanceDomains() []*DescribeInstanceDomainsResponseBodyInstanceDomains
	SetPageNumber(v int32) *DescribeInstanceDomainsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceDomainsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceDomainsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeInstanceDomainsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeInstanceDomainsResponseBody
	GetTotalPages() *int32
}

type DescribeInstanceDomainsResponseBody struct {
	// A list of domain names attached to the instance.
	InstanceDomains []*DescribeInstanceDomainsResponseBodyInstanceDomains `json:"InstanceDomains,omitempty" xml:"InstanceDomains,omitempty" type:"Repeated"`
	// The page number. The value starts from **1**. Default: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 68386699-8B9E-4D5B-BC4C-75A28F6C2A00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 2
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeInstanceDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsResponseBody) GetInstanceDomains() []*DescribeInstanceDomainsResponseBodyInstanceDomains {
	return s.InstanceDomains
}

func (s *DescribeInstanceDomainsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceDomainsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDomainsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeInstanceDomainsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeInstanceDomainsResponseBody) SetInstanceDomains(v []*DescribeInstanceDomainsResponseBodyInstanceDomains) *DescribeInstanceDomainsResponseBody {
	s.InstanceDomains = v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetPageNumber(v int32) *DescribeInstanceDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetPageSize(v int32) *DescribeInstanceDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetRequestId(v string) *DescribeInstanceDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetTotalItems(v int32) *DescribeInstanceDomainsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) SetTotalPages(v int32) *DescribeInstanceDomainsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBody) Validate() error {
	if s.InstanceDomains != nil {
		for _, item := range s.InstanceDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceDomainsResponseBodyInstanceDomains struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 2020-03-09T02:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The UNIX timestamp that indicates when the instance was created.
	//
	// example:
	//
	// 1583720154000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.top
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeInstanceDomainsResponseBodyInstanceDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDomainsResponseBodyInstanceDomains) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) SetCreateTime(v string) *DescribeInstanceDomainsResponseBodyInstanceDomains {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) SetCreateTimestamp(v int64) *DescribeInstanceDomainsResponseBodyInstanceDomains {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) SetDomainName(v string) *DescribeInstanceDomainsResponseBodyInstanceDomains {
	s.DomainName = &v
	return s
}

func (s *DescribeInstanceDomainsResponseBodyInstanceDomains) Validate() error {
	return dara.Validate(s)
}
