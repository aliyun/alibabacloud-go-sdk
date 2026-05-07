// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestoreDBInstanceResponseBody) *RestoreDBInstanceResponse
	GetBody() *RestoreDBInstanceResponseBody
}

type RestoreDBInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestoreDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreDBInstanceResponse) GetBody() *RestoreDBInstanceResponseBody {
	return s.Body
}

func (s *RestoreDBInstanceResponse) SetHeaders(v map[string]*string) *RestoreDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestoreDBInstanceResponse) SetStatusCode(v int32) *RestoreDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreDBInstanceResponse) SetBody(v *RestoreDBInstanceResponseBody) *RestoreDBInstanceResponse {
	s.Body = v
	return s
}

func (s *RestoreDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
