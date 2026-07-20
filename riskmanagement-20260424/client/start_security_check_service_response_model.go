// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSecurityCheckServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSecurityCheckServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSecurityCheckServiceResponse
	GetStatusCode() *int32
	SetBody(v *StartSecurityCheckServiceResponseBody) *StartSecurityCheckServiceResponse
	GetBody() *StartSecurityCheckServiceResponseBody
}

type StartSecurityCheckServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSecurityCheckServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSecurityCheckServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSecurityCheckServiceResponse) GoString() string {
	return s.String()
}

func (s *StartSecurityCheckServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSecurityCheckServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSecurityCheckServiceResponse) GetBody() *StartSecurityCheckServiceResponseBody {
	return s.Body
}

func (s *StartSecurityCheckServiceResponse) SetHeaders(v map[string]*string) *StartSecurityCheckServiceResponse {
	s.Headers = v
	return s
}

func (s *StartSecurityCheckServiceResponse) SetStatusCode(v int32) *StartSecurityCheckServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSecurityCheckServiceResponse) SetBody(v *StartSecurityCheckServiceResponseBody) *StartSecurityCheckServiceResponse {
	s.Body = v
	return s
}

func (s *StartSecurityCheckServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
