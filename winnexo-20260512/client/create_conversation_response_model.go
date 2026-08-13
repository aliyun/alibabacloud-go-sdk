// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConversationResponse
	GetStatusCode() *int32
	SetBody(v *CreateConversationResponseBody) *CreateConversationResponse
	GetBody() *CreateConversationResponseBody
}

type CreateConversationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConversationResponse) GetBody() *CreateConversationResponseBody {
	return s.Body
}

func (s *CreateConversationResponse) SetHeaders(v map[string]*string) *CreateConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateConversationResponse) SetStatusCode(v int32) *CreateConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConversationResponse) SetBody(v *CreateConversationResponseBody) *CreateConversationResponse {
	s.Body = v
	return s
}

func (s *CreateConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
