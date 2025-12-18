// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRemediationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateRemediationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateRemediationsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateRemediationsResponseBody) *ListAggregateRemediationsResponse
	GetBody() *ListAggregateRemediationsResponseBody
}

type ListAggregateRemediationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateRemediationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateRemediationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateRemediationsResponse) GetBody() *ListAggregateRemediationsResponseBody {
	return s.Body
}

func (s *ListAggregateRemediationsResponse) SetHeaders(v map[string]*string) *ListAggregateRemediationsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateRemediationsResponse) SetStatusCode(v int32) *ListAggregateRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateRemediationsResponse) SetBody(v *ListAggregateRemediationsResponseBody) *ListAggregateRemediationsResponse {
	s.Body = v
	return s
}

func (s *ListAggregateRemediationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
