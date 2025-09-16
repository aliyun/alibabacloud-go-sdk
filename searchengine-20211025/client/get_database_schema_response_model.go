// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatabaseSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatabaseSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetDatabaseSchemaResponseBody) *GetDatabaseSchemaResponse
	GetBody() *GetDatabaseSchemaResponseBody
}

type GetDatabaseSchemaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatabaseSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatabaseSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatabaseSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatabaseSchemaResponse) GetBody() *GetDatabaseSchemaResponseBody {
	return s.Body
}

func (s *GetDatabaseSchemaResponse) SetHeaders(v map[string]*string) *GetDatabaseSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseSchemaResponse) SetStatusCode(v int32) *GetDatabaseSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseSchemaResponse) SetBody(v *GetDatabaseSchemaResponseBody) *GetDatabaseSchemaResponse {
	s.Body = v
	return s
}

func (s *GetDatabaseSchemaResponse) Validate() error {
	return dara.Validate(s)
}
