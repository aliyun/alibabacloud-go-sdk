// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelAuthTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelAuthTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelAuthTokenResponse
	GetStatusCode() *int32
	SetBody(v *DelAuthTokenResponseBody) *DelAuthTokenResponse
	GetBody() *DelAuthTokenResponseBody
}

type DelAuthTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelAuthTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DelAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *DelAuthTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelAuthTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelAuthTokenResponse) GetBody() *DelAuthTokenResponseBody {
	return s.Body
}

func (s *DelAuthTokenResponse) SetHeaders(v map[string]*string) *DelAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *DelAuthTokenResponse) SetStatusCode(v int32) *DelAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DelAuthTokenResponse) SetBody(v *DelAuthTokenResponseBody) *DelAuthTokenResponse {
	s.Body = v
	return s
}

func (s *DelAuthTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
