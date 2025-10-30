// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcEndpointAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpcEndpointAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpcEndpointAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetVpcEndpointAttributeResponseBody) *GetVpcEndpointAttributeResponse
	GetBody() *GetVpcEndpointAttributeResponseBody
}

type GetVpcEndpointAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpcEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpcEndpointAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpcEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpcEndpointAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpcEndpointAttributeResponse) GetBody() *GetVpcEndpointAttributeResponseBody {
	return s.Body
}

func (s *GetVpcEndpointAttributeResponse) SetHeaders(v map[string]*string) *GetVpcEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetVpcEndpointAttributeResponse) SetStatusCode(v int32) *GetVpcEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpcEndpointAttributeResponse) SetBody(v *GetVpcEndpointAttributeResponseBody) *GetVpcEndpointAttributeResponse {
	s.Body = v
	return s
}

func (s *GetVpcEndpointAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
