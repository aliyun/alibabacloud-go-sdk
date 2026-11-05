// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonSwitchConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *GetCommonSwitchConfigRequest
	GetType() *string
}

type GetCommonSwitchConfigRequest struct {
	// The general switch type.
	//
	// >You can obtain this parameter from the SwitchId response parameter of the [ListClientUserDefineRules](~~ListClientUserDefineRules~~) and [ListSystemClientRules](~~ListSystemClientRules~~) operations.
	//
	// example:
	//
	// USER-DEFINE-RULE-SWITCH-TYPE_180****
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCommonSwitchConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommonSwitchConfigRequest) GoString() string {
	return s.String()
}

func (s *GetCommonSwitchConfigRequest) GetType() *string {
	return s.Type
}

func (s *GetCommonSwitchConfigRequest) SetType(v string) *GetCommonSwitchConfigRequest {
	s.Type = &v
	return s
}

func (s *GetCommonSwitchConfigRequest) Validate() error {
	return dara.Validate(s)
}
