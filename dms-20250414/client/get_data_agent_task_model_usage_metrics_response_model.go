// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataAgentTaskModelUsageMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataAgentTaskModelUsageMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetDataAgentTaskModelUsageMetricsResponseBody) *GetDataAgentTaskModelUsageMetricsResponse
	GetBody() *GetDataAgentTaskModelUsageMetricsResponseBody
}

type GetDataAgentTaskModelUsageMetricsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataAgentTaskModelUsageMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataAgentTaskModelUsageMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataAgentTaskModelUsageMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataAgentTaskModelUsageMetricsResponse) GetBody() *GetDataAgentTaskModelUsageMetricsResponseBody {
	return s.Body
}

func (s *GetDataAgentTaskModelUsageMetricsResponse) SetHeaders(v map[string]*string) *GetDataAgentTaskModelUsageMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponse) SetStatusCode(v int32) *GetDataAgentTaskModelUsageMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponse) SetBody(v *GetDataAgentTaskModelUsageMetricsResponseBody) *GetDataAgentTaskModelUsageMetricsResponse {
	s.Body = v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
