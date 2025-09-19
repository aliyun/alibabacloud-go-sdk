// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulMetaCountStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppCount(v int32) *DescribeVulMetaCountStatisticsResponseBody
	GetAppCount() *int32
	SetCveCount(v int32) *DescribeVulMetaCountStatisticsResponseBody
	GetCveCount() *int32
	SetRaspDefendCount(v int32) *DescribeVulMetaCountStatisticsResponseBody
	GetRaspDefendCount() *int32
	SetRequestId(v string) *DescribeVulMetaCountStatisticsResponseBody
	GetRequestId() *string
	SetSysCount(v int32) *DescribeVulMetaCountStatisticsResponseBody
	GetSysCount() *int32
}

type DescribeVulMetaCountStatisticsResponseBody struct {
	// The number of application vulnerabilities.
	//
	// example:
	//
	// 0
	AppCount *int32 `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	// The number of Linux software vulnerabilities.
	//
	// example:
	//
	// 10
	CveCount *int32 `json:"CveCount,omitempty" xml:"CveCount,omitempty"`
	// The number of vulnerabilities that can be defended by the application protection feature.
	//
	// example:
	//
	// 10
	RaspDefendCount *int32 `json:"RaspDefendCount,omitempty" xml:"RaspDefendCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 571B2642-BF51-5BDD-906B-D2340DB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of Windows system vulnerabilities.
	//
	// example:
	//
	// 10
	SysCount *int32 `json:"SysCount,omitempty" xml:"SysCount,omitempty"`
}

func (s DescribeVulMetaCountStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulMetaCountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulMetaCountStatisticsResponseBody) GetAppCount() *int32 {
	return s.AppCount
}

func (s *DescribeVulMetaCountStatisticsResponseBody) GetCveCount() *int32 {
	return s.CveCount
}

func (s *DescribeVulMetaCountStatisticsResponseBody) GetRaspDefendCount() *int32 {
	return s.RaspDefendCount
}

func (s *DescribeVulMetaCountStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulMetaCountStatisticsResponseBody) GetSysCount() *int32 {
	return s.SysCount
}

func (s *DescribeVulMetaCountStatisticsResponseBody) SetAppCount(v int32) *DescribeVulMetaCountStatisticsResponseBody {
	s.AppCount = &v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponseBody) SetCveCount(v int32) *DescribeVulMetaCountStatisticsResponseBody {
	s.CveCount = &v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponseBody) SetRaspDefendCount(v int32) *DescribeVulMetaCountStatisticsResponseBody {
	s.RaspDefendCount = &v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponseBody) SetRequestId(v string) *DescribeVulMetaCountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponseBody) SetSysCount(v int32) *DescribeVulMetaCountStatisticsResponseBody {
	s.SysCount = &v
	return s
}

func (s *DescribeVulMetaCountStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
