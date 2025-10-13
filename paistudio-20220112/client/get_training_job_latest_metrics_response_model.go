// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainingJobLatestMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrainingJobLatestMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrainingJobLatestMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetTrainingJobLatestMetricsResponseBody) *GetTrainingJobLatestMetricsResponse
	GetBody() *GetTrainingJobLatestMetricsResponseBody
}

type GetTrainingJobLatestMetricsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrainingJobLatestMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrainingJobLatestMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobLatestMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrainingJobLatestMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrainingJobLatestMetricsResponse) GetBody() *GetTrainingJobLatestMetricsResponseBody {
	return s.Body
}

func (s *GetTrainingJobLatestMetricsResponse) SetHeaders(v map[string]*string) *GetTrainingJobLatestMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetTrainingJobLatestMetricsResponse) SetStatusCode(v int32) *GetTrainingJobLatestMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrainingJobLatestMetricsResponse) SetBody(v *GetTrainingJobLatestMetricsResponseBody) *GetTrainingJobLatestMetricsResponse {
	s.Body = v
	return s
}

func (s *GetTrainingJobLatestMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
