// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneEffectiveScopeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionZoneEffectiveScopeShrinkRequest
	GetClientToken() *string
	SetEffectiveScopesShrink(v string) *UpdateRecursionZoneEffectiveScopeShrinkRequest
	GetEffectiveScopesShrink() *string
	SetZoneId(v string) *UpdateRecursionZoneEffectiveScopeShrinkRequest
	GetZoneId() *string
}

type UpdateRecursionZoneEffectiveScopeShrinkRequest struct {
	// A client-generated token to ensure request idempotence. The token must be unique for each request. It can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// A list of effective scopes.
	EffectiveScopesShrink *string `json:"EffectiveScopes,omitempty" xml:"EffectiveScopes,omitempty"`
	// The unique ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateRecursionZoneEffectiveScopeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneEffectiveScopeShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneEffectiveScopeShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionZoneEffectiveScopeShrinkRequest) GetEffectiveScopesShrink() *string {
	return s.EffectiveScopesShrink
}

func (s *UpdateRecursionZoneEffectiveScopeShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateRecursionZoneEffectiveScopeShrinkRequest) SetClientToken(v string) *UpdateRecursionZoneEffectiveScopeShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeShrinkRequest) SetEffectiveScopesShrink(v string) *UpdateRecursionZoneEffectiveScopeShrinkRequest {
	s.EffectiveScopesShrink = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeShrinkRequest) SetZoneId(v string) *UpdateRecursionZoneEffectiveScopeShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
