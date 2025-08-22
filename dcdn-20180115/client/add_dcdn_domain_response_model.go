// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddDcdnDomainResponseBody) *AddDcdnDomainResponse
	GetBody() *AddDcdnDomainResponseBody
}

type AddDcdnDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDcdnDomainResponse) GetBody() *AddDcdnDomainResponseBody {
	return s.Body
}

func (s *AddDcdnDomainResponse) SetHeaders(v map[string]*string) *AddDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDcdnDomainResponse) SetStatusCode(v int32) *AddDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDcdnDomainResponse) SetBody(v *AddDcdnDomainResponseBody) *AddDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *AddDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
