// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainOfDnsProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeDomainOfDnsProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeDomainOfDnsProductResponse
	GetStatusCode() *int32
	SetBody(v *ChangeDomainOfDnsProductResponseBody) *ChangeDomainOfDnsProductResponse
	GetBody() *ChangeDomainOfDnsProductResponseBody
}

type ChangeDomainOfDnsProductResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDomainOfDnsProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDomainOfDnsProductResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainOfDnsProductResponse) GoString() string {
	return s.String()
}

func (s *ChangeDomainOfDnsProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeDomainOfDnsProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeDomainOfDnsProductResponse) GetBody() *ChangeDomainOfDnsProductResponseBody {
	return s.Body
}

func (s *ChangeDomainOfDnsProductResponse) SetHeaders(v map[string]*string) *ChangeDomainOfDnsProductResponse {
	s.Headers = v
	return s
}

func (s *ChangeDomainOfDnsProductResponse) SetStatusCode(v int32) *ChangeDomainOfDnsProductResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDomainOfDnsProductResponse) SetBody(v *ChangeDomainOfDnsProductResponseBody) *ChangeDomainOfDnsProductResponse {
	s.Body = v
	return s
}

func (s *ChangeDomainOfDnsProductResponse) Validate() error {
	return dara.Validate(s)
}
