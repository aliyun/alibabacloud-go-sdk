// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnIpaDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDcdnIpaDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDcdnIpaDomainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDcdnIpaDomainResponseBody) *UpdateDcdnIpaDomainResponse
	GetBody() *UpdateDcdnIpaDomainResponseBody
}

type UpdateDcdnIpaDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDcdnIpaDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDcdnIpaDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnIpaDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDcdnIpaDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDcdnIpaDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDcdnIpaDomainResponse) GetBody() *UpdateDcdnIpaDomainResponseBody {
	return s.Body
}

func (s *UpdateDcdnIpaDomainResponse) SetHeaders(v map[string]*string) *UpdateDcdnIpaDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDcdnIpaDomainResponse) SetStatusCode(v int32) *UpdateDcdnIpaDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDcdnIpaDomainResponse) SetBody(v *UpdateDcdnIpaDomainResponseBody) *UpdateDcdnIpaDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateDcdnIpaDomainResponse) Validate() error {
	return dara.Validate(s)
}
