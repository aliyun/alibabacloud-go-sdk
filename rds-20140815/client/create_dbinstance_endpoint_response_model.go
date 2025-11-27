// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstanceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstanceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstanceEndpointResponseBody) *CreateDBInstanceEndpointResponse
	GetBody() *CreateDBInstanceEndpointResponseBody
}

type CreateDBInstanceEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstanceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstanceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstanceEndpointResponse) GetBody() *CreateDBInstanceEndpointResponseBody {
	return s.Body
}

func (s *CreateDBInstanceEndpointResponse) SetHeaders(v map[string]*string) *CreateDBInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceEndpointResponse) SetStatusCode(v int32) *CreateDBInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceEndpointResponse) SetBody(v *CreateDBInstanceEndpointResponseBody) *CreateDBInstanceEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstanceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
