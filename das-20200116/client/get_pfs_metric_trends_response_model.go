// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsMetricTrendsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPfsMetricTrendsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPfsMetricTrendsResponse
	GetStatusCode() *int32
	SetBody(v *GetPfsMetricTrendsResponseBody) *GetPfsMetricTrendsResponse
	GetBody() *GetPfsMetricTrendsResponseBody
}

type GetPfsMetricTrendsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPfsMetricTrendsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPfsMetricTrendsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPfsMetricTrendsResponse) GoString() string {
	return s.String()
}

func (s *GetPfsMetricTrendsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPfsMetricTrendsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPfsMetricTrendsResponse) GetBody() *GetPfsMetricTrendsResponseBody {
	return s.Body
}

func (s *GetPfsMetricTrendsResponse) SetHeaders(v map[string]*string) *GetPfsMetricTrendsResponse {
	s.Headers = v
	return s
}

func (s *GetPfsMetricTrendsResponse) SetStatusCode(v int32) *GetPfsMetricTrendsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPfsMetricTrendsResponse) SetBody(v *GetPfsMetricTrendsResponseBody) *GetPfsMetricTrendsResponse {
	s.Body = v
	return s
}

func (s *GetPfsMetricTrendsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
