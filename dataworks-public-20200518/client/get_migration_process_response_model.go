// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMigrationProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMigrationProcessResponse
	GetStatusCode() *int32
	SetBody(v *GetMigrationProcessResponseBody) *GetMigrationProcessResponse
	GetBody() *GetMigrationProcessResponseBody
}

type GetMigrationProcessResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMigrationProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMigrationProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationProcessResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMigrationProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMigrationProcessResponse) GetBody() *GetMigrationProcessResponseBody {
	return s.Body
}

func (s *GetMigrationProcessResponse) SetHeaders(v map[string]*string) *GetMigrationProcessResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationProcessResponse) SetStatusCode(v int32) *GetMigrationProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationProcessResponse) SetBody(v *GetMigrationProcessResponseBody) *GetMigrationProcessResponse {
	s.Body = v
	return s
}

func (s *GetMigrationProcessResponse) Validate() error {
	return dara.Validate(s)
}
