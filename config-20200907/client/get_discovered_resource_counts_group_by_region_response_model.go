// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceCountsGroupByRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDiscoveredResourceCountsGroupByRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDiscoveredResourceCountsGroupByRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetDiscoveredResourceCountsGroupByRegionResponseBody) *GetDiscoveredResourceCountsGroupByRegionResponse
	GetBody() *GetDiscoveredResourceCountsGroupByRegionResponseBody
}

type GetDiscoveredResourceCountsGroupByRegionResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiscoveredResourceCountsGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) GetBody() *GetDiscoveredResourceCountsGroupByRegionResponseBody {
	return s.Body
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceCountsGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) SetStatusCode(v int32) *GetDiscoveredResourceCountsGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) SetBody(v *GetDiscoveredResourceCountsGroupByRegionResponseBody) *GetDiscoveredResourceCountsGroupByRegionResponse {
	s.Body = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
