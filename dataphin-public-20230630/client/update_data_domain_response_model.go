// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataDomainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataDomainResponseBody) *UpdateDataDomainResponse
	GetBody() *UpdateDataDomainResponseBody
}

type UpdateDataDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataDomainResponse) GetBody() *UpdateDataDomainResponseBody {
	return s.Body
}

func (s *UpdateDataDomainResponse) SetHeaders(v map[string]*string) *UpdateDataDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataDomainResponse) SetStatusCode(v int32) *UpdateDataDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataDomainResponse) SetBody(v *UpdateDataDomainResponseBody) *UpdateDataDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateDataDomainResponse) Validate() error {
	return dara.Validate(s)
}
