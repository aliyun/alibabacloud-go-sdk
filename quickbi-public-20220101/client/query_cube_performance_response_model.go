// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCubePerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCubePerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCubePerformanceResponse
	GetStatusCode() *int32
	SetBody(v *QueryCubePerformanceResponseBody) *QueryCubePerformanceResponse
	GetBody() *QueryCubePerformanceResponseBody
}

type QueryCubePerformanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCubePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCubePerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCubePerformanceResponse) GoString() string {
	return s.String()
}

func (s *QueryCubePerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCubePerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCubePerformanceResponse) GetBody() *QueryCubePerformanceResponseBody {
	return s.Body
}

func (s *QueryCubePerformanceResponse) SetHeaders(v map[string]*string) *QueryCubePerformanceResponse {
	s.Headers = v
	return s
}

func (s *QueryCubePerformanceResponse) SetStatusCode(v int32) *QueryCubePerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCubePerformanceResponse) SetBody(v *QueryCubePerformanceResponseBody) *QueryCubePerformanceResponse {
	s.Body = v
	return s
}

func (s *QueryCubePerformanceResponse) Validate() error {
	return dara.Validate(s)
}
