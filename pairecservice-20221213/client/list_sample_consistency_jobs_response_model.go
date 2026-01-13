// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSampleConsistencyJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSampleConsistencyJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSampleConsistencyJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListSampleConsistencyJobsResponseBody) *ListSampleConsistencyJobsResponse
	GetBody() *ListSampleConsistencyJobsResponseBody
}

type ListSampleConsistencyJobsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSampleConsistencyJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSampleConsistencyJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSampleConsistencyJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSampleConsistencyJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSampleConsistencyJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSampleConsistencyJobsResponse) GetBody() *ListSampleConsistencyJobsResponseBody {
	return s.Body
}

func (s *ListSampleConsistencyJobsResponse) SetHeaders(v map[string]*string) *ListSampleConsistencyJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSampleConsistencyJobsResponse) SetStatusCode(v int32) *ListSampleConsistencyJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSampleConsistencyJobsResponse) SetBody(v *ListSampleConsistencyJobsResponseBody) *ListSampleConsistencyJobsResponse {
	s.Body = v
	return s
}

func (s *ListSampleConsistencyJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
