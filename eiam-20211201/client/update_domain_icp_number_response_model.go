// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainIcpNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDomainIcpNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDomainIcpNumberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDomainIcpNumberResponseBody) *UpdateDomainIcpNumberResponse
	GetBody() *UpdateDomainIcpNumberResponseBody
}

type UpdateDomainIcpNumberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDomainIcpNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainIcpNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainIcpNumberResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainIcpNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDomainIcpNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDomainIcpNumberResponse) GetBody() *UpdateDomainIcpNumberResponseBody {
	return s.Body
}

func (s *UpdateDomainIcpNumberResponse) SetHeaders(v map[string]*string) *UpdateDomainIcpNumberResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainIcpNumberResponse) SetStatusCode(v int32) *UpdateDomainIcpNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainIcpNumberResponse) SetBody(v *UpdateDomainIcpNumberResponseBody) *UpdateDomainIcpNumberResponse {
	s.Body = v
	return s
}

func (s *UpdateDomainIcpNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
