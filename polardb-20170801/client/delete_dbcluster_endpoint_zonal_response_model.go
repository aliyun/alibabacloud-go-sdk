// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterEndpointZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBClusterEndpointZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBClusterEndpointZonalResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBClusterEndpointZonalResponseBody) *DeleteDBClusterEndpointZonalResponse
	GetBody() *DeleteDBClusterEndpointZonalResponseBody
}

type DeleteDBClusterEndpointZonalResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBClusterEndpointZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBClusterEndpointZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterEndpointZonalResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBClusterEndpointZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBClusterEndpointZonalResponse) GetBody() *DeleteDBClusterEndpointZonalResponseBody {
	return s.Body
}

func (s *DeleteDBClusterEndpointZonalResponse) SetHeaders(v map[string]*string) *DeleteDBClusterEndpointZonalResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterEndpointZonalResponse) SetStatusCode(v int32) *DeleteDBClusterEndpointZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBClusterEndpointZonalResponse) SetBody(v *DeleteDBClusterEndpointZonalResponseBody) *DeleteDBClusterEndpointZonalResponse {
	s.Body = v
	return s
}

func (s *DeleteDBClusterEndpointZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
