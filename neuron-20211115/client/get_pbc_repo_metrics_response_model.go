// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcRepoMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcRepoMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcRepoMetricsResponse
	GetStatusCode() *int32
	SetBody(v *PbcRepoMetricResult) *GetPbcRepoMetricsResponse
	GetBody() *PbcRepoMetricResult
}

type GetPbcRepoMetricsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcRepoMetricResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcRepoMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcRepoMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetPbcRepoMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcRepoMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcRepoMetricsResponse) GetBody() *PbcRepoMetricResult {
	return s.Body
}

func (s *GetPbcRepoMetricsResponse) SetHeaders(v map[string]*string) *GetPbcRepoMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetPbcRepoMetricsResponse) SetStatusCode(v int32) *GetPbcRepoMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcRepoMetricsResponse) SetBody(v *PbcRepoMetricResult) *GetPbcRepoMetricsResponse {
	s.Body = v
	return s
}

func (s *GetPbcRepoMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
