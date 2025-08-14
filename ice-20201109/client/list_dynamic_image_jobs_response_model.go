// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicImageJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDynamicImageJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDynamicImageJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListDynamicImageJobsResponseBody) *ListDynamicImageJobsResponse
	GetBody() *ListDynamicImageJobsResponseBody
}

type ListDynamicImageJobsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicImageJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicImageJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicImageJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDynamicImageJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDynamicImageJobsResponse) GetBody() *ListDynamicImageJobsResponseBody {
	return s.Body
}

func (s *ListDynamicImageJobsResponse) SetHeaders(v map[string]*string) *ListDynamicImageJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicImageJobsResponse) SetStatusCode(v int32) *ListDynamicImageJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicImageJobsResponse) SetBody(v *ListDynamicImageJobsResponseBody) *ListDynamicImageJobsResponse {
	s.Body = v
	return s
}

func (s *ListDynamicImageJobsResponse) Validate() error {
	return dara.Validate(s)
}
