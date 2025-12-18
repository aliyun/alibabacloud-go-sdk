// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateRemediationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAggregateRemediationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAggregateRemediationsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAggregateRemediationsResponseBody) *DeleteAggregateRemediationsResponse
	GetBody() *DeleteAggregateRemediationsResponseBody
}

type DeleteAggregateRemediationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAggregateRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAggregateRemediationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateRemediationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAggregateRemediationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAggregateRemediationsResponse) GetBody() *DeleteAggregateRemediationsResponseBody {
	return s.Body
}

func (s *DeleteAggregateRemediationsResponse) SetHeaders(v map[string]*string) *DeleteAggregateRemediationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregateRemediationsResponse) SetStatusCode(v int32) *DeleteAggregateRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregateRemediationsResponse) SetBody(v *DeleteAggregateRemediationsResponseBody) *DeleteAggregateRemediationsResponse {
	s.Body = v
	return s
}

func (s *DeleteAggregateRemediationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
