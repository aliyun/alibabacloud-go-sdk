// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevealServiceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevealServiceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevealServiceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *RevealServiceCredentialResponseBody) *RevealServiceCredentialResponse
	GetBody() *RevealServiceCredentialResponseBody
}

type RevealServiceCredentialResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevealServiceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevealServiceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s RevealServiceCredentialResponse) GoString() string {
	return s.String()
}

func (s *RevealServiceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevealServiceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevealServiceCredentialResponse) GetBody() *RevealServiceCredentialResponseBody {
	return s.Body
}

func (s *RevealServiceCredentialResponse) SetHeaders(v map[string]*string) *RevealServiceCredentialResponse {
	s.Headers = v
	return s
}

func (s *RevealServiceCredentialResponse) SetStatusCode(v int32) *RevealServiceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *RevealServiceCredentialResponse) SetBody(v *RevealServiceCredentialResponseBody) *RevealServiceCredentialResponse {
	s.Body = v
	return s
}

func (s *RevealServiceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
