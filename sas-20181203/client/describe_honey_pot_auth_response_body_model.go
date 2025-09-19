// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHoneyPotAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHoneyPotAuthCount(v int64) *DescribeHoneyPotAuthResponseBody
	GetHoneyPotAuthCount() *int64
	SetHoneyPotCount(v int32) *DescribeHoneyPotAuthResponseBody
	GetHoneyPotCount() *int32
	SetRequestId(v string) *DescribeHoneyPotAuthResponseBody
	GetRequestId() *string
}

type DescribeHoneyPotAuthResponseBody struct {
	// The total quota.
	//
	// example:
	//
	// 10
	HoneyPotAuthCount *int64 `json:"HoneyPotAuthCount,omitempty" xml:"HoneyPotAuthCount,omitempty"`
	// The quota that is consumed.
	//
	// example:
	//
	// 4
	HoneyPotCount *int32 `json:"HoneyPotCount,omitempty" xml:"HoneyPotCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0F5023B6-9C1F-459F-ACCC-8B4636804037
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHoneyPotAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHoneyPotAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHoneyPotAuthResponseBody) GetHoneyPotAuthCount() *int64 {
	return s.HoneyPotAuthCount
}

func (s *DescribeHoneyPotAuthResponseBody) GetHoneyPotCount() *int32 {
	return s.HoneyPotCount
}

func (s *DescribeHoneyPotAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHoneyPotAuthResponseBody) SetHoneyPotAuthCount(v int64) *DescribeHoneyPotAuthResponseBody {
	s.HoneyPotAuthCount = &v
	return s
}

func (s *DescribeHoneyPotAuthResponseBody) SetHoneyPotCount(v int32) *DescribeHoneyPotAuthResponseBody {
	s.HoneyPotCount = &v
	return s
}

func (s *DescribeHoneyPotAuthResponseBody) SetRequestId(v string) *DescribeHoneyPotAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHoneyPotAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
