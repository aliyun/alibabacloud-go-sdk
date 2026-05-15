// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentStatisticsResponseBody) *GetAgentStatisticsResponse
	GetBody() *GetAgentStatisticsResponseBody
}

type GetAgentStatisticsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetAgentStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentStatisticsResponse) GetBody() *GetAgentStatisticsResponseBody {
	return s.Body
}

func (s *GetAgentStatisticsResponse) SetHeaders(v map[string]*string) *GetAgentStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetAgentStatisticsResponse) SetStatusCode(v int32) *GetAgentStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentStatisticsResponse) SetBody(v *GetAgentStatisticsResponseBody) *GetAgentStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetAgentStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
