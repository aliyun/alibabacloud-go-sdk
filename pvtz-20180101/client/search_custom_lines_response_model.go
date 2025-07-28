// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCustomLinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchCustomLinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchCustomLinesResponse
	GetStatusCode() *int32
	SetBody(v *SearchCustomLinesResponseBody) *SearchCustomLinesResponse
	GetBody() *SearchCustomLinesResponseBody
}

type SearchCustomLinesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCustomLinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCustomLinesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchCustomLinesResponse) GoString() string {
	return s.String()
}

func (s *SearchCustomLinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchCustomLinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchCustomLinesResponse) GetBody() *SearchCustomLinesResponseBody {
	return s.Body
}

func (s *SearchCustomLinesResponse) SetHeaders(v map[string]*string) *SearchCustomLinesResponse {
	s.Headers = v
	return s
}

func (s *SearchCustomLinesResponse) SetStatusCode(v int32) *SearchCustomLinesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCustomLinesResponse) SetBody(v *SearchCustomLinesResponseBody) *SearchCustomLinesResponse {
	s.Body = v
	return s
}

func (s *SearchCustomLinesResponse) Validate() error {
	return dara.Validate(s)
}
