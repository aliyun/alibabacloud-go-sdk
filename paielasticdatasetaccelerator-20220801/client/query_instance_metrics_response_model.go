// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstanceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstanceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstanceMetricsResponseBody) *QueryInstanceMetricsResponse
	GetBody() *QueryInstanceMetricsResponseBody
}

type QueryInstanceMetricsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstanceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstanceMetricsResponse) GetBody() *QueryInstanceMetricsResponseBody {
	return s.Body
}

func (s *QueryInstanceMetricsResponse) SetHeaders(v map[string]*string) *QueryInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceMetricsResponse) SetStatusCode(v int32) *QueryInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceMetricsResponse) SetBody(v *QueryInstanceMetricsResponseBody) *QueryInstanceMetricsResponse {
	s.Body = v
	return s
}

func (s *QueryInstanceMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
