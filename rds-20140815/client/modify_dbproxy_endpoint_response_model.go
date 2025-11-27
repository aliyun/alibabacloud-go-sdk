// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBProxyEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBProxyEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBProxyEndpointResponseBody) *ModifyDBProxyEndpointResponse
	GetBody() *ModifyDBProxyEndpointResponseBody
}

type ModifyDBProxyEndpointResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBProxyEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBProxyEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBProxyEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBProxyEndpointResponse) GetBody() *ModifyDBProxyEndpointResponseBody {
	return s.Body
}

func (s *ModifyDBProxyEndpointResponse) SetHeaders(v map[string]*string) *ModifyDBProxyEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBProxyEndpointResponse) SetStatusCode(v int32) *ModifyDBProxyEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBProxyEndpointResponse) SetBody(v *ModifyDBProxyEndpointResponseBody) *ModifyDBProxyEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyDBProxyEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
