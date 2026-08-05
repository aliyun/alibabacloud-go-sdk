// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperienceDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExperienceDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExperienceDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExperienceDataResponseBody) *DeleteExperienceDataResponse
	GetBody() *DeleteExperienceDataResponseBody
}

type DeleteExperienceDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperienceDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperienceDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperienceDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperienceDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExperienceDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExperienceDataResponse) GetBody() *DeleteExperienceDataResponseBody {
	return s.Body
}

func (s *DeleteExperienceDataResponse) SetHeaders(v map[string]*string) *DeleteExperienceDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperienceDataResponse) SetStatusCode(v int32) *DeleteExperienceDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperienceDataResponse) SetBody(v *DeleteExperienceDataResponseBody) *DeleteExperienceDataResponse {
	s.Body = v
	return s
}

func (s *DeleteExperienceDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
