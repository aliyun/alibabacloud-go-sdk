// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAggregatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAggregatorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAggregatorResponseBody) *UpdateAggregatorResponse
	GetBody() *UpdateAggregatorResponseBody
}

type UpdateAggregatorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAggregatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAggregatorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregatorResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAggregatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAggregatorResponse) GetBody() *UpdateAggregatorResponseBody {
	return s.Body
}

func (s *UpdateAggregatorResponse) SetHeaders(v map[string]*string) *UpdateAggregatorResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregatorResponse) SetStatusCode(v int32) *UpdateAggregatorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregatorResponse) SetBody(v *UpdateAggregatorResponseBody) *UpdateAggregatorResponse {
	s.Body = v
	return s
}

func (s *UpdateAggregatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
