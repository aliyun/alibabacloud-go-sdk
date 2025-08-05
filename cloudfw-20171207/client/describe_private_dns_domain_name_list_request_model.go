// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsDomainNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DescribePrivateDnsDomainNameListRequest
	GetAccessInstanceId() *string
	SetDomainName(v string) *DescribePrivateDnsDomainNameListRequest
	GetDomainName() *string
	SetPageNo(v int32) *DescribePrivateDnsDomainNameListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribePrivateDnsDomainNameListRequest
	GetPageSize() *int32
	SetRegionNo(v string) *DescribePrivateDnsDomainNameListRequest
	GetRegionNo() *string
}

type DescribePrivateDnsDomainNameListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pd-12345
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribePrivateDnsDomainNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsDomainNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsDomainNameListRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DescribePrivateDnsDomainNameListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePrivateDnsDomainNameListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribePrivateDnsDomainNameListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePrivateDnsDomainNameListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePrivateDnsDomainNameListRequest) SetAccessInstanceId(v string) *DescribePrivateDnsDomainNameListRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListRequest) SetDomainName(v string) *DescribePrivateDnsDomainNameListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListRequest) SetPageNo(v int32) *DescribePrivateDnsDomainNameListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListRequest) SetPageSize(v int32) *DescribePrivateDnsDomainNameListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListRequest) SetRegionNo(v string) *DescribePrivateDnsDomainNameListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListRequest) Validate() error {
	return dara.Validate(s)
}
