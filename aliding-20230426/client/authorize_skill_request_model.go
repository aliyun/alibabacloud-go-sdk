// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionCodes(v []*string) *AuthorizeSkillRequest
	GetPermissionCodes() []*string
	SetSourceIdOfAssistantId(v string) *AuthorizeSkillRequest
	GetSourceIdOfAssistantId() *string
}

type AuthorizeSkillRequest struct {
	PermissionCodes       []*string `json:"PermissionCodes,omitempty" xml:"PermissionCodes,omitempty" type:"Repeated"`
	SourceIdOfAssistantId *string   `json:"SourceIdOfAssistantId,omitempty" xml:"SourceIdOfAssistantId,omitempty"`
}

func (s AuthorizeSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSkillRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSkillRequest) GetPermissionCodes() []*string {
	return s.PermissionCodes
}

func (s *AuthorizeSkillRequest) GetSourceIdOfAssistantId() *string {
	return s.SourceIdOfAssistantId
}

func (s *AuthorizeSkillRequest) SetPermissionCodes(v []*string) *AuthorizeSkillRequest {
	s.PermissionCodes = v
	return s
}

func (s *AuthorizeSkillRequest) SetSourceIdOfAssistantId(v string) *AuthorizeSkillRequest {
	s.SourceIdOfAssistantId = &v
	return s
}

func (s *AuthorizeSkillRequest) Validate() error {
	return dara.Validate(s)
}
