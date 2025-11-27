// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTempDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTempDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateTempDBInstanceResponseBody) *CreateTempDBInstanceResponse
	GetBody() *CreateTempDBInstanceResponseBody
}

type CreateTempDBInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTempDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTempDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTempDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateTempDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTempDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTempDBInstanceResponse) GetBody() *CreateTempDBInstanceResponseBody {
	return s.Body
}

func (s *CreateTempDBInstanceResponse) SetHeaders(v map[string]*string) *CreateTempDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateTempDBInstanceResponse) SetStatusCode(v int32) *CreateTempDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTempDBInstanceResponse) SetBody(v *CreateTempDBInstanceResponseBody) *CreateTempDBInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateTempDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
