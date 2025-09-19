// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDefendCountStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRaspDefendedCount(v int32) *DescribeVulDefendCountStatisticsResponseBody
	GetRaspDefendedCount() *int32
	SetRaspDefensibleCount(v int32) *DescribeVulDefendCountStatisticsResponseBody
	GetRaspDefensibleCount() *int32
	SetRequestId(v string) *DescribeVulDefendCountStatisticsResponseBody
	GetRequestId() *string
}

type DescribeVulDefendCountStatisticsResponseBody struct {
	// The number of defended vulnerabilities.
	//
	// example:
	//
	// 10
	RaspDefendedCount *int32 `json:"RaspDefendedCount,omitempty" xml:"RaspDefendedCount,omitempty"`
	// The number of supported vulnerabilities.
	//
	// example:
	//
	// 100
	RaspDefensibleCount *int32 `json:"RaspDefensibleCount,omitempty" xml:"RaspDefensibleCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVulDefendCountStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDefendCountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulDefendCountStatisticsResponseBody) GetRaspDefendedCount() *int32 {
	return s.RaspDefendedCount
}

func (s *DescribeVulDefendCountStatisticsResponseBody) GetRaspDefensibleCount() *int32 {
	return s.RaspDefensibleCount
}

func (s *DescribeVulDefendCountStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulDefendCountStatisticsResponseBody) SetRaspDefendedCount(v int32) *DescribeVulDefendCountStatisticsResponseBody {
	s.RaspDefendedCount = &v
	return s
}

func (s *DescribeVulDefendCountStatisticsResponseBody) SetRaspDefensibleCount(v int32) *DescribeVulDefendCountStatisticsResponseBody {
	s.RaspDefensibleCount = &v
	return s
}

func (s *DescribeVulDefendCountStatisticsResponseBody) SetRequestId(v string) *DescribeVulDefendCountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulDefendCountStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
