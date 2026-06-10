// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSkillSpaceRequest
	GetClientToken() *string
	SetSkillSpaceId(v string) *DeleteSkillSpaceRequest
	GetSkillSpaceId() *string
}

type DeleteSkillSpaceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// SkillSpace  ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ss-11111111111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
}

func (s DeleteSkillSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSkillSpaceRequest) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *DeleteSkillSpaceRequest) SetClientToken(v string) *DeleteSkillSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSkillSpaceRequest) SetSkillSpaceId(v string) *DeleteSkillSpaceRequest {
	s.SkillSpaceId = &v
	return s
}

func (s *DeleteSkillSpaceRequest) Validate() error {
	return dara.Validate(s)
}
