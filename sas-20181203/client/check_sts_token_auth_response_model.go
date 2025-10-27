// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckStsTokenAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckStsTokenAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckStsTokenAuthResponse
	GetStatusCode() *int32
	SetBody(v *CheckStsTokenAuthResponseBody) *CheckStsTokenAuthResponse
	GetBody() *CheckStsTokenAuthResponseBody
}

type CheckStsTokenAuthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckStsTokenAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckStsTokenAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckStsTokenAuthResponse) GoString() string {
	return s.String()
}

func (s *CheckStsTokenAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckStsTokenAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckStsTokenAuthResponse) GetBody() *CheckStsTokenAuthResponseBody {
	return s.Body
}

func (s *CheckStsTokenAuthResponse) SetHeaders(v map[string]*string) *CheckStsTokenAuthResponse {
	s.Headers = v
	return s
}

func (s *CheckStsTokenAuthResponse) SetStatusCode(v int32) *CheckStsTokenAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckStsTokenAuthResponse) SetBody(v *CheckStsTokenAuthResponseBody) *CheckStsTokenAuthResponse {
	s.Body = v
	return s
}

func (s *CheckStsTokenAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
