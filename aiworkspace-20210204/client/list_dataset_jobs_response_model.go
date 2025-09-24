// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasetJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasetJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasetJobsResponseBody) *ListDatasetJobsResponse
	GetBody() *ListDatasetJobsResponseBody
}

type ListDatasetJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasetJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasetJobsResponse) GetBody() *ListDatasetJobsResponseBody {
	return s.Body
}

func (s *ListDatasetJobsResponse) SetHeaders(v map[string]*string) *ListDatasetJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetJobsResponse) SetStatusCode(v int32) *ListDatasetJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetJobsResponse) SetBody(v *ListDatasetJobsResponseBody) *ListDatasetJobsResponse {
	s.Body = v
	return s
}

func (s *ListDatasetJobsResponse) Validate() error {
	return dara.Validate(s)
}
