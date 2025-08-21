// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeasurementDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeasurementDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeasurementDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeasurementDataResponseBody) *DescribeMeasurementDataResponse
	GetBody() *DescribeMeasurementDataResponseBody
}

type DescribeMeasurementDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeasurementDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeasurementDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeasurementDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeasurementDataResponse) GetBody() *DescribeMeasurementDataResponseBody {
	return s.Body
}

func (s *DescribeMeasurementDataResponse) SetHeaders(v map[string]*string) *DescribeMeasurementDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeasurementDataResponse) SetStatusCode(v int32) *DescribeMeasurementDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeasurementDataResponse) SetBody(v *DescribeMeasurementDataResponseBody) *DescribeMeasurementDataResponse {
	s.Body = v
	return s
}

func (s *DescribeMeasurementDataResponse) Validate() error {
	return dara.Validate(s)
}
