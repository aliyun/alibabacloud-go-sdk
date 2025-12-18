// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourcesByAdvancedSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateResourcesByAdvancedSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateResourcesByAdvancedSearchResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateResourcesByAdvancedSearchResponseBody) *ListAggregateResourcesByAdvancedSearchResponse
	GetBody() *ListAggregateResourcesByAdvancedSearchResponseBody
}

type ListAggregateResourcesByAdvancedSearchResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateResourcesByAdvancedSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateResourcesByAdvancedSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourcesByAdvancedSearchResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateResourcesByAdvancedSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateResourcesByAdvancedSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateResourcesByAdvancedSearchResponse) GetBody() *ListAggregateResourcesByAdvancedSearchResponseBody {
	return s.Body
}

func (s *ListAggregateResourcesByAdvancedSearchResponse) SetHeaders(v map[string]*string) *ListAggregateResourcesByAdvancedSearchResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponse) SetStatusCode(v int32) *ListAggregateResourcesByAdvancedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponse) SetBody(v *ListAggregateResourcesByAdvancedSearchResponseBody) *ListAggregateResourcesByAdvancedSearchResponse {
	s.Body = v
	return s
}

func (s *ListAggregateResourcesByAdvancedSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
