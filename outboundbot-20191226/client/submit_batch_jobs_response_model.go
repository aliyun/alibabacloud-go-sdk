// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitBatchJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitBatchJobsResponse
	GetStatusCode() *int32
	SetBody(v *SubmitBatchJobsResponseBody) *SubmitBatchJobsResponse
	GetBody() *SubmitBatchJobsResponseBody
}

type SubmitBatchJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBatchJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBatchJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchJobsResponse) GoString() string {
	return s.String()
}

func (s *SubmitBatchJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitBatchJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitBatchJobsResponse) GetBody() *SubmitBatchJobsResponseBody {
	return s.Body
}

func (s *SubmitBatchJobsResponse) SetHeaders(v map[string]*string) *SubmitBatchJobsResponse {
	s.Headers = v
	return s
}

func (s *SubmitBatchJobsResponse) SetStatusCode(v int32) *SubmitBatchJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBatchJobsResponse) SetBody(v *SubmitBatchJobsResponseBody) *SubmitBatchJobsResponse {
	s.Body = v
	return s
}

func (s *SubmitBatchJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
