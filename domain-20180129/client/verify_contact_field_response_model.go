// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyContactFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyContactFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyContactFieldResponse
	GetStatusCode() *int32
	SetBody(v *VerifyContactFieldResponseBody) *VerifyContactFieldResponse
	GetBody() *VerifyContactFieldResponseBody
}

type VerifyContactFieldResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyContactFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyContactFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyContactFieldResponse) GoString() string {
	return s.String()
}

func (s *VerifyContactFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyContactFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyContactFieldResponse) GetBody() *VerifyContactFieldResponseBody {
	return s.Body
}

func (s *VerifyContactFieldResponse) SetHeaders(v map[string]*string) *VerifyContactFieldResponse {
	s.Headers = v
	return s
}

func (s *VerifyContactFieldResponse) SetStatusCode(v int32) *VerifyContactFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyContactFieldResponse) SetBody(v *VerifyContactFieldResponseBody) *VerifyContactFieldResponse {
	s.Body = v
	return s
}

func (s *VerifyContactFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
