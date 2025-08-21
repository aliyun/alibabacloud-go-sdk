// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSolutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSolutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSolutionResponse
	GetStatusCode() *int32
	SetBody(v *ListSolutionResponseBody) *ListSolutionResponse
	GetBody() *ListSolutionResponseBody
}

type ListSolutionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSolutionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSolutionResponse) GoString() string {
	return s.String()
}

func (s *ListSolutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSolutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSolutionResponse) GetBody() *ListSolutionResponseBody {
	return s.Body
}

func (s *ListSolutionResponse) SetHeaders(v map[string]*string) *ListSolutionResponse {
	s.Headers = v
	return s
}

func (s *ListSolutionResponse) SetStatusCode(v int32) *ListSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSolutionResponse) SetBody(v *ListSolutionResponseBody) *ListSolutionResponse {
	s.Body = v
	return s
}

func (s *ListSolutionResponse) Validate() error {
	return dara.Validate(s)
}
