// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathWhitelistId(v string) *GetAttackPathWhitelistRequest
	GetAttackPathWhitelistId() *string
}

type GetAttackPathWhitelistRequest struct {
	// Attack path whitelist ID.
	//
	// > You can call [ListAttackPathWhitelist](~~ListAttackPathWhitelist~~) to query the attack path whitelist ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// apwl-b33dec0acf9b42aabde032d656c0****
	AttackPathWhitelistId *string `json:"AttackPathWhitelistId,omitempty" xml:"AttackPathWhitelistId,omitempty"`
}

func (s GetAttackPathWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathWhitelistRequest) GoString() string {
	return s.String()
}

func (s *GetAttackPathWhitelistRequest) GetAttackPathWhitelistId() *string {
	return s.AttackPathWhitelistId
}

func (s *GetAttackPathWhitelistRequest) SetAttackPathWhitelistId(v string) *GetAttackPathWhitelistRequest {
	s.AttackPathWhitelistId = &v
	return s
}

func (s *GetAttackPathWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
