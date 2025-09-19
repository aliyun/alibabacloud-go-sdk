// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackPathWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathWhitelistId(v string) *DeleteAttackPathWhitelistRequest
	GetAttackPathWhitelistId() *string
}

type DeleteAttackPathWhitelistRequest struct {
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

func (s DeleteAttackPathWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackPathWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteAttackPathWhitelistRequest) GetAttackPathWhitelistId() *string {
	return s.AttackPathWhitelistId
}

func (s *DeleteAttackPathWhitelistRequest) SetAttackPathWhitelistId(v string) *DeleteAttackPathWhitelistRequest {
	s.AttackPathWhitelistId = &v
	return s
}

func (s *DeleteAttackPathWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
