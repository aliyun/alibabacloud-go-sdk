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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the SkillSpace.
	//
	// This parameter is required.
	//
	// example:
	//
	// All-in-one office toolkit — generate reports, process data, manage files, and streamline workflows effortlessly.
	SkillSpaceDescription *string `json:"SkillSpaceDescription,omitempty" xml:"SkillSpaceDescription,omitempty"`
	// The name of the SkillSpace.
	//
	// This parameter is required.
	//
	// example:
	//
	// office-toolkit-skills
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
