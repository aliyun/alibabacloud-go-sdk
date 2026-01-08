// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupInviteLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatGroupInviteLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatGroupInviteLinkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatGroupInviteLinkResponseBody) *DeleteChatGroupInviteLinkResponse
	GetBody() *DeleteChatGroupInviteLinkResponseBody
}

type DeleteChatGroupInviteLinkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatGroupInviteLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatGroupInviteLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupInviteLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupInviteLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatGroupInviteLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatGroupInviteLinkResponse) GetBody() *DeleteChatGroupInviteLinkResponseBody {
	return s.Body
}

func (s *DeleteChatGroupInviteLinkResponse) SetHeaders(v map[string]*string) *DeleteChatGroupInviteLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatGroupInviteLinkResponse) SetStatusCode(v int32) *DeleteChatGroupInviteLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatGroupInviteLinkResponse) SetBody(v *DeleteChatGroupInviteLinkResponseBody) *DeleteChatGroupInviteLinkResponse {
	s.Body = v
	return s
}

func (s *DeleteChatGroupInviteLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
