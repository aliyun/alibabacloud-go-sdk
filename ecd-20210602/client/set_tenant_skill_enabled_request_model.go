// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTenantSkillEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *SetTenantSkillEnabledRequest
	GetEnabled() *bool
	SetSkillChannel(v string) *SetTenantSkillEnabledRequest
	GetSkillChannel() *string
	SetSkillIds(v []*string) *SetTenantSkillEnabledRequest
	GetSkillIds() []*string
}

type SetTenantSkillEnabledRequest struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// BUSINESS
	SkillChannel *string   `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
	SkillIds     []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
}

func (s SetTenantSkillEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s SetTenantSkillEnabledRequest) GoString() string {
	return s.String()
}

func (s *SetTenantSkillEnabledRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *SetTenantSkillEnabledRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *SetTenantSkillEnabledRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *SetTenantSkillEnabledRequest) SetEnabled(v bool) *SetTenantSkillEnabledRequest {
	s.Enabled = &v
	return s
}

func (s *SetTenantSkillEnabledRequest) SetSkillChannel(v string) *SetTenantSkillEnabledRequest {
	s.SkillChannel = &v
	return s
}

func (s *SetTenantSkillEnabledRequest) SetSkillIds(v []*string) *SetTenantSkillEnabledRequest {
	s.SkillIds = v
	return s
}

func (s *SetTenantSkillEnabledRequest) Validate() error {
	return dara.Validate(s)
}
