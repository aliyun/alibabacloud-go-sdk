// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedTextToImageQueryPreModelInferenceJobInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
	GetStatusCode() *int32
	SetBody(v *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse
	GetBody() *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody
}

type PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse struct {
	Headers    map[string]*string                                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) GoString() string {
	return s.String()
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) GetBody() *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody {
	return s.Body
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetHeaders(v map[string]*string) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.Headers = v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetStatusCode(v int32) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) SetBody(v *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponseBody) *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse {
	s.Body = v
	return s
}

func (s *PersonalizedTextToImageQueryPreModelInferenceJobInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
