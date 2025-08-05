// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDnsDomainNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrivateDnsDomainNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrivateDnsDomainNameResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrivateDnsDomainNameResponseBody) *DeletePrivateDnsDomainNameResponse
	GetBody() *DeletePrivateDnsDomainNameResponseBody
}

type DeletePrivateDnsDomainNameResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateDnsDomainNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateDnsDomainNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDnsDomainNameResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateDnsDomainNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrivateDnsDomainNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrivateDnsDomainNameResponse) GetBody() *DeletePrivateDnsDomainNameResponseBody {
	return s.Body
}

func (s *DeletePrivateDnsDomainNameResponse) SetHeaders(v map[string]*string) *DeletePrivateDnsDomainNameResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateDnsDomainNameResponse) SetStatusCode(v int32) *DeletePrivateDnsDomainNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateDnsDomainNameResponse) SetBody(v *DeletePrivateDnsDomainNameResponseBody) *DeletePrivateDnsDomainNameResponse {
	s.Body = v
	return s
}

func (s *DeletePrivateDnsDomainNameResponse) Validate() error {
	return dara.Validate(s)
}
