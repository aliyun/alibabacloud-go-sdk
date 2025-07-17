// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInspectProxyAccessSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InspectProxyAccessSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InspectProxyAccessSecretResponse
	GetStatusCode() *int32
	SetBody(v *InspectProxyAccessSecretResponseBody) *InspectProxyAccessSecretResponse
	GetBody() *InspectProxyAccessSecretResponseBody
}

type InspectProxyAccessSecretResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InspectProxyAccessSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InspectProxyAccessSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s InspectProxyAccessSecretResponse) GoString() string {
	return s.String()
}

func (s *InspectProxyAccessSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InspectProxyAccessSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InspectProxyAccessSecretResponse) GetBody() *InspectProxyAccessSecretResponseBody {
	return s.Body
}

func (s *InspectProxyAccessSecretResponse) SetHeaders(v map[string]*string) *InspectProxyAccessSecretResponse {
	s.Headers = v
	return s
}

func (s *InspectProxyAccessSecretResponse) SetStatusCode(v int32) *InspectProxyAccessSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *InspectProxyAccessSecretResponse) SetBody(v *InspectProxyAccessSecretResponseBody) *InspectProxyAccessSecretResponse {
	s.Body = v
	return s
}

func (s *InspectProxyAccessSecretResponse) Validate() error {
	return dara.Validate(s)
}
