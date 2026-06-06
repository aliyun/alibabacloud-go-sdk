// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskExecutionStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskExecutionStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskExecutionStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskExecutionStatisticsResponseBody) *GetTaskExecutionStatisticsResponse
	GetBody() *GetTaskExecutionStatisticsResponseBody
}

type GetTaskExecutionStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskExecutionStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskExecutionStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskExecutionStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetTaskExecutionStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskExecutionStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskExecutionStatisticsResponse) GetBody() *GetTaskExecutionStatisticsResponseBody {
	return s.Body
}

func (s *GetTaskExecutionStatisticsResponse) SetHeaders(v map[string]*string) *GetTaskExecutionStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetTaskExecutionStatisticsResponse) SetStatusCode(v int32) *GetTaskExecutionStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskExecutionStatisticsResponse) SetBody(v *GetTaskExecutionStatisticsResponseBody) *GetTaskExecutionStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetTaskExecutionStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
