// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignJobsAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignJobsAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignJobsAsyncResponse
	GetStatusCode() *int32
	SetBody(v *AssignJobsAsyncResponseBody) *AssignJobsAsyncResponse
	GetBody() *AssignJobsAsyncResponseBody
}

type AssignJobsAsyncResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignJobsAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignJobsAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignJobsAsyncResponse) GoString() string {
	return s.String()
}

func (s *AssignJobsAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignJobsAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignJobsAsyncResponse) GetBody() *AssignJobsAsyncResponseBody {
	return s.Body
}

func (s *AssignJobsAsyncResponse) SetHeaders(v map[string]*string) *AssignJobsAsyncResponse {
	s.Headers = v
	return s
}

func (s *AssignJobsAsyncResponse) SetStatusCode(v int32) *AssignJobsAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignJobsAsyncResponse) SetBody(v *AssignJobsAsyncResponseBody) *AssignJobsAsyncResponse {
	s.Body = v
	return s
}

func (s *AssignJobsAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
