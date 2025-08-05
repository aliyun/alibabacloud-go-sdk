// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsAllDomainNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateDnsAllDomainNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateDnsAllDomainNameResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateDnsAllDomainNameResponseBody) *DeletePrivateDnsAllDomainNameResponse
	GetBody() *DeletePrivateDnsAllDomainNameResponseBody
}

type DeletePrivateDnsAllDomainNameResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateDnsAllDomainNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateDnsAllDomainNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsAllDomainNameResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsAllDomainNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateDnsAllDomainNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateDnsAllDomainNameResponse) GetBody() *DeletePrivateDnsAllDomainNameResponseBody {
	return s.Body
}

func (s *DeletePrivateDnsAllDomainNameResponse) SetHeaders(v map[string]*string) *DeletePrivateDnsAllDomainNameResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateDnsAllDomainNameResponse) SetStatusCode(v int32) *DeletePrivateDnsAllDomainNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateDnsAllDomainNameResponse) SetBody(v *DeletePrivateDnsAllDomainNameResponseBody) *DeletePrivateDnsAllDomainNameResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateDnsAllDomainNameResponse) Validate() error {
	return dara.Validate(s)
}
