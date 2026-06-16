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
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// A client token that you provide to ensure the idempotence of the request. Make sure that the client token is unique for each request. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The time range of the validity period. This parameter is valid only when **ValidityType*	- is set to **time_bound**.
	ValidityPeriod *AddUserToAuthorizationRuleRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The type of the validity period for the relationship. Valid values:
	//
	// - permanent: The authorization is permanent.
	//
	// - time_bound: The authorization is valid for a custom time range.
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
	// The end time of the validity period. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1704062061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the validity period. The value is a UNIX timestamp. Unit: milliseconds.
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
