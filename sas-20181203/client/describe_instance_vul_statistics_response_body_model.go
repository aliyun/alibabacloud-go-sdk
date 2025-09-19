// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVulStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceVulStatisticsResponseBody
	GetRequestId() *string
	SetVulStat(v *DescribeInstanceVulStatisticsResponseBodyVulStat) *DescribeInstanceVulStatisticsResponseBody
	GetVulStat() *DescribeInstanceVulStatisticsResponseBodyVulStat
}

type DescribeInstanceVulStatisticsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of the vulnerabilities.
	VulStat *DescribeInstanceVulStatisticsResponseBodyVulStat `json:"VulStat,omitempty" xml:"VulStat,omitempty" type:"Struct"`
}

func (s DescribeInstanceVulStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVulStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVulStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceVulStatisticsResponseBody) GetVulStat() *DescribeInstanceVulStatisticsResponseBodyVulStat {
	return s.VulStat
}

func (s *DescribeInstanceVulStatisticsResponseBody) SetRequestId(v string) *DescribeInstanceVulStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceVulStatisticsResponseBody) SetVulStat(v *DescribeInstanceVulStatisticsResponseBodyVulStat) *DescribeInstanceVulStatisticsResponseBody {
	s.VulStat = v
	return s
}

func (s *DescribeInstanceVulStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceVulStatisticsResponseBodyVulStat struct {
	// The number of high-risk vulnerabilities.
	//
	// example:
	//
	// 0
	AsapCount *string `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// The number of medium-risk vulnerabilities.
	//
	// example:
	//
	// 0
	LaterCount *string `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// The number of low-risk vulnerabilities.
	//
	// example:
	//
	// 0
	NntfCount *string `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
}

func (s DescribeInstanceVulStatisticsResponseBodyVulStat) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVulStatisticsResponseBodyVulStat) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVulStatisticsResponseBodyVulStat) GetAsapCount() *string {
	return s.AsapCount
}

func (s *DescribeInstanceVulStatisticsResponseBodyVulStat) GetLaterCount() *string {
	return s.LaterCount
}

func (s *DescribeInstanceVulStatisticsResponseBodyVulStat) GetNntfCount() *string {
	return s.NntfCount
}

func (s *DescribeInstanceVulStatisticsResponseBodyVulStat) SetAsapCount(v string) *DescribeInstanceVulStatisticsResponseBodyVulStat {
	s.AsapCount = &v
	return s
}

func (s *DescribeInstanceVulStatisticsResponseBodyVulStat) SetLaterCount(v string) *DescribeInstanceVulStatisticsResponseBodyVulStat {
	s.LaterCount = &v
	return s
}

func (s *DescribeInstanceVulStatisticsResponseBodyVulStat) SetNntfCount(v string) *DescribeInstanceVulStatisticsResponseBodyVulStat {
	s.NntfCount = &v
	return s
}

func (s *DescribeInstanceVulStatisticsResponseBodyVulStat) Validate() error {
	return dara.Validate(s)
}
