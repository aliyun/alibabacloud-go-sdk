// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneEffectiveScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionZoneEffectiveScopeRequest
	GetClientToken() *string
	SetEffectiveScopes(v []*UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) *UpdateRecursionZoneEffectiveScopeRequest
	GetEffectiveScopes() []*UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes
	SetZoneId(v string) *UpdateRecursionZoneEffectiveScopeRequest
	GetZoneId() *string
}

type UpdateRecursionZoneEffectiveScopeRequest struct {
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken     *string                                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EffectiveScopes []*UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes `json:"EffectiveScopes,omitempty" xml:"EffectiveScopes,omitempty" type:"Repeated"`
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateRecursionZoneEffectiveScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneEffectiveScopeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneEffectiveScopeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionZoneEffectiveScopeRequest) GetEffectiveScopes() []*UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes {
	return s.EffectiveScopes
}

func (s *UpdateRecursionZoneEffectiveScopeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateRecursionZoneEffectiveScopeRequest) SetClientToken(v string) *UpdateRecursionZoneEffectiveScopeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeRequest) SetEffectiveScopes(v []*UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) *UpdateRecursionZoneEffectiveScopeRequest {
	s.EffectiveScopes = v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeRequest) SetZoneId(v string) *UpdateRecursionZoneEffectiveScopeRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeRequest) Validate() error {
	if s.EffectiveScopes != nil {
		for _, item := range s.EffectiveScopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes struct {
	// example:
	//
	// account
	EffectiveType *string   `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	Scope         []*string `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) GetEffectiveType() *string {
	return s.EffectiveType
}

func (s *UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) GetScope() []*string {
	return s.Scope
}

func (s *UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) SetEffectiveType(v string) *UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes {
	s.EffectiveType = &v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) SetScope(v []*string) *UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes {
	s.Scope = v
	return s
}

func (s *UpdateRecursionZoneEffectiveScopeRequestEffectiveScopes) Validate() error {
	return dara.Validate(s)
}
