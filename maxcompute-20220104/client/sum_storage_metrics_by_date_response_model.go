// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumStorageMetricsByDateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SumStorageMetricsByDateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SumStorageMetricsByDateResponse
	GetStatusCode() *int32
	SetBody(v *SumStorageMetricsByDateResponseBody) *SumStorageMetricsByDateResponse
	GetBody() *SumStorageMetricsByDateResponseBody
}

type SumStorageMetricsByDateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SumStorageMetricsByDateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SumStorageMetricsByDateResponse) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByDateResponse) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByDateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SumStorageMetricsByDateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SumStorageMetricsByDateResponse) GetBody() *SumStorageMetricsByDateResponseBody {
	return s.Body
}

func (s *SumStorageMetricsByDateResponse) SetHeaders(v map[string]*string) *SumStorageMetricsByDateResponse {
	s.Headers = v
	return s
}

func (s *SumStorageMetricsByDateResponse) SetStatusCode(v int32) *SumStorageMetricsByDateResponse {
	s.StatusCode = &v
	return s
}

func (s *SumStorageMetricsByDateResponse) SetBody(v *SumStorageMetricsByDateResponseBody) *SumStorageMetricsByDateResponse {
	s.Body = v
	return s
}

func (s *SumStorageMetricsByDateResponse) Validate() error {
	return dara.Validate(s)
}
