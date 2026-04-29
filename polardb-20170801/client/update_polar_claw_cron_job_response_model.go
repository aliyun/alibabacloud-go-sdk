// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawCronJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolarClawCronJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolarClawCronJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolarClawCronJobResponseBody) *UpdatePolarClawCronJobResponse
	GetBody() *UpdatePolarClawCronJobResponseBody
}

type UpdatePolarClawCronJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolarClawCronJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolarClawCronJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawCronJobResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawCronJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolarClawCronJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolarClawCronJobResponse) GetBody() *UpdatePolarClawCronJobResponseBody {
	return s.Body
}

func (s *UpdatePolarClawCronJobResponse) SetHeaders(v map[string]*string) *UpdatePolarClawCronJobResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolarClawCronJobResponse) SetStatusCode(v int32) *UpdatePolarClawCronJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolarClawCronJobResponse) SetBody(v *UpdatePolarClawCronJobResponseBody) *UpdatePolarClawCronJobResponse {
	s.Body = v
	return s
}

func (s *UpdatePolarClawCronJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
