// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryComponentPerformanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryComponentPerformanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryComponentPerformanceResponse
	GetStatusCode() *int32
	SetBody(v *QueryComponentPerformanceResponseBody) *QueryComponentPerformanceResponse
	GetBody() *QueryComponentPerformanceResponseBody
}

type QueryComponentPerformanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryComponentPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryComponentPerformanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryComponentPerformanceResponse) GoString() string {
	return s.String()
}

func (s *QueryComponentPerformanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryComponentPerformanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryComponentPerformanceResponse) GetBody() *QueryComponentPerformanceResponseBody {
	return s.Body
}

func (s *QueryComponentPerformanceResponse) SetHeaders(v map[string]*string) *QueryComponentPerformanceResponse {
	s.Headers = v
	return s
}

func (s *QueryComponentPerformanceResponse) SetStatusCode(v int32) *QueryComponentPerformanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryComponentPerformanceResponse) SetBody(v *QueryComponentPerformanceResponseBody) *QueryComponentPerformanceResponse {
	s.Body = v
	return s
}

func (s *QueryComponentPerformanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
