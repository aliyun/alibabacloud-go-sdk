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
	// Specifies whether the switch takes effect by default. Valid values:
	//
	// - **add**: Takes effect by default for newly added servers.
	//
	// - **del**: Does not take effect by default for newly added servers.
	//
	// example:
	//
	// add
	TargetDefault *string `json:"TargetDefault,omitempty" xml:"TargetDefault,omitempty"`
	// The general switch type.
	//
	// >You can obtain this parameter from the SwitchId response parameter of the [ListClientUserDefineRules](~~ListClientUserDefineRules~~) and [ListSystemClientRules](~~ListSystemClientRules~~) operations.
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
