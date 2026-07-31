// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticityAssuranceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *ModifyElasticityAssuranceRequestPrivatePoolOptions) *ModifyElasticityAssuranceRequest
	GetPrivatePoolOptions() *ModifyElasticityAssuranceRequestPrivatePoolOptions
	SetClientToken(v string) *ModifyElasticityAssuranceRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyElasticityAssuranceRequest
	GetDescription() *string
	SetInstanceAmount(v int32) *ModifyElasticityAssuranceRequest
	GetInstanceAmount() *int32
	SetOwnerAccount(v string) *ModifyElasticityAssuranceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyElasticityAssuranceRequest
	GetOwnerId() *int64
	SetRecurrenceRules(v []*ModifyElasticityAssuranceRequestRecurrenceRules) *ModifyElasticityAssuranceRequest
	GetRecurrenceRules() []*ModifyElasticityAssuranceRequestRecurrenceRules
	SetRegionId(v string) *ModifyElasticityAssuranceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyElasticityAssuranceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyElasticityAssuranceRequest
	GetResourceOwnerId() *int64
}

type ModifyElasticityAssuranceRequest struct {
	PrivatePoolOptions *ModifyElasticityAssuranceRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the elasticity assurance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total number of instances to be reserved by the elasticity assurance. Valid values: number of used instances to 1000. This parameter cannot be modified together with other parameters.
	//
	// example:
	//
	// 10
	InstanceAmount *int32  `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of recurrence rules for the time-sharing elasticity assurance.
	RecurrenceRules []*ModifyElasticityAssuranceRequestRecurrenceRules `json:"RecurrenceRules,omitempty" xml:"RecurrenceRules,omitempty" type:"Repeated"`
	// The region ID of the elasticity assurance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyElasticityAssuranceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceRequest) GetPrivatePoolOptions() *ModifyElasticityAssuranceRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyElasticityAssuranceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyElasticityAssuranceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyElasticityAssuranceRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *ModifyElasticityAssuranceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyElasticityAssuranceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyElasticityAssuranceRequest) GetRecurrenceRules() []*ModifyElasticityAssuranceRequestRecurrenceRules {
	return s.RecurrenceRules
}

func (s *ModifyElasticityAssuranceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyElasticityAssuranceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyElasticityAssuranceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyElasticityAssuranceRequest) SetPrivatePoolOptions(v *ModifyElasticityAssuranceRequestPrivatePoolOptions) *ModifyElasticityAssuranceRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetClientToken(v string) *ModifyElasticityAssuranceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetDescription(v string) *ModifyElasticityAssuranceRequest {
	s.Description = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetInstanceAmount(v int32) *ModifyElasticityAssuranceRequest {
	s.InstanceAmount = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetOwnerAccount(v string) *ModifyElasticityAssuranceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetOwnerId(v int64) *ModifyElasticityAssuranceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetRecurrenceRules(v []*ModifyElasticityAssuranceRequestRecurrenceRules) *ModifyElasticityAssuranceRequest {
	s.RecurrenceRules = v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetRegionId(v string) *ModifyElasticityAssuranceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetResourceOwnerAccount(v string) *ModifyElasticityAssuranceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) SetResourceOwnerId(v int64) *ModifyElasticityAssuranceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyElasticityAssuranceRequest) Validate() error {
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
	return nil
}

type ModifyElasticityAssuranceRequestPrivatePoolOptions struct {
	// The ID of the elasticity assurance.
	//
	// This parameter is required.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the elasticity assurance. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with http:// or https://. The name can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// eapTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyElasticityAssuranceRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ModifyElasticityAssuranceRequestPrivatePoolOptions) GetName() *string {
	return s.Name
}

func (s *ModifyElasticityAssuranceRequestPrivatePoolOptions) SetId(v string) *ModifyElasticityAssuranceRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyElasticityAssuranceRequestPrivatePoolOptions) SetName(v string) *ModifyElasticityAssuranceRequestPrivatePoolOptions {
	s.Name = &v
	return s
}

func (s *ModifyElasticityAssuranceRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyElasticityAssuranceRequestRecurrenceRules struct {
	// The end time of the time-sharing assurance. The value must be on the hour.
	//
	// example:
	//
	// 10
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The policy type of the recurrence rule. Valid values:
	//
	// - Daily: repeats on a daily basis.
	//
	// - Weekly: repeats on a weekly basis.
	//
	// - Monthly: repeats on a monthly basis.
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
	// 5
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The effective period start time of the time-sharing assurance. The value must be on the hour.
	//
	// > You must specify both `StartHour` and `EndHour`, and the difference between them must be at least 4 hours.
	//
	// example:
	//
	// 4
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s ModifyElasticityAssuranceRequestRecurrenceRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticityAssuranceRequestRecurrenceRules) GoString() string {
	return s.String()
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) GetEndHour() *int32 {
	return s.EndHour
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) GetStartHour() *int32 {
	return s.StartHour
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) SetEndHour(v int32) *ModifyElasticityAssuranceRequestRecurrenceRules {
	s.EndHour = &v
	return s
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) SetRecurrenceType(v string) *ModifyElasticityAssuranceRequestRecurrenceRules {
	s.RecurrenceType = &v
	return s
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) SetRecurrenceValue(v string) *ModifyElasticityAssuranceRequestRecurrenceRules {
	s.RecurrenceValue = &v
	return s
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) SetStartHour(v int32) *ModifyElasticityAssuranceRequestRecurrenceRules {
	s.StartHour = &v
	return s
}

func (s *ModifyElasticityAssuranceRequestRecurrenceRules) Validate() error {
	return dara.Validate(s)
}
