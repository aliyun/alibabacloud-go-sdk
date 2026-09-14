// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLensMonitorDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskCategory(v string) *DescribeLensMonitorDisksRequest
	GetDiskCategory() *string
	SetDiskIdPattern(v string) *DescribeLensMonitorDisksRequest
	GetDiskIdPattern() *string
	SetDiskIds(v []*string) *DescribeLensMonitorDisksRequest
	GetDiskIds() []*string
	SetEcsInstanceId(v string) *DescribeLensMonitorDisksRequest
	GetEcsInstanceId() *string
	SetLensTags(v []*string) *DescribeLensMonitorDisksRequest
	GetLensTags() []*string
	SetMaxResults(v int32) *DescribeLensMonitorDisksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeLensMonitorDisksRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeLensMonitorDisksRequest
	GetRegionId() *string
}

type DescribeLensMonitorDisksRequest struct {
	// The cloud disk type. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: Enterprise SSD (ESSD).
	//
	// - cloud_auto: ESSD AutoPL cloud disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// example:
	//
	// cloud_auto
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The regular expression pattern used for fuzzy match filtering of cloud disk IDs.
	//
	// example:
	//
	// d-cd40hxfu0v*
	DiskIdPattern *string `json:"DiskIdPattern,omitempty" xml:"DiskIdPattern,omitempty"`
	// The list of cloud disk IDs.
	//
	// example:
	//
	// [\\"d-1\\", \\"d-2\\"]
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// The ECS instance ID.
	//
	// example:
	//
	// i-2zedroc0yv8z19ubnyos
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The list of cloud disk event tags, used to filter cloud disks that have experienced these event types within the last 24 hours. Valid values:
	//
	// - NoSnapshot: data protection
	//
	// - BurstIOTriggered: burst I/O
	//
	// - CostOptimizationNeeded: cost optimization
	//
	// - DiskSpecNotMatchedWithInstance: instance and cloud disk specifications do not match
	//
	// - DiskIONo4kAligned: non-4K aligned read/write
	//
	// - DiskIOHang: I/O hang occurred on the cloud disk
	//
	// - InstanceIOPSExceedInstanceMaxLimit: instance IOPS reached the upper limit
	//
	// - InstanceBPSExceedInstanceMaxLimit: instance BPS reached the upper limit
	//
	// - DiskIOPSExceedInstanceMaxLimit: cloud disk IOPS reached the instance upper limit
	//
	// - DiskBPSExceedInstanceMaxLimit: cloud disk BPS reached the instance upper limit
	//
	// - DiskIOPSExceedDiskMaxLimit: cloud disk IOPS reached the disk upper limit
	//
	// - DiskBPSExceedDiskMaxLimit: cloud disk BPS reached the disk upper limit
	LensTags []*string `json:"LensTags,omitempty" xml:"LensTags,omitempty" type:"Repeated"`
	// The maximum number of entries per page for a paged query. Maximum value: 100.
	//
	// Default value:
	//
	// - The default value is 10.
	//
	// - If the specified value is greater than 100, the default value of 100 is used.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call DescribeRegions to query the list of regions supported by EBS Lens.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLensMonitorDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLensMonitorDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeLensMonitorDisksRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *DescribeLensMonitorDisksRequest) GetDiskIdPattern() *string {
	return s.DiskIdPattern
}

func (s *DescribeLensMonitorDisksRequest) GetDiskIds() []*string {
	return s.DiskIds
}

func (s *DescribeLensMonitorDisksRequest) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeLensMonitorDisksRequest) GetLensTags() []*string {
	return s.LensTags
}

func (s *DescribeLensMonitorDisksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeLensMonitorDisksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLensMonitorDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLensMonitorDisksRequest) SetDiskCategory(v string) *DescribeLensMonitorDisksRequest {
	s.DiskCategory = &v
	return s
}

func (s *DescribeLensMonitorDisksRequest) SetDiskIdPattern(v string) *DescribeLensMonitorDisksRequest {
	s.DiskIdPattern = &v
	return s
}

func (s *DescribeLensMonitorDisksRequest) SetDiskIds(v []*string) *DescribeLensMonitorDisksRequest {
	s.DiskIds = v
	return s
}

func (s *DescribeLensMonitorDisksRequest) SetEcsInstanceId(v string) *DescribeLensMonitorDisksRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeLensMonitorDisksRequest) SetLensTags(v []*string) *DescribeLensMonitorDisksRequest {
	s.LensTags = v
	return s
}

func (s *DescribeLensMonitorDisksRequest) SetMaxResults(v int32) *DescribeLensMonitorDisksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeLensMonitorDisksRequest) SetNextToken(v string) *DescribeLensMonitorDisksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLensMonitorDisksRequest) SetRegionId(v string) *DescribeLensMonitorDisksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLensMonitorDisksRequest) Validate() error {
	return dara.Validate(s)
}
