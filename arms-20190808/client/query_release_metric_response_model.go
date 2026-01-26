// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReleaseMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReleaseMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReleaseMetricResponse
	GetStatusCode() *int32
	SetBody(v *QueryReleaseMetricResponseBody) *QueryReleaseMetricResponse
	GetBody() *QueryReleaseMetricResponseBody
}

type QueryReleaseMetricResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReleaseMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReleaseMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReleaseMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryReleaseMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReleaseMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReleaseMetricResponse) GetBody() *QueryReleaseMetricResponseBody {
	return s.Body
}

func (s *QueryReleaseMetricResponse) SetHeaders(v map[string]*string) *QueryReleaseMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryReleaseMetricResponse) SetStatusCode(v int32) *QueryReleaseMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReleaseMetricResponse) SetBody(v *QueryReleaseMetricResponseBody) *QueryReleaseMetricResponse {
	s.Body = v
	return s
}

func (s *QueryReleaseMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
