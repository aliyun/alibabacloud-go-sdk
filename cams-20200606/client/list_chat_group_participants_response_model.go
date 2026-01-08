// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupParticipantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListChatGroupParticipantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListChatGroupParticipantsResponse
	GetStatusCode() *int32
	SetBody(v *ListChatGroupParticipantsResponseBody) *ListChatGroupParticipantsResponse
	GetBody() *ListChatGroupParticipantsResponseBody
}

type ListChatGroupParticipantsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChatGroupParticipantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChatGroupParticipantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupParticipantsResponse) GoString() string {
	return s.String()
}

func (s *ListChatGroupParticipantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListChatGroupParticipantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListChatGroupParticipantsResponse) GetBody() *ListChatGroupParticipantsResponseBody {
	return s.Body
}

func (s *ListChatGroupParticipantsResponse) SetHeaders(v map[string]*string) *ListChatGroupParticipantsResponse {
	s.Headers = v
	return s
}

func (s *ListChatGroupParticipantsResponse) SetStatusCode(v int32) *ListChatGroupParticipantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChatGroupParticipantsResponse) SetBody(v *ListChatGroupParticipantsResponseBody) *ListChatGroupParticipantsResponse {
	s.Body = v
	return s
}

func (s *ListChatGroupParticipantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
