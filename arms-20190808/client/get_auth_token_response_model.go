// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse
	GetBody() *GetAuthTokenResponseBody
}

type GetAuthTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthTokenResponse) GetBody() *GetAuthTokenResponseBody {
	return s.Body
}

func (s *GetAuthTokenResponse) SetHeaders(v map[string]*string) *GetAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenResponse) SetStatusCode(v int32) *GetAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthTokenResponse) SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse {
	s.Body = v
	return s
}

func (s *GetAuthTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
