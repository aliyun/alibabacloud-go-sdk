// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRenewalPriceRequest
	GetInstanceId() *string
	SetInstanceIds(v []*string) *DescribeRenewalPriceRequest
	GetInstanceIds() []*string
	SetPeriod(v int32) *DescribeRenewalPriceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeRenewalPriceRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeRenewalPriceRequest
	GetPromotionId() *string
	SetRegionId(v string) *DescribeRenewalPriceRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *DescribeRenewalPriceRequest
	GetResellerOwnerUid() *int64
	SetResourceType(v string) *DescribeRenewalPriceRequest
	GetResourceType() *string
}

type DescribeRenewalPriceRequest struct {
	// The instance ID. The value you specify depends on the resource type (ResourceType) you\\"re querying the renewal price for.
	//
	// 	- When `ResourceType` is set to `Desktop`, you must provide the cloud computer ID as the value of `InstanceId`.
	//
	// 	- When `ResourceType` is set to `DesktopGroup`, you must provide the share ID as the value of `InstanceId`.
	//
	// 	- When `ResourceType` is set to `Bandwidth`, you must provide the ID of the premium bandwidth plan as the value of `InstanceId`.
	//
	// example:
	//
	// ecd-6ldllk9zxcpfhs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IDs. The value you specify depends on the resource type (ResourceType) you\\"re querying the renewal price for.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The renewal duration. The valid values for this parameter depend on the value of `PeriodUnit`.
	//
	// 	- If you set `PeriodUnit` to `Month`, set the value of this parameter to 1, 2, 3, or 6.
	//
	// 	- If you set `PeriodUnit` to `Year`, set the value of this parameter to 1, 2, or 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration specified by `Period`.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit  *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- Desktop (default): cloud computers.
	//
	// 	- Bandwidth: premium bandwidth plans.
	//
	// 	- DesktopGroup: cloud computer shares.
	//
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRenewalPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRenewalPriceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeRenewalPriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeRenewalPriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeRenewalPriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeRenewalPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRenewalPriceRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DescribeRenewalPriceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRenewalPriceRequest) SetInstanceId(v string) *DescribeRenewalPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetInstanceIds(v []*string) *DescribeRenewalPriceRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPeriod(v int32) *DescribeRenewalPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPeriodUnit(v string) *DescribeRenewalPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetPromotionId(v string) *DescribeRenewalPriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetRegionId(v string) *DescribeRenewalPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResellerOwnerUid(v int64) *DescribeRenewalPriceRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DescribeRenewalPriceRequest) SetResourceType(v string) *DescribeRenewalPriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRenewalPriceRequest) Validate() error {
	return dara.Validate(s)
}
