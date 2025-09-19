// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaliciousFileWhitelistConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventName(v string) *CreateMaliciousFileWhitelistConfigRequest
	GetEventName() *string
	SetField(v string) *CreateMaliciousFileWhitelistConfigRequest
	GetField() *string
	SetFieldValue(v string) *CreateMaliciousFileWhitelistConfigRequest
	GetFieldValue() *string
	SetOperator(v string) *CreateMaliciousFileWhitelistConfigRequest
	GetOperator() *string
	SetSource(v string) *CreateMaliciousFileWhitelistConfigRequest
	GetSource() *string
	SetTargetType(v string) *CreateMaliciousFileWhitelistConfigRequest
	GetTargetType() *string
	SetTargetValue(v string) *CreateMaliciousFileWhitelistConfigRequest
	GetTargetValue() *string
}

type CreateMaliciousFileWhitelistConfigRequest struct {
	// The name of the alert.
	//
	// 	- Set the value to ALL, which indicates all alert types.
	//
	// example:
	//
	// ALL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The field that you want to use in the whitelist rule.
	//
	// example:
	//
	// fileMd5
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The value of the field that you want to use in the whitelist rule.
	//
	// example:
	//
	// b2cf9747ee49d8d9b105cf16e078cc16
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The logical operator that you want to use in the whitelist rule.
	//
	// 	- Set the value to strEqual, which indicates the equality operator (=).
	//
	// example:
	//
	// strEqual
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The feature to which this operation belongs.
	//
	// 	- Set the value to agentless, which indicates the agentless detection feature.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the assets on which you want the whitelist rule to take effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- SELECTION_KEY: selected assets
	//
	// example:
	//
	// ALL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The assets on which you want the whitelist rule to take effect. Valid values:
	//
	// 	- ALL: all assets
	//
	// 	- Others: selected assets
	//
	// example:
	//
	// ALL
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s CreateMaliciousFileWhitelistConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMaliciousFileWhitelistConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateMaliciousFileWhitelistConfigRequest) GetEventName() *string {
	return s.EventName
}

func (s *CreateMaliciousFileWhitelistConfigRequest) GetField() *string {
	return s.Field
}

func (s *CreateMaliciousFileWhitelistConfigRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateMaliciousFileWhitelistConfigRequest) GetOperator() *string {
	return s.Operator
}

func (s *CreateMaliciousFileWhitelistConfigRequest) GetSource() *string {
	return s.Source
}

func (s *CreateMaliciousFileWhitelistConfigRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateMaliciousFileWhitelistConfigRequest) GetTargetValue() *string {
	return s.TargetValue
}

func (s *CreateMaliciousFileWhitelistConfigRequest) SetEventName(v string) *CreateMaliciousFileWhitelistConfigRequest {
	s.EventName = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigRequest) SetField(v string) *CreateMaliciousFileWhitelistConfigRequest {
	s.Field = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigRequest) SetFieldValue(v string) *CreateMaliciousFileWhitelistConfigRequest {
	s.FieldValue = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigRequest) SetOperator(v string) *CreateMaliciousFileWhitelistConfigRequest {
	s.Operator = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigRequest) SetSource(v string) *CreateMaliciousFileWhitelistConfigRequest {
	s.Source = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigRequest) SetTargetType(v string) *CreateMaliciousFileWhitelistConfigRequest {
	s.TargetType = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigRequest) SetTargetValue(v string) *CreateMaliciousFileWhitelistConfigRequest {
	s.TargetValue = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigRequest) Validate() error {
	return dara.Validate(s)
}
