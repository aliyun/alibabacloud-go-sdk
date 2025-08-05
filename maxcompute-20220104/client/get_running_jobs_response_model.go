// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRunningJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRunningJobsResponse
	GetStatusCode() *int32
	SetBody(v *GetRunningJobsResponseBody) *GetRunningJobsResponse
	GetBody() *GetRunningJobsResponseBody
}

type GetRunningJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRunningJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunningJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRunningJobsResponse) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRunningJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRunningJobsResponse) GetBody() *GetRunningJobsResponseBody {
	return s.Body
}

func (s *GetRunningJobsResponse) SetHeaders(v map[string]*string) *GetRunningJobsResponse {
	s.Headers = v
	return s
}

func (s *GetRunningJobsResponse) SetStatusCode(v int32) *GetRunningJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunningJobsResponse) SetBody(v *GetRunningJobsResponseBody) *GetRunningJobsResponse {
	s.Body = v
	return s
}

func (s *GetRunningJobsResponse) Validate() error {
	return dara.Validate(s)
}
