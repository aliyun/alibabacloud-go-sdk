// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentMetricTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvironmentMetricTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvironmentMetricTargetsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvironmentMetricTargetsResponseBody) *ListEnvironmentMetricTargetsResponse
	GetBody() *ListEnvironmentMetricTargetsResponseBody
}

type ListEnvironmentMetricTargetsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentMetricTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentMetricTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentMetricTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentMetricTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvironmentMetricTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvironmentMetricTargetsResponse) GetBody() *ListEnvironmentMetricTargetsResponseBody {
	return s.Body
}

func (s *ListEnvironmentMetricTargetsResponse) SetHeaders(v map[string]*string) *ListEnvironmentMetricTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponse) SetStatusCode(v int32) *ListEnvironmentMetricTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentMetricTargetsResponse) SetBody(v *ListEnvironmentMetricTargetsResponseBody) *ListEnvironmentMetricTargetsResponse {
	s.Body = v
	return s
}

func (s *ListEnvironmentMetricTargetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
