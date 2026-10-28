// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePeripheralDriversResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribePeripheralDriversResponseBody
	GetCount() *int32
	SetDriverInfos(v []*DescribePeripheralDriversResponseBodyDriverInfos) *DescribePeripheralDriversResponseBody
	GetDriverInfos() []*DescribePeripheralDriversResponseBodyDriverInfos
	SetMaxResults(v int32) *DescribePeripheralDriversResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePeripheralDriversResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribePeripheralDriversResponseBody
	GetRequestId() *string
}

type DescribePeripheralDriversResponseBody struct {
	// The total number of matched drivers, not the length of the list on the current page. If no data exists on the current page, 0 may be returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of driver information on the current page. An empty list is returned if no data exists.
	DriverInfos []*DescribePeripheralDriversResponseBodyDriverInfos `json:"DriverInfos,omitempty" xml:"DriverInfos,omitempty" type:"Repeated"`
	// The reserved field. This field does not provide a valid return value and may not be returned. This operation uses PageSize and PageNumber for pagination. Do not rely on this field. The sample value 20 is used only to demonstrate the integer type and does not represent the actual return value, default value, or page size of the current query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The reserved field. This field does not provide token-based pagination and may not be returned. Do not rely on this field for subsequent queries. The sample value token-for-format-only is used only to demonstrate the string type and is not an actual return value or a usable pagination token.
	//
	// example:
	//
	// token-for-format-only
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 00000000-1111-4222-8333-444444444444
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePeripheralDriversResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePeripheralDriversResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePeripheralDriversResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribePeripheralDriversResponseBody) GetDriverInfos() []*DescribePeripheralDriversResponseBodyDriverInfos {
	return s.DriverInfos
}

func (s *DescribePeripheralDriversResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePeripheralDriversResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePeripheralDriversResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePeripheralDriversResponseBody) SetCount(v int32) *DescribePeripheralDriversResponseBody {
	s.Count = &v
	return s
}

func (s *DescribePeripheralDriversResponseBody) SetDriverInfos(v []*DescribePeripheralDriversResponseBodyDriverInfos) *DescribePeripheralDriversResponseBody {
	s.DriverInfos = v
	return s
}

func (s *DescribePeripheralDriversResponseBody) SetMaxResults(v int32) *DescribePeripheralDriversResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribePeripheralDriversResponseBody) SetNextToken(v string) *DescribePeripheralDriversResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePeripheralDriversResponseBody) SetRequestId(v string) *DescribePeripheralDriversResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePeripheralDriversResponseBody) Validate() error {
	if s.DriverInfos != nil {
		for _, item := range s.DriverInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePeripheralDriversResponseBodyDriverInfos struct {
	// The brand to which the driver belongs.
	//
	// example:
	//
	// hp
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// The time when the driver record was created. The time is in the ISO 8601 (RFC 3339) format with time zone offset. The time zone offset is subject to the returned value. If the time information does not exist, this field may be empty or not returned.
	//
	// example:
	//
	// 2026-09-01T10:30:00+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The device type to which the driver applies.
	//
	// example:
	//
	// printer
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The URL of the brand icon. If no icon is configured, this field may be empty or not returned. The example is a placeholder address.
	//
	// example:
	//
	// https://example.com/icons/printer.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The driver ID, which can be used for subsequent queries.
	//
	// example:
	//
	// 11111111-2222-4333-8444-555555555555
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The driver name.
	//
	// example:
	//
	// HP Universal Printing PCL 6
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system to which the driver applies, such as Windows. The actual returned value prevails.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The driver ownership. Valid values: WUYING: Wuying official drivers. CUSTOMER: custom drivers of the current account.
	//
	// example:
	//
	// WUYING
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The driver source. Valid values: OpsApp: uploaded from the management console. WuyingHelper: uploaded from Wuying Assistant. Wuying: from Wuying. Unrecognized sources may also be classified as Wuying. To distinguish between official and custom drivers, use OwnerType.
	//
	// example:
	//
	// Wuying
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePeripheralDriversResponseBodyDriverInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribePeripheralDriversResponseBodyDriverInfos) GoString() string {
	return s.String()
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetBrand() *string {
	return s.Brand
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetIcon() *string {
	return s.Icon
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetId() *string {
	return s.Id
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetName() *string {
	return s.Name
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetOsType() *string {
	return s.OsType
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) GetSource() *string {
	return s.Source
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetBrand(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.Brand = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetCreateTime(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.CreateTime = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetDeviceType(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.DeviceType = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetIcon(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.Icon = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetId(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.Id = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetName(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.Name = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetOsType(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.OsType = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetOwnerType(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.OwnerType = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) SetSource(v string) *DescribePeripheralDriversResponseBodyDriverInfos {
	s.Source = &v
	return s
}

func (s *DescribePeripheralDriversResponseBodyDriverInfos) Validate() error {
	return dara.Validate(s)
}
