// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdpConfigId(v string) *GetClientUserRequest
	GetIdpConfigId() *string
	SetUsername(v string) *GetClientUserRequest
	GetUsername() *string
}

type GetClientUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 598
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// This parameter is required.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientUserRequest) GoString() string {
	return s.String()
}

func (s *GetClientUserRequest) GetIdpConfigId() *string {
	return s.IdpConfigId
}

func (s *GetClientUserRequest) GetUsername() *string {
	return s.Username
}

func (s *GetClientUserRequest) SetIdpConfigId(v string) *GetClientUserRequest {
	s.IdpConfigId = &v
	return s
}

func (s *GetClientUserRequest) SetUsername(v string) *GetClientUserRequest {
	s.Username = &v
	return s
}

func (s *GetClientUserRequest) Validate() error {
	return dara.Validate(s)
}
