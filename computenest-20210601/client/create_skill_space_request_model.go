// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSkillSpaceRequest
	GetClientToken() *string
	SetSkillSpaceDescription(v string) *CreateSkillSpaceRequest
	GetSkillSpaceDescription() *string
	SetSkillSpaceName(v string) *CreateSkillSpaceRequest
	GetSkillSpaceName() *string
}

type CreateSkillSpaceRequest struct {
	// A client-generated token to ensure the idempotence of the request. The token must be unique across requests. The **ClientToken*	- value can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the skill space.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111
	SkillSpaceDescription *string `json:"SkillSpaceDescription,omitempty" xml:"SkillSpaceDescription,omitempty"`
	// The name of the skill space.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111
	SkillSpaceName *string `json:"SkillSpaceName,omitempty" xml:"SkillSpaceName,omitempty"`
}

func (s CreateSkillSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSkillSpaceRequest) GetSkillSpaceDescription() *string {
	return s.SkillSpaceDescription
}

func (s *CreateSkillSpaceRequest) GetSkillSpaceName() *string {
	return s.SkillSpaceName
}

func (s *CreateSkillSpaceRequest) SetClientToken(v string) *CreateSkillSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSkillSpaceRequest) SetSkillSpaceDescription(v string) *CreateSkillSpaceRequest {
	s.SkillSpaceDescription = &v
	return s
}

func (s *CreateSkillSpaceRequest) SetSkillSpaceName(v string) *CreateSkillSpaceRequest {
	s.SkillSpaceName = &v
	return s
}

func (s *CreateSkillSpaceRequest) Validate() error {
	return dara.Validate(s)
}
