// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageTranslationProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageTranslationProResponse
	GetStatusCode() *int32
	SetBody(v *ImageTranslationProResponseBody) *ImageTranslationProResponse
	GetBody() *ImageTranslationProResponseBody
}

type ImageTranslationProResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageTranslationProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageTranslationProResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponse) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageTranslationProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageTranslationProResponse) GetBody() *ImageTranslationProResponseBody {
	return s.Body
}

func (s *ImageTranslationProResponse) SetHeaders(v map[string]*string) *ImageTranslationProResponse {
	s.Headers = v
	return s
}

func (s *ImageTranslationProResponse) SetStatusCode(v int32) *ImageTranslationProResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageTranslationProResponse) SetBody(v *ImageTranslationProResponseBody) *ImageTranslationProResponse {
	s.Body = v
	return s
}

func (s *ImageTranslationProResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
