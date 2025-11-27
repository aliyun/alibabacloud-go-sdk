// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcuType(v string) *RunRCInstancesShrinkRequest
	GetAcuType() *string
	SetAmount(v int32) *RunRCInstancesShrinkRequest
	GetAmount() *int32
	SetAutoPay(v bool) *RunRCInstancesShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RunRCInstancesShrinkRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *RunRCInstancesShrinkRequest
	GetAutoUseCoupon() *bool
	SetClientToken(v string) *RunRCInstancesShrinkRequest
	GetClientToken() *string
	SetCreateAckEdgeParamShrink(v string) *RunRCInstancesShrinkRequest
	GetCreateAckEdgeParamShrink() *string
	SetCreateExtraParam(v string) *RunRCInstancesShrinkRequest
	GetCreateExtraParam() *string
	SetCreateMode(v string) *RunRCInstancesShrinkRequest
	GetCreateMode() *string
	SetDataDiskShrink(v string) *RunRCInstancesShrinkRequest
	GetDataDiskShrink() *string
	SetDeletionProtection(v bool) *RunRCInstancesShrinkRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *RunRCInstancesShrinkRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *RunRCInstancesShrinkRequest
	GetDescription() *string
	SetDryRun(v bool) *RunRCInstancesShrinkRequest
	GetDryRun() *bool
	SetHostName(v string) *RunRCInstancesShrinkRequest
	GetHostName() *string
	SetImageId(v string) *RunRCInstancesShrinkRequest
	GetImageId() *string
	SetInstanceChargeType(v string) *RunRCInstancesShrinkRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *RunRCInstancesShrinkRequest
	GetInstanceName() *string
	SetInstanceType(v string) *RunRCInstancesShrinkRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *RunRCInstancesShrinkRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int32) *RunRCInstancesShrinkRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *RunRCInstancesShrinkRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *RunRCInstancesShrinkRequest
	GetKeyPairName() *string
	SetPassword(v string) *RunRCInstancesShrinkRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *RunRCInstancesShrinkRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *RunRCInstancesShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RunRCInstancesShrinkRequest
	GetPeriodUnit() *string
	SetPromotionCode(v string) *RunRCInstancesShrinkRequest
	GetPromotionCode() *string
	SetRegionId(v string) *RunRCInstancesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RunRCInstancesShrinkRequest
	GetResourceGroupId() *string
	SetScheduledRule(v string) *RunRCInstancesShrinkRequest
	GetScheduledRule() *string
	SetSecurityEnhancementStrategy(v string) *RunRCInstancesShrinkRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *RunRCInstancesShrinkRequest
	GetSecurityGroupId() *string
	SetSpotStrategy(v string) *RunRCInstancesShrinkRequest
	GetSpotStrategy() *string
	SetSupportCase(v string) *RunRCInstancesShrinkRequest
	GetSupportCase() *string
	SetSystemDiskShrink(v string) *RunRCInstancesShrinkRequest
	GetSystemDiskShrink() *string
	SetTag(v []*RunRCInstancesShrinkRequestTag) *RunRCInstancesShrinkRequest
	GetTag() []*RunRCInstancesShrinkRequestTag
	SetUserData(v string) *RunRCInstancesShrinkRequest
	GetUserData() *string
	SetUserDataInBase64(v bool) *RunRCInstancesShrinkRequest
	GetUserDataInBase64() *bool
	SetVSwitchId(v string) *RunRCInstancesShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *RunRCInstancesShrinkRequest
	GetZoneId() *string
}

