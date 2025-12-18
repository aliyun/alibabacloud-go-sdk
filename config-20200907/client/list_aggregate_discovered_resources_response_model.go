// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateDiscoveredResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateDiscoveredResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateDiscoveredResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateDiscoveredResourcesResponseBody) *ListAggregateDiscoveredResourcesResponse
	GetBody() *ListAggregateDiscoveredResourcesResponseBody
}

type ListAggregateDiscoveredResourcesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateDiscoveredResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateDiscoveredResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateDiscoveredResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateDiscoveredResourcesResponse) GetBody() *ListAggregateDiscoveredResourcesResponseBody {
	return s.Body
}

func (s *ListAggregateDiscoveredResourcesResponse) SetHeaders(v map[string]*string) *ListAggregateDiscoveredResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponse) SetStatusCode(v int32) *ListAggregateDiscoveredResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponse) SetBody(v *ListAggregateDiscoveredResourcesResponseBody) *ListAggregateDiscoveredResourcesResponse {
	s.Body = v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
