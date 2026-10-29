// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribePurchasedDevicesRequest
	GetGroupId() *string
	SetId(v string) *DescribePurchasedDevicesRequest
	GetId() *string
	SetName(v string) *DescribePurchasedDevicesRequest
	GetName() *string
	SetOwnerId(v int64) *DescribePurchasedDevicesRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribePurchasedDevicesRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribePurchasedDevicesRequest
	GetPageSize() *int64
	SetSortBy(v string) *DescribePurchasedDevicesRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribePurchasedDevicesRequest
	GetSortDirection() *string
	SetSubType(v string) *DescribePurchasedDevicesRequest
	GetSubType() *string
	SetType(v string) *DescribePurchasedDevicesRequest
	GetType() *string
	SetVendor(v string) *DescribePurchasedDevicesRequest
	GetVendor() *string
}

type DescribePurchasedDevicesRequest struct {
	// The ID of the workspace to which the device belongs.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The device ID.
	//
	// example:
	//
	// 4070*****1132-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The device name.
	//
	// example:
	//
	// XXX intersection camera
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
	// The field by which to sort the results. Valid values:
	//
	// - id (default)
	//
	// example:
	//
	// id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Default value: asc. Valid values:
	//
	// - asc: Ascending order.
	//
	// - desc: Descending order.
	//
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// The device subtype. Valid values:
	//
	// - bullet: Bullet camera.
	//
	// - dome: Dome camera.
	//
	// - ptz: PTZ camera.
	//
	// example:
	//
	// dome
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The device type. Valid values:
	//
	// - ipc: Camera.
	//
	// - platform: Platform.
	//
	// - ied: Intelligent device.
	//
	// example:
	//
	// ipc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The device vendor.
	//
	// example:
	//
	// Company A
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribePurchasedDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePurchasedDevicesRequest) GetId() *string {
	return s.Id
}

func (s *DescribePurchasedDevicesRequest) GetName() *string {
	return s.Name
}

func (s *DescribePurchasedDevicesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePurchasedDevicesRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribePurchasedDevicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePurchasedDevicesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribePurchasedDevicesRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribePurchasedDevicesRequest) GetSubType() *string {
	return s.SubType
}

func (s *DescribePurchasedDevicesRequest) GetType() *string {
	return s.Type
}

func (s *DescribePurchasedDevicesRequest) GetVendor() *string {
	return s.Vendor
}

func (s *DescribePurchasedDevicesRequest) SetGroupId(v string) *DescribePurchasedDevicesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetId(v string) *DescribePurchasedDevicesRequest {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetName(v string) *DescribePurchasedDevicesRequest {
	s.Name = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetOwnerId(v int64) *DescribePurchasedDevicesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetPageNum(v int64) *DescribePurchasedDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetPageSize(v int64) *DescribePurchasedDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetSortBy(v string) *DescribePurchasedDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetSortDirection(v string) *DescribePurchasedDevicesRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetSubType(v string) *DescribePurchasedDevicesRequest {
	s.SubType = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetType(v string) *DescribePurchasedDevicesRequest {
	s.Type = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) SetVendor(v string) *DescribePurchasedDevicesRequest {
	s.Vendor = &v
	return s
}

func (s *DescribePurchasedDevicesRequest) Validate() error {
	return dara.Validate(s)
}
