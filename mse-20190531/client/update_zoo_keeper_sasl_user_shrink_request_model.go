// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZooKeeperSaslUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateZooKeeperSaslUserShrinkRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *UpdateZooKeeperSaslUserShrinkRequest
	GetInstanceId() *string
	SetReload(v bool) *UpdateZooKeeperSaslUserShrinkRequest
	GetReload() *bool
	SetSaslUserShrink(v string) *UpdateZooKeeperSaslUserShrinkRequest
	GetSaslUserShrink() *string
}

type UpdateZooKeeperSaslUserShrinkRequest struct {
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
	SaslUserShrink *string `json:"SaslUser,omitempty" xml:"SaslUser,omitempty"`
}

func (s UpdateZooKeeperSaslUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateZooKeeperSaslUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) GetReload() *bool {
	return s.Reload
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) GetSaslUserShrink() *string {
	return s.SaslUserShrink
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) SetAcceptLanguage(v string) *UpdateZooKeeperSaslUserShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) SetInstanceId(v string) *UpdateZooKeeperSaslUserShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) SetReload(v bool) *UpdateZooKeeperSaslUserShrinkRequest {
	s.Reload = &v
	return s
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) SetSaslUserShrink(v string) *UpdateZooKeeperSaslUserShrinkRequest {
	s.SaslUserShrink = &v
	return s
}

func (s *UpdateZooKeeperSaslUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
