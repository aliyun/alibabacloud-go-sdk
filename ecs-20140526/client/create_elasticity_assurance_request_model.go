// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticityAssuranceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *CreateElasticityAssuranceRequestPrivatePoolOptions) *CreateElasticityAssuranceRequest
	GetPrivatePoolOptions() *CreateElasticityAssuranceRequestPrivatePoolOptions
	SetAssuranceTimes(v string) *CreateElasticityAssuranceRequest
	GetAssuranceTimes() *string
	SetAutoRenew(v bool) *CreateElasticityAssuranceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateElasticityAssuranceRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateElasticityAssuranceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateElasticityAssuranceRequest
	GetDescription() *string
	SetInstanceAmount(v int32) *CreateElasticityAssuranceRequest
	GetInstanceAmount() *int32
	SetInstanceCpuCoreCount(v int32) *CreateElasticityAssuranceRequest
	GetInstanceCpuCoreCount() *int32
	SetInstanceType(v []*string) *CreateElasticityAssuranceRequest
	GetInstanceType() []*string
	SetOwnerAccount(v string) *CreateElasticityAssuranceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateElasticityAssuranceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateElasticityAssuranceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateElasticityAssuranceRequest
	GetPeriodUnit() *string
	SetRecurrenceRules(v []*CreateElasticityAssuranceRequestRecurrenceRules) *CreateElasticityAssuranceRequest
	GetRecurrenceRules() []*CreateElasticityAssuranceRequestRecurrenceRules
	SetRegionId(v string) *CreateElasticityAssuranceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateElasticityAssuranceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateElasticityAssuranceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateElasticityAssuranceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *CreateElasticityAssuranceRequest
	GetStartTime() *string
	SetTag(v []*CreateElasticityAssuranceRequestTag) *CreateElasticityAssuranceRequest
	GetTag() []*CreateElasticityAssuranceRequestTag
	SetZoneId(v []*string) *CreateElasticityAssuranceRequest
	GetZoneId() []*string
}

