// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSkillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionCodesShrink(v string) *AuthorizeSkillShrinkRequest
	GetPermissionCodesShrink() *string
	SetSourceIdOfAssistantId(v string) *AuthorizeSkillShrinkRequest
	GetSourceIdOfAssistantId() *string
}

type AuthorizeSkillShrinkRequest struct {
	PermissionCodesShrink *string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty"`
	SourceIdOfAssistantId *string `json:"SourceIdOfAssistantId,omitempty" xml:"SourceIdOfAssistantId,omitempty"`
}

func (s AuthorizeSkillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSkillShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSkillShrinkRequest) GetPermissionCodesShrink() *string {
	return s.PermissionCodesShrink
}

func (s *AuthorizeSkillShrinkRequest) GetSourceIdOfAssistantId() *string {
	return s.SourceIdOfAssistantId
}

func (s *AuthorizeSkillShrinkRequest) SetPermissionCodesShrink(v string) *AuthorizeSkillShrinkRequest {
	s.PermissionCodesShrink = &v
	return s
}

func (s *AuthorizeSkillShrinkRequest) SetSourceIdOfAssistantId(v string) *AuthorizeSkillShrinkRequest {
	s.SourceIdOfAssistantId = &v
	return s
}

func (s *AuthorizeSkillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
