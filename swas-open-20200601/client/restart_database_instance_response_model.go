// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDatabaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDatabaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDatabaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestartDatabaseInstanceResponseBody) *RestartDatabaseInstanceResponse
	GetBody() *RestartDatabaseInstanceResponseBody
}

type RestartDatabaseInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDatabaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDatabaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDatabaseInstanceResponse) GetBody() *RestartDatabaseInstanceResponseBody {
	return s.Body
}

func (s *RestartDatabaseInstanceResponse) SetHeaders(v map[string]*string) *RestartDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDatabaseInstanceResponse) SetStatusCode(v int32) *RestartDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDatabaseInstanceResponse) SetBody(v *RestartDatabaseInstanceResponseBody) *RestartDatabaseInstanceResponse {
	s.Body = v
	return s
}

func (s *RestartDatabaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