type CreateElasticityAssuranceRequest struct {
	PrivatePoolOptions *CreateElasticityAssuranceRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The total number of times that the elasticity assurance can be applied. Valid values: Unlimited. Currently, only the unlimited mode is supported within the service validity period.
	//
	// Default value: Unlimited.
	//
	// example:
	//
	// Unlimited
	AssuranceTimes *string `json:"AssuranceTimes,omitempty" xml:"AssuranceTimes,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - true: Auto-renewal is enabled.
	//
	// - false: Auto-renewal is disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Unit: months. Valid values: 1, 2, 3, 6, 12, 24, and 36.
	//
	//
	//
	// - If `PeriodUnit=Month`, the default value is 1.
	//
	// - If `PeriodUnit=Year`, the default value is 12.
	//
	//
	// > This parameter is required when `AutoRenew` is set to `True`.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The `ClientToken` value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the elasticity assurance service. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total number of instances to be reserved for a single instance type.
	//
	// Valid values: 1 to 1000.
	//
	// 	Notice: This parameter is required.
	//
	// example:
	//
	// 2
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	InstanceCpuCoreCount *int32 `json:"InstanceCpuCoreCount,omitempty" xml:"InstanceCpuCoreCount,omitempty"`
	// The instance type. Currently, you can configure an elasticity assurance service for only one instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.c6.xlarge
	InstanceType []*string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The purchase duration. The unit of the duration is determined by the `PeriodUnit` parameter. Valid values:
	//
	// - If `PeriodUnit` is set to `Month`: 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// - If `PeriodUnit` is set to `Year`: 1, 2, 3, 4, and 5.
	//
	// - If `PeriodUnit` is set to `Day`: 1 to 365.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the purchase duration. Valid values:
	//
	// - Month: month.
	//
	// - Year: year.
	//
	// - Day: day.
	//
	//   > When `PeriodUnit` is set to `Day`, you must also specify RecurrenceRules to create a time-sharing elasticity assurance.
	//
	// Default value: Year.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The recurrence rules for the time-sharing elasticity assurance.
	RecurrenceRules []*CreateElasticityAssuranceRequestRecurrenceRules `json:"RecurrenceRules,omitempty" xml:"RecurrenceRules,omitempty" type:"Repeated"`
	// The region ID of the elasticity assurance service. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the elasticity assurance service belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The effective period start time of the elasticity assurance service. By default, the service takes effect when the operation is invoked. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2020-10-30T06:32:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tags of the elasticity assurance service.
	Tag []*CreateElasticityAssuranceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID within the region of the elasticity assurance service. Currently, you can create an elasticity assurance service in only one zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s CreateElasticityAssuranceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequest) GetPrivatePoolOptions() *CreateElasticityAssuranceRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateElasticityAssuranceRequest) GetAssuranceTimes() *string {
	return s.AssuranceTimes
}

func (s *CreateElasticityAssuranceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateElasticityAssuranceRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateElasticityAssuranceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateElasticityAssuranceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateElasticityAssuranceRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *CreateElasticityAssuranceRequest) GetInstanceCpuCoreCount() *int32 {
	return s.InstanceCpuCoreCount
}

func (s *CreateElasticityAssuranceRequest) GetInstanceType() []*string {
	return s.InstanceType
}

func (s *CreateElasticityAssuranceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateElasticityAssuranceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateElasticityAssuranceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateElasticityAssuranceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateElasticityAssuranceRequest) GetRecurrenceRules() []*CreateElasticityAssuranceRequestRecurrenceRules {
	return s.RecurrenceRules
}

func (s *CreateElasticityAssuranceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateElasticityAssuranceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateElasticityAssuranceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateElasticityAssuranceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateElasticityAssuranceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateElasticityAssuranceRequest) GetTag() []*CreateElasticityAssuranceRequestTag {
	return s.Tag
}

func (s *CreateElasticityAssuranceRequest) GetZoneId() []*string {
	return s.ZoneId
}

func (s *CreateElasticityAssuranceRequest) SetPrivatePoolOptions(v *CreateElasticityAssuranceRequestPrivatePoolOptions) *CreateElasticityAssuranceRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetAssuranceTimes(v string) *CreateElasticityAssuranceRequest {
	s.AssuranceTimes = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetAutoRenew(v bool) *CreateElasticityAssuranceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetAutoRenewPeriod(v int32) *CreateElasticityAssuranceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetClientToken(v string) *CreateElasticityAssuranceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetDescription(v string) *CreateElasticityAssuranceRequest {
	s.Description = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetInstanceAmount(v int32) *CreateElasticityAssuranceRequest {
	s.InstanceAmount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetInstanceCpuCoreCount(v int32) *CreateElasticityAssuranceRequest {
	s.InstanceCpuCoreCount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetInstanceType(v []*string) *CreateElasticityAssuranceRequest {
	s.InstanceType = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetOwnerAccount(v string) *CreateElasticityAssuranceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetOwnerId(v int64) *CreateElasticityAssuranceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetPeriod(v int32) *CreateElasticityAssuranceRequest {
	s.Period = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetPeriodUnit(v string) *CreateElasticityAssuranceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetRecurrenceRules(v []*CreateElasticityAssuranceRequestRecurrenceRules) *CreateElasticityAssuranceRequest {
	s.RecurrenceRules = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetRegionId(v string) *CreateElasticityAssuranceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetResourceGroupId(v string) *CreateElasticityAssuranceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetResourceOwnerAccount(v string) *CreateElasticityAssuranceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetResourceOwnerId(v int64) *CreateElasticityAssuranceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetStartTime(v string) *CreateElasticityAssuranceRequest {
	s.StartTime = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetTag(v []*CreateElasticityAssuranceRequestTag) *CreateElasticityAssuranceRequest {
	s.Tag = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetZoneId(v []*string) *CreateElasticityAssuranceRequest {
	s.ZoneId = v
	return s
}

func (s *CreateElasticityAssuranceRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.RecurrenceRules != nil {
		for _, item := range s.RecurrenceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateElasticityAssuranceRequestPrivatePoolOptions struct {
	// The match mode of the elasticity assurance service. Valid values:
	//
	// - Open: open mode. The system automatically matches the capacity of open private pools when instances are started. If no matching private pool capacity is available, public pool resources are used to start the instances.
	//
	// - Target: targeted mode. Instances are started by using the capacity of the specified private pool. If the specified private pool capacity is unavailable, the instances fail to start.
	//
	// Default value: Open.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	// The name of the elasticity assurance service. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// eapTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateElasticityAssuranceRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) GetName() *string {
	return s.Name
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateElasticityAssuranceRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) SetName(v string) *CreateElasticityAssuranceRequestPrivatePoolOptions {
	s.Name = &v
	return s
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateElasticityAssuranceRequestRecurrenceRules struct {
	// The end time of the time-sharing assurance. The value must be on the hour.
	//
	// example:
	//
	// 10
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The type of the recurrence rule. Valid values:
	//
	// - Daily: daily recurrence.
	//
	// - Weekly: weekly recurrence.
	//
	// - Monthly: monthly recurrence.
	//
	// > You must specify both `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The value of the recurrence rule.
	//
	// - If `RecurrenceType` is set to `Daily`, you can specify only one value. Valid values: 1 to 31. The value specifies the interval in days between recurrences.
	//
	// - If `RecurrenceType` is set to `Weekly`, you can specify multiple values separated by commas (,). The values for Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday are 0, 1, 2, 3, 4, 5, and 6. For example, `1,2` specifies Monday and Tuesday.
	//
	// - If `RecurrenceType` is set to `Monthly`, the format is `A-B`. Valid values of A and B: 1 to 31. B must be greater than or equal to A. For example, `1-5` specifies the 1st to 5th day of each month.
	//
	// > You must specify both `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The effective period start hour of the time-sharing assurance. The value must be on the hour.
	//
	// > You must specify both `StartHour` and `EndHour`, and the difference between them must be at least 4 hours.
	//
	// example:
	//
	// 4
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s CreateElasticityAssuranceRequestRecurrenceRules) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequestRecurrenceRules) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetEndHour() *int32 {
	return s.EndHour
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetStartHour() *int32 {
	return s.StartHour
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetEndHour(v int32) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.EndHour = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetRecurrenceType(v string) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.RecurrenceType = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetRecurrenceValue(v string) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.RecurrenceValue = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetStartHour(v int32) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.StartHour = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) Validate() error {
	return dara.Validate(s)
}

type CreateElasticityAssuranceRequestTag struct {
	// The tag key of the elasticity assurance service. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the elasticity assurance service. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateElasticityAssuranceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateElasticityAssuranceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateElasticityAssuranceRequestTag) SetKey(v string) *CreateElasticityAssuranceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateElasticityAssuranceRequestTag) SetValue(v string) *CreateElasticityAssuranceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateElasticityAssuranceRequestTag) Validate() error {
	return dara.Validate(s)
}
