// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetRegistrationPolicyResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetRegistrationPolicyResponseBody
	GetDescription() *string
	SetLimitDetail(v []*GetRegistrationPolicyResponseBodyLimitDetail) *GetRegistrationPolicyResponseBody
	GetLimitDetail() []*GetRegistrationPolicyResponseBodyLimitDetail
	SetMatchMode(v string) *GetRegistrationPolicyResponseBody
	GetMatchMode() *string
	SetName(v string) *GetRegistrationPolicyResponseBody
	GetName() *string
	SetPolicyId(v string) *GetRegistrationPolicyResponseBody
	GetPolicyId() *string
	SetPriority(v int64) *GetRegistrationPolicyResponseBody
	GetPriority() *int64
	SetRequestId(v string) *GetRegistrationPolicyResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetRegistrationPolicyResponseBody
	GetStatus() *string
	SetUserGroupIds(v []*string) *GetRegistrationPolicyResponseBody
	GetUserGroupIds() []*string
	SetWhitelist(v []*string) *GetRegistrationPolicyResponseBody
	GetWhitelist() []*string
}

type GetRegistrationPolicyResponseBody struct {
	// The time when the device registration policy was created. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the device registration policy.
	//
	// example:
	//
	// This is a device registration policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of limit details for the device registration policy.
	LimitDetail []*GetRegistrationPolicyResponseBodyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// The target type for policy matching. Valid values:
	//
	// - **UserGroupAll**: All users.
	//
	// - **UserGroupNormal**: Specific user groups.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The policy name of the device registration policy.
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
	// The policy priority of the device registration policy. A value of 0 indicates the highest priority, and a value of 99 indicates the lowest priority.
	//
	// example:
	//
	// 99
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 47363C2B-1AAA-5954-8847-0E50FCC54117
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The IDs of the user groups associated with the device registration policy. This field has a value only when the target type for policy matching is **UserGroupNormal**.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The list of whitelisted users for the device registration policy.
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s GetRegistrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetRegistrationPolicyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetRegistrationPolicyResponseBody) GetLimitDetail() []*GetRegistrationPolicyResponseBodyLimitDetail {
	return s.LimitDetail
}

func (s *GetRegistrationPolicyResponseBody) GetMatchMode() *string {
	return s.MatchMode
}

func (s *GetRegistrationPolicyResponseBody) GetName() *string {
	return s.Name
}

func (s *GetRegistrationPolicyResponseBody) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetRegistrationPolicyResponseBody) GetPriority() *int64 {
	return s.Priority
}

func (s *GetRegistrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegistrationPolicyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetRegistrationPolicyResponseBody) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *GetRegistrationPolicyResponseBody) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *GetRegistrationPolicyResponseBody) SetCreateTime(v string) *GetRegistrationPolicyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetDescription(v string) *GetRegistrationPolicyResponseBody {
	s.Description = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetLimitDetail(v []*GetRegistrationPolicyResponseBodyLimitDetail) *GetRegistrationPolicyResponseBody {
	s.LimitDetail = v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetMatchMode(v string) *GetRegistrationPolicyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetName(v string) *GetRegistrationPolicyResponseBody {
	s.Name = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetPolicyId(v string) *GetRegistrationPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetPriority(v int64) *GetRegistrationPolicyResponseBody {
	s.Priority = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetRequestId(v string) *GetRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetStatus(v string) *GetRegistrationPolicyResponseBody {
	s.Status = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetUserGroupIds(v []*string) *GetRegistrationPolicyResponseBody {
	s.UserGroupIds = v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetWhitelist(v []*string) *GetRegistrationPolicyResponseBody {
	s.Whitelist = v
	return s
}

func (s *GetRegistrationPolicyResponseBody) Validate() error {
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

type GetRegistrationPolicyResponseBodyLimitDetail struct {
	// The device ownership. Valid values:
	//
	// - **Company**: Corporate device.
	//
	// - **Personal**: Personal device.
	//
	// example:
	//
	// Personal
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// The limit on the number of device registrations.
	LimitCount *GetRegistrationPolicyResponseBodyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// The type of device registration limit. Valid values:
	//
	// - **Unlimited**: Unlimited.
	//
	// - **LimitAll**: Limited by total count.
	//
	// - **LimitDiff**: Limited by endpoint category.
	//
	// example:
	//
	// LimitDiff
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s GetRegistrationPolicyResponseBodyLimitDetail) String() string {
	return dara.Prettify(s)
}

func (s GetRegistrationPolicyResponseBodyLimitDetail) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) GetLimitCount() *GetRegistrationPolicyResponseBodyLimitDetailLimitCount {
	return s.LimitCount
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) GetLimitType() *string {
	return s.LimitType
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) SetDeviceBelong(v string) *GetRegistrationPolicyResponseBodyLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) SetLimitCount(v *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) *GetRegistrationPolicyResponseBodyLimitDetail {
	s.LimitCount = v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) SetLimitType(v string) *GetRegistrationPolicyResponseBodyLimitDetail {
	s.LimitType = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) Validate() error {
	if s.LimitCount != nil {
		if err := s.LimitCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRegistrationPolicyResponseBodyLimitDetailLimitCount struct {
	// The total limit on the number of device registrations. This field is valid only when the type of device registration limit is **LimitAll**.
	//
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// The limit on the number of mobile logons for device registration. This field is valid only when the type of device registration limit is **LimitDiff**.
	//
	// example:
	//
	// 2
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The limit on the number of PC logons for device registration. This field is valid only when the type of device registration limit is **LimitDiff**.
	//
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s GetRegistrationPolicyResponseBodyLimitDetailLimitCount) String() string {
	return dara.Prettify(s)
}

func (s GetRegistrationPolicyResponseBodyLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) GetAll() *int32 {
	return s.All
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) SetAll(v int32) *GetRegistrationPolicyResponseBodyLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) SetMobile(v int32) *GetRegistrationPolicyResponseBodyLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) SetPC(v int32) *GetRegistrationPolicyResponseBodyLimitDetailLimitCount {
	s.PC = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) Validate() error {
	return dara.Validate(s)
}
