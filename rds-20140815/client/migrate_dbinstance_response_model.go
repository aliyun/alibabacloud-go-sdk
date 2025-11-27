// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *MigrateDBInstanceResponseBody) *MigrateDBInstanceResponse
	GetBody() *MigrateDBInstanceResponseBody
}

type MigrateDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *MigrateDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateDBInstanceResponse) GetBody() *MigrateDBInstanceResponseBody {
	return s.Body
}

func (s *MigrateDBInstanceResponse) SetHeaders(v map[string]*string) *MigrateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *MigrateDBInstanceResponse) SetStatusCode(v int32) *MigrateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateDBInstanceResponse) SetBody(v *MigrateDBInstanceResponseBody) *MigrateDBInstanceResponse {
	s.Body = v
	return s
}

func (s *MigrateDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
