// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseChatInstanceSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseChatInstanceSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseChatInstanceSessionsResponse
	GetStatusCode() *int32
	SetBody(v *CloseChatInstanceSessionsResponseBody) *CloseChatInstanceSessionsResponse
	GetBody() *CloseChatInstanceSessionsResponseBody
}

type CloseChatInstanceSessionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseChatInstanceSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseChatInstanceSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseChatInstanceSessionsResponse) GoString() string {
	return s.String()
}

func (s *CloseChatInstanceSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseChatInstanceSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseChatInstanceSessionsResponse) GetBody() *CloseChatInstanceSessionsResponseBody {
	return s.Body
}

func (s *CloseChatInstanceSessionsResponse) SetHeaders(v map[string]*string) *CloseChatInstanceSessionsResponse {
	s.Headers = v
	return s
}

func (s *CloseChatInstanceSessionsResponse) SetStatusCode(v int32) *CloseChatInstanceSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseChatInstanceSessionsResponse) SetBody(v *CloseChatInstanceSessionsResponseBody) *CloseChatInstanceSessionsResponse {
	s.Body = v
	return s
}

func (s *CloseChatInstanceSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
