// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcListCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlackCount(v int32) *DescribeAutoCcListCountResponseBody
	GetBlackCount() *int32
	SetRequestId(v string) *DescribeAutoCcListCountResponseBody
	GetRequestId() *string
	SetWhiteCount(v int32) *DescribeAutoCcListCountResponseBody
	GetWhiteCount() *int32
}

type DescribeAutoCcListCountResponseBody struct {
	// The total number of IP addresses in the blacklist.
	//
	// example:
	//
	// 0
	BlackCount *int32 `json:"BlackCount,omitempty" xml:"BlackCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5AC3785F-C789-4622-87A4-F58BE7F6B184
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of IP addresses in the whitelist.
	//
	// example:
	//
	// 2
	WhiteCount *int32 `json:"WhiteCount,omitempty" xml:"WhiteCount,omitempty"`
}

func (s DescribeAutoCcListCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcListCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcListCountResponseBody) GetBlackCount() *int32 {
	return s.BlackCount
}

func (s *DescribeAutoCcListCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoCcListCountResponseBody) GetWhiteCount() *int32 {
	return s.WhiteCount
}

func (s *DescribeAutoCcListCountResponseBody) SetBlackCount(v int32) *DescribeAutoCcListCountResponseBody {
	s.BlackCount = &v
	return s
}

func (s *DescribeAutoCcListCountResponseBody) SetRequestId(v string) *DescribeAutoCcListCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoCcListCountResponseBody) SetWhiteCount(v int32) *DescribeAutoCcListCountResponseBody {
	s.WhiteCount = &v
	return s
}

func (s *DescribeAutoCcListCountResponseBody) Validate() error {
	return dara.Validate(s)
}
