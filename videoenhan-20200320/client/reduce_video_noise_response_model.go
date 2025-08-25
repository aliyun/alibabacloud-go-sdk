// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceVideoNoiseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReduceVideoNoiseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReduceVideoNoiseResponse
	GetStatusCode() *int32
	SetBody(v *ReduceVideoNoiseResponseBody) *ReduceVideoNoiseResponse
	GetBody() *ReduceVideoNoiseResponseBody
}

type ReduceVideoNoiseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReduceVideoNoiseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReduceVideoNoiseResponse) String() string {
	return dara.Prettify(s)
}

func (s ReduceVideoNoiseResponse) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReduceVideoNoiseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReduceVideoNoiseResponse) GetBody() *ReduceVideoNoiseResponseBody {
	return s.Body
}

func (s *ReduceVideoNoiseResponse) SetHeaders(v map[string]*string) *ReduceVideoNoiseResponse {
	s.Headers = v
	return s
}

func (s *ReduceVideoNoiseResponse) SetStatusCode(v int32) *ReduceVideoNoiseResponse {
	s.StatusCode = &v
	return s
}

func (s *ReduceVideoNoiseResponse) SetBody(v *ReduceVideoNoiseResponseBody) *ReduceVideoNoiseResponse {
	s.Body = v
	return s
}

func (s *ReduceVideoNoiseResponse) Validate() error {
	return dara.Validate(s)
}
