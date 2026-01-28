// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGroupToAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *AddGroupToAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *AddGroupToAuthorizationRuleRequest
	GetClientToken() *string
	SetGroupId(v string) *AddGroupToAuthorizationRuleRequest
	GetGroupId() *string
	SetInstanceId(v string) *AddGroupToAuthorizationRuleRequest
	GetInstanceId() *string
	SetValidityPeriod(v *AddGroupToAuthorizationRuleRequestValidityPeriod) *AddGroupToAuthorizationRuleRequest
	GetValidityPeriod() *AddGroupToAuthorizationRuleRequestValidityPeriod
	SetValidityType(v string) *AddGroupToAuthorizationRuleRequest
	GetValidityType() *string
}

type AddGroupToAuthorizationRuleRequest struct {
	// 授权规则标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 组ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 有效周期，当validityPeriodType为custom有效。
	ValidityPeriod *AddGroupToAuthorizationRuleRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
	//
	// This parameter is required.
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s AddGroupToAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGroupToAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *AddGroupToAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *AddGroupToAuthorizationRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddGroupToAuthorizationRuleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddGroupToAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddGroupToAuthorizationRuleRequest) GetValidityPeriod() *AddGroupToAuthorizationRuleRequestValidityPeriod {
	return s.ValidityPeriod
}

func (s *AddGroupToAuthorizationRuleRequest) GetValidityType() *string {
	return s.ValidityType
}

func (s *AddGroupToAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *AddGroupToAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *AddGroupToAuthorizationRuleRequest) SetClientToken(v string) *AddGroupToAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *AddGroupToAuthorizationRuleRequest) SetGroupId(v string) *AddGroupToAuthorizationRuleRequest {
	s.GroupId = &v
	return s
}

func (s *AddGroupToAuthorizationRuleRequest) SetInstanceId(v string) *AddGroupToAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *AddGroupToAuthorizationRuleRequest) SetValidityPeriod(v *AddGroupToAuthorizationRuleRequestValidityPeriod) *AddGroupToAuthorizationRuleRequest {
	s.ValidityPeriod = v
	return s
}

func (s *AddGroupToAuthorizationRuleRequest) SetValidityType(v string) *AddGroupToAuthorizationRuleRequest {
	s.ValidityType = &v
	return s
}

func (s *AddGroupToAuthorizationRuleRequest) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddGroupToAuthorizationRuleRequestValidityPeriod struct {
	// 授权规则生效结束时间，采用unix纪元精确到毫秒。
	//
	// example:
	//
	// 1704062061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 授权规则生效开始时间，采用unix纪元精确到毫秒。
	//
	// example:
	//
	// 1704042061000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s AddGroupToAuthorizationRuleRequestValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s AddGroupToAuthorizationRuleRequestValidityPeriod) GoString() string {
	return s.String()
}

func (s *AddGroupToAuthorizationRuleRequestValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AddGroupToAuthorizationRuleRequestValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AddGroupToAuthorizationRuleRequestValidityPeriod) SetEndTime(v int64) *AddGroupToAuthorizationRuleRequestValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *AddGroupToAuthorizationRuleRequestValidityPeriod) SetStartTime(v int64) *AddGroupToAuthorizationRuleRequestValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *AddGroupToAuthorizationRuleRequestValidityPeriod) Validate() error {
	return dara.Validate(s)
}
