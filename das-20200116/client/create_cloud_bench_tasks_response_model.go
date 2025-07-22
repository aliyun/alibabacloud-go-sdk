// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudBenchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudBenchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudBenchTasksResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudBenchTasksResponseBody) *CreateCloudBenchTasksResponse
	GetBody() *CreateCloudBenchTasksResponseBody
}

type CreateCloudBenchTasksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudBenchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudBenchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudBenchTasksResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudBenchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudBenchTasksResponse) GetBody() *CreateCloudBenchTasksResponseBody {
	return s.Body
}

func (s *CreateCloudBenchTasksResponse) SetHeaders(v map[string]*string) *CreateCloudBenchTasksResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudBenchTasksResponse) SetStatusCode(v int32) *CreateCloudBenchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudBenchTasksResponse) SetBody(v *CreateCloudBenchTasksResponseBody) *CreateCloudBenchTasksResponse {
	s.Body = v
	return s
}

func (s *CreateCloudBenchTasksResponse) Validate() error {
	return dara.Validate(s)
}
