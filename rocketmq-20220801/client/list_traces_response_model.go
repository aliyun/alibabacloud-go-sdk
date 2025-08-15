// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTracesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTracesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTracesResponse
	GetStatusCode() *int32
	SetBody(v *ListTracesResponseBody) *ListTracesResponse
	GetBody() *ListTracesResponseBody
}

type ListTracesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTracesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTracesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTracesResponse) GoString() string {
	return s.String()
}

func (s *ListTracesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTracesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTracesResponse) GetBody() *ListTracesResponseBody {
	return s.Body
}

func (s *ListTracesResponse) SetHeaders(v map[string]*string) *ListTracesResponse {
	s.Headers = v
	return s
}

func (s *ListTracesResponse) SetStatusCode(v int32) *ListTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTracesResponse) SetBody(v *ListTracesResponseBody) *ListTracesResponse {
	s.Body = v
	return s
}

func (s *ListTracesResponse) Validate() error {
	return dara.Validate(s)
}
