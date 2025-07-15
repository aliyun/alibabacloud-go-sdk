// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatMediaUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChatMediaUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChatMediaUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetChatMediaUrlResponseBody) *GetChatMediaUrlResponse
	GetBody() *GetChatMediaUrlResponseBody
}

type GetChatMediaUrlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChatMediaUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChatMediaUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChatMediaUrlResponse) GoString() string {
	return s.String()
}

func (s *GetChatMediaUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChatMediaUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChatMediaUrlResponse) GetBody() *GetChatMediaUrlResponseBody {
	return s.Body
}

func (s *GetChatMediaUrlResponse) SetHeaders(v map[string]*string) *GetChatMediaUrlResponse {
	s.Headers = v
	return s
}

func (s *GetChatMediaUrlResponse) SetStatusCode(v int32) *GetChatMediaUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChatMediaUrlResponse) SetBody(v *GetChatMediaUrlResponseBody) *GetChatMediaUrlResponse {
	s.Body = v
	return s
}

func (s *GetChatMediaUrlResponse) Validate() error {
	return dara.Validate(s)
}
