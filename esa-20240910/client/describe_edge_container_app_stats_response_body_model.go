// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeContainerAppStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCpuUsageSecondsQuotaRateAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody
	GetCpuUsageSecondsQuotaRateAvg() *float64
	SetCpuUsageSecondsTotalAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody
	GetCpuUsageSecondsTotalAvg() *float64
	SetFsReadsBytesAvgAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody
	GetFsReadsBytesAvgAvg() *float64
	SetFsWritesBytesAvgAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody
	GetFsWritesBytesAvgAvg() *float64
	SetMemoryRssAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody
	GetMemoryRssAvg() *float64
	SetMemoryRssQuotaRateAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody
	GetMemoryRssQuotaRateAvg() *float64
	SetPodReadyRateAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody
	GetPodReadyRateAvg() *float64
	SetPoints(v []*DescribeEdgeContainerAppStatsResponseBodyPoints) *DescribeEdgeContainerAppStatsResponseBody
	GetPoints() []*DescribeEdgeContainerAppStatsResponseBodyPoints
	SetRequestId(v string) *DescribeEdgeContainerAppStatsResponseBody
	GetRequestId() *string
}

type DescribeEdgeContainerAppStatsResponseBody struct {
	// Average CPU limit ratio
	//
	// example:
	//
	// 0.1
	CpuUsageSecondsQuotaRateAvg *float64 `json:"CpuUsageSecondsQuotaRateAvg,omitempty" xml:"CpuUsageSecondsQuotaRateAvg,omitempty"`
	// Average number of CPU cores
	//
	// example:
	//
	// 2
	CpuUsageSecondsTotalAvg *float64 `json:"CpuUsageSecondsTotalAvg,omitempty" xml:"CpuUsageSecondsTotalAvg,omitempty"`
	// Average read IO
	//
	// example:
	//
	// 0
	FsReadsBytesAvgAvg *float64 `json:"FsReadsBytesAvgAvg,omitempty" xml:"FsReadsBytesAvgAvg,omitempty"`
	// Average write IO
	//
	// example:
	//
	// 0
	FsWritesBytesAvgAvg *float64 `json:"FsWritesBytesAvgAvg,omitempty" xml:"FsWritesBytesAvgAvg,omitempty"`
	// Average memory usage
	//
	// example:
	//
	// 0.1
	MemoryRssAvg *float64 `json:"MemoryRssAvg,omitempty" xml:"MemoryRssAvg,omitempty"`
	// Average memory limit proportion
	//
	// example:
	//
	// 1
	MemoryRssQuotaRateAvg *float64 `json:"MemoryRssQuotaRateAvg,omitempty" xml:"MemoryRssQuotaRateAvg,omitempty"`
	// Average PodReady rate
	//
	// example:
	//
	// 100
	PodReadyRateAvg *float64                                           `json:"PodReadyRateAvg,omitempty" xml:"PodReadyRateAvg,omitempty"`
	Points          []*DescribeEdgeContainerAppStatsResponseBodyPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 35C66C7B-671H-4297-9187-2C4477247A78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEdgeContainerAppStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeContainerAppStatsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetCpuUsageSecondsQuotaRateAvg() *float64 {
	return s.CpuUsageSecondsQuotaRateAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetCpuUsageSecondsTotalAvg() *float64 {
	return s.CpuUsageSecondsTotalAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetFsReadsBytesAvgAvg() *float64 {
	return s.FsReadsBytesAvgAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetFsWritesBytesAvgAvg() *float64 {
	return s.FsWritesBytesAvgAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetMemoryRssAvg() *float64 {
	return s.MemoryRssAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetMemoryRssQuotaRateAvg() *float64 {
	return s.MemoryRssQuotaRateAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetPodReadyRateAvg() *float64 {
	return s.PodReadyRateAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetPoints() []*DescribeEdgeContainerAppStatsResponseBodyPoints {
	return s.Points
}

func (s *DescribeEdgeContainerAppStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetCpuUsageSecondsQuotaRateAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody {
	s.CpuUsageSecondsQuotaRateAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetCpuUsageSecondsTotalAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody {
	s.CpuUsageSecondsTotalAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetFsReadsBytesAvgAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody {
	s.FsReadsBytesAvgAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetFsWritesBytesAvgAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody {
	s.FsWritesBytesAvgAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetMemoryRssAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody {
	s.MemoryRssAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetMemoryRssQuotaRateAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody {
	s.MemoryRssQuotaRateAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetPodReadyRateAvg(v float64) *DescribeEdgeContainerAppStatsResponseBody {
	s.PodReadyRateAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetPoints(v []*DescribeEdgeContainerAppStatsResponseBodyPoints) *DescribeEdgeContainerAppStatsResponseBody {
	s.Points = v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) SetRequestId(v string) *DescribeEdgeContainerAppStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEdgeContainerAppStatsResponseBodyPoints struct {
	// example:
	//
	// 0.1
	ContainerCpuUsageSecondsQuotaRate *float64 `json:"ContainerCpuUsageSecondsQuotaRate,omitempty" xml:"ContainerCpuUsageSecondsQuotaRate,omitempty"`
	// example:
	//
	// 2
	ContainerCpuUsageSecondsTotal *float64 `json:"ContainerCpuUsageSecondsTotal,omitempty" xml:"ContainerCpuUsageSecondsTotal,omitempty"`
	// example:
	//
	// 0
	ContainerFsReadsBytesAvg *float64 `json:"ContainerFsReadsBytesAvg,omitempty" xml:"ContainerFsReadsBytesAvg,omitempty"`
	// example:
	//
	// 0
	ContainerFsWritesBytesAvg *float64 `json:"ContainerFsWritesBytesAvg,omitempty" xml:"ContainerFsWritesBytesAvg,omitempty"`
	// example:
	//
	// 0.1
	ContainerMemoryRss *float64 `json:"ContainerMemoryRss,omitempty" xml:"ContainerMemoryRss,omitempty"`
	// example:
	//
	// 1
	ContainerMemoryRssQuotaRate *float64 `json:"ContainerMemoryRssQuotaRate,omitempty" xml:"ContainerMemoryRssQuotaRate,omitempty"`
	// example:
	//
	// 100
	PodReadyRate *float64 `json:"PodReadyRate,omitempty" xml:"PodReadyRate,omitempty"`
	// example:
	//
	// 2024-01-18T15:04:05Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeEdgeContainerAppStatsResponseBodyPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeContainerAppStatsResponseBodyPoints) GoString() string {
	return s.String()
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetContainerCpuUsageSecondsQuotaRate() *float64 {
	return s.ContainerCpuUsageSecondsQuotaRate
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetContainerCpuUsageSecondsTotal() *float64 {
	return s.ContainerCpuUsageSecondsTotal
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetContainerFsReadsBytesAvg() *float64 {
	return s.ContainerFsReadsBytesAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetContainerFsWritesBytesAvg() *float64 {
	return s.ContainerFsWritesBytesAvg
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetContainerMemoryRss() *float64 {
	return s.ContainerMemoryRss
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetContainerMemoryRssQuotaRate() *float64 {
	return s.ContainerMemoryRssQuotaRate
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetPodReadyRate() *float64 {
	return s.PodReadyRate
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) GetTime() *string {
	return s.Time
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetContainerCpuUsageSecondsQuotaRate(v float64) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.ContainerCpuUsageSecondsQuotaRate = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetContainerCpuUsageSecondsTotal(v float64) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.ContainerCpuUsageSecondsTotal = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetContainerFsReadsBytesAvg(v float64) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.ContainerFsReadsBytesAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetContainerFsWritesBytesAvg(v float64) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.ContainerFsWritesBytesAvg = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetContainerMemoryRss(v float64) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.ContainerMemoryRss = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetContainerMemoryRssQuotaRate(v float64) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.ContainerMemoryRssQuotaRate = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetPodReadyRate(v float64) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.PodReadyRate = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) SetTime(v string) *DescribeEdgeContainerAppStatsResponseBodyPoints {
	s.Time = &v
	return s
}

func (s *DescribeEdgeContainerAppStatsResponseBodyPoints) Validate() error {
	return dara.Validate(s)
}
