// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLensMonitorDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskInfos(v []*DescribeLensMonitorDisksResponseBodyDiskInfos) *DescribeLensMonitorDisksResponseBody
	GetDiskInfos() []*DescribeLensMonitorDisksResponseBodyDiskInfos
	SetNextToken(v string) *DescribeLensMonitorDisksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeLensMonitorDisksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeLensMonitorDisksResponseBody
	GetTotalCount() *int64
}

type DescribeLensMonitorDisksResponseBody struct {
	// The list of cloud disk information.
	DiskInfos []*DescribeLensMonitorDisksResponseBodyDiskInfos `json:"DiskInfos,omitempty" xml:"DiskInfos,omitempty" type:"Repeated"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID. A request ID is returned regardless of whether the API call succeeds.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLensMonitorDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLensMonitorDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLensMonitorDisksResponseBody) GetDiskInfos() []*DescribeLensMonitorDisksResponseBodyDiskInfos {
	return s.DiskInfos
}

func (s *DescribeLensMonitorDisksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLensMonitorDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLensMonitorDisksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLensMonitorDisksResponseBody) SetDiskInfos(v []*DescribeLensMonitorDisksResponseBodyDiskInfos) *DescribeLensMonitorDisksResponseBody {
	s.DiskInfos = v
	return s
}

func (s *DescribeLensMonitorDisksResponseBody) SetNextToken(v string) *DescribeLensMonitorDisksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBody) SetRequestId(v string) *DescribeLensMonitorDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBody) SetTotalCount(v int64) *DescribeLensMonitorDisksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBody) Validate() error {
	if s.DiskInfos != nil {
		for _, item := range s.DiskInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLensMonitorDisksResponseBodyDiskInfos struct {
	// The maximum data throughput for read/write (I/O) operations per second. Unit: MB/s.
	//
	// example:
	//
	// 300
	Bps *int32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// Indicates whether burst (performance bursting) is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL cloud disk](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
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
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The cloud disk ID.
	//
	// example:
	//
	// d-cd401****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The cloud disk name.
	//
	// example:
	//
	// disk-28c6b****
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The cloud disk status. Valid values:
	//
	// - Available: in use.
	//
	// - Deleted: deleted.
	//
	// example:
	//
	// Available
	DiskStatus *string `json:"DiskStatus,omitempty" xml:"DiskStatus,omitempty"`
	// The cloud disk type. Valid values:
	//
	// - system: system cloud disk.
	//
	// - data: data cloud disk.
	//
	// example:
	//
	// system
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The maximum number of read/write (I/O) operations per second. Unit: operations/s.
	//
	// example:
	//
	// 4000
	Iops *int32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The collection of event tags for the cloud disk. Event tags display events that occurred on the cloud disk within the last 24 hours, with a delay of up to 1 hour compared to the actual events.
	LensTags []*string `json:"LensTags,omitempty" xml:"LensTags,omitempty" type:"Repeated"`
	// The performance level (PL) of the ESSD cloud disk. Valid values:
	//
	// - PL0: maximum random read/write IOPS of 10,000 per standard SSD.
	//
	// - PL1: maximum random read/write IOPS of 50,000 per standard SSD.
	//
	// - PL2: maximum random read/write IOPS of 100,000 per standard SSD.
	//
	// - PL3: maximum random read/write IOPS of 1,000,000 per standard SSD.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL cloud disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL cloud disk](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 4000
	ProvisionedIops *int32 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether the cloud disk is a shared cloud disk.
	//
	// example:
	//
	// true
	SharingEnabled *string `json:"SharingEnabled,omitempty" xml:"SharingEnabled,omitempty"`
	// The cloud disk size. Unit: GiB.
	//
	// example:
	//
	// 64
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The collection of tags for the cloud disk.
	Tags []*DescribeLensMonitorDisksResponseBodyDiskInfosTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID of the cloud disk.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeLensMonitorDisksResponseBodyDiskInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLensMonitorDisksResponseBodyDiskInfos) GoString() string {
	return s.String()
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetBps() *int32 {
	return s.Bps
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetDiskStatus() *string {
	return s.DiskStatus
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetIops() *int32 {
	return s.Iops
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetLensTags() []*string {
	return s.LensTags
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetProvisionedIops() *int32 {
	return s.ProvisionedIops
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetSharingEnabled() *string {
	return s.SharingEnabled
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetSize() *int32 {
	return s.Size
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetTags() []*DescribeLensMonitorDisksResponseBodyDiskInfosTags {
	return s.Tags
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetBps(v int32) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.Bps = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetBurstingEnabled(v bool) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetDiskCategory(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.DiskCategory = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetDiskId(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.DiskId = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetDiskName(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.DiskName = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetDiskStatus(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.DiskStatus = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetDiskType(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.DiskType = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetIops(v int32) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.Iops = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetLensTags(v []*string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.LensTags = v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetPerformanceLevel(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetProvisionedIops(v int32) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetRegionId(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetSharingEnabled(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.SharingEnabled = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetSize(v int32) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.Size = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetTags(v []*DescribeLensMonitorDisksResponseBodyDiskInfosTags) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.Tags = v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) SetZoneId(v string) *DescribeLensMonitorDisksResponseBodyDiskInfos {
	s.ZoneId = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfos) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLensMonitorDisksResponseBodyDiskInfosTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// user
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeLensMonitorDisksResponseBodyDiskInfosTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLensMonitorDisksResponseBodyDiskInfosTags) GoString() string {
	return s.String()
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfosTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfosTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfosTags) SetTagKey(v string) *DescribeLensMonitorDisksResponseBodyDiskInfosTags {
	s.TagKey = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfosTags) SetTagValue(v string) *DescribeLensMonitorDisksResponseBodyDiskInfosTags {
	s.TagValue = &v
	return s
}

func (s *DescribeLensMonitorDisksResponseBodyDiskInfosTags) Validate() error {
	return dara.Validate(s)
}
