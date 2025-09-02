// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListFileTypeResponseBody) *ListFileTypeResponse
	GetBody() *ListFileTypeResponseBody
}

type ListFileTypeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileTypeResponse) GoString() string {
	return s.String()
}

func (s *ListFileTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileTypeResponse) GetBody() *ListFileTypeResponseBody {
	return s.Body
}

func (s *ListFileTypeResponse) SetHeaders(v map[string]*string) *ListFileTypeResponse {
	s.Headers = v
	return s
}

func (s *ListFileTypeResponse) SetStatusCode(v int32) *ListFileTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileTypeResponse) SetBody(v *ListFileTypeResponseBody) *ListFileTypeResponse {
	s.Body = v
	return s
}

func (s *ListFileTypeResponse) Validate() error {
	return dara.Validate(s)
}
