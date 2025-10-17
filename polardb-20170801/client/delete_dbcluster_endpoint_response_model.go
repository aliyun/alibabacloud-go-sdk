// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBClusterEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBClusterEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBClusterEndpointResponseBody) *DeleteDBClusterEndpointResponse
	GetBody() *DeleteDBClusterEndpointResponseBody
}

type DeleteDBClusterEndpointResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBClusterEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBClusterEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBClusterEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBClusterEndpointResponse) GetBody() *DeleteDBClusterEndpointResponseBody {
	return s.Body
}

func (s *DeleteDBClusterEndpointResponse) SetHeaders(v map[string]*string) *DeleteDBClusterEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterEndpointResponse) SetStatusCode(v int32) *DeleteDBClusterEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBClusterEndpointResponse) SetBody(v *DeleteDBClusterEndpointResponseBody) *DeleteDBClusterEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteDBClusterEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
