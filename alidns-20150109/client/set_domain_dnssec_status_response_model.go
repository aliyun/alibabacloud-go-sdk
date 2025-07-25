// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainDnssecStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDomainDnssecStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDomainDnssecStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDomainDnssecStatusResponseBody) *SetDomainDnssecStatusResponse
	GetBody() *SetDomainDnssecStatusResponseBody
}

type SetDomainDnssecStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainDnssecStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainDnssecStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDomainDnssecStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDomainDnssecStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDomainDnssecStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDomainDnssecStatusResponse) GetBody() *SetDomainDnssecStatusResponseBody {
	return s.Body
}

func (s *SetDomainDnssecStatusResponse) SetHeaders(v map[string]*string) *SetDomainDnssecStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDomainDnssecStatusResponse) SetStatusCode(v int32) *SetDomainDnssecStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainDnssecStatusResponse) SetBody(v *SetDomainDnssecStatusResponseBody) *SetDomainDnssecStatusResponse {
	s.Body = v
	return s
}

func (s *SetDomainDnssecStatusResponse) Validate() error {
	return dara.Validate(s)
}
