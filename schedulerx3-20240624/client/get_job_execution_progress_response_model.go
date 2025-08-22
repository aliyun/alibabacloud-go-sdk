// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobExecutionProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobExecutionProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetJobExecutionProgressResponseBody) *GetJobExecutionProgressResponse
	GetBody() *GetJobExecutionProgressResponseBody
}

type GetJobExecutionProgressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobExecutionProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobExecutionProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponse) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobExecutionProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobExecutionProgressResponse) GetBody() *GetJobExecutionProgressResponseBody {
	return s.Body
}

func (s *GetJobExecutionProgressResponse) SetHeaders(v map[string]*string) *GetJobExecutionProgressResponse {
	s.Headers = v
	return s
}

func (s *GetJobExecutionProgressResponse) SetStatusCode(v int32) *GetJobExecutionProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobExecutionProgressResponse) SetBody(v *GetJobExecutionProgressResponseBody) *GetJobExecutionProgressResponse {
	s.Body = v
	return s
}

func (s *GetJobExecutionProgressResponse) Validate() error {
	return dara.Validate(s)
}
