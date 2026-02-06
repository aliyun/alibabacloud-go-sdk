// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendJobsResponse
	GetStatusCode() *int32
	SetBody(v *SuspendJobsResponseBody) *SuspendJobsResponse
	GetBody() *SuspendJobsResponseBody
}

type SuspendJobsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendJobsResponse) GoString() string {
	return s.String()
}

func (s *SuspendJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendJobsResponse) GetBody() *SuspendJobsResponseBody {
	return s.Body
}

func (s *SuspendJobsResponse) SetHeaders(v map[string]*string) *SuspendJobsResponse {
	s.Headers = v
	return s
}

func (s *SuspendJobsResponse) SetStatusCode(v int32) *SuspendJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendJobsResponse) SetBody(v *SuspendJobsResponseBody) *SuspendJobsResponse {
	s.Body = v
	return s
}

func (s *SuspendJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
