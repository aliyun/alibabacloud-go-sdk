// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *UpdateRegistrationPolicyResponseBodyPolicy) *UpdateRegistrationPolicyResponseBody
	GetPolicy() *UpdateRegistrationPolicyResponseBodyPolicy
	SetRequestId(v string) *UpdateRegistrationPolicyResponseBody
	GetRequestId() *string
}

type UpdateRegistrationPolicyResponseBody struct {
	// The device registration policy.
	Policy *UpdateRegistrationPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 27064ECA-0936-59F3-8A98-EC821E5BD08F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRegistrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBody) GetPolicy() *UpdateRegistrationPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *UpdateRegistrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRegistrationPolicyResponseBody) SetPolicy(v *UpdateRegistrationPolicyResponseBodyPolicy) *UpdateRegistrationPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBody) SetRequestId(v string) *UpdateRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRegistrationPolicyResponseBodyPolicy struct {
	// The time when the device registration policy was created.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the device registration policy.
	//
	// example:
	//
	// 这是一条设备注册策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of limit details for the device registration policy.
	LimitDetail []*UpdateRegistrationPolicyResponseBodyPolicyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// The matching target type of the policy. Valid values:
	//
	// - **UserGroupAll**: Associates with all users.
	//
	// - **UserGroupNormal**: Associates with specific user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The name of the device registration policy.
	//
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the device registration policy.
	//
	// example:
	//
	// 设备注册策略ID
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The priority of the device registration policy. A smaller value indicates a higher priority. The value 0 indicates the highest priority, and 99 indicates the lowest priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the device registration policy. Valid values:
	//
	// - **Enabled**
	//
	// - **Disabled**
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of user group IDs for the device registration policy. This field has a value when the policy target type is **UserGroupNormal**.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of whitelisted users for the device registration policy.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateRegistrationPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetLimitDetail() []*UpdateRegistrationPolicyResponseBodyPolicyLimitDetail {
	return s.LimitDetail
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetMatchMode() *string {
	return s.MatchMode
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetName() *string {
	return s.Name
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetPriority() *string {
	return s.Priority
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetStatus() *string {
	return s.Status
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetCreateTime(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetDescription(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetLimitDetail(v []*UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.LimitDetail = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetMatchMode(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.MatchMode = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetName(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetPolicyId(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetPriority(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetStatus(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Status = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetWhitelist(v []*string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Whitelist = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) Validate() error {
	if s.LimitDetail != nil {
		for _, item := range s.LimitDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRegistrationPolicyResponseBodyPolicyLimitDetail struct {
	// The owner of the device. Valid values:
	//
	// - **Company**: Corporate device.
	//
	// - **Personal**: Personal device.
	//
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// The registration limit count.
	LimitCount *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// The registration limit type. Valid values:
	//
	// - **Unlimited**: No limit.
	//
	// - **LimitAll**: Limits the total number of devices.
	//
	// - **LimitDiff**: Limits devices by terminal type.
	//
	// example:
	//
	// LimitAll
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) GetLimitCount() *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	return s.LimitCount
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) SetDeviceBelong(v string) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitCount(v *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitCount = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitType(v string) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) Validate() error {
	if s.LimitCount != nil {
		if err := s.LimitCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount struct {
	// The total number of devices that can be registered. This parameter is valid only when LimitType is set to **LimitAll**.
	//
	// example:
	//
	// 1
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// The number of mobile devices that can be registered. This parameter is valid only when LimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The number of PCs that can be registered. This parameter is valid only when LimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GetAll() *int32 {
	return s.All
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetAll(v int32) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetMobile(v int32) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetPC(v int32) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.PC = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) Validate() error {
	return dara.Validate(s)
}
