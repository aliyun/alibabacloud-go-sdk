// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignJobsResponse
	GetStatusCode() *int32
	SetBody(v *AssignJobsResponseBody) *AssignJobsResponse
	GetBody() *AssignJobsResponseBody
}

type AssignJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignJobsResponse) GoString() string {
	return s.String()
}

func (s *AssignJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignJobsResponse) GetBody() *AssignJobsResponseBody {
	return s.Body
}

func (s *AssignJobsResponse) SetHeaders(v map[string]*string) *AssignJobsResponse {
	s.Headers = v
	return s
}

func (s *AssignJobsResponse) SetStatusCode(v int32) *AssignJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignJobsResponse) SetBody(v *AssignJobsResponseBody) *AssignJobsResponse {
	s.Body = v
	return s
}

func (s *AssignJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
