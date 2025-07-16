// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnFullDomainsBlockIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCdnFullDomainsBlockIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCdnFullDomainsBlockIPResponse
	GetStatusCode() *int32
	SetBody(v *SetCdnFullDomainsBlockIPResponseBody) *SetCdnFullDomainsBlockIPResponse
	GetBody() *SetCdnFullDomainsBlockIPResponseBody
}

type SetCdnFullDomainsBlockIPResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCdnFullDomainsBlockIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCdnFullDomainsBlockIPResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCdnFullDomainsBlockIPResponse) GoString() string {
	return s.String()
}

func (s *SetCdnFullDomainsBlockIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCdnFullDomainsBlockIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCdnFullDomainsBlockIPResponse) GetBody() *SetCdnFullDomainsBlockIPResponseBody {
	return s.Body
}

func (s *SetCdnFullDomainsBlockIPResponse) SetHeaders(v map[string]*string) *SetCdnFullDomainsBlockIPResponse {
	s.Headers = v
	return s
}

func (s *SetCdnFullDomainsBlockIPResponse) SetStatusCode(v int32) *SetCdnFullDomainsBlockIPResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPResponse) SetBody(v *SetCdnFullDomainsBlockIPResponseBody) *SetCdnFullDomainsBlockIPResponse {
	s.Body = v
	return s
}

func (s *SetCdnFullDomainsBlockIPResponse) Validate() error {
	return dara.Validate(s)
}
