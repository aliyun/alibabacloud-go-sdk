// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregatorResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregatorResponseBody) *GetAggregatorResponse
	GetBody() *GetAggregatorResponseBody
}

type GetAggregatorResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregatorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregatorResponse) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregatorResponse) GetBody() *GetAggregatorResponseBody {
	return s.Body
}

func (s *GetAggregatorResponse) SetHeaders(v map[string]*string) *GetAggregatorResponse {
	s.Headers = v
	return s
}

func (s *GetAggregatorResponse) SetStatusCode(v int32) *GetAggregatorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregatorResponse) SetBody(v *GetAggregatorResponseBody) *GetAggregatorResponse {
	s.Body = v
	return s
}

func (s *GetAggregatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
