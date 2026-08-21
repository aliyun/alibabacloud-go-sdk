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
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length.
	//
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
	// ss-xxxxx
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
