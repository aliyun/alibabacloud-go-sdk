// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatListResponse
	GetStatusCode() *int32
	SetBody(v *GetChatListResponseBody) *GetChatListResponse
	GetBody() *GetChatListResponseBody
}

type GetChatListResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatListResponse) GoString() string {
	return s.String()
}

func (s *GetChatListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatListResponse) GetBody() *GetChatListResponseBody {
	return s.Body
}

func (s *GetChatListResponse) SetHeaders(v map[string]*string) *GetChatListResponse {
	s.Headers = v
	return s
}

func (s *GetChatListResponse) SetStatusCode(v int32) *GetChatListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatListResponse) SetBody(v *GetChatListResponseBody) *GetChatListResponse {
	s.Body = v
	return s
}

func (s *GetChatListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
