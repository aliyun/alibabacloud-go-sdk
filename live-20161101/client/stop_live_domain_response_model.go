// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLiveDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLiveDomainResponse
	GetStatusCode() *int32
	SetBody(v *StopLiveDomainResponseBody) *StopLiveDomainResponse
	GetBody() *StopLiveDomainResponseBody
}

type StopLiveDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLiveDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLiveDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLiveDomainResponse) GoString() string {
	return s.String()
}

func (s *StopLiveDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLiveDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLiveDomainResponse) GetBody() *StopLiveDomainResponseBody {
	return s.Body
}

func (s *StopLiveDomainResponse) SetHeaders(v map[string]*string) *StopLiveDomainResponse {
	s.Headers = v
	return s
}

func (s *StopLiveDomainResponse) SetStatusCode(v int32) *StopLiveDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLiveDomainResponse) SetBody(v *StopLiveDomainResponseBody) *StopLiveDomainResponse {
	s.Body = v
	return s
}

func (s *StopLiveDomainResponse) Validate() error {
	return dara.Validate(s)
}
