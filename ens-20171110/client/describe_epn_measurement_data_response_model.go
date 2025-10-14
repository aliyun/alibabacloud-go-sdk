// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnMeasurementDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEpnMeasurementDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEpnMeasurementDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEpnMeasurementDataResponseBody) *DescribeEpnMeasurementDataResponse
	GetBody() *DescribeEpnMeasurementDataResponseBody
}

type DescribeEpnMeasurementDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEpnMeasurementDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEpnMeasurementDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEpnMeasurementDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEpnMeasurementDataResponse) GetBody() *DescribeEpnMeasurementDataResponseBody {
	return s.Body
}

func (s *DescribeEpnMeasurementDataResponse) SetHeaders(v map[string]*string) *DescribeEpnMeasurementDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEpnMeasurementDataResponse) SetStatusCode(v int32) *DescribeEpnMeasurementDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEpnMeasurementDataResponse) SetBody(v *DescribeEpnMeasurementDataResponseBody) *DescribeEpnMeasurementDataResponse {
	s.Body = v
	return s
}

func (s *DescribeEpnMeasurementDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
