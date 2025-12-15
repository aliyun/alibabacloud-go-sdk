// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockDBInstanceWriteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LockDBInstanceWriteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LockDBInstanceWriteResponse
	GetStatusCode() *int32
	SetBody(v *LockDBInstanceWriteResponseBody) *LockDBInstanceWriteResponse
	GetBody() *LockDBInstanceWriteResponseBody
}

type LockDBInstanceWriteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockDBInstanceWriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockDBInstanceWriteResponse) String() string {
	return dara.Prettify(s)
}

func (s LockDBInstanceWriteResponse) GoString() string {
	return s.String()
}

func (s *LockDBInstanceWriteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LockDBInstanceWriteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LockDBInstanceWriteResponse) GetBody() *LockDBInstanceWriteResponseBody {
	return s.Body
}

func (s *LockDBInstanceWriteResponse) SetHeaders(v map[string]*string) *LockDBInstanceWriteResponse {
	s.Headers = v
	return s
}

func (s *LockDBInstanceWriteResponse) SetStatusCode(v int32) *LockDBInstanceWriteResponse {
	s.StatusCode = &v
	return s
}

func (s *LockDBInstanceWriteResponse) SetBody(v *LockDBInstanceWriteResponseBody) *LockDBInstanceWriteResponse {
	s.Body = v
	return s
}

func (s *LockDBInstanceWriteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
