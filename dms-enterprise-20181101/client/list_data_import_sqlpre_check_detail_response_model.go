// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataImportSQLPreCheckDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataImportSQLPreCheckDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataImportSQLPreCheckDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListDataImportSQLPreCheckDetailResponseBody) *ListDataImportSQLPreCheckDetailResponse
	GetBody() *ListDataImportSQLPreCheckDetailResponseBody
}

type ListDataImportSQLPreCheckDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataImportSQLPreCheckDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataImportSQLPreCheckDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataImportSQLPreCheckDetailResponse) GoString() string {
	return s.String()
}

func (s *ListDataImportSQLPreCheckDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataImportSQLPreCheckDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataImportSQLPreCheckDetailResponse) GetBody() *ListDataImportSQLPreCheckDetailResponseBody {
	return s.Body
}

func (s *ListDataImportSQLPreCheckDetailResponse) SetHeaders(v map[string]*string) *ListDataImportSQLPreCheckDetailResponse {
	s.Headers = v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponse) SetStatusCode(v int32) *ListDataImportSQLPreCheckDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponse) SetBody(v *ListDataImportSQLPreCheckDetailResponseBody) *ListDataImportSQLPreCheckDetailResponse {
	s.Body = v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponse) Validate() error {
	return dara.Validate(s)
}
