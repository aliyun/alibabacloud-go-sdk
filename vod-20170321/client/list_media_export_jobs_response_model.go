// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaExportJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaExportJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaExportJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaExportJobsResponseBody) *ListMediaExportJobsResponse
	GetBody() *ListMediaExportJobsResponseBody
}

type ListMediaExportJobsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaExportJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaExportJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaExportJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaExportJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaExportJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaExportJobsResponse) GetBody() *ListMediaExportJobsResponseBody {
	return s.Body
}

func (s *ListMediaExportJobsResponse) SetHeaders(v map[string]*string) *ListMediaExportJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaExportJobsResponse) SetStatusCode(v int32) *ListMediaExportJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaExportJobsResponse) SetBody(v *ListMediaExportJobsResponseBody) *ListMediaExportJobsResponse {
	s.Body = v
	return s
}

func (s *ListMediaExportJobsResponse) Validate() error {
	return dara.Validate(s)
}
