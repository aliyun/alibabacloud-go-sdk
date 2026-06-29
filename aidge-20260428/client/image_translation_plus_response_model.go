// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationPlusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageTranslationPlusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageTranslationPlusResponse
	GetStatusCode() *int32
	SetBody(v *ImageTranslationPlusResponseBody) *ImageTranslationPlusResponse
	GetBody() *ImageTranslationPlusResponseBody
}

type ImageTranslationPlusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageTranslationPlusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageTranslationPlusResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusResponse) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageTranslationPlusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageTranslationPlusResponse) GetBody() *ImageTranslationPlusResponseBody {
	return s.Body
}

func (s *ImageTranslationPlusResponse) SetHeaders(v map[string]*string) *ImageTranslationPlusResponse {
	s.Headers = v
	return s
}

func (s *ImageTranslationPlusResponse) SetStatusCode(v int32) *ImageTranslationPlusResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageTranslationPlusResponse) SetBody(v *ImageTranslationPlusResponseBody) *ImageTranslationPlusResponse {
	s.Body = v
	return s
}

func (s *ImageTranslationPlusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
