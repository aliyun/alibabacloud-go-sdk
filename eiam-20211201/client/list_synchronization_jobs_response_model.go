// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSynchronizationJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSynchronizationJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSynchronizationJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListSynchronizationJobsResponseBody) *ListSynchronizationJobsResponse
	GetBody() *ListSynchronizationJobsResponseBody
}

type ListSynchronizationJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSynchronizationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSynchronizationJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSynchronizationJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSynchronizationJobsResponse) GetBody() *ListSynchronizationJobsResponseBody {
	return s.Body
}

func (s *ListSynchronizationJobsResponse) SetHeaders(v map[string]*string) *ListSynchronizationJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSynchronizationJobsResponse) SetStatusCode(v int32) *ListSynchronizationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSynchronizationJobsResponse) SetBody(v *ListSynchronizationJobsResponseBody) *ListSynchronizationJobsResponse {
	s.Body = v
	return s
}

func (s *ListSynchronizationJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
