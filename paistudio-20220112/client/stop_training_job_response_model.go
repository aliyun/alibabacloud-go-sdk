// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *StopTrainingJobResponseBody) *StopTrainingJobResponse
	GetBody() *StopTrainingJobResponseBody
}

type StopTrainingJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *StopTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTrainingJobResponse) GetBody() *StopTrainingJobResponseBody {
	return s.Body
}

func (s *StopTrainingJobResponse) SetHeaders(v map[string]*string) *StopTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *StopTrainingJobResponse) SetStatusCode(v int32) *StopTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTrainingJobResponse) SetBody(v *StopTrainingJobResponseBody) *StopTrainingJobResponse {
	s.Body = v
	return s
}

func (s *StopTrainingJobResponse) Validate() error {
	return dara.Validate(s)
}
