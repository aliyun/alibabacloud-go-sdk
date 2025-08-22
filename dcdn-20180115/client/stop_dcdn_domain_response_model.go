// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *StopDcdnDomainResponseBody) *StopDcdnDomainResponse
	GetBody() *StopDcdnDomainResponseBody
}

type StopDcdnDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StopDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDcdnDomainResponse) GetBody() *StopDcdnDomainResponseBody {
	return s.Body
}

func (s *StopDcdnDomainResponse) SetHeaders(v map[string]*string) *StopDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StopDcdnDomainResponse) SetStatusCode(v int32) *StopDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDcdnDomainResponse) SetBody(v *StopDcdnDomainResponseBody) *StopDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *StopDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
