// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStsTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStsTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStsTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetStsTokenResponseBody) *GetStsTokenResponse
	GetBody() *GetStsTokenResponseBody
}

type GetStsTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStsTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStsTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStsTokenResponse) GoString() string {
	return s.String()
}

func (s *GetStsTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStsTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStsTokenResponse) GetBody() *GetStsTokenResponseBody {
	return s.Body
}

func (s *GetStsTokenResponse) SetHeaders(v map[string]*string) *GetStsTokenResponse {
	s.Headers = v
	return s
}

func (s *GetStsTokenResponse) SetStatusCode(v int32) *GetStsTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStsTokenResponse) SetBody(v *GetStsTokenResponseBody) *GetStsTokenResponse {
	s.Body = v
	return s
}

func (s *GetStsTokenResponse) Validate() error {
	return dara.Validate(s)
}
