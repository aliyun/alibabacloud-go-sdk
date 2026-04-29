// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentMonitorStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudAgentMonitorStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudAgentMonitorStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *CloudAgentMonitorStatisticsResponseBody) *CloudAgentMonitorStatisticsResponse
	GetBody() *CloudAgentMonitorStatisticsResponseBody
}

type CloudAgentMonitorStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudAgentMonitorStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudAgentMonitorStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentMonitorStatisticsResponse) GoString() string {
	return s.String()
}

func (s *CloudAgentMonitorStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudAgentMonitorStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudAgentMonitorStatisticsResponse) GetBody() *CloudAgentMonitorStatisticsResponseBody {
	return s.Body
}

func (s *CloudAgentMonitorStatisticsResponse) SetHeaders(v map[string]*string) *CloudAgentMonitorStatisticsResponse {
	s.Headers = v
	return s
}

func (s *CloudAgentMonitorStatisticsResponse) SetStatusCode(v int32) *CloudAgentMonitorStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponse) SetBody(v *CloudAgentMonitorStatisticsResponseBody) *CloudAgentMonitorStatisticsResponse {
	s.Body = v
	return s
}

func (s *CloudAgentMonitorStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
