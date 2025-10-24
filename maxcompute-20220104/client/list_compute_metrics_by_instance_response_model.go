// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeMetricsByInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeMetricsByInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeMetricsByInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeMetricsByInstanceResponseBody) *ListComputeMetricsByInstanceResponse
	GetBody() *ListComputeMetricsByInstanceResponseBody
}

type ListComputeMetricsByInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeMetricsByInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeMetricsByInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeMetricsByInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeMetricsByInstanceResponse) GetBody() *ListComputeMetricsByInstanceResponseBody {
	return s.Body
}

func (s *ListComputeMetricsByInstanceResponse) SetHeaders(v map[string]*string) *ListComputeMetricsByInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListComputeMetricsByInstanceResponse) SetStatusCode(v int32) *ListComputeMetricsByInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponse) SetBody(v *ListComputeMetricsByInstanceResponseBody) *ListComputeMetricsByInstanceResponse {
	s.Body = v
	return s
}

func (s *ListComputeMetricsByInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
