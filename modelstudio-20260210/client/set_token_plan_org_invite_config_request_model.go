// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTokenPlanOrgInviteConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultRoleId(v string) *SetTokenPlanOrgInviteConfigRequest
	GetDefaultRoleId() *string
	SetSeatAssignStrategy(v string) *SetTokenPlanOrgInviteConfigRequest
	GetSeatAssignStrategy() *string
}

type SetTokenPlanOrgInviteConfigRequest struct {
	// The ID of the organization role assigned by default. Valid values:
	//
	// - SYSTEM_ROLE_ORG_ADMIN
	//
	// - SYSTEM_ROLE_ORG_MEMBER
	//
	// This parameter is required.
	//
	// example:
	//
	// ORG_MEMBER
	DefaultRoleId *string `json:"DefaultRoleId,omitempty" xml:"DefaultRoleId,omitempty"`
	// The default seat allocation policy. Valid values:
	//
	// - HIGH_TO_LOW
	//
	// - LOW_TO_HIGH
	//
	// - NONE
	//
	// This parameter is required.
	//
	// example:
	//
	// NONE
	SeatAssignStrategy *string `json:"SeatAssignStrategy,omitempty" xml:"SeatAssignStrategy,omitempty"`
}

func (s SetTokenPlanOrgInviteConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetTokenPlanOrgInviteConfigRequest) GoString() string {
	return s.String()
}

func (s *SetTokenPlanOrgInviteConfigRequest) GetDefaultRoleId() *string {
	return s.DefaultRoleId
}

func (s *SetTokenPlanOrgInviteConfigRequest) GetSeatAssignStrategy() *string {
	return s.SeatAssignStrategy
}

func (s *SetTokenPlanOrgInviteConfigRequest) SetDefaultRoleId(v string) *SetTokenPlanOrgInviteConfigRequest {
	s.DefaultRoleId = &v
	return s
}

func (s *SetTokenPlanOrgInviteConfigRequest) SetSeatAssignStrategy(v string) *SetTokenPlanOrgInviteConfigRequest {
	s.SeatAssignStrategy = &v
	return s
}

func (s *SetTokenPlanOrgInviteConfigRequest) Validate() error {
	return dara.Validate(s)
}
