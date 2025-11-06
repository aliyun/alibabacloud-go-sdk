// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayDomainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayDomainResponseBody) *UpdateGatewayDomainResponse
	GetBody() *UpdateGatewayDomainResponseBody
}

type UpdateGatewayDomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayDomainResponse) GetBody() *UpdateGatewayDomainResponseBody {
	return s.Body
}

func (s *UpdateGatewayDomainResponse) SetHeaders(v map[string]*string) *UpdateGatewayDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayDomainResponse) SetStatusCode(v int32) *UpdateGatewayDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayDomainResponse) SetBody(v *UpdateGatewayDomainResponseBody) *UpdateGatewayDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
