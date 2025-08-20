// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePerformanceViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePerformanceViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePerformanceViewResponse
	GetStatusCode() *int32
	SetBody(v *DeletePerformanceViewResponseBody) *DeletePerformanceViewResponse
	GetBody() *DeletePerformanceViewResponseBody
}

type DeletePerformanceViewResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePerformanceViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePerformanceViewResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePerformanceViewResponse) GoString() string {
	return s.String()
}

func (s *DeletePerformanceViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePerformanceViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePerformanceViewResponse) GetBody() *DeletePerformanceViewResponseBody {
	return s.Body
}

func (s *DeletePerformanceViewResponse) SetHeaders(v map[string]*string) *DeletePerformanceViewResponse {
	s.Headers = v
	return s
}

func (s *DeletePerformanceViewResponse) SetStatusCode(v int32) *DeletePerformanceViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePerformanceViewResponse) SetBody(v *DeletePerformanceViewResponseBody) *DeletePerformanceViewResponse {
	s.Body = v
	return s
}

func (s *DeletePerformanceViewResponse) Validate() error {
	return dara.Validate(s)
}
