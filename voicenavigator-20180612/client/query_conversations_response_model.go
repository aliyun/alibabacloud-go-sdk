// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConversationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConversationsResponse
	GetStatusCode() *int32
	SetBody(v *QueryConversationsResponseBody) *QueryConversationsResponse
	GetBody() *QueryConversationsResponseBody
}

type QueryConversationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConversationsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationsResponse) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConversationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConversationsResponse) GetBody() *QueryConversationsResponseBody {
	return s.Body
}

func (s *QueryConversationsResponse) SetHeaders(v map[string]*string) *QueryConversationsResponse {
	s.Headers = v
	return s
}

func (s *QueryConversationsResponse) SetStatusCode(v int32) *QueryConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationsResponse) SetBody(v *QueryConversationsResponseBody) *QueryConversationsResponse {
	s.Body = v
	return s
}

func (s *QueryConversationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
