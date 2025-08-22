// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDcdnIpaDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDcdnIpaDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDcdnIpaDomainResponse
	GetStatusCode() *int32
	SetBody(v *StopDcdnIpaDomainResponseBody) *StopDcdnIpaDomainResponse
	GetBody() *StopDcdnIpaDomainResponseBody
}

type StopDcdnIpaDomainResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDcdnIpaDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *StopDcdnIpaDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDcdnIpaDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDcdnIpaDomainResponse) GetBody() *StopDcdnIpaDomainResponseBody {
	return s.Body
}

func (s *StopDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *StopDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *StopDcdnIpaDomainResponse) SetStatusCode(v int32) *StopDcdnIpaDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDcdnIpaDomainResponse) SetBody(v *StopDcdnIpaDomainResponseBody) *StopDcdnIpaDomainResponse {
	s.Body = v
	return s
}

func (s *StopDcdnIpaDomainResponse) Validate() error {
	return dara.Validate(s)
}
