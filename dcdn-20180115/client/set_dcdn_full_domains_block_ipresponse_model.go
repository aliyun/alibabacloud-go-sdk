// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnFullDomainsBlockIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDcdnFullDomainsBlockIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDcdnFullDomainsBlockIPResponse
	GetStatusCode() *int32
	SetBody(v *SetDcdnFullDomainsBlockIPResponseBody) *SetDcdnFullDomainsBlockIPResponse
	GetBody() *SetDcdnFullDomainsBlockIPResponseBody
}

type SetDcdnFullDomainsBlockIPResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDcdnFullDomainsBlockIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDcdnFullDomainsBlockIPResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnFullDomainsBlockIPResponse) GoString() string {
	return s.String()
}

func (s *SetDcdnFullDomainsBlockIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDcdnFullDomainsBlockIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDcdnFullDomainsBlockIPResponse) GetBody() *SetDcdnFullDomainsBlockIPResponseBody {
	return s.Body
}

func (s *SetDcdnFullDomainsBlockIPResponse) SetHeaders(v map[string]*string) *SetDcdnFullDomainsBlockIPResponse {
	s.Headers = v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponse) SetStatusCode(v int32) *SetDcdnFullDomainsBlockIPResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponse) SetBody(v *SetDcdnFullDomainsBlockIPResponseBody) *SetDcdnFullDomainsBlockIPResponse {
	s.Body = v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponse) Validate() error {
	return dara.Validate(s)
}
