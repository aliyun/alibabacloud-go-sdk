// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataLakeDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataLakeDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *GetDataLakeDatabaseResponseBody) *GetDataLakeDatabaseResponse
	GetBody() *GetDataLakeDatabaseResponseBody
}

type GetDataLakeDatabaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataLakeDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataLakeDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetDataLakeDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataLakeDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataLakeDatabaseResponse) GetBody() *GetDataLakeDatabaseResponseBody {
	return s.Body
}

func (s *GetDataLakeDatabaseResponse) SetHeaders(v map[string]*string) *GetDataLakeDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetDataLakeDatabaseResponse) SetStatusCode(v int32) *GetDataLakeDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataLakeDatabaseResponse) SetBody(v *GetDataLakeDatabaseResponseBody) *GetDataLakeDatabaseResponse {
	s.Body = v
	return s
}

func (s *GetDataLakeDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
