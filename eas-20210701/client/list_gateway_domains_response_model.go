// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayDomainsResponseBody) *ListGatewayDomainsResponse
	GetBody() *ListGatewayDomainsResponseBody
}

type ListGatewayDomainsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayDomainsResponse) GetBody() *ListGatewayDomainsResponseBody {
	return s.Body
}

func (s *ListGatewayDomainsResponse) SetHeaders(v map[string]*string) *ListGatewayDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayDomainsResponse) SetStatusCode(v int32) *ListGatewayDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayDomainsResponse) SetBody(v *ListGatewayDomainsResponseBody) *ListGatewayDomainsResponse {
	s.Body = v
	return s
}

func (s *ListGatewayDomainsResponse) Validate() error {
	return dara.Validate(s)
}
