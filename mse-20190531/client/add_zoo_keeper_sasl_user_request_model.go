// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZooKeeperSaslUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddZooKeeperSaslUserRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *AddZooKeeperSaslUserRequest
	GetInstanceId() *string
	SetReload(v bool) *AddZooKeeperSaslUserRequest
	GetReload() *bool
	SetSaslUser(v []*AddZooKeeperSaslUserRequestSaslUser) *AddZooKeeperSaslUserRequest
	GetSaslUser() []*AddZooKeeperSaslUserRequestSaslUser
}

type AddZooKeeperSaslUserRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Reload *bool `json:"Reload,omitempty" xml:"Reload,omitempty"`
	// This parameter is required.
	SaslUser []*AddZooKeeperSaslUserRequestSaslUser `json:"SaslUser,omitempty" xml:"SaslUser,omitempty" type:"Repeated"`
}

func (s AddZooKeeperSaslUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddZooKeeperSaslUserRequest) GoString() string {
	return s.String()
}

func (s *AddZooKeeperSaslUserRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddZooKeeperSaslUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddZooKeeperSaslUserRequest) GetReload() *bool {
	return s.Reload
}

func (s *AddZooKeeperSaslUserRequest) GetSaslUser() []*AddZooKeeperSaslUserRequestSaslUser {
	return s.SaslUser
}

func (s *AddZooKeeperSaslUserRequest) SetAcceptLanguage(v string) *AddZooKeeperSaslUserRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddZooKeeperSaslUserRequest) SetInstanceId(v string) *AddZooKeeperSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AddZooKeeperSaslUserRequest) SetReload(v bool) *AddZooKeeperSaslUserRequest {
	s.Reload = &v
	return s
}

func (s *AddZooKeeperSaslUserRequest) SetSaslUser(v []*AddZooKeeperSaslUserRequestSaslUser) *AddZooKeeperSaslUserRequest {
	s.SaslUser = v
	return s
}

func (s *AddZooKeeperSaslUserRequest) Validate() error {
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

type AddZooKeeperSaslUserRequestSaslUser struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// traaqwrasdf
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AddZooKeeperSaslUserRequestSaslUser) String() string {
	return dara.Prettify(s)
}

func (s AddZooKeeperSaslUserRequestSaslUser) GoString() string {
	return s.String()
}

func (s *AddZooKeeperSaslUserRequestSaslUser) GetDescription() *string {
	return s.Description
}

func (s *AddZooKeeperSaslUserRequestSaslUser) GetPassword() *string {
	return s.Password
}

func (s *AddZooKeeperSaslUserRequestSaslUser) GetUserName() *string {
	return s.UserName
}

func (s *AddZooKeeperSaslUserRequestSaslUser) SetDescription(v string) *AddZooKeeperSaslUserRequestSaslUser {
	s.Description = &v
	return s
}

func (s *AddZooKeeperSaslUserRequestSaslUser) SetPassword(v string) *AddZooKeeperSaslUserRequestSaslUser {
	s.Password = &v
	return s
}

func (s *AddZooKeeperSaslUserRequestSaslUser) SetUserName(v string) *AddZooKeeperSaslUserRequestSaslUser {
	s.UserName = &v
	return s
}

func (s *AddZooKeeperSaslUserRequestSaslUser) Validate() error {
	return dara.Validate(s)
}
