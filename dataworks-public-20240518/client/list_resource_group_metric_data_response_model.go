// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceGroupMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceGroupMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceGroupMetricDataResponseBody) *ListResourceGroupMetricDataResponse
	GetBody() *ListResourceGroupMetricDataResponseBody
}

type ListResourceGroupMetricDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMetricDataResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceGroupMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceGroupMetricDataResponse) GetBody() *ListResourceGroupMetricDataResponseBody {
	return s.Body
}

func (s *ListResourceGroupMetricDataResponse) SetHeaders(v map[string]*string) *ListResourceGroupMetricDataResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupMetricDataResponse) SetStatusCode(v int32) *ListResourceGroupMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupMetricDataResponse) SetBody(v *ListResourceGroupMetricDataResponseBody) *ListResourceGroupMetricDataResponse {
	s.Body = v
	return s
}

func (s *ListResourceGroupMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
