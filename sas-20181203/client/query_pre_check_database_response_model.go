// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPreCheckDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPreCheckDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPreCheckDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *QueryPreCheckDatabaseResponseBody) *QueryPreCheckDatabaseResponse
	GetBody() *QueryPreCheckDatabaseResponseBody
}

type QueryPreCheckDatabaseResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPreCheckDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPreCheckDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPreCheckDatabaseResponse) GoString() string {
	return s.String()
}

func (s *QueryPreCheckDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPreCheckDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPreCheckDatabaseResponse) GetBody() *QueryPreCheckDatabaseResponseBody {
	return s.Body
}

func (s *QueryPreCheckDatabaseResponse) SetHeaders(v map[string]*string) *QueryPreCheckDatabaseResponse {
	s.Headers = v
	return s
}

func (s *QueryPreCheckDatabaseResponse) SetStatusCode(v int32) *QueryPreCheckDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPreCheckDatabaseResponse) SetBody(v *QueryPreCheckDatabaseResponseBody) *QueryPreCheckDatabaseResponse {
	s.Body = v
	return s
}

func (s *QueryPreCheckDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
