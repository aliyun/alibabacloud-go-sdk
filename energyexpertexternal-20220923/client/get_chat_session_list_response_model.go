// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatSessionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatSessionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatSessionListResponse
	GetStatusCode() *int32
	SetBody(v *GetChatSessionListResponseBody) *GetChatSessionListResponse
	GetBody() *GetChatSessionListResponseBody
}

type GetChatSessionListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatSessionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatSessionListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionListResponse) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatSessionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatSessionListResponse) GetBody() *GetChatSessionListResponseBody {
	return s.Body
}

func (s *GetChatSessionListResponse) SetHeaders(v map[string]*string) *GetChatSessionListResponse {
	s.Headers = v
	return s
}

func (s *GetChatSessionListResponse) SetStatusCode(v int32) *GetChatSessionListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatSessionListResponse) SetBody(v *GetChatSessionListResponseBody) *GetChatSessionListResponse {
	s.Body = v
	return s
}

func (s *GetChatSessionListResponse) Validate() error {
	return dara.Validate(s)
}
