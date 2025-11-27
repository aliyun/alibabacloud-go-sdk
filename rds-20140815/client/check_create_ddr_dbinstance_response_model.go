// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateDdrDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCreateDdrDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCreateDdrDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CheckCreateDdrDBInstanceResponseBody) *CheckCreateDdrDBInstanceResponse
	GetBody() *CheckCreateDdrDBInstanceResponseBody
}

type CheckCreateDdrDBInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCreateDdrDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCreateDdrDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateDdrDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CheckCreateDdrDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCreateDdrDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCreateDdrDBInstanceResponse) GetBody() *CheckCreateDdrDBInstanceResponseBody {
	return s.Body
}

func (s *CheckCreateDdrDBInstanceResponse) SetHeaders(v map[string]*string) *CheckCreateDdrDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CheckCreateDdrDBInstanceResponse) SetStatusCode(v int32) *CheckCreateDdrDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCreateDdrDBInstanceResponse) SetBody(v *CheckCreateDdrDBInstanceResponseBody) *CheckCreateDdrDBInstanceResponse {
	s.Body = v
	return s
}

func (s *CheckCreateDdrDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
