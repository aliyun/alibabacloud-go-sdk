// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddCdnDomainResponseBody) *AddCdnDomainResponse
	GetBody() *AddCdnDomainResponseBody
}

type AddCdnDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCdnDomainResponse) GoString() string {
	return s.String()
}

func (s *AddCdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCdnDomainResponse) GetBody() *AddCdnDomainResponseBody {
	return s.Body
}

func (s *AddCdnDomainResponse) SetHeaders(v map[string]*string) *AddCdnDomainResponse {
	s.Headers = v
	return s
}

func (s *AddCdnDomainResponse) SetStatusCode(v int32) *AddCdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCdnDomainResponse) SetBody(v *AddCdnDomainResponseBody) *AddCdnDomainResponse {
	s.Body = v
	return s
}

func (s *AddCdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
