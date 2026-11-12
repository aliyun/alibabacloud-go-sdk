// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDatabaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDatabaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDatabaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartDatabaseInstanceResponseBody) *StartDatabaseInstanceResponse
	GetBody() *StartDatabaseInstanceResponseBody
}

type StartDatabaseInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDatabaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDatabaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDatabaseInstanceResponse) GetBody() *StartDatabaseInstanceResponseBody {
	return s.Body
}

func (s *StartDatabaseInstanceResponse) SetHeaders(v map[string]*string) *StartDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartDatabaseInstanceResponse) SetStatusCode(v int32) *StartDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDatabaseInstanceResponse) SetBody(v *StartDatabaseInstanceResponseBody) *StartDatabaseInstanceResponse {
	s.Body = v
	return s
}

func (s *StartDatabaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
