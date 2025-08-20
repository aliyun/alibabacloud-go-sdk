// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResultExportJobHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResultExportJobHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResultExportJobHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListResultExportJobHistoryResponseBody) *ListResultExportJobHistoryResponse
	GetBody() *ListResultExportJobHistoryResponseBody
}

type ListResultExportJobHistoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResultExportJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResultExportJobHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResultExportJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListResultExportJobHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResultExportJobHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResultExportJobHistoryResponse) GetBody() *ListResultExportJobHistoryResponseBody {
	return s.Body
}

func (s *ListResultExportJobHistoryResponse) SetHeaders(v map[string]*string) *ListResultExportJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListResultExportJobHistoryResponse) SetStatusCode(v int32) *ListResultExportJobHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResultExportJobHistoryResponse) SetBody(v *ListResultExportJobHistoryResponseBody) *ListResultExportJobHistoryResponse {
	s.Body = v
	return s
}

func (s *ListResultExportJobHistoryResponse) Validate() error {
	return dara.Validate(s)
}
