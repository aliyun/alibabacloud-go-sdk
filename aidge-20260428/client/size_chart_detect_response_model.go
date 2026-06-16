// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SizeChartDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SizeChartDetectResponse
	GetStatusCode() *int32
	SetBody(v *SizeChartDetectResponseBody) *SizeChartDetectResponse
	GetBody() *SizeChartDetectResponseBody
}

type SizeChartDetectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SizeChartDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SizeChartDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s SizeChartDetectResponse) GoString() string {
	return s.String()
}

func (s *SizeChartDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SizeChartDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SizeChartDetectResponse) GetBody() *SizeChartDetectResponseBody {
	return s.Body
}

func (s *SizeChartDetectResponse) SetHeaders(v map[string]*string) *SizeChartDetectResponse {
	s.Headers = v
	return s
}

func (s *SizeChartDetectResponse) SetStatusCode(v int32) *SizeChartDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *SizeChartDetectResponse) SetBody(v *SizeChartDetectResponseBody) *SizeChartDetectResponse {
	s.Body = v
	return s
}

func (s *SizeChartDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
