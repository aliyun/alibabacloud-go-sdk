// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartDBInstanceResponseBody) *StartDBInstanceResponse
	GetBody() *StartDBInstanceResponseBody
}

type StartDBInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDBInstanceResponse) GetBody() *StartDBInstanceResponseBody {
	return s.Body
}

func (s *StartDBInstanceResponse) SetHeaders(v map[string]*string) *StartDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartDBInstanceResponse) SetStatusCode(v int32) *StartDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDBInstanceResponse) SetBody(v *StartDBInstanceResponseBody) *StartDBInstanceResponse {
	s.Body = v
	return s
}

func (s *StartDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
