// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentPlatformDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevices(v []*DescribeParentPlatformDevicesResponseBodyDevices) *DescribeParentPlatformDevicesResponseBody
	GetDevices() []*DescribeParentPlatformDevicesResponseBodyDevices
	SetPageCount(v int64) *DescribeParentPlatformDevicesResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribeParentPlatformDevicesResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeParentPlatformDevicesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeParentPlatformDevicesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeParentPlatformDevicesResponseBody
	GetTotalCount() *int64
}

type DescribeParentPlatformDevicesResponseBody struct {
	Devices []*DescribeParentPlatformDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParentPlatformDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesResponseBody) GetDevices() []*DescribeParentPlatformDevicesResponseBodyDevices {
	return s.Devices
}

func (s *DescribeParentPlatformDevicesResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeParentPlatformDevicesResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeParentPlatformDevicesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeParentPlatformDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParentPlatformDevicesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeParentPlatformDevicesResponseBody) SetDevices(v []*DescribeParentPlatformDevicesResponseBodyDevices) *DescribeParentPlatformDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetPageCount(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetPageNum(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetPageSize(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetRequestId(v string) *DescribeParentPlatformDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) SetTotalCount(v int64) *DescribeParentPlatformDevicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParentPlatformDevicesResponseBodyDevices struct {
	// example:
	//
	// 310101*****7542007
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// example:
	//
	// 3484*****8732174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 3487*****323380-cn-qingdao
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 3614*****766212-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s DescribeParentPlatformDevicesResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentPlatformDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) GetGbId() *string {
	return s.GbId
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) GetId() *string {
	return s.Id
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) GetName() *string {
	return s.Name
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetGbId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.GbId = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetGroupId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.GroupId = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.Id = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetName(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) SetParentId(v string) *DescribeParentPlatformDevicesResponseBodyDevices {
	s.ParentId = &v
	return s
}

func (s *DescribeParentPlatformDevicesResponseBodyDevices) Validate() error {
	return dara.Validate(s)
}
