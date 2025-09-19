// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFixUsedCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeFixUsedCountResponseBody
	GetRequestId() *string
	SetUsedCount(v int32) *DescribeFixUsedCountResponseBody
	GetUsedCount() *int32
	SetUsedCountCn(v int32) *DescribeFixUsedCountResponseBody
	GetUsedCountCn() *int32
	SetUsedCountSg(v int32) *DescribeFixUsedCountResponseBody
	GetUsedCountSg() *int32
}

type DescribeFixUsedCountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1CE824-7F80-546D-8AF8-4A5209F9B698
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of vulnerabilities that are fixed by the vulnerability fixing feature.
	//
	// example:
	//
	// 10
	UsedCount *int32 `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	// The number of vulnerabilities that are fixed by the vulnerability fixing feature in China.
	//
	// example:
	//
	// 5
	UsedCountCn *int32 `json:"UsedCountCn,omitempty" xml:"UsedCountCn,omitempty"`
	// The number of vulnerabilities that are fixed by the vulnerability fixing feature outside China.
	//
	// example:
	//
	// 5
	UsedCountSg *int32 `json:"UsedCountSg,omitempty" xml:"UsedCountSg,omitempty"`
}

func (s DescribeFixUsedCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFixUsedCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFixUsedCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFixUsedCountResponseBody) GetUsedCount() *int32 {
	return s.UsedCount
}

func (s *DescribeFixUsedCountResponseBody) GetUsedCountCn() *int32 {
	return s.UsedCountCn
}

func (s *DescribeFixUsedCountResponseBody) GetUsedCountSg() *int32 {
	return s.UsedCountSg
}

func (s *DescribeFixUsedCountResponseBody) SetRequestId(v string) *DescribeFixUsedCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFixUsedCountResponseBody) SetUsedCount(v int32) *DescribeFixUsedCountResponseBody {
	s.UsedCount = &v
	return s
}

func (s *DescribeFixUsedCountResponseBody) SetUsedCountCn(v int32) *DescribeFixUsedCountResponseBody {
	s.UsedCountCn = &v
	return s
}

func (s *DescribeFixUsedCountResponseBody) SetUsedCountSg(v int32) *DescribeFixUsedCountResponseBody {
	s.UsedCountSg = &v
	return s
}

func (s *DescribeFixUsedCountResponseBody) Validate() error {
	return dara.Validate(s)
}
