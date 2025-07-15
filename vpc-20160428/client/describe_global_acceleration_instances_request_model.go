// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalAccelerationInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthType(v string) *DescribeGlobalAccelerationInstancesRequest
	GetBandwidthType() *string
	SetGlobalAccelerationInstanceId(v string) *DescribeGlobalAccelerationInstancesRequest
	GetGlobalAccelerationInstanceId() *string
	SetIncludeReservationData(v bool) *DescribeGlobalAccelerationInstancesRequest
	GetIncludeReservationData() *bool
	SetIpAddress(v string) *DescribeGlobalAccelerationInstancesRequest
	GetIpAddress() *string
	SetName(v string) *DescribeGlobalAccelerationInstancesRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeGlobalAccelerationInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGlobalAccelerationInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeGlobalAccelerationInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGlobalAccelerationInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGlobalAccelerationInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeGlobalAccelerationInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGlobalAccelerationInstancesRequest
	GetResourceOwnerId() *int64
	SetServerId(v string) *DescribeGlobalAccelerationInstancesRequest
	GetServerId() *string
	SetServiceLocation(v string) *DescribeGlobalAccelerationInstancesRequest
	GetServiceLocation() *string
	SetStatus(v string) *DescribeGlobalAccelerationInstancesRequest
	GetStatus() *string
}

type DescribeGlobalAccelerationInstancesRequest struct {
	// The bandwidth type of the GA instance. Valid values:
	//
	// 	- **Sharing**
	//
	// 	- **Exclusive*	- (default)
	//
	// example:
	//
	// Exclusive
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-234sljmxaz****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	// Specifies whether to return information about pending orders. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IncludeReservationData *bool `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	// The public IP address of the GA instance.
	//
	// example:
	//
	// 12.xx.xx.78
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the GA instance.
	//
	// example:
	//
	// GA-1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the GA instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the backend service instance.
	//
	// example:
	//
	// i-sxjblddejj9x****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The region of the backend service. Valid values:
	//
	// 	- **china-mainland**
	//
	// 	- **north-america**
	//
	// 	- **asia-pacific**
	//
	// 	- **europe**
	//
	// example:
	//
	// china-mainland
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
	// The status of the GA instance. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Inuse**
	//
	// 	- **Associating**
	//
	// 	- **Unassociating**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGlobalAccelerationInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalAccelerationInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *DescribeGlobalAccelerationInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetBandwidthType(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.BandwidthType = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetGlobalAccelerationInstanceId(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetIncludeReservationData(v bool) *DescribeGlobalAccelerationInstancesRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetIpAddress(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.IpAddress = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetName(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.Name = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetOwnerAccount(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetOwnerId(v int64) *DescribeGlobalAccelerationInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetPageNumber(v int32) *DescribeGlobalAccelerationInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetPageSize(v int32) *DescribeGlobalAccelerationInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetRegionId(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetResourceOwnerAccount(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetResourceOwnerId(v int64) *DescribeGlobalAccelerationInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetServerId(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.ServerId = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetServiceLocation(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.ServiceLocation = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) SetStatus(v string) *DescribeGlobalAccelerationInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeGlobalAccelerationInstancesRequest) Validate() error {
	return dara.Validate(s)
}
