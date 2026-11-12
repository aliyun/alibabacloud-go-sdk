// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateInstancesRequest
	GetAmount() *int32
	SetAutoRenew(v bool) *CreateInstancesRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateInstancesRequest
	GetAutoRenewPeriod() *int32
	SetChargeType(v string) *CreateInstancesRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateInstancesRequest
	GetClientToken() *string
	SetDataDiskSize(v int64) *CreateInstancesRequest
	GetDataDiskSize() *int64
	SetImageId(v string) *CreateInstancesRequest
	GetImageId() *string
	SetPeriod(v int32) *CreateInstancesRequest
	GetPeriod() *int32
	SetPlanId(v string) *CreateInstancesRequest
	GetPlanId() *string
	SetRegionId(v string) *CreateInstancesRequest
	GetRegionId() *string
}

type CreateInstancesRequest struct {
	// The number of Simple Application Server instances to create. Valid values: 1 to 20.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable auto-renewal upon expiration. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. This parameter is required only when AutoRenew is set to true. Unit: months. Valid values: 1, 3, 6, 12, 24, and 36.
	//
	// example:
	//
	// 3
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The billing method. Set the value to PrePaid. Currently, only the subscription billing method is supported.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The size of the data cloud disk to be mounted. Unit: GB. Valid values: 0 to 16380. The value must be a multiple of 20.
	//
	// <props="china">
	//
	// - A value of 0 indicates that no data cloud disk is mounted.
	//
	// - If the selected plan uses a PL0 ESSD, the minimum size of the data cloud disk is 40 GB.
	//
	// - If the selected plan uses a standard SSD, the minimum size of the data cloud disk is 20 GB.
	//
	//
	//
	// <props="intl">
	//
	// - A value of 0 indicates that no data cloud disk is mounted.
	//
	// - If the selected plan uses a standard SSD, the minimum size of the data cloud disk is 20 GB.
	//
	//
	//
	// Default value: 0.
	//
	// example:
	//
	// 20
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The image ID. You can call the [ListImages](https://help.aliyun.com/document_detail/189313.html) operation to query available image IDs in a specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// e2c9c365024a44369c9b955a998a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The subscription period of the resource. Unit: months. Valid values: 1, 3, 6, 12, 24, and 36.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The plan ID. You can call the [ListPlans](https://help.aliyun.com/document_detail/189314.html) operation to query all plans in a specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// swas.s1.c1m1s40b3t05
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateInstancesRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstancesRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstancesRequest) GetDataDiskSize() *int64 {
	return s.DataDiskSize
}

func (s *CreateInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateInstancesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstancesRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstancesRequest) SetAmount(v int32) *CreateInstancesRequest {
	s.Amount = &v
	return s
}

func (s *CreateInstancesRequest) SetAutoRenew(v bool) *CreateInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstancesRequest) SetAutoRenewPeriod(v int32) *CreateInstancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstancesRequest) SetChargeType(v string) *CreateInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstancesRequest) SetClientToken(v string) *CreateInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstancesRequest) SetDataDiskSize(v int64) *CreateInstancesRequest {
	s.DataDiskSize = &v
	return s
}

func (s *CreateInstancesRequest) SetImageId(v string) *CreateInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstancesRequest) SetPeriod(v int32) *CreateInstancesRequest {
	s.Period = &v
	return s
}

func (s *CreateInstancesRequest) SetPlanId(v string) *CreateInstancesRequest {
	s.PlanId = &v
	return s
}

func (s *CreateInstancesRequest) SetRegionId(v string) *CreateInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstancesRequest) Validate() error {
	return dara.Validate(s)
}
