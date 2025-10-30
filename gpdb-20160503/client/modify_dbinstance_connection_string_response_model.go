// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBInstanceConnectionStringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBInstanceConnectionStringResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBInstanceConnectionStringResponseBody) *ModifyDBInstanceConnectionStringResponse
	GetBody() *ModifyDBInstanceConnectionStringResponseBody
}

type ModifyDBInstanceConnectionStringResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBInstanceConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBInstanceConnectionStringResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBInstanceConnectionStringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBInstanceConnectionStringResponse) GetBody() *ModifyDBInstanceConnectionStringResponseBody {
	return s.Body
}

func (s *ModifyDBInstanceConnectionStringResponse) SetHeaders(v map[string]*string) *ModifyDBInstanceConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetStatusCode(v int32) *ModifyDBInstanceConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) SetBody(v *ModifyDBInstanceConnectionStringResponseBody) *ModifyDBInstanceConnectionStringResponse {
	s.Body = v
	return s
}

func (s *ModifyDBInstanceConnectionStringResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
