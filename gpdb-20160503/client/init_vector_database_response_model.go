// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitVectorDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitVectorDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitVectorDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *InitVectorDatabaseResponseBody) *InitVectorDatabaseResponse
	GetBody() *InitVectorDatabaseResponseBody
}

type InitVectorDatabaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitVectorDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitVectorDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s InitVectorDatabaseResponse) GoString() string {
	return s.String()
}

func (s *InitVectorDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitVectorDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitVectorDatabaseResponse) GetBody() *InitVectorDatabaseResponseBody {
	return s.Body
}

func (s *InitVectorDatabaseResponse) SetHeaders(v map[string]*string) *InitVectorDatabaseResponse {
	s.Headers = v
	return s
}

func (s *InitVectorDatabaseResponse) SetStatusCode(v int32) *InitVectorDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *InitVectorDatabaseResponse) SetBody(v *InitVectorDatabaseResponseBody) *InitVectorDatabaseResponse {
	s.Body = v
	return s
}

func (s *InitVectorDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
