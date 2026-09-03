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
	// The instance ID. The value depends on the resource type (ResourceType) for which you want to query the renewal price:
	//
	// 	- If `ResourceType` is set to `Desktop` (to query the renewal price of a cloud computer), set `InstanceId` to the cloud computer ID.
	//
	// 	- If `ResourceType` is set to `DesktopGroup` (to query the renewal price of a cloud computer pool), set `InstanceId` to the cloud computer pool ID.
	//
	// 	- If `ResourceType` is set to `Bandwidth` (to query the renewal price of premium Internet bandwidth), set `InstanceId` to the premium Internet bandwidth ID.
	//
	// example:
	//
	// ecd-6ldllk9zxcpfhs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance IDs. The values depend on the resource type (ResourceType) for which you want to query the renewal price.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The renewal duration. Valid values of this parameter are determined by the value of `PeriodUnit`.
	//
	// - If `PeriodUnit` is set to `Month`, valid values are 1, 2, 3, and 6.
	//
	// - If `PeriodUnit` is set to `Year`, valid values are 1, 2, and 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration, which is the unit of the `Period` parameter.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Wuying Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the resource ownership user in reseller pattern. You do not need to specify this parameter in non-reseller pattern.
	//
	// example:
	//
	// 1017457975738750
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The resource type.
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
