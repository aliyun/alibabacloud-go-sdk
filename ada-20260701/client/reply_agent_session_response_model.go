// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyAgentSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplyAgentSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplyAgentSessionResponse
	GetStatusCode() *int32
	SetBody(v *ReplyAgentSessionResponseBody) *ReplyAgentSessionResponse
	GetBody() *ReplyAgentSessionResponseBody
}

type ReplyAgentSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplyAgentSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplyAgentSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionResponse) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplyAgentSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplyAgentSessionResponse) GetBody() *ReplyAgentSessionResponseBody {
	return s.Body
}

func (s *ReplyAgentSessionResponse) SetHeaders(v map[string]*string) *ReplyAgentSessionResponse {
	s.Headers = v
	return s
}

func (s *ReplyAgentSessionResponse) SetStatusCode(v int32) *ReplyAgentSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplyAgentSessionResponse) SetBody(v *ReplyAgentSessionResponseBody) *ReplyAgentSessionResponse {
	s.Body = v
	return s
}

func (s *ReplyAgentSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
