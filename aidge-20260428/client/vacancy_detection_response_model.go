// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVacancyDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VacancyDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VacancyDetectionResponse
	GetStatusCode() *int32
	SetBody(v *VacancyDetectionResponseBody) *VacancyDetectionResponse
	GetBody() *VacancyDetectionResponseBody
}

type VacancyDetectionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VacancyDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VacancyDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s VacancyDetectionResponse) GoString() string {
	return s.String()
}

func (s *VacancyDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VacancyDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VacancyDetectionResponse) GetBody() *VacancyDetectionResponseBody {
	return s.Body
}

func (s *VacancyDetectionResponse) SetHeaders(v map[string]*string) *VacancyDetectionResponse {
	s.Headers = v
	return s
}

func (s *VacancyDetectionResponse) SetStatusCode(v int32) *VacancyDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *VacancyDetectionResponse) SetBody(v *VacancyDetectionResponseBody) *VacancyDetectionResponse {
	s.Body = v
	return s
}

func (s *VacancyDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
