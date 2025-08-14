// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTableUnderstandingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitTableUnderstandingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitTableUnderstandingJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitTableUnderstandingJobResponseBody) *SubmitTableUnderstandingJobResponse
	GetBody() *SubmitTableUnderstandingJobResponseBody
}

type SubmitTableUnderstandingJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitTableUnderstandingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitTableUnderstandingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitTableUnderstandingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitTableUnderstandingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitTableUnderstandingJobResponse) GetBody() *SubmitTableUnderstandingJobResponseBody {
	return s.Body
}

func (s *SubmitTableUnderstandingJobResponse) SetHeaders(v map[string]*string) *SubmitTableUnderstandingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitTableUnderstandingJobResponse) SetStatusCode(v int32) *SubmitTableUnderstandingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTableUnderstandingJobResponse) SetBody(v *SubmitTableUnderstandingJobResponseBody) *SubmitTableUnderstandingJobResponse {
	s.Body = v
	return s
}

func (s *SubmitTableUnderstandingJobResponse) Validate() error {
	return dara.Validate(s)
}
