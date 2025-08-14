// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaProducingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitMediaProducingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitMediaProducingJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitMediaProducingJobResponseBody) *SubmitMediaProducingJobResponse
	GetBody() *SubmitMediaProducingJobResponseBody
}

type SubmitMediaProducingJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitMediaProducingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitMediaProducingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitMediaProducingJobResponse) GetBody() *SubmitMediaProducingJobResponseBody {
	return s.Body
}

func (s *SubmitMediaProducingJobResponse) SetHeaders(v map[string]*string) *SubmitMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaProducingJobResponse) SetStatusCode(v int32) *SubmitMediaProducingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitMediaProducingJobResponse) SetBody(v *SubmitMediaProducingJobResponseBody) *SubmitMediaProducingJobResponse {
	s.Body = v
	return s
}

func (s *SubmitMediaProducingJobResponse) Validate() error {
	return dara.Validate(s)
}
