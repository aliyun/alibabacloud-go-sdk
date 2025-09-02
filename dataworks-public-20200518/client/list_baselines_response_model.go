// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaselinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaselinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaselinesResponse
	GetStatusCode() *int32
	SetBody(v *ListBaselinesResponseBody) *ListBaselinesResponse
	GetBody() *ListBaselinesResponseBody
}

type ListBaselinesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaselinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaselinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaselinesResponse) GoString() string {
	return s.String()
}

func (s *ListBaselinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaselinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaselinesResponse) GetBody() *ListBaselinesResponseBody {
	return s.Body
}

func (s *ListBaselinesResponse) SetHeaders(v map[string]*string) *ListBaselinesResponse {
	s.Headers = v
	return s
}

func (s *ListBaselinesResponse) SetStatusCode(v int32) *ListBaselinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaselinesResponse) SetBody(v *ListBaselinesResponseBody) *ListBaselinesResponse {
	s.Body = v
	return s
}

func (s *ListBaselinesResponse) Validate() error {
	return dara.Validate(s)
}
