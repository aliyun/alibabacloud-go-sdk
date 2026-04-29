// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawCronJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolarClawCronJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolarClawCronJobResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolarClawCronJobResponseBody) *CreatePolarClawCronJobResponse
	GetBody() *CreatePolarClawCronJobResponseBody
}

type CreatePolarClawCronJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolarClawCronJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolarClawCronJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawCronJobResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarClawCronJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolarClawCronJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolarClawCronJobResponse) GetBody() *CreatePolarClawCronJobResponseBody {
	return s.Body
}

func (s *CreatePolarClawCronJobResponse) SetHeaders(v map[string]*string) *CreatePolarClawCronJobResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarClawCronJobResponse) SetStatusCode(v int32) *CreatePolarClawCronJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolarClawCronJobResponse) SetBody(v *CreatePolarClawCronJobResponseBody) *CreatePolarClawCronJobResponse {
	s.Body = v
	return s
}

func (s *CreatePolarClawCronJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
