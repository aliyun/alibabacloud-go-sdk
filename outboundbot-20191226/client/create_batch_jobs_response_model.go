// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBatchJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBatchJobsResponse
	GetStatusCode() *int32
	SetBody(v *CreateBatchJobsResponseBody) *CreateBatchJobsResponse
	GetBody() *CreateBatchJobsResponseBody
}

type CreateBatchJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchJobsResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBatchJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBatchJobsResponse) GetBody() *CreateBatchJobsResponseBody {
	return s.Body
}

func (s *CreateBatchJobsResponse) SetHeaders(v map[string]*string) *CreateBatchJobsResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchJobsResponse) SetStatusCode(v int32) *CreateBatchJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchJobsResponse) SetBody(v *CreateBatchJobsResponseBody) *CreateBatchJobsResponse {
	s.Body = v
	return s
}

func (s *CreateBatchJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
