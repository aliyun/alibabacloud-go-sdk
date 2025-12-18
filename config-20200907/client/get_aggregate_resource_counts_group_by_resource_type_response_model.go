// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceCountsGroupByResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceCountsGroupByResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceCountsGroupByResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceCountsGroupByResourceTypeResponseBody) *GetAggregateResourceCountsGroupByResourceTypeResponse
	GetBody() *GetAggregateResourceCountsGroupByResourceTypeResponseBody
}

type GetAggregateResourceCountsGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceCountsGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) GetBody() *GetAggregateResourceCountsGroupByResourceTypeResponseBody {
	return s.Body
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetAggregateResourceCountsGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) SetStatusCode(v int32) *GetAggregateResourceCountsGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) SetBody(v *GetAggregateResourceCountsGroupByResourceTypeResponseBody) *GetAggregateResourceCountsGroupByResourceTypeResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
