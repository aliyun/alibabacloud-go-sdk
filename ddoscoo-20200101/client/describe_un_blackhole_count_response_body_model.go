// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnBlackholeCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemainCount(v int32) *DescribeUnBlackholeCountResponseBody
	GetRemainCount() *int32
	SetRequestId(v string) *DescribeUnBlackholeCountResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeUnBlackholeCountResponseBody
	GetTotalCount() *int32
}

type DescribeUnBlackholeCountResponseBody struct {
	// The remaining quota that you can deactivate blackhole filtering.
	//
	// example:
	//
	// 5
	RemainCount *int32 `json:"RemainCount,omitempty" xml:"RemainCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 232929FA-40B6-4C53-9476-EE335ABA44CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total quota that you can deactivate blackhole filtering.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUnBlackholeCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnBlackholeCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnBlackholeCountResponseBody) GetRemainCount() *int32 {
	return s.RemainCount
}

func (s *DescribeUnBlackholeCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUnBlackholeCountResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUnBlackholeCountResponseBody) SetRemainCount(v int32) *DescribeUnBlackholeCountResponseBody {
	s.RemainCount = &v
	return s
}

func (s *DescribeUnBlackholeCountResponseBody) SetRequestId(v string) *DescribeUnBlackholeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnBlackholeCountResponseBody) SetTotalCount(v int32) *DescribeUnBlackholeCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUnBlackholeCountResponseBody) Validate() error {
	return dara.Validate(s)
}
