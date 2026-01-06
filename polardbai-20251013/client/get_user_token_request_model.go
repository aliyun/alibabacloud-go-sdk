// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GetUserTokenRequest
	GetDBClusterId() *string
	SetPassword(v string) *GetUserTokenRequest
	GetPassword() *string
	SetUsername(v string) *GetUserTokenRequest
	GetUsername() *string
}

type GetUserTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUserTokenRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetUserTokenRequest) GetPassword() *string {
	return s.Password
}

func (s *GetUserTokenRequest) GetUsername() *string {
	return s.Username
}

func (s *GetUserTokenRequest) SetDBClusterId(v string) *GetUserTokenRequest {
	s.DBClusterId = &v
	return s
}

func (s *GetUserTokenRequest) SetPassword(v string) *GetUserTokenRequest {
	s.Password = &v
	return s
}

func (s *GetUserTokenRequest) SetUsername(v string) *GetUserTokenRequest {
	s.Username = &v
	return s
}

func (s *GetUserTokenRequest) Validate() error {
	return dara.Validate(s)
}
