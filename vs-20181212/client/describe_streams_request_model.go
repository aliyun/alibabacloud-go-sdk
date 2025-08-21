// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *DescribeStreamsRequest
	GetApp() *string
	SetDeviceId(v string) *DescribeStreamsRequest
	GetDeviceId() *string
	SetDomain(v string) *DescribeStreamsRequest
	GetDomain() *string
	SetGroupId(v string) *DescribeStreamsRequest
	GetGroupId() *string
	SetId(v string) *DescribeStreamsRequest
	GetId() *string
	SetName(v string) *DescribeStreamsRequest
	GetName() *string
	SetOwnerId(v int64) *DescribeStreamsRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeStreamsRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeStreamsRequest
	GetPageSize() *int64
	SetParentId(v string) *DescribeStreamsRequest
	GetParentId() *string
	SetSortBy(v string) *DescribeStreamsRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeStreamsRequest
	GetSortDirection() *string
}

type DescribeStreamsRequest struct {
	// example:
	//
	// live
	App *string `json:"App,omitempty" xml:"App,omitempty"`
	// example:
	//
	// 348*****380-cn-qingdao
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// 323*****997-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 3100000*****00000002
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
}

func (s DescribeStreamsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamsRequest) GetApp() *string {
	return s.App
}

func (s *DescribeStreamsRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeStreamsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeStreamsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeStreamsRequest) GetId() *string {
	return s.Id
}

func (s *DescribeStreamsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeStreamsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStreamsRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeStreamsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeStreamsRequest) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeStreamsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeStreamsRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeStreamsRequest) SetApp(v string) *DescribeStreamsRequest {
	s.App = &v
	return s
}

func (s *DescribeStreamsRequest) SetDeviceId(v string) *DescribeStreamsRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeStreamsRequest) SetDomain(v string) *DescribeStreamsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeStreamsRequest) SetGroupId(v string) *DescribeStreamsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeStreamsRequest) SetId(v string) *DescribeStreamsRequest {
	s.Id = &v
	return s
}

func (s *DescribeStreamsRequest) SetName(v string) *DescribeStreamsRequest {
	s.Name = &v
	return s
}

func (s *DescribeStreamsRequest) SetOwnerId(v int64) *DescribeStreamsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStreamsRequest) SetPageNum(v int64) *DescribeStreamsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeStreamsRequest) SetPageSize(v int64) *DescribeStreamsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStreamsRequest) SetParentId(v string) *DescribeStreamsRequest {
	s.ParentId = &v
	return s
}

func (s *DescribeStreamsRequest) SetSortBy(v string) *DescribeStreamsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeStreamsRequest) SetSortDirection(v string) *DescribeStreamsRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeStreamsRequest) Validate() error {
	return dara.Validate(s)
}
