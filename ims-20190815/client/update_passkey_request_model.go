// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePasskeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPasskeyId(v string) *UpdatePasskeyRequest
	GetPasskeyId() *string
	SetPasskeyName(v string) *UpdatePasskeyRequest
	GetPasskeyName() *string
	SetUserPrincipalName(v string) *UpdatePasskeyRequest
	GetUserPrincipalName() *string
}

type UpdatePasskeyRequest struct {
	// The ID of the passkey.
	//
	// example:
	//
	// PASSKEY-CuZjEHhWcr7GIQOMGvkS
	PasskeyId *string `json:"PasskeyId,omitempty" xml:"PasskeyId,omitempty"`
	// The name of the passkey.
	//
	// example:
	//
	// device1
	PasskeyName *string `json:"PasskeyName,omitempty" xml:"PasskeyName,omitempty"`
	// The logon name of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdatePasskeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePasskeyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePasskeyRequest) GetPasskeyId() *string {
	return s.PasskeyId
}

func (s *UpdatePasskeyRequest) GetPasskeyName() *string {
	return s.PasskeyName
}

func (s *UpdatePasskeyRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *UpdatePasskeyRequest) SetPasskeyId(v string) *UpdatePasskeyRequest {
	s.PasskeyId = &v
	return s
}

func (s *UpdatePasskeyRequest) SetPasskeyName(v string) *UpdatePasskeyRequest {
	s.PasskeyName = &v
	return s
}

func (s *UpdatePasskeyRequest) SetUserPrincipalName(v string) *UpdatePasskeyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdatePasskeyRequest) Validate() error {
	return dara.Validate(s)
}
