// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamDiscoveredResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamDiscoveredResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamDiscoveredResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamDiscoveredResourceResponseBody) *ListIpamDiscoveredResourceResponse
	GetBody() *ListIpamDiscoveredResourceResponseBody
}

type ListIpamDiscoveredResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamDiscoveredResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamDiscoveredResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamDiscoveredResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamDiscoveredResourceResponse) GetBody() *ListIpamDiscoveredResourceResponseBody {
	return s.Body
}

func (s *ListIpamDiscoveredResourceResponse) SetHeaders(v map[string]*string) *ListIpamDiscoveredResourceResponse {
	s.Headers = v
	return s
}

func (s *ListIpamDiscoveredResourceResponse) SetStatusCode(v int32) *ListIpamDiscoveredResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamDiscoveredResourceResponse) SetBody(v *ListIpamDiscoveredResourceResponseBody) *ListIpamDiscoveredResourceResponse {
	s.Body = v
	return s
}

func (s *ListIpamDiscoveredResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
