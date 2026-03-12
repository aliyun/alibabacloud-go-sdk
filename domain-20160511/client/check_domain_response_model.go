// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDomainResponse
	GetStatusCode() *int32
	SetBody(v *CheckDomainResponseBody) *CheckDomainResponse
	GetBody() *CheckDomainResponseBody
}

type CheckDomainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDomainResponse) GetBody() *CheckDomainResponseBody {
	return s.Body
}

func (s *CheckDomainResponse) SetHeaders(v map[string]*string) *CheckDomainResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainResponse) SetStatusCode(v int32) *CheckDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainResponse) SetBody(v *CheckDomainResponseBody) *CheckDomainResponse {
	s.Body = v
	return s
}

func (s *CheckDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
