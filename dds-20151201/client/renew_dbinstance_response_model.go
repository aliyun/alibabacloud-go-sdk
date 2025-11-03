// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewDBInstanceResponseBody) *RenewDBInstanceResponse
	GetBody() *RenewDBInstanceResponseBody
}

type RenewDBInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewDBInstanceResponse) GetBody() *RenewDBInstanceResponseBody {
	return s.Body
}

func (s *RenewDBInstanceResponse) SetHeaders(v map[string]*string) *RenewDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewDBInstanceResponse) SetStatusCode(v int32) *RenewDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewDBInstanceResponse) SetBody(v *RenewDBInstanceResponseBody) *RenewDBInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
