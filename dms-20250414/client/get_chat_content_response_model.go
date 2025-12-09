// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatContentResponse
	GetStatusCode() *int32
	SetBody(v *GetChatContentResponseBody) *GetChatContentResponse
	GetBody() *GetChatContentResponseBody
}

type GetChatContentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatContentResponse) GoString() string {
	return s.String()
}

func (s *GetChatContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatContentResponse) GetBody() *GetChatContentResponseBody {
	return s.Body
}

func (s *GetChatContentResponse) SetHeaders(v map[string]*string) *GetChatContentResponse {
	s.Headers = v
	return s
}

func (s *GetChatContentResponse) SetStatusCode(v int32) *GetChatContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatContentResponse) SetBody(v *GetChatContentResponseBody) *GetChatContentResponse {
	s.Body = v
	return s
}

func (s *GetChatContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
