// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDBInstanceReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDBInstanceReplicationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDBInstanceReplicationResponseBody) *UpdateDBInstanceReplicationResponse
	GetBody() *UpdateDBInstanceReplicationResponseBody
}

type UpdateDBInstanceReplicationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDBInstanceReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDBInstanceReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceReplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDBInstanceReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDBInstanceReplicationResponse) GetBody() *UpdateDBInstanceReplicationResponseBody {
	return s.Body
}

func (s *UpdateDBInstanceReplicationResponse) SetHeaders(v map[string]*string) *UpdateDBInstanceReplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstanceReplicationResponse) SetStatusCode(v int32) *UpdateDBInstanceReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstanceReplicationResponse) SetBody(v *UpdateDBInstanceReplicationResponseBody) *UpdateDBInstanceReplicationResponse {
	s.Body = v
	return s
}

func (s *UpdateDBInstanceReplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
