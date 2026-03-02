// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcDeveloperRepoMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcDeveloperRepoMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcDeveloperRepoMetricsResponse
	GetStatusCode() *int32
	SetBody(v *PbcDeveloperRepoMetricResult) *GetPbcDeveloperRepoMetricsResponse
	GetBody() *PbcDeveloperRepoMetricResult
}

type GetPbcDeveloperRepoMetricsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcDeveloperRepoMetricResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcDeveloperRepoMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcDeveloperRepoMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetPbcDeveloperRepoMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcDeveloperRepoMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcDeveloperRepoMetricsResponse) GetBody() *PbcDeveloperRepoMetricResult {
	return s.Body
}

func (s *GetPbcDeveloperRepoMetricsResponse) SetHeaders(v map[string]*string) *GetPbcDeveloperRepoMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetPbcDeveloperRepoMetricsResponse) SetStatusCode(v int32) *GetPbcDeveloperRepoMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcDeveloperRepoMetricsResponse) SetBody(v *PbcDeveloperRepoMetricResult) *GetPbcDeveloperRepoMetricsResponse {
	s.Body = v
	return s
}

func (s *GetPbcDeveloperRepoMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
