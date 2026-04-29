// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolarClawCronJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisablePolarClawCronJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisablePolarClawCronJobResponse
	GetStatusCode() *int32
	SetBody(v *DisablePolarClawCronJobResponseBody) *DisablePolarClawCronJobResponse
	GetBody() *DisablePolarClawCronJobResponseBody
}

type DisablePolarClawCronJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisablePolarClawCronJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisablePolarClawCronJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DisablePolarClawCronJobResponse) GoString() string {
	return s.String()
}

func (s *DisablePolarClawCronJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisablePolarClawCronJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisablePolarClawCronJobResponse) GetBody() *DisablePolarClawCronJobResponseBody {
	return s.Body
}

func (s *DisablePolarClawCronJobResponse) SetHeaders(v map[string]*string) *DisablePolarClawCronJobResponse {
	s.Headers = v
	return s
}

func (s *DisablePolarClawCronJobResponse) SetStatusCode(v int32) *DisablePolarClawCronJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DisablePolarClawCronJobResponse) SetBody(v *DisablePolarClawCronJobResponseBody) *DisablePolarClawCronJobResponse {
	s.Body = v
	return s
}

func (s *DisablePolarClawCronJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
