// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTracesByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTracesByPageResponse
	GetStatusCode() *int32
	SetBody(v *SearchTracesByPageResponseBody) *SearchTracesByPageResponse
	GetBody() *SearchTracesByPageResponseBody
}

type SearchTracesByPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTracesByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTracesByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTracesByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTracesByPageResponse) GetBody() *SearchTracesByPageResponseBody {
	return s.Body
}

func (s *SearchTracesByPageResponse) SetHeaders(v map[string]*string) *SearchTracesByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesByPageResponse) SetStatusCode(v int32) *SearchTracesByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesByPageResponse) SetBody(v *SearchTracesByPageResponseBody) *SearchTracesByPageResponse {
	s.Body = v
	return s
}

func (s *SearchTracesByPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
