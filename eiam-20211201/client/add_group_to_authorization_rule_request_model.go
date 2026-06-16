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
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// A client token that you generate to ensure the idempotence of the request. Make sure that the value of this parameter is unique across different requests. The client token can contain only ASCII characters and must be no more than 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/en/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time range of the validity period. This parameter takes effect only when **ValidityType*	- is set to **time_bound**.
	ValidityPeriod *AddGroupToAuthorizationRuleRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The type of the validity period. Valid values:
	//
	// - permanent: The relationship is permanent.
	//
	// - time_bound: The relationship is valid for a custom time range.
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
	// The end time of the validity period. This is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1704062061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the validity period. This is a UNIX timestamp. Unit: milliseconds.
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
