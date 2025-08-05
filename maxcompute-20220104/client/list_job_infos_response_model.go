// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListJobInfosResponseBody) *ListJobInfosResponse
	GetBody() *ListJobInfosResponseBody
}

type ListJobInfosResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobInfosResponse) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobInfosResponse) GetBody() *ListJobInfosResponseBody {
	return s.Body
}

func (s *ListJobInfosResponse) SetHeaders(v map[string]*string) *ListJobInfosResponse {
	s.Headers = v
	return s
}

func (s *ListJobInfosResponse) SetStatusCode(v int32) *ListJobInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobInfosResponse) SetBody(v *ListJobInfosResponseBody) *ListJobInfosResponse {
	s.Body = v
	return s
}

func (s *ListJobInfosResponse) Validate() error {
	return dara.Validate(s)
}
