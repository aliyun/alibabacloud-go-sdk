// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevices(v []*DescribePurchasedDevicesResponseBodyDevices) *DescribePurchasedDevicesResponseBody
	GetDevices() []*DescribePurchasedDevicesResponseBodyDevices
	SetPageCount(v int64) *DescribePurchasedDevicesResponseBody
	GetPageCount() *int64
	SetPageNum(v int64) *DescribePurchasedDevicesResponseBody
	GetPageNum() *int64
	SetPageSize(v int64) *DescribePurchasedDevicesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribePurchasedDevicesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePurchasedDevicesResponseBody
	GetTotalCount() *int64
}

type DescribePurchasedDevicesResponseBody struct {
	Devices []*DescribePurchasedDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
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

func (s DescribePurchasedDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesResponseBody) GetDevices() []*DescribePurchasedDevicesResponseBodyDevices {
	return s.Devices
}

func (s *DescribePurchasedDevicesResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribePurchasedDevicesResponseBody) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribePurchasedDevicesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePurchasedDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePurchasedDevicesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePurchasedDevicesResponseBody) SetDevices(v []*DescribePurchasedDevicesResponseBodyDevices) *DescribePurchasedDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetPageCount(v int64) *DescribePurchasedDevicesResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetPageNum(v int64) *DescribePurchasedDevicesResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetPageSize(v int64) *DescribePurchasedDevicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetRequestId(v string) *DescribePurchasedDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) SetTotalCount(v int64) *DescribePurchasedDevicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePurchasedDevicesResponseBodyDevices struct {
	// example:
	//
	// 2019-02-28T17:00:17Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 348*****380-cn-qingdao
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2117*****0447
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1234*****67890
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	// example:
	//
	// dome
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// ipc
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribePurchasedDevicesResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetDescription() *string {
	return s.Description
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetId() *string {
	return s.Id
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetName() *string {
	return s.Name
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetRegion() *string {
	return s.Region
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetRegisterCode() *string {
	return s.RegisterCode
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetSubType() *string {
	return s.SubType
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetType() *string {
	return s.Type
}

func (s *DescribePurchasedDevicesResponseBodyDevices) GetVendor() *string {
	return s.Vendor
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetCreatedTime(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetDescription(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Description = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetGroupId(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetGroupName(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetId(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetName(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Name = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetOrderId(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.OrderId = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetRegion(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Region = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetRegisterCode(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.RegisterCode = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetSubType(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.SubType = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetType(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Type = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) SetVendor(v string) *DescribePurchasedDevicesResponseBodyDevices {
	s.Vendor = &v
	return s
}

func (s *DescribePurchasedDevicesResponseBodyDevices) Validate() error {
	return dara.Validate(s)
}
