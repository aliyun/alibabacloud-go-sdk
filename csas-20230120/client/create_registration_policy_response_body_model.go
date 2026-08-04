// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *CreateRegistrationPolicyResponseBodyPolicy) *CreateRegistrationPolicyResponseBody
	GetPolicy() *CreateRegistrationPolicyResponseBodyPolicy
	SetRequestId(v string) *CreateRegistrationPolicyResponseBody
	GetRequestId() *string
}

type CreateRegistrationPolicyResponseBody struct {
	// The device registration policy.
	Policy *CreateRegistrationPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FEF1144C-95D1-5F7C-81EF-9DB70EA49FCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRegistrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBody) GetPolicy() *CreateRegistrationPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *CreateRegistrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRegistrationPolicyResponseBody) SetPolicy(v *CreateRegistrationPolicyResponseBodyPolicy) *CreateRegistrationPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *CreateRegistrationPolicyResponseBody) SetRequestId(v string) *CreateRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRegistrationPolicyResponseBodyPolicy struct {
	// The time when the device registration policy was created.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// A description of the device registration policy.
	//
	// example:
	//
	// 这是一条设备注册策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of device registration limits.
	LimitDetail []*CreateRegistrationPolicyResponseBodyPolicyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// The target type for policy matching. Valid values:
	//
	// - **UserGroupAll**: Apply to all users.
	//
	// - **UserGroupNormal**: Apply to selected user groups.
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
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The priority of the device registration policy. A value of 0 indicates the highest priority. A value of 99 indicates the lowest priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status of the device registration policy. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of user groups to which the device registration policy applies. This field has a value only when MatchMode is set to **UserGroupNormal**.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of usernames in the whitelist for the device registration policy.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateRegistrationPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetLimitDetail() []*CreateRegistrationPolicyResponseBodyPolicyLimitDetail {
	return s.LimitDetail
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetMatchMode() *string {
	return s.MatchMode
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetName() *string {
	return s.Name
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetPriority() *string {
	return s.Priority
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetStatus() *string {
	return s.Status
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetCreateTime(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetDescription(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetLimitDetail(v []*CreateRegistrationPolicyResponseBodyPolicyLimitDetail) *CreateRegistrationPolicyResponseBodyPolicy {
	s.LimitDetail = v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetMatchMode(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.MatchMode = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetName(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetPolicyId(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetPriority(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetStatus(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Status = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetWhitelist(v []*string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Whitelist = v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) Validate() error {
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

type CreateRegistrationPolicyResponseBodyPolicyLimitDetail struct {
	// The ownership of the device. Valid values:
	//
	// - **Company**: Company device.
	//
	// - **Personal**: Personal device.
	//
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// The restriction count for device registration.
	LimitCount *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// The restriction type for device registration. Valid values:
	//
	// - **Unlimited**: No restrictions.
	//
	// - **LimitAll**: Limit by total count.
	//
	// - **LimitDiff**: Limit by device category.
	//
	// example:
	//
	// LimitDiff
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetail) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) GetLimitCount() *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	return s.LimitCount
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) GetLimitType() *string {
	return s.LimitType
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) SetDeviceBelong(v string) *CreateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitCount(v *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) *CreateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitCount = v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitType(v string) *CreateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitType = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) Validate() error {
	if s.LimitCount != nil {
		if err := s.LimitCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount struct {
	// The total restriction count for device registration. This parameter takes effect only when LimitType is set to **LimitAll**.
	//
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// The restriction count for mobile logins. This parameter takes effect only when LimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 3
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The restriction count for PC logins. This parameter takes effect only when LimitType is set to **LimitDiff**.
	//
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GetAll() *int32 {
	return s.All
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetAll(v int32) *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetMobile(v int32) *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetPC(v int32) *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.PC = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) Validate() error {
	return dara.Validate(s)
}
