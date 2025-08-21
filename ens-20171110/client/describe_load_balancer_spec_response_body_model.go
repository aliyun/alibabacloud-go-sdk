// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerSpecs(v []*DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) *DescribeLoadBalancerSpecResponseBody
	GetLoadBalancerSpecs() []*DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs
	SetPageNumber(v int32) *DescribeLoadBalancerSpecResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancerSpecResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLoadBalancerSpecResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLoadBalancerSpecResponseBody
	GetTotalCount() *int32
}

type DescribeLoadBalancerSpecResponseBody struct {
	// The specifications. Valid values:
	LoadBalancerSpecs []*DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs `json:"LoadBalancerSpecs,omitempty" xml:"LoadBalancerSpecs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadBalancerSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerSpecResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerSpecResponseBody) GetLoadBalancerSpecs() []*DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs {
	return s.LoadBalancerSpecs
}

func (s *DescribeLoadBalancerSpecResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancerSpecResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancerSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerSpecResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLoadBalancerSpecResponseBody) SetLoadBalancerSpecs(v []*DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) *DescribeLoadBalancerSpecResponseBody {
	s.LoadBalancerSpecs = v
	return s
}

func (s *DescribeLoadBalancerSpecResponseBody) SetPageNumber(v int32) *DescribeLoadBalancerSpecResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancerSpecResponseBody) SetPageSize(v int32) *DescribeLoadBalancerSpecResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancerSpecResponseBody) SetRequestId(v string) *DescribeLoadBalancerSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerSpecResponseBody) SetTotalCount(v int32) *DescribeLoadBalancerSpecResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoadBalancerSpecResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs struct {
	// The display name of the instance type.
	//
	// example:
	//
	// elb.s1.small
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The specifications of the ELB instance.
	//
	// example:
	//
	// elb.s1.small
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
}

func (s DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) SetDisplayName(v string) *DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs {
	s.DisplayName = &v
	return s
}

func (s *DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) SetLoadBalancerSpec(v string) *DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeLoadBalancerSpecResponseBodyLoadBalancerSpecs) Validate() error {
	return dara.Validate(s)
}
