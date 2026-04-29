// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawCronJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarClawCronJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarClawCronJobResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarClawCronJobResponseBody) *DeletePolarClawCronJobResponse
	GetBody() *DeletePolarClawCronJobResponseBody
}

type DeletePolarClawCronJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarClawCronJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarClawCronJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawCronJobResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarClawCronJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarClawCronJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarClawCronJobResponse) GetBody() *DeletePolarClawCronJobResponseBody {
	return s.Body
}

func (s *DeletePolarClawCronJobResponse) SetHeaders(v map[string]*string) *DeletePolarClawCronJobResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarClawCronJobResponse) SetStatusCode(v int32) *DeletePolarClawCronJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarClawCronJobResponse) SetBody(v *DeletePolarClawCronJobResponseBody) *DeletePolarClawCronJobResponse {
	s.Body = v
	return s
}

func (s *DeletePolarClawCronJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
