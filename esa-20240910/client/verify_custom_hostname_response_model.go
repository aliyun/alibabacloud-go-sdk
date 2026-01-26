// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCustomHostnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCustomHostnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCustomHostnameResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCustomHostnameResponseBody) *VerifyCustomHostnameResponse
	GetBody() *VerifyCustomHostnameResponseBody
}

type VerifyCustomHostnameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCustomHostnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCustomHostnameResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCustomHostnameResponse) GoString() string {
	return s.String()
}

func (s *VerifyCustomHostnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCustomHostnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCustomHostnameResponse) GetBody() *VerifyCustomHostnameResponseBody {
	return s.Body
}

func (s *VerifyCustomHostnameResponse) SetHeaders(v map[string]*string) *VerifyCustomHostnameResponse {
	s.Headers = v
	return s
}

func (s *VerifyCustomHostnameResponse) SetStatusCode(v int32) *VerifyCustomHostnameResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCustomHostnameResponse) SetBody(v *VerifyCustomHostnameResponseBody) *VerifyCustomHostnameResponse {
	s.Body = v
	return s
}

func (s *VerifyCustomHostnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
