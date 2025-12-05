// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSampleMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJMeterSampleMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJMeterSampleMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetJMeterSampleMetricsResponseBody) *GetJMeterSampleMetricsResponse
	GetBody() *GetJMeterSampleMetricsResponseBody
}

type GetJMeterSampleMetricsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJMeterSampleMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJMeterSampleMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSampleMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterSampleMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJMeterSampleMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJMeterSampleMetricsResponse) GetBody() *GetJMeterSampleMetricsResponseBody {
	return s.Body
}

func (s *GetJMeterSampleMetricsResponse) SetHeaders(v map[string]*string) *GetJMeterSampleMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterSampleMetricsResponse) SetStatusCode(v int32) *GetJMeterSampleMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJMeterSampleMetricsResponse) SetBody(v *GetJMeterSampleMetricsResponseBody) *GetJMeterSampleMetricsResponse {
	s.Body = v
	return s
}

func (s *GetJMeterSampleMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
