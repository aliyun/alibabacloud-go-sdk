// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteSkillRequest
	GetClientToken() *string
	SetSkillId(v string) *DeleteSkillRequest
	GetSkillId() *string
}

type DeleteSkillRequest struct {
	// A token that you provide to ensure request idempotence. The value must be unique for each request. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The skill ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-06e9dca2-0ac9-4d2e-a965-e9db9c057e00
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s DeleteSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteSkillRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *DeleteSkillRequest) SetClientToken(v string) *DeleteSkillRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSkillRequest) SetSkillId(v string) *DeleteSkillRequest {
	s.SkillId = &v
	return s
}

func (s *DeleteSkillRequest) Validate() error {
	return dara.Validate(s)
}
