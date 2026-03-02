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
	PasskeyId         *string `json:"PasskeyId,omitempty" xml:"PasskeyId,omitempty"`
	PasskeyName       *string `json:"PasskeyName,omitempty" xml:"PasskeyName,omitempty"`
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
