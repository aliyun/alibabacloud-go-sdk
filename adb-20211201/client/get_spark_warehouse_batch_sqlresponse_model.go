// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkWarehouseBatchSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkWarehouseBatchSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkWarehouseBatchSQLResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkWarehouseBatchSQLResponseBody) *GetSparkWarehouseBatchSQLResponse
	GetBody() *GetSparkWarehouseBatchSQLResponseBody
}

type GetSparkWarehouseBatchSQLResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkWarehouseBatchSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkWarehouseBatchSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkWarehouseBatchSQLResponse) GoString() string {
	return s.String()
}

func (s *GetSparkWarehouseBatchSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkWarehouseBatchSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkWarehouseBatchSQLResponse) GetBody() *GetSparkWarehouseBatchSQLResponseBody {
	return s.Body
}

func (s *GetSparkWarehouseBatchSQLResponse) SetHeaders(v map[string]*string) *GetSparkWarehouseBatchSQLResponse {
	s.Headers = v
	return s
}

func (s *GetSparkWarehouseBatchSQLResponse) SetStatusCode(v int32) *GetSparkWarehouseBatchSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkWarehouseBatchSQLResponse) SetBody(v *GetSparkWarehouseBatchSQLResponseBody) *GetSparkWarehouseBatchSQLResponse {
	s.Body = v
	return s
}

func (s *GetSparkWarehouseBatchSQLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
