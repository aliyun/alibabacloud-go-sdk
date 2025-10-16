// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrivateDnsEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPrivateDnsEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPrivateDnsEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPrivateDnsEndpointResponseBody) *ModifyPrivateDnsEndpointResponse
	GetBody() *ModifyPrivateDnsEndpointResponseBody
}

type ModifyPrivateDnsEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPrivateDnsEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPrivateDnsEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrivateDnsEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyPrivateDnsEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPrivateDnsEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPrivateDnsEndpointResponse) GetBody() *ModifyPrivateDnsEndpointResponseBody {
	return s.Body
}

func (s *ModifyPrivateDnsEndpointResponse) SetHeaders(v map[string]*string) *ModifyPrivateDnsEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyPrivateDnsEndpointResponse) SetStatusCode(v int32) *ModifyPrivateDnsEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPrivateDnsEndpointResponse) SetBody(v *ModifyPrivateDnsEndpointResponseBody) *ModifyPrivateDnsEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyPrivateDnsEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
