// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataImportSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataImportSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataImportSQLResponse
	GetStatusCode() *int32
	SetBody(v *GetDataImportSQLResponseBody) *GetDataImportSQLResponse
	GetBody() *GetDataImportSQLResponseBody
}

type GetDataImportSQLResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataImportSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataImportSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataImportSQLResponse) GoString() string {
	return s.String()
}

func (s *GetDataImportSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataImportSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataImportSQLResponse) GetBody() *GetDataImportSQLResponseBody {
	return s.Body
}

func (s *GetDataImportSQLResponse) SetHeaders(v map[string]*string) *GetDataImportSQLResponse {
	s.Headers = v
	return s
}

func (s *GetDataImportSQLResponse) SetStatusCode(v int32) *GetDataImportSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataImportSQLResponse) SetBody(v *GetDataImportSQLResponseBody) *GetDataImportSQLResponse {
	s.Body = v
	return s
}

func (s *GetDataImportSQLResponse) Validate() error {
	return dara.Validate(s)
}
