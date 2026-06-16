// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApplicationToAuthorizationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AddApplicationToAuthorizationRuleRequest
	GetApplicationId() *string
	SetAuthorizationRuleId(v string) *AddApplicationToAuthorizationRuleRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *AddApplicationToAuthorizationRuleRequest
	GetClientToken() *string
	SetInstanceId(v string) *AddApplicationToAuthorizationRuleRequest
	GetInstanceId() *string
	SetValidityPeriod(v *AddApplicationToAuthorizationRuleRequestValidityPeriod) *AddApplicationToAuthorizationRuleRequest
	GetValidityPeriod() *AddApplicationToAuthorizationRuleRequestValidityPeriod
	SetValidityType(v string) *AddApplicationToAuthorizationRuleRequest
	GetValidityType() *string
}

type AddApplicationToAuthorizationRuleRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// A client token used to ensure the idempotence of the request. Generate a value from your client to make sure that the value is unique among different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
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
	// The time range of the validity period. This parameter is valid only when you set **ValidityType*	- to **time_bound**.
	ValidityPeriod *AddApplicationToAuthorizationRuleRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The validity period type. Valid values:
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

func (s AddApplicationToAuthorizationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddApplicationToAuthorizationRuleRequest) GoString() string {
	return s.String()
}

func (s *AddApplicationToAuthorizationRuleRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AddApplicationToAuthorizationRuleRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *AddApplicationToAuthorizationRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddApplicationToAuthorizationRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddApplicationToAuthorizationRuleRequest) GetValidityPeriod() *AddApplicationToAuthorizationRuleRequestValidityPeriod {
	return s.ValidityPeriod
}

func (s *AddApplicationToAuthorizationRuleRequest) GetValidityType() *string {
	return s.ValidityType
}

func (s *AddApplicationToAuthorizationRuleRequest) SetApplicationId(v string) *AddApplicationToAuthorizationRuleRequest {
	s.ApplicationId = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequest) SetAuthorizationRuleId(v string) *AddApplicationToAuthorizationRuleRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequest) SetClientToken(v string) *AddApplicationToAuthorizationRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequest) SetInstanceId(v string) *AddApplicationToAuthorizationRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequest) SetValidityPeriod(v *AddApplicationToAuthorizationRuleRequestValidityPeriod) *AddApplicationToAuthorizationRuleRequest {
	s.ValidityPeriod = v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequest) SetValidityType(v string) *AddApplicationToAuthorizationRuleRequest {
	s.ValidityType = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequest) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddApplicationToAuthorizationRuleRequestValidityPeriod struct {
	// The end time of the validity period. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1704062061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the validity period. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1704042061000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s AddApplicationToAuthorizationRuleRequestValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s AddApplicationToAuthorizationRuleRequestValidityPeriod) GoString() string {
	return s.String()
}

func (s *AddApplicationToAuthorizationRuleRequestValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *AddApplicationToAuthorizationRuleRequestValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *AddApplicationToAuthorizationRuleRequestValidityPeriod) SetEndTime(v int64) *AddApplicationToAuthorizationRuleRequestValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequestValidityPeriod) SetStartTime(v int64) *AddApplicationToAuthorizationRuleRequestValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *AddApplicationToAuthorizationRuleRequestValidityPeriod) Validate() error {
	return dara.Validate(s)
}
