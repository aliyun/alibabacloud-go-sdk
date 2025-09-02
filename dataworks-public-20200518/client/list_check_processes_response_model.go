// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckProcessesResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckProcessesResponseBody) *ListCheckProcessesResponse
	GetBody() *ListCheckProcessesResponseBody
}

type ListCheckProcessesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckProcessesResponse) GoString() string {
	return s.String()
}

func (s *ListCheckProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckProcessesResponse) GetBody() *ListCheckProcessesResponseBody {
	return s.Body
}

func (s *ListCheckProcessesResponse) SetHeaders(v map[string]*string) *ListCheckProcessesResponse {
	s.Headers = v
	return s
}

func (s *ListCheckProcessesResponse) SetStatusCode(v int32) *ListCheckProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckProcessesResponse) SetBody(v *ListCheckProcessesResponseBody) *ListCheckProcessesResponse {
	s.Body = v
	return s
}

func (s *ListCheckProcessesResponse) Validate() error {
	return dara.Validate(s)
}
