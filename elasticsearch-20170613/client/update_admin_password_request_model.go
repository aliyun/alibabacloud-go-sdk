// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdminPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEsAdminPassword(v string) *UpdateAdminPasswordRequest
	GetEsAdminPassword() *string
	SetClientToken(v string) *UpdateAdminPasswordRequest
	GetClientToken() *string
}

type UpdateAdminPasswordRequest struct {
	// example:
	//
	// es_password
	EsAdminPassword *string `json:"esAdminPassword,omitempty" xml:"esAdminPassword,omitempty"`
	// Indicates whether the password was updated. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAdminPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdminPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdminPasswordRequest) GetEsAdminPassword() *string {
	return s.EsAdminPassword
}

func (s *UpdateAdminPasswordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAdminPasswordRequest) SetEsAdminPassword(v string) *UpdateAdminPasswordRequest {
	s.EsAdminPassword = &v
	return s
}

func (s *UpdateAdminPasswordRequest) SetClientToken(v string) *UpdateAdminPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAdminPasswordRequest) Validate() error {
	return dara.Validate(s)
}
