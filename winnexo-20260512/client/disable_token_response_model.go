// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableTokenResponse
	GetStatusCode() *int32
	SetBody(v *DisableTokenResponseBody) *DisableTokenResponse
	GetBody() *DisableTokenResponseBody
}

type DisableTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableTokenResponse) GoString() string {
	return s.String()
}

func (s *DisableTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableTokenResponse) GetBody() *DisableTokenResponseBody {
	return s.Body
}

func (s *DisableTokenResponse) SetHeaders(v map[string]*string) *DisableTokenResponse {
	s.Headers = v
	return s
}

func (s *DisableTokenResponse) SetStatusCode(v int32) *DisableTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableTokenResponse) SetBody(v *DisableTokenResponseBody) *DisableTokenResponse {
	s.Body = v
	return s
}

func (s *DisableTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
