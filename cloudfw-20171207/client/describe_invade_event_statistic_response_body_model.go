// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHighOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody
	GetHighOpenEventCnt() *int32
	SetLowOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody
	GetLowOpenEventCnt() *int32
	SetMiddleOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody
	GetMiddleOpenEventCnt() *int32
	SetRequestId(v string) *DescribeInvadeEventStatisticResponseBody
	GetRequestId() *string
	SetTotalOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody
	GetTotalOpenEventCnt() *int32
}

type DescribeInvadeEventStatisticResponseBody struct {
	// example:
	//
	// 1
	HighOpenEventCnt *int32 `json:"HighOpenEventCnt,omitempty" xml:"HighOpenEventCnt,omitempty"`
	// example:
	//
	// 1
	LowOpenEventCnt *int32 `json:"LowOpenEventCnt,omitempty" xml:"LowOpenEventCnt,omitempty"`
	// example:
	//
	// 0
	MiddleOpenEventCnt *int32 `json:"MiddleOpenEventCnt,omitempty" xml:"MiddleOpenEventCnt,omitempty"`
	// example:
	//
	// 1530A01A-6901-5B72-AB88-28B6E96B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalOpenEventCnt *int32 `json:"TotalOpenEventCnt,omitempty" xml:"TotalOpenEventCnt,omitempty"`
}

func (s DescribeInvadeEventStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventStatisticResponseBody) GetHighOpenEventCnt() *int32 {
	return s.HighOpenEventCnt
}

func (s *DescribeInvadeEventStatisticResponseBody) GetLowOpenEventCnt() *int32 {
	return s.LowOpenEventCnt
}

func (s *DescribeInvadeEventStatisticResponseBody) GetMiddleOpenEventCnt() *int32 {
	return s.MiddleOpenEventCnt
}

func (s *DescribeInvadeEventStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvadeEventStatisticResponseBody) GetTotalOpenEventCnt() *int32 {
	return s.TotalOpenEventCnt
}

func (s *DescribeInvadeEventStatisticResponseBody) SetHighOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody {
	s.HighOpenEventCnt = &v
	return s
}

func (s *DescribeInvadeEventStatisticResponseBody) SetLowOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody {
	s.LowOpenEventCnt = &v
	return s
}

func (s *DescribeInvadeEventStatisticResponseBody) SetMiddleOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody {
	s.MiddleOpenEventCnt = &v
	return s
}

func (s *DescribeInvadeEventStatisticResponseBody) SetRequestId(v string) *DescribeInvadeEventStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvadeEventStatisticResponseBody) SetTotalOpenEventCnt(v int32) *DescribeInvadeEventStatisticResponseBody {
	s.TotalOpenEventCnt = &v
	return s
}

func (s *DescribeInvadeEventStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}
