// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistrationPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v []*ListRegistrationPoliciesResponseBodyPolicies) *ListRegistrationPoliciesResponseBody
	GetPolicies() []*ListRegistrationPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListRegistrationPoliciesResponseBody
	GetRequestId() *string
	SetTotalNum(v string) *ListRegistrationPoliciesResponseBody
	GetTotalNum() *string
}

type ListRegistrationPoliciesResponseBody struct {
	Policies []*ListRegistrationPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// 7A8FE38A-E29C-5678-B84A-FEDBCB83552F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListRegistrationPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBody) GetPolicies() []*ListRegistrationPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListRegistrationPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegistrationPoliciesResponseBody) GetTotalNum() *string {
	return s.TotalNum
}

func (s *ListRegistrationPoliciesResponseBody) SetPolicies(v []*ListRegistrationPoliciesResponseBodyPolicies) *ListRegistrationPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListRegistrationPoliciesResponseBody) SetRequestId(v string) *ListRegistrationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBody) SetTotalNum(v string) *ListRegistrationPoliciesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRegistrationPoliciesResponseBodyPolicies struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*ListRegistrationPoliciesResponseBodyPoliciesLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
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
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetLimitDetail() []*ListRegistrationPoliciesResponseBodyPoliciesLimitDetail {
	return s.LimitDetail
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetName() *string {
	return s.Name
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetPriority() *int64 {
	return s.Priority
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) GetWhitelist() []*string {
	return s.Whitelist
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetCreateTime(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetDescription(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetLimitDetail(v []*ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) *ListRegistrationPoliciesResponseBodyPolicies {
	s.LimitDetail = v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetMatchMode(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.MatchMode = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetName(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetPolicyId(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetPriority(v int64) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Priority = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetStatus(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Status = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetUserGroupIds(v []*string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.UserGroupIds = v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetWhitelist(v []*string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Whitelist = v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}

type ListRegistrationPoliciesResponseBodyPoliciesLimitDetail struct {
	// example:
	//
	// Company
	DeviceBelong *string                                                            `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitAll
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) GetDeviceBelong() *string {
	return s.DeviceBelong
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) GetLimitCount() *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount {
	return s.LimitCount
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) GetLimitType() *string {
	return s.LimitType
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) SetDeviceBelong(v string) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) SetLimitCount(v *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail {
	s.LimitCount = v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) SetLimitType(v string) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail {
	s.LimitType = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) Validate() error {
	return dara.Validate(s)
}

type ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount struct {
	// example:
	//
	// 3
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) GetAll() *int32 {
	return s.All
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) GetMobile() *int32 {
	return s.Mobile
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) GetPC() *int32 {
	return s.PC
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) SetAll(v int32) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) SetMobile(v int32) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) SetPC(v int32) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount {
	s.PC = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) Validate() error {
	return dara.Validate(s)
}
