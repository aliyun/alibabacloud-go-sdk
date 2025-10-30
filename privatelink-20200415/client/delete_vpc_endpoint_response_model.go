// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcEndpointResponseBody) *DeleteVpcEndpointResponse
	GetBody() *DeleteVpcEndpointResponseBody
}

type DeleteVpcEndpointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcEndpointResponse) GetBody() *DeleteVpcEndpointResponseBody {
	return s.Body
}

func (s *DeleteVpcEndpointResponse) SetHeaders(v map[string]*string) *DeleteVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcEndpointResponse) SetStatusCode(v int32) *DeleteVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcEndpointResponse) SetBody(v *DeleteVpcEndpointResponseBody) *DeleteVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
