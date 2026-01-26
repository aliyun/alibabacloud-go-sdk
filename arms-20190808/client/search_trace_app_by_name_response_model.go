// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTraceAppByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTraceAppByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTraceAppByNameResponse
	GetStatusCode() *int32
	SetBody(v *SearchTraceAppByNameResponseBody) *SearchTraceAppByNameResponse
	GetBody() *SearchTraceAppByNameResponseBody
}

type SearchTraceAppByNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTraceAppByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTraceAppByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTraceAppByNameResponse) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTraceAppByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTraceAppByNameResponse) GetBody() *SearchTraceAppByNameResponseBody {
	return s.Body
}

func (s *SearchTraceAppByNameResponse) SetHeaders(v map[string]*string) *SearchTraceAppByNameResponse {
	s.Headers = v
	return s
}

func (s *SearchTraceAppByNameResponse) SetStatusCode(v int32) *SearchTraceAppByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTraceAppByNameResponse) SetBody(v *SearchTraceAppByNameResponseBody) *SearchTraceAppByNameResponse {
	s.Body = v
	return s
}

func (s *SearchTraceAppByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
