// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoveryDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoveryDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RecoveryDBInstanceResponseBody) *RecoveryDBInstanceResponse
	GetBody() *RecoveryDBInstanceResponseBody
}

type RecoveryDBInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoveryDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoveryDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoveryDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RecoveryDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoveryDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoveryDBInstanceResponse) GetBody() *RecoveryDBInstanceResponseBody {
	return s.Body
}

func (s *RecoveryDBInstanceResponse) SetHeaders(v map[string]*string) *RecoveryDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RecoveryDBInstanceResponse) SetStatusCode(v int32) *RecoveryDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoveryDBInstanceResponse) SetBody(v *RecoveryDBInstanceResponseBody) *RecoveryDBInstanceResponse {
	s.Body = v
	return s
}

func (s *RecoveryDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
