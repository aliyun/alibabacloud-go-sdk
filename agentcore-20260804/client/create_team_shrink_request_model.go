// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateTeamShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateTeamShrinkRequest
	GetClientToken() *string
}

type CreateTeamShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateTeamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTeamShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateTeamShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTeamShrinkRequest) SetBodyShrink(v string) *CreateTeamShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateTeamShrinkRequest) SetClientToken(v string) *CreateTeamShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTeamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
