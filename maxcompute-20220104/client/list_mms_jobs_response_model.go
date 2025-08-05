// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsJobsResponseBody) *ListMmsJobsResponse
	GetBody() *ListMmsJobsResponseBody
}

type ListMmsJobsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsJobsResponse) GetBody() *ListMmsJobsResponseBody {
	return s.Body
}

func (s *ListMmsJobsResponse) SetHeaders(v map[string]*string) *ListMmsJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsJobsResponse) SetStatusCode(v int32) *ListMmsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsJobsResponse) SetBody(v *ListMmsJobsResponseBody) *ListMmsJobsResponse {
	s.Body = v
	return s
}

func (s *ListMmsJobsResponse) Validate() error {
	return dara.Validate(s)
}
