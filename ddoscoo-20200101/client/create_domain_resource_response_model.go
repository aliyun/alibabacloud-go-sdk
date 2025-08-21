// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDomainResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDomainResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDomainResourceResponseBody) *CreateDomainResourceResponse
	GetBody() *CreateDomainResourceResponseBody
}

type CreateDomainResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDomainResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDomainResourceResponse) GetBody() *CreateDomainResourceResponseBody {
	return s.Body
}

func (s *CreateDomainResourceResponse) SetHeaders(v map[string]*string) *CreateDomainResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResourceResponse) SetStatusCode(v int32) *CreateDomainResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResourceResponse) SetBody(v *CreateDomainResourceResponseBody) *CreateDomainResourceResponse {
	s.Body = v
	return s
}

func (s *CreateDomainResourceResponse) Validate() error {
	return dara.Validate(s)
}
