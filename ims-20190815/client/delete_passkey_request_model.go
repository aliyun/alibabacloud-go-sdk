// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePasskeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPasskeyId(v string) *DeletePasskeyRequest
	GetPasskeyId() *string
	SetUserPrincipalName(v string) *DeletePasskeyRequest
	GetUserPrincipalName() *string
}

type DeletePasskeyRequest struct {
	// The ID of the passkey.
	//
	// example:
	//
	// PASSKEY-CuZjEHhWcr7GIQOMGvkS
	PasskeyId *string `json:"PasskeyId,omitempty" xml:"PasskeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeletePasskeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePasskeyRequest) GoString() string {
	return s.String()
}

func (s *DeletePasskeyRequest) GetPasskeyId() *string {
	return s.PasskeyId
}

func (s *DeletePasskeyRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DeletePasskeyRequest) SetPasskeyId(v string) *DeletePasskeyRequest {
	s.PasskeyId = &v
	return s
}

func (s *DeletePasskeyRequest) SetUserPrincipalName(v string) *DeletePasskeyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *DeletePasskeyRequest) Validate() error {
	return dara.Validate(s)
}
