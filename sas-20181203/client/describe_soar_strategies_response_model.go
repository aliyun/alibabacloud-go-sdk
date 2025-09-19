// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarStrategiesResponseBody) *DescribeSoarStrategiesResponse
	GetBody() *DescribeSoarStrategiesResponseBody
}

type DescribeSoarStrategiesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarStrategiesResponse) GetBody() *DescribeSoarStrategiesResponseBody {
	return s.Body
}

func (s *DescribeSoarStrategiesResponse) SetHeaders(v map[string]*string) *DescribeSoarStrategiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarStrategiesResponse) SetStatusCode(v int32) *DescribeSoarStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarStrategiesResponse) SetBody(v *DescribeSoarStrategiesResponseBody) *DescribeSoarStrategiesResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarStrategiesResponse) Validate() error {
	return dara.Validate(s)
}
