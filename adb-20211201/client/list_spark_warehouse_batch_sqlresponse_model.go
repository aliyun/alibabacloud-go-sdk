// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSparkWarehouseBatchSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSparkWarehouseBatchSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSparkWarehouseBatchSQLResponse
	GetStatusCode() *int32
	SetBody(v *ListSparkWarehouseBatchSQLResponseBody) *ListSparkWarehouseBatchSQLResponse
	GetBody() *ListSparkWarehouseBatchSQLResponseBody
}

type ListSparkWarehouseBatchSQLResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkWarehouseBatchSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSparkWarehouseBatchSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSparkWarehouseBatchSQLResponse) GoString() string {
	return s.String()
}

func (s *ListSparkWarehouseBatchSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSparkWarehouseBatchSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSparkWarehouseBatchSQLResponse) GetBody() *ListSparkWarehouseBatchSQLResponseBody {
	return s.Body
}

func (s *ListSparkWarehouseBatchSQLResponse) SetHeaders(v map[string]*string) *ListSparkWarehouseBatchSQLResponse {
	s.Headers = v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponse) SetStatusCode(v int32) *ListSparkWarehouseBatchSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponse) SetBody(v *ListSparkWarehouseBatchSQLResponseBody) *ListSparkWarehouseBatchSQLResponse {
	s.Body = v
	return s
}

func (s *ListSparkWarehouseBatchSQLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
