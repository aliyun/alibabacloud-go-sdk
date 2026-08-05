// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperienceDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExperienceDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExperienceDataResponse
	GetStatusCode() *int32
	SetBody(v *CreateExperienceDataResponseBody) *CreateExperienceDataResponse
	GetBody() *CreateExperienceDataResponseBody
}

type CreateExperienceDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperienceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperienceDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExperienceDataResponse) GoString() string {
	return s.String()
}

func (s *CreateExperienceDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExperienceDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExperienceDataResponse) GetBody() *CreateExperienceDataResponseBody {
	return s.Body
}

func (s *CreateExperienceDataResponse) SetHeaders(v map[string]*string) *CreateExperienceDataResponse {
	s.Headers = v
	return s
}

func (s *CreateExperienceDataResponse) SetStatusCode(v int32) *CreateExperienceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperienceDataResponse) SetBody(v *CreateExperienceDataResponseBody) *CreateExperienceDataResponse {
	s.Body = v
	return s
}

func (s *CreateExperienceDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
