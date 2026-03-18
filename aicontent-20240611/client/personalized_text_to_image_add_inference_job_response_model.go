// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageAddInferenceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PersonalizedTextToImageAddInferenceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PersonalizedTextToImageAddInferenceJobResponse
	GetStatusCode() *int32
	SetBody(v *PersonalizedTextToImageAddInferenceJobResponseBody) *PersonalizedTextToImageAddInferenceJobResponse
	GetBody() *PersonalizedTextToImageAddInferenceJobResponseBody
}

type PersonalizedTextToImageAddInferenceJobResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PersonalizedTextToImageAddInferenceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageAddInferenceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageAddInferenceJobResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) GetBody() *PersonalizedTextToImageAddInferenceJobResponseBody {
	return s.Body
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageAddInferenceJobResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetStatusCode(v int32) *PersonalizedTextToImageAddInferenceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) SetBody(v *PersonalizedTextToImageAddInferenceJobResponseBody) *PersonalizedTextToImageAddInferenceJobResponse {
	s.Body = v
	return s
}

func (s *PersonalizedTextToImageAddInferenceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
