// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserIsGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *CheckUserIsGroupMemberRequest
	GetOpenConversationId() *string
}

type CheckUserIsGroupMemberRequest struct {
	// example:
	//
	// cidB8Pz*******FIWPv2PMA==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
}

func (s CheckUserIsGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUserIsGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *CheckUserIsGroupMemberRequest) SetOpenConversationId(v string) *CheckUserIsGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CheckUserIsGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
