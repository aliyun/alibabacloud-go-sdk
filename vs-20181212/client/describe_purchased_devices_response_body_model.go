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
	// The list of devices.
	Devices []*DescribePurchasedDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of devices.
	//
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
	if s.Devices != nil {
		for _, item := range s.Devices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePurchasedDevicesResponseBodyDevices struct {
	// The time the device was created.
	//
	// example:
	//
	// 2019-02-28T17:00:17Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The device description.
	//
	// example:
	//
	// xxx路口摄像头
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group to which the device belongs.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Group name.
	//
	// example:
	//
	// 测试空间
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The device ID.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The device name.
	//
	// example:
	//
	// xxx路口摄像头
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2117*****0447
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The region to which the space belongs. This is the service center.
	//
	// example:
	//
	// cn-qingdao
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The device registration code.
	//
	// example:
	//
	// 1234*****67890
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	// The device subtype. Valid values:
	//
	// - bullet (bullet camera)
	//
	// - dome (dome camera)
	//
	// - ptz (PTZ camera)
	//
	// example:
	//
	// dome
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The device type. Valid values:
	//
	// - ipc (camera)
	//
	// - platform (platform)
	//
	// - ied (intelligent edge device)
	//
	// example:
	//
	// ipc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The device vendor.
	//
	// example:
	//
	// 公司A
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