type RunRCInstancesShrinkRequest struct {
	// example:
	//
	// gn8is
	AcuType *string `json:"AcuType,omitempty" xml:"AcuType,omitempty"`
	// The number of RDS Custom instances that you want to create. The parameter is available if you want to create multiple RDS Custom instances at a time.
	//
	// Valid values: **1*	- to **10**. Default value: **1**.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable the automatic payment feature. Valid values:
	//
	// 	- **true*	- (default): enables the feature. Make sure that your account balance is sufficient.
	//
	// 	- **false**: disables the feature. An unpaid order is generated.
	//
	// >  If your account balance is insufficient, you can set the AutoPay parameter to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenew     *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreateAckEdgeParamShrink *string `json:"CreateAckEdgeParam,omitempty" xml:"CreateAckEdgeParam,omitempty"`
	CreateExtraParam         *string `json:"CreateExtraParam,omitempty" xml:"CreateExtraParam,omitempty"`
	CreateMode               *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	// The information about the data disks.
	DataDiskShrink     *string `json:"DataDisk,omitempty" xml:"DataDisk,omitempty"`
	DeletionProtection *bool   `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-uf6670sipmph5j5b6ke4
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance description. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// Instance_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and insufficient inventory errors.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, the instance is directly created.
	//
	// example:
	//
	// false
	DryRun   *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the image used by the instance.
	//
	// example:
	//
	// image-dsvjzw2ii8n4fvr6de
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing method of the instance. Set the value to **Prepaid**, which indicates the subscription billing method.
	//
	// example:
	//
	// Prepaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name.
	//
	// example:
	//
	// ceshi
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. For more information about the instance types that are supported by RDS Custom instances, see [Instance types for RDS Custom instances](https://help.aliyun.com/document_detail/2844823.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.i8.large.2cm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// null
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// null
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// null
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the AccessKey pair. You can specify only one name.
	//
	// example:
	//
	// dell5502
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The password of the account that is used to log on to the instance.
	//
	// example:
	//
	// 2F9e9@a69c!e18b569c8
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordInherit *bool   `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription duration of the instance. Default value: **1**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month*	- (default)
	//
	// example:
	//
	// Year
	PeriodUnit    *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// {"rule":[{"beginTime":"09:00","endTime":"17:00","acu":4}]}
	ScheduledRule *string `json:"ScheduledRule,omitempty" xml:"ScheduledRule,omitempty"`
	// The reserved parameter. This parameter is not supported.
	//
	// example:
	//
	// null
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which you want to add the new instance. Instances in the same security group can communicate with each other. The maximum number of instances allowed in a security group varies based on the type of the security group. For more information, see the "Security group limits" section in [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// >  The network type of the instance is determined by the security group specified by the SecurityGroupId parameter. For example, if the network type of the specified security group is VPC, the instance is a VPC-type instance. In this case, you must specify the VSwitchId parameter.
	//
	// example:
	//
	// sg-uf6av412xaxixuezol6w
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotStrategy    *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SupportCase     *string `json:"SupportCase,omitempty" xml:"SupportCase,omitempty"`
	// The specification of the system disk.
	SystemDiskShrink *string                           `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	Tag              []*RunRCInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserData         *string                           `json:"UserData,omitempty" xml:"UserData,omitempty"`
	UserDataInBase64 *bool                             `json:"UserDataInBase64,omitempty" xml:"UserDataInBase64,omitempty"`
	// The vSwitch ID of the instance. You must specify this parameter when you create an instance of the virtual private cloud (VPC) type. The specified vSwitch and security group must belong to the same VPC.
	//
	// >  If you specify the VSwitchId parameter, the zone specified by the ZoneId parameter must be the same as the zone in which the specified vSwitch resides. You can leave the ZoneId parameter empty. In this case, the system uses the zone in which the specified vSwitch resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2vcd61ngm890sk****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instance. You can call the DescribeZones operation to query the zone IDs.
	//
	// >  If you specify the VSwitchId parameter, the zone specified by the ZoneId parameter must be the same as the zone in which the specified vSwitch resides. You can leave the ZoneId parameter empty. In this case, the system uses the zone in which the specified vSwitch resides.
	//
	// example:
	//
	// cn-beijing-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RunRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunRCInstancesShrinkRequest) GetAcuType() *string {
	return s.AcuType
}

func (s *RunRCInstancesShrinkRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *RunRCInstancesShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RunRCInstancesShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RunRCInstancesShrinkRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *RunRCInstancesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunRCInstancesShrinkRequest) GetCreateAckEdgeParamShrink() *string {
	return s.CreateAckEdgeParamShrink
}

func (s *RunRCInstancesShrinkRequest) GetCreateExtraParam() *string {
	return s.CreateExtraParam
}

func (s *RunRCInstancesShrinkRequest) GetCreateMode() *string {
	return s.CreateMode
}

func (s *RunRCInstancesShrinkRequest) GetDataDiskShrink() *string {
	return s.DataDiskShrink
}

func (s *RunRCInstancesShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunRCInstancesShrinkRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RunRCInstancesShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *RunRCInstancesShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RunRCInstancesShrinkRequest) GetHostName() *string {
	return s.HostName
}

func (s *RunRCInstancesShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RunRCInstancesShrinkRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *RunRCInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RunRCInstancesShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunRCInstancesShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *RunRCInstancesShrinkRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *RunRCInstancesShrinkRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *RunRCInstancesShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *RunRCInstancesShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *RunRCInstancesShrinkRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *RunRCInstancesShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RunRCInstancesShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RunRCInstancesShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *RunRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunRCInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunRCInstancesShrinkRequest) GetScheduledRule() *string {
	return s.ScheduledRule
}

func (s *RunRCInstancesShrinkRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *RunRCInstancesShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RunRCInstancesShrinkRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *RunRCInstancesShrinkRequest) GetSupportCase() *string {
	return s.SupportCase
}

func (s *RunRCInstancesShrinkRequest) GetSystemDiskShrink() *string {
	return s.SystemDiskShrink
}

func (s *RunRCInstancesShrinkRequest) GetTag() []*RunRCInstancesShrinkRequestTag {
	return s.Tag
}

func (s *RunRCInstancesShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *RunRCInstancesShrinkRequest) GetUserDataInBase64() *bool {
	return s.UserDataInBase64
}

func (s *RunRCInstancesShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunRCInstancesShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RunRCInstancesShrinkRequest) SetAcuType(v string) *RunRCInstancesShrinkRequest {
	s.AcuType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAmount(v int32) *RunRCInstancesShrinkRequest {
	s.Amount = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAutoPay(v bool) *RunRCInstancesShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAutoRenew(v bool) *RunRCInstancesShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetAutoUseCoupon(v bool) *RunRCInstancesShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetClientToken(v string) *RunRCInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetCreateAckEdgeParamShrink(v string) *RunRCInstancesShrinkRequest {
	s.CreateAckEdgeParamShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetCreateExtraParam(v string) *RunRCInstancesShrinkRequest {
	s.CreateExtraParam = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetCreateMode(v string) *RunRCInstancesShrinkRequest {
	s.CreateMode = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDataDiskShrink(v string) *RunRCInstancesShrinkRequest {
	s.DataDiskShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDeletionProtection(v bool) *RunRCInstancesShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDeploymentSetId(v string) *RunRCInstancesShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDescription(v string) *RunRCInstancesShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetDryRun(v bool) *RunRCInstancesShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetHostName(v string) *RunRCInstancesShrinkRequest {
	s.HostName = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetImageId(v string) *RunRCInstancesShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInstanceChargeType(v string) *RunRCInstancesShrinkRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInstanceName(v string) *RunRCInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInstanceType(v string) *RunRCInstancesShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInternetChargeType(v string) *RunRCInstancesShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetInternetMaxBandwidthOut(v int32) *RunRCInstancesShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetIoOptimized(v string) *RunRCInstancesShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetKeyPairName(v string) *RunRCInstancesShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPassword(v string) *RunRCInstancesShrinkRequest {
	s.Password = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPasswordInherit(v bool) *RunRCInstancesShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPeriod(v int32) *RunRCInstancesShrinkRequest {
	s.Period = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPeriodUnit(v string) *RunRCInstancesShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetPromotionCode(v string) *RunRCInstancesShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetRegionId(v string) *RunRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetResourceGroupId(v string) *RunRCInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetScheduledRule(v string) *RunRCInstancesShrinkRequest {
	s.ScheduledRule = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSecurityEnhancementStrategy(v string) *RunRCInstancesShrinkRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSecurityGroupId(v string) *RunRCInstancesShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSpotStrategy(v string) *RunRCInstancesShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSupportCase(v string) *RunRCInstancesShrinkRequest {
	s.SupportCase = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetSystemDiskShrink(v string) *RunRCInstancesShrinkRequest {
	s.SystemDiskShrink = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetTag(v []*RunRCInstancesShrinkRequestTag) *RunRCInstancesShrinkRequest {
	s.Tag = v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetUserData(v string) *RunRCInstancesShrinkRequest {
	s.UserData = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetUserDataInBase64(v bool) *RunRCInstancesShrinkRequest {
	s.UserDataInBase64 = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetVSwitchId(v string) *RunRCInstancesShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) SetZoneId(v string) *RunRCInstancesShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *RunRCInstancesShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunRCInstancesShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunRCInstancesShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunRCInstancesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *RunRCInstancesShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunRCInstancesShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunRCInstancesShrinkRequestTag) SetKey(v string) *RunRCInstancesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *RunRCInstancesShrinkRequestTag) SetValue(v string) *RunRCInstancesShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *RunRCInstancesShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
