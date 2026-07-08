// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *DescribeDevicesRequest
	GetDirectoryId() *string
	SetDsn(v string) *DescribeDevicesRequest
	GetDsn() *string
	SetGbId(v string) *DescribeDevicesRequest
	GetGbId() *string
	SetGroupId(v string) *DescribeDevicesRequest
	GetGroupId() *string
	SetId(v string) *DescribeDevicesRequest
	GetId() *string
	SetIncludeDirectory(v bool) *DescribeDevicesRequest
	GetIncludeDirectory() *bool
	SetIncludeStats(v bool) *DescribeDevicesRequest
	GetIncludeStats() *bool
	SetName(v string) *DescribeDevicesRequest
	GetName() *string
	SetOwnerId(v int64) *DescribeDevicesRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeDevicesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeDevicesRequest
	GetPageSize() *int64
	SetParentId(v string) *DescribeDevicesRequest
	GetParentId() *string
	SetSortBy(v string) *DescribeDevicesRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeDevicesRequest
	GetSortDirection() *string
	SetStatus(v string) *DescribeDevicesRequest
	GetStatus() *string
	SetType(v string) *DescribeDevicesRequest
	GetType() *string
	SetVendor(v string) *DescribeDevicesRequest
	GetVendor() *string
}

type DescribeDevicesRequest struct {
	// The ID of the directory to which the device belongs.
	//
	// example:
	//
	// 399*****488-cn-qingdao
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The serial number of the device. The value must be unique.
	//
	// example:
	//
	// 7D0*****4C0
	Dsn *string `json:"Dsn,omitempty" xml:"Dsn,omitempty"`
	// You can query by device national standard ID.
	//
	// example:
	//
	// 310000000****0000002
	GbId *string `json:"GbId,omitempty" xml:"GbId,omitempty"`
	// Query by device Space ID.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The device ID.
	//
	// > Specify multiple IDs. Separate them with commas (,).
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to return directory information. Default value: false.
	//
	// example:
	//
	// false
	IncludeDirectory *bool `json:"IncludeDirectory,omitempty" xml:"IncludeDirectory,omitempty"`
	// Specifies whether to return stream statistics. Default value: false.
	//
	// example:
	//
	// false
	IncludeStats *bool `json:"IncludeStats,omitempty" xml:"IncludeStats,omitempty"`
	// The device name.
	//
	// > Specify multiple names. Separate them with commas (,).
	//
	// example:
	//
	// 摄像头A
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent device.
	//
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The field by which to sort the results. Valid value:
	//
	// > id (default)
	//
	// example:
	//
	// id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values:
	//
	// - asc (ascending) (default)
	//
	// - desc (descending)
	//
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// Query devices by status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The device type. Valid values:
	//
	// - ipc (camera)
	//
	// - platform
	//
	// - ied (intelligent edge device)
	//
	// example:
	//
	// ipc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Query by device manufacturer.
	//
	// example:
	//
	// 8yd*****qem
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDevicesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDevicesRequest) GetDsn() *string {
	return s.Dsn
}

func (s *DescribeDevicesRequest) GetGbId() *string {
	return s.GbId
}

func (s *DescribeDevicesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDevicesRequest) GetId() *string {
	return s.Id
}

func (s *DescribeDevicesRequest) GetIncludeDirectory() *bool {
	return s.IncludeDirectory
}

func (s *DescribeDevicesRequest) GetIncludeStats() *bool {
	return s.IncludeStats
}

func (s *DescribeDevicesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDevicesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeDevicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDevicesRequest) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDevicesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeDevicesRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeDevicesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDevicesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDevicesRequest) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeDevicesRequest) SetDirectoryId(v string) *DescribeDevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDevicesRequest) SetDsn(v string) *DescribeDevicesRequest {
	s.Dsn = &v
	return s
}

func (s *DescribeDevicesRequest) SetGbId(v string) *DescribeDevicesRequest {
	s.GbId = &v
	return s
}

func (s *DescribeDevicesRequest) SetGroupId(v string) *DescribeDevicesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDevicesRequest) SetId(v string) *DescribeDevicesRequest {
	s.Id = &v
	return s
}

func (s *DescribeDevicesRequest) SetIncludeDirectory(v bool) *DescribeDevicesRequest {
	s.IncludeDirectory = &v
	return s
}

func (s *DescribeDevicesRequest) SetIncludeStats(v bool) *DescribeDevicesRequest {
	s.IncludeStats = &v
	return s
}

func (s *DescribeDevicesRequest) SetName(v string) *DescribeDevicesRequest {
	s.Name = &v
	return s
}

func (s *DescribeDevicesRequest) SetOwnerId(v int64) *DescribeDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageNum(v int64) *DescribeDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageSize(v int64) *DescribeDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesRequest) SetParentId(v string) *DescribeDevicesRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeDevicesRequest) SetSortBy(v string) *DescribeDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeDevicesRequest) SetSortDirection(v string) *DescribeDevicesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeDevicesRequest) SetStatus(v string) *DescribeDevicesRequest {
	s.Status = &v
	return s
}

func (s *DescribeDevicesRequest) SetType(v string) *DescribeDevicesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDevicesRequest) SetVendor(v string) *DescribeDevicesRequest {
	s.Vendor = &v
	return s
}

func (s *DescribeDevicesRequest) Validate() error {
	return dara.Validate(s)
}
