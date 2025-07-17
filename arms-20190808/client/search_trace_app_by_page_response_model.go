// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTraceAppByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTraceAppByPageResponse
	GetStatusCode() *int32
	SetBody(v *SearchTraceAppByPageResponseBody) *SearchTraceAppByPageResponse
	GetBody() *SearchTraceAppByPageResponseBody
}

type SearchTraceAppByPageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTraceAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTraceAppByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTraceAppByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTraceAppByPageResponse) GetBody() *SearchTraceAppByPageResponseBody {
	return s.Body
}

func (s *SearchTraceAppByPageResponse) SetHeaders(v map[string]*string) *SearchTraceAppByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchTraceAppByPageResponse) SetStatusCode(v int32) *SearchTraceAppByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTraceAppByPageResponse) SetBody(v *SearchTraceAppByPageResponseBody) *SearchTraceAppByPageResponse {
	s.Body = v
	return s
}

func (s *SearchTraceAppByPageResponse) Validate() error {
	return dara.Validate(s)
}
