// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAggregatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAggregatorResponse
	GetStatusCode() *int32
	SetBody(v *CreateAggregatorResponseBody) *CreateAggregatorResponse
	GetBody() *CreateAggregatorResponseBody
}

type CreateAggregatorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAggregatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAggregatorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregatorResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAggregatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAggregatorResponse) GetBody() *CreateAggregatorResponseBody {
	return s.Body
}

func (s *CreateAggregatorResponse) SetHeaders(v map[string]*string) *CreateAggregatorResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregatorResponse) SetStatusCode(v int32) *CreateAggregatorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregatorResponse) SetBody(v *CreateAggregatorResponseBody) *CreateAggregatorResponse {
	s.Body = v
	return s
}

func (s *CreateAggregatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
