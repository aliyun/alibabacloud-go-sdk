// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperienceDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExperienceDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExperienceDataResponse
	GetStatusCode() *int32
	SetBody(v *ListExperienceDataResponseBody) *ListExperienceDataResponse
	GetBody() *ListExperienceDataResponseBody
}

type ListExperienceDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperienceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperienceDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExperienceDataResponse) GoString() string {
	return s.String()
}

func (s *ListExperienceDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExperienceDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExperienceDataResponse) GetBody() *ListExperienceDataResponseBody {
	return s.Body
}

func (s *ListExperienceDataResponse) SetHeaders(v map[string]*string) *ListExperienceDataResponse {
	s.Headers = v
	return s
}

func (s *ListExperienceDataResponse) SetStatusCode(v int32) *ListExperienceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperienceDataResponse) SetBody(v *ListExperienceDataResponseBody) *ListExperienceDataResponse {
	s.Body = v
	return s
}

func (s *ListExperienceDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
