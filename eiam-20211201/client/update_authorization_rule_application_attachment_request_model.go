// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleApplicationAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest
	GetApplicationId() *string
	SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest
	GetInstanceId() *string
	SetValidityPeriod(v *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) *UpdateAuthorizationRuleApplicationAttachmentRequest
	GetValidityPeriod() *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod
	SetValidityType(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest
	GetValidityType() *string
}

type UpdateAuthorizationRuleApplicationAttachmentRequest struct {
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
	// A client token to ensure the idempotence of the request. Generate a unique value from your client for this parameter. The ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
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
	// The time range of the validity period. This parameter takes effect only when **ValidityType*	- is set to **time_bound**.
	ValidityPeriod *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The validity type of the relationship. Valid values:
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

func (s UpdateAuthorizationRuleApplicationAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleApplicationAttachmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) GetValidityPeriod() *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod {
	return s.ValidityPeriod
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) GetValidityType() *string {
	return s.ValidityType
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) SetApplicationId(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) SetClientToken(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) SetInstanceId(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) SetValidityPeriod(v *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) *UpdateAuthorizationRuleApplicationAttachmentRequest {
	s.ValidityPeriod = v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) SetValidityType(v string) *UpdateAuthorizationRuleApplicationAttachmentRequest {
	s.ValidityType = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequest) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod struct {
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

func (s UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) SetEndTime(v int64) *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) SetStartTime(v int64) *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod) Validate() error {
	return dara.Validate(s)
}
