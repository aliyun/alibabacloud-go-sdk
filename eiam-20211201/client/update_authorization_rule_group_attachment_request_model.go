// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleGroupAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleGroupAttachmentRequest
	GetAuthorizationRuleId() *string
	SetClientToken(v string) *UpdateAuthorizationRuleGroupAttachmentRequest
	GetClientToken() *string
	SetGroupId(v string) *UpdateAuthorizationRuleGroupAttachmentRequest
	GetGroupId() *string
	SetInstanceId(v string) *UpdateAuthorizationRuleGroupAttachmentRequest
	GetInstanceId() *string
	SetValidityPeriod(v *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) *UpdateAuthorizationRuleGroupAttachmentRequest
	GetValidityPeriod() *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod
	SetValidityType(v string) *UpdateAuthorizationRuleGroupAttachmentRequest
	GetValidityType() *string
}

type UpdateAuthorizationRuleGroupAttachmentRequest struct {
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
	// 组标识。
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
	ValidityPeriod *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
	//
	// This parameter is required.
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s UpdateAuthorizationRuleGroupAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleGroupAttachmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) GetValidityPeriod() *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod {
	return s.ValidityPeriod
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) GetValidityType() *string {
	return s.ValidityType
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) SetAuthorizationRuleId(v string) *UpdateAuthorizationRuleGroupAttachmentRequest {
	s.AuthorizationRuleId = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) SetClientToken(v string) *UpdateAuthorizationRuleGroupAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) SetGroupId(v string) *UpdateAuthorizationRuleGroupAttachmentRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) SetInstanceId(v string) *UpdateAuthorizationRuleGroupAttachmentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) SetValidityPeriod(v *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) *UpdateAuthorizationRuleGroupAttachmentRequest {
	s.ValidityPeriod = v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) SetValidityType(v string) *UpdateAuthorizationRuleGroupAttachmentRequest {
	s.ValidityType = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequest) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod struct {
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

func (s UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) SetEndTime(v int64) *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) SetStartTime(v int64) *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *UpdateAuthorizationRuleGroupAttachmentRequestValidityPeriod) Validate() error {
	return dara.Validate(s)
}
