// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTenantSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSkillChannel(v string) *DeleteTenantSkillsRequest
	GetSkillChannel() *string
	SetSkillIds(v []*string) *DeleteTenantSkillsRequest
	GetSkillIds() []*string
}

type DeleteTenantSkillsRequest struct {
	// The skill channel. Valid values:
	//
	// - ENTERPRISE: Enterprise edition.
	//
	// - BUSINESS: Business edition.
	//
	// example:
	//
	// ENTERPRISE
	SkillChannel *string `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
	// The list of skill IDs.
	SkillIds []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
}

func (s DeleteTenantSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTenantSkillsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTenantSkillsRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *DeleteTenantSkillsRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *DeleteTenantSkillsRequest) SetSkillChannel(v string) *DeleteTenantSkillsRequest {
	s.SkillChannel = &v
	return s
}

func (s *DeleteTenantSkillsRequest) SetSkillIds(v []*string) *DeleteTenantSkillsRequest {
	s.SkillIds = v
	return s
}

func (s *DeleteTenantSkillsRequest) Validate() error {
	return dara.Validate(s)
}
