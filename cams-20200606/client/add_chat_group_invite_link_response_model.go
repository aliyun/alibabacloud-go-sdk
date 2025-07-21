// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatGroupInviteLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddChatGroupInviteLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddChatGroupInviteLinkResponse
	GetStatusCode() *int32
	SetBody(v *AddChatGroupInviteLinkResponseBody) *AddChatGroupInviteLinkResponse
	GetBody() *AddChatGroupInviteLinkResponseBody
}

type AddChatGroupInviteLinkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddChatGroupInviteLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddChatGroupInviteLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s AddChatGroupInviteLinkResponse) GoString() string {
	return s.String()
}

func (s *AddChatGroupInviteLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddChatGroupInviteLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddChatGroupInviteLinkResponse) GetBody() *AddChatGroupInviteLinkResponseBody {
	return s.Body
}

func (s *AddChatGroupInviteLinkResponse) SetHeaders(v map[string]*string) *AddChatGroupInviteLinkResponse {
	s.Headers = v
	return s
}

func (s *AddChatGroupInviteLinkResponse) SetStatusCode(v int32) *AddChatGroupInviteLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *AddChatGroupInviteLinkResponse) SetBody(v *AddChatGroupInviteLinkResponseBody) *AddChatGroupInviteLinkResponse {
	s.Body = v
	return s
}

func (s *AddChatGroupInviteLinkResponse) Validate() error {
	return dara.Validate(s)
}
