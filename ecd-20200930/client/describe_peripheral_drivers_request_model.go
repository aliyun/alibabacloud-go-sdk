// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePeripheralDriversRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrand(v string) *DescribePeripheralDriversRequest
	GetBrand() *string
	SetDeviceType(v string) *DescribePeripheralDriversRequest
	GetDeviceType() *string
	SetDriverIds(v []*string) *DescribePeripheralDriversRequest
	GetDriverIds() []*string
	SetFilter(v string) *DescribePeripheralDriversRequest
	GetFilter() *string
	SetMaxResults(v int32) *DescribePeripheralDriversRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePeripheralDriversRequest
	GetNextToken() *string
	SetOwnerType(v string) *DescribePeripheralDriversRequest
	GetOwnerType() *string
	SetPageNumber(v int32) *DescribePeripheralDriversRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePeripheralDriversRequest
	GetPageSize() *int32
}

type DescribePeripheralDriversRequest struct {
	// The brand identifier. Exact match. The value is based on the actual configuration and has no fixed enumeration. If this parameter is not specified, the brand is not limited. The example is used only to illustrate the format.
	//
	// example:
	//
	// hp
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// The device type identifier. Exact match. The value is based on the actual configuration. For example, printer indicates a printer. If this parameter is not specified, the device type is not limited.
	//
	// example:
	//
	// printer
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The list of driver IDs. If this parameter is not specified or an empty array is passed, the driver ID is not limited. Only the drivers that are visible to the current account and match the conditions are returned. Unmatched IDs do not generate corresponding records.
	DriverIds []*string `json:"DriverIds,omitempty" xml:"DriverIds,omitempty" type:"Repeated"`
	// The search keyword. Matches the driver ID, brand identifier, driver name, description, device type, or brand display name. A match in any field is considered a hit. The percent sign (%) matches characters of any length, and the underscore (_) matches a single character. If this parameter is not specified, keyword filtering is not performed.
	//
	// example:
	//
	// LaserJet
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The reserved parameter. This parameter is not used for querying or pagination. Do not specify this parameter. Use PageSize to set the number of entries to return on each page. The sample value 20 is used only to demonstrate the integer type. It is not the default value of this parameter and does not take effect even if specified.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The reserved parameter. Token-based pagination is not supported. Do not specify this parameter. Use PageNumber to specify the page number. The sample value token-for-format-only is used only to demonstrate the string type and is not a usable pagination token.
	//
	// example:
	//
	// token-for-format-only
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The driver ownership. Valid values: WUYING: Wuying official drivers. CUSTOMER: custom drivers of the current account. If this parameter is not specified, both types of drivers are queried.
	//
	// example:
	//
	// CUSTOMER
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The page number. Specify a value starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 500. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePeripheralDriversRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePeripheralDriversRequest) GoString() string {
	return s.String()
}

func (s *DescribePeripheralDriversRequest) GetBrand() *string {
	return s.Brand
}

func (s *DescribePeripheralDriversRequest) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribePeripheralDriversRequest) GetDriverIds() []*string {
	return s.DriverIds
}

func (s *DescribePeripheralDriversRequest) GetFilter() *string {
	return s.Filter
}

func (s *DescribePeripheralDriversRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePeripheralDriversRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePeripheralDriversRequest) GetOwnerType() *string {
	return s.OwnerType
}

func (s *DescribePeripheralDriversRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePeripheralDriversRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePeripheralDriversRequest) SetBrand(v string) *DescribePeripheralDriversRequest {
	s.Brand = &v
	return s
}

func (s *DescribePeripheralDriversRequest) SetDeviceType(v string) *DescribePeripheralDriversRequest {
	s.DeviceType = &v
	return s
}

func (s *DescribePeripheralDriversRequest) SetDriverIds(v []*string) *DescribePeripheralDriversRequest {
	s.DriverIds = v
	return s
}

func (s *DescribePeripheralDriversRequest) SetFilter(v string) *DescribePeripheralDriversRequest {
	s.Filter = &v
	return s
}

func (s *DescribePeripheralDriversRequest) SetMaxResults(v int32) *DescribePeripheralDriversRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePeripheralDriversRequest) SetNextToken(v string) *DescribePeripheralDriversRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePeripheralDriversRequest) SetOwnerType(v string) *DescribePeripheralDriversRequest {
	s.OwnerType = &v
	return s
}

func (s *DescribePeripheralDriversRequest) SetPageNumber(v int32) *DescribePeripheralDriversRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePeripheralDriversRequest) SetPageSize(v int32) *DescribePeripheralDriversRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePeripheralDriversRequest) Validate() error {
	return dara.Validate(s)
}
