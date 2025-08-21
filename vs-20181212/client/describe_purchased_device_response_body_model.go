// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *DescribePurchasedDeviceResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribePurchasedDeviceResponseBody
	GetDescription() *string
	SetGroupId(v string) *DescribePurchasedDeviceResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribePurchasedDeviceResponseBody
	GetGroupName() *string
	SetId(v string) *DescribePurchasedDeviceResponseBody
	GetId() *string
	SetName(v string) *DescribePurchasedDeviceResponseBody
	GetName() *string
	SetOrderId(v string) *DescribePurchasedDeviceResponseBody
	GetOrderId() *string
	SetRegion(v string) *DescribePurchasedDeviceResponseBody
	GetRegion() *string
	SetRegisterCode(v string) *DescribePurchasedDeviceResponseBody
	GetRegisterCode() *string
	SetRequestId(v string) *DescribePurchasedDeviceResponseBody
	GetRequestId() *string
	SetSubType(v string) *DescribePurchasedDeviceResponseBody
	GetSubType() *string
	SetType(v string) *DescribePurchasedDeviceResponseBody
	GetType() *string
	SetVendor(v string) *DescribePurchasedDeviceResponseBody
	GetVendor() *string
}

type DescribePurchasedDeviceResponseBody struct {
	// example:
	//
	// 2018-12-10T21:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 33763****77224964-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
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
	// 12345*****67890
	RegisterCode *string `json:"RegisterCode,omitempty" xml:"RegisterCode,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// dome
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// ipc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 8yd*****qem
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribePurchasedDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedDeviceResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribePurchasedDeviceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribePurchasedDeviceResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePurchasedDeviceResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePurchasedDeviceResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribePurchasedDeviceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribePurchasedDeviceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribePurchasedDeviceResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribePurchasedDeviceResponseBody) GetRegisterCode() *string {
	return s.RegisterCode
}

func (s *DescribePurchasedDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePurchasedDeviceResponseBody) GetSubType() *string {
	return s.SubType
}

func (s *DescribePurchasedDeviceResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribePurchasedDeviceResponseBody) GetVendor() *string {
	return s.Vendor
}

func (s *DescribePurchasedDeviceResponseBody) SetCreatedTime(v string) *DescribePurchasedDeviceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetDescription(v string) *DescribePurchasedDeviceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetGroupId(v string) *DescribePurchasedDeviceResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetGroupName(v string) *DescribePurchasedDeviceResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetId(v string) *DescribePurchasedDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetName(v string) *DescribePurchasedDeviceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetOrderId(v string) *DescribePurchasedDeviceResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetRegion(v string) *DescribePurchasedDeviceResponseBody {
	s.Region = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetRegisterCode(v string) *DescribePurchasedDeviceResponseBody {
	s.RegisterCode = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetRequestId(v string) *DescribePurchasedDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetSubType(v string) *DescribePurchasedDeviceResponseBody {
	s.SubType = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetType(v string) *DescribePurchasedDeviceResponseBody {
	s.Type = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) SetVendor(v string) *DescribePurchasedDeviceResponseBody {
	s.Vendor = &v
	return s
}

func (s *DescribePurchasedDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
