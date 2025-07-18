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
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*GetRegistrationPolicyResponseBodyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupAll
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
	// 99
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 47363C2B-1AAA-5954-8847-0E50FCC54117
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type GetRegistrationPolicyResponseBodyLimitDetail struct {
	// example:
	//
	// Personal
	DeviceBelong *string                                                 `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *GetRegistrationPolicyResponseBodyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type GetRegistrationPolicyResponseBodyLimitDetailLimitCount struct {
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 2
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
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
