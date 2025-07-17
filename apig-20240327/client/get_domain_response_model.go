// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDomainResponse
	GetStatusCode() *int32
	SetBody(v *GetDomainResponseBody) *GetDomainResponse
	GetBody() *GetDomainResponseBody
}

type GetDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDomainResponse) GetBody() *GetDomainResponseBody {
	return s.Body
}

func (s *GetDomainResponse) SetHeaders(v map[string]*string) *GetDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDomainResponse) SetStatusCode(v int32) *GetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainResponse) SetBody(v *GetDomainResponseBody) *GetDomainResponse {
	s.Body = v
	return s
}

func (s *GetDomainResponse) Validate() error {
	return dara.Validate(s)
}
