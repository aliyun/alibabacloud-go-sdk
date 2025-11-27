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
	// Specifies whether to enable the automatic payment feature. Valid values:
	//
	// 	- **true*	- (default): enables the feature. You must make sure that your account balance is sufficient.
	//
	// 	- **false**: disables the feature. An unpaid order is generated.
	//
	// >  If your account balance is insufficient, you can set AutoPay to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.
	//
	// example:
	//
	// true
	AutoPay       *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The type of the change that you want to perform on the instance. Valid values:
	//
	// >  This parameter is optional. The system can automatically determine whether the instance change is an upgrade or a downgrade. If you want to specify this parameter, take note of the following items:
	//
	// 	- **Upgrade*	- (default): upgrades the instance type. Make sure that your account balance is sufficient.
	//
	// 	- **Down**: downgrades the instance type. If the new instance type specified by InstanceType has lower specifications than the current instance type, set Direction to Down.
	//
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and resource inventory.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, the operation is performed.
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
	// The new instance type. For more information about the instance types that are supported by RDS Custom instances, see [Instance types of RDS Custom instances](https://help.aliyun.com/document_detail/2844823.html).
	//
	// example:
	//
	// mysql.i8.large.2cm
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PromotionCode      *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	RebootTime         *string `json:"RebootTime,omitempty" xml:"RebootTime,omitempty"`
	RebootWhenFinished *bool   `json:"RebootWhenFinished,omitempty" xml:"RebootWhenFinished,omitempty"`
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
