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
	// The type of the disk. Valid values:
	//
	// - cloud
	//
	// - cloud_efficiency
	//
	// - cloud_ssd
	//
	// - cloud_essd
	//
	// - cloud_auto
	//
	// - cloud_essd_entry
	//
	// example:
	//
	// cloud_auto
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// Regular matching fuzzy query to filter cloud disk IDs.
	//
	// example:
	//
	// d-cd40hxfu0v**
	DiskIdPattern *string `json:"DiskIdPattern,omitempty" xml:"DiskIdPattern,omitempty"`
	// The list of disks.
	//
	// example:
	//
	// [\\"d-1\\", \\"d-2\\"]
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// Event tags of the disk, which are used to filter the disks on which the events associated with the specified tags occurred in the previous 24 hours. Valid values:
	//
	// 	- NoSnapshot: specifies the event that is triggered because no snapshot is created for the disk to protect data on the disk.
	//
	// 	- BurstIOTriggered: specifies the event that is triggered when a burst I/O operation is performed on the disk.
	//
	// 	- CostOptimizationNeeded: specifies the event that is triggered when cost optimization is required.
	//
	// 	- DiskSpecNotMatchedWithInstance: specifies the event that is triggered if the disk specifications do not match the instance to which the disk is attached.
	//
	// 	- DiskIONo4kAligned: specifies the event that is triggered if the physical and logical sectors involved in a read or write operation are not 4K aligned.
	//
	// 	- DiskIOHang: specifies the event that is triggered when an I/O hang occurs on the disk.
	//
	// 	- InstanceIOPSExceedInstanceMaxLimit: specifies the event that is triggered when the number of IOPS on the instance reaches the upper limit.
	//
	// 	- InstanceBPSExceedInstanceMaxLimit: specifies the event that is triggered when the number of BPS on the instance reaches the upper limit.
	//
	// 	- DiskIOPSExceedInstanceMaxLimit: specifies the event that is triggered when the number of IOPS on the disk reaches the upper limit of the instance.
	//
	// 	- DiskBPSExceedInstanceMaxLimit: specifies the event that is triggered when the number of BPS on the disk reaches the upper limit of the instance.
	//
	// 	- DiskIOPSExceedDiskMaxLimit: specifies the event that is triggered when the number of IOPS on the disk reaches the upper limit of the disk.
	//
	// 	- DiskBPSExceedDiskMaxLimit: specifies the event that is triggered when the number of BPS on the disk reaches the upper limit of the disk.
	LensTags []*string `json:"LensTags,omitempty" xml:"LensTags,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query to retrieve more results.
	//
	// >The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
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
