// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsDomainNameListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNameList(v []*string) *DescribePrivateDnsDomainNameListResponseBody
	GetDomainNameList() []*string
	SetPageNo(v int32) *DescribePrivateDnsDomainNameListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribePrivateDnsDomainNameListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePrivateDnsDomainNameListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePrivateDnsDomainNameListResponseBody
	GetTotalCount() *int64
}

type DescribePrivateDnsDomainNameListResponseBody struct {
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 822B9125-6E1A-551C-8EAF-6E7AE74****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 132
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePrivateDnsDomainNameListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsDomainNameListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsDomainNameListResponseBody) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *DescribePrivateDnsDomainNameListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribePrivateDnsDomainNameListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePrivateDnsDomainNameListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrivateDnsDomainNameListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePrivateDnsDomainNameListResponseBody) SetDomainNameList(v []*string) *DescribePrivateDnsDomainNameListResponseBody {
	s.DomainNameList = v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponseBody) SetPageNo(v int32) *DescribePrivateDnsDomainNameListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponseBody) SetPageSize(v int32) *DescribePrivateDnsDomainNameListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponseBody) SetRequestId(v string) *DescribePrivateDnsDomainNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponseBody) SetTotalCount(v int64) *DescribePrivateDnsDomainNameListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePrivateDnsDomainNameListResponseBody) Validate() error {
	return dara.Validate(s)
}
