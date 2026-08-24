// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanScheduledStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirusScanScheduledStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirusScanScheduledStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *ListVirusScanScheduledStrategiesResponseBody) *ListVirusScanScheduledStrategiesResponse
	GetBody() *ListVirusScanScheduledStrategiesResponseBody
}

type ListVirusScanScheduledStrategiesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirusScanScheduledStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirusScanScheduledStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanScheduledStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListVirusScanScheduledStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirusScanScheduledStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirusScanScheduledStrategiesResponse) GetBody() *ListVirusScanScheduledStrategiesResponseBody {
	return s.Body
}

func (s *ListVirusScanScheduledStrategiesResponse) SetHeaders(v map[string]*string) *ListVirusScanScheduledStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponse) SetStatusCode(v int32) *ListVirusScanScheduledStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponse) SetBody(v *ListVirusScanScheduledStrategiesResponseBody) *ListVirusScanScheduledStrategiesResponse {
	s.Body = v
	return s
}

func (s *ListVirusScanScheduledStrategiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
