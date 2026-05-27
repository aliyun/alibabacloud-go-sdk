// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumStorageMetricsByTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SumStorageMetricsByTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SumStorageMetricsByTypeResponse
	GetStatusCode() *int32
	SetBody(v *SumStorageMetricsByTypeResponseBody) *SumStorageMetricsByTypeResponse
	GetBody() *SumStorageMetricsByTypeResponseBody
}

type SumStorageMetricsByTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SumStorageMetricsByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SumStorageMetricsByTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByTypeResponse) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SumStorageMetricsByTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SumStorageMetricsByTypeResponse) GetBody() *SumStorageMetricsByTypeResponseBody {
	return s.Body
}

func (s *SumStorageMetricsByTypeResponse) SetHeaders(v map[string]*string) *SumStorageMetricsByTypeResponse {
	s.Headers = v
	return s
}

func (s *SumStorageMetricsByTypeResponse) SetStatusCode(v int32) *SumStorageMetricsByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SumStorageMetricsByTypeResponse) SetBody(v *SumStorageMetricsByTypeResponseBody) *SumStorageMetricsByTypeResponse {
	s.Body = v
	return s
}

func (s *SumStorageMetricsByTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
