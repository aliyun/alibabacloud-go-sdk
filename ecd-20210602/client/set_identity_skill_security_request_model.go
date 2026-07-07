// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentitySkillSecurityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *SetIdentitySkillSecurityRequest
	GetEnabled() *bool
	SetIdentityIds(v []*SetIdentitySkillSecurityRequestIdentityIds) *SetIdentitySkillSecurityRequest
	GetIdentityIds() []*SetIdentitySkillSecurityRequestIdentityIds
	SetSkillChannel(v string) *SetIdentitySkillSecurityRequest
	GetSkillChannel() *string
}

type SetIdentitySkillSecurityRequest struct {
	// Specifies whether to enable the skill installation permission. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of resource information.
	//
	// This parameter is required.
	IdentityIds []*SetIdentitySkillSecurityRequestIdentityIds `json:"IdentityIds,omitempty" xml:"IdentityIds,omitempty" type:"Repeated"`
	// The skill channel. Valid values:
	//
	// - ENTERPRISE: enterprise edition.
	//
	// - BUSINESS: business edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTERPRISE
	SkillChannel *string `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
}

func (s SetIdentitySkillSecurityRequest) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillSecurityRequest) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillSecurityRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *SetIdentitySkillSecurityRequest) GetIdentityIds() []*SetIdentitySkillSecurityRequestIdentityIds {
	return s.IdentityIds
}

func (s *SetIdentitySkillSecurityRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *SetIdentitySkillSecurityRequest) SetEnabled(v bool) *SetIdentitySkillSecurityRequest {
	s.Enabled = &v
	return s
}

func (s *SetIdentitySkillSecurityRequest) SetIdentityIds(v []*SetIdentitySkillSecurityRequestIdentityIds) *SetIdentitySkillSecurityRequest {
	s.IdentityIds = v
	return s
}

func (s *SetIdentitySkillSecurityRequest) SetSkillChannel(v string) *SetIdentitySkillSecurityRequest {
	s.SkillChannel = &v
	return s
}

func (s *SetIdentitySkillSecurityRequest) Validate() error {
	if s.IdentityIds != nil {
		for _, item := range s.IdentityIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetIdentitySkillSecurityRequestIdentityIds struct {
	// The resource information ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-b9ej3xiok4tjbgf9x
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetIdentitySkillSecurityRequestIdentityIds) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillSecurityRequestIdentityIds) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillSecurityRequestIdentityIds) GetIdentityId() *string {
	return s.IdentityId
}

func (s *SetIdentitySkillSecurityRequestIdentityIds) GetRegionId() *string {
	return s.RegionId
}

func (s *SetIdentitySkillSecurityRequestIdentityIds) SetIdentityId(v string) *SetIdentitySkillSecurityRequestIdentityIds {
	s.IdentityId = &v
	return s
}

func (s *SetIdentitySkillSecurityRequestIdentityIds) SetRegionId(v string) *SetIdentitySkillSecurityRequestIdentityIds {
	s.RegionId = &v
	return s
}

func (s *SetIdentitySkillSecurityRequestIdentityIds) Validate() error {
	return dara.Validate(s)
}
