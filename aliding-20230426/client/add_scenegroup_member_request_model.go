// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScenegroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *AddScenegroupMemberRequest
	GetOpenConversationId() *string
	SetUserIds(v string) *AddScenegroupMemberRequest
	GetUserIds() *string
}

type AddScenegroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123xx,224xx
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s AddScenegroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s AddScenegroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddScenegroupMemberRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *AddScenegroupMemberRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *AddScenegroupMemberRequest) SetOpenConversationId(v string) *AddScenegroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *AddScenegroupMemberRequest) SetUserIds(v string) *AddScenegroupMemberRequest {
	s.UserIds = &v
	return s
}

func (s *AddScenegroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
