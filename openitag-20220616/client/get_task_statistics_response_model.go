// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskStatisticsResponseBody) *GetTaskStatisticsResponse
	GetBody() *GetTaskStatisticsResponseBody
}

type GetTaskStatisticsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskStatisticsResponse) GetBody() *GetTaskStatisticsResponseBody {
	return s.Body
}

func (s *GetTaskStatisticsResponse) SetHeaders(v map[string]*string) *GetTaskStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatisticsResponse) SetStatusCode(v int32) *GetTaskStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatisticsResponse) SetBody(v *GetTaskStatisticsResponseBody) *GetTaskStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetTaskStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
