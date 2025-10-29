// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *GetDatabaseResponseBody) *GetDatabaseResponse
	GetBody() *GetDatabaseResponseBody
}

type GetDatabaseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatabaseResponse) GetBody() *GetDatabaseResponseBody {
	return s.Body
}

func (s *GetDatabaseResponse) SetHeaders(v map[string]*string) *GetDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseResponse) SetStatusCode(v int32) *GetDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseResponse) SetBody(v *GetDatabaseResponseBody) *GetDatabaseResponse {
	s.Body = v
	return s
}

func (s *GetDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
