// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayDomainResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayDomainResponseBody) *ListGatewayDomainResponse
	GetBody() *ListGatewayDomainResponseBody
}

type ListGatewayDomainResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayDomainResponse) GetBody() *ListGatewayDomainResponseBody {
	return s.Body
}

func (s *ListGatewayDomainResponse) SetHeaders(v map[string]*string) *ListGatewayDomainResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayDomainResponse) SetStatusCode(v int32) *ListGatewayDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayDomainResponse) SetBody(v *ListGatewayDomainResponseBody) *ListGatewayDomainResponse {
	s.Body = v
	return s
}

func (s *ListGatewayDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
