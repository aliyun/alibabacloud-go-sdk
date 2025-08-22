// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDcdnDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDcdnDomainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDcdnDomainResponseBody) *UpdateDcdnDomainResponse
	GetBody() *UpdateDcdnDomainResponseBody
}

type UpdateDcdnDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDcdnDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDcdnDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDcdnDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDcdnDomainResponse) GetBody() *UpdateDcdnDomainResponseBody {
	return s.Body
}

func (s *UpdateDcdnDomainResponse) SetHeaders(v map[string]*string) *UpdateDcdnDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnDomainResponse) SetStatusCode(v int32) *UpdateDcdnDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDcdnDomainResponse) SetBody(v *UpdateDcdnDomainResponseBody) *UpdateDcdnDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateDcdnDomainResponse) Validate() error {
	return dara.Validate(s)
}
