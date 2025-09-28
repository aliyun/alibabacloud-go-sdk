// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsAuthTokensResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmsAuthTokensResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmsAuthTokensResponse
	GetStatusCode() *int32
	SetBody(v *GetSmsAuthTokensResponseBody) *GetSmsAuthTokensResponse
	GetBody() *GetSmsAuthTokensResponseBody
}

type GetSmsAuthTokensResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsAuthTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsAuthTokensResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmsAuthTokensResponse) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmsAuthTokensResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmsAuthTokensResponse) GetBody() *GetSmsAuthTokensResponseBody {
	return s.Body
}

func (s *GetSmsAuthTokensResponse) SetHeaders(v map[string]*string) *GetSmsAuthTokensResponse {
	s.Headers = v
	return s
}

func (s *GetSmsAuthTokensResponse) SetStatusCode(v int32) *GetSmsAuthTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsAuthTokensResponse) SetBody(v *GetSmsAuthTokensResponseBody) *GetSmsAuthTokensResponse {
	s.Body = v
	return s
}

func (s *GetSmsAuthTokensResponse) Validate() error {
	return dara.Validate(s)
}
