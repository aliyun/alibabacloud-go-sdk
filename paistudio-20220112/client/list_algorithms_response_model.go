// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlgorithmsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlgorithmsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlgorithmsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlgorithmsResponseBody) *ListAlgorithmsResponse
	GetBody() *ListAlgorithmsResponseBody
}

type ListAlgorithmsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlgorithmsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlgorithmsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlgorithmsResponse) GoString() string {
	return s.String()
}

func (s *ListAlgorithmsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlgorithmsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlgorithmsResponse) GetBody() *ListAlgorithmsResponseBody {
	return s.Body
}

func (s *ListAlgorithmsResponse) SetHeaders(v map[string]*string) *ListAlgorithmsResponse {
	s.Headers = v
	return s
}

func (s *ListAlgorithmsResponse) SetStatusCode(v int32) *ListAlgorithmsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlgorithmsResponse) SetBody(v *ListAlgorithmsResponseBody) *ListAlgorithmsResponse {
	s.Body = v
	return s
}

func (s *ListAlgorithmsResponse) Validate() error {
	return dara.Validate(s)
}
