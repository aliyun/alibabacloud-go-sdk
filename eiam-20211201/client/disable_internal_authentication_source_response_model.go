// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInternalAuthenticationSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableInternalAuthenticationSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableInternalAuthenticationSourceResponse
	GetStatusCode() *int32
	SetBody(v *DisableInternalAuthenticationSourceResponseBody) *DisableInternalAuthenticationSourceResponse
	GetBody() *DisableInternalAuthenticationSourceResponseBody
}

type DisableInternalAuthenticationSourceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInternalAuthenticationSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInternalAuthenticationSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableInternalAuthenticationSourceResponse) GoString() string {
	return s.String()
}

func (s *DisableInternalAuthenticationSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableInternalAuthenticationSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableInternalAuthenticationSourceResponse) GetBody() *DisableInternalAuthenticationSourceResponseBody {
	return s.Body
}

func (s *DisableInternalAuthenticationSourceResponse) SetHeaders(v map[string]*string) *DisableInternalAuthenticationSourceResponse {
	s.Headers = v
	return s
}

func (s *DisableInternalAuthenticationSourceResponse) SetStatusCode(v int32) *DisableInternalAuthenticationSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInternalAuthenticationSourceResponse) SetBody(v *DisableInternalAuthenticationSourceResponseBody) *DisableInternalAuthenticationSourceResponse {
	s.Body = v
	return s
}

func (s *DisableInternalAuthenticationSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
