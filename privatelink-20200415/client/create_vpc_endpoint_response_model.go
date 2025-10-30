// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcEndpointResponseBody) *CreateVpcEndpointResponse
	GetBody() *CreateVpcEndpointResponseBody
}

type CreateVpcEndpointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcEndpointResponse) GetBody() *CreateVpcEndpointResponseBody {
	return s.Body
}

func (s *CreateVpcEndpointResponse) SetHeaders(v map[string]*string) *CreateVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcEndpointResponse) SetStatusCode(v int32) *CreateVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcEndpointResponse) SetBody(v *CreateVpcEndpointResponseBody) *CreateVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
