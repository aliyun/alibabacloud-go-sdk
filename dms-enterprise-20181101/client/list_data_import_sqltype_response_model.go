// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataImportSQLTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataImportSQLTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataImportSQLTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListDataImportSQLTypeResponseBody) *ListDataImportSQLTypeResponse
	GetBody() *ListDataImportSQLTypeResponseBody
}

type ListDataImportSQLTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataImportSQLTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataImportSQLTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataImportSQLTypeResponse) GoString() string {
	return s.String()
}

func (s *ListDataImportSQLTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataImportSQLTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataImportSQLTypeResponse) GetBody() *ListDataImportSQLTypeResponseBody {
	return s.Body
}

func (s *ListDataImportSQLTypeResponse) SetHeaders(v map[string]*string) *ListDataImportSQLTypeResponse {
	s.Headers = v
	return s
}

func (s *ListDataImportSQLTypeResponse) SetStatusCode(v int32) *ListDataImportSQLTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataImportSQLTypeResponse) SetBody(v *ListDataImportSQLTypeResponseBody) *ListDataImportSQLTypeResponse {
	s.Body = v
	return s
}

func (s *ListDataImportSQLTypeResponse) Validate() error {
	return dara.Validate(s)
}
