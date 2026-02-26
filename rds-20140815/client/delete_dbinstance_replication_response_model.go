// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBInstanceReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBInstanceReplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBInstanceReplicationResponseBody) *DeleteDBInstanceReplicationResponse
	GetBody() *DeleteDBInstanceReplicationResponseBody
}

type DeleteDBInstanceReplicationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBInstanceReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBInstanceReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceReplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBInstanceReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBInstanceReplicationResponse) GetBody() *DeleteDBInstanceReplicationResponseBody {
	return s.Body
}

func (s *DeleteDBInstanceReplicationResponse) SetHeaders(v map[string]*string) *DeleteDBInstanceReplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBInstanceReplicationResponse) SetStatusCode(v int32) *DeleteDBInstanceReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBInstanceReplicationResponse) SetBody(v *DeleteDBInstanceReplicationResponseBody) *DeleteDBInstanceReplicationResponse {
	s.Body = v
	return s
}

func (s *DeleteDBInstanceReplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
