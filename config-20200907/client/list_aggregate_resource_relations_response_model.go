// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourceRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateResourceRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateResourceRelationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateResourceRelationsResponseBody) *ListAggregateResourceRelationsResponse
	GetBody() *ListAggregateResourceRelationsResponseBody
}

type ListAggregateResourceRelationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateResourceRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateResourceRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourceRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateResourceRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateResourceRelationsResponse) GetBody() *ListAggregateResourceRelationsResponseBody {
	return s.Body
}

func (s *ListAggregateResourceRelationsResponse) SetHeaders(v map[string]*string) *ListAggregateResourceRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateResourceRelationsResponse) SetStatusCode(v int32) *ListAggregateResourceRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateResourceRelationsResponse) SetBody(v *ListAggregateResourceRelationsResponseBody) *ListAggregateResourceRelationsResponse {
	s.Body = v
	return s
}

func (s *ListAggregateResourceRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
