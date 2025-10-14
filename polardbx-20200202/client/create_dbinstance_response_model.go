// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse
	GetBody() *CreateDBInstanceResponseBody
}

type CreateDBInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBInstanceResponse) GetBody() *CreateDBInstanceResponseBody {
	return s.Body
}

func (s *CreateDBInstanceResponse) SetHeaders(v map[string]*string) *CreateDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDBInstanceResponse) SetStatusCode(v int32) *CreateDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBInstanceResponse) SetBody(v *CreateDBInstanceResponseBody) *CreateDBInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
