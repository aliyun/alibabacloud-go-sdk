// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBClusterEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBClusterEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBClusterEndpointResponseBody) *CreateDBClusterEndpointResponse
	GetBody() *CreateDBClusterEndpointResponseBody
}

type CreateDBClusterEndpointResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBClusterEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBClusterEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBClusterEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBClusterEndpointResponse) GetBody() *CreateDBClusterEndpointResponseBody {
	return s.Body
}

func (s *CreateDBClusterEndpointResponse) SetHeaders(v map[string]*string) *CreateDBClusterEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterEndpointResponse) SetStatusCode(v int32) *CreateDBClusterEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBClusterEndpointResponse) SetBody(v *CreateDBClusterEndpointResponseBody) *CreateDBClusterEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateDBClusterEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
