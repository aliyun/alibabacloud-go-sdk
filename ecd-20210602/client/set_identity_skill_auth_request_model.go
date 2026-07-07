// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdentitySkillAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoInstall(v bool) *SetIdentitySkillAuthRequest
	GetAutoInstall() *bool
	SetIdentities(v []*SetIdentitySkillAuthRequestIdentities) *SetIdentitySkillAuthRequest
	GetIdentities() []*SetIdentitySkillAuthRequestIdentities
	SetOperationType(v string) *SetIdentitySkillAuthRequest
	GetOperationType() *string
	SetSkillChannel(v string) *SetIdentitySkillAuthRequest
	GetSkillChannel() *string
	SetSkillIds(v []*string) *SetIdentitySkillAuthRequest
	GetSkillIds() []*string
}

type SetIdentitySkillAuthRequest struct {
	// Specifies whether to automatically install. Valid values:
	//
	// - true: yes
	//
	// - false: no
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The list of authorized objects.
	//
	// This parameter is required.
	Identities []*SetIdentitySkillAuthRequestIdentities `json:"Identities,omitempty" xml:"Identities,omitempty" type:"Repeated"`
	// The operation type.
	//
	// This parameter is required.
	//
	// example:
	//
	// SET_AUTH
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The skill channel. Valid values:
	//
	// - ENTERPRISE: enterprise edition
	//
	// - BUSINESS: business edition
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTERPRISE
	SkillChannel *string `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
	// The list of skill IDs.
	//
	// This parameter is required.
	SkillIds []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
}

func (s SetIdentitySkillAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillAuthRequest) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillAuthRequest) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *SetIdentitySkillAuthRequest) GetIdentities() []*SetIdentitySkillAuthRequestIdentities {
	return s.Identities
}

func (s *SetIdentitySkillAuthRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *SetIdentitySkillAuthRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *SetIdentitySkillAuthRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *SetIdentitySkillAuthRequest) SetAutoInstall(v bool) *SetIdentitySkillAuthRequest {
	s.AutoInstall = &v
	return s
}

func (s *SetIdentitySkillAuthRequest) SetIdentities(v []*SetIdentitySkillAuthRequestIdentities) *SetIdentitySkillAuthRequest {
	s.Identities = v
	return s
}

func (s *SetIdentitySkillAuthRequest) SetOperationType(v string) *SetIdentitySkillAuthRequest {
	s.OperationType = &v
	return s
}

func (s *SetIdentitySkillAuthRequest) SetSkillChannel(v string) *SetIdentitySkillAuthRequest {
	s.SkillChannel = &v
	return s
}

func (s *SetIdentitySkillAuthRequest) SetSkillIds(v []*string) *SetIdentitySkillAuthRequest {
	s.SkillIds = v
	return s
}

func (s *SetIdentitySkillAuthRequest) Validate() error {
	if s.Identities != nil {
		for _, item := range s.Identities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetIdentitySkillAuthRequestIdentities struct {
	// The ID of the authorized object.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-av4u9m5ghko26****
	IdentityId *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetIdentitySkillAuthRequestIdentities) String() string {
	return dara.Prettify(s)
}

func (s SetIdentitySkillAuthRequestIdentities) GoString() string {
	return s.String()
}

func (s *SetIdentitySkillAuthRequestIdentities) GetIdentityId() *string {
	return s.IdentityId
}

func (s *SetIdentitySkillAuthRequestIdentities) GetRegionId() *string {
	return s.RegionId
}

func (s *SetIdentitySkillAuthRequestIdentities) SetIdentityId(v string) *SetIdentitySkillAuthRequestIdentities {
	s.IdentityId = &v
	return s
}

func (s *SetIdentitySkillAuthRequestIdentities) SetRegionId(v string) *SetIdentitySkillAuthRequestIdentities {
	s.RegionId = &v
	return s
}

func (s *SetIdentitySkillAuthRequestIdentities) Validate() error {
	return dara.Validate(s)
}
