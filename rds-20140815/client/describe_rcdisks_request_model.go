// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIds(v string) *DescribeRCDisksRequest
	GetDiskIds() *string
	SetInstanceId(v string) *DescribeRCDisksRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeRCDisksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRCDisksRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeRCDisksRequest
	GetRegionId() *string
	SetTag(v []*DescribeRCDisksRequestTag) *DescribeRCDisksRequest
	GetTag() []*DescribeRCDisksRequestTag
}

type DescribeRCDisksRequest struct {
	// The disk ID. The value is a JSON array that consists of up to 100 disk IDs. Separate the disk IDs with commas (,). Format: `["Disk ID1","Disk ID2"]`.
	//
	// example:
	//
	// ["rcd-bp67acfmxazb4p****", "rcd-bp67acfmxazb4g****", … "rcd-bp67acfmxazb4d****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
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
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of the tags.
	Tag []*DescribeRCDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeRCDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCDisksRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeRCDisksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCDisksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRCDisksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRCDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCDisksRequest) GetTag() []*DescribeRCDisksRequestTag {
	return s.Tag
}

func (s *DescribeRCDisksRequest) SetDiskIds(v string) *DescribeRCDisksRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeRCDisksRequest) SetInstanceId(v string) *DescribeRCDisksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCDisksRequest) SetPageNumber(v int64) *DescribeRCDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCDisksRequest) SetPageSize(v int64) *DescribeRCDisksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCDisksRequest) SetRegionId(v string) *DescribeRCDisksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCDisksRequest) SetTag(v []*DescribeRCDisksRequestTag) *DescribeRCDisksRequest {
	s.Tag = v
	return s
}

func (s *DescribeRCDisksRequest) Validate() error {
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

type DescribeRCDisksRequestTag struct {
	// The key of the tag. The tag key **cannot be*	- an empty string or a duplicate value.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The tag value **can be*	- an empty string.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRCDisksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDisksRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeRCDisksRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeRCDisksRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeRCDisksRequestTag) SetKey(v string) *DescribeRCDisksRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeRCDisksRequestTag) SetValue(v string) *DescribeRCDisksRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeRCDisksRequestTag) Validate() error {
	return dara.Validate(s)
}
