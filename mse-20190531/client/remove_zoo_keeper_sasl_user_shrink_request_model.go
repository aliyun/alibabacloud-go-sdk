// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveZooKeeperSaslUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *RemoveZooKeeperSaslUserShrinkRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *RemoveZooKeeperSaslUserShrinkRequest
	GetInstanceId() *string
	SetReload(v bool) *RemoveZooKeeperSaslUserShrinkRequest
	GetReload() *bool
	SetUserNamesShrink(v string) *RemoveZooKeeperSaslUserShrinkRequest
	GetUserNamesShrink() *string
}

type RemoveZooKeeperSaslUserShrinkRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Reload *bool `json:"Reload,omitempty" xml:"Reload,omitempty"`
	// This parameter is required.
	UserNamesShrink *string `json:"UserNames,omitempty" xml:"UserNames,omitempty"`
}

func (s RemoveZooKeeperSaslUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveZooKeeperSaslUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) GetReload() *bool {
	return s.Reload
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) GetUserNamesShrink() *string {
	return s.UserNamesShrink
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) SetAcceptLanguage(v string) *RemoveZooKeeperSaslUserShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) SetInstanceId(v string) *RemoveZooKeeperSaslUserShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) SetReload(v bool) *RemoveZooKeeperSaslUserShrinkRequest {
	s.Reload = &v
	return s
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) SetUserNamesShrink(v string) *RemoveZooKeeperSaslUserShrinkRequest {
	s.UserNamesShrink = &v
	return s
}

func (s *RemoveZooKeeperSaslUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
