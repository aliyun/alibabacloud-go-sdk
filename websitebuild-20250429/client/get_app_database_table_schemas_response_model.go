// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppDatabaseTableSchemasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppDatabaseTableSchemasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppDatabaseTableSchemasResponse
	GetStatusCode() *int32
	SetBody(v *GetAppDatabaseTableSchemasResponseBody) *GetAppDatabaseTableSchemasResponse
	GetBody() *GetAppDatabaseTableSchemasResponseBody
}

type GetAppDatabaseTableSchemasResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppDatabaseTableSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppDatabaseTableSchemasResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppDatabaseTableSchemasResponse) GoString() string {
	return s.String()
}

func (s *GetAppDatabaseTableSchemasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppDatabaseTableSchemasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppDatabaseTableSchemasResponse) GetBody() *GetAppDatabaseTableSchemasResponseBody {
	return s.Body
}

func (s *GetAppDatabaseTableSchemasResponse) SetHeaders(v map[string]*string) *GetAppDatabaseTableSchemasResponse {
	s.Headers = v
	return s
}

func (s *GetAppDatabaseTableSchemasResponse) SetStatusCode(v int32) *GetAppDatabaseTableSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppDatabaseTableSchemasResponse) SetBody(v *GetAppDatabaseTableSchemasResponseBody) *GetAppDatabaseTableSchemasResponse {
	s.Body = v
	return s
}

func (s *GetAppDatabaseTableSchemasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
