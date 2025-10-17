// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterEndpointZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBClusterEndpointZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBClusterEndpointZonalResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBClusterEndpointZonalResponseBody) *CreateDBClusterEndpointZonalResponse
	GetBody() *CreateDBClusterEndpointZonalResponseBody
}

type CreateDBClusterEndpointZonalResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBClusterEndpointZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBClusterEndpointZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterEndpointZonalResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBClusterEndpointZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBClusterEndpointZonalResponse) GetBody() *CreateDBClusterEndpointZonalResponseBody {
	return s.Body
}

func (s *CreateDBClusterEndpointZonalResponse) SetHeaders(v map[string]*string) *CreateDBClusterEndpointZonalResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterEndpointZonalResponse) SetStatusCode(v int32) *CreateDBClusterEndpointZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBClusterEndpointZonalResponse) SetBody(v *CreateDBClusterEndpointZonalResponseBody) *CreateDBClusterEndpointZonalResponse {
	s.Body = v
	return s
}

func (s *CreateDBClusterEndpointZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
