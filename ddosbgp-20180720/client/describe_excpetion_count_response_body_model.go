// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExcpetionCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExceptionIpCount(v int32) *DescribeExcpetionCountResponseBody
	GetExceptionIpCount() *int32
	SetExpireTimeCount(v int32) *DescribeExcpetionCountResponseBody
	GetExpireTimeCount() *int32
	SetRequestId(v string) *DescribeExcpetionCountResponseBody
	GetRequestId() *string
}

type DescribeExcpetionCountResponseBody struct {
	// The number of abnormal IP addresses (such as those in the Black Hole Activated state) among protected assets that are assigned public IP addresses (such as a cloud server ECS instance, a Server Load Balancer (SLB) instance, and an elastic IP address (EIP)).
	//
	// example:
	//
	// 0
	ExceptionIpCount *int32 `json:"ExceptionIpCount,omitempty" xml:"ExceptionIpCount,omitempty"`
	// The number of Anti-DDoS Origin instances that are about to expire (with a remaining validity period of less than 7 days).
	//
	// example:
	//
	// 1
	ExpireTimeCount *int32 `json:"ExpireTimeCount,omitempty" xml:"ExpireTimeCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4B45279A-B1BE-5EEE-87CA-58AF4183EA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeExcpetionCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExcpetionCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcpetionCountResponseBody) GetExceptionIpCount() *int32 {
	return s.ExceptionIpCount
}

func (s *DescribeExcpetionCountResponseBody) GetExpireTimeCount() *int32 {
	return s.ExpireTimeCount
}

func (s *DescribeExcpetionCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExcpetionCountResponseBody) SetExceptionIpCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExceptionIpCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetExpireTimeCount(v int32) *DescribeExcpetionCountResponseBody {
	s.ExpireTimeCount = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) SetRequestId(v string) *DescribeExcpetionCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExcpetionCountResponseBody) Validate() error {
	return dara.Validate(s)
}
