// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnBlockCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemainCount(v int32) *DescribeUnBlockCountResponseBody
	GetRemainCount() *int32
	SetRequestId(v string) *DescribeUnBlockCountResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeUnBlockCountResponseBody
	GetTotalCount() *int32
}

type DescribeUnBlockCountResponseBody struct {
	// The remaining number of times that you can enable the near-origin traffic diversion feature.
	//
	// example:
	//
	// 7
	RemainCount *int32 `json:"RemainCount,omitempty" xml:"RemainCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of times that you can enable the near-origin traffic diversion feature.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUnBlockCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnBlockCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnBlockCountResponseBody) GetRemainCount() *int32 {
	return s.RemainCount
}

func (s *DescribeUnBlockCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUnBlockCountResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUnBlockCountResponseBody) SetRemainCount(v int32) *DescribeUnBlockCountResponseBody {
	s.RemainCount = &v
	return s
}

func (s *DescribeUnBlockCountResponseBody) SetRequestId(v string) *DescribeUnBlockCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnBlockCountResponseBody) SetTotalCount(v int32) *DescribeUnBlockCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUnBlockCountResponseBody) Validate() error {
	return dara.Validate(s)
}
