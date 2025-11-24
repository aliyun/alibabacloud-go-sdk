// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReusableSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeReusableSlbResponseBody
	GetRequestId() *string
	SetReusableSlbList(v []*DescribeReusableSlbResponseBodyReusableSlbList) *DescribeReusableSlbResponseBody
	GetReusableSlbList() []*DescribeReusableSlbResponseBodyReusableSlbList
}

type DescribeReusableSlbResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EDDC0D86-2FC3-56FB-9213-96EB0A3523F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of SLB instances that can be reused.
	ReusableSlbList []*DescribeReusableSlbResponseBodyReusableSlbList `json:"ReusableSlbList,omitempty" xml:"ReusableSlbList,omitempty" type:"Repeated"`
}

func (s DescribeReusableSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReusableSlbResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReusableSlbResponseBody) GetReusableSlbList() []*DescribeReusableSlbResponseBodyReusableSlbList {
	return s.ReusableSlbList
}

func (s *DescribeReusableSlbResponseBody) SetRequestId(v string) *DescribeReusableSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReusableSlbResponseBody) SetReusableSlbList(v []*DescribeReusableSlbResponseBodyReusableSlbList) *DescribeReusableSlbResponseBody {
	s.ReusableSlbList = v
	return s
}

func (s *DescribeReusableSlbResponseBody) Validate() error {
	if s.ReusableSlbList != nil {
		for _, item := range s.ReusableSlbList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReusableSlbResponseBodyReusableSlbList struct {
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp17xygzsxth0bwlz****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the SLB instance.
	//
	// example:
	//
	// ad6044b015ac54484a20c47241175bb3
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
}

func (s DescribeReusableSlbResponseBodyReusableSlbList) String() string {
	return dara.Prettify(s)
}

func (s DescribeReusableSlbResponseBodyReusableSlbList) GoString() string {
	return s.String()
}

func (s *DescribeReusableSlbResponseBodyReusableSlbList) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeReusableSlbResponseBodyReusableSlbList) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeReusableSlbResponseBodyReusableSlbList) SetLoadBalancerId(v string) *DescribeReusableSlbResponseBodyReusableSlbList {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeReusableSlbResponseBodyReusableSlbList) SetLoadBalancerName(v string) *DescribeReusableSlbResponseBodyReusableSlbList {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeReusableSlbResponseBodyReusableSlbList) Validate() error {
	return dara.Validate(s)
}
