// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse
	GetBody() *ListApplicationsResponseBody
}

type ListApplicationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsResponse) GetBody() *ListApplicationsResponseBody {
	return s.Body
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetStatusCode(v int32) *ListApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
