// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterVulStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeClusterVulStatisticsResponseBody
	GetRequestId() *string
	SetVulStat(v *DescribeClusterVulStatisticsResponseBodyVulStat) *DescribeClusterVulStatisticsResponseBody
	GetVulStat() *DescribeClusterVulStatisticsResponseBodyVulStat
}

type DescribeClusterVulStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics of the vulnerabilities.
	VulStat *DescribeClusterVulStatisticsResponseBodyVulStat `json:"VulStat,omitempty" xml:"VulStat,omitempty" type:"Struct"`
}

func (s DescribeClusterVulStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterVulStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterVulStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterVulStatisticsResponseBody) GetVulStat() *DescribeClusterVulStatisticsResponseBodyVulStat {
	return s.VulStat
}

func (s *DescribeClusterVulStatisticsResponseBody) SetRequestId(v string) *DescribeClusterVulStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterVulStatisticsResponseBody) SetVulStat(v *DescribeClusterVulStatisticsResponseBodyVulStat) *DescribeClusterVulStatisticsResponseBody {
	s.VulStat = v
	return s
}

func (s *DescribeClusterVulStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterVulStatisticsResponseBodyVulStat struct {
	// The number of high-risk vulnerabilities.
	//
	// example:
	//
	// 13
	AsapCount *string `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// The number of medium-risk vulnerabilities.
	//
	// example:
	//
	// 21
	LaterCount *string `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// The number of low-risk vulnerabilities.
	//
	// example:
	//
	// 0
	NntfCount *string `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
}

func (s DescribeClusterVulStatisticsResponseBodyVulStat) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterVulStatisticsResponseBodyVulStat) GoString() string {
	return s.String()
}

func (s *DescribeClusterVulStatisticsResponseBodyVulStat) GetAsapCount() *string {
	return s.AsapCount
}

func (s *DescribeClusterVulStatisticsResponseBodyVulStat) GetLaterCount() *string {
	return s.LaterCount
}

func (s *DescribeClusterVulStatisticsResponseBodyVulStat) GetNntfCount() *string {
	return s.NntfCount
}

func (s *DescribeClusterVulStatisticsResponseBodyVulStat) SetAsapCount(v string) *DescribeClusterVulStatisticsResponseBodyVulStat {
	s.AsapCount = &v
	return s
}

func (s *DescribeClusterVulStatisticsResponseBodyVulStat) SetLaterCount(v string) *DescribeClusterVulStatisticsResponseBodyVulStat {
	s.LaterCount = &v
	return s
}

func (s *DescribeClusterVulStatisticsResponseBodyVulStat) SetNntfCount(v string) *DescribeClusterVulStatisticsResponseBodyVulStat {
	s.NntfCount = &v
	return s
}

func (s *DescribeClusterVulStatisticsResponseBodyVulStat) Validate() error {
	return dara.Validate(s)
}
