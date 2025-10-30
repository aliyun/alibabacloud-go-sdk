// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcEndpointServiceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcEndpointServiceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcEndpointServiceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcEndpointServiceAttributeResponseBody) *GetVpcEndpointServiceAttributeResponse
	GetBody() *GetVpcEndpointServiceAttributeResponseBody
}

type GetVpcEndpointServiceAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcEndpointServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcEndpointServiceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcEndpointServiceAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointServiceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcEndpointServiceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcEndpointServiceAttributeResponse) GetBody() *GetVpcEndpointServiceAttributeResponseBody {
	return s.Body
}

func (s *GetVpcEndpointServiceAttributeResponse) SetHeaders(v map[string]*string) *GetVpcEndpointServiceAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponse) SetStatusCode(v int32) *GetVpcEndpointServiceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponse) SetBody(v *GetVpcEndpointServiceAttributeResponseBody) *GetVpcEndpointServiceAttributeResponse {
	s.Body = v
	return s
}

func (s *GetVpcEndpointServiceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
