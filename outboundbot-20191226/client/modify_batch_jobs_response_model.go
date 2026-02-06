// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBatchJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBatchJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBatchJobsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBatchJobsResponseBody) *ModifyBatchJobsResponse
	GetBody() *ModifyBatchJobsResponseBody
}

type ModifyBatchJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBatchJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBatchJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBatchJobsResponse) GoString() string {
	return s.String()
}

func (s *ModifyBatchJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBatchJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBatchJobsResponse) GetBody() *ModifyBatchJobsResponseBody {
	return s.Body
}

func (s *ModifyBatchJobsResponse) SetHeaders(v map[string]*string) *ModifyBatchJobsResponse {
	s.Headers = v
	return s
}

func (s *ModifyBatchJobsResponse) SetStatusCode(v int32) *ModifyBatchJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBatchJobsResponse) SetBody(v *ModifyBatchJobsResponseBody) *ModifyBatchJobsResponse {
	s.Body = v
	return s
}

func (s *ModifyBatchJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
