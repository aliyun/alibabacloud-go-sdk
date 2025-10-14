// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDBInstanceTDEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDBInstanceTDEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDBInstanceTDEResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDBInstanceTDEResponseBody) *UpdateDBInstanceTDEResponse
	GetBody() *UpdateDBInstanceTDEResponseBody
}

type UpdateDBInstanceTDEResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDBInstanceTDEResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDBInstanceTDEResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDBInstanceTDEResponse) GoString() string {
	return s.String()
}

func (s *UpdateDBInstanceTDEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDBInstanceTDEResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDBInstanceTDEResponse) GetBody() *UpdateDBInstanceTDEResponseBody {
	return s.Body
}

func (s *UpdateDBInstanceTDEResponse) SetHeaders(v map[string]*string) *UpdateDBInstanceTDEResponse {
	s.Headers = v
	return s
}

func (s *UpdateDBInstanceTDEResponse) SetStatusCode(v int32) *UpdateDBInstanceTDEResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDBInstanceTDEResponse) SetBody(v *UpdateDBInstanceTDEResponseBody) *UpdateDBInstanceTDEResponse {
	s.Body = v
	return s
}

func (s *UpdateDBInstanceTDEResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
