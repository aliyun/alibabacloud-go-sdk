// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDtsJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDtsJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDtsJobsResponse
	GetStatusCode() *int32
	SetBody(v *StartDtsJobsResponseBody) *StartDtsJobsResponse
	GetBody() *StartDtsJobsResponseBody
}

type StartDtsJobsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDtsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDtsJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDtsJobsResponse) GoString() string {
	return s.String()
}

func (s *StartDtsJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDtsJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDtsJobsResponse) GetBody() *StartDtsJobsResponseBody {
	return s.Body
}

func (s *StartDtsJobsResponse) SetHeaders(v map[string]*string) *StartDtsJobsResponse {
	s.Headers = v
	return s
}

func (s *StartDtsJobsResponse) SetStatusCode(v int32) *StartDtsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDtsJobsResponse) SetBody(v *StartDtsJobsResponseBody) *StartDtsJobsResponse {
	s.Body = v
	return s
}

func (s *StartDtsJobsResponse) Validate() error {
	return dara.Validate(s)
}
