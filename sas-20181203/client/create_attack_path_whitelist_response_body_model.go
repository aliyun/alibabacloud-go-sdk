// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttackPathWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackPathWhitelist(v *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist) *CreateAttackPathWhitelistResponseBody
	GetAttackPathWhitelist() *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist
	SetRequestId(v string) *CreateAttackPathWhitelistResponseBody
	GetRequestId() *string
}

type CreateAttackPathWhitelistResponseBody struct {
	// The attack path whitelist.
	AttackPathWhitelist *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist `json:"AttackPathWhitelist,omitempty" xml:"AttackPathWhitelist,omitempty" type:"Struct"`
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAttackPathWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAttackPathWhitelistResponseBody) GetAttackPathWhitelist() *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist {
	return s.AttackPathWhitelist
}

func (s *CreateAttackPathWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAttackPathWhitelistResponseBody) SetAttackPathWhitelist(v *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist) *CreateAttackPathWhitelistResponseBody {
	s.AttackPathWhitelist = v
	return s
}

func (s *CreateAttackPathWhitelistResponseBody) SetRequestId(v string) *CreateAttackPathWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAttackPathWhitelistResponseBody) Validate() error {
	if s.AttackPathWhitelist != nil {
		if err := s.AttackPathWhitelist.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAttackPathWhitelistResponseBodyAttackPathWhitelist struct {
	// The ID of the attack path whitelist.
	//
	// example:
	//
	// apwl-b33dec0acf9b42aabde032d656c0****
	AttackPathWhitelistId *string `json:"AttackPathWhitelistId,omitempty" xml:"AttackPathWhitelistId,omitempty"`
}

func (s CreateAttackPathWhitelistResponseBodyAttackPathWhitelist) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackPathWhitelistResponseBodyAttackPathWhitelist) GoString() string {
	return s.String()
}

func (s *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist) GetAttackPathWhitelistId() *string {
	return s.AttackPathWhitelistId
}

func (s *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist) SetAttackPathWhitelistId(v string) *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist {
	s.AttackPathWhitelistId = &v
	return s
}

func (s *CreateAttackPathWhitelistResponseBodyAttackPathWhitelist) Validate() error {
	return dara.Validate(s)
}
