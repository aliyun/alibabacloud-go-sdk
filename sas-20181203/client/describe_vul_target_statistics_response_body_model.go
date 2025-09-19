// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulTargetStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeVulTargetStatisticsResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeVulTargetStatisticsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVulTargetStatisticsResponseBody
	GetRequestId() *string
	SetTargetStats(v []*DescribeVulTargetStatisticsResponseBodyTargetStats) *DescribeVulTargetStatisticsResponseBody
	GetTargetStats() []*DescribeVulTargetStatisticsResponseBodyTargetStats
	SetTotalCount(v int32) *DescribeVulTargetStatisticsResponseBody
	GetTotalCount() *int32
}

type DescribeVulTargetStatisticsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 23AD0BD2-8771-5647-819E-6BA51E212F80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the configurations of the vulnerability scan feature.
	TargetStats []*DescribeVulTargetStatisticsResponseBodyTargetStats `json:"TargetStats,omitempty" xml:"TargetStats,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVulTargetStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetStatisticsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVulTargetStatisticsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVulTargetStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulTargetStatisticsResponseBody) GetTargetStats() []*DescribeVulTargetStatisticsResponseBodyTargetStats {
	return s.TargetStats
}

func (s *DescribeVulTargetStatisticsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulTargetStatisticsResponseBody) SetCurrentPage(v int32) *DescribeVulTargetStatisticsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBody) SetPageSize(v int32) *DescribeVulTargetStatisticsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBody) SetRequestId(v string) *DescribeVulTargetStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBody) SetTargetStats(v []*DescribeVulTargetStatisticsResponseBodyTargetStats) *DescribeVulTargetStatisticsResponseBody {
	s.TargetStats = v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBody) SetTotalCount(v int32) *DescribeVulTargetStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVulTargetStatisticsResponseBodyTargetStats struct {
	// An array that consists of available servers.
	Targets []*DescribeVulTargetStatisticsResponseBodyTargetStatsTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The total number of servers.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of servers to which the configurations are applied.
	//
	// example:
	//
	// 1
	UuidCount *int32 `json:"UuidCount,omitempty" xml:"UuidCount,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- cve: Linux software vulnerabilities
	//
	// 	- sys: Windows system vulnerabilities
	//
	// 	- cms: Web-CMS vulnerabilities
	//
	// 	- emg: urgent vulnerabilities
	//
	// example:
	//
	// cve
	VulType *string `json:"VulType,omitempty" xml:"VulType,omitempty"`
}

func (s DescribeVulTargetStatisticsResponseBodyTargetStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetStatisticsResponseBodyTargetStats) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) GetTargets() []*DescribeVulTargetStatisticsResponseBodyTargetStatsTargets {
	return s.Targets
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) GetUuidCount() *int32 {
	return s.UuidCount
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) GetVulType() *string {
	return s.VulType
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) SetTargets(v []*DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) *DescribeVulTargetStatisticsResponseBodyTargetStats {
	s.Targets = v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) SetTotalCount(v int32) *DescribeVulTargetStatisticsResponseBodyTargetStats {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) SetUuidCount(v int32) *DescribeVulTargetStatisticsResponseBodyTargetStats {
	s.UuidCount = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) SetVulType(v string) *DescribeVulTargetStatisticsResponseBodyTargetStats {
	s.VulType = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStats) Validate() error {
	return dara.Validate(s)
}

type DescribeVulTargetStatisticsResponseBodyTargetStatsTargets struct {
	// Indicates whether the configurations are applied to the server. Valid values:
	//
	// 	- **add**: yes
	//
	// 	- **del**: no
	//
	// example:
	//
	// add
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The group ID or UUID of the server to which the configurations are applied.
	//
	// example:
	//
	// 0011ea53-738c-4bff-93be-ce6a1cc9****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The condition by which the configurations are applied to the server. Valid values:
	//
	// 	- **uuid**: the UUID of the server
	//
	// 	- **groupId**: the ID of the server group
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) GetFlag() *string {
	return s.Flag
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) GetTarget() *string {
	return s.Target
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) SetFlag(v string) *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets {
	s.Flag = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) SetTarget(v string) *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets {
	s.Target = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) SetTargetType(v string) *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets {
	s.TargetType = &v
	return s
}

func (s *DescribeVulTargetStatisticsResponseBodyTargetStatsTargets) Validate() error {
	return dara.Validate(s)
}
