// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPolarClawCronJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunPolarClawCronJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunPolarClawCronJobResponse
	GetStatusCode() *int32
	SetBody(v *RunPolarClawCronJobResponseBody) *RunPolarClawCronJobResponse
	GetBody() *RunPolarClawCronJobResponseBody
}

type RunPolarClawCronJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPolarClawCronJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPolarClawCronJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RunPolarClawCronJobResponse) GoString() string {
	return s.String()
}

func (s *RunPolarClawCronJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunPolarClawCronJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPolarClawCronJobResponse) GetBody() *RunPolarClawCronJobResponseBody {
	return s.Body
}

func (s *RunPolarClawCronJobResponse) SetHeaders(v map[string]*string) *RunPolarClawCronJobResponse {
	s.Headers = v
	return s
}

func (s *RunPolarClawCronJobResponse) SetStatusCode(v int32) *RunPolarClawCronJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPolarClawCronJobResponse) SetBody(v *RunPolarClawCronJobResponseBody) *RunPolarClawCronJobResponse {
	s.Body = v
	return s
}

func (s *RunPolarClawCronJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
