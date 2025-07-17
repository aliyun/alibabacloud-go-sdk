// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ResumeTaskInstancesResponseBody) *ResumeTaskInstancesResponse
	GetBody() *ResumeTaskInstancesResponseBody
}

type ResumeTaskInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ResumeTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeTaskInstancesResponse) GetBody() *ResumeTaskInstancesResponseBody {
	return s.Body
}

func (s *ResumeTaskInstancesResponse) SetHeaders(v map[string]*string) *ResumeTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ResumeTaskInstancesResponse) SetStatusCode(v int32) *ResumeTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeTaskInstancesResponse) SetBody(v *ResumeTaskInstancesResponseBody) *ResumeTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *ResumeTaskInstancesResponse) Validate() error {
	return dara.Validate(s)
}
