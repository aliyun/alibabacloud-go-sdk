// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageTranslationStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageTranslationStandardResponse
	GetStatusCode() *int32
	SetBody(v *ImageTranslationStandardResponseBody) *ImageTranslationStandardResponse
	GetBody() *ImageTranslationStandardResponseBody
}

type ImageTranslationStandardResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageTranslationStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageTranslationStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponse) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageTranslationStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageTranslationStandardResponse) GetBody() *ImageTranslationStandardResponseBody {
	return s.Body
}

func (s *ImageTranslationStandardResponse) SetHeaders(v map[string]*string) *ImageTranslationStandardResponse {
	s.Headers = v
	return s
}

func (s *ImageTranslationStandardResponse) SetStatusCode(v int32) *ImageTranslationStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageTranslationStandardResponse) SetBody(v *ImageTranslationStandardResponseBody) *ImageTranslationStandardResponse {
	s.Body = v
	return s
}

func (s *ImageTranslationStandardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
