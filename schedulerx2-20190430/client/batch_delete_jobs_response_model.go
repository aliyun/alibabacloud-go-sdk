// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteJobsResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteJobsResponseBody) *BatchDeleteJobsResponse
	GetBody() *BatchDeleteJobsResponseBody
}

type BatchDeleteJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteJobsResponse) GetBody() *BatchDeleteJobsResponseBody {
	return s.Body
}

func (s *BatchDeleteJobsResponse) SetHeaders(v map[string]*string) *BatchDeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteJobsResponse) SetStatusCode(v int32) *BatchDeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteJobsResponse) SetBody(v *BatchDeleteJobsResponseBody) *BatchDeleteJobsResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
