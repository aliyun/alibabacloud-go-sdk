// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPackageJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPackageJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPackageJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListPackageJobsResponseBody) *ListPackageJobsResponse
	GetBody() *ListPackageJobsResponseBody
}

type ListPackageJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPackageJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPackageJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsResponse) GoString() string {
	return s.String()
}

func (s *ListPackageJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPackageJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPackageJobsResponse) GetBody() *ListPackageJobsResponseBody {
	return s.Body
}

func (s *ListPackageJobsResponse) SetHeaders(v map[string]*string) *ListPackageJobsResponse {
	s.Headers = v
	return s
}

func (s *ListPackageJobsResponse) SetStatusCode(v int32) *ListPackageJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPackageJobsResponse) SetBody(v *ListPackageJobsResponseBody) *ListPackageJobsResponse {
	s.Body = v
	return s
}

func (s *ListPackageJobsResponse) Validate() error {
	return dara.Validate(s)
}
