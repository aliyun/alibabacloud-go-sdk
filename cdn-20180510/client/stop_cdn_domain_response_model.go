// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *StopCdnDomainResponseBody) *StopCdnDomainResponse
	GetBody() *StopCdnDomainResponseBody
}

type StopCdnDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *StopCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCdnDomainResponse) GetBody() *StopCdnDomainResponseBody {
	return s.Body
}

func (s *StopCdnDomainResponse) SetHeaders(v map[string]*string) *StopCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *StopCdnDomainResponse) SetStatusCode(v int32) *StopCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCdnDomainResponse) SetBody(v *StopCdnDomainResponseBody) *StopCdnDomainResponse {
	s.Body = v
	return s
}

func (s *StopCdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
