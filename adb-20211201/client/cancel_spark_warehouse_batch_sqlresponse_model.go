// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSparkWarehouseBatchSQLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSparkWarehouseBatchSQLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSparkWarehouseBatchSQLResponse
	GetStatusCode() *int32
	SetBody(v *CancelSparkWarehouseBatchSQLResponseBody) *CancelSparkWarehouseBatchSQLResponse
	GetBody() *CancelSparkWarehouseBatchSQLResponseBody
}

type CancelSparkWarehouseBatchSQLResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSparkWarehouseBatchSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSparkWarehouseBatchSQLResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSparkWarehouseBatchSQLResponse) GoString() string {
	return s.String()
}

func (s *CancelSparkWarehouseBatchSQLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSparkWarehouseBatchSQLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSparkWarehouseBatchSQLResponse) GetBody() *CancelSparkWarehouseBatchSQLResponseBody {
	return s.Body
}

func (s *CancelSparkWarehouseBatchSQLResponse) SetHeaders(v map[string]*string) *CancelSparkWarehouseBatchSQLResponse {
	s.Headers = v
	return s
}

func (s *CancelSparkWarehouseBatchSQLResponse) SetStatusCode(v int32) *CancelSparkWarehouseBatchSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSparkWarehouseBatchSQLResponse) SetBody(v *CancelSparkWarehouseBatchSQLResponseBody) *CancelSparkWarehouseBatchSQLResponse {
	s.Body = v
	return s
}

func (s *CancelSparkWarehouseBatchSQLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
