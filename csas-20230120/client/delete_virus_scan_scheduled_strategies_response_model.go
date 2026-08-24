// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirusScanScheduledStrategiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirusScanScheduledStrategiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirusScanScheduledStrategiesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirusScanScheduledStrategiesResponseBody) *DeleteVirusScanScheduledStrategiesResponse
	GetBody() *DeleteVirusScanScheduledStrategiesResponseBody
}

type DeleteVirusScanScheduledStrategiesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirusScanScheduledStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirusScanScheduledStrategiesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirusScanScheduledStrategiesResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirusScanScheduledStrategiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirusScanScheduledStrategiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirusScanScheduledStrategiesResponse) GetBody() *DeleteVirusScanScheduledStrategiesResponseBody {
	return s.Body
}

func (s *DeleteVirusScanScheduledStrategiesResponse) SetHeaders(v map[string]*string) *DeleteVirusScanScheduledStrategiesResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirusScanScheduledStrategiesResponse) SetStatusCode(v int32) *DeleteVirusScanScheduledStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirusScanScheduledStrategiesResponse) SetBody(v *DeleteVirusScanScheduledStrategiesResponseBody) *DeleteVirusScanScheduledStrategiesResponse {
	s.Body = v
	return s
}

func (s *DeleteVirusScanScheduledStrategiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
