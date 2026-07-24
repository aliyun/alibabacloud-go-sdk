// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICloudPhoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateAICloudPhoneRequest
	GetAmount() *int32
	SetAutoPay(v bool) *CreateAICloudPhoneRequest
	GetAutoPay() *bool
	SetBandwidthPackageId(v string) *CreateAICloudPhoneRequest
	GetBandwidthPackageId() *string
	SetBizRegionId(v string) *CreateAICloudPhoneRequest
	GetBizRegionId() *string
	SetImageId(v string) *CreateAICloudPhoneRequest
	GetImageId() *string
	SetInstanceGroupName(v string) *CreateAICloudPhoneRequest
	GetInstanceGroupName() *string
	SetInstanceGroupSpec(v string) *CreateAICloudPhoneRequest
	GetInstanceGroupSpec() *string
	SetPeriod(v int32) *CreateAICloudPhoneRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateAICloudPhoneRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateAICloudPhoneRequest
	GetPolicyGroupId() *string
	SetPromotionId(v string) *CreateAICloudPhoneRequest
	GetPromotionId() *string
}

type CreateAICloudPhoneRequest struct {
	// The quantity to purchase.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The bandwidth package ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// np-791ncq8qcuoopxxxx
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The region ID for the purchase.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// imgc-0aae4rxwoktvr851h
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The instance group specification. Valid values:
	//
	// - STANDARD: standard.
	//
	// - MEDIUM: advanced.
	//
	// This parameter is required.
	//
	// example:
	//
	// MEDIUM
	InstanceGroupSpec *string `json:"InstanceGroupSpec,omitempty" xml:"InstanceGroupSpec,omitempty"`
	// The purchase duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the purchase duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The policy group ID.
	//
	// example:
	//
	// pg-0bjrh3oxk2q0xxxxx
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// youhuiquan_promotion_xxx
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s CreateAICloudPhoneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAICloudPhoneRequest) GoString() string {
	return s.String()
}

func (s *CreateAICloudPhoneRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateAICloudPhoneRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateAICloudPhoneRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateAICloudPhoneRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAICloudPhoneRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateAICloudPhoneRequest) GetInstanceGroupName() *string {
	return s.InstanceGroupName
}

func (s *CreateAICloudPhoneRequest) GetInstanceGroupSpec() *string {
	return s.InstanceGroupSpec
}

func (s *CreateAICloudPhoneRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAICloudPhoneRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateAICloudPhoneRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateAICloudPhoneRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateAICloudPhoneRequest) SetAmount(v int32) *CreateAICloudPhoneRequest {
	s.Amount = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetAutoPay(v bool) *CreateAICloudPhoneRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetBandwidthPackageId(v string) *CreateAICloudPhoneRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetBizRegionId(v string) *CreateAICloudPhoneRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetImageId(v string) *CreateAICloudPhoneRequest {
	s.ImageId = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetInstanceGroupName(v string) *CreateAICloudPhoneRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetInstanceGroupSpec(v string) *CreateAICloudPhoneRequest {
	s.InstanceGroupSpec = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetPeriod(v int32) *CreateAICloudPhoneRequest {
	s.Period = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetPeriodUnit(v string) *CreateAICloudPhoneRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetPolicyGroupId(v string) *CreateAICloudPhoneRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateAICloudPhoneRequest) SetPromotionId(v string) *CreateAICloudPhoneRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAICloudPhoneRequest) Validate() error {
	return dara.Validate(s)
}
