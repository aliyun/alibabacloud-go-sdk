// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZooKeeperSaslUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateZooKeeperSaslUserRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *UpdateZooKeeperSaslUserRequest
	GetInstanceId() *string
	SetReload(v bool) *UpdateZooKeeperSaslUserRequest
	GetReload() *bool
	SetSaslUser(v []*UpdateZooKeeperSaslUserRequestSaslUser) *UpdateZooKeeperSaslUserRequest
	GetSaslUser() []*UpdateZooKeeperSaslUserRequestSaslUser
}

type UpdateZooKeeperSaslUserRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl327w****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Reload *bool `json:"Reload,omitempty" xml:"Reload,omitempty"`
	// This parameter is required.
	SaslUser []*UpdateZooKeeperSaslUserRequestSaslUser `json:"SaslUser,omitempty" xml:"SaslUser,omitempty" type:"Repeated"`
}

func (s UpdateZooKeeperSaslUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateZooKeeperSaslUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateZooKeeperSaslUserRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateZooKeeperSaslUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateZooKeeperSaslUserRequest) GetReload() *bool {
	return s.Reload
}

func (s *UpdateZooKeeperSaslUserRequest) GetSaslUser() []*UpdateZooKeeperSaslUserRequestSaslUser {
	return s.SaslUser
}

func (s *UpdateZooKeeperSaslUserRequest) SetAcceptLanguage(v string) *UpdateZooKeeperSaslUserRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateZooKeeperSaslUserRequest) SetInstanceId(v string) *UpdateZooKeeperSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateZooKeeperSaslUserRequest) SetReload(v bool) *UpdateZooKeeperSaslUserRequest {
	s.Reload = &v
	return s
}

func (s *UpdateZooKeeperSaslUserRequest) SetSaslUser(v []*UpdateZooKeeperSaslUserRequestSaslUser) *UpdateZooKeeperSaslUserRequest {
	s.SaslUser = v
	return s
}

func (s *UpdateZooKeeperSaslUserRequest) Validate() error {
	if s.SaslUser != nil {
		for _, item := range s.SaslUser {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateZooKeeperSaslUserRequestSaslUser struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// origin-password
	OriginPassword *string `json:"OriginPassword,omitempty" xml:"OriginPassword,omitempty"`
	// example:
	//
	// xxxxxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateZooKeeperSaslUserRequestSaslUser) String() string {
	return dara.Prettify(s)
}

func (s UpdateZooKeeperSaslUserRequestSaslUser) GoString() string {
	return s.String()
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) GetDescription() *string {
	return s.Description
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) GetOriginPassword() *string {
	return s.OriginPassword
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) GetPassword() *string {
	return s.Password
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) GetUserName() *string {
	return s.UserName
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) SetDescription(v string) *UpdateZooKeeperSaslUserRequestSaslUser {
	s.Description = &v
	return s
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) SetOriginPassword(v string) *UpdateZooKeeperSaslUserRequestSaslUser {
	s.OriginPassword = &v
	return s
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) SetPassword(v string) *UpdateZooKeeperSaslUserRequestSaslUser {
	s.Password = &v
	return s
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) SetUserName(v string) *UpdateZooKeeperSaslUserRequestSaslUser {
	s.UserName = &v
	return s
}

func (s *UpdateZooKeeperSaslUserRequestSaslUser) Validate() error {
	return dara.Validate(s)
}
