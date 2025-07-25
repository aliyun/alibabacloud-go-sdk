// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMainDomainNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMainDomainNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMainDomainNameResponse
	GetStatusCode() *int32
	SetBody(v *GetMainDomainNameResponseBody) *GetMainDomainNameResponse
	GetBody() *GetMainDomainNameResponseBody
}

type GetMainDomainNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMainDomainNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMainDomainNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMainDomainNameResponse) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMainDomainNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMainDomainNameResponse) GetBody() *GetMainDomainNameResponseBody {
	return s.Body
}

func (s *GetMainDomainNameResponse) SetHeaders(v map[string]*string) *GetMainDomainNameResponse {
	s.Headers = v
	return s
}

func (s *GetMainDomainNameResponse) SetStatusCode(v int32) *GetMainDomainNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMainDomainNameResponse) SetBody(v *GetMainDomainNameResponseBody) *GetMainDomainNameResponse {
	s.Body = v
	return s
}

func (s *GetMainDomainNameResponse) Validate() error {
	return dara.Validate(s)
}
