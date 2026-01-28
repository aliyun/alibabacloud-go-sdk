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
	ValidityPeriod *UpdateAuthorizationRuleUserAttachmentRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
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
