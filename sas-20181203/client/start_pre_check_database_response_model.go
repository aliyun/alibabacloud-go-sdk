// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPreCheckDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPreCheckDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPreCheckDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *StartPreCheckDatabaseResponseBody) *StartPreCheckDatabaseResponse
	GetBody() *StartPreCheckDatabaseResponseBody
}

type StartPreCheckDatabaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPreCheckDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPreCheckDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPreCheckDatabaseResponse) GoString() string {
	return s.String()
}

func (s *StartPreCheckDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPreCheckDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPreCheckDatabaseResponse) GetBody() *StartPreCheckDatabaseResponseBody {
	return s.Body
}

func (s *StartPreCheckDatabaseResponse) SetHeaders(v map[string]*string) *StartPreCheckDatabaseResponse {
	s.Headers = v
	return s
}

func (s *StartPreCheckDatabaseResponse) SetStatusCode(v int32) *StartPreCheckDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPreCheckDatabaseResponse) SetBody(v *StartPreCheckDatabaseResponseBody) *StartPreCheckDatabaseResponse {
	s.Body = v
	return s
}

func (s *StartPreCheckDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
