// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateDiscoveredResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateDiscoveredResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateDiscoveredResourceResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateDiscoveredResourceResponseBody) *GetAggregateDiscoveredResourceResponse
	GetBody() *GetAggregateDiscoveredResourceResponseBody
}

type GetAggregateDiscoveredResourceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateDiscoveredResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateDiscoveredResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateDiscoveredResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateDiscoveredResourceResponse) GetBody() *GetAggregateDiscoveredResourceResponseBody {
	return s.Body
}

func (s *GetAggregateDiscoveredResourceResponse) SetHeaders(v map[string]*string) *GetAggregateDiscoveredResourceResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateDiscoveredResourceResponse) SetStatusCode(v int32) *GetAggregateDiscoveredResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponse) SetBody(v *GetAggregateDiscoveredResourceResponseBody) *GetAggregateDiscoveredResourceResponse {
	s.Body = v
	return s
}

func (s *GetAggregateDiscoveredResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
