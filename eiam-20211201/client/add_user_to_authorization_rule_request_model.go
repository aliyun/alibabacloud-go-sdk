// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *AddUserToAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *AddUserToAuthorizationRuleRequest
	GetClientToken() *string
	SetInstanceId(v string) *AddUserToAuthorizationRuleRequest
	GetInstanceId() *string
	SetUserId(v string) *AddUserToAuthorizationRuleRequest
	GetUserId() *string
	SetValidityPeriod(v *AddUserToAuthorizationRuleRequestValidityPeriod) *AddUserToAuthorizationRuleRequest
	GetValidityPeriod() *AddUserToAuthorizationRuleRequestValidityPeriod
	SetValidityType(v string) *AddUserToAuthorizationRuleRequest
	GetValidityType() *string
}

type AddUserToAuthorizationRuleRequest struct {
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
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 账户ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 有效周期，当validityPeriodType为custom有效。
	ValidityPeriod *AddUserToAuthorizationRuleRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
	//
	// This parameter is required.
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s AddUserToAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *AddUserToAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *AddUserToAuthorizationRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddUserToAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddUserToAuthorizationRuleRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddUserToAuthorizationRuleRequest) GetValidityPeriod() *AddUserToAuthorizationRuleRequestValidityPeriod {
	return s.ValidityPeriod
}

func (s *AddUserToAuthorizationRuleRequest) GetValidityType() *string {
	return s.ValidityType
}

func (s *AddUserToAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *AddUserToAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *AddUserToAuthorizationRuleRequest) SetClientToken(v string) *AddUserToAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *AddUserToAuthorizationRuleRequest) SetInstanceId(v string) *AddUserToAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUserToAuthorizationRuleRequest) SetUserId(v string) *AddUserToAuthorizationRuleRequest {
	s.UserId = &v
	return s
}

func (s *AddUserToAuthorizationRuleRequest) SetValidityPeriod(v *AddUserToAuthorizationRuleRequestValidityPeriod) *AddUserToAuthorizationRuleRequest {
	s.ValidityPeriod = v
	return s
}

func (s *AddUserToAuthorizationRuleRequest) SetValidityType(v string) *AddUserToAuthorizationRuleRequest {
	s.ValidityType = &v
	return s
}

func (s *AddUserToAuthorizationRuleRequest) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddUserToAuthorizationRuleRequestValidityPeriod struct {
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

func (s AddUserToAuthorizationRuleRequestValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s AddUserToAuthorizationRuleRequestValidityPeriod) GoString() string {
	return s.String()
}

func (s *AddUserToAuthorizationRuleRequestValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AddUserToAuthorizationRuleRequestValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AddUserToAuthorizationRuleRequestValidityPeriod) SetEndTime(v int64) *AddUserToAuthorizationRuleRequestValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *AddUserToAuthorizationRuleRequestValidityPeriod) SetStartTime(v int64) *AddUserToAuthorizationRuleRequestValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *AddUserToAuthorizationRuleRequestValidityPeriod) Validate() error {
	return dara.Validate(s)
}
