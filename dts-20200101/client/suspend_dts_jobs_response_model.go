// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendDtsJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendDtsJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendDtsJobsResponse
	GetStatusCode() *int32
	SetBody(v *SuspendDtsJobsResponseBody) *SuspendDtsJobsResponse
	GetBody() *SuspendDtsJobsResponseBody
}

type SuspendDtsJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendDtsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendDtsJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendDtsJobsResponse) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendDtsJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendDtsJobsResponse) GetBody() *SuspendDtsJobsResponseBody {
	return s.Body
}

func (s *SuspendDtsJobsResponse) SetHeaders(v map[string]*string) *SuspendDtsJobsResponse {
	s.Headers = v
	return s
}

func (s *SuspendDtsJobsResponse) SetStatusCode(v int32) *SuspendDtsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendDtsJobsResponse) SetBody(v *SuspendDtsJobsResponseBody) *SuspendDtsJobsResponse {
	s.Body = v
	return s
}

func (s *SuspendDtsJobsResponse) Validate() error {
	return dara.Validate(s)
}
