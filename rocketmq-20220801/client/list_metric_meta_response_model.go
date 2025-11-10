// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetricMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetricMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetricMetaResponse
	GetStatusCode() *int32
	SetBody(v *ListMetricMetaResponseBody) *ListMetricMetaResponse
	GetBody() *ListMetricMetaResponseBody
}

type ListMetricMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetricMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetricMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetricMetaResponse) GoString() string {
	return s.String()
}

func (s *ListMetricMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetricMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetricMetaResponse) GetBody() *ListMetricMetaResponseBody {
	return s.Body
}

func (s *ListMetricMetaResponse) SetHeaders(v map[string]*string) *ListMetricMetaResponse {
	s.Headers = v
	return s
}

func (s *ListMetricMetaResponse) SetStatusCode(v int32) *ListMetricMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetricMetaResponse) SetBody(v *ListMetricMetaResponseBody) *ListMetricMetaResponse {
	s.Body = v
	return s
}

func (s *ListMetricMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
