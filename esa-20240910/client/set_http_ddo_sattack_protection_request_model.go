// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalMode(v string) *SetHttpDDoSAttackProtectionRequest
	GetGlobalMode() *string
	SetSiteId(v int64) *SetHttpDDoSAttackProtectionRequest
	GetSiteId() *int64
}

type SetHttpDDoSAttackProtectionRequest struct {
	// This parameter is required.
	GlobalMode *string `json:"GlobalMode,omitempty" xml:"GlobalMode,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetHttpDDoSAttackProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackProtectionRequest) GetGlobalMode() *string {
	return s.GlobalMode
}

func (s *SetHttpDDoSAttackProtectionRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetHttpDDoSAttackProtectionRequest) SetGlobalMode(v string) *SetHttpDDoSAttackProtectionRequest {
	s.GlobalMode = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionRequest) SetSiteId(v int64) *SetHttpDDoSAttackProtectionRequest {
	s.SiteId = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionRequest) Validate() error {
	return dara.Validate(s)
}
