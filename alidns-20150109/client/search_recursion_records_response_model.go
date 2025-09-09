// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchRecursionRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchRecursionRecordsResponse
	GetStatusCode() *int32
	SetBody(v *SearchRecursionRecordsResponseBody) *SearchRecursionRecordsResponse
	GetBody() *SearchRecursionRecordsResponseBody
}

type SearchRecursionRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchRecursionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchRecursionRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionRecordsResponse) GoString() string {
	return s.String()
}

func (s *SearchRecursionRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchRecursionRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchRecursionRecordsResponse) GetBody() *SearchRecursionRecordsResponseBody {
	return s.Body
}

func (s *SearchRecursionRecordsResponse) SetHeaders(v map[string]*string) *SearchRecursionRecordsResponse {
	s.Headers = v
	return s
}

func (s *SearchRecursionRecordsResponse) SetStatusCode(v int32) *SearchRecursionRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchRecursionRecordsResponse) SetBody(v *SearchRecursionRecordsResponseBody) *SearchRecursionRecordsResponse {
	s.Body = v
	return s
}

func (s *SearchRecursionRecordsResponse) Validate() error {
	return dara.Validate(s)
}
