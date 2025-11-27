// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReceiveDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReceiveDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReceiveDBInstanceResponseBody) *ReceiveDBInstanceResponse
	GetBody() *ReceiveDBInstanceResponseBody
}

type ReceiveDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReceiveDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReceiveDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReceiveDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReceiveDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReceiveDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReceiveDBInstanceResponse) GetBody() *ReceiveDBInstanceResponseBody {
	return s.Body
}

func (s *ReceiveDBInstanceResponse) SetHeaders(v map[string]*string) *ReceiveDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReceiveDBInstanceResponse) SetStatusCode(v int32) *ReceiveDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReceiveDBInstanceResponse) SetBody(v *ReceiveDBInstanceResponseBody) *ReceiveDBInstanceResponse {
	s.Body = v
	return s
}

func (s *ReceiveDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
