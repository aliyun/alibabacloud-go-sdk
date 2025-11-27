// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DestroyDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DestroyDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DestroyDBInstanceResponseBody) *DestroyDBInstanceResponse
	GetBody() *DestroyDBInstanceResponseBody
}

type DestroyDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DestroyDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DestroyDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DestroyDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DestroyDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DestroyDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DestroyDBInstanceResponse) GetBody() *DestroyDBInstanceResponseBody {
	return s.Body
}

func (s *DestroyDBInstanceResponse) SetHeaders(v map[string]*string) *DestroyDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DestroyDBInstanceResponse) SetStatusCode(v int32) *DestroyDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DestroyDBInstanceResponse) SetBody(v *DestroyDBInstanceResponseBody) *DestroyDBInstanceResponse {
	s.Body = v
	return s
}

func (s *DestroyDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
