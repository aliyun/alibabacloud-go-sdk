// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceCountsGroupByResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDiscoveredResourceCountsGroupByResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDiscoveredResourceCountsGroupByResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) *GetDiscoveredResourceCountsGroupByResourceTypeResponse
	GetBody() *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody
}

type GetDiscoveredResourceCountsGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) GetBody() *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody {
	return s.Body
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceCountsGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) SetStatusCode(v int32) *GetDiscoveredResourceCountsGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) SetBody(v *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) *GetDiscoveredResourceCountsGroupByResourceTypeResponse {
	s.Body = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
