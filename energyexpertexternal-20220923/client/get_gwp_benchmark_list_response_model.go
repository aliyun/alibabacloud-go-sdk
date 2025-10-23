// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpBenchmarkListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGwpBenchmarkListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGwpBenchmarkListResponse
	GetStatusCode() *int32
	SetBody(v *GetGwpBenchmarkListResponseBody) *GetGwpBenchmarkListResponse
	GetBody() *GetGwpBenchmarkListResponseBody
}

type GetGwpBenchmarkListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpBenchmarkListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpBenchmarkListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkListResponse) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGwpBenchmarkListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGwpBenchmarkListResponse) GetBody() *GetGwpBenchmarkListResponseBody {
	return s.Body
}

func (s *GetGwpBenchmarkListResponse) SetHeaders(v map[string]*string) *GetGwpBenchmarkListResponse {
	s.Headers = v
	return s
}

func (s *GetGwpBenchmarkListResponse) SetStatusCode(v int32) *GetGwpBenchmarkListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpBenchmarkListResponse) SetBody(v *GetGwpBenchmarkListResponseBody) *GetGwpBenchmarkListResponse {
	s.Body = v
	return s
}

func (s *GetGwpBenchmarkListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
