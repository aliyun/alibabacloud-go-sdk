// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMFAAuthenticationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetMFAAuthenticationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetMFAAuthenticationStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetMFAAuthenticationStatusResponseBody) *SetMFAAuthenticationStatusResponse
	GetBody() *SetMFAAuthenticationStatusResponseBody
}

type SetMFAAuthenticationStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMFAAuthenticationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMFAAuthenticationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetMFAAuthenticationStatusResponse) GoString() string {
	return s.String()
}

func (s *SetMFAAuthenticationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetMFAAuthenticationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetMFAAuthenticationStatusResponse) GetBody() *SetMFAAuthenticationStatusResponseBody {
	return s.Body
}

func (s *SetMFAAuthenticationStatusResponse) SetHeaders(v map[string]*string) *SetMFAAuthenticationStatusResponse {
	s.Headers = v
	return s
}

func (s *SetMFAAuthenticationStatusResponse) SetStatusCode(v int32) *SetMFAAuthenticationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMFAAuthenticationStatusResponse) SetBody(v *SetMFAAuthenticationStatusResponseBody) *SetMFAAuthenticationStatusResponse {
	s.Body = v
	return s
}

func (s *SetMFAAuthenticationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
