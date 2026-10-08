// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyRCInstanceRequest
	GetAutoPay() *bool
	SetAutoUseCoupon(v bool) *ModifyRCInstanceRequest
	GetAutoUseCoupon() *bool
	SetBusinessInfo(v string) *ModifyRCInstanceRequest
	GetBusinessInfo() *string
	SetDirection(v string) *ModifyRCInstanceRequest
	GetDirection() *string
	SetDryRun(v bool) *ModifyRCInstanceRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ModifyRCInstanceRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyRCInstanceRequest
	GetInstanceType() *string
	SetPromotionCode(v string) *ModifyRCInstanceRequest
	GetPromotionCode() *string
	SetRebootTime(v string) *ModifyRCInstanceRequest
	GetRebootTime() *string
	SetRebootWhenFinished(v bool) *ModifyRCInstanceRequest
	GetRebootWhenFinished() *bool
	SetRegionId(v string) *ModifyRCInstanceRequest
	GetRegionId() *string
}

type ModifyRCInstanceRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **true*	- (default): Automatic payment is enabled. Make sure that your account balance is sufficient.
	//
	// - **false**: An order is generated but payment is not automatically made.
	//
	// > If your payment method balance is insufficient, set the parameter AutoPay to false. An unpaid order is generated, and you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- **true*	- (default): Coupons are automatically used.
	//
	// 	- **false**: Coupons are not used.
	//
	// > If you use coupons and then perform a downgrade, the amount deducted by coupons is not refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BusinessInfo  *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The type of the Upgrade/Downgrade. Valid values:
	//
	// > This parameter does not need to be uploaded. The system can automatically determine whether the change is an upgrade or a downgrade. If you upload this parameter, follow the rules below.
	//
	// - **Up*	- (default): Upgrades the instance type. Make sure that your account payment method balance is sufficient.
	//
	// - **Down**: Downgrades the instance type. Set Direction to down when the instance type specified by InstanceType is lower than the current instance type.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: Performs a dry run without creating the instance. The system checks items such as the request parameters, request format, service limits, and available resources.
	//
	// 	- **false*	- (default): Sends the request. If the request passes the check, the instance is created.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf62br2491p5l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The target instance type. For information about the instance types supported by RDS Custom instances, see [RDS Custom instance types](https://help.aliyun.com/document_detail/2844823.html).
	//
	// example:
	//
	// mysql.i8.large.2cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 72329885****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The restart time of the instance.
	//
	// - If **RebootWhenFinished*	- is set to **false*	- and the instance status is **Running**, you **must*	- set a restart time within 48 hours.
	//
	// - The time follows the ISO 8601 standard in UTC+0. Format: `yyyy-MM-ddTHH:mmZ`.
	//
	// example:
	//
	// 2025-04-03T12:05Z
	RebootTime *string `json:"RebootTime,omitempty" xml:"RebootTime,omitempty"`
	// Specifies whether to immediately restart the instance after the specification change is complete. Valid values:
	//
	// - **true*	- (default): The instance is restarted immediately.
	//
	// - **false**: The instance is not restarted.
	//
	// > If the instance is in the **Stopped*	- state, the instance remains in the Stopped state and is not restarted even if you set `RebootWhenFinished=true`.
	//
	// example:
	//
	// true
	RebootWhenFinished *bool `json:"RebootWhenFinished,omitempty" xml:"RebootWhenFinished,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hagnzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyRCInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyRCInstanceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *ModifyRCInstanceRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyRCInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyRCInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyRCInstanceRequest) GetRebootTime() *string {
	return s.RebootTime
}

func (s *ModifyRCInstanceRequest) GetRebootWhenFinished() *bool {
	return s.RebootWhenFinished
}

func (s *ModifyRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceRequest) SetAutoPay(v bool) *ModifyRCInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetAutoUseCoupon(v bool) *ModifyRCInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetBusinessInfo(v string) *ModifyRCInstanceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetDirection(v string) *ModifyRCInstanceRequest {
	s.Direction = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetDryRun(v bool) *ModifyRCInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetInstanceId(v string) *ModifyRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetInstanceType(v string) *ModifyRCInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetPromotionCode(v string) *ModifyRCInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetRebootTime(v string) *ModifyRCInstanceRequest {
	s.RebootTime = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetRebootWhenFinished(v bool) *ModifyRCInstanceRequest {
	s.RebootWhenFinished = &v
	return s
}

func (s *ModifyRCInstanceRequest) SetRegionId(v string) *ModifyRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
