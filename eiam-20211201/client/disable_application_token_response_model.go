// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationTokenResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationTokenResponseBody) *DisableApplicationTokenResponse
	GetBody() *DisableApplicationTokenResponseBody
}

type DisableApplicationTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationTokenResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationTokenResponse) GetBody() *DisableApplicationTokenResponseBody {
	return s.Body
}

func (s *DisableApplicationTokenResponse) SetHeaders(v map[string]*string) *DisableApplicationTokenResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationTokenResponse) SetStatusCode(v int32) *DisableApplicationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationTokenResponse) SetBody(v *DisableApplicationTokenResponseBody) *DisableApplicationTokenResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
