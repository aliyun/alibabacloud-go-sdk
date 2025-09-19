// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommonSwitchConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetDefault(v string) *UpdateCommonSwitchConfigRequest
	GetTargetDefault() *string
	SetType(v string) *UpdateCommonSwitchConfigRequest
	GetType() *string
}

type UpdateCommonSwitchConfigRequest struct {
	// Specifies whether to turn on the switch for newly added servers. Valid values:
	//
	// 	- **add**: yes
	//
	// 	- **del**: no
	//
	// example:
	//
	// add
	TargetDefault *string `json:"TargetDefault,omitempty" xml:"TargetDefault,omitempty"`
	// The type of the switch.
	//
	// >  You can call the [ListClientUserDefineRules](~~ListClientUserDefineRules~~) or [ListSystemClientRules](~~ListSystemClientRules~~) operation to obtain the type from the response parameter SwitchId.
	//
	// example:
	//
	// USER-DEFINE-RULE-SWITCH-TYPE_190****
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateCommonSwitchConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommonSwitchConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommonSwitchConfigRequest) GetTargetDefault() *string {
	return s.TargetDefault
}

func (s *UpdateCommonSwitchConfigRequest) GetType() *string {
	return s.Type
}

func (s *UpdateCommonSwitchConfigRequest) SetTargetDefault(v string) *UpdateCommonSwitchConfigRequest {
	s.TargetDefault = &v
	return s
}

func (s *UpdateCommonSwitchConfigRequest) SetType(v string) *UpdateCommonSwitchConfigRequest {
	s.Type = &v
	return s
}

func (s *UpdateCommonSwitchConfigRequest) Validate() error {
	return dara.Validate(s)
}
