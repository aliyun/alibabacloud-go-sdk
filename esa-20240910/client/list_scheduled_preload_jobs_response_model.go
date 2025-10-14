// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPreloadJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduledPreloadJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduledPreloadJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduledPreloadJobsResponseBody) *ListScheduledPreloadJobsResponse
	GetBody() *ListScheduledPreloadJobsResponseBody
}

type ListScheduledPreloadJobsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPreloadJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPreloadJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadJobsResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduledPreloadJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduledPreloadJobsResponse) GetBody() *ListScheduledPreloadJobsResponseBody {
	return s.Body
}

func (s *ListScheduledPreloadJobsResponse) SetHeaders(v map[string]*string) *ListScheduledPreloadJobsResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPreloadJobsResponse) SetStatusCode(v int32) *ListScheduledPreloadJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPreloadJobsResponse) SetBody(v *ListScheduledPreloadJobsResponseBody) *ListScheduledPreloadJobsResponse {
	s.Body = v
	return s
}

func (s *ListScheduledPreloadJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
