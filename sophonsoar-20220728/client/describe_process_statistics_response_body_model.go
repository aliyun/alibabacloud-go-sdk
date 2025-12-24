// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v *DescribeProcessStatisticsResponseBodyMetrics) *DescribeProcessStatisticsResponseBody
	GetMetrics() *DescribeProcessStatisticsResponseBodyMetrics
	SetRequestId(v string) *DescribeProcessStatisticsResponseBody
	GetRequestId() *string
}

type DescribeProcessStatisticsResponseBody struct {
	// The data returned.
	Metrics *DescribeProcessStatisticsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4CFC0F8A-D600-5FFF-A0DF-3121C4C1B90F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProcessStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsResponseBody) GetMetrics() *DescribeProcessStatisticsResponseBodyMetrics {
	return s.Metrics
}

func (s *DescribeProcessStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProcessStatisticsResponseBody) SetMetrics(v *DescribeProcessStatisticsResponseBodyMetrics) *DescribeProcessStatisticsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribeProcessStatisticsResponseBody) SetRequestId(v string) *DescribeProcessStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBody) Validate() error {
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProcessStatisticsResponseBodyMetrics struct {
	// The number of blocked files.
	//
	// example:
	//
	// 1
	BanFileNum *int32 `json:"BanFileNum,omitempty" xml:"BanFileNum,omitempty"`
	// The number of blocked IP addresses.
	//
	// example:
	//
	// 1
	BanIpNum *int32 `json:"BanIpNum,omitempty" xml:"BanIpNum,omitempty"`
	// The number of blocked processes.
	//
	// example:
	//
	// 1
	BanProcessNum *int32 `json:"BanProcessNum,omitempty" xml:"BanProcessNum,omitempty"`
	// The number of handling tasks.
	//
	// example:
	//
	// 1
	TaskNum *int32 `json:"TaskNum,omitempty" xml:"TaskNum,omitempty"`
}

func (s DescribeProcessStatisticsResponseBodyMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessStatisticsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) GetBanFileNum() *int32 {
	return s.BanFileNum
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) GetBanIpNum() *int32 {
	return s.BanIpNum
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) GetBanProcessNum() *int32 {
	return s.BanProcessNum
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetBanFileNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.BanFileNum = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetBanIpNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.BanIpNum = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetBanProcessNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.BanProcessNum = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) SetTaskNum(v int32) *DescribeProcessStatisticsResponseBodyMetrics {
	s.TaskNum = &v
	return s
}

func (s *DescribeProcessStatisticsResponseBodyMetrics) Validate() error {
	return dara.Validate(s)
}
