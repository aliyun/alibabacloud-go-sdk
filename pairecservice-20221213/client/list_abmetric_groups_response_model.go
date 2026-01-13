// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABMetricGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListABMetricGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListABMetricGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListABMetricGroupsResponseBody) *ListABMetricGroupsResponse
	GetBody() *ListABMetricGroupsResponseBody
}

type ListABMetricGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABMetricGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABMetricGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListABMetricGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListABMetricGroupsResponse) GetBody() *ListABMetricGroupsResponseBody {
	return s.Body
}

func (s *ListABMetricGroupsResponse) SetHeaders(v map[string]*string) *ListABMetricGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListABMetricGroupsResponse) SetStatusCode(v int32) *ListABMetricGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABMetricGroupsResponse) SetBody(v *ListABMetricGroupsResponseBody) *ListABMetricGroupsResponse {
	s.Body = v
	return s
}

func (s *ListABMetricGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
