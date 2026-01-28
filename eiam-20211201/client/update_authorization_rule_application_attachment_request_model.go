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
	// 应用 ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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
	// 有效周期，当validityPeriodType为custom有效。
	ValidityPeriod *UpdateAuthorizationRuleApplicationAttachmentRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
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
