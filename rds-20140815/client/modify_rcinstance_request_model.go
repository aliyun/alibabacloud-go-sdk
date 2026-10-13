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
	// - **true*	- (default): enables automatic payment. Make sure that your account balance is sufficient.
	//
	// - **false**: generates an order without deducting fees.
	//
	// > If your account balance is insufficient, you can set the AutoPay parameter to false. An unpaid order is generated. You can log on to the ApsaraDB RDS console to complete the payment.
	//
	// >
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to automatically use coupons. Valid values:
	//
	// 	- **true**: automatically uses coupons. This is the default value.
	//
	// 	- **false**: does not automatically use coupons.
	//
	// > If you use a coupon and then downgrade the instance type, the amount deducted by the voucher is not refunded.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BusinessInfo  *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The instance upgrade or downgrade type. Valid values:
	//
	// > You do not need to specify this parameter. The system automatically determines whether to upgrade or downgrade the instance type. If you want to specify it, follow these rules:
	//
	// - **Up*	- (default): upgrades the instance type. Make sure that your account balance is sufficient.
	//
	// - **Down**: downgrades the instance type. When the instance type specified in InstanceType is lower than the current instance type, set Direction to Down.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Specifies whether to perform a precheck on the instance creation operation. Valid values:
	//
	// 	- **true**: performs a precheck without creating an instance. Check items include request parameters, request format, service limits, and inventory.
	//
	// 	- **false**: sends a normal request and directly creates the instance after the request passes the precheck. This is the default value.
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
	// The destination instance type. For more information about the instance types supported by RDS Target instances, see [RDS Target instance types](https://help.aliyun.com/document_detail/2844823.html).
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
	// - If **RebootWhenFinished*	- is set to **false*	- and the instance is in the **Running*	- state, you **must*	- specify a restart time within 48 hours.
	//
	// - Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-04-03T12:05Z
	RebootTime *string `json:"RebootTime,omitempty" xml:"RebootTime,omitempty"`
	// Specifies whether to immediately restart the instance after the instance type change is complete. Valid values:
	//
	// - **true**: restarts the instance immediately. This is the default value.
	//
	// - **false**: does not restart the instance immediately.
	//
	// > If the instance is in the **Stopped*	- state, it remains in the Stopped state and is not restarted even if you set RebootWhenFinished to true.
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
