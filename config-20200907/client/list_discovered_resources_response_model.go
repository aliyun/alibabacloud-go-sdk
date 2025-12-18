// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiscoveredResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiscoveredResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiscoveredResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListDiscoveredResourcesResponseBody) *ListDiscoveredResourcesResponse
	GetBody() *ListDiscoveredResourcesResponseBody
}

type ListDiscoveredResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiscoveredResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiscoveredResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiscoveredResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiscoveredResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiscoveredResourcesResponse) GetBody() *ListDiscoveredResourcesResponseBody {
	return s.Body
}

func (s *ListDiscoveredResourcesResponse) SetHeaders(v map[string]*string) *ListDiscoveredResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDiscoveredResourcesResponse) SetStatusCode(v int32) *ListDiscoveredResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiscoveredResourcesResponse) SetBody(v *ListDiscoveredResourcesResponseBody) *ListDiscoveredResourcesResponse {
	s.Body = v
	return s
}

func (s *ListDiscoveredResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
