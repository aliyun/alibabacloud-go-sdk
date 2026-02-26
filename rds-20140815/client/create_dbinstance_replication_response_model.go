// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstanceReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstanceReplicationResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstanceReplicationResponseBody) *CreateDBInstanceReplicationResponse
	GetBody() *CreateDBInstanceReplicationResponseBody
}

type CreateDBInstanceReplicationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstanceReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceReplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstanceReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstanceReplicationResponse) GetBody() *CreateDBInstanceReplicationResponseBody {
	return s.Body
}

func (s *CreateDBInstanceReplicationResponse) SetHeaders(v map[string]*string) *CreateDBInstanceReplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceReplicationResponse) SetStatusCode(v int32) *CreateDBInstanceReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceReplicationResponse) SetBody(v *CreateDBInstanceReplicationResponseBody) *CreateDBInstanceReplicationResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstanceReplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
