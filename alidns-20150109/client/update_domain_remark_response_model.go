// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainRemarkResponseBody) *UpdateDomainRemarkResponse
	GetBody() *UpdateDomainRemarkResponseBody
}

type UpdateDomainRemarkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainRemarkResponse) GetBody() *UpdateDomainRemarkResponseBody {
	return s.Body
}

func (s *UpdateDomainRemarkResponse) SetHeaders(v map[string]*string) *UpdateDomainRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainRemarkResponse) SetStatusCode(v int32) *UpdateDomainRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainRemarkResponse) SetBody(v *UpdateDomainRemarkResponseBody) *UpdateDomainRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainRemarkResponse) Validate() error {
	return dara.Validate(s)
}
