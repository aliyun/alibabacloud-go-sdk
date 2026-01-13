// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeRCSnapshotsRequest
	GetDiskId() *string
	SetInstanceId(v string) *DescribeRCSnapshotsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeRCSnapshotsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRCSnapshotsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeRCSnapshotsRequest
	GetRegionId() *string
	SetSnapshotIds(v string) *DescribeRCSnapshotsRequest
	GetSnapshotIds() *string
	SetTag(v []*DescribeRCSnapshotsRequestTag) *DescribeRCSnapshotsRequest
	GetTag() []*DescribeRCSnapshotsRequestTag
}

type DescribeRCSnapshotsRequest struct {
	// The cloud disk ID.
	//
	// example:
	//
	// rcd-wz9c8isqly8637zw****
	DiskId     *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30*	- to **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot IDs.
	//
	// You can specify a maximum of 100 IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["rcds-bp67acfmxazb4p****", "rcds-bp67acfmxazb5p****", … "rcds-bp67acfmxazb6p****"]
	SnapshotIds *string                          `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	Tag         []*DescribeRCSnapshotsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRCSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCSnapshotsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeRCSnapshotsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCSnapshotsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRCSnapshotsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRCSnapshotsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCSnapshotsRequest) GetSnapshotIds() *string {
	return s.SnapshotIds
}

func (s *DescribeRCSnapshotsRequest) GetTag() []*DescribeRCSnapshotsRequestTag {
	return s.Tag
}

func (s *DescribeRCSnapshotsRequest) SetDiskId(v string) *DescribeRCSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeRCSnapshotsRequest) SetInstanceId(v string) *DescribeRCSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCSnapshotsRequest) SetPageNumber(v int64) *DescribeRCSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCSnapshotsRequest) SetPageSize(v int64) *DescribeRCSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCSnapshotsRequest) SetRegionId(v string) *DescribeRCSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCSnapshotsRequest) SetSnapshotIds(v string) *DescribeRCSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *DescribeRCSnapshotsRequest) SetTag(v []*DescribeRCSnapshotsRequestTag) *DescribeRCSnapshotsRequest {
	s.Tag = v
	return s
}

func (s *DescribeRCSnapshotsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCSnapshotsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCSnapshotsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSnapshotsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeRCSnapshotsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRCSnapshotsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRCSnapshotsRequestTag) SetKey(v string) *DescribeRCSnapshotsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeRCSnapshotsRequestTag) SetValue(v string) *DescribeRCSnapshotsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeRCSnapshotsRequestTag) Validate() error {
	return dara.Validate(s)
}
