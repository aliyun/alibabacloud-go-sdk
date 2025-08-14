// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProgramsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProgramsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProgramsResponse
	GetStatusCode() *int32
	SetBody(v *ListProgramsResponseBody) *ListProgramsResponse
	GetBody() *ListProgramsResponseBody
}

type ListProgramsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProgramsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProgramsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProgramsResponse) GoString() string {
	return s.String()
}

func (s *ListProgramsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProgramsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProgramsResponse) GetBody() *ListProgramsResponseBody {
	return s.Body
}

func (s *ListProgramsResponse) SetHeaders(v map[string]*string) *ListProgramsResponse {
	s.Headers = v
	return s
}

func (s *ListProgramsResponse) SetStatusCode(v int32) *ListProgramsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProgramsResponse) SetBody(v *ListProgramsResponseBody) *ListProgramsResponse {
	s.Body = v
	return s
}

func (s *ListProgramsResponse) Validate() error {
	return dara.Validate(s)
}
