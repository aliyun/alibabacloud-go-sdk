// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupParticipantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChatGroupParticipantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChatGroupParticipantsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChatGroupParticipantsResponseBody) *DeleteChatGroupParticipantsResponse
	GetBody() *DeleteChatGroupParticipantsResponseBody
}

type DeleteChatGroupParticipantsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChatGroupParticipantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChatGroupParticipantsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupParticipantsResponse) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupParticipantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChatGroupParticipantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChatGroupParticipantsResponse) GetBody() *DeleteChatGroupParticipantsResponseBody {
	return s.Body
}

func (s *DeleteChatGroupParticipantsResponse) SetHeaders(v map[string]*string) *DeleteChatGroupParticipantsResponse {
	s.Headers = v
	return s
}

func (s *DeleteChatGroupParticipantsResponse) SetStatusCode(v int32) *DeleteChatGroupParticipantsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChatGroupParticipantsResponse) SetBody(v *DeleteChatGroupParticipantsResponseBody) *DeleteChatGroupParticipantsResponse {
	s.Body = v
	return s
}

func (s *DeleteChatGroupParticipantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
