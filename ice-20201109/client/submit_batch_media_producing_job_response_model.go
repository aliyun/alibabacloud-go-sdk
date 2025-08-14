// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchMediaProducingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitBatchMediaProducingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitBatchMediaProducingJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitBatchMediaProducingJobResponseBody) *SubmitBatchMediaProducingJobResponse
	GetBody() *SubmitBatchMediaProducingJobResponseBody
}

type SubmitBatchMediaProducingJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBatchMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBatchMediaProducingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitBatchMediaProducingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitBatchMediaProducingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitBatchMediaProducingJobResponse) GetBody() *SubmitBatchMediaProducingJobResponseBody {
	return s.Body
}

func (s *SubmitBatchMediaProducingJobResponse) SetHeaders(v map[string]*string) *SubmitBatchMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitBatchMediaProducingJobResponse) SetStatusCode(v int32) *SubmitBatchMediaProducingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBatchMediaProducingJobResponse) SetBody(v *SubmitBatchMediaProducingJobResponseBody) *SubmitBatchMediaProducingJobResponse {
	s.Body = v
	return s
}

func (s *SubmitBatchMediaProducingJobResponse) Validate() error {
	return dara.Validate(s)
}
