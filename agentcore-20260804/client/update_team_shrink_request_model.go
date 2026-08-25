// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateTeamShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateTeamShrinkRequest
	GetClientToken() *string
}

type UpdateTeamShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateTeamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTeamShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateTeamShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTeamShrinkRequest) SetBodyShrink(v string) *UpdateTeamShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateTeamShrinkRequest) SetClientToken(v string) *UpdateTeamShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTeamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
