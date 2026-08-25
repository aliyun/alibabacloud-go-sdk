// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMFAAuthenticationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMFAAuthenticationStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetMFAAuthenticationStatusResponseBody) *GetMFAAuthenticationStatusResponse
	GetBody() *GetMFAAuthenticationStatusResponseBody
}

type GetMFAAuthenticationStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMFAAuthenticationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMFAAuthenticationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMFAAuthenticationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMFAAuthenticationStatusResponse) GetBody() *GetMFAAuthenticationStatusResponseBody {
	return s.Body
}

func (s *GetMFAAuthenticationStatusResponse) SetHeaders(v map[string]*string) *GetMFAAuthenticationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMFAAuthenticationStatusResponse) SetStatusCode(v int32) *GetMFAAuthenticationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMFAAuthenticationStatusResponse) SetBody(v *GetMFAAuthenticationStatusResponseBody) *GetMFAAuthenticationStatusResponse {
	s.Body = v
	return s
}

func (s *GetMFAAuthenticationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
