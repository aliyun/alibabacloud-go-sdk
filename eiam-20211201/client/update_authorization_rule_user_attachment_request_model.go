// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleUserAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleUserAttachmentRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *UpdateAuthorizationRuleUserAttachmentRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateAuthorizationRuleUserAttachmentRequest
	GetInstanceId() *string
	SetUserId(v string) *UpdateAuthorizationRuleUserAttachmentRequest
	GetUserId() *string
	SetValidityPeriod(v *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) *UpdateAuthorizationRuleUserAttachmentRequest
	GetValidityPeriod() *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod
	SetValidityType(v string) *UpdateAuthorizationRuleUserAttachmentRequest
	GetValidityType() *string
}

type UpdateAuthorizationRuleUserAttachmentRequest struct {
	// The authorization rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// A client token used to ensure the idempotence of the request. Generate a unique value from your client for this parameter. ClientToken supports only ASCII characters and must be no more than 64 characters long. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
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
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The time range of the validity period. This parameter takes effect when ValidityType is set to time_bound.
	ValidityPeriod *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The validity period type of the association. Valid values:
	//
	// - permanent: The association is permanent.
	//
	// - time_bound: The association is valid for a custom time range.
	//
	// This parameter is required.
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s UpdateAuthorizationRuleUserAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleUserAttachmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) GetValidityPeriod() *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod {
	return s.ValidityPeriod
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) GetValidityType() *string {
	return s.ValidityType
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleUserAttachmentRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) SetClientToken(v string) *UpdateAuthorizationRuleUserAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) SetInstanceId(v string) *UpdateAuthorizationRuleUserAttachmentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) SetUserId(v string) *UpdateAuthorizationRuleUserAttachmentRequest {
	s.UserId = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) SetValidityPeriod(v *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) *UpdateAuthorizationRuleUserAttachmentRequest {
	s.ValidityPeriod = v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) SetValidityType(v string) *UpdateAuthorizationRuleUserAttachmentRequest {
	s.ValidityType = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequest) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod struct {
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

func (s UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) SetEndTime(v int64) *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) SetStartTime(v int64) *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod) Validate() error {
	return dara.Validate(s)
}
