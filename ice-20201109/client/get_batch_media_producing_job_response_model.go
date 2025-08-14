// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchMediaProducingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchMediaProducingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchMediaProducingJobResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchMediaProducingJobResponseBody) *GetBatchMediaProducingJobResponse
	GetBody() *GetBatchMediaProducingJobResponseBody
}

type GetBatchMediaProducingJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchMediaProducingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *GetBatchMediaProducingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchMediaProducingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchMediaProducingJobResponse) GetBody() *GetBatchMediaProducingJobResponseBody {
	return s.Body
}

func (s *GetBatchMediaProducingJobResponse) SetHeaders(v map[string]*string) *GetBatchMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *GetBatchMediaProducingJobResponse) SetStatusCode(v int32) *GetBatchMediaProducingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchMediaProducingJobResponse) SetBody(v *GetBatchMediaProducingJobResponseBody) *GetBatchMediaProducingJobResponse {
	s.Body = v
	return s
}

func (s *GetBatchMediaProducingJobResponse) Validate() error {
	return dara.Validate(s)
}
