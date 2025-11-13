// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCreateDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCreateDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CheckCreateDBInstanceResponseBody) *CheckCreateDBInstanceResponse
	GetBody() *CheckCreateDBInstanceResponseBody
}

type CheckCreateDBInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCreateDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CheckCreateDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCreateDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCreateDBInstanceResponse) GetBody() *CheckCreateDBInstanceResponseBody {
	return s.Body
}

func (s *CheckCreateDBInstanceResponse) SetHeaders(v map[string]*string) *CheckCreateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CheckCreateDBInstanceResponse) SetStatusCode(v int32) *CheckCreateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCreateDBInstanceResponse) SetBody(v *CheckCreateDBInstanceResponseBody) *CheckCreateDBInstanceResponse {
	s.Body = v
	return s
}

func (s *CheckCreateDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
