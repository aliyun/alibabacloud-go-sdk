// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperienceDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperienceDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperienceDataResponse
	GetStatusCode() *int32
	SetBody(v *GetExperienceDataResponseBody) *GetExperienceDataResponse
	GetBody() *GetExperienceDataResponseBody
}

type GetExperienceDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperienceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperienceDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperienceDataResponse) GoString() string {
	return s.String()
}

func (s *GetExperienceDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperienceDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperienceDataResponse) GetBody() *GetExperienceDataResponseBody {
	return s.Body
}

func (s *GetExperienceDataResponse) SetHeaders(v map[string]*string) *GetExperienceDataResponse {
	s.Headers = v
	return s
}

func (s *GetExperienceDataResponse) SetStatusCode(v int32) *GetExperienceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperienceDataResponse) SetBody(v *GetExperienceDataResponseBody) *GetExperienceDataResponse {
	s.Body = v
	return s
}

func (s *GetExperienceDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
