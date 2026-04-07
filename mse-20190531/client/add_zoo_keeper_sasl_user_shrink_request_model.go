// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZooKeeperSaslUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddZooKeeperSaslUserShrinkRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *AddZooKeeperSaslUserShrinkRequest
	GetInstanceId() *string
	SetReload(v bool) *AddZooKeeperSaslUserShrinkRequest
	GetReload() *bool
	SetSaslUserShrink(v string) *AddZooKeeperSaslUserShrinkRequest
	GetSaslUserShrink() *string
}

type AddZooKeeperSaslUserShrinkRequest struct {
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
	SaslUserShrink *string `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
}

func (s AddZooKeeperSaslUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddZooKeeperSaslUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddZooKeeperSaslUserShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddZooKeeperSaslUserShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddZooKeeperSaslUserShrinkRequest) GetReload() *bool {
	return s.Reload
}

func (s *AddZooKeeperSaslUserShrinkRequest) GetSaslUserShrink() *string {
	return s.SaslUserShrink
}

func (s *AddZooKeeperSaslUserShrinkRequest) SetAcceptLanguage(v string) *AddZooKeeperSaslUserShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddZooKeeperSaslUserShrinkRequest) SetInstanceId(v string) *AddZooKeeperSaslUserShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AddZooKeeperSaslUserShrinkRequest) SetReload(v bool) *AddZooKeeperSaslUserShrinkRequest {
	s.Reload = &v
	return s
}

func (s *AddZooKeeperSaslUserShrinkRequest) SetSaslUserShrink(v string) *AddZooKeeperSaslUserShrinkRequest {
	s.SaslUserShrink = &v
	return s
}

func (s *AddZooKeeperSaslUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
