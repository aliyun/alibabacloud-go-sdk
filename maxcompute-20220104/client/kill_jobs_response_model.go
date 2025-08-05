// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillJobsResponse
	GetStatusCode() *int32
	SetBody(v *KillJobsResponseBody) *KillJobsResponse
	GetBody() *KillJobsResponseBody
}

type KillJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s KillJobsResponse) GoString() string {
	return s.String()
}

func (s *KillJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillJobsResponse) GetBody() *KillJobsResponseBody {
	return s.Body
}

func (s *KillJobsResponse) SetHeaders(v map[string]*string) *KillJobsResponse {
	s.Headers = v
	return s
}

func (s *KillJobsResponse) SetStatusCode(v int32) *KillJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *KillJobsResponse) SetBody(v *KillJobsResponseBody) *KillJobsResponse {
	s.Body = v
	return s
}

func (s *KillJobsResponse) Validate() error {
	return dara.Validate(s)
}
