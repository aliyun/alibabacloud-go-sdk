// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefineMaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefineMaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefineMaskResponse
	GetStatusCode() *int32
	SetBody(v *RefineMaskResponseBody) *RefineMaskResponse
	GetBody() *RefineMaskResponseBody
}

type RefineMaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefineMaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefineMaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RefineMaskResponse) GoString() string {
	return s.String()
}

func (s *RefineMaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefineMaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefineMaskResponse) GetBody() *RefineMaskResponseBody {
	return s.Body
}

func (s *RefineMaskResponse) SetHeaders(v map[string]*string) *RefineMaskResponse {
	s.Headers = v
	return s
}

func (s *RefineMaskResponse) SetStatusCode(v int32) *RefineMaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RefineMaskResponse) SetBody(v *RefineMaskResponseBody) *RefineMaskResponse {
	s.Body = v
	return s
}

func (s *RefineMaskResponse) Validate() error {
	return dara.Validate(s)
}
