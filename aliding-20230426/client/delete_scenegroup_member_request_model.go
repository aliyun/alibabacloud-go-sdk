// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenegroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *DeleteScenegroupMemberRequest
	GetOpenConversationId() *string
	SetUserIds(v string) *DeleteScenegroupMemberRequest
	GetUserIds() *string
}

type DeleteScenegroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidvkLfbOyIiSYqjgvAiWwFow==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123xxx,223xxx
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s DeleteScenegroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenegroupMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteScenegroupMemberRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *DeleteScenegroupMemberRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *DeleteScenegroupMemberRequest) SetOpenConversationId(v string) *DeleteScenegroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DeleteScenegroupMemberRequest) SetUserIds(v string) *DeleteScenegroupMemberRequest {
	s.UserIds = &v
	return s
}

func (s *DeleteScenegroupMemberRequest) Validate() error {
	return dara.Validate(s)
}
