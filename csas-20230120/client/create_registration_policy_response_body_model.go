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
	Policy *CreateRegistrationPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateRegistrationPolicyResponseBodyPolicy struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*CreateRegistrationPolicyResponseBodyPolicyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type CreateRegistrationPolicyResponseBodyPolicyLimitDetail struct {
	// example:
	//
	// Company
	DeviceBelong *string                                                          `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount struct {
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 3
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
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
