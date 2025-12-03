// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConversationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConversationsResponse
	GetStatusCode() *int32
	SetBody(v *GetConversationsResponseBody) *GetConversationsResponse
	GetBody() *GetConversationsResponseBody
}

type GetConversationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConversationsResponse) GoString() string {
	return s.String()
}

func (s *GetConversationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConversationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConversationsResponse) GetBody() *GetConversationsResponseBody {
	return s.Body
}

func (s *GetConversationsResponse) SetHeaders(v map[string]*string) *GetConversationsResponse {
	s.Headers = v
	return s
}

func (s *GetConversationsResponse) SetStatusCode(v int32) *GetConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationsResponse) SetBody(v *GetConversationsResponseBody) *GetConversationsResponse {
	s.Body = v
	return s
}

func (s *GetConversationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
