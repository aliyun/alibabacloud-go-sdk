// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationStandardAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImageTranslationStandardAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImageTranslationStandardAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ImageTranslationStandardAsyncResponseBody) *ImageTranslationStandardAsyncResponse
	GetBody() *ImageTranslationStandardAsyncResponseBody
}

type ImageTranslationStandardAsyncResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageTranslationStandardAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImageTranslationStandardAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImageTranslationStandardAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImageTranslationStandardAsyncResponse) GetBody() *ImageTranslationStandardAsyncResponseBody {
	return s.Body
}

func (s *ImageTranslationStandardAsyncResponse) SetHeaders(v map[string]*string) *ImageTranslationStandardAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImageTranslationStandardAsyncResponse) SetStatusCode(v int32) *ImageTranslationStandardAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageTranslationStandardAsyncResponse) SetBody(v *ImageTranslationStandardAsyncResponseBody) *ImageTranslationStandardAsyncResponse {
	s.Body = v
	return s
}

func (s *ImageTranslationStandardAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
