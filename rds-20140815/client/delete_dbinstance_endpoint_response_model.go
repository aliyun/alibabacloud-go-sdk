// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBInstanceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBInstanceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBInstanceEndpointResponseBody) *DeleteDBInstanceEndpointResponse
	GetBody() *DeleteDBInstanceEndpointResponseBody
}

type DeleteDBInstanceEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBInstanceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBInstanceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBInstanceEndpointResponse) GetBody() *DeleteDBInstanceEndpointResponseBody {
	return s.Body
}

func (s *DeleteDBInstanceEndpointResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceEndpointResponse) SetStatusCode(v int32) *DeleteDBInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceEndpointResponse) SetBody(v *DeleteDBInstanceEndpointResponseBody) *DeleteDBInstanceEndpointResponse {
	s.Body = v
	return s
}

func (s *DeleteDBInstanceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
