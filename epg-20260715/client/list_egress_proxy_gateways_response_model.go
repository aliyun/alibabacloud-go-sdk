// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEgressProxyGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEgressProxyGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *ListEgressProxyGatewaysResponseBody) *ListEgressProxyGatewaysResponse
	GetBody() *ListEgressProxyGatewaysResponseBody
}

type ListEgressProxyGatewaysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEgressProxyGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEgressProxyGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyGatewaysResponse) GoString() string {
	return s.String()
}

func (s *ListEgressProxyGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEgressProxyGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEgressProxyGatewaysResponse) GetBody() *ListEgressProxyGatewaysResponseBody {
	return s.Body
}

func (s *ListEgressProxyGatewaysResponse) SetHeaders(v map[string]*string) *ListEgressProxyGatewaysResponse {
	s.Headers = v
	return s
}

func (s *ListEgressProxyGatewaysResponse) SetStatusCode(v int32) *ListEgressProxyGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEgressProxyGatewaysResponse) SetBody(v *ListEgressProxyGatewaysResponseBody) *ListEgressProxyGatewaysResponse {
	s.Body = v
	return s
}

func (s *ListEgressProxyGatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
