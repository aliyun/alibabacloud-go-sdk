// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDcdnIpaDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDcdnIpaDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDcdnIpaDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddDcdnIpaDomainResponseBody) *AddDcdnIpaDomainResponse
	GetBody() *AddDcdnIpaDomainResponseBody
}

type AddDcdnIpaDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDcdnIpaDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDcdnIpaDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDcdnIpaDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDcdnIpaDomainResponse) GetBody() *AddDcdnIpaDomainResponseBody {
	return s.Body
}

func (s *AddDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *AddDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDcdnIpaDomainResponse) SetStatusCode(v int32) *AddDcdnIpaDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDcdnIpaDomainResponse) SetBody(v *AddDcdnIpaDomainResponseBody) *AddDcdnIpaDomainResponse {
	s.Body = v
	return s
}

func (s *AddDcdnIpaDomainResponse) Validate() error {
	return dara.Validate(s)
}
