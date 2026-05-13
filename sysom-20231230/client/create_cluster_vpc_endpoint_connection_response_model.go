// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterVpcEndpointConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClusterVpcEndpointConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClusterVpcEndpointConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateClusterVpcEndpointConnectionResponseBody) *CreateClusterVpcEndpointConnectionResponse
	GetBody() *CreateClusterVpcEndpointConnectionResponseBody
}

type CreateClusterVpcEndpointConnectionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterVpcEndpointConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterVpcEndpointConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterVpcEndpointConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterVpcEndpointConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClusterVpcEndpointConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClusterVpcEndpointConnectionResponse) GetBody() *CreateClusterVpcEndpointConnectionResponseBody {
	return s.Body
}

func (s *CreateClusterVpcEndpointConnectionResponse) SetHeaders(v map[string]*string) *CreateClusterVpcEndpointConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponse) SetStatusCode(v int32) *CreateClusterVpcEndpointConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponse) SetBody(v *CreateClusterVpcEndpointConnectionResponseBody) *CreateClusterVpcEndpointConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateClusterVpcEndpointConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
