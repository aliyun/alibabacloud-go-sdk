// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatInstanceSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryChatInstanceSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryChatInstanceSessionsResponse
	GetStatusCode() *int32
	SetBody(v *QueryChatInstanceSessionsResponseBody) *QueryChatInstanceSessionsResponse
	GetBody() *QueryChatInstanceSessionsResponseBody
}

type QueryChatInstanceSessionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChatInstanceSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChatInstanceSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryChatInstanceSessionsResponse) GoString() string {
	return s.String()
}

func (s *QueryChatInstanceSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryChatInstanceSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryChatInstanceSessionsResponse) GetBody() *QueryChatInstanceSessionsResponseBody {
	return s.Body
}

func (s *QueryChatInstanceSessionsResponse) SetHeaders(v map[string]*string) *QueryChatInstanceSessionsResponse {
	s.Headers = v
	return s
}

func (s *QueryChatInstanceSessionsResponse) SetStatusCode(v int32) *QueryChatInstanceSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChatInstanceSessionsResponse) SetBody(v *QueryChatInstanceSessionsResponseBody) *QueryChatInstanceSessionsResponse {
	s.Body = v
	return s
}

func (s *QueryChatInstanceSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
