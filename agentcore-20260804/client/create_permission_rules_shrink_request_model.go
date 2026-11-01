// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePermissionRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreatePermissionRulesShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreatePermissionRulesShrinkRequest
	GetClientToken() *string
}

type CreatePermissionRulesShrinkRequest struct {
	// The request body for creating permission rules.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreatePermissionRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePermissionRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePermissionRulesShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreatePermissionRulesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePermissionRulesShrinkRequest) SetBodyShrink(v string) *CreatePermissionRulesShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreatePermissionRulesShrinkRequest) SetClientToken(v string) *CreatePermissionRulesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePermissionRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
