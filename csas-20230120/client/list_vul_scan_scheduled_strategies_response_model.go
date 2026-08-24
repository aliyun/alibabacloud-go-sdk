// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulScanScheduledStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVulScanScheduledStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVulScanScheduledStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *ListVulScanScheduledStrategiesResponseBody) *ListVulScanScheduledStrategiesResponse
	GetBody() *ListVulScanScheduledStrategiesResponseBody
}

type ListVulScanScheduledStrategiesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVulScanScheduledStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVulScanScheduledStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanScheduledStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListVulScanScheduledStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVulScanScheduledStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVulScanScheduledStrategiesResponse) GetBody() *ListVulScanScheduledStrategiesResponseBody {
	return s.Body
}

func (s *ListVulScanScheduledStrategiesResponse) SetHeaders(v map[string]*string) *ListVulScanScheduledStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListVulScanScheduledStrategiesResponse) SetStatusCode(v int32) *ListVulScanScheduledStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVulScanScheduledStrategiesResponse) SetBody(v *ListVulScanScheduledStrategiesResponseBody) *ListVulScanScheduledStrategiesResponse {
	s.Body = v
	return s
}

func (s *ListVulScanScheduledStrategiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
