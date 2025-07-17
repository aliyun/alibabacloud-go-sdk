// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTracesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTracesResponse
	GetStatusCode() *int32
	SetBody(v *SearchTracesResponseBody) *SearchTracesResponse
	GetBody() *SearchTracesResponseBody
}

type SearchTracesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTracesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTracesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTracesResponse) GetBody() *SearchTracesResponseBody {
	return s.Body
}

func (s *SearchTracesResponse) SetHeaders(v map[string]*string) *SearchTracesResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesResponse) SetStatusCode(v int32) *SearchTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesResponse) SetBody(v *SearchTracesResponseBody) *SearchTracesResponse {
	s.Body = v
	return s
}

func (s *SearchTracesResponse) Validate() error {
	return dara.Validate(s)
}
