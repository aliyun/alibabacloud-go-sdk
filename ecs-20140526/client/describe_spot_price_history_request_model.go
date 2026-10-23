// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotPriceHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSpotPriceHistoryRequest
	GetEndTime() *string
	SetInstanceType(v string) *DescribeSpotPriceHistoryRequest
	GetInstanceType() *string
	SetIoOptimized(v string) *DescribeSpotPriceHistoryRequest
	GetIoOptimized() *string
	SetNetworkType(v string) *DescribeSpotPriceHistoryRequest
	GetNetworkType() *string
	SetOSType(v string) *DescribeSpotPriceHistoryRequest
	GetOSType() *string
	SetOffset(v int32) *DescribeSpotPriceHistoryRequest
	GetOffset() *int32
	SetOwnerAccount(v string) *DescribeSpotPriceHistoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSpotPriceHistoryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSpotPriceHistoryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSpotPriceHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSpotPriceHistoryRequest
	GetResourceOwnerId() *int64
	SetSpotDuration(v int32) *DescribeSpotPriceHistoryRequest
	GetSpotDuration() *int32
	SetStartTime(v string) *DescribeSpotPriceHistoryRequest
	GetStartTime() *string
	SetZoneId(v string) *DescribeSpotPriceHistoryRequest
	GetZoneId() *string
}

type DescribeSpotPriceHistoryRequest struct {
	// The end time to query the historical prices of spot instances. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// Default value: empty. An empty value indicates the current time.
	//
	// example:
	//
	// 2017-08-22T08:45:08Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.t1.xsmall
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the spot instance is an I/O optimized instance. Valid values:
	//
	// - optimized: The spot instance is an I/O optimized instance.
	//
	// - none: The spot instance is a non-I/O optimized instance.
	//
	// Default value for Generation I instances: none.
	//
	// Default value for other instance family types: optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The network type of the spot instance. Valid values:
	//
	// - vpc: VPC.
	//
	// - classic: classic network. This feature is retired. For more information, see [Retirement notice](https://help.aliyun.com/document_detail/2833134.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The operating system platform of the instance. Valid values:
	//
	// - linux
	//
	// - windows
	//
	// example:
	//
	// linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The starting row for the query.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	Offset       *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// - 1: Alibaba Cloud guarantees that the instance is not automatically released within 1 hour after it is created. After 1 hour, the system automatically compares your bid with the market price and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// - 0: Alibaba Cloud does not guarantee that the instance runs for 1 hour after it is created. The system automatically compares your bid with the market price and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// Alibaba Cloud sends you a notification through an ECS system event 5 minutes before the instance is revoked. Spot instances are billed by the second. Select an appropriate protection period based on the execution duration of specific tasks.
	//
	// > This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit or SpotAsPriceGo.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The start time to query the historical prices of spot instances. The time span between the start time and the specified end time cannot exceed 30 days. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// Default value: empty. An empty value indicates 3 hours before the end time.
	//
	// example:
	//
	// 2017-08-22T08:45:08Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSpotPriceHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotPriceHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpotPriceHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSpotPriceHistoryRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSpotPriceHistoryRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeSpotPriceHistoryRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeSpotPriceHistoryRequest) GetOSType() *string {
	return s.OSType
}

func (s *DescribeSpotPriceHistoryRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeSpotPriceHistoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSpotPriceHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSpotPriceHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSpotPriceHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSpotPriceHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSpotPriceHistoryRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeSpotPriceHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSpotPriceHistoryRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSpotPriceHistoryRequest) SetEndTime(v string) *DescribeSpotPriceHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetInstanceType(v string) *DescribeSpotPriceHistoryRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetIoOptimized(v string) *DescribeSpotPriceHistoryRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetNetworkType(v string) *DescribeSpotPriceHistoryRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOSType(v string) *DescribeSpotPriceHistoryRequest {
	s.OSType = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOffset(v int32) *DescribeSpotPriceHistoryRequest {
	s.Offset = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOwnerAccount(v string) *DescribeSpotPriceHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOwnerId(v int64) *DescribeSpotPriceHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetRegionId(v string) *DescribeSpotPriceHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetResourceOwnerAccount(v string) *DescribeSpotPriceHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetResourceOwnerId(v int64) *DescribeSpotPriceHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetSpotDuration(v int32) *DescribeSpotPriceHistoryRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetStartTime(v string) *DescribeSpotPriceHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetZoneId(v string) *DescribeSpotPriceHistoryRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) Validate() error {
	return dara.Validate(s)
}
