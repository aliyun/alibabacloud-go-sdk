// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecallManagementJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecallManagementJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecallManagementJobsResponseBody) *ListRecallManagementJobsResponse
	GetBody() *ListRecallManagementJobsResponseBody
}

type ListRecallManagementJobsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecallManagementJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecallManagementJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementJobsResponse) GoString() string {
	return s.String()
}

func (s *ListRecallManagementJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecallManagementJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecallManagementJobsResponse) GetBody() *ListRecallManagementJobsResponseBody {
	return s.Body
}

func (s *ListRecallManagementJobsResponse) SetHeaders(v map[string]*string) *ListRecallManagementJobsResponse {
	s.Headers = v
	return s
}

func (s *ListRecallManagementJobsResponse) SetStatusCode(v int32) *ListRecallManagementJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecallManagementJobsResponse) SetBody(v *ListRecallManagementJobsResponseBody) *ListRecallManagementJobsResponse {
	s.Body = v
	return s
}

func (s *ListRecallManagementJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
