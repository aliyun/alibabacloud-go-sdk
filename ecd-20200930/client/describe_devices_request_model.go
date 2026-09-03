// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDevicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdDomain(v string) *DescribeDevicesRequest
	GetAdDomain() *string
	SetClientType(v string) *DescribeDevicesRequest
	GetClientType() *string
	SetDeviceId(v string) *DescribeDevicesRequest
	GetDeviceId() *string
	SetDirectoryId(v string) *DescribeDevicesRequest
	GetDirectoryId() *string
	SetEndUserId(v string) *DescribeDevicesRequest
	GetEndUserId() *string
	SetPageNumber(v int32) *DescribeDevicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDevicesRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeDevicesRequest
	GetRegion() *string
	SetUserType(v string) *DescribeDevicesRequest
	GetUserType() *string
}

type DescribeDevicesRequest struct {
	// The address of the Active Directory (AD) office network.
	//
	// example:
	//
	// xn--0zw****
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The type of the client.
	//
	// Valid values:
	//
	// 	- 1: hardware client.
	//
	// 	- 2: software client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The ID of the device. The serial number (SN) of the hardware client or the UUID of the software client.
	//
	// example:
	//
	// 5F52817BE267A43C608D245070D2****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The ID of the convenient office network.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the bound user.
	//
	// example:
	//
	// moli
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by WUYING Workspace.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The account type of the user.
	//
	// Valid values:
	//
	// 	- AD: enterprise AD account.
	//
	// 	- SIMPLE: convenience account
	//
	// example:
	//
	// SIMPLE
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeDevicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDevicesRequest) GetAdDomain() *string {
	return s.AdDomain
}

func (s *DescribeDevicesRequest) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeDevicesRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeDevicesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDevicesRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDevicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDevicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDevicesRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDevicesRequest) GetUserType() *string {
	return s.UserType
}

func (s *DescribeDevicesRequest) SetAdDomain(v string) *DescribeDevicesRequest {
	s.AdDomain = &v
	return s
}

func (s *DescribeDevicesRequest) SetClientType(v string) *DescribeDevicesRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeDevicesRequest) SetDeviceId(v string) *DescribeDevicesRequest {
	s.DeviceId = &v
	return s
}

func (s *DescribeDevicesRequest) SetDirectoryId(v string) *DescribeDevicesRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDevicesRequest) SetEndUserId(v string) *DescribeDevicesRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageNumber(v int32) *DescribeDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageSize(v int32) *DescribeDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesRequest) SetRegion(v string) *DescribeDevicesRequest {
	s.Region = &v
	return s
}

func (s *DescribeDevicesRequest) SetUserType(v string) *DescribeDevicesRequest {
	s.UserType = &v
	return s
}

func (s *DescribeDevicesRequest) Validate() error {
	return dara.Validate(s)
}
