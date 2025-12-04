// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortsForClickHouseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePortsForClickHouseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePortsForClickHouseResponse
	GetStatusCode() *int32
	SetBody(v *CreatePortsForClickHouseResponseBody) *CreatePortsForClickHouseResponse
	GetBody() *CreatePortsForClickHouseResponseBody
}

type CreatePortsForClickHouseResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePortsForClickHouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePortsForClickHouseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePortsForClickHouseResponse) GoString() string {
	return s.String()
}

func (s *CreatePortsForClickHouseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePortsForClickHouseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePortsForClickHouseResponse) GetBody() *CreatePortsForClickHouseResponseBody {
	return s.Body
}

func (s *CreatePortsForClickHouseResponse) SetHeaders(v map[string]*string) *CreatePortsForClickHouseResponse {
	s.Headers = v
	return s
}

func (s *CreatePortsForClickHouseResponse) SetStatusCode(v int32) *CreatePortsForClickHouseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePortsForClickHouseResponse) SetBody(v *CreatePortsForClickHouseResponseBody) *CreatePortsForClickHouseResponse {
	s.Body = v
	return s
}

func (s *CreatePortsForClickHouseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
