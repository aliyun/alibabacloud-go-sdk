// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDomainResponse
	GetStatusCode() *int32
	SetBody(v *CreateDomainResponseBody) *CreateDomainResponse
	GetBody() *CreateDomainResponseBody
}

type CreateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDomainResponse) GetBody() *CreateDomainResponseBody {
	return s.Body
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

func (s *CreateDomainResponse) Validate() error {
	return dara.Validate(s)
}
