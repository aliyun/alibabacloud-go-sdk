// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDedicatedBlockStorageClusterInventoryDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody
	GetData() []*QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData
	SetDbscId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody
	GetDbscId() *string
	SetDbscName(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody
	GetDbscName() *string
	SetDiskCategory(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody
	GetDiskCategory() *string
	SetRequestId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody
	GetTotalCount() *int32
}

type QueryDedicatedBlockStorageClusterInventoryDataResponseBody struct {
	// The returned data.
	Data []*QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the dedicated block storage cluster.
	//
	// example:
	//
	// dbsc-xxx
	DbscId *string `json:"DbscId,omitempty" xml:"DbscId,omitempty"`
	// The name of the dedicated block storage cluster.
	//
	// example:
	//
	// myDBSCCluster
	DbscName *string `json:"DbscName,omitempty" xml:"DbscName,omitempty"`
	// The type of the disk. Valid values:
	//
	// 	- cloud_essd: enhanced SSD (ESSD).
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F1A4258A-0C8C-5329-B495-BC5AD7AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 60
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GetData() []*QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData {
	return s.Data
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GetDbscId() *string {
	return s.DbscId
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GetDbscName() *string {
	return s.DbscName
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetData(v []*QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.Data = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetDbscId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.DbscId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetDbscName(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.DbscName = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetDiskCategory(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetRequestId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) SetTotalCount(v int32) *QueryDedicatedBlockStorageClusterInventoryDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData struct {
	// The returned metrics.
	MonitorItems *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems `json:"MonitorItems,omitempty" xml:"MonitorItems,omitempty" type:"Struct"`
	// The ID list of the resource.
	//
	// example:
	//
	// dbsc-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The timestamp when the data is collected.
	//
	// example:
	//
	// 1606403800
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) GetMonitorItems() *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems {
	return s.MonitorItems
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) GetTimestamp() *string {
	return s.Timestamp
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) SetMonitorItems(v *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData {
	s.MonitorItems = v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) SetResourceId(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) SetTimestamp(v string) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyData) Validate() error {
	if s.MonitorItems != nil {
		if err := s.MonitorItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems struct {
	// Available capacity size of the dedicated block storage cluster.
	//
	// example:
	//
	// 61360
	AvailableSize *int64 `json:"AvailableSize,omitempty" xml:"AvailableSize,omitempty"`
	// Total capacity size of the dedicated block storage cluster.
	//
	// example:
	//
	// 61440
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) String() string {
	return dara.Prettify(s)
}

func (s QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) GoString() string {
	return s.String()
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) GetAvailableSize() *int64 {
	return s.AvailableSize
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) SetAvailableSize(v int64) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems {
	s.AvailableSize = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) SetTotalSize(v int64) *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems {
	s.TotalSize = &v
	return s
}

func (s *QueryDedicatedBlockStorageClusterInventoryDataResponseBodyDataMonitorItems) Validate() error {
	return dara.Validate(s)
}
