// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaliciousFileWhitelistConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateMaliciousFileWhitelistConfigRequest
	GetConfigId() *int64
	SetEventName(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetEventName() *string
	SetField(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetField() *string
	SetFieldValue(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetFieldValue() *string
	SetOperator(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetOperator() *string
	SetRemark(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetRemark() *string
	SetSource(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetSource() *string
	SetTargetType(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetTargetType() *string
	SetTargetValue(v string) *UpdateMaliciousFileWhitelistConfigRequest
	GetTargetValue() *string
}

type UpdateMaliciousFileWhitelistConfigRequest struct {
	// The rule ID. This parameter is optional. If you do not specify this parameter, a whitelist rule is created.
	//
	// example:
	//
	// 1
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The alerting name. Valid values:
	//
	// - ALL: all Alarm Metric.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The field to be whitelisted.
	//
	// example:
	//
	// fileMd5
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The value of the field to be whitelisted.
	//
	// example:
	//
	// b2cf9747ee49d8d9b105cf16e078cc16
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The operator used for rule matching. Valid values:
	//
	// - strEqual: string equals.
	//
	// example:
	//
	// strEqual
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The business source. Valid values:
	//
	// - agentless: agentless detection.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the target scope. Valid values:
	//
	// - ALL: all assets
	//
	// - SELECTION_KEY: assets selected by using the asset selection component.
	//
	// example:
	//
	// ALL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The target scope. Valid values:
	//
	// - ALL: all assets
	//
	// - Other values: the key of the asset scope selected by using the asset selection component.
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s UpdateMaliciousFileWhitelistConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaliciousFileWhitelistConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetEventName() *string {
	return s.EventName
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetField() *string {
	return s.Field
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetOperator() *string {
	return s.Operator
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetSource() *string {
	return s.Source
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetConfigId(v int64) *UpdateMaliciousFileWhitelistConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetEventName(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.EventName = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetField(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.Field = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetFieldValue(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.FieldValue = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetOperator(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.Operator = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetRemark(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.Remark = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetSource(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.Source = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetTargetType(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.TargetType = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) SetTargetValue(v string) *UpdateMaliciousFileWhitelistConfigRequest {
	s.TargetValue = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigRequest) Validate() error {
	return dara.Validate(s)
}
