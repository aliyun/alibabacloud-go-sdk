// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceCountsGroupByRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceCountsGroupByRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceCountsGroupByRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceCountsGroupByRegionResponseBody) *GetAggregateResourceCountsGroupByRegionResponse
	GetBody() *GetAggregateResourceCountsGroupByRegionResponseBody
}

type GetAggregateResourceCountsGroupByRegionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceCountsGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) GetBody() *GetAggregateResourceCountsGroupByRegionResponseBody {
	return s.Body
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) SetHeaders(v map[string]*string) *GetAggregateResourceCountsGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) SetStatusCode(v int32) *GetAggregateResourceCountsGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) SetBody(v *GetAggregateResourceCountsGroupByRegionResponseBody) *GetAggregateResourceCountsGroupByRegionResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
